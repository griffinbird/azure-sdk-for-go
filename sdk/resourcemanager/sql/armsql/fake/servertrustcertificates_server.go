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

// ServerTrustCertificatesServer is a fake server for instances of the armsql.ServerTrustCertificatesClient type.
type ServerTrustCertificatesServer struct {
	// BeginCreateOrUpdate is the fake for method ServerTrustCertificatesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, parameters armsql.ServerTrustCertificate, options *armsql.ServerTrustCertificatesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ServerTrustCertificatesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ServerTrustCertificatesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, options *armsql.ServerTrustCertificatesClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.ServerTrustCertificatesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServerTrustCertificatesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, certificateName string, options *armsql.ServerTrustCertificatesClientGetOptions) (resp azfake.Responder[armsql.ServerTrustCertificatesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByInstancePager is the fake for method ServerTrustCertificatesClient.NewListByInstancePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInstancePager func(resourceGroupName string, managedInstanceName string, options *armsql.ServerTrustCertificatesClientListByInstanceOptions) (resp azfake.PagerResponder[armsql.ServerTrustCertificatesClientListByInstanceResponse])
}

// NewServerTrustCertificatesServerTransport creates a new instance of ServerTrustCertificatesServerTransport with the provided implementation.
// The returned ServerTrustCertificatesServerTransport instance is connected to an instance of armsql.ServerTrustCertificatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTrustCertificatesServerTransport(srv *ServerTrustCertificatesServer) *ServerTrustCertificatesServerTransport {
	return &ServerTrustCertificatesServerTransport{
		srv:                    srv,
		beginCreateOrUpdate:    newTracker[azfake.PollerResponder[armsql.ServerTrustCertificatesClientCreateOrUpdateResponse]](),
		beginDelete:            newTracker[azfake.PollerResponder[armsql.ServerTrustCertificatesClientDeleteResponse]](),
		newListByInstancePager: newTracker[azfake.PagerResponder[armsql.ServerTrustCertificatesClientListByInstanceResponse]](),
	}
}

// ServerTrustCertificatesServerTransport connects instances of armsql.ServerTrustCertificatesClient to instances of ServerTrustCertificatesServer.
// Don't use this type directly, use NewServerTrustCertificatesServerTransport instead.
type ServerTrustCertificatesServerTransport struct {
	srv                    *ServerTrustCertificatesServer
	beginCreateOrUpdate    *tracker[azfake.PollerResponder[armsql.ServerTrustCertificatesClientCreateOrUpdateResponse]]
	beginDelete            *tracker[azfake.PollerResponder[armsql.ServerTrustCertificatesClientDeleteResponse]]
	newListByInstancePager *tracker[azfake.PagerResponder[armsql.ServerTrustCertificatesClientListByInstanceResponse]]
}

// Do implements the policy.Transporter interface for ServerTrustCertificatesServerTransport.
func (s *ServerTrustCertificatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServerTrustCertificatesClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "ServerTrustCertificatesClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "ServerTrustCertificatesClient.Get":
		resp, err = s.dispatchGet(req)
	case "ServerTrustCertificatesClient.NewListByInstancePager":
		resp, err = s.dispatchNewListByInstancePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTrustCertificatesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverTrustCertificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ServerTrustCertificate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, certificateNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *ServerTrustCertificatesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverTrustCertificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, managedInstanceNameParam, certificateNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *ServerTrustCertificatesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverTrustCertificates/(?P<certificateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	certificateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("certificateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, certificateNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerTrustCertificate, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTrustCertificatesServerTransport) dispatchNewListByInstancePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByInstancePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInstancePager not implemented")}
	}
	newListByInstancePager := s.newListByInstancePager.get(req)
	if newListByInstancePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverTrustCertificates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByInstancePager(resourceGroupNameParam, managedInstanceNameParam, nil)
		newListByInstancePager = &resp
		s.newListByInstancePager.add(req, newListByInstancePager)
		server.PagerResponderInjectNextLinks(newListByInstancePager, req, func(page *armsql.ServerTrustCertificatesClientListByInstanceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInstancePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByInstancePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInstancePager) {
		s.newListByInstancePager.remove(req)
	}
	return resp, nil
}
