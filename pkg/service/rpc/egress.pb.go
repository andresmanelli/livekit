// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: pkg/service/rpc/egress.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListActiveEgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListActiveEgressRequest) Reset() {
	*x = ListActiveEgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_rpc_egress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveEgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveEgressRequest) ProtoMessage() {}

func (x *ListActiveEgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_rpc_egress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveEgressRequest.ProtoReflect.Descriptor instead.
func (*ListActiveEgressRequest) Descriptor() ([]byte, []int) {
	return file_pkg_service_rpc_egress_proto_rawDescGZIP(), []int{0}
}

type ListActiveEgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EgressIds []string `protobuf:"bytes,1,rep,name=egress_ids,json=egressIds,proto3" json:"egress_ids,omitempty"`
}

func (x *ListActiveEgressResponse) Reset() {
	*x = ListActiveEgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_service_rpc_egress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveEgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveEgressResponse) ProtoMessage() {}

func (x *ListActiveEgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_service_rpc_egress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveEgressResponse.ProtoReflect.Descriptor instead.
func (*ListActiveEgressResponse) Descriptor() ([]byte, []int) {
	return file_pkg_service_rpc_egress_proto_rawDescGZIP(), []int{1}
}

func (x *ListActiveEgressResponse) GetEgressIds() []string {
	if x != nil {
		return x.EgressIds
	}
	return nil
}

var File_pkg_service_rpc_egress_proto protoreflect.FileDescriptor

var file_pkg_service_rpc_egress_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x39, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x73, 0x32, 0xb2, 0x01, 0x0a, 0x0e, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x47, 0x0a,
	0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x06,
	0xb2, 0x89, 0x01, 0x02, 0x20, 0x01, 0x12, 0x57, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x08, 0x01, 0x32,
	0xa1, 0x01, 0x0a, 0x0d, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x18, 0x01, 0x12, 0x45, 0x0a, 0x0a,
	0x53, 0x74, 0x6f, 0x70, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x06, 0xb2, 0x89, 0x01,
	0x02, 0x18, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_service_rpc_egress_proto_rawDescOnce sync.Once
	file_pkg_service_rpc_egress_proto_rawDescData = file_pkg_service_rpc_egress_proto_rawDesc
)

func file_pkg_service_rpc_egress_proto_rawDescGZIP() []byte {
	file_pkg_service_rpc_egress_proto_rawDescOnce.Do(func() {
		file_pkg_service_rpc_egress_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_service_rpc_egress_proto_rawDescData)
	})
	return file_pkg_service_rpc_egress_proto_rawDescData
}

var file_pkg_service_rpc_egress_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_service_rpc_egress_proto_goTypes = []interface{}{
	(*ListActiveEgressRequest)(nil),     // 0: rpc.ListActiveEgressRequest
	(*ListActiveEgressResponse)(nil),    // 1: rpc.ListActiveEgressResponse
	(*livekit.StartEgressRequest)(nil),  // 2: livekit.StartEgressRequest
	(*livekit.UpdateStreamRequest)(nil), // 3: livekit.UpdateStreamRequest
	(*livekit.StopEgressRequest)(nil),   // 4: livekit.StopEgressRequest
	(*livekit.EgressInfo)(nil),          // 5: livekit.EgressInfo
}
var file_pkg_service_rpc_egress_proto_depIdxs = []int32{
	2, // 0: rpc.EgressInternal.StartEgress:input_type -> livekit.StartEgressRequest
	0, // 1: rpc.EgressInternal.ListActiveEgress:input_type -> rpc.ListActiveEgressRequest
	3, // 2: rpc.EgressHandler.UpdateStream:input_type -> livekit.UpdateStreamRequest
	4, // 3: rpc.EgressHandler.StopEgress:input_type -> livekit.StopEgressRequest
	5, // 4: rpc.EgressInternal.StartEgress:output_type -> livekit.EgressInfo
	1, // 5: rpc.EgressInternal.ListActiveEgress:output_type -> rpc.ListActiveEgressResponse
	5, // 6: rpc.EgressHandler.UpdateStream:output_type -> livekit.EgressInfo
	5, // 7: rpc.EgressHandler.StopEgress:output_type -> livekit.EgressInfo
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_service_rpc_egress_proto_init() }
func file_pkg_service_rpc_egress_proto_init() {
	if File_pkg_service_rpc_egress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_service_rpc_egress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveEgressRequest); i {
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
		file_pkg_service_rpc_egress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveEgressResponse); i {
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
			RawDescriptor: file_pkg_service_rpc_egress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_service_rpc_egress_proto_goTypes,
		DependencyIndexes: file_pkg_service_rpc_egress_proto_depIdxs,
		MessageInfos:      file_pkg_service_rpc_egress_proto_msgTypes,
	}.Build()
	File_pkg_service_rpc_egress_proto = out.File
	file_pkg_service_rpc_egress_proto_rawDesc = nil
	file_pkg_service_rpc_egress_proto_goTypes = nil
	file_pkg_service_rpc_egress_proto_depIdxs = nil
}
