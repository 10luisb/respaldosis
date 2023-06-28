// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package baremetalsolution aliases all exported identifiers in package
// "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb".
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package baremetalsolution

import (
	src "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
const (
	Instance_DELETED                                                  = src.Instance_DELETED
	Instance_PROVISIONING                                             = src.Instance_PROVISIONING
	Instance_RUNNING                                                  = src.Instance_RUNNING
	Instance_STATE_UNSPECIFIED                                        = src.Instance_STATE_UNSPECIFIED
	Lun_CREATING                                                      = src.Lun_CREATING
	Lun_DELETING                                                      = src.Lun_DELETING
	Lun_HDD                                                           = src.Lun_HDD
	Lun_LINUX                                                         = src.Lun_LINUX
	Lun_MULTIPROTOCOL_TYPE_UNSPECIFIED                                = src.Lun_MULTIPROTOCOL_TYPE_UNSPECIFIED
	Lun_READY                                                         = src.Lun_READY
	Lun_SSD                                                           = src.Lun_SSD
	Lun_STATE_UNSPECIFIED                                             = src.Lun_STATE_UNSPECIFIED
	Lun_STORAGE_TYPE_UNSPECIFIED                                      = src.Lun_STORAGE_TYPE_UNSPECIFIED
	Lun_UPDATING                                                      = src.Lun_UPDATING
	Network_CLIENT                                                    = src.Network_CLIENT
	Network_PRIVATE                                                   = src.Network_PRIVATE
	Network_PROVISIONED                                               = src.Network_PROVISIONED
	Network_PROVISIONING                                              = src.Network_PROVISIONING
	Network_STATE_UNSPECIFIED                                         = src.Network_STATE_UNSPECIFIED
	Network_TYPE_UNSPECIFIED                                          = src.Network_TYPE_UNSPECIFIED
	NfsShare_MOUNT_PERMISSIONS_UNSPECIFIED                            = src.NfsShare_MOUNT_PERMISSIONS_UNSPECIFIED
	NfsShare_PROVISIONED                                              = src.NfsShare_PROVISIONED
	NfsShare_READ                                                     = src.NfsShare_READ
	NfsShare_READ_WRITE                                               = src.NfsShare_READ_WRITE
	NfsShare_STATE_UNSPECIFIED                                        = src.NfsShare_STATE_UNSPECIFIED
	ServerNetworkTemplate_LogicalInterface_BOND                       = src.ServerNetworkTemplate_LogicalInterface_BOND
	ServerNetworkTemplate_LogicalInterface_INTERFACE_TYPE_UNSPECIFIED = src.ServerNetworkTemplate_LogicalInterface_INTERFACE_TYPE_UNSPECIFIED
	ServerNetworkTemplate_LogicalInterface_NIC                        = src.ServerNetworkTemplate_LogicalInterface_NIC
	VRF_PROVISIONED                                                   = src.VRF_PROVISIONED
	VRF_PROVISIONING                                                  = src.VRF_PROVISIONING
	VRF_STATE_UNSPECIFIED                                             = src.VRF_STATE_UNSPECIFIED
	Volume_CREATING                                                   = src.Volume_CREATING
	Volume_DELETING                                                   = src.Volume_DELETING
	Volume_DISABLED                                                   = src.Volume_DISABLED
	Volume_HDD                                                        = src.Volume_HDD
	Volume_NEWEST_FIRST                                               = src.Volume_NEWEST_FIRST
	Volume_OLDEST_FIRST                                               = src.Volume_OLDEST_FIRST
	Volume_READY                                                      = src.Volume_READY
	Volume_SNAPSHOT_AUTO_DELETE_BEHAVIOR_UNSPECIFIED                  = src.Volume_SNAPSHOT_AUTO_DELETE_BEHAVIOR_UNSPECIFIED
	Volume_SSD                                                        = src.Volume_SSD
	Volume_STATE_UNSPECIFIED                                          = src.Volume_STATE_UNSPECIFIED
	Volume_STORAGE_TYPE_UNSPECIFIED                                   = src.Volume_STORAGE_TYPE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
var (
	File_google_cloud_baremetalsolution_v2_baremetalsolution_proto = src.File_google_cloud_baremetalsolution_v2_baremetalsolution_proto
	File_google_cloud_baremetalsolution_v2_instance_proto          = src.File_google_cloud_baremetalsolution_v2_instance_proto
	File_google_cloud_baremetalsolution_v2_lun_proto               = src.File_google_cloud_baremetalsolution_v2_lun_proto
	File_google_cloud_baremetalsolution_v2_network_proto           = src.File_google_cloud_baremetalsolution_v2_network_proto
	File_google_cloud_baremetalsolution_v2_nfs_share_proto         = src.File_google_cloud_baremetalsolution_v2_nfs_share_proto
	File_google_cloud_baremetalsolution_v2_volume_proto            = src.File_google_cloud_baremetalsolution_v2_volume_proto
	Instance_State_name                                            = src.Instance_State_name
	Instance_State_value                                           = src.Instance_State_value
	Lun_MultiprotocolType_name                                     = src.Lun_MultiprotocolType_name
	Lun_MultiprotocolType_value                                    = src.Lun_MultiprotocolType_value
	Lun_State_name                                                 = src.Lun_State_name
	Lun_State_value                                                = src.Lun_State_value
	Lun_StorageType_name                                           = src.Lun_StorageType_name
	Lun_StorageType_value                                          = src.Lun_StorageType_value
	Network_State_name                                             = src.Network_State_name
	Network_State_value                                            = src.Network_State_value
	Network_Type_name                                              = src.Network_Type_name
	Network_Type_value                                             = src.Network_Type_value
	NfsShare_MountPermissions_name                                 = src.NfsShare_MountPermissions_name
	NfsShare_MountPermissions_value                                = src.NfsShare_MountPermissions_value
	NfsShare_State_name                                            = src.NfsShare_State_name
	NfsShare_State_value                                           = src.NfsShare_State_value
	ServerNetworkTemplate_LogicalInterface_InterfaceType_name      = src.ServerNetworkTemplate_LogicalInterface_InterfaceType_name
	ServerNetworkTemplate_LogicalInterface_InterfaceType_value     = src.ServerNetworkTemplate_LogicalInterface_InterfaceType_value
	VRF_State_name                                                 = src.VRF_State_name
	VRF_State_value                                                = src.VRF_State_value
	Volume_SnapshotAutoDeleteBehavior_name                         = src.Volume_SnapshotAutoDeleteBehavior_name
	Volume_SnapshotAutoDeleteBehavior_value                        = src.Volume_SnapshotAutoDeleteBehavior_value
	Volume_State_name                                              = src.Volume_State_name
	Volume_State_value                                             = src.Volume_State_value
	Volume_StorageType_name                                        = src.Volume_StorageType_name
	Volume_StorageType_value                                       = src.Volume_StorageType_value
)

// BareMetalSolutionClient is the client API for BareMetalSolution service.
// For semantics around ctx use and closing/ending streaming RPCs, please refer
// to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type BareMetalSolutionClient = src.BareMetalSolutionClient

// BareMetalSolutionServer is the server API for BareMetalSolution service.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type BareMetalSolutionServer = src.BareMetalSolutionServer

// Message for detach specific LUN from an Instance.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type DetachLunRequest = src.DetachLunRequest

// Message for requesting server information.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type GetInstanceRequest = src.GetInstanceRequest

// Message for requesting storage lun information.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type GetLunRequest = src.GetLunRequest

// Message for requesting network information.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type GetNetworkRequest = src.GetNetworkRequest

// Message for requesting NFS share information.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type GetNfsShareRequest = src.GetNfsShareRequest

// Message for requesting storage volume information.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type GetVolumeRequest = src.GetVolumeRequest

// A server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Instance = src.Instance

// The possible states for this server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Instance_State = src.Instance_State

// Message for requesting the list of servers.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListInstancesRequest = src.ListInstancesRequest

// Response message for the list of servers.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListInstancesResponse = src.ListInstancesResponse

// Message for requesting a list of storage volume luns.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListLunsRequest = src.ListLunsRequest

// Response message containing the list of storage volume luns.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListLunsResponse = src.ListLunsResponse

// Request to get networks with IPs.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListNetworkUsageRequest = src.ListNetworkUsageRequest

// Response with Networks with IPs
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListNetworkUsageResponse = src.ListNetworkUsageResponse

// Message for requesting a list of networks.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListNetworksRequest = src.ListNetworksRequest

// Response message containing the list of networks.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListNetworksResponse = src.ListNetworksResponse

// Message for requesting a list of NFS shares.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListNfsSharesRequest = src.ListNfsSharesRequest

// Response message containing the list of NFS shares.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListNfsSharesResponse = src.ListNfsSharesResponse

// Message for requesting a list of storage volumes.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListVolumesRequest = src.ListVolumesRequest

// Response message containing the list of storage volumes.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ListVolumesResponse = src.ListVolumesResponse

// Each logical interface represents a logical abstraction of the underlying
// physical interface (for eg. bond, nic) of the instance. Each logical
// interface can effectively map to multiple network-IP pairs and still be
// mapped to one underlying physical interface.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type LogicalInterface = src.LogicalInterface

// Each logical network interface is effectively a network and IP pair.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type LogicalInterface_LogicalNetworkInterface = src.LogicalInterface_LogicalNetworkInterface

// A storage volume logical unit number (LUN).
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Lun = src.Lun

// Display the operating systems present for the LUN multiprotocol type.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Lun_MultiprotocolType = src.Lun_MultiprotocolType

// The possible states for the LUN.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Lun_State = src.Lun_State

// The storage types for a LUN.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Lun_StorageType = src.Lun_StorageType

// A Network.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Network = src.Network

// A reservation of one or more addresses in a network.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type NetworkAddressReservation = src.NetworkAddressReservation

// Network with all used IP addresses.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type NetworkUsage = src.NetworkUsage

// The possible states for this Network.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Network_State = src.Network_State

// Network type.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Network_Type = src.Network_Type

// An NFS share.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type NfsShare = src.NfsShare

// Represents an 'access point' for the share.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type NfsShare_AllowedClient = src.NfsShare_AllowedClient

// The possible mount permissions.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type NfsShare_MountPermissions = src.NfsShare_MountPermissions

// The possible states for this NFS share.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type NfsShare_State = src.NfsShare_State

// Represents the metadata from a long-running operation.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type OperationMetadata = src.OperationMetadata

// Message requesting to reset a server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ResetInstanceRequest = src.ResetInstanceRequest

// Response message from resetting a server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ResetInstanceResponse = src.ResetInstanceResponse

// Request for emergency resize Volume.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ResizeVolumeRequest = src.ResizeVolumeRequest

// Network template.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ServerNetworkTemplate = src.ServerNetworkTemplate

// Logical interface.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ServerNetworkTemplate_LogicalInterface = src.ServerNetworkTemplate_LogicalInterface

// Interface type.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type ServerNetworkTemplate_LogicalInterface_InterfaceType = src.ServerNetworkTemplate_LogicalInterface_InterfaceType

// Message requesting to start a server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type StartInstanceRequest = src.StartInstanceRequest

// Response message from starting a server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type StartInstanceResponse = src.StartInstanceResponse

// Message requesting to stop a server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type StopInstanceRequest = src.StopInstanceRequest

// Response message from stopping a server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type StopInstanceResponse = src.StopInstanceResponse

// UnimplementedBareMetalSolutionServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type UnimplementedBareMetalSolutionServer = src.UnimplementedBareMetalSolutionServer

// Message requesting to updating a server.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type UpdateInstanceRequest = src.UpdateInstanceRequest

// Message requesting to updating a network.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type UpdateNetworkRequest = src.UpdateNetworkRequest

// Message requesting to updating a NFS share.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type UpdateNfsShareRequest = src.UpdateNfsShareRequest

// Message for updating a volume.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type UpdateVolumeRequest = src.UpdateVolumeRequest

// A network VRF.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type VRF = src.VRF

// QOS policy parameters.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type VRF_QosPolicy = src.VRF_QosPolicy

// The possible states for this VRF.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type VRF_State = src.VRF_State

// VLAN attachment details.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type VRF_VlanAttachment = src.VRF_VlanAttachment

// A storage volume.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Volume = src.Volume

// The kinds of auto delete behavior to use when snapshot reserved space is
// full.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Volume_SnapshotAutoDeleteBehavior = src.Volume_SnapshotAutoDeleteBehavior

// Details about snapshot space reservation and usage on the storage volume.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Volume_SnapshotReservationDetail = src.Volume_SnapshotReservationDetail

// The possible states for a storage volume.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Volume_State = src.Volume_State

// The storage type for a volume.
//
// Deprecated: Please use types in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
type Volume_StorageType = src.Volume_StorageType

// Deprecated: Please use funcs in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
func NewBareMetalSolutionClient(cc grpc.ClientConnInterface) BareMetalSolutionClient {
	return src.NewBareMetalSolutionClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb
func RegisterBareMetalSolutionServer(s *grpc.Server, srv BareMetalSolutionServer) {
	src.RegisterBareMetalSolutionServer(s, srv)
}