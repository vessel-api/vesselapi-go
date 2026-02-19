// Package vesselapi provides a Go client for the Vessel Tracking API.
//
// Usage:
//
//	client, err := vesselapi.NewVesselClient("your-api-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	vessel, err := client.Vessels.Get(ctx, "9363728", nil)
package vesselapi

const (
	// Version is the SDK version string.
	Version = "1.0.0"

	// DefaultBaseURL is the default Vessel API base URL.
	DefaultBaseURL = "https://api.vesselapi.com/v1"

	// DefaultUserAgent is the default User-Agent header value.
	DefaultUserAgent = "vesselapi-go/" + Version
)
