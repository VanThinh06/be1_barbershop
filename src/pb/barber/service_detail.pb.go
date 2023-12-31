// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/service_detail.proto

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

type ServiceDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId   string   `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	NameCategory string   `protobuf:"bytes,8,opt,name=name_category,json=nameCategory,proto3" json:"name_category,omitempty"`
	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Timer        *int32   `protobuf:"varint,4,opt,name=timer,proto3,oneof" json:"timer,omitempty"`
	Price        *float32 `protobuf:"fixed32,5,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Description  *string  `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Image        *string  `protobuf:"bytes,7,opt,name=image,proto3,oneof" json:"image,omitempty"`
}

func (x *ServiceDetail) Reset() {
	*x = ServiceDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_service_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDetail) ProtoMessage() {}

func (x *ServiceDetail) ProtoReflect() protoreflect.Message {
	mi := &file_barber_service_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDetail.ProtoReflect.Descriptor instead.
func (*ServiceDetail) Descriptor() ([]byte, []int) {
	return file_barber_service_detail_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceDetail) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *ServiceDetail) GetNameCategory() string {
	if x != nil {
		return x.NameCategory
	}
	return ""
}

func (x *ServiceDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceDetail) GetTimer() int32 {
	if x != nil && x.Timer != nil {
		return *x.Timer
	}
	return 0
}

func (x *ServiceDetail) GetPrice() float32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *ServiceDetail) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ServiceDetail) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

var File_barber_service_detail_proto protoreflect.FileDescriptor

var file_barber_service_detail_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x9f, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x61, 0x6d,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f,
	0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_service_detail_proto_rawDescOnce sync.Once
	file_barber_service_detail_proto_rawDescData = file_barber_service_detail_proto_rawDesc
)

func file_barber_service_detail_proto_rawDescGZIP() []byte {
	file_barber_service_detail_proto_rawDescOnce.Do(func() {
		file_barber_service_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_service_detail_proto_rawDescData)
	})
	return file_barber_service_detail_proto_rawDescData
}

var file_barber_service_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_barber_service_detail_proto_goTypes = []interface{}{
	(*ServiceDetail)(nil), // 0: pb.ServiceDetail
}
var file_barber_service_detail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_barber_service_detail_proto_init() }
func file_barber_service_detail_proto_init() {
	if File_barber_service_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_service_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDetail); i {
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
	file_barber_service_detail_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_service_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_service_detail_proto_goTypes,
		DependencyIndexes: file_barber_service_detail_proto_depIdxs,
		MessageInfos:      file_barber_service_detail_proto_msgTypes,
	}.Build()
	File_barber_service_detail_proto = out.File
	file_barber_service_detail_proto_rawDesc = nil
	file_barber_service_detail_proto_goTypes = nil
	file_barber_service_detail_proto_depIdxs = nil
}
