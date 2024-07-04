// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_barbershop_chains_get.proto

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

type GetBarberShopChainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBarberShopChainsRequest) Reset() {
	*x = GetBarberShopChainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershop_chains_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBarberShopChainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBarberShopChainsRequest) ProtoMessage() {}

func (x *GetBarberShopChainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershop_chains_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBarberShopChainsRequest.ProtoReflect.Descriptor instead.
func (*GetBarberShopChainsRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershop_chains_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetBarberShopChainsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBarberShopChainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarbershopChain *BarberShopChains `protobuf:"bytes,1,opt,name=barbershop_chain,json=barbershopChain,proto3" json:"barbershop_chain,omitempty"`
}

func (x *GetBarberShopChainsResponse) Reset() {
	*x = GetBarberShopChainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershop_chains_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBarberShopChainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBarberShopChainsResponse) ProtoMessage() {}

func (x *GetBarberShopChainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershop_chains_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBarberShopChainsResponse.ProtoReflect.Descriptor instead.
func (*GetBarberShopChainsResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershop_chains_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetBarberShopChainsResponse) GetBarbershopChain() *BarberShopChains {
	if x != nil {
		return x.BarbershopChain
	}
	return nil
}

var File_barber_rpc_barbershop_chains_get_proto protoreflect.FileDescriptor

var file_barber_rpc_barbershop_chains_get_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x5f, 0x67,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1e, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x0f, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbershop_chains_get_proto_rawDescOnce sync.Once
	file_barber_rpc_barbershop_chains_get_proto_rawDescData = file_barber_rpc_barbershop_chains_get_proto_rawDesc
)

func file_barber_rpc_barbershop_chains_get_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbershop_chains_get_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbershop_chains_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbershop_chains_get_proto_rawDescData)
	})
	return file_barber_rpc_barbershop_chains_get_proto_rawDescData
}

var file_barber_rpc_barbershop_chains_get_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbershop_chains_get_proto_goTypes = []interface{}{
	(*GetBarberShopChainsRequest)(nil),  // 0: pb.GetBarberShopChainsRequest
	(*GetBarberShopChainsResponse)(nil), // 1: pb.GetBarberShopChainsResponse
	(*BarberShopChains)(nil),            // 2: pb.BarberShopChains
}
var file_barber_rpc_barbershop_chains_get_proto_depIdxs = []int32{
	2, // 0: pb.GetBarberShopChainsResponse.barbershop_chain:type_name -> pb.BarberShopChains
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbershop_chains_get_proto_init() }
func file_barber_rpc_barbershop_chains_get_proto_init() {
	if File_barber_rpc_barbershop_chains_get_proto != nil {
		return
	}
	file_barber_barbershop_chains_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbershop_chains_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBarberShopChainsRequest); i {
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
		file_barber_rpc_barbershop_chains_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBarberShopChainsResponse); i {
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
			RawDescriptor: file_barber_rpc_barbershop_chains_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbershop_chains_get_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbershop_chains_get_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbershop_chains_get_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbershop_chains_get_proto = out.File
	file_barber_rpc_barbershop_chains_get_proto_rawDesc = nil
	file_barber_rpc_barbershop_chains_get_proto_goTypes = nil
	file_barber_rpc_barbershop_chains_get_proto_depIdxs = nil
}
