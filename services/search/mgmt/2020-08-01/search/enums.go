package search

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AdminKeyKind enumerates the values for admin key kind.
type AdminKeyKind string

const (
	// Primary ...
	Primary AdminKeyKind = "primary"
	// Secondary ...
	Secondary AdminKeyKind = "secondary"
)

// PossibleAdminKeyKindValues returns an array of possible values for the AdminKeyKind const type.
func PossibleAdminKeyKindValues() []AdminKeyKind {
	return []AdminKeyKind{Primary, Secondary}
}

// HostingMode enumerates the values for hosting mode.
type HostingMode string

const (
	// Default ...
	Default HostingMode = "default"
	// HighDensity ...
	HighDensity HostingMode = "highDensity"
)

// PossibleHostingModeValues returns an array of possible values for the HostingMode const type.
func PossibleHostingModeValues() []HostingMode {
	return []HostingMode{Default, HighDensity}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// None ...
	None IdentityType = "None"
	// SystemAssigned ...
	SystemAssigned IdentityType = "SystemAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{None, SystemAssigned}
}

// PrivateLinkServiceConnectionStatus enumerates the values for private link service connection status.
type PrivateLinkServiceConnectionStatus string

const (
	// Approved ...
	Approved PrivateLinkServiceConnectionStatus = "Approved"
	// Disconnected ...
	Disconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	// Pending ...
	Pending PrivateLinkServiceConnectionStatus = "Pending"
	// Rejected ...
	Rejected PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns an array of possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{Approved, Disconnected, Pending, Rejected}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Failed ...
	Failed ProvisioningState = "failed"
	// Provisioning ...
	Provisioning ProvisioningState = "provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Failed, Provisioning, Succeeded}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// Disabled ...
	Disabled PublicNetworkAccess = "disabled"
	// Enabled ...
	Enabled PublicNetworkAccess = "enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{Disabled, Enabled}
}

// ServiceStatus enumerates the values for service status.
type ServiceStatus string

const (
	// ServiceStatusDegraded ...
	ServiceStatusDegraded ServiceStatus = "degraded"
	// ServiceStatusDeleting ...
	ServiceStatusDeleting ServiceStatus = "deleting"
	// ServiceStatusDisabled ...
	ServiceStatusDisabled ServiceStatus = "disabled"
	// ServiceStatusError ...
	ServiceStatusError ServiceStatus = "error"
	// ServiceStatusProvisioning ...
	ServiceStatusProvisioning ServiceStatus = "provisioning"
	// ServiceStatusRunning ...
	ServiceStatusRunning ServiceStatus = "running"
)

// PossibleServiceStatusValues returns an array of possible values for the ServiceStatus const type.
func PossibleServiceStatusValues() []ServiceStatus {
	return []ServiceStatus{ServiceStatusDegraded, ServiceStatusDeleting, ServiceStatusDisabled, ServiceStatusError, ServiceStatusProvisioning, ServiceStatusRunning}
}

// SharedPrivateLinkResourceAsyncOperationResult enumerates the values for shared private link resource async
// operation result.
type SharedPrivateLinkResourceAsyncOperationResult string

const (
	// SharedPrivateLinkResourceAsyncOperationResultFailed ...
	SharedPrivateLinkResourceAsyncOperationResultFailed SharedPrivateLinkResourceAsyncOperationResult = "Failed"
	// SharedPrivateLinkResourceAsyncOperationResultRunning ...
	SharedPrivateLinkResourceAsyncOperationResultRunning SharedPrivateLinkResourceAsyncOperationResult = "Running"
	// SharedPrivateLinkResourceAsyncOperationResultSucceeded ...
	SharedPrivateLinkResourceAsyncOperationResultSucceeded SharedPrivateLinkResourceAsyncOperationResult = "Succeeded"
)

// PossibleSharedPrivateLinkResourceAsyncOperationResultValues returns an array of possible values for the SharedPrivateLinkResourceAsyncOperationResult const type.
func PossibleSharedPrivateLinkResourceAsyncOperationResultValues() []SharedPrivateLinkResourceAsyncOperationResult {
	return []SharedPrivateLinkResourceAsyncOperationResult{SharedPrivateLinkResourceAsyncOperationResultFailed, SharedPrivateLinkResourceAsyncOperationResultRunning, SharedPrivateLinkResourceAsyncOperationResultSucceeded}
}

// SharedPrivateLinkResourceProvisioningState enumerates the values for shared private link resource
// provisioning state.
type SharedPrivateLinkResourceProvisioningState string

const (
	// SharedPrivateLinkResourceProvisioningStateDeleting ...
	SharedPrivateLinkResourceProvisioningStateDeleting SharedPrivateLinkResourceProvisioningState = "Deleting"
	// SharedPrivateLinkResourceProvisioningStateFailed ...
	SharedPrivateLinkResourceProvisioningStateFailed SharedPrivateLinkResourceProvisioningState = "Failed"
	// SharedPrivateLinkResourceProvisioningStateIncomplete ...
	SharedPrivateLinkResourceProvisioningStateIncomplete SharedPrivateLinkResourceProvisioningState = "Incomplete"
	// SharedPrivateLinkResourceProvisioningStateSucceeded ...
	SharedPrivateLinkResourceProvisioningStateSucceeded SharedPrivateLinkResourceProvisioningState = "Succeeded"
	// SharedPrivateLinkResourceProvisioningStateUpdating ...
	SharedPrivateLinkResourceProvisioningStateUpdating SharedPrivateLinkResourceProvisioningState = "Updating"
)

// PossibleSharedPrivateLinkResourceProvisioningStateValues returns an array of possible values for the SharedPrivateLinkResourceProvisioningState const type.
func PossibleSharedPrivateLinkResourceProvisioningStateValues() []SharedPrivateLinkResourceProvisioningState {
	return []SharedPrivateLinkResourceProvisioningState{SharedPrivateLinkResourceProvisioningStateDeleting, SharedPrivateLinkResourceProvisioningStateFailed, SharedPrivateLinkResourceProvisioningStateIncomplete, SharedPrivateLinkResourceProvisioningStateSucceeded, SharedPrivateLinkResourceProvisioningStateUpdating}
}

// SharedPrivateLinkResourceStatus enumerates the values for shared private link resource status.
type SharedPrivateLinkResourceStatus string

const (
	// SharedPrivateLinkResourceStatusApproved ...
	SharedPrivateLinkResourceStatusApproved SharedPrivateLinkResourceStatus = "Approved"
	// SharedPrivateLinkResourceStatusDisconnected ...
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	// SharedPrivateLinkResourceStatusPending ...
	SharedPrivateLinkResourceStatusPending SharedPrivateLinkResourceStatus = "Pending"
	// SharedPrivateLinkResourceStatusRejected ...
	SharedPrivateLinkResourceStatusRejected SharedPrivateLinkResourceStatus = "Rejected"
)

// PossibleSharedPrivateLinkResourceStatusValues returns an array of possible values for the SharedPrivateLinkResourceStatus const type.
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return []SharedPrivateLinkResourceStatus{SharedPrivateLinkResourceStatusApproved, SharedPrivateLinkResourceStatusDisconnected, SharedPrivateLinkResourceStatusPending, SharedPrivateLinkResourceStatusRejected}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "basic"
	// Free ...
	Free SkuName = "free"
	// Standard ...
	Standard SkuName = "standard"
	// Standard2 ...
	Standard2 SkuName = "standard2"
	// Standard3 ...
	Standard3 SkuName = "standard3"
	// StorageOptimizedL1 ...
	StorageOptimizedL1 SkuName = "storage_optimized_l1"
	// StorageOptimizedL2 ...
	StorageOptimizedL2 SkuName = "storage_optimized_l2"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Free, Standard, Standard2, Standard3, StorageOptimizedL1, StorageOptimizedL2}
}

// UnavailableNameReason enumerates the values for unavailable name reason.
type UnavailableNameReason string

const (
	// AlreadyExists ...
	AlreadyExists UnavailableNameReason = "AlreadyExists"
	// Invalid ...
	Invalid UnavailableNameReason = "Invalid"
)

// PossibleUnavailableNameReasonValues returns an array of possible values for the UnavailableNameReason const type.
func PossibleUnavailableNameReasonValues() []UnavailableNameReason {
	return []UnavailableNameReason{AlreadyExists, Invalid}
}
