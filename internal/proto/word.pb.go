// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: internal/proto/word.proto

package proto

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

// *
// 単語の型
type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID  string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Word    string `protobuf:"bytes,3,opt,name=word,proto3" json:"word,omitempty"`
	Meaning string `protobuf:"bytes,4,opt,name=meaning,proto3" json:"meaning,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{0}
}

func (x *Word) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Word) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Word) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *Word) GetMeaning() string {
	if x != nil {
		return x.Meaning
	}
	return ""
}

// *
// 単語を作成する際のリクエスト
type CreateWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Word    string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	Meaning string `protobuf:"bytes,3,opt,name=meaning,proto3" json:"meaning,omitempty"`
}

func (x *CreateWordRequest) Reset() {
	*x = CreateWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWordRequest) ProtoMessage() {}

func (x *CreateWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWordRequest.ProtoReflect.Descriptor instead.
func (*CreateWordRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWordRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateWordRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *CreateWordRequest) GetMeaning() string {
	if x != nil {
		return x.Meaning
	}
	return ""
}

// *
// 単語の一覧を取得する際のリクエスト
type ListWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *ListWordRequest) Reset() {
	*x = ListWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWordRequest) ProtoMessage() {}

func (x *ListWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWordRequest.ProtoReflect.Descriptor instead.
func (*ListWordRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{2}
}

func (x *ListWordRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

// *
// 単語を取得する際のリクエスト
type GetWordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetWordRequest) Reset() {
	*x = GetWordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordRequest) ProtoMessage() {}

func (x *GetWordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordRequest.ProtoReflect.Descriptor instead.
func (*GetWordRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{3}
}

func (x *GetWordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetWordRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

// *
// 指定したIDの単語を更新する際のリクエスト
type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Meaning string `protobuf:"bytes,2,opt,name=meaning,proto3" json:"meaning,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetMeaning() string {
	if x != nil {
		return x.Meaning
	}
	return ""
}

// *
// 指定したIDの単語を削除する際のリクエスト
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// *
// 指定したユーザIDの単語一覧を返す
type ListWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words []*Word `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *ListWordResponse) Reset() {
	*x = ListWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWordResponse) ProtoMessage() {}

func (x *ListWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWordResponse.ProtoReflect.Descriptor instead.
func (*ListWordResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{6}
}

func (x *ListWordResponse) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

type GetWordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word *Word `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *GetWordResponse) Reset() {
	*x = GetWordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_word_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWordResponse) ProtoMessage() {}

func (x *GetWordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_word_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWordResponse.ProtoReflect.Descriptor instead.
func (*GetWordResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_word_proto_rawDescGZIP(), []int{7}
}

func (x *GetWordResponse) GetWord() *Word {
	if x != nil {
		return x.Word
	}
	return nil
}

var File_internal_proto_word_proto protoreflect.FileDescriptor

var file_internal_proto_word_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c,
	0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x59, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x29, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x61, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x77, 0x6f,
	0x72, 0x64, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x31,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x32, 0xb5, 0x02, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12,
	0x17, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x39, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x77,
	0x6f, 0x72, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x77,
	0x6f, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x64, 0x12, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x13, 0x2e, 0x77,
	0x6f, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_word_proto_rawDescOnce sync.Once
	file_internal_proto_word_proto_rawDescData = file_internal_proto_word_proto_rawDesc
)

func file_internal_proto_word_proto_rawDescGZIP() []byte {
	file_internal_proto_word_proto_rawDescOnce.Do(func() {
		file_internal_proto_word_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_word_proto_rawDescData)
	})
	return file_internal_proto_word_proto_rawDescData
}

var file_internal_proto_word_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_proto_word_proto_goTypes = []interface{}{
	(*Word)(nil),              // 0: word.Word
	(*CreateWordRequest)(nil), // 1: word.CreateWordRequest
	(*ListWordRequest)(nil),   // 2: word.ListWordRequest
	(*GetWordRequest)(nil),    // 3: word.GetWordRequest
	(*UpdateRequest)(nil),     // 4: word.UpdateRequest
	(*DeleteRequest)(nil),     // 5: word.DeleteRequest
	(*ListWordResponse)(nil),  // 6: word.ListWordResponse
	(*GetWordResponse)(nil),   // 7: word.GetWordResponse
	(*emptypb.Empty)(nil),     // 8: google.protobuf.Empty
}
var file_internal_proto_word_proto_depIdxs = []int32{
	0, // 0: word.ListWordResponse.words:type_name -> word.Word
	0, // 1: word.GetWordResponse.word:type_name -> word.Word
	1, // 2: word.WordService.CreateWord:input_type -> word.CreateWordRequest
	2, // 3: word.WordService.ListWord:input_type -> word.ListWordRequest
	3, // 4: word.WordService.GetWord:input_type -> word.GetWordRequest
	4, // 5: word.WordService.UpdateWord:input_type -> word.UpdateRequest
	5, // 6: word.WordService.DeleteWord:input_type -> word.DeleteRequest
	8, // 7: word.WordService.CreateWord:output_type -> google.protobuf.Empty
	6, // 8: word.WordService.ListWord:output_type -> word.ListWordResponse
	7, // 9: word.WordService.GetWord:output_type -> word.GetWordResponse
	8, // 10: word.WordService.UpdateWord:output_type -> google.protobuf.Empty
	8, // 11: word.WordService.DeleteWord:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_word_proto_init() }
func file_internal_proto_word_proto_init() {
	if File_internal_proto_word_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_word_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
		file_internal_proto_word_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWordRequest); i {
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
		file_internal_proto_word_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWordRequest); i {
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
		file_internal_proto_word_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordRequest); i {
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
		file_internal_proto_word_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_internal_proto_word_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_internal_proto_word_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWordResponse); i {
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
		file_internal_proto_word_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWordResponse); i {
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
			RawDescriptor: file_internal_proto_word_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_word_proto_goTypes,
		DependencyIndexes: file_internal_proto_word_proto_depIdxs,
		MessageInfos:      file_internal_proto_word_proto_msgTypes,
	}.Build()
	File_internal_proto_word_proto = out.File
	file_internal_proto_word_proto_rawDesc = nil
	file_internal_proto_word_proto_goTypes = nil
	file_internal_proto_word_proto_depIdxs = nil
}
