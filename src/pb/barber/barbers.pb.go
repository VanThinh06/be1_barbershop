// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/barbers.proto

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

type Barbers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GenderId   int32  `protobuf:"varint,2,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone      string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	NickName   string `protobuf:"bytes,5,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	FullName   string `protobuf:"bytes,6,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Haircut    bool   `protobuf:"varint,7,opt,name=haircut,proto3" json:"haircut,omitempty"`
	AvatarUrl  string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	WorkStatus bool   `protobuf:"varint,9,opt,name=work_status,json=workStatus,proto3" json:"work_status,omitempty"`
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

func (x *Barbers) GetGenderId() int32 {
	if x != nil {
		return x.GenderId
	}
	return 0
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

func (x *Barbers) GetWorkStatus() bool {
	if x != nil {
		return x.WorkStatus
	}
	return false
}

type BarberEmployee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone    string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *BarberEmployee) Reset() {
	*x = BarberEmployee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarberEmployee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarberEmployee) ProtoMessage() {}

func (x *BarberEmployee) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarberEmployee.ProtoReflect.Descriptor instead.
func (*BarberEmployee) Descriptor() ([]byte, []int) {
	return file_barber_barbers_proto_rawDescGZIP(), []int{1}
}

func (x *BarberEmployee) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BarberEmployee) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *BarberEmployee) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type BarberDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Barber     *Barbers     `protobuf:"bytes,1,opt,name=barber,proto3" json:"barber,omitempty"`
	BarberRole *BarberRoles `protobuf:"bytes,2,opt,name=barber_role,json=barberRole,proto3" json:"barber_role,omitempty"`
}

func (x *BarberDetail) Reset() {
	*x = BarberDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarberDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarberDetail) ProtoMessage() {}

func (x *BarberDetail) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarberDetail.ProtoReflect.Descriptor instead.
func (*BarberDetail) Descriptor() ([]byte, []int) {
	return file_barber_barbers_proto_rawDescGZIP(), []int{2}
}

func (x *BarberDetail) GetBarber() *Barbers {
	if x != nil {
		return x.Barber
	}
	return nil
}

func (x *BarberDetail) GetBarberRole() *BarberRoles {
	if x != nil {
		return x.BarberRole
	}
	return nil
}

type BarberUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GenderId   *int32  `protobuf:"varint,2,opt,name=gender_id,json=genderId,proto3,oneof" json:"gender_id,omitempty"`
	Email      *string `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Phone      *string `protobuf:"bytes,4,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Nickname   *string `protobuf:"bytes,5,opt,name=nickname,proto3,oneof" json:"nickname,omitempty"`
	FullName   *string `protobuf:"bytes,6,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Haircut    *bool   `protobuf:"varint,7,opt,name=haircut,proto3,oneof" json:"haircut,omitempty"`
	WorkStatus *bool   `protobuf:"varint,8,opt,name=work_status,json=workStatus,proto3,oneof" json:"work_status,omitempty"`
	AvatarUrl  *string `protobuf:"bytes,9,opt,name=avatar_url,json=avatarUrl,proto3,oneof" json:"avatar_url,omitempty"`
}

func (x *BarberUpdate) Reset() {
	*x = BarberUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarberUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarberUpdate) ProtoMessage() {}

func (x *BarberUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarberUpdate.ProtoReflect.Descriptor instead.
func (*BarberUpdate) Descriptor() ([]byte, []int) {
	return file_barber_barbers_proto_rawDescGZIP(), []int{3}
}

func (x *BarberUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BarberUpdate) GetGenderId() int32 {
	if x != nil && x.GenderId != nil {
		return *x.GenderId
	}
	return 0
}

func (x *BarberUpdate) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *BarberUpdate) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *BarberUpdate) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *BarberUpdate) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *BarberUpdate) GetHaircut() bool {
	if x != nil && x.Haircut != nil {
		return *x.Haircut
	}
	return false
}

func (x *BarberUpdate) GetWorkStatus() bool {
	if x != nil && x.WorkStatus != nil {
		return *x.WorkStatus
	}
	return false
}

func (x *BarberUpdate) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

var File_barber_barbers_proto protoreflect.FileDescriptor

var file_barber_barbers_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69,
	0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x69, 0x72, 0x63, 0x75, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x69, 0x72, 0x63, 0x75, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x60,
	0x0a, 0x0e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x65, 0x0a, 0x0c, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x23, 0x0a, 0x06, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x52, 0x06, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x0a, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x8a, 0x03, 0x0a, 0x0c, 0x42, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1f, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x68, 0x61, 0x69, 0x72, 0x63, 0x75, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x07, 0x68, 0x61, 0x69, 0x72, 0x63, 0x75, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x68, 0x61,
	0x69, 0x72, 0x63, 0x75, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_barber_barbers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_barber_barbers_proto_goTypes = []interface{}{
	(*Barbers)(nil),        // 0: pb.Barbers
	(*BarberEmployee)(nil), // 1: pb.BarberEmployee
	(*BarberDetail)(nil),   // 2: pb.BarberDetail
	(*BarberUpdate)(nil),   // 3: pb.BarberUpdate
	(*BarberRoles)(nil),    // 4: pb.BarberRoles
}
var file_barber_barbers_proto_depIdxs = []int32{
	0, // 0: pb.BarberDetail.barber:type_name -> pb.Barbers
	4, // 1: pb.BarberDetail.barber_role:type_name -> pb.BarberRoles
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_barber_barbers_proto_init() }
func file_barber_barbers_proto_init() {
	if File_barber_barbers_proto != nil {
		return
	}
	file_barber_barber_roles_proto_init()
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
		file_barber_barbers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarberEmployee); i {
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
		file_barber_barbers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarberDetail); i {
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
		file_barber_barbers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarberUpdate); i {
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
	file_barber_barbers_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_barbers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
