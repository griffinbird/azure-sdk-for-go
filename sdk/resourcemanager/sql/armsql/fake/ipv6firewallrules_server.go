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

// IPv6FirewallRulesServer is a fake server for instances of the armsql.IPv6FirewallRulesClient type.
type IPv6FirewallRulesServer struct {
	// CreateOrUpdate is the fake for method IPv6FirewallRulesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters armsql.IPv6FirewallRule, options *armsql.IPv6FirewallRulesClientCreateOrUpdateOptions) (resp azfake.Responder[armsql.IPv6FirewallRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method IPv6FirewallRulesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *armsql.IPv6FirewallRulesClientDeleteOptions) (resp azfake.Responder[armsql.IPv6FirewallRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IPv6FirewallRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *armsql.IPv6FirewallRulesClientGetOptions) (resp azfake.Responder[armsql.IPv6FirewallRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method IPv6FirewallRulesClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armsql.IPv6FirewallRulesClientListByServerOptions) (resp azfake.PagerResponder[armsql.IPv6FirewallRulesClientListByServerResponse])
}

// NewIPv6FirewallRulesServerTransport creates a new instance of IPv6FirewallRulesServerTransport with the provided implementation.
// The returned IPv6FirewallRulesServerTransport instance is connected to an instance of armsql.IPv6FirewallRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIPv6FirewallRulesServerTransport(srv *IPv6FirewallRulesServer) *IPv6FirewallRulesServerTransport {
	return &IPv6FirewallRulesServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armsql.IPv6FirewallRulesClientListByServerResponse]](),
	}
}

// IPv6FirewallRulesServerTransport connects instances of armsql.IPv6FirewallRulesClient to instances of IPv6FirewallRulesServer.
// Don't use this type directly, use NewIPv6FirewallRulesServerTransport instead.
type IPv6FirewallRulesServerTransport struct {
	srv                  *IPv6FirewallRulesServer
	newListByServerPager *tracker[azfake.PagerResponder[armsql.IPv6FirewallRulesClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for IPv6FirewallRulesServerTransport.
func (i *IPv6FirewallRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IPv6FirewallRulesClient.CreateOrUpdate":
		resp, err = i.dispatchCreateOrUpdate(req)
	case "IPv6FirewallRulesClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "IPv6FirewallRulesClient.Get":
		resp, err = i.dispatchGet(req)
	case "IPv6FirewallRulesClient.NewListByServerPager":
		resp, err = i.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IPv6FirewallRulesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipv6FirewallRules/(?P<firewallRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsql.IPv6FirewallRule](req)
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
	firewallRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, firewallRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IPv6FirewallRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IPv6FirewallRulesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipv6FirewallRules/(?P<firewallRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	firewallRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, serverNameParam, firewallRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IPv6FirewallRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipv6FirewallRules/(?P<firewallRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	firewallRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, firewallRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IPv6FirewallRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IPv6FirewallRulesServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := i.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipv6FirewallRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		resp := i.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		i.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armsql.IPv6FirewallRulesClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		i.newListByServerPager.remove(req)
	}
	return resp, nil
}
