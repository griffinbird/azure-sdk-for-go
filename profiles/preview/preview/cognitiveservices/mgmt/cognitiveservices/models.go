// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package cognitiveservices

import original "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/mgmt/2016-02-01-preview/cognitiveservices"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type KeyName = original.KeyName

const (
	Key1 KeyName = original.Key1
	Key2 KeyName = original.Key2
)

type Kind = original.Kind

const (
	Academic           Kind = original.Academic
	BingAutosuggest    Kind = original.BingAutosuggest
	BingSearch         Kind = original.BingSearch
	BingSpeech         Kind = original.BingSpeech
	BingSpellCheck     Kind = original.BingSpellCheck
	ComputerVision     Kind = original.ComputerVision
	ContentModerator   Kind = original.ContentModerator
	Emotion            Kind = original.Emotion
	Face               Kind = original.Face
	LUIS               Kind = original.LUIS
	Recommendations    Kind = original.Recommendations
	SpeakerRecognition Kind = original.SpeakerRecognition
	Speech             Kind = original.Speech
	SpeechTranslation  Kind = original.SpeechTranslation
	TextAnalytics      Kind = original.TextAnalytics
	TextTranslation    Kind = original.TextTranslation
	WebLM              Kind = original.WebLM
)

type ProvisioningState = original.ProvisioningState

const (
	Creating     ProvisioningState = original.Creating
	Failed       ProvisioningState = original.Failed
	ResolvingDNS ProvisioningState = original.ResolvingDNS
	Succeeded    ProvisioningState = original.Succeeded
)

type SkuName = original.SkuName

const (
	F0 SkuName = original.F0
	P0 SkuName = original.P0
	P1 SkuName = original.P1
	P2 SkuName = original.P2
	S0 SkuName = original.S0
	S1 SkuName = original.S1
	S2 SkuName = original.S2
	S3 SkuName = original.S3
	S4 SkuName = original.S4
	S5 SkuName = original.S5
	S6 SkuName = original.S6
)

type SkuTier = original.SkuTier

const (
	Free     SkuTier = original.Free
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type Account = original.Account
type AccountCreateParameters = original.AccountCreateParameters
type AccountEnumerateSkusResult = original.AccountEnumerateSkusResult
type AccountKeys = original.AccountKeys
type AccountListResult = original.AccountListResult
type AccountProperties = original.AccountProperties
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountsClient = original.AccountsClient
type BaseClient = original.BaseClient
type Error = original.Error
type ErrorBody = original.ErrorBody
type RegenerateKeyParameters = original.RegenerateKeyParameters
type ResourceAndSku = original.ResourceAndSku
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleKeyNameValues() []KeyName {
	return original.PossibleKeyNameValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
