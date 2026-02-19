package vesselapi

import "context"

// fetchFunc is a function that fetches a page of items and returns
// the items, an optional next-page token, and any error.
type fetchFunc[T any] func() (items []T, nextToken *string, err error)

// Iterator provides lazy, sequential access to paginated API results.
// Use Next to advance, Value to read the current item, and Err to check
// for errors. Collect returns all remaining items.
type Iterator[T any] struct {
	fetch   fetchFunc[T]
	items   []T
	index   int
	done    bool
	err     error
	started bool
}

func newIterator[T any](fetch fetchFunc[T]) *Iterator[T] {
	return &Iterator[T]{fetch: fetch}
}

// Next advances the iterator to the next item. It returns true if there
// is another item available, or false when iteration is complete or an
// error has occurred.
func (it *Iterator[T]) Next() bool {
	if it.err != nil {
		return false
	}

	if it.started {
		it.index++
	}
	it.started = true

	// If we have buffered items remaining, use them.
	if it.index < len(it.items) {
		return true
	}

	// No more pages to fetch.
	if it.done {
		return false
	}

	// Fetch the next page.
	items, nextToken, err := it.fetch()
	if err != nil {
		it.err = err
		return false
	}

	it.items = items
	it.index = 0

	if len(items) == 0 {
		it.done = true
		return false
	}

	// If there is no next token, this is the last page.
	if nextToken == nil || *nextToken == "" {
		it.done = true
	}

	return true
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
func (it *Iterator[T]) Collect() ([]T, error) {
	var all []T
	for it.Next() {
		all = append(all, it.Value())
	}
	if it.err != nil {
		return nil, it.err
	}
	return all, nil
}

// derefSlice safely dereferences a pointer to a slice.
func derefSlice[T any](p *[]T) []T {
	if p == nil {
		return nil
	}
	return *p
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
		params = &GetPorteventsVesselIdParams{FilterIdType: GetPorteventsVesselIdParamsFilterIdTypeImo}
	}
	p := *params
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
		params = &GetVesselIdCasualtiesParams{FilterIdType: GetVesselIdCasualtiesParamsFilterIdTypeImo}
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
		params = &GetVesselIdEmissionsParams{FilterIdType: GetVesselIdEmissionsParamsFilterIdTypeImo}
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
		params = &GetVesselsPositionsParams{FilterIdType: GetVesselsPositionsParamsFilterIdTypeImo}
	}
	p := *params
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

// --- Navtex ---

// ListAll returns an iterator over all NAVTEX messages.
func (s *NavtexService) ListAll(ctx context.Context, params *GetNavtexParams) *Iterator[Navtex] {
	if params == nil {
		params = &GetNavtexParams{}
	}
	p := *params
	return newIterator(func() ([]Navtex, *string, error) {
		resp, err := s.List(ctx, &p)
		if err != nil {
			return nil, nil, err
		}
		items := derefSlice(resp.NavtexMessages)
		p.PaginationNextToken = resp.NextToken
		return items, resp.NextToken, nil
	})
}
