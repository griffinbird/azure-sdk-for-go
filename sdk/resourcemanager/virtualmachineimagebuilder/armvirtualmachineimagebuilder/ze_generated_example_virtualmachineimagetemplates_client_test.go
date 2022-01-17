//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder"
)

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/ListImageTemplates.json
func ExampleVirtualMachineImageTemplatesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/ListImageTemplatesByRg.json
func ExampleVirtualMachineImageTemplatesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/CreateImageTemplateLinux.json
func ExampleVirtualMachineImageTemplatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		armvirtualmachineimagebuilder.ImageTemplate{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"imagetemplate_tag1": to.StringPtr("IT_T1"),
				"imagetemplate_tag2": to.StringPtr("IT_T2"),
			},
			Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
				Type: armvirtualmachineimagebuilder.ResourceIdentityTypeUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armvirtualmachineimagebuilder.ComponentsVrq145SchemasImagetemplateidentityPropertiesUserassignedidentitiesAdditionalproperties{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1": {},
				},
			},
			Properties: &armvirtualmachineimagebuilder.ImageTemplateProperties{
				Customize: []armvirtualmachineimagebuilder.ImageTemplateCustomizerClassification{
					&armvirtualmachineimagebuilder.ImageTemplateShellCustomizer{
						Name:      to.StringPtr("<name>"),
						Type:      to.StringPtr("<type>"),
						ScriptURI: to.StringPtr("<script-uri>"),
					}},
				Distribute: []armvirtualmachineimagebuilder.ImageTemplateDistributorClassification{
					&armvirtualmachineimagebuilder.ImageTemplateManagedImageDistributor{
						Type: to.StringPtr("<type>"),
						ArtifactTags: map[string]*string{
							"tagName": to.StringPtr("value"),
						},
						RunOutputName: to.StringPtr("<run-output-name>"),
						ImageID:       to.StringPtr("<image-id>"),
						Location:      to.StringPtr("<location>"),
					}},
				Source: &armvirtualmachineimagebuilder.ImageTemplateManagedImageSource{
					Type:    to.StringPtr("<type>"),
					ImageID: to.StringPtr("<image-id>"),
				},
				VMProfile: &armvirtualmachineimagebuilder.ImageTemplateVMProfile{
					OSDiskSizeGB: to.Int32Ptr(64),
					VMSize:       to.StringPtr("<vmsize>"),
					VnetConfig: &armvirtualmachineimagebuilder.VirtualNetworkConfig{
						SubnetID: to.StringPtr("<subnet-id>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineImageTemplatesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/UpdateImageTemplateToRemoveIdentities.json
func ExampleVirtualMachineImageTemplatesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		armvirtualmachineimagebuilder.ImageTemplateUpdateParameters{
			Identity: &armvirtualmachineimagebuilder.ImageTemplateIdentity{
				Type: armvirtualmachineimagebuilder.ResourceIdentityTypeNone.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineImageTemplatesClientUpdateResult)
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/GetImageTemplate.json
func ExampleVirtualMachineImageTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineImageTemplatesClientGetResult)
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/DeleteImageTemplate.json
func ExampleVirtualMachineImageTemplatesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/RunImageTemplate.json
func ExampleVirtualMachineImageTemplatesClient_BeginRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRun(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/CancelImageBuild.json
func ExampleVirtualMachineImageTemplatesClient_BeginCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCancel(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/ListRunOutputs.json
func ExampleVirtualMachineImageTemplatesClient_ListRunOutputs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	pager := client.ListRunOutputs("<resource-group-name>",
		"<image-template-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/GetRunOutput.json
func ExampleVirtualMachineImageTemplatesClient_GetRunOutput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient("<subscription-id>", cred, nil)
	res, err := client.GetRunOutput(ctx,
		"<resource-group-name>",
		"<image-template-name>",
		"<run-output-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineImageTemplatesClientGetRunOutputResult)
}