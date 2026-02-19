package vesselapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestIterator_MultiplePages(t *testing.T) {
	page := 0
	it := newIterator(func() ([]string, *string, error) {
		page++
		switch page {
		case 1:
			tok := "page2"
			return []string{"a", "b"}, &tok, nil
		case 2:
			tok := "page3"
			return []string{"c"}, &tok, nil
		case 3:
			return []string{"d"}, nil, nil
		default:
			return nil, nil, fmt.Errorf("unexpected page %d", page)
		}
	})

	var results []string
	for it.Next() {
		results = append(results, it.Value())
	}
	if it.Err() != nil {
		t.Fatalf("unexpected error: %v", it.Err())
	}
	expected := []string{"a", "b", "c", "d"}
	if len(results) != len(expected) {
		t.Fatalf("expected %d items, got %d", len(expected), len(results))
	}
	for i, v := range results {
		if v != expected[i] {
			t.Errorf("item %d: expected %s, got %s", i, expected[i], v)
		}
	}
}

func TestIterator_EmptyResult(t *testing.T) {
	it := newIterator(func() ([]string, *string, error) {
		return nil, nil, nil
	})

	if it.Next() {
		t.Error("expected Next() to return false for empty result")
	}
	if it.Err() != nil {
		t.Errorf("unexpected error: %v", it.Err())
	}
}

func TestIterator_ErrorOnFirstPage(t *testing.T) {
	expectedErr := fmt.Errorf("network error")
	it := newIterator(func() ([]string, *string, error) {
		return nil, nil, expectedErr
	})

	if it.Next() {
		t.Error("expected Next() to return false on error")
	}
	if it.Err() != expectedErr {
		t.Errorf("expected %v, got %v", expectedErr, it.Err())
	}
}

func TestIterator_ErrorOnSubsequentPage(t *testing.T) {
	page := 0
	expectedErr := fmt.Errorf("page 2 error")
	it := newIterator(func() ([]string, *string, error) {
		page++
		if page == 1 {
			tok := "next"
			return []string{"a"}, &tok, nil
		}
		return nil, nil, expectedErr
	})

	// First item succeeds.
	if !it.Next() {
		t.Fatal("expected Next() to return true for first item")
	}
	if it.Value() != "a" {
		t.Errorf("expected 'a', got %s", it.Value())
	}

	// Second call should fail.
	if it.Next() {
		t.Error("expected Next() to return false after error")
	}
	if it.Err() != expectedErr {
		t.Errorf("expected %v, got %v", expectedErr, it.Err())
	}
}

func TestIterator_Collect(t *testing.T) {
	page := 0
	it := newIterator(func() ([]int, *string, error) {
		page++
		switch page {
		case 1:
			tok := "next"
			return []int{1, 2, 3}, &tok, nil
		case 2:
			return []int{4, 5}, nil, nil
		default:
			return nil, nil, fmt.Errorf("unexpected page")
		}
	})

	items, err := it.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 5 {
		t.Fatalf("expected 5 items, got %d", len(items))
	}
	for i, v := range items {
		if v != i+1 {
			t.Errorf("item %d: expected %d, got %d", i, i+1, v)
		}
	}
}

func TestIterator_CollectError(t *testing.T) {
	page := 0
	it := newIterator(func() ([]int, *string, error) {
		page++
		if page == 1 {
			tok := "next"
			return []int{1}, &tok, nil
		}
		return nil, nil, fmt.Errorf("collect error")
	})

	items, err := it.Collect()
	if err == nil {
		t.Fatal("expected error from Collect")
	}
	if items != nil {
		t.Errorf("expected nil items on error, got %v", items)
	}
}

func TestIterator_DoesNotMutateOriginalParams(t *testing.T) {
	originalToken := "original"
	params := &GetSearchVesselsParams{
		FilterName:          Ptr("test"),
		PaginationNextToken: &originalToken,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := FindVesselsResponse{
			Vessels:   &[]Vessel{},
			NextToken: nil,
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	it := vc.Search.AllVessels(context.Background(), params)

	// Consume the iterator.
	for it.Next() {
		// no-op
	}

	// Original params should not be mutated.
	if params.PaginationNextToken == nil || *params.PaginationNextToken != "original" {
		t.Error("original params PaginationNextToken was mutated")
	}
}

func TestIterator_ValueBeforeNextReturnsZero(t *testing.T) {
	it := newIterator(func() ([]string, *string, error) {
		return []string{"a", "b"}, nil, nil
	})

	// Value before Next should return zero value.
	if v := it.Value(); v != "" {
		t.Errorf("expected empty string before Next(), got %q", v)
	}

	// Exhaust the iterator.
	for it.Next() {
	}

	// Value after exhaustion should return zero value.
	if v := it.Value(); v != "" {
		t.Errorf("expected empty string after exhaustion, got %q", v)
	}
}

func TestIterator_SearchVesselsIntegration(t *testing.T) {
	var page atomic.Int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := page.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		switch n {
		case 1:
			resp := FindVesselsResponse{
				Vessels:   &[]Vessel{{Name: Ptr("Vessel A")}, {Name: Ptr("Vessel B")}},
				NextToken: Ptr("token2"),
			}
			json.NewEncoder(w).Encode(resp)
		case 2:
			resp := FindVesselsResponse{
				Vessels:   &[]Vessel{{Name: Ptr("Vessel C")}},
				NextToken: nil,
			}
			json.NewEncoder(w).Encode(resp)
		default:
			t.Error("too many requests")
		}
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	it := vc.Search.AllVessels(context.Background(), &GetSearchVesselsParams{
		FilterName: Ptr("Vessel"),
	})

	vessels, err := it.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(vessels) != 3 {
		t.Fatalf("expected 3 vessels, got %d", len(vessels))
	}
	names := make([]string, len(vessels))
	for i, v := range vessels {
		names[i] = Deref(v.Name)
	}
	expected := []string{"Vessel A", "Vessel B", "Vessel C"}
	for i, n := range names {
		if n != expected[i] {
			t.Errorf("vessel %d: expected %s, got %s", i, expected[i], n)
		}
	}
}

func TestIterator_PortEventsIntegration(t *testing.T) {
	var page atomic.Int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := page.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		switch n {
		case 1:
			resp := PortEventsResponse{
				PortEvents: &[]PortEvent{{Event: Ptr("Arrival")}, {Event: Ptr("Departure")}},
				NextToken:  Ptr("next"),
			}
			json.NewEncoder(w).Encode(resp)
		case 2:
			resp := PortEventsResponse{
				PortEvents: &[]PortEvent{{Event: Ptr("Arrival")}},
				NextToken:  nil,
			}
			json.NewEncoder(w).Encode(resp)
		default:
			t.Error("too many requests")
		}
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	it := vc.PortEvents.ListAll(context.Background(), &GetPorteventsParams{})

	events, err := it.Collect()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 3 {
		t.Fatalf("expected 3 events, got %d", len(events))
	}
}

func TestIterator_ErrorResponseIntegration(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"error":{"message":"invalid api key","type":"authentication_error"}}`)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("bad-key",
		WithVesselBaseURL(ts.URL),
		WithVesselRetry(0),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	it := vc.Search.AllVessels(context.Background(), &GetSearchVesselsParams{
		FilterName: Ptr("test"),
	})

	if it.Next() {
		t.Error("expected Next() to return false on auth error")
	}
	if it.Err() == nil {
		t.Fatal("expected error")
	}
	var apiErr *APIError
	if !errors.As(it.Err(), &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", it.Err(), it.Err())
	}
	if apiErr.StatusCode != 401 {
		t.Errorf("expected 401, got %d", apiErr.StatusCode)
	}
}

func TestIterator_NilParamsDoesNotPanic(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := FindVesselsResponse{
			Vessels:   &[]Vessel{},
			NextToken: nil,
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Passing nil params must not panic.
	it := vc.Search.AllVessels(context.Background(), nil)
	for it.Next() {
		// no-op
	}
	if it.Err() != nil {
		t.Fatalf("unexpected error: %v", it.Err())
	}
}

func TestDerefSlice(t *testing.T) {
	// nil pointer
	var nilSlice *[]string
	result := derefSlice(nilSlice)
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}

	// non-nil pointer
	s := []string{"a", "b"}
	result = derefSlice(&s)
	if len(result) != 2 || result[0] != "a" {
		t.Errorf("expected [a b], got %v", result)
	}
}
