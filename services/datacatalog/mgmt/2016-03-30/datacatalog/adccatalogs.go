package datacatalog

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

// ADCCatalogsClient is the the Azure Data Catalog management API provides a RESTful set of web services that interact
// with Azure Data Catalog services.
type ADCCatalogsClient struct {
	BaseClient
}

// NewADCCatalogsClient creates an instance of the ADCCatalogsClient client.
func NewADCCatalogsClient(subscriptionID string, catalogName string) ADCCatalogsClient {
	return NewADCCatalogsClientWithBaseURI(DefaultBaseURI, subscriptionID, catalogName)
}

// NewADCCatalogsClientWithBaseURI creates an instance of the ADCCatalogsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewADCCatalogsClientWithBaseURI(baseURI string, subscriptionID string, catalogName string) ADCCatalogsClient {
	return ADCCatalogsClient{NewWithBaseURI(baseURI, subscriptionID, catalogName)}
}

// CreateOrUpdate the Create Azure Data Catalog service operation creates a new data catalog service with the specified
// parameters. If the specific service already exists, then any patchable properties will be updated and any immutable
// properties will remain unchanged.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// properties - properties supplied to the Create or Update a data catalog.
func (client ADCCatalogsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, properties ADCCatalog) (result ADCCatalog, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ADCCatalogsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datacatalog.ADCCatalogsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, properties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ADCCatalogsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, properties ADCCatalog) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"catalogName":       autorest.Encode("path", client.CatalogName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", pathParameters),
		autorest.WithJSON(properties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ADCCatalogsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ADCCatalogsClient) CreateOrUpdateResponder(resp *http.Response) (result ADCCatalog, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete Azure Data Catalog Service operation deletes an existing data catalog.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
func (client ADCCatalogsClient) Delete(ctx context.Context, resourceGroupName string) (result ADCCatalogsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ADCCatalogsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datacatalog.ADCCatalogsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ADCCatalogsClient) DeletePreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"catalogName":       autorest.Encode("path", client.CatalogName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ADCCatalogsClient) DeleteSender(req *http.Request) (future ADCCatalogsDeleteFuture, err error) {
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
func (client ADCCatalogsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get Azure Data Catalog Service operation retrieves a json representation of the data catalog.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
func (client ADCCatalogsClient) Get(ctx context.Context, resourceGroupName string) (result ADCCatalog, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ADCCatalogsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datacatalog.ADCCatalogsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ADCCatalogsClient) GetPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"catalogName":       autorest.Encode("path", client.CatalogName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ADCCatalogsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ADCCatalogsClient) GetResponder(resp *http.Response) (result ADCCatalog, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListtByResourceGroup the List catalogs in Resource Group operation lists all the Azure Data Catalogs available under
// the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
func (client ADCCatalogsClient) ListtByResourceGroup(ctx context.Context, resourceGroupName string) (result ADCCatalogsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ADCCatalogsClient.ListtByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datacatalog.ADCCatalogsClient", "ListtByResourceGroup", err.Error())
	}

	req, err := client.ListtByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "ListtByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListtByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "ListtByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListtByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "ListtByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListtByResourceGroupPreparer prepares the ListtByResourceGroup request.
func (client ADCCatalogsClient) ListtByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListtByResourceGroupSender sends the ListtByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ADCCatalogsClient) ListtByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListtByResourceGroupResponder handles the response to the ListtByResourceGroup request. The method always
// closes the http.Response Body.
func (client ADCCatalogsClient) ListtByResourceGroupResponder(resp *http.Response) (result ADCCatalogsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update the Update Azure Data Catalog Service operation can be used to update the existing deployment. The update
// call only supports the properties listed in the PATCH body.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// properties - properties supplied to the Update a data catalog.
func (client ADCCatalogsClient) Update(ctx context.Context, resourceGroupName string, properties ADCCatalog) (result ADCCatalog, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ADCCatalogsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datacatalog.ADCCatalogsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, properties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacatalog.ADCCatalogsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ADCCatalogsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, properties ADCCatalog) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"catalogName":       autorest.Encode("path", client.CatalogName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", pathParameters),
		autorest.WithJSON(properties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ADCCatalogsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ADCCatalogsClient) UpdateResponder(resp *http.Response) (result ADCCatalog, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
