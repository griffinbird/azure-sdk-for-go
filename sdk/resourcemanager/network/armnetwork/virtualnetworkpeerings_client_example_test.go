//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringDelete.json
func ExampleVirtualNetworkPeeringsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkPeeringsClient().BeginDelete(ctx, "peerTest", "vnet1", "peer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringGet.json
func ExampleVirtualNetworkPeeringsClient_Get_getPeering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkPeeringsClient().Get(ctx, "peerTest", "vnet1", "peer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkPeering = armnetwork.VirtualNetworkPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
	// 	Name: to.Ptr("peer"),
	// 	Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
	// 		AllowForwardedTraffic: to.Ptr(true),
	// 		AllowGatewayTransit: to.Ptr(false),
	// 		AllowVirtualNetworkAccess: to.Ptr(true),
	// 		PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
	// 		PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("12.0.0.0/8")},
	// 			},
	// 			RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
	// 				RegionalCommunity: to.Ptr("12076:50004"),
	// 				VirtualNetworkCommunity: to.Ptr("12076:20002"),
	// 			},
	// 			RemoteVirtualNetwork: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 			},
	// 			RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("12.0.0.0/8")},
	// 				},
	// 				UseRemoteGateways: to.Ptr(false),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringGetWithRemoteVirtualNetworkEncryption.json
func ExampleVirtualNetworkPeeringsClient_Get_getPeeringWithRemoteVirtualNetworkEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkPeeringsClient().Get(ctx, "peerTest", "vnet1", "peer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkPeering = armnetwork.VirtualNetworkPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
	// 	Name: to.Ptr("peer"),
	// 	Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
	// 		AllowForwardedTraffic: to.Ptr(true),
	// 		AllowGatewayTransit: to.Ptr(false),
	// 		AllowVirtualNetworkAccess: to.Ptr(true),
	// 		PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("12.0.0.0/8")},
	// 			},
	// 			RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
	// 				RegionalCommunity: to.Ptr("12076:50004"),
	// 				VirtualNetworkCommunity: to.Ptr("12076:20002"),
	// 			},
	// 			RemoteVirtualNetwork: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 			},
	// 			RemoteVirtualNetworkEncryption: &armnetwork.VirtualNetworkEncryption{
	// 				Enabled: to.Ptr(true),
	// 				Enforcement: to.Ptr(armnetwork.VirtualNetworkEncryptionEnforcementAllowUnencrypted),
	// 			},
	// 			UseRemoteGateways: to.Ptr(false),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringCreate.json
func ExampleVirtualNetworkPeeringsClient_BeginCreateOrUpdate_createPeering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkPeeringsClient().BeginCreateOrUpdate(ctx, "peerTest", "vnet1", "peer", armnetwork.VirtualNetworkPeering{
		Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
			AllowForwardedTraffic:     to.Ptr(true),
			AllowGatewayTransit:       to.Ptr(false),
			AllowVirtualNetworkAccess: to.Ptr(true),
			RemoteVirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
			},
			UseRemoteGateways: to.Ptr(false),
		},
	}, &armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions{SyncRemoteAddressSpace: nil})
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
	// res.VirtualNetworkPeering = armnetwork.VirtualNetworkPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
	// 	Name: to.Ptr("peer"),
	// 	Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
	// 		AllowForwardedTraffic: to.Ptr(true),
	// 		AllowGatewayTransit: to.Ptr(false),
	// 		AllowVirtualNetworkAccess: to.Ptr(true),
	// 		PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
	// 		PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("12.0.0.0/8")},
	// 			},
	// 			RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
	// 				RegionalCommunity: to.Ptr("12076:50004"),
	// 				VirtualNetworkCommunity: to.Ptr("12076:20002"),
	// 			},
	// 			RemoteVirtualNetwork: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 			},
	// 			RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("12.0.0.0/8")},
	// 				},
	// 				UseRemoteGateways: to.Ptr(false),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringCreateWithRemoteVirtualNetworkEncryption.json
func ExampleVirtualNetworkPeeringsClient_BeginCreateOrUpdate_createPeeringWithRemoteVirtualNetworkEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkPeeringsClient().BeginCreateOrUpdate(ctx, "peerTest", "vnet1", "peer", armnetwork.VirtualNetworkPeering{
		Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
			AllowForwardedTraffic:     to.Ptr(true),
			AllowGatewayTransit:       to.Ptr(false),
			AllowVirtualNetworkAccess: to.Ptr(true),
			RemoteVirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
			},
			UseRemoteGateways: to.Ptr(false),
		},
	}, &armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions{SyncRemoteAddressSpace: nil})
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
	// res.VirtualNetworkPeering = armnetwork.VirtualNetworkPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
	// 	Name: to.Ptr("peer"),
	// 	Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
	// 		AllowForwardedTraffic: to.Ptr(true),
	// 		AllowGatewayTransit: to.Ptr(false),
	// 		AllowVirtualNetworkAccess: to.Ptr(true),
	// 		PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("12.0.0.0/8")},
	// 			},
	// 			RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
	// 				RegionalCommunity: to.Ptr("12076:50004"),
	// 				VirtualNetworkCommunity: to.Ptr("12076:20002"),
	// 			},
	// 			RemoteVirtualNetwork: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 			},
	// 			RemoteVirtualNetworkEncryption: &armnetwork.VirtualNetworkEncryption{
	// 				Enabled: to.Ptr(true),
	// 				Enforcement: to.Ptr(armnetwork.VirtualNetworkEncryptionEnforcementAllowUnencrypted),
	// 			},
	// 			UseRemoteGateways: to.Ptr(false),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringSync.json
func ExampleVirtualNetworkPeeringsClient_BeginCreateOrUpdate_syncPeering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkPeeringsClient().BeginCreateOrUpdate(ctx, "peerTest", "vnet1", "peer", armnetwork.VirtualNetworkPeering{
		Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
			AllowForwardedTraffic:     to.Ptr(true),
			AllowGatewayTransit:       to.Ptr(false),
			AllowVirtualNetworkAccess: to.Ptr(true),
			RemoteVirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
			},
			UseRemoteGateways: to.Ptr(false),
		},
	}, &armnetwork.VirtualNetworkPeeringsClientBeginCreateOrUpdateOptions{SyncRemoteAddressSpace: to.Ptr(armnetwork.SyncRemoteAddressSpaceTrue)})
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
	// res.VirtualNetworkPeering = armnetwork.VirtualNetworkPeering{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
	// 	Name: to.Ptr("peer"),
	// 	Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
	// 		AllowForwardedTraffic: to.Ptr(true),
	// 		AllowGatewayTransit: to.Ptr(false),
	// 		AllowVirtualNetworkAccess: to.Ptr(true),
	// 		PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
	// 		PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteAddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("12.0.0.0/8")},
	// 			},
	// 			RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
	// 				RegionalCommunity: to.Ptr("12076:50004"),
	// 				VirtualNetworkCommunity: to.Ptr("12076:20002"),
	// 			},
	// 			RemoteVirtualNetwork: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
	// 			},
	// 			RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("12.0.0.0/8")},
	// 				},
	// 				UseRemoteGateways: to.Ptr(false),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringList.json
func ExampleVirtualNetworkPeeringsClient_NewListPager_listPeerings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkPeeringsClient().NewListPager("peerTest", "vnet1", nil)
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
		// page.VirtualNetworkPeeringListResult = armnetwork.VirtualNetworkPeeringListResult{
		// 	Value: []*armnetwork.VirtualNetworkPeering{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
		// 			Name: to.Ptr("peer"),
		// 			Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
		// 				AllowForwardedTraffic: to.Ptr(true),
		// 				AllowGatewayTransit: to.Ptr(false),
		// 				AllowVirtualNetworkAccess: to.Ptr(true),
		// 				PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
		// 				PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RemoteAddressSpace: &armnetwork.AddressSpace{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("12.0.0.0/8")},
		// 					},
		// 					RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
		// 						RegionalCommunity: to.Ptr("12076:50004"),
		// 						VirtualNetworkCommunity: to.Ptr("12076:20002"),
		// 					},
		// 					RemoteVirtualNetwork: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
		// 					},
		// 					RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("12.0.0.0/8")},
		// 						},
		// 						UseRemoteGateways: to.Ptr(false),
		// 					},
		// 				},
		// 				{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer2"),
		// 					Name: to.Ptr("peer"),
		// 					Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
		// 						AllowForwardedTraffic: to.Ptr(false),
		// 						AllowGatewayTransit: to.Ptr(false),
		// 						AllowVirtualNetworkAccess: to.Ptr(true),
		// 						PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
		// 						PeeringSyncLevel: to.Ptr(armnetwork.VirtualNetworkPeeringLevelFullyInSync),
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						RemoteAddressSpace: &armnetwork.AddressSpace{
		// 							AddressPrefixes: []*string{
		// 								to.Ptr("13.0.0.0/8")},
		// 							},
		// 							RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
		// 								RegionalCommunity: to.Ptr("12076:50004"),
		// 								VirtualNetworkCommunity: to.Ptr("12076:20003"),
		// 							},
		// 							RemoteVirtualNetwork: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet3"),
		// 							},
		// 							RemoteVirtualNetworkAddressSpace: &armnetwork.AddressSpace{
		// 								AddressPrefixes: []*string{
		// 									to.Ptr("13.0.0.0/8")},
		// 								},
		// 								UseRemoteGateways: to.Ptr(false),
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VirtualNetworkPeeringListWithRemoteVirtualNetworkEncryption.json
func ExampleVirtualNetworkPeeringsClient_NewListPager_listPeeringsWithRemoteVirtualNetworkEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkPeeringsClient().NewListPager("peerTest", "vnet1", nil)
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
		// page.VirtualNetworkPeeringListResult = armnetwork.VirtualNetworkPeeringListResult{
		// 	Value: []*armnetwork.VirtualNetworkPeering{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer"),
		// 			Name: to.Ptr("peer"),
		// 			Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
		// 				AllowForwardedTraffic: to.Ptr(true),
		// 				AllowGatewayTransit: to.Ptr(false),
		// 				AllowVirtualNetworkAccess: to.Ptr(true),
		// 				PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RemoteAddressSpace: &armnetwork.AddressSpace{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("12.0.0.0/8")},
		// 					},
		// 					RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
		// 						RegionalCommunity: to.Ptr("12076:50004"),
		// 						VirtualNetworkCommunity: to.Ptr("12076:20002"),
		// 					},
		// 					RemoteVirtualNetwork: &armnetwork.SubResource{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
		// 					},
		// 					RemoteVirtualNetworkEncryption: &armnetwork.VirtualNetworkEncryption{
		// 						Enabled: to.Ptr(true),
		// 						Enforcement: to.Ptr(armnetwork.VirtualNetworkEncryptionEnforcementAllowUnencrypted),
		// 					},
		// 					UseRemoteGateways: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet1/virtualNetworkPeerings/peer2"),
		// 				Name: to.Ptr("peer"),
		// 				Properties: &armnetwork.VirtualNetworkPeeringPropertiesFormat{
		// 					AllowForwardedTraffic: to.Ptr(false),
		// 					AllowGatewayTransit: to.Ptr(false),
		// 					AllowVirtualNetworkAccess: to.Ptr(true),
		// 					PeeringState: to.Ptr(armnetwork.VirtualNetworkPeeringStateInitiated),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RemoteAddressSpace: &armnetwork.AddressSpace{
		// 						AddressPrefixes: []*string{
		// 							to.Ptr("13.0.0.0/8")},
		// 						},
		// 						RemoteBgpCommunities: &armnetwork.VirtualNetworkBgpCommunities{
		// 							RegionalCommunity: to.Ptr("12076:50004"),
		// 							VirtualNetworkCommunity: to.Ptr("12076:20003"),
		// 						},
		// 						RemoteVirtualNetwork: &armnetwork.SubResource{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet3"),
		// 						},
		// 						RemoteVirtualNetworkEncryption: &armnetwork.VirtualNetworkEncryption{
		// 							Enabled: to.Ptr(true),
		// 							Enforcement: to.Ptr(armnetwork.VirtualNetworkEncryptionEnforcementAllowUnencrypted),
		// 						},
		// 						UseRemoteGateways: to.Ptr(false),
		// 					},
		// 			}},
		// 		}
	}
}
