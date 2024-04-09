// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_barbershops_list.proto

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

type ListBarberShopsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBarberShopsRequest) Reset() {
	*x = ListBarberShopsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBarberShopsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBarberShopsRequest) ProtoMessage() {}

func (x *ListBarberShopsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBarberShopsRequest.ProtoReflect.Descriptor instead.
func (*ListBarberShopsRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_list_proto_rawDescGZIP(), []int{0}
}

type ListBarberShopsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShops []*BarberShopInList `protobuf:"bytes,1,rep,name=barber_shops,json=barberShops,proto3" json:"barber_shops,omitempty"`
}

func (x *ListBarberShopsResponse) Reset() {
	*x = ListBarberShopsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBarberShopsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBarberShopsResponse) ProtoMessage() {}

func (x *ListBarberShopsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBarberShopsResponse.ProtoReflect.Descriptor instead.
func (*ListBarberShopsResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_list_proto_rawDescGZIP(), []int{1}
}

func (x *ListBarberShopsResponse) GetBarberShops() []*BarberShopInList {
	if x != nil {
		return x.BarberShops
	}
	return nil
}

var File_barber_rpc_barbershops_list_proto protoreflect.FileDescriptor

var file_barber_rpc_barbershops_list_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x17,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73,
	0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbershops_list_proto_rawDescOnce sync.Once
	file_barber_rpc_barbershops_list_proto_rawDescData = file_barber_rpc_barbershops_list_proto_rawDesc
)

func file_barber_rpc_barbershops_list_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbershops_list_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbershops_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbershops_list_proto_rawDescData)
	})
	return file_barber_rpc_barbershops_list_proto_rawDescData
}

var file_barber_rpc_barbershops_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbershops_list_proto_goTypes = []interface{}{
	(*ListBarberShopsRequest)(nil),  // 0: pb.ListBarberShopsRequest
	(*ListBarberShopsResponse)(nil), // 1: pb.ListBarberShopsResponse
	(*BarberShopInList)(nil),        // 2: pb.BarberShopInList
}
var file_barber_rpc_barbershops_list_proto_depIdxs = []int32{
	2, // 0: pb.ListBarberShopsResponse.barber_shops:type_name -> pb.BarberShopInList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbershops_list_proto_init() }
func file_barber_rpc_barbershops_list_proto_init() {
	if File_barber_rpc_barbershops_list_proto != nil {
		return
	}
	file_barber_barber_shops_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbershops_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBarberShopsRequest); i {
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
		file_barber_rpc_barbershops_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBarberShopsResponse); i {
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
			RawDescriptor: file_barber_rpc_barbershops_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbershops_list_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbershops_list_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbershops_list_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbershops_list_proto = out.File
	file_barber_rpc_barbershops_list_proto_rawDesc = nil
	file_barber_rpc_barbershops_list_proto_goTypes = nil
	file_barber_rpc_barbershops_list_proto_depIdxs = nil
}
