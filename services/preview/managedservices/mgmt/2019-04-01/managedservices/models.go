package managedservices

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
	"github.com/gofrs/uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/managedservices/mgmt/2019-04-01/managedservices"

// Authorization authorization tuple containing principal Id (of user/service principal/security group) and
// role definition id.
type Authorization struct {
	// PrincipalID - Principal Id of the security group/service principal/user that would be assigned permissions to the projected subscription
	PrincipalID *string `json:"principalId,omitempty"`
	// PrincipalIDDisplayName - Display name of the principal Id.
	PrincipalIDDisplayName *string `json:"principalIdDisplayName,omitempty"`
	// RoleDefinitionID - The role definition identifier. This role will define all the permissions that the security group/service principal/user must have on the projected subscription. This role cannot be an owner role.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// DelegatedRoleDefinitionIds - The delegatedRoleDefinitionIds field is required when the roleDefinitionId refers to the User Access Administrator Role. It is the list of role definition ids which define all the permissions that the user in the authorization can assign to other security groups/service principals/users.
	DelegatedRoleDefinitionIds *[]uuid.UUID `json:"delegatedRoleDefinitionIds,omitempty"`
}

// ErrorDefinition error response indicates Azure Resource Manager is not able to process the incoming
// request. The reason is provided in the error message.
type ErrorDefinition struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
	// Details - Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// Operation object that describes a single Microsoft.ManagedServices operation.
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - READ-ONLY; The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.ManagedServices
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Registration definition, registration assignment etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationList list of the operations.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of Microsoft.ManagedServices operations.
	Value *[]Operation `json:"value,omitempty"`
}

// Plan plan details for the managed services.
type Plan struct {
	// Name - The plan name.
	Name *string `json:"name,omitempty"`
	// Publisher - The publisher ID.
	Publisher *string `json:"publisher,omitempty"`
	// Product - The product code.
	Product *string `json:"product,omitempty"`
	// Version - The plan's version.
	Version *string `json:"version,omitempty"`
}

// RegistrationAssignment registration assignment.
type RegistrationAssignment struct {
	autorest.Response `json:"-"`
	// Properties - Properties of a registration assignment.
	Properties *RegistrationAssignmentProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The fully qualified path of the registration assignment.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Type of the resource.
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Name of the registration assignment.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for RegistrationAssignment.
func (ra RegistrationAssignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ra.Properties != nil {
		objectMap["properties"] = ra.Properties
	}
	return json.Marshal(objectMap)
}

// RegistrationAssignmentList list of registration assignments.
type RegistrationAssignmentList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of registration assignments.
	Value *[]RegistrationAssignment `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to next page of registration assignments.
	NextLink *string `json:"nextLink,omitempty"`
}

// RegistrationAssignmentListIterator provides access to a complete listing of RegistrationAssignment
// values.
type RegistrationAssignmentListIterator struct {
	i    int
	page RegistrationAssignmentListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RegistrationAssignmentListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationAssignmentListIterator.NextWithContext")
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
func (iter *RegistrationAssignmentListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RegistrationAssignmentListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RegistrationAssignmentListIterator) Response() RegistrationAssignmentList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RegistrationAssignmentListIterator) Value() RegistrationAssignment {
	if !iter.page.NotDone() {
		return RegistrationAssignment{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the RegistrationAssignmentListIterator type.
func NewRegistrationAssignmentListIterator(page RegistrationAssignmentListPage) RegistrationAssignmentListIterator {
	return RegistrationAssignmentListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ral RegistrationAssignmentList) IsEmpty() bool {
	return ral.Value == nil || len(*ral.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ral RegistrationAssignmentList) hasNextLink() bool {
	return ral.NextLink != nil && len(*ral.NextLink) != 0
}

// registrationAssignmentListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ral RegistrationAssignmentList) registrationAssignmentListPreparer(ctx context.Context) (*http.Request, error) {
	if !ral.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ral.NextLink)))
}

// RegistrationAssignmentListPage contains a page of RegistrationAssignment values.
type RegistrationAssignmentListPage struct {
	fn  func(context.Context, RegistrationAssignmentList) (RegistrationAssignmentList, error)
	ral RegistrationAssignmentList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RegistrationAssignmentListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationAssignmentListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ral)
		if err != nil {
			return err
		}
		page.ral = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *RegistrationAssignmentListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RegistrationAssignmentListPage) NotDone() bool {
	return !page.ral.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RegistrationAssignmentListPage) Response() RegistrationAssignmentList {
	return page.ral
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RegistrationAssignmentListPage) Values() []RegistrationAssignment {
	if page.ral.IsEmpty() {
		return nil
	}
	return *page.ral.Value
}

// Creates a new instance of the RegistrationAssignmentListPage type.
func NewRegistrationAssignmentListPage(cur RegistrationAssignmentList, getNextPage func(context.Context, RegistrationAssignmentList) (RegistrationAssignmentList, error)) RegistrationAssignmentListPage {
	return RegistrationAssignmentListPage{
		fn:  getNextPage,
		ral: cur,
	}
}

// RegistrationAssignmentProperties properties of a registration assignment.
type RegistrationAssignmentProperties struct {
	// RegistrationDefinitionID - Fully qualified path of the registration definition.
	RegistrationDefinitionID *string `json:"registrationDefinitionId,omitempty"`
	// ProvisioningState - READ-ONLY; Current state of the registration assignment. Possible values include: 'NotSpecified', 'Accepted', 'Running', 'Ready', 'Creating', 'Created', 'Deleting', 'Deleted', 'Canceled', 'Failed', 'Succeeded', 'Updating'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// RegistrationDefinition - READ-ONLY; Registration definition inside registration assignment.
	RegistrationDefinition *RegistrationAssignmentPropertiesRegistrationDefinition `json:"registrationDefinition,omitempty"`
}

// MarshalJSON is the custom marshaler for RegistrationAssignmentProperties.
func (rap RegistrationAssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rap.RegistrationDefinitionID != nil {
		objectMap["registrationDefinitionId"] = rap.RegistrationDefinitionID
	}
	return json.Marshal(objectMap)
}

// RegistrationAssignmentPropertiesRegistrationDefinition registration definition inside registration
// assignment.
type RegistrationAssignmentPropertiesRegistrationDefinition struct {
	// Properties - Properties of registration definition inside registration assignment.
	Properties *RegistrationAssignmentPropertiesRegistrationDefinitionProperties `json:"properties,omitempty"`
	// Plan - Plan details for the managed services.
	Plan *Plan `json:"plan,omitempty"`
	// ID - READ-ONLY; Fully qualified path of the registration definition.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Type of the resource (Microsoft.ManagedServices/registrationDefinitions).
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Name of the registration definition.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for RegistrationAssignmentPropertiesRegistrationDefinition.
func (rapD RegistrationAssignmentPropertiesRegistrationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rapD.Properties != nil {
		objectMap["properties"] = rapD.Properties
	}
	if rapD.Plan != nil {
		objectMap["plan"] = rapD.Plan
	}
	return json.Marshal(objectMap)
}

// RegistrationAssignmentPropertiesRegistrationDefinitionProperties properties of registration definition
// inside registration assignment.
type RegistrationAssignmentPropertiesRegistrationDefinitionProperties struct {
	// Description - Description of the registration definition.
	Description *string `json:"description,omitempty"`
	// Authorizations - Authorization tuple containing principal id of the user/security group or service principal and id of the build-in role.
	Authorizations *[]Authorization `json:"authorizations,omitempty"`
	// RegistrationDefinitionName - Name of the registration definition.
	RegistrationDefinitionName *string `json:"registrationDefinitionName,omitempty"`
	// ProvisioningState - Current state of the registration definition. Possible values include: 'NotSpecified', 'Accepted', 'Running', 'Ready', 'Creating', 'Created', 'Deleting', 'Deleted', 'Canceled', 'Failed', 'Succeeded', 'Updating'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// ManageeTenantID - Id of the home tenant.
	ManageeTenantID *string `json:"manageeTenantId,omitempty"`
	// ManageeTenantName - Name of the home tenant.
	ManageeTenantName *string `json:"manageeTenantName,omitempty"`
	// ManagedByTenantID - Id of the managedBy tenant.
	ManagedByTenantID *string `json:"managedByTenantId,omitempty"`
	// ManagedByTenantName - Name of the managedBy tenant.
	ManagedByTenantName *string `json:"managedByTenantName,omitempty"`
}

// RegistrationAssignmentsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of
// a long-running operation.
type RegistrationAssignmentsCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(RegistrationAssignmentsClient) (RegistrationAssignment, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *RegistrationAssignmentsCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for RegistrationAssignmentsCreateOrUpdateFuture.Result.
func (future *RegistrationAssignmentsCreateOrUpdateFuture) result(client RegistrationAssignmentsClient) (ra RegistrationAssignment, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("managedservices.RegistrationAssignmentsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if ra.Response.Response, err = future.GetResult(sender); err == nil && ra.Response.Response.StatusCode != http.StatusNoContent {
		ra, err = client.CreateOrUpdateResponder(ra.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsCreateOrUpdateFuture", "Result", ra.Response.Response, "Failure responding to request")
		}
	}
	return
}

// RegistrationAssignmentsDeleteFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type RegistrationAssignmentsDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(RegistrationAssignmentsClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *RegistrationAssignmentsDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for RegistrationAssignmentsDeleteFuture.Result.
func (future *RegistrationAssignmentsDeleteFuture) result(client RegistrationAssignmentsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationAssignmentsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("managedservices.RegistrationAssignmentsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// RegistrationDefinition registration definition.
type RegistrationDefinition struct {
	autorest.Response `json:"-"`
	// Properties - Properties of a registration definition.
	Properties *RegistrationDefinitionProperties `json:"properties,omitempty"`
	// Plan - Plan details for the managed services.
	Plan *Plan `json:"plan,omitempty"`
	// ID - READ-ONLY; Fully qualified path of the registration definition.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Type of the resource.
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Name of the registration definition.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for RegistrationDefinition.
func (rd RegistrationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rd.Properties != nil {
		objectMap["properties"] = rd.Properties
	}
	if rd.Plan != nil {
		objectMap["plan"] = rd.Plan
	}
	return json.Marshal(objectMap)
}

// RegistrationDefinitionList list of registration definitions.
type RegistrationDefinitionList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of registration definitions.
	Value *[]RegistrationDefinition `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to next page of registration definitions.
	NextLink *string `json:"nextLink,omitempty"`
}

// RegistrationDefinitionListIterator provides access to a complete listing of RegistrationDefinition
// values.
type RegistrationDefinitionListIterator struct {
	i    int
	page RegistrationDefinitionListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *RegistrationDefinitionListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationDefinitionListIterator.NextWithContext")
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
func (iter *RegistrationDefinitionListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter RegistrationDefinitionListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter RegistrationDefinitionListIterator) Response() RegistrationDefinitionList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter RegistrationDefinitionListIterator) Value() RegistrationDefinition {
	if !iter.page.NotDone() {
		return RegistrationDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the RegistrationDefinitionListIterator type.
func NewRegistrationDefinitionListIterator(page RegistrationDefinitionListPage) RegistrationDefinitionListIterator {
	return RegistrationDefinitionListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rdl RegistrationDefinitionList) IsEmpty() bool {
	return rdl.Value == nil || len(*rdl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rdl RegistrationDefinitionList) hasNextLink() bool {
	return rdl.NextLink != nil && len(*rdl.NextLink) != 0
}

// registrationDefinitionListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rdl RegistrationDefinitionList) registrationDefinitionListPreparer(ctx context.Context) (*http.Request, error) {
	if !rdl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rdl.NextLink)))
}

// RegistrationDefinitionListPage contains a page of RegistrationDefinition values.
type RegistrationDefinitionListPage struct {
	fn  func(context.Context, RegistrationDefinitionList) (RegistrationDefinitionList, error)
	rdl RegistrationDefinitionList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *RegistrationDefinitionListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegistrationDefinitionListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rdl)
		if err != nil {
			return err
		}
		page.rdl = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *RegistrationDefinitionListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page RegistrationDefinitionListPage) NotDone() bool {
	return !page.rdl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page RegistrationDefinitionListPage) Response() RegistrationDefinitionList {
	return page.rdl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page RegistrationDefinitionListPage) Values() []RegistrationDefinition {
	if page.rdl.IsEmpty() {
		return nil
	}
	return *page.rdl.Value
}

// Creates a new instance of the RegistrationDefinitionListPage type.
func NewRegistrationDefinitionListPage(cur RegistrationDefinitionList, getNextPage func(context.Context, RegistrationDefinitionList) (RegistrationDefinitionList, error)) RegistrationDefinitionListPage {
	return RegistrationDefinitionListPage{
		fn:  getNextPage,
		rdl: cur,
	}
}

// RegistrationDefinitionProperties properties of a registration definition.
type RegistrationDefinitionProperties struct {
	// Description - Description of the registration definition.
	Description *string `json:"description,omitempty"`
	// Authorizations - Authorization tuple containing principal id of the user/security group or service principal and id of the build-in role.
	Authorizations *[]Authorization `json:"authorizations,omitempty"`
	// RegistrationDefinitionName - Name of the registration definition.
	RegistrationDefinitionName *string `json:"registrationDefinitionName,omitempty"`
	// ManagedByTenantID - Id of the managedBy tenant.
	ManagedByTenantID *string `json:"managedByTenantId,omitempty"`
	// ProvisioningState - READ-ONLY; Current state of the registration definition. Possible values include: 'NotSpecified', 'Accepted', 'Running', 'Ready', 'Creating', 'Created', 'Deleting', 'Deleted', 'Canceled', 'Failed', 'Succeeded', 'Updating'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// ManagedByTenantName - READ-ONLY; Name of the managedBy tenant.
	ManagedByTenantName *string `json:"managedByTenantName,omitempty"`
}

// MarshalJSON is the custom marshaler for RegistrationDefinitionProperties.
func (rdp RegistrationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rdp.Description != nil {
		objectMap["description"] = rdp.Description
	}
	if rdp.Authorizations != nil {
		objectMap["authorizations"] = rdp.Authorizations
	}
	if rdp.RegistrationDefinitionName != nil {
		objectMap["registrationDefinitionName"] = rdp.RegistrationDefinitionName
	}
	if rdp.ManagedByTenantID != nil {
		objectMap["managedByTenantId"] = rdp.ManagedByTenantID
	}
	return json.Marshal(objectMap)
}

// RegistrationDefinitionsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of
// a long-running operation.
type RegistrationDefinitionsCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(RegistrationDefinitionsClient) (RegistrationDefinition, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *RegistrationDefinitionsCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for RegistrationDefinitionsCreateOrUpdateFuture.Result.
func (future *RegistrationDefinitionsCreateOrUpdateFuture) result(client RegistrationDefinitionsClient) (rd RegistrationDefinition, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("managedservices.RegistrationDefinitionsCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if rd.Response.Response, err = future.GetResult(sender); err == nil && rd.Response.Response.StatusCode != http.StatusNoContent {
		rd, err = client.CreateOrUpdateResponder(rd.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedservices.RegistrationDefinitionsCreateOrUpdateFuture", "Result", rd.Response.Response, "Failure responding to request")
		}
	}
	return
}
