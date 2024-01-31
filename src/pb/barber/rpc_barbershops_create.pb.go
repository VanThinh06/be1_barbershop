// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_barbershops_create.proto

package barber

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBarberShopsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarbershopChainId *string                 `protobuf:"bytes,1,opt,name=barbershop_chain_id,json=barbershopChainId,proto3,oneof" json:"barbershop_chain_id,omitempty"`
	Name              string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsMainBranch      *bool                   `protobuf:"varint,3,opt,name=is_main_branch,json=isMainBranch,proto3,oneof" json:"is_main_branch,omitempty"`
	BranchCount       int32                   `protobuf:"varint,4,opt,name=branch_count,json=branchCount,proto3" json:"branch_count,omitempty"`
	Longitude         *wrapperspb.DoubleValue `protobuf:"bytes,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude          *wrapperspb.DoubleValue `protobuf:"bytes,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Address           string                  `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Image             *string                 `protobuf:"bytes,8,opt,name=image,proto3,oneof" json:"image,omitempty"`
}

func (x *CreateBarberShopsRequest) Reset() {
	*x = CreateBarberShopsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberShopsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberShopsRequest) ProtoMessage() {}

func (x *CreateBarberShopsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberShopsRequest.ProtoReflect.Descriptor instead.
func (*CreateBarberShopsRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBarberShopsRequest) GetBarbershopChainId() string {
	if x != nil && x.BarbershopChainId != nil {
		return *x.BarbershopChainId
	}
	return ""
}

func (x *CreateBarberShopsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBarberShopsRequest) GetIsMainBranch() bool {
	if x != nil && x.IsMainBranch != nil {
		return *x.IsMainBranch
	}
	return false
}

func (x *CreateBarberShopsRequest) GetBranchCount() int32 {
	if x != nil {
		return x.BranchCount
	}
	return 0
}

func (x *CreateBarberShopsRequest) GetLongitude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Longitude
	}
	return nil
}

func (x *CreateBarberShopsRequest) GetLatitude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Latitude
	}
	return nil
}

func (x *CreateBarberShopsRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateBarberShopsRequest) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

type CreateBarberShopsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShop *BarberShops `protobuf:"bytes,1,opt,name=barber_shop,json=barberShop,proto3" json:"barber_shop,omitempty"`
}

func (x *CreateBarberShopsResponse) Reset() {
	*x = CreateBarberShopsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberShopsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberShopsResponse) ProtoMessage() {}

func (x *CreateBarberShopsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberShopsResponse.ProtoReflect.Descriptor instead.
func (*CreateBarberShopsResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBarberShopsResponse) GetBarberShop() *BarberShops {
	if x != nil {
		return x.BarberShop
	}
	return nil
}

var File_barber_rpc_barbershops_create_proto protoreflect.FileDescriptor

var file_barber_rpc_barbershops_create_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x13, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x11, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0e, 0x69, 0x73,
	0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x01, 0x52, 0x0c, 0x69, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x69, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x0a, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbershops_create_proto_rawDescOnce sync.Once
	file_barber_rpc_barbershops_create_proto_rawDescData = file_barber_rpc_barbershops_create_proto_rawDesc
)

func file_barber_rpc_barbershops_create_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbershops_create_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbershops_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbershops_create_proto_rawDescData)
	})
	return file_barber_rpc_barbershops_create_proto_rawDescData
}

var file_barber_rpc_barbershops_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbershops_create_proto_goTypes = []interface{}{
	(*CreateBarberShopsRequest)(nil),  // 0: pb.CreateBarberShopsRequest
	(*CreateBarberShopsResponse)(nil), // 1: pb.CreateBarberShopsResponse
	(*wrapperspb.DoubleValue)(nil),    // 2: google.protobuf.DoubleValue
	(*BarberShops)(nil),               // 3: pb.BarberShops
}
var file_barber_rpc_barbershops_create_proto_depIdxs = []int32{
	2, // 0: pb.CreateBarberShopsRequest.longitude:type_name -> google.protobuf.DoubleValue
	2, // 1: pb.CreateBarberShopsRequest.latitude:type_name -> google.protobuf.DoubleValue
	3, // 2: pb.CreateBarberShopsResponse.barber_shop:type_name -> pb.BarberShops
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbershops_create_proto_init() }
func file_barber_rpc_barbershops_create_proto_init() {
	if File_barber_rpc_barbershops_create_proto != nil {
		return
	}
	file_barber_barber_shops_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbershops_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberShopsRequest); i {
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
		file_barber_rpc_barbershops_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberShopsResponse); i {
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
	file_barber_rpc_barbershops_create_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_rpc_barbershops_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbershops_create_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbershops_create_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbershops_create_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbershops_create_proto = out.File
	file_barber_rpc_barbershops_create_proto_rawDesc = nil
	file_barber_rpc_barbershops_create_proto_goTypes = nil
	file_barber_rpc_barbershops_create_proto_depIdxs = nil
}