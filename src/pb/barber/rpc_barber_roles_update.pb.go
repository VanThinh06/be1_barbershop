// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_barber_roles_update.proto

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

type UpdateBarberRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoleId int32  `protobuf:"varint,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *UpdateBarberRolesRequest) Reset() {
	*x = UpdateBarberRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barber_roles_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarberRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarberRolesRequest) ProtoMessage() {}

func (x *UpdateBarberRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barber_roles_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarberRolesRequest.ProtoReflect.Descriptor instead.
func (*UpdateBarberRolesRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barber_roles_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBarberRolesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBarberRolesRequest) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type UpdateBarberRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberRoles *BarberRoles `protobuf:"bytes,1,opt,name=barber_roles,json=barberRoles,proto3" json:"barber_roles,omitempty"`
}

func (x *UpdateBarberRolesResponse) Reset() {
	*x = UpdateBarberRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barber_roles_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarberRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarberRolesResponse) ProtoMessage() {}

func (x *UpdateBarberRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barber_roles_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarberRolesResponse.ProtoReflect.Descriptor instead.
func (*UpdateBarberRolesResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barber_roles_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBarberRolesResponse) GetBarberRoles() *BarberRoles {
	if x != nil {
		return x.BarberRoles
	}
	return nil
}

var File_barber_rpc_barber_roles_update_proto protoreflect.FileDescriptor

var file_barber_rpc_barber_roles_update_proto_rawDesc = []byte{
	0x0a, 0x24, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x0b,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x1a, 0x5a, 0x18, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62,
	0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barber_roles_update_proto_rawDescOnce sync.Once
	file_barber_rpc_barber_roles_update_proto_rawDescData = file_barber_rpc_barber_roles_update_proto_rawDesc
)

func file_barber_rpc_barber_roles_update_proto_rawDescGZIP() []byte {
	file_barber_rpc_barber_roles_update_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barber_roles_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barber_roles_update_proto_rawDescData)
	})
	return file_barber_rpc_barber_roles_update_proto_rawDescData
}

var file_barber_rpc_barber_roles_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barber_roles_update_proto_goTypes = []interface{}{
	(*UpdateBarberRolesRequest)(nil),  // 0: pb.UpdateBarberRolesRequest
	(*UpdateBarberRolesResponse)(nil), // 1: pb.UpdateBarberRolesResponse
	(*BarberRoles)(nil),               // 2: pb.BarberRoles
}
var file_barber_rpc_barber_roles_update_proto_depIdxs = []int32{
	2, // 0: pb.UpdateBarberRolesResponse.barber_roles:type_name -> pb.BarberRoles
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barber_roles_update_proto_init() }
func file_barber_rpc_barber_roles_update_proto_init() {
	if File_barber_rpc_barber_roles_update_proto != nil {
		return
	}
	file_barber_barber_roles_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barber_roles_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarberRolesRequest); i {
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
		file_barber_rpc_barber_roles_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarberRolesResponse); i {
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
			RawDescriptor: file_barber_rpc_barber_roles_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barber_roles_update_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barber_roles_update_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barber_roles_update_proto_msgTypes,
	}.Build()
	File_barber_rpc_barber_roles_update_proto = out.File
	file_barber_rpc_barber_roles_update_proto_rawDesc = nil
	file_barber_rpc_barber_roles_update_proto_goTypes = nil
	file_barber_rpc_barber_roles_update_proto_depIdxs = nil
}
