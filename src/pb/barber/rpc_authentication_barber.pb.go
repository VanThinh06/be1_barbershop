// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_authentication_barber.proto

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

	FullName string  `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone    string  `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    *string `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Password string  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateBarberRequest) Reset() {
	*x = CreateBarberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberRequest) ProtoMessage() {}

func (x *CreateBarberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[0]
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
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{0}
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
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *CreateBarberRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateBarberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barber *Barbers `protobuf:"bytes,1,opt,name=barber,proto3" json:"barber,omitempty"`
}

func (x *CreateBarberResponse) Reset() {
	*x = CreateBarberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberResponse) ProtoMessage() {}

func (x *CreateBarberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[1]
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
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBarberResponse) GetBarber() *Barbers {
	if x != nil {
		return x.Barber
	}
	return nil
}

type ForgotPasswordBarberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ForgotPasswordBarberRequest) Reset() {
	*x = ForgotPasswordBarberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordBarberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordBarberRequest) ProtoMessage() {}

func (x *ForgotPasswordBarberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordBarberRequest.ProtoReflect.Descriptor instead.
func (*ForgotPasswordBarberRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{2}
}

func (x *ForgotPasswordBarberRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ForgotPasswordBarberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ForgotPasswordBarberResponse) Reset() {
	*x = ForgotPasswordBarberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordBarberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordBarberResponse) ProtoMessage() {}

func (x *ForgotPasswordBarberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordBarberResponse.ProtoReflect.Descriptor instead.
func (*ForgotPasswordBarberResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{3}
}

func (x *ForgotPasswordBarberResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ForgotPasswordBarberResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyOtpBarberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	BarberId string `protobuf:"bytes,2,opt,name=barber_id,json=barberId,proto3" json:"barber_id,omitempty"`
}

func (x *VerifyOtpBarberRequest) Reset() {
	*x = VerifyOtpBarberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOtpBarberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOtpBarberRequest) ProtoMessage() {}

func (x *VerifyOtpBarberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOtpBarberRequest.ProtoReflect.Descriptor instead.
func (*VerifyOtpBarberRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyOtpBarberRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyOtpBarberRequest) GetBarberId() string {
	if x != nil {
		return x.BarberId
	}
	return ""
}

type VerifyOtpBarberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VerifyOtpBarberResponse) Reset() {
	*x = VerifyOtpBarberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOtpBarberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOtpBarberResponse) ProtoMessage() {}

func (x *VerifyOtpBarberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOtpBarberResponse.ProtoReflect.Descriptor instead.
func (*VerifyOtpBarberResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyOtpBarberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResetPasswordBarberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberId string `protobuf:"bytes,1,opt,name=barber_id,json=barberId,proto3" json:"barber_id,omitempty"`
	Code     string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ResetPasswordBarberRequest) Reset() {
	*x = ResetPasswordBarberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordBarberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordBarberRequest) ProtoMessage() {}

func (x *ResetPasswordBarberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordBarberRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordBarberRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{6}
}

func (x *ResetPasswordBarberRequest) GetBarberId() string {
	if x != nil {
		return x.BarberId
	}
	return ""
}

func (x *ResetPasswordBarberRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResetPasswordBarberRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResetPasswordBarberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResetPasswordBarberResponse) Reset() {
	*x = ResetPasswordBarberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_authentication_barber_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordBarberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordBarberResponse) ProtoMessage() {}

func (x *ResetPasswordBarberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_authentication_barber_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordBarberResponse.ProtoReflect.Descriptor instead.
func (*ResetPasswordBarberResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_authentication_barber_proto_rawDescGZIP(), []int{7}
}

func (x *ResetPasswordBarberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_barber_rpc_authentication_barber_proto protoreflect.FileDescriptor

var file_barber_rpc_authentication_barber_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3b, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x06, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x1b, 0x46, 0x6f,
	0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x1c, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x49, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74,
	0x70, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x33, 0x0a, 0x17, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x74, 0x70, 0x42, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x1a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x37, 0x0a, 0x1b, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_authentication_barber_proto_rawDescOnce sync.Once
	file_barber_rpc_authentication_barber_proto_rawDescData = file_barber_rpc_authentication_barber_proto_rawDesc
)

func file_barber_rpc_authentication_barber_proto_rawDescGZIP() []byte {
	file_barber_rpc_authentication_barber_proto_rawDescOnce.Do(func() {
		file_barber_rpc_authentication_barber_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_authentication_barber_proto_rawDescData)
	})
	return file_barber_rpc_authentication_barber_proto_rawDescData
}

var file_barber_rpc_authentication_barber_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_barber_rpc_authentication_barber_proto_goTypes = []interface{}{
	(*CreateBarberRequest)(nil),          // 0: pb.CreateBarberRequest
	(*CreateBarberResponse)(nil),         // 1: pb.CreateBarberResponse
	(*ForgotPasswordBarberRequest)(nil),  // 2: pb.ForgotPasswordBarberRequest
	(*ForgotPasswordBarberResponse)(nil), // 3: pb.ForgotPasswordBarberResponse
	(*VerifyOtpBarberRequest)(nil),       // 4: pb.VerifyOtpBarberRequest
	(*VerifyOtpBarberResponse)(nil),      // 5: pb.VerifyOtpBarberResponse
	(*ResetPasswordBarberRequest)(nil),   // 6: pb.ResetPasswordBarberRequest
	(*ResetPasswordBarberResponse)(nil),  // 7: pb.ResetPasswordBarberResponse
	(*Barbers)(nil),                      // 8: pb.Barbers
}
var file_barber_rpc_authentication_barber_proto_depIdxs = []int32{
	8, // 0: pb.CreateBarberResponse.barber:type_name -> pb.Barbers
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_barber_rpc_authentication_barber_proto_init() }
func file_barber_rpc_authentication_barber_proto_init() {
	if File_barber_rpc_authentication_barber_proto != nil {
		return
	}
	file_barber_barber_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_authentication_barber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_barber_rpc_authentication_barber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_barber_rpc_authentication_barber_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordBarberRequest); i {
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
		file_barber_rpc_authentication_barber_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordBarberResponse); i {
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
		file_barber_rpc_authentication_barber_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOtpBarberRequest); i {
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
		file_barber_rpc_authentication_barber_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyOtpBarberResponse); i {
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
		file_barber_rpc_authentication_barber_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordBarberRequest); i {
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
		file_barber_rpc_authentication_barber_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordBarberResponse); i {
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
	file_barber_rpc_authentication_barber_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_rpc_authentication_barber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_authentication_barber_proto_goTypes,
		DependencyIndexes: file_barber_rpc_authentication_barber_proto_depIdxs,
		MessageInfos:      file_barber_rpc_authentication_barber_proto_msgTypes,
	}.Build()
	File_barber_rpc_authentication_barber_proto = out.File
	file_barber_rpc_authentication_barber_proto_rawDesc = nil
	file_barber_rpc_authentication_barber_proto_goTypes = nil
	file_barber_rpc_authentication_barber_proto_depIdxs = nil
}
