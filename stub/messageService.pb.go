// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: protobuf/service/messageService.proto

package stub

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

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	To      string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Chatid  int32  `protobuf:"varint,4,opt,name=chatid,proto3" json:"chatid,omitempty"`
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

func (x *ChatMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ChatMessage) GetChatid() int32 {
	if x != nil {
		return x.Chatid
	}
	return 0
}

type ChatMessageFromServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string       `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   *ChatMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Chatid    *ChatMessage `protobuf:"bytes,3,opt,name=chatid,proto3" json:"chatid,omitempty"`
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

func (x *ChatMessageFromServer) GetChatid() *ChatMessage {
	if x != nil {
		return x.Chatid
	}
	return nil
}

type ChatMessageToServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string       `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   *ChatMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	To        *ChatMessage `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Chatid    *ChatMessage `protobuf:"bytes,4,opt,name=chatid,proto3" json:"chatid,omitempty"`
}

func (x *ChatMessageToServer) Reset() {
	*x = ChatMessageToServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessageToServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessageToServer) ProtoMessage() {}

func (x *ChatMessageToServer) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessageToServer.ProtoReflect.Descriptor instead.
func (*ChatMessageToServer) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{2}
}

func (x *ChatMessageToServer) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *ChatMessageToServer) GetMessage() *ChatMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ChatMessageToServer) GetTo() *ChatMessage {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *ChatMessageToServer) GetChatid() *ChatMessage {
	if x != nil {
		return x.Chatid
	}
	return nil
}

type RegisterUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	PublicKey string `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *RegisterUser) Reset() {
	*x = RegisterUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUser) ProtoMessage() {}

func (x *RegisterUser) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUser.ProtoReflect.Descriptor instead.
func (*RegisterUser) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterUser) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type LoginUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *LoginUser) Reset() {
	*x = LoginUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUser) ProtoMessage() {}

func (x *LoginUser) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUser.ProtoReflect.Descriptor instead.
func (*LoginUser) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{4}
}

func (x *LoginUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{5}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Edit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Edit) Reset() {
	*x = Edit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edit) ProtoMessage() {}

func (x *Edit) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edit.ProtoReflect.Descriptor instead.
func (*Edit) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{6}
}

func (x *Edit) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type FriendList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendsEmail string        `protobuf:"bytes,1,opt,name=FriendsEmail,proto3" json:"FriendsEmail,omitempty"`
	Email        *RegisterUser `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username     *RegisterUser `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	PublicKey    string        `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (x *FriendList) Reset() {
	*x = FriendList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendList) ProtoMessage() {}

func (x *FriendList) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendList.ProtoReflect.Descriptor instead.
func (*FriendList) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{7}
}

func (x *FriendList) GetFriendsEmail() string {
	if x != nil {
		return x.FriendsEmail
	}
	return ""
}

func (x *FriendList) GetEmail() *RegisterUser {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *FriendList) GetUsername() *RegisterUser {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *FriendList) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type AddFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Myemail string      `protobuf:"bytes,1,opt,name=myemail,proto3" json:"myemail,omitempty"`
	Detail  *FriendList `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *AddFriendReq) Reset() {
	*x = AddFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendReq) ProtoMessage() {}

func (x *AddFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendReq.ProtoReflect.Descriptor instead.
func (*AddFriendReq) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{8}
}

func (x *AddFriendReq) GetMyemail() string {
	if x != nil {
		return x.Myemail
	}
	return ""
}

func (x *AddFriendReq) GetDetail() *FriendList {
	if x != nil {
		return x.Detail
	}
	return nil
}

type ViewFriends struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Myemail           string          `protobuf:"bytes,1,opt,name=myemail,proto3" json:"myemail,omitempty"`
	FriendsInList     []*RegisterUser `protobuf:"bytes,2,rep,name=friendsInList,proto3" json:"friendsInList,omitempty"`
	FriendsNameInList []*RegisterUser `protobuf:"bytes,3,rep,name=friendsNameInList,proto3" json:"friendsNameInList,omitempty"`
}

func (x *ViewFriends) Reset() {
	*x = ViewFriends{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_service_messageService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewFriends) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewFriends) ProtoMessage() {}

func (x *ViewFriends) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_messageService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewFriends.ProtoReflect.Descriptor instead.
func (*ViewFriends) Descriptor() ([]byte, []int) {
	return file_protobuf_service_messageService_proto_rawDescGZIP(), []int{9}
}

func (x *ViewFriends) GetMyemail() string {
	if x != nil {
		return x.Myemail
	}
	return ""
}

func (x *ViewFriends) GetFriendsInList() []*RegisterUser {
	if x != nil {
		return x.FriendsInList
	}
	return nil
}

func (x *ViewFriends) GetFriendsNameInList() []*RegisterUser {
	if x != nil {
		return x.FriendsNameInList
	}
	return nil
}

var File_protobuf_service_messageService_proto protoreflect.FileDescriptor

var file_protobuf_service_messageService_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x69, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2e,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x69, 0x64, 0x22, 0x7a, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x59,
	0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xae, 0x01, 0x0a,
	0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x2b, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x55, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x79, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x79, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x79, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x79, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3b,
	0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0d, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x11, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x11, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x32, 0x4f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x28, 0x01, 0x30,
	0x01, 0x32, 0x7a, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x2b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb5, 0x01,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x39, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x38, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x1a,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x42, 0x5e, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x74, 0x75, 0x62, 0x42, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x69, 0x47,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x73, 0x74, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protobuf_service_messageService_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protobuf_service_messageService_proto_goTypes = []interface{}{
	(*ChatMessage)(nil),           // 0: service.ChatMessage
	(*ChatMessageFromServer)(nil), // 1: service.ChatMessageFromServer
	(*ChatMessageToServer)(nil),   // 2: service.ChatMessageToServer
	(*RegisterUser)(nil),          // 3: service.RegisterUser
	(*LoginUser)(nil),             // 4: service.LoginUser
	(*Token)(nil),                 // 5: service.Token
	(*Edit)(nil),                  // 6: service.Edit
	(*FriendList)(nil),            // 7: service.FriendList
	(*AddFriendReq)(nil),          // 8: service.AddFriendReq
	(*ViewFriends)(nil),           // 9: service.ViewFriends
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_protobuf_service_messageService_proto_depIdxs = []int32{
	0,  // 0: service.ChatMessageFromServer.message:type_name -> service.ChatMessage
	0,  // 1: service.ChatMessageFromServer.chatid:type_name -> service.ChatMessage
	0,  // 2: service.ChatMessageToServer.message:type_name -> service.ChatMessage
	0,  // 3: service.ChatMessageToServer.to:type_name -> service.ChatMessage
	0,  // 4: service.ChatMessageToServer.chatid:type_name -> service.ChatMessage
	3,  // 5: service.FriendList.email:type_name -> service.RegisterUser
	3,  // 6: service.FriendList.username:type_name -> service.RegisterUser
	7,  // 7: service.AddFriendReq.detail:type_name -> service.FriendList
	3,  // 8: service.ViewFriends.friendsInList:type_name -> service.RegisterUser
	3,  // 9: service.ViewFriends.friendsNameInList:type_name -> service.RegisterUser
	0,  // 10: service.ChatService.Chat:input_type -> service.ChatMessage
	3,  // 11: service.AuthenticateUser.Register:input_type -> service.RegisterUser
	4,  // 12: service.AuthenticateUser.Login:input_type -> service.LoginUser
	6,  // 13: service.UpdateUser.UpdateName:input_type -> service.Edit
	8,  // 14: service.UpdateUser.AddFriend:input_type -> service.AddFriendReq
	9,  // 15: service.UpdateUser.GetFriends:input_type -> service.ViewFriends
	1,  // 16: service.ChatService.Chat:output_type -> service.ChatMessageFromServer
	10, // 17: service.AuthenticateUser.Register:output_type -> google.protobuf.Empty
	5,  // 18: service.AuthenticateUser.Login:output_type -> service.Token
	3,  // 19: service.UpdateUser.UpdateName:output_type -> service.RegisterUser
	8,  // 20: service.UpdateUser.AddFriend:output_type -> service.AddFriendReq
	9,  // 21: service.UpdateUser.GetFriends:output_type -> service.ViewFriends
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
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
		file_protobuf_service_messageService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessageToServer); i {
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
		file_protobuf_service_messageService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUser); i {
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
		file_protobuf_service_messageService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUser); i {
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
		file_protobuf_service_messageService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_protobuf_service_messageService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edit); i {
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
		file_protobuf_service_messageService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendList); i {
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
		file_protobuf_service_messageService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendReq); i {
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
		file_protobuf_service_messageService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewFriends); i {
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
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   3,
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
