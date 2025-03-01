/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// ResourceApiApiService is a service that implements the logic for the ResourceApiApiServicer
// This service should implement the business logic for every endpoint for the ResourceApiApi API.
// Include any external packages or services that will be required by this service.
type ResourceApiApiService struct {
}

// NewResourceApiApiService creates a default api service
func NewResourceApiApiService() ResourceApiApiServicer {
	return &ResourceApiApiService{}
}

// CreateDemo - Create a Demonstration setting
func (s *ResourceApiApiService) CreateDemo(ctx context.Context, demo Demo) (ImplResponse, error) {
	// TODO - update CreateDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Demo{}) or use other options such as http.Ok ...
	//return Response(201, Demo{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateDemo method not implemented")
}

// CreatePollingProfile - Create a Polling Profile
func (s *ResourceApiApiService) CreatePollingProfile(ctx context.Context, pollingProfile PollingProfile) (ImplResponse, error) {
	// TODO - update CreatePollingProfile with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, PollingProfile{}) or use other options such as http.Ok ...
	//return Response(201, PollingProfile{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreatePollingProfile method not implemented")
}

// DeleteDemo - Delete a Demonstration setting
func (s *ResourceApiApiService) DeleteDemo(ctx context.Context, demoId string) (ImplResponse, error) {
	// TODO - update DeleteDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteDemo method not implemented")
}

// DeletePollingProfile - Delete a Polling Profile
func (s *ResourceApiApiService) DeletePollingProfile(ctx context.Context, pollingProfileId string) (ImplResponse, error) {
	// TODO - update DeletePollingProfile with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePollingProfile method not implemented")
}

// GetConfiguration - Get Tracetest configuration
func (s *ResourceApiApiService) GetConfiguration(ctx context.Context, configId string) (ImplResponse, error) {
	// TODO - update GetConfiguration with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ConfigurationResource{}) or use other options such as http.Ok ...
	//return Response(200, ConfigurationResource{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetConfiguration method not implemented")
}

// GetDemo - Get Demonstration setting
func (s *ResourceApiApiService) GetDemo(ctx context.Context, demoId string) (ImplResponse, error) {
	// TODO - update GetDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Demo{}) or use other options such as http.Ok ...
	//return Response(200, Demo{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetDemo method not implemented")
}

// GetPollingProfile - Get Polling Profile
func (s *ResourceApiApiService) GetPollingProfile(ctx context.Context, pollingProfileId string) (ImplResponse, error) {
	// TODO - update GetPollingProfile with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PollingProfile{}) or use other options such as http.Ok ...
	//return Response(200, PollingProfile{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPollingProfile method not implemented")
}

// ListDemos - List Demonstrations
func (s *ResourceApiApiService) ListDemos(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListDemos with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Demo{}) or use other options such as http.Ok ...
	//return Response(200, []Demo{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListDemos method not implemented")
}

// ListPollingProfiles - List Polling Profiles
func (s *ResourceApiApiService) ListPollingProfiles(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListPollingProfiles with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ListPollingProfiles200Response{}) or use other options such as http.Ok ...
	//return Response(200, ListPollingProfiles200Response{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListPollingProfiles method not implemented")
}

// UpdateConfiguration - Update Tracetest configuration
func (s *ResourceApiApiService) UpdateConfiguration(ctx context.Context, configId string, configurationResource ConfigurationResource) (ImplResponse, error) {
	// TODO - update UpdateConfiguration with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ConfigurationResource{}) or use other options such as http.Ok ...
	//return Response(200, ConfigurationResource{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateConfiguration method not implemented")
}

// UpdateDemo - Update a Demonstration setting
func (s *ResourceApiApiService) UpdateDemo(ctx context.Context, demoId string, demo Demo) (ImplResponse, error) {
	// TODO - update UpdateDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Demo{}) or use other options such as http.Ok ...
	//return Response(200, Demo{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateDemo method not implemented")
}

// UpdatePollingProfile - Update a Polling Profile
func (s *ResourceApiApiService) UpdatePollingProfile(ctx context.Context, pollingProfileId string, pollingProfile PollingProfile) (ImplResponse, error) {
	// TODO - update UpdatePollingProfile with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PollingProfile{}) or use other options such as http.Ok ...
	//return Response(200, PollingProfile{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePollingProfile method not implemented")
}
