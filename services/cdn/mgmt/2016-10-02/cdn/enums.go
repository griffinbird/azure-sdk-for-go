package cdn

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CustomDomainResourceState enumerates the values for custom domain resource state.
type CustomDomainResourceState string

const (
	// Active ...
	Active CustomDomainResourceState = "Active"
	// Creating ...
	Creating CustomDomainResourceState = "Creating"
	// Deleting ...
	Deleting CustomDomainResourceState = "Deleting"
)

// PossibleCustomDomainResourceStateValues returns an array of possible values for the CustomDomainResourceState const type.
func PossibleCustomDomainResourceStateValues() []CustomDomainResourceState {
	return []CustomDomainResourceState{Active, Creating, Deleting}
}

// CustomHTTPSProvisioningState enumerates the values for custom https provisioning state.
type CustomHTTPSProvisioningState string

const (
	// Disabled ...
	Disabled CustomHTTPSProvisioningState = "Disabled"
	// Disabling ...
	Disabling CustomHTTPSProvisioningState = "Disabling"
	// Enabled ...
	Enabled CustomHTTPSProvisioningState = "Enabled"
	// Enabling ...
	Enabling CustomHTTPSProvisioningState = "Enabling"
	// Failed ...
	Failed CustomHTTPSProvisioningState = "Failed"
)

// PossibleCustomHTTPSProvisioningStateValues returns an array of possible values for the CustomHTTPSProvisioningState const type.
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return []CustomHTTPSProvisioningState{Disabled, Disabling, Enabled, Enabling, Failed}
}

// EndpointResourceState enumerates the values for endpoint resource state.
type EndpointResourceState string

const (
	// EndpointResourceStateCreating ...
	EndpointResourceStateCreating EndpointResourceState = "Creating"
	// EndpointResourceStateDeleting ...
	EndpointResourceStateDeleting EndpointResourceState = "Deleting"
	// EndpointResourceStateRunning ...
	EndpointResourceStateRunning EndpointResourceState = "Running"
	// EndpointResourceStateStarting ...
	EndpointResourceStateStarting EndpointResourceState = "Starting"
	// EndpointResourceStateStopped ...
	EndpointResourceStateStopped EndpointResourceState = "Stopped"
	// EndpointResourceStateStopping ...
	EndpointResourceStateStopping EndpointResourceState = "Stopping"
)

// PossibleEndpointResourceStateValues returns an array of possible values for the EndpointResourceState const type.
func PossibleEndpointResourceStateValues() []EndpointResourceState {
	return []EndpointResourceState{EndpointResourceStateCreating, EndpointResourceStateDeleting, EndpointResourceStateRunning, EndpointResourceStateStarting, EndpointResourceStateStopped, EndpointResourceStateStopping}
}

// GeoFilterActions enumerates the values for geo filter actions.
type GeoFilterActions string

const (
	// Allow ...
	Allow GeoFilterActions = "Allow"
	// Block ...
	Block GeoFilterActions = "Block"
)

// PossibleGeoFilterActionsValues returns an array of possible values for the GeoFilterActions const type.
func PossibleGeoFilterActionsValues() []GeoFilterActions {
	return []GeoFilterActions{Allow, Block}
}

// OptimizationType enumerates the values for optimization type.
type OptimizationType string

const (
	// DynamicSiteAcceleration ...
	DynamicSiteAcceleration OptimizationType = "DynamicSiteAcceleration"
	// GeneralMediaStreaming ...
	GeneralMediaStreaming OptimizationType = "GeneralMediaStreaming"
	// GeneralWebDelivery ...
	GeneralWebDelivery OptimizationType = "GeneralWebDelivery"
	// LargeFileDownload ...
	LargeFileDownload OptimizationType = "LargeFileDownload"
	// VideoOnDemandMediaStreaming ...
	VideoOnDemandMediaStreaming OptimizationType = "VideoOnDemandMediaStreaming"
)

// PossibleOptimizationTypeValues returns an array of possible values for the OptimizationType const type.
func PossibleOptimizationTypeValues() []OptimizationType {
	return []OptimizationType{DynamicSiteAcceleration, GeneralMediaStreaming, GeneralWebDelivery, LargeFileDownload, VideoOnDemandMediaStreaming}
}

// OriginResourceState enumerates the values for origin resource state.
type OriginResourceState string

const (
	// OriginResourceStateActive ...
	OriginResourceStateActive OriginResourceState = "Active"
	// OriginResourceStateCreating ...
	OriginResourceStateCreating OriginResourceState = "Creating"
	// OriginResourceStateDeleting ...
	OriginResourceStateDeleting OriginResourceState = "Deleting"
)

// PossibleOriginResourceStateValues returns an array of possible values for the OriginResourceState const type.
func PossibleOriginResourceStateValues() []OriginResourceState {
	return []OriginResourceState{OriginResourceStateActive, OriginResourceStateCreating, OriginResourceStateDeleting}
}

// ProfileResourceState enumerates the values for profile resource state.
type ProfileResourceState string

const (
	// ProfileResourceStateActive ...
	ProfileResourceStateActive ProfileResourceState = "Active"
	// ProfileResourceStateCreating ...
	ProfileResourceStateCreating ProfileResourceState = "Creating"
	// ProfileResourceStateDeleting ...
	ProfileResourceStateDeleting ProfileResourceState = "Deleting"
	// ProfileResourceStateDisabled ...
	ProfileResourceStateDisabled ProfileResourceState = "Disabled"
)

// PossibleProfileResourceStateValues returns an array of possible values for the ProfileResourceState const type.
func PossibleProfileResourceStateValues() []ProfileResourceState {
	return []ProfileResourceState{ProfileResourceStateActive, ProfileResourceStateCreating, ProfileResourceStateDeleting, ProfileResourceStateDisabled}
}

// QueryStringCachingBehavior enumerates the values for query string caching behavior.
type QueryStringCachingBehavior string

const (
	// BypassCaching ...
	BypassCaching QueryStringCachingBehavior = "BypassCaching"
	// IgnoreQueryString ...
	IgnoreQueryString QueryStringCachingBehavior = "IgnoreQueryString"
	// NotSet ...
	NotSet QueryStringCachingBehavior = "NotSet"
	// UseQueryString ...
	UseQueryString QueryStringCachingBehavior = "UseQueryString"
)

// PossibleQueryStringCachingBehaviorValues returns an array of possible values for the QueryStringCachingBehavior const type.
func PossibleQueryStringCachingBehaviorValues() []QueryStringCachingBehavior {
	return []QueryStringCachingBehavior{BypassCaching, IgnoreQueryString, NotSet, UseQueryString}
}

// ResourceType enumerates the values for resource type.
type ResourceType string

const (
	// MicrosoftCdnProfilesEndpoints ...
	MicrosoftCdnProfilesEndpoints ResourceType = "Microsoft.Cdn/Profiles/Endpoints"
)

// PossibleResourceTypeValues returns an array of possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{MicrosoftCdnProfilesEndpoints}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// CustomVerizon ...
	CustomVerizon SkuName = "Custom_Verizon"
	// PremiumVerizon ...
	PremiumVerizon SkuName = "Premium_Verizon"
	// StandardAkamai ...
	StandardAkamai SkuName = "Standard_Akamai"
	// StandardChinaCdn ...
	StandardChinaCdn SkuName = "Standard_ChinaCdn"
	// StandardVerizon ...
	StandardVerizon SkuName = "Standard_Verizon"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{CustomVerizon, PremiumVerizon, StandardAkamai, StandardChinaCdn, StandardVerizon}
}
