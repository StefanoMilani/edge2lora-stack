// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: lorawan-stack/api/gateway_services.proto

package ttnpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PullGatewayConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayIds *GatewayIdentifiers    `protobuf:"bytes,1,opt,name=gateway_ids,json=gatewayIds,proto3" json:"gateway_ids,omitempty"`
	FieldMask  *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
}

func (x *PullGatewayConfigurationRequest) Reset() {
	*x = PullGatewayConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lorawan_stack_api_gateway_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullGatewayConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullGatewayConfigurationRequest) ProtoMessage() {}

func (x *PullGatewayConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lorawan_stack_api_gateway_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullGatewayConfigurationRequest.ProtoReflect.Descriptor instead.
func (*PullGatewayConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_lorawan_stack_api_gateway_services_proto_rawDescGZIP(), []int{0}
}

func (x *PullGatewayConfigurationRequest) GetGatewayIds() *GatewayIdentifiers {
	if x != nil {
		return x.GatewayIds
	}
	return nil
}

func (x *PullGatewayConfigurationRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

var File_lorawan_stack_api_gateway_services_proto protoreflect.FileDescriptor

var file_lorawan_stack_api_gateway_services_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01,
	0x0a, 0x1f, 0x50, 0x75, 0x6c, 0x6c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x43, 0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x49, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x32, 0xd9, 0x08, 0x0a, 0x0f, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0xd3, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x24, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22,
	0x89, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x82, 0x01, 0x22, 0x2f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x3a, 0x01, 0x2a, 0x5a, 0x4c, 0x22,
	0x47, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x21, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x2a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x45,
	0x55, 0x49, 0x12, 0x32, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x45, 0x55, 0x49, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72,
	0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0xd8, 0x01, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x73, 0x22, 0x90, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x89, 0x01, 0x12, 0x09, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x5a, 0x31, 0x12, 0x2f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x5a, 0x49, 0x12, 0x47, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x73, 0x12, 0x76, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x24, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61,
	0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x2d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x1a, 0x22, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x22,
	0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x22, 0x1e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x69, 0x0a, 0x05, 0x50, 0x75, 0x72, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a,
	0x1c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x32, 0xcd, 0x0b,
	0x0a, 0x0d, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x6f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x22, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x1a, 0x16, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x12, 0x1d, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x4b, 0x65,
	0x79, 0x12, 0x2a, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x2b, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x29, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c,
	0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79,
	0x73, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64,
	0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x8a, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x97, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d,
	0x1a, 0x38, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x61,
	0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x97, 0x01,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x2a,
	0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x50, 0x49,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x50, 0x49, 0x4b,
	0x65, 0x79, 0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x2a, 0x38, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69,
	0x64, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79,
	0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xbb, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x2e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xcf, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xc8, 0x01, 0x5a, 0x56, 0x12,
	0x54, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x5a, 0x6e, 0x12, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x95, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x1a, 0x30, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x73,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x9d, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x2f, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77,
	0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x2f, 0x7b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x32, 0x76, 0x0a,
	0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x5f, 0x0a, 0x11, 0x50, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x74, 0x74, 0x6e, 0x2e,
	0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x74, 0x6e,
	0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f,
	0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x74, 0x74, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lorawan_stack_api_gateway_services_proto_rawDescOnce sync.Once
	file_lorawan_stack_api_gateway_services_proto_rawDescData = file_lorawan_stack_api_gateway_services_proto_rawDesc
)

func file_lorawan_stack_api_gateway_services_proto_rawDescGZIP() []byte {
	file_lorawan_stack_api_gateway_services_proto_rawDescOnce.Do(func() {
		file_lorawan_stack_api_gateway_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_lorawan_stack_api_gateway_services_proto_rawDescData)
	})
	return file_lorawan_stack_api_gateway_services_proto_rawDescData
}

var file_lorawan_stack_api_gateway_services_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lorawan_stack_api_gateway_services_proto_goTypes = []interface{}{
	(*PullGatewayConfigurationRequest)(nil),    // 0: ttn.lorawan.v3.PullGatewayConfigurationRequest
	(*GatewayIdentifiers)(nil),                 // 1: ttn.lorawan.v3.GatewayIdentifiers
	(*fieldmaskpb.FieldMask)(nil),              // 2: google.protobuf.FieldMask
	(*CreateGatewayRequest)(nil),               // 3: ttn.lorawan.v3.CreateGatewayRequest
	(*GetGatewayRequest)(nil),                  // 4: ttn.lorawan.v3.GetGatewayRequest
	(*GetGatewayIdentifiersForEUIRequest)(nil), // 5: ttn.lorawan.v3.GetGatewayIdentifiersForEUIRequest
	(*ListGatewaysRequest)(nil),                // 6: ttn.lorawan.v3.ListGatewaysRequest
	(*UpdateGatewayRequest)(nil),               // 7: ttn.lorawan.v3.UpdateGatewayRequest
	(*CreateGatewayAPIKeyRequest)(nil),         // 8: ttn.lorawan.v3.CreateGatewayAPIKeyRequest
	(*ListGatewayAPIKeysRequest)(nil),          // 9: ttn.lorawan.v3.ListGatewayAPIKeysRequest
	(*GetGatewayAPIKeyRequest)(nil),            // 10: ttn.lorawan.v3.GetGatewayAPIKeyRequest
	(*UpdateGatewayAPIKeyRequest)(nil),         // 11: ttn.lorawan.v3.UpdateGatewayAPIKeyRequest
	(*DeleteGatewayAPIKeyRequest)(nil),         // 12: ttn.lorawan.v3.DeleteGatewayAPIKeyRequest
	(*GetGatewayCollaboratorRequest)(nil),      // 13: ttn.lorawan.v3.GetGatewayCollaboratorRequest
	(*SetGatewayCollaboratorRequest)(nil),      // 14: ttn.lorawan.v3.SetGatewayCollaboratorRequest
	(*ListGatewayCollaboratorsRequest)(nil),    // 15: ttn.lorawan.v3.ListGatewayCollaboratorsRequest
	(*Gateway)(nil),                            // 16: ttn.lorawan.v3.Gateway
	(*Gateways)(nil),                           // 17: ttn.lorawan.v3.Gateways
	(*emptypb.Empty)(nil),                      // 18: google.protobuf.Empty
	(*Rights)(nil),                             // 19: ttn.lorawan.v3.Rights
	(*APIKey)(nil),                             // 20: ttn.lorawan.v3.APIKey
	(*APIKeys)(nil),                            // 21: ttn.lorawan.v3.APIKeys
	(*GetCollaboratorResponse)(nil),            // 22: ttn.lorawan.v3.GetCollaboratorResponse
	(*Collaborators)(nil),                      // 23: ttn.lorawan.v3.Collaborators
}
var file_lorawan_stack_api_gateway_services_proto_depIdxs = []int32{
	1,  // 0: ttn.lorawan.v3.PullGatewayConfigurationRequest.gateway_ids:type_name -> ttn.lorawan.v3.GatewayIdentifiers
	2,  // 1: ttn.lorawan.v3.PullGatewayConfigurationRequest.field_mask:type_name -> google.protobuf.FieldMask
	3,  // 2: ttn.lorawan.v3.GatewayRegistry.Create:input_type -> ttn.lorawan.v3.CreateGatewayRequest
	4,  // 3: ttn.lorawan.v3.GatewayRegistry.Get:input_type -> ttn.lorawan.v3.GetGatewayRequest
	5,  // 4: ttn.lorawan.v3.GatewayRegistry.GetIdentifiersForEUI:input_type -> ttn.lorawan.v3.GetGatewayIdentifiersForEUIRequest
	6,  // 5: ttn.lorawan.v3.GatewayRegistry.List:input_type -> ttn.lorawan.v3.ListGatewaysRequest
	7,  // 6: ttn.lorawan.v3.GatewayRegistry.Update:input_type -> ttn.lorawan.v3.UpdateGatewayRequest
	1,  // 7: ttn.lorawan.v3.GatewayRegistry.Delete:input_type -> ttn.lorawan.v3.GatewayIdentifiers
	1,  // 8: ttn.lorawan.v3.GatewayRegistry.Restore:input_type -> ttn.lorawan.v3.GatewayIdentifiers
	1,  // 9: ttn.lorawan.v3.GatewayRegistry.Purge:input_type -> ttn.lorawan.v3.GatewayIdentifiers
	1,  // 10: ttn.lorawan.v3.GatewayAccess.ListRights:input_type -> ttn.lorawan.v3.GatewayIdentifiers
	8,  // 11: ttn.lorawan.v3.GatewayAccess.CreateAPIKey:input_type -> ttn.lorawan.v3.CreateGatewayAPIKeyRequest
	9,  // 12: ttn.lorawan.v3.GatewayAccess.ListAPIKeys:input_type -> ttn.lorawan.v3.ListGatewayAPIKeysRequest
	10, // 13: ttn.lorawan.v3.GatewayAccess.GetAPIKey:input_type -> ttn.lorawan.v3.GetGatewayAPIKeyRequest
	11, // 14: ttn.lorawan.v3.GatewayAccess.UpdateAPIKey:input_type -> ttn.lorawan.v3.UpdateGatewayAPIKeyRequest
	12, // 15: ttn.lorawan.v3.GatewayAccess.DeleteAPIKey:input_type -> ttn.lorawan.v3.DeleteGatewayAPIKeyRequest
	13, // 16: ttn.lorawan.v3.GatewayAccess.GetCollaborator:input_type -> ttn.lorawan.v3.GetGatewayCollaboratorRequest
	14, // 17: ttn.lorawan.v3.GatewayAccess.SetCollaborator:input_type -> ttn.lorawan.v3.SetGatewayCollaboratorRequest
	15, // 18: ttn.lorawan.v3.GatewayAccess.ListCollaborators:input_type -> ttn.lorawan.v3.ListGatewayCollaboratorsRequest
	0,  // 19: ttn.lorawan.v3.GatewayConfigurator.PullConfiguration:input_type -> ttn.lorawan.v3.PullGatewayConfigurationRequest
	16, // 20: ttn.lorawan.v3.GatewayRegistry.Create:output_type -> ttn.lorawan.v3.Gateway
	16, // 21: ttn.lorawan.v3.GatewayRegistry.Get:output_type -> ttn.lorawan.v3.Gateway
	1,  // 22: ttn.lorawan.v3.GatewayRegistry.GetIdentifiersForEUI:output_type -> ttn.lorawan.v3.GatewayIdentifiers
	17, // 23: ttn.lorawan.v3.GatewayRegistry.List:output_type -> ttn.lorawan.v3.Gateways
	16, // 24: ttn.lorawan.v3.GatewayRegistry.Update:output_type -> ttn.lorawan.v3.Gateway
	18, // 25: ttn.lorawan.v3.GatewayRegistry.Delete:output_type -> google.protobuf.Empty
	18, // 26: ttn.lorawan.v3.GatewayRegistry.Restore:output_type -> google.protobuf.Empty
	18, // 27: ttn.lorawan.v3.GatewayRegistry.Purge:output_type -> google.protobuf.Empty
	19, // 28: ttn.lorawan.v3.GatewayAccess.ListRights:output_type -> ttn.lorawan.v3.Rights
	20, // 29: ttn.lorawan.v3.GatewayAccess.CreateAPIKey:output_type -> ttn.lorawan.v3.APIKey
	21, // 30: ttn.lorawan.v3.GatewayAccess.ListAPIKeys:output_type -> ttn.lorawan.v3.APIKeys
	20, // 31: ttn.lorawan.v3.GatewayAccess.GetAPIKey:output_type -> ttn.lorawan.v3.APIKey
	20, // 32: ttn.lorawan.v3.GatewayAccess.UpdateAPIKey:output_type -> ttn.lorawan.v3.APIKey
	20, // 33: ttn.lorawan.v3.GatewayAccess.DeleteAPIKey:output_type -> ttn.lorawan.v3.APIKey
	22, // 34: ttn.lorawan.v3.GatewayAccess.GetCollaborator:output_type -> ttn.lorawan.v3.GetCollaboratorResponse
	18, // 35: ttn.lorawan.v3.GatewayAccess.SetCollaborator:output_type -> google.protobuf.Empty
	23, // 36: ttn.lorawan.v3.GatewayAccess.ListCollaborators:output_type -> ttn.lorawan.v3.Collaborators
	16, // 37: ttn.lorawan.v3.GatewayConfigurator.PullConfiguration:output_type -> ttn.lorawan.v3.Gateway
	20, // [20:38] is the sub-list for method output_type
	2,  // [2:20] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_lorawan_stack_api_gateway_services_proto_init() }
func file_lorawan_stack_api_gateway_services_proto_init() {
	if File_lorawan_stack_api_gateway_services_proto != nil {
		return
	}
	file_lorawan_stack_api_gateway_proto_init()
	file_lorawan_stack_api_identifiers_proto_init()
	file_lorawan_stack_api_rights_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lorawan_stack_api_gateway_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullGatewayConfigurationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lorawan_stack_api_gateway_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_lorawan_stack_api_gateway_services_proto_goTypes,
		DependencyIndexes: file_lorawan_stack_api_gateway_services_proto_depIdxs,
		MessageInfos:      file_lorawan_stack_api_gateway_services_proto_msgTypes,
	}.Build()
	File_lorawan_stack_api_gateway_services_proto = out.File
	file_lorawan_stack_api_gateway_services_proto_rawDesc = nil
	file_lorawan_stack_api_gateway_services_proto_goTypes = nil
	file_lorawan_stack_api_gateway_services_proto_depIdxs = nil
}
