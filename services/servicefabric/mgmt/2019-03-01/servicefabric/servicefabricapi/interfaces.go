package servicefabricapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2019-03-01/servicefabric"
	"github.com/Azure/go-autorest/autorest"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters servicefabric.Cluster) (result servicefabric.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result servicefabric.Cluster, err error)
	List(ctx context.Context) (result servicefabric.ClusterListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicefabric.ClusterListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters servicefabric.ClusterUpdateParameters) (result servicefabric.ClustersUpdateFuture, err error)
}

var _ ClustersClientAPI = (*servicefabric.ClustersClient)(nil)

// ClusterVersionsClientAPI contains the set of methods on the ClusterVersionsClient type.
type ClusterVersionsClientAPI interface {
	Get(ctx context.Context, location string, clusterVersion string) (result servicefabric.ClusterCodeVersionsListResult, err error)
	GetByEnvironment(ctx context.Context, location string, environment string, clusterVersion string) (result servicefabric.ClusterCodeVersionsListResult, err error)
	List(ctx context.Context, location string) (result servicefabric.ClusterCodeVersionsListResult, err error)
	ListByEnvironment(ctx context.Context, location string, environment string) (result servicefabric.ClusterCodeVersionsListResult, err error)
}

var _ ClusterVersionsClientAPI = (*servicefabric.ClusterVersionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context, APIVersion string) (result servicefabric.OperationListResultPage, err error)
	ListComplete(ctx context.Context, APIVersion string) (result servicefabric.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*servicefabric.OperationsClient)(nil)

// ApplicationTypesClientAPI contains the set of methods on the ApplicationTypesClient type.
type ApplicationTypesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters servicefabric.ApplicationTypeResource) (result servicefabric.ApplicationTypeResource, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (result servicefabric.ApplicationTypesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (result servicefabric.ApplicationTypeResource, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string) (result servicefabric.ApplicationTypeResourceList, err error)
}

var _ ApplicationTypesClientAPI = (*servicefabric.ApplicationTypesClient)(nil)

// ApplicationTypeVersionsClientAPI contains the set of methods on the ApplicationTypeVersionsClient type.
type ApplicationTypeVersionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters servicefabric.ApplicationTypeVersionResource) (result servicefabric.ApplicationTypeVersionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result servicefabric.ApplicationTypeVersionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result servicefabric.ApplicationTypeVersionResource, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string) (result servicefabric.ApplicationTypeVersionResourceList, err error)
}

var _ ApplicationTypeVersionsClientAPI = (*servicefabric.ApplicationTypeVersionsClient)(nil)

// ApplicationsClientAPI contains the set of methods on the ApplicationsClient type.
type ApplicationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters servicefabric.ApplicationResource) (result servicefabric.ApplicationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result servicefabric.ApplicationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result servicefabric.ApplicationResource, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string) (result servicefabric.ApplicationResourceList, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters servicefabric.ApplicationResourceUpdate) (result servicefabric.ApplicationsUpdateFuture, err error)
}

var _ ApplicationsClientAPI = (*servicefabric.ApplicationsClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters servicefabric.ServiceResource) (result servicefabric.ServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result servicefabric.ServicesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result servicefabric.ServiceResource, err error)
	List(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result servicefabric.ServiceResourceList, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters servicefabric.ServiceResourceUpdate) (result servicefabric.ServicesUpdateFuture, err error)
}

var _ ServicesClientAPI = (*servicefabric.ServicesClient)(nil)
