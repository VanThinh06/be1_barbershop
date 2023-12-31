// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/barber_shop.proto

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

type BarberShop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId            string                  `protobuf:"bytes,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	Status            int32                   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	OwnerId           string                  `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Name              string                  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Coordinates       string                  `protobuf:"bytes,5,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Address           string                  `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Image             string                  `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	ListImage         []string                `protobuf:"bytes,8,rep,name=list_image,json=listImage,proto3" json:"list_image,omitempty"`
	CreatedAt         *timestamppb.Timestamp  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Distance          float32                 `protobuf:"fixed32,10,opt,name=distance,proto3" json:"distance,omitempty"`
	Longitude         *wrapperspb.DoubleValue `protobuf:"bytes,11,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude          *wrapperspb.DoubleValue `protobuf:"bytes,12,opt,name=latitude,proto3" json:"latitude,omitempty"`
	IsReputation      bool                    `protobuf:"varint,13,opt,name=is_reputation,json=isReputation,proto3" json:"is_reputation,omitempty"`
	Rate              float32                 `protobuf:"fixed32,14,opt,name=rate,proto3" json:"rate,omitempty"`
	Facility          int32                   `protobuf:"varint,15,opt,name=facility,proto3" json:"facility,omitempty"`
	ChainId           *string                 `protobuf:"bytes,16,opt,name=chain_id,json=chainId,proto3,oneof" json:"chain_id,omitempty"`
	StartTime         *TimeOfDay              `protobuf:"bytes,17,opt,name=start_time,json=startTime,proto3,oneof" json:"start_time,omitempty"`
	EndTime           *TimeOfDay              `protobuf:"bytes,18,opt,name=end_time,json=endTime,proto3,oneof" json:"end_time,omitempty"`
	BreakTime         *TimeOfDay              `protobuf:"bytes,19,opt,name=break_time,json=breakTime,proto3,oneof" json:"break_time,omitempty"`
	BreakMinutes      int32                   `protobuf:"varint,20,opt,name=break_minutes,json=breakMinutes,proto3" json:"break_minutes,omitempty"`
	IntervalScheduler int32                   `protobuf:"varint,21,opt,name=interval_scheduler,json=intervalScheduler,proto3" json:"interval_scheduler,omitempty"`
}

func (x *BarberShop) Reset() {
	*x = BarberShop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barber_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarberShop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarberShop) ProtoMessage() {}

func (x *BarberShop) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barber_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarberShop.ProtoReflect.Descriptor instead.
func (*BarberShop) Descriptor() ([]byte, []int) {
	return file_barber_barber_shop_proto_rawDescGZIP(), []int{0}
}

func (x *BarberShop) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *BarberShop) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BarberShop) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *BarberShop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BarberShop) GetCoordinates() string {
	if x != nil {
		return x.Coordinates
	}
	return ""
}

func (x *BarberShop) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BarberShop) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BarberShop) GetListImage() []string {
	if x != nil {
		return x.ListImage
	}
	return nil
}

func (x *BarberShop) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BarberShop) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *BarberShop) GetLongitude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Longitude
	}
	return nil
}

func (x *BarberShop) GetLatitude() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Latitude
	}
	return nil
}

func (x *BarberShop) GetIsReputation() bool {
	if x != nil {
		return x.IsReputation
	}
	return false
}

func (x *BarberShop) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *BarberShop) GetFacility() int32 {
	if x != nil {
		return x.Facility
	}
	return 0
}

func (x *BarberShop) GetChainId() string {
	if x != nil && x.ChainId != nil {
		return *x.ChainId
	}
	return ""
}

func (x *BarberShop) GetStartTime() *TimeOfDay {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *BarberShop) GetEndTime() *TimeOfDay {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *BarberShop) GetBreakTime() *TimeOfDay {
	if x != nil {
		return x.BreakTime
	}
	return nil
}

func (x *BarberShop) GetBreakMinutes() int32 {
	if x != nil {
		return x.BreakMinutes
	}
	return 0
}

func (x *BarberShop) GetIntervalScheduler() int32 {
	if x != nil {
		return x.IntervalScheduler
	}
	return 0
}

var File_barber_barber_shop_proto protoreflect.FileDescriptor

var file_barber_barber_shop_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f,
	0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x06, 0x0a, 0x0a, 0x42, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x08, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48, 0x01,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x48,
	0x02, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a,
	0x0a, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79,
	0x48, 0x03, 0x52, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x4d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x1a, 0x5a, 0x18,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70,
	0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_barber_shop_proto_rawDescOnce sync.Once
	file_barber_barber_shop_proto_rawDescData = file_barber_barber_shop_proto_rawDesc
)

func file_barber_barber_shop_proto_rawDescGZIP() []byte {
	file_barber_barber_shop_proto_rawDescOnce.Do(func() {
		file_barber_barber_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_barber_shop_proto_rawDescData)
	})
	return file_barber_barber_shop_proto_rawDescData
}

var file_barber_barber_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_barber_barber_shop_proto_goTypes = []interface{}{
	(*BarberShop)(nil),             // 0: pb.BarberShop
	(*timestamppb.Timestamp)(nil),  // 1: google.protobuf.Timestamp
	(*wrapperspb.DoubleValue)(nil), // 2: google.protobuf.DoubleValue
	(*TimeOfDay)(nil),              // 3: pb.TimeOfDay
}
var file_barber_barber_shop_proto_depIdxs = []int32{
	1, // 0: pb.BarberShop.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: pb.BarberShop.longitude:type_name -> google.protobuf.DoubleValue
	2, // 2: pb.BarberShop.latitude:type_name -> google.protobuf.DoubleValue
	3, // 3: pb.BarberShop.start_time:type_name -> pb.TimeOfDay
	3, // 4: pb.BarberShop.end_time:type_name -> pb.TimeOfDay
	3, // 5: pb.BarberShop.break_time:type_name -> pb.TimeOfDay
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_barber_barber_shop_proto_init() }
func file_barber_barber_shop_proto_init() {
	if File_barber_barber_shop_proto != nil {
		return
	}
	file_barber_time_of_day_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_barber_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarberShop); i {
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
	file_barber_barber_shop_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_barber_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_barber_shop_proto_goTypes,
		DependencyIndexes: file_barber_barber_shop_proto_depIdxs,
		MessageInfos:      file_barber_barber_shop_proto_msgTypes,
	}.Build()
	File_barber_barber_shop_proto = out.File
	file_barber_barber_shop_proto_rawDesc = nil
	file_barber_barber_shop_proto_goTypes = nil
	file_barber_barber_shop_proto_depIdxs = nil
}
