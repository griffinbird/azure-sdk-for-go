//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// EventChannelsClient contains the methods for the EventChannels group.
// Don't use this type directly, use NewEventChannelsClient() instead.
type EventChannelsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewEventChannelsClient creates a new instance of EventChannelsClient with the specified values.
func NewEventChannelsClient(con *arm.Connection, subscriptionID string) *EventChannelsClient {
	return &EventChannelsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Asynchronously creates a new event channel with the specified parameters.
// If the operation fails it returns a generic error.
func (client *EventChannelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, eventChannelInfo EventChannel, options *EventChannelsCreateOrUpdateOptions) (EventChannelsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerNamespaceName, eventChannelName, eventChannelInfo, options)
	if err != nil {
		return EventChannelsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EventChannelsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EventChannelsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EventChannelsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, eventChannelInfo EventChannel, options *EventChannelsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels/{eventChannelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if eventChannelName == "" {
		return nil, errors.New("parameter eventChannelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventChannelName}", url.PathEscape(eventChannelName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, eventChannelInfo)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EventChannelsClient) createOrUpdateHandleResponse(resp *http.Response) (EventChannelsCreateOrUpdateResponse, error) {
	result := EventChannelsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventChannel); err != nil {
		return EventChannelsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *EventChannelsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete existing event channel.
// If the operation fails it returns a generic error.
func (client *EventChannelsClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, options *EventChannelsBeginDeleteOptions) (EventChannelsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, partnerNamespaceName, eventChannelName, options)
	if err != nil {
		return EventChannelsDeletePollerResponse{}, err
	}
	result := EventChannelsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("EventChannelsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return EventChannelsDeletePollerResponse{}, err
	}
	result.Poller = &EventChannelsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete existing event channel.
// If the operation fails it returns a generic error.
func (client *EventChannelsClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, options *EventChannelsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerNamespaceName, eventChannelName, options)
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
func (client *EventChannelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, options *EventChannelsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels/{eventChannelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if eventChannelName == "" {
		return nil, errors.New("parameter eventChannelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventChannelName}", url.PathEscape(eventChannelName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *EventChannelsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get properties of an event channel.
// If the operation fails it returns a generic error.
func (client *EventChannelsClient) Get(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, options *EventChannelsGetOptions) (EventChannelsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerNamespaceName, eventChannelName, options)
	if err != nil {
		return EventChannelsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EventChannelsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EventChannelsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EventChannelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, options *EventChannelsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels/{eventChannelName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	if eventChannelName == "" {
		return nil, errors.New("parameter eventChannelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventChannelName}", url.PathEscape(eventChannelName))
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
func (client *EventChannelsClient) getHandleResponse(resp *http.Response) (EventChannelsGetResponse, error) {
	result := EventChannelsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventChannel); err != nil {
		return EventChannelsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *EventChannelsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByPartnerNamespace - List all the event channels in a partner namespace.
// If the operation fails it returns a generic error.
func (client *EventChannelsClient) ListByPartnerNamespace(resourceGroupName string, partnerNamespaceName string, options *EventChannelsListByPartnerNamespaceOptions) *EventChannelsListByPartnerNamespacePager {
	return &EventChannelsListByPartnerNamespacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByPartnerNamespaceCreateRequest(ctx, resourceGroupName, partnerNamespaceName, options)
		},
		advancer: func(ctx context.Context, resp EventChannelsListByPartnerNamespaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EventChannelsListResult.NextLink)
		},
	}
}

// listByPartnerNamespaceCreateRequest creates the ListByPartnerNamespace request.
func (client *EventChannelsClient) listByPartnerNamespaceCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *EventChannelsListByPartnerNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPartnerNamespaceHandleResponse handles the ListByPartnerNamespace response.
func (client *EventChannelsClient) listByPartnerNamespaceHandleResponse(resp *http.Response) (EventChannelsListByPartnerNamespaceResponse, error) {
	result := EventChannelsListByPartnerNamespaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventChannelsListResult); err != nil {
		return EventChannelsListByPartnerNamespaceResponse{}, err
	}
	return result, nil
}

// listByPartnerNamespaceHandleError handles the ListByPartnerNamespace error response.
func (client *EventChannelsClient) listByPartnerNamespaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
