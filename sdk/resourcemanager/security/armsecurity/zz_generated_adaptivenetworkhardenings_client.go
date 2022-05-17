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

// AdaptiveNetworkHardeningsClient contains the methods for the AdaptiveNetworkHardenings group.
// Don't use this type directly, use NewAdaptiveNetworkHardeningsClient() instead.
type AdaptiveNetworkHardeningsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAdaptiveNetworkHardeningsClient creates a new instance of AdaptiveNetworkHardeningsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAdaptiveNetworkHardeningsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AdaptiveNetworkHardeningsClient, error) {
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
	client := &AdaptiveNetworkHardeningsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginEnforce - Enforces the given rules on the NSG(s) listed in the request
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// resourceNamespace - The Namespace of the resource.
// resourceType - The type of the resource.
// resourceName - Name of the resource.
// adaptiveNetworkHardeningResourceName - The name of the Adaptive Network Hardening resource.
// options - AdaptiveNetworkHardeningsClientBeginEnforceOptions contains the optional parameters for the AdaptiveNetworkHardeningsClient.BeginEnforce
// method.
func (client *AdaptiveNetworkHardeningsClient) BeginEnforce(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body AdaptiveNetworkHardeningEnforceRequest, options *AdaptiveNetworkHardeningsClientBeginEnforceOptions) (*runtime.Poller[AdaptiveNetworkHardeningsClientEnforceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.enforce(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, adaptiveNetworkHardeningResourceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AdaptiveNetworkHardeningsClientEnforceResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AdaptiveNetworkHardeningsClientEnforceResponse](options.ResumeToken, client.pl, nil)
	}
}

// Enforce - Enforces the given rules on the NSG(s) listed in the request
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
func (client *AdaptiveNetworkHardeningsClient) enforce(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body AdaptiveNetworkHardeningEnforceRequest, options *AdaptiveNetworkHardeningsClientBeginEnforceOptions) (*http.Response, error) {
	req, err := client.enforceCreateRequest(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, adaptiveNetworkHardeningResourceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// enforceCreateRequest creates the Enforce request.
func (client *AdaptiveNetworkHardeningsClient) enforceCreateRequest(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body AdaptiveNetworkHardeningEnforceRequest, options *AdaptiveNetworkHardeningsClientBeginEnforceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/adaptiveNetworkHardenings/{adaptiveNetworkHardeningResourceName}/{adaptiveNetworkHardeningEnforceAction}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceNamespace == "" {
		return nil, errors.New("parameter resourceNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceNamespace}", url.PathEscape(resourceNamespace))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if adaptiveNetworkHardeningResourceName == "" {
		return nil, errors.New("parameter adaptiveNetworkHardeningResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{adaptiveNetworkHardeningResourceName}", url.PathEscape(adaptiveNetworkHardeningResourceName))
	urlPath = strings.ReplaceAll(urlPath, "{adaptiveNetworkHardeningEnforceAction}", url.PathEscape("enforce"))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// Get - Gets a single Adaptive Network Hardening resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// resourceNamespace - The Namespace of the resource.
// resourceType - The type of the resource.
// resourceName - Name of the resource.
// adaptiveNetworkHardeningResourceName - The name of the Adaptive Network Hardening resource.
// options - AdaptiveNetworkHardeningsClientGetOptions contains the optional parameters for the AdaptiveNetworkHardeningsClient.Get
// method.
func (client *AdaptiveNetworkHardeningsClient) Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, options *AdaptiveNetworkHardeningsClientGetOptions) (AdaptiveNetworkHardeningsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, adaptiveNetworkHardeningResourceName, options)
	if err != nil {
		return AdaptiveNetworkHardeningsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AdaptiveNetworkHardeningsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AdaptiveNetworkHardeningsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AdaptiveNetworkHardeningsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, options *AdaptiveNetworkHardeningsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/adaptiveNetworkHardenings/{adaptiveNetworkHardeningResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceNamespace == "" {
		return nil, errors.New("parameter resourceNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceNamespace}", url.PathEscape(resourceNamespace))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if adaptiveNetworkHardeningResourceName == "" {
		return nil, errors.New("parameter adaptiveNetworkHardeningResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{adaptiveNetworkHardeningResourceName}", url.PathEscape(adaptiveNetworkHardeningResourceName))
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

// getHandleResponse handles the Get response.
func (client *AdaptiveNetworkHardeningsClient) getHandleResponse(resp *http.Response) (AdaptiveNetworkHardeningsClientGetResponse, error) {
	result := AdaptiveNetworkHardeningsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdaptiveNetworkHardening); err != nil {
		return AdaptiveNetworkHardeningsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByExtendedResourcePager - Gets a list of Adaptive Network Hardenings resources in scope of an extended resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// resourceNamespace - The Namespace of the resource.
// resourceType - The type of the resource.
// resourceName - Name of the resource.
// options - AdaptiveNetworkHardeningsClientListByExtendedResourceOptions contains the optional parameters for the AdaptiveNetworkHardeningsClient.ListByExtendedResource
// method.
func (client *AdaptiveNetworkHardeningsClient) NewListByExtendedResourcePager(resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, options *AdaptiveNetworkHardeningsClientListByExtendedResourceOptions) *runtime.Pager[AdaptiveNetworkHardeningsClientListByExtendedResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[AdaptiveNetworkHardeningsClientListByExtendedResourceResponse]{
		More: func(page AdaptiveNetworkHardeningsClientListByExtendedResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AdaptiveNetworkHardeningsClientListByExtendedResourceResponse) (AdaptiveNetworkHardeningsClientListByExtendedResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByExtendedResourceCreateRequest(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AdaptiveNetworkHardeningsClientListByExtendedResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AdaptiveNetworkHardeningsClientListByExtendedResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AdaptiveNetworkHardeningsClientListByExtendedResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByExtendedResourceHandleResponse(resp)
		},
	})
}

// listByExtendedResourceCreateRequest creates the ListByExtendedResource request.
func (client *AdaptiveNetworkHardeningsClient) listByExtendedResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, options *AdaptiveNetworkHardeningsClientListByExtendedResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/adaptiveNetworkHardenings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceNamespace == "" {
		return nil, errors.New("parameter resourceNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceNamespace}", url.PathEscape(resourceNamespace))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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

// listByExtendedResourceHandleResponse handles the ListByExtendedResource response.
func (client *AdaptiveNetworkHardeningsClient) listByExtendedResourceHandleResponse(resp *http.Response) (AdaptiveNetworkHardeningsClientListByExtendedResourceResponse, error) {
	result := AdaptiveNetworkHardeningsClientListByExtendedResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdaptiveNetworkHardeningsList); err != nil {
		return AdaptiveNetworkHardeningsClientListByExtendedResourceResponse{}, err
	}
	return result, nil
}
