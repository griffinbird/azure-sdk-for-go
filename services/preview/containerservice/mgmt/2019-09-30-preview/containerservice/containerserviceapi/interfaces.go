package containerserviceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/containerservice/mgmt/2019-09-30-preview/containerservice"
)

// OpenShiftManagedClustersClientAPI contains the set of methods on the OpenShiftManagedClustersClient type.
type OpenShiftManagedClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.OpenShiftManagedCluster) (result containerservice.OpenShiftManagedClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.OpenShiftManagedClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.OpenShiftManagedCluster, err error)
	List(ctx context.Context) (result containerservice.OpenShiftManagedClusterListResultPage, err error)
	ListComplete(ctx context.Context) (result containerservice.OpenShiftManagedClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.OpenShiftManagedClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerservice.OpenShiftManagedClusterListResultIterator, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.TagsObject) (result containerservice.OpenShiftManagedClustersUpdateTagsFuture, err error)
}

var _ OpenShiftManagedClustersClientAPI = (*containerservice.OpenShiftManagedClustersClient)(nil)

// ContainerServicesClientAPI contains the set of methods on the ContainerServicesClient type.
type ContainerServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters containerservice.ContainerService) (result containerservice.ContainerServicesCreateOrUpdateFutureType, err error)
	Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerServicesDeleteFutureType, err error)
	Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result containerservice.ContainerService, err error)
	List(ctx context.Context) (result containerservice.ListResultPage, err error)
	ListComplete(ctx context.Context) (result containerservice.ListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerservice.ListResultIterator, err error)
	ListOrchestrators(ctx context.Context, location string, resourceType string) (result containerservice.OrchestratorVersionProfileListResult, err error)
}

var _ ContainerServicesClientAPI = (*containerservice.ContainerServicesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerservice.OperationListResult, err error)
}

var _ OperationsClientAPI = (*containerservice.OperationsClient)(nil)

// ManagedClustersClientAPI contains the set of methods on the ManagedClustersClient type.
type ManagedClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedCluster) (result containerservice.ManagedClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedCluster, err error)
	GetAccessProfile(ctx context.Context, resourceGroupName string, resourceName string, roleName string) (result containerservice.ManagedClusterAccessProfile, err error)
	GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClusterUpgradeProfile, err error)
	List(ctx context.Context) (result containerservice.ManagedClusterListResultPage, err error)
	ListComplete(ctx context.Context) (result containerservice.ManagedClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ManagedClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerservice.ManagedClusterListResultIterator, err error)
	ListClusterAdminCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.CredentialResults, err error)
	ListClusterUserCredentials(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.CredentialResults, err error)
	ResetAADProfile(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedClusterAADProfile) (result containerservice.ManagedClustersResetAADProfileFuture, err error)
	ResetServicePrincipalProfile(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedClusterServicePrincipalProfile) (result containerservice.ManagedClustersResetServicePrincipalProfileFuture, err error)
	RotateClusterCertificates(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClustersRotateClusterCertificatesFuture, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.TagsObject) (result containerservice.ManagedClustersUpdateTagsFuture, err error)
}

var _ ManagedClustersClientAPI = (*containerservice.ManagedClustersClient)(nil)

// AgentPoolsClientAPI contains the set of methods on the AgentPoolsClient type.
type AgentPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters containerservice.AgentPool) (result containerservice.AgentPoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result containerservice.AgentPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result containerservice.AgentPool, err error)
	GetAvailableAgentPoolVersions(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.AgentPoolAvailableVersions, err error)
	GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result containerservice.AgentPoolUpgradeProfile, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.AgentPoolListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.AgentPoolListResultIterator, err error)
}

var _ AgentPoolsClientAPI = (*containerservice.AgentPoolsClient)(nil)
