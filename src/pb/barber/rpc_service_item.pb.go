// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_service_item.proto

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

// insert
type CreateServiceItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId   int32   `protobuf:"varint,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	BarberShopId string  `protobuf:"bytes,2,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
	GenderId     int32   `protobuf:"varint,3,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Name         string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Timer        int32   `protobuf:"varint,5,opt,name=timer,proto3" json:"timer,omitempty"`
	Price        float32 `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Description  *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	ImageUrl     *string `protobuf:"bytes,8,opt,name=image_url,json=imageUrl,proto3,oneof" json:"image_url,omitempty"`
}

func (x *CreateServiceItemRequest) Reset() {
	*x = CreateServiceItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceItemRequest) ProtoMessage() {}

func (x *CreateServiceItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceItemRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceItemRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServiceItemRequest) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreateServiceItemRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *CreateServiceItemRequest) GetGenderId() int32 {
	if x != nil {
		return x.GenderId
	}
	return 0
}

func (x *CreateServiceItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServiceItemRequest) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *CreateServiceItemRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateServiceItemRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreateServiceItemRequest) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

type CreateServiceItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *ServiceItem `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *CreateServiceItemResponse) Reset() {
	*x = CreateServiceItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceItemResponse) ProtoMessage() {}

func (x *CreateServiceItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceItemResponse.ProtoReflect.Descriptor instead.
func (*CreateServiceItemResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServiceItemResponse) GetService() *ServiceItem {
	if x != nil {
		return x.Service
	}
	return nil
}

// detail
type GetServiceItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetServiceItemRequest) Reset() {
	*x = GetServiceItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceItemRequest) ProtoMessage() {}

func (x *GetServiceItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceItemRequest.ProtoReflect.Descriptor instead.
func (*GetServiceItemRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{2}
}

func (x *GetServiceItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetServiceItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *ServiceItem `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetServiceItemResponse) Reset() {
	*x = GetServiceItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceItemResponse) ProtoMessage() {}

func (x *GetServiceItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceItemResponse.ProtoReflect.Descriptor instead.
func (*GetServiceItemResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{3}
}

func (x *GetServiceItemResponse) GetService() *ServiceItem {
	if x != nil {
		return x.Service
	}
	return nil
}

// list
type ListServiceItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShopId string `protobuf:"bytes,1,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
}

func (x *ListServiceItemRequest) Reset() {
	*x = ListServiceItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceItemRequest) ProtoMessage() {}

func (x *ListServiceItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceItemRequest.ProtoReflect.Descriptor instead.
func (*ListServiceItemRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{4}
}

func (x *ListServiceItemRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

type ListServiceItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*ServiceItem `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListServiceItemResponse) Reset() {
	*x = ListServiceItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceItemResponse) ProtoMessage() {}

func (x *ListServiceItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceItemResponse.ProtoReflect.Descriptor instead.
func (*ListServiceItemResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{5}
}

func (x *ListServiceItemResponse) GetServices() []*ServiceItem {
	if x != nil {
		return x.Services
	}
	return nil
}

// update
type UpdateServiceItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarberShopId      string                 `protobuf:"bytes,2,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
	GenderId          *int32                 `protobuf:"varint,3,opt,name=gender_id,json=genderId,proto3,oneof" json:"gender_id,omitempty"`
	Name              *string                `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Timer             *int32                 `protobuf:"varint,5,opt,name=timer,proto3,oneof" json:"timer,omitempty"`
	Price             *float32               `protobuf:"fixed32,6,opt,name=price,proto3,oneof" json:"price,omitempty"`
	Description       *string                `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	ImageUrl          *string                `protobuf:"bytes,8,opt,name=image_url,json=imageUrl,proto3,oneof" json:"image_url,omitempty"`
	IsActive          *bool                  `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
	CategoryId        *int32                 `protobuf:"varint,10,opt,name=category_id,json=categoryId,proto3,oneof" json:"category_id,omitempty"`
	DiscountPrice     *float32               `protobuf:"fixed32,12,opt,name=discount_price,json=discountPrice,proto3,oneof" json:"discount_price,omitempty"`
	DiscountStartTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=discount_start_time,json=discountStartTime,proto3,oneof" json:"discount_start_time,omitempty"`
	DiscountEndTime   *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=discount_end_time,json=discountEndTime,proto3,oneof" json:"discount_end_time,omitempty"`
}

func (x *UpdateServiceItemRequest) Reset() {
	*x = UpdateServiceItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServiceItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServiceItemRequest) ProtoMessage() {}

func (x *UpdateServiceItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServiceItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateServiceItemRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateServiceItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateServiceItemRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *UpdateServiceItemRequest) GetGenderId() int32 {
	if x != nil && x.GenderId != nil {
		return *x.GenderId
	}
	return 0
}

func (x *UpdateServiceItemRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateServiceItemRequest) GetTimer() int32 {
	if x != nil && x.Timer != nil {
		return *x.Timer
	}
	return 0
}

func (x *UpdateServiceItemRequest) GetPrice() float32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *UpdateServiceItemRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateServiceItemRequest) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

func (x *UpdateServiceItemRequest) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

func (x *UpdateServiceItemRequest) GetCategoryId() int32 {
	if x != nil && x.CategoryId != nil {
		return *x.CategoryId
	}
	return 0
}

func (x *UpdateServiceItemRequest) GetDiscountPrice() float32 {
	if x != nil && x.DiscountPrice != nil {
		return *x.DiscountPrice
	}
	return 0
}

func (x *UpdateServiceItemRequest) GetDiscountStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountStartTime
	}
	return nil
}

func (x *UpdateServiceItemRequest) GetDiscountEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountEndTime
	}
	return nil
}

type UpdateServiceItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *ServiceItem `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *UpdateServiceItemResponse) Reset() {
	*x = UpdateServiceItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServiceItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServiceItemResponse) ProtoMessage() {}

func (x *UpdateServiceItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServiceItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateServiceItemResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateServiceItemResponse) GetService() *ServiceItem {
	if x != nil {
		return x.Service
	}
	return nil
}

// delete
type DeleteServiceItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarberShopId string `protobuf:"bytes,2,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
}

func (x *DeleteServiceItemRequest) Reset() {
	*x = DeleteServiceItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceItemRequest) ProtoMessage() {}

func (x *DeleteServiceItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteServiceItemRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteServiceItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteServiceItemRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

type DeleteServiceItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteServiceItemResponse) Reset() {
	*x = DeleteServiceItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_item_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceItemResponse) ProtoMessage() {}

func (x *DeleteServiceItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_item_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteServiceItemResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_item_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteServiceItemResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_barber_rpc_service_item_proto protoreflect.FileDescriptor

var file_barber_rpc_service_item_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x19, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa5, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x46, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3e, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x46, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xc4, 0x05, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x48, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x48, 0x07, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x02, 0x48, 0x08, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x09, 0x52, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x0a,
	0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x16,
	0x0a, 0x14, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x19,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x50, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x5a,
	0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_barber_rpc_service_item_proto_rawDescOnce sync.Once
	file_barber_rpc_service_item_proto_rawDescData = file_barber_rpc_service_item_proto_rawDesc
)

func file_barber_rpc_service_item_proto_rawDescGZIP() []byte {
	file_barber_rpc_service_item_proto_rawDescOnce.Do(func() {
		file_barber_rpc_service_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_service_item_proto_rawDescData)
	})
	return file_barber_rpc_service_item_proto_rawDescData
}

var file_barber_rpc_service_item_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_barber_rpc_service_item_proto_goTypes = []interface{}{
	(*CreateServiceItemRequest)(nil),  // 0: pb.CreateServiceItemRequest
	(*CreateServiceItemResponse)(nil), // 1: pb.CreateServiceItemResponse
	(*GetServiceItemRequest)(nil),     // 2: pb.GetServiceItemRequest
	(*GetServiceItemResponse)(nil),    // 3: pb.GetServiceItemResponse
	(*ListServiceItemRequest)(nil),    // 4: pb.ListServiceItemRequest
	(*ListServiceItemResponse)(nil),   // 5: pb.ListServiceItemResponse
	(*UpdateServiceItemRequest)(nil),  // 6: pb.UpdateServiceItemRequest
	(*UpdateServiceItemResponse)(nil), // 7: pb.UpdateServiceItemResponse
	(*DeleteServiceItemRequest)(nil),  // 8: pb.DeleteServiceItemRequest
	(*DeleteServiceItemResponse)(nil), // 9: pb.DeleteServiceItemResponse
	(*ServiceItem)(nil),               // 10: pb.ServiceItem
	(*timestamppb.Timestamp)(nil),     // 11: google.protobuf.Timestamp
}
var file_barber_rpc_service_item_proto_depIdxs = []int32{
	10, // 0: pb.CreateServiceItemResponse.service:type_name -> pb.ServiceItem
	10, // 1: pb.GetServiceItemResponse.service:type_name -> pb.ServiceItem
	10, // 2: pb.ListServiceItemResponse.services:type_name -> pb.ServiceItem
	11, // 3: pb.UpdateServiceItemRequest.discount_start_time:type_name -> google.protobuf.Timestamp
	11, // 4: pb.UpdateServiceItemRequest.discount_end_time:type_name -> google.protobuf.Timestamp
	10, // 5: pb.UpdateServiceItemResponse.service:type_name -> pb.ServiceItem
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_barber_rpc_service_item_proto_init() }
func file_barber_rpc_service_item_proto_init() {
	if File_barber_rpc_service_item_proto != nil {
		return
	}
	file_barber_service_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_service_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceItemRequest); i {
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
		file_barber_rpc_service_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceItemResponse); i {
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
		file_barber_rpc_service_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceItemRequest); i {
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
		file_barber_rpc_service_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceItemResponse); i {
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
		file_barber_rpc_service_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceItemRequest); i {
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
		file_barber_rpc_service_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceItemResponse); i {
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
		file_barber_rpc_service_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServiceItemRequest); i {
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
		file_barber_rpc_service_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServiceItemResponse); i {
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
		file_barber_rpc_service_item_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceItemRequest); i {
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
		file_barber_rpc_service_item_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceItemResponse); i {
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
	file_barber_rpc_service_item_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_barber_rpc_service_item_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_rpc_service_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_service_item_proto_goTypes,
		DependencyIndexes: file_barber_rpc_service_item_proto_depIdxs,
		MessageInfos:      file_barber_rpc_service_item_proto_msgTypes,
	}.Build()
	File_barber_rpc_service_item_proto = out.File
	file_barber_rpc_service_item_proto_rawDesc = nil
	file_barber_rpc_service_item_proto_goTypes = nil
	file_barber_rpc_service_item_proto_depIdxs = nil
}
