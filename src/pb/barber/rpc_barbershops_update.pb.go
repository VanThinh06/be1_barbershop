// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/rpc_barbershops_update.proto

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

type UpdateBarberShopsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             *string                 `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	IsMainBranch     *string                 `protobuf:"bytes,2,opt,name=is_main_branch,json=isMainBranch,proto3,oneof" json:"is_main_branch,omitempty"`
	BranchCount      *int32                  `protobuf:"varint,3,opt,name=branch_count,json=branchCount,proto3,oneof" json:"branch_count,omitempty"`
	Latitude         *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longtude         *wrapperspb.DoubleValue `protobuf:"bytes,5,opt,name=longtude,proto3,oneof" json:"longtude,omitempty"`
	Address          *string                 `protobuf:"bytes,6,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Image            *string                 `protobuf:"bytes,7,opt,name=image,proto3,oneof" json:"image,omitempty"`
	Status           *int32                  `protobuf:"varint,8,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Rate             *int32                  `protobuf:"varint,9,opt,name=rate,proto3,oneof" json:"rate,omitempty"`
	StartTime        *TimeOfDay              `protobuf:"bytes,10,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time,omitempty"`
	EndTime          *TimeOfDay              `protobuf:"bytes,11,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
	BreakTime        *TimeOfDay              `protobuf:"bytes,12,opt,name=break_time,json=breakTime,proto3,oneof" json:"break_time,omitempty"`
	IntervalSheduler *int32                  `protobuf:"varint,13,opt,name=interval_sheduler,json=intervalSheduler,proto3,oneof" json:"interval_sheduler,omitempty"`
	BreakMinutes     *int32                  `protobuf:"varint,14,opt,name=break_minutes,json=breakMinutes,proto3,oneof" json:"break_minutes,omitempty"`
	Id               string                  `protobuf:"bytes,15,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateBarberShopsRequest) Reset() {
	*x = UpdateBarberShopsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarberShopsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarberShopsRequest) ProtoMessage() {}

func (x *UpdateBarberShopsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarberShopsRequest.ProtoReflect.Descriptor instead.
func (*UpdateBarberShopsRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBarberShopsRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateBarberShopsRequest) GetIsMainBranch() string {
	if x != nil && x.IsMainBranch != nil {
		return *x.IsMainBranch
	}
	return ""
}

func (x *UpdateBarberShopsRequest) GetBranchCount() int32 {
	if x != nil && x.BranchCount != nil {
		return *x.BranchCount
	}
	return 0
}

func (x *UpdateBarberShopsRequest) GetLatitude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Latitude
	}
	return nil
}

func (x *UpdateBarberShopsRequest) GetLongtude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Longtude
	}
	return nil
}

func (x *UpdateBarberShopsRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *UpdateBarberShopsRequest) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *UpdateBarberShopsRequest) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *UpdateBarberShopsRequest) GetRate() int32 {
	if x != nil && x.Rate != nil {
		return *x.Rate
	}
	return 0
}

func (x *UpdateBarberShopsRequest) GetStartTime() *TimeOfDay {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *UpdateBarberShopsRequest) GetEndTime() *TimeOfDay {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *UpdateBarberShopsRequest) GetBreakTime() *TimeOfDay {
	if x != nil {
		return x.BreakTime
	}
	return nil
}

func (x *UpdateBarberShopsRequest) GetIntervalSheduler() int32 {
	if x != nil && x.IntervalSheduler != nil {
		return *x.IntervalSheduler
	}
	return 0
}

func (x *UpdateBarberShopsRequest) GetBreakMinutes() int32 {
	if x != nil && x.BreakMinutes != nil {
		return *x.BreakMinutes
	}
	return 0
}

func (x *UpdateBarberShopsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBarberShopsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShop *BarberShops `protobuf:"bytes,1,opt,name=barber_shop,json=barberShop,proto3" json:"barber_shop,omitempty"`
}

func (x *UpdateBarberShopsResponse) Reset() {
	*x = UpdateBarberShopsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBarberShopsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBarberShopsResponse) ProtoMessage() {}

func (x *UpdateBarberShopsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBarberShopsResponse.ProtoReflect.Descriptor instead.
func (*UpdateBarberShopsResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBarberShopsResponse) GetBarberShop() *BarberShops {
	if x != nil {
		return x.BarberShop
	}
	return nil
}

var File_barber_rpc_barbershops_update_proto protoreflect.FileDescriptor

var file_barber_rpc_barbershops_update_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9,
	0x06, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c,
	0x69, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12,
	0x26, 0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x74, 0x75,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x07, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x48, 0x08, 0x52, 0x04, 0x72, 0x61,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48, 0x0a, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0a, 0x62, 0x72, 0x65, 0x61, 0x6b,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48, 0x0b, 0x52, 0x09, 0x62, 0x72,
	0x65, 0x61, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0c, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x53, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x0d, 0x52, 0x0c, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f,
	0x73, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x0a, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbershops_update_proto_rawDescOnce sync.Once
	file_barber_rpc_barbershops_update_proto_rawDescData = file_barber_rpc_barbershops_update_proto_rawDesc
)

func file_barber_rpc_barbershops_update_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbershops_update_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbershops_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbershops_update_proto_rawDescData)
	})
	return file_barber_rpc_barbershops_update_proto_rawDescData
}

var file_barber_rpc_barbershops_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbershops_update_proto_goTypes = []interface{}{
	(*UpdateBarberShopsRequest)(nil),  // 0: pb.UpdateBarberShopsRequest
	(*UpdateBarberShopsResponse)(nil), // 1: pb.UpdateBarberShopsResponse
	(*wrapperspb.DoubleValue)(nil),    // 2: google.protobuf.DoubleValue
	(*TimeOfDay)(nil),                 // 3: pb.TimeOfDay
	(*BarberShops)(nil),               // 4: pb.BarberShops
}
var file_barber_rpc_barbershops_update_proto_depIdxs = []int32{
	2, // 0: pb.UpdateBarberShopsRequest.latitude:type_name -> google.protobuf.DoubleValue
	2, // 1: pb.UpdateBarberShopsRequest.longtude:type_name -> google.protobuf.DoubleValue
	3, // 2: pb.UpdateBarberShopsRequest.start_time:type_name -> pb.TimeOfDay
	3, // 3: pb.UpdateBarberShopsRequest.end_time:type_name -> pb.TimeOfDay
	3, // 4: pb.UpdateBarberShopsRequest.break_time:type_name -> pb.TimeOfDay
	4, // 5: pb.UpdateBarberShopsResponse.barber_shop:type_name -> pb.BarberShops
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbershops_update_proto_init() }
func file_barber_rpc_barbershops_update_proto_init() {
	if File_barber_rpc_barbershops_update_proto != nil {
		return
	}
	file_barber_barber_shops_proto_init()
	file_barber_time_of_day_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbershops_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarberShopsRequest); i {
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
		file_barber_rpc_barbershops_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBarberShopsResponse); i {
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
	file_barber_rpc_barbershops_update_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_rpc_barbershops_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbershops_update_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbershops_update_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbershops_update_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbershops_update_proto = out.File
	file_barber_rpc_barbershops_update_proto_rawDesc = nil
	file_barber_rpc_barbershops_update_proto_goTypes = nil
	file_barber_rpc_barbershops_update_proto_depIdxs = nil
}
