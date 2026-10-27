package vesselapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const positionFullJSON = `{
	"vesselPosition": {"imo": 9321483, "mmsi": 477045900, "latitude": 51.9, "longitude": 4.1, "timestamp": "2026-10-25T10:00:00Z"},
	"vessel": {"imo": 9321483, "mmsi": 477045900, "name": "EXAMPLE", "speed_observed_max": 20.8},
	"vesselEta": {"imo": 9321483, "destination": "NLRTM", "eta": "2026-10-26T06:00:00Z"},
	"positionCheck": {
		"previous_position": {"timestamp": "2026-10-25T03:08:00Z", "latitude": 51.0, "longitude": 2.0},
		"minutes_since_previous": 412,
		"distance_nm": 96.3,
		"implied_speed_kn": 14,
		"implied_course": 87,
		"exceeds_typical_speed": false,
		"identity_match": true,
		"lookback_hours": 80
	},
	"_meta": {}
}`

const positionOldJSON = `{
	"vesselPosition": {"imo": 9321483, "mmsi": 477045900, "latitude": 51.9, "longitude": 4.1, "timestamp": "2026-10-25T10:00:00Z"}
}`

func positionServer(t *testing.T, body string, check func(*http.Request)) *VesselClient {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if check != nil {
			check(r)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, body)
	}))
	t.Cleanup(ts.Close)
	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL), WithVesselRetry(0))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return vc
}

func TestPosition_DecodesEnrichment(t *testing.T) {
	vc := positionServer(t, positionFullJSON, nil)
	resp, err := vc.Vessels.Position(context.Background(), "9321483", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Vessel == nil || resp.Vessel.Name == nil || *resp.Vessel.Name != "EXAMPLE" {
		t.Errorf("vessel not decoded: %+v", resp.Vessel)
	}
	if resp.VesselEta == nil || resp.VesselEta.Destination == nil || *resp.VesselEta.Destination != "NLRTM" {
		t.Errorf("vesselEta not decoded: %+v", resp.VesselEta)
	}
	pc := resp.PositionCheck
	if pc == nil {
		t.Fatal("positionCheck not decoded")
	}
	if pc.LookbackHours == nil || *pc.LookbackHours != 80 {
		t.Errorf("lookback_hours: got %v", pc.LookbackHours)
	}
	if pc.PreviousPosition == nil || pc.PreviousPosition.Latitude == nil || *pc.PreviousPosition.Latitude != 51.0 {
		t.Errorf("previous_position not decoded: %+v", pc.PreviousPosition)
	}
	if pc.MinutesSincePrevious == nil || *pc.MinutesSincePrevious != 412 {
		t.Errorf("minutes_since_previous: got %v", pc.MinutesSincePrevious)
	}
	if pc.DistanceNm == nil || *pc.DistanceNm != 96.3 {
		t.Errorf("distance_nm: got %v", pc.DistanceNm)
	}
	if pc.ImpliedCourse == nil || *pc.ImpliedCourse != 87 {
		t.Errorf("implied_course: got %v", pc.ImpliedCourse)
	}
	if pc.ExceedsTypicalSpeed == nil || *pc.ExceedsTypicalSpeed {
		t.Errorf("exceeds_typical_speed: got %v", pc.ExceedsTypicalSpeed)
	}
	if pc.IdentityMatch == nil || !*pc.IdentityMatch {
		t.Errorf("identity_match: got %v", pc.IdentityMatch)
	}
}

func TestPosition_DecodesOldShape(t *testing.T) {
	vc := positionServer(t, positionOldJSON, nil)
	resp, err := vc.Vessels.Position(context.Background(), "9321483", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.VesselPosition == nil || resp.VesselPosition.Latitude == nil {
		t.Fatal("vesselPosition not decoded")
	}
	if resp.Vessel != nil || resp.VesselEta != nil || resp.PositionCheck != nil {
		t.Errorf("expected nil enrichment fields, got %+v %+v %+v", resp.Vessel, resp.VesselEta, resp.PositionCheck)
	}
}

func TestPosition_SatelliteWindowParams(t *testing.T) {
	vc := positionServer(t, positionOldJSON, func(r *http.Request) {
		q := r.URL.Query()
		if got := q.Get("filter.satLookbackMinutes"); got != "480" {
			t.Errorf("filter.satLookbackMinutes: got %q", got)
		}
		if got := q.Get("filter.satMaxAgeMinutes"); got != "30" {
			t.Errorf("filter.satMaxAgeMinutes: got %q", got)
		}
	})
	_, err := vc.Vessels.Position(context.Background(), "9321483", &GetVesselIdPositionParams{
		FilterSat:                Ptr(true),
		FilterSatLookbackMinutes: Ptr(480),
		FilterSatMaxAgeMinutes:   Ptr(30),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
