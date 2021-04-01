package storagesync

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CloudEndpointsClient is the microsoft Storage Sync Service API
type CloudEndpointsClient struct {
	BaseClient
}

// NewCloudEndpointsClient creates an instance of the CloudEndpointsClient client.
func NewCloudEndpointsClient(subscriptionID string) CloudEndpointsClient {
	return NewCloudEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCloudEndpointsClientWithBaseURI creates an instance of the CloudEndpointsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCloudEndpointsClientWithBaseURI(baseURI string, subscriptionID string) CloudEndpointsClient {
	return CloudEndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a new CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
// parameters - body of Cloud Endpoint resource.
func (client CloudEndpointsClient) Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters CloudEndpointCreateParameters) (result CloudEndpointsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client CloudEndpointsClient) CreatePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters CloudEndpointCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) CreateSender(req *http.Request) (future CloudEndpointsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) CreateResponder(resp *http.Response) (result CloudEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a given CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
func (client CloudEndpointsClient) Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (result CloudEndpointsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CloudEndpointsClient) DeletePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) DeleteSender(req *http.Request) (future CloudEndpointsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a given CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
func (client CloudEndpointsClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (result CloudEndpoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CloudEndpointsClient) GetPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) GetResponder(resp *http.Response) (result CloudEndpoint, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySyncGroup get a CloudEndpoint List.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
func (client CloudEndpointsClient) ListBySyncGroup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result CloudEndpointArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.ListBySyncGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "ListBySyncGroup", err.Error())
	}

	req, err := client.ListBySyncGroupPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "ListBySyncGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySyncGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "ListBySyncGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySyncGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "ListBySyncGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListBySyncGroupPreparer prepares the ListBySyncGroup request.
func (client CloudEndpointsClient) ListBySyncGroupPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySyncGroupSender sends the ListBySyncGroup request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) ListBySyncGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySyncGroupResponder handles the response to the ListBySyncGroup request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) ListBySyncGroupResponder(resp *http.Response) (result CloudEndpointArray, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PostBackup post Backup a given CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
// parameters - body of Backup request.
func (client CloudEndpointsClient) PostBackup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters BackupRequest) (result CloudEndpointsPostBackupFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.PostBackup")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "PostBackup", err.Error())
	}

	req, err := client.PostBackupPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PostBackup", nil, "Failure preparing request")
		return
	}

	result, err = client.PostBackupSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PostBackup", nil, "Failure sending request")
		return
	}

	return
}

// PostBackupPreparer prepares the PostBackup request.
func (client CloudEndpointsClient) PostBackupPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters BackupRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}/postbackup", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostBackupSender sends the PostBackup request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) PostBackupSender(req *http.Request) (future CloudEndpointsPostBackupFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// PostBackupResponder handles the response to the PostBackup request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) PostBackupResponder(resp *http.Response) (result PostBackupResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PostRestore post Restore a given CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
// parameters - body of Cloud Endpoint object.
func (client CloudEndpointsClient) PostRestore(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters PostRestoreRequest) (result CloudEndpointsPostRestoreFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.PostRestore")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "PostRestore", err.Error())
	}

	req, err := client.PostRestorePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PostRestore", nil, "Failure preparing request")
		return
	}

	result, err = client.PostRestoreSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PostRestore", nil, "Failure sending request")
		return
	}

	return
}

// PostRestorePreparer prepares the PostRestore request.
func (client CloudEndpointsClient) PostRestorePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters PostRestoreRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}/postrestore", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostRestoreSender sends the PostRestore request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) PostRestoreSender(req *http.Request) (future CloudEndpointsPostRestoreFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// PostRestoreResponder handles the response to the PostRestore request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) PostRestoreResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PreBackup pre Backup a given CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
// parameters - body of Backup request.
func (client CloudEndpointsClient) PreBackup(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters BackupRequest) (result CloudEndpointsPreBackupFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.PreBackup")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "PreBackup", err.Error())
	}

	req, err := client.PreBackupPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PreBackup", nil, "Failure preparing request")
		return
	}

	result, err = client.PreBackupSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PreBackup", nil, "Failure sending request")
		return
	}

	return
}

// PreBackupPreparer prepares the PreBackup request.
func (client CloudEndpointsClient) PreBackupPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters BackupRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}/prebackup", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PreBackupSender sends the PreBackup request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) PreBackupSender(req *http.Request) (future CloudEndpointsPreBackupFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// PreBackupResponder handles the response to the PreBackup request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) PreBackupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PreRestore pre Restore a given CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
// parameters - body of Cloud Endpoint object.
func (client CloudEndpointsClient) PreRestore(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters PreRestoreRequest) (result CloudEndpointsPreRestoreFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.PreRestore")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "PreRestore", err.Error())
	}

	req, err := client.PreRestorePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PreRestore", nil, "Failure preparing request")
		return
	}

	result, err = client.PreRestoreSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "PreRestore", nil, "Failure sending request")
		return
	}

	return
}

// PreRestorePreparer prepares the PreRestore request.
func (client CloudEndpointsClient) PreRestorePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string, parameters PreRestoreRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}/prerestore", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PreRestoreSender sends the PreRestore request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) PreRestoreSender(req *http.Request) (future CloudEndpointsPreRestoreFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// PreRestoreResponder handles the response to the PreRestore request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) PreRestoreResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Restoreheartbeat restore Heartbeat a given CloudEndpoint.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// cloudEndpointName - name of Cloud Endpoint object.
func (client CloudEndpointsClient) Restoreheartbeat(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudEndpointsClient.Restoreheartbeat")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.CloudEndpointsClient", "Restoreheartbeat", err.Error())
	}

	req, err := client.RestoreheartbeatPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Restoreheartbeat", nil, "Failure preparing request")
		return
	}

	resp, err := client.RestoreheartbeatSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Restoreheartbeat", resp, "Failure sending request")
		return
	}

	result, err = client.RestoreheartbeatResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.CloudEndpointsClient", "Restoreheartbeat", resp, "Failure responding to request")
		return
	}

	return
}

// RestoreheartbeatPreparer prepares the Restoreheartbeat request.
func (client CloudEndpointsClient) RestoreheartbeatPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, cloudEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudEndpointName":      autorest.Encode("path", cloudEndpointName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}/restoreheartbeat", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestoreheartbeatSender sends the Restoreheartbeat request. The method will close the
// http.Response Body if it receives an error.
func (client CloudEndpointsClient) RestoreheartbeatSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RestoreheartbeatResponder handles the response to the Restoreheartbeat request. The method always
// closes the http.Response Body.
func (client CloudEndpointsClient) RestoreheartbeatResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
