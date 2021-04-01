// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package policy

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-12-01/policy"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Mode = original.Mode

const (
	All          Mode = original.All
	Indexed      Mode = original.Indexed
	NotSpecified Mode = original.NotSpecified
)

type Type = original.Type

const (
	TypeBuiltIn      Type = original.TypeBuiltIn
	TypeCustom       Type = original.TypeCustom
	TypeNotSpecified Type = original.TypeNotSpecified
)

type Assignment = original.Assignment
type AssignmentListResult = original.AssignmentListResult
type AssignmentListResultIterator = original.AssignmentListResultIterator
type AssignmentListResultPage = original.AssignmentListResultPage
type AssignmentProperties = original.AssignmentProperties
type AssignmentsClient = original.AssignmentsClient
type BaseClient = original.BaseClient
type Definition = original.Definition
type DefinitionListResult = original.DefinitionListResult
type DefinitionListResultIterator = original.DefinitionListResultIterator
type DefinitionListResultPage = original.DefinitionListResultPage
type DefinitionProperties = original.DefinitionProperties
type DefinitionsClient = original.DefinitionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAssignmentListResultIterator(page AssignmentListResultPage) AssignmentListResultIterator {
	return original.NewAssignmentListResultIterator(page)
}
func NewAssignmentListResultPage(cur AssignmentListResult, getNextPage func(context.Context, AssignmentListResult) (AssignmentListResult, error)) AssignmentListResultPage {
	return original.NewAssignmentListResultPage(cur, getNextPage)
}
func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClient(subscriptionID)
}
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDefinitionListResultIterator(page DefinitionListResultPage) DefinitionListResultIterator {
	return original.NewDefinitionListResultIterator(page)
}
func NewDefinitionListResultPage(cur DefinitionListResult, getNextPage func(context.Context, DefinitionListResult) (DefinitionListResult, error)) DefinitionListResultPage {
	return original.NewDefinitionListResultPage(cur, getNextPage)
}
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClient(subscriptionID)
}
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleModeValues() []Mode {
	return original.PossibleModeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/2020-09-01"
}
func Version() string {
	return original.Version()
}
