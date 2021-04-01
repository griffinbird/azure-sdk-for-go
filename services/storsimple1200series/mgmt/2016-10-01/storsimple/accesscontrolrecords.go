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

// AccessControlRecordsClient is the client for the AccessControlRecords methods of the Storsimple service.
type AccessControlRecordsClient struct {
	BaseClient
}

// NewAccessControlRecordsClient creates an instance of the AccessControlRecordsClient client.
func NewAccessControlRecordsClient(subscriptionID string) AccessControlRecordsClient {
	return NewAccessControlRecordsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAccessControlRecordsClientWithBaseURI creates an instance of the AccessControlRecordsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAccessControlRecordsClientWithBaseURI(baseURI string, subscriptionID string) AccessControlRecordsClient {
	return AccessControlRecordsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or Updates an access control record.
// Parameters:
// accessControlRecordName - the name of the access control record.
// accessControlRecord - the access control record to be added or updated.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client AccessControlRecordsClient) CreateOrUpdate(ctx context.Context, accessControlRecordName string, accessControlRecord AccessControlRecord, resourceGroupName string, managerName string) (result AccessControlRecordsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessControlRecordsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accessControlRecord,
			Constraints: []validation.Constraint{{Target: "accessControlRecord.AccessControlRecordProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "accessControlRecord.AccessControlRecordProperties.InitiatorName", Name: validation.Null, Rule: true, Chain: nil}}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.AccessControlRecordsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, accessControlRecordName, accessControlRecord, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AccessControlRecordsClient) CreateOrUpdatePreparer(ctx context.Context, accessControlRecordName string, accessControlRecord AccessControlRecord, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessControlRecordName": autorest.Encode("path", accessControlRecordName),
		"managerName":             autorest.Encode("path", managerName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", pathParameters),
		autorest.WithJSON(accessControlRecord),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AccessControlRecordsClient) CreateOrUpdateSender(req *http.Request) (future AccessControlRecordsCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AccessControlRecordsClient) CreateOrUpdateResponder(resp *http.Response) (result AccessControlRecord, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the access control record.
// Parameters:
// accessControlRecordName - the name of the access control record to delete.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client AccessControlRecordsClient) Delete(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string) (result AccessControlRecordsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessControlRecordsClient.Delete")
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
		return result, validation.NewError("storsimple.AccessControlRecordsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, accessControlRecordName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AccessControlRecordsClient) DeletePreparer(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessControlRecordName": autorest.Encode("path", accessControlRecordName),
		"managerName":             autorest.Encode("path", managerName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AccessControlRecordsClient) DeleteSender(req *http.Request) (future AccessControlRecordsDeleteFuture, err error) {
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
func (client AccessControlRecordsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the properties of the specified access control record name.
// Parameters:
// accessControlRecordName - name of access control record to be fetched.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client AccessControlRecordsClient) Get(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string) (result AccessControlRecord, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessControlRecordsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.AccessControlRecordsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, accessControlRecordName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AccessControlRecordsClient) GetPreparer(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessControlRecordName": autorest.Encode("path", accessControlRecordName),
		"managerName":             autorest.Encode("path", managerName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AccessControlRecordsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AccessControlRecordsClient) GetResponder(resp *http.Response) (result AccessControlRecord, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManager retrieves all the access control records in a manager.
// Parameters:
// resourceGroupName - the resource group name
// managerName - the manager name
func (client AccessControlRecordsClient) ListByManager(ctx context.Context, resourceGroupName string, managerName string) (result AccessControlRecordList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessControlRecordsClient.ListByManager")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.AccessControlRecordsClient", "ListByManager", err.Error())
	}

	req, err := client.ListByManagerPreparer(ctx, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "ListByManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "ListByManager", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.AccessControlRecordsClient", "ListByManager", resp, "Failure responding to request")
		return
	}

	return
}

// ListByManagerPreparer prepares the ListByManager request.
func (client AccessControlRecordsClient) ListByManagerPreparer(ctx context.Context, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagerSender sends the ListByManager request. The method will close the
// http.Response Body if it receives an error.
func (client AccessControlRecordsClient) ListByManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByManagerResponder handles the response to the ListByManager request. The method always
// closes the http.Response Body.
func (client AccessControlRecordsClient) ListByManagerResponder(resp *http.Response) (result AccessControlRecordList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
