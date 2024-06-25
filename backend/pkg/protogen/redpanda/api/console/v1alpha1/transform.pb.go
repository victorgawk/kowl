// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: redpanda/api/console/v1alpha1/transform.proto

package consolev1alpha1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1alpha1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListTransformsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *v1alpha1.ListTransformsRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *ListTransformsRequest) Reset() {
	*x = ListTransformsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransformsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransformsRequest) ProtoMessage() {}

func (x *ListTransformsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransformsRequest.ProtoReflect.Descriptor instead.
func (*ListTransformsRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_transform_proto_rawDescGZIP(), []int{0}
}

func (x *ListTransformsRequest) GetRequest() *v1alpha1.ListTransformsRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type ListTransformsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *v1alpha1.ListTransformsResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ListTransformsResponse) Reset() {
	*x = ListTransformsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransformsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransformsResponse) ProtoMessage() {}

func (x *ListTransformsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransformsResponse.ProtoReflect.Descriptor instead.
func (*ListTransformsResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_transform_proto_rawDescGZIP(), []int{1}
}

func (x *ListTransformsResponse) GetResponse() *v1alpha1.ListTransformsResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetTransformRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *v1alpha1.GetTransformRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetTransformRequest) Reset() {
	*x = GetTransformRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransformRequest) ProtoMessage() {}

func (x *GetTransformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransformRequest.ProtoReflect.Descriptor instead.
func (*GetTransformRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_transform_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransformRequest) GetRequest() *v1alpha1.GetTransformRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetTransformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *v1alpha1.GetTransformResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetTransformResponse) Reset() {
	*x = GetTransformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransformResponse) ProtoMessage() {}

func (x *GetTransformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransformResponse.ProtoReflect.Descriptor instead.
func (*GetTransformResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_transform_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransformResponse) GetResponse() *v1alpha1.GetTransformResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type DeleteTransformRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *v1alpha1.DeleteTransformRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *DeleteTransformRequest) Reset() {
	*x = DeleteTransformRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTransformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTransformRequest) ProtoMessage() {}

func (x *DeleteTransformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTransformRequest.ProtoReflect.Descriptor instead.
func (*DeleteTransformRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_transform_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTransformRequest) GetRequest() *v1alpha1.DeleteTransformRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type DeleteTransformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *v1alpha1.DeleteTransformResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteTransformResponse) Reset() {
	*x = DeleteTransformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTransformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTransformResponse) ProtoMessage() {}

func (x *DeleteTransformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTransformResponse.ProtoReflect.Descriptor instead.
func (*DeleteTransformResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_transform_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTransformResponse) GetResponse() *v1alpha1.DeleteTransformResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_redpanda_api_console_v1alpha1_transform_proto protoreflect.FileDescriptor

var file_redpanda_api_console_v1alpha1_transform_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2f,
	0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x69, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x72, 0x65, 0x64, 0x70,
	0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x69, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6f, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x93, 0x03, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73,
	0x12, 0x34, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x79, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x32, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x35,
	0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0xaf, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72,
	0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52,
	0x41, 0x43, 0xaa, 0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xe2, 0x02, 0x29, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x20, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_console_v1alpha1_transform_proto_rawDescOnce sync.Once
	file_redpanda_api_console_v1alpha1_transform_proto_rawDescData = file_redpanda_api_console_v1alpha1_transform_proto_rawDesc
)

func file_redpanda_api_console_v1alpha1_transform_proto_rawDescGZIP() []byte {
	file_redpanda_api_console_v1alpha1_transform_proto_rawDescOnce.Do(func() {
		file_redpanda_api_console_v1alpha1_transform_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_console_v1alpha1_transform_proto_rawDescData)
	})
	return file_redpanda_api_console_v1alpha1_transform_proto_rawDescData
}

var file_redpanda_api_console_v1alpha1_transform_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_redpanda_api_console_v1alpha1_transform_proto_goTypes = []interface{}{
	(*ListTransformsRequest)(nil),            // 0: redpanda.api.console.v1alpha1.ListTransformsRequest
	(*ListTransformsResponse)(nil),           // 1: redpanda.api.console.v1alpha1.ListTransformsResponse
	(*GetTransformRequest)(nil),              // 2: redpanda.api.console.v1alpha1.GetTransformRequest
	(*GetTransformResponse)(nil),             // 3: redpanda.api.console.v1alpha1.GetTransformResponse
	(*DeleteTransformRequest)(nil),           // 4: redpanda.api.console.v1alpha1.DeleteTransformRequest
	(*DeleteTransformResponse)(nil),          // 5: redpanda.api.console.v1alpha1.DeleteTransformResponse
	(*v1alpha1.ListTransformsRequest)(nil),   // 6: redpanda.api.dataplane.v1alpha1.ListTransformsRequest
	(*v1alpha1.ListTransformsResponse)(nil),  // 7: redpanda.api.dataplane.v1alpha1.ListTransformsResponse
	(*v1alpha1.GetTransformRequest)(nil),     // 8: redpanda.api.dataplane.v1alpha1.GetTransformRequest
	(*v1alpha1.GetTransformResponse)(nil),    // 9: redpanda.api.dataplane.v1alpha1.GetTransformResponse
	(*v1alpha1.DeleteTransformRequest)(nil),  // 10: redpanda.api.dataplane.v1alpha1.DeleteTransformRequest
	(*v1alpha1.DeleteTransformResponse)(nil), // 11: redpanda.api.dataplane.v1alpha1.DeleteTransformResponse
}
var file_redpanda_api_console_v1alpha1_transform_proto_depIdxs = []int32{
	6,  // 0: redpanda.api.console.v1alpha1.ListTransformsRequest.request:type_name -> redpanda.api.dataplane.v1alpha1.ListTransformsRequest
	7,  // 1: redpanda.api.console.v1alpha1.ListTransformsResponse.response:type_name -> redpanda.api.dataplane.v1alpha1.ListTransformsResponse
	8,  // 2: redpanda.api.console.v1alpha1.GetTransformRequest.request:type_name -> redpanda.api.dataplane.v1alpha1.GetTransformRequest
	9,  // 3: redpanda.api.console.v1alpha1.GetTransformResponse.response:type_name -> redpanda.api.dataplane.v1alpha1.GetTransformResponse
	10, // 4: redpanda.api.console.v1alpha1.DeleteTransformRequest.request:type_name -> redpanda.api.dataplane.v1alpha1.DeleteTransformRequest
	11, // 5: redpanda.api.console.v1alpha1.DeleteTransformResponse.response:type_name -> redpanda.api.dataplane.v1alpha1.DeleteTransformResponse
	0,  // 6: redpanda.api.console.v1alpha1.TransformService.ListTransforms:input_type -> redpanda.api.console.v1alpha1.ListTransformsRequest
	2,  // 7: redpanda.api.console.v1alpha1.TransformService.GetTransform:input_type -> redpanda.api.console.v1alpha1.GetTransformRequest
	4,  // 8: redpanda.api.console.v1alpha1.TransformService.DeleteTransform:input_type -> redpanda.api.console.v1alpha1.DeleteTransformRequest
	1,  // 9: redpanda.api.console.v1alpha1.TransformService.ListTransforms:output_type -> redpanda.api.console.v1alpha1.ListTransformsResponse
	3,  // 10: redpanda.api.console.v1alpha1.TransformService.GetTransform:output_type -> redpanda.api.console.v1alpha1.GetTransformResponse
	5,  // 11: redpanda.api.console.v1alpha1.TransformService.DeleteTransform:output_type -> redpanda.api.console.v1alpha1.DeleteTransformResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_redpanda_api_console_v1alpha1_transform_proto_init() }
func file_redpanda_api_console_v1alpha1_transform_proto_init() {
	if File_redpanda_api_console_v1alpha1_transform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransformsRequest); i {
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
		file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTransformsResponse); i {
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
		file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransformRequest); i {
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
		file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransformResponse); i {
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
		file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTransformRequest); i {
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
		file_redpanda_api_console_v1alpha1_transform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTransformResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_console_v1alpha1_transform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redpanda_api_console_v1alpha1_transform_proto_goTypes,
		DependencyIndexes: file_redpanda_api_console_v1alpha1_transform_proto_depIdxs,
		MessageInfos:      file_redpanda_api_console_v1alpha1_transform_proto_msgTypes,
	}.Build()
	File_redpanda_api_console_v1alpha1_transform_proto = out.File
	file_redpanda_api_console_v1alpha1_transform_proto_rawDesc = nil
	file_redpanda_api_console_v1alpha1_transform_proto_goTypes = nil
	file_redpanda_api_console_v1alpha1_transform_proto_depIdxs = nil
}
