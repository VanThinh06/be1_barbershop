// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: barber/service_barber_shop.proto

package barber

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_barber_service_barber_shop_proto protoreflect.FileDescriptor

var file_barber_service_barber_shop_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe6, 0x0c, 0x0a, 0x0a, 0x42, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x98, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x36, 0x12, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x1a,
	0x23, 0x41, 0x50, 0x49, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x74, 0x6f,
	0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x62, 0x61,
	0x72, 0x62, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x12, 0xa4, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x64, 0x92, 0x41, 0x3f, 0x12, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x1a, 0x26, 0x41,
	0x50, 0x49, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x12, 0xb3, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x12, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x92,
	0x41, 0x3f, 0x12, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x20, 0x61, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x1a, 0x26, 0x41, 0x50, 0x49, 0x20, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x62, 0x61, 0x72, 0x62,
	0x65, 0x72, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xa7, 0x01,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x12, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x64, 0x92, 0x41, 0x3f, 0x12, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x1a, 0x26, 0x41,
	0x50, 0x49, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x32, 0x17,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x5f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x12, 0xc5, 0x01, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x62,
	0x79, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x61, 0x72, 0x62, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x92, 0x41, 0x3e, 0x12, 0x17, 0x66, 0x69, 0x6e,
	0x64, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x20, 0x73, 0x68, 0x6f, 0x70, 0x20, 0x6e, 0x65,
	0x61, 0x72, 0x62, 0x79, 0x1a, 0x23, 0x41, 0x50, 0x49, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e,
	0x65, 0x77, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a,
	0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x12,
	0xcc, 0x01, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x71, 0x92, 0x41, 0x40,
	0x12, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x20, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x23, 0x41, 0x50, 0x49,
	0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xaf,
	0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41, 0x40, 0x12, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x23, 0x41, 0x50, 0x49, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20,
	0x6e, 0x65, 0x77, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0xb1, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68,
	0x6f, 0x70, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x72,
	0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x62, 0x92, 0x41, 0x36, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20,
	0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x1a, 0x23, 0x41, 0x50, 0x49, 0x20, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x2d,
	0x73, 0x68, 0x6f, 0x70, 0x12, 0xb8, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x12,
	0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x92, 0x41, 0x40, 0x12, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x23, 0x41, 0x50, 0x49, 0x20, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x42,
	0x59, 0x92, 0x41, 0x3c, 0x12, 0x3a, 0x0a, 0x0f, 0x42, 0x61, 0x72, 0x62, 0x65, 0x72, 0x20, 0x53,
	0x68, 0x6f, 0x70, 0x20, 0x41, 0x50, 0x49, 0x22, 0x22, 0x0a, 0x09, 0x54, 0x68, 0x69, 0x6e, 0x68,
	0x20, 0x4b, 0x61, 0x69, 0x1a, 0x15, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x68, 0x69, 0x6e, 0x68, 0x30,
	0x36, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30,
	0x5a, 0x18, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_barber_service_barber_shop_proto_goTypes = []interface{}{
	(*CreateBarberRequest)(nil),           // 0: pb.CreateBarberRequest
	(*LoginBarberRequest)(nil),            // 1: pb.LoginBarberRequest
	(*RefreshTokenBarberRequest)(nil),     // 2: pb.RefreshTokenBarberRequest
	(*UpdateBarberRequest)(nil),           // 3: pb.UpdateBarberRequest
	(*FindBarberShopsNearbyRequest)(nil),  // 4: pb.FindBarberShopsNearbyRequest
	(*CreateServiceCategoryRequest)(nil),  // 5: pb.CreateServiceCategoryRequest
	(*CreateServicesRequest)(nil),         // 6: pb.CreateServicesRequest
	(*CreateBarberShopsRequest)(nil),      // 7: pb.CreateBarberShopsRequest
	(*CodeBarberShopRequest)(nil),         // 8: pb.CodeBarberShopRequest
	(*CreateBarberResponse)(nil),          // 9: pb.CreateBarberResponse
	(*LoginBarberResponse)(nil),           // 10: pb.LoginBarberResponse
	(*RefreshTokenBarberResponse)(nil),    // 11: pb.RefreshTokenBarberResponse
	(*UpdateBarberResponse)(nil),          // 12: pb.UpdateBarberResponse
	(*FindBarberShopsNearbyResponse)(nil), // 13: pb.FindBarberShopsNearbyResponse
	(*CreateServiceCategoryResponse)(nil), // 14: pb.CreateServiceCategoryResponse
	(*CreateServicesResponse)(nil),        // 15: pb.CreateServicesResponse
	(*CreateBarberShopsResponse)(nil),     // 16: pb.CreateBarberShopsResponse
	(*CodeBarberShopResponse)(nil),        // 17: pb.CodeBarberShopResponse
}
var file_barber_service_barber_shop_proto_depIdxs = []int32{
	0,  // 0: pb.BarberShop.CreateBarber:input_type -> pb.CreateBarberRequest
	1,  // 1: pb.BarberShop.LoginBarber:input_type -> pb.LoginBarberRequest
	2,  // 2: pb.BarberShop.RefreshTokenBarber:input_type -> pb.RefreshTokenBarberRequest
	3,  // 3: pb.BarberShop.UpdateBarber:input_type -> pb.UpdateBarberRequest
	4,  // 4: pb.BarberShop.FindBarberShopsNearby:input_type -> pb.FindBarberShopsNearbyRequest
	5,  // 5: pb.BarberShop.NewServiceCategory:input_type -> pb.CreateServiceCategoryRequest
	6,  // 6: pb.BarberShop.NewServices:input_type -> pb.CreateServicesRequest
	7,  // 7: pb.BarberShop.NewBarberShops:input_type -> pb.CreateBarberShopsRequest
	8,  // 8: pb.BarberShop.GenerateCodeBarberShop:input_type -> pb.CodeBarberShopRequest
	9,  // 9: pb.BarberShop.CreateBarber:output_type -> pb.CreateBarberResponse
	10, // 10: pb.BarberShop.LoginBarber:output_type -> pb.LoginBarberResponse
	11, // 11: pb.BarberShop.RefreshTokenBarber:output_type -> pb.RefreshTokenBarberResponse
	12, // 12: pb.BarberShop.UpdateBarber:output_type -> pb.UpdateBarberResponse
	13, // 13: pb.BarberShop.FindBarberShopsNearby:output_type -> pb.FindBarberShopsNearbyResponse
	14, // 14: pb.BarberShop.NewServiceCategory:output_type -> pb.CreateServiceCategoryResponse
	15, // 15: pb.BarberShop.NewServices:output_type -> pb.CreateServicesResponse
	16, // 16: pb.BarberShop.NewBarberShops:output_type -> pb.CreateBarberShopsResponse
	17, // 17: pb.BarberShop.GenerateCodeBarberShop:output_type -> pb.CodeBarberShopResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_barber_service_barber_shop_proto_init() }
func file_barber_service_barber_shop_proto_init() {
	if File_barber_service_barber_shop_proto != nil {
		return
	}
	file_barber_rpc_create_barber_proto_init()
	file_barber_rpc_login_barber_proto_init()
	file_barber_rpc_refresh_token_barber_proto_init()
	file_barber_rpc_update_barber_proto_init()
	file_barber_rpc_create_barbershop_proto_init()
	file_barber_rpc_create_service_category_proto_init()
	file_barber_rpc_create_services_proto_init()
	file_barber_rpc_find_barber_shop_nearby_proto_init()
	file_barber_rpc_code_barbershop_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barber_service_barber_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_barber_service_barber_shop_proto_goTypes,
		DependencyIndexes: file_barber_service_barber_shop_proto_depIdxs,
	}.Build()
	File_barber_service_barber_shop_proto = out.File
	file_barber_service_barber_shop_proto_rawDesc = nil
	file_barber_service_barber_shop_proto_goTypes = nil
	file_barber_service_barber_shop_proto_depIdxs = nil
}
