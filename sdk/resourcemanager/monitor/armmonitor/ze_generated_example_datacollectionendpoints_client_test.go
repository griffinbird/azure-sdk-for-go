//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsListByResourceGroup.json
func ExampleDataCollectionEndpointsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DataCollectionEndpointResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsListBySubscription.json
func ExampleDataCollectionEndpointsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DataCollectionEndpointResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsGet.json
func ExampleDataCollectionEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<data-collection-endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataCollectionEndpointResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsCreate.json
func ExampleDataCollectionEndpointsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<data-collection-endpoint-name>",
		&armmonitor.DataCollectionEndpointsCreateOptions{Body: &armmonitor.DataCollectionEndpointResource{
			Location: to.StringPtr("<location>"),
			Properties: &armmonitor.DataCollectionEndpointResourceProperties{
				DataCollectionEndpoint: armmonitor.DataCollectionEndpoint{
					NetworkACLs: &armmonitor.DataCollectionEndpointNetworkACLs{
						NetworkRuleSet: armmonitor.NetworkRuleSet{
							PublicNetworkAccess: armmonitor.KnownPublicNetworkAccessOptionsEnabled.ToPtr(),
						},
					},
				},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataCollectionEndpointResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsUpdate.json
func ExampleDataCollectionEndpointsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<data-collection-endpoint-name>",
		&armmonitor.DataCollectionEndpointsUpdateOptions{Body: &armmonitor.ResourceForUpdate{
			Tags: map[string]*string{
				"tag1": to.StringPtr("A"),
				"tag2": to.StringPtr("B"),
				"tag3": to.StringPtr("C"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataCollectionEndpointResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionEndpointsDelete.json
func ExampleDataCollectionEndpointsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionEndpointsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<data-collection-endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}