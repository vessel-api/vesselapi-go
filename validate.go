package vesselapi

import (
	"errors"
	"fmt"
	"math"
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

// requireString rejects an empty value for a required string parameter.
func requireString(param, value string) error {
	if value == "" {
		return &ParamError{Param: param, Reason: "is required"}
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
// A bounding box may legitimately cross the antimeridian, so lonLeft is allowed
// to be greater than lonRight. Latitude does not wrap, so an inverted latitude
// pair is always a mistake and is rejected.
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
	if err := checkLongitude("filter.longitude", longitude); err != nil {
		return err
	}
	return checkLatitude("filter.latitude", latitude)
}
