// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_barber_managers_create.proto

package barber

import (
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

type CreateBarberManagersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberId  string `protobuf:"bytes,1,opt,name=barber_id,json=barberId,proto3" json:"barber_id,omitempty"`
	ManagerId string `protobuf:"bytes,2,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
}

func (x *CreateBarberManagersRequest) Reset() {
	*x = CreateBarberManagersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barber_managers_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberManagersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberManagersRequest) ProtoMessage() {}

func (x *CreateBarberManagersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barber_managers_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberManagersRequest.ProtoReflect.Descriptor instead.
func (*CreateBarberManagersRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barber_managers_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBarberManagersRequest) GetBarberId() string {
	if x != nil {
		return x.BarberId
	}
	return ""
}

func (x *CreateBarberManagersRequest) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

type CreateBarberManagersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberManager *BarberManagers `protobuf:"bytes,1,opt,name=barber_manager,json=barberManager,proto3" json:"barber_manager,omitempty"`
}

func (x *CreateBarberManagersResponse) Reset() {
	*x = CreateBarberManagersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barber_managers_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberManagersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberManagersResponse) ProtoMessage() {}

func (x *CreateBarberManagersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barber_managers_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberManagersResponse.ProtoReflect.Descriptor instead.
func (*CreateBarberManagersResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barber_managers_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBarberManagersResponse) GetBarberManager() *BarberManagers {
	if x != nil {
		return x.BarberManager
	}
	return nil
}

var File_barber_rpc_barber_managers_create_proto protoreflect.FileDescriptor

var file_barber_rpc_barber_managers_create_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x1b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x52, 0x0d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barber_managers_create_proto_rawDescOnce sync.Once
	file_barber_rpc_barber_managers_create_proto_rawDescData = file_barber_rpc_barber_managers_create_proto_rawDesc
)

func file_barber_rpc_barber_managers_create_proto_rawDescGZIP() []byte {
	file_barber_rpc_barber_managers_create_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barber_managers_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barber_managers_create_proto_rawDescData)
	})
	return file_barber_rpc_barber_managers_create_proto_rawDescData
}

var file_barber_rpc_barber_managers_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barber_managers_create_proto_goTypes = []interface{}{
	(*CreateBarberManagersRequest)(nil),  // 0: pb.CreateBarberManagersRequest
	(*CreateBarberManagersResponse)(nil), // 1: pb.CreateBarberManagersResponse
	(*BarberManagers)(nil),               // 2: pb.BarberManagers
}
var file_barber_rpc_barber_managers_create_proto_depIdxs = []int32{
	2, // 0: pb.CreateBarberManagersResponse.barber_manager:type_name -> pb.BarberManagers
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barber_managers_create_proto_init() }
func file_barber_rpc_barber_managers_create_proto_init() {
	if File_barber_rpc_barber_managers_create_proto != nil {
		return
	}
	file_barber_barber_managers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barber_managers_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberManagersRequest); i {
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
		file_barber_rpc_barber_managers_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberManagersResponse); i {
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
			RawDescriptor: file_barber_rpc_barber_managers_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barber_managers_create_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barber_managers_create_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barber_managers_create_proto_msgTypes,
	}.Build()
	File_barber_rpc_barber_managers_create_proto = out.File
	file_barber_rpc_barber_managers_create_proto_rawDesc = nil
	file_barber_rpc_barber_managers_create_proto_goTypes = nil
	file_barber_rpc_barber_managers_create_proto_depIdxs = nil
}