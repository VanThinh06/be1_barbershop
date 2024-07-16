// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_barbers_delete_employee.proto

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

type DeleteBarberEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarberShopId string `protobuf:"bytes,2,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
}

func (x *DeleteBarberEmployeeRequest) Reset() {
	*x = DeleteBarberEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbers_delete_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBarberEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBarberEmployeeRequest) ProtoMessage() {}

func (x *DeleteBarberEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbers_delete_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBarberEmployeeRequest.ProtoReflect.Descriptor instead.
func (*DeleteBarberEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbers_delete_employee_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteBarberEmployeeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteBarberEmployeeRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

type DeleteBarberEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteBarberEmployeeResponse) Reset() {
	*x = DeleteBarberEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbers_delete_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBarberEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBarberEmployeeResponse) ProtoMessage() {}

func (x *DeleteBarberEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbers_delete_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBarberEmployeeResponse.ProtoReflect.Descriptor instead.
func (*DeleteBarberEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbers_delete_employee_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteBarberEmployeeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_barber_rpc_barbers_delete_employee_proto protoreflect.FileDescriptor

var file_barber_rpc_barbers_delete_employee_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x53,
	0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x5a,
	0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_barber_rpc_barbers_delete_employee_proto_rawDescOnce sync.Once
	file_barber_rpc_barbers_delete_employee_proto_rawDescData = file_barber_rpc_barbers_delete_employee_proto_rawDesc
)

func file_barber_rpc_barbers_delete_employee_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbers_delete_employee_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbers_delete_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbers_delete_employee_proto_rawDescData)
	})
	return file_barber_rpc_barbers_delete_employee_proto_rawDescData
}

var file_barber_rpc_barbers_delete_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbers_delete_employee_proto_goTypes = []interface{}{
	(*DeleteBarberEmployeeRequest)(nil),  // 0: pb.DeleteBarberEmployeeRequest
	(*DeleteBarberEmployeeResponse)(nil), // 1: pb.DeleteBarberEmployeeResponse
}
var file_barber_rpc_barbers_delete_employee_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbers_delete_employee_proto_init() }
func file_barber_rpc_barbers_delete_employee_proto_init() {
	if File_barber_rpc_barbers_delete_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbers_delete_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBarberEmployeeRequest); i {
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
		file_barber_rpc_barbers_delete_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBarberEmployeeResponse); i {
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
			RawDescriptor: file_barber_rpc_barbers_delete_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbers_delete_employee_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbers_delete_employee_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbers_delete_employee_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbers_delete_employee_proto = out.File
	file_barber_rpc_barbers_delete_employee_proto_rawDesc = nil
	file_barber_rpc_barbers_delete_employee_proto_goTypes = nil
	file_barber_rpc_barbers_delete_employee_proto_depIdxs = nil
}
