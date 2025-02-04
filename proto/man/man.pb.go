// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: man.proto

package gen

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

type ArticleRequest_SelectType int32

const (
	ArticleRequest_SELECT_FULL ArticleRequest_SelectType = 0
	ArticleRequest_SELECT_HEAD ArticleRequest_SelectType = 1
	ArticleRequest_SELECT_TAIL ArticleRequest_SelectType = 2
)

// Enum value maps for ArticleRequest_SelectType.
var (
	ArticleRequest_SelectType_name = map[int32]string{
		0: "SELECT_FULL",
		1: "SELECT_HEAD",
		2: "SELECT_TAIL",
	}
	ArticleRequest_SelectType_value = map[string]int32{
		"SELECT_FULL": 0,
		"SELECT_HEAD": 1,
		"SELECT_TAIL": 2,
	}
)

func (x ArticleRequest_SelectType) Enum() *ArticleRequest_SelectType {
	p := new(ArticleRequest_SelectType)
	*p = x
	return p
}

func (x ArticleRequest_SelectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArticleRequest_SelectType) Descriptor() protoreflect.EnumDescriptor {
	return file_man_proto_enumTypes[0].Descriptor()
}

func (ArticleRequest_SelectType) Type() protoreflect.EnumType {
	return &file_man_proto_enumTypes[0]
}

func (x ArticleRequest_SelectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArticleRequest_SelectType.Descriptor instead.
func (ArticleRequest_SelectType) EnumDescriptor() ([]byte, []int) {
	return file_man_proto_rawDescGZIP(), []int{3, 0}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardName string `protobuf:"bytes,1,opt,name=board_name,json=boardName,proto3" json:"board_name,omitempty"`
	Path      string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_man_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_man_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_man_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetBoardName() string {
	if x != nil {
		return x.BoardName
	}
	return ""
}

func (x *ListRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool     `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Entries   []*Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ListReply) Reset() {
	*x = ListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_man_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReply) ProtoMessage() {}

func (x *ListReply) ProtoReflect() protoreflect.Message {
	mi := &file_man_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReply.ProtoReflect.Descriptor instead.
func (*ListReply) Descriptor() ([]byte, []int) {
	return file_man_proto_rawDescGZIP(), []int{1}
}

func (x *ListReply) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *ListReply) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardName string `protobuf:"bytes,1,opt,name=board_name,json=boardName,proto3" json:"board_name,omitempty"`
	Path      string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IsDir     bool   `protobuf:"varint,4,opt,name=is_dir,json=isDir,proto3" json:"is_dir,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_man_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_man_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_man_proto_rawDescGZIP(), []int{2}
}

func (x *Entry) GetBoardName() string {
	if x != nil {
		return x.BoardName
	}
	return ""
}

func (x *Entry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Entry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Entry) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

type ArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardName  string                    `protobuf:"bytes,1,opt,name=board_name,json=boardName,proto3" json:"board_name,omitempty"`
	Path       string                    `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	SelectType ArticleRequest_SelectType `protobuf:"varint,3,opt,name=select_type,json=selectType,proto3,enum=pttbbs.man.ArticleRequest_SelectType" json:"select_type,omitempty"`
	CacheKey   string                    `protobuf:"bytes,4,opt,name=cache_key,json=cacheKey,proto3" json:"cache_key,omitempty"`
	Offset     int64                     `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	MaxLength  int64                     `protobuf:"varint,6,opt,name=max_length,json=maxLength,proto3" json:"max_length,omitempty"`
}

func (x *ArticleRequest) Reset() {
	*x = ArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_man_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleRequest) ProtoMessage() {}

func (x *ArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_man_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleRequest.ProtoReflect.Descriptor instead.
func (*ArticleRequest) Descriptor() ([]byte, []int) {
	return file_man_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleRequest) GetBoardName() string {
	if x != nil {
		return x.BoardName
	}
	return ""
}

func (x *ArticleRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ArticleRequest) GetSelectType() ArticleRequest_SelectType {
	if x != nil {
		return x.SelectType
	}
	return ArticleRequest_SELECT_FULL
}

func (x *ArticleRequest) GetCacheKey() string {
	if x != nil {
		return x.CacheKey
	}
	return ""
}

func (x *ArticleRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ArticleRequest) GetMaxLength() int64 {
	if x != nil {
		return x.MaxLength
	}
	return 0
}

type ArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content        []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	CacheKey       string `protobuf:"bytes,2,opt,name=cache_key,json=cacheKey,proto3" json:"cache_key,omitempty"`
	FileSize       int64  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	SelectedOffset int64  `protobuf:"varint,4,opt,name=selected_offset,json=selectedOffset,proto3" json:"selected_offset,omitempty"`
	SelectedSize   int64  `protobuf:"varint,5,opt,name=selected_size,json=selectedSize,proto3" json:"selected_size,omitempty"`
}

func (x *ArticleReply) Reset() {
	*x = ArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_man_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleReply) ProtoMessage() {}

func (x *ArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_man_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleReply.ProtoReflect.Descriptor instead.
func (*ArticleReply) Descriptor() ([]byte, []int) {
	return file_man_proto_rawDescGZIP(), []int{4}
}

func (x *ArticleReply) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ArticleReply) GetCacheKey() string {
	if x != nil {
		return x.CacheKey
	}
	return ""
}

func (x *ArticleReply) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *ArticleReply) GetSelectedOffset() int64 {
	if x != nil {
		return x.SelectedOffset
	}
	return 0
}

func (x *ArticleReply) GetSelectedSize() int64 {
	if x != nil {
		return x.SelectedSize
	}
	return 0
}

var File_man_proto protoreflect.FileDescriptor

var file_man_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x74, 0x74,
	0x62, 0x62, 0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x22, 0x40, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x57, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x74, 0x74, 0x62, 0x62, 0x73, 0x2e,
	0x6d, 0x61, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x67, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x22, 0xa0, 0x02, 0x0a, 0x0e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x46, 0x0a, 0x0b, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x70, 0x74, 0x74, 0x62, 0x62, 0x73, 0x2e,
	0x6d, 0x61, 0x6e, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x3f, 0x0a,
	0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x22, 0xb0,
	0x01, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a,
	0x65, 0x32, 0x89, 0x01, 0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x74, 0x74, 0x62, 0x62,
	0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x74, 0x74, 0x62, 0x62, 0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x74, 0x74, 0x62, 0x62, 0x73, 0x2e, 0x6d,
	0x61, 0x6e, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x74, 0x74, 0x62, 0x62, 0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_man_proto_rawDescOnce sync.Once
	file_man_proto_rawDescData = file_man_proto_rawDesc
)

func file_man_proto_rawDescGZIP() []byte {
	file_man_proto_rawDescOnce.Do(func() {
		file_man_proto_rawDescData = protoimpl.X.CompressGZIP(file_man_proto_rawDescData)
	})
	return file_man_proto_rawDescData
}

var file_man_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_man_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_man_proto_goTypes = []interface{}{
	(ArticleRequest_SelectType)(0), // 0: pttbbs.man.ArticleRequest.SelectType
	(*ListRequest)(nil),            // 1: pttbbs.man.ListRequest
	(*ListReply)(nil),              // 2: pttbbs.man.ListReply
	(*Entry)(nil),                  // 3: pttbbs.man.Entry
	(*ArticleRequest)(nil),         // 4: pttbbs.man.ArticleRequest
	(*ArticleReply)(nil),           // 5: pttbbs.man.ArticleReply
}
var file_man_proto_depIdxs = []int32{
	3, // 0: pttbbs.man.ListReply.entries:type_name -> pttbbs.man.Entry
	0, // 1: pttbbs.man.ArticleRequest.select_type:type_name -> pttbbs.man.ArticleRequest.SelectType
	1, // 2: pttbbs.man.ManService.List:input_type -> pttbbs.man.ListRequest
	4, // 3: pttbbs.man.ManService.Article:input_type -> pttbbs.man.ArticleRequest
	2, // 4: pttbbs.man.ManService.List:output_type -> pttbbs.man.ListReply
	5, // 5: pttbbs.man.ManService.Article:output_type -> pttbbs.man.ArticleReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_man_proto_init() }
func file_man_proto_init() {
	if File_man_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_man_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_man_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReply); i {
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
		file_man_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_man_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleRequest); i {
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
		file_man_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleReply); i {
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
			RawDescriptor: file_man_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_man_proto_goTypes,
		DependencyIndexes: file_man_proto_depIdxs,
		EnumInfos:         file_man_proto_enumTypes,
		MessageInfos:      file_man_proto_msgTypes,
	}.Build()
	File_man_proto = out.File
	file_man_proto_rawDesc = nil
	file_man_proto_goTypes = nil
	file_man_proto_depIdxs = nil
}
