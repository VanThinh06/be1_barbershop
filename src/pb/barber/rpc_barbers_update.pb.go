// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_barbers_update.proto

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

type UpdateBarbersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarbershopId *string `protobuf:"bytes,2,opt,name=barbershop_id,json=barbershopId,proto3,oneof" json:"barbershop_id,omitempty"`
	GenderId     *int32  `protobuf:"varint,3,opt,name=gender_id,json=genderId,proto3,oneof" json:"gender_id,omitempty"`
	Email        *string `protobuf:"bytes,4,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Phone        *string `protobuf:"bytes,5,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Nickname     *string `protobuf:"bytes,6,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"`
	FullName     *string `protobuf:"bytes,7,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Haircut      bool    `protobuf:"varint,8,opt,name=haircut,proto3" json:"haircut,omitempty"`
	AvatarUrl    *string `protobuf:"bytes,9,opt,name=avatar_url,json=avatarUrl,proto3,oneof" json:"avatar_url,omitempty"`
}

func (x *UpdateBarbersRequest) Reset() {
	*x = UpdateBarbersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbers_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarbersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarbersRequest) ProtoMessage() {}

func (x *UpdateBarbersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbers_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarbersRequest.ProtoReflect.Descriptor instead.
func (*UpdateBarbersRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbers_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBarbersRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBarbersRequest) GetBarbershopId() string {
	if x != nil && x.BarbershopId != nil {
		return *x.BarbershopId
	}
	return ""
}

func (x *UpdateBarbersRequest) GetGenderId() int32 {
	if x != nil && x.GenderId != nil {
		return *x.GenderId
	}
	return 0
}

func (x *UpdateBarbersRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateBarbersRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateBarbersRequest) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *UpdateBarbersRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UpdateBarbersRequest) GetHaircut() bool {
	if x != nil {
		return x.Haircut
	}
	return false
}

func (x *UpdateBarbersRequest) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

type UpdateBarbersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barber *Barbers `protobuf:"bytes,1,opt,name=barber,proto3" json:"barber,omitempty"`
}

func (x *UpdateBarbersResponse) Reset() {
	*x = UpdateBarbersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbers_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarbersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarbersResponse) ProtoMessage() {}

func (x *UpdateBarbersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbers_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarbersResponse.ProtoReflect.Descriptor instead.
func (*UpdateBarbersResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbers_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBarbersResponse) GetBarber() *Barbers {
	if x != nil {
		return x.Barber
	}
	return nil
}

var File_barber_rpc_barbers_update_proto protoreflect.FileDescriptor

var file_barber_rpc_barbers_update_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x14, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x69,
	0x72, 0x63, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x69, 0x72,
	0x63, 0x75, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x3c, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x06, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x52, 0x06, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f,
	0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbers_update_proto_rawDescOnce sync.Once
	file_barber_rpc_barbers_update_proto_rawDescData = file_barber_rpc_barbers_update_proto_rawDesc
)

func file_barber_rpc_barbers_update_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbers_update_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbers_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbers_update_proto_rawDescData)
	})
	return file_barber_rpc_barbers_update_proto_rawDescData
}

var file_barber_rpc_barbers_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbers_update_proto_goTypes = []interface{}{
	(*UpdateBarbersRequest)(nil),  // 0: pb.UpdateBarbersRequest
	(*UpdateBarbersResponse)(nil), // 1: pb.UpdateBarbersResponse
	(*Barbers)(nil),               // 2: pb.Barbers
}
var file_barber_rpc_barbers_update_proto_depIdxs = []int32{
	2, // 0: pb.UpdateBarbersResponse.barber:type_name -> pb.Barbers
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbers_update_proto_init() }
func file_barber_rpc_barbers_update_proto_init() {
	if File_barber_rpc_barbers_update_proto != nil {
		return
	}
	file_barber_barbers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbers_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarbersRequest); i {
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
		file_barber_rpc_barbers_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarbersResponse); i {
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
	file_barber_rpc_barbers_update_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_rpc_barbers_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbers_update_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbers_update_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbers_update_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbers_update_proto = out.File
	file_barber_rpc_barbers_update_proto_rawDesc = nil
	file_barber_rpc_barbers_update_proto_goTypes = nil
	file_barber_rpc_barbers_update_proto_depIdxs = nil
}
