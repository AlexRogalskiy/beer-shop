// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: v1/cart.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetCartReq) Reset() {
	*x = GetCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReq) ProtoMessage() {}

func (x *GetCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReq.ProtoReflect.Descriptor instead.
func (*GetCartReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{0}
}

func (x *GetCartReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetCartReply_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetCartReply) Reset() {
	*x = GetCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReply) ProtoMessage() {}

func (x *GetCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReply.ProtoReflect.Descriptor instead.
func (*GetCartReply) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{1}
}

func (x *GetCartReply) GetItems() []*GetCartReply_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteCartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteCartReq) Reset() {
	*x = DeleteCartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartReq) ProtoMessage() {}

func (x *DeleteCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartReq.ProtoReflect.Descriptor instead.
func (*DeleteCartReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCartReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteCartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCartReply) Reset() {
	*x = DeleteCartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartReply) ProtoMessage() {}

func (x *DeleteCartReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartReply.ProtoReflect.Descriptor instead.
func (*DeleteCartReply) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{3}
}

type AddItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *AddItemReq) Reset() {
	*x = AddItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReq) ProtoMessage() {}

func (x *AddItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemReq.ProtoReflect.Descriptor instead.
func (*AddItemReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{4}
}

func (x *AddItemReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddItemReq) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type AddItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddItemReply) Reset() {
	*x = AddItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReply) ProtoMessage() {}

func (x *AddItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemReply.ProtoReflect.Descriptor instead.
func (*AddItemReply) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{5}
}

func (x *AddItemReply) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *AddItemReply) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId   int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateItemReq) Reset() {
	*x = UpdateItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemReq) ProtoMessage() {}

func (x *UpdateItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemReq.ProtoReflect.Descriptor instead.
func (*UpdateItemReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateItemReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateItemReq) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *UpdateItemReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateItemReply) Reset() {
	*x = UpdateItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemReply) ProtoMessage() {}

func (x *UpdateItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemReply.ProtoReflect.Descriptor instead.
func (*UpdateItemReply) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateItemReply) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *UpdateItemReply) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DeleteItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteItemReq) Reset() {
	*x = DeleteItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemReq) ProtoMessage() {}

func (x *DeleteItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemReq.ProtoReflect.Descriptor instead.
func (*DeleteItemReq) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteItemReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteItemReply) Reset() {
	*x = DeleteItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemReply) ProtoMessage() {}

func (x *DeleteItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemReply.ProtoReflect.Descriptor instead.
func (*DeleteItemReply) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{9}
}

type GetCartReply_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *GetCartReply_Item) Reset() {
	*x = GetCartReply_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartReply_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReply_Item) ProtoMessage() {}

func (x *GetCartReply_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReply_Item.ProtoReflect.Descriptor instead.
func (*GetCartReply_Item) Descriptor() ([]byte, []int) {
	return file_v1_cart_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetCartReply_Item) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GetCartReply_Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_v1_cart_proto protoreflect.FileDescriptor

var file_v1_cart_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x1a, 0x3b, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x28, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3e, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x0c, 0x41, 0x64,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x5d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x46,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x28, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x8e, 0x03, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x47, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x20,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x1a, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_cart_proto_rawDescOnce sync.Once
	file_v1_cart_proto_rawDescData = file_v1_cart_proto_rawDesc
)

func file_v1_cart_proto_rawDescGZIP() []byte {
	file_v1_cart_proto_rawDescOnce.Do(func() {
		file_v1_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_cart_proto_rawDescData)
	})
	return file_v1_cart_proto_rawDescData
}

var file_v1_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_v1_cart_proto_goTypes = []interface{}{
	(*GetCartReq)(nil),        // 0: cart.service.v1.GetCartReq
	(*GetCartReply)(nil),      // 1: cart.service.v1.GetCartReply
	(*DeleteCartReq)(nil),     // 2: cart.service.v1.DeleteCartReq
	(*DeleteCartReply)(nil),   // 3: cart.service.v1.DeleteCartReply
	(*AddItemReq)(nil),        // 4: cart.service.v1.AddItemReq
	(*AddItemReply)(nil),      // 5: cart.service.v1.AddItemReply
	(*UpdateItemReq)(nil),     // 6: cart.service.v1.UpdateItemReq
	(*UpdateItemReply)(nil),   // 7: cart.service.v1.UpdateItemReply
	(*DeleteItemReq)(nil),     // 8: cart.service.v1.DeleteItemReq
	(*DeleteItemReply)(nil),   // 9: cart.service.v1.DeleteItemReply
	(*GetCartReply_Item)(nil), // 10: cart.service.v1.GetCartReply.Item
}
var file_v1_cart_proto_depIdxs = []int32{
	10, // 0: cart.service.v1.GetCartReply.items:type_name -> cart.service.v1.GetCartReply.Item
	0,  // 1: cart.service.v1.Cart.GetCart:input_type -> cart.service.v1.GetCartReq
	2,  // 2: cart.service.v1.Cart.DeleteCart:input_type -> cart.service.v1.DeleteCartReq
	4,  // 3: cart.service.v1.Cart.AddItem:input_type -> cart.service.v1.AddItemReq
	6,  // 4: cart.service.v1.Cart.UpdateItem:input_type -> cart.service.v1.UpdateItemReq
	8,  // 5: cart.service.v1.Cart.DeleteItem:input_type -> cart.service.v1.DeleteItemReq
	1,  // 6: cart.service.v1.Cart.GetCart:output_type -> cart.service.v1.GetCartReply
	3,  // 7: cart.service.v1.Cart.DeleteCart:output_type -> cart.service.v1.DeleteCartReply
	5,  // 8: cart.service.v1.Cart.AddItem:output_type -> cart.service.v1.AddItemReply
	7,  // 9: cart.service.v1.Cart.UpdateItem:output_type -> cart.service.v1.UpdateItemReply
	9,  // 10: cart.service.v1.Cart.DeleteItem:output_type -> cart.service.v1.DeleteItemReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_v1_cart_proto_init() }
func file_v1_cart_proto_init() {
	if File_v1_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReq); i {
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
		file_v1_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReply); i {
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
		file_v1_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartReq); i {
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
		file_v1_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartReply); i {
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
		file_v1_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemReq); i {
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
		file_v1_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemReply); i {
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
		file_v1_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemReq); i {
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
		file_v1_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemReply); i {
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
		file_v1_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemReq); i {
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
		file_v1_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemReply); i {
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
		file_v1_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartReply_Item); i {
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
			RawDescriptor: file_v1_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_cart_proto_goTypes,
		DependencyIndexes: file_v1_cart_proto_depIdxs,
		MessageInfos:      file_v1_cart_proto_msgTypes,
	}.Build()
	File_v1_cart_proto = out.File
	file_v1_cart_proto_rawDesc = nil
	file_v1_cart_proto_goTypes = nil
	file_v1_cart_proto_depIdxs = nil
}
