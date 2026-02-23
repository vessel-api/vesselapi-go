# vesselapi-go

[![CI](https://github.com/vessel-api/vesselapi-go/actions/workflows/ci.yml/badge.svg)](https://github.com/vessel-api/vesselapi-go/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/vessel-api/vesselapi-go/v3.svg)](https://pkg.go.dev/github.com/vessel-api/vesselapi-go/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/vessel-api/vesselapi-go/v3)](https://goreportcard.com/report/github.com/vessel-api/vesselapi-go/v3)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Go client for the [Vessel Tracking API](https://vesselapi.com) — maritime vessel tracking, port events, emissions, and navigation data.

**Resources**: [Documentation](https://vesselapi.com/docs) | [API Explorer](https://vesselapi.com/api-reference) | [Dashboard](https://dashboard.vesselapi.com) | [Contact Support](mailto:support@vesselapi.com)

## Install

```bash
go get github.com/vessel-api/vesselapi-go/v3
```

Requires Go 1.22+.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vesselapi "github.com/vessel-api/vesselapi-go/v3"
)

func main() {
	client, err := vesselapi.NewVesselClient(os.Getenv("VESSELAPI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Search for a vessel by name.
	result, err := client.Search.Vessels(ctx, &vesselapi.GetSearchVesselsParams{
		FilterName: vesselapi.Ptr("Ever Given"),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range vesselapi.Deref(result.Vessels) {
		fmt.Printf("%s (IMO %d)\n", vesselapi.Deref(v.Name), vesselapi.Deref(v.Imo))
	}

	// Get a port by UN/LOCODE.
	port, err := client.Ports.Get(ctx, "NLRTM")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(vesselapi.Deref(port.Port.Name))

	// Auto-paginate through port events.
	it := client.PortEvents.ListAll(ctx, &vesselapi.GetPorteventsParams{
		PaginationLimit: vesselapi.Ptr(10),
	})
	for it.Next() {
		event := it.Value()
		fmt.Printf("%s at %s\n", vesselapi.Deref(event.Event), vesselapi.Deref(event.Timestamp))
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}
```

## Available Services

| Service | Methods | Description |
|---------|---------|-------------|
| `Vessels` | `Get`, `Position`, `Casualties`, `Classification`, `Emissions`, `ETA`, `Inspections`, `InspectionDetail`, `Ownership`, `Positions` | Vessel details, positions, and records ([docs](https://vesselapi.com/docs/vessels)) |
| `Ports` | `Get` | Port lookup by UN/LOCODE ([docs](https://vesselapi.com/docs/ports)) |
| `PortEvents` | `List`, `ByPort`, `ByPorts`, `ByVessel`, `LastByVessel`, `ByVessels` | Vessel arrival/departure events ([docs](https://vesselapi.com/docs/port-events)) |
| `Emissions` | `List` | EU MRV emissions data ([docs](https://vesselapi.com/docs/emissions)) |
| `Search` | `Vessels`, `Ports`, `DGPS`, `LightAids`, `MODUs`, `RadioBeacons` | Full-text search across entity types |
| `Location` | `VesselsBoundingBox`, `VesselsRadius`, `PortsBoundingBox`, `PortsRadius`, `DGPSBoundingBox`, `DGPSRadius`, `LightAidsBoundingBox`, `LightAidsRadius`, `MODUsBoundingBox`, `MODUsRadius`, `RadioBeaconsBoundingBox`, `RadioBeaconsRadius` | Geo queries by bounding box or radius ([docs](https://vesselapi.com/docs/navigation)) |
| `Navtex` | `List` | NAVTEX maritime safety messages ([docs](https://vesselapi.com/docs/navtex)) |

**37 methods total.**

## Vessel Lookup & Location

```go
// Get vessel details by IMO number (nil defaults to IMO; pass FilterIdType for MMSI).
vessel, err := client.Vessels.Get(ctx, "9811000", nil)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("%s (%s)\n", vesselapi.Deref(vessel.Vessel.Name), vesselapi.Deref(vessel.Vessel.VesselType))

// Get the vessel's latest AIS position.
pos, err := client.Vessels.Position(ctx, "9811000", nil)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Position: %f, %f\n",
	vesselapi.Deref(pos.VesselPosition.Latitude),
	vesselapi.Deref(pos.VesselPosition.Longitude),
)

// Find all vessels within 10 km of Rotterdam.
nearby, err := client.Location.VesselsRadius(ctx, &vesselapi.GetLocationVesselsRadiusParams{
	FilterLatitude:  vesselapi.Ptr(51.9225),
	FilterLongitude: vesselapi.Ptr(4.47917),
	FilterRadius:    10000,
})
if err != nil {
	log.Fatal(err)
}
for _, v := range vesselapi.Deref(nearby.Vessels) {
	fmt.Printf("%s at %f, %f\n",
		vesselapi.Deref(v.VesselName),
		vesselapi.Deref(v.Latitude),
		vesselapi.Deref(v.Longitude),
	)
}
```

## Error Handling

All methods return `*APIError` on non-2xx responses. Use `errors.As` to inspect:

```go
var apiErr *vesselapi.APIError
if errors.As(err, &apiErr) {
	if apiErr.IsNotFound() {
		// Handle 404
	}
	if apiErr.IsRateLimited() {
		// Back off (automatic retries handle most 429s)
	}
	if apiErr.IsAuthError() {
		// Check API key
	}
	fmt.Println(apiErr.StatusCode, apiErr.Message)
}
```

## Auto-Pagination

Every list endpoint has an `All*` variant returning an `Iterator`:

```go
it := client.Search.AllVessels(ctx, &vesselapi.GetSearchVesselsParams{
	FilterVesselType: vesselapi.Ptr("Tanker"),
})
for it.Next() {
	vessel := it.Value()
	// ...
}
if err := it.Err(); err != nil {
	log.Fatal(err)
}

// Or collect a bounded set at once:
vessels, err := client.Search.AllVessels(ctx, &vesselapi.GetSearchVesselsParams{
	FilterVesselType: vesselapi.Ptr("Tanker"),
	PaginationLimit:  vesselapi.Ptr(50),
}).Collect()
```

## Configuration

```go
client, err := vesselapi.NewVesselClient(apiKey,
	vesselapi.WithVesselBaseURL("https://custom-endpoint.example.com/v1"),
	vesselapi.WithVesselHTTPClient(&http.Client{Timeout: 60 * time.Second}),
	vesselapi.WithVesselUserAgent("my-app/1.0"),
	vesselapi.WithVesselRetry(5), // default: 3
)
```

Retries use exponential backoff with jitter on 429 and 5xx responses. The `Retry-After` header is respected.

## Documentation

- [API Documentation](https://vesselapi.com/docs) — endpoint guides, request/response schemas, and usage examples
- [API Explorer](https://vesselapi.com/api-reference) — interactive API reference to try endpoints in the browser
- [Dashboard](https://dashboard.vesselapi.com) — manage API keys and monitor usage
- [pkg.go.dev](https://pkg.go.dev/github.com/vessel-api/vesselapi-go/v3) — Go package reference

## Contributing & Support

Found a bug, have a feature request, or need help? You're welcome to [open an issue](https://github.com/vessel-api/vesselapi-go/issues). For API-level bugs and feature requests, please use the [main VesselAPI repository](https://github.com/vessel-api/VesselApi/issues) — see the [contributing guide](https://github.com/vessel-api/VesselApi/blob/master/CONTRIBUTING.md) for details.

For security vulnerabilities, **do not** open a public issue — email security@vesselapi.com instead. See [SECURITY.md](SECURITY.md).

For account or billing questions, contact support@vesselapi.com.

## Generation

Types and low-level client generated by [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI spec. Wrapper layer, retry logic, pagination, and tests designed and reviewed through iterative AI-assisted development.

## License

[MIT](LICENSE)
