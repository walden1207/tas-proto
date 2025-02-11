// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: forward.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnswerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn           string `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
	AnswerType   string `protobuf:"bytes,2,opt,name=answerType,proto3" json:"answerType,omitempty"`
	Success      bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Err          string `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	ForwardData  []byte `protobuf:"bytes,5,opt,name=forwardData,proto3" json:"forwardData,omitempty"`
	DeviceLogID  int64  `protobuf:"varint,6,opt,name=deviceLogID,proto3" json:"deviceLogID,omitempty"`
	ResponseData []byte `protobuf:"bytes,7,opt,name=responseData,proto3" json:"responseData,omitempty"`
}

func (x *AnswerResponse) Reset() {
	*x = AnswerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerResponse) ProtoMessage() {}

func (x *AnswerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerResponse.ProtoReflect.Descriptor instead.
func (*AnswerResponse) Descriptor() ([]byte, []int) {
	return file_forward_proto_rawDescGZIP(), []int{0}
}

func (x *AnswerResponse) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *AnswerResponse) GetAnswerType() string {
	if x != nil {
		return x.AnswerType
	}
	return ""
}

func (x *AnswerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AnswerResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *AnswerResponse) GetForwardData() []byte {
	if x != nil {
		return x.ForwardData
	}
	return nil
}

func (x *AnswerResponse) GetDeviceLogID() int64 {
	if x != nil {
		return x.DeviceLogID
	}
	return 0
}

func (x *AnswerResponse) GetResponseData() []byte {
	if x != nil {
		return x.ResponseData
	}
	return nil
}

type ForwardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	DeviceLogID int64  `protobuf:"varint,2,opt,name=deviceLogID,proto3" json:"deviceLogID,omitempty"`
	Value       []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ForwardRequest) Reset() {
	*x = ForwardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardRequest) ProtoMessage() {}

func (x *ForwardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardRequest.ProtoReflect.Descriptor instead.
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return file_forward_proto_rawDescGZIP(), []int{1}
}

func (x *ForwardRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ForwardRequest) GetDeviceLogID() int64 {
	if x != nil {
		return x.DeviceLogID
	}
	return 0
}

func (x *ForwardRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type AnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnswerRequest) Reset() {
	*x = AnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerRequest) ProtoMessage() {}

func (x *AnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerRequest.ProtoReflect.Descriptor instead.
func (*AnswerRequest) Descriptor() ([]byte, []int) {
	return file_forward_proto_rawDescGZIP(), []int{2}
}

var File_forward_proto protoreflect.FileDescriptor

var file_forward_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd4, 0x01, 0x0a, 0x0e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x73, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x44, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0xa5, 0x01, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x12, 0x21, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x06,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x2e,
	0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_forward_proto_rawDescOnce sync.Once
	file_forward_proto_rawDescData = file_forward_proto_rawDesc
)

func file_forward_proto_rawDescGZIP() []byte {
	file_forward_proto_rawDescOnce.Do(func() {
		file_forward_proto_rawDescData = protoimpl.X.CompressGZIP(file_forward_proto_rawDescData)
	})
	return file_forward_proto_rawDescData
}

var file_forward_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_forward_proto_goTypes = []interface{}{
	(*AnswerResponse)(nil), // 0: forward_server.v1.AnswerResponse
	(*ForwardRequest)(nil), // 1: forward_server.v1.ForwardRequest
	(*AnswerRequest)(nil),  // 2: forward_server.v1.AnswerRequest
	(*emptypb.Empty)(nil),  // 3: google.protobuf.Empty
}
var file_forward_proto_depIdxs = []int32{
	1, // 0: forward_server.v1.ForwardService.Forward:input_type -> forward_server.v1.ForwardRequest
	2, // 1: forward_server.v1.ForwardService.Answer:input_type -> forward_server.v1.AnswerRequest
	3, // 2: forward_server.v1.ForwardService.Forward:output_type -> google.protobuf.Empty
	0, // 3: forward_server.v1.ForwardService.Answer:output_type -> forward_server.v1.AnswerResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_forward_proto_init() }
func file_forward_proto_init() {
	if File_forward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_forward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerResponse); i {
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
		file_forward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardRequest); i {
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
		file_forward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerRequest); i {
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
			RawDescriptor: file_forward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_forward_proto_goTypes,
		DependencyIndexes: file_forward_proto_depIdxs,
		MessageInfos:      file_forward_proto_msgTypes,
	}.Build()
	File_forward_proto = out.File
	file_forward_proto_rawDesc = nil
	file_forward_proto_goTypes = nil
	file_forward_proto_depIdxs = nil
}
