package resourcehealthapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resourcehealth/mgmt/2020-05-01/resourcehealth"
)

// AvailabilityStatusesClientAPI contains the set of methods on the AvailabilityStatusesClient type.
type AvailabilityStatusesClientAPI interface {
	GetByResource(ctx context.Context, resourceURI string, filter string, expand string) (result resourcehealth.AvailabilityStatus, err error)
	List(ctx context.Context, resourceURI string, filter string, expand string) (result resourcehealth.AvailabilityStatusListResultPage, err error)
	ListComplete(ctx context.Context, resourceURI string, filter string, expand string) (result resourcehealth.AvailabilityStatusListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, expand string) (result resourcehealth.AvailabilityStatusListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, expand string) (result resourcehealth.AvailabilityStatusListResultIterator, err error)
	ListBySubscriptionID(ctx context.Context, filter string, expand string) (result resourcehealth.AvailabilityStatusListResultPage, err error)
	ListBySubscriptionIDComplete(ctx context.Context, filter string, expand string) (result resourcehealth.AvailabilityStatusListResultIterator, err error)
}

var _ AvailabilityStatusesClientAPI = (*resourcehealth.AvailabilityStatusesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result resourcehealth.OperationListResult, err error)
}

var _ OperationsClientAPI = (*resourcehealth.OperationsClient)(nil)