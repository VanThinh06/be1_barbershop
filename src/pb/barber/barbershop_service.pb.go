// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/barbershop_service.proto

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

type BarberShopService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarberShopId      string                 `protobuf:"bytes,2,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
	CategoryId        int32                  `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	GenderId          int32                  `protobuf:"varint,4,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Name              string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Timer             int32                  `protobuf:"varint,6,opt,name=timer,proto3" json:"timer,omitempty"`
	Price             float32                `protobuf:"fixed32,7,opt,name=price,proto3" json:"price,omitempty"`
	DiscountPrice     *float32               `protobuf:"fixed32,8,opt,name=discount_price,json=discountPrice,proto3,oneof" json:"discount_price,omitempty"`
	DiscountStartTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=discount_start_time,json=discountStartTime,proto3,oneof" json:"discount_start_time,omitempty"`
	DiscountEndTime   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=discount_end_time,json=discountEndTime,proto3,oneof" json:"discount_end_time,omitempty"`
	Description       string                 `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl          string                 `protobuf:"bytes,12,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	CategoryName      string                 `protobuf:"bytes,13,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	IsActive          bool                   `protobuf:"varint,14,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *BarberShopService) Reset() {
	*x = BarberShopService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbershop_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarberShopService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarberShopService) ProtoMessage() {}

func (x *BarberShopService) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbershop_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarberShopService.ProtoReflect.Descriptor instead.
func (*BarberShopService) Descriptor() ([]byte, []int) {
	return file_barber_barbershop_service_proto_rawDescGZIP(), []int{0}
}

func (x *BarberShopService) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BarberShopService) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *BarberShopService) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *BarberShopService) GetGenderId() int32 {
	if x != nil {
		return x.GenderId
	}
	return 0
}

func (x *BarberShopService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BarberShopService) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *BarberShopService) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BarberShopService) GetDiscountPrice() float32 {
	if x != nil && x.DiscountPrice != nil {
		return *x.DiscountPrice
	}
	return 0
}

func (x *BarberShopService) GetDiscountStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountStartTime
	}
	return nil
}

func (x *BarberShopService) GetDiscountEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountEndTime
	}
	return nil
}

func (x *BarberShopService) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BarberShopService) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *BarberShopService) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *BarberShopService) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type ComboService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarberShopId      string                 `protobuf:"bytes,2,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
	GenderId          int32                  `protobuf:"varint,3,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Name              string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Timer             int32                  `protobuf:"varint,5,opt,name=timer,proto3" json:"timer,omitempty"`
	Price             float32                `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	DiscountPrice     *float32               `protobuf:"fixed32,7,opt,name=discount_price,json=discountPrice,proto3,oneof" json:"discount_price,omitempty"`
	DiscountStartTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=discount_start_time,json=discountStartTime,proto3,oneof" json:"discount_start_time,omitempty"`
	DiscountEndTime   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=discount_end_time,json=discountEndTime,proto3,oneof" json:"discount_end_time,omitempty"`
	Description       string                 `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl          string                 `protobuf:"bytes,11,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	IsActive          bool                   `protobuf:"varint,12,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Services          []*BarberShopService   `protobuf:"bytes,13,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ComboService) Reset() {
	*x = ComboService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbershop_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComboService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComboService) ProtoMessage() {}

func (x *ComboService) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbershop_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComboService.ProtoReflect.Descriptor instead.
func (*ComboService) Descriptor() ([]byte, []int) {
	return file_barber_barbershop_service_proto_rawDescGZIP(), []int{1}
}

func (x *ComboService) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComboService) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *ComboService) GetGenderId() int32 {
	if x != nil {
		return x.GenderId
	}
	return 0
}

func (x *ComboService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComboService) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *ComboService) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ComboService) GetDiscountPrice() float32 {
	if x != nil && x.DiscountPrice != nil {
		return *x.DiscountPrice
	}
	return 0
}

func (x *ComboService) GetDiscountStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountStartTime
	}
	return nil
}

func (x *ComboService) GetDiscountEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountEndTime
	}
	return nil
}

func (x *ComboService) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ComboService) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *ComboService) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ComboService) GetServices() []*BarberShopService {
	if x != nil {
		return x.Services
	}
	return nil
}

type CategoryPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId int32 `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Position   int32 `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
	Visible    bool  `protobuf:"varint,3,opt,name=visible,proto3" json:"visible,omitempty"`
}

func (x *CategoryPosition) Reset() {
	*x = CategoryPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_barbershop_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryPosition) ProtoMessage() {}

func (x *CategoryPosition) ProtoReflect() protoreflect.Message {
	mi := &file_barber_barbershop_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryPosition.ProtoReflect.Descriptor instead.
func (*CategoryPosition) Descriptor() ([]byte, []int) {
	return file_barber_barbershop_service_proto_rawDescGZIP(), []int{2}
}

func (x *CategoryPosition) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CategoryPosition) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *CategoryPosition) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

var File_barber_barbershop_service_proto protoreflect.FileDescriptor

var file_barber_barbershop_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x04, 0x0a, 0x11, 0x42, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x13,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a,
	0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x16,
	0x0a, 0x14, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xbb, 0x04, 0x0a,
	0x0c, 0x43, 0x6f, 0x6d, 0x62, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x13,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a,
	0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x10, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_barbershop_service_proto_rawDescOnce sync.Once
	file_barber_barbershop_service_proto_rawDescData = file_barber_barbershop_service_proto_rawDesc
)

func file_barber_barbershop_service_proto_rawDescGZIP() []byte {
	file_barber_barbershop_service_proto_rawDescOnce.Do(func() {
		file_barber_barbershop_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_barbershop_service_proto_rawDescData)
	})
	return file_barber_barbershop_service_proto_rawDescData
}

var file_barber_barbershop_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_barber_barbershop_service_proto_goTypes = []interface{}{
	(*BarberShopService)(nil),     // 0: pb.BarberShopService
	(*ComboService)(nil),          // 1: pb.ComboService
	(*CategoryPosition)(nil),      // 2: pb.CategoryPosition
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_barber_barbershop_service_proto_depIdxs = []int32{
	3, // 0: pb.BarberShopService.discount_start_time:type_name -> google.protobuf.Timestamp
	3, // 1: pb.BarberShopService.discount_end_time:type_name -> google.protobuf.Timestamp
	3, // 2: pb.ComboService.discount_start_time:type_name -> google.protobuf.Timestamp
	3, // 3: pb.ComboService.discount_end_time:type_name -> google.protobuf.Timestamp
	0, // 4: pb.ComboService.services:type_name -> pb.BarberShopService
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_barber_barbershop_service_proto_init() }
func file_barber_barbershop_service_proto_init() {
	if File_barber_barbershop_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_barbershop_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarberShopService); i {
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
		file_barber_barbershop_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComboService); i {
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
		file_barber_barbershop_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryPosition); i {
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
	file_barber_barbershop_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_barber_barbershop_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_barbershop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_barbershop_service_proto_goTypes,
		DependencyIndexes: file_barber_barbershop_service_proto_depIdxs,
		MessageInfos:      file_barber_barbershop_service_proto_msgTypes,
	}.Build()
	File_barber_barbershop_service_proto = out.File
	file_barber_barbershop_service_proto_rawDesc = nil
	file_barber_barbershop_service_proto_goTypes = nil
	file_barber_barbershop_service_proto_depIdxs = nil
}
