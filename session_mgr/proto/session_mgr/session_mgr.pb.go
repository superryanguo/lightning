// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.3
// source: proto/session_mgr/session_mgr.proto

package micro_super_lightning_service_session_mgr

import (
	proto "github.com/golang/protobuf/proto"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSay() string {
	if x != nil {
		return x.Say
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId   string `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	SessionData []byte `protobuf:"bytes,2,opt,name=SessionData,proto3" json:"SessionData,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Session) GetSessionData() []byte {
	if x != nil {
		return x.SessionData
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	Data   string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *Response) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StreamingRequest) Reset() {
	*x = StreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingRequest) ProtoMessage() {}

func (x *StreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingRequest.ProtoReflect.Descriptor instead.
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{4}
}

func (x *StreamingRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StreamingResponse) Reset() {
	*x = StreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingResponse) ProtoMessage() {}

func (x *StreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingResponse.ProtoReflect.Descriptor instead.
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{5}
}

func (x *StreamingResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{6}
}

func (x *Ping) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_proto_session_mgr_session_mgr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_proto_session_mgr_session_mgr_proto_rawDescGZIP(), []int{7}
}

func (x *Pong) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

var File_proto_session_mgr_session_mgr_proto protoreflect.FileDescriptor

var file_proto_session_mgr_session_mgr_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x67, 0x72, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72,
	0x22, 0x1b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x61, 0x79, 0x22, 0x27, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x4c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72,
	0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x28, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1e, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x6f, 0x6b, 0x65, 0x22, 0x1e, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x6f, 0x6b, 0x65, 0x32, 0xfa, 0x04, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x67, 0x72, 0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x67, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a,
	0x0b, 0x53, 0x61, 0x76, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0x33, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3b,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x72, 0x0a,
	0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x2f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_session_mgr_session_mgr_proto_rawDescOnce sync.Once
	file_proto_session_mgr_session_mgr_proto_rawDescData = file_proto_session_mgr_session_mgr_proto_rawDesc
)

func file_proto_session_mgr_session_mgr_proto_rawDescGZIP() []byte {
	file_proto_session_mgr_session_mgr_proto_rawDescOnce.Do(func() {
		file_proto_session_mgr_session_mgr_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_session_mgr_session_mgr_proto_rawDescData)
	})
	return file_proto_session_mgr_session_mgr_proto_rawDescData
}

var file_proto_session_mgr_session_mgr_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_session_mgr_session_mgr_proto_goTypes = []interface{}{
	(*Message)(nil),           // 0: micro.super.lightning.service.session_mgr.Message
	(*Request)(nil),           // 1: micro.super.lightning.service.session_mgr.Request
	(*Session)(nil),           // 2: micro.super.lightning.service.session_mgr.Session
	(*Response)(nil),          // 3: micro.super.lightning.service.session_mgr.Response
	(*StreamingRequest)(nil),  // 4: micro.super.lightning.service.session_mgr.StreamingRequest
	(*StreamingResponse)(nil), // 5: micro.super.lightning.service.session_mgr.StreamingResponse
	(*Ping)(nil),              // 6: micro.super.lightning.service.session_mgr.Ping
	(*Pong)(nil),              // 7: micro.super.lightning.service.session_mgr.Pong
}
var file_proto_session_mgr_session_mgr_proto_depIdxs = []int32{
	1, // 0: micro.super.lightning.service.session_mgr.Session_mgr.GetSession:input_type -> micro.super.lightning.service.session_mgr.Request
	2, // 1: micro.super.lightning.service.session_mgr.Session_mgr.SaveSession:input_type -> micro.super.lightning.service.session_mgr.Session
	1, // 2: micro.super.lightning.service.session_mgr.Session_mgr.DeleteSession:input_type -> micro.super.lightning.service.session_mgr.Request
	4, // 3: micro.super.lightning.service.session_mgr.Session_mgr.Stream:input_type -> micro.super.lightning.service.session_mgr.StreamingRequest
	6, // 4: micro.super.lightning.service.session_mgr.Session_mgr.PingPong:input_type -> micro.super.lightning.service.session_mgr.Ping
	3, // 5: micro.super.lightning.service.session_mgr.Session_mgr.GetSession:output_type -> micro.super.lightning.service.session_mgr.Response
	3, // 6: micro.super.lightning.service.session_mgr.Session_mgr.SaveSession:output_type -> micro.super.lightning.service.session_mgr.Response
	3, // 7: micro.super.lightning.service.session_mgr.Session_mgr.DeleteSession:output_type -> micro.super.lightning.service.session_mgr.Response
	5, // 8: micro.super.lightning.service.session_mgr.Session_mgr.Stream:output_type -> micro.super.lightning.service.session_mgr.StreamingResponse
	7, // 9: micro.super.lightning.service.session_mgr.Session_mgr.PingPong:output_type -> micro.super.lightning.service.session_mgr.Pong
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_session_mgr_session_mgr_proto_init() }
func file_proto_session_mgr_session_mgr_proto_init() {
	if File_proto_session_mgr_session_mgr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_session_mgr_session_mgr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_session_mgr_session_mgr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_session_mgr_session_mgr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_proto_session_mgr_session_mgr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_session_mgr_session_mgr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingRequest); i {
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
		file_proto_session_mgr_session_mgr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingResponse); i {
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
		file_proto_session_mgr_session_mgr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_proto_session_mgr_session_mgr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_proto_session_mgr_session_mgr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_session_mgr_session_mgr_proto_goTypes,
		DependencyIndexes: file_proto_session_mgr_session_mgr_proto_depIdxs,
		MessageInfos:      file_proto_session_mgr_session_mgr_proto_msgTypes,
	}.Build()
	File_proto_session_mgr_session_mgr_proto = out.File
	file_proto_session_mgr_session_mgr_proto_rawDesc = nil
	file_proto_session_mgr_session_mgr_proto_goTypes = nil
	file_proto_session_mgr_session_mgr_proto_depIdxs = nil
}
