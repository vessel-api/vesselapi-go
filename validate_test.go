package vesselapi

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// noRequestClient returns a client whose transport fails the test if it is ever
// used. Validation must reject the call before any request is sent.
func noRequestClient(t *testing.T) *VesselClient {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("request was sent to %s, but validation should have rejected it first", r.URL)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	t.Cleanup(ts.Close)

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return vc
}

// TestRequiredParamsRejectedBeforeSending covers every method that the spec
// marks as having a required query parameter. Each call passes a zero-valued
// params struct, which is what a caller gets by forgetting to set the fields.
func TestRequiredParamsRejectedBeforeSending(t *testing.T) {
	vc := noRequestClient(t)
	ctx := context.Background()

	cases := []struct {
		name  string
		call  func() error
		param string
	}{
		{"Location.VesselsBoundingBox", func() error {
			_, err := vc.Location.VesselsBoundingBox(ctx, &GetLocationVesselsBoundingBoxParams{})
			return err
		}, "filter.lonLeft"},
		{"Location.PortsBoundingBox", func() error {
			_, err := vc.Location.PortsBoundingBox(ctx, &GetLocationPortsBoundingBoxParams{})
			return err
		}, "filter.lonLeft"},
		{"Location.DGPSBoundingBox", func() error {
			_, err := vc.Location.DGPSBoundingBox(ctx, &GetLocationDgpsBoundingBoxParams{})
			return err
		}, "filter.lonLeft"},
		{"Location.LightAidsBoundingBox", func() error {
			_, err := vc.Location.LightAidsBoundingBox(ctx, &GetLocationLightaidsBoundingBoxParams{})
			return err
		}, "filter.lonLeft"},
		{"Location.MODUsBoundingBox", func() error {
			_, err := vc.Location.MODUsBoundingBox(ctx, &GetLocationModuBoundingBoxParams{})
			return err
		}, "filter.lonLeft"},
		{"Location.RadioBeaconsBoundingBox", func() error {
			_, err := vc.Location.RadioBeaconsBoundingBox(ctx, &GetLocationRadiobeaconsBoundingBoxParams{})
			return err
		}, "filter.lonLeft"},

		{"Location.VesselsRadius", func() error {
			_, err := vc.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{})
			return err
		}, "filter.radius"},
		{"Location.PortsRadius", func() error {
			_, err := vc.Location.PortsRadius(ctx, &GetLocationPortsRadiusParams{})
			return err
		}, "filter.radius"},
		{"Location.DGPSRadius", func() error {
			_, err := vc.Location.DGPSRadius(ctx, &GetLocationDgpsRadiusParams{})
			return err
		}, "filter.radius"},
		{"Location.LightAidsRadius", func() error {
			_, err := vc.Location.LightAidsRadius(ctx, &GetLocationLightaidsRadiusParams{})
			return err
		}, "filter.radius"},
		{"Location.MODUsRadius", func() error {
			_, err := vc.Location.MODUsRadius(ctx, &GetLocationModuRadiusParams{})
			return err
		}, "filter.radius"},
		{"Location.RadioBeaconsRadius", func() error {
			_, err := vc.Location.RadioBeaconsRadius(ctx, &GetLocationRadiobeaconsRadiusParams{})
			return err
		}, "filter.radius"},

		{"Search.DGPS", func() error {
			_, err := vc.Search.DGPS(ctx, &GetSearchDgpsParams{})
			return err
		}, "filter.name"},
		{"Search.LightAids", func() error {
			_, err := vc.Search.LightAids(ctx, &GetSearchLightaidsParams{})
			return err
		}, "filter.name"},
		{"Search.MODUs", func() error {
			_, err := vc.Search.MODUs(ctx, &GetSearchModusParams{})
			return err
		}, "filter.name"},
		{"Search.RadioBeacons", func() error {
			_, err := vc.Search.RadioBeacons(ctx, &GetSearchRadiobeaconsParams{})
			return err
		}, "filter.name"},

		{"PortEvents.ByPorts", func() error {
			_, err := vc.PortEvents.ByPorts(ctx, &GetPorteventsPortsParams{})
			return err
		}, "filter.portName"},
		{"PortEvents.ByVessels", func() error {
			_, err := vc.PortEvents.ByVessels(ctx, &GetPorteventsVesselsParams{})
			return err
		}, "filter.vesselName"},

		{"Vessels.Positions", func() error {
			_, err := vc.Vessels.Positions(ctx, &GetVesselsPositionsParams{})
			return err
		}, "filter.ids"},
	}

	// A lower bound rather than an exact count: this should fail when coverage is
	// dropped, not when a method gains validation.
	if len(cases) < 19 {
		t.Fatalf("expected at least 19 methods with required parameters, have %d", len(cases))
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call()
			if err == nil {
				t.Fatal("expected an error for zero-valued required parameters, got nil")
			}
			if !errors.Is(err, ErrInvalidParams) {
				t.Errorf("errors.Is(err, ErrInvalidParams) = false, want true (err = %v)", err)
			}
			var pe *ParamError
			if !errors.As(err, &pe) {
				t.Fatalf("expected a *ParamError, got %T", err)
			}
			if !strings.Contains(pe.Param, tc.param) {
				t.Errorf("ParamError.Param = %q, want it to name %q", pe.Param, tc.param)
			}
		})
	}
}

// TestNilParamsRejected checks that passing nil for a method with required
// parameters is an error rather than a request full of zero values.
func TestNilParamsRejected(t *testing.T) {
	vc := noRequestClient(t)
	ctx := context.Background()

	if _, err := vc.Location.VesselsBoundingBox(ctx, nil); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("VesselsBoundingBox(nil) = %v, want ErrInvalidParams", err)
	}
	if _, err := vc.Search.DGPS(ctx, nil); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("DGPS(nil) = %v, want ErrInvalidParams", err)
	}
	if _, err := vc.Vessels.Positions(ctx, nil); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("Positions(nil) = %v, want ErrInvalidParams", err)
	}
}

func TestBoundingBoxRangeChecks(t *testing.T) {
	vc := noRequestClient(t)
	ctx := context.Background()

	cases := []struct {
		name   string
		params GetLocationVesselsBoundingBoxParams
		want   string
	}{
		{"latitude above 90", GetLocationVesselsBoundingBoxParams{
			FilterLonLeft: 4, FilterLonRight: 5, FilterLatBottom: 50, FilterLatTop: 91}, "filter.latTop"},
		{"latitude below -90", GetLocationVesselsBoundingBoxParams{
			FilterLonLeft: 4, FilterLonRight: 5, FilterLatBottom: -91, FilterLatTop: 50}, "filter.latBottom"},
		{"longitude above 180", GetLocationVesselsBoundingBoxParams{
			FilterLonLeft: 181, FilterLonRight: 5, FilterLatBottom: 50, FilterLatTop: 55}, "filter.lonLeft"},
		{"inverted latitude", GetLocationVesselsBoundingBoxParams{
			FilterLonLeft: 4, FilterLonRight: 5, FilterLatBottom: 55, FilterLatTop: 50}, "filter.latBottom"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := tc.params
			_, err := vc.Location.VesselsBoundingBox(ctx, &p)
			if !errors.Is(err, ErrInvalidParams) {
				t.Fatalf("got %v, want ErrInvalidParams", err)
			}
			var pe *ParamError
			errors.As(err, &pe)
			if !strings.Contains(pe.Param, tc.want) {
				t.Errorf("ParamError.Param = %q, want %q", pe.Param, tc.want)
			}
		})
	}
}

func TestRadiusMustBePositive(t *testing.T) {
	vc := noRequestClient(t)
	ctx := context.Background()

	for _, r := range []float64{0, -1} {
		_, err := vc.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
			FilterLongitude: 4, FilterLatitude: 52, FilterRadius: r,
		})
		if !errors.Is(err, ErrInvalidParams) {
			t.Errorf("radius %g: got %v, want ErrInvalidParams", r, err)
		}
	}
}

// TestValidParamsStillReachTheWire guards against validation that is too
// strict. A bounding box crossing the antimeridian is legitimate, and so is one
// that touches the equator or the prime meridian.
func TestValidParamsStillReachTheWire(t *testing.T) {
	var got string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VesselsWithinLocationResponse{Vessels: &[]VesselPosition{}})
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Crosses the antimeridian: lonLeft is east of lonRight, which is valid.
	if _, err := vc.Location.VesselsBoundingBox(context.Background(), &GetLocationVesselsBoundingBoxParams{
		FilterLonLeft: 170, FilterLonRight: -170, FilterLatBottom: -10, FilterLatTop: 10,
	}); err != nil {
		t.Fatalf("antimeridian bounding box was rejected: %v", err)
	}
	for _, want := range []string{"filter.lonLeft=170", "filter.lonRight=-170"} {
		if !strings.Contains(got, want) {
			t.Errorf("query %q missing %q", got, want)
		}
	}

	// A box with one edge on zero is fine; only all four being zero is rejected.
	if _, err := vc.Location.VesselsBoundingBox(context.Background(), &GetLocationVesselsBoundingBoxParams{
		FilterLonLeft: 0, FilterLonRight: 5, FilterLatBottom: 0, FilterLatTop: 10,
	}); err != nil {
		t.Errorf("bounding box touching zero was rejected: %v", err)
	}
}

// TestPositionsDefaultsIDType checks that filter.idType defaults to imo whenever
// the field is empty, and that the caller's struct is not mutated in the
// process.
func TestPositionsDefaultsIDType(t *testing.T) {
	var got string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VesselPositionsResponse{VesselPositions: &[]VesselPosition{}})
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	params := &GetVesselsPositionsParams{FilterIds: "9321483"}
	if _, err := vc.Vessels.Positions(context.Background(), params); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "filter.idType=imo") {
		t.Errorf("query %q does not carry the default filter.idType=imo", got)
	}
	if params.FilterIdType != "" {
		t.Errorf("caller's params were mutated: FilterIdType = %q", params.FilterIdType)
	}
}

// TestIteratorSurfacesValidationError checks that the auto-paginating helpers
// report the same failure through Err() rather than looping on a bad request.
func TestIteratorSurfacesValidationError(t *testing.T) {
	vc := noRequestClient(t)

	it := vc.Location.AllVesselsBoundingBox(context.Background(), &GetLocationVesselsBoundingBoxParams{})
	if it.Next() {
		t.Fatal("iterator yielded a value despite invalid parameters")
	}
	if !errors.Is(it.Err(), ErrInvalidParams) {
		t.Errorf("it.Err() = %v, want ErrInvalidParams", it.Err())
	}

	it2 := vc.Search.AllDGPS(context.Background(), &GetSearchDgpsParams{})
	if it2.Next() {
		t.Fatal("iterator yielded a value despite invalid parameters")
	}
	if !errors.Is(it2.Err(), ErrInvalidParams) {
		t.Errorf("it2.Err() = %v, want ErrInvalidParams", it2.Err())
	}
}

// TestNaNIsRejected covers the case that defeats a naive range check: every
// comparison against NaN is false, so it slips through unless tested for.
func TestNaNIsRejected(t *testing.T) {
	vc := noRequestClient(t)
	ctx := context.Background()
	n := math.NaN()

	if _, err := vc.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
		FilterLongitude: n, FilterLatitude: n, FilterRadius: n,
	}); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("NaN radius: got %v, want ErrInvalidParams", err)
	}
	if _, err := vc.Location.VesselsBoundingBox(ctx, &GetLocationVesselsBoundingBoxParams{
		FilterLonLeft: n, FilterLonRight: n, FilterLatBottom: n, FilterLatTop: n,
	}); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("NaN bounding box: got %v, want ErrInvalidParams", err)
	}
}

// TestDegenerateBoundingBoxRejected covers forgetting one pair of edges, which
// leaves them at zero and yields a box that can never match anything.
func TestDegenerateBoundingBoxRejected(t *testing.T) {
	vc := noRequestClient(t)
	ctx := context.Background()

	if _, err := vc.Location.VesselsBoundingBox(ctx, &GetLocationVesselsBoundingBoxParams{
		FilterLatBottom: 51, FilterLatTop: 52, // longitudes forgotten
	}); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("zero-width box: got %v, want ErrInvalidParams", err)
	}
	if _, err := vc.Location.VesselsBoundingBox(ctx, &GetLocationVesselsBoundingBoxParams{
		FilterLonLeft: 3, FilterLonRight: 5, // latitudes forgotten
	}); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("zero-height box: got %v, want ErrInvalidParams", err)
	}
}

// TestRadiusCentreRequired covers forgetting the centre point, which leaves it
// at 0,0 and would otherwise search the Gulf of Guinea without complaint.
func TestRadiusCentreRequired(t *testing.T) {
	vc := noRequestClient(t)
	ctx := context.Background()

	if _, err := vc.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
		FilterRadius: 10000, // centre forgotten
	}); !errors.Is(err, ErrInvalidParams) {
		t.Errorf("missing centre: got %v, want ErrInvalidParams", err)
	}

	// A centre on one axis only is still a real place, so it must reach the
	// wire. That needs a server, since noRequestClient fails on any request.
	var reached bool
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reached = true
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(VesselsWithinLocationResponse{Vessels: &[]VesselPosition{}})
	}))
	defer ts.Close()
	sender, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := sender.Location.VesselsRadius(ctx, &GetLocationVesselsRadiusParams{
		FilterLatitude: 51.9, FilterRadius: 10000,
	}); err != nil {
		t.Errorf("latitude-only centre was rejected: %v", err)
	}
	if !reached {
		t.Error("a valid latitude-only centre never reached the wire")
	}
}
