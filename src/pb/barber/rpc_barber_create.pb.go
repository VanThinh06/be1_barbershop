// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_barber_create.proto

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

type CreateBarberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname       string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	FullName       string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone          string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email          string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Gender         int32  `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Password       string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	CodeBarberShop string `protobuf:"bytes,7,opt,name=code_barber_shop,json=codeBarberShop,proto3" json:"code_barber_shop,omitempty"`
	Role           int32  `protobuf:"varint,8,opt,name=role,proto3" json:"role,omitempty"`
	Haircut        bool   `protobuf:"varint,9,opt,name=haircut,proto3" json:"haircut,omitempty"`
}

func (x *CreateBarberRequest) Reset() {
	*x = CreateBarberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barber_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberRequest) ProtoMessage() {}

func (x *CreateBarberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barber_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberRequest.ProtoReflect.Descriptor instead.
func (*CreateBarberRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barber_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBarberRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateBarberRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateBarberRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateBarberRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateBarberRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *CreateBarberRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateBarberRequest) GetCodeBarberShop() string {
	if x != nil {
		return x.CodeBarberShop
	}
	return ""
}

func (x *CreateBarberRequest) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *CreateBarberRequest) GetHaircut() bool {
	if x != nil {
		return x.Haircut
	}
	return false
}

type CreateBarberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barber *Barber `protobuf:"bytes,1,opt,name=barber,proto3" json:"barber,omitempty"`
}

func (x *CreateBarberResponse) Reset() {
	*x = CreateBarberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barber_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberResponse) ProtoMessage() {}

func (x *CreateBarberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barber_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberResponse.ProtoReflect.Descriptor instead.
func (*CreateBarberResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barber_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBarberResponse) GetBarber() *Barber {
	if x != nil {
		return x.Barber
	}
	return nil
}

var File_barber_rpc_barber_create_proto protoreflect.FileDescriptor

var file_barber_rpc_barber_create_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x02, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x53, 0x68, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x69, 0x72,
	0x63, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x69, 0x72, 0x63,
	0x75, 0x74, 0x22, 0x3a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x06, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x42, 0x1a,
	0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_barber_rpc_barber_create_proto_rawDescOnce sync.Once
	file_barber_rpc_barber_create_proto_rawDescData = file_barber_rpc_barber_create_proto_rawDesc
)

func file_barber_rpc_barber_create_proto_rawDescGZIP() []byte {
	file_barber_rpc_barber_create_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barber_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barber_create_proto_rawDescData)
	})
	return file_barber_rpc_barber_create_proto_rawDescData
}

var file_barber_rpc_barber_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barber_create_proto_goTypes = []interface{}{
	(*CreateBarberRequest)(nil),  // 0: pb.CreateBarberRequest
	(*CreateBarberResponse)(nil), // 1: pb.CreateBarberResponse
	(*Barber)(nil),               // 2: pb.Barber
}
var file_barber_rpc_barber_create_proto_depIdxs = []int32{
	2, // 0: pb.CreateBarberResponse.barber:type_name -> pb.Barber
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barber_create_proto_init() }
func file_barber_rpc_barber_create_proto_init() {
	if File_barber_rpc_barber_create_proto != nil {
		return
	}
	file_barber_barber_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barber_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberRequest); i {
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
		file_barber_rpc_barber_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberResponse); i {
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
			RawDescriptor: file_barber_rpc_barber_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barber_create_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barber_create_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barber_create_proto_msgTypes,
	}.Build()
	File_barber_rpc_barber_create_proto = out.File
	file_barber_rpc_barber_create_proto_rawDesc = nil
	file_barber_rpc_barber_create_proto_goTypes = nil
	file_barber_rpc_barber_create_proto_depIdxs = nil
}
