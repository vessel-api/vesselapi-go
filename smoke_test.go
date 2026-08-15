//go:build smoke

package vesselapi

import (
	"context"
	"errors"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	smokeClient     *VesselClient
	smokeClientOnce sync.Once
	smokeClientErr  error
)

func getSmokeClient(t *testing.T) *VesselClient {
	t.Helper()
	smokeClientOnce.Do(func() {
		key := os.Getenv("VESSELAPI_API_KEY")
		if key == "" {
			return
		}
		var opts []VesselClientOption
		if base := os.Getenv("VESSELAPI_BASE_URL"); base != "" {
			opts = append(opts, WithVesselBaseURL(base))
		}
		smokeClient, smokeClientErr = NewVesselClient(key, opts...)
	})
	if os.Getenv("VESSELAPI_API_KEY") == "" {
		t.Skip("VESSELAPI_API_KEY not set")
	}
	if smokeClientErr != nil {
		t.Fatalf("create smoke client: %v", smokeClientErr)
	}
	return smokeClient
}

func smokeCtx(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)
	return ctx
}

// ---------------------------------------------------------------------------
// Vessels (10 subtests)
// ---------------------------------------------------------------------------

func TestSmoke_Vessels(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Get(ctx, "9321483", &GetVesselIdParams{
			FilterIdType: GetVesselIdParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.Get: %v", err)
		}
		if resp == nil || resp.Vessel == nil {
			t.Fatal("expected non-nil vessel response")
		}
		if Deref(resp.Vessel.Imo) == 0 {
			t.Error("expected non-zero IMO")
		}
	})

	t.Run("Position", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Position(ctx, "232003239", &GetVesselIdPositionParams{
			FilterIdType: GetVesselIdPositionParamsFilterIdTypeMmsi,
		})
		if err != nil {
			t.Fatalf("Vessels.Position: %v", err)
		}
		if resp == nil || resp.VesselPosition == nil {
			t.Fatal("expected non-nil position response")
		}
	})

	t.Run("Casualties", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Casualties(ctx, "9321483", &GetVesselIdCasualtiesParams{
			FilterIdType: GetVesselIdCasualtiesParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.Casualties: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil casualties response")
		}
	})

	t.Run("Emissions", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Emissions(ctx, "1045356", &GetVesselIdEmissionsParams{
			FilterIdType: GetVesselIdEmissionsParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.Emissions: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil emissions response")
		}
	})

	t.Run("ETA", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.ETA(ctx, "232003239", &GetVesselIdEtaParams{
			FilterIdType: GetVesselIdEtaParamsFilterIdTypeMmsi,
		})
		if err != nil {
			t.Fatalf("Vessels.ETA: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil ETA response")
		}
	})

	t.Run("Positions", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Positions(ctx, &GetVesselsPositionsParams{
			FilterIds:    "232003239,246497000",
			FilterIdType: GetVesselsPositionsParamsFilterIdTypeMmsi,
		})
		if err != nil {
			t.Fatalf("Vessels.Positions: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil positions response")
		}
	})
}

// ---------------------------------------------------------------------------
// Ports (1 subtest)
// ---------------------------------------------------------------------------

func TestSmoke_Ports(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Ports.Get(ctx, "NLRTM")
		if err != nil {
			t.Fatalf("Ports.Get: %v", err)
		}
		if resp == nil || resp.Port == nil {
			t.Fatal("expected non-nil port response")
		}
		if Deref(resp.Port.UnloCode) != "NLRTM" {
			t.Errorf("unexpected UNLOCODE: got %q", Deref(resp.Port.UnloCode))
		}
	})

	t.Run("Inbound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		now := time.Now().UTC()
		etaFrom := now.Format(time.RFC3339)
		etaTo := now.Add(48 * time.Hour).Format(time.RFC3339)
		resp, err := client.Ports.Inbound(ctx, "NLRTM", &GetPortUnlocodeInboundParams{
			FilterEtaFrom:   &etaFrom,
			FilterEtaTo:     &etaTo,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Ports.Inbound: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil inbound response")
		}
		// VesselETAs may be nil if no vessels heading to port right now, that's OK
		t.Logf("got %d inbound vessels", len(derefSlice(resp.VesselETAs)))
	})
}

// ---------------------------------------------------------------------------
// VesselETA destination_port and _meta fallback
// ---------------------------------------------------------------------------

func TestSmoke_VesselETA_DestinationPort(t *testing.T) {
	client := getSmokeClient(t)
	ctx := smokeCtx(t)
	resp, err := client.Vessels.ETA(ctx, "9925837", nil)
	if err != nil {
		t.Skipf("Vessels.ETA: %v", err)
	}
	if resp == nil || resp.VesselEta == nil {
		t.Skip("no ETA data for test vessel")
	}
	dp := Deref(resp.VesselEta.DestinationPort)
	if dp == "" {
		t.Fatal("expected destination_port to be non-empty")
	}
	t.Logf("destination_port=%q", dp)
}

func TestSmoke_Meta_Fallback(t *testing.T) {
	client := getSmokeClient(t)
	ctx := smokeCtx(t)
	// Use an MMSI value (477045900) but claim idType=imo to trigger fallback
	mmsiAsImo := "477045900"
	resp, err := client.Vessels.Get(ctx, mmsiAsImo, &GetVesselIdParams{
		FilterIdType: GetVesselIdParamsFilterIdTypeImo,
	})
	if err != nil {
		t.Skipf("fallback lookup failed (vessel may not exist): %v", err)
	}
	if resp == nil || resp.Vessel == nil {
		t.Skip("no vessel returned")
	}
	if resp.Meta == nil {
		t.Skip("no _meta in response — API may not have triggered fallback")
	}
	if Deref(resp.Meta.ResolvedIdType) != "mmsi" {
		t.Errorf("expected resolvedIdType=mmsi, got %q", Deref(resp.Meta.ResolvedIdType))
	}
	t.Logf("_meta fallback: requested=%s resolved=%s id=%d",
		Deref(resp.Meta.RequestedIdType), Deref(resp.Meta.ResolvedIdType), Deref(resp.Meta.ResolvedId))
}

// ---------------------------------------------------------------------------
// PortEvents (6 subtests)
// ---------------------------------------------------------------------------

func TestSmoke_PortEvents(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("List", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		now := time.Now().UTC()
		from := now.Add(-24 * time.Hour).Format(time.RFC3339)
		to := now.Format(time.RFC3339)
		resp, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			TimeFrom:        Ptr(from),
			TimeTo:          Ptr(to),
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.List: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})

	t.Run("List_FilterCountry", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			FilterCountry:   Ptr("Singapore"),
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.List (country): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})

	t.Run("List_FilterEventType", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			FilterEventType: Ptr(GetPorteventsParamsFilterEventType("arrival")),
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.List (eventType): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})

	t.Run("List_CombinedFilters", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			FilterCountry:   Ptr("Singapore"),
			FilterEventType: Ptr(GetPorteventsParamsFilterEventType("arrival")),
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.List (combined): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})

	t.Run("ByPort", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.ByPort(ctx, "NLRTM", &GetPorteventsPortUnlocodeParams{
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.ByPort: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})

	t.Run("ByPorts", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.ByPorts(ctx, &GetPorteventsPortsParams{
			FilterPortName:  "Rotterdam",
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.ByPorts: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})

	t.Run("ByVessel", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.ByVessel(ctx, "9863132", &GetPorteventsVesselIdParams{
			FilterIdType:    GetPorteventsVesselIdParamsFilterIdTypeImo,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.ByVessel: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})

	t.Run("LastByVessel", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.LastByVessel(ctx, "9863132", &GetPorteventsVesselIdLastParams{
			FilterIdType: GetPorteventsVesselIdLastParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("PortEvents.LastByVessel: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port event response")
		}
	})

	t.Run("ByVessels", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.ByVessels(ctx, &GetPorteventsVesselsParams{
			FilterVesselName: "strangford 2",
			PaginationLimit:  Ptr(5),
		})
		if err != nil {
			t.Fatalf("PortEvents.ByVessels: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil port events response")
		}
	})
}

// ---------------------------------------------------------------------------
// Emissions (1 subtest)
// ---------------------------------------------------------------------------

func TestSmoke_Emissions(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("List", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Emissions.List(ctx, &GetEmissionsParams{
			FilterPeriod:    Ptr(2024),
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Emissions.List: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil emissions response")
		}
		if len(Deref(resp.Emissions)) == 0 {
			t.Error("expected at least one emission record")
		}
	})
}

// ---------------------------------------------------------------------------
// Search (6 subtests)
// ---------------------------------------------------------------------------

func TestSmoke_Search(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("Vessels", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Vessels(ctx, &GetSearchVesselsParams{
			FilterName: Ptr("EVER GIVEN"),
		})
		if err != nil {
			t.Fatalf("Search.Vessels: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Vessels)) == 0 {
			t.Error("expected at least one vessel")
		}
	})

	t.Run("Vessels_FilterFlag", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Vessels(ctx, &GetSearchVesselsParams{
			FilterFlag:      &[]string{"PA"},
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Search.Vessels (flag): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Vessels)) == 0 {
			t.Error("expected at least one vessel with flag PA")
		}
	})

	t.Run("Vessels_FilterVesselType", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Vessels(ctx, &GetSearchVesselsParams{
			FilterVesselType: &[]string{"Container Ship"},
			PaginationLimit:  Ptr(5),
		})
		if err != nil {
			t.Fatalf("Search.Vessels (vesselType): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Vessels)) == 0 {
			t.Error("expected at least one container ship")
		}
	})

	t.Run("Vessels_CombinedFilters", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Vessels(ctx, &GetSearchVesselsParams{
			FilterFlag:       &[]string{"PA"},
			FilterVesselType: &[]string{"Container Ship"},
			PaginationLimit:  Ptr(5),
		})
		if err != nil {
			t.Fatalf("Search.Vessels (combined): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("Ports", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Ports(ctx, &GetSearchPortsParams{
			FilterName: Ptr("Rotterdam"),
		})
		if err != nil {
			t.Fatalf("Search.Ports: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Ports)) == 0 {
			t.Error("expected at least one port")
		}
	})

	t.Run("Ports_FilterCountry", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Ports(ctx, &GetSearchPortsParams{
			FilterCountry:   &[]string{"NL"},
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Search.Ports (country): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Ports)) == 0 {
			t.Error("expected at least one port in NL")
		}
	})

	t.Run("Ports_FilterType", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Ports(ctx, &GetSearchPortsParams{
			FilterType:      &[]GetSearchPortsParamsFilterType{"Port"},
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Search.Ports (type): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Ports)) == 0 {
			t.Error("expected at least one port with type Port")
		}
	})

	t.Run("Ports_CombinedFilters", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Ports(ctx, &GetSearchPortsParams{
			FilterCountry:    &[]string{"NL"},
			FilterHarborSize: &[]string{"L"},
			PaginationLimit:  Ptr(5),
		})
		if err != nil {
			t.Fatalf("Search.Ports (combined): %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("DGPS", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.DGPS(ctx, &GetSearchDgpsParams{
			FilterName: "Hammer Odde",
		})
		if err != nil {
			t.Fatalf("Search.DGPS: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.DgpsStations)) == 0 {
			t.Error("expected at least one DGPS station")
		}
	})

	t.Run("LightAids", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.LightAids(ctx, &GetSearchLightaidsParams{
			FilterName: "Creach",
		})
		if err != nil {
			t.Fatalf("Search.LightAids: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.LightAids)) == 0 {
			t.Error("expected at least one light aid")
		}
	})

	t.Run("MODUs", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.MODUs(ctx, &GetSearchModusParams{
			FilterName: "ABAN",
		})
		if err != nil {
			t.Fatalf("Search.MODUs: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Modus)) == 0 {
			t.Error("expected at least one MODU")
		}
	})

	t.Run("RadioBeacons", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.RadioBeacons(ctx, &GetSearchRadiobeaconsParams{
			FilterName: "Brighton",
		})
		if err != nil {
			t.Fatalf("Search.RadioBeacons: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.RadioBeacons)) == 0 {
			t.Error("expected at least one radio beacon")
		}
	})
}

// ---------------------------------------------------------------------------
// Location (12 subtests — bounding box + radius for each entity)
// ---------------------------------------------------------------------------

func TestSmoke_Location(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("VesselsBoundingBox", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.VesselsBoundingBox(ctx, &GetLocationVesselsBoundingBoxParams{
			FilterLonLeft:   4.0,
			FilterLonRight:  5.0,
			FilterLatBottom: 51.0,
			FilterLatTop:    52.0,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.VesselsBoundingBox: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("VesselsRadius", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
			FilterLongitude: 4.5,
			FilterLatitude:  51.5,
			FilterRadius:    100000,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.VesselsRadius: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("PortsBoundingBox", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.PortsBoundingBox(ctx, &GetLocationPortsBoundingBoxParams{
			FilterLonLeft:   4.0,
			FilterLonRight:  5.0,
			FilterLatBottom: 51.0,
			FilterLatTop:    52.0,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.PortsBoundingBox: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Ports)) == 0 {
			t.Error("expected at least one port")
		}
	})

	t.Run("PortsRadius", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.PortsRadius(ctx, &GetLocationPortsRadiusParams{
			FilterLongitude: 4.5,
			FilterLatitude:  51.5,
			FilterRadius:    100000,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.PortsRadius: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(Deref(resp.Ports)) == 0 {
			t.Error("expected at least one port")
		}
	})

	t.Run("DGPSBoundingBox", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.DGPSBoundingBox(ctx, &GetLocationDgpsBoundingBoxParams{
			FilterLonLeft:   7.0,
			FilterLonRight:  9.0,
			FilterLatBottom: 55.0,
			FilterLatTop:    56.0,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.DGPSBoundingBox: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("DGPSRadius", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.DGPSRadius(ctx, &GetLocationDgpsRadiusParams{
			FilterLongitude: 8.084,
			FilterLatitude:  55.558,
			FilterRadius:    10000,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.DGPSRadius: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("LightAidsBoundingBox", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.LightAidsBoundingBox(ctx, &GetLocationLightaidsBoundingBoxParams{
			FilterLonLeft:   4.0,
			FilterLonRight:  5.0,
			FilterLatBottom: 51.0,
			FilterLatTop:    52.0,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.LightAidsBoundingBox: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("LightAidsRadius", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.LightAidsRadius(ctx, &GetLocationLightaidsRadiusParams{
			FilterLongitude: 4.5,
			FilterLatitude:  51.5,
			FilterRadius:    100000,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.LightAidsRadius: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("MODUsBoundingBox", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.MODUsBoundingBox(ctx, &GetLocationModuBoundingBoxParams{
			FilterLonLeft:   -89.0,
			FilterLonRight:  -88.0,
			FilterLatBottom: 28.0,
			FilterLatTop:    29.0,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.MODUsBoundingBox: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("MODUsRadius", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.MODUsRadius(ctx, &GetLocationModuRadiusParams{
			FilterLongitude: -88.5,
			FilterLatitude:  28.2,
			FilterRadius:    50000,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.MODUsRadius: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("RadioBeaconsBoundingBox", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.RadioBeaconsBoundingBox(ctx, &GetLocationRadiobeaconsBoundingBoxParams{
			FilterLonLeft:   -1.0,
			FilterLonRight:  1.0,
			FilterLatBottom: 50.0,
			FilterLatTop:    51.0,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.RadioBeaconsBoundingBox: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})

	t.Run("RadioBeaconsRadius", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Location.RadioBeaconsRadius(ctx, &GetLocationRadiobeaconsRadiusParams{
			FilterLongitude: -0.1,
			FilterLatitude:  50.8,
			FilterRadius:    100000,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Location.RadioBeaconsRadius: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
	})
}

// ---------------------------------------------------------------------------
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Helper: assert an APIError with a specific status code.
// ---------------------------------------------------------------------------

func requireAPIError(t *testing.T, err error, wantStatus int) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error with status %d, got nil", wantStatus)
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != wantStatus {
		t.Errorf("expected status %d, got %d: %s", wantStatus, apiErr.StatusCode, apiErr.Message)
	}
}

// ---------------------------------------------------------------------------
// Bad-param: Vessels (non-existent IDs → 404, invalid pagination → 400)
// ---------------------------------------------------------------------------

func TestSmoke_Vessels_BadParams(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("Get_NotFound_IMO", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Vessels.Get(ctx, "0000000", &GetVesselIdParams{
			FilterIdType: GetVesselIdParamsFilterIdTypeImo,
		})
		requireAPIError(t, err, 404)
	})

	t.Run("Get_NotFound_MMSI", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Vessels.Get(ctx, "000000000", &GetVesselIdParams{
			FilterIdType: GetVesselIdParamsFilterIdTypeMmsi,
		})
		requireAPIError(t, err, 404)
	})

	t.Run("Position_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Vessels.Position(ctx, "000000000", &GetVesselIdPositionParams{
			FilterIdType: GetVesselIdPositionParamsFilterIdTypeMmsi,
		})
		requireAPIError(t, err, 404)
	})

	t.Run("ETA_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Vessels.ETA(ctx, "0000000", &GetVesselIdEtaParams{
			FilterIdType: GetVesselIdEtaParamsFilterIdTypeImo,
		})
		requireAPIError(t, err, 404)
	})

	t.Run("Casualties_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Vessels.Casualties(ctx, "0000000", &GetVesselIdCasualtiesParams{
			FilterIdType: GetVesselIdCasualtiesParamsFilterIdTypeImo,
		})
		requireAPIError(t, err, 404)
	})

	// Vessel exists but has zero casualty records → 200 with empty list
	t.Run("Casualties_ExistsButEmpty", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Casualties(ctx, "9778791", &GetVesselIdCasualtiesParams{
			FilterIdType: GetVesselIdCasualtiesParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Casualties_ExistsButEmpty: expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("Casualties_ExistsButEmpty: expected non-nil response")
		}
		if items := derefSlice(resp.Casualties); len(items) != 0 {
			t.Fatalf("Casualties_ExistsButEmpty: expected empty list, got %d items", len(items))
		}
	})

	t.Run("Emissions_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Vessels.Emissions(ctx, "0000000", &GetVesselIdEmissionsParams{
			FilterIdType: GetVesselIdEmissionsParamsFilterIdTypeImo,
		})
		requireAPIError(t, err, 404)
	})

	// Vessel exists but has zero emission records → 200 with empty list
	t.Run("Emissions_ExistsButEmpty", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Emissions(ctx, "9363728", &GetVesselIdEmissionsParams{
			FilterIdType: GetVesselIdEmissionsParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Emissions_ExistsButEmpty: expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("Emissions_ExistsButEmpty: expected non-nil response")
		}
		if items := derefSlice(resp.Emissions); len(items) != 0 {
			t.Fatalf("Emissions_ExistsButEmpty: expected empty list, got %d items", len(items))
		}
	})
}

// ---------------------------------------------------------------------------
// Bad-param: Ports (non-existent UNLOCODE → 404)
// ---------------------------------------------------------------------------

func TestSmoke_Ports_BadParams(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("Get_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Ports.Get(ctx, "ZZZZZ")
		requireAPIError(t, err, 404)
	})
}

// ---------------------------------------------------------------------------
// Bad-param: PortEvents (malformed timestamps, non-existent port, bad pagination)
// ---------------------------------------------------------------------------

func TestSmoke_PortEvents_BadParams(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("List_MalformedTimeFrom", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			TimeFrom: Ptr("not-a-date"),
		})
		requireAPIError(t, err, 400)
	})

	t.Run("List_InvertedTimeRange", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			TimeFrom: Ptr("2025-01-02T00:00:00Z"),
			TimeTo:   Ptr("2025-01-01T00:00:00Z"),
		})
		requireAPIError(t, err, 400)
	})

	t.Run("List_PaginationLimitTooHigh", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			PaginationLimit: Ptr(999),
		})
		requireAPIError(t, err, 400)
	})

	t.Run("List_PaginationLimitNegative", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.List(ctx, &GetPorteventsParams{
			PaginationLimit: Ptr(-1),
		})
		requireAPIError(t, err, 400)
	})

	t.Run("ByPort_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.ByPort(ctx, "ZZZZZ", &GetPorteventsPortUnlocodeParams{
			PaginationLimit: Ptr(5),
		})
		requireAPIError(t, err, 404)
	})

	t.Run("ByVessel_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.ByVessel(ctx, "000000000", &GetPorteventsVesselIdParams{
			FilterIdType:    GetPorteventsVesselIdParamsFilterIdTypeMmsi,
			PaginationLimit: Ptr(5),
		})
		requireAPIError(t, err, 404)
	})

	// Vessel exists but has zero port event records → 200 with empty list (MADRID MAERSK)
	t.Run("ByVessel_ExistsButEmpty", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.PortEvents.ByVessel(ctx, "219836000", &GetPorteventsVesselIdParams{
			FilterIdType:    GetPorteventsVesselIdParamsFilterIdTypeMmsi,
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("ByVessel_ExistsButEmpty: expected no error, got %v", err)
		}
		if resp == nil {
			t.Fatal("ByVessel_ExistsButEmpty: expected non-nil response")
		}
		if items := derefSlice(resp.PortEvents); len(items) != 0 {
			t.Fatalf("ByVessel_ExistsButEmpty: expected empty list, got %d items", len(items))
		}
	})

	t.Run("LastByVessel_NotFound", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.LastByVessel(ctx, "000000000", &GetPorteventsVesselIdLastParams{
			FilterIdType: GetPorteventsVesselIdLastParamsFilterIdTypeMmsi,
		})
		requireAPIError(t, err, 404)
	})

	t.Run("ByPorts_EmptyName", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.ByPorts(ctx, &GetPorteventsPortsParams{
			FilterPortName: "",
		})
		requireAPIError(t, err, 400)
	})

	t.Run("ByVessels_EmptyName", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.PortEvents.ByVessels(ctx, &GetPorteventsVesselsParams{
			FilterVesselName: "",
		})
		requireAPIError(t, err, 400)
	})
}

// ---------------------------------------------------------------------------
// Bad-param: Emissions (invalid pagination)
// ---------------------------------------------------------------------------

func TestSmoke_Emissions_BadParams(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("List_PaginationLimitTooHigh", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Emissions.List(ctx, &GetEmissionsParams{
			PaginationLimit: Ptr(999),
		})
		requireAPIError(t, err, 400)
	})
}

// ---------------------------------------------------------------------------
// Bad-param: Search (empty required params, invalid pagination)
// ---------------------------------------------------------------------------

func TestSmoke_Search_BadParams(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("Vessels_NoFilters", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Search.Vessels(ctx, &GetSearchVesselsParams{})
		requireAPIError(t, err, 400)
	})

	t.Run("Vessels_PaginationTooHigh", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Search.Vessels(ctx, &GetSearchVesselsParams{
			FilterName:      Ptr("EVER GIVEN"),
			PaginationLimit: Ptr(999),
		})
		requireAPIError(t, err, 400)
	})

	t.Run("Ports_NoFilters", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Search.Ports(ctx, &GetSearchPortsParams{})
		requireAPIError(t, err, 400)
	})

	t.Run("DGPS_EmptyName", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Search.DGPS(ctx, &GetSearchDgpsParams{
			FilterName: "",
		})
		requireAPIError(t, err, 400)
	})

	t.Run("LightAids_EmptyName", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Search.LightAids(ctx, &GetSearchLightaidsParams{
			FilterName: "",
		})
		requireAPIError(t, err, 400)
	})

	t.Run("MODUs_EmptyName", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Search.MODUs(ctx, &GetSearchModusParams{
			FilterName: "",
		})
		requireAPIError(t, err, 400)
	})

	t.Run("RadioBeacons_EmptyName", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Search.RadioBeacons(ctx, &GetSearchRadiobeaconsParams{
			FilterName: "",
		})
		requireAPIError(t, err, 400)
	})
}

// ---------------------------------------------------------------------------
// Bad-param: Location (invalid coords, over-limit radius, bad pagination)
// ---------------------------------------------------------------------------

func TestSmoke_Location_BadParams(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("VesselsRadius_LatitudeTooHigh", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
			FilterLongitude: 4.5,
			FilterLatitude:  91.0,
			FilterRadius:    10000,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("VesselsRadius_LongitudeTooHigh", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
			FilterLongitude: 181.0,
			FilterLatitude:  51.5,
			FilterRadius:    10000,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("VesselsRadius_RadiusTooLarge", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
			FilterLongitude: 4.5,
			FilterLatitude:  51.5,
			FilterRadius:    200000,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("VesselsRadius_NegativeRadius", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
			FilterLongitude: 4.5,
			FilterLatitude:  51.5,
			FilterRadius:    -1,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("VesselsBoundingBox_InvertedLat", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.VesselsBoundingBox(ctx, &GetLocationVesselsBoundingBoxParams{
			FilterLonLeft:   4.0,
			FilterLonRight:  5.0,
			FilterLatBottom: 52.0,
			FilterLatTop:    51.0, // inverted: bottom > top
		})
		requireAPIError(t, err, 400)
	})

	t.Run("VesselsBoundingBox_PaginationTooHigh", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.VesselsBoundingBox(ctx, &GetLocationVesselsBoundingBoxParams{
			FilterLonLeft:   4.0,
			FilterLonRight:  5.0,
			FilterLatBottom: 51.0,
			FilterLatTop:    52.0,
			PaginationLimit: Ptr(999),
		})
		requireAPIError(t, err, 400)
	})

	t.Run("PortsRadius_RadiusTooLarge", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.PortsRadius(ctx, &GetLocationPortsRadiusParams{
			FilterLongitude: 4.5,
			FilterLatitude:  51.5,
			FilterRadius:    200000,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("PortsBoundingBox_InvertedLon", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.PortsBoundingBox(ctx, &GetLocationPortsBoundingBoxParams{
			FilterLonLeft:   5.0,
			FilterLonRight:  4.0, // inverted: left > right
			FilterLatBottom: 51.0,
			FilterLatTop:    52.0,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("DGPSRadius_LatitudeTooLow", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.DGPSRadius(ctx, &GetLocationDgpsRadiusParams{
			FilterLongitude: 8.0,
			FilterLatitude:  -91.0,
			FilterRadius:    10000,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("LightAidsRadius_LongitudeTooLow", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.LightAidsRadius(ctx, &GetLocationLightaidsRadiusParams{
			FilterLongitude: -181.0,
			FilterLatitude:  51.5,
			FilterRadius:    10000,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("MODUsRadius_RadiusTooLarge", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.MODUsRadius(ctx, &GetLocationModuRadiusParams{
			FilterLongitude: -88.5,
			FilterLatitude:  28.2,
			FilterRadius:    200000,
		})
		requireAPIError(t, err, 400)
	})

	t.Run("RadioBeaconsRadius_RadiusTooLarge", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		_, err := client.Location.RadioBeaconsRadius(ctx, &GetLocationRadiobeaconsRadiusParams{
			FilterLongitude: -0.1,
			FilterLatitude:  50.8,
			FilterRadius:    200000,
		})
		requireAPIError(t, err, 400)
	})
}

// ---------------------------------------------------------------------------
// ---------------------------------------------------------------------------
