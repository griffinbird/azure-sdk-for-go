package labservices

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

// ProviderOperationsClient is the the Managed Labs Client.
type ProviderOperationsClient struct {
	BaseClient
}

// NewProviderOperationsClient creates an instance of the ProviderOperationsClient client.
func NewProviderOperationsClient(subscriptionID string) ProviderOperationsClient {
	return NewProviderOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderOperationsClientWithBaseURI creates an instance of the ProviderOperationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewProviderOperationsClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsClient {
	return ProviderOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List result of the request to list REST API operations
func (client ProviderOperationsClient) List(ctx context.Context) (result ProviderOperationResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderOperationsClient.List")
		defer func() {
			sc := -1
			if result.por.Response.Response != nil {
				sc = result.por.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.ProviderOperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.por.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.ProviderOperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.por, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.ProviderOperationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.por.hasNextLink() && result.por.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ProviderOperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.LabServices/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProviderOperationsClient) ListResponder(resp *http.Response) (result ProviderOperationResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProviderOperationsClient) listNextResults(ctx context.Context, lastResults ProviderOperationResult) (result ProviderOperationResult, err error) {
	req, err := lastResults.providerOperationResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.ProviderOperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.ProviderOperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.ProviderOperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderOperationsClient) ListComplete(ctx context.Context) (result ProviderOperationResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderOperationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
