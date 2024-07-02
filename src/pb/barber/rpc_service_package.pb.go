// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_service_package.proto

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

type CreateServicePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShopId string   `protobuf:"bytes,1,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GenderId     int32    `protobuf:"varint,3,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Timer        int32    `protobuf:"varint,4,opt,name=timer,proto3" json:"timer,omitempty"`
	Price        float32  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Description  *string  `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	ImageUrl     *string  `protobuf:"bytes,7,opt,name=image_url,json=imageUrl,proto3,oneof" json:"image_url,omitempty"`
	ServiceItems []string `protobuf:"bytes,8,rep,name=service_items,json=serviceItems,proto3" json:"service_items,omitempty"`
}

func (x *CreateServicePackageRequest) Reset() {
	*x = CreateServicePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServicePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServicePackageRequest) ProtoMessage() {}

func (x *CreateServicePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServicePackageRequest.ProtoReflect.Descriptor instead.
func (*CreateServicePackageRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServicePackageRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *CreateServicePackageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServicePackageRequest) GetGenderId() int32 {
	if x != nil {
		return x.GenderId
	}
	return 0
}

func (x *CreateServicePackageRequest) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *CreateServicePackageRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateServicePackageRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreateServicePackageRequest) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

func (x *CreateServicePackageRequest) GetServiceItems() []string {
	if x != nil {
		return x.ServiceItems
	}
	return nil
}

type CreateServicePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServicePackage *ServicePackage `protobuf:"bytes,1,opt,name=service_package,json=servicePackage,proto3" json:"service_package,omitempty"`
}

func (x *CreateServicePackageResponse) Reset() {
	*x = CreateServicePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServicePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServicePackageResponse) ProtoMessage() {}

func (x *CreateServicePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServicePackageResponse.ProtoReflect.Descriptor instead.
func (*CreateServicePackageResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServicePackageResponse) GetServicePackage() *ServicePackage {
	if x != nil {
		return x.ServicePackage
	}
	return nil
}

// detail
type GetServicePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetServicePackageRequest) Reset() {
	*x = GetServicePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicePackageRequest) ProtoMessage() {}

func (x *GetServicePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicePackageRequest.ProtoReflect.Descriptor instead.
func (*GetServicePackageRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{2}
}

func (x *GetServicePackageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetServicePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServicePackageItem *ServicePackageItem `protobuf:"bytes,1,opt,name=service_package_item,json=servicePackageItem,proto3" json:"service_package_item,omitempty"`
}

func (x *GetServicePackageResponse) Reset() {
	*x = GetServicePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicePackageResponse) ProtoMessage() {}

func (x *GetServicePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicePackageResponse.ProtoReflect.Descriptor instead.
func (*GetServicePackageResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{3}
}

func (x *GetServicePackageResponse) GetServicePackageItem() *ServicePackageItem {
	if x != nil {
		return x.ServicePackageItem
	}
	return nil
}

// list
type ListServicePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShopId string `protobuf:"bytes,1,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
}

func (x *ListServicePackageRequest) Reset() {
	*x = ListServicePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicePackageRequest) ProtoMessage() {}

func (x *ListServicePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicePackageRequest.ProtoReflect.Descriptor instead.
func (*ListServicePackageRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{4}
}

func (x *ListServicePackageRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

type ListServicePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServicePackages []*ServicePackage `protobuf:"bytes,1,rep,name=service_packages,json=servicePackages,proto3" json:"service_packages,omitempty"`
}

func (x *ListServicePackageResponse) Reset() {
	*x = ListServicePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServicePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServicePackageResponse) ProtoMessage() {}

func (x *ListServicePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServicePackageResponse.ProtoReflect.Descriptor instead.
func (*ListServicePackageResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{5}
}

func (x *ListServicePackageResponse) GetServicePackages() []*ServicePackage {
	if x != nil {
		return x.ServicePackages
	}
	return nil
}

type UpdateServicePackageRequest struct {
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
	DiscountPrice     *float32               `protobuf:"fixed32,10,opt,name=discount_price,json=discountPrice,proto3,oneof" json:"discount_price,omitempty"`
	DiscountStartTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=discount_start_time,json=discountStartTime,proto3,oneof" json:"discount_start_time,omitempty"`
	DiscountEndTime   *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=discount_end_time,json=discountEndTime,proto3,oneof" json:"discount_end_time,omitempty"`
	ServiceItems      []string               `protobuf:"bytes,13,rep,name=service_items,json=serviceItems,proto3" json:"service_items,omitempty"`
}

func (x *UpdateServicePackageRequest) Reset() {
	*x = UpdateServicePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServicePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServicePackageRequest) ProtoMessage() {}

func (x *UpdateServicePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServicePackageRequest.ProtoReflect.Descriptor instead.
func (*UpdateServicePackageRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateServicePackageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateServicePackageRequest) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *UpdateServicePackageRequest) GetGenderId() int32 {
	if x != nil && x.GenderId != nil {
		return *x.GenderId
	}
	return 0
}

func (x *UpdateServicePackageRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateServicePackageRequest) GetTimer() int32 {
	if x != nil && x.Timer != nil {
		return *x.Timer
	}
	return 0
}

func (x *UpdateServicePackageRequest) GetPrice() float32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *UpdateServicePackageRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateServicePackageRequest) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

func (x *UpdateServicePackageRequest) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

func (x *UpdateServicePackageRequest) GetDiscountPrice() float32 {
	if x != nil && x.DiscountPrice != nil {
		return *x.DiscountPrice
	}
	return 0
}

func (x *UpdateServicePackageRequest) GetDiscountStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountStartTime
	}
	return nil
}

func (x *UpdateServicePackageRequest) GetDiscountEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountEndTime
	}
	return nil
}

func (x *UpdateServicePackageRequest) GetServiceItems() []string {
	if x != nil {
		return x.ServiceItems
	}
	return nil
}

type UpdateServicePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BarberShopId      string                 `protobuf:"bytes,2,opt,name=barber_shop_id,json=barberShopId,proto3" json:"barber_shop_id,omitempty"`
	GenderId          int32                  `protobuf:"varint,3,opt,name=gender_id,json=genderId,proto3" json:"gender_id,omitempty"`
	Name              string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Timer             int32                  `protobuf:"varint,5,opt,name=timer,proto3" json:"timer,omitempty"`
	Price             float32                `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Description       string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl          string                 `protobuf:"bytes,8,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	IsActive          bool                   `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	DiscountPrice     float32                `protobuf:"fixed32,10,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	DiscountStartTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=discount_start_time,json=discountStartTime,proto3,oneof" json:"discount_start_time,omitempty"`
	DiscountEndTime   *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=discount_end_time,json=discountEndTime,proto3,oneof" json:"discount_end_time,omitempty"`
	ServiceItems      []string               `protobuf:"bytes,13,rep,name=service_items,json=serviceItems,proto3" json:"service_items,omitempty"`
}

func (x *UpdateServicePackageResponse) Reset() {
	*x = UpdateServicePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_service_package_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServicePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServicePackageResponse) ProtoMessage() {}

func (x *UpdateServicePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_service_package_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServicePackageResponse.ProtoReflect.Descriptor instead.
func (*UpdateServicePackageResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_service_package_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateServicePackageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateServicePackageResponse) GetBarberShopId() string {
	if x != nil {
		return x.BarberShopId
	}
	return ""
}

func (x *UpdateServicePackageResponse) GetGenderId() int32 {
	if x != nil {
		return x.GenderId
	}
	return 0
}

func (x *UpdateServicePackageResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateServicePackageResponse) GetTimer() int32 {
	if x != nil {
		return x.Timer
	}
	return 0
}

func (x *UpdateServicePackageResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateServicePackageResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateServicePackageResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *UpdateServicePackageResponse) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *UpdateServicePackageResponse) GetDiscountPrice() float32 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *UpdateServicePackageResponse) GetDiscountStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountStartTime
	}
	return nil
}

func (x *UpdateServicePackageResponse) GetDiscountEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DiscountEndTime
	}
	return nil
}

func (x *UpdateServicePackageResponse) GetServiceItems() []string {
	if x != nil {
		return x.ServiceItems
	}
	return nil
}

var File_barber_rpc_service_package_proto protoreflect.FileDescriptor

var file_barber_rpc_service_package_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x22, 0x5b, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x14, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x41, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x22, 0xb6, 0x05, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73,
	0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x48,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x48, 0x07,
	0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x4f, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x08, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x09, 0x52, 0x0f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42,
	0x16, 0x0a, 0x14, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xa5, 0x04,
	0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68,
	0x6f, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52,
	0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42,
	0x14, 0x0a, 0x12, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_service_package_proto_rawDescOnce sync.Once
	file_barber_rpc_service_package_proto_rawDescData = file_barber_rpc_service_package_proto_rawDesc
)

func file_barber_rpc_service_package_proto_rawDescGZIP() []byte {
	file_barber_rpc_service_package_proto_rawDescOnce.Do(func() {
		file_barber_rpc_service_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_service_package_proto_rawDescData)
	})
	return file_barber_rpc_service_package_proto_rawDescData
}

var file_barber_rpc_service_package_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_barber_rpc_service_package_proto_goTypes = []interface{}{
	(*CreateServicePackageRequest)(nil),  // 0: pb.CreateServicePackageRequest
	(*CreateServicePackageResponse)(nil), // 1: pb.CreateServicePackageResponse
	(*GetServicePackageRequest)(nil),     // 2: pb.GetServicePackageRequest
	(*GetServicePackageResponse)(nil),    // 3: pb.GetServicePackageResponse
	(*ListServicePackageRequest)(nil),    // 4: pb.ListServicePackageRequest
	(*ListServicePackageResponse)(nil),   // 5: pb.ListServicePackageResponse
	(*UpdateServicePackageRequest)(nil),  // 6: pb.UpdateServicePackageRequest
	(*UpdateServicePackageResponse)(nil), // 7: pb.UpdateServicePackageResponse
	(*ServicePackage)(nil),               // 8: pb.ServicePackage
	(*ServicePackageItem)(nil),           // 9: pb.ServicePackageItem
	(*timestamppb.Timestamp)(nil),        // 10: google.protobuf.Timestamp
}
var file_barber_rpc_service_package_proto_depIdxs = []int32{
	8,  // 0: pb.CreateServicePackageResponse.service_package:type_name -> pb.ServicePackage
	9,  // 1: pb.GetServicePackageResponse.service_package_item:type_name -> pb.ServicePackageItem
	8,  // 2: pb.ListServicePackageResponse.service_packages:type_name -> pb.ServicePackage
	10, // 3: pb.UpdateServicePackageRequest.discount_start_time:type_name -> google.protobuf.Timestamp
	10, // 4: pb.UpdateServicePackageRequest.discount_end_time:type_name -> google.protobuf.Timestamp
	10, // 5: pb.UpdateServicePackageResponse.discount_start_time:type_name -> google.protobuf.Timestamp
	10, // 6: pb.UpdateServicePackageResponse.discount_end_time:type_name -> google.protobuf.Timestamp
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_barber_rpc_service_package_proto_init() }
func file_barber_rpc_service_package_proto_init() {
	if File_barber_rpc_service_package_proto != nil {
		return
	}
	file_barber_service_package_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_service_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServicePackageRequest); i {
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
		file_barber_rpc_service_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServicePackageResponse); i {
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
		file_barber_rpc_service_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicePackageRequest); i {
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
		file_barber_rpc_service_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicePackageResponse); i {
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
		file_barber_rpc_service_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicePackageRequest); i {
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
		file_barber_rpc_service_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServicePackageResponse); i {
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
		file_barber_rpc_service_package_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServicePackageRequest); i {
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
		file_barber_rpc_service_package_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServicePackageResponse); i {
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
	file_barber_rpc_service_package_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_barber_rpc_service_package_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_barber_rpc_service_package_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_rpc_service_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_service_package_proto_goTypes,
		DependencyIndexes: file_barber_rpc_service_package_proto_depIdxs,
		MessageInfos:      file_barber_rpc_service_package_proto_msgTypes,
	}.Build()
	File_barber_rpc_service_package_proto = out.File
	file_barber_rpc_service_package_proto_rawDesc = nil
	file_barber_rpc_service_package_proto_goTypes = nil
	file_barber_rpc_service_package_proto_depIdxs = nil
}
