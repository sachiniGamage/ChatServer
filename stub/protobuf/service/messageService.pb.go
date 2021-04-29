// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protobuf/service/messageService.proto

package stub

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

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChatMessageFromServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string       `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   *ChatMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatMessageFromServer) Reset() {
	*x = ChatMessageFromServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageFromServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageFromServer) ProtoMessage() {}

func (x *ChatMessageFromServer) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageFromServer.ProtoReflect.Descriptor instead.
func (*ChatMessageFromServer) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{1}
}

func (x *ChatMessageFromServer) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *ChatMessageFromServer) GetMessage() *ChatMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_protobuf_service_messageService_proto protoreflect.FileDescriptor

var file_protobuf_service_messageService_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x3b, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x65, 0x0a,
	0x15, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x4f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x75,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_service_messageService_proto_rawDescOnce sync.Once
	file_protobuf_service_messageService_proto_rawDescData = file_protobuf_service_messageService_proto_rawDesc
)

func file_protobuf_service_messageService_proto_rawDescGZIP() []byte {
	file_protobuf_service_messageService_proto_rawDescOnce.Do(func() {
		file_protobuf_service_messageService_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_service_messageService_proto_rawDescData)
	})
	return file_protobuf_service_messageService_proto_rawDescData
}

var file_protobuf_service_messageService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_service_messageService_proto_goTypes = []interface{}{
	(*ChatMessage)(nil),           // 0: service.ChatMessage
	(*ChatMessageFromServer)(nil), // 1: service.ChatMessageFromServer
}
var file_protobuf_service_messageService_proto_depIdxs = []int32{
	0, // 0: service.ChatMessageFromServer.message:type_name -> service.ChatMessage
	0, // 1: service.ChatService.Chat:input_type -> service.ChatMessage
	1, // 2: service.ChatService.Chat:output_type -> service.ChatMessageFromServer
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_service_messageService_proto_init() }
func file_protobuf_service_messageService_proto_init() {
	if File_protobuf_service_messageService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_service_messageService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_protobuf_service_messageService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessageFromServer); i {
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
			RawDescriptor: file_protobuf_service_messageService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_service_messageService_proto_goTypes,
		DependencyIndexes: file_protobuf_service_messageService_proto_depIdxs,
		MessageInfos:      file_protobuf_service_messageService_proto_msgTypes,
	}.Build()
	File_protobuf_service_messageService_proto = out.File
	file_protobuf_service_messageService_proto_rawDesc = nil
	file_protobuf_service_messageService_proto_goTypes = nil
	file_protobuf_service_messageService_proto_depIdxs = nil
}