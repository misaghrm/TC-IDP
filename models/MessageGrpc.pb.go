// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: MessageGrpc.proto

package models

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GrpcMessageType int32

const (
	GrpcMessageType_MessageTypeUndefined GrpcMessageType = 0
	GrpcMessageType_Otp                  GrpcMessageType = 1
	GrpcMessageType_General              GrpcMessageType = 2
)

// Enum value maps for GrpcMessageType.
var (
	GrpcMessageType_name = map[int32]string{
		0: "MessageTypeUndefined",
		1: "Otp",
		2: "General",
	}
	GrpcMessageType_value = map[string]int32{
		"MessageTypeUndefined": 0,
		"Otp":                  1,
		"General":              2,
	}
)

func (x GrpcMessageType) Enum() *GrpcMessageType {
	p := new(GrpcMessageType)
	*p = x
	return p
}

func (x GrpcMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrpcMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_MessageGrpc_proto_enumTypes[0].Descriptor()
}

func (GrpcMessageType) Type() protoreflect.EnumType {
	return &file_MessageGrpc_proto_enumTypes[0]
}

func (x GrpcMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrpcMessageType.Descriptor instead.
func (GrpcMessageType) EnumDescriptor() ([]byte, []int) {
	return file_MessageGrpc_proto_rawDescGZIP(), []int{0}
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipient  string              `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Type       GrpcMessageType     `protobuf:"varint,2,opt,name=type,proto3,enum=models.GrpcMessageType" json:"type,omitempty"`
	Parameters []*ParamaterMessage `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageGrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MessageGrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_MessageGrpc_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *SendMessageRequest) GetType() GrpcMessageType {
	if x != nil {
		return x.Type
	}
	return GrpcMessageType_MessageTypeUndefined
}

func (x *SendMessageRequest) GetParameters() []*ParamaterMessage {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type ParamaterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ParamaterMessage) Reset() {
	*x = ParamaterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageGrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamaterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamaterMessage) ProtoMessage() {}

func (x *ParamaterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_MessageGrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamaterMessage.ProtoReflect.Descriptor instead.
func (*ParamaterMessage) Descriptor() ([]byte, []int) {
	return file_MessageGrpc_proto_rawDescGZIP(), []int{1}
}

func (x *ParamaterMessage) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ParamaterMessage) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok           bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageGrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MessageGrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_MessageGrpc_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *SendMessageResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type CreateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTemplateRequest) Reset() {
	*x = CreateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageGrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateRequest) ProtoMessage() {}

func (x *CreateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MessageGrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_MessageGrpc_proto_rawDescGZIP(), []int{3}
}

type CreateTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTemplateResponse) Reset() {
	*x = CreateTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MessageGrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateResponse) ProtoMessage() {}

func (x *CreateTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MessageGrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateResponse.ProtoReflect.Descriptor instead.
func (*CreateTemplateResponse) Descriptor() ([]byte, []int) {
	return file_MessageGrpc_proto_rawDescGZIP(), []int{4}
}

var File_MessageGrpc_proto protoreflect.FileDescriptor

var file_MessageGrpc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x12,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x17,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2a, 0x41, 0x0a, 0x0f, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x4f, 0x74, 0x70, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x10, 0x02, 0x32, 0xa6, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x47, 0x72, 0x70, 0x63, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a,
	0x09, 0x2e, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02, 0x17, 0x54, 0x63, 0x2e,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x43, 0x67, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MessageGrpc_proto_rawDescOnce sync.Once
	file_MessageGrpc_proto_rawDescData = file_MessageGrpc_proto_rawDesc
)

func file_MessageGrpc_proto_rawDescGZIP() []byte {
	file_MessageGrpc_proto_rawDescOnce.Do(func() {
		file_MessageGrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_MessageGrpc_proto_rawDescData)
	})
	return file_MessageGrpc_proto_rawDescData
}

var file_MessageGrpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MessageGrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_MessageGrpc_proto_goTypes = []interface{}{
	(GrpcMessageType)(0),           // 0: models.GrpcMessageType
	(*SendMessageRequest)(nil),     // 1: models.SendMessageRequest
	(*ParamaterMessage)(nil),       // 2: models.ParamaterMessage
	(*SendMessageResponse)(nil),    // 3: models.SendMessageResponse
	(*CreateTemplateRequest)(nil),  // 4: models.CreateTemplateRequest
	(*CreateTemplateResponse)(nil), // 5: models.CreateTemplateResponse
}
var file_MessageGrpc_proto_depIdxs = []int32{
	0, // 0: models.SendMessageRequest.type:type_name -> models.GrpcMessageType
	2, // 1: models.SendMessageRequest.parameters:type_name -> models.ParamaterMessage
	1, // 2: models.MessageGrpc.SendMessage:input_type -> models.SendMessageRequest
	4, // 3: models.MessageGrpc.CreateTemplate:input_type -> models.CreateTemplateRequest
	3, // 4: models.MessageGrpc.SendMessage:output_type -> models.SendMessageResponse
	5, // 5: models.MessageGrpc.CreateTemplate:output_type -> models.CreateTemplateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MessageGrpc_proto_init() }
func file_MessageGrpc_proto_init() {
	if File_MessageGrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MessageGrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_MessageGrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamaterMessage); i {
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
		file_MessageGrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_MessageGrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateRequest); i {
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
		file_MessageGrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateResponse); i {
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
			RawDescriptor: file_MessageGrpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MessageGrpc_proto_goTypes,
		DependencyIndexes: file_MessageGrpc_proto_depIdxs,
		EnumInfos:         file_MessageGrpc_proto_enumTypes,
		MessageInfos:      file_MessageGrpc_proto_msgTypes,
	}.Build()
	File_MessageGrpc_proto = out.File
	file_MessageGrpc_proto_rawDesc = nil
	file_MessageGrpc_proto_goTypes = nil
	file_MessageGrpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageGrpcClient is the client API for MessageGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageGrpcClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
}

type messageGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageGrpcClient(cc grpc.ClientConnInterface) MessageGrpcClient {
	return &messageGrpcClient{cc}
}

func (c *messageGrpcClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/models.MessageGrpc/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageGrpcClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	out := new(CreateTemplateResponse)
	err := c.cc.Invoke(ctx, "/models.MessageGrpc/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageGrpcServer is the server API for MessageGrpc service.
type MessageGrpcServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
}

// UnimplementedMessageGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedMessageGrpcServer struct {
}

func (*UnimplementedMessageGrpcServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedMessageGrpcServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}

func RegisterMessageGrpcServer(s *grpc.Server, srv MessageGrpcServer) {
	s.RegisterService(&_MessageGrpc_serviceDesc, srv)
}

func _MessageGrpc_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGrpcServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.MessageGrpc/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGrpcServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageGrpc_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGrpcServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.MessageGrpc/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGrpcServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.MessageGrpc",
	HandlerType: (*MessageGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessageGrpc_SendMessage_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _MessageGrpc_CreateTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MessageGrpc.proto",
}
