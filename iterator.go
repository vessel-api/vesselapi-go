package vesselapi

import (
	"context"
	"fmt"
	"time"
)

// fetchFunc is a function that fetches a page of items and returns
// the items, an optional next-page token, and any error.
type fetchFunc[T any] func() (items []T, nextToken *string, err error)

// maxConsecutiveEmptyPages bounds how many empty pages the iterator will walk
// through before giving up. A service that keeps returning an empty page with a
// fresh token would otherwise be followed indefinitely, one billed request per
// turn.
const maxConsecutiveEmptyPages = 20

// Iterator provides lazy, sequential access to paginated API results.
// Use Next to advance, Value to read the current item, and Err to check
// for errors. Collect returns all remaining items.
//
// An Iterator is not safe for concurrent use. Drive it from one goroutine, or
// guard it yourself.
type Iterator[T any] struct {
	fetch   fetchFunc[T]
	items   []T
	index   int
	done    bool
	err     error
	started bool

	// seenTokens records the tokens the service has already issued. A cursor
	// that cycles (a, b, a, b) never repeats consecutively, so comparing only
	// against the previous one would follow it without limit.
	seenTokens map[string]struct{}

	// emptyPages counts consecutive pages that arrived with no items.
	emptyPages int
}

func newIterator[T any](fetch fetchFunc[T]) *Iterator[T] {
	return &Iterator[T]{fetch: fetch}
}

// Next advances the iterator to the next item. It returns true if there
// is another item available, or false when iteration is complete or an
// error has occurred.
func (it *Iterator[T]) Next() bool {
	if it.started {
		it.index++
	}
	it.started = true

	// Serve buffered items before acting on an error. The page holding them was
	// already fetched and billed, and an early stop records its error while the
	// page that triggered it is still buffered. Checking err first would drop
	// those items, which is what Collect promises not to do.
	if it.index < len(it.items) {
		return true
	}

	// Buffer drained: report a recorded failure, or the end of the pages.
	if it.err != nil || it.done {
		return false
	}

	// Keep fetching while the service hands back a token. A page can come back
	// empty while later pages still hold items, so an empty page only ends the
	// iteration when it arrives without a token. Stopping on the first empty
	// page would drop the rest silently.
	//
	// Both escapes from this loop set err as well as done. Stopping early is a
	// short read, and a caller that is told nothing cannot tell it apart from a
	// complete one.
	for !it.done {
		items, nextToken, err := it.fetch()
		if err != nil {
			it.err = err
			return false
		}

		it.items = items
		it.index = 0

		switch {
		case nextToken == nil || *nextToken == "":
			it.done = true
		default:
			if _, seen := it.seenTokens[*nextToken]; seen {
				// The service reissued a token it has already given us.
				// Following it would walk the same pages again indefinitely.
				// The loop ends on the next condition check; the page fetched
				// alongside this token is still served below.
				it.done = true
				it.err = fmt.Errorf("vesselapi: pagination stopped, the service reissued token %q; results may be incomplete", *nextToken)
			} else {
				if it.seenTokens == nil {
					it.seenTokens = make(map[string]struct{})
				}
				it.seenTokens[*nextToken] = struct{}{}
			}
		}

		if len(items) > 0 {
			it.emptyPages = 0
			return true
		}

		// An empty page with a fresh token each time would otherwise loop
		// without limit, spending a request every turn.
		it.emptyPages++
		if !it.done && it.emptyPages >= maxConsecutiveEmptyPages {
			it.done = true
			it.err = fmt.Errorf("vesselapi: pagination stopped after %d consecutive empty pages; results may be incomplete", it.emptyPages)
			return false
		}
	}

	return false
}

// Value returns the current item. Returns the zero value of T if called
// before Next() or after iteration is exhausted.
func (it *Iterator[T]) Value() T {
	if it.index < len(it.items) {
		return it.items[it.index]
	}
	var zero T
	return zero
}

// Err returns the first error encountered during iteration.
func (it *Iterator[T]) Err() error {
	return it.err
}

// Collect consumes the iterator and returns all remaining items.
//
// On failure it returns the items gathered before the error alongside it. Those
// pages have already been fetched and paid for, and the iterator cannot be
// rewound, so discarding them would lose them for good. Check the error before
// treating the result as complete.
func (it *Iterator[T]) Collect() ([]T, error) {
	var all []T
	for it.Next() {
		all = append(all, it.Value())
	}
	return all, it.err
}

// pinTimeWindow fixes a relative time window before the first page is requested.
//
// Several endpoints resolve time.from and time.to server-side relative to the
// moment each request arrives. Left unset, every page of a walk would resolve a
// slightly different window, so the pages would not describe one coherent
// result set. The vessel location endpoints go further and document that a
// nextToken is only valid alongside the bounds it was issued with, and that
// mismatched bounds are not rejected: they silently return rows from a
// different slice. Pinning once, before the first request, keeps every page on
// the same window.
//
// lookback is the endpoint's own documented default, so pinning reproduces the
// behaviour a caller already expects rather than narrowing it.
func pinTimeWindow(from, to **string, lookback time.Duration) {
	if *to == nil {
		s := time.Now().UTC().Format(time.RFC3339)
		*to = &s
	}
	if *from == nil {
		// Derive the start from the resolved end, not from now. A caller asking
		// for a historical slice supplies only TimeTo, and measuring back from
		// now would hand the service a window running backwards.
		end, err := time.Parse(time.RFC3339, **to)
		if err != nil {
			// Caller-supplied value in some other shape: leave it alone and let
			// the service judge it rather than invent a start.
			return
		}
		s := end.Add(-lookback).UTC().Format(time.RFC3339)
		*from = &s
	}
}

// pinForwardWindow fixes a forward-looking window, used by the port inbound
// endpoint whose ETA filter defaults to now through a fixed horizon.
func pinForwardWindow(from, to **string, ahead time.Duration) {
	if *from == nil {
		now := time.Now().UTC()
		// The documented default for the start of a forward window is now, and
		// pinning has to reproduce that: a caller asking for arrivals in the next
		// six hours must not also be handed the previous seventy-two. Measuring
		// back from the caller's end only makes sense when that end is already
		// past, where a window starting now would be empty.
		if *to != nil {
			if end, err := time.Parse(time.RFC3339, **to); err == nil && end.Before(now) {
				s := end.Add(-ahead).UTC().Format(time.RFC3339)
				*from = &s
			} else if err != nil {
				return
			} else {
				s := now.Format(time.RFC3339)
				*from = &s
			}
			return
		}
		s := now.Format(time.RFC3339)
		*from = &s
	}
	if *to == nil {
		start, err := time.Parse(time.RFC3339, **from)
		if err != nil {
			return
		}
		s := start.Add(ahead).UTC().Format(time.RFC3339)
		*to = &s
	}
}

// derefSlice safely dereferences a pointer to a slice.
func derefSlice[T any](p *[]T) []T {
	if p == nil {
		return nil
	}
	return *p
}

// --- Ports ---

// InboundAll returns an iterator over all inbound vessels for a port across pages.
func (s *PortsService) InboundAll(ctx context.Context, unlocode string, params *GetPortUnlocodeInboundParams) *Iterator[VesselETA] {
	if params == nil {
		params = &GetPortUnlocodeInboundParams{}
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 24*time.Hour)
	pinForwardWindow(&p.FilterEtaFrom, &p.FilterEtaTo, 72*time.Hour)
	return newIterator(func() ([]VesselETA, *string, error) {
		resp, err := s.Inbound(ctx, unlocode, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.VesselETAs)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// --- Emissions ---

// ListAll returns an iterator over all emissions across all pages.
func (s *EmissionsService) ListAll(ctx context.Context, params *GetEmissionsParams) *Iterator[VesselEmission] {
	if params == nil {
		params = &GetEmissionsParams{}
	}
	p := *params
	return newIterator(func() ([]VesselEmission, *string, error) {
		resp, err := s.List(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Emissions)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// --- Search ---

// AllVessels returns an iterator over all vessel search results.
func (s *SearchService) AllVessels(ctx context.Context, params *GetSearchVesselsParams) *Iterator[Vessel] {
	if params == nil {
		params = &GetSearchVesselsParams{}
	}
	p := *params
	return newIterator(func() ([]Vessel, *string, error) {
		resp, err := s.Vessels(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Vessels)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllPorts returns an iterator over all port search results.
func (s *SearchService) AllPorts(ctx context.Context, params *GetSearchPortsParams) *Iterator[Port] {
	if params == nil {
		params = &GetSearchPortsParams{}
	}
	p := *params
	return newIterator(func() ([]Port, *string, error) {
		resp, err := s.Ports(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Ports)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllDGPS returns an iterator over all DGPS station search results.
func (s *SearchService) AllDGPS(ctx context.Context, params *GetSearchDgpsParams) *Iterator[DGPSStation] {
	if params == nil {
		params = &GetSearchDgpsParams{}
	}
	p := *params
	return newIterator(func() ([]DGPSStation, *string, error) {
		resp, err := s.DGPS(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.DgpsStations)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllLightAids returns an iterator over all light aid search results.
func (s *SearchService) AllLightAids(ctx context.Context, params *GetSearchLightaidsParams) *Iterator[LightAid] {
	if params == nil {
		params = &GetSearchLightaidsParams{}
	}
	p := *params
	return newIterator(func() ([]LightAid, *string, error) {
		resp, err := s.LightAids(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.LightAids)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllMODUs returns an iterator over all MODU search results.
func (s *SearchService) AllMODUs(ctx context.Context, params *GetSearchModusParams) *Iterator[MODU] {
	if params == nil {
		params = &GetSearchModusParams{}
	}
	p := *params
	return newIterator(func() ([]MODU, *string, error) {
		resp, err := s.MODUs(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Modus)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllRadioBeacons returns an iterator over all radio beacon search results.
func (s *SearchService) AllRadioBeacons(ctx context.Context, params *GetSearchRadiobeaconsParams) *Iterator[RadioBeacon] {
	if params == nil {
		params = &GetSearchRadiobeaconsParams{}
	}
	p := *params
	return newIterator(func() ([]RadioBeacon, *string, error) {
		resp, err := s.RadioBeacons(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.RadioBeacons)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// --- PortEvents ---

// ListAll returns an iterator over all port events.
func (s *PortEventsService) ListAll(ctx context.Context, params *GetPorteventsParams) *Iterator[PortEvent] {
	if params == nil {
		params = &GetPorteventsParams{}
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]PortEvent, *string, error) {
		resp, err := s.List(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.PortEvents)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllByPort returns an iterator over all port events for a specific port.
func (s *PortEventsService) AllByPort(ctx context.Context, unlocode string, params *GetPorteventsPortUnlocodeParams) *Iterator[PortEvent] {
	if params == nil {
		params = &GetPorteventsPortUnlocodeParams{}
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]PortEvent, *string, error) {
		resp, err := s.ByPort(ctx, unlocode, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.PortEvents)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllByPorts returns an iterator over all port events by port name search.
func (s *PortEventsService) AllByPorts(ctx context.Context, params *GetPorteventsPortsParams) *Iterator[PortEvent] {
	if params == nil {
		params = &GetPorteventsPortsParams{}
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]PortEvent, *string, error) {
		resp, err := s.ByPorts(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.PortEvents)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllByVessel returns an iterator over all port events for a vessel.
func (s *PortEventsService) AllByVessel(ctx context.Context, id string, params *GetPorteventsVesselIdParams) *Iterator[PortEvent] {
	if params == nil {
		params = &GetPorteventsVesselIdParams{}
	}
	if params.FilterIdType == "" {
		p := *params
		p.FilterIdType = GetPorteventsVesselIdParamsFilterIdTypeImo
		params = &p
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]PortEvent, *string, error) {
		resp, err := s.ByVessel(ctx, id, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.PortEvents)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllByVessels returns an iterator over all port events by vessel name search.
func (s *PortEventsService) AllByVessels(ctx context.Context, params *GetPorteventsVesselsParams) *Iterator[PortEvent] {
	if params == nil {
		params = &GetPorteventsVesselsParams{}
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]PortEvent, *string, error) {
		resp, err := s.ByVessels(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.PortEvents)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// --- Vessels (paginated) ---

// AllCasualties returns an iterator over all casualties for a vessel.
func (s *VesselsService) AllCasualties(ctx context.Context, id string, params *GetVesselIdCasualtiesParams) *Iterator[MarineCasualty] {
	if params == nil {
		params = &GetVesselIdCasualtiesParams{}
	}
	if params.FilterIdType == "" {
		p := *params
		p.FilterIdType = GetVesselIdCasualtiesParamsFilterIdTypeImo
		params = &p
	}
	p := *params
	return newIterator(func() ([]MarineCasualty, *string, error) {
		resp, err := s.Casualties(ctx, id, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Casualties)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllEmissions returns an iterator over all emissions for a vessel.
func (s *VesselsService) AllEmissions(ctx context.Context, id string, params *GetVesselIdEmissionsParams) *Iterator[VesselEmission] {
	if params == nil {
		params = &GetVesselIdEmissionsParams{}
	}
	if params.FilterIdType == "" {
		p := *params
		p.FilterIdType = GetVesselIdEmissionsParamsFilterIdTypeImo
		params = &p
	}
	p := *params
	return newIterator(func() ([]VesselEmission, *string, error) {
		resp, err := s.Emissions(ctx, id, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Emissions)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllPositions returns an iterator over all positions for multiple vessels.
func (s *VesselsService) AllPositions(ctx context.Context, params *GetVesselsPositionsParams) *Iterator[VesselPosition] {
	if params == nil {
		params = &GetVesselsPositionsParams{}
	}
	if params.FilterIdType == "" {
		p := *params
		p.FilterIdType = GetVesselsPositionsParamsFilterIdTypeImo
		params = &p
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]VesselPosition, *string, error) {
		resp, err := s.Positions(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.VesselPositions)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// --- Location (paginated) ---

// AllVesselsBoundingBox returns an iterator over all vessel positions in a bounding box.
func (s *LocationService) AllVesselsBoundingBox(ctx context.Context, params *GetLocationVesselsBoundingBoxParams) *Iterator[VesselPosition] {
	if params == nil {
		params = &GetLocationVesselsBoundingBoxParams{}
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]VesselPosition, *string, error) {
		resp, err := s.VesselsBoundingBox(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Vessels)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllVesselsRadius returns an iterator over all vessel positions within a radius.
func (s *LocationService) AllVesselsRadius(ctx context.Context, params *GetLocationVesselsRadiusParams) *Iterator[VesselPosition] {
	if params == nil {
		params = &GetLocationVesselsRadiusParams{}
	}
	p := *params
	pinTimeWindow(&p.TimeFrom, &p.TimeTo, 2*time.Hour)
	return newIterator(func() ([]VesselPosition, *string, error) {
		resp, err := s.VesselsRadius(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Vessels)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllPortsBoundingBox returns an iterator over all ports in a bounding box.
func (s *LocationService) AllPortsBoundingBox(ctx context.Context, params *GetLocationPortsBoundingBoxParams) *Iterator[Port] {
	if params == nil {
		params = &GetLocationPortsBoundingBoxParams{}
	}
	p := *params
	return newIterator(func() ([]Port, *string, error) {
		resp, err := s.PortsBoundingBox(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Ports)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllPortsRadius returns an iterator over all ports within a radius.
func (s *LocationService) AllPortsRadius(ctx context.Context, params *GetLocationPortsRadiusParams) *Iterator[Port] {
	if params == nil {
		params = &GetLocationPortsRadiusParams{}
	}
	p := *params
	return newIterator(func() ([]Port, *string, error) {
		resp, err := s.PortsRadius(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Ports)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllDGPSBoundingBox returns an iterator over all DGPS stations in a bounding box.
func (s *LocationService) AllDGPSBoundingBox(ctx context.Context, params *GetLocationDgpsBoundingBoxParams) *Iterator[DGPSStation] {
	if params == nil {
		params = &GetLocationDgpsBoundingBoxParams{}
	}
	p := *params
	return newIterator(func() ([]DGPSStation, *string, error) {
		resp, err := s.DGPSBoundingBox(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.DgpsStations)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllDGPSRadius returns an iterator over all DGPS stations within a radius.
func (s *LocationService) AllDGPSRadius(ctx context.Context, params *GetLocationDgpsRadiusParams) *Iterator[DGPSStation] {
	if params == nil {
		params = &GetLocationDgpsRadiusParams{}
	}
	p := *params
	return newIterator(func() ([]DGPSStation, *string, error) {
		resp, err := s.DGPSRadius(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.DgpsStations)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllLightAidsBoundingBox returns an iterator over all light aids in a bounding box.
func (s *LocationService) AllLightAidsBoundingBox(ctx context.Context, params *GetLocationLightaidsBoundingBoxParams) *Iterator[LightAid] {
	if params == nil {
		params = &GetLocationLightaidsBoundingBoxParams{}
	}
	p := *params
	return newIterator(func() ([]LightAid, *string, error) {
		resp, err := s.LightAidsBoundingBox(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.LightAids)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllLightAidsRadius returns an iterator over all light aids within a radius.
func (s *LocationService) AllLightAidsRadius(ctx context.Context, params *GetLocationLightaidsRadiusParams) *Iterator[LightAid] {
	if params == nil {
		params = &GetLocationLightaidsRadiusParams{}
	}
	p := *params
	return newIterator(func() ([]LightAid, *string, error) {
		resp, err := s.LightAidsRadius(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.LightAids)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllMODUsBoundingBox returns an iterator over all MODUs in a bounding box.
func (s *LocationService) AllMODUsBoundingBox(ctx context.Context, params *GetLocationModuBoundingBoxParams) *Iterator[MODU] {
	if params == nil {
		params = &GetLocationModuBoundingBoxParams{}
	}
	p := *params
	return newIterator(func() ([]MODU, *string, error) {
		resp, err := s.MODUsBoundingBox(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Modus)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllMODUsRadius returns an iterator over all MODUs within a radius.
func (s *LocationService) AllMODUsRadius(ctx context.Context, params *GetLocationModuRadiusParams) *Iterator[MODU] {
	if params == nil {
		params = &GetLocationModuRadiusParams{}
	}
	p := *params
	return newIterator(func() ([]MODU, *string, error) {
		resp, err := s.MODUsRadius(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.Modus)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllRadioBeaconsBoundingBox returns an iterator over all radio beacons in a bounding box.
func (s *LocationService) AllRadioBeaconsBoundingBox(ctx context.Context, params *GetLocationRadiobeaconsBoundingBoxParams) *Iterator[RadioBeacon] {
	if params == nil {
		params = &GetLocationRadiobeaconsBoundingBoxParams{}
	}
	p := *params
	return newIterator(func() ([]RadioBeacon, *string, error) {
		resp, err := s.RadioBeaconsBoundingBox(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.RadioBeacons)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}

// AllRadioBeaconsRadius returns an iterator over all radio beacons within a radius.
func (s *LocationService) AllRadioBeaconsRadius(ctx context.Context, params *GetLocationRadiobeaconsRadiusParams) *Iterator[RadioBeacon] {
	if params == nil {
		params = &GetLocationRadiobeaconsRadiusParams{}
	}
	p := *params
	return newIterator(func() ([]RadioBeacon, *string, error) {
		resp, err := s.RadioBeaconsRadius(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.RadioBeacons)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}
