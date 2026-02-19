package vesselapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// VesselsService wraps vessel-related API endpoints.
type VesselsService struct {
	client *Client
}

// Get retrieves vessel details by ID (IMO or MMSI).
func (s *VesselsService) Get(ctx context.Context, id string, params *GetVesselIdParams) (*VesselResponse, error) {
	if params == nil {
		params = &GetVesselIdParams{FilterIdType: GetVesselIdParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselId(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Position retrieves the latest position for a vessel.
func (s *VesselsService) Position(ctx context.Context, id string, params *GetVesselIdPositionParams) (*VesselPositionResponse, error) {
	if params == nil {
		params = &GetVesselIdPositionParams{FilterIdType: GetVesselIdPositionParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdPosition(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdPositionResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Casualties retrieves marine casualty records for a vessel.
func (s *VesselsService) Casualties(ctx context.Context, id string, params *GetVesselIdCasualtiesParams) (*MarineCasualtiesResponse, error) {
	if params == nil {
		params = &GetVesselIdCasualtiesParams{FilterIdType: GetVesselIdCasualtiesParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdCasualties(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdCasualtiesResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Classification retrieves classification data for a vessel.
func (s *VesselsService) Classification(ctx context.Context, id string, params *GetVesselIdClassificationParams) (*ClassificationResponse, error) {
	if params == nil {
		params = &GetVesselIdClassificationParams{FilterIdType: GetVesselIdClassificationParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdClassification(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdClassificationResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Emissions retrieves emissions data for a vessel.
func (s *VesselsService) Emissions(ctx context.Context, id string, params *GetVesselIdEmissionsParams) (*VesselEmissionsResponse, error) {
	if params == nil {
		params = &GetVesselIdEmissionsParams{FilterIdType: GetVesselIdEmissionsParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdEmissions(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdEmissionsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// ETA retrieves the estimated time of arrival for a vessel.
func (s *VesselsService) ETA(ctx context.Context, id string, params *GetVesselIdEtaParams) (*VesselETAResponse, error) {
	if params == nil {
		params = &GetVesselIdEtaParams{FilterIdType: GetVesselIdEtaParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdEta(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdEtaResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Inspections retrieves inspection records for a vessel.
func (s *VesselsService) Inspections(ctx context.Context, id string, params *GetVesselIdInspectionsParams) (*TypesInspectionsResponse, error) {
	if params == nil {
		params = &GetVesselIdInspectionsParams{FilterIdType: GetVesselIdInspectionsParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdInspections(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdInspectionsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// InspectionDetail retrieves detailed inspection data.
func (s *VesselsService) InspectionDetail(ctx context.Context, id, detailId string, params *GetVesselIdInspectionsDetailIdParams) (*TypesInspectionDetailResponse, error) {
	if params == nil {
		params = &GetVesselIdInspectionsDetailIdParams{FilterIdType: GetVesselIdInspectionsDetailIdParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdInspectionsDetailId(ctx, id, detailId, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdInspectionsDetailIdResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Ownership retrieves ownership data for a vessel.
func (s *VesselsService) Ownership(ctx context.Context, id string, params *GetVesselIdOwnershipParams) (*TypesOwnershipResponse, error) {
	if params == nil {
		params = &GetVesselIdOwnershipParams{FilterIdType: GetVesselIdOwnershipParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselIdOwnership(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselIdOwnershipResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Positions retrieves positions for multiple vessels.
func (s *VesselsService) Positions(ctx context.Context, params *GetVesselsPositionsParams) (*VesselPositionsResponse, error) {
	if params == nil {
		params = &GetVesselsPositionsParams{FilterIdType: GetVesselsPositionsParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetVesselsPositions(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetVesselsPositionsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// PortsService wraps port-related API endpoints.
type PortsService struct {
	client *Client
}

// Get retrieves a port by its UN/LOCODE.
func (s *PortsService) Get(ctx context.Context, unlocode string) (*PortResponse, error) {
	rsp, err := s.client.GetPortUnlocode(ctx, unlocode)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetPortUnlocodeResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// PortEventsService wraps port event API endpoints.
type PortEventsService struct {
	client *Client
}

// List retrieves port events with optional filtering.
func (s *PortEventsService) List(ctx context.Context, params *GetPorteventsParams) (*PortEventsResponse, error) {
	if params == nil {
		params = &GetPorteventsParams{}
	}
	rsp, err := s.client.GetPortevents(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetPorteventsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// ByPort retrieves port events for a specific port by UNLOCODE.
func (s *PortEventsService) ByPort(ctx context.Context, unlocode string, params *GetPorteventsPortUnlocodeParams) (*PortEventsResponse, error) {
	if params == nil {
		params = &GetPorteventsPortUnlocodeParams{}
	}
	rsp, err := s.client.GetPorteventsPortUnlocode(ctx, unlocode, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetPorteventsPortUnlocodeResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// ByPorts retrieves port events by port name search.
func (s *PortEventsService) ByPorts(ctx context.Context, params *GetPorteventsPortsParams) (*PortEventsResponse, error) {
	if params == nil {
		params = &GetPorteventsPortsParams{}
	}
	rsp, err := s.client.GetPorteventsPorts(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetPorteventsPortsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// ByVessel retrieves port events for a specific vessel.
func (s *PortEventsService) ByVessel(ctx context.Context, id string, params *GetPorteventsVesselIdParams) (*PortEventsResponse, error) {
	if params == nil {
		params = &GetPorteventsVesselIdParams{FilterIdType: GetPorteventsVesselIdParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetPorteventsVesselId(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetPorteventsVesselIdResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// LastByVessel retrieves the last port event for a vessel.
func (s *PortEventsService) LastByVessel(ctx context.Context, id string, params *GetPorteventsVesselIdLastParams) (*PortEventResponse, error) {
	if params == nil {
		params = &GetPorteventsVesselIdLastParams{FilterIdType: GetPorteventsVesselIdLastParamsFilterIdTypeImo}
	}
	rsp, err := s.client.GetPorteventsVesselIdLast(ctx, id, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetPorteventsVesselIdLastResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// ByVessels retrieves port events by vessel name search.
func (s *PortEventsService) ByVessels(ctx context.Context, params *GetPorteventsVesselsParams) (*PortEventsResponse, error) {
	if params == nil {
		params = &GetPorteventsVesselsParams{}
	}
	rsp, err := s.client.GetPorteventsVessels(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetPorteventsVesselsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// EmissionsService wraps emissions API endpoints.
type EmissionsService struct {
	client *Client
}

// List retrieves vessel emissions data.
func (s *EmissionsService) List(ctx context.Context, params *GetEmissionsParams) (*VesselEmissionsResponse, error) {
	if params == nil {
		params = &GetEmissionsParams{}
	}
	rsp, err := s.client.GetEmissions(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetEmissionsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// SearchService wraps search API endpoints.
type SearchService struct {
	client *Client
}

// Vessels searches for vessels by name or callsign.
func (s *SearchService) Vessels(ctx context.Context, params *GetSearchVesselsParams) (*FindVesselsResponse, error) {
	if params == nil {
		params = &GetSearchVesselsParams{}
	}
	rsp, err := s.client.GetSearchVessels(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetSearchVesselsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// Ports searches for ports by name.
func (s *SearchService) Ports(ctx context.Context, params *GetSearchPortsParams) (*FindPortsResponse, error) {
	if params == nil {
		params = &GetSearchPortsParams{}
	}
	rsp, err := s.client.GetSearchPorts(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetSearchPortsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// DGPS searches for DGPS stations by name.
func (s *SearchService) DGPS(ctx context.Context, params *GetSearchDgpsParams) (*FindDGPSStationsResponse, error) {
	if params == nil {
		params = &GetSearchDgpsParams{}
	}
	rsp, err := s.client.GetSearchDgps(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetSearchDgpsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// LightAids searches for light aids by name.
func (s *SearchService) LightAids(ctx context.Context, params *GetSearchLightaidsParams) (*FindLightAidsResponse, error) {
	if params == nil {
		params = &GetSearchLightaidsParams{}
	}
	rsp, err := s.client.GetSearchLightaids(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetSearchLightaidsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// MODUs searches for MODUs (Mobile Offshore Drilling Units) by name.
func (s *SearchService) MODUs(ctx context.Context, params *GetSearchModusParams) (*FindMODUsResponse, error) {
	if params == nil {
		params = &GetSearchModusParams{}
	}
	rsp, err := s.client.GetSearchModus(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetSearchModusResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// RadioBeacons searches for radio beacons by name.
func (s *SearchService) RadioBeacons(ctx context.Context, params *GetSearchRadiobeaconsParams) (*FindRadioBeaconsResponse, error) {
	if params == nil {
		params = &GetSearchRadiobeaconsParams{}
	}
	rsp, err := s.client.GetSearchRadiobeacons(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetSearchRadiobeaconsResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// LocationService wraps location-based API endpoints.
type LocationService struct {
	client *Client
}

// VesselsBoundingBox retrieves vessel positions within a bounding box.
func (s *LocationService) VesselsBoundingBox(ctx context.Context, params *GetLocationVesselsBoundingBoxParams) (*VesselsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationVesselsBoundingBoxParams{}
	}
	rsp, err := s.client.GetLocationVesselsBoundingBox(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationVesselsBoundingBoxResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// VesselsRadius retrieves vessel positions within a radius.
func (s *LocationService) VesselsRadius(ctx context.Context, params *GetLocationVesselsRadiusParams) (*VesselsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationVesselsRadiusParams{}
	}
	rsp, err := s.client.GetLocationVesselsRadius(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationVesselsRadiusResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// PortsBoundingBox retrieves ports within a bounding box.
func (s *LocationService) PortsBoundingBox(ctx context.Context, params *GetLocationPortsBoundingBoxParams) (*PortsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationPortsBoundingBoxParams{}
	}
	rsp, err := s.client.GetLocationPortsBoundingBox(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationPortsBoundingBoxResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// PortsRadius retrieves ports within a radius.
func (s *LocationService) PortsRadius(ctx context.Context, params *GetLocationPortsRadiusParams) (*PortsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationPortsRadiusParams{}
	}
	rsp, err := s.client.GetLocationPortsRadius(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationPortsRadiusResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// DGPSBoundingBox retrieves DGPS stations within a bounding box.
func (s *LocationService) DGPSBoundingBox(ctx context.Context, params *GetLocationDgpsBoundingBoxParams) (*DGPSStationsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationDgpsBoundingBoxParams{}
	}
	rsp, err := s.client.GetLocationDgpsBoundingBox(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationDgpsBoundingBoxResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// DGPSRadius retrieves DGPS stations within a radius.
func (s *LocationService) DGPSRadius(ctx context.Context, params *GetLocationDgpsRadiusParams) (*DGPSStationsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationDgpsRadiusParams{}
	}
	rsp, err := s.client.GetLocationDgpsRadius(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationDgpsRadiusResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// LightAidsBoundingBox retrieves light aids within a bounding box.
func (s *LocationService) LightAidsBoundingBox(ctx context.Context, params *GetLocationLightaidsBoundingBoxParams) (*LightAidsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationLightaidsBoundingBoxParams{}
	}
	rsp, err := s.client.GetLocationLightaidsBoundingBox(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationLightaidsBoundingBoxResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// LightAidsRadius retrieves light aids within a radius.
func (s *LocationService) LightAidsRadius(ctx context.Context, params *GetLocationLightaidsRadiusParams) (*LightAidsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationLightaidsRadiusParams{}
	}
	rsp, err := s.client.GetLocationLightaidsRadius(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationLightaidsRadiusResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// MODUsBoundingBox retrieves MODUs within a bounding box.
func (s *LocationService) MODUsBoundingBox(ctx context.Context, params *GetLocationModuBoundingBoxParams) (*MODUsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationModuBoundingBoxParams{}
	}
	rsp, err := s.client.GetLocationModuBoundingBox(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationModuBoundingBoxResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// MODUsRadius retrieves MODUs within a radius.
func (s *LocationService) MODUsRadius(ctx context.Context, params *GetLocationModuRadiusParams) (*MODUsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationModuRadiusParams{}
	}
	rsp, err := s.client.GetLocationModuRadius(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationModuRadiusResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// RadioBeaconsBoundingBox retrieves radio beacons within a bounding box.
func (s *LocationService) RadioBeaconsBoundingBox(ctx context.Context, params *GetLocationRadiobeaconsBoundingBoxParams) (*RadioBeaconsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationRadiobeaconsBoundingBoxParams{}
	}
	rsp, err := s.client.GetLocationRadiobeaconsBoundingBox(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationRadiobeaconsBoundingBoxResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// RadioBeaconsRadius retrieves radio beacons within a radius.
func (s *LocationService) RadioBeaconsRadius(ctx context.Context, params *GetLocationRadiobeaconsRadiusParams) (*RadioBeaconsWithinLocationResponse, error) {
	if params == nil {
		params = &GetLocationRadiobeaconsRadiusParams{}
	}
	rsp, err := s.client.GetLocationRadiobeaconsRadius(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetLocationRadiobeaconsRadiusResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// NavtexService wraps NAVTEX message API endpoints.
type NavtexService struct {
	client *Client
}

// List retrieves NAVTEX maritime safety messages.
func (s *NavtexService) List(ctx context.Context, params *GetNavtexParams) (*NavtexMessagesResponse, error) {
	if params == nil {
		params = &GetNavtexParams{}
	}
	rsp, err := s.client.GetNavtex(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	parsed, err := ParseGetNavtexResponse(rsp)
	if err != nil {
		return nil, fmt.Errorf("vesselapi: %w", err)
	}
	if err := errFromStatus(parsed.StatusCode(), parsed.Body); err != nil {
		return nil, err
	}
	if parsed.JSON200 == nil {
		return nil, &APIError{StatusCode: parsed.StatusCode(), Message: "unexpected empty response", Body: parsed.Body}
	}
	return parsed.JSON200, nil
}

// --- Error checking helpers ---

func errFromStatus(statusCode int, body []byte) error {
	if statusCode >= 200 && statusCode < 300 {
		return nil
	}
	msg := http.StatusText(statusCode)
	if len(body) > 0 {
		// Try {"error":{"message":"..."}} (Vessel API standard shape).
		var nested struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if json.Unmarshal(body, &nested) == nil && nested.Error.Message != "" {
			msg = nested.Error.Message
		} else {
			// Try {"message":"..."} (common alternative shape).
			var flat struct {
				Message string `json:"message"`
			}
			if json.Unmarshal(body, &flat) == nil && flat.Message != "" {
				msg = flat.Message
			}
		}
		// If both fail, msg stays as http.StatusText. Raw body is in APIError.Body.
	}
	return &APIError{StatusCode: statusCode, Message: msg, Body: body}
}
