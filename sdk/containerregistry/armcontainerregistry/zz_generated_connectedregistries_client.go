//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ConnectedRegistriesClient contains the methods for the ConnectedRegistries group.
// Don't use this type directly, use NewConnectedRegistriesClient() instead.
type ConnectedRegistriesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConnectedRegistriesClient creates a new instance of ConnectedRegistriesClient with the specified values.
func NewConnectedRegistriesClient(con *arm.Connection, subscriptionID string) *ConnectedRegistriesClient {
	return &ConnectedRegistriesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreate - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesBeginCreateOptions) (ConnectedRegistriesCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters, options)
	if err != nil {
		return ConnectedRegistriesCreatePollerResponse{}, err
	}
	result := ConnectedRegistriesCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectedRegistriesClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return ConnectedRegistriesCreatePollerResponse{}, err
	}
	result.Poller = &ConnectedRegistriesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) create(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ConnectedRegistriesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, connectedRegistryCreateParameters)
}

// createHandleError handles the Create error response.
func (client *ConnectedRegistriesClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDeactivate - Deactivates the connected registry instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginDeactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeactivateOptions) (ConnectedRegistriesDeactivatePollerResponse, error) {
	resp, err := client.deactivate(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return ConnectedRegistriesDeactivatePollerResponse{}, err
	}
	result := ConnectedRegistriesDeactivatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectedRegistriesClient.Deactivate", "", resp, client.pl, client.deactivateHandleError)
	if err != nil {
		return ConnectedRegistriesDeactivatePollerResponse{}, err
	}
	result.Poller = &ConnectedRegistriesDeactivatePoller{
		pt: pt,
	}
	return result, nil
}

// Deactivate - Deactivates the connected registry instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) deactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeactivateOptions) (*http.Response, error) {
	req, err := client.deactivateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.deactivateHandleError(resp)
	}
	return resp, nil
}

// deactivateCreateRequest creates the Deactivate request.
func (client *ConnectedRegistriesClient) deactivateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeactivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}/deactivate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deactivateHandleError handles the Deactivate error response.
func (client *ConnectedRegistriesClient) deactivateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a connected registry from a container registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeleteOptions) (ConnectedRegistriesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return ConnectedRegistriesDeletePollerResponse{}, err
	}
	result := ConnectedRegistriesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectedRegistriesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ConnectedRegistriesDeletePollerResponse{}, err
	}
	result.Poller = &ConnectedRegistriesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a connected registry from a container registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedRegistriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ConnectedRegistriesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the properties of the connected registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) Get(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesGetOptions) (ConnectedRegistriesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return ConnectedRegistriesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedRegistriesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedRegistriesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectedRegistriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedRegistriesClient) getHandleResponse(resp *http.Response) (ConnectedRegistriesGetResponse, error) {
	result := ConnectedRegistriesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedRegistry); err != nil {
		return ConnectedRegistriesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ConnectedRegistriesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all connected registries for the specified container registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) List(resourceGroupName string, registryName string, options *ConnectedRegistriesListOptions) *ConnectedRegistriesListPager {
	return &ConnectedRegistriesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		advancer: func(ctx context.Context, resp ConnectedRegistriesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConnectedRegistryListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ConnectedRegistriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ConnectedRegistriesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedRegistriesClient) listHandleResponse(resp *http.Response) (ConnectedRegistriesListResponse, error) {
	result := ConnectedRegistriesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedRegistryListResult); err != nil {
		return ConnectedRegistriesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ConnectedRegistriesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates a connected registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesBeginUpdateOptions) (ConnectedRegistriesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters, options)
	if err != nil {
		return ConnectedRegistriesUpdatePollerResponse{}, err
	}
	result := ConnectedRegistriesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectedRegistriesClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return ConnectedRegistriesUpdatePollerResponse{}, err
	}
	result.Poller = &ConnectedRegistriesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a connected registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) update(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ConnectedRegistriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, connectedRegistryUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *ConnectedRegistriesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
