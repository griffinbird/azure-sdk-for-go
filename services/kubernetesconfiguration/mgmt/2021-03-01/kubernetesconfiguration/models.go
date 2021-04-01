package kubernetesconfiguration

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/kubernetesconfiguration/mgmt/2021-03-01/kubernetesconfiguration"

// AzureEntityResource the resource model definition for an Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// ComplianceStatus compliance Status details
type ComplianceStatus struct {
	// ComplianceState - READ-ONLY; The compliance state of the configuration. Possible values include: 'Pending', 'Compliant', 'Noncompliant', 'Installed', 'Failed'
	ComplianceState ComplianceStateType `json:"complianceState,omitempty"`
	// LastConfigApplied - Datetime the configuration was last applied.
	LastConfigApplied *date.Time `json:"lastConfigApplied,omitempty"`
	// Message - Message from when the configuration was applied.
	Message *string `json:"message,omitempty"`
	// MessageLevel - Level of the message. Possible values include: 'Error', 'Warning', 'Information'
	MessageLevel MessageLevelType `json:"messageLevel,omitempty"`
}

// MarshalJSON is the custom marshaler for ComplianceStatus.
func (cs ComplianceStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cs.LastConfigApplied != nil {
		objectMap["lastConfigApplied"] = cs.LastConfigApplied
	}
	if cs.Message != nil {
		objectMap["message"] = cs.Message
	}
	if cs.MessageLevel != "" {
		objectMap["messageLevel"] = cs.MessageLevel
	}
	return json.Marshal(objectMap)
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - Description of the error.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - Error definition.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// HelmOperatorProperties properties for Helm operator.
type HelmOperatorProperties struct {
	// ChartVersion - Version of the operator Helm chart.
	ChartVersion *string `json:"chartVersion,omitempty"`
	// ChartValues - Values override for the operator Helm chart.
	ChartValues *string `json:"chartValues,omitempty"`
}

// ProxyResource the resource model definition for a Azure Resource Manager proxy resource. It will not
// have tags and a location
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Resource common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// ResourceProviderOperation supported operation of this resource provider.
type ResourceProviderOperation struct {
	// Name - Operation name, in format of {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`
	// IsDataAction - READ-ONLY; The flag that indicates whether the operation applies to data plane.
	IsDataAction *bool `json:"isDataAction,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceProviderOperation.
func (rpo ResourceProviderOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rpo.Name != nil {
		objectMap["name"] = rpo.Name
	}
	if rpo.Display != nil {
		objectMap["display"] = rpo.Display
	}
	return json.Marshal(objectMap)
}

// ResourceProviderOperationDisplay display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Provider - Resource provider: Microsoft KubernetesConfiguration.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of this operation.
	Description *string `json:"description,omitempty"`
}

// ResourceProviderOperationList result of the request to list operations.
type ResourceProviderOperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by this resource provider.
	Value *[]ResourceProviderOperation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceProviderOperationList.
func (rpol ResourceProviderOperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rpol.Value != nil {
		objectMap["value"] = rpol.Value
	}
	return json.Marshal(objectMap)
}

// ResourceProviderOperationListIterator provides access to a complete listing of ResourceProviderOperation
// values.
type ResourceProviderOperationListIterator struct {
	i    int
	page ResourceProviderOperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceProviderOperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListIterator.NextWithContext")
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
func (iter *ResourceProviderOperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceProviderOperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceProviderOperationListIterator) Response() ResourceProviderOperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceProviderOperationListIterator) Value() ResourceProviderOperation {
	if !iter.page.NotDone() {
		return ResourceProviderOperation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceProviderOperationListIterator type.
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return ResourceProviderOperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rpol ResourceProviderOperationList) IsEmpty() bool {
	return rpol.Value == nil || len(*rpol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rpol ResourceProviderOperationList) hasNextLink() bool {
	return rpol.NextLink != nil && len(*rpol.NextLink) != 0
}

// resourceProviderOperationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rpol ResourceProviderOperationList) resourceProviderOperationListPreparer(ctx context.Context) (*http.Request, error) {
	if !rpol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rpol.NextLink)))
}

// ResourceProviderOperationListPage contains a page of ResourceProviderOperation values.
type ResourceProviderOperationListPage struct {
	fn   func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)
	rpol ResourceProviderOperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceProviderOperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rpol)
		if err != nil {
			return err
		}
		page.rpol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceProviderOperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceProviderOperationListPage) NotDone() bool {
	return !page.rpol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceProviderOperationListPage) Response() ResourceProviderOperationList {
	return page.rpol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceProviderOperationListPage) Values() []ResourceProviderOperation {
	if page.rpol.IsEmpty() {
		return nil
	}
	return *page.rpol.Value
}

// Creates a new instance of the ResourceProviderOperationListPage type.
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return ResourceProviderOperationListPage{
		fn:   getNextPage,
		rpol: cur,
	}
}

// Result sample result definition
type Result struct {
	// SampleProperty - Sample property of type string
	SampleProperty *string `json:"sampleProperty,omitempty"`
}

// SourceControlConfiguration the SourceControl Configuration object returned in Get & Put response.
type SourceControlConfiguration struct {
	autorest.Response `json:"-"`
	// SourceControlConfigurationProperties - Properties to create a Source Control Configuration resource
	*SourceControlConfigurationProperties `json:"properties,omitempty"`
	// SystemData - Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData *SystemData `json:"systemData,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for SourceControlConfiguration.
func (scc SourceControlConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if scc.SourceControlConfigurationProperties != nil {
		objectMap["properties"] = scc.SourceControlConfigurationProperties
	}
	if scc.SystemData != nil {
		objectMap["systemData"] = scc.SystemData
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SourceControlConfiguration struct.
func (scc *SourceControlConfiguration) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var sourceControlConfigurationProperties SourceControlConfigurationProperties
				err = json.Unmarshal(*v, &sourceControlConfigurationProperties)
				if err != nil {
					return err
				}
				scc.SourceControlConfigurationProperties = &sourceControlConfigurationProperties
			}
		case "systemData":
			if v != nil {
				var systemData SystemData
				err = json.Unmarshal(*v, &systemData)
				if err != nil {
					return err
				}
				scc.SystemData = &systemData
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				scc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				scc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				scc.Type = &typeVar
			}
		}
	}

	return nil
}

// SourceControlConfigurationList result of the request to list Source Control Configurations.  It contains
// a list of SourceControlConfiguration objects and a URL link to get the next set of results.
type SourceControlConfigurationList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of Source Control Configurations within a Kubernetes cluster.
	Value *[]SourceControlConfiguration `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of configuration objects, if any.
	NextLink *string `json:"nextLink,omitempty"`
}

// SourceControlConfigurationListIterator provides access to a complete listing of
// SourceControlConfiguration values.
type SourceControlConfigurationListIterator struct {
	i    int
	page SourceControlConfigurationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SourceControlConfigurationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationListIterator.NextWithContext")
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
func (iter *SourceControlConfigurationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SourceControlConfigurationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SourceControlConfigurationListIterator) Response() SourceControlConfigurationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SourceControlConfigurationListIterator) Value() SourceControlConfiguration {
	if !iter.page.NotDone() {
		return SourceControlConfiguration{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SourceControlConfigurationListIterator type.
func NewSourceControlConfigurationListIterator(page SourceControlConfigurationListPage) SourceControlConfigurationListIterator {
	return SourceControlConfigurationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (sccl SourceControlConfigurationList) IsEmpty() bool {
	return sccl.Value == nil || len(*sccl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (sccl SourceControlConfigurationList) hasNextLink() bool {
	return sccl.NextLink != nil && len(*sccl.NextLink) != 0
}

// sourceControlConfigurationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (sccl SourceControlConfigurationList) sourceControlConfigurationListPreparer(ctx context.Context) (*http.Request, error) {
	if !sccl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(sccl.NextLink)))
}

// SourceControlConfigurationListPage contains a page of SourceControlConfiguration values.
type SourceControlConfigurationListPage struct {
	fn   func(context.Context, SourceControlConfigurationList) (SourceControlConfigurationList, error)
	sccl SourceControlConfigurationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SourceControlConfigurationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SourceControlConfigurationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.sccl)
		if err != nil {
			return err
		}
		page.sccl = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SourceControlConfigurationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SourceControlConfigurationListPage) NotDone() bool {
	return !page.sccl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SourceControlConfigurationListPage) Response() SourceControlConfigurationList {
	return page.sccl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SourceControlConfigurationListPage) Values() []SourceControlConfiguration {
	if page.sccl.IsEmpty() {
		return nil
	}
	return *page.sccl.Value
}

// Creates a new instance of the SourceControlConfigurationListPage type.
func NewSourceControlConfigurationListPage(cur SourceControlConfigurationList, getNextPage func(context.Context, SourceControlConfigurationList) (SourceControlConfigurationList, error)) SourceControlConfigurationListPage {
	return SourceControlConfigurationListPage{
		fn:   getNextPage,
		sccl: cur,
	}
}

// SourceControlConfigurationProperties properties to create a Source Control Configuration resource
type SourceControlConfigurationProperties struct {
	// RepositoryURL - Url of the SourceControl Repository.
	RepositoryURL *string `json:"repositoryUrl,omitempty"`
	// OperatorNamespace - The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
	OperatorNamespace *string `json:"operatorNamespace,omitempty"`
	// OperatorInstanceName - Instance name of the operator - identifying the specific configuration.
	OperatorInstanceName *string `json:"operatorInstanceName,omitempty"`
	// OperatorType - Type of the operator. Possible values include: 'Flux'
	OperatorType OperatorType `json:"operatorType,omitempty"`
	// OperatorParams - Any Parameters for the Operator instance in string format.
	OperatorParams *string `json:"operatorParams,omitempty"`
	// ConfigurationProtectedSettings - Name-value pairs of protected configuration settings for the configuration
	ConfigurationProtectedSettings map[string]*string `json:"configurationProtectedSettings"`
	// OperatorScope - Scope at which the operator will be installed. Possible values include: 'Cluster', 'Namespace'
	OperatorScope OperatorScopeType `json:"operatorScope,omitempty"`
	// RepositoryPublicKey - READ-ONLY; Public Key associated with this SourceControl configuration (either generated within the cluster or provided by the user).
	RepositoryPublicKey *string `json:"repositoryPublicKey,omitempty"`
	// SSHKnownHostsContents - Base64-encoded known_hosts contents containing public SSH keys required to access private Git instances
	SSHKnownHostsContents *string `json:"sshKnownHostsContents,omitempty"`
	// EnableHelmOperator - Option to enable Helm Operator for this git configuration.
	EnableHelmOperator *bool `json:"enableHelmOperator,omitempty"`
	// HelmOperatorProperties - Properties for Helm operator.
	HelmOperatorProperties *HelmOperatorProperties `json:"helmOperatorProperties,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state of the resource provider. Possible values include: 'ProvisioningStateTypeAccepted', 'ProvisioningStateTypeDeleting', 'ProvisioningStateTypeRunning', 'ProvisioningStateTypeSucceeded', 'ProvisioningStateTypeFailed'
	ProvisioningState ProvisioningStateType `json:"provisioningState,omitempty"`
	// ComplianceStatus - READ-ONLY; Compliance Status of the Configuration
	ComplianceStatus *ComplianceStatus `json:"complianceStatus,omitempty"`
}

// MarshalJSON is the custom marshaler for SourceControlConfigurationProperties.
func (scc SourceControlConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if scc.RepositoryURL != nil {
		objectMap["repositoryUrl"] = scc.RepositoryURL
	}
	if scc.OperatorNamespace != nil {
		objectMap["operatorNamespace"] = scc.OperatorNamespace
	}
	if scc.OperatorInstanceName != nil {
		objectMap["operatorInstanceName"] = scc.OperatorInstanceName
	}
	if scc.OperatorType != "" {
		objectMap["operatorType"] = scc.OperatorType
	}
	if scc.OperatorParams != nil {
		objectMap["operatorParams"] = scc.OperatorParams
	}
	if scc.ConfigurationProtectedSettings != nil {
		objectMap["configurationProtectedSettings"] = scc.ConfigurationProtectedSettings
	}
	if scc.OperatorScope != "" {
		objectMap["operatorScope"] = scc.OperatorScope
	}
	if scc.SSHKnownHostsContents != nil {
		objectMap["sshKnownHostsContents"] = scc.SSHKnownHostsContents
	}
	if scc.EnableHelmOperator != nil {
		objectMap["enableHelmOperator"] = scc.EnableHelmOperator
	}
	if scc.HelmOperatorProperties != nil {
		objectMap["helmOperatorProperties"] = scc.HelmOperatorProperties
	}
	return json.Marshal(objectMap)
}

// SourceControlConfigurationsDeleteFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type SourceControlConfigurationsDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(SourceControlConfigurationsClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *SourceControlConfigurationsDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for SourceControlConfigurationsDeleteFuture.Result.
func (future *SourceControlConfigurationsDeleteFuture) result(client SourceControlConfigurationsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetesconfiguration.SourceControlConfigurationsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("kubernetesconfiguration.SourceControlConfigurationsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The timestamp of resource last modification (UTC)
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for an Azure Resource Manager tracked top level resource
// which has 'tags' and a 'location'
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
