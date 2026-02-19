//go:build smoke

package vesselapi

import (
	"context"
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
		smokeClient, smokeClientErr = NewVesselClient(key)
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

	t.Run("Classification", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Classification(ctx, "9121998", &GetVesselIdClassificationParams{
			FilterIdType: GetVesselIdClassificationParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.Classification: %v", err)
		}
		if resp == nil || resp.Classification == nil {
			t.Fatal("expected non-nil classification response")
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

	t.Run("Inspections", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Inspections(ctx, "9121998", &GetVesselIdInspectionsParams{
			FilterIdType: GetVesselIdInspectionsParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.Inspections: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil inspections response")
		}
		if Deref(resp.InspectionCount) == 0 {
			t.Log("warning: no inspections returned")
		}
	})

	t.Run("InspectionDetail", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)

		// Dynamically discover a real detail_id from the inspections list.
		inspResp, err := client.Vessels.Inspections(ctx, "9121998", &GetVesselIdInspectionsParams{
			FilterIdType: GetVesselIdInspectionsParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.Inspections (for detail discovery): %v", err)
		}
		inspections := Deref(inspResp.Inspections)
		if len(inspections) == 0 {
			t.Skip("no inspections available to test detail endpoint")
		}
		detailID := Deref(inspections[0].DetailId)
		if detailID == "" {
			t.Skip("first inspection has no detail_id")
		}

		resp, err := client.Vessels.InspectionDetail(ctx, "9121998", detailID, &GetVesselIdInspectionsDetailIdParams{
			FilterIdType: GetVesselIdInspectionsDetailIdParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.InspectionDetail: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil inspection detail response")
		}
		if Deref(resp.DetailId) != detailID {
			t.Errorf("detail_id mismatch: got %q, want %q", Deref(resp.DetailId), detailID)
		}
	})

	t.Run("Ownership", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Vessels.Ownership(ctx, "9121998", &GetVesselIdOwnershipParams{
			FilterIdType: GetVesselIdOwnershipParamsFilterIdTypeImo,
		})
		if err != nil {
			t.Fatalf("Vessels.Ownership: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil ownership response")
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
		resp, err := client.PortEvents.ByVessel(ctx, "232003239", &GetPorteventsVesselIdParams{
			FilterIdType:    GetPorteventsVesselIdParamsFilterIdTypeMmsi,
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
		resp, err := client.PortEvents.LastByVessel(ctx, "232003239", &GetPorteventsVesselIdLastParams{
			FilterIdType: GetPorteventsVesselIdLastParamsFilterIdTypeMmsi,
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

	t.Run("Ports", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		resp, err := client.Search.Ports(ctx, &GetSearchPortsParams{
			FilterName: "Rotterdam",
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
			FilterLonLeft:   Ptr(4.0),
			FilterLonRight:  Ptr(5.0),
			FilterLatBottom: Ptr(51.0),
			FilterLatTop:    Ptr(52.0),
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
			FilterLongitude: Ptr(4.5),
			FilterLatitude:  Ptr(51.5),
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
			FilterLonLeft:   Ptr(4.0),
			FilterLonRight:  Ptr(5.0),
			FilterLatBottom: Ptr(51.0),
			FilterLatTop:    Ptr(52.0),
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
			FilterLongitude: Ptr(4.5),
			FilterLatitude:  Ptr(51.5),
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
			FilterLonLeft:   Ptr(7.0),
			FilterLonRight:  Ptr(9.0),
			FilterLatBottom: Ptr(55.0),
			FilterLatTop:    Ptr(56.0),
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
			FilterLongitude: Ptr(8.084),
			FilterLatitude:  Ptr(55.558),
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
			FilterLonLeft:   Ptr(4.0),
			FilterLonRight:  Ptr(5.0),
			FilterLatBottom: Ptr(51.0),
			FilterLatTop:    Ptr(52.0),
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
			FilterLongitude: Ptr(4.5),
			FilterLatitude:  Ptr(51.5),
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
			FilterLonLeft:   Ptr(-89.0),
			FilterLonRight:  Ptr(-88.0),
			FilterLatBottom: Ptr(28.0),
			FilterLatTop:    Ptr(29.0),
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
			FilterLongitude: Ptr(-88.5),
			FilterLatitude:  Ptr(28.2),
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
			FilterLonLeft:   Ptr(-1.0),
			FilterLonRight:  Ptr(1.0),
			FilterLatBottom: Ptr(50.0),
			FilterLatTop:    Ptr(51.0),
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
			FilterLongitude: Ptr(-0.1),
			FilterLatitude:  Ptr(50.8),
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
// Navtex (1 subtest)
// ---------------------------------------------------------------------------

func TestSmoke_Navtex(t *testing.T) {
	client := getSmokeClient(t)

	t.Run("List", func(t *testing.T) {
		t.Parallel()
		ctx := smokeCtx(t)
		now := time.Now().UTC()
		from := now.Add(-24 * time.Hour).Format(time.RFC3339)
		to := now.Format(time.RFC3339)
		resp, err := client.Navtex.List(ctx, &GetNavtexParams{
			TimeFrom:        Ptr(from),
			TimeTo:          Ptr(to),
			PaginationLimit: Ptr(5),
		})
		if err != nil {
			t.Fatalf("Navtex.List: %v", err)
		}
		if resp == nil {
			t.Fatal("expected non-nil navtex response")
		}
		// NAVTEX messages are frequent; there should be some in the last 24h.
		if len(Deref(resp.NavtexMessages)) == 0 {
			t.Log("warning: no NAVTEX messages in last 24h")
		}
	})
}

