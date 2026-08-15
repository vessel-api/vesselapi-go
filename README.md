# vesselapi-go

[![CI](https://github.com/vessel-api/vesselapi-go/actions/workflows/ci.yml/badge.svg)](https://github.com/vessel-api/vesselapi-go/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/vessel-api/vesselapi-go/v4.svg)](https://pkg.go.dev/github.com/vessel-api/vesselapi-go/v4)
[![Go Report Card](https://goreportcard.com/badge/github.com/vessel-api/vesselapi-go/v4)](https://goreportcard.com/report/github.com/vessel-api/vesselapi-go/v4)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Go client for the [Vessel Tracking API](https://vesselapi.com): maritime vessel tracking, port events, emissions, and navigation data.

**Resources**: [Documentation](https://vesselapi.com/docs) | [API Explorer](https://vesselapi.com/api-reference) | [Dashboard](https://dashboard.vesselapi.com) | [Contact Support](mailto:support@vesselapi.com)

## Install

```bash
go get github.com/vessel-api/vesselapi-go/v4
```

Requires Go 1.24+ (raised in v4: the generated client depends on oapi-codegen/runtime v1.6.0, which requires Go 1.24).

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vesselapi "github.com/vessel-api/vesselapi-go/v4"
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
| `Vessels` | `Get`, `Position`, `Casualties`, `Emissions`, `ETA`, `Positions` | Vessel details, positions, and records ([docs](https://vesselapi.com/docs/vessels)) |
| `Ports` | `Get`, `Inbound` | Port lookup by UN/LOCODE and inbound vessels ([docs](https://vesselapi.com/docs/ports)) |
| `PortEvents` | `List`, `ByPort`, `ByPorts`, `ByVessel`, `LastByVessel`, `ByVessels` | Vessel arrival/departure events ([docs](https://vesselapi.com/docs/port-events)) |
| `Emissions` | `List` | EU MRV emissions data ([docs](https://vesselapi.com/docs/emissions)) |
| `Search` | `Vessels`, `Ports`, `DGPS`, `LightAids`, `MODUs`, `RadioBeacons` | Full-text search across entity types |
| `Location` | `VesselsBoundingBox`, `VesselsRadius`, `PortsBoundingBox`, `PortsRadius`, `DGPSBoundingBox`, `DGPSRadius`, `LightAidsBoundingBox`, `LightAidsRadius`, `MODUsBoundingBox`, `MODUsRadius`, `RadioBeaconsBoundingBox`, `RadioBeaconsRadius` | Geo queries by bounding box or radius ([docs](https://vesselapi.com/docs/navigation)) |

**33 methods total**, one per API endpoint, plus 28 auto-pagination iterators.

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
	FilterLatitude:  51.9225,
	FilterLongitude: 4.47917,
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
	if apiErr.IsPaymentRequired() {
		// Out of satellite credits
	}
	if apiErr.IsForbidden() {
		// Key suspended, or feature not on your plan
	}
	fmt.Println(apiErr.StatusCode, apiErr.Message)
}
```

| Status | Helper | Meaning |
|--------|--------|---------|
| 401 | `IsAuthError()` | API key missing, invalid or revoked |
| 402 | `IsPaymentRequired()` | Out of satellite credits, with no stored position to fall back on |
| 403 | `IsForbidden()` | Key suspended for sustained quota abuse, or feature not on your plan |
| 404 | `IsNotFound()` | Resource does not exist |
| 429 | `IsRateLimited()` | Rate or quota limit hit (retried automatically) |

`APIError` also carries the API's machine-readable `Code` and `Type` fields.
Branch on `Code` rather than on `Message`, which is free-form text and may be
reworded:

```go
_, err := client.Vessels.Position(ctx, "9811000", &vesselapi.GetVesselIdPositionParams{
	FilterSat: vesselapi.Ptr(true),
})

var apiErr *vesselapi.APIError
if errors.As(err, &apiErr) {
	switch {
	case apiErr.Code == vesselapi.ErrorCodeInsufficientCredits:
		// 402: top up credits, or retry without the satellite fallback.
		pos, err := client.Vessels.Position(ctx, "9811000", nil)
		_, _ = pos, err
	case apiErr.Code == vesselapi.ErrorCodeFeatureNotAvailable:
		// 403: the endpoint is not part of this plan.
	case apiErr.IsForbidden():
		// 403 with ErrorCodeForbidden: key suspended; see the Retry-After header.
	}
}
```

`Code` and `Type` are empty when a response carries no structured error body.
402 is only returned by `Vessels.Position` when `FilterSat` is true; 403 can come
from any endpoint. Neither status is retried, since retrying cannot succeed.

## Auto-Pagination

Every list endpoint has an iterator variant. Most are named `All*`
(`AllVessels`, `AllByPort`); a few read better the other way round
(`ListAll`, `InboundAll`).

```go
it := client.Search.AllVessels(ctx, &vesselapi.GetSearchVesselsParams{
	FilterVesselType: &[]string{"Tanker"},
})
for it.Next() {
	vessel := it.Value()
	// ...
}
if err := it.Err(); err != nil {
	log.Fatal(err)
}

// Collect gathers every page, so it makes as many requests as the result set
// needs. PaginationLimit sets the page size, not a total, and does not bound
// the walk. Use it on result sets you know are small:
port, err := client.Search.AllPorts(ctx, &vesselapi.GetSearchPortsParams{
	FilterCountry: &[]string{"NL"},
}).Collect()

// To take a fixed number, stop the loop yourself:
var first50 []vesselapi.Vessel
tankers := client.Search.AllVessels(ctx, &vesselapi.GetSearchVesselsParams{
	FilterVesselType: &[]string{"Tanker"},
})
for tankers.Next() && len(first50) < 50 {
	first50 = append(first50, tankers.Value())
}
if err := tankers.Err(); err != nil {
	log.Fatal(err)
}
```

`Collect` returns whatever it gathered before a failure alongside the error, so
check the error before treating the result as complete.

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

- [API Documentation](https://vesselapi.com/docs): endpoint guides, request/response schemas, and usage examples
- [API Explorer](https://vesselapi.com/api-reference): interactive API reference to try endpoints in the browser
- [Dashboard](https://dashboard.vesselapi.com): manage API keys and monitor usage
- [pkg.go.dev](https://pkg.go.dev/github.com/vessel-api/vesselapi-go/v4): Go package reference

## Contributing & Support

Found a bug, have a feature request, or need help? You're welcome to [open an issue](https://github.com/vessel-api/vesselapi-go/issues). For API-level bugs and feature requests, please use the [main VesselAPI repository](https://github.com/vessel-api/VesselApi/issues). See the [contributing guide](https://github.com/vessel-api/VesselApi/blob/master/CONTRIBUTING.md) for details.

For security vulnerabilities, **do not** open a public issue. Email security@vesselapi.com instead. See [SECURITY.md](SECURITY.md).

For account or billing questions, contact support@vesselapi.com.

## Generation

Types and low-level client are generated by [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from `openapi/openapi.json`, which is committed to this repository. Regenerate with `make generate`. The wrapper layer, retry logic, pagination and tests are hand-written.

## Data Sources & Attribution

Emissions and casualty data: © European Union. Source: European Maritime Safety Agency
(EMSA): THETIS-MRV (EU MRV, Regulation (EU) 2015/757) and the European Marine Casualty
Information Platform (EMCIP). Reused under the European Commission reuse notice
(Commission Decision 2011/833/EU), which authorises reuse for commercial and
non-commercial purposes with acknowledgement of the source. Data may be transformed and
combined; EMSA does not endorse this service.

## License

[MIT](LICENSE)
