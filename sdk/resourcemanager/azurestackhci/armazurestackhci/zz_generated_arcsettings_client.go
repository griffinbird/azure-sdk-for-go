//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ArcSettingsClient contains the methods for the ArcSettings group.
// Don't use this type directly, use NewArcSettingsClient() instead.
type ArcSettingsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewArcSettingsClient creates a new instance of ArcSettingsClient with the specified values.
func NewArcSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ArcSettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ArcSettingsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Create ArcSetting for HCI cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ArcSettingsClient) Create(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSetting, options *ArcSettingsCreateOptions) (ArcSettingsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, arcSetting, options)
	if err != nil {
		return ArcSettingsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArcSettingsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArcSettingsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ArcSettingsClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSetting, options *ArcSettingsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, arcSetting)
}

// createHandleResponse handles the Create response.
func (client *ArcSettingsClient) createHandleResponse(resp *http.Response) (ArcSettingsCreateResponse, error) {
	result := ArcSettingsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSetting); err != nil {
		return ArcSettingsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *ArcSettingsClient) createHandleError(resp *http.Response) error {
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

// BeginDelete - Delete ArcSetting resource details of HCI Cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ArcSettingsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsBeginDeleteOptions) (ArcSettingsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, arcSettingName, options)
	if err != nil {
		return ArcSettingsDeletePollerResponse{}, err
	}
	result := ArcSettingsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ArcSettingsClient.Delete", "azure-async-operation", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ArcSettingsDeletePollerResponse{}, err
	}
	result.Poller = &ArcSettingsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete ArcSetting resource details of HCI Cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ArcSettingsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
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
func (client *ArcSettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ArcSettingsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get ArcSetting resource details of HCI Cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ArcSettingsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsGetOptions) (ArcSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
	if err != nil {
		return ArcSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArcSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArcSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArcSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArcSettingsClient) getHandleResponse(resp *http.Response) (ArcSettingsGetResponse, error) {
	result := ArcSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSetting); err != nil {
		return ArcSettingsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ArcSettingsClient) getHandleError(resp *http.Response) error {
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

// ListByCluster - Get ArcSetting resources of HCI Cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ArcSettingsClient) ListByCluster(resourceGroupName string, clusterName string, options *ArcSettingsListByClusterOptions) *ArcSettingsListByClusterPager {
	return &ArcSettingsListByClusterPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
		},
		advancer: func(ctx context.Context, resp ArcSettingsListByClusterResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ArcSettingList.NextLink)
		},
	}
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *ArcSettingsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ArcSettingsListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *ArcSettingsClient) listByClusterHandleResponse(resp *http.Response) (ArcSettingsListByClusterResponse, error) {
	result := ArcSettingsListByClusterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSettingList); err != nil {
		return ArcSettingsListByClusterResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByClusterHandleError handles the ListByCluster error response.
func (client *ArcSettingsClient) listByClusterHandleError(resp *http.Response) error {
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