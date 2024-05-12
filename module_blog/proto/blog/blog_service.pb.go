// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc2
// source: blog_service.proto

package blog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Blog_SystemStatus int32

const (
	Blog_DRAFT     Blog_SystemStatus = 0
	Blog_PUBLISHED Blog_SystemStatus = 1
	Blog_CLOSED    Blog_SystemStatus = 2
)

// Enum value maps for Blog_SystemStatus.
var (
	Blog_SystemStatus_name = map[int32]string{
		0: "DRAFT",
		1: "PUBLISHED",
		2: "CLOSED",
	}
	Blog_SystemStatus_value = map[string]int32{
		"DRAFT":     0,
		"PUBLISHED": 1,
		"CLOSED":    2,
	}
)

func (x Blog_SystemStatus) Enum() *Blog_SystemStatus {
	p := new(Blog_SystemStatus)
	*p = x
	return p
}

func (x Blog_SystemStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Blog_SystemStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_blog_service_proto_enumTypes[0].Descriptor()
}

func (Blog_SystemStatus) Type() protoreflect.EnumType {
	return &file_blog_service_proto_enumTypes[0]
}

func (x Blog_SystemStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Blog_SystemStatus.Descriptor instead.
func (Blog_SystemStatus) EnumDescriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{4, 0}
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{1}
}

type BlogStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlogId int32  `protobuf:"varint,2,opt,name=blogId,proto3" json:"blogId,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BlogStatus) Reset() {
	*x = BlogStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogStatus) ProtoMessage() {}

func (x *BlogStatus) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogStatus.ProtoReflect.Descriptor instead.
func (*BlogStatus) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{2}
}

func (x *BlogStatus) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlogStatus) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *BlogStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BlogRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlogId       int32                  `protobuf:"varint,2,opt,name=blogId,proto3" json:"blogId,omitempty"`
	UserId       int32                  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Rating       string                 `protobuf:"bytes,4,opt,name=rating,proto3" json:"rating,omitempty"`
	CreationTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
}

func (x *BlogRating) Reset() {
	*x = BlogRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogRating) ProtoMessage() {}

func (x *BlogRating) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogRating.ProtoReflect.Descriptor instead.
func (*BlogRating) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{3}
}

func (x *BlogRating) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlogRating) GetBlogId() int32 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *BlogRating) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BlogRating) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *BlogRating) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kita         int64                  `protobuf:"varint,1,opt,name=kita,proto3" json:"kita,omitempty"`
	CreatorId    int64                  `protobuf:"varint,2,opt,name=creatorId,proto3" json:"creatorId,omitempty"`
	Title        string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description  string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	SystemStatus Blog_SystemStatus      `protobuf:"varint,5,opt,name=systemStatus,proto3,enum=blog.Blog_SystemStatus" json:"systemStatus,omitempty"`
	ImageLinks   string                 `protobuf:"bytes,6,opt,name=imageLinks,proto3" json:"imageLinks,omitempty"`
	CreationDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	BlogStatuses []*BlogStatus          `protobuf:"bytes,8,rep,name=blogStatuses,proto3" json:"blogStatuses,omitempty"`
	BlogRatings  []*BlogRating          `protobuf:"bytes,9,rep,name=blogRatings,proto3" json:"blogRatings,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{4}
}

func (x *Blog) GetKita() int64 {
	if x != nil {
		return x.Kita
	}
	return 0
}

func (x *Blog) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Blog) GetSystemStatus() Blog_SystemStatus {
	if x != nil {
		return x.SystemStatus
	}
	return Blog_DRAFT
}

func (x *Blog) GetImageLinks() string {
	if x != nil {
		return x.ImageLinks
	}
	return ""
}

func (x *Blog) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *Blog) GetBlogStatuses() []*BlogStatus {
	if x != nil {
		return x.BlogStatuses
	}
	return nil
}

func (x *Blog) GetBlogRatings() []*BlogRating {
	if x != nil {
		return x.BlogRatings
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{5}
}

type BlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *BlogResponse) Reset() {
	*x = BlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogResponse) ProtoMessage() {}

func (x *BlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogResponse.ProtoReflect.Descriptor instead.
func (*BlogResponse) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{6}
}

func (x *BlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type MultiBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog []*Blog `protobuf:"bytes,1,rep,name=blog,proto3" json:"blog,omitempty"`
}

func (x *MultiBlogResponse) Reset() {
	*x = MultiBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiBlogResponse) ProtoMessage() {}

func (x *MultiBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiBlogResponse.ProtoReflect.Descriptor instead.
func (*MultiBlogResponse) Descriptor() ([]byte, []int) {
	return file_blog_service_proto_rawDescGZIP(), []int{7}
}

func (x *MultiBlogResponse) GetBlog() []*Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

var File_blog_service_proto protoreflect.FileDescriptor

var file_blog_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x48, 0x0a, 0x0a, 0x42, 0x6c,
	0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xad, 0x03, 0x0a, 0x04,
	0x42, 0x6c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x6b, 0x69, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x62,
	0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x34, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x02, 0x22, 0x0f, 0x0a, 0x0d, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x0c,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x33, 0x0a, 0x11,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f,
	0x67, 0x32, 0xc0, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x29, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x08, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x08, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x64,
	0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x64, 0x1a, 0x13, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x62,
	0x6c, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_service_proto_rawDescOnce sync.Once
	file_blog_service_proto_rawDescData = file_blog_service_proto_rawDesc
)

func file_blog_service_proto_rawDescGZIP() []byte {
	file_blog_service_proto_rawDescOnce.Do(func() {
		file_blog_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_service_proto_rawDescData)
	})
	return file_blog_service_proto_rawDescData
}

var file_blog_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blog_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_blog_service_proto_goTypes = []interface{}{
	(Blog_SystemStatus)(0),        // 0: blog.Blog.SystemStatus
	(*Id)(nil),                    // 1: blog.Id
	(*Empty)(nil),                 // 2: blog.Empty
	(*BlogStatus)(nil),            // 3: blog.BlogStatus
	(*BlogRating)(nil),            // 4: blog.BlogRating
	(*Blog)(nil),                  // 5: blog.Blog
	(*EmptyResponse)(nil),         // 6: blog.EmptyResponse
	(*BlogResponse)(nil),          // 7: blog.BlogResponse
	(*MultiBlogResponse)(nil),     // 8: blog.MultiBlogResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_blog_service_proto_depIdxs = []int32{
	9,  // 0: blog.BlogRating.creationTime:type_name -> google.protobuf.Timestamp
	0,  // 1: blog.Blog.systemStatus:type_name -> blog.Blog.SystemStatus
	9,  // 2: blog.Blog.creationDate:type_name -> google.protobuf.Timestamp
	3,  // 3: blog.Blog.blogStatuses:type_name -> blog.BlogStatus
	4,  // 4: blog.Blog.blogRatings:type_name -> blog.BlogRating
	5,  // 5: blog.BlogResponse.blog:type_name -> blog.Blog
	5,  // 6: blog.MultiBlogResponse.blog:type_name -> blog.Blog
	1,  // 7: blog.BlogService.GetById:input_type -> blog.Id
	1,  // 8: blog.BlogService.GetAll:input_type -> blog.Id
	2,  // 9: blog.BlogService.Create:input_type -> blog.Empty
	1,  // 10: blog.BlogService.Delete:input_type -> blog.Id
	7,  // 11: blog.BlogService.GetById:output_type -> blog.BlogResponse
	8,  // 12: blog.BlogService.GetAll:output_type -> blog.MultiBlogResponse
	6,  // 13: blog.BlogService.Create:output_type -> blog.EmptyResponse
	6,  // 14: blog.BlogService.Delete:output_type -> blog.EmptyResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_blog_service_proto_init() }
func file_blog_service_proto_init() {
	if File_blog_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_blog_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_blog_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogStatus); i {
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
		file_blog_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogRating); i {
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
		file_blog_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_blog_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_blog_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogResponse); i {
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
		file_blog_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiBlogResponse); i {
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
			RawDescriptor: file_blog_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_service_proto_goTypes,
		DependencyIndexes: file_blog_service_proto_depIdxs,
		EnumInfos:         file_blog_service_proto_enumTypes,
		MessageInfos:      file_blog_service_proto_msgTypes,
	}.Build()
	File_blog_service_proto = out.File
	file_blog_service_proto_rawDesc = nil
	file_blog_service_proto_goTypes = nil
	file_blog_service_proto_depIdxs = nil
}