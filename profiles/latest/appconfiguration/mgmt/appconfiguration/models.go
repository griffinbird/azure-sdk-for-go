// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package appconfiguration

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/appconfiguration/mgmt/2020-06-01/appconfiguration"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionsRequired = original.ActionsRequired

const (
	None     ActionsRequired = original.None
	Recreate ActionsRequired = original.Recreate
)

type ConnectionStatus = original.ConnectionStatus

const (
	Approved     ConnectionStatus = original.Approved
	Disconnected ConnectionStatus = original.Disconnected
	Pending      ConnectionStatus = original.Pending
	Rejected     ConnectionStatus = original.Rejected
)

type IdentityType = original.IdentityType

const (
	IdentityTypeNone                       IdentityType = original.IdentityTypeNone
	IdentityTypeSystemAssigned             IdentityType = original.IdentityTypeSystemAssigned
	IdentityTypeSystemAssignedUserAssigned IdentityType = original.IdentityTypeSystemAssignedUserAssigned
	IdentityTypeUserAssigned               IdentityType = original.IdentityTypeUserAssigned
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	Disabled PublicNetworkAccess = original.Disabled
	Enabled  PublicNetworkAccess = original.Enabled
)

type APIKey = original.APIKey
type APIKeyListResult = original.APIKeyListResult
type APIKeyListResultIterator = original.APIKeyListResultIterator
type APIKeyListResultPage = original.APIKeyListResultPage
type BaseClient = original.BaseClient
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type ConfigurationStore = original.ConfigurationStore
type ConfigurationStoreListResult = original.ConfigurationStoreListResult
type ConfigurationStoreListResultIterator = original.ConfigurationStoreListResultIterator
type ConfigurationStoreListResultPage = original.ConfigurationStoreListResultPage
type ConfigurationStoreProperties = original.ConfigurationStoreProperties
type ConfigurationStorePropertiesUpdateParameters = original.ConfigurationStorePropertiesUpdateParameters
type ConfigurationStoreUpdateParameters = original.ConfigurationStoreUpdateParameters
type ConfigurationStoresClient = original.ConfigurationStoresClient
type ConfigurationStoresCreateFuture = original.ConfigurationStoresCreateFuture
type ConfigurationStoresDeleteFuture = original.ConfigurationStoresDeleteFuture
type ConfigurationStoresUpdateFuture = original.ConfigurationStoresUpdateFuture
type EncryptionProperties = original.EncryptionProperties
type Error = original.Error
type KeyValue = original.KeyValue
type KeyVaultProperties = original.KeyVaultProperties
type ListKeyValueParameters = original.ListKeyValueParameters
type NameAvailabilityStatus = original.NameAvailabilityStatus
type OperationDefinition = original.OperationDefinition
type OperationDefinitionDisplay = original.OperationDefinitionDisplay
type OperationDefinitionListResult = original.OperationDefinitionListResult
type OperationDefinitionListResultIterator = original.OperationDefinitionListResultIterator
type OperationDefinitionListResultPage = original.OperationDefinitionListResultPage
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionReference = original.PrivateEndpointConnectionReference
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceListResultIterator = original.PrivateLinkResourceListResultIterator
type PrivateLinkResourceListResultPage = original.PrivateLinkResourceListResultPage
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type ResourceIdentity = original.ResourceIdentity
type Sku = original.Sku
type UserIdentity = original.UserIdentity

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAPIKeyListResultIterator(page APIKeyListResultPage) APIKeyListResultIterator {
	return original.NewAPIKeyListResultIterator(page)
}
func NewAPIKeyListResultPage(cur APIKeyListResult, getNextPage func(context.Context, APIKeyListResult) (APIKeyListResult, error)) APIKeyListResultPage {
	return original.NewAPIKeyListResultPage(cur, getNextPage)
}
func NewConfigurationStoreListResultIterator(page ConfigurationStoreListResultPage) ConfigurationStoreListResultIterator {
	return original.NewConfigurationStoreListResultIterator(page)
}
func NewConfigurationStoreListResultPage(cur ConfigurationStoreListResult, getNextPage func(context.Context, ConfigurationStoreListResult) (ConfigurationStoreListResult, error)) ConfigurationStoreListResultPage {
	return original.NewConfigurationStoreListResultPage(cur, getNextPage)
}
func NewConfigurationStoresClient(subscriptionID string) ConfigurationStoresClient {
	return original.NewConfigurationStoresClient(subscriptionID)
}
func NewConfigurationStoresClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationStoresClient {
	return original.NewConfigurationStoresClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationDefinitionListResultIterator(page OperationDefinitionListResultPage) OperationDefinitionListResultIterator {
	return original.NewOperationDefinitionListResultIterator(page)
}
func NewOperationDefinitionListResultPage(cur OperationDefinitionListResult, getNextPage func(context.Context, OperationDefinitionListResult) (OperationDefinitionListResult, error)) OperationDefinitionListResultPage {
	return original.NewOperationDefinitionListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(cur PrivateEndpointConnectionListResult, getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourceListResultIterator(page PrivateLinkResourceListResultPage) PrivateLinkResourceListResultIterator {
	return original.NewPrivateLinkResourceListResultIterator(page)
}
func NewPrivateLinkResourceListResultPage(cur PrivateLinkResourceListResult, getNextPage func(context.Context, PrivateLinkResourceListResult) (PrivateLinkResourceListResult, error)) PrivateLinkResourceListResultPage {
	return original.NewPrivateLinkResourceListResultPage(cur, getNextPage)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionsRequiredValues() []ActionsRequired {
	return original.PossibleActionsRequiredValues()
}
func PossibleConnectionStatusValues() []ConnectionStatus {
	return original.PossibleConnectionStatusValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
