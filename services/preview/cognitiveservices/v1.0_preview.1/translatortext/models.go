package translatortext

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/gofrs/uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0_preview.1/translatortext"

// BatchRequest definition for the input batch translation request
type BatchRequest struct {
	Source *SourceInput `json:"source,omitempty"`
	// Targets - Location of the destination for the output
	Targets *[]TargetInput `json:"targets,omitempty"`
	// StorageType - Possible values include: 'Folder', 'File'
	StorageType StorageType `json:"storageType,omitempty"`
}

// BatchStatusDetail job status response
type BatchStatusDetail struct {
	autorest.Response `json:"-"`
	// ID - Id of the operation.
	ID *uuid.UUID `json:"id,omitempty"`
	// CreatedDateTimeUtc - Operation created date time
	CreatedDateTimeUtc *date.Time `json:"createdDateTimeUtc,omitempty"`
	// LastActionDateTimeUtc - Date time in which the operation's status has been updated
	LastActionDateTimeUtc *date.Time `json:"lastActionDateTimeUtc,omitempty"`
	// Status - Possible values include: 'NotStarted', 'Running', 'Succeeded', 'Failed', 'Cancelled', 'Cancelling'
	Status  Status         `json:"status,omitempty"`
	Summary *StatusSummary `json:"summary,omitempty"`
}

// BatchStatusResponse document Status Response
type BatchStatusResponse struct {
	autorest.Response `json:"-"`
	// Value - The summary status of individual operation
	Value *[]BatchStatusDetail `json:"value,omitempty"`
	// NextLink - Url for the next page.  Null if no more pages available
	NextLink *string `json:"@nextLink,omitempty"`
}

// BatchSubmissionRequest job submission batch request
type BatchSubmissionRequest struct {
	// Inputs - The input list of documents or folders containing documents
	Inputs *[]BatchRequest `json:"inputs,omitempty"`
}

// DocumentFilter ...
type DocumentFilter struct {
	// Prefix - A case-sensitive prefix string to filter documents in the source path for translation.
	// For example, when using a Azure storage blob Uri, use the prefix to restrict sub folders for translation.
	Prefix *string `json:"prefix,omitempty"`
	// Suffix - A case-sensitive suffix string to filter documents in the source path for translation.
	// This is most often use for file extensions
	Suffix *string `json:"suffix,omitempty"`
}

// DocumentStatusDetail ...
type DocumentStatusDetail struct {
	autorest.Response `json:"-"`
	// Path - Location of the document or folder
	Path *string `json:"path,omitempty"`
	// CreatedDateTimeUtc - Operation created date time
	CreatedDateTimeUtc *date.Time `json:"createdDateTimeUtc,omitempty"`
	// LastActionDateTimeUtc - Date time in which the operation's status has been updated
	LastActionDateTimeUtc *date.Time `json:"lastActionDateTimeUtc,omitempty"`
	// Status - Possible values include: 'Status1NotStarted', 'Status1Running', 'Status1Succeeded', 'Status1Failed', 'Status1Cancelled', 'Status1Cancelling'
	Status Status1 `json:"status,omitempty"`
	// DetectedLanguage - Detected language of the original document (to be implemented)
	DetectedLanguage *string `json:"detectedLanguage,omitempty"`
	// To - To language
	To    *string  `json:"to,omitempty"`
	Error *ErrorV2 `json:"error,omitempty"`
	// Progress - Progress of the translation if available
	Progress *float64 `json:"progress,omitempty"`
	// ID - Document Id
	ID *uuid.UUID `json:"id,omitempty"`
}

// DocumentStatusResponse document Status Response
type DocumentStatusResponse struct {
	autorest.Response `json:"-"`
	// Value - The detail status of individual documents
	Value *[]DocumentStatusDetail `json:"value,omitempty"`
	// NextLink - Url for the next page.  Null if no more pages available
	NextLink *string `json:"@nextLink,omitempty"`
}

// ErrorResponseV2 contains unified error information used for HTTP responses across any Cognitive Service.
// Instances
// can be created either through Microsoft.CloudAI.Containers.HttpStatusExceptionV2 or by returning it
// directly from
// a controller.
type ErrorResponseV2 struct {
	Error *ErrorV2 `json:"error,omitempty"`
}

// ErrorV2 this contains an outer error with error code, message, details, target and an inner error with
// more descriptive details.
type ErrorV2 struct {
	// Code - Possible values include: 'InvalidRequest', 'InvalidArgument', 'InternalServerError', 'ServiceUnavailable', 'ResourceNotFound', 'Unauthorized', 'RequestRateTooHigh'
	Code Code `json:"code,omitempty"`
	// Message - READ-ONLY; Gets high level error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; Gets the source of the error.
	// For example it would be "documents" or "document id" in case of invalid document.
	Target     *string       `json:"target,omitempty"`
	InnerError *InnerErrorV2 `json:"innerError,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorV2.
func (ev ErrorV2) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ev.Code != "" {
		objectMap["code"] = ev.Code
	}
	if ev.InnerError != nil {
		objectMap["innerError"] = ev.InnerError
	}
	return json.Marshal(objectMap)
}

// FileFormat ...
type FileFormat struct {
	// Format - Name of the format
	Format *string `json:"format,omitempty"`
	// FileExtensions - Supported file extension for this format
	FileExtensions *[]string `json:"fileExtensions,omitempty"`
	// ContentTypes - Supported Content-Types for this format
	ContentTypes *[]string `json:"contentTypes,omitempty"`
	// Versions - Supported Version
	Versions *[]string `json:"versions,omitempty"`
}

// FileFormatListResult base type for List return in our api
type FileFormatListResult struct {
	autorest.Response `json:"-"`
	// Value - list of objects
	Value *[]FileFormat `json:"value,omitempty"`
}

// Glossary glossary / translation memory for the request
type Glossary struct {
	// GlossaryURL - Location of the glossary.
	// We will use the file extension to extract the formating if the format parameter is not supplied.
	//
	// If the translation language pair is not present in the glossary, it will not be applied
	GlossaryURL *string `json:"glossaryUrl,omitempty"`
	// Format - Format
	Format *string `json:"format,omitempty"`
	// Version - Version
	Version *string `json:"version,omitempty"`
}

// InnerErrorV2 new Inner Error format which conforms to Cognitive Services API Guidelines which is
// available at
// https://microsoft.sharepoint.com/%3Aw%3A/t/CognitiveServicesPMO/EUoytcrjuJdKpeOKIK_QRC8BPtUYQpKBi8JsWyeDMRsWlQ?e=CPq8ow.
// This contains required properties ErrorCode, message and optional properties target, details(key value
// pair), inner error(this can be nested).
type InnerErrorV2 struct {
	// Code - READ-ONLY; Gets detailed error code.
	Code *int32 `json:"code,omitempty"`
	// Error - READ-ONLY; Gets detailed error string.
	Error *string `json:"error,omitempty"`
	// Message - READ-ONLY; Gets high level error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; Gets the source of the error.
	// For example it would be "documents" or "document id" in case of invalid document.
	Target     *string       `json:"target,omitempty"`
	InnerError *InnerErrorV2 `json:"innerError,omitempty"`
}

// MarshalJSON is the custom marshaler for InnerErrorV2.
func (iev InnerErrorV2) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if iev.InnerError != nil {
		objectMap["innerError"] = iev.InnerError
	}
	return json.Marshal(objectMap)
}

// SourceInput source of the input documents
type SourceInput struct {
	// SourceURL - Location of the folder / container or single file with your documents
	SourceURL *string         `json:"sourceUrl,omitempty"`
	Filter    *DocumentFilter `json:"filter,omitempty"`
	// Language - Language code
	// If none is specified, we will perform auto detect on the document
	Language *string `json:"language,omitempty"`
	// StorageSource - Possible values include: 'AzureBlob'
	StorageSource StorageSource `json:"storageSource,omitempty"`
}

// StatusSummary ...
type StatusSummary struct {
	// Total - Total count
	Total *int32 `json:"total,omitempty"`
	// Failed - Failed count
	Failed *int32 `json:"failed,omitempty"`
	// Success - Number of Success
	Success *int32 `json:"success,omitempty"`
	// InProgress - Number of in progress
	InProgress *int32 `json:"inProgress,omitempty"`
	// NotYetStarted - Count of not yet started
	NotYetStarted *int32 `json:"notYetStarted,omitempty"`
	// Cancelled - Number of cancelled
	Cancelled *int32 `json:"cancelled,omitempty"`
}

// StorageSourceListResult base type for List return in our api
type StorageSourceListResult struct {
	autorest.Response `json:"-"`
	// Value - list of objects
	Value *[]string `json:"value,omitempty"`
}

// TargetInput destination for the finished translated documents
type TargetInput struct {
	// TargetURL - Location of the folder / container with your documents
	TargetURL *string `json:"targetUrl,omitempty"`
	// Category - Category / custom system for translation request
	Category *string `json:"category,omitempty"`
	// Language - Target Language
	Language *string `json:"language,omitempty"`
	// Glossaries - List of Glossary
	Glossaries *[]Glossary `json:"glossaries,omitempty"`
	// StorageSource - Possible values include: 'StorageSource1AzureBlob'
	StorageSource StorageSource1 `json:"storageSource,omitempty"`
}
