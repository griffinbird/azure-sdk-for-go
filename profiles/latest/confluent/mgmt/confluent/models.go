// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package confluent

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/confluent/mgmt/2020-03-01/confluent"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisionState = original.ProvisionState

const (
	Accepted     ProvisionState = original.Accepted
	Canceled     ProvisionState = original.Canceled
	Creating     ProvisionState = original.Creating
	Deleted      ProvisionState = original.Deleted
	Deleting     ProvisionState = original.Deleting
	Failed       ProvisionState = original.Failed
	NotSpecified ProvisionState = original.NotSpecified
	Succeeded    ProvisionState = original.Succeeded
	Updating     ProvisionState = original.Updating
)

type SaaSOfferStatus = original.SaaSOfferStatus

const (
	SaaSOfferStatusFailed                  SaaSOfferStatus = original.SaaSOfferStatusFailed
	SaaSOfferStatusInProgress              SaaSOfferStatus = original.SaaSOfferStatusInProgress
	SaaSOfferStatusPendingFulfillmentStart SaaSOfferStatus = original.SaaSOfferStatusPendingFulfillmentStart
	SaaSOfferStatusReinstated              SaaSOfferStatus = original.SaaSOfferStatusReinstated
	SaaSOfferStatusStarted                 SaaSOfferStatus = original.SaaSOfferStatusStarted
	SaaSOfferStatusSubscribed              SaaSOfferStatus = original.SaaSOfferStatusSubscribed
	SaaSOfferStatusSucceeded               SaaSOfferStatus = original.SaaSOfferStatusSucceeded
	SaaSOfferStatusSuspended               SaaSOfferStatus = original.SaaSOfferStatusSuspended
	SaaSOfferStatusUnsubscribed            SaaSOfferStatus = original.SaaSOfferStatusUnsubscribed
	SaaSOfferStatusUpdating                SaaSOfferStatus = original.SaaSOfferStatusUpdating
)

type AgreementProperties = original.AgreementProperties
type AgreementResource = original.AgreementResource
type AgreementResourceListResponse = original.AgreementResourceListResponse
type AgreementResourceListResponseIterator = original.AgreementResourceListResponseIterator
type AgreementResourceListResponsePage = original.AgreementResourceListResponsePage
type BaseClient = original.BaseClient
type ErrorResponseBody = original.ErrorResponseBody
type MarketplaceAgreementsClient = original.MarketplaceAgreementsClient
type OfferDetail = original.OfferDetail
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OrganizationClient = original.OrganizationClient
type OrganizationCreateFuture = original.OrganizationCreateFuture
type OrganizationDeleteFuture = original.OrganizationDeleteFuture
type OrganizationOperationsClient = original.OrganizationOperationsClient
type OrganizationResource = original.OrganizationResource
type OrganizationResourceListResult = original.OrganizationResourceListResult
type OrganizationResourceListResultIterator = original.OrganizationResourceListResultIterator
type OrganizationResourceListResultPage = original.OrganizationResourceListResultPage
type OrganizationResourceProperties = original.OrganizationResourceProperties
type OrganizationResourcePropertiesModel = original.OrganizationResourcePropertiesModel
type OrganizationResourcePropertiesOfferDetail = original.OrganizationResourcePropertiesOfferDetail
type OrganizationResourcePropertiesUserDetail = original.OrganizationResourcePropertiesUserDetail
type OrganizationResourceUpdate = original.OrganizationResourceUpdate
type ResourceProviderDefaultErrorResponse = original.ResourceProviderDefaultErrorResponse
type UserDetail = original.UserDetail

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAgreementResourceListResponseIterator(page AgreementResourceListResponsePage) AgreementResourceListResponseIterator {
	return original.NewAgreementResourceListResponseIterator(page)
}
func NewAgreementResourceListResponsePage(cur AgreementResourceListResponse, getNextPage func(context.Context, AgreementResourceListResponse) (AgreementResourceListResponse, error)) AgreementResourceListResponsePage {
	return original.NewAgreementResourceListResponsePage(cur, getNextPage)
}
func NewMarketplaceAgreementsClient(subscriptionID string) MarketplaceAgreementsClient {
	return original.NewMarketplaceAgreementsClient(subscriptionID)
}
func NewMarketplaceAgreementsClientWithBaseURI(baseURI string, subscriptionID string) MarketplaceAgreementsClient {
	return original.NewMarketplaceAgreementsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOrganizationClient(subscriptionID string) OrganizationClient {
	return original.NewOrganizationClient(subscriptionID)
}
func NewOrganizationClientWithBaseURI(baseURI string, subscriptionID string) OrganizationClient {
	return original.NewOrganizationClientWithBaseURI(baseURI, subscriptionID)
}
func NewOrganizationOperationsClient(subscriptionID string) OrganizationOperationsClient {
	return original.NewOrganizationOperationsClient(subscriptionID)
}
func NewOrganizationOperationsClientWithBaseURI(baseURI string, subscriptionID string) OrganizationOperationsClient {
	return original.NewOrganizationOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOrganizationResourceListResultIterator(page OrganizationResourceListResultPage) OrganizationResourceListResultIterator {
	return original.NewOrganizationResourceListResultIterator(page)
}
func NewOrganizationResourceListResultPage(cur OrganizationResourceListResult, getNextPage func(context.Context, OrganizationResourceListResult) (OrganizationResourceListResult, error)) OrganizationResourceListResultPage {
	return original.NewOrganizationResourceListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleProvisionStateValues() []ProvisionState {
	return original.PossibleProvisionStateValues()
}
func PossibleSaaSOfferStatusValues() []SaaSOfferStatus {
	return original.PossibleSaaSOfferStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
