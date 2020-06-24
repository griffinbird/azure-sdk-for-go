// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// TagsOperations contains the methods for the Tags group.
type TagsOperations interface {
	// CreateOrUpdate - The tag name can have a maximum of 512 characters and is case insensitive. Tag names created by Azure have prefixes of microsoft, azure, or windows. You cannot create tags with one of these prefixes.
	CreateOrUpdate(ctx context.Context, tagName string) (*TagDetailsResponse, error)
	// CreateOrUpdateValue - Creates a tag value. The name of the tag must already exist.
	CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string) (*TagValueResponse, error)
	// Delete - You must remove all values from a resource tag before you can delete it.
	Delete(ctx context.Context, tagName string) (*http.Response, error)
	// DeleteValue - Deletes a tag value.
	DeleteValue(ctx context.Context, tagName string, tagValue string) (*http.Response, error)
	// List - Gets the names and values of all resource tags that are defined in a subscription.
	List() (TagsListResultPager, error)
}

// tagsOperations implements the TagsOperations interface.
type tagsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - The tag name can have a maximum of 512 characters and is case insensitive. Tag names created by Azure have prefixes of microsoft, azure, or windows. You cannot create tags with one of these prefixes.
func (client *tagsOperations) CreateOrUpdate(ctx context.Context, tagName string) (*TagDetailsResponse, error) {
	req, err := client.createOrUpdateCreateRequest(tagName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *tagsOperations) createOrUpdateCreateRequest(tagName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-05-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *tagsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*TagDetailsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := TagDetailsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TagDetails)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *tagsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// CreateOrUpdateValue - Creates a tag value. The name of the tag must already exist.
func (client *tagsOperations) CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string) (*TagValueResponse, error) {
	req, err := client.createOrUpdateValueCreateRequest(tagName, tagValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateValueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createOrUpdateValueCreateRequest creates the CreateOrUpdateValue request.
func (client *tagsOperations) createOrUpdateValueCreateRequest(tagName string, tagValue string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-05-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, nil
}

// createOrUpdateValueHandleResponse handles the CreateOrUpdateValue response.
func (client *tagsOperations) createOrUpdateValueHandleResponse(resp *azcore.Response) (*TagValueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateValueHandleError(resp)
	}
	result := TagValueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TagValue)
}

// createOrUpdateValueHandleError handles the CreateOrUpdateValue error response.
func (client *tagsOperations) createOrUpdateValueHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Delete - You must remove all values from a resource tag before you can delete it.
func (client *tagsOperations) Delete(ctx context.Context, tagName string) (*http.Response, error) {
	req, err := client.deleteCreateRequest(tagName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// deleteCreateRequest creates the Delete request.
func (client *tagsOperations) deleteCreateRequest(tagName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-05-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *tagsOperations) deleteHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteHandleError handles the Delete error response.
func (client *tagsOperations) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// DeleteValue - Deletes a tag value.
func (client *tagsOperations) DeleteValue(ctx context.Context, tagName string, tagValue string) (*http.Response, error) {
	req, err := client.deleteValueCreateRequest(tagName, tagValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteValueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// deleteValueCreateRequest creates the DeleteValue request.
func (client *tagsOperations) deleteValueCreateRequest(tagName string, tagValue string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-05-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteValueHandleResponse handles the DeleteValue response.
func (client *tagsOperations) deleteValueHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteValueHandleError(resp)
	}
	return resp.Response, nil
}

// deleteValueHandleError handles the DeleteValue error response.
func (client *tagsOperations) deleteValueHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// List - Gets the names and values of all resource tags that are defined in a subscription.
func (client *tagsOperations) List() (TagsListResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &tagsListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *TagsListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.TagsListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.TagsListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *tagsOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-05-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *tagsOperations) listHandleResponse(resp *azcore.Response) (*TagsListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := TagsListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.TagsListResult)
}

// listHandleError handles the List error response.
func (client *tagsOperations) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}
