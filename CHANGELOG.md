# Changelog

All notable changes to this module are documented here.
This project follows [Semantic Versioning](https://semver.org/).

## [4.0.0] - 2026-08-15

Vessel ownership, classification and inspection data, and NAVTEX messages, were
retired from the VesselAPI service on **10 August 2026**. This release removes
the corresponding client surface and brings the module back in line with the API.

### Import path

v4 changes the module path, as Go requires for a major version:

```go
import vesselapi "github.com/vessel-api/vesselapi-go/v4"
```

```
go get github.com/vessel-api/vesselapi-go/v4
```

### Removed (breaking)

These endpoints no longer exist. Calling them returned HTTP 410 from 10 August
2026 onward; the methods are now gone so the failure surfaces at compile time.

| Removed | Replacement |
|---|---|
| `Vessels.Classification()` | none |
| `Vessels.Inspections()` | none |
| `Vessels.InspectionDetail()` | none |
| `Vessels.Ownership()` | none |
| `client.Navtex` (whole service: `List()`, `ListAll()`) | none |

Also removed: the `FilterClassSociety` and `FilterOwner` search parameters (the
API began rejecting both with HTTP 400), the `Builder`, `ClassSociety`,
`OwnerName` and `ManagerName` fields on `Vessel`, and the generated types for
the removed endpoints.

**Migrating from v3:** update the import path to `/v4`, then remove calls to the
methods above. There is no alternative source for that data through this API.

Three further changes break at compile time, so the compiler will find them for
you: the six search filters listed under Changed now take a slice, `Ports.Inbound`
takes a different params type, and `ResolutionMeta` is now `TypesResolutionMeta`.

### Requirements

**Go 1.24 or later** (was 1.22). The generated client now depends on
`oapi-codegen/runtime` v1.6.0, which requires Go 1.24.

### Added

- **Search**: `Q` (free-text across name, IMO, MMSI and callsign) and
  `FilterEni`.
- **`FilterSat`** on `Vessels.Position()`. ⚠️ Satellite lookups draw on a
  prepaid balance and are charged per call. Omit it to serve stored positions.
- **`TimeFrom` / `TimeTo`** on the port-event methods.
- **Vessel fields**: `Eni`, `Teu`, `VesselSubtype`, `NameAis`, `SummerDraught`,
  `DraughtCalculatedAvg`, `DraughtObservedMax`, `SpeedCalculatedAvg`,
  `SpeedObservedMax`.
- **Error handling**: `(*APIError).IsPaymentRequired()` (402, satellite credits
  exhausted) and `(*APIError).IsForbidden()` (403, key suspended or feature not
  on your plan), alongside the existing `IsNotFound()` / `IsRateLimited()` /
  `IsAuthError()`.
  `APIError` now carries the API's machine-readable `Code` and `Type`. Branch
  on those rather than on the message, which may be reworded.

### Changed (breaking)

- **Required parameters are now checked before the request is sent.** A
  zero-valued params struct used to compile and then query a zero-area bounding
  box. Those calls now return a `*ParamError` without sending anything. Match
  any of them with `errors.Is(err, vesselapi.ErrInvalidParams)`.
  - `Location.*BoundingBox()`: all four edges.
  - `Location.*Radius()`: the centre point, and a radius greater than zero.
  - `Search.DGPS()`, `Search.LightAids()`, `Search.MODUs()`,
    `Search.RadioBeacons()`: `filter.name`.
  - `PortEvents.ByPorts()`: `filter.portName`.
    `PortEvents.ByVessels()`: `filter.vesselName`.
  - `Vessels.Positions()`: `filter.ids`.
  - The `All*` iterators report the same failure through `Err()`.
- **Passing `nil` params to those 19 methods is now an error** instead of a
  request full of zero values.
- **`Ports.Inbound()` and `Ports.InboundAll()` take a different params type.**
  They existed in v3; what changed is the type, from `*PortInboundParams` to
  `*GetPortUnlocodeInboundParams`, matching every other method. `EtaFrom` and
  `EtaTo` become `FilterEtaFrom` and `FilterEtaTo`, and are now optional
  pointers: omit them and the service covers now to 72 hours ahead.
- **Six search filters became slices.** `FilterFlag` and `FilterVesselType` on
  `Search.Vessels()`, and `FilterCountry`, `FilterSize`, `FilterHarborSize` and
  `FilterType` on `Search.Ports()`, changed from `*string` to `*[]string`, so a
  filter can match any of several values. Existing call sites need the value
  wrapped in a slice. `FilterCountry` on the port-event methods is unchanged.
- Coordinates are range-checked (latitude within -90 to 90, longitude within
  -180 to 180). NaN and infinities are rejected, as are an inverted latitude
  pair and a box of zero width or height. A box crossing the antimeridian is
  still valid, so `lonLeft` may be greater than `lonRight`.

### Fixed

- **An empty `filter.idType` was sent whenever a params struct was supplied.**
  Twelve methods and four iterators applied the `imo` default only when the
  whole struct was nil, so passing `&Params{}`, or setting any other field,
  put an empty value on the wire for a parameter the API requires. The
  default now applies whenever the field is empty, and the caller's struct is
  not modified.
- **Auto-paginating iterators stopped at the first empty page**, even when it
  carried a token, dropping every later page with no error. They now continue
  until the service stops issuing a token, and stop if it ever repeats one.
- **The API key was reattached to redirects that left the configured origin.**
  A redirect to another host, or a downgrade from https to http on the same
  host, is now refused rather than followed with a valid token attached.
- **Auto-pagination could mix results from different moments.** Nine endpoints
  resolve their time window server-side relative to each request, so every page
  of a walk covered a slightly different span. On the vessel bounding box and
  radius endpoints this is worse than untidy: their tokens are only valid
  alongside the bounds they were issued with, and mismatched bounds are not
  rejected. Each iterator now fixes its window, using that endpoint's own
  default, before the first page and reuses it for the rest.
- **An iterator stopping early now reports it.** Previously a short read looked
  identical to a complete one. `Collect` also returns the items gathered before
  a failure instead of discarding them.
- Retry backoff could overflow to a negative duration at high retry counts,
  which made the timer fire at once and turned backoff into a tight loop.
- **29 numeric fields were `float32` and are now `float64`**, among them `Sog`,
  `Cog`, `Draught`, the port depth figures and every emissions figure. Latitude
  and longitude were always `float64` and are unaffected.
- The `Version` constant reported `1.0.0` across v2 and v3, so the User-Agent
  misreported the installed version.

### Attribution

Emissions and casualty data: © European Union. Source: European Maritime Safety
Agency (EMSA), THETIS-MRV (EU MRV, Regulation (EU) 2015/757) and the European
Marine Casualty Information Platform (EMCIP). See the README.
