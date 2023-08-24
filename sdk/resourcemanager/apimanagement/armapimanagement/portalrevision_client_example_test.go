//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListPortalRevisions.json
func ExamplePortalRevisionClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPortalRevisionClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.PortalRevisionClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PortalRevisionCollection = armapimanagement.PortalRevisionCollection{
		// 	Value: []*armapimanagement.PortalRevisionContract{
		// 		{
		// 			Name: to.Ptr("20201112000000"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/portalRevisions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalRevisions/20201112000000"),
		// 			Properties: &armapimanagement.PortalRevisionContractProperties{
		// 				Description: to.Ptr("portal revision"),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:10:09.673Z"); return t}()),
		// 				IsCurrent: to.Ptr(false),
		// 				Status: to.Ptr(armapimanagement.PortalRevisionStatusCompleted),
		// 				UpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:12:41.46Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("20201112101010"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/portalRevisions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalRevisions/20201112101010"),
		// 			Properties: &armapimanagement.PortalRevisionContractProperties{
		// 				Description: to.Ptr("portal revision 1"),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:51:36.47Z"); return t}()),
		// 				IsCurrent: to.Ptr(true),
		// 				Status: to.Ptr(armapimanagement.PortalRevisionStatusCompleted),
		// 				UpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:52:00.097Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadPortalRevision.json
func ExamplePortalRevisionClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPortalRevisionClient().GetEntityTag(ctx, "rg1", "apimService1", "20201112101010", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetPortalRevision.json
func ExamplePortalRevisionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPortalRevisionClient().Get(ctx, "rg1", "apimService1", "20201112101010", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PortalRevisionContract = armapimanagement.PortalRevisionContract{
	// 	Name: to.Ptr("20201112101010"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/portalRevisions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/portalRevisions/20201112101010"),
	// 	Properties: &armapimanagement.PortalRevisionContractProperties{
	// 		Description: to.Ptr("portal revision 1"),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:51:36.47Z"); return t}()),
	// 		IsCurrent: to.Ptr(true),
	// 		Status: to.Ptr(armapimanagement.PortalRevisionStatusCompleted),
	// 		UpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-12T22:52:00.097Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreatePortalRevision.json
func ExamplePortalRevisionClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPortalRevisionClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "20201112101010", armapimanagement.PortalRevisionContract{
		Properties: &armapimanagement.PortalRevisionContractProperties{
			Description: to.Ptr("portal revision 1"),
			IsCurrent:   to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdatePortalRevision.json
func ExamplePortalRevisionClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPortalRevisionClient().BeginUpdate(ctx, "rg1", "apimService1", "20201112101010", "*", armapimanagement.PortalRevisionContract{
		Properties: &armapimanagement.PortalRevisionContractProperties{
			Description: to.Ptr("portal revision update"),
			IsCurrent:   to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PortalRevisionContract = armapimanagement.PortalRevisionContract{
	// 	Name: to.Ptr("20201112101010"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/portalRevisions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/namedValues/testprop2"),
	// 	Properties: &armapimanagement.PortalRevisionContractProperties{
	// 		Description: to.Ptr("portal revision update"),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-13T22:47:13.397Z"); return t}()),
	// 		IsCurrent: to.Ptr(true),
	// 		Status: to.Ptr(armapimanagement.PortalRevisionStatusCompleted),
	// 		UpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-13T23:29:25.34Z"); return t}()),
	// 	},
	// }
}
