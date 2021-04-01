package storsimple

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

// BackupsClient is the client for the Backups methods of the Storsimple service.
type BackupsClient struct {
	BaseClient
}

// NewBackupsClient creates an instance of the BackupsClient client.
func NewBackupsClient(subscriptionID string) BackupsClient {
	return NewBackupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupsClientWithBaseURI creates an instance of the BackupsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return BackupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Clone clones the given backup element to a new disk or share with given details.
// Parameters:
// deviceName - the device name.
// backupName - the backup name.
// elementName - the backup element name.
// cloneRequest - the clone request.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupsClient) Clone(ctx context.Context, deviceName string, backupName string, elementName string, cloneRequest CloneRequest, resourceGroupName string, managerName string) (result BackupsCloneFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Clone")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cloneRequest,
			Constraints: []validation.Constraint{{Target: "cloneRequest.CloneRequestProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "cloneRequest.CloneRequestProperties.TargetDeviceID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "cloneRequest.CloneRequestProperties.TargetAccessPointID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "cloneRequest.CloneRequestProperties.NewEndpointName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "cloneRequest.CloneRequestProperties.Share", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "cloneRequest.CloneRequestProperties.Share.FileShareProperties", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "cloneRequest.CloneRequestProperties.Share.FileShareProperties.AdminUser", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "cloneRequest.CloneRequestProperties.Share.FileShareProperties.ProvisionedCapacityInBytes", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
					{Target: "cloneRequest.CloneRequestProperties.Disk", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "cloneRequest.CloneRequestProperties.Disk.ISCSIDiskProperties", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "cloneRequest.CloneRequestProperties.Disk.ISCSIDiskProperties.AccessControlRecords", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "cloneRequest.CloneRequestProperties.Disk.ISCSIDiskProperties.ProvisionedCapacityInBytes", Name: validation.Null, Rule: true, Chain: nil},
							}},
						}},
				}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "Clone", err.Error())
	}

	req, err := client.ClonePreparer(ctx, deviceName, backupName, elementName, cloneRequest, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Clone", nil, "Failure preparing request")
		return
	}

	result, err = client.CloneSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Clone", nil, "Failure sending request")
		return
	}

	return
}

// ClonePreparer prepares the Clone request.
func (client BackupsClient) ClonePreparer(ctx context.Context, deviceName string, backupName string, elementName string, cloneRequest CloneRequest, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupName":        autorest.Encode("path", backupName),
		"deviceName":        autorest.Encode("path", deviceName),
		"elementName":       autorest.Encode("path", elementName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/elements/{elementName}/clone", pathParameters),
		autorest.WithJSON(cloneRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CloneSender sends the Clone request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) CloneSender(req *http.Request) (future BackupsCloneFuture, err error) {
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

// CloneResponder handles the response to the Clone request. The method always
// closes the http.Response Body.
func (client BackupsClient) CloneResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the backup.
// Parameters:
// deviceName - the device name.
// backupName - the backup name.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupsClient) Delete(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string) (result BackupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, deviceName, backupName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackupsClient) DeletePreparer(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupName":        autorest.Encode("path", backupName),
		"deviceName":        autorest.Encode("path", deviceName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) DeleteSender(req *http.Request) (future BackupsDeleteFuture, err error) {
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
func (client BackupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByDevice retrieves all the backups in a device. Can be used to get the backups for failover also.
// Parameters:
// deviceName - the device name.
// resourceGroupName - the resource group name
// managerName - the manager name
// forFailover - set to true if you need backups which can be used for failover.
// filter - oData Filter options
func (client BackupsClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string, forFailover *bool, filter string) (result BackupListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.ListByDevice")
		defer func() {
			sc := -1
			if result.bl.Response.Response != nil {
				sc = result.bl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "ListByDevice", err.Error())
	}

	result.fn = client.listByDeviceNextResults
	req, err := client.ListByDevicePreparer(ctx, deviceName, resourceGroupName, managerName, forFailover, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.bl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result.bl, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByDevice", resp, "Failure responding to request")
		return
	}
	if result.bl.hasNextLink() && result.bl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client BackupsClient) ListByDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string, managerName string, forFailover *bool, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if forFailover != nil {
		queryParameters["forFailover"] = autorest.Encode("query", *forFailover)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client BackupsClient) ListByDeviceResponder(resp *http.Response) (result BackupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDeviceNextResults retrieves the next set of results, if any.
func (client BackupsClient) listByDeviceNextResults(ctx context.Context, lastResults BackupList) (result BackupList, err error) {
	req, err := lastResults.backupListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByDeviceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByDeviceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByDeviceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDeviceComplete enumerates all values, automatically crossing page boundaries as required.
func (client BackupsClient) ListByDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string, managerName string, forFailover *bool, filter string) (result BackupListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.ListByDevice")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDevice(ctx, deviceName, resourceGroupName, managerName, forFailover, filter)
	return
}

// ListByManager retrieves all the backups in a manager.
// Parameters:
// resourceGroupName - the resource group name
// managerName - the manager name
// filter - oData Filter options
func (client BackupsClient) ListByManager(ctx context.Context, resourceGroupName string, managerName string, filter string) (result BackupListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.ListByManager")
		defer func() {
			sc := -1
			if result.bl.Response.Response != nil {
				sc = result.bl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupsClient", "ListByManager", err.Error())
	}

	result.fn = client.listByManagerNextResults
	req, err := client.ListByManagerPreparer(ctx, resourceGroupName, managerName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.bl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByManager", resp, "Failure sending request")
		return
	}

	result.bl, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "ListByManager", resp, "Failure responding to request")
		return
	}
	if result.bl.hasNextLink() && result.bl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByManagerPreparer prepares the ListByManager request.
func (client BackupsClient) ListByManagerPreparer(ctx context.Context, resourceGroupName string, managerName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/backups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagerSender sends the ListByManager request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) ListByManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByManagerResponder handles the response to the ListByManager request. The method always
// closes the http.Response Body.
func (client BackupsClient) ListByManagerResponder(resp *http.Response) (result BackupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByManagerNextResults retrieves the next set of results, if any.
func (client BackupsClient) listByManagerNextResults(ctx context.Context, lastResults BackupList) (result BackupList, err error) {
	req, err := lastResults.backupListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByManagerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByManagerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupsClient", "listByManagerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByManagerComplete enumerates all values, automatically crossing page boundaries as required.
func (client BackupsClient) ListByManagerComplete(ctx context.Context, resourceGroupName string, managerName string, filter string) (result BackupListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.ListByManager")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByManager(ctx, resourceGroupName, managerName, filter)
	return
}
