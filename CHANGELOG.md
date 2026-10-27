# Changelog

Notable changes in this release. Earlier versions predate this file.
This project follows [Semantic Versioning](https://semver.org/).

## [4.1.0] - 2026-10-26

### Added

- `VesselPositionResponse` has three new fields, which the API returns on every
  position response since 25 October 2026: `Vessel` (the same object as
  `Vessels.Get`), `VesselEta` (the same object as `Vessels.ETA`) and
  `PositionCheck`. `Vessel` and `VesselEta` are nil when the API holds no record
  for them. Earlier versions drop these keys when decoding.
- `PositionCheck` compares the position with the vessel's previous report from
  the last 80 hours. It gives the time and great-circle distance since then, and
  the speed and course that would cover that distance. It also says whether that
  speed is more than 1.2 times the vessel's usual top speed, and whether the MMSI
  and IMO match.
- `GetVesselIdPositionParams` has `FilterSatLookbackMinutes` and
  `FilterSatMaxAgeMinutes`, which set the satellite fallback's time windows.
- `GetVesselIdEtaResponse` and `GetPortUnlocodeInboundResponse` decode a 503 body
  into `JSON503` (`ServiceUnavailableErrorResponse`).

### Changed

- The bundled OpenAPI spec is synced with the API. Some field descriptions and
  examples changed with it.

### Fixed

- `make generate` failed at its first `sed` step, so the generator's package
  comment and placeholder server URLs stayed in `generated.go`.

## [4.0.1] - 2026-10-08

Documentation only. No function, parameter or type changed.

### Fixed

- Doc comments and the bundled OpenAPI spec named the satellite fallback
  parameter `sat`; the API's parameter is `filter.sat`.
- The 4.0.0 security advisory now lists 3.1.1 among the affected versions.

### Changed

- README: the radius example describes the time window it searches (default
  2 hours, at most 4), and a note explains that v1 to v3 are deprecated and that
  their pkg.go.dev pages still list retired methods.

## [4.0.0] - 2026-08-15

Vessel ownership, classification and inspection data, and NAVTEX messages, were
retired from the VesselAPI service on **10 August 2026**. This release removes
the corresponding client surface and brings the module back in line with the API.

### Security

Releases before 4.0.0 attach the API key inside an HTTP round-tripper, so it is
applied again to each request made while following a redirect. Go's own
protection, which drops sensitive headers when a redirect changes host, covers
only headers set on the original request, so it does not apply here. A redirect
to another host therefore arrives carrying a valid key.

| Module path | Affected |
|---|---|
| `github.com/vessel-api/vesselapi-go` | 1.0.0, and 2.0.0 if vendored |
| `github.com/vessel-api/vesselapi-go/v3` | 3.0.0, 3.1.0, 3.1.1 |

2.0.0 cannot be resolved by `go get` at any path, so it is only reachable by
vendoring or a `replace` directive.

4.0.0 refuses any redirect leaving the configured origin, including a downgrade
from https to http on the same host. A redirect policy you supply yourself now
runs after that check rather than replacing it, so a permissive policy no longer
disables it.

If you cannot upgrade yet, pass a client that refuses off-origin redirects. This
works on 1.x, 2.x and 3.x:

```go
guard := &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if req.URL.Scheme != via[0].URL.Scheme || req.URL.Host != via[0].URL.Host {
			return fmt.Errorf("refusing off-origin redirect")
		}
		return nil
	},
}
client, err := vesselapi.NewVesselClient(apiKey, vesselapi.WithVesselHTTPClient(guard))
```

The disclosure is silent: the call returns no error, so logs will not show it.
Rotate the key if your deployment could have followed a redirect off your
configured host. Report security issues to security@vesselapi.com rather than a
public issue; see [SECURITY.md](SECURITY.md).

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

Beyond the removals, several signature and type changes break at compile time.
They are listed under Changed (breaking) below, and the compiler will point at
each one. The two most likely to affect an existing application are the location
coordinate params, which are now plain `float64` rather than `*float64` and so
touch every geo call site, and the seven search filters that now take slices.

### Requirements

**Go 1.24 or later** (was 1.22). The generated client now depends on
`oapi-codegen/runtime` v1.6.0, which requires Go 1.24.

### Added

- **Search**: `Q` (free-text across name, IMO, MMSI, ENI and callsign) and
  `FilterEni`.
- **`FilterSat`** on `Vessels.Position()`. ⚠️ Satellite lookups draw on a
  prepaid balance and are charged per call. Omit it to serve stored positions.
- **`TimeFrom` / `TimeTo`** on `PortEvents.ByPort()`, `ByPorts()` and
  `ByVessels()`. `List()` and `ByVessel()` already had them.
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

- **Required parameters are now checked before the request is sent.** Because
  the coordinate fields stopped being pointers, a zero-valued params struct now
  compiles and would send explicit zeros where v3 sent nothing at all. Those
  calls return a `*ParamError` without sending anything. Match any of them with
  `errors.Is(err, vesselapi.ErrInvalidParams)`.
  - `Location.*BoundingBox()`: all four edges.
  - `Location.*Radius()`: the centre point, and a radius greater than zero.
  - `Search.DGPS()`, `Search.LightAids()`, `Search.MODUs()`,
    `Search.RadioBeacons()`: `filter.name`.
  - `PortEvents.ByPorts()`: `filter.portName`.
    `PortEvents.ByVessels()`: `filter.vesselName`.
  - `Vessels.Positions()`: `filter.ids`.
  - The `All*` iterators report the same failure through `Err()`.
- **The vessel id and UN/LOCODE arguments are now checked too.** An empty value
  used to be interpolated into the path and requested a route that does not
  exist, arriving as a puzzling 404; a value of `..` resolved to the parent and
  sent the request to the API root instead of the endpoint named. Both are now
  a `*ParamError`, across all fifteen methods that take one.
- **An unusable base URL is rejected by `NewVesselClient`** rather than failing
  on every call. A query string or fragment on it is dropped, since the endpoint
  path is appended to it and anything after would be discarded along with the
  base path.
- **Passing `nil` params to those 19 methods is now an error** instead of a
  request full of zero values.
- **`Ports.Inbound()` and `Ports.InboundAll()` take a different params type.**
  They existed in v3; what changed is the type, from `*PortInboundParams` to
  `*GetPortUnlocodeInboundParams`, matching every other method. `EtaFrom` and
  `EtaTo` become `FilterEtaFrom` and `FilterEtaTo`, and are now optional
  pointers: omit them and the service covers now to 72 hours ahead.
- **Coordinates on the twelve location methods are no longer pointers.**
  `FilterLonLeft`, `FilterLonRight`, `FilterLatBottom`, `FilterLatTop`,
  `FilterLatitude` and `FilterLongitude` changed from `*float64` to `float64`
  across every `Location.*BoundingBox()` and `Location.*Radius()` params type.
  `FilterRadius` was already a plain `float64`. Pass the value directly instead of a pointer. This is why those
  methods now validate: an unset field is no longer distinguishable from a
  deliberate zero, so a forgotten edge would otherwise be sent as 0.
- **`FilterEventType` on `PortEvents.List()`** is now
  `*GetPorteventsParamsFilterEventType` rather than `*string`. `ByVessel()`
  already used a named type.
- **`FeatureNumber` on `DGPSStation` and `RadioBeacon`** changed from `*int` to
  `*float64`.
- **Six generated types were renamed.** Five shed a long
  `GithubComVesselapiCommonVesselDataContracts…` prefix and are now
  `ContractsPortCountry`, `ContractsVesselFormerName`, `GeoJSON`,
  `PortReference` and `VesselReference`. The sixth went the other way:
  `ResolutionMeta` is now `TypesResolutionMeta`.
- **Seven search filters became slices**, so a filter can match any of several
  values. Existing call sites need the value wrapped in a slice.
  - `Search.Vessels()`: `FilterFlag`, `FilterVesselType`, `*string` to `*[]string`.
  - `Search.Ports()`: `FilterCountry`, `FilterSize`, `FilterHarborSize`,
    `FilterHarborUse`, `*string` to `*[]string`.
  - `Search.Ports()`: `FilterType`, `*string` to
    `*[]GetSearchPortsParamsFilterType`, a named type rather than a plain string.
  - `FilterCountry` on the port-event methods is unchanged.
- Coordinates are range-checked (latitude within -90 to 90, longitude within
  -180 to 180). NaN and infinities are rejected, as are a box of zero width or
  height and either edge pair given in the wrong order. The service has no
  antimeridian handling and rejects `lonLeft` greater than `lonRight`, so a box
  crossing 180° has to be issued as two requests, one either side. Note that
  the service also caps a box at four degrees of total span
  (`|dLat| + |dLon| ≤ 4`), which the SDK does not check because it is a service
  limit rather than a malformed request; it surfaces as an HTTP 400.

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
  See Security above for affected versions and a workaround for anyone still on
  1.x, 2.x or 3.x.
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
- Retry backoff could overflow to a negative duration and turn into a tight
  loop, both at high retry counts and on a large `Retry-After` value, where the
  wrapped duration was clamped to zero and retried at once.
- **An error response whose body could not be decoded lost its status code.** A
  declared error status arriving with an empty or malformed body while claiming
  a JSON content type surfaced as a bare syntax error, so `errors.As` did not
  match `*APIError` and `IsNotFound()`, `IsRateLimited()` and the rest never
  fired. Such a response now returns an `*APIError` carrying the status. This
  needs something between you and the service to produce it, such as a proxy or
  rate limiter answering in its own format; the API itself does not.
- **29 numeric fields were `float32` and are now `float64`**, among them `Sog`,
  `Cog`, `Draught`, the port depth figures and every emissions figure, plus
  `GeoJSON.Coordinates` from `*[]float32` to `*[]float64`. Latitude and
  longitude were always `float64` and are unaffected.
- The `Version` constant reported `1.0.0` across v2 and v3, so the User-Agent
  misreported the installed version.

### Attribution

Emissions and casualty data: © European Union. Source: European Maritime Safety
Agency (EMSA), THETIS-MRV (EU MRV, Regulation (EU) 2015/757) and the European
Marine Casualty Information Platform (EMCIP). See the README.
