package vesselapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// Defines values for ErrorCode.
const (
	ErrorCodeBoundingBoxTooDense  ErrorCode = "bounding_box_too_dense"
	ErrorCodeConflict             ErrorCode = "conflict"
	ErrorCodeDatabaseError        ErrorCode = "database_error"
	ErrorCodeEndpointRetired      ErrorCode = "endpoint_retired"
	ErrorCodeFeatureNotAvailable  ErrorCode = "feature_not_available"
	ErrorCodeForbidden            ErrorCode = "forbidden"
	ErrorCodeInsufficientCredits  ErrorCode = "insufficient_credits"
	ErrorCodeInternalError        ErrorCode = "internal_error"
	ErrorCodeInvalidAPIKey        ErrorCode = "invalid_api_key"
	ErrorCodeInvalidCoordinates   ErrorCode = "invalid_coordinates"
	ErrorCodeInvalidIMO           ErrorCode = "invalid_imo"
	ErrorCodeInvalidMMSI          ErrorCode = "invalid_mmsi"
	ErrorCodeInvalidParameter     ErrorCode = "invalid_parameter"
	ErrorCodeInvalidTimeRange     ErrorCode = "invalid_time_range"
	ErrorCodeMissingParameter     ErrorCode = "missing_parameter"
	ErrorCodeNotificationInactive ErrorCode = "notification_inactive"
	ErrorCodePrefillPending       ErrorCode = "prefill_pending"
	ErrorCodeRateLimitExceeded    ErrorCode = "rate_limit_exceeded"
	ErrorCodeResourceMissing      ErrorCode = "resource_missing"
	ErrorCodeServiceUnavailable   ErrorCode = "service_unavailable"
)

// Valid indicates whether the value is a known member of the ErrorCode enum.
func (e ErrorCode) Valid() bool {
	switch e {
	case ErrorCodeBoundingBoxTooDense:
		return true
	case ErrorCodeConflict:
		return true
	case ErrorCodeDatabaseError:
		return true
	case ErrorCodeEndpointRetired:
		return true
	case ErrorCodeFeatureNotAvailable:
		return true
	case ErrorCodeForbidden:
		return true
	case ErrorCodeInsufficientCredits:
		return true
	case ErrorCodeInternalError:
		return true
	case ErrorCodeInvalidAPIKey:
		return true
	case ErrorCodeInvalidCoordinates:
		return true
	case ErrorCodeInvalidIMO:
		return true
	case ErrorCodeInvalidMMSI:
		return true
	case ErrorCodeInvalidParameter:
		return true
	case ErrorCodeInvalidTimeRange:
		return true
	case ErrorCodeMissingParameter:
		return true
	case ErrorCodeNotificationInactive:
		return true
	case ErrorCodePrefillPending:
		return true
	case ErrorCodeRateLimitExceeded:
		return true
	case ErrorCodeResourceMissing:
		return true
	case ErrorCodeServiceUnavailable:
		return true
	default:
		return false
	}
}

// Defines values for ErrorType.
const (
	ErrorTypeAPIError            ErrorType = "api_error"
	ErrorTypeAuthenticationError ErrorType = "authentication_error"
	ErrorTypeConflict            ErrorType = "conflict_error"
	ErrorTypeForbidden           ErrorType = "forbidden_error"
	ErrorTypeGone                ErrorType = "gone_error"
	ErrorTypeInvalidRequest      ErrorType = "invalid_request_error"
	ErrorTypeNotFoundError       ErrorType = "not_found_error"
	ErrorTypePaymentRequired     ErrorType = "payment_required_error"
	ErrorTypeRateLimitError      ErrorType = "rate_limit_error"
	ErrorTypeServiceUnavailable  ErrorType = "service_unavailable_error"
)

// Valid indicates whether the value is a known member of the ErrorType enum.
func (e ErrorType) Valid() bool {
	switch e {
	case ErrorTypeAPIError:
		return true
	case ErrorTypeAuthenticationError:
		return true
	case ErrorTypeConflict:
		return true
	case ErrorTypeForbidden:
		return true
	case ErrorTypeGone:
		return true
	case ErrorTypeInvalidRequest:
		return true
	case ErrorTypeNotFoundError:
		return true
	case ErrorTypePaymentRequired:
		return true
	case ErrorTypeRateLimitError:
		return true
	case ErrorTypeServiceUnavailable:
		return true
	default:
		return false
	}
}

// Defines values for GetPorteventsParamsFilterEventType.
const (
	GetPorteventsParamsFilterEventTypeAll       GetPorteventsParamsFilterEventType = "all"
	GetPorteventsParamsFilterEventTypeArrival   GetPorteventsParamsFilterEventType = "arrival"
	GetPorteventsParamsFilterEventTypeDeparture GetPorteventsParamsFilterEventType = "departure"
)

// Valid indicates whether the value is a known member of the GetPorteventsParamsFilterEventType enum.
func (e GetPorteventsParamsFilterEventType) Valid() bool {
	switch e {
	case GetPorteventsParamsFilterEventTypeAll:
		return true
	case GetPorteventsParamsFilterEventTypeArrival:
		return true
	case GetPorteventsParamsFilterEventTypeDeparture:
		return true
	default:
		return false
	}
}

// Defines values for GetPorteventsVesselIdParamsFilterIdType.
const (
	GetPorteventsVesselIdParamsFilterIdTypeImo  GetPorteventsVesselIdParamsFilterIdType = "imo"
	GetPorteventsVesselIdParamsFilterIdTypeMmsi GetPorteventsVesselIdParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetPorteventsVesselIdParamsFilterIdType enum.
func (e GetPorteventsVesselIdParamsFilterIdType) Valid() bool {
	switch e {
	case GetPorteventsVesselIdParamsFilterIdTypeImo:
		return true
	case GetPorteventsVesselIdParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// Defines values for GetPorteventsVesselIdParamsFilterEventType.
const (
	GetPorteventsVesselIdParamsFilterEventTypeAll       GetPorteventsVesselIdParamsFilterEventType = "all"
	GetPorteventsVesselIdParamsFilterEventTypeArrival   GetPorteventsVesselIdParamsFilterEventType = "arrival"
	GetPorteventsVesselIdParamsFilterEventTypeDeparture GetPorteventsVesselIdParamsFilterEventType = "departure"
)

// Valid indicates whether the value is a known member of the GetPorteventsVesselIdParamsFilterEventType enum.
func (e GetPorteventsVesselIdParamsFilterEventType) Valid() bool {
	switch e {
	case GetPorteventsVesselIdParamsFilterEventTypeAll:
		return true
	case GetPorteventsVesselIdParamsFilterEventTypeArrival:
		return true
	case GetPorteventsVesselIdParamsFilterEventTypeDeparture:
		return true
	default:
		return false
	}
}

// Defines values for GetPorteventsVesselIdParamsFilterSortOrder.
const (
	Asc  GetPorteventsVesselIdParamsFilterSortOrder = "asc"
	Desc GetPorteventsVesselIdParamsFilterSortOrder = "desc"
)

// Valid indicates whether the value is a known member of the GetPorteventsVesselIdParamsFilterSortOrder enum.
func (e GetPorteventsVesselIdParamsFilterSortOrder) Valid() bool {
	switch e {
	case Asc:
		return true
	case Desc:
		return true
	default:
		return false
	}
}

// Defines values for GetPorteventsVesselIdLastParamsFilterIdType.
const (
	GetPorteventsVesselIdLastParamsFilterIdTypeImo  GetPorteventsVesselIdLastParamsFilterIdType = "imo"
	GetPorteventsVesselIdLastParamsFilterIdTypeMmsi GetPorteventsVesselIdLastParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetPorteventsVesselIdLastParamsFilterIdType enum.
func (e GetPorteventsVesselIdLastParamsFilterIdType) Valid() bool {
	switch e {
	case GetPorteventsVesselIdLastParamsFilterIdTypeImo:
		return true
	case GetPorteventsVesselIdLastParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// Defines values for GetSearchPortsParamsFilterType.
const (
	GetSearchPortsParamsFilterTypePort GetSearchPortsParamsFilterType = "Port"
)

// Valid indicates whether the value is a known member of the GetSearchPortsParamsFilterType enum.
func (e GetSearchPortsParamsFilterType) Valid() bool {
	switch e {
	case GetSearchPortsParamsFilterTypePort:
		return true
	default:
		return false
	}
}

// Defines values for GetVesselIdParamsFilterIdType.
const (
	GetVesselIdParamsFilterIdTypeImo  GetVesselIdParamsFilterIdType = "imo"
	GetVesselIdParamsFilterIdTypeMmsi GetVesselIdParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetVesselIdParamsFilterIdType enum.
func (e GetVesselIdParamsFilterIdType) Valid() bool {
	switch e {
	case GetVesselIdParamsFilterIdTypeImo:
		return true
	case GetVesselIdParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// Defines values for GetVesselIdCasualtiesParamsFilterIdType.
const (
	GetVesselIdCasualtiesParamsFilterIdTypeImo  GetVesselIdCasualtiesParamsFilterIdType = "imo"
	GetVesselIdCasualtiesParamsFilterIdTypeMmsi GetVesselIdCasualtiesParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetVesselIdCasualtiesParamsFilterIdType enum.
func (e GetVesselIdCasualtiesParamsFilterIdType) Valid() bool {
	switch e {
	case GetVesselIdCasualtiesParamsFilterIdTypeImo:
		return true
	case GetVesselIdCasualtiesParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// Defines values for GetVesselIdEmissionsParamsFilterIdType.
const (
	GetVesselIdEmissionsParamsFilterIdTypeImo  GetVesselIdEmissionsParamsFilterIdType = "imo"
	GetVesselIdEmissionsParamsFilterIdTypeMmsi GetVesselIdEmissionsParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetVesselIdEmissionsParamsFilterIdType enum.
func (e GetVesselIdEmissionsParamsFilterIdType) Valid() bool {
	switch e {
	case GetVesselIdEmissionsParamsFilterIdTypeImo:
		return true
	case GetVesselIdEmissionsParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// Defines values for GetVesselIdEtaParamsFilterIdType.
const (
	GetVesselIdEtaParamsFilterIdTypeImo  GetVesselIdEtaParamsFilterIdType = "imo"
	GetVesselIdEtaParamsFilterIdTypeMmsi GetVesselIdEtaParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetVesselIdEtaParamsFilterIdType enum.
func (e GetVesselIdEtaParamsFilterIdType) Valid() bool {
	switch e {
	case GetVesselIdEtaParamsFilterIdTypeImo:
		return true
	case GetVesselIdEtaParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// Defines values for GetVesselIdPositionParamsFilterIdType.
const (
	GetVesselIdPositionParamsFilterIdTypeImo  GetVesselIdPositionParamsFilterIdType = "imo"
	GetVesselIdPositionParamsFilterIdTypeMmsi GetVesselIdPositionParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetVesselIdPositionParamsFilterIdType enum.
func (e GetVesselIdPositionParamsFilterIdType) Valid() bool {
	switch e {
	case GetVesselIdPositionParamsFilterIdTypeImo:
		return true
	case GetVesselIdPositionParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// Defines values for GetVesselsPositionsParamsFilterIdType.
const (
	GetVesselsPositionsParamsFilterIdTypeImo  GetVesselsPositionsParamsFilterIdType = "imo"
	GetVesselsPositionsParamsFilterIdTypeMmsi GetVesselsPositionsParamsFilterIdType = "mmsi"
)

// Valid indicates whether the value is a known member of the GetVesselsPositionsParamsFilterIdType enum.
func (e GetVesselsPositionsParamsFilterIdType) Valid() bool {
	switch e {
	case GetVesselsPositionsParamsFilterIdTypeImo:
		return true
	case GetVesselsPositionsParamsFilterIdTypeMmsi:
		return true
	default:
		return false
	}
}

// AuthenticationErrorDetail defines model for AuthenticationErrorDetail.
type AuthenticationErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: invalid_api_key
	Code *ErrorCode `json:"code,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: api key is invalid or not found
	Message *string `json:"message,omitempty"`

	// Type Type categorizes the error (always "authentication_error" for 401s)
	//
	// Example: authentication_error
	Type *ErrorType `json:"type,omitempty"`
}

// AuthenticationErrorResponse defines model for AuthenticationErrorResponse.
type AuthenticationErrorResponse struct {
	Error *AuthenticationErrorDetail `json:"error,omitempty"`
}

// BadRequestErrorDetail defines model for BadRequestErrorDetail.
type BadRequestErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: invalid_mmsi
	Code *ErrorCode `json:"code,omitempty"`

	// DocUrl DocURL is a link to documentation for more information
	//
	// Example: https://vesselapi.com/docs#invalid_mmsi
	DocUrl *string `json:"doc_url,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: Invalid MMSI: abc123
	Message *string `json:"message,omitempty"`

	// Param Param identifies the parameter that caused the error (if applicable)
	//
	// Example: mmsi
	Param *string `json:"param,omitempty"`

	// Type Type categorizes the error (e.g., "invalid_request_error")
	//
	// Example: invalid_request_error
	Type *ErrorType `json:"type,omitempty"`
}

// BadRequestErrorResponse defines model for BadRequestErrorResponse.
type BadRequestErrorResponse struct {
	Error *BadRequestErrorDetail `json:"error,omitempty"`
}

// DGPSStation Differential GPS (DGPS) correction signal broadcast station
type DGPSStation struct {
	// AidType Type of navigational aid
	//
	// Example: DGPS
	AidType *string `json:"aid_type,omitempty"`

	// DeleteFlag Deletion status flag
	DeleteFlag *string `json:"delete_flag,omitempty"`

	// FeatureNumber NGA feature number identifier
	//
	// Example: 1234.5
	FeatureNumber *float64 `json:"feature_number,omitempty"`

	// Frequency Broadcast frequency in kHz
	//
	// Example: 290
	Frequency *float64 `json:"frequency,omitempty"`

	// GeopoliticalHeading Country or major geographic area
	//
	// Example: UNITED STATES
	GeopoliticalHeading *string `json:"geopolitical_heading,omitempty"`

	// Location GeoJSON point for geospatial queries
	Location *GeoJSON `json:"location,omitempty"`

	// Name Station name
	//
	// Example: Cape Henry DGPS
	Name *string `json:"name,omitempty"`

	// NoticeNumber Notice to Mariners number
	//
	// Example: 12
	NoticeNumber *int `json:"notice_number,omitempty"`

	// NoticeWeek Week of the notice
	//
	// Example: 15
	NoticeWeek *string `json:"notice_week,omitempty"`

	// NoticeYear Year of the notice
	//
	// Example: 2024
	NoticeYear *string `json:"notice_year,omitempty"`

	// Position Human-readable position description
	//
	// Example: 36°55.5'N 76°00.3'W
	Position *string `json:"position,omitempty"`

	// PostNote Notes appearing after the main entry
	PostNote *string `json:"post_note,omitempty"`

	// PrecedingNote Notes appearing before the main entry
	PrecedingNote *string `json:"preceding_note,omitempty"`

	// Range Signal range in nautical miles
	//
	// Example: 100
	Range *int `json:"range,omitempty"`

	// RegionHeading Regional geographic subdivision
	//
	// Example: EAST COAST
	RegionHeading *string `json:"region_heading,omitempty"`

	// Remarks Additional remarks about the station
	Remarks *string `json:"remarks,omitempty"`

	// RemoveFromList Flag indicating if entry should be removed
	RemoveFromList *string `json:"remove_from_list,omitempty"`

	// StationId Station identifier code
	//
	// Example: 852
	StationId *string `json:"station_id,omitempty"`

	// TransferRate Data transfer rate in bits per second
	//
	// Example: 200
	TransferRate *int `json:"transfer_rate,omitempty"`

	// VolumeNumber NGA publication volume number
	//
	// Example: PUB 117
	VolumeNumber *string `json:"volume_number,omitempty"`
}

// DGPSStationsWithinLocationResponse Response containing DGPS stations within location data
type DGPSStationsWithinLocationResponse struct {
	DgpsStations *[]DGPSStation `json:"dgpsStations,omitempty"`
	NextToken    *string        `json:"nextToken,omitempty"`
}

// ErrorCode defines model for ErrorCode.
type ErrorCode string

// ErrorType defines model for ErrorType.
type ErrorType string

// FindDGPSStationsResponse Response containing dgps station data
type FindDGPSStationsResponse struct {
	DgpsStations *[]DGPSStation `json:"dgpsStations,omitempty"`
	NextToken    *string        `json:"nextToken,omitempty"`
}

// FindLightAidsResponse Response containing light aid data
type FindLightAidsResponse struct {
	LightAids *[]LightAid `json:"lightAids,omitempty"`
	NextToken *string     `json:"nextToken,omitempty"`
}

// FindMODUsResponse Query parameters for finding modu by name
type FindMODUsResponse struct {
	Modus     *[]MODU `json:"modus,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
}

// FindPortsResponse Response containing port data
type FindPortsResponse struct {
	NextToken *string `json:"nextToken,omitempty"`
	Ports     *[]Port `json:"ports,omitempty"`
}

// FindRadioBeaconsResponse Response containing radio beacon data
type FindRadioBeaconsResponse struct {
	NextToken    *string        `json:"nextToken,omitempty"`
	RadioBeacons *[]RadioBeacon `json:"radioBeacons,omitempty"`
}

// FindVesselsResponse Query parameters for finding vessels by name
type FindVesselsResponse struct {
	Meta      *VesselSearchMeta `json:"_meta,omitempty"`
	NextToken *string           `json:"nextToken,omitempty"`
	Vessels   *[]Vessel         `json:"vessels,omitempty"`
}

// ForbiddenErrorDetail defines model for ForbiddenErrorDetail.
type ForbiddenErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: feature_not_available
	Code *ErrorCode `json:"code,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: this feature is not available on the "free" plan
	Message *string `json:"message,omitempty"`

	// Type Type categorizes the error (always "forbidden_error" for 403s)
	//
	// Example: forbidden_error
	Type *ErrorType `json:"type,omitempty"`
}

// ForbiddenErrorResponse defines model for ForbiddenErrorResponse.
type ForbiddenErrorResponse struct {
	Error *ForbiddenErrorDetail `json:"error,omitempty"`
}

// GeoJSON GeoJSON Point geometry for geospatial coordinates
type GeoJSON struct {
	// Coordinates Array of [longitude, latitude] in decimal degrees
	//
	// Example: [103.8215,1.2644]
	Coordinates *[]float64 `json:"coordinates,omitempty"`

	// Type GeoJSON geometry type, always "Point" for location data
	//
	// Example: Point
	Type *string `json:"type,omitempty"`
}

// InternalServerErrorDetail defines model for InternalServerErrorDetail.
type InternalServerErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: internal_error
	Code *ErrorCode `json:"code,omitempty"`

	// ErrorId ErrorID is a unique identifier for tracking this error (always present for 500 errors)
	//
	// Example: 550e8400-e29b-41d4-a716-446655440000
	ErrorId *string `json:"error_id,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: An internal error occurred. Please reference error ID 550e8400-e29b-41d4-a716-446655440000 when contacting support if this issue persists.
	Message *string `json:"message,omitempty"`

	// Timestamp Timestamp when the error occurred (ISO 8601 format)
	//
	// Example: 2025-10-31T15:30:45Z
	Timestamp *string `json:"timestamp,omitempty"`

	// Type Type categorizes the error (always "api_error" for 500s)
	//
	// Example: api_error
	Type *ErrorType `json:"type,omitempty"`
}

// InternalServerErrorResponse defines model for InternalServerErrorResponse.
type InternalServerErrorResponse struct {
	Error *InternalServerErrorDetail `json:"error,omitempty"`
}

// LightAid Navigational light aid including lighthouses, buoys, and beacons
type LightAid struct {
	// AidType Type of navigational aid (Light, Buoy, Beacon, etc.)
	//
	// Example: Light
	AidType *string `json:"aid_type,omitempty"`

	// Characteristic Light flash pattern description (e.g., Fl W 7.5s)
	//
	// Example: Fl W 7.5s
	Characteristic *string `json:"characteristic,omitempty"`

	// CharacteristicNumber Light characteristic code number
	CharacteristicNumber *int `json:"characteristic_number,omitempty"`

	// DeleteFlag Deletion status flag
	DeleteFlag *string `json:"delete_flag,omitempty"`

	// FeatureNumber NGA feature number identifier
	//
	// Example: 590
	FeatureNumber *string `json:"feature_number,omitempty"`

	// GeopoliticalHeading Country or major geographic area
	//
	// Example: UNITED STATES
	GeopoliticalHeading *string `json:"geopolitical_heading,omitempty"`

	// HeightFeetMeters Height of light above water in feet and meters
	//
	// Example: 192ft 59m
	HeightFeetMeters *string `json:"height_feet_meters,omitempty"`

	// LocalHeading Local area description
	//
	// Example: Cape Hatteras
	LocalHeading *string `json:"local_heading,omitempty"`

	// Location GeoJSON point for geospatial queries
	Location *GeoJSON `json:"location,omitempty"`

	// Name Name of the light aid
	//
	// Example: Cape Hatteras Light
	Name *string `json:"name,omitempty"`

	// NoticeNumber Notice to Mariners number
	NoticeNumber *int `json:"notice_number,omitempty"`

	// NoticeWeek Week of the notice
	NoticeWeek *string `json:"notice_week,omitempty"`

	// NoticeYear Year of the notice
	NoticeYear *string `json:"notice_year,omitempty"`

	// Position Human-readable position description
	//
	// Example: 35°15.1'N 75°31.6'W
	Position *string `json:"position,omitempty"`

	// PostNote Notes appearing after the main entry
	PostNote *string `json:"post_note,omitempty"`

	// PrecedingNote Notes appearing before the main entry
	PrecedingNote *string `json:"preceding_note,omitempty"`

	// Range Nominal range of light in nautical miles
	//
	// Example: 24
	Range *string `json:"range,omitempty"`

	// RegionHeading Regional geographic subdivision
	//
	// Example: EAST COAST
	RegionHeading *string `json:"region_heading,omitempty"`

	// Remarks Additional remarks about the light
	Remarks *string `json:"remarks,omitempty"`

	// RemoveFromList Flag indicating if entry should be removed
	RemoveFromList *string `json:"remove_from_list,omitempty"`

	// Structure Description of the physical structure
	//
	// Example: Black and white spiral bands
	Structure *string `json:"structure,omitempty"`

	// SubregionHeading Sub-regional geographic area
	//
	// Example: North Carolina
	SubregionHeading *string `json:"subregion_heading,omitempty"`

	// VolumeNumber NGA publication volume number
	//
	// Example: PUB 110
	VolumeNumber *string `json:"volume_number,omitempty"`
}

// LightAidsWithinLocationResponse Response containing light aids within location data
type LightAidsWithinLocationResponse struct {
	LightAids *[]LightAid `json:"lightAids,omitempty"`
	NextToken *string     `json:"nextToken,omitempty"`
}

// MODU Mobile Offshore Drilling Unit (MODU) location and status information
type MODU struct {
	// Date Date of the position report (parsed from API string format YYYY-MM-DD)
	Date *string `json:"date,omitempty"`

	// Distance Distance from reference point in nautical miles
	Distance *float64 `json:"distance,omitempty"`

	// Latitude Geographic latitude in decimal degrees
	//
	// Example: 28.7381
	Latitude *float64 `json:"latitude,omitempty"`

	// Location GeoJSON point for geospatial queries
	Location *GeoJSON `json:"location,omitempty"`

	// Longitude Geographic longitude in decimal degrees
	//
	// Example: -88.3659
	Longitude *float64 `json:"longitude,omitempty"`

	// Name Name of the drilling unit
	//
	// Example: DEEPWATER HORIZON
	Name *string `json:"name,omitempty"`

	// NavigationArea NAVAREA designation
	//
	// Example: IV
	NavigationArea *string `json:"navigation_area,omitempty"`

	// Position Human-readable position description
	//
	// Example: Gulf of Mexico
	Position *string `json:"position,omitempty"`

	// Region NGA region code
	//
	// Example: 8
	Region *int `json:"region,omitempty"`

	// RigStatus Current operational status of the rig
	//
	// Example: Drilling
	RigStatus *string `json:"rig_status,omitempty"`

	// SpecialStatus Any special status or notes
	SpecialStatus *string `json:"special_status,omitempty"`

	// SubRegion NGA sub-region code
	//
	// Example: 81
	SubRegion *int `json:"sub_region,omitempty"`
}

// MODUsWithinLocationResponse Response containing MODUs within location data
type MODUsWithinLocationResponse struct {
	Modus     *[]MODU `json:"modus,omitempty"`
	NextToken *string `json:"nextToken,omitempty"`
}

// MarineCasualtiesResponse Response containing marine casualty data
type MarineCasualtiesResponse struct {
	Meta       *TypesResolutionMeta `json:"_meta,omitempty"`
	Casualties *[]MarineCasualty    `json:"casualties,omitempty"`
	NextToken  *string              `json:"nextToken,omitempty"`
}

// MarineCasualty defines model for MarineCasualty.
type MarineCasualty struct {
	AtCoding         *[]string `json:"atCoding,omitempty"`
	CasualtyReportNr *string   `json:"casualtyReportNr,omitempty"`
	CfCoding         *[]string `json:"cfCoding,omitempty"`

	// CollectedAt Record metadata (added when the record was stored)
	CollectedAt        *string   `json:"collectedAt,omitempty"`
	CompetentAuthority *[]string `json:"competentAuthority,omitempty"`

	// DateOfOccurrence Event details
	DateOfOccurrence      *string   `json:"dateOfOccurrence,omitempty"`
	Deviation             *[]string `json:"deviation,omitempty"`
	EventType             *[]string `json:"eventType,omitempty"`
	FinishedInvestigation *bool     `json:"finishedInvestigation,omitempty"`
	ImoNr                 *[]string `json:"imoNr,omitempty"`
	InterimReport         *bool     `json:"interimReport,omitempty"`

	// InvestigatingState Investigation
	InvestigatingState *string `json:"investigatingState,omitempty"`

	// LivesLostTotal Consequences
	LivesLostTotal *string `json:"livesLostTotal,omitempty"`

	// NameOfShip Vessel info (arrays since multiple vessels can be involved)
	NameOfShip         *[]string `json:"nameOfShip,omitempty"`
	OccurrenceSeverity *string   `json:"occurrenceSeverity,omitempty"`

	// OccurrenceUuid External IDs
	OccurrenceUuid        *string   `json:"occurrenceUuid,omitempty"`
	OccurrenceWithPersons *[]string `json:"occurrenceWithPersons,omitempty"`

	// OccurrenceWithShips Taxonomy classifications
	OccurrenceWithShips *[]string `json:"occurrenceWithShips,omitempty"`
	PeopleInjuredTotal  *string   `json:"peopleInjuredTotal,omitempty"`
	Pollution           *bool     `json:"pollution,omitempty"`
	ShipCraftType       *[]string `json:"shipCraftType,omitempty"`
	SrCoding            *[]string `json:"srCoding,omitempty"`
}

// NotFoundErrorDetail defines model for NotFoundErrorDetail.
type NotFoundErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: resource_missing
	Code *ErrorCode `json:"code,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: Vessel not found: 123456789
	Message *string `json:"message,omitempty"`

	// Type Type categorizes the error (always "not_found_error" for 404s)
	//
	// Example: not_found_error
	Type *ErrorType `json:"type,omitempty"`
}

// NotFoundErrorResponse defines model for NotFoundErrorResponse.
type NotFoundErrorResponse struct {
	Error *NotFoundErrorDetail `json:"error,omitempty"`
}

// PaymentRequiredErrorDetail defines model for PaymentRequiredErrorDetail.
type PaymentRequiredErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: insufficient_credits
	Code *ErrorCode `json:"code,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: insufficient satellite credits
	Message *string `json:"message,omitempty"`

	// Type Type categorizes the error (always "payment_required_error" for 402s)
	//
	// Example: payment_required_error
	Type *ErrorType `json:"type,omitempty"`
}

// PaymentRequiredErrorResponse defines model for PaymentRequiredErrorResponse.
type PaymentRequiredErrorResponse struct {
	Error *PaymentRequiredErrorDetail `json:"error,omitempty"`
}

// Port Port information including facilities, services, and characteristics where available
type Port struct {
	// AnchorageDepth Depth at the anchorage area
	//
	// Example: 20
	AnchorageDepth *float64 `json:"anchorage_depth,omitempty"`

	// AnchorageDepthUnit Unit for anchorage depth measurement
	//
	// Example: m
	AnchorageDepthUnit *string `json:"anchorage_depth_unit,omitempty"`

	// CargoHandlingDepth Depth at cargo handling berths
	//
	// Example: 18
	CargoHandlingDepth *float64 `json:"cargo_handling_depth,omitempty"`

	// CargoHandlingDepthUnit Unit for cargo handling depth measurement
	//
	// Example: m
	CargoHandlingDepthUnit *string `json:"cargo_handling_depth_unit,omitempty"`

	// ChannelDepth Depth of the approach channel
	//
	// Example: 23
	ChannelDepth *float64 `json:"channel_depth,omitempty"`

	// ChannelDepthUnit Unit for channel depth measurement
	//
	// Example: m
	ChannelDepthUnit *string `json:"channel_depth_unit,omitempty"`

	// Country Country information for the port
	Country *ContractsPortCountry `json:"country,omitempty"`

	// GarbageDisposal Whether garbage disposal services are available
	//
	// Example: true
	GarbageDisposal *bool `json:"garbage_disposal,omitempty"`

	// HarborSize Harbor size classification (Large/Medium/Small/Very Small)
	//
	// Example: Large
	HarborSize *string `json:"harbor_size,omitempty"`

	// HarborType Type of harbor (CB=Coastal Breakwater, CN=Coastal Natural, etc.)
	//
	// Example: CN
	HarborType *string `json:"harbor_type,omitempty"`

	// HarborUse Primary use of the harbor (FISH/MIL/CARGO/FERRY/UNK)
	//
	// Example: CARGO
	HarborUse *string `json:"harbor_use,omitempty"`

	// HasDrydock Whether the port has drydock facilities
	//
	// Example: true
	HasDrydock *bool `json:"has_drydock,omitempty"`

	// Latitude Geographic latitude in decimal degrees
	//
	// Example: 1.2644
	Latitude *float64 `json:"latitude,omitempty"`

	// Location GeoJSON point for geospatial queries
	Location *GeoJSON `json:"location,omitempty"`

	// Longitude Geographic longitude in decimal degrees
	//
	// Example: 103.8215
	Longitude *float64 `json:"longitude,omitempty"`

	// MaxVesselBeam Maximum beam (width) of vessel that can be accommodated
	//
	// Example: 60
	MaxVesselBeam *float64 `json:"max_vessel_beam,omitempty"`

	// MaxVesselBeamUnit Unit for maximum vessel beam
	//
	// Example: m
	MaxVesselBeamUnit *string `json:"max_vessel_beam_unit,omitempty"`

	// MaxVesselDraft Maximum draft of vessel that can be accommodated
	//
	// Example: 16
	MaxVesselDraft *float64 `json:"max_vessel_draft,omitempty"`

	// MaxVesselDraftUnit Unit for maximum vessel draft
	//
	// Example: m
	MaxVesselDraftUnit *string `json:"max_vessel_draft_unit,omitempty"`

	// MaxVesselLength Maximum length of vessel that can be accommodated
	//
	// Example: 400
	MaxVesselLength *float64 `json:"max_vessel_length,omitempty"`

	// MaxVesselLengthUnit Unit for maximum vessel length
	//
	// Example: m
	MaxVesselLengthUnit *string `json:"max_vessel_length_unit,omitempty"`

	// MedicalFacilities Whether medical facilities are available at the port
	//
	// Example: true
	MedicalFacilities *bool `json:"medical_facilities,omitempty"`

	// Name The port's official name
	//
	// Example: Singapore
	Name *string `json:"name,omitempty"`

	// NavigationArea NAVAREA designation for maritime safety communications
	//
	// Example: XI
	NavigationArea *string `json:"navigation_area,omitempty"`

	// PilotageAvailable Whether pilotage services are available
	//
	// Example: true
	PilotageAvailable *bool `json:"pilotage_available,omitempty"`

	// PilotageCompulsory Whether pilotage is mandatory for vessel entry
	//
	// Example: true
	PilotageCompulsory *bool `json:"pilotage_compulsory,omitempty"`

	// PortSecurity Whether ISPS (International Ship and Port Facility Security) compliant
	//
	// Example: true
	PortSecurity *bool `json:"port_security,omitempty"`

	// RegionName Geographic region where the port is located
	//
	// Example: Southeast Asia
	RegionName *string `json:"region_name,omitempty"`

	// RepairCapability Level of ship repair capability (Major/Moderate/Limited/Emergency/None)
	//
	// Example: Major
	RepairCapability *string `json:"repair_capability,omitempty"`

	// Shelter Quality of shelter from weather (Excellent/Good/Fair/Poor/None)
	//
	// Example: Excellent
	Shelter *string `json:"shelter,omitempty"`

	// Size Size classification of the port
	//
	// Example: Large
	Size *string `json:"size,omitempty"`

	// SupplyDiesel Whether diesel supply is available
	//
	// Example: true
	SupplyDiesel *bool `json:"supply_diesel,omitempty"`

	// SupplyFuel Whether fuel oil supply is available
	//
	// Example: true
	SupplyFuel *bool `json:"supply_fuel,omitempty"`

	// SupplyWater Whether fresh water supply is available
	//
	// Example: true
	SupplyWater *bool `json:"supply_water,omitempty"`

	// TrafficSeparationScheme Whether a TSS (Traffic Separation Scheme) is in place
	//
	// Example: true
	TrafficSeparationScheme *bool `json:"traffic_separation_scheme,omitempty"`

	// TugsAvailable Whether tug services are available
	//
	// Example: true
	TugsAvailable *bool `json:"tugs_available,omitempty"`

	// Type Port classification type
	//
	// Example: Seaport
	Type *string `json:"type,omitempty"`

	// UnloCode UN Location Code (LOCODE) - unique port identifier
	//
	// Example: SGSIN
	UnloCode *string `json:"unlo_code,omitempty"`

	// VesselTrafficService Whether VTS (Vessel Traffic Service) is operational
	//
	// Example: true
	VesselTrafficService *bool `json:"vessel_traffic_service,omitempty"`
}

// PortEvent Vessel port call event including arrivals and departures
type PortEvent struct {
	// Event Type of port event - either "Arrival" or "Departure"
	//
	// Example: Arrival
	Event *string `json:"event,omitempty"`

	// Port Reference to the port where the event occurred
	Port *PortReference `json:"port,omitempty"`

	// Timestamp UTC timestamp when the event occurred
	Timestamp *string `json:"timestamp,omitempty"`

	// Vessel Reference to the vessel involved in the port event
	Vessel *VesselReference `json:"vessel,omitempty"`
}

// PortEventResponse Response containing a single port event
type PortEventResponse struct {
	Meta *TypesResolutionMeta `json:"_meta,omitempty"`

	// PortEvent Vessel port call event including arrivals and departures
	PortEvent *PortEvent `json:"portEvent,omitempty"`
}

// PortEventsResponse defines model for PortEventsResponse.
type PortEventsResponse struct {
	Meta       *TypesResolutionMeta `json:"_meta,omitempty"`
	NextToken  *string              `json:"nextToken,omitempty"`
	PortEvents *[]PortEvent         `json:"portEvents,omitempty"`
}

// PortInboundResponse Response containing vessels heading to a port
type PortInboundResponse struct {
	NextToken  *string      `json:"nextToken,omitempty"`
	VesselETAs *[]VesselETA `json:"vesselETAs,omitempty"`
}

// PortReference Port identification details for port event reference
type PortReference struct {
	// Country Country where the port is located
	//
	// Example: Singapore
	Country *string `json:"country,omitempty"`

	// Name The port's official name
	//
	// Example: Singapore
	Name *string `json:"name,omitempty"`

	// UnloCode UN Location Code (LOCODE) for the port
	//
	// Example: SGSIN
	UnloCode *string `json:"unlo_code,omitempty"`
}

// PortResponse Response containing a single port
type PortResponse struct {
	// Port Port information including facilities, services, and characteristics where available
	Port *Port `json:"port,omitempty"`
}

// PortsWithinLocationResponse Response containing ports within location data
type PortsWithinLocationResponse struct {
	NextToken *string `json:"nextToken,omitempty"`
	Ports     *[]Port `json:"ports,omitempty"`
}

// PositionCheck defines model for PositionCheck.
type PositionCheck struct {
	// DistanceNm Great-circle distance from the previous position, in nautical miles
	//
	// Example: 96.3
	DistanceNm *float64 `json:"distance_nm,omitempty"`

	// ExceedsTypicalSpeed True when the implied speed is more than 1.2 times the vessel's usual top speed (99th percentile of its recent speeds)
	//
	// Example: false
	ExceedsTypicalSpeed *bool `json:"exceeds_typical_speed,omitempty"`

	// IdentityMatch True when the returned position's MMSI and IMO match the identifiers in our vessel record. Omitted when we hold no record or nothing can be compared.
	//
	// Example: true
	IdentityMatch *bool `json:"identity_match,omitempty"`

	// ImpliedCourse Initial great-circle bearing from the previous position, in degrees from true north
	//
	// Example: 87
	ImpliedCourse *int `json:"implied_course,omitempty"`

	// ImpliedSpeedKn Average speed needed to cover that distance in that time, in knots
	//
	// Example: 14
	ImpliedSpeedKn *float64 `json:"implied_speed_kn,omitempty"`

	// LookbackHours How far back the previous position was searched for
	//
	// Example: 80
	LookbackHours *int `json:"lookback_hours,omitempty"`

	// MinutesSincePrevious Minutes between the previous position and the returned one
	//
	// Example: 412
	MinutesSincePrevious *int `json:"minutes_since_previous,omitempty"`

	// PreviousPosition The vessel's newest earlier position report within lookback_hours that is not a suspected glitch
	PreviousPosition *PreviousPosition `json:"previous_position,omitempty"`
}

// PreviousPosition defines model for PreviousPosition.
type PreviousPosition struct {
	// Latitude Latitude of the previous position in decimal degrees
	//
	// Example: 1.2644
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude Longitude of the previous position in decimal degrees
	//
	// Example: 103.8215
	Longitude *float64 `json:"longitude,omitempty"`

	// Timestamp When the previous position was reported (UTC)
	//
	// Example: 2026-10-24T03:10:00Z
	Timestamp *string `json:"timestamp,omitempty"`
}

// RadioBeacon Navigational radio beacon for maritime direction finding
type RadioBeacon struct {
	// AidType Type of radio aid
	//
	// Example: RBn
	AidType *string `json:"aid_type,omitempty"`

	// Characteristic Beacon transmission characteristic
	//
	// Example: A (.--)
	Characteristic *string `json:"characteristic,omitempty"`

	// DeleteFlag Deletion status flag
	DeleteFlag *string `json:"delete_flag,omitempty"`

	// FeatureNumber NGA feature number identifier
	//
	// Example: 456
	FeatureNumber *float64 `json:"feature_number,omitempty"`

	// Frequency Broadcast frequency
	//
	// Example: 286 kHz
	Frequency *string `json:"frequency,omitempty"`

	// GeopoliticalHeading Country or major geographic area
	//
	// Example: UNITED STATES
	GeopoliticalHeading *string `json:"geopolitical_heading,omitempty"`

	// Location GeoJSON point for geospatial queries
	Location *GeoJSON `json:"location,omitempty"`

	// Name Name of the radio beacon
	//
	// Example: Ambrose Light
	Name *string `json:"name,omitempty"`

	// NoticeNumber Notice to Mariners number
	NoticeNumber *int `json:"notice_number,omitempty"`

	// NoticeWeek Week of the notice
	NoticeWeek *string `json:"notice_week,omitempty"`

	// NoticeYear Year of the notice
	NoticeYear *string `json:"notice_year,omitempty"`

	// Position Human-readable position description
	//
	// Example: 40°27.1'N 73°49.5'W
	Position *string `json:"position,omitempty"`

	// PostNote Notes appearing after the main entry
	PostNote *string `json:"post_note,omitempty"`

	// PrecedingNote Notes appearing before the main entry
	PrecedingNote *string `json:"preceding_note,omitempty"`

	// Range Signal range in nautical miles
	//
	// Example: 20
	Range *string `json:"range,omitempty"`

	// RegionHeading Regional geographic subdivision
	//
	// Example: EAST COAST
	RegionHeading *string `json:"region_heading,omitempty"`

	// RemoveFromList Flag indicating if entry should be removed
	RemoveFromList *string `json:"remove_from_list,omitempty"`

	// SequenceText Transmission sequence description
	SequenceText *string `json:"sequence_text,omitempty"`

	// StationRemark Remarks specific to the station
	StationRemark *string `json:"station_remark,omitempty"`

	// VolumeNumber NGA publication volume number
	//
	// Example: PUB 117
	VolumeNumber *string `json:"volume_number,omitempty"`
}

// RadioBeaconsWithinLocationResponse Response containing radio beacons within location data
type RadioBeaconsWithinLocationResponse struct {
	NextToken    *string        `json:"nextToken,omitempty"`
	RadioBeacons *[]RadioBeacon `json:"radioBeacons,omitempty"`
}

// RateLimitErrorDetail defines model for RateLimitErrorDetail.
type RateLimitErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: rate_limit_exceeded
	Code *ErrorCode `json:"code,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: API monthly quota exceeded
	Message *string `json:"message,omitempty"`

	// Type Type categorizes the error (always "rate_limit_error" for 429s)
	//
	// Example: rate_limit_error
	Type *ErrorType `json:"type,omitempty"`
}

// RateLimitErrorResponse defines model for RateLimitErrorResponse.
type RateLimitErrorResponse struct {
	Error *RateLimitErrorDetail `json:"error,omitempty"`
}

// ServiceUnavailableErrorDetail defines model for ServiceUnavailableErrorDetail.
type ServiceUnavailableErrorDetail struct {
	// Code Code is a short string identifier for this error for programmatic handling
	//
	// Example: service_unavailable
	Code *ErrorCode `json:"code,omitempty"`

	// Message Message is a human-readable message providing more details about the error
	//
	// Example: The service is temporarily unavailable. Retry after a brief delay.
	Message *string `json:"message,omitempty"`

	// Type Type categorizes the error (always "service_unavailable_error" for 503s)
	//
	// Example: service_unavailable_error
	Type *ErrorType `json:"type,omitempty"`
}

// ServiceUnavailableErrorResponse defines model for ServiceUnavailableErrorResponse.
type ServiceUnavailableErrorResponse struct {
	Error *ServiceUnavailableErrorDetail `json:"error,omitempty"`
}

// Vessel Vessel static data including identification, dimensions, and capacities where available
type Vessel struct {
	// Breadth Maximum beam (width) of the vessel
	//
	// Example: 59
	Breadth *int `json:"breadth,omitempty"`

	// BreadthUnit Unit of measurement for breadth (typically meters)
	//
	// Example: m
	BreadthUnit *string `json:"breadth_unit,omitempty"`

	// CallSign International radio call sign assigned to the vessel
	//
	// Example: H3RC
	CallSign *string `json:"call_sign,omitempty"`

	// Country Country of registration (flag state)
	//
	// Example: Panama
	Country *string `json:"country,omitempty"`

	// CountryCode ISO 2-letter country code of flag state
	//
	// Example: PA
	CountryCode *string `json:"country_code,omitempty"`

	// DeadweightTonnage Deadweight tonnage (DWT) - maximum cargo capacity in metric tons
	//
	// Example: 199629
	DeadweightTonnage *int `json:"deadweight_tonnage,omitempty"`

	// Draft Maximum draft (depth below waterline) of the vessel
	//
	// Example: 16
	Draft *int `json:"draft,omitempty"`

	// DraftUnit Unit of measurement for draft (typically meters)
	//
	// Example: m
	DraftUnit *string `json:"draft_unit,omitempty"`

	// DraughtCalculatedAvg Arithmetic mean of reported draught in metres over the last 31 days of ETA messages (zero readings excluded).
	//
	// Example: 13.8
	DraughtCalculatedAvg *float64 `json:"draught_calculated_avg,omitempty"`

	// DraughtObservedMax Largest reported draught in metres over the last 31 days of ETA messages. Distinct from summer_draught (the design value).
	//
	// Example: 16
	DraughtObservedMax *float64 `json:"draught_observed_max,omitempty"`

	// EngineModelName Model name of the main engine
	EngineModelName *string `json:"engine_model_name,omitempty"`

	// EngineType Type code for the main engine
	EngineType *int `json:"engine_type,omitempty"`

	// Eni European Number of Identification for inland waterway vessels (8 digits)
	Eni *string `json:"eni,omitempty"`

	// FormerNames List of previous names the vessel has operated under
	FormerNames *[]ContractsVesselFormerName `json:"former_names,omitempty"`

	// GrossTonnage Gross tonnage (GT) - measure of vessel's overall internal volume
	//
	// Example: 220940
	GrossTonnage *int `json:"gross_tonnage,omitempty"`

	// HomePort Port of registry for the vessel
	//
	// Example: Panama City
	HomePort *string `json:"home_port,omitempty"`

	// Imo International Maritime Organization number - permanent 7-digit vessel identifier
	//
	// Example: 9811000
	Imo *int `json:"imo,omitempty"`

	// KilowattPower Main engine power in kilowatts
	//
	// Example: 58000
	KilowattPower *int `json:"kilowatt_power,omitempty"`

	// Length Overall length of the vessel
	//
	// Example: 400
	Length *int `json:"length,omitempty"`

	// LengthUnit Unit of measurement for length (typically meters)
	//
	// Example: m
	LengthUnit *string `json:"length_unit,omitempty"`

	// Mmsi Maritime Mobile Service Identity - 9-digit radio identifier
	//
	// Example: 353136000
	Mmsi *int `json:"mmsi,omitempty"`

	// Name The vessel's current registered name
	//
	// Example: EVER GIVEN
	Name *string `json:"name,omitempty"`

	// NameAis The name as currently broadcast on AIS. It may differ from the registered name for a short time after a rename, until our next registry update.
	//
	// Example: EVER GIVEN
	NameAis *string `json:"name_ais,omitempty"`

	// OperatingStatus Current operational status (e.g., Active, Laid Up, Scrapped)
	//
	// Example: Active
	OperatingStatus *string `json:"operating_status,omitempty"`

	// SpeedCalculatedAvg Arithmetic mean of speed over ground in knots over a recent 31-day window of position reports, computed from this vessel's AIS history (stops excluded). It is derived from our observations and is distinct from the design service speed.
	//
	// Example: 12.4
	SpeedCalculatedAvg *float64 `json:"speed_calculated_avg,omitempty"`

	// SpeedObservedMax Usual top speed: the 99th percentile of speed over ground in knots over a recent 31-day window, capped at 40 knots to discard GPS glitches.
	//
	// Example: 22.1
	SpeedObservedMax *float64 `json:"speed_observed_max,omitempty"`

	// SummerDraught Maximum design draught (summer-load line). Static vessel attribute, distinct from draught_observed_max.
	//
	// Example: 16.5
	SummerDraught *float64 `json:"summer_draught,omitempty"`

	// Teu Twenty-foot equivalent unit container capacity. Populated only for container ships.
	//
	// Example: 23992
	Teu *int `json:"teu,omitempty"`

	// VesselSubtype Finer-grained classification beyond vessel_type (e.g., inland motor cargo or tanker vessel).
	VesselSubtype *string `json:"vessel_subtype,omitempty"`

	// VesselType Classification of vessel type (e.g., Container Ship, Bulk Carrier)
	//
	// Example: Cargo A
	VesselType *string `json:"vessel_type,omitempty"`

	// YearBuilt Year the vessel was constructed
	//
	// Example: 2018
	YearBuilt *int `json:"year_built,omitempty"`
}

// VesselETA Vessel Estimated Time of Arrival information from AIS static/voyage data
type VesselETA struct {
	// Destination Reported destination port or area as entered by the vessel
	//
	// Example: SINGAPORE
	Destination *string `json:"destination,omitempty"`

	// DestinationPort Resolved UN/LOCODE of the destination port (empty if unresolved)
	//
	// Example: NLRTM
	DestinationPort *string `json:"destination_port,omitempty"`

	// Draught Current draught (draft) of the vessel in meters
	//
	// Example: 14.5
	Draught *float64 `json:"draught,omitempty"`

	// Eta Estimated time of arrival at destination as reported by the vessel
	Eta *string `json:"eta,omitempty"`

	// Imo International Maritime Organization number - permanent vessel identifier
	//
	// Example: 9811000
	Imo *int `json:"imo,omitempty"`

	// Mmsi Maritime Mobile Service Identity - unique 9-digit vessel identifier
	//
	// Example: 353136000
	Mmsi *int `json:"mmsi,omitempty"`

	// Timestamp UTC timestamp when this ETA information was received
	Timestamp *string `json:"timestamp,omitempty"`

	// VesselName The vessel's registered name as reported in AIS
	//
	// Example: EVER GIVEN
	VesselName *string `json:"vessel_name,omitempty"`
}

// VesselETAResponse Response containing vessel ETA data
type VesselETAResponse struct {
	Meta *TypesResolutionMeta `json:"_meta,omitempty"`

	// VesselEta Vessel Estimated Time of Arrival information from AIS static/voyage data
	VesselEta *VesselETA `json:"vesselEta,omitempty"`
}

// VesselEmission defines model for VesselEmission.
type VesselEmission struct {
	Co2EmissionsAtBerth        *float64 `json:"co2_emissions_at_berth,omitempty"`
	Co2EmissionsOnLadenVoyages *float64 `json:"co2_emissions_on_laden_voyages,omitempty"`

	// Co2EmissionsTotal CO2 emissions
	Co2EmissionsTotal *float64 `json:"co2_emissions_total,omitempty"`

	// Co2PerDistance kg CO₂ / n mile
	Co2PerDistance *float64 `json:"co2_per_distance,omitempty"`

	// Co2PerTransportWork g CO₂ / m tonnes · n miles
	Co2PerTransportWork *float64 `json:"co2_per_transport_work,omitempty"`
	CollectedAt         *string  `json:"collected_at,omitempty"`

	// DistanceThroughIce nautical miles
	DistanceThroughIce *float64 `json:"distance_through_ice,omitempty"`
	DocExpiryDate      *string  `json:"doc_expiry_date,omitempty"`

	// DocIssueDate DOC info
	DocIssueDate         *string  `json:"doc_issue_date,omitempty"`
	FlagCode             *string  `json:"flag_code,omitempty"`
	FlagName             *string  `json:"flag_name,omitempty"`
	FuelConsumptionHfo   *float64 `json:"fuel_consumption_hfo,omitempty"`
	FuelConsumptionLfo   *float64 `json:"fuel_consumption_lfo,omitempty"`
	FuelConsumptionLng   *float64 `json:"fuel_consumption_lng,omitempty"`
	FuelConsumptionMdo   *float64 `json:"fuel_consumption_mdo,omitempty"`
	FuelConsumptionMgo   *float64 `json:"fuel_consumption_mgo,omitempty"`
	FuelConsumptionOther *float64 `json:"fuel_consumption_other,omitempty"`

	// FuelConsumptionTotal Fuel consumption (tonnes)
	FuelConsumptionTotal *float64 `json:"fuel_consumption_total,omitempty"`

	// FuelPerDistance Efficiency metrics (pre-calculated)
	FuelPerDistance *float64 `json:"fuel_per_distance,omitempty"`

	// FuelPerTransportWork g / m tonnes · n miles
	FuelPerTransportWork *float64 `json:"fuel_per_transport_work,omitempty"`
	HomePort             *string  `json:"home_port,omitempty"`
	IceClass             *string  `json:"ice_class,omitempty"`

	// Imo Identifiers
	Imo *int `json:"imo,omitempty"`

	// MonitoringMethodA Monitoring method
	MonitoringMethodA  *string `json:"monitoring_method_a,omitempty"`
	MonitoringMethodB  *string `json:"monitoring_method_b,omitempty"`
	MonitoringMethodC  *string `json:"monitoring_method_c,omitempty"`
	MonitoringMethodD  *string `json:"monitoring_method_d,omitempty"`
	Name               *string `json:"name,omitempty"`
	PortCallsOutsideEu *int    `json:"port_calls_outside_eu,omitempty"`

	// PortCallsWithinEu EU specific
	PortCallsWithinEu *int `json:"port_calls_within_eu,omitempty"`

	// ReportingPeriod Reporting period
	ReportingPeriod *string `json:"reporting_period,omitempty"`

	// TechnicalEfficiency Technical efficiency
	TechnicalEfficiency      *string  `json:"technical_efficiency,omitempty"`
	TechnicalEfficiencyValue *float64 `json:"technical_efficiency_value,omitempty"`

	// TimeAtSeaThroughIce hours
	TimeAtSeaThroughIce *float64 `json:"time_at_sea_through_ice,omitempty"`

	// TotalTimeAtSea Distance and time
	TotalTimeAtSea *float64 `json:"total_time_at_sea,omitempty"`

	// UniqueKey Stable identifier for this emission record (IMO + reporting period)
	UniqueKey             *string `json:"unique_key,omitempty"`
	VerifierAccreditation *string `json:"verifier_accreditation,omitempty"`
	VerifierAddress       *string `json:"verifier_address,omitempty"`

	// VerifierName Verifier info
	VerifierName *string `json:"verifier_name,omitempty"`

	// VesselType Vessel identification
	VesselType *string `json:"vessel_type,omitempty"`
}

// VesselEmissionsResponse Response containing vessel emissions data
type VesselEmissionsResponse struct {
	Meta      *TypesResolutionMeta `json:"_meta,omitempty"`
	Emissions *[]VesselEmission    `json:"emissions,omitempty"`
	NextToken *string              `json:"nextToken,omitempty"`
}

// VesselPosition Vessel position report derived from AIS messages
type VesselPosition struct {
	// Cog Course Over Ground in degrees (0-359.9), null when unavailable
	//
	// Example: 231.5
	Cog *float64 `json:"cog,omitempty"`

	// Heading True heading in degrees (0-359), null when unavailable
	//
	// Example: 230
	Heading *int `json:"heading,omitempty"`

	// Imo International Maritime Organization number - permanent vessel identifier
	//
	// Example: 9811000
	Imo *int `json:"imo,omitempty"`

	// Latitude Geographic latitude in decimal degrees (-90 to 90)
	//
	// Example: 1.2644
	Latitude *float64 `json:"latitude,omitempty"`

	// Location GeoJSON point for geospatial queries
	Location *GeoJSON `json:"location,omitempty"`

	// Longitude Geographic longitude in decimal degrees (-180 to 180)
	//
	// Example: 103.8215
	Longitude *float64 `json:"longitude,omitempty"`

	// Mmsi Maritime Mobile Service Identity - unique 9-digit vessel identifier
	//
	// Example: 353136000
	Mmsi *int `json:"mmsi,omitempty"`

	// NavStatus AIS navigational status (0=under way using engine, 1=at anchor, 2=not under command, 3=restricted manoeuvrability, 5=moored, 8=under way sailing, etc.), null when unavailable
	//
	// Example: 0
	NavStatus *int `json:"nav_status,omitempty"`

	// ProcessedTimestamp UTC timestamp when the position was processed by the system
	ProcessedTimestamp *string `json:"processed_timestamp,omitempty"`

	// Sog Speed Over Ground in knots (0-102.2), null when unavailable
	//
	// Example: 14.1
	Sog *float64 `json:"sog,omitempty"`

	// SuspectedGlitch True when the position implies a speed over 500 knots from the vessel's previous clean report, which no vessel reaches. Such a fix is likely a GPS glitch or a corrupted AIS transmission.
	//
	// Example: false
	SuspectedGlitch *bool `json:"suspected_glitch,omitempty"`

	// Timestamp UTC timestamp when the AIS message was transmitted by the vessel
	Timestamp *string `json:"timestamp,omitempty"`

	// VesselName The vessel's registered name as reported in AIS
	//
	// Example: EVER GIVEN
	VesselName *string `json:"vessel_name,omitempty"`
}

// VesselPositionResponse Response containing a single vessel position
type VesselPositionResponse struct {
	Meta *TypesResolutionMeta `json:"_meta,omitempty"`

	// PositionCheck How the returned position compares with the vessel's previous position report. The lookback_hours field is always present. The identity_match field is present when there is a record to compare against. The other fields need a previous report.
	PositionCheck *PositionCheck `json:"positionCheck,omitempty"`

	// Vessel The vessel's particulars, as returned by GET /vessel/{id}. Omitted when the vessel is not in the registry.
	Vessel *Vessel `json:"vessel,omitempty"`

	// VesselEta The latest destination and ETA the vessel reported, as returned by GET /vessel/{id}/eta. Omitted when no destination or ETA is available, including when the lookup does not finish in time.
	VesselEta *VesselETA `json:"vesselEta,omitempty"`

	// VesselPosition Vessel position report derived from AIS messages
	VesselPosition *VesselPosition `json:"vesselPosition,omitempty"`
}

// VesselPositionsResponse Response containing vessel position data
type VesselPositionsResponse struct {
	NextToken       *string           `json:"nextToken,omitempty"`
	VesselPositions *[]VesselPosition `json:"vesselPositions,omitempty"`
}

// VesselReference Vessel identification details for port event reference
type VesselReference struct {
	// Imo International Maritime Organization number
	//
	// Example: 9811000
	Imo *int `json:"imo,omitempty"`

	// Mmsi Maritime Mobile Service Identity number
	//
	// Example: 353136000
	Mmsi *int `json:"mmsi,omitempty"`

	// Name The vessel's registered name
	//
	// Example: EVER GIVEN
	Name *string `json:"name,omitempty"`
}

// VesselResponse Response containing vessel data
type VesselResponse struct {
	Meta *TypesResolutionMeta `json:"_meta,omitempty"`

	// Vessel Vessel static data including identification, dimensions, and capacities where available
	Vessel *Vessel `json:"vessel,omitempty"`
}

// VesselSearchMeta defines model for VesselSearchMeta.
type VesselSearchMeta struct {
	// MatchedOn MatchedOn maps the index of a vessel in the vessels array to the fields
	// that matched it, e.g. {"0":["eni"],"1":["imo"]}. Values are arrays because
	// one vessel can match on several fields at once (a vessel whose ENI and IMO
	// are the same digits). Field order within an entry is stable: eni, imo,
	// mmsi, callsign, name.
	MatchedOn *map[string][]string `json:"matchedOn,omitempty"`

	// Query Query echoes the q value the fields were matched against.
	//
	// Example: 4606770
	Query *string `json:"query,omitempty"`
}

// VesselsWithinLocationResponse Response containing vessel within location data
type VesselsWithinLocationResponse struct {
	NextToken *string           `json:"nextToken,omitempty"`
	Vessels   *[]VesselPosition `json:"vessels,omitempty"`
}

// ContractsPortCountry Country information for port location
type ContractsPortCountry struct {
	// Code ISO 2-letter country code
	//
	// Example: SG
	Code *string `json:"code,omitempty"`

	// Name Full country name
	//
	// Example: Singapore
	Name *string `json:"name,omitempty"`
}

// ContractsVesselFormerName Historical vessel name record
type ContractsVesselFormerName struct {
	// Name Previous name of the vessel
	//
	// Example: EVER GREEN
	Name *string `json:"name,omitempty"`

	// YearUntil Year until which this name was used
	//
	// Example: 2018
	YearUntil *string `json:"year_until,omitempty"`
}

// TypesResolutionMeta defines model for types.ResolutionMeta.
type TypesResolutionMeta struct {
	RequestedIdType *string `json:"requestedIdType,omitempty"`
	ResolvedId      *int    `json:"resolvedId,omitempty"`
	ResolvedIdType  *string `json:"resolvedIdType,omitempty"`

	// SuggestedIdType When set, the caller may get a result by retrying with the opposite
	// identifier type. It is set only on a total miss where the counterpart id
	// was neither supplied nor derivable. It is phrased conditionally and does
	// not assert that the other id exists.
	SuggestedIdType *string `json:"suggestedIdType,omitempty"`
}

// GetEmissionsParams defines parameters for GetEmissions.
type GetEmissionsParams struct {
	// FilterPeriod Reporting year filter (e.g. 2024)
	FilterPeriod *int `form:"filter.period,omitempty" json:"filter.period,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Pagination token for retrieving the next page of results
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationDgpsBoundingBoxParams defines parameters for GetLocationDgpsBoundingBox.
type GetLocationDgpsBoundingBoxParams struct {
	// FilterLonLeft Longitude of the left (western) edge of the bounding box
	FilterLonLeft float64 `form:"filter.lonLeft" json:"filter.lonLeft"`

	// FilterLonRight Longitude of the right (eastern) edge of the bounding box
	FilterLonRight float64 `form:"filter.lonRight" json:"filter.lonRight"`

	// FilterLatBottom Latitude of the bottom (southern) edge of the bounding box
	FilterLatBottom float64 `form:"filter.latBottom" json:"filter.latBottom"`

	// FilterLatTop Latitude of the top (northern) edge of the bounding box
	FilterLatTop float64 `form:"filter.latTop" json:"filter.latTop"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationDgpsRadiusParams defines parameters for GetLocationDgpsRadius.
type GetLocationDgpsRadiusParams struct {
	// FilterLongitude Longitude of the center point
	FilterLongitude float64 `form:"filter.longitude" json:"filter.longitude"`

	// FilterLatitude Latitude of the center point
	FilterLatitude float64 `form:"filter.latitude" json:"filter.latitude"`

	// FilterRadius Search radius in meters (max 100,000 = 100 km)
	FilterRadius float64 `form:"filter.radius" json:"filter.radius"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationLightaidsBoundingBoxParams defines parameters for GetLocationLightaidsBoundingBox.
type GetLocationLightaidsBoundingBoxParams struct {
	// FilterLonLeft Longitude of the left (western) edge of the bounding box
	FilterLonLeft float64 `form:"filter.lonLeft" json:"filter.lonLeft"`

	// FilterLonRight Longitude of the right (eastern) edge of the bounding box
	FilterLonRight float64 `form:"filter.lonRight" json:"filter.lonRight"`

	// FilterLatBottom Latitude of the bottom (southern) edge of the bounding box
	FilterLatBottom float64 `form:"filter.latBottom" json:"filter.latBottom"`

	// FilterLatTop Latitude of the top (northern) edge of the bounding box
	FilterLatTop float64 `form:"filter.latTop" json:"filter.latTop"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationLightaidsRadiusParams defines parameters for GetLocationLightaidsRadius.
type GetLocationLightaidsRadiusParams struct {
	// FilterLongitude Longitude of the center point
	FilterLongitude float64 `form:"filter.longitude" json:"filter.longitude"`

	// FilterLatitude Latitude of the center point
	FilterLatitude float64 `form:"filter.latitude" json:"filter.latitude"`

	// FilterRadius Search radius in meters (max 100,000 = 100 km)
	FilterRadius float64 `form:"filter.radius" json:"filter.radius"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationModuBoundingBoxParams defines parameters for GetLocationModuBoundingBox.
type GetLocationModuBoundingBoxParams struct {
	// FilterLonLeft Longitude of the left (western) edge of the bounding box
	FilterLonLeft float64 `form:"filter.lonLeft" json:"filter.lonLeft"`

	// FilterLonRight Longitude of the right (eastern) edge of the bounding box
	FilterLonRight float64 `form:"filter.lonRight" json:"filter.lonRight"`

	// FilterLatBottom Latitude of the bottom (southern) edge of the bounding box
	FilterLatBottom float64 `form:"filter.latBottom" json:"filter.latBottom"`

	// FilterLatTop Latitude of the top (northern) edge of the bounding box
	FilterLatTop float64 `form:"filter.latTop" json:"filter.latTop"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationModuRadiusParams defines parameters for GetLocationModuRadius.
type GetLocationModuRadiusParams struct {
	// FilterLongitude Longitude of the center point
	FilterLongitude float64 `form:"filter.longitude" json:"filter.longitude"`

	// FilterLatitude Latitude of the center point
	FilterLatitude float64 `form:"filter.latitude" json:"filter.latitude"`

	// FilterRadius Search radius in meters (max 100,000 = 100 km)
	FilterRadius float64 `form:"filter.radius" json:"filter.radius"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationPortsBoundingBoxParams defines parameters for GetLocationPortsBoundingBox.
type GetLocationPortsBoundingBoxParams struct {
	// FilterLonLeft Longitude of the left (western) edge of the bounding box
	FilterLonLeft float64 `form:"filter.lonLeft" json:"filter.lonLeft"`

	// FilterLonRight Longitude of the right (eastern) edge of the bounding box
	FilterLonRight float64 `form:"filter.lonRight" json:"filter.lonRight"`

	// FilterLatBottom Latitude of the bottom (southern) edge of the bounding box
	FilterLatBottom float64 `form:"filter.latBottom" json:"filter.latBottom"`

	// FilterLatTop Latitude of the top (northern) edge of the bounding box
	FilterLatTop float64 `form:"filter.latTop" json:"filter.latTop"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationPortsRadiusParams defines parameters for GetLocationPortsRadius.
type GetLocationPortsRadiusParams struct {
	// FilterLongitude Longitude of the center point
	FilterLongitude float64 `form:"filter.longitude" json:"filter.longitude"`

	// FilterLatitude Latitude of the center point
	FilterLatitude float64 `form:"filter.latitude" json:"filter.latitude"`

	// FilterRadius Search radius in meters (max 100,000 = 100 km)
	FilterRadius float64 `form:"filter.radius" json:"filter.radius"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationRadiobeaconsBoundingBoxParams defines parameters for GetLocationRadiobeaconsBoundingBox.
type GetLocationRadiobeaconsBoundingBoxParams struct {
	// FilterLonLeft Longitude of the left (western) edge of the bounding box
	FilterLonLeft float64 `form:"filter.lonLeft" json:"filter.lonLeft"`

	// FilterLonRight Longitude of the right (eastern) edge of the bounding box
	FilterLonRight float64 `form:"filter.lonRight" json:"filter.lonRight"`

	// FilterLatBottom Latitude of the bottom (southern) edge of the bounding box
	FilterLatBottom float64 `form:"filter.latBottom" json:"filter.latBottom"`

	// FilterLatTop Latitude of the top (northern) edge of the bounding box
	FilterLatTop float64 `form:"filter.latTop" json:"filter.latTop"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationRadiobeaconsRadiusParams defines parameters for GetLocationRadiobeaconsRadius.
type GetLocationRadiobeaconsRadiusParams struct {
	// FilterLongitude Longitude of the center point
	FilterLongitude float64 `form:"filter.longitude" json:"filter.longitude"`

	// FilterLatitude Latitude of the center point
	FilterLatitude float64 `form:"filter.latitude" json:"filter.latitude"`

	// FilterRadius Search radius in meters (max 100,000 = 100 km)
	FilterRadius float64 `form:"filter.radius" json:"filter.radius"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationVesselsBoundingBoxParams defines parameters for GetLocationVesselsBoundingBox.
type GetLocationVesselsBoundingBoxParams struct {
	// FilterLonLeft Longitude of the left (western) edge of the bounding box
	FilterLonLeft float64 `form:"filter.lonLeft" json:"filter.lonLeft"`

	// FilterLonRight Longitude of the right (eastern) edge of the bounding box
	FilterLonRight float64 `form:"filter.lonRight" json:"filter.lonRight"`

	// FilterLatBottom Latitude of the bottom (southern) edge of the bounding box
	FilterLatBottom float64 `form:"filter.latBottom" json:"filter.latBottom"`

	// FilterLatTop Latitude of the top (northern) edge of the bounding box
	FilterLatTop float64 `form:"filter.latTop" json:"filter.latTop"`

	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago). Max window 4 hours.
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time). Max window 4 hours.
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetLocationVesselsRadiusParams defines parameters for GetLocationVesselsRadius.
type GetLocationVesselsRadiusParams struct {
	// FilterLongitude Longitude of the center point
	FilterLongitude float64 `form:"filter.longitude" json:"filter.longitude"`

	// FilterLatitude Latitude of the center point
	FilterLatitude float64 `form:"filter.latitude" json:"filter.latitude"`

	// FilterRadius Search radius in meters (max 100,000 = 100 km)
	FilterRadius float64 `form:"filter.radius" json:"filter.radius"`

	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago). Max window 4 hours.
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time). Max window 4 hours.
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetPortUnlocodeInboundParams defines parameters for GetPortUnlocodeInbound.
type GetPortUnlocodeInboundParams struct {
	// FilterEtaFrom Start of ETA arrival window in RFC3339 format (defaults to now)
	FilterEtaFrom *string `form:"filter.etaFrom,omitempty" json:"filter.etaFrom,omitempty"`

	// FilterEtaTo End of ETA arrival window in RFC3339 format (defaults to 72 hours ahead)
	FilterEtaTo *string `form:"filter.etaTo,omitempty" json:"filter.etaTo,omitempty"`

	// TimeFrom Report freshness start (defaults to 24 hours ago)
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo Report freshness end (defaults to current time)
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetPorteventsParams defines parameters for GetPortevents.
type GetPorteventsParams struct {
	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago)
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time)
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// FilterCountry Filter by port country (case-insensitive)
	FilterCountry *string `form:"filter.country,omitempty" json:"filter.country,omitempty"`

	// FilterUnlocode Filter by port UN/LOCODE
	FilterUnlocode *string `form:"filter.unlocode,omitempty" json:"filter.unlocode,omitempty"`

	// FilterEventType Filter by event type
	FilterEventType *GetPorteventsParamsFilterEventType `form:"filter.eventType,omitempty" json:"filter.eventType,omitempty"`

	// FilterVesselName Filter by vessel name (full-text search)
	FilterVesselName *string `form:"filter.vesselName,omitempty" json:"filter.vesselName,omitempty"`

	// FilterPortName Filter by port name (full-text search)
	FilterPortName *string `form:"filter.portName,omitempty" json:"filter.portName,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetPorteventsParamsFilterEventType defines parameters for GetPortevents.
type GetPorteventsParamsFilterEventType string

// GetPorteventsPortUnlocodeParams defines parameters for GetPorteventsPortUnlocode.
type GetPorteventsPortUnlocodeParams struct {
	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago)
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time)
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetPorteventsPortsParams defines parameters for GetPorteventsPorts.
type GetPorteventsPortsParams struct {
	// FilterPortName Port name
	FilterPortName string `form:"filter.portName" json:"filter.portName"`

	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago)
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time)
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetPorteventsVesselIdParams defines parameters for GetPorteventsVesselId.
type GetPorteventsVesselIdParams struct {
	// FilterIdType Identifier type (mmsi or imo)
	FilterIdType GetPorteventsVesselIdParamsFilterIdType `form:"filter.idType" json:"filter.idType"`

	// FilterEventType Filter by event type
	FilterEventType *GetPorteventsVesselIdParamsFilterEventType `form:"filter.eventType,omitempty" json:"filter.eventType,omitempty"`

	// FilterSortOrder Sort order by timestamp
	FilterSortOrder *GetPorteventsVesselIdParamsFilterSortOrder `form:"filter.sortOrder,omitempty" json:"filter.sortOrder,omitempty"`

	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago)
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time)
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetPorteventsVesselIdParamsFilterIdType defines parameters for GetPorteventsVesselId.
type GetPorteventsVesselIdParamsFilterIdType string

// GetPorteventsVesselIdParamsFilterEventType defines parameters for GetPorteventsVesselId.
type GetPorteventsVesselIdParamsFilterEventType string

// GetPorteventsVesselIdParamsFilterSortOrder defines parameters for GetPorteventsVesselId.
type GetPorteventsVesselIdParamsFilterSortOrder string

// GetPorteventsVesselIdLastParams defines parameters for GetPorteventsVesselIdLast.
type GetPorteventsVesselIdLastParams struct {
	// FilterIdType Identifier type (mmsi or imo)
	FilterIdType GetPorteventsVesselIdLastParamsFilterIdType `form:"filter.idType" json:"filter.idType"`
}

// GetPorteventsVesselIdLastParamsFilterIdType defines parameters for GetPorteventsVesselIdLast.
type GetPorteventsVesselIdLastParamsFilterIdType string

// GetPorteventsVesselsParams defines parameters for GetPorteventsVessels.
type GetPorteventsVesselsParams struct {
	// FilterVesselName Vessel name
	FilterVesselName string `form:"filter.vesselName" json:"filter.vesselName"`

	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago)
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time)
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetSearchDgpsParams defines parameters for GetSearchDgps.
type GetSearchDgpsParams struct {
	// FilterName Name of the DGPS station
	FilterName string `form:"filter.name" json:"filter.name"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetSearchLightaidsParams defines parameters for GetSearchLightaids.
type GetSearchLightaidsParams struct {
	// FilterName Name of the Light Aid
	FilterName string `form:"filter.name" json:"filter.name"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetSearchModusParams defines parameters for GetSearchModus.
type GetSearchModusParams struct {
	// FilterName Name of the MODU
	FilterName string `form:"filter.name" json:"filter.name"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetSearchPortsParams defines parameters for GetSearchPorts.
type GetSearchPortsParams struct {
	// FilterName Name of the port
	FilterName *string `form:"filter.name,omitempty" json:"filter.name,omitempty"`

	// FilterCountry ISO 2-letter country code(s) or country name(s). Repeat the parameter to match multiple.
	FilterCountry *[]string `form:"filter.country,omitempty" json:"filter.country,omitempty"`

	// FilterType Port type classification(s) (case-insensitive). Repeat the parameter to match multiple. Note: not all ports have a type assigned.
	FilterType *[]GetSearchPortsParamsFilterType `form:"filter.type,omitempty" json:"filter.type,omitempty"`

	// FilterSize Port size classification(s). Repeat the parameter to match multiple.
	FilterSize *[]string `form:"filter.size,omitempty" json:"filter.size,omitempty"`

	// FilterRegion Geographic region name (partial match)
	FilterRegion *string `form:"filter.region,omitempty" json:"filter.region,omitempty"`

	// FilterHarborSize Harbor size classification(s). Repeat the parameter to match multiple.
	FilterHarborSize *[]string `form:"filter.harborSize,omitempty" json:"filter.harborSize,omitempty"`

	// FilterHarborUse Primary harbor use(s). Repeat the parameter to match multiple.
	FilterHarborUse *[]string `form:"filter.harborUse,omitempty" json:"filter.harborUse,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetSearchPortsParamsFilterType defines parameters for GetSearchPorts.
type GetSearchPortsParamsFilterType string

// GetSearchRadiobeaconsParams defines parameters for GetSearchRadiobeacons.
type GetSearchRadiobeaconsParams struct {
	// FilterName Name of the Radio Beacon
	FilterName string `form:"filter.name" json:"filter.name"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetSearchVesselsParams defines parameters for GetSearchVessels.
type GetSearchVesselsParams struct {
	// Q Unified search across IMO, MMSI, ENI, callsign, and vessel name. Matches a vessel if any of those identifiers matches, so one value can return more than one vessel; the _meta.matchedOn field on the response names the fields that matched. An ENI matches with or without its leading zeros. Name and callsign also accept SQL LIKE wildcards (% and _). Can be combined with filter.* parameters.
	Q *string `form:"q,omitempty" json:"q,omitempty"`

	// FilterName Name of the vessel
	FilterName *string `form:"filter.name,omitempty" json:"filter.name,omitempty"`

	// FilterCallsign Radio callsign of the vessel
	FilterCallsign *string `form:"filter.callsign,omitempty" json:"filter.callsign,omitempty"`

	// FilterFlag ISO 2-letter country code(s) of the flag state. Repeat the parameter to match multiple.
	FilterFlag *[]string `form:"filter.flag,omitempty" json:"filter.flag,omitempty"`

	// FilterVesselType Vessel type classification(s) (case-insensitive). Repeat the parameter to match multiple.
	FilterVesselType *[]string `form:"filter.vesselType,omitempty" json:"filter.vesselType,omitempty"`

	// FilterMmsi MMSI number
	FilterMmsi *int `form:"filter.mmsi,omitempty" json:"filter.mmsi,omitempty"`

	// FilterImo IMO number
	FilterImo *int `form:"filter.imo,omitempty" json:"filter.imo,omitempty"`

	// FilterYearBuiltMin Minimum year built
	FilterYearBuiltMin *int `form:"filter.yearBuiltMin,omitempty" json:"filter.yearBuiltMin,omitempty"`

	// FilterYearBuiltMax Maximum year built
	FilterYearBuiltMax *int `form:"filter.yearBuiltMax,omitempty" json:"filter.yearBuiltMax,omitempty"`

	// FilterEni European Number of Identification, the inland waterway equivalent of an IMO. Leading zeros optional: 4606770 and 04606770 both match.
	FilterEni *string `form:"filter.eni,omitempty" json:"filter.eni,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Token for next page
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetVesselIdParams defines parameters for GetVesselId.
type GetVesselIdParams struct {
	// FilterIdType Identifier type: 'mmsi' or 'imo'
	FilterIdType GetVesselIdParamsFilterIdType `form:"filter.idType" json:"filter.idType"`
}

// GetVesselIdParamsFilterIdType defines parameters for GetVesselId.
type GetVesselIdParamsFilterIdType string

// GetVesselIdCasualtiesParams defines parameters for GetVesselIdCasualties.
type GetVesselIdCasualtiesParams struct {
	// FilterIdType Identifier type: 'mmsi' or 'imo'
	FilterIdType GetVesselIdCasualtiesParamsFilterIdType `form:"filter.idType" json:"filter.idType"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Pagination token for retrieving the next page of results
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetVesselIdCasualtiesParamsFilterIdType defines parameters for GetVesselIdCasualties.
type GetVesselIdCasualtiesParamsFilterIdType string

// GetVesselIdEmissionsParams defines parameters for GetVesselIdEmissions.
type GetVesselIdEmissionsParams struct {
	// FilterIdType Identifier type: 'mmsi' or 'imo'
	FilterIdType GetVesselIdEmissionsParamsFilterIdType `form:"filter.idType" json:"filter.idType"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Pagination token for retrieving the next page of results
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetVesselIdEmissionsParamsFilterIdType defines parameters for GetVesselIdEmissions.
type GetVesselIdEmissionsParamsFilterIdType string

// GetVesselIdEtaParams defines parameters for GetVesselIdEta.
type GetVesselIdEtaParams struct {
	// FilterIdType Identifier type: 'mmsi' or 'imo'
	FilterIdType GetVesselIdEtaParamsFilterIdType `form:"filter.idType" json:"filter.idType"`
}

// GetVesselIdEtaParamsFilterIdType defines parameters for GetVesselIdEta.
type GetVesselIdEtaParamsFilterIdType string

// GetVesselIdPositionParams defines parameters for GetVesselIdPosition.
type GetVesselIdPositionParams struct {
	// FilterIdType Identifier type: 'mmsi' or 'imo'
	FilterIdType GetVesselIdPositionParamsFilterIdType `form:"filter.idType" json:"filter.idType"`

	// FilterSat Deprecated: the satellite fallback is withdrawn in its current form. Until the cutoff date (see the Sunset header) it enables the satellite AIS fallback for existing credit holders; from that date it is accepted and ignored and no credit is charged (default: false)
	// Deprecated: The satellite fallback is withdrawn in its current form. Accepted and ignored from the cutoff date in the Sunset header.
	FilterSat *bool `form:"filter.sat,omitempty" json:"filter.sat,omitempty"`

	// FilterSatLookbackMinutes Deprecated: the satellite fallback is withdrawn in its current form. Ignored from the cutoff date. With filter.sat=true: how old a satellite-lookup fix may be and still be returned, 60-4800 minutes (default: 180)
	// Deprecated: The satellite fallback is withdrawn in its current form. Accepted and ignored from the cutoff date in the Sunset header.
	FilterSatLookbackMinutes *int `form:"filter.satLookbackMinutes,omitempty" json:"filter.satLookbackMinutes,omitempty"`

	// FilterSatMaxAgeMinutes Deprecated: the satellite fallback is withdrawn in its current form. Ignored from the cutoff date. With filter.sat=true: how old a stored position may be before a satellite lookup is made, 15-4800 minutes (default: 60). Paid plans only
	// Deprecated: The satellite fallback is withdrawn in its current form. Accepted and ignored from the cutoff date in the Sunset header.
	FilterSatMaxAgeMinutes *int `form:"filter.satMaxAgeMinutes,omitempty" json:"filter.satMaxAgeMinutes,omitempty"`
}

// GetVesselIdPositionParamsFilterIdType defines parameters for GetVesselIdPosition.
type GetVesselIdPositionParamsFilterIdType string

// GetVesselsPositionsParams defines parameters for GetVesselsPositions.
type GetVesselsPositionsParams struct {
	// FilterIds MMSI or IMO number(s): comma-separated in one param, or repeat filter.ids for multiple
	FilterIds string `form:"filter.ids" json:"filter.ids"`

	// FilterIdType Identifier type: 'mmsi' or 'imo'
	FilterIdType GetVesselsPositionsParamsFilterIdType `form:"filter.idType" json:"filter.idType"`

	// TimeFrom Start timestamp in RFC3339 format (defaults to 2 hours ago)
	TimeFrom *string `form:"time.from,omitempty" json:"time.from,omitempty"`

	// TimeTo End timestamp in RFC3339 format (defaults to current time)
	TimeTo *string `form:"time.to,omitempty" json:"time.to,omitempty"`

	// PaginationLimit Maximum number of items to return, must be between 1 and 50
	PaginationLimit *int `form:"pagination.limit,omitempty" json:"pagination.limit,omitempty"`

	// PaginationNextToken Pagination token for retrieving the next page of results
	PaginationNextToken *string `form:"pagination.nextToken,omitempty" json:"pagination.nextToken,omitempty"`
}

// GetVesselsPositionsParamsFilterIdType defines parameters for GetVesselsPositions.
type GetVesselsPositionsParamsFilterIdType string

// RequestEditorFn is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.vesselapi.com/v1 for example. This can contain a path relative
	// to the server, such as https://api.vesselapi.com/v1, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {

	// GetEmissions List emissions data
	//
	// Retrieves emissions data with optional filtering by reporting period.
	//
	// Corresponds with GET /emissions (the `GetEmissions` operationId).
	GetEmissions(ctx context.Context, params *GetEmissionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationDgpsBoundingBox Get DGPS Stations within a bounding box
	//
	// Retrieves Dgps stations within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Corresponds with GET /location/dgps/bounding-box (the `GetLocationDgpsBoundingBox` operationId).
	GetLocationDgpsBoundingBox(ctx context.Context, params *GetLocationDgpsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationDgpsRadius Get DGPS Stations within a radius
	//
	// Retrieves Dgps stations within a specified radius of a given point.
	//
	// Corresponds with GET /location/dgps/radius (the `GetLocationDgpsRadius` operationId).
	GetLocationDgpsRadius(ctx context.Context, params *GetLocationDgpsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationLightaidsBoundingBox Get Light Aids to Navigation within a bounding box
	//
	// Retrieves Light Aids to Navigation within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Corresponds with GET /location/lightaids/bounding-box (the `GetLocationLightaidsBoundingBox` operationId).
	GetLocationLightaidsBoundingBox(ctx context.Context, params *GetLocationLightaidsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationLightaidsRadius Get Light Aids to Navigation within a radius
	//
	// Retrieves Light Aids to Navigation within a specified radius of a given point.
	//
	// Corresponds with GET /location/lightaids/radius (the `GetLocationLightaidsRadius` operationId).
	GetLocationLightaidsRadius(ctx context.Context, params *GetLocationLightaidsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationModuBoundingBox Get Mobile Offshore Drilling Units within a bounding box
	//
	// Retrieves Mobile Offshore Drilling Units within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Corresponds with GET /location/modu/bounding-box (the `GetLocationModuBoundingBox` operationId).
	GetLocationModuBoundingBox(ctx context.Context, params *GetLocationModuBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationModuRadius Get Mobile Offshore Drilling Units within a radius
	//
	// Retrieves Mobile Offshore Drilling Units within a specified radius of a given point.
	//
	// Corresponds with GET /location/modu/radius (the `GetLocationModuRadius` operationId).
	GetLocationModuRadius(ctx context.Context, params *GetLocationModuRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationPortsBoundingBox Get Ports within a bounding box
	//
	// Retrieves Ports within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Corresponds with GET /location/ports/bounding-box (the `GetLocationPortsBoundingBox` operationId).
	GetLocationPortsBoundingBox(ctx context.Context, params *GetLocationPortsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationPortsRadius Get Ports within a radius
	//
	// Retrieves Ports within a specified radius of a given point.
	//
	// Corresponds with GET /location/ports/radius (the `GetLocationPortsRadius` operationId).
	GetLocationPortsRadius(ctx context.Context, params *GetLocationPortsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationRadiobeaconsBoundingBox Get Radio Beacons within a bounding box
	//
	// Retrieves Radio Beacons within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Corresponds with GET /location/radiobeacons/bounding-box (the `GetLocationRadiobeaconsBoundingBox` operationId).
	GetLocationRadiobeaconsBoundingBox(ctx context.Context, params *GetLocationRadiobeaconsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationRadiobeaconsRadius Get Radio Beacons within a radius
	//
	// Retrieves Radio Beacons within a specified radius of a given point.
	//
	// Corresponds with GET /location/radiobeacons/radius (the `GetLocationRadiobeaconsRadius` operationId).
	GetLocationRadiobeaconsRadius(ctx context.Context, params *GetLocationRadiobeaconsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationVesselsBoundingBox Get vessels within a bounding box
	//
	// Retrieves vessels within a specified bounding box and time window. Max bbox span: |dLat| + |dLon| ≤ 4 degrees. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
	//
	// Corresponds with GET /location/vessels/bounding-box (the `GetLocationVesselsBoundingBox` operationId).
	GetLocationVesselsBoundingBox(ctx context.Context, params *GetLocationVesselsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetLocationVesselsRadius Get vessels within a radius
	//
	// Retrieves vessels within a specified radius of a given point and a given time range. Max radius: 100 km. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
	//
	// Corresponds with GET /location/vessels/radius (the `GetLocationVesselsRadius` operationId).
	GetLocationVesselsRadius(ctx context.Context, params *GetLocationVesselsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPortUnlocode Get Port by UNLOCODE
	//
	// Retrieves port details by its UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
	//
	// Corresponds with GET /port/{unlocode} (the `GetPortUnlocode` operationId).
	GetPortUnlocode(ctx context.Context, unlocode string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPortUnlocodeInbound Get Inbound Vessels for Port
	//
	// Retrieves vessels heading to a specific port within an ETA window.
	//
	// Corresponds with GET /port/{unlocode}/inbound (the `GetPortUnlocodeInbound` operationId).
	GetPortUnlocodeInbound(ctx context.Context, unlocode string, params *GetPortUnlocodeInboundParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPortevents Get Port Events in a time range
	//
	// Get Port Events, such as Arrivals and Departures, in a time range with optional filters
	//
	// Corresponds with GET /portevents (the `GetPortevents` operationId).
	GetPortevents(ctx context.Context, params *GetPorteventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPorteventsPortUnlocode Get Port Events by port UNLOCODE
	//
	// Get Port Events, such as Arrivals and Departures by supplying the port UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
	//
	// Corresponds with GET /portevents/port/{unlocode} (the `GetPorteventsPortUnlocode` operationId).
	GetPorteventsPortUnlocode(ctx context.Context, unlocode string, params *GetPorteventsPortUnlocodeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPorteventsPorts Get Port Events by port name
	//
	// Get Port Events, such as Arrivals and Departures by supplying the port name
	//
	// Corresponds with GET /portevents/ports (the `GetPorteventsPorts` operationId).
	GetPorteventsPorts(ctx context.Context, params *GetPorteventsPortsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPorteventsVesselId Get Port Events by Vessel ID
	//
	// Get all port events (arrivals and departures) for a vessel identified by MMSI or IMO, with optional filtering by event type and time range
	//
	// Corresponds with GET /portevents/vessel/{id} (the `GetPorteventsVesselId` operationId).
	GetPorteventsVesselId(ctx context.Context, id string, params *GetPorteventsVesselIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPorteventsVesselIdLast Get Last Port Event by ID
	//
	// Get the most recent port event (arrival or departure) for a vessel identified by MMSI or IMO.
	//
	// Corresponds with GET /portevents/vessel/{id}/last (the `GetPorteventsVesselIdLast` operationId).
	GetPorteventsVesselIdLast(ctx context.Context, id string, params *GetPorteventsVesselIdLastParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPorteventsVessels Get Port Events by vessel name
	//
	// Corresponds with GET /portevents/vessels (the `GetPorteventsVessels` operationId).
	GetPorteventsVessels(ctx context.Context, params *GetPorteventsVesselsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSearchDgps Search for DGPS Stations
	//
	// Retrieves a list of DGPS stations for the given query parameters.
	//
	// Corresponds with GET /search/dgps (the `GetSearchDgps` operationId).
	GetSearchDgps(ctx context.Context, params *GetSearchDgpsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSearchLightaids Search for Light Aids to Navigation
	//
	// Retrieves a list of Light Aids to Navigation for the given query parameters.
	//
	// Corresponds with GET /search/lightaids (the `GetSearchLightaids` operationId).
	GetSearchLightaids(ctx context.Context, params *GetSearchLightaidsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSearchModus Search for MODUs
	//
	// Retrieves a list of MODUs for the given query parameters.
	//
	// Corresponds with GET /search/modus (the `GetSearchModus` operationId).
	GetSearchModus(ctx context.Context, params *GetSearchModusParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSearchPorts Search for Ports
	//
	// Retrieves a list of ports matching the given filters. At least one filter parameter is required.
	//
	// Corresponds with GET /search/ports (the `GetSearchPorts` operationId).
	GetSearchPorts(ctx context.Context, params *GetSearchPortsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSearchRadiobeacons Search for Radio Beacons
	//
	// Retrieves a list of Radio Beacons for the given query parameters.
	//
	// Corresponds with GET /search/radiobeacons (the `GetSearchRadiobeacons` operationId).
	GetSearchRadiobeacons(ctx context.Context, params *GetSearchRadiobeaconsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSearchVessels Search for Vessels
	//
	// Retrieves a list of vessels matching the given filters. At least one filter parameter (or the unified `q` parameter) is required.
	//
	// Corresponds with GET /search/vessels (the `GetSearchVessels` operationId).
	GetSearchVessels(ctx context.Context, params *GetSearchVesselsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVesselId Get vessel information by MMSI or IMO
	//
	// Retrieves static vessel data including name, type, dimensions, and registration information for a vessel identified by its MMSI or IMO number
	//
	// Corresponds with GET /vessel/{id} (the `GetVesselId` operationId).
	GetVesselId(ctx context.Context, id string, params *GetVesselIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVesselIdCasualties Get marine casualties involving a vessel
	//
	// Retrieves marine casualty records involving the specified vessel.
	//
	// Corresponds with GET /vessel/{id}/casualties (the `GetVesselIdCasualties` operationId).
	GetVesselIdCasualties(ctx context.Context, id string, params *GetVesselIdCasualtiesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVesselIdEmissions Get emissions data for a vessel
	//
	// Retrieves emissions reports for the specified vessel including CO2 emissions, fuel consumption, and efficiency metrics
	//
	// Corresponds with GET /vessel/{id}/emissions (the `GetVesselIdEmissions` operationId).
	GetVesselIdEmissions(ctx context.Context, id string, params *GetVesselIdEmissionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVesselIdEta Get vessel latest ETA
	//
	// Retrieves the most recent Estimated Time of Arrival (ETA) reported by the vessel.
	//
	// Corresponds with GET /vessel/{id}/eta (the `GetVesselIdEta` operationId).
	GetVesselIdEta(ctx context.Context, id string, params *GetVesselIdEtaParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVesselIdPosition Get last known vessel position
	//
	// Retrieves the most recent AIS position report for a vessel, including coordinates, vessel identifiers, and timestamps. The filter.sat satellite fallback is deprecated and withdrawn in its current form: it works for existing credit holders until the cutoff date, after which it is accepted and ignored. The response also includes the vessel's particulars (vessel), the latest destination and ETA it reported (vesselEta), and positionCheck, which compares the position with the vessel's previous position report: time and distance since then, implied speed and course, and whether the identifiers match.
	//
	// Corresponds with GET /vessel/{id}/position (the `GetVesselIdPosition` operationId).
	GetVesselIdPosition(ctx context.Context, id string, params *GetVesselIdPositionParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVesselsPositions Get positions for multiple vessels
	//
	// Retrieves AIS position data for multiple vessels identified by MMSI or IMO numbers within a specified time range (defaults to past 2 hours). Provide multiple IDs either as a comma-separated list in one filter.ids param, or by repeating filter.ids; both forms (and a mix) are accepted.
	//
	// Corresponds with GET /vessels/positions (the `GetVesselsPositions` operationId).
	GetVesselsPositions(ctx context.Context, params *GetVesselsPositionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

// GetEmissions List emissions data
//
// Retrieves emissions data with optional filtering by reporting period.
//
// Corresponds with GET /emissions (the `GetEmissions` operationId).
func (c *Client) GetEmissions(ctx context.Context, params *GetEmissionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetEmissionsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationDgpsBoundingBox Get DGPS Stations within a bounding box
//
// Retrieves Dgps stations within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Corresponds with GET /location/dgps/bounding-box (the `GetLocationDgpsBoundingBox` operationId).
func (c *Client) GetLocationDgpsBoundingBox(ctx context.Context, params *GetLocationDgpsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationDgpsBoundingBoxRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationDgpsRadius Get DGPS Stations within a radius
//
// Retrieves Dgps stations within a specified radius of a given point.
//
// Corresponds with GET /location/dgps/radius (the `GetLocationDgpsRadius` operationId).
func (c *Client) GetLocationDgpsRadius(ctx context.Context, params *GetLocationDgpsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationDgpsRadiusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationLightaidsBoundingBox Get Light Aids to Navigation within a bounding box
//
// Retrieves Light Aids to Navigation within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Corresponds with GET /location/lightaids/bounding-box (the `GetLocationLightaidsBoundingBox` operationId).
func (c *Client) GetLocationLightaidsBoundingBox(ctx context.Context, params *GetLocationLightaidsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationLightaidsBoundingBoxRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationLightaidsRadius Get Light Aids to Navigation within a radius
//
// Retrieves Light Aids to Navigation within a specified radius of a given point.
//
// Corresponds with GET /location/lightaids/radius (the `GetLocationLightaidsRadius` operationId).
func (c *Client) GetLocationLightaidsRadius(ctx context.Context, params *GetLocationLightaidsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationLightaidsRadiusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationModuBoundingBox Get Mobile Offshore Drilling Units within a bounding box
//
// Retrieves Mobile Offshore Drilling Units within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Corresponds with GET /location/modu/bounding-box (the `GetLocationModuBoundingBox` operationId).
func (c *Client) GetLocationModuBoundingBox(ctx context.Context, params *GetLocationModuBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationModuBoundingBoxRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationModuRadius Get Mobile Offshore Drilling Units within a radius
//
// Retrieves Mobile Offshore Drilling Units within a specified radius of a given point.
//
// Corresponds with GET /location/modu/radius (the `GetLocationModuRadius` operationId).
func (c *Client) GetLocationModuRadius(ctx context.Context, params *GetLocationModuRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationModuRadiusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationPortsBoundingBox Get Ports within a bounding box
//
// Retrieves Ports within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Corresponds with GET /location/ports/bounding-box (the `GetLocationPortsBoundingBox` operationId).
func (c *Client) GetLocationPortsBoundingBox(ctx context.Context, params *GetLocationPortsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationPortsBoundingBoxRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationPortsRadius Get Ports within a radius
//
// Retrieves Ports within a specified radius of a given point.
//
// Corresponds with GET /location/ports/radius (the `GetLocationPortsRadius` operationId).
func (c *Client) GetLocationPortsRadius(ctx context.Context, params *GetLocationPortsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationPortsRadiusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationRadiobeaconsBoundingBox Get Radio Beacons within a bounding box
//
// Retrieves Radio Beacons within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Corresponds with GET /location/radiobeacons/bounding-box (the `GetLocationRadiobeaconsBoundingBox` operationId).
func (c *Client) GetLocationRadiobeaconsBoundingBox(ctx context.Context, params *GetLocationRadiobeaconsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationRadiobeaconsBoundingBoxRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationRadiobeaconsRadius Get Radio Beacons within a radius
//
// Retrieves Radio Beacons within a specified radius of a given point.
//
// Corresponds with GET /location/radiobeacons/radius (the `GetLocationRadiobeaconsRadius` operationId).
func (c *Client) GetLocationRadiobeaconsRadius(ctx context.Context, params *GetLocationRadiobeaconsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationRadiobeaconsRadiusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationVesselsBoundingBox Get vessels within a bounding box
//
// Retrieves vessels within a specified bounding box and time window. Max bbox span: |dLat| + |dLon| ≤ 4 degrees. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
//
// Corresponds with GET /location/vessels/bounding-box (the `GetLocationVesselsBoundingBox` operationId).
func (c *Client) GetLocationVesselsBoundingBox(ctx context.Context, params *GetLocationVesselsBoundingBoxParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationVesselsBoundingBoxRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetLocationVesselsRadius Get vessels within a radius
//
// Retrieves vessels within a specified radius of a given point and a given time range. Max radius: 100 km. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
//
// Corresponds with GET /location/vessels/radius (the `GetLocationVesselsRadius` operationId).
func (c *Client) GetLocationVesselsRadius(ctx context.Context, params *GetLocationVesselsRadiusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLocationVesselsRadiusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPortUnlocode Get Port by UNLOCODE
//
// Retrieves port details by its UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
//
// Corresponds with GET /port/{unlocode} (the `GetPortUnlocode` operationId).
func (c *Client) GetPortUnlocode(ctx context.Context, unlocode string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPortUnlocodeRequest(c.Server, unlocode)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPortUnlocodeInbound Get Inbound Vessels for Port
//
// Retrieves vessels heading to a specific port within an ETA window.
//
// Corresponds with GET /port/{unlocode}/inbound (the `GetPortUnlocodeInbound` operationId).
func (c *Client) GetPortUnlocodeInbound(ctx context.Context, unlocode string, params *GetPortUnlocodeInboundParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPortUnlocodeInboundRequest(c.Server, unlocode, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPortevents Get Port Events in a time range
//
// # Get Port Events, such as Arrivals and Departures, in a time range with optional filters
//
// Corresponds with GET /portevents (the `GetPortevents` operationId).
func (c *Client) GetPortevents(ctx context.Context, params *GetPorteventsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPorteventsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPorteventsPortUnlocode Get Port Events by port UNLOCODE
//
// Get Port Events, such as Arrivals and Departures by supplying the port UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
//
// Corresponds with GET /portevents/port/{unlocode} (the `GetPorteventsPortUnlocode` operationId).
func (c *Client) GetPorteventsPortUnlocode(ctx context.Context, unlocode string, params *GetPorteventsPortUnlocodeParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPorteventsPortUnlocodeRequest(c.Server, unlocode, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPorteventsPorts Get Port Events by port name
//
// # Get Port Events, such as Arrivals and Departures by supplying the port name
//
// Corresponds with GET /portevents/ports (the `GetPorteventsPorts` operationId).
func (c *Client) GetPorteventsPorts(ctx context.Context, params *GetPorteventsPortsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPorteventsPortsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPorteventsVesselId Get Port Events by Vessel ID
//
// # Get all port events (arrivals and departures) for a vessel identified by MMSI or IMO, with optional filtering by event type and time range
//
// Corresponds with GET /portevents/vessel/{id} (the `GetPorteventsVesselId` operationId).
func (c *Client) GetPorteventsVesselId(ctx context.Context, id string, params *GetPorteventsVesselIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPorteventsVesselIdRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPorteventsVesselIdLast Get Last Port Event by ID
//
// Get the most recent port event (arrival or departure) for a vessel identified by MMSI or IMO.
//
// Corresponds with GET /portevents/vessel/{id}/last (the `GetPorteventsVesselIdLast` operationId).
func (c *Client) GetPorteventsVesselIdLast(ctx context.Context, id string, params *GetPorteventsVesselIdLastParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPorteventsVesselIdLastRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetPorteventsVessels Get Port Events by vessel name
//
// Corresponds with GET /portevents/vessels (the `GetPorteventsVessels` operationId).
func (c *Client) GetPorteventsVessels(ctx context.Context, params *GetPorteventsVesselsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPorteventsVesselsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetSearchDgps Search for DGPS Stations
//
// Retrieves a list of DGPS stations for the given query parameters.
//
// Corresponds with GET /search/dgps (the `GetSearchDgps` operationId).
func (c *Client) GetSearchDgps(ctx context.Context, params *GetSearchDgpsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSearchDgpsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetSearchLightaids Search for Light Aids to Navigation
//
// Retrieves a list of Light Aids to Navigation for the given query parameters.
//
// Corresponds with GET /search/lightaids (the `GetSearchLightaids` operationId).
func (c *Client) GetSearchLightaids(ctx context.Context, params *GetSearchLightaidsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSearchLightaidsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetSearchModus Search for MODUs
//
// Retrieves a list of MODUs for the given query parameters.
//
// Corresponds with GET /search/modus (the `GetSearchModus` operationId).
func (c *Client) GetSearchModus(ctx context.Context, params *GetSearchModusParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSearchModusRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetSearchPorts Search for Ports
//
// Retrieves a list of ports matching the given filters. At least one filter parameter is required.
//
// Corresponds with GET /search/ports (the `GetSearchPorts` operationId).
func (c *Client) GetSearchPorts(ctx context.Context, params *GetSearchPortsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSearchPortsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetSearchRadiobeacons Search for Radio Beacons
//
// Retrieves a list of Radio Beacons for the given query parameters.
//
// Corresponds with GET /search/radiobeacons (the `GetSearchRadiobeacons` operationId).
func (c *Client) GetSearchRadiobeacons(ctx context.Context, params *GetSearchRadiobeaconsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSearchRadiobeaconsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetSearchVessels Search for Vessels
//
// Retrieves a list of vessels matching the given filters. At least one filter parameter (or the unified `q` parameter) is required.
//
// Corresponds with GET /search/vessels (the `GetSearchVessels` operationId).
func (c *Client) GetSearchVessels(ctx context.Context, params *GetSearchVesselsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSearchVesselsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetVesselId Get vessel information by MMSI or IMO
//
// # Retrieves static vessel data including name, type, dimensions, and registration information for a vessel identified by its MMSI or IMO number
//
// Corresponds with GET /vessel/{id} (the `GetVesselId` operationId).
func (c *Client) GetVesselId(ctx context.Context, id string, params *GetVesselIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVesselIdRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetVesselIdCasualties Get marine casualties involving a vessel
//
// Retrieves marine casualty records involving the specified vessel.
//
// Corresponds with GET /vessel/{id}/casualties (the `GetVesselIdCasualties` operationId).
func (c *Client) GetVesselIdCasualties(ctx context.Context, id string, params *GetVesselIdCasualtiesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVesselIdCasualtiesRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetVesselIdEmissions Get emissions data for a vessel
//
// # Retrieves emissions reports for the specified vessel including CO2 emissions, fuel consumption, and efficiency metrics
//
// Corresponds with GET /vessel/{id}/emissions (the `GetVesselIdEmissions` operationId).
func (c *Client) GetVesselIdEmissions(ctx context.Context, id string, params *GetVesselIdEmissionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVesselIdEmissionsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetVesselIdEta Get vessel latest ETA
//
// Retrieves the most recent Estimated Time of Arrival (ETA) reported by the vessel.
//
// Corresponds with GET /vessel/{id}/eta (the `GetVesselIdEta` operationId).
func (c *Client) GetVesselIdEta(ctx context.Context, id string, params *GetVesselIdEtaParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVesselIdEtaRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetVesselIdPosition Get last known vessel position
//
// Retrieves the most recent AIS position report for a vessel, including coordinates, vessel identifiers, and timestamps. The filter.sat satellite fallback is deprecated and withdrawn in its current form: it works for existing credit holders until the cutoff date, after which it is accepted and ignored. The response also includes the vessel's particulars (vessel), the latest destination and ETA it reported (vesselEta), and positionCheck, which compares the position with the vessel's previous position report: time and distance since then, implied speed and course, and whether the identifiers match.
//
// Corresponds with GET /vessel/{id}/position (the `GetVesselIdPosition` operationId).
func (c *Client) GetVesselIdPosition(ctx context.Context, id string, params *GetVesselIdPositionParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVesselIdPositionRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetVesselsPositions Get positions for multiple vessels
//
// Retrieves AIS position data for multiple vessels identified by MMSI or IMO numbers within a specified time range (defaults to past 2 hours). Provide multiple IDs either as a comma-separated list in one filter.ids param, or by repeating filter.ids; both forms (and a mix) are accepted.
//
// Corresponds with GET /vessels/positions (the `GetVesselsPositions` operationId).
func (c *Client) GetVesselsPositions(ctx context.Context, params *GetVesselsPositionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVesselsPositionsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetEmissionsRequest constructs an http.Request for the GetEmissions method
func NewGetEmissionsRequest(server string, params *GetEmissionsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/emissions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.FilterPeriod != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.period", *params.FilterPeriod, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationDgpsBoundingBoxRequest constructs an http.Request for the GetLocationDgpsBoundingBox method
func NewGetLocationDgpsBoundingBoxRequest(server string, params *GetLocationDgpsBoundingBoxParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/dgps/bounding-box")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonLeft", params.FilterLonLeft, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonRight", params.FilterLonRight, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latBottom", params.FilterLatBottom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latTop", params.FilterLatTop, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationDgpsRadiusRequest constructs an http.Request for the GetLocationDgpsRadius method
func NewGetLocationDgpsRadiusRequest(server string, params *GetLocationDgpsRadiusParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/dgps/radius")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.longitude", params.FilterLongitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latitude", params.FilterLatitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.radius", params.FilterRadius, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationLightaidsBoundingBoxRequest constructs an http.Request for the GetLocationLightaidsBoundingBox method
func NewGetLocationLightaidsBoundingBoxRequest(server string, params *GetLocationLightaidsBoundingBoxParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/lightaids/bounding-box")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonLeft", params.FilterLonLeft, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonRight", params.FilterLonRight, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latBottom", params.FilterLatBottom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latTop", params.FilterLatTop, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationLightaidsRadiusRequest constructs an http.Request for the GetLocationLightaidsRadius method
func NewGetLocationLightaidsRadiusRequest(server string, params *GetLocationLightaidsRadiusParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/lightaids/radius")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.longitude", params.FilterLongitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latitude", params.FilterLatitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.radius", params.FilterRadius, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationModuBoundingBoxRequest constructs an http.Request for the GetLocationModuBoundingBox method
func NewGetLocationModuBoundingBoxRequest(server string, params *GetLocationModuBoundingBoxParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/modu/bounding-box")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonLeft", params.FilterLonLeft, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonRight", params.FilterLonRight, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latBottom", params.FilterLatBottom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latTop", params.FilterLatTop, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationModuRadiusRequest constructs an http.Request for the GetLocationModuRadius method
func NewGetLocationModuRadiusRequest(server string, params *GetLocationModuRadiusParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/modu/radius")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.longitude", params.FilterLongitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latitude", params.FilterLatitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.radius", params.FilterRadius, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationPortsBoundingBoxRequest constructs an http.Request for the GetLocationPortsBoundingBox method
func NewGetLocationPortsBoundingBoxRequest(server string, params *GetLocationPortsBoundingBoxParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/ports/bounding-box")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonLeft", params.FilterLonLeft, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonRight", params.FilterLonRight, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latBottom", params.FilterLatBottom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latTop", params.FilterLatTop, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationPortsRadiusRequest constructs an http.Request for the GetLocationPortsRadius method
func NewGetLocationPortsRadiusRequest(server string, params *GetLocationPortsRadiusParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/ports/radius")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.longitude", params.FilterLongitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latitude", params.FilterLatitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.radius", params.FilterRadius, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationRadiobeaconsBoundingBoxRequest constructs an http.Request for the GetLocationRadiobeaconsBoundingBox method
func NewGetLocationRadiobeaconsBoundingBoxRequest(server string, params *GetLocationRadiobeaconsBoundingBoxParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/radiobeacons/bounding-box")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonLeft", params.FilterLonLeft, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonRight", params.FilterLonRight, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latBottom", params.FilterLatBottom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latTop", params.FilterLatTop, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationRadiobeaconsRadiusRequest constructs an http.Request for the GetLocationRadiobeaconsRadius method
func NewGetLocationRadiobeaconsRadiusRequest(server string, params *GetLocationRadiobeaconsRadiusParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/radiobeacons/radius")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.longitude", params.FilterLongitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latitude", params.FilterLatitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.radius", params.FilterRadius, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationVesselsBoundingBoxRequest constructs an http.Request for the GetLocationVesselsBoundingBox method
func NewGetLocationVesselsBoundingBoxRequest(server string, params *GetLocationVesselsBoundingBoxParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/vessels/bounding-box")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonLeft", params.FilterLonLeft, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.lonRight", params.FilterLonRight, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latBottom", params.FilterLatBottom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latTop", params.FilterLatTop, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetLocationVesselsRadiusRequest constructs an http.Request for the GetLocationVesselsRadius method
func NewGetLocationVesselsRadiusRequest(server string, params *GetLocationVesselsRadiusParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/location/vessels/radius")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.longitude", params.FilterLongitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.latitude", params.FilterLatitude, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.radius", params.FilterRadius, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "number", Format: "double"}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPortUnlocodeRequest constructs an http.Request for the GetPortUnlocode method
func NewGetPortUnlocodeRequest(server string, unlocode string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "unlocode", unlocode, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/port/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPortUnlocodeInboundRequest constructs an http.Request for the GetPortUnlocodeInbound method
func NewGetPortUnlocodeInboundRequest(server string, unlocode string, params *GetPortUnlocodeInboundParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "unlocode", unlocode, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/port/%s/inbound", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.FilterEtaFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.etaFrom", *params.FilterEtaFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterEtaTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.etaTo", *params.FilterEtaTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPorteventsRequest constructs an http.Request for the GetPortevents method
func NewGetPorteventsRequest(server string, params *GetPorteventsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/portevents")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterCountry != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.country", *params.FilterCountry, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterUnlocode != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.unlocode", *params.FilterUnlocode, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterEventType != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.eventType", *params.FilterEventType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterVesselName != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.vesselName", *params.FilterVesselName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterPortName != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.portName", *params.FilterPortName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPorteventsPortUnlocodeRequest constructs an http.Request for the GetPorteventsPortUnlocode method
func NewGetPorteventsPortUnlocodeRequest(server string, unlocode string, params *GetPorteventsPortUnlocodeParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "unlocode", unlocode, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/portevents/port/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPorteventsPortsRequest constructs an http.Request for the GetPorteventsPorts method
func NewGetPorteventsPortsRequest(server string, params *GetPorteventsPortsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/portevents/ports")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.portName", params.FilterPortName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPorteventsVesselIdRequest constructs an http.Request for the GetPorteventsVesselId method
func NewGetPorteventsVesselIdRequest(server string, id string, params *GetPorteventsVesselIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "id", id, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/portevents/vessel/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.FilterEventType != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.eventType", *params.FilterEventType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterSortOrder != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.sortOrder", *params.FilterSortOrder, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPorteventsVesselIdLastRequest constructs an http.Request for the GetPorteventsVesselIdLast method
func NewGetPorteventsVesselIdLastRequest(server string, id string, params *GetPorteventsVesselIdLastParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "id", id, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/portevents/vessel/%s/last", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPorteventsVesselsRequest constructs an http.Request for the GetPorteventsVessels method
func NewGetPorteventsVesselsRequest(server string, params *GetPorteventsVesselsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/portevents/vessels")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.vesselName", params.FilterVesselName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSearchDgpsRequest constructs an http.Request for the GetSearchDgps method
func NewGetSearchDgpsRequest(server string, params *GetSearchDgpsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/search/dgps")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.name", params.FilterName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSearchLightaidsRequest constructs an http.Request for the GetSearchLightaids method
func NewGetSearchLightaidsRequest(server string, params *GetSearchLightaidsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/search/lightaids")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.name", params.FilterName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSearchModusRequest constructs an http.Request for the GetSearchModus method
func NewGetSearchModusRequest(server string, params *GetSearchModusParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/search/modus")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.name", params.FilterName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSearchPortsRequest constructs an http.Request for the GetSearchPorts method
func NewGetSearchPortsRequest(server string, params *GetSearchPortsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/search/ports")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.FilterName != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.name", *params.FilterName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterCountry != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.country", *params.FilterCountry, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "array", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterType != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.type", *params.FilterType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "array", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterSize != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.size", *params.FilterSize, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "array", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterRegion != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.region", *params.FilterRegion, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterHarborSize != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.harborSize", *params.FilterHarborSize, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "array", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterHarborUse != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.harborUse", *params.FilterHarborUse, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "array", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSearchRadiobeaconsRequest constructs an http.Request for the GetSearchRadiobeacons method
func NewGetSearchRadiobeaconsRequest(server string, params *GetSearchRadiobeaconsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/search/radiobeacons")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.name", params.FilterName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSearchVesselsRequest constructs an http.Request for the GetSearchVessels method
func NewGetSearchVesselsRequest(server string, params *GetSearchVesselsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/search/vessels")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Q != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "q", *params.Q, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterName != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.name", *params.FilterName, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterCallsign != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.callsign", *params.FilterCallsign, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterFlag != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.flag", *params.FilterFlag, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "array", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterVesselType != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.vesselType", *params.FilterVesselType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "array", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterMmsi != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.mmsi", *params.FilterMmsi, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterImo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.imo", *params.FilterImo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterYearBuiltMin != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.yearBuiltMin", *params.FilterYearBuiltMin, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterYearBuiltMax != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.yearBuiltMax", *params.FilterYearBuiltMax, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterEni != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.eni", *params.FilterEni, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVesselIdRequest constructs an http.Request for the GetVesselId method
func NewGetVesselIdRequest(server string, id string, params *GetVesselIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "id", id, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vessel/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVesselIdCasualtiesRequest constructs an http.Request for the GetVesselIdCasualties method
func NewGetVesselIdCasualtiesRequest(server string, id string, params *GetVesselIdCasualtiesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "id", id, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vessel/%s/casualties", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVesselIdEmissionsRequest constructs an http.Request for the GetVesselIdEmissions method
func NewGetVesselIdEmissionsRequest(server string, id string, params *GetVesselIdEmissionsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "id", id, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vessel/%s/emissions", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVesselIdEtaRequest constructs an http.Request for the GetVesselIdEta method
func NewGetVesselIdEtaRequest(server string, id string, params *GetVesselIdEtaParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "id", id, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vessel/%s/eta", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVesselIdPositionRequest constructs an http.Request for the GetVesselIdPosition method
func NewGetVesselIdPositionRequest(server string, id string, params *GetVesselIdPositionParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "id", id, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vessel/%s/position", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.FilterSat != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.sat", *params.FilterSat, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "boolean", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterSatLookbackMinutes != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.satLookbackMinutes", *params.FilterSatLookbackMinutes, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.FilterSatMaxAgeMinutes != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.satMaxAgeMinutes", *params.FilterSatMaxAgeMinutes, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVesselsPositionsRequest constructs an http.Request for the GetVesselsPositions method
func NewGetVesselsPositionsRequest(server string, params *GetVesselsPositionsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vessels/positions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.ids", params.FilterIds, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "filter.idType", params.FilterIdType, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.TimeFrom != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.from", *params.TimeFrom, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.TimeTo != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "time.to", *params.TimeTo, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationLimit != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.limit", *params.PaginationLimit, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "integer", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.PaginationNextToken != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "pagination.nextToken", *params.PaginationNextToken, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {

	// GetEmissionsWithResponse List emissions data
	//
	// Retrieves emissions data with optional filtering by reporting period.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /emissions (the `GetEmissions` operationId).
	GetEmissionsWithResponse(ctx context.Context, params *GetEmissionsParams, reqEditors ...RequestEditorFn) (*GetEmissionsResponse, error)

	// GetLocationDgpsBoundingBoxWithResponse Get DGPS Stations within a bounding box
	//
	// Retrieves Dgps stations within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/dgps/bounding-box (the `GetLocationDgpsBoundingBox` operationId).
	GetLocationDgpsBoundingBoxWithResponse(ctx context.Context, params *GetLocationDgpsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationDgpsBoundingBoxResponse, error)

	// GetLocationDgpsRadiusWithResponse Get DGPS Stations within a radius
	//
	// Retrieves Dgps stations within a specified radius of a given point.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/dgps/radius (the `GetLocationDgpsRadius` operationId).
	GetLocationDgpsRadiusWithResponse(ctx context.Context, params *GetLocationDgpsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationDgpsRadiusResponse, error)

	// GetLocationLightaidsBoundingBoxWithResponse Get Light Aids to Navigation within a bounding box
	//
	// Retrieves Light Aids to Navigation within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/lightaids/bounding-box (the `GetLocationLightaidsBoundingBox` operationId).
	GetLocationLightaidsBoundingBoxWithResponse(ctx context.Context, params *GetLocationLightaidsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationLightaidsBoundingBoxResponse, error)

	// GetLocationLightaidsRadiusWithResponse Get Light Aids to Navigation within a radius
	//
	// Retrieves Light Aids to Navigation within a specified radius of a given point.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/lightaids/radius (the `GetLocationLightaidsRadius` operationId).
	GetLocationLightaidsRadiusWithResponse(ctx context.Context, params *GetLocationLightaidsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationLightaidsRadiusResponse, error)

	// GetLocationModuBoundingBoxWithResponse Get Mobile Offshore Drilling Units within a bounding box
	//
	// Retrieves Mobile Offshore Drilling Units within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/modu/bounding-box (the `GetLocationModuBoundingBox` operationId).
	GetLocationModuBoundingBoxWithResponse(ctx context.Context, params *GetLocationModuBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationModuBoundingBoxResponse, error)

	// GetLocationModuRadiusWithResponse Get Mobile Offshore Drilling Units within a radius
	//
	// Retrieves Mobile Offshore Drilling Units within a specified radius of a given point.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/modu/radius (the `GetLocationModuRadius` operationId).
	GetLocationModuRadiusWithResponse(ctx context.Context, params *GetLocationModuRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationModuRadiusResponse, error)

	// GetLocationPortsBoundingBoxWithResponse Get Ports within a bounding box
	//
	// Retrieves Ports within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/ports/bounding-box (the `GetLocationPortsBoundingBox` operationId).
	GetLocationPortsBoundingBoxWithResponse(ctx context.Context, params *GetLocationPortsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationPortsBoundingBoxResponse, error)

	// GetLocationPortsRadiusWithResponse Get Ports within a radius
	//
	// Retrieves Ports within a specified radius of a given point.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/ports/radius (the `GetLocationPortsRadius` operationId).
	GetLocationPortsRadiusWithResponse(ctx context.Context, params *GetLocationPortsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationPortsRadiusResponse, error)

	// GetLocationRadiobeaconsBoundingBoxWithResponse Get Radio Beacons within a bounding box
	//
	// Retrieves Radio Beacons within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/radiobeacons/bounding-box (the `GetLocationRadiobeaconsBoundingBox` operationId).
	GetLocationRadiobeaconsBoundingBoxWithResponse(ctx context.Context, params *GetLocationRadiobeaconsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationRadiobeaconsBoundingBoxResponse, error)

	// GetLocationRadiobeaconsRadiusWithResponse Get Radio Beacons within a radius
	//
	// Retrieves Radio Beacons within a specified radius of a given point.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/radiobeacons/radius (the `GetLocationRadiobeaconsRadius` operationId).
	GetLocationRadiobeaconsRadiusWithResponse(ctx context.Context, params *GetLocationRadiobeaconsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationRadiobeaconsRadiusResponse, error)

	// GetLocationVesselsBoundingBoxWithResponse Get vessels within a bounding box
	//
	// Retrieves vessels within a specified bounding box and time window. Max bbox span: |dLat| + |dLon| ≤ 4 degrees. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/vessels/bounding-box (the `GetLocationVesselsBoundingBox` operationId).
	GetLocationVesselsBoundingBoxWithResponse(ctx context.Context, params *GetLocationVesselsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationVesselsBoundingBoxResponse, error)

	// GetLocationVesselsRadiusWithResponse Get vessels within a radius
	//
	// Retrieves vessels within a specified radius of a given point and a given time range. Max radius: 100 km. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /location/vessels/radius (the `GetLocationVesselsRadius` operationId).
	GetLocationVesselsRadiusWithResponse(ctx context.Context, params *GetLocationVesselsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationVesselsRadiusResponse, error)

	// GetPortUnlocodeWithResponse Get Port by UNLOCODE
	//
	// Retrieves port details by its UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /port/{unlocode} (the `GetPortUnlocode` operationId).
	GetPortUnlocodeWithResponse(ctx context.Context, unlocode string, reqEditors ...RequestEditorFn) (*GetPortUnlocodeResponse, error)

	// GetPortUnlocodeInboundWithResponse Get Inbound Vessels for Port
	//
	// Retrieves vessels heading to a specific port within an ETA window.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /port/{unlocode}/inbound (the `GetPortUnlocodeInbound` operationId).
	GetPortUnlocodeInboundWithResponse(ctx context.Context, unlocode string, params *GetPortUnlocodeInboundParams, reqEditors ...RequestEditorFn) (*GetPortUnlocodeInboundResponse, error)

	// GetPorteventsWithResponse Get Port Events in a time range
	//
	// Get Port Events, such as Arrivals and Departures, in a time range with optional filters
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /portevents (the `GetPortevents` operationId).
	GetPorteventsWithResponse(ctx context.Context, params *GetPorteventsParams, reqEditors ...RequestEditorFn) (*GetPorteventsResponse, error)

	// GetPorteventsPortUnlocodeWithResponse Get Port Events by port UNLOCODE
	//
	// Get Port Events, such as Arrivals and Departures by supplying the port UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /portevents/port/{unlocode} (the `GetPorteventsPortUnlocode` operationId).
	GetPorteventsPortUnlocodeWithResponse(ctx context.Context, unlocode string, params *GetPorteventsPortUnlocodeParams, reqEditors ...RequestEditorFn) (*GetPorteventsPortUnlocodeResponse, error)

	// GetPorteventsPortsWithResponse Get Port Events by port name
	//
	// Get Port Events, such as Arrivals and Departures by supplying the port name
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /portevents/ports (the `GetPorteventsPorts` operationId).
	GetPorteventsPortsWithResponse(ctx context.Context, params *GetPorteventsPortsParams, reqEditors ...RequestEditorFn) (*GetPorteventsPortsResponse, error)

	// GetPorteventsVesselIdWithResponse Get Port Events by Vessel ID
	//
	// Get all port events (arrivals and departures) for a vessel identified by MMSI or IMO, with optional filtering by event type and time range
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /portevents/vessel/{id} (the `GetPorteventsVesselId` operationId).
	GetPorteventsVesselIdWithResponse(ctx context.Context, id string, params *GetPorteventsVesselIdParams, reqEditors ...RequestEditorFn) (*GetPorteventsVesselIdResponse, error)

	// GetPorteventsVesselIdLastWithResponse Get Last Port Event by ID
	//
	// Get the most recent port event (arrival or departure) for a vessel identified by MMSI or IMO.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /portevents/vessel/{id}/last (the `GetPorteventsVesselIdLast` operationId).
	GetPorteventsVesselIdLastWithResponse(ctx context.Context, id string, params *GetPorteventsVesselIdLastParams, reqEditors ...RequestEditorFn) (*GetPorteventsVesselIdLastResponse, error)

	// GetPorteventsVesselsWithResponse Get Port Events by vessel name
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /portevents/vessels (the `GetPorteventsVessels` operationId).
	GetPorteventsVesselsWithResponse(ctx context.Context, params *GetPorteventsVesselsParams, reqEditors ...RequestEditorFn) (*GetPorteventsVesselsResponse, error)

	// GetSearchDgpsWithResponse Search for DGPS Stations
	//
	// Retrieves a list of DGPS stations for the given query parameters.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /search/dgps (the `GetSearchDgps` operationId).
	GetSearchDgpsWithResponse(ctx context.Context, params *GetSearchDgpsParams, reqEditors ...RequestEditorFn) (*GetSearchDgpsResponse, error)

	// GetSearchLightaidsWithResponse Search for Light Aids to Navigation
	//
	// Retrieves a list of Light Aids to Navigation for the given query parameters.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /search/lightaids (the `GetSearchLightaids` operationId).
	GetSearchLightaidsWithResponse(ctx context.Context, params *GetSearchLightaidsParams, reqEditors ...RequestEditorFn) (*GetSearchLightaidsResponse, error)

	// GetSearchModusWithResponse Search for MODUs
	//
	// Retrieves a list of MODUs for the given query parameters.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /search/modus (the `GetSearchModus` operationId).
	GetSearchModusWithResponse(ctx context.Context, params *GetSearchModusParams, reqEditors ...RequestEditorFn) (*GetSearchModusResponse, error)

	// GetSearchPortsWithResponse Search for Ports
	//
	// Retrieves a list of ports matching the given filters. At least one filter parameter is required.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /search/ports (the `GetSearchPorts` operationId).
	GetSearchPortsWithResponse(ctx context.Context, params *GetSearchPortsParams, reqEditors ...RequestEditorFn) (*GetSearchPortsResponse, error)

	// GetSearchRadiobeaconsWithResponse Search for Radio Beacons
	//
	// Retrieves a list of Radio Beacons for the given query parameters.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /search/radiobeacons (the `GetSearchRadiobeacons` operationId).
	GetSearchRadiobeaconsWithResponse(ctx context.Context, params *GetSearchRadiobeaconsParams, reqEditors ...RequestEditorFn) (*GetSearchRadiobeaconsResponse, error)

	// GetSearchVesselsWithResponse Search for Vessels
	//
	// Retrieves a list of vessels matching the given filters. At least one filter parameter (or the unified `q` parameter) is required.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /search/vessels (the `GetSearchVessels` operationId).
	GetSearchVesselsWithResponse(ctx context.Context, params *GetSearchVesselsParams, reqEditors ...RequestEditorFn) (*GetSearchVesselsResponse, error)

	// GetVesselIdWithResponse Get vessel information by MMSI or IMO
	//
	// Retrieves static vessel data including name, type, dimensions, and registration information for a vessel identified by its MMSI or IMO number
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /vessel/{id} (the `GetVesselId` operationId).
	GetVesselIdWithResponse(ctx context.Context, id string, params *GetVesselIdParams, reqEditors ...RequestEditorFn) (*GetVesselIdResponse, error)

	// GetVesselIdCasualtiesWithResponse Get marine casualties involving a vessel
	//
	// Retrieves marine casualty records involving the specified vessel.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /vessel/{id}/casualties (the `GetVesselIdCasualties` operationId).
	GetVesselIdCasualtiesWithResponse(ctx context.Context, id string, params *GetVesselIdCasualtiesParams, reqEditors ...RequestEditorFn) (*GetVesselIdCasualtiesResponse, error)

	// GetVesselIdEmissionsWithResponse Get emissions data for a vessel
	//
	// Retrieves emissions reports for the specified vessel including CO2 emissions, fuel consumption, and efficiency metrics
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /vessel/{id}/emissions (the `GetVesselIdEmissions` operationId).
	GetVesselIdEmissionsWithResponse(ctx context.Context, id string, params *GetVesselIdEmissionsParams, reqEditors ...RequestEditorFn) (*GetVesselIdEmissionsResponse, error)

	// GetVesselIdEtaWithResponse Get vessel latest ETA
	//
	// Retrieves the most recent Estimated Time of Arrival (ETA) reported by the vessel.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /vessel/{id}/eta (the `GetVesselIdEta` operationId).
	GetVesselIdEtaWithResponse(ctx context.Context, id string, params *GetVesselIdEtaParams, reqEditors ...RequestEditorFn) (*GetVesselIdEtaResponse, error)

	// GetVesselIdPositionWithResponse Get last known vessel position
	//
	// Retrieves the most recent AIS position report for a vessel, including coordinates, vessel identifiers, and timestamps. The filter.sat satellite fallback is deprecated and withdrawn in its current form: it works for existing credit holders until the cutoff date, after which it is accepted and ignored. The response also includes the vessel's particulars (vessel), the latest destination and ETA it reported (vesselEta), and positionCheck, which compares the position with the vessel's previous position report: time and distance since then, implied speed and course, and whether the identifiers match.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /vessel/{id}/position (the `GetVesselIdPosition` operationId).
	GetVesselIdPositionWithResponse(ctx context.Context, id string, params *GetVesselIdPositionParams, reqEditors ...RequestEditorFn) (*GetVesselIdPositionResponse, error)

	// GetVesselsPositionsWithResponse Get positions for multiple vessels
	//
	// Retrieves AIS position data for multiple vessels identified by MMSI or IMO numbers within a specified time range (defaults to past 2 hours). Provide multiple IDs either as a comma-separated list in one filter.ids param, or by repeating filter.ids; both forms (and a mix) are accepted.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /vessels/positions (the `GetVesselsPositions` operationId).
	GetVesselsPositionsWithResponse(ctx context.Context, params *GetVesselsPositionsParams, reqEditors ...RequestEditorFn) (*GetVesselsPositionsResponse, error)
}

type GetEmissionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselEmissionsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetEmissionsResponse) GetJSON200() *VesselEmissionsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetEmissionsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetEmissionsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetEmissionsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetEmissionsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetEmissionsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetEmissionsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetEmissionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetEmissionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetEmissionsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationDgpsBoundingBoxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *DGPSStationsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationDgpsBoundingBoxResponse) GetJSON200() *DGPSStationsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationDgpsBoundingBoxResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationDgpsBoundingBoxResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationDgpsBoundingBoxResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationDgpsBoundingBoxResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationDgpsBoundingBoxResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationDgpsBoundingBoxResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationDgpsBoundingBoxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationDgpsBoundingBoxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationDgpsBoundingBoxResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationDgpsRadiusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *DGPSStationsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationDgpsRadiusResponse) GetJSON200() *DGPSStationsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationDgpsRadiusResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationDgpsRadiusResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationDgpsRadiusResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationDgpsRadiusResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationDgpsRadiusResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationDgpsRadiusResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationDgpsRadiusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationDgpsRadiusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationDgpsRadiusResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationLightaidsBoundingBoxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *LightAidsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationLightaidsBoundingBoxResponse) GetJSON200() *LightAidsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationLightaidsBoundingBoxResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationLightaidsBoundingBoxResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationLightaidsBoundingBoxResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationLightaidsBoundingBoxResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationLightaidsBoundingBoxResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationLightaidsBoundingBoxResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationLightaidsBoundingBoxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationLightaidsBoundingBoxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationLightaidsBoundingBoxResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationLightaidsRadiusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *LightAidsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationLightaidsRadiusResponse) GetJSON200() *LightAidsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationLightaidsRadiusResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationLightaidsRadiusResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationLightaidsRadiusResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationLightaidsRadiusResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationLightaidsRadiusResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationLightaidsRadiusResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationLightaidsRadiusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationLightaidsRadiusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationLightaidsRadiusResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationModuBoundingBoxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *MODUsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationModuBoundingBoxResponse) GetJSON200() *MODUsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationModuBoundingBoxResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationModuBoundingBoxResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationModuBoundingBoxResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationModuBoundingBoxResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationModuBoundingBoxResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationModuBoundingBoxResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationModuBoundingBoxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationModuBoundingBoxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationModuBoundingBoxResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationModuRadiusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *MODUsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationModuRadiusResponse) GetJSON200() *MODUsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationModuRadiusResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationModuRadiusResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationModuRadiusResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationModuRadiusResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationModuRadiusResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationModuRadiusResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationModuRadiusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationModuRadiusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationModuRadiusResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationPortsBoundingBoxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationPortsBoundingBoxResponse) GetJSON200() *PortsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationPortsBoundingBoxResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationPortsBoundingBoxResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationPortsBoundingBoxResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationPortsBoundingBoxResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationPortsBoundingBoxResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationPortsBoundingBoxResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationPortsBoundingBoxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationPortsBoundingBoxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationPortsBoundingBoxResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationPortsRadiusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationPortsRadiusResponse) GetJSON200() *PortsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationPortsRadiusResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationPortsRadiusResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationPortsRadiusResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationPortsRadiusResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationPortsRadiusResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationPortsRadiusResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationPortsRadiusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationPortsRadiusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationPortsRadiusResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationRadiobeaconsBoundingBoxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *RadioBeaconsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationRadiobeaconsBoundingBoxResponse) GetJSON200() *RadioBeaconsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationRadiobeaconsBoundingBoxResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationRadiobeaconsBoundingBoxResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationRadiobeaconsBoundingBoxResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationRadiobeaconsBoundingBoxResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationRadiobeaconsBoundingBoxResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationRadiobeaconsBoundingBoxResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationRadiobeaconsBoundingBoxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationRadiobeaconsBoundingBoxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationRadiobeaconsBoundingBoxResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationRadiobeaconsRadiusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *RadioBeaconsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationRadiobeaconsRadiusResponse) GetJSON200() *RadioBeaconsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationRadiobeaconsRadiusResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationRadiobeaconsRadiusResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationRadiobeaconsRadiusResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationRadiobeaconsRadiusResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationRadiobeaconsRadiusResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationRadiobeaconsRadiusResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationRadiobeaconsRadiusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationRadiobeaconsRadiusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationRadiobeaconsRadiusResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationVesselsBoundingBoxResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationVesselsBoundingBoxResponse) GetJSON200() *VesselsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationVesselsBoundingBoxResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationVesselsBoundingBoxResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationVesselsBoundingBoxResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationVesselsBoundingBoxResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationVesselsBoundingBoxResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationVesselsBoundingBoxResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationVesselsBoundingBoxResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationVesselsBoundingBoxResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationVesselsBoundingBoxResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetLocationVesselsRadiusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselsWithinLocationResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetLocationVesselsRadiusResponse) GetJSON200() *VesselsWithinLocationResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetLocationVesselsRadiusResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetLocationVesselsRadiusResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetLocationVesselsRadiusResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetLocationVesselsRadiusResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetLocationVesselsRadiusResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetLocationVesselsRadiusResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetLocationVesselsRadiusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLocationVesselsRadiusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetLocationVesselsRadiusResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPortUnlocodeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPortUnlocodeResponse) GetJSON200() *PortResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPortUnlocodeResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPortUnlocodeResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPortUnlocodeResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetPortUnlocodeResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPortUnlocodeResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPortUnlocodeResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetPortUnlocodeResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPortUnlocodeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPortUnlocodeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPortUnlocodeResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPortUnlocodeInboundResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortInboundResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
	// JSON503 the response for an HTTP 503 `application/json` response
	JSON503 *ServiceUnavailableErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON200() *PortInboundResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetJSON503 returns the response for an HTTP 503 `application/json` response
func (r GetPortUnlocodeInboundResponse) GetJSON503() *ServiceUnavailableErrorResponse {
	return r.JSON503
}

// GetBody returns the raw response body bytes
func (r GetPortUnlocodeInboundResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPortUnlocodeInboundResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPortUnlocodeInboundResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPortUnlocodeInboundResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPorteventsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortEventsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPorteventsResponse) GetJSON200() *PortEventsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPorteventsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPorteventsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPorteventsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPorteventsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPorteventsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetPorteventsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPorteventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPorteventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPorteventsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPorteventsPortUnlocodeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortEventsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPorteventsPortUnlocodeResponse) GetJSON200() *PortEventsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPorteventsPortUnlocodeResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPorteventsPortUnlocodeResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPorteventsPortUnlocodeResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetPorteventsPortUnlocodeResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPorteventsPortUnlocodeResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPorteventsPortUnlocodeResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetPorteventsPortUnlocodeResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPorteventsPortUnlocodeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPorteventsPortUnlocodeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPorteventsPortUnlocodeResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPorteventsPortsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortEventsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPorteventsPortsResponse) GetJSON200() *PortEventsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPorteventsPortsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPorteventsPortsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPorteventsPortsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPorteventsPortsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPorteventsPortsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetPorteventsPortsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPorteventsPortsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPorteventsPortsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPorteventsPortsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPorteventsVesselIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortEventsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPorteventsVesselIdResponse) GetJSON200() *PortEventsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPorteventsVesselIdResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPorteventsVesselIdResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPorteventsVesselIdResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetPorteventsVesselIdResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPorteventsVesselIdResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPorteventsVesselIdResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetPorteventsVesselIdResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPorteventsVesselIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPorteventsVesselIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPorteventsVesselIdResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPorteventsVesselIdLastResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortEventResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPorteventsVesselIdLastResponse) GetJSON200() *PortEventResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPorteventsVesselIdLastResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPorteventsVesselIdLastResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPorteventsVesselIdLastResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetPorteventsVesselIdLastResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPorteventsVesselIdLastResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPorteventsVesselIdLastResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetPorteventsVesselIdLastResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPorteventsVesselIdLastResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPorteventsVesselIdLastResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPorteventsVesselIdLastResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetPorteventsVesselsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *PortEventsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetPorteventsVesselsResponse) GetJSON200() *PortEventsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetPorteventsVesselsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetPorteventsVesselsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetPorteventsVesselsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetPorteventsVesselsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetPorteventsVesselsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetPorteventsVesselsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetPorteventsVesselsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPorteventsVesselsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetPorteventsVesselsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetSearchDgpsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *FindDGPSStationsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetSearchDgpsResponse) GetJSON200() *FindDGPSStationsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetSearchDgpsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetSearchDgpsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetSearchDgpsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetSearchDgpsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetSearchDgpsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetSearchDgpsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetSearchDgpsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSearchDgpsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetSearchDgpsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetSearchLightaidsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *FindLightAidsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetSearchLightaidsResponse) GetJSON200() *FindLightAidsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetSearchLightaidsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetSearchLightaidsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetSearchLightaidsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetSearchLightaidsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetSearchLightaidsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetSearchLightaidsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetSearchLightaidsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSearchLightaidsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetSearchLightaidsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetSearchModusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *FindMODUsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetSearchModusResponse) GetJSON200() *FindMODUsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetSearchModusResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetSearchModusResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetSearchModusResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetSearchModusResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetSearchModusResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetSearchModusResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetSearchModusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSearchModusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetSearchModusResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetSearchPortsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *FindPortsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetSearchPortsResponse) GetJSON200() *FindPortsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetSearchPortsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetSearchPortsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetSearchPortsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetSearchPortsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetSearchPortsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetSearchPortsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetSearchPortsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSearchPortsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetSearchPortsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetSearchRadiobeaconsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *FindRadioBeaconsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetSearchRadiobeaconsResponse) GetJSON200() *FindRadioBeaconsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetSearchRadiobeaconsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetSearchRadiobeaconsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetSearchRadiobeaconsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetSearchRadiobeaconsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetSearchRadiobeaconsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetSearchRadiobeaconsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetSearchRadiobeaconsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSearchRadiobeaconsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetSearchRadiobeaconsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetSearchVesselsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *FindVesselsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetSearchVesselsResponse) GetJSON200() *FindVesselsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetSearchVesselsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetSearchVesselsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetSearchVesselsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetSearchVesselsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetSearchVesselsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetSearchVesselsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetSearchVesselsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSearchVesselsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetSearchVesselsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetVesselIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetVesselIdResponse) GetJSON200() *VesselResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetVesselIdResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetVesselIdResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetVesselIdResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetVesselIdResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetVesselIdResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetVesselIdResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetVesselIdResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetVesselIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVesselIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetVesselIdResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetVesselIdCasualtiesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *MarineCasualtiesResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetVesselIdCasualtiesResponse) GetJSON200() *MarineCasualtiesResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetVesselIdCasualtiesResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetVesselIdCasualtiesResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetVesselIdCasualtiesResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetVesselIdCasualtiesResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetVesselIdCasualtiesResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetVesselIdCasualtiesResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetVesselIdCasualtiesResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetVesselIdCasualtiesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVesselIdCasualtiesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetVesselIdCasualtiesResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetVesselIdEmissionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselEmissionsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetVesselIdEmissionsResponse) GetJSON200() *VesselEmissionsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetVesselIdEmissionsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetVesselIdEmissionsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetVesselIdEmissionsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetVesselIdEmissionsResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetVesselIdEmissionsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetVesselIdEmissionsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetVesselIdEmissionsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetVesselIdEmissionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVesselIdEmissionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetVesselIdEmissionsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetVesselIdEtaResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselETAResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
	// JSON503 the response for an HTTP 503 `application/json` response
	JSON503 *ServiceUnavailableErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON200() *VesselETAResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetJSON503 returns the response for an HTTP 503 `application/json` response
func (r GetVesselIdEtaResponse) GetJSON503() *ServiceUnavailableErrorResponse {
	return r.JSON503
}

// GetBody returns the raw response body bytes
func (r GetVesselIdEtaResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetVesselIdEtaResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVesselIdEtaResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetVesselIdEtaResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// GetVesselIdPositionResponse200Headers the declared response headers of an HTTP 200 response for GetVesselIdPosition
type GetVesselIdPositionResponse200Headers struct {
	XDataSource                *string
	XSatelliteCreditCharged    *string
	XSatelliteCreditsRemaining *string
	XSatelliteLookbackMinutes  *string
	XSatelliteMaxAgeMinutes    *string
	XSatelliteStatus           *string
}

type GetVesselIdPositionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselPositionResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON402 the response for an HTTP 402 `application/json` response
	JSON402 *PaymentRequiredErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFoundErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *GetVesselIdPositionResponse200Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON200() *VesselPositionResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON402 returns the response for an HTTP 402 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON402() *PaymentRequiredErrorResponse {
	return r.JSON402
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON404() *NotFoundErrorResponse {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetVesselIdPositionResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetVesselIdPositionResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetVesselIdPositionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVesselIdPositionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetVesselIdPositionResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

type GetVesselsPositionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *VesselPositionsResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequestErrorResponse
	// JSON401 the response for an HTTP 401 `application/json` response
	JSON401 *AuthenticationErrorResponse
	// JSON403 the response for an HTTP 403 `application/json` response
	JSON403 *ForbiddenErrorResponse
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *RateLimitErrorResponse
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalServerErrorResponse
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetVesselsPositionsResponse) GetJSON200() *VesselPositionsResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetVesselsPositionsResponse) GetJSON400() *BadRequestErrorResponse {
	return r.JSON400
}

// GetJSON401 returns the response for an HTTP 401 `application/json` response
func (r GetVesselsPositionsResponse) GetJSON401() *AuthenticationErrorResponse {
	return r.JSON401
}

// GetJSON403 returns the response for an HTTP 403 `application/json` response
func (r GetVesselsPositionsResponse) GetJSON403() *ForbiddenErrorResponse {
	return r.JSON403
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetVesselsPositionsResponse) GetJSON429() *RateLimitErrorResponse {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r GetVesselsPositionsResponse) GetJSON500() *InternalServerErrorResponse {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r GetVesselsPositionsResponse) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetVesselsPositionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVesselsPositionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetVesselsPositionsResponse) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// GetEmissionsWithResponse List emissions data
//
// Retrieves emissions data with optional filtering by reporting period.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /emissions (the `GetEmissions` operationId).
func (c *ClientWithResponses) GetEmissionsWithResponse(ctx context.Context, params *GetEmissionsParams, reqEditors ...RequestEditorFn) (*GetEmissionsResponse, error) {
	rsp, err := c.GetEmissions(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetEmissionsResponse(rsp)
}

// GetLocationDgpsBoundingBoxWithResponse Get DGPS Stations within a bounding box
//
// Retrieves Dgps stations within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/dgps/bounding-box (the `GetLocationDgpsBoundingBox` operationId).
func (c *ClientWithResponses) GetLocationDgpsBoundingBoxWithResponse(ctx context.Context, params *GetLocationDgpsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationDgpsBoundingBoxResponse, error) {
	rsp, err := c.GetLocationDgpsBoundingBox(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationDgpsBoundingBoxResponse(rsp)
}

// GetLocationDgpsRadiusWithResponse Get DGPS Stations within a radius
//
// Retrieves Dgps stations within a specified radius of a given point.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/dgps/radius (the `GetLocationDgpsRadius` operationId).
func (c *ClientWithResponses) GetLocationDgpsRadiusWithResponse(ctx context.Context, params *GetLocationDgpsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationDgpsRadiusResponse, error) {
	rsp, err := c.GetLocationDgpsRadius(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationDgpsRadiusResponse(rsp)
}

// GetLocationLightaidsBoundingBoxWithResponse Get Light Aids to Navigation within a bounding box
//
// Retrieves Light Aids to Navigation within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/lightaids/bounding-box (the `GetLocationLightaidsBoundingBox` operationId).
func (c *ClientWithResponses) GetLocationLightaidsBoundingBoxWithResponse(ctx context.Context, params *GetLocationLightaidsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationLightaidsBoundingBoxResponse, error) {
	rsp, err := c.GetLocationLightaidsBoundingBox(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationLightaidsBoundingBoxResponse(rsp)
}

// GetLocationLightaidsRadiusWithResponse Get Light Aids to Navigation within a radius
//
// Retrieves Light Aids to Navigation within a specified radius of a given point.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/lightaids/radius (the `GetLocationLightaidsRadius` operationId).
func (c *ClientWithResponses) GetLocationLightaidsRadiusWithResponse(ctx context.Context, params *GetLocationLightaidsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationLightaidsRadiusResponse, error) {
	rsp, err := c.GetLocationLightaidsRadius(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationLightaidsRadiusResponse(rsp)
}

// GetLocationModuBoundingBoxWithResponse Get Mobile Offshore Drilling Units within a bounding box
//
// Retrieves Mobile Offshore Drilling Units within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/modu/bounding-box (the `GetLocationModuBoundingBox` operationId).
func (c *ClientWithResponses) GetLocationModuBoundingBoxWithResponse(ctx context.Context, params *GetLocationModuBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationModuBoundingBoxResponse, error) {
	rsp, err := c.GetLocationModuBoundingBox(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationModuBoundingBoxResponse(rsp)
}

// GetLocationModuRadiusWithResponse Get Mobile Offshore Drilling Units within a radius
//
// Retrieves Mobile Offshore Drilling Units within a specified radius of a given point.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/modu/radius (the `GetLocationModuRadius` operationId).
func (c *ClientWithResponses) GetLocationModuRadiusWithResponse(ctx context.Context, params *GetLocationModuRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationModuRadiusResponse, error) {
	rsp, err := c.GetLocationModuRadius(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationModuRadiusResponse(rsp)
}

// GetLocationPortsBoundingBoxWithResponse Get Ports within a bounding box
//
// Retrieves Ports within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/ports/bounding-box (the `GetLocationPortsBoundingBox` operationId).
func (c *ClientWithResponses) GetLocationPortsBoundingBoxWithResponse(ctx context.Context, params *GetLocationPortsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationPortsBoundingBoxResponse, error) {
	rsp, err := c.GetLocationPortsBoundingBox(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationPortsBoundingBoxResponse(rsp)
}

// GetLocationPortsRadiusWithResponse Get Ports within a radius
//
// Retrieves Ports within a specified radius of a given point.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/ports/radius (the `GetLocationPortsRadius` operationId).
func (c *ClientWithResponses) GetLocationPortsRadiusWithResponse(ctx context.Context, params *GetLocationPortsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationPortsRadiusResponse, error) {
	rsp, err := c.GetLocationPortsRadius(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationPortsRadiusResponse(rsp)
}

// GetLocationRadiobeaconsBoundingBoxWithResponse Get Radio Beacons within a bounding box
//
// Retrieves Radio Beacons within a specified bounding box. Max span: |dLat| + |dLon| ≤ 4 degrees
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/radiobeacons/bounding-box (the `GetLocationRadiobeaconsBoundingBox` operationId).
func (c *ClientWithResponses) GetLocationRadiobeaconsBoundingBoxWithResponse(ctx context.Context, params *GetLocationRadiobeaconsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationRadiobeaconsBoundingBoxResponse, error) {
	rsp, err := c.GetLocationRadiobeaconsBoundingBox(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationRadiobeaconsBoundingBoxResponse(rsp)
}

// GetLocationRadiobeaconsRadiusWithResponse Get Radio Beacons within a radius
//
// Retrieves Radio Beacons within a specified radius of a given point.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/radiobeacons/radius (the `GetLocationRadiobeaconsRadius` operationId).
func (c *ClientWithResponses) GetLocationRadiobeaconsRadiusWithResponse(ctx context.Context, params *GetLocationRadiobeaconsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationRadiobeaconsRadiusResponse, error) {
	rsp, err := c.GetLocationRadiobeaconsRadius(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationRadiobeaconsRadiusResponse(rsp)
}

// GetLocationVesselsBoundingBoxWithResponse Get vessels within a bounding box
//
// Retrieves vessels within a specified bounding box and time window. Max bbox span: |dLat| + |dLon| ≤ 4 degrees. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/vessels/bounding-box (the `GetLocationVesselsBoundingBox` operationId).
func (c *ClientWithResponses) GetLocationVesselsBoundingBoxWithResponse(ctx context.Context, params *GetLocationVesselsBoundingBoxParams, reqEditors ...RequestEditorFn) (*GetLocationVesselsBoundingBoxResponse, error) {
	rsp, err := c.GetLocationVesselsBoundingBox(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationVesselsBoundingBoxResponse(rsp)
}

// GetLocationVesselsRadiusWithResponse Get vessels within a radius
//
// Retrieves vessels within a specified radius of a given point and a given time range. Max radius: 100 km. Max time window (time.to - time.from): 4 hours; for longer ranges issue sequential 4h calls and stitch results client-side using half-open intervals (one slice's time.to = T, next slice's time.from = T + 1ms) to avoid boundary duplicates. A `nextToken` is only valid when reused with the same `time.from`/`time.to`; the API does not reject mismatched bounds and silently returns rows from the wrong slice.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /location/vessels/radius (the `GetLocationVesselsRadius` operationId).
func (c *ClientWithResponses) GetLocationVesselsRadiusWithResponse(ctx context.Context, params *GetLocationVesselsRadiusParams, reqEditors ...RequestEditorFn) (*GetLocationVesselsRadiusResponse, error) {
	rsp, err := c.GetLocationVesselsRadius(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLocationVesselsRadiusResponse(rsp)
}

// GetPortUnlocodeWithResponse Get Port by UNLOCODE
//
// Retrieves port details by its UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /port/{unlocode} (the `GetPortUnlocode` operationId).
func (c *ClientWithResponses) GetPortUnlocodeWithResponse(ctx context.Context, unlocode string, reqEditors ...RequestEditorFn) (*GetPortUnlocodeResponse, error) {
	rsp, err := c.GetPortUnlocode(ctx, unlocode, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPortUnlocodeResponse(rsp)
}

// GetPortUnlocodeInboundWithResponse Get Inbound Vessels for Port
//
// Retrieves vessels heading to a specific port within an ETA window.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /port/{unlocode}/inbound (the `GetPortUnlocodeInbound` operationId).
func (c *ClientWithResponses) GetPortUnlocodeInboundWithResponse(ctx context.Context, unlocode string, params *GetPortUnlocodeInboundParams, reqEditors ...RequestEditorFn) (*GetPortUnlocodeInboundResponse, error) {
	rsp, err := c.GetPortUnlocodeInbound(ctx, unlocode, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPortUnlocodeInboundResponse(rsp)
}

// GetPorteventsWithResponse Get Port Events in a time range
//
// # Get Port Events, such as Arrivals and Departures, in a time range with optional filters
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /portevents (the `GetPortevents` operationId).
func (c *ClientWithResponses) GetPorteventsWithResponse(ctx context.Context, params *GetPorteventsParams, reqEditors ...RequestEditorFn) (*GetPorteventsResponse, error) {
	rsp, err := c.GetPortevents(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPorteventsResponse(rsp)
}

// GetPorteventsPortUnlocodeWithResponse Get Port Events by port UNLOCODE
//
// Get Port Events, such as Arrivals and Departures by supplying the port UN/LOCODE (e.g., NLRTM for Rotterdam, SGSIN for Singapore)
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /portevents/port/{unlocode} (the `GetPorteventsPortUnlocode` operationId).
func (c *ClientWithResponses) GetPorteventsPortUnlocodeWithResponse(ctx context.Context, unlocode string, params *GetPorteventsPortUnlocodeParams, reqEditors ...RequestEditorFn) (*GetPorteventsPortUnlocodeResponse, error) {
	rsp, err := c.GetPorteventsPortUnlocode(ctx, unlocode, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPorteventsPortUnlocodeResponse(rsp)
}

// GetPorteventsPortsWithResponse Get Port Events by port name
//
// # Get Port Events, such as Arrivals and Departures by supplying the port name
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /portevents/ports (the `GetPorteventsPorts` operationId).
func (c *ClientWithResponses) GetPorteventsPortsWithResponse(ctx context.Context, params *GetPorteventsPortsParams, reqEditors ...RequestEditorFn) (*GetPorteventsPortsResponse, error) {
	rsp, err := c.GetPorteventsPorts(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPorteventsPortsResponse(rsp)
}

// GetPorteventsVesselIdWithResponse Get Port Events by Vessel ID
//
// # Get all port events (arrivals and departures) for a vessel identified by MMSI or IMO, with optional filtering by event type and time range
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /portevents/vessel/{id} (the `GetPorteventsVesselId` operationId).
func (c *ClientWithResponses) GetPorteventsVesselIdWithResponse(ctx context.Context, id string, params *GetPorteventsVesselIdParams, reqEditors ...RequestEditorFn) (*GetPorteventsVesselIdResponse, error) {
	rsp, err := c.GetPorteventsVesselId(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPorteventsVesselIdResponse(rsp)
}

// GetPorteventsVesselIdLastWithResponse Get Last Port Event by ID
//
// Get the most recent port event (arrival or departure) for a vessel identified by MMSI or IMO.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /portevents/vessel/{id}/last (the `GetPorteventsVesselIdLast` operationId).
func (c *ClientWithResponses) GetPorteventsVesselIdLastWithResponse(ctx context.Context, id string, params *GetPorteventsVesselIdLastParams, reqEditors ...RequestEditorFn) (*GetPorteventsVesselIdLastResponse, error) {
	rsp, err := c.GetPorteventsVesselIdLast(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPorteventsVesselIdLastResponse(rsp)
}

// GetPorteventsVesselsWithResponse Get Port Events by vessel name
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /portevents/vessels (the `GetPorteventsVessels` operationId).
func (c *ClientWithResponses) GetPorteventsVesselsWithResponse(ctx context.Context, params *GetPorteventsVesselsParams, reqEditors ...RequestEditorFn) (*GetPorteventsVesselsResponse, error) {
	rsp, err := c.GetPorteventsVessels(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPorteventsVesselsResponse(rsp)
}

// GetSearchDgpsWithResponse Search for DGPS Stations
//
// Retrieves a list of DGPS stations for the given query parameters.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /search/dgps (the `GetSearchDgps` operationId).
func (c *ClientWithResponses) GetSearchDgpsWithResponse(ctx context.Context, params *GetSearchDgpsParams, reqEditors ...RequestEditorFn) (*GetSearchDgpsResponse, error) {
	rsp, err := c.GetSearchDgps(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSearchDgpsResponse(rsp)
}

// GetSearchLightaidsWithResponse Search for Light Aids to Navigation
//
// Retrieves a list of Light Aids to Navigation for the given query parameters.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /search/lightaids (the `GetSearchLightaids` operationId).
func (c *ClientWithResponses) GetSearchLightaidsWithResponse(ctx context.Context, params *GetSearchLightaidsParams, reqEditors ...RequestEditorFn) (*GetSearchLightaidsResponse, error) {
	rsp, err := c.GetSearchLightaids(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSearchLightaidsResponse(rsp)
}

// GetSearchModusWithResponse Search for MODUs
//
// Retrieves a list of MODUs for the given query parameters.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /search/modus (the `GetSearchModus` operationId).
func (c *ClientWithResponses) GetSearchModusWithResponse(ctx context.Context, params *GetSearchModusParams, reqEditors ...RequestEditorFn) (*GetSearchModusResponse, error) {
	rsp, err := c.GetSearchModus(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSearchModusResponse(rsp)
}

// GetSearchPortsWithResponse Search for Ports
//
// Retrieves a list of ports matching the given filters. At least one filter parameter is required.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /search/ports (the `GetSearchPorts` operationId).
func (c *ClientWithResponses) GetSearchPortsWithResponse(ctx context.Context, params *GetSearchPortsParams, reqEditors ...RequestEditorFn) (*GetSearchPortsResponse, error) {
	rsp, err := c.GetSearchPorts(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSearchPortsResponse(rsp)
}

// GetSearchRadiobeaconsWithResponse Search for Radio Beacons
//
// Retrieves a list of Radio Beacons for the given query parameters.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /search/radiobeacons (the `GetSearchRadiobeacons` operationId).
func (c *ClientWithResponses) GetSearchRadiobeaconsWithResponse(ctx context.Context, params *GetSearchRadiobeaconsParams, reqEditors ...RequestEditorFn) (*GetSearchRadiobeaconsResponse, error) {
	rsp, err := c.GetSearchRadiobeacons(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSearchRadiobeaconsResponse(rsp)
}

// GetSearchVesselsWithResponse Search for Vessels
//
// Retrieves a list of vessels matching the given filters. At least one filter parameter (or the unified `q` parameter) is required.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /search/vessels (the `GetSearchVessels` operationId).
func (c *ClientWithResponses) GetSearchVesselsWithResponse(ctx context.Context, params *GetSearchVesselsParams, reqEditors ...RequestEditorFn) (*GetSearchVesselsResponse, error) {
	rsp, err := c.GetSearchVessels(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSearchVesselsResponse(rsp)
}

// GetVesselIdWithResponse Get vessel information by MMSI or IMO
//
// # Retrieves static vessel data including name, type, dimensions, and registration information for a vessel identified by its MMSI or IMO number
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /vessel/{id} (the `GetVesselId` operationId).
func (c *ClientWithResponses) GetVesselIdWithResponse(ctx context.Context, id string, params *GetVesselIdParams, reqEditors ...RequestEditorFn) (*GetVesselIdResponse, error) {
	rsp, err := c.GetVesselId(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVesselIdResponse(rsp)
}

// GetVesselIdCasualtiesWithResponse Get marine casualties involving a vessel
//
// Retrieves marine casualty records involving the specified vessel.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /vessel/{id}/casualties (the `GetVesselIdCasualties` operationId).
func (c *ClientWithResponses) GetVesselIdCasualtiesWithResponse(ctx context.Context, id string, params *GetVesselIdCasualtiesParams, reqEditors ...RequestEditorFn) (*GetVesselIdCasualtiesResponse, error) {
	rsp, err := c.GetVesselIdCasualties(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVesselIdCasualtiesResponse(rsp)
}

// GetVesselIdEmissionsWithResponse Get emissions data for a vessel
//
// # Retrieves emissions reports for the specified vessel including CO2 emissions, fuel consumption, and efficiency metrics
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /vessel/{id}/emissions (the `GetVesselIdEmissions` operationId).
func (c *ClientWithResponses) GetVesselIdEmissionsWithResponse(ctx context.Context, id string, params *GetVesselIdEmissionsParams, reqEditors ...RequestEditorFn) (*GetVesselIdEmissionsResponse, error) {
	rsp, err := c.GetVesselIdEmissions(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVesselIdEmissionsResponse(rsp)
}

// GetVesselIdEtaWithResponse Get vessel latest ETA
//
// Retrieves the most recent Estimated Time of Arrival (ETA) reported by the vessel.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /vessel/{id}/eta (the `GetVesselIdEta` operationId).
func (c *ClientWithResponses) GetVesselIdEtaWithResponse(ctx context.Context, id string, params *GetVesselIdEtaParams, reqEditors ...RequestEditorFn) (*GetVesselIdEtaResponse, error) {
	rsp, err := c.GetVesselIdEta(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVesselIdEtaResponse(rsp)
}

// GetVesselIdPositionWithResponse Get last known vessel position
//
// Retrieves the most recent AIS position report for a vessel, including coordinates, vessel identifiers, and timestamps. The filter.sat satellite fallback is deprecated and withdrawn in its current form: it works for existing credit holders until the cutoff date, after which it is accepted and ignored. The response also includes the vessel's particulars (vessel), the latest destination and ETA it reported (vesselEta), and positionCheck, which compares the position with the vessel's previous position report: time and distance since then, implied speed and course, and whether the identifiers match.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /vessel/{id}/position (the `GetVesselIdPosition` operationId).
func (c *ClientWithResponses) GetVesselIdPositionWithResponse(ctx context.Context, id string, params *GetVesselIdPositionParams, reqEditors ...RequestEditorFn) (*GetVesselIdPositionResponse, error) {
	rsp, err := c.GetVesselIdPosition(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVesselIdPositionResponse(rsp)
}

// GetVesselsPositionsWithResponse Get positions for multiple vessels
//
// Retrieves AIS position data for multiple vessels identified by MMSI or IMO numbers within a specified time range (defaults to past 2 hours). Provide multiple IDs either as a comma-separated list in one filter.ids param, or by repeating filter.ids; both forms (and a mix) are accepted.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /vessels/positions (the `GetVesselsPositions` operationId).
func (c *ClientWithResponses) GetVesselsPositionsWithResponse(ctx context.Context, params *GetVesselsPositionsParams, reqEditors ...RequestEditorFn) (*GetVesselsPositionsResponse, error) {
	rsp, err := c.GetVesselsPositions(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVesselsPositionsResponse(rsp)
}

// ParseGetEmissionsResponse parses an HTTP response from a GetEmissionsWithResponse call
func ParseGetEmissionsResponse(rsp *http.Response) (*GetEmissionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetEmissionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselEmissionsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationDgpsBoundingBoxResponse parses an HTTP response from a GetLocationDgpsBoundingBoxWithResponse call
func ParseGetLocationDgpsBoundingBoxResponse(rsp *http.Response) (*GetLocationDgpsBoundingBoxResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationDgpsBoundingBoxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DGPSStationsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationDgpsRadiusResponse parses an HTTP response from a GetLocationDgpsRadiusWithResponse call
func ParseGetLocationDgpsRadiusResponse(rsp *http.Response) (*GetLocationDgpsRadiusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationDgpsRadiusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DGPSStationsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationLightaidsBoundingBoxResponse parses an HTTP response from a GetLocationLightaidsBoundingBoxWithResponse call
func ParseGetLocationLightaidsBoundingBoxResponse(rsp *http.Response) (*GetLocationLightaidsBoundingBoxResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationLightaidsBoundingBoxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LightAidsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationLightaidsRadiusResponse parses an HTTP response from a GetLocationLightaidsRadiusWithResponse call
func ParseGetLocationLightaidsRadiusResponse(rsp *http.Response) (*GetLocationLightaidsRadiusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationLightaidsRadiusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LightAidsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationModuBoundingBoxResponse parses an HTTP response from a GetLocationModuBoundingBoxWithResponse call
func ParseGetLocationModuBoundingBoxResponse(rsp *http.Response) (*GetLocationModuBoundingBoxResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationModuBoundingBoxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MODUsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationModuRadiusResponse parses an HTTP response from a GetLocationModuRadiusWithResponse call
func ParseGetLocationModuRadiusResponse(rsp *http.Response) (*GetLocationModuRadiusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationModuRadiusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MODUsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationPortsBoundingBoxResponse parses an HTTP response from a GetLocationPortsBoundingBoxWithResponse call
func ParseGetLocationPortsBoundingBoxResponse(rsp *http.Response) (*GetLocationPortsBoundingBoxResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationPortsBoundingBoxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationPortsRadiusResponse parses an HTTP response from a GetLocationPortsRadiusWithResponse call
func ParseGetLocationPortsRadiusResponse(rsp *http.Response) (*GetLocationPortsRadiusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationPortsRadiusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationRadiobeaconsBoundingBoxResponse parses an HTTP response from a GetLocationRadiobeaconsBoundingBoxWithResponse call
func ParseGetLocationRadiobeaconsBoundingBoxResponse(rsp *http.Response) (*GetLocationRadiobeaconsBoundingBoxResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationRadiobeaconsBoundingBoxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RadioBeaconsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationRadiobeaconsRadiusResponse parses an HTTP response from a GetLocationRadiobeaconsRadiusWithResponse call
func ParseGetLocationRadiobeaconsRadiusResponse(rsp *http.Response) (*GetLocationRadiobeaconsRadiusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationRadiobeaconsRadiusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RadioBeaconsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationVesselsBoundingBoxResponse parses an HTTP response from a GetLocationVesselsBoundingBoxWithResponse call
func ParseGetLocationVesselsBoundingBoxResponse(rsp *http.Response) (*GetLocationVesselsBoundingBoxResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationVesselsBoundingBoxResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetLocationVesselsRadiusResponse parses an HTTP response from a GetLocationVesselsRadiusWithResponse call
func ParseGetLocationVesselsRadiusResponse(rsp *http.Response) (*GetLocationVesselsRadiusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLocationVesselsRadiusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselsWithinLocationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetPortUnlocodeResponse parses an HTTP response from a GetPortUnlocodeWithResponse call
func ParseGetPortUnlocodeResponse(rsp *http.Response) (*GetPortUnlocodeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPortUnlocodeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetPortUnlocodeInboundResponse parses an HTTP response from a GetPortUnlocodeInboundWithResponse call
func ParseGetPortUnlocodeInboundResponse(rsp *http.Response) (*GetPortUnlocodeInboundResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPortUnlocodeInboundResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortInboundResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ServiceUnavailableErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetPorteventsResponse parses an HTTP response from a GetPorteventsWithResponse call
func ParseGetPorteventsResponse(rsp *http.Response) (*GetPorteventsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPorteventsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetPorteventsPortUnlocodeResponse parses an HTTP response from a GetPorteventsPortUnlocodeWithResponse call
func ParseGetPorteventsPortUnlocodeResponse(rsp *http.Response) (*GetPorteventsPortUnlocodeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPorteventsPortUnlocodeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetPorteventsPortsResponse parses an HTTP response from a GetPorteventsPortsWithResponse call
func ParseGetPorteventsPortsResponse(rsp *http.Response) (*GetPorteventsPortsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPorteventsPortsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetPorteventsVesselIdResponse parses an HTTP response from a GetPorteventsVesselIdWithResponse call
func ParseGetPorteventsVesselIdResponse(rsp *http.Response) (*GetPorteventsVesselIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPorteventsVesselIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetPorteventsVesselIdLastResponse parses an HTTP response from a GetPorteventsVesselIdLastWithResponse call
func ParseGetPorteventsVesselIdLastResponse(rsp *http.Response) (*GetPorteventsVesselIdLastResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPorteventsVesselIdLastResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortEventResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetPorteventsVesselsResponse parses an HTTP response from a GetPorteventsVesselsWithResponse call
func ParseGetPorteventsVesselsResponse(rsp *http.Response) (*GetPorteventsVesselsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPorteventsVesselsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PortEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetSearchDgpsResponse parses an HTTP response from a GetSearchDgpsWithResponse call
func ParseGetSearchDgpsResponse(rsp *http.Response) (*GetSearchDgpsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSearchDgpsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindDGPSStationsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetSearchLightaidsResponse parses an HTTP response from a GetSearchLightaidsWithResponse call
func ParseGetSearchLightaidsResponse(rsp *http.Response) (*GetSearchLightaidsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSearchLightaidsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindLightAidsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetSearchModusResponse parses an HTTP response from a GetSearchModusWithResponse call
func ParseGetSearchModusResponse(rsp *http.Response) (*GetSearchModusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSearchModusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindMODUsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetSearchPortsResponse parses an HTTP response from a GetSearchPortsWithResponse call
func ParseGetSearchPortsResponse(rsp *http.Response) (*GetSearchPortsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSearchPortsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindPortsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetSearchRadiobeaconsResponse parses an HTTP response from a GetSearchRadiobeaconsWithResponse call
func ParseGetSearchRadiobeaconsResponse(rsp *http.Response) (*GetSearchRadiobeaconsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSearchRadiobeaconsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindRadioBeaconsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetSearchVesselsResponse parses an HTTP response from a GetSearchVesselsWithResponse call
func ParseGetSearchVesselsResponse(rsp *http.Response) (*GetSearchVesselsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSearchVesselsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest FindVesselsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetVesselIdResponse parses an HTTP response from a GetVesselIdWithResponse call
func ParseGetVesselIdResponse(rsp *http.Response) (*GetVesselIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVesselIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetVesselIdCasualtiesResponse parses an HTTP response from a GetVesselIdCasualtiesWithResponse call
func ParseGetVesselIdCasualtiesResponse(rsp *http.Response) (*GetVesselIdCasualtiesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVesselIdCasualtiesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MarineCasualtiesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetVesselIdEmissionsResponse parses an HTTP response from a GetVesselIdEmissionsWithResponse call
func ParseGetVesselIdEmissionsResponse(rsp *http.Response) (*GetVesselIdEmissionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVesselIdEmissionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselEmissionsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetVesselIdEtaResponse parses an HTTP response from a GetVesselIdEtaWithResponse call
func ParseGetVesselIdEtaResponse(rsp *http.Response) (*GetVesselIdEtaResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVesselIdEtaResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselETAResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ServiceUnavailableErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetVesselIdPositionResponse parses an HTTP response from a GetVesselIdPositionWithResponse call
func ParseGetVesselIdPositionResponse(rsp *http.Response) (*GetVesselIdPositionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVesselIdPositionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselPositionResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 402:
		var dest PaymentRequiredErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON402 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFoundErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers GetVesselIdPositionResponse200Headers
		if values := rsp.Header.Values("X-Data-Source"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-Data-Source", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XDataSource = &value
		}
		if values := rsp.Header.Values("X-Satellite-Credit-Charged"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-Satellite-Credit-Charged", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XSatelliteCreditCharged = &value
		}
		if values := rsp.Header.Values("X-Satellite-Credits-Remaining"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-Satellite-Credits-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XSatelliteCreditsRemaining = &value
		}
		if values := rsp.Header.Values("X-Satellite-Lookback-Minutes"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-Satellite-Lookback-Minutes", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XSatelliteLookbackMinutes = &value
		}
		if values := rsp.Header.Values("X-Satellite-Max-Age-Minutes"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-Satellite-Max-Age-Minutes", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XSatelliteMaxAgeMinutes = &value
		}
		if values := rsp.Header.Values("X-Satellite-Status"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-Satellite-Status", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XSatelliteStatus = &value
		}
		response.Headers200 = &headers
	}

	return response, nil
}

// ParseGetVesselsPositionsResponse parses an HTTP response from a GetVesselsPositionsWithResponse call
func ParseGetVesselsPositionsResponse(rsp *http.Response) (*GetVesselsPositionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVesselsPositionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VesselPositionsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequestErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest AuthenticationErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ForbiddenErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest RateLimitErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
