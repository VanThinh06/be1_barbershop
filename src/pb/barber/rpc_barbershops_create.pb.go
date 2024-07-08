// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: barber/rpc_barbershops_create.proto

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

type CreateBarberShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarberShopChainId        *string `protobuf:"bytes,1,opt,name=barber_shop_chain_id,json=barberShopChainId,proto3,oneof" json:"barber_shop_chain_id,omitempty"`
	Name                     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProvinceId               int32   `protobuf:"varint,3,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	DistrictId               int32   `protobuf:"varint,4,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	WardId                   int32   `protobuf:"varint,5,opt,name=ward_id,json=wardId,proto3" json:"ward_id,omitempty"`
	SpecificLocation         string  `protobuf:"bytes,6,opt,name=specific_location,json=specificLocation,proto3" json:"specific_location,omitempty"`
	Phone                    string  `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                    string  `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	WebsiteUrl               *string `protobuf:"bytes,9,opt,name=website_url,json=websiteUrl,proto3,oneof" json:"website_url,omitempty"`
	AvatarUrl                string  `protobuf:"bytes,10,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	CoverPhotoUrl            string  `protobuf:"bytes,11,opt,name=cover_photo_url,json=coverPhotoUrl,proto3" json:"cover_photo_url,omitempty"`
	FacadePhotoUrl           string  `protobuf:"bytes,12,opt,name=facade_photo_url,json=facadePhotoUrl,proto3" json:"facade_photo_url,omitempty"`
	RepresentativeName       string  `protobuf:"bytes,13,opt,name=representative_name,json=representativeName,proto3" json:"representative_name,omitempty"`
	CitizenId                string  `protobuf:"bytes,14,opt,name=citizen_id,json=citizenId,proto3" json:"citizen_id,omitempty"`
	RepresentativePhone      string  `protobuf:"bytes,15,opt,name=representative_phone,json=representativePhone,proto3" json:"representative_phone,omitempty"`
	RepresentativeEmail      string  `protobuf:"bytes,16,opt,name=representative_email,json=representativeEmail,proto3" json:"representative_email,omitempty"`
	RepresentativePhoneOther *string `protobuf:"bytes,17,opt,name=representative_phone_other,json=representativePhoneOther,proto3,oneof" json:"representative_phone_other,omitempty"`
	StartTimeMonday          int64   `protobuf:"varint,18,opt,name=start_time_monday,json=startTimeMonday,proto3" json:"start_time_monday,omitempty"`
	EndTimeMonday            int64   `protobuf:"varint,19,opt,name=end_time_monday,json=endTimeMonday,proto3" json:"end_time_monday,omitempty"`
	StartTimeTuesday         int64   `protobuf:"varint,20,opt,name=start_time_tuesday,json=startTimeTuesday,proto3" json:"start_time_tuesday,omitempty"`
	EndTimeTuesday           int64   `protobuf:"varint,21,opt,name=end_time_tuesday,json=endTimeTuesday,proto3" json:"end_time_tuesday,omitempty"`
	StartTimeWednesday       int64   `protobuf:"varint,22,opt,name=start_time_wednesday,json=startTimeWednesday,proto3" json:"start_time_wednesday,omitempty"`
	EndTimeWednesday         int64   `protobuf:"varint,23,opt,name=end_time_wednesday,json=endTimeWednesday,proto3" json:"end_time_wednesday,omitempty"`
	StartTimeThursday        int64   `protobuf:"varint,24,opt,name=start_time_thursday,json=startTimeThursday,proto3" json:"start_time_thursday,omitempty"`
	EndTimeThursday          int64   `protobuf:"varint,25,opt,name=end_time_thursday,json=endTimeThursday,proto3" json:"end_time_thursday,omitempty"`
	StartTimeFriday          int64   `protobuf:"varint,26,opt,name=start_time_friday,json=startTimeFriday,proto3" json:"start_time_friday,omitempty"`
	EndTimeFriday            int64   `protobuf:"varint,27,opt,name=end_time_friday,json=endTimeFriday,proto3" json:"end_time_friday,omitempty"`
	StartTimeSaturday        int64   `protobuf:"varint,28,opt,name=start_time_saturday,json=startTimeSaturday,proto3" json:"start_time_saturday,omitempty"`
	EndTimeSaturday          int64   `protobuf:"varint,29,opt,name=end_time_saturday,json=endTimeSaturday,proto3" json:"end_time_saturday,omitempty"`
	StartTimeSunday          int64   `protobuf:"varint,30,opt,name=start_time_sunday,json=startTimeSunday,proto3" json:"start_time_sunday,omitempty"`
	EndTimeSunday            int64   `protobuf:"varint,31,opt,name=end_time_sunday,json=endTimeSunday,proto3" json:"end_time_sunday,omitempty"`
	Longitude                float32 `protobuf:"fixed32,32,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude                 float32 `protobuf:"fixed32,33,opt,name=latitude,proto3" json:"latitude,omitempty"`
	IntervalScheduler        int32   `protobuf:"varint,34,opt,name=interval_scheduler,json=intervalScheduler,proto3" json:"interval_scheduler,omitempty"`
}

func (x *CreateBarberShopRequest) Reset() {
	*x = CreateBarberShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberShopRequest) ProtoMessage() {}

func (x *CreateBarberShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberShopRequest.ProtoReflect.Descriptor instead.
func (*CreateBarberShopRequest) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBarberShopRequest) GetBarberShopChainId() string {
	if x != nil && x.BarberShopChainId != nil {
		return *x.BarberShopChainId
	}
	return ""
}

func (x *CreateBarberShopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBarberShopRequest) GetProvinceId() int32 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *CreateBarberShopRequest) GetDistrictId() int32 {
	if x != nil {
		return x.DistrictId
	}
	return 0
}

func (x *CreateBarberShopRequest) GetWardId() int32 {
	if x != nil {
		return x.WardId
	}
	return 0
}

func (x *CreateBarberShopRequest) GetSpecificLocation() string {
	if x != nil {
		return x.SpecificLocation
	}
	return ""
}

func (x *CreateBarberShopRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateBarberShopRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateBarberShopRequest) GetWebsiteUrl() string {
	if x != nil && x.WebsiteUrl != nil {
		return *x.WebsiteUrl
	}
	return ""
}

func (x *CreateBarberShopRequest) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *CreateBarberShopRequest) GetCoverPhotoUrl() string {
	if x != nil {
		return x.CoverPhotoUrl
	}
	return ""
}

func (x *CreateBarberShopRequest) GetFacadePhotoUrl() string {
	if x != nil {
		return x.FacadePhotoUrl
	}
	return ""
}

func (x *CreateBarberShopRequest) GetRepresentativeName() string {
	if x != nil {
		return x.RepresentativeName
	}
	return ""
}

func (x *CreateBarberShopRequest) GetCitizenId() string {
	if x != nil {
		return x.CitizenId
	}
	return ""
}

func (x *CreateBarberShopRequest) GetRepresentativePhone() string {
	if x != nil {
		return x.RepresentativePhone
	}
	return ""
}

func (x *CreateBarberShopRequest) GetRepresentativeEmail() string {
	if x != nil {
		return x.RepresentativeEmail
	}
	return ""
}

func (x *CreateBarberShopRequest) GetRepresentativePhoneOther() string {
	if x != nil && x.RepresentativePhoneOther != nil {
		return *x.RepresentativePhoneOther
	}
	return ""
}

func (x *CreateBarberShopRequest) GetStartTimeMonday() int64 {
	if x != nil {
		return x.StartTimeMonday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetEndTimeMonday() int64 {
	if x != nil {
		return x.EndTimeMonday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetStartTimeTuesday() int64 {
	if x != nil {
		return x.StartTimeTuesday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetEndTimeTuesday() int64 {
	if x != nil {
		return x.EndTimeTuesday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetStartTimeWednesday() int64 {
	if x != nil {
		return x.StartTimeWednesday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetEndTimeWednesday() int64 {
	if x != nil {
		return x.EndTimeWednesday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetStartTimeThursday() int64 {
	if x != nil {
		return x.StartTimeThursday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetEndTimeThursday() int64 {
	if x != nil {
		return x.EndTimeThursday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetStartTimeFriday() int64 {
	if x != nil {
		return x.StartTimeFriday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetEndTimeFriday() int64 {
	if x != nil {
		return x.EndTimeFriday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetStartTimeSaturday() int64 {
	if x != nil {
		return x.StartTimeSaturday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetEndTimeSaturday() int64 {
	if x != nil {
		return x.EndTimeSaturday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetStartTimeSunday() int64 {
	if x != nil {
		return x.StartTimeSunday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetEndTimeSunday() int64 {
	if x != nil {
		return x.EndTimeSunday
	}
	return 0
}

func (x *CreateBarberShopRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *CreateBarberShopRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CreateBarberShopRequest) GetIntervalScheduler() int32 {
	if x != nil {
		return x.IntervalScheduler
	}
	return 0
}

type CreateBarberShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateBarberShopResponse) Reset() {
	*x = CreateBarberShopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_barber_rpc_barbershops_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBarberShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBarberShopResponse) ProtoMessage() {}

func (x *CreateBarberShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barber_rpc_barbershops_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBarberShopResponse.ProtoReflect.Descriptor instead.
func (*CreateBarberShopResponse) Descriptor() ([]byte, []int) {
	return file_barber_rpc_barbershops_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBarberShopResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_barber_rpc_barbershops_create_proto protoreflect.FileDescriptor

var file_barber_rpc_barbershops_create_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xc4, 0x0b, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x14, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c,
	0x12, 0x28, 0x0a, 0x10, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x63, 0x61,
	0x64, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x72, 0x65,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x31, 0x0a,
	0x14, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x41, 0x0a, 0x1a, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x18, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x12,
	0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x64,
	0x61, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x4d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x75,
	0x65, 0x73, 0x64, 0x61, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x74, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79, 0x12,
	0x30, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x65,
	0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61,
	0x79, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77, 0x65,
	0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x12,
	0x2e, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x68,
	0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x12,
	0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x68, 0x75, 0x72,
	0x73, 0x64, 0x61, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x54, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x69, 0x64, 0x61, 0x79,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x46, 0x72, 0x69, 0x64, 0x61, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72, 0x69, 0x64, 0x61, 0x79, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x69, 0x64, 0x61, 0x79, 0x12,
	0x2e, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x61,
	0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x12,
	0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x61, 0x74, 0x75,
	0x72, 0x64, 0x61, 0x79, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x79,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18,
	0x22, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x22, 0x34, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barber_rpc_barbershops_create_proto_rawDescOnce sync.Once
	file_barber_rpc_barbershops_create_proto_rawDescData = file_barber_rpc_barbershops_create_proto_rawDesc
)

func file_barber_rpc_barbershops_create_proto_rawDescGZIP() []byte {
	file_barber_rpc_barbershops_create_proto_rawDescOnce.Do(func() {
		file_barber_rpc_barbershops_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_barber_rpc_barbershops_create_proto_rawDescData)
	})
	return file_barber_rpc_barbershops_create_proto_rawDescData
}

var file_barber_rpc_barbershops_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barber_rpc_barbershops_create_proto_goTypes = []interface{}{
	(*CreateBarberShopRequest)(nil),  // 0: pb.CreateBarberShopRequest
	(*CreateBarberShopResponse)(nil), // 1: pb.CreateBarberShopResponse
}
var file_barber_rpc_barbershops_create_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_barber_rpc_barbershops_create_proto_init() }
func file_barber_rpc_barbershops_create_proto_init() {
	if File_barber_rpc_barbershops_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barber_rpc_barbershops_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberShopRequest); i {
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
		file_barber_rpc_barbershops_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBarberShopResponse); i {
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
	file_barber_rpc_barbershops_create_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_rpc_barbershops_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_barber_rpc_barbershops_create_proto_goTypes,
		DependencyIndexes: file_barber_rpc_barbershops_create_proto_depIdxs,
		MessageInfos:      file_barber_rpc_barbershops_create_proto_msgTypes,
	}.Build()
	File_barber_rpc_barbershops_create_proto = out.File
	file_barber_rpc_barbershops_create_proto_rawDesc = nil
	file_barber_rpc_barbershops_create_proto_goTypes = nil
	file_barber_rpc_barbershops_create_proto_depIdxs = nil
}
