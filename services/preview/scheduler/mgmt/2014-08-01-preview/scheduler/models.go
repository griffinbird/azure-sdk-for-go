package scheduler

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler"

// BasicAuthentication ...
type BasicAuthentication struct {
	// Username - Gets or sets the username.
	Username *string `json:"username,omitempty"`
	// Password - Gets or sets the password.
	Password *string `json:"password,omitempty"`
	// Type - Gets or sets the http authentication type. Possible values include: 'NotSpecified', 'ClientCertificate', 'ActiveDirectoryOAuth', 'Basic'
	Type HTTPAuthenticationType `json:"type,omitempty"`
}

// ClientCertAuthentication ...
type ClientCertAuthentication struct {
	// Password - Gets or sets the password.
	Password *string `json:"password,omitempty"`
	// Pfx - Gets or sets the pfx.
	Pfx *string `json:"pfx,omitempty"`
	// CertificateThumbprint - Gets or sets the certificate thumbprint.
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// CertificateExpirationDate - Gets or sets the certificate expiration date.
	CertificateExpirationDate *date.Time `json:"certificateExpirationDate,omitempty"`
	// CertificateSubjectName - Gets or sets the certificate subject name.
	CertificateSubjectName *string `json:"certificateSubjectName,omitempty"`
	// Type - Gets or sets the http authentication type. Possible values include: 'NotSpecified', 'ClientCertificate', 'ActiveDirectoryOAuth', 'Basic'
	Type HTTPAuthenticationType `json:"type,omitempty"`
}

// HTTPAuthentication ...
type HTTPAuthentication struct {
	// Type - Gets or sets the http authentication type. Possible values include: 'NotSpecified', 'ClientCertificate', 'ActiveDirectoryOAuth', 'Basic'
	Type HTTPAuthenticationType `json:"type,omitempty"`
}

// HTTPRequest ...
type HTTPRequest struct {
	// Authentication - Gets or sets the http authentication.
	Authentication *HTTPAuthentication `json:"authentication,omitempty"`
	// URI - Gets or sets the Uri.
	URI *string `json:"uri,omitempty"`
	// Method - Gets or sets the method of the request.
	Method *string `json:"method,omitempty"`
	// Body - Gets or sets the request body.
	Body *string `json:"body,omitempty"`
	// Headers - Gets or sets the headers.
	Headers map[string]*string `json:"headers"`
}

// MarshalJSON is the custom marshaler for HTTPRequest.
func (hr HTTPRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if hr.Authentication != nil {
		objectMap["authentication"] = hr.Authentication
	}
	if hr.URI != nil {
		objectMap["uri"] = hr.URI
	}
	if hr.Method != nil {
		objectMap["method"] = hr.Method
	}
	if hr.Body != nil {
		objectMap["body"] = hr.Body
	}
	if hr.Headers != nil {
		objectMap["headers"] = hr.Headers
	}
	return json.Marshal(objectMap)
}

// JobAction ...
type JobAction struct {
	// Type - Gets or sets the job action type. Possible values include: 'HTTP', 'HTTPS', 'StorageQueue', 'ServiceBusQueue', 'ServiceBusTopic'
	Type JobActionType `json:"type,omitempty"`
	// Request - Gets or sets the http requests.
	Request *HTTPRequest `json:"request,omitempty"`
	// QueueMessage - Gets or sets the storage queue message.
	QueueMessage *StorageQueueMessage `json:"queueMessage,omitempty"`
	// ServiceBusQueueMessage - Gets or sets the service bus queue message.
	ServiceBusQueueMessage *ServiceBusQueueMessage `json:"serviceBusQueueMessage,omitempty"`
	// ServiceBusTopicMessage - Gets or sets the service bus topic message.
	ServiceBusTopicMessage *ServiceBusTopicMessage `json:"serviceBusTopicMessage,omitempty"`
	// RetryPolicy - Gets or sets the retry policy.
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`
	// ErrorAction - Gets or sets the error action.
	ErrorAction *JobErrorAction `json:"errorAction,omitempty"`
}

// JobCollectionDefinition ...
type JobCollectionDefinition struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets the job collection resource identifier.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Gets the job collection resource type.
	Type *string `json:"type,omitempty"`
	// Name - Gets or sets the job collection resource name.
	Name *string `json:"name,omitempty"`
	// Location - Gets or sets the storage account location.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets the tags.
	Tags map[string]*string `json:"tags"`
	// Properties - Gets or sets the job collection properties.
	Properties *JobCollectionProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for JobCollectionDefinition.
func (jcd JobCollectionDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if jcd.Name != nil {
		objectMap["name"] = jcd.Name
	}
	if jcd.Location != nil {
		objectMap["location"] = jcd.Location
	}
	if jcd.Tags != nil {
		objectMap["tags"] = jcd.Tags
	}
	if jcd.Properties != nil {
		objectMap["properties"] = jcd.Properties
	}
	return json.Marshal(objectMap)
}

// JobCollectionListResult ...
type JobCollectionListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Gets the job collections.
	Value *[]JobCollectionDefinition `json:"value,omitempty"`
	// NextLink - Gets or sets the URL to get the next set of job collections.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for JobCollectionListResult.
func (jclr JobCollectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if jclr.NextLink != nil {
		objectMap["nextLink"] = jclr.NextLink
	}
	return json.Marshal(objectMap)
}

// JobCollectionListResultIterator provides access to a complete listing of JobCollectionDefinition values.
type JobCollectionListResultIterator struct {
	i    int
	page JobCollectionListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *JobCollectionListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCollectionListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *JobCollectionListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter JobCollectionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter JobCollectionListResultIterator) Response() JobCollectionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter JobCollectionListResultIterator) Value() JobCollectionDefinition {
	if !iter.page.NotDone() {
		return JobCollectionDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the JobCollectionListResultIterator type.
func NewJobCollectionListResultIterator(page JobCollectionListResultPage) JobCollectionListResultIterator {
	return JobCollectionListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (jclr JobCollectionListResult) IsEmpty() bool {
	return jclr.Value == nil || len(*jclr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (jclr JobCollectionListResult) hasNextLink() bool {
	return jclr.NextLink != nil && len(*jclr.NextLink) != 0
}

// jobCollectionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (jclr JobCollectionListResult) jobCollectionListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !jclr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(jclr.NextLink)))
}

// JobCollectionListResultPage contains a page of JobCollectionDefinition values.
type JobCollectionListResultPage struct {
	fn   func(context.Context, JobCollectionListResult) (JobCollectionListResult, error)
	jclr JobCollectionListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *JobCollectionListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobCollectionListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.jclr)
		if err != nil {
			return err
		}
		page.jclr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *JobCollectionListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page JobCollectionListResultPage) NotDone() bool {
	return !page.jclr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page JobCollectionListResultPage) Response() JobCollectionListResult {
	return page.jclr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page JobCollectionListResultPage) Values() []JobCollectionDefinition {
	if page.jclr.IsEmpty() {
		return nil
	}
	return *page.jclr.Value
}

// Creates a new instance of the JobCollectionListResultPage type.
func NewJobCollectionListResultPage(cur JobCollectionListResult, getNextPage func(context.Context, JobCollectionListResult) (JobCollectionListResult, error)) JobCollectionListResultPage {
	return JobCollectionListResultPage{
		fn:   getNextPage,
		jclr: cur,
	}
}

// JobCollectionProperties ...
type JobCollectionProperties struct {
	// Sku - Gets or sets the SKU.
	Sku *Sku `json:"sku,omitempty"`
	// State - Gets or sets the state. Possible values include: 'Enabled', 'Disabled', 'Suspended', 'Deleted'
	State JobCollectionState `json:"state,omitempty"`
	// Quota - Gets or sets the job collection quota.
	Quota *JobCollectionQuota `json:"quota,omitempty"`
}

// JobCollectionQuota ...
type JobCollectionQuota struct {
	// MaxJobCount - Gets or set the maximum job count.
	MaxJobCount *int32 `json:"maxJobCount,omitempty"`
	// MaxJobOccurrence - Gets or sets the maximum job occurrence.
	MaxJobOccurrence *int32 `json:"maxJobOccurrence,omitempty"`
	// MaxRecurrence - Gets or set the maximum recurrence.
	MaxRecurrence *JobMaxRecurrence `json:"maxRecurrence,omitempty"`
}

// JobDefinition ...
type JobDefinition struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets the job resource identifier.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Gets the job resource type.
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Gets the job resource name.
	Name *string `json:"name,omitempty"`
	// Properties - Gets or sets the job properties.
	Properties *JobProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for JobDefinition.
func (jd JobDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if jd.Properties != nil {
		objectMap["properties"] = jd.Properties
	}
	return json.Marshal(objectMap)
}

// JobErrorAction ...
type JobErrorAction struct {
	// Type - Gets or sets the job error action type. Possible values include: 'HTTP', 'HTTPS', 'StorageQueue', 'ServiceBusQueue', 'ServiceBusTopic'
	Type JobActionType `json:"type,omitempty"`
	// Request - Gets or sets the http requests.
	Request *HTTPRequest `json:"request,omitempty"`
	// QueueMessage - Gets or sets the storage queue message.
	QueueMessage *StorageQueueMessage `json:"queueMessage,omitempty"`
	// ServiceBusQueueMessage - Gets or sets the service bus queue message.
	ServiceBusQueueMessage *ServiceBusQueueMessage `json:"serviceBusQueueMessage,omitempty"`
	// ServiceBusTopicMessage - Gets or sets the service bus topic message.
	ServiceBusTopicMessage *ServiceBusTopicMessage `json:"serviceBusTopicMessage,omitempty"`
	// RetryPolicy - Gets or sets the retry policy.
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`
}

// JobHistoryDefinition ...
type JobHistoryDefinition struct {
	// ID - READ-ONLY; Gets the job history identifier.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; Gets the job history resource type.
	Type *string `json:"type,omitempty"`
	// Name - READ-ONLY; Gets the job history name.
	Name *string `json:"name,omitempty"`
	// Properties - READ-ONLY; Gets or sets the job history properties.
	Properties *JobHistoryDefinitionProperties `json:"properties,omitempty"`
}

// JobHistoryDefinitionProperties ...
type JobHistoryDefinitionProperties struct {
	// StartTime - READ-ONLY; Gets the start time for this job.
	StartTime *date.Time `json:"startTime,omitempty"`
	// EndTime - READ-ONLY; Gets the end time for this job.
	EndTime *date.Time `json:"endTime,omitempty"`
	// ExpectedExecutionTime - READ-ONLY; Gets the expected execution time for this job.
	ExpectedExecutionTime *date.Time `json:"expectedExecutionTime,omitempty"`
	// ActionName - READ-ONLY; Gets the job history action name. Possible values include: 'MainAction', 'ErrorAction'
	ActionName JobHistoryActionName `json:"actionName,omitempty"`
	// Status - READ-ONLY; Gets the job history status. Possible values include: 'Completed', 'Failed', 'Postponed'
	Status JobExecutionStatus `json:"status,omitempty"`
	// Message - READ-ONLY; Gets the message for the job history.
	Message *string `json:"message,omitempty"`
	// RetryCount - READ-ONLY; Gets the retry count for job.
	RetryCount *int32 `json:"retryCount,omitempty"`
	// RepeatCount - READ-ONLY; Gets the repeat count for the job.
	RepeatCount *int32 `json:"repeatCount,omitempty"`
}

// JobHistoryFilter ...
type JobHistoryFilter struct {
	// Status - Gets or sets the job execution status. Possible values include: 'Completed', 'Failed', 'Postponed'
	Status JobExecutionStatus `json:"status,omitempty"`
}

// JobHistoryListResult ...
type JobHistoryListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets the job histories under job.
	Value *[]JobHistoryDefinition `json:"value,omitempty"`
	// NextLink - Gets or sets the URL to get the next set of job histories.
	NextLink *string `json:"nextLink,omitempty"`
}

// JobHistoryListResultIterator provides access to a complete listing of JobHistoryDefinition values.
type JobHistoryListResultIterator struct {
	i    int
	page JobHistoryListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *JobHistoryListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobHistoryListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *JobHistoryListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter JobHistoryListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter JobHistoryListResultIterator) Response() JobHistoryListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter JobHistoryListResultIterator) Value() JobHistoryDefinition {
	if !iter.page.NotDone() {
		return JobHistoryDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the JobHistoryListResultIterator type.
func NewJobHistoryListResultIterator(page JobHistoryListResultPage) JobHistoryListResultIterator {
	return JobHistoryListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (jhlr JobHistoryListResult) IsEmpty() bool {
	return jhlr.Value == nil || len(*jhlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (jhlr JobHistoryListResult) hasNextLink() bool {
	return jhlr.NextLink != nil && len(*jhlr.NextLink) != 0
}

// jobHistoryListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (jhlr JobHistoryListResult) jobHistoryListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !jhlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(jhlr.NextLink)))
}

// JobHistoryListResultPage contains a page of JobHistoryDefinition values.
type JobHistoryListResultPage struct {
	fn   func(context.Context, JobHistoryListResult) (JobHistoryListResult, error)
	jhlr JobHistoryListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *JobHistoryListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobHistoryListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.jhlr)
		if err != nil {
			return err
		}
		page.jhlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *JobHistoryListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page JobHistoryListResultPage) NotDone() bool {
	return !page.jhlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page JobHistoryListResultPage) Response() JobHistoryListResult {
	return page.jhlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page JobHistoryListResultPage) Values() []JobHistoryDefinition {
	if page.jhlr.IsEmpty() {
		return nil
	}
	return *page.jhlr.Value
}

// Creates a new instance of the JobHistoryListResultPage type.
func NewJobHistoryListResultPage(cur JobHistoryListResult, getNextPage func(context.Context, JobHistoryListResult) (JobHistoryListResult, error)) JobHistoryListResultPage {
	return JobHistoryListResultPage{
		fn:   getNextPage,
		jhlr: cur,
	}
}

// JobListResult ...
type JobListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets all jobs under job collection.
	Value *[]JobDefinition `json:"value,omitempty"`
	// NextLink - Gets or sets the URL to get the next set of jobs.
	NextLink *string `json:"nextLink,omitempty"`
}

// JobListResultIterator provides access to a complete listing of JobDefinition values.
type JobListResultIterator struct {
	i    int
	page JobListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *JobListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *JobListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter JobListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter JobListResultIterator) Response() JobListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter JobListResultIterator) Value() JobDefinition {
	if !iter.page.NotDone() {
		return JobDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the JobListResultIterator type.
func NewJobListResultIterator(page JobListResultPage) JobListResultIterator {
	return JobListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (jlr JobListResult) IsEmpty() bool {
	return jlr.Value == nil || len(*jlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (jlr JobListResult) hasNextLink() bool {
	return jlr.NextLink != nil && len(*jlr.NextLink) != 0
}

// jobListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (jlr JobListResult) jobListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !jlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(jlr.NextLink)))
}

// JobListResultPage contains a page of JobDefinition values.
type JobListResultPage struct {
	fn  func(context.Context, JobListResult) (JobListResult, error)
	jlr JobListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *JobListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.jlr)
		if err != nil {
			return err
		}
		page.jlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *JobListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page JobListResultPage) NotDone() bool {
	return !page.jlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page JobListResultPage) Response() JobListResult {
	return page.jlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page JobListResultPage) Values() []JobDefinition {
	if page.jlr.IsEmpty() {
		return nil
	}
	return *page.jlr.Value
}

// Creates a new instance of the JobListResultPage type.
func NewJobListResultPage(cur JobListResult, getNextPage func(context.Context, JobListResult) (JobListResult, error)) JobListResultPage {
	return JobListResultPage{
		fn:  getNextPage,
		jlr: cur,
	}
}

// JobMaxRecurrence ...
type JobMaxRecurrence struct {
	// Frequency - Gets or sets the frequency of recurrence (second, minute, hour, day, week, month). Possible values include: 'Minute', 'Hour', 'Day', 'Week', 'Month'
	Frequency RecurrenceFrequency `json:"frequency,omitempty"`
	// Interval - Gets or sets the interval between retries.
	Interval *int32 `json:"interval,omitempty"`
}

// JobProperties ...
type JobProperties struct {
	// StartTime - Gets or sets the job start time.
	StartTime *date.Time `json:"startTime,omitempty"`
	// Action - Gets or sets the job action.
	Action *JobAction `json:"action,omitempty"`
	// Recurrence - Gets or sets the job recurrence.
	Recurrence *JobRecurrence `json:"recurrence,omitempty"`
	// State - Gets or set the job state. Possible values include: 'JobStateEnabled', 'JobStateDisabled', 'JobStateFaulted', 'JobStateCompleted'
	State JobState `json:"state,omitempty"`
	// Status - READ-ONLY; Gets the job status.
	Status *JobStatus `json:"status,omitempty"`
}

// MarshalJSON is the custom marshaler for JobProperties.
func (jp JobProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if jp.StartTime != nil {
		objectMap["startTime"] = jp.StartTime
	}
	if jp.Action != nil {
		objectMap["action"] = jp.Action
	}
	if jp.Recurrence != nil {
		objectMap["recurrence"] = jp.Recurrence
	}
	if jp.State != "" {
		objectMap["state"] = jp.State
	}
	return json.Marshal(objectMap)
}

// JobRecurrence ...
type JobRecurrence struct {
	// Frequency - Gets or sets the frequency of recurrence (second, minute, hour, day, week, month). Possible values include: 'Minute', 'Hour', 'Day', 'Week', 'Month'
	Frequency RecurrenceFrequency `json:"frequency,omitempty"`
	// Interval - Gets or sets the interval between retries.
	Interval *int32 `json:"interval,omitempty"`
	// Count - Gets or sets the maximum number of times that the job should run.
	Count *int32 `json:"count,omitempty"`
	// EndTime - Gets or sets the time at which the job will complete.
	EndTime  *date.Time             `json:"endTime,omitempty"`
	Schedule *JobRecurrenceSchedule `json:"schedule,omitempty"`
}

// JobRecurrenceSchedule ...
type JobRecurrenceSchedule struct {
	// WeekDays - Gets or sets the days of the week that the job should execute on.
	WeekDays *[]DayOfWeek `json:"weekDays,omitempty"`
	// Hours - Gets or sets the hours of the day that the job should execute at.
	Hours *[]int32 `json:"hours,omitempty"`
	// Minutes - Gets or sets the minutes of the hour that the job should execute at.
	Minutes *[]int32 `json:"minutes,omitempty"`
	// MonthDays - Gets or sets the days of the month that the job should execute on. Must be between 1 and 31.
	MonthDays *[]int32 `json:"monthDays,omitempty"`
	// MonthlyOccurrences - Gets or sets the occurrences of days within a month.
	MonthlyOccurrences *[]JobRecurrenceScheduleMonthlyOccurrence `json:"monthlyOccurrences,omitempty"`
}

// JobRecurrenceScheduleMonthlyOccurrence ...
type JobRecurrenceScheduleMonthlyOccurrence struct {
	// Day - Gets or sets the day. Must be one of monday, tuesday, wednesday, thursday, friday, saturday, sunday. Possible values include: 'JobScheduleDayMonday', 'JobScheduleDayTuesday', 'JobScheduleDayWednesday', 'JobScheduleDayThursday', 'JobScheduleDayFriday', 'JobScheduleDaySaturday', 'JobScheduleDaySunday'
	Day JobScheduleDay `json:"day,omitempty"`
	// Occurrence - Gets or sets the occurrence. Must be between -5 and 5.
	Occurrence *int32 `json:"Occurrence,omitempty"`
}

// JobStateFilter ...
type JobStateFilter struct {
	// State - Gets or sets the job state. Possible values include: 'JobStateEnabled', 'JobStateDisabled', 'JobStateFaulted', 'JobStateCompleted'
	State JobState `json:"state,omitempty"`
}

// JobStatus ...
type JobStatus struct {
	// ExecutionCount - READ-ONLY; Gets the number of times this job has executed.
	ExecutionCount *int32 `json:"executionCount,omitempty"`
	// FailureCount - READ-ONLY; Gets the number of times this job has failed.
	FailureCount *int32 `json:"failureCount,omitempty"`
	// FaultedCount - READ-ONLY; Gets the number of faulted occurrences (occurrences that were retried and failed as many times as the retry policy states).
	FaultedCount *int32 `json:"faultedCount,omitempty"`
	// LastExecutionTime - READ-ONLY; Gets the time the last occurrence executed in ISO-8601 format.  Could be empty if job has not run yet.
	LastExecutionTime *date.Time `json:"lastExecutionTime,omitempty"`
	// NextExecutionTime - READ-ONLY; Gets the time of the next occurrence in ISO-8601 format. Could be empty if the job is completed.
	NextExecutionTime *date.Time `json:"nextExecutionTime,omitempty"`
}

// OAuthAuthentication ...
type OAuthAuthentication struct {
	// Secret - Gets or sets the secret.
	Secret *string `json:"secret,omitempty"`
	// Tenant - Gets or sets the tenant.
	Tenant *string `json:"tenant,omitempty"`
	// Audience - Gets or sets the audience.
	Audience *string `json:"audience,omitempty"`
	// ClientID - Gets or sets the client identifier.
	ClientID *string `json:"clientId,omitempty"`
	// Type - Gets or sets the http authentication type. Possible values include: 'NotSpecified', 'ClientCertificate', 'ActiveDirectoryOAuth', 'Basic'
	Type HTTPAuthenticationType `json:"type,omitempty"`
}

// RetryPolicy ...
type RetryPolicy struct {
	// RetryType - Gets or sets the retry strategy to be used. Possible values include: 'None', 'Fixed'
	RetryType RetryType `json:"retryType,omitempty"`
	// RetryInterval - Gets or sets the retry interval between retries.
	RetryInterval *string `json:"retryInterval,omitempty"`
	// RetryCount - Gets or sets the number of times a retry should be attempted.
	RetryCount *int32 `json:"retryCount,omitempty"`
}

// ServiceBusAuthentication ...
type ServiceBusAuthentication struct {
	// SasKey - Gets or sets the SAS key.
	SasKey *string `json:"sasKey,omitempty"`
	// SasKeyName - Gets or sets the SAS key name.
	SasKeyName *string `json:"sasKeyName,omitempty"`
	// Type - Gets or sets the authentication type. Possible values include: 'ServiceBusAuthenticationTypeNotSpecified', 'ServiceBusAuthenticationTypeSharedAccessKey'
	Type ServiceBusAuthenticationType `json:"type,omitempty"`
}

// ServiceBusBrokeredMessageProperties ...
type ServiceBusBrokeredMessageProperties struct {
	// ContentType - Gets or sets the content type.
	ContentType *string `json:"contentType,omitempty"`
	// CorrelationID - Gets or sets the correlation id.
	CorrelationID *string `json:"correlationId,omitempty"`
	// ForcePersistence - Gets or sets the force persistence.
	ForcePersistence *bool `json:"forcePersistence,omitempty"`
	// Label - Gets or sets the label.
	Label *string `json:"label,omitempty"`
	// MessageID - Gets or sets the message id.
	MessageID *string `json:"messageId,omitempty"`
	// PartitionKey - Gets or sets the partition key.
	PartitionKey *string `json:"partitionKey,omitempty"`
	// ReplyTo - Gets or sets the reply to.
	ReplyTo *string `json:"replyTo,omitempty"`
	// ReplyToSessionID - Gets or sets the reply to session id.
	ReplyToSessionID *string `json:"replyToSessionId,omitempty"`
	// ScheduledEnqueueTimeUtc - Gets or sets the scheduled enqueue time UTC.
	ScheduledEnqueueTimeUtc *date.Time `json:"scheduledEnqueueTimeUtc,omitempty"`
	// SessionID - Gets or sets the session id.
	SessionID *string `json:"sessionId,omitempty"`
	// TimeToLive - Gets or sets the time to live.
	TimeToLive *date.Time `json:"timeToLive,omitempty"`
	// To - Gets or sets the to.
	To *string `json:"to,omitempty"`
	// ViaPartitionKey - Gets or sets the via partition key.
	ViaPartitionKey *string `json:"viaPartitionKey,omitempty"`
}

// ServiceBusMessage ...
type ServiceBusMessage struct {
	// Authentication - Gets or sets the authentication.
	Authentication *ServiceBusAuthentication `json:"authentication,omitempty"`
	// BrokeredMessageProperties - Gets or sets the brokered message properties.
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `json:"brokeredMessageProperties,omitempty"`
	// CustomMessageProperties - Gets or sets the custom message properties.
	CustomMessageProperties map[string]*string `json:"customMessageProperties"`
	// Message - Gets or sets the message.
	Message *string `json:"message,omitempty"`
	// Namespace - Gets or sets the namespace.
	Namespace *string `json:"namespace,omitempty"`
	// TransportType - Gets or sets the transport type. Possible values include: 'ServiceBusTransportTypeNotSpecified', 'ServiceBusTransportTypeNetMessaging', 'ServiceBusTransportTypeAMQP'
	TransportType ServiceBusTransportType `json:"transportType,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceBusMessage.
func (sbm ServiceBusMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sbm.Authentication != nil {
		objectMap["authentication"] = sbm.Authentication
	}
	if sbm.BrokeredMessageProperties != nil {
		objectMap["brokeredMessageProperties"] = sbm.BrokeredMessageProperties
	}
	if sbm.CustomMessageProperties != nil {
		objectMap["customMessageProperties"] = sbm.CustomMessageProperties
	}
	if sbm.Message != nil {
		objectMap["message"] = sbm.Message
	}
	if sbm.Namespace != nil {
		objectMap["namespace"] = sbm.Namespace
	}
	if sbm.TransportType != "" {
		objectMap["transportType"] = sbm.TransportType
	}
	return json.Marshal(objectMap)
}

// ServiceBusQueueMessage ...
type ServiceBusQueueMessage struct {
	// QueueName - Gets or sets the queue name.
	QueueName *string `json:"queueName,omitempty"`
	// Authentication - Gets or sets the authentication.
	Authentication *ServiceBusAuthentication `json:"authentication,omitempty"`
	// BrokeredMessageProperties - Gets or sets the brokered message properties.
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `json:"brokeredMessageProperties,omitempty"`
	// CustomMessageProperties - Gets or sets the custom message properties.
	CustomMessageProperties map[string]*string `json:"customMessageProperties"`
	// Message - Gets or sets the message.
	Message *string `json:"message,omitempty"`
	// Namespace - Gets or sets the namespace.
	Namespace *string `json:"namespace,omitempty"`
	// TransportType - Gets or sets the transport type. Possible values include: 'ServiceBusTransportTypeNotSpecified', 'ServiceBusTransportTypeNetMessaging', 'ServiceBusTransportTypeAMQP'
	TransportType ServiceBusTransportType `json:"transportType,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceBusQueueMessage.
func (sbqm ServiceBusQueueMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sbqm.QueueName != nil {
		objectMap["queueName"] = sbqm.QueueName
	}
	if sbqm.Authentication != nil {
		objectMap["authentication"] = sbqm.Authentication
	}
	if sbqm.BrokeredMessageProperties != nil {
		objectMap["brokeredMessageProperties"] = sbqm.BrokeredMessageProperties
	}
	if sbqm.CustomMessageProperties != nil {
		objectMap["customMessageProperties"] = sbqm.CustomMessageProperties
	}
	if sbqm.Message != nil {
		objectMap["message"] = sbqm.Message
	}
	if sbqm.Namespace != nil {
		objectMap["namespace"] = sbqm.Namespace
	}
	if sbqm.TransportType != "" {
		objectMap["transportType"] = sbqm.TransportType
	}
	return json.Marshal(objectMap)
}

// ServiceBusTopicMessage ...
type ServiceBusTopicMessage struct {
	// TopicPath - Gets or sets the topic path.
	TopicPath *string `json:"topicPath,omitempty"`
	// Authentication - Gets or sets the authentication.
	Authentication *ServiceBusAuthentication `json:"authentication,omitempty"`
	// BrokeredMessageProperties - Gets or sets the brokered message properties.
	BrokeredMessageProperties *ServiceBusBrokeredMessageProperties `json:"brokeredMessageProperties,omitempty"`
	// CustomMessageProperties - Gets or sets the custom message properties.
	CustomMessageProperties map[string]*string `json:"customMessageProperties"`
	// Message - Gets or sets the message.
	Message *string `json:"message,omitempty"`
	// Namespace - Gets or sets the namespace.
	Namespace *string `json:"namespace,omitempty"`
	// TransportType - Gets or sets the transport type. Possible values include: 'ServiceBusTransportTypeNotSpecified', 'ServiceBusTransportTypeNetMessaging', 'ServiceBusTransportTypeAMQP'
	TransportType ServiceBusTransportType `json:"transportType,omitempty"`
}

// MarshalJSON is the custom marshaler for ServiceBusTopicMessage.
func (sbtm ServiceBusTopicMessage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sbtm.TopicPath != nil {
		objectMap["topicPath"] = sbtm.TopicPath
	}
	if sbtm.Authentication != nil {
		objectMap["authentication"] = sbtm.Authentication
	}
	if sbtm.BrokeredMessageProperties != nil {
		objectMap["brokeredMessageProperties"] = sbtm.BrokeredMessageProperties
	}
	if sbtm.CustomMessageProperties != nil {
		objectMap["customMessageProperties"] = sbtm.CustomMessageProperties
	}
	if sbtm.Message != nil {
		objectMap["message"] = sbtm.Message
	}
	if sbtm.Namespace != nil {
		objectMap["namespace"] = sbtm.Namespace
	}
	if sbtm.TransportType != "" {
		objectMap["transportType"] = sbtm.TransportType
	}
	return json.Marshal(objectMap)
}

// Sku ...
type Sku struct {
	// Name - Gets or set the SKU. Possible values include: 'Standard', 'Free', 'Premium'
	Name SkuDefinition `json:"name,omitempty"`
}

// StorageQueueMessage ...
type StorageQueueMessage struct {
	// StorageAccount - Gets or sets the storage account name.
	StorageAccount *string `json:"storageAccount,omitempty"`
	// QueueName - Gets or sets the queue name.
	QueueName *string `json:"queueName,omitempty"`
	// SasToken - Gets or sets the SAS key.
	SasToken *string `json:"sasToken,omitempty"`
	// Message - Gets or sets the message.
	Message *string `json:"message,omitempty"`
}
