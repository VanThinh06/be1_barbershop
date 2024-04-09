// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_service_categories_create.proto

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

type CreateServiceCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsGlobal bool   `protobuf:"varint,2,opt,name=is_global,json=isGlobal,proto3" json:"is_global,omitempty"`
}

func (x *CreateServiceCategoriesRequest) Reset() {
	*x = CreateServiceCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_categories_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceCategoriesRequest) ProtoMessage() {}

func (x *CreateServiceCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_categories_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceCategoriesRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_categories_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServiceCategoriesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServiceCategoriesRequest) GetIsGlobal() bool {
	if x != nil {
		return x.IsGlobal
	}
	return false
}

type CreateServiceCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceCategories *ServiceCategories `protobuf:"bytes,1,opt,name=service_categories,json=serviceCategories,proto3" json:"service_categories,omitempty"`
}

func (x *CreateServiceCategoriesResponse) Reset() {
	*x = CreateServiceCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_categories_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceCategoriesResponse) ProtoMessage() {}

func (x *CreateServiceCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_categories_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceCategoriesResponse.ProtoReflect.Descriptor instead.
func (*CreateServiceCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_categories_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServiceCategoriesResponse) GetServiceCategories() *ServiceCategories {
	if x != nil {
		return x.ServiceCategories
	}
	return nil
}

var File_barber_rpc_service_categories_create_proto protoreflect.FileDescriptor

var file_barber_rpc_service_categories_create_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x51, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x22, 0x67, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x1a, 0x5a,
	0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_barber_rpc_service_categories_create_proto_rawDescOnce sync.Once
	file_barber_rpc_service_categories_create_proto_rawDescData = file_barber_rpc_service_categories_create_proto_rawDesc
)

func file_barber_rpc_service_categories_create_proto_rawDescGZIP() []byte {
	file_barber_rpc_service_categories_create_proto_rawDescOnce.Do(func() {
		file_barber_rpc_service_categories_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_service_categories_create_proto_rawDescData)
	})
	return file_barber_rpc_service_categories_create_proto_rawDescData
}

var file_barber_rpc_service_categories_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_service_categories_create_proto_goTypes = []interface{}{
	(*CreateServiceCategoriesRequest)(nil),  // 0: pb.CreateServiceCategoriesRequest
	(*CreateServiceCategoriesResponse)(nil), // 1: pb.CreateServiceCategoriesResponse
	(*ServiceCategories)(nil),               // 2: pb.ServiceCategories
}
var file_barber_rpc_service_categories_create_proto_depIdxs = []int32{
	2, // 0: pb.CreateServiceCategoriesResponse.service_categories:type_name -> pb.ServiceCategories
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_service_categories_create_proto_init() }
func file_barber_rpc_service_categories_create_proto_init() {
	if File_barber_rpc_service_categories_create_proto != nil {
		return
	}
	file_barber_service_categories_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_service_categories_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceCategoriesRequest); i {
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
		file_barber_rpc_service_categories_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceCategoriesResponse); i {
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
			RawDescriptor: file_barber_rpc_service_categories_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_service_categories_create_proto_goTypes,
		DependencyIndexes: file_barber_rpc_service_categories_create_proto_depIdxs,
		MessageInfos:      file_barber_rpc_service_categories_create_proto_msgTypes,
	}.Build()
	File_barber_rpc_service_categories_create_proto = out.File
	file_barber_rpc_service_categories_create_proto_rawDesc = nil
	file_barber_rpc_service_categories_create_proto_goTypes = nil
	file_barber_rpc_service_categories_create_proto_depIdxs = nil
}
