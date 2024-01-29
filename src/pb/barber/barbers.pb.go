// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/barbers.proto

package barber

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Barbers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GenderId     string                 `protobuf:"bytes,2,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Email        string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone        string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	NickName     string                 `protobuf:"bytes,5,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	FullName     string                 `protobuf:"bytes,6,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Haircut      bool                   `protobuf:"varint,7,opt,name=haircut,proto3" json:"haircut,omitempty"`
	AvatarUrl    string                 `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	CreateAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	BarberRoleId *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=barber_role_id,json=barberRoleId,proto3" json:"barber_role_id,omitempty"`
}

func (x *Barbers) Reset() {
	*x = Barbers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Barbers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Barbers) ProtoMessage() {}

func (x *Barbers) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Barbers.ProtoReflect.Descriptor instead.
func (*Barbers) Descriptor() ([]byte, []int) {
	return file_barber_barbers_proto_rawDescGZIP(), []int{0}
}

func (x *Barbers) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Barbers) GetGenderId() string {
	if x != nil {
		return x.GenderId
	}
	return ""
}

func (x *Barbers) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Barbers) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Barbers) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *Barbers) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Barbers) GetHaircut() bool {
	if x != nil {
		return x.Haircut
	}
	return false
}

func (x *Barbers) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Barbers) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *Barbers) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

func (x *Barbers) GetBarberRoleId() *timestamppb.Timestamp {
	if x != nil {
		return x.BarberRoleId
	}
	return nil
}

var File_barber_barbers_proto protoreflect.FileDescriptor

var file_barber_barbers_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03, 0x0a, 0x07,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61,
	0x69, 0x72, 0x63, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x69,
	0x72, 0x63, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_barbers_proto_rawDescOnce sync.Once
	file_barber_barbers_proto_rawDescData = file_barber_barbers_proto_rawDesc
)

func file_barber_barbers_proto_rawDescGZIP() []byte {
	file_barber_barbers_proto_rawDescOnce.Do(func() {
		file_barber_barbers_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_barbers_proto_rawDescData)
	})
	return file_barber_barbers_proto_rawDescData
}

var file_barber_barbers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_barber_barbers_proto_goTypes = []interface{}{
	(*Barbers)(nil),               // 0: pb.Barbers
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_barber_barbers_proto_depIdxs = []int32{
	1, // 0: pb.Barbers.create_at:type_name -> google.protobuf.Timestamp
	1, // 1: pb.Barbers.update_at:type_name -> google.protobuf.Timestamp
	1, // 2: pb.Barbers.barber_role_id:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_barber_barbers_proto_init() }
func file_barber_barbers_proto_init() {
	if File_barber_barbers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_barbers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Barbers); i {
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
			RawDescriptor: file_barber_barbers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_barbers_proto_goTypes,
		DependencyIndexes: file_barber_barbers_proto_depIdxs,
		MessageInfos:      file_barber_barbers_proto_msgTypes,
	}.Build()
	File_barber_barbers_proto = out.File
	file_barber_barbers_proto_rawDesc = nil
	file_barber_barbers_proto_goTypes = nil
	file_barber_barbers_proto_depIdxs = nil
}
