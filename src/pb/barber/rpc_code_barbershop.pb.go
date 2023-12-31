// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_code_barbershop.proto

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

type CodeBarberShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role int32 `protobuf:"varint,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CodeBarberShopRequest) Reset() {
	*x = CodeBarberShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_code_barbershop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeBarberShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeBarberShopRequest) ProtoMessage() {}

func (x *CodeBarberShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_code_barbershop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeBarberShopRequest.ProtoReflect.Descriptor instead.
func (*CodeBarberShopRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_code_barbershop_proto_rawDescGZIP(), []int{0}
}

func (x *CodeBarberShopRequest) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type CodeBarberShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CodeBarberShopResponse) Reset() {
	*x = CodeBarberShopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_code_barbershop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeBarberShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeBarberShopResponse) ProtoMessage() {}

func (x *CodeBarberShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_code_barbershop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeBarberShopResponse.ProtoReflect.Descriptor instead.
func (*CodeBarberShopResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_code_barbershop_proto_rawDescGZIP(), []int{1}
}

func (x *CodeBarberShopResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_barber_rpc_code_barbershop_proto protoreflect.FileDescriptor

var file_barber_rpc_code_barbershop_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2b, 0x0a, 0x15, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_code_barbershop_proto_rawDescOnce sync.Once
	file_barber_rpc_code_barbershop_proto_rawDescData = file_barber_rpc_code_barbershop_proto_rawDesc
)

func file_barber_rpc_code_barbershop_proto_rawDescGZIP() []byte {
	file_barber_rpc_code_barbershop_proto_rawDescOnce.Do(func() {
		file_barber_rpc_code_barbershop_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_code_barbershop_proto_rawDescData)
	})
	return file_barber_rpc_code_barbershop_proto_rawDescData
}

var file_barber_rpc_code_barbershop_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_code_barbershop_proto_goTypes = []interface{}{
	(*CodeBarberShopRequest)(nil),  // 0: pb.CodeBarberShopRequest
	(*CodeBarberShopResponse)(nil), // 1: pb.CodeBarberShopResponse
}
var file_barber_rpc_code_barbershop_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_barber_rpc_code_barbershop_proto_init() }
func file_barber_rpc_code_barbershop_proto_init() {
	if File_barber_rpc_code_barbershop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_code_barbershop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeBarberShopRequest); i {
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
		file_barber_rpc_code_barbershop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeBarberShopResponse); i {
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
			RawDescriptor: file_barber_rpc_code_barbershop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_code_barbershop_proto_goTypes,
		DependencyIndexes: file_barber_rpc_code_barbershop_proto_depIdxs,
		MessageInfos:      file_barber_rpc_code_barbershop_proto_msgTypes,
	}.Build()
	File_barber_rpc_code_barbershop_proto = out.File
	file_barber_rpc_code_barbershop_proto_rawDesc = nil
	file_barber_rpc_code_barbershop_proto_goTypes = nil
	file_barber_rpc_code_barbershop_proto_depIdxs = nil
}
