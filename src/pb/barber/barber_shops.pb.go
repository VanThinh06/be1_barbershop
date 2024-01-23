// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/barber_shops.proto

package barber

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type BarberShops struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarbershopChainId *string                 `protobuf:"bytes,2,opt,name=barbershop_chain_id,json=barbershopChainId,proto3,oneof" json:"barbershop_chain_id,omitempty"`
	Name              string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	IsMainBranch      bool                    `protobuf:"varint,4,opt,name=is_main_branch,json=isMainBranch,proto3" json:"is_main_branch,omitempty"`
	BranchCount       int32                   `protobuf:"varint,5,opt,name=branch_count,json=branchCount,proto3" json:"branch_count,omitempty"`
	Coordinates       string                  `protobuf:"bytes,6,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Address           string                  `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Image             string                  `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`
	Status            int32                   `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Rate              float32                 `protobuf:"fixed32,10,opt,name=rate,proto3" json:"rate,omitempty"`
	StartTime         *TimeOfDay              `protobuf:"bytes,11,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time,omitempty"`
	EndTime           *TimeOfDay              `protobuf:"bytes,12,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
	BreakTime         *TimeOfDay              `protobuf:"bytes,13,opt,name=break_time,json=breakTime,proto3,oneof" json:"break_time,omitempty"`
	BreakMinutes      int32                   `protobuf:"varint,14,opt,name=break_minutes,json=breakMinutes,proto3" json:"break_minutes,omitempty"`
	IntervalScheduler int32                   `protobuf:"varint,15,opt,name=interval_scheduler,json=intervalScheduler,proto3" json:"interval_scheduler,omitempty"`
	Reputation        bool                    `protobuf:"varint,16,opt,name=reputation,proto3" json:"reputation,omitempty"`
	Distance          float32                 `protobuf:"fixed32,17,opt,name=distance,proto3" json:"distance,omitempty"`
	Longitude         *wrapperspb.DoubleValue `protobuf:"bytes,18,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude          *wrapperspb.DoubleValue `protobuf:"bytes,19,opt,name=latitude,proto3" json:"latitude,omitempty"`
	CreateAt          *timestamppb.Timestamp  `protobuf:"bytes,20,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt          *timestamppb.Timestamp  `protobuf:"bytes,21,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *BarberShops) Reset() {
	*x = BarberShops{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barber_shops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarberShops) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarberShops) ProtoMessage() {}

func (x *BarberShops) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barber_shops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarberShops.ProtoReflect.Descriptor instead.
func (*BarberShops) Descriptor() ([]byte, []int) {
	return file_barber_barber_shops_proto_rawDescGZIP(), []int{0}
}

func (x *BarberShops) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BarberShops) GetBarbershopChainId() string {
	if x != nil && x.BarbershopChainId != nil {
		return *x.BarbershopChainId
	}
	return ""
}

func (x *BarberShops) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BarberShops) GetIsMainBranch() bool {
	if x != nil {
		return x.IsMainBranch
	}
	return false
}

func (x *BarberShops) GetBranchCount() int32 {
	if x != nil {
		return x.BranchCount
	}
	return 0
}

func (x *BarberShops) GetCoordinates() string {
	if x != nil {
		return x.Coordinates
	}
	return ""
}

func (x *BarberShops) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BarberShops) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BarberShops) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BarberShops) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *BarberShops) GetStartTime() *TimeOfDay {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *BarberShops) GetEndTime() *TimeOfDay {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *BarberShops) GetBreakTime() *TimeOfDay {
	if x != nil {
		return x.BreakTime
	}
	return nil
}

func (x *BarberShops) GetBreakMinutes() int32 {
	if x != nil {
		return x.BreakMinutes
	}
	return 0
}

func (x *BarberShops) GetIntervalScheduler() int32 {
	if x != nil {
		return x.IntervalScheduler
	}
	return 0
}

func (x *BarberShops) GetReputation() bool {
	if x != nil {
		return x.Reputation
	}
	return false
}

func (x *BarberShops) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *BarberShops) GetLongitude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Longitude
	}
	return nil
}

func (x *BarberShops) GetLatitude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Latitude
	}
	return nil
}

func (x *BarberShops) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *BarberShops) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

var File_barber_barber_shops_proto protoreflect.FileDescriptor

var file_barber_barber_shops_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66,
	0x5f, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x06, 0x0a, 0x0b, 0x42,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x13, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x4d,
	0x61, 0x69, 0x6e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48, 0x01, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48, 0x02,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0a,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48,
	0x03, 0x52, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x0d, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x37,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_barber_shops_proto_rawDescOnce sync.Once
	file_barber_barber_shops_proto_rawDescData = file_barber_barber_shops_proto_rawDesc
)

func file_barber_barber_shops_proto_rawDescGZIP() []byte {
	file_barber_barber_shops_proto_rawDescOnce.Do(func() {
		file_barber_barber_shops_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_barber_shops_proto_rawDescData)
	})
	return file_barber_barber_shops_proto_rawDescData
}

var file_barber_barber_shops_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_barber_barber_shops_proto_goTypes = []interface{}{
	(*BarberShops)(nil),            // 0: pb.BarberShops
	(*TimeOfDay)(nil),              // 1: pb.TimeOfDay
	(*wrapperspb.DoubleValue)(nil), // 2: google.protobuf.DoubleValue
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_barber_barber_shops_proto_depIdxs = []int32{
	1, // 0: pb.BarberShops.start_time:type_name -> pb.TimeOfDay
	1, // 1: pb.BarberShops.end_time:type_name -> pb.TimeOfDay
	1, // 2: pb.BarberShops.break_time:type_name -> pb.TimeOfDay
	2, // 3: pb.BarberShops.longitude:type_name -> google.protobuf.DoubleValue
	2, // 4: pb.BarberShops.latitude:type_name -> google.protobuf.DoubleValue
	3, // 5: pb.BarberShops.create_at:type_name -> google.protobuf.Timestamp
	3, // 6: pb.BarberShops.update_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_barber_barber_shops_proto_init() }
func file_barber_barber_shops_proto_init() {
	if File_barber_barber_shops_proto != nil {
		return
	}
	file_barber_time_of_day_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_barber_shops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarberShops); i {
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
	file_barber_barber_shops_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_barber_shops_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_barber_shops_proto_goTypes,
		DependencyIndexes: file_barber_barber_shops_proto_depIdxs,
		MessageInfos:      file_barber_barber_shops_proto_msgTypes,
	}.Build()
	File_barber_barber_shops_proto = out.File
	file_barber_barber_shops_proto_rawDesc = nil
	file_barber_barber_shops_proto_goTypes = nil
	file_barber_barber_shops_proto_depIdxs = nil
}
