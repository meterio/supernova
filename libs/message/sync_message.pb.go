// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.2
// source: sync_message.proto

package message

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

type NewBlockEnvelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RawBlock      []byte                 `protobuf:"bytes,1,opt,name=rawBlock,proto3" json:"rawBlock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewBlockEnvelope) Reset() {
	*x = NewBlockEnvelope{}
	mi := &file_sync_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewBlockEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBlockEnvelope) ProtoMessage() {}

func (x *NewBlockEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_sync_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBlockEnvelope.ProtoReflect.Descriptor instead.
func (*NewBlockEnvelope) Descriptor() ([]byte, []int) {
	return file_sync_message_proto_rawDescGZIP(), []int{0}
}

func (x *NewBlockEnvelope) GetRawBlock() []byte {
	if x != nil {
		return x.RawBlock
	}
	return nil
}

type NewBlockIDEnvelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlockID       []byte                 `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewBlockIDEnvelope) Reset() {
	*x = NewBlockIDEnvelope{}
	mi := &file_sync_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewBlockIDEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBlockIDEnvelope) ProtoMessage() {}

func (x *NewBlockIDEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_sync_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBlockIDEnvelope.ProtoReflect.Descriptor instead.
func (*NewBlockIDEnvelope) Descriptor() ([]byte, []int) {
	return file_sync_message_proto_rawDescGZIP(), []int{1}
}

func (x *NewBlockIDEnvelope) GetBlockID() []byte {
	if x != nil {
		return x.BlockID
	}
	return nil
}

type NewTxsEnvelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RawTxs        [][]byte               `protobuf:"bytes,1,rep,name=rawTxs,proto3" json:"rawTxs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewTxsEnvelope) Reset() {
	*x = NewTxsEnvelope{}
	mi := &file_sync_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTxsEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTxsEnvelope) ProtoMessage() {}

func (x *NewTxsEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_sync_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTxsEnvelope.ProtoReflect.Descriptor instead.
func (*NewTxsEnvelope) Descriptor() ([]byte, []int) {
	return file_sync_message_proto_rawDescGZIP(), []int{2}
}

func (x *NewTxsEnvelope) GetRawTxs() [][]byte {
	if x != nil {
		return x.RawTxs
	}
	return nil
}

type GetBlockByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Hash          []byte                 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBlockByIDRequest) Reset() {
	*x = GetBlockByIDRequest{}
	mi := &file_sync_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlockByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockByIDRequest) ProtoMessage() {}

func (x *GetBlockByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockByIDRequest.ProtoReflect.Descriptor instead.
func (*GetBlockByIDRequest) Descriptor() ([]byte, []int) {
	return file_sync_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlockByIDRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type GetBlockByIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RawBlock      []byte                 `protobuf:"bytes,1,opt,name=rawBlock,proto3" json:"rawBlock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBlockByIDResponse) Reset() {
	*x = GetBlockByIDResponse{}
	mi := &file_sync_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlockByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockByIDResponse) ProtoMessage() {}

func (x *GetBlockByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockByIDResponse.ProtoReflect.Descriptor instead.
func (*GetBlockByIDResponse) Descriptor() ([]byte, []int) {
	return file_sync_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlockByIDResponse) GetRawBlock() []byte {
	if x != nil {
		return x.RawBlock
	}
	return nil
}

var File_sync_message_proto protoreflect.FileDescriptor

var file_sync_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a,
	0x10, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x2e, 0x0a,
	0x12, 0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x45, 0x6e, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x22, 0x28, 0x0a,
	0x0e, 0x4e, 0x65, 0x77, 0x54, 0x78, 0x73, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x54, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x72, 0x61, 0x77, 0x54, 0x78, 0x73, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x2f, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x6e, 0x6f, 0x76, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sync_message_proto_rawDescOnce sync.Once
	file_sync_message_proto_rawDescData = file_sync_message_proto_rawDesc
)

func file_sync_message_proto_rawDescGZIP() []byte {
	file_sync_message_proto_rawDescOnce.Do(func() {
		file_sync_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_message_proto_rawDescData)
	})
	return file_sync_message_proto_rawDescData
}

var file_sync_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sync_message_proto_goTypes = []any{
	(*NewBlockEnvelope)(nil),     // 0: message.NewBlockEnvelope
	(*NewBlockIDEnvelope)(nil),   // 1: message.NewBlockIDEnvelope
	(*NewTxsEnvelope)(nil),       // 2: message.NewTxsEnvelope
	(*GetBlockByIDRequest)(nil),  // 3: message.GetBlockByIDRequest
	(*GetBlockByIDResponse)(nil), // 4: message.GetBlockByIDResponse
}
var file_sync_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sync_message_proto_init() }
func file_sync_message_proto_init() {
	if File_sync_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sync_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_message_proto_goTypes,
		DependencyIndexes: file_sync_message_proto_depIdxs,
		MessageInfos:      file_sync_message_proto_msgTypes,
	}.Build()
	File_sync_message_proto = out.File
	file_sync_message_proto_rawDesc = nil
	file_sync_message_proto_goTypes = nil
	file_sync_message_proto_depIdxs = nil
}