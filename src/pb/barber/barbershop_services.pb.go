// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/barbershop_services.proto

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

type BarberShopServices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarbershopCategoryId string                 `protobuf:"bytes,2,opt,name=barbershop_category_id,json=barbershopCategoryId,proto3" json:"barbershop_category_id,omitempty"`
	BarbershopChainId    string                 `protobuf:"bytes,3,opt,name=barbershop_chain_id,json=barbershopChainId,proto3" json:"barbershop_chain_id,omitempty"`
	BarbershopId         string                 `protobuf:"bytes,4,opt,name=barbershop_id,json=barbershopId,proto3" json:"barbershop_id,omitempty"`
	GenderId             int32                  `protobuf:"varint,5,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Name                 string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Timer                int32                  `protobuf:"varint,7,opt,name=timer,proto3" json:"timer,omitempty"`
	Price                float32                `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
	Description          string                 `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Image                string                 `protobuf:"bytes,10,opt,name=image,proto3" json:"image,omitempty"`
	IsCustom             bool                   `protobuf:"varint,11,opt,name=is_custom,json=isCustom,proto3" json:"is_custom,omitempty"`
	IsRemoved            bool                   `protobuf:"varint,12,opt,name=is_removed,json=isRemoved,proto3" json:"is_removed,omitempty"`
	CreateAt             *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt             *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *BarberShopServices) Reset() {
	*x = BarberShopServices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbershop_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarberShopServices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarberShopServices) ProtoMessage() {}

func (x *BarberShopServices) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbershop_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarberShopServices.ProtoReflect.Descriptor instead.
func (*BarberShopServices) Descriptor() ([]byte, []int) {
	return file_barber_barbershop_services_proto_rawDescGZIP(), []int{0}
}

func (x *BarberShopServices) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BarberShopServices) GetBarbershopCategoryId() string {
	if x != nil {
		return x.BarbershopCategoryId
	}
	return ""
}

func (x *BarberShopServices) GetBarbershopChainId() string {
	if x != nil {
		return x.BarbershopChainId
	}
	return ""
}

func (x *BarberShopServices) GetBarbershopId() string {
	if x != nil {
		return x.BarbershopId
	}
	return ""
}

func (x *BarberShopServices) GetGenderId() int32 {
	if x != nil {
		return x.GenderId
	}
	return 0
}

func (x *BarberShopServices) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BarberShopServices) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *BarberShopServices) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BarberShopServices) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BarberShopServices) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BarberShopServices) GetIsCustom() bool {
	if x != nil {
		return x.IsCustom
	}
	return false
}

func (x *BarberShopServices) GetIsRemoved() bool {
	if x != nil {
		return x.IsRemoved
	}
	return false
}

func (x *BarberShopServices) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *BarberShopServices) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

var File_barber_barbershop_services_proto protoreflect.FileDescriptor

var file_barber_barbershop_services_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x03, 0x0a, 0x12, 0x42, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34,
	0x0a, 0x16, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x6d, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x42, 0x1a, 0x5a, 0x18,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70,
	0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_barbershop_services_proto_rawDescOnce sync.Once
	file_barber_barbershop_services_proto_rawDescData = file_barber_barbershop_services_proto_rawDesc
)

func file_barber_barbershop_services_proto_rawDescGZIP() []byte {
	file_barber_barbershop_services_proto_rawDescOnce.Do(func() {
		file_barber_barbershop_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_barbershop_services_proto_rawDescData)
	})
	return file_barber_barbershop_services_proto_rawDescData
}

var file_barber_barbershop_services_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_barber_barbershop_services_proto_goTypes = []interface{}{
	(*BarberShopServices)(nil),    // 0: pb.BarberShopServices
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_barber_barbershop_services_proto_depIdxs = []int32{
	1, // 0: pb.BarberShopServices.create_at:type_name -> google.protobuf.Timestamp
	1, // 1: pb.BarberShopServices.update_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_barber_barbershop_services_proto_init() }
func file_barber_barbershop_services_proto_init() {
	if File_barber_barbershop_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_barbershop_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarberShopServices); i {
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
			RawDescriptor: file_barber_barbershop_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_barbershop_services_proto_goTypes,
		DependencyIndexes: file_barber_barbershop_services_proto_depIdxs,
		MessageInfos:      file_barber_barbershop_services_proto_msgTypes,
	}.Build()
	File_barber_barbershop_services_proto = out.File
	file_barber_barbershop_services_proto_rawDesc = nil
	file_barber_barbershop_services_proto_goTypes = nil
	file_barber_barbershop_services_proto_depIdxs = nil
}
