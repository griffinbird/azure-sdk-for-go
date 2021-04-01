package policy

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

// AssignmentsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type AssignmentsClient struct {
	BaseClient
}

// NewAssignmentsClient creates an instance of the AssignmentsClient client.
func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return NewAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAssignmentsClientWithBaseURI creates an instance of the AssignmentsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return AssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create policy assignments are inherited by child resources. For example, when you apply a policy to a resource group
// that policy is assigned to all resources in the group.
// Parameters:
// scope - the scope of the policy assignment.
// policyAssignmentName - the name of the policy assignment.
// parameters - parameters for the policy assignment.
func (client AssignmentsClient) Create(ctx context.Context, scope string, policyAssignmentName string, parameters Assignment) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, scope, policyAssignmentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AssignmentsClient) CreatePreparer(ctx context.Context, scope string, policyAssignmentName string, parameters Assignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": autorest.Encode("path", policyAssignmentName),
		"scope":                scope,
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) CreateResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateByID policy assignments are inherited by child resources. For example, when you apply a policy to a resource
// group that policy is assigned to all resources in the group. When providing a scope for the assignment, use
// '/subscriptions/{subscription-id}/' for subscriptions,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for resource groups, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider-namespace}/{resource-type}/{resource-name}'
// for resources.
// Parameters:
// policyAssignmentID - the ID of the policy assignment to create. Use the format
// '/{scope}/providers/Microsoft.Authorization/policyAssignments/{policy-assignment-name}'.
// parameters - parameters for policy assignment.
func (client AssignmentsClient) CreateByID(ctx context.Context, policyAssignmentID string, parameters Assignment) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.CreateByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateByIDPreparer(ctx, policyAssignmentID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "CreateByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "CreateByID", resp, "Failure sending request")
		return
	}

	result, err = client.CreateByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "CreateByID", resp, "Failure responding to request")
		return
	}

	return
}

// CreateByIDPreparer prepares the CreateByID request.
func (client AssignmentsClient) CreateByIDPreparer(ctx context.Context, policyAssignmentID string, parameters Assignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{policyAssignmentId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateByIDSender sends the CreateByID request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) CreateByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateByIDResponder handles the response to the CreateByID request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) CreateByIDResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a policy assignment.
// Parameters:
// scope - the scope of the policy assignment.
// policyAssignmentName - the name of the policy assignment to delete.
func (client AssignmentsClient) Delete(ctx context.Context, scope string, policyAssignmentName string) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, policyAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AssignmentsClient) DeletePreparer(ctx context.Context, scope string, policyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": autorest.Encode("path", policyAssignmentName),
		"scope":                scope,
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) DeleteResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteByID when providing a scope for the assignment, use '/subscriptions/{subscription-id}/' for subscriptions,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for resource groups, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider-namespace}/{resource-type}/{resource-name}'
// for resources.
// Parameters:
// policyAssignmentID - the ID of the policy assignment to delete. Use the format
// '/{scope}/providers/Microsoft.Authorization/policyAssignments/{policy-assignment-name}'.
func (client AssignmentsClient) DeleteByID(ctx context.Context, policyAssignmentID string) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.DeleteByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByIDPreparer(ctx, policyAssignmentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "DeleteByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "DeleteByID", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "DeleteByID", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client AssignmentsClient) DeleteByIDPreparer(ctx context.Context, policyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{policyAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) DeleteByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) DeleteByIDResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a policy assignment.
// Parameters:
// scope - the scope of the policy assignment.
// policyAssignmentName - the name of the policy assignment to get.
func (client AssignmentsClient) Get(ctx context.Context, scope string, policyAssignmentName string) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, policyAssignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssignmentsClient) GetPreparer(ctx context.Context, scope string, policyAssignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentName": autorest.Encode("path", policyAssignmentName),
		"scope":                scope,
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) GetResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID when providing a scope for the assignment, use '/subscriptions/{subscription-id}/' for subscriptions,
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for resource groups, and
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider-namespace}/{resource-type}/{resource-name}'
// for resources.
// Parameters:
// policyAssignmentID - the ID of the policy assignment to get. Use the format
// '/{scope}/providers/Microsoft.Authorization/policyAssignments/{policy-assignment-name}'.
func (client AssignmentsClient) GetByID(ctx context.Context, policyAssignmentID string) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByIDPreparer(ctx, policyAssignmentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "GetByID", resp, "Failure responding to request")
		return
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client AssignmentsClient) GetByIDPreparer(ctx context.Context, policyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyAssignmentId": policyAssignmentID,
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{policyAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) GetByIDResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the policy assignments for a subscription.
// Parameters:
// filter - the filter to apply on the operation.
func (client AssignmentsClient) List(ctx context.Context, filter string) (result AssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.List")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AssignmentsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) ListResponder(resp *http.Response) (result AssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AssignmentsClient) listNextResults(ctx context.Context, lastResults AssignmentListResult) (result AssignmentListResult, err error) {
	req, err := lastResults.assignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssignmentsClient) ListComplete(ctx context.Context, filter string) (result AssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}

// ListForResource gets policy assignments for a resource.
// Parameters:
// resourceGroupName - the name of the resource group containing the resource. The name is case insensitive.
// resourceProviderNamespace - the namespace of the resource provider.
// parentResourcePath - the parent resource path.
// resourceType - the resource type.
// resourceName - the name of the resource with policy assignments.
// filter - the filter to apply on the operation.
func (client AssignmentsClient) ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result AssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.ListForResource")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("policy.AssignmentsClient", "ListForResource", err.Error())
	}

	result.fn = client.listForResourceNextResults
	req, err := client.ListForResourcePreparer(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResource", resp, "Failure responding to request")
		return
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForResourcePreparer prepares the ListForResource request.
func (client AssignmentsClient) ListForResourcePreparer(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/policyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForResourceSender sends the ListForResource request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) ListForResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceResponder handles the response to the ListForResource request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) ListForResourceResponder(resp *http.Response) (result AssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForResourceNextResults retrieves the next set of results, if any.
func (client AssignmentsClient) listForResourceNextResults(ctx context.Context, lastResults AssignmentListResult) (result AssignmentListResult, err error) {
	req, err := lastResults.assignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listForResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listForResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listForResourceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssignmentsClient) ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result AssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.ListForResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForResource(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	return
}

// ListForResourceGroup gets policy assignments for the resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains policy assignments.
// filter - the filter to apply on the operation.
func (client AssignmentsClient) ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result AssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.ListForResourceGroup")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("policy.AssignmentsClient", "ListForResourceGroup", err.Error())
	}

	result.fn = client.listForResourceGroupNextResults
	req, err := client.ListForResourceGroupPreparer(ctx, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "ListForResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client AssignmentsClient) ListForResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForResourceGroupSender sends the ListForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) ListForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceGroupResponder handles the response to the ListForResourceGroup request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) ListForResourceGroupResponder(resp *http.Response) (result AssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForResourceGroupNextResults retrieves the next set of results, if any.
func (client AssignmentsClient) listForResourceGroupNextResults(ctx context.Context, lastResults AssignmentListResult) (result AssignmentListResult, err error) {
	req, err := lastResults.assignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listForResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listForResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.AssignmentsClient", "listForResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssignmentsClient) ListForResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result AssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.ListForResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForResourceGroup(ctx, resourceGroupName, filter)
	return
}
