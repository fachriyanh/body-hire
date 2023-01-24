// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: proto/publisher.proto

package publisher

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

type CreatePublisherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreatePublisherRequest) Reset() {
	*x = CreatePublisherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePublisherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePublisherRequest) ProtoMessage() {}

func (x *CreatePublisherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePublisherRequest.ProtoReflect.Descriptor instead.
func (*CreatePublisherRequest) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePublisherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatePublisherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreatePublisherResponse) Reset() {
	*x = CreatePublisherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePublisherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePublisherResponse) ProtoMessage() {}

func (x *CreatePublisherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePublisherResponse.ProtoReflect.Descriptor instead.
func (*CreatePublisherResponse) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePublisherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReadPublisherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadPublisherRequest) Reset() {
	*x = ReadPublisherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPublisherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPublisherRequest) ProtoMessage() {}

func (x *ReadPublisherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPublisherRequest.ProtoReflect.Descriptor instead.
func (*ReadPublisherRequest) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{2}
}

func (x *ReadPublisherRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReadPublisherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReadPublisherResponse) Reset() {
	*x = ReadPublisherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPublisherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPublisherResponse) ProtoMessage() {}

func (x *ReadPublisherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPublisherResponse.ProtoReflect.Descriptor instead.
func (*ReadPublisherResponse) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPublisherResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdatePublisherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdatePublisherRequest) Reset() {
	*x = UpdatePublisherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePublisherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePublisherRequest) ProtoMessage() {}

func (x *UpdatePublisherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePublisherRequest.ProtoReflect.Descriptor instead.
func (*UpdatePublisherRequest) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePublisherRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePublisherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdatePublisherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdatePublisherResponse) Reset() {
	*x = UpdatePublisherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePublisherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePublisherResponse) ProtoMessage() {}

func (x *UpdatePublisherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePublisherResponse.ProtoReflect.Descriptor instead.
func (*UpdatePublisherResponse) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePublisherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeletePublisherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePublisherRequest) Reset() {
	*x = DeletePublisherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePublisherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePublisherRequest) ProtoMessage() {}

func (x *DeletePublisherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePublisherRequest.ProtoReflect.Descriptor instead.
func (*DeletePublisherRequest) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePublisherRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePublisherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeletePublisherResponse) Reset() {
	*x = DeletePublisherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publisher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePublisherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePublisherResponse) ProtoMessage() {}

func (x *DeletePublisherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publisher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePublisherResponse.ProtoReflect.Descriptor instead.
func (*DeletePublisherResponse) Descriptor() ([]byte, []int) {
	return file_proto_publisher_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePublisherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_publisher_proto protoreflect.FileDescriptor

var file_proto_publisher_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2b, 0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3c, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xac, 0x02, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_publisher_proto_rawDescOnce sync.Once
	file_proto_publisher_proto_rawDescData = file_proto_publisher_proto_rawDesc
)

func file_proto_publisher_proto_rawDescGZIP() []byte {
	file_proto_publisher_proto_rawDescOnce.Do(func() {
		file_proto_publisher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_publisher_proto_rawDescData)
	})
	return file_proto_publisher_proto_rawDescData
}

var file_proto_publisher_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_publisher_proto_goTypes = []interface{}{
	(*CreatePublisherRequest)(nil),  // 0: CreatePublisherRequest
	(*CreatePublisherResponse)(nil), // 1: CreatePublisherResponse
	(*ReadPublisherRequest)(nil),    // 2: ReadPublisherRequest
	(*ReadPublisherResponse)(nil),   // 3: ReadPublisherResponse
	(*UpdatePublisherRequest)(nil),  // 4: UpdatePublisherRequest
	(*UpdatePublisherResponse)(nil), // 5: UpdatePublisherResponse
	(*DeletePublisherRequest)(nil),  // 6: DeletePublisherRequest
	(*DeletePublisherResponse)(nil), // 7: DeletePublisherResponse
}
var file_proto_publisher_proto_depIdxs = []int32{
	0, // 0: PublisherService.CreatePublisher:input_type -> CreatePublisherRequest
	2, // 1: PublisherService.ReadPublisher:input_type -> ReadPublisherRequest
	4, // 2: PublisherService.UpdatePublisher:input_type -> UpdatePublisherRequest
	6, // 3: PublisherService.DeletePublisher:input_type -> DeletePublisherRequest
	1, // 4: PublisherService.CreatePublisher:output_type -> CreatePublisherResponse
	3, // 5: PublisherService.ReadPublisher:output_type -> ReadPublisherResponse
	5, // 6: PublisherService.UpdatePublisher:output_type -> UpdatePublisherResponse
	7, // 7: PublisherService.DeletePublisher:output_type -> DeletePublisherResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_publisher_proto_init() }
func file_proto_publisher_proto_init() {
	if File_proto_publisher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_publisher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePublisherRequest); i {
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
		file_proto_publisher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePublisherResponse); i {
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
		file_proto_publisher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPublisherRequest); i {
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
		file_proto_publisher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPublisherResponse); i {
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
		file_proto_publisher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePublisherRequest); i {
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
		file_proto_publisher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePublisherResponse); i {
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
		file_proto_publisher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePublisherRequest); i {
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
		file_proto_publisher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePublisherResponse); i {
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
			RawDescriptor: file_proto_publisher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_publisher_proto_goTypes,
		DependencyIndexes: file_proto_publisher_proto_depIdxs,
		MessageInfos:      file_proto_publisher_proto_msgTypes,
	}.Build()
	File_proto_publisher_proto = out.File
	file_proto_publisher_proto_rawDesc = nil
	file_proto_publisher_proto_goTypes = nil
	file_proto_publisher_proto_depIdxs = nil
}