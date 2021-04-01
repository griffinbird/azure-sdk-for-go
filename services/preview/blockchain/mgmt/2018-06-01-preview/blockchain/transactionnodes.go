package blockchain

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TransactionNodesClient is the REST API for Azure Blockchain Service
type TransactionNodesClient struct {
	BaseClient
}

// NewTransactionNodesClient creates an instance of the TransactionNodesClient client.
func NewTransactionNodesClient(subscriptionID string) TransactionNodesClient {
	return NewTransactionNodesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTransactionNodesClientWithBaseURI creates an instance of the TransactionNodesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewTransactionNodesClientWithBaseURI(baseURI string, subscriptionID string) TransactionNodesClient {
	return TransactionNodesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create or update the transaction node.
// Parameters:
// blockchainMemberName - blockchain member name.
// transactionNodeName - transaction node name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// transactionNode - payload to create the transaction node.
func (client TransactionNodesClient) Create(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, transactionNode *TransactionNode) (result TransactionNodesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, transactionNode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client TransactionNodesClient) CreatePreparer(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, transactionNode *TransactionNode) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"transactionNodeName":  autorest.Encode("path", transactionNodeName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if transactionNode != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(transactionNode))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionNodesClient) CreateSender(req *http.Request) (future TransactionNodesCreateFuture, err error) {
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
func (client TransactionNodesClient) CreateResponder(resp *http.Response) (result TransactionNode, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the transaction node.
// Parameters:
// blockchainMemberName - blockchain member name.
// transactionNodeName - transaction node name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client TransactionNodesClient) Delete(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (result TransactionNodesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, blockchainMemberName, transactionNodeName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TransactionNodesClient) DeletePreparer(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"transactionNodeName":  autorest.Encode("path", transactionNodeName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionNodesClient) DeleteSender(req *http.Request) (future TransactionNodesDeleteFuture, err error) {
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
func (client TransactionNodesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the details of the transaction node.
// Parameters:
// blockchainMemberName - blockchain member name.
// transactionNodeName - transaction node name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client TransactionNodesClient) Get(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (result TransactionNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, blockchainMemberName, transactionNodeName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TransactionNodesClient) GetPreparer(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"transactionNodeName":  autorest.Encode("path", transactionNodeName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionNodesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TransactionNodesClient) GetResponder(resp *http.Response) (result TransactionNode, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the transaction nodes for a blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client TransactionNodesClient) List(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result TransactionNodeCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.List")
		defer func() {
			sc := -1
			if result.tnc.Response.Response != nil {
				sc = result.tnc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, blockchainMemberName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tnc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "List", resp, "Failure sending request")
		return
	}

	result.tnc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.tnc.hasNextLink() && result.tnc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client TransactionNodesClient) ListPreparer(ctx context.Context, blockchainMemberName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionNodesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TransactionNodesClient) ListResponder(resp *http.Response) (result TransactionNodeCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TransactionNodesClient) listNextResults(ctx context.Context, lastResults TransactionNodeCollection) (result TransactionNodeCollection, err error) {
	req, err := lastResults.transactionNodeCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TransactionNodesClient) ListComplete(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result TransactionNodeCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, blockchainMemberName, resourceGroupName)
	return
}

// ListAPIKeys list the API keys for the transaction node.
// Parameters:
// blockchainMemberName - blockchain member name.
// transactionNodeName - transaction node name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client TransactionNodesClient) ListAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (result APIKeyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.ListAPIKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListAPIKeysPreparer(ctx, blockchainMemberName, transactionNodeName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "ListAPIKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAPIKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "ListAPIKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListAPIKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "ListAPIKeys", resp, "Failure responding to request")
		return
	}

	return
}

// ListAPIKeysPreparer prepares the ListAPIKeys request.
func (client TransactionNodesClient) ListAPIKeysPreparer(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"transactionNodeName":  autorest.Encode("path", transactionNodeName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}/listApiKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAPIKeysSender sends the ListAPIKeys request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionNodesClient) ListAPIKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAPIKeysResponder handles the response to the ListAPIKeys request. The method always
// closes the http.Response Body.
func (client TransactionNodesClient) ListAPIKeysResponder(resp *http.Response) (result APIKeyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListRegenerateAPIKeys regenerate the API keys for the blockchain member.
// Parameters:
// blockchainMemberName - blockchain member name.
// transactionNodeName - transaction node name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// APIKey - api key to be regenerated
func (client TransactionNodesClient) ListRegenerateAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, APIKey *APIKey) (result APIKeyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.ListRegenerateAPIKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListRegenerateAPIKeysPreparer(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, APIKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "ListRegenerateAPIKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRegenerateAPIKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "ListRegenerateAPIKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListRegenerateAPIKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "ListRegenerateAPIKeys", resp, "Failure responding to request")
		return
	}

	return
}

// ListRegenerateAPIKeysPreparer prepares the ListRegenerateAPIKeys request.
func (client TransactionNodesClient) ListRegenerateAPIKeysPreparer(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, APIKey *APIKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"transactionNodeName":  autorest.Encode("path", transactionNodeName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}/regenerateApiKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if APIKey != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(APIKey))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRegenerateAPIKeysSender sends the ListRegenerateAPIKeys request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionNodesClient) ListRegenerateAPIKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListRegenerateAPIKeysResponder handles the response to the ListRegenerateAPIKeys request. The method always
// closes the http.Response Body.
func (client TransactionNodesClient) ListRegenerateAPIKeysResponder(resp *http.Response) (result APIKeyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update the transaction node.
// Parameters:
// blockchainMemberName - blockchain member name.
// transactionNodeName - transaction node name.
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// transactionNode - payload to create the transaction node.
func (client TransactionNodesClient) Update(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, transactionNode *TransactionNodeUpdate) (result TransactionNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionNodesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, blockchainMemberName, transactionNodeName, resourceGroupName, transactionNode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.TransactionNodesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client TransactionNodesClient) UpdatePreparer(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, transactionNode *TransactionNodeUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"blockchainMemberName": autorest.Encode("path", blockchainMemberName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
		"transactionNodeName":  autorest.Encode("path", transactionNodeName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if transactionNode != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(transactionNode))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionNodesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client TransactionNodesClient) UpdateResponder(resp *http.Response) (result TransactionNode, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
