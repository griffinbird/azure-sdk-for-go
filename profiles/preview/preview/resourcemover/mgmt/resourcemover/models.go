//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package resourcemover

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/resourcemover/mgmt/2019-10-01-preview/resourcemover"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DependencyType = original.DependencyType

const (
	RequiredForMove    DependencyType = original.RequiredForMove
	RequiredForPrepare DependencyType = original.RequiredForPrepare
)

type JobName = original.JobName

const (
	InitialSync JobName = original.InitialSync
)

type MoveResourceInputType = original.MoveResourceInputType

const (
	MoveResourceID       MoveResourceInputType = original.MoveResourceID
	MoveResourceSourceID MoveResourceInputType = original.MoveResourceSourceID
)

type MoveState = original.MoveState

const (
	AssignmentPending MoveState = original.AssignmentPending
	CommitFailed      MoveState = original.CommitFailed
	CommitInProgress  MoveState = original.CommitInProgress
	CommitPending     MoveState = original.CommitPending
	Committed         MoveState = original.Committed
	DiscardFailed     MoveState = original.DiscardFailed
	DiscardInProgress MoveState = original.DiscardInProgress
	MoveFailed        MoveState = original.MoveFailed
	MoveInProgress    MoveState = original.MoveInProgress
	MovePending       MoveState = original.MovePending
	PrepareFailed     MoveState = original.PrepareFailed
	PrepareInProgress MoveState = original.PrepareInProgress
	PreparePending    MoveState = original.PreparePending
)

type ProvisioningState = original.ProvisioningState

const (
	Creating  ProvisioningState = original.Creating
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type ResolutionType = original.ResolutionType

const (
	Automatic ResolutionType = original.Automatic
	Manual    ResolutionType = original.Manual
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None           ResourceIdentityType = original.None
	SystemAssigned ResourceIdentityType = original.SystemAssigned
	UserAssigned   ResourceIdentityType = original.UserAssigned
)

type ResourceType = original.ResourceType

const (
	ResourceTypeMicrosoftComputeavailabilitySets      ResourceType = original.ResourceTypeMicrosoftComputeavailabilitySets
	ResourceTypeMicrosoftComputevirtualMachines       ResourceType = original.ResourceTypeMicrosoftComputevirtualMachines
	ResourceTypeMicrosoftNetworkloadBalancers         ResourceType = original.ResourceTypeMicrosoftNetworkloadBalancers
	ResourceTypeMicrosoftNetworknetworkInterfaces     ResourceType = original.ResourceTypeMicrosoftNetworknetworkInterfaces
	ResourceTypeMicrosoftNetworknetworkSecurityGroups ResourceType = original.ResourceTypeMicrosoftNetworknetworkSecurityGroups
	ResourceTypeMicrosoftNetworkpublicIPAddresses     ResourceType = original.ResourceTypeMicrosoftNetworkpublicIPAddresses
	ResourceTypeMicrosoftNetworkvirtualNetworks       ResourceType = original.ResourceTypeMicrosoftNetworkvirtualNetworks
	ResourceTypeMicrosoftSqlservers                   ResourceType = original.ResourceTypeMicrosoftSqlservers
	ResourceTypeMicrosoftSqlserversdatabases          ResourceType = original.ResourceTypeMicrosoftSqlserversdatabases
	ResourceTypeMicrosoftSqlserverselasticPools       ResourceType = original.ResourceTypeMicrosoftSqlserverselasticPools
	ResourceTypeResourceGroups                        ResourceType = original.ResourceTypeResourceGroups
	ResourceTypeResourceSettings                      ResourceType = original.ResourceTypeResourceSettings
)

type TargetAvailabilityZone = original.TargetAvailabilityZone

const (
	NA    TargetAvailabilityZone = original.NA
	One   TargetAvailabilityZone = original.One
	Three TargetAvailabilityZone = original.Three
	Two   TargetAvailabilityZone = original.Two
)

type ZoneRedundant = original.ZoneRedundant

const (
	Disable ZoneRedundant = original.Disable
	Enable  ZoneRedundant = original.Enable
)

type AffectedMoveResource = original.AffectedMoveResource
type AutomaticResolutionProperties = original.AutomaticResolutionProperties
type AvailabilitySetResourceSettings = original.AvailabilitySetResourceSettings
type AzureResourceReference = original.AzureResourceReference
type BaseClient = original.BaseClient
type BasicResourceSettings = original.BasicResourceSettings
type BulkRemoveRequest = original.BulkRemoveRequest
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CommitRequest = original.CommitRequest
type DiscardRequest = original.DiscardRequest
type Display = original.Display
type Identity = original.Identity
type JobStatus = original.JobStatus
type LBBackendAddressPoolResourceSettings = original.LBBackendAddressPoolResourceSettings
type LBFrontendIPConfigurationResourceSettings = original.LBFrontendIPConfigurationResourceSettings
type LoadBalancerBackendAddressPoolReference = original.LoadBalancerBackendAddressPoolReference
type LoadBalancerNatRuleReference = original.LoadBalancerNatRuleReference
type LoadBalancerResourceSettings = original.LoadBalancerResourceSettings
type ManualResolutionProperties = original.ManualResolutionProperties
type MoveCollection = original.MoveCollection
type MoveCollectionProperties = original.MoveCollectionProperties
type MoveCollectionResultList = original.MoveCollectionResultList
type MoveCollectionResultListIterator = original.MoveCollectionResultListIterator
type MoveCollectionResultListPage = original.MoveCollectionResultListPage
type MoveCollectionsBulkRemoveFuture = original.MoveCollectionsBulkRemoveFuture
type MoveCollectionsClient = original.MoveCollectionsClient
type MoveCollectionsCommitFuture = original.MoveCollectionsCommitFuture
type MoveCollectionsDeleteFuture = original.MoveCollectionsDeleteFuture
type MoveCollectionsDiscardFuture = original.MoveCollectionsDiscardFuture
type MoveCollectionsInitiateMoveFuture = original.MoveCollectionsInitiateMoveFuture
type MoveCollectionsPrepareFuture = original.MoveCollectionsPrepareFuture
type MoveCollectionsResolveDependenciesFuture = original.MoveCollectionsResolveDependenciesFuture
type MoveErrorInfo = original.MoveErrorInfo
type MoveResource = original.MoveResource
type MoveResourceCollection = original.MoveResourceCollection
type MoveResourceCollectionIterator = original.MoveResourceCollectionIterator
type MoveResourceCollectionPage = original.MoveResourceCollectionPage
type MoveResourceDependency = original.MoveResourceDependency
type MoveResourceDependencyOverride = original.MoveResourceDependencyOverride
type MoveResourceError = original.MoveResourceError
type MoveResourceErrorBody = original.MoveResourceErrorBody
type MoveResourceFilter = original.MoveResourceFilter
type MoveResourceFilterProperties = original.MoveResourceFilterProperties
type MoveResourceProperties = original.MoveResourceProperties
type MoveResourcePropertiesErrors = original.MoveResourcePropertiesErrors
type MoveResourcePropertiesMoveStatus = original.MoveResourcePropertiesMoveStatus
type MoveResourceStatus = original.MoveResourceStatus
type MoveResourcesClient = original.MoveResourcesClient
type MoveResourcesCreateFuture = original.MoveResourcesCreateFuture
type MoveResourcesDeleteFuture = original.MoveResourcesDeleteFuture
type NetworkInterfaceResourceSettings = original.NetworkInterfaceResourceSettings
type NetworkSecurityGroupResourceSettings = original.NetworkSecurityGroupResourceSettings
type NicIPConfigurationResourceSettings = original.NicIPConfigurationResourceSettings
type NsgSecurityRule = original.NsgSecurityRule
type OperationErrorAdditionalInfo = original.OperationErrorAdditionalInfo
type OperationStatus = original.OperationStatus
type OperationStatusError = original.OperationStatusError
type OperationsDiscovery = original.OperationsDiscovery
type OperationsDiscoveryClient = original.OperationsDiscoveryClient
type OperationsDiscoveryCollection = original.OperationsDiscoveryCollection
type PrepareRequest = original.PrepareRequest
type ProxyResourceReference = original.ProxyResourceReference
type PublicIPAddressResourceSettings = original.PublicIPAddressResourceSettings
type ResourceGroupResourceSettings = original.ResourceGroupResourceSettings
type ResourceMoveRequestType = original.ResourceMoveRequestType
type ResourceSettings = original.ResourceSettings
type SQLDatabaseResourceSettings = original.SQLDatabaseResourceSettings
type SQLElasticPoolResourceSettings = original.SQLElasticPoolResourceSettings
type SQLServerResourceSettings = original.SQLServerResourceSettings
type SubnetReference = original.SubnetReference
type SubnetResourceSettings = original.SubnetResourceSettings
type SummaryItem = original.SummaryItem
type UnresolvedDependenciesClient = original.UnresolvedDependenciesClient
type UnresolvedDependency = original.UnresolvedDependency
type UnresolvedDependencyCollection = original.UnresolvedDependencyCollection
type UpdateMoveCollectionRequest = original.UpdateMoveCollectionRequest
type VirtualMachineResourceSettings = original.VirtualMachineResourceSettings
type VirtualNetworkResourceSettings = original.VirtualNetworkResourceSettings

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewMoveCollectionResultListIterator(page MoveCollectionResultListPage) MoveCollectionResultListIterator {
	return original.NewMoveCollectionResultListIterator(page)
}
func NewMoveCollectionResultListPage(cur MoveCollectionResultList, getNextPage func(context.Context, MoveCollectionResultList) (MoveCollectionResultList, error)) MoveCollectionResultListPage {
	return original.NewMoveCollectionResultListPage(cur, getNextPage)
}
func NewMoveCollectionsClient(subscriptionID string) MoveCollectionsClient {
	return original.NewMoveCollectionsClient(subscriptionID)
}
func NewMoveCollectionsClientWithBaseURI(baseURI string, subscriptionID string) MoveCollectionsClient {
	return original.NewMoveCollectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMoveResourceCollectionIterator(page MoveResourceCollectionPage) MoveResourceCollectionIterator {
	return original.NewMoveResourceCollectionIterator(page)
}
func NewMoveResourceCollectionPage(cur MoveResourceCollection, getNextPage func(context.Context, MoveResourceCollection) (MoveResourceCollection, error)) MoveResourceCollectionPage {
	return original.NewMoveResourceCollectionPage(cur, getNextPage)
}
func NewMoveResourcesClient(subscriptionID string) MoveResourcesClient {
	return original.NewMoveResourcesClient(subscriptionID)
}
func NewMoveResourcesClientWithBaseURI(baseURI string, subscriptionID string) MoveResourcesClient {
	return original.NewMoveResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsDiscoveryClient(subscriptionID string) OperationsDiscoveryClient {
	return original.NewOperationsDiscoveryClient(subscriptionID)
}
func NewOperationsDiscoveryClientWithBaseURI(baseURI string, subscriptionID string) OperationsDiscoveryClient {
	return original.NewOperationsDiscoveryClientWithBaseURI(baseURI, subscriptionID)
}
func NewUnresolvedDependenciesClient(subscriptionID string) UnresolvedDependenciesClient {
	return original.NewUnresolvedDependenciesClient(subscriptionID)
}
func NewUnresolvedDependenciesClientWithBaseURI(baseURI string, subscriptionID string) UnresolvedDependenciesClient {
	return original.NewUnresolvedDependenciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDependencyTypeValues() []DependencyType {
	return original.PossibleDependencyTypeValues()
}
func PossibleJobNameValues() []JobName {
	return original.PossibleJobNameValues()
}
func PossibleMoveResourceInputTypeValues() []MoveResourceInputType {
	return original.PossibleMoveResourceInputTypeValues()
}
func PossibleMoveStateValues() []MoveState {
	return original.PossibleMoveStateValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResolutionTypeValues() []ResolutionType {
	return original.PossibleResolutionTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleTargetAvailabilityZoneValues() []TargetAvailabilityZone {
	return original.PossibleTargetAvailabilityZoneValues()
}
func PossibleZoneRedundantValues() []ZoneRedundant {
	return original.PossibleZoneRedundantValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
