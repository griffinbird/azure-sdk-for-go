package subscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/subscription/mgmt/2017-11-01-preview/subscription"

// Definition the subscription definition used to create the subscription.
type Definition struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// DefinitionProperties - the subscription definition properties
	*DefinitionProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Definition.
func (d Definition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DefinitionProperties != nil {
		objectMap["properties"] = d.DefinitionProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Definition struct.
func (d *Definition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				d.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				d.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				d.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var definitionProperties DefinitionProperties
				err = json.Unmarshal(*v, &definitionProperties)
				if err != nil {
					return err
				}
				d.DefinitionProperties = &definitionProperties
			}
		}
	}

	return nil
}

// DefinitionList subscription Definition List operation response.
type DefinitionList struct {
	autorest.Response `json:"-"`
	// Value - An array of subscriptionDefinitions
	Value *[]Definition `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// DefinitionListIterator provides access to a complete listing of Definition values.
type DefinitionListIterator struct {
	i    int
	page DefinitionListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DefinitionListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *DefinitionListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DefinitionListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DefinitionListIterator) Response() DefinitionList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DefinitionListIterator) Value() Definition {
	if !iter.page.NotDone() {
		return Definition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DefinitionListIterator type.
func NewDefinitionListIterator(page DefinitionListPage) DefinitionListIterator {
	return DefinitionListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dl DefinitionList) IsEmpty() bool {
	return dl.Value == nil || len(*dl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (dl DefinitionList) hasNextLink() bool {
	return dl.NextLink != nil && len(*dl.NextLink) != 0
}

// definitionListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dl DefinitionList) definitionListPreparer(ctx context.Context) (*http.Request, error) {
	if !dl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dl.NextLink)))
}

// DefinitionListPage contains a page of Definition values.
type DefinitionListPage struct {
	fn func(context.Context, DefinitionList) (DefinitionList, error)
	dl DefinitionList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DefinitionListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.dl)
		if err != nil {
			return err
		}
		page.dl = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DefinitionListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DefinitionListPage) NotDone() bool {
	return !page.dl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DefinitionListPage) Response() DefinitionList {
	return page.dl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DefinitionListPage) Values() []Definition {
	if page.dl.IsEmpty() {
		return nil
	}
	return *page.dl.Value
}

// Creates a new instance of the DefinitionListPage type.
func NewDefinitionListPage(cur DefinitionList, getNextPage func(context.Context, DefinitionList) (DefinitionList, error)) DefinitionListPage {
	return DefinitionListPage{
		fn: getNextPage,
		dl: cur,
	}
}

// DefinitionProperties the subscription definition properties.
type DefinitionProperties struct {
	// SubscriptionID - READ-ONLY; The ID of the subscription.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// SubscriptionDisplayName - The display name of the subscription.
	SubscriptionDisplayName *string `json:"subscriptionDisplayName,omitempty"`
	// OfferType - The offer type of the subscription. For example, MS-AZR-0017P (EnterpriseAgreement) and MS-AZR-0148P (EnterpriseAgreement devTest) are available.
	OfferType *string `json:"offerType,omitempty"`
	// Etag - The etag the subscription definition.
	Etag *string `json:"etag,omitempty"`
}

// MarshalJSON is the custom marshaler for DefinitionProperties.
func (dp DefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dp.SubscriptionDisplayName != nil {
		objectMap["subscriptionDisplayName"] = dp.SubscriptionDisplayName
	}
	if dp.OfferType != nil {
		objectMap["offerType"] = dp.OfferType
	}
	if dp.Etag != nil {
		objectMap["etag"] = dp.Etag
	}
	return json.Marshal(objectMap)
}

// DefinitionsCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type DefinitionsCreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(DefinitionsClient) (Definition, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *DefinitionsCreateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for DefinitionsCreateFuture.Result.
func (future *DefinitionsCreateFuture) result(client DefinitionsClient) (d Definition, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.DefinitionsCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("subscription.DefinitionsCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if d.Response.Response, err = future.GetResult(sender); err == nil && d.Response.Response.StatusCode != http.StatusNoContent {
		d, err = client.CreateResponder(d.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.DefinitionsCreateFuture", "Result", d.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// Operation REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Subscription
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}
