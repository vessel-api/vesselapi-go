package vesselapi

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// ErrInvalidParams matches any error returned when a request is rejected before
// it is sent, because a required parameter is missing or a value is out of
// range. Match it with errors.Is:
//
//	if errors.Is(err, vesselapi.ErrInvalidParams) {
//		// the request was never sent
//	}
var ErrInvalidParams = errors.New("invalid parameters")

// ParamError reports a parameter that failed validation before the request was
// sent.
//
// The API marks these parameters required, but Go cannot distinguish a field
// that was never set from one deliberately set to zero or to an empty string,
// so a missing value cannot be caught at compile time. It is caught here
// instead, which turns what would have been an HTTP 400 into a local error that
// names the parameter.
type ParamError struct {
	// Param is the API parameter name, for example "filter.latTop".
	Param string

	// Reason describes what was wrong with the value.
	Reason string
}

func (e *ParamError) Error() string {
	return fmt.Sprintf("vesselapi: %s %s", e.Param, e.Reason)
}

// Is reports whether target is ErrInvalidParams, so that any validation failure
// can be matched with a single errors.Is check.
func (e *ParamError) Is(target error) bool { return target == ErrInvalidParams }

// The checks below run at request time. A required value can only be enforced
// at compile time through a positional argument, and taking four bare float64
// coordinates positionally would make a silent lat/lon ordering mistake easy.

// requireString rejects an empty value for a required string parameter. A value
// of nothing but whitespace is empty for this purpose: the service rejects it,
// and it is what an unfilled form field arrives as.
func requireString(param, value string) error {
	if strings.TrimSpace(value) == "" {
		return &ParamError{Param: param, Reason: "is required"}
	}
	return nil
}

// requirePathParam rejects a path segment that is missing, or that would change
// which endpoint the request reaches.
//
// These are interpolated into the request path rather than sent as query
// parameters, so an empty value silently produces a different, usually
// nonexistent route: Vessels.Get(ctx, "") requests /vessel/ and comes back as a
// puzzling 404. Worse, url.PathEscape escapes a slash but not a dot, so ".."
// survives into the path and resolves to the parent, sending the request to the
// API root instead of the endpoint the caller named.
func requirePathParam(param, value string) error {
	if strings.TrimSpace(value) == "" {
		return &ParamError{Param: param, Reason: "is required"}
	}
	if strings.ContainsAny(value, "/\\") || strings.Trim(value, ".") == "" {
		return &ParamError{
			Param:  param,
			Reason: fmt.Sprintf("must be a single path segment, got %q", value),
		}
	}
	return nil
}

// checkLatitude rejects a latitude outside the valid range.
func checkLatitude(param string, v float64) error {
	// NaN fails every comparison, so it has to be rejected explicitly or it
	// passes straight through the range check.
	if math.IsNaN(v) || v < -90 || v > 90 {
		return &ParamError{Param: param, Reason: fmt.Sprintf("must be between -90 and 90, got %g", v)}
	}
	return nil
}

// checkLongitude rejects a longitude outside the valid range.
func checkLongitude(param string, v float64) error {
	if math.IsNaN(v) || v < -180 || v > 180 {
		return &ParamError{Param: param, Reason: fmt.Sprintf("must be between -180 and 180, got %g", v)}
	}
	return nil
}

// validateBoundingBox checks the four edges the API requires.
//
// The service requires lonRight greater than lonLeft and latTop greater than
// latBottom, and rejects anything else with HTTP 400. It has no antimeridian
// handling: a box crossing 180° has to be issued as two requests, one either
// side. Both inversions are therefore caught here rather than sent.
func validateBoundingBox(lonLeft, lonRight, latBottom, latTop float64) error {
	if lonLeft == 0 && lonRight == 0 && latBottom == 0 && latTop == 0 {
		return &ParamError{
			Param:  "filter.lonLeft, filter.lonRight, filter.latBottom, filter.latTop",
			Reason: "are required, and all four are zero; set the edges of the bounding box",
		}
	}
	if err := checkLongitude("filter.lonLeft", lonLeft); err != nil {
		return err
	}
	if err := checkLongitude("filter.lonRight", lonRight); err != nil {
		return err
	}
	if err := checkLatitude("filter.latBottom", latBottom); err != nil {
		return err
	}
	if err := checkLatitude("filter.latTop", latTop); err != nil {
		return err
	}
	// A degenerate box can never match anything, and it is what forgetting one
	// pair of edges looks like: set only the latitudes and both longitudes stay
	// at zero, giving a zero-width box on the prime meridian.
	if lonLeft == lonRight {
		return &ParamError{
			Param:  "filter.lonLeft, filter.lonRight",
			Reason: fmt.Sprintf("are equal (%g), which is a box of zero width", lonLeft),
		}
	}
	if latBottom == latTop {
		return &ParamError{
			Param:  "filter.latBottom, filter.latTop",
			Reason: fmt.Sprintf("are equal (%g), which is a box of zero height", latBottom),
		}
	}
	if latBottom > latTop {
		return &ParamError{
			Param:  "filter.latBottom",
			Reason: fmt.Sprintf("must not be north of filter.latTop (%g > %g)", latBottom, latTop),
		}
	}
	if lonLeft > lonRight {
		return &ParamError{
			Param: "filter.lonLeft",
			Reason: fmt.Sprintf("must not be east of filter.lonRight (%g > %g); "+
				"a box crossing the antimeridian has to be issued either side of 180°", lonLeft, lonRight),
		}
	}
	return nil
}

// validateRadius checks the centre point and the radius the API requires.
func validateRadius(longitude, latitude, radius float64) error {
	if math.IsNaN(radius) || math.IsInf(radius, 0) || radius <= 0 {
		return &ParamError{
			Param:  "filter.radius",
			Reason: fmt.Sprintf("is required and must be greater than zero, got %g", radius),
		}
	}
	// A centre of exactly 0,0 is the signature of a forgotten field rather than
	// a real query: it is a point in the Gulf of Guinea, and a caller who wants
	// it can nudge one coordinate. Without this, omitting the centre searches
	// there silently, which is the failure this validation exists to prevent.
	if longitude == 0 && latitude == 0 {
		return &ParamError{
			Param: "filter.longitude, filter.latitude",
			Reason: "are required, and both are zero; set the centre point. " +
				"To search at 0°,0° deliberately, offset a coordinate slightly (0.000001)",
		}
	}
	if err := checkLongitude("filter.longitude", longitude); err != nil {
		return err
	}
	return checkLatitude("filter.latitude", latitude)
}
