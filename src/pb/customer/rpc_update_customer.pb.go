// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: customer/rpc_update_customer.proto

package customer

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

type UpdateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId        string                 `protobuf:"bytes,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Name              *string                `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Email             *string                `protobuf:"bytes,2,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Phone             *string                `protobuf:"bytes,3,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Gender            *int32                 `protobuf:"varint,4,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	Password          *string                `protobuf:"bytes,5,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Avatar            *string                `protobuf:"bytes,6,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=password_changed_at,json=passwordChangedAt,proto3,oneof" json:"password_changed_at,omitempty"`
	UpdateAt          *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *UpdateCustomerRequest) Reset() {
	*x = UpdateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_rpc_update_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerRequest) ProtoMessage() {}

func (x *UpdateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_rpc_update_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerRequest.ProtoReflect.Descriptor instead.
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_rpc_update_customer_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *UpdateCustomerRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateCustomerRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateCustomerRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateCustomerRequest) GetGender() int32 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *UpdateCustomerRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *UpdateCustomerRequest) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

func (x *UpdateCustomerRequest) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

func (x *UpdateCustomerRequest) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

type UpdateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *UpdateCustomerResponse) Reset() {
	*x = UpdateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_rpc_update_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerResponse) ProtoMessage() {}

func (x *UpdateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_rpc_update_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerResponse.ProtoReflect.Descriptor instead.
func (*UpdateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_rpc_update_customer_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

var File_customer_rpc_update_customer_proto protoreflect.FileDescriptor

var file_customer_rpc_update_customer_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc4, 0x03, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x06, 0x52, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x42, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x1c, 0x5a,
	0x1a, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_customer_rpc_update_customer_proto_rawDescOnce sync.Once
	file_customer_rpc_update_customer_proto_rawDescData = file_customer_rpc_update_customer_proto_rawDesc
)

func file_customer_rpc_update_customer_proto_rawDescGZIP() []byte {
	file_customer_rpc_update_customer_proto_rawDescOnce.Do(func() {
		file_customer_rpc_update_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_rpc_update_customer_proto_rawDescData)
	})
	return file_customer_rpc_update_customer_proto_rawDescData
}

var file_customer_rpc_update_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_customer_rpc_update_customer_proto_goTypes = []interface{}{
	(*UpdateCustomerRequest)(nil),  // 0: pb.UpdateCustomerRequest
	(*UpdateCustomerResponse)(nil), // 1: pb.UpdateCustomerResponse
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*Customer)(nil),               // 3: pb.Customer
}
var file_customer_rpc_update_customer_proto_depIdxs = []int32{
	2, // 0: pb.UpdateCustomerRequest.password_changed_at:type_name -> google.protobuf.Timestamp
	2, // 1: pb.UpdateCustomerRequest.update_at:type_name -> google.protobuf.Timestamp
	3, // 2: pb.UpdateCustomerResponse.customer:type_name -> pb.Customer
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customer_rpc_update_customer_proto_init() }
func file_customer_rpc_update_customer_proto_init() {
	if File_customer_rpc_update_customer_proto != nil {
		return
	}
	file_customer_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_customer_rpc_update_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCustomerRequest); i {
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
		file_customer_rpc_update_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCustomerResponse); i {
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
	file_customer_rpc_update_customer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_customer_rpc_update_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_customer_rpc_update_customer_proto_goTypes,
		DependencyIndexes: file_customer_rpc_update_customer_proto_depIdxs,
		MessageInfos:      file_customer_rpc_update_customer_proto_msgTypes,
	}.Build()
	File_customer_rpc_update_customer_proto = out.File
	file_customer_rpc_update_customer_proto_rawDesc = nil
	file_customer_rpc_update_customer_proto_goTypes = nil
	file_customer_rpc_update_customer_proto_depIdxs = nil
}
