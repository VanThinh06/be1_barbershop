// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_barbers_get_employee.proto

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

type GetBarberEmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShopId string `protobuf:"bytes,1,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
	Limit        int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page         int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetBarberEmployeeRequest) Reset() {
	*x = GetBarberEmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbers_get_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBarberEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBarberEmployeeRequest) ProtoMessage() {}

func (x *GetBarberEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbers_get_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBarberEmployeeRequest.ProtoReflect.Descriptor instead.
func (*GetBarberEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbers_get_employee_proto_rawDescGZIP(), []int{0}
}

func (x *GetBarberEmployeeRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *GetBarberEmployeeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetBarberEmployeeRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetBarberEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberDetails  []*BarberDetail `protobuf:"bytes,1,rep,name=barber_details,json=barberDetails,proto3" json:"barber_details,omitempty"`
	TotalEmployees int32           `protobuf:"varint,2,opt,name=total_employees,json=totalEmployees,proto3" json:"total_employees,omitempty"`
}

func (x *GetBarberEmployeeResponse) Reset() {
	*x = GetBarberEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbers_get_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBarberEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBarberEmployeeResponse) ProtoMessage() {}

func (x *GetBarberEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbers_get_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBarberEmployeeResponse.ProtoReflect.Descriptor instead.
func (*GetBarberEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbers_get_employee_proto_rawDescGZIP(), []int{1}
}

func (x *GetBarberEmployeeResponse) GetBarberDetails() []*BarberDetail {
	if x != nil {
		return x.BarberDetails
	}
	return nil
}

func (x *GetBarberEmployeeResponse) GetTotalEmployees() int32 {
	if x != nil {
		return x.TotalEmployees
	}
	return 0
}

var File_barber_rpc_barbers_get_employee_proto protoreflect.FileDescriptor

var file_barber_rpc_barbers_get_employee_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x14, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x42, 0x1a, 0x5a, 0x18,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70,
	0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbers_get_employee_proto_rawDescOnce sync.Once
	file_barber_rpc_barbers_get_employee_proto_rawDescData = file_barber_rpc_barbers_get_employee_proto_rawDesc
)

func file_barber_rpc_barbers_get_employee_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbers_get_employee_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbers_get_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbers_get_employee_proto_rawDescData)
	})
	return file_barber_rpc_barbers_get_employee_proto_rawDescData
}

var file_barber_rpc_barbers_get_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbers_get_employee_proto_goTypes = []interface{}{
	(*GetBarberEmployeeRequest)(nil),  // 0: pb.GetBarberEmployeeRequest
	(*GetBarberEmployeeResponse)(nil), // 1: pb.GetBarberEmployeeResponse
	(*BarberDetail)(nil),              // 2: pb.BarberDetail
}
var file_barber_rpc_barbers_get_employee_proto_depIdxs = []int32{
	2, // 0: pb.GetBarberEmployeeResponse.barber_details:type_name -> pb.BarberDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbers_get_employee_proto_init() }
func file_barber_rpc_barbers_get_employee_proto_init() {
	if File_barber_rpc_barbers_get_employee_proto != nil {
		return
	}
	file_barber_barbers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbers_get_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBarberEmployeeRequest); i {
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
		file_barber_rpc_barbers_get_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBarberEmployeeResponse); i {
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
			RawDescriptor: file_barber_rpc_barbers_get_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbers_get_employee_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbers_get_employee_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbers_get_employee_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbers_get_employee_proto = out.File
	file_barber_rpc_barbers_get_employee_proto_rawDesc = nil
	file_barber_rpc_barbers_get_employee_proto_goTypes = nil
	file_barber_rpc_barbers_get_employee_proto_depIdxs = nil
}
