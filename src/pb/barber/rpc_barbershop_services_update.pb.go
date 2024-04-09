// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_barbershop_services_update.proto

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

type UpdateBarberShopServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GenderId    string  `protobuf:"bytes,2,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Timer       int32   `protobuf:"varint,4,opt,name=timer,proto3" json:"timer,omitempty"`
	Price       float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Description string  `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Image       string  `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *UpdateBarberShopServicesRequest) Reset() {
	*x = UpdateBarberShopServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershop_services_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarberShopServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarberShopServicesRequest) ProtoMessage() {}

func (x *UpdateBarberShopServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershop_services_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarberShopServicesRequest.ProtoReflect.Descriptor instead.
func (*UpdateBarberShopServicesRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershop_services_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBarberShopServicesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBarberShopServicesRequest) GetGenderId() string {
	if x != nil {
		return x.GenderId
	}
	return ""
}

func (x *UpdateBarberShopServicesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBarberShopServicesRequest) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *UpdateBarberShopServicesRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateBarberShopServicesRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateBarberShopServicesRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type UpdateBarberShopServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarbershopService *BarberShopServices `protobuf:"bytes,1,opt,name=barbershop_service,json=barbershopService,proto3" json:"barbershop_service,omitempty"`
}

func (x *UpdateBarberShopServicesResponse) Reset() {
	*x = UpdateBarberShopServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershop_services_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarberShopServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarberShopServicesResponse) ProtoMessage() {}

func (x *UpdateBarberShopServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershop_services_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarberShopServicesResponse.ProtoReflect.Descriptor instead.
func (*UpdateBarberShopServicesResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershop_services_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBarberShopServicesResponse) GetBarbershopService() *BarberShopServices {
	if x != nil {
		return x.BarbershopService
	}
	return nil
}

var File_barber_rpc_barbershop_services_update_proto protoreflect.FileDescriptor

var file_barber_rpc_barbershop_services_update_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x20,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x12, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x11, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbershop_services_update_proto_rawDescOnce sync.Once
	file_barber_rpc_barbershop_services_update_proto_rawDescData = file_barber_rpc_barbershop_services_update_proto_rawDesc
)

func file_barber_rpc_barbershop_services_update_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbershop_services_update_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbershop_services_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbershop_services_update_proto_rawDescData)
	})
	return file_barber_rpc_barbershop_services_update_proto_rawDescData
}

var file_barber_rpc_barbershop_services_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbershop_services_update_proto_goTypes = []interface{}{
	(*UpdateBarberShopServicesRequest)(nil),  // 0: pb.UpdateBarberShopServicesRequest
	(*UpdateBarberShopServicesResponse)(nil), // 1: pb.UpdateBarberShopServicesResponse
	(*BarberShopServices)(nil),               // 2: pb.BarberShopServices
}
var file_barber_rpc_barbershop_services_update_proto_depIdxs = []int32{
	2, // 0: pb.UpdateBarberShopServicesResponse.barbershop_service:type_name -> pb.BarberShopServices
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbershop_services_update_proto_init() }
func file_barber_rpc_barbershop_services_update_proto_init() {
	if File_barber_rpc_barbershop_services_update_proto != nil {
		return
	}
	file_barber_barbershop_services_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbershop_services_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarberShopServicesRequest); i {
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
		file_barber_rpc_barbershop_services_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarberShopServicesResponse); i {
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
			RawDescriptor: file_barber_rpc_barbershop_services_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbershop_services_update_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbershop_services_update_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbershop_services_update_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbershop_services_update_proto = out.File
	file_barber_rpc_barbershop_services_update_proto_rawDesc = nil
	file_barber_rpc_barbershop_services_update_proto_goTypes = nil
	file_barber_rpc_barbershop_services_update_proto_depIdxs = nil
}
