// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// WebApplicationFirewallPoliciesOperations contains the methods for the WebApplicationFirewallPolicies group.
type WebApplicationFirewallPoliciesOperations interface {
	// CreateOrUpdate - Creates or update policy with specified rule set name within a resource group.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (*WebApplicationFirewallPolicyResponse, error)
	// BeginDelete - Deletes Policy.
	BeginDelete(ctx context.Context, resourceGroupName string, policyName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieve protection policy with specified name within a resource group.
	Get(ctx context.Context, resourceGroupName string, policyName string) (*WebApplicationFirewallPolicyResponse, error)
	// List - Lists all of the protection policies within a resource group.
	List(resourceGroupName string) (WebApplicationFirewallPolicyListResultPager, error)
	// ListAll - Gets all the WAF policies in a subscription.
	ListAll() (WebApplicationFirewallPolicyListResultPager, error)
}

// webApplicationFirewallPoliciesOperations implements the WebApplicationFirewallPoliciesOperations interface.
type webApplicationFirewallPoliciesOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or update policy with specified rule set name within a resource group.
func (client *webApplicationFirewallPoliciesOperations) CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (*WebApplicationFirewallPolicyResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, policyName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *webApplicationFirewallPoliciesOperations) createOrUpdateCreateRequest(resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *webApplicationFirewallPoliciesOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := WebApplicationFirewallPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicy)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *webApplicationFirewallPoliciesOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes Policy.
func (client *webApplicationFirewallPoliciesOperations) BeginDelete(ctx context.Context, resourceGroupName string, policyName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, policyName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("webApplicationFirewallPoliciesOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *webApplicationFirewallPoliciesOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("webApplicationFirewallPoliciesOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *webApplicationFirewallPoliciesOperations) deleteCreateRequest(resourceGroupName string, policyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *webApplicationFirewallPoliciesOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *webApplicationFirewallPoliciesOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Retrieve protection policy with specified name within a resource group.
func (client *webApplicationFirewallPoliciesOperations) Get(ctx context.Context, resourceGroupName string, policyName string) (*WebApplicationFirewallPolicyResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, policyName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *webApplicationFirewallPoliciesOperations) getCreateRequest(resourceGroupName string, policyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *webApplicationFirewallPoliciesOperations) getHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := WebApplicationFirewallPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicy)
}

// getHandleError handles the Get error response.
func (client *webApplicationFirewallPoliciesOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all of the protection policies within a resource group.
func (client *webApplicationFirewallPoliciesOperations) List(resourceGroupName string) (WebApplicationFirewallPolicyListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &webApplicationFirewallPolicyListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *WebApplicationFirewallPolicyListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.WebApplicationFirewallPolicyListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.WebApplicationFirewallPolicyListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *webApplicationFirewallPoliciesOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *webApplicationFirewallPoliciesOperations) listHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := WebApplicationFirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicyListResult)
}

// listHandleError handles the List error response.
func (client *webApplicationFirewallPoliciesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the WAF policies in a subscription.
func (client *webApplicationFirewallPoliciesOperations) ListAll() (WebApplicationFirewallPolicyListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &webApplicationFirewallPolicyListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *WebApplicationFirewallPolicyListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.WebApplicationFirewallPolicyListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.WebApplicationFirewallPolicyListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *webApplicationFirewallPoliciesOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *webApplicationFirewallPoliciesOperations) listAllHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := WebApplicationFirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicyListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *webApplicationFirewallPoliciesOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
