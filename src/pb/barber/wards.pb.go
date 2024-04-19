// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/wards.proto

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

type Wards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DistrictId int32  `protobuf:"varint,3,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
}

func (x *Wards) Reset() {
	*x = Wards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_wards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wards) ProtoMessage() {}

func (x *Wards) ProtoReflect() protoreflect.Message {
	mi := &file_barber_wards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wards.ProtoReflect.Descriptor instead.
func (*Wards) Descriptor() ([]byte, []int) {
	return file_barber_wards_proto_rawDescGZIP(), []int{0}
}

func (x *Wards) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Wards) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Wards) GetDistrictId() int32 {
	if x != nil {
		return x.DistrictId
	}
	return 0
}

var File_barber_wards_proto protoreflect.FileDescriptor

var file_barber_wards_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x4c, 0x0a, 0x05, 0x57, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_wards_proto_rawDescOnce sync.Once
	file_barber_wards_proto_rawDescData = file_barber_wards_proto_rawDesc
)

func file_barber_wards_proto_rawDescGZIP() []byte {
	file_barber_wards_proto_rawDescOnce.Do(func() {
		file_barber_wards_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_wards_proto_rawDescData)
	})
	return file_barber_wards_proto_rawDescData
}

var file_barber_wards_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_barber_wards_proto_goTypes = []interface{}{
	(*Wards)(nil), // 0: pb.Wards
}
var file_barber_wards_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_barber_wards_proto_init() }
func file_barber_wards_proto_init() {
	if File_barber_wards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_wards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wards); i {
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
			RawDescriptor: file_barber_wards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_wards_proto_goTypes,
		DependencyIndexes: file_barber_wards_proto_depIdxs,
		MessageInfos:      file_barber_wards_proto_msgTypes,
	}.Build()
	File_barber_wards_proto = out.File
	file_barber_wards_proto_rawDesc = nil
	file_barber_wards_proto_goTypes = nil
	file_barber_wards_proto_depIdxs = nil
}
