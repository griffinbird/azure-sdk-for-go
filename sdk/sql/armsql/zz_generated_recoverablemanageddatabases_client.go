//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RecoverableManagedDatabasesClient contains the methods for the RecoverableManagedDatabases group.
// Don't use this type directly, use NewRecoverableManagedDatabasesClient() instead.
type RecoverableManagedDatabasesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRecoverableManagedDatabasesClient creates a new instance of RecoverableManagedDatabasesClient with the specified values.
func NewRecoverableManagedDatabasesClient(con *arm.Connection, subscriptionID string) *RecoverableManagedDatabasesClient {
	return &RecoverableManagedDatabasesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets a recoverable managed database.
// If the operation fails it returns a generic error.
func (client *RecoverableManagedDatabasesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string, options *RecoverableManagedDatabasesGetOptions) (RecoverableManagedDatabasesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, recoverableDatabaseName, options)
	if err != nil {
		return RecoverableManagedDatabasesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecoverableManagedDatabasesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecoverableManagedDatabasesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecoverableManagedDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string, options *RecoverableManagedDatabasesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/recoverableDatabases/{recoverableDatabaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if recoverableDatabaseName == "" {
		return nil, errors.New("parameter recoverableDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoverableDatabaseName}", url.PathEscape(recoverableDatabaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoverableManagedDatabasesClient) getHandleResponse(resp *http.Response) (RecoverableManagedDatabasesGetResponse, error) {
	result := RecoverableManagedDatabasesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableManagedDatabase); err != nil {
		return RecoverableManagedDatabasesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RecoverableManagedDatabasesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByInstance - Gets a list of recoverable managed databases.
// If the operation fails it returns a generic error.
func (client *RecoverableManagedDatabasesClient) ListByInstance(resourceGroupName string, managedInstanceName string, options *RecoverableManagedDatabasesListByInstanceOptions) *RecoverableManagedDatabasesListByInstancePager {
	return &RecoverableManagedDatabasesListByInstancePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
		},
		advancer: func(ctx context.Context, resp RecoverableManagedDatabasesListByInstanceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RecoverableManagedDatabaseListResult.NextLink)
		},
	}
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *RecoverableManagedDatabasesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *RecoverableManagedDatabasesListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/recoverableDatabases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *RecoverableManagedDatabasesClient) listByInstanceHandleResponse(resp *http.Response) (RecoverableManagedDatabasesListByInstanceResponse, error) {
	result := RecoverableManagedDatabasesListByInstanceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableManagedDatabaseListResult); err != nil {
		return RecoverableManagedDatabasesListByInstanceResponse{}, err
	}
	return result, nil
}

// listByInstanceHandleError handles the ListByInstance error response.
func (client *RecoverableManagedDatabasesClient) listByInstanceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
