// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.2
// source: consensus_message.proto

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

type ConsensusEnvelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Raw           []byte                 `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	SenderAddr    []byte                 `protobuf:"bytes,2,opt,name=senderAddr,proto3" json:"senderAddr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsensusEnvelope) Reset() {
	*x = ConsensusEnvelope{}
	mi := &file_consensus_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsensusEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusEnvelope) ProtoMessage() {}

func (x *ConsensusEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusEnvelope.ProtoReflect.Descriptor instead.
func (*ConsensusEnvelope) Descriptor() ([]byte, []int) {
	return file_consensus_message_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusEnvelope) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *ConsensusEnvelope) GetSenderAddr() []byte {
	if x != nil {
		return x.SenderAddr
	}
	return nil
}

var File_consensus_message_proto protoreflect.FileDescriptor

var file_consensus_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x2f,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x76, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_message_proto_rawDescOnce sync.Once
	file_consensus_message_proto_rawDescData = file_consensus_message_proto_rawDesc
)

func file_consensus_message_proto_rawDescGZIP() []byte {
	file_consensus_message_proto_rawDescOnce.Do(func() {
		file_consensus_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_message_proto_rawDescData)
	})
	return file_consensus_message_proto_rawDescData
}

var file_consensus_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_consensus_message_proto_goTypes = []any{
	(*ConsensusEnvelope)(nil), // 0: message.ConsensusEnvelope
}
var file_consensus_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_consensus_message_proto_init() }
func file_consensus_message_proto_init() {
	if File_consensus_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_consensus_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consensus_message_proto_goTypes,
		DependencyIndexes: file_consensus_message_proto_depIdxs,
		MessageInfos:      file_consensus_message_proto_msgTypes,
	}.Build()
	File_consensus_message_proto = out.File
	file_consensus_message_proto_rawDesc = nil
	file_consensus_message_proto_goTypes = nil
	file_consensus_message_proto_depIdxs = nil
}