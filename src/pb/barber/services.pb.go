// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/services.proto

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

type Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId  string                 `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Timer       *int32                 `protobuf:"varint,4,opt,name=timer,proto3,oneof" json:"timer,omitempty"`
	Price       *int32                 `protobuf:"varint,5,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Description *string                `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Image       *string                `protobuf:"bytes,7,opt,name=image,proto3,oneof" json:"image,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdateAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *Services) Reset() {
	*x = Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Services) ProtoMessage() {}

func (x *Services) ProtoReflect() protoreflect.Message {
	mi := &file_barber_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Services.ProtoReflect.Descriptor instead.
func (*Services) Descriptor() ([]byte, []int) {
	return file_barber_services_proto_rawDescGZIP(), []int{0}
}

func (x *Services) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Services) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Services) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Services) GetTimer() int32 {
	if x != nil && x.Timer != nil {
		return *x.Timer
	}
	return 0
}

func (x *Services) GetPrice() int32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *Services) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Services) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *Services) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Services) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

var File_barber_services_proto protoreflect.FileDescriptor

var file_barber_services_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02, 0x0a,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_services_proto_rawDescOnce sync.Once
	file_barber_services_proto_rawDescData = file_barber_services_proto_rawDesc
)

func file_barber_services_proto_rawDescGZIP() []byte {
	file_barber_services_proto_rawDescOnce.Do(func() {
		file_barber_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_services_proto_rawDescData)
	})
	return file_barber_services_proto_rawDescData
}

var file_barber_services_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_barber_services_proto_goTypes = []interface{}{
	(*Services)(nil),              // 0: pb.Services
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_barber_services_proto_depIdxs = []int32{
	1, // 0: pb.Services.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: pb.Services.update_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_barber_services_proto_init() }
func file_barber_services_proto_init() {
	if File_barber_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Services); i {
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
	file_barber_services_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_services_proto_goTypes,
		DependencyIndexes: file_barber_services_proto_depIdxs,
		MessageInfos:      file_barber_services_proto_msgTypes,
	}.Build()
	File_barber_services_proto = out.File
	file_barber_services_proto_rawDesc = nil
	file_barber_services_proto_goTypes = nil
	file_barber_services_proto_depIdxs = nil
}
