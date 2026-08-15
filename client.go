package vesselapi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// VesselClient is the high-level wrapper around the generated API client.
// It provides resource-oriented service accessors for interacting with the
// Vessel Tracking API.
type VesselClient struct {
	// gen is the underlying oapi-codegen generated client.
	gen *Client

	// Vessels provides access to vessel-related endpoints.
	Vessels *VesselsService

	// Ports provides access to port lookup endpoints.
	Ports *PortsService

	// PortEvents provides access to port event endpoints.
	PortEvents *PortEventsService

	// Emissions provides access to emissions endpoints.
	Emissions *EmissionsService

	// Search provides access to search endpoints.
	Search *SearchService

	// Location provides access to location-based endpoints.
	Location *LocationService
}

// VesselClientOption configures a VesselClient.
type VesselClientOption func(*clientConfig)

type clientConfig struct {
	baseURL    string
	httpClient *http.Client
	userAgent  string
	maxRetries int
}

// WithVesselBaseURL sets the API base URL. Defaults to DefaultBaseURL.
func WithVesselBaseURL(url string) VesselClientOption {
	return func(c *clientConfig) {
		c.baseURL = url
	}
}

// WithVesselHTTPClient sets the underlying HTTP client used for transport.
//
// The client's Transport is used as the base round-tripper, with auth and retry
// layered on top. Its Timeout, Jar and CheckRedirect are adopted too. Note that
// Timeout covers the whole call including retries and the waits between them,
// not one attempt.
//
// A supplied CheckRedirect runs after the built-in check that refuses redirects
// leaving the configured origin. That check cannot be turned off, because the
// API key is attached below this layer and would otherwise travel to whatever
// host the redirect names.
func WithVesselHTTPClient(hc *http.Client) VesselClientOption {
	return func(c *clientConfig) {
		c.httpClient = hc
	}
}

// WithVesselUserAgent sets the User-Agent header value.
func WithVesselUserAgent(ua string) VesselClientOption {
	return func(c *clientConfig) {
		c.userAgent = ua
	}
}

// WithVesselRetry sets the maximum number of retries on 429 and 5xx responses.
// Defaults to 3.
func WithVesselRetry(maxRetries int) VesselClientOption {
	return func(c *clientConfig) {
		c.maxRetries = maxRetries
	}
}

// NewVesselClient creates a new high-level Vessel API client.
// The apiKey is used as a Bearer token for authentication.
func NewVesselClient(apiKey string, opts ...VesselClientOption) (*VesselClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("vesselapi: API key must not be empty")
	}

	cfg := &clientConfig{
		baseURL:    DefaultBaseURL,
		userAgent:  DefaultUserAgent,
		maxRetries: 3,
	}
	for _, o := range opts {
		o(cfg)
	}
	if cfg.maxRetries < 0 {
		cfg.maxRetries = 0
	}
	// Reject a base URL that cannot address the API here, rather than letting
	// every call fail somewhere confusing. A query string or fragment is
	// dropped: the generated client appends the endpoint path to this value, so
	// anything after it would be silently discarded along with the base path.
	parsedBase, err := url.Parse(cfg.baseURL)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: base URL %q is not a valid URL: %w", cfg.baseURL, err)
	}
	if parsedBase.Scheme != "http" && parsedBase.Scheme != "https" {
		return nil, fmt.Errorf("vesselapi: base URL %q must be http or https", cfg.baseURL)
	}
	if parsedBase.Host == "" {
		return nil, fmt.Errorf("vesselapi: base URL %q has no host", cfg.baseURL)
	}
	parsedBase.RawQuery, parsedBase.Fragment = "", ""
	cfg.baseURL = parsedBase.String()

	base := http.DefaultTransport
	if cfg.httpClient != nil && cfg.httpClient.Transport != nil {
		base = cfg.httpClient.Transport
	}

	transport := &retryTransport{
		base: &authTransport{
			base:      base,
			apiKey:    apiKey,
			userAgent: cfg.userAgent,
		},
		maxRetries: cfg.maxRetries,
	}

	hc := &http.Client{Transport: transport, CheckRedirect: refuseOffHostRedirect}
	if cfg.httpClient != nil {
		hc.Timeout = cfg.httpClient.Timeout
		hc.Jar = cfg.httpClient.Jar
		if user := cfg.httpClient.CheckRedirect; user != nil {
			// Run the key protection first, then defer to the caller. Replacing
			// it outright would leave a caller who set CheckRedirect for an
			// unrelated reason weaker than the standard library's own default,
			// because the key is attached below this layer.
			hc.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				if err := refuseOffHostRedirect(req, via); err != nil {
					return err
				}
				return user(req, via)
			}
		}
	}

	gen, err := NewClient(cfg.baseURL, WithHTTPClient(hc))
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}

	vc := &VesselClient{gen: gen}
	vc.Vessels = &VesselsService{client: gen}
	vc.Ports = &PortsService{client: gen}
	vc.PortEvents = &PortEventsService{client: gen}
	vc.Emissions = &EmissionsService{client: gen}
	vc.Search = &SearchService{client: gen}
	vc.Location = &LocationService{client: gen}

	return vc, nil
}

// refuseOffHostRedirect stops a redirect that would send the request to another
// host.
//
// The API key is attached inside the round-tripper, so it is applied afresh to
// each redirected request. Go's usual protection, which drops sensitive headers
// when a redirect crosses hosts, only covers headers set on the original
// request and therefore does not apply here. Without this check a redirect to a
// third party would arrive carrying a valid Bearer token.
func refuseOffHostRedirect(req *http.Request, via []*http.Request) error {
	if len(via) == 0 {
		return nil
	}
	if len(via) >= 10 {
		return fmt.Errorf("vesselapi: stopped after 10 redirects")
	}
	from, to := via[0].URL, req.URL
	// Compare normalised origins. A bare host comparison both over-refuses
	// (case, explicit default port) and under-refuses: it treats an https to
	// http downgrade on the same host as safe, which would put the key on the
	// wire in clear.
	if !sameOrigin(from, to) {
		return fmt.Errorf("vesselapi: refusing to follow a redirect from %s://%s to %s://%s, "+
			"which would disclose the API key", from.Scheme, from.Host, to.Scheme, to.Host)
	}
	return nil
}

// sameOrigin reports whether two URLs share a scheme and host, ignoring case
// and an explicitly written default port. A downgrade from https is never the
// same origin, even to the same host.
func sameOrigin(a, b *url.URL) bool {
	as, bs := strings.ToLower(a.Scheme), strings.ToLower(b.Scheme)
	if as != bs {
		return false
	}
	return normalisedHost(a) == normalisedHost(b)
}

func normalisedHost(u *url.URL) string {
	h := strings.ToLower(u.Hostname())
	port := u.Port()
	switch {
	case port == "",
		strings.EqualFold(u.Scheme, "https") && port == "443",
		strings.EqualFold(u.Scheme, "http") && port == "80":
		return h
	default:
		return h + ":" + port
	}
}

// authTransport adds Bearer token authentication and User-Agent headers.
type authTransport struct {
	base      http.RoundTripper
	apiKey    string
	userAgent string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Set("Authorization", "Bearer "+t.apiKey)
	r.Header.Set("User-Agent", t.userAgent)
	return t.base.RoundTrip(r)
}

// retryTransport retries requests on 429 (rate limit), 5xx responses, and
// transient network errors using exponential backoff with jitter. It respects
// the Retry-After header (both seconds and HTTP-date formats) and caps backoff
// at 30 seconds.
type retryTransport struct {
	base       http.RoundTripper
	maxRetries int
}

const maxBackoff = 30 * time.Second

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		// Clone the request per attempt to satisfy the RoundTripper contract
		// and ensure the body is fresh for retries.
		r := req.Clone(req.Context())
		if req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return nil, fmt.Errorf("vesselapi: reset request body: %w", err)
			}
			r.Body = body
		}

		resp, err := t.base.RoundTrip(r)

		// Handle network errors — retry transient ones for idempotent methods.
		if err != nil {
			if !isTemporaryErr(err) || attempt >= t.maxRetries || !isIdempotent(req.Method) {
				return nil, err
			}
			if err := sleepCtx(req.Context(), calcExpBackoff(attempt)); err != nil {
				return nil, err
			}
			continue
		}

		// Success or non-retryable status — return immediately.
		if !isRetryable(resp.StatusCode) || attempt >= t.maxRetries {
			return resp, nil
		}

		// Don't retry non-idempotent methods on 5xx — the server may have
		// processed the request. Only retry non-idempotent on 429 (rate limit)
		// where the server guarantees it was not processed.
		if resp.StatusCode != http.StatusTooManyRequests && !isIdempotent(req.Method) {
			return resp, nil
		}

		// Retryable status — compute wait from headers, then drain body and sleep.
		wait := calcBackoff(attempt, resp)
		io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<20)) //nolint:errcheck // 1 MB max drain
		resp.Body.Close()

		if err := sleepCtx(req.Context(), wait); err != nil {
			return nil, err
		}
	}
}

// sleepCtx sleeps for d, returning the context error if cancelled first.
// Uses time.NewTimer to avoid leaking timers.
func sleepCtx(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
		timer.Stop()
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func isRetryable(statusCode int) bool {
	return statusCode == http.StatusTooManyRequests || statusCode >= 500
}

// isIdempotent returns true for HTTP methods that are safe to retry.
func isIdempotent(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPut, http.MethodDelete:
		return true
	default:
		return false
	}
}

// isTemporaryErr returns true for transient network errors worth retrying.
// It returns false for context cancellation and deadline exceeded.
func isTemporaryErr(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var netErr net.Error
	return errors.As(err, &netErr)
}

func calcBackoff(attempt int, resp *http.Response) time.Duration {
	if ra := resp.Header.Get("Retry-After"); ra != "" {
		// Try seconds format.
		if seconds, err := strconv.Atoi(ra); err == nil {
			// Clamp before converting. Multiplying a large value by time.Second
			// overflows int64 and wraps to a negative duration, which would then
			// be clamped to zero and turn backoff into an immediate retry.
			switch {
			case seconds <= 0:
				return 0
			case seconds > int(maxBackoff/time.Second):
				return maxBackoff
			}
			return time.Duration(seconds) * time.Second
		}
		// Try HTTP-date format (RFC 7231 section 7.1.3).
		if t, err := http.ParseTime(ra); err == nil {
			d := time.Until(t)
			if d < 0 {
				d = 0
			}
			if d > maxBackoff {
				d = maxBackoff
			}
			return d
		}
	}
	return calcExpBackoff(attempt)
}

// calcExpBackoff returns an exponential backoff duration with jitter,
// capped at maxBackoff. Used for both retryable status codes and
// transient network errors.
func calcExpBackoff(attempt int) time.Duration {
	// Clamp the exponent first. Converting a large float to time.Duration
	// overflows int64 and wraps negative, and a negative duration makes the
	// timer fire at once, turning the backoff into a hot retry loop. The cap is
	// reached well before this bound, so clamping costs nothing.
	if attempt > 16 {
		attempt = 16
	}
	base := math.Pow(2, float64(attempt))
	jitter := rand.Float64() * base //nolint:gosec
	ms := (base + jitter) * 500
	if ms > float64(maxBackoff/time.Millisecond) {
		return maxBackoff
	}
	d := time.Duration(ms) * time.Millisecond
	if d > maxBackoff {
		d = maxBackoff
	}
	return d
}

// APIError represents an error response from the Vessel API.
type APIError struct {
	// StatusCode is the HTTP status code.
	StatusCode int

	// Message is the human-readable error message.
	Message string

	// Body is the raw response body, available for re-parsing if needed.
	Body []byte

	// Code is the machine-readable error code from the response body
	// ("error.code"), for example ErrorCodeInsufficientCredits. Empty when the
	// response carried no structured error. Branch on this rather than on
	// Message, which is free-form text and may be reworded.
	Code ErrorCode

	// Type is the error category from the response body ("error.type"), for
	// example ErrorTypePaymentRequired. Empty when the response carried no
	// structured error.
	Type ErrorType
}

func (e *APIError) Error() string {
	return fmt.Sprintf("vesselapi: %s (status %d)", e.Message, e.StatusCode)
}

// IsNotFound returns true if the error is a 404 Not Found response.
func (e *APIError) IsNotFound() bool { return e.StatusCode == 404 }

// IsRateLimited returns true if the error is a 429 Too Many Requests response.
func (e *APIError) IsRateLimited() bool { return e.StatusCode == 429 }

// IsAuthError returns true if the error is a 401 Unauthorized response.
func (e *APIError) IsAuthError() bool { return e.StatusCode == 401 }

// IsPaymentRequired returns true if the error is a 402 Payment Required
// response, returned when a satellite position is requested with no satellite
// credits left and no stored position to fall back on. Top up credits, or retry
// without the satellite fallback.
func (e *APIError) IsPaymentRequired() bool { return e.StatusCode == 402 }

// IsForbidden returns true if the error is a 403 Forbidden response, returned
// when the API key is suspended for sustained quota abuse (see the Retry-After
// header) or the feature is not on the caller's plan. Check Code to tell the
// two apart: ErrorCodeForbidden for a suspended key,
// ErrorCodeFeatureNotAvailable for a plan restriction.
func (e *APIError) IsForbidden() bool { return e.StatusCode == 403 }

// Ptr returns a pointer to the given value. Useful for constructing
// request parameters with optional fields.
func Ptr[T any](v T) *T {
	return &v
}

// Deref safely dereferences a pointer. Returns the zero value of T if p is nil.
func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
