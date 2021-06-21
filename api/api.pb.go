// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/api.proto

package api

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TranslationStatusEnum int32

const (
	TranslationStatusEnum_WAITING TranslationStatusEnum = 0
	TranslationStatusEnum_SUCCESS TranslationStatusEnum = 1
	TranslationStatusEnum_ERROR   TranslationStatusEnum = 2
)

// Enum value maps for TranslationStatusEnum.
var (
	TranslationStatusEnum_name = map[int32]string{
		0: "WAITING",
		1: "SUCCESS",
		2: "ERROR",
	}
	TranslationStatusEnum_value = map[string]int32{
		"WAITING": 0,
		"SUCCESS": 1,
		"ERROR":   2,
	}
)

func (x TranslationStatusEnum) Enum() *TranslationStatusEnum {
	p := new(TranslationStatusEnum)
	*p = x
	return p
}

func (x TranslationStatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TranslationStatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_proto_enumTypes[0].Descriptor()
}

func (TranslationStatusEnum) Type() protoreflect.EnumType {
	return &file_api_api_proto_enumTypes[0]
}

func (x TranslationStatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TranslationStatusEnum.Descriptor instead.
func (TranslationStatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

type TranslationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TranslationRequest) Reset() {
	*x = TranslationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationRequest) ProtoMessage() {}

func (x *TranslationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationRequest.ProtoReflect.Descriptor instead.
func (*TranslationRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *TranslationRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *TranslationRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TranslationID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranslationID string `protobuf:"bytes,1,opt,name=translationID,proto3" json:"translationID,omitempty"`
}

func (x *TranslationID) Reset() {
	*x = TranslationID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationID) ProtoMessage() {}

func (x *TranslationID) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationID.ProtoReflect.Descriptor instead.
func (*TranslationID) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *TranslationID) GetTranslationID() string {
	if x != nil {
		return x.TranslationID
	}
	return ""
}

type TranslationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranslationStatus TranslationStatusEnum `protobuf:"varint,1,opt,name=translationStatus,proto3,enum=api.TranslationStatusEnum" json:"translationStatus,omitempty"`
}

func (x *TranslationStatus) Reset() {
	*x = TranslationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationStatus) ProtoMessage() {}

func (x *TranslationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationStatus.ProtoReflect.Descriptor instead.
func (*TranslationStatus) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *TranslationStatus) GetTranslationStatus() TranslationStatusEnum {
	if x != nil {
		return x.TranslationStatus
	}
	return TranslationStatusEnum_WAITING
}

type TranslationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TranslationData) Reset() {
	*x = TranslationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationData) ProtoMessage() {}

func (x *TranslationData) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationData.ProtoReflect.Descriptor instead.
func (*TranslationData) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *TranslationData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x35, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x5d, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x48, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x2a, 0x3c, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41,
	0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32,
	0xbc, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x64,
	0x64, 0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_api_proto_goTypes = []interface{}{
	(TranslationStatusEnum)(0), // 0: api.TranslationStatusEnum
	(*TranslationRequest)(nil), // 1: api.TranslationRequest
	(*TranslationID)(nil),      // 2: api.TranslationID
	(*TranslationStatus)(nil),  // 3: api.TranslationStatus
	(*TranslationData)(nil),    // 4: api.TranslationData
}
var file_api_api_proto_depIdxs = []int32{
	0, // 0: api.TranslationStatus.translationStatus:type_name -> api.TranslationStatusEnum
	1, // 1: api.TranslationService.Translate:input_type -> api.TranslationRequest
	2, // 2: api.TranslationService.GetTranslationStatus:input_type -> api.TranslationID
	2, // 3: api.TranslationService.GetTranslationData:input_type -> api.TranslationID
	2, // 4: api.TranslationService.Translate:output_type -> api.TranslationID
	3, // 5: api.TranslationService.GetTranslationStatus:output_type -> api.TranslationStatus
	4, // 6: api.TranslationService.GetTranslationData:output_type -> api.TranslationData
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationRequest); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationID); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationStatus); i {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationData); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		EnumInfos:         file_api_api_proto_enumTypes,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TranslationServiceClient is the client API for TranslationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TranslationServiceClient interface {
	Translate(ctx context.Context, in *TranslationRequest, opts ...grpc.CallOption) (*TranslationID, error)
	GetTranslationStatus(ctx context.Context, in *TranslationID, opts ...grpc.CallOption) (*TranslationStatus, error)
	GetTranslationData(ctx context.Context, in *TranslationID, opts ...grpc.CallOption) (*TranslationData, error)
}

type translationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslationServiceClient(cc grpc.ClientConnInterface) TranslationServiceClient {
	return &translationServiceClient{cc}
}

func (c *translationServiceClient) Translate(ctx context.Context, in *TranslationRequest, opts ...grpc.CallOption) (*TranslationID, error) {
	out := new(TranslationID)
	err := c.cc.Invoke(ctx, "/api.TranslationService/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationServiceClient) GetTranslationStatus(ctx context.Context, in *TranslationID, opts ...grpc.CallOption) (*TranslationStatus, error) {
	out := new(TranslationStatus)
	err := c.cc.Invoke(ctx, "/api.TranslationService/GetTranslationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationServiceClient) GetTranslationData(ctx context.Context, in *TranslationID, opts ...grpc.CallOption) (*TranslationData, error) {
	out := new(TranslationData)
	err := c.cc.Invoke(ctx, "/api.TranslationService/GetTranslationData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslationServiceServer is the server API for TranslationService service.
type TranslationServiceServer interface {
	Translate(context.Context, *TranslationRequest) (*TranslationID, error)
	GetTranslationStatus(context.Context, *TranslationID) (*TranslationStatus, error)
	GetTranslationData(context.Context, *TranslationID) (*TranslationData, error)
}

// UnimplementedTranslationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTranslationServiceServer struct {
}

func (*UnimplementedTranslationServiceServer) Translate(context.Context, *TranslationRequest) (*TranslationID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (*UnimplementedTranslationServiceServer) GetTranslationStatus(context.Context, *TranslationID) (*TranslationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranslationStatus not implemented")
}
func (*UnimplementedTranslationServiceServer) GetTranslationData(context.Context, *TranslationID) (*TranslationData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranslationData not implemented")
}

func RegisterTranslationServiceServer(s *grpc.Server, srv TranslationServiceServer) {
	s.RegisterService(&_TranslationService_serviceDesc, srv)
}

func _TranslationService_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TranslationService/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).Translate(ctx, req.(*TranslationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationService_GetTranslationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).GetTranslationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TranslationService/GetTranslationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).GetTranslationStatus(ctx, req.(*TranslationID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationService_GetTranslationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslationID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).GetTranslationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TranslationService/GetTranslationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).GetTranslationData(ctx, req.(*TranslationID))
	}
	return interceptor(ctx, in, info, handler)
}

var _TranslationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TranslationService",
	HandlerType: (*TranslationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _TranslationService_Translate_Handler,
		},
		{
			MethodName: "GetTranslationStatus",
			Handler:    _TranslationService_GetTranslationStatus_Handler,
		},
		{
			MethodName: "GetTranslationData",
			Handler:    _TranslationService_GetTranslationData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
