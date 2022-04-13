//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ComplianceStatus.
func (c ComplianceStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceState", c.ComplianceState)
	populateTimeRFC3339(objectMap, "lastConfigApplied", c.LastConfigApplied)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "messageLevel", c.MessageLevel)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComplianceStatus.
func (c *ComplianceStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceState":
			err = unpopulate(val, &c.ComplianceState)
			delete(rawMsg, key)
		case "lastConfigApplied":
			err = unpopulateTimeRFC3339(val, &c.LastConfigApplied)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &c.Message)
			delete(rawMsg, key)
		case "messageLevel":
			err = unpopulate(val, &c.MessageLevel)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ExtensionProperties.
func (e ExtensionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aksAssignedIdentity", e.AksAssignedIdentity)
	populate(objectMap, "autoUpgradeMinorVersion", e.AutoUpgradeMinorVersion)
	populate(objectMap, "configurationProtectedSettings", e.ConfigurationProtectedSettings)
	populate(objectMap, "configurationSettings", e.ConfigurationSettings)
	populate(objectMap, "customLocationSettings", e.CustomLocationSettings)
	populate(objectMap, "errorInfo", e.ErrorInfo)
	populate(objectMap, "extensionType", e.ExtensionType)
	populate(objectMap, "installedVersion", e.InstalledVersion)
	populate(objectMap, "packageUri", e.PackageURI)
	populate(objectMap, "provisioningState", e.ProvisioningState)
	populate(objectMap, "releaseTrain", e.ReleaseTrain)
	populate(objectMap, "scope", e.Scope)
	populate(objectMap, "statuses", e.Statuses)
	populate(objectMap, "version", e.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ExtensionsList.
func (e ExtensionsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", e.NextLink)
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FluxConfigurationPatch.
func (f FluxConfigurationPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", f.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FluxConfigurationPatchProperties.
func (f FluxConfigurationPatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bucket", f.Bucket)
	populate(objectMap, "configurationProtectedSettings", f.ConfigurationProtectedSettings)
	populate(objectMap, "gitRepository", f.GitRepository)
	populate(objectMap, "kustomizations", f.Kustomizations)
	populate(objectMap, "sourceKind", f.SourceKind)
	populate(objectMap, "suspend", f.Suspend)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FluxConfigurationProperties.
func (f FluxConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bucket", f.Bucket)
	populate(objectMap, "complianceState", f.ComplianceState)
	populate(objectMap, "configurationProtectedSettings", f.ConfigurationProtectedSettings)
	populate(objectMap, "errorMessage", f.ErrorMessage)
	populate(objectMap, "gitRepository", f.GitRepository)
	populate(objectMap, "kustomizations", f.Kustomizations)
	populate(objectMap, "namespace", f.Namespace)
	populate(objectMap, "provisioningState", f.ProvisioningState)
	populate(objectMap, "repositoryPublicKey", f.RepositoryPublicKey)
	populate(objectMap, "scope", f.Scope)
	populate(objectMap, "sourceKind", f.SourceKind)
	populate(objectMap, "sourceSyncedCommitId", f.SourceSyncedCommitID)
	populateTimeRFC3339(objectMap, "sourceUpdatedAt", f.SourceUpdatedAt)
	populateTimeRFC3339(objectMap, "statusUpdatedAt", f.StatusUpdatedAt)
	populate(objectMap, "statuses", f.Statuses)
	populate(objectMap, "suspend", f.Suspend)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FluxConfigurationProperties.
func (f *FluxConfigurationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "bucket":
			err = unpopulate(val, &f.Bucket)
			delete(rawMsg, key)
		case "complianceState":
			err = unpopulate(val, &f.ComplianceState)
			delete(rawMsg, key)
		case "configurationProtectedSettings":
			err = unpopulate(val, &f.ConfigurationProtectedSettings)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, &f.ErrorMessage)
			delete(rawMsg, key)
		case "gitRepository":
			err = unpopulate(val, &f.GitRepository)
			delete(rawMsg, key)
		case "kustomizations":
			err = unpopulate(val, &f.Kustomizations)
			delete(rawMsg, key)
		case "namespace":
			err = unpopulate(val, &f.Namespace)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &f.ProvisioningState)
			delete(rawMsg, key)
		case "repositoryPublicKey":
			err = unpopulate(val, &f.RepositoryPublicKey)
			delete(rawMsg, key)
		case "scope":
			err = unpopulate(val, &f.Scope)
			delete(rawMsg, key)
		case "sourceKind":
			err = unpopulate(val, &f.SourceKind)
			delete(rawMsg, key)
		case "sourceSyncedCommitId":
			err = unpopulate(val, &f.SourceSyncedCommitID)
			delete(rawMsg, key)
		case "sourceUpdatedAt":
			err = unpopulateTimeRFC3339(val, &f.SourceUpdatedAt)
			delete(rawMsg, key)
		case "statusUpdatedAt":
			err = unpopulateTimeRFC3339(val, &f.StatusUpdatedAt)
			delete(rawMsg, key)
		case "statuses":
			err = unpopulate(val, &f.Statuses)
			delete(rawMsg, key)
		case "suspend":
			err = unpopulate(val, &f.Suspend)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FluxConfigurationsList.
func (f FluxConfigurationsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", f.NextLink)
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KustomizationDefinition.
func (k KustomizationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dependsOn", k.DependsOn)
	populate(objectMap, "force", k.Force)
	populate(objectMap, "name", k.Name)
	populate(objectMap, "path", k.Path)
	populate(objectMap, "prune", k.Prune)
	populate(objectMap, "retryIntervalInSeconds", k.RetryIntervalInSeconds)
	populate(objectMap, "syncIntervalInSeconds", k.SyncIntervalInSeconds)
	populate(objectMap, "timeoutInSeconds", k.TimeoutInSeconds)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KustomizationPatchDefinition.
func (k KustomizationPatchDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dependsOn", k.DependsOn)
	populate(objectMap, "force", k.Force)
	populate(objectMap, "path", k.Path)
	populate(objectMap, "prune", k.Prune)
	populate(objectMap, "retryIntervalInSeconds", k.RetryIntervalInSeconds)
	populate(objectMap, "syncIntervalInSeconds", k.SyncIntervalInSeconds)
	populate(objectMap, "timeoutInSeconds", k.TimeoutInSeconds)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ObjectStatusConditionDefinition.
func (o ObjectStatusConditionDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "lastTransitionTime", o.LastTransitionTime)
	populate(objectMap, "message", o.Message)
	populate(objectMap, "reason", o.Reason)
	populate(objectMap, "status", o.Status)
	populate(objectMap, "type", o.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ObjectStatusConditionDefinition.
func (o *ObjectStatusConditionDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "lastTransitionTime":
			err = unpopulateTimeRFC3339(val, &o.LastTransitionTime)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &o.Message)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &o.Reason)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &o.Status)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &o.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ObjectStatusDefinition.
func (o ObjectStatusDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "appliedBy", o.AppliedBy)
	populate(objectMap, "complianceState", o.ComplianceState)
	populate(objectMap, "helmReleaseProperties", o.HelmReleaseProperties)
	populate(objectMap, "kind", o.Kind)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "namespace", o.Namespace)
	populate(objectMap, "statusConditions", o.StatusConditions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationStatusList.
func (o OperationStatusList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationStatusResult.
func (o OperationStatusResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", o.Error)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "status", o.Status)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PatchExtension.
func (p PatchExtension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", p.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PatchExtensionProperties.
func (p PatchExtensionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoUpgradeMinorVersion", p.AutoUpgradeMinorVersion)
	populate(objectMap, "configurationProtectedSettings", p.ConfigurationProtectedSettings)
	populate(objectMap, "configurationSettings", p.ConfigurationSettings)
	populate(objectMap, "releaseTrain", p.ReleaseTrain)
	populate(objectMap, "version", p.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderOperationList.
func (r ResourceProviderOperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SourceControlConfigurationList.
func (s SourceControlConfigurationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SourceControlConfigurationProperties.
func (s SourceControlConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceStatus", s.ComplianceStatus)
	populate(objectMap, "configurationProtectedSettings", s.ConfigurationProtectedSettings)
	populate(objectMap, "enableHelmOperator", s.EnableHelmOperator)
	populate(objectMap, "helmOperatorProperties", s.HelmOperatorProperties)
	populate(objectMap, "operatorInstanceName", s.OperatorInstanceName)
	populate(objectMap, "operatorNamespace", s.OperatorNamespace)
	populate(objectMap, "operatorParams", s.OperatorParams)
	populate(objectMap, "operatorScope", s.OperatorScope)
	populate(objectMap, "operatorType", s.OperatorType)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "repositoryPublicKey", s.RepositoryPublicKey)
	populate(objectMap, "repositoryUrl", s.RepositoryURL)
	populate(objectMap, "sshKnownHostsContents", s.SSHKnownHostsContents)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}