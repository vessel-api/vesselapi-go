package vesselapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewVesselClient_DefaultOptions(t *testing.T) {
	vc, err := NewVesselClient("test-key")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if vc == nil {
		t.Fatal("expected non-nil VesselClient")
	}
	if vc.gen == nil {
		t.Fatal("expected non-nil generated client")
	}
	if vc.Vessels == nil {
		t.Fatal("expected non-nil Vessels service")
	}
	if vc.Ports == nil {
		t.Fatal("expected non-nil Ports service")
	}
	if vc.PortEvents == nil {
		t.Fatal("expected non-nil PortEvents service")
	}
	if vc.Emissions == nil {
		t.Fatal("expected non-nil Emissions service")
	}
	if vc.Search == nil {
		t.Fatal("expected non-nil Search service")
	}
	if vc.Location == nil {
		t.Fatal("expected non-nil Location service")
	}
	if vc.Navtex == nil {
		t.Fatal("expected non-nil Navtex service")
	}
}

func TestNewVesselClient_WithOptions(t *testing.T) {
	hc := &http.Client{Timeout: 5 * time.Second}
	vc, err := NewVesselClient("test-key",
		WithVesselBaseURL("https://custom.api.com"),
		WithVesselHTTPClient(hc),
		WithVesselUserAgent("custom-agent/1.0"),
		WithVesselRetry(5),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if vc == nil {
		t.Fatal("expected non-nil VesselClient")
	}
	// Verify base URL was applied to the generated client.
	if !strings.HasPrefix(vc.gen.Server, "https://custom.api.com") {
		t.Errorf("expected base URL to start with https://custom.api.com, got %s", vc.gen.Server)
	}
}

func TestAuthTransport_SetsHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Bearer my-secret-key" {
			t.Errorf("expected Bearer my-secret-key, got %s", auth)
		}
		ua := r.Header.Get("User-Agent")
		if ua != "test-agent/1.0" {
			t.Errorf("expected test-agent/1.0, got %s", ua)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{}`)
	}))
	defer ts.Close()

	at := &authTransport{
		base:      http.DefaultTransport,
		apiKey:    "my-secret-key",
		userAgent: "test-agent/1.0",
	}

	hc := &http.Client{Transport: at}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestNewVesselClient_PreservesHTTPClientSettings(t *testing.T) {
	// Server that blocks long enough to trigger the short timeout.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{}`)
	}))
	defer ts.Close()

	hc := &http.Client{Timeout: 50 * time.Millisecond}
	vc, err := NewVesselClient("test-key",
		WithVesselBaseURL(ts.URL),
		WithVesselHTTPClient(hc),
		WithVesselRetry(0), // no retries to avoid delay
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = vc.Vessels.Get(context.Background(), "123", nil)
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
	if !strings.Contains(err.Error(), "Client.Timeout") && !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Errorf("expected timeout-related error, got: %v", err)
	}
}

func TestRetryTransport_RetriesOn429(t *testing.T) {
	var attempts int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 3 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, `{"error":"rate limited"}`)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if atomic.LoadInt32(&attempts) != 3 {
		t.Errorf("expected 3 attempts, got %d", atomic.LoadInt32(&attempts))
	}
}

func TestRetryTransport_RetriesOn5xx(t *testing.T) {
	var attempts int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, `{"error":"server error"}`)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if atomic.LoadInt32(&attempts) != 2 {
		t.Errorf("expected 2 attempts, got %d", atomic.LoadInt32(&attempts))
	}
}

func TestRetryTransport_RespectsRetryAfter(t *testing.T) {
	var attempts int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, `{}`)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 2,
	}
	hc := &http.Client{Transport: rt}

	start := time.Now()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	elapsed := time.Since(start)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	// Should have waited at least 1 second due to Retry-After.
	if elapsed < 900*time.Millisecond {
		t.Errorf("expected at least ~1s wait due to Retry-After, got %v", elapsed)
	}
}

func TestRetryTransport_RespectsContextCancellation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "60")
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, `{}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 5,
	}
	hc := &http.Client{Transport: rt}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, ts.URL, nil)
	_, err := hc.Do(req)
	if err == nil {
		t.Fatal("expected error from context cancellation")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestRetryTransport_BackoffCapEnforced(t *testing.T) {
	// Test that the calcBackoff function caps at maxBackoff.
	resp := &http.Response{
		Header: http.Header{},
	}
	for attempt := 0; attempt < 20; attempt++ {
		d := calcBackoff(attempt, resp)
		if d > maxBackoff {
			t.Errorf("attempt %d: backoff %v exceeds max %v", attempt, d, maxBackoff)
		}
	}
}

func TestRetryTransport_RetryAfterCapAtMaxBackoff(t *testing.T) {
	resp := &http.Response{
		Header: http.Header{},
	}
	resp.Header.Set("Retry-After", "120")
	d := calcBackoff(0, resp)
	if d > maxBackoff {
		t.Errorf("expected backoff capped at %v, got %v", maxBackoff, d)
	}
}

func TestPtr(t *testing.T) {
	s := Ptr("hello")
	if *s != "hello" {
		t.Errorf("expected hello, got %s", *s)
	}

	i := Ptr(42)
	if *i != 42 {
		t.Errorf("expected 42, got %d", *i)
	}

	f := Ptr(3.14)
	if *f != 3.14 {
		t.Errorf("expected 3.14, got %f", *f)
	}
}

func TestDeref(t *testing.T) {
	s := "hello"
	if Deref(&s) != "hello" {
		t.Error("expected hello")
	}
	var sp *string
	if Deref(sp) != "" {
		t.Error("expected empty string for nil pointer")
	}

	i := 42
	if Deref(&i) != 42 {
		t.Error("expected 42")
	}
	var ip *int
	if Deref(ip) != 0 {
		t.Error("expected 0 for nil pointer")
	}
}

func TestAPIError_ImplementsError(t *testing.T) {
	var err error = &APIError{StatusCode: 400, Message: "bad request"}
	if err.Error() != "vesselapi: bad request (status 400)" {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}

func TestAPIError_Prefix(t *testing.T) {
	e := &APIError{StatusCode: 401, Message: "unauthorized"}
	if !strings.HasPrefix(e.Error(), "vesselapi:") {
		t.Errorf("expected vesselapi: prefix, got: %s", e.Error())
	}
}

func TestServiceEndToEnd_SearchVessels(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify auth header was set.
		if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") {
			t.Error("missing Bearer token")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := FindVesselsResponse{
			Vessels:   &[]Vessel{{Name: Ptr("Ever Given")}},
			NextToken: nil,
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result, err := vc.Search.Vessels(context.Background(), &GetSearchVesselsParams{
		FilterName: Ptr("Ever Given"),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	vessels := Deref(result.Vessels)
	if len(vessels) != 1 {
		t.Fatalf("expected 1 vessel, got %d", len(vessels))
	}
	if Deref(vessels[0].Name) != "Ever Given" {
		t.Errorf("expected Ever Given, got %s", Deref(vessels[0].Name))
	}
}

func TestServiceEndToEnd_ErrorResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":{"message":"missing parameter","type":"invalid_request_error"}}`)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key",
		WithVesselBaseURL(ts.URL),
		WithVesselRetry(0),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = vc.Search.Vessels(context.Background(), &GetSearchVesselsParams{})
	if err == nil {
		t.Fatal("expected error for 400 response")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", apiErr.StatusCode)
	}
}

func TestServiceEndToEnd_NilJSON200ReturnsError(t *testing.T) {
	// A 204 No Content is 2xx (passes errFromStatus) but JSON200 is nil.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key",
		WithVesselBaseURL(ts.URL),
		WithVesselRetry(0),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = vc.Vessels.Get(context.Background(), "123", nil)
	if err == nil {
		t.Fatal("expected error for nil JSON200 response")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 204 {
		t.Errorf("expected status 204, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "unexpected empty response" {
		t.Errorf("expected 'unexpected empty response', got %q", apiErr.Message)
	}
}

func TestServiceEndToEnd_GetPort(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := PortResponse{
			Port: &Port{Name: Ptr("Rotterdam"), UnloCode: Ptr("NLRTM")},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result, err := vc.Ports.Get(context.Background(), "NLRTM")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil || result.Port == nil {
		t.Fatal("expected non-nil result with port")
	}
	if Deref(result.Port.Name) != "Rotterdam" {
		t.Errorf("expected Rotterdam, got %s", Deref(result.Port.Name))
	}
}

func TestServiceEndToEnd_GetVessel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := VesselResponse{
			Vessel: &Vessel{Name: Ptr("Ever Given"), Imo: Ptr(9811000)},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	vc, err := NewVesselClient("test-key", WithVesselBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result, err := vc.Vessels.Get(context.Background(), "9811000", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil || result.Vessel == nil {
		t.Fatal("expected non-nil result with vessel")
	}
	if Deref(result.Vessel.Name) != "Ever Given" {
		t.Errorf("expected Ever Given, got %s", Deref(result.Vessel.Name))
	}
}

func TestRetryTransport_NoRetryOnSuccess(t *testing.T) {
	var attempts int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)

	if atomic.LoadInt32(&attempts) != 1 {
		t.Errorf("expected 1 attempt (no retries), got %d", atomic.LoadInt32(&attempts))
	}
}

func TestRetryTransport_NoRetryOn4xx(t *testing.T) {
	var attempts int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, ts.URL, nil)
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)

	if atomic.LoadInt32(&attempts) != 1 {
		t.Errorf("expected 1 attempt (no retries for 400), got %d", atomic.LoadInt32(&attempts))
	}
}

func TestAPIError_IsNotFound(t *testing.T) {
	e := &APIError{StatusCode: 404, Message: "not found"}
	if !e.IsNotFound() {
		t.Error("expected IsNotFound() to return true for 404")
	}
	if e.IsRateLimited() {
		t.Error("expected IsRateLimited() to return false for 404")
	}
	if e.IsAuthError() {
		t.Error("expected IsAuthError() to return false for 404")
	}
}

func TestAPIError_IsRateLimited(t *testing.T) {
	e := &APIError{StatusCode: 429, Message: "rate limited"}
	if !e.IsRateLimited() {
		t.Error("expected IsRateLimited() to return true for 429")
	}
	if e.IsNotFound() {
		t.Error("expected IsNotFound() to return false for 429")
	}
	if e.IsAuthError() {
		t.Error("expected IsAuthError() to return false for 429")
	}
}

func TestAPIError_IsAuthError(t *testing.T) {
	e := &APIError{StatusCode: 401, Message: "unauthorized"}
	if !e.IsAuthError() {
		t.Error("expected IsAuthError() to return true for 401")
	}
	if e.IsNotFound() {
		t.Error("expected IsNotFound() to return false for 401")
	}
	if e.IsRateLimited() {
		t.Error("expected IsRateLimited() to return false for 401")
	}
}

func TestAPIError_Body(t *testing.T) {
	body := []byte(`{"error":{"message":"bad request"}}`)
	e := &APIError{StatusCode: 400, Message: "bad request", Body: body}
	if string(e.Body) != string(body) {
		t.Errorf("expected body %s, got %s", body, e.Body)
	}
}

func TestNewVesselClient_EmptyKey(t *testing.T) {
	_, err := NewVesselClient("")
	if err == nil {
		t.Fatal("expected error for empty API key")
	}
	if !strings.Contains(err.Error(), "API key must not be empty") {
		t.Errorf("expected 'API key must not be empty' message, got: %v", err)
	}
}

func TestNewVesselClient_NegativeRetry(t *testing.T) {
	vc, err := NewVesselClient("test-key", WithVesselRetry(-1))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if vc == nil {
		t.Fatal("expected non-nil VesselClient")
	}
	// Negative retries should be clamped to 0, not cause a panic.
}

func TestCalcBackoff_NegativeRetryAfter(t *testing.T) {
	resp := &http.Response{
		Header: http.Header{},
	}
	resp.Header.Set("Retry-After", "-5")
	d := calcBackoff(0, resp)
	if d < 0 {
		t.Errorf("expected non-negative backoff, got %v", d)
	}
}

func TestCalcBackoff_HTTPDateRetryAfter(t *testing.T) {
	resp := &http.Response{
		Header: http.Header{},
	}
	// Set Retry-After to 2 seconds from now in HTTP-date format.
	future := time.Now().Add(2 * time.Second)
	resp.Header.Set("Retry-After", future.UTC().Format(http.TimeFormat))
	d := calcBackoff(0, resp)
	// Should be approximately 2 seconds (allow some slack for test execution).
	if d < 1*time.Second || d > 3*time.Second {
		t.Errorf("expected ~2s backoff for HTTP-date, got %v", d)
	}
}

func TestCalcBackoff_HTTPDatePastRetryAfter(t *testing.T) {
	resp := &http.Response{
		Header: http.Header{},
	}
	// Set Retry-After to a time in the past — should clamp to 0.
	past := time.Now().Add(-10 * time.Second)
	resp.Header.Set("Retry-After", past.UTC().Format(http.TimeFormat))
	d := calcBackoff(0, resp)
	if d != 0 {
		t.Errorf("expected 0 backoff for past HTTP-date, got %v", d)
	}
}

func TestRetryTransport_RetryWithPOSTBody(t *testing.T) {
	var attempts atomic.Int32
	var lastBody atomic.Value
	lastBody.Store("")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := attempts.Add(1)
		body, _ := io.ReadAll(r.Body)
		lastBody.Store(string(body))
		if n < 3 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, `{}`)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		ts.URL,
		strings.NewReader(`{"query":"test"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if attempts.Load() != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts.Load())
	}
	// Verify the POST body was correctly re-sent on the successful attempt.
	if lastBody.Load().(string) != `{"query":"test"}` {
		t.Errorf("expected POST body to be preserved on retry, got: %s", lastBody.Load().(string))
	}
}

func TestRetryTransport_RetriesOnNetworkError(t *testing.T) {
	var attempts atomic.Int32

	// Create a transport that returns network errors for the first 2 attempts.
	rt := &retryTransport{
		base: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			n := attempts.Add(1)
			if n < 3 {
				return nil, &net.OpError{
					Op:  "dial",
					Net: "tcp",
					Err: fmt.Errorf("connection refused"),
				}
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"ok":true}`)),
				Header:     http.Header{},
			}, nil
		}),
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://example.com", nil)
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if attempts.Load() != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts.Load())
	}
}

func TestRetryTransport_NoRetryOnNonTemporaryError(t *testing.T) {
	var attempts atomic.Int32

	rt := &retryTransport{
		base: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			attempts.Add(1)
			return nil, fmt.Errorf("tls: certificate is not trusted")
		}),
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://example.com", nil)
	_, err := hc.Do(req)
	if err == nil {
		t.Fatal("expected error")
	}
	if attempts.Load() != 1 {
		t.Errorf("expected 1 attempt (no retry on non-temporary error), got %d", attempts.Load())
	}
}

func TestErrFromStatus_FlatMessageJSON(t *testing.T) {
	body := []byte(`{"message":"invalid request"}`)
	err := errFromStatus(400, body)
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Message != "invalid request" {
		t.Errorf("expected 'invalid request', got %q", apiErr.Message)
	}
}

func TestErrFromStatus_FallsBackToStatusText(t *testing.T) {
	body := []byte(`<html>Server Error</html>`)
	err := errFromStatus(500, body)
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Message != "Internal Server Error" {
		t.Errorf("expected 'Internal Server Error', got %q", apiErr.Message)
	}
	if string(apiErr.Body) != "<html>Server Error</html>" {
		t.Errorf("expected raw body preserved, got %q", apiErr.Body)
	}
}

func TestRetryTransport_NoRetryOnPOST5xx(t *testing.T) {
	var attempts atomic.Int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"server error"}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		ts.URL,
		strings.NewReader(`{"data":"test"}`),
	)

	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	// POST should NOT be retried on 5xx — server may have processed the request.
	if attempts.Load() != 1 {
		t.Errorf("expected 1 attempt (no retry for POST on 5xx), got %d", attempts.Load())
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", resp.StatusCode)
	}
}

func TestRetryTransport_RetriesOnPOST429(t *testing.T) {
	var attempts atomic.Int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := attempts.Add(1)
		if n < 2 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, `{}`)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer ts.Close()

	rt := &retryTransport{
		base:       http.DefaultTransport,
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		ts.URL,
		strings.NewReader(`{"data":"test"}`),
	)

	resp, err := hc.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	// POST SHOULD be retried on 429 — rate limit means request was NOT processed.
	if attempts.Load() != 2 {
		t.Errorf("expected 2 attempts (retry POST on 429), got %d", attempts.Load())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestRetryTransport_NoRetryOnPOSTNetworkError(t *testing.T) {
	var attempts atomic.Int32

	rt := &retryTransport{
		base: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			attempts.Add(1)
			return nil, &net.OpError{
				Op:  "dial",
				Net: "tcp",
				Err: fmt.Errorf("connection refused"),
			}
		}),
		maxRetries: 3,
	}
	hc := &http.Client{Transport: rt}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://example.com", strings.NewReader(`{}`))
	_, err := hc.Do(req)
	if err == nil {
		t.Fatal("expected error")
	}
	// POST should NOT be retried on network errors.
	if attempts.Load() != 1 {
		t.Errorf("expected 1 attempt (no retry for POST on network error), got %d", attempts.Load())
	}
}

// roundTripFunc is an adapter to allow ordinary functions to be used as
// http.RoundTripper for testing.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
