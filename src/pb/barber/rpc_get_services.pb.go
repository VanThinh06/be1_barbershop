// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_get_services.proto

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

type GetCategoryAndServiceDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	ShopId     string `protobuf:"bytes,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GetCategoryAndServiceDetailsRequest) Reset() {
	*x = GetCategoryAndServiceDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_get_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryAndServiceDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryAndServiceDetailsRequest) ProtoMessage() {}

func (x *GetCategoryAndServiceDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_get_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryAndServiceDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryAndServiceDetailsRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_get_services_proto_rawDescGZIP(), []int{0}
}

func (x *GetCategoryAndServiceDetailsRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *GetCategoryAndServiceDetailsRequest) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

type GetServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	ShopId     string `protobuf:"bytes,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GetServicesRequest) Reset() {
	*x = GetServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_get_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesRequest) ProtoMessage() {}

func (x *GetServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_get_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesRequest.ProtoReflect.Descriptor instead.
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_get_services_proto_rawDescGZIP(), []int{1}
}

func (x *GetServicesRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *GetServicesRequest) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

type GetServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *GetServicesResponse) Reset() {
	*x = GetServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_get_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesResponse) ProtoMessage() {}

func (x *GetServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_get_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesResponse.ProtoReflect.Descriptor instead.
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_get_services_proto_rawDescGZIP(), []int{2}
}

func (x *GetServicesResponse) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_barber_rpc_get_services_proto protoreflect.FileDescriptor

var file_barber_rpc_get_services_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x14, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x23, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_get_services_proto_rawDescOnce sync.Once
	file_barber_rpc_get_services_proto_rawDescData = file_barber_rpc_get_services_proto_rawDesc
)

func file_barber_rpc_get_services_proto_rawDescGZIP() []byte {
	file_barber_rpc_get_services_proto_rawDescOnce.Do(func() {
		file_barber_rpc_get_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_get_services_proto_rawDescData)
	})
	return file_barber_rpc_get_services_proto_rawDescData
}

var file_barber_rpc_get_services_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_barber_rpc_get_services_proto_goTypes = []interface{}{
	(*GetCategoryAndServiceDetailsRequest)(nil), // 0: pb.GetCategoryAndServiceDetailsRequest
	(*GetServicesRequest)(nil),                  // 1: pb.GetServicesRequest
	(*GetServicesResponse)(nil),                 // 2: pb.GetServicesResponse
	(*Service)(nil),                             // 3: pb.Service
}
var file_barber_rpc_get_services_proto_depIdxs = []int32{
	3, // 0: pb.GetServicesResponse.services:type_name -> pb.Service
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_get_services_proto_init() }
func file_barber_rpc_get_services_proto_init() {
	if File_barber_rpc_get_services_proto != nil {
		return
	}
	file_barber_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_get_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryAndServiceDetailsRequest); i {
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
		file_barber_rpc_get_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicesRequest); i {
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
		file_barber_rpc_get_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicesResponse); i {
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
			RawDescriptor: file_barber_rpc_get_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_get_services_proto_goTypes,
		DependencyIndexes: file_barber_rpc_get_services_proto_depIdxs,
		MessageInfos:      file_barber_rpc_get_services_proto_msgTypes,
	}.Build()
	File_barber_rpc_get_services_proto = out.File
	file_barber_rpc_get_services_proto_rawDesc = nil
	file_barber_rpc_get_services_proto_goTypes = nil
	file_barber_rpc_get_services_proto_depIdxs = nil
}
