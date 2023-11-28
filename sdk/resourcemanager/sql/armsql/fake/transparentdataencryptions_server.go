//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
)

// TransparentDataEncryptionsServer is a fake server for instances of the armsql.TransparentDataEncryptionsClient type.
type TransparentDataEncryptionsServer struct {
	// CreateOrUpdate is the fake for method TransparentDataEncryptionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, tdeName armsql.TransparentDataEncryptionName, parameters armsql.LogicalDatabaseTransparentDataEncryption, options *armsql.TransparentDataEncryptionsClientCreateOrUpdateOptions) (resp azfake.Responder[armsql.TransparentDataEncryptionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TransparentDataEncryptionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, tdeName armsql.TransparentDataEncryptionName, options *armsql.TransparentDataEncryptionsClientGetOptions) (resp azfake.Responder[armsql.TransparentDataEncryptionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabasePager is the fake for method TransparentDataEncryptionsClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.TransparentDataEncryptionsClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.TransparentDataEncryptionsClientListByDatabaseResponse])
}

// NewTransparentDataEncryptionsServerTransport creates a new instance of TransparentDataEncryptionsServerTransport with the provided implementation.
// The returned TransparentDataEncryptionsServerTransport instance is connected to an instance of armsql.TransparentDataEncryptionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTransparentDataEncryptionsServerTransport(srv *TransparentDataEncryptionsServer) *TransparentDataEncryptionsServerTransport {
	return &TransparentDataEncryptionsServerTransport{
		srv:                    srv,
		newListByDatabasePager: newTracker[azfake.PagerResponder[armsql.TransparentDataEncryptionsClientListByDatabaseResponse]](),
	}
}

// TransparentDataEncryptionsServerTransport connects instances of armsql.TransparentDataEncryptionsClient to instances of TransparentDataEncryptionsServer.
// Don't use this type directly, use NewTransparentDataEncryptionsServerTransport instead.
type TransparentDataEncryptionsServerTransport struct {
	srv                    *TransparentDataEncryptionsServer
	newListByDatabasePager *tracker[azfake.PagerResponder[armsql.TransparentDataEncryptionsClientListByDatabaseResponse]]
}

// Do implements the policy.Transporter interface for TransparentDataEncryptionsServerTransport.
func (t *TransparentDataEncryptionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TransparentDataEncryptionsClient.CreateOrUpdate":
		resp, err = t.dispatchCreateOrUpdate(req)
	case "TransparentDataEncryptionsClient.Get":
		resp, err = t.dispatchGet(req)
	case "TransparentDataEncryptionsClient.NewListByDatabasePager":
		resp, err = t.dispatchNewListByDatabasePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransparentDataEncryptionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transparentDataEncryption/(?P<tdeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsql.LogicalDatabaseTransparentDataEncryption](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	tdeNameParam, err := parseWithCast(matches[regex.SubexpIndex("tdeName")], func(v string) (armsql.TransparentDataEncryptionName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.TransparentDataEncryptionName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, tdeNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LogicalDatabaseTransparentDataEncryption, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TransparentDataEncryptionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transparentDataEncryption/(?P<tdeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	tdeNameParam, err := parseWithCast(matches[regex.SubexpIndex("tdeName")], func(v string) (armsql.TransparentDataEncryptionName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.TransparentDataEncryptionName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, tdeNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LogicalDatabaseTransparentDataEncryption, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TransparentDataEncryptionsServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := t.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transparentDataEncryption`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, nil)
		newListByDatabasePager = &resp
		t.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.TransparentDataEncryptionsClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		t.newListByDatabasePager.remove(req)
	}
	return resp, nil
}
