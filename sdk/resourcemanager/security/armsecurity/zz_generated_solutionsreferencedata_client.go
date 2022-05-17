//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SolutionsReferenceDataClient contains the methods for the SecuritySolutionsReferenceData group.
// Don't use this type directly, use NewSolutionsReferenceDataClient() instead.
type SolutionsReferenceDataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSolutionsReferenceDataClient creates a new instance of SolutionsReferenceDataClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSolutionsReferenceDataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SolutionsReferenceDataClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SolutionsReferenceDataClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Gets a list of all supported Security Solutions for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
// options - SolutionsReferenceDataClientListOptions contains the optional parameters for the SolutionsReferenceDataClient.List
// method.
func (client *SolutionsReferenceDataClient) List(ctx context.Context, options *SolutionsReferenceDataClientListOptions) (SolutionsReferenceDataClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return SolutionsReferenceDataClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SolutionsReferenceDataClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SolutionsReferenceDataClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SolutionsReferenceDataClient) listCreateRequest(ctx context.Context, options *SolutionsReferenceDataClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/securitySolutionsReferenceData"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SolutionsReferenceDataClient) listHandleResponse(resp *http.Response) (SolutionsReferenceDataClientListResponse, error) {
	result := SolutionsReferenceDataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionsReferenceDataList); err != nil {
		return SolutionsReferenceDataClientListResponse{}, err
	}
	return result, nil
}

// ListByHomeRegion - Gets list of all supported Security Solutions for subscription and location.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
// ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
// options - SolutionsReferenceDataClientListByHomeRegionOptions contains the optional parameters for the SolutionsReferenceDataClient.ListByHomeRegion
// method.
func (client *SolutionsReferenceDataClient) ListByHomeRegion(ctx context.Context, ascLocation string, options *SolutionsReferenceDataClientListByHomeRegionOptions) (SolutionsReferenceDataClientListByHomeRegionResponse, error) {
	req, err := client.listByHomeRegionCreateRequest(ctx, ascLocation, options)
	if err != nil {
		return SolutionsReferenceDataClientListByHomeRegionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SolutionsReferenceDataClientListByHomeRegionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SolutionsReferenceDataClientListByHomeRegionResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByHomeRegionHandleResponse(resp)
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *SolutionsReferenceDataClient) listByHomeRegionCreateRequest(ctx context.Context, ascLocation string, options *SolutionsReferenceDataClientListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/securitySolutionsReferenceData"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *SolutionsReferenceDataClient) listByHomeRegionHandleResponse(resp *http.Response) (SolutionsReferenceDataClientListByHomeRegionResponse, error) {
	result := SolutionsReferenceDataClientListByHomeRegionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionsReferenceDataList); err != nil {
		return SolutionsReferenceDataClientListByHomeRegionResponse{}, err
	}
	return result, nil
}
