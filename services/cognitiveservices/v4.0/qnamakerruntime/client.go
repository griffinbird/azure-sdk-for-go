// Package qnamakerruntime implements the Azure ARM Qnamakerruntime service API version 4.0.
//
// An API for QnAMaker runtime
package qnamakerruntime

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// BaseClient is the base client for Qnamakerruntime.
type BaseClient struct {
	autorest.Client
	RuntimeEndpoint string
}

// New creates an instance of the BaseClient client.
func New(runtimeEndpoint string) BaseClient {
	return NewWithoutDefaults(runtimeEndpoint)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(runtimeEndpoint string) BaseClient {
	return BaseClient{
		Client:          autorest.NewClientWithUserAgent(UserAgent()),
		RuntimeEndpoint: runtimeEndpoint,
	}
}
