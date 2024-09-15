// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: curp_command.proto

package curp_proto

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

type CurpClientCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op       uint32 `protobuf:"varint,1,opt,name=Op,proto3" json:"Op,omitempty"`
	ClientId uint64 `protobuf:"varint,2,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	SeqId    uint64 `protobuf:"varint,3,opt,name=SeqId,proto3" json:"SeqId,omitempty"`
	Key      string `protobuf:"bytes,4,opt,name=Key,proto3" json:"Key,omitempty"`
	Value    string `protobuf:"bytes,5,opt,name=Value,proto3" json:"Value,omitempty"`
	Sync     uint32 `protobuf:"varint,6,opt,name=Sync,proto3" json:"Sync,omitempty"`
}

func (x *CurpClientCommand) Reset() {
	*x = CurpClientCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_curp_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurpClientCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurpClientCommand) ProtoMessage() {}

func (x *CurpClientCommand) ProtoReflect() protoreflect.Message {
	mi := &file_curp_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurpClientCommand.ProtoReflect.Descriptor instead.
func (*CurpClientCommand) Descriptor() ([]byte, []int) {
	return file_curp_command_proto_rawDescGZIP(), []int{0}
}

func (x *CurpClientCommand) GetOp() uint32 {
	if x != nil {
		return x.Op
	}
	return 0
}

func (x *CurpClientCommand) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *CurpClientCommand) GetSeqId() uint64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *CurpClientCommand) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CurpClientCommand) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CurpClientCommand) GetSync() uint32 {
	if x != nil {
		return x.Sync
	}
	return 0
}

type CurpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content    string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	StatusCode uint32 `protobuf:"varint,2,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	SeqId      uint64 `protobuf:"varint,3,opt,name=SeqId,proto3" json:"SeqId,omitempty"`
}

func (x *CurpReply) Reset() {
	*x = CurpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_curp_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurpReply) ProtoMessage() {}

func (x *CurpReply) ProtoReflect() protoreflect.Message {
	mi := &file_curp_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurpReply.ProtoReflect.Descriptor instead.
func (*CurpReply) Descriptor() ([]byte, []int) {
	return file_curp_command_proto_rawDescGZIP(), []int{1}
}

func (x *CurpReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CurpReply) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CurpReply) GetSeqId() uint64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

type LeaderAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaderAck) Reset() {
	*x = LeaderAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_curp_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderAck) ProtoMessage() {}

func (x *LeaderAck) ProtoReflect() protoreflect.Message {
	mi := &file_curp_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderAck.ProtoReflect.Descriptor instead.
func (*LeaderAck) Descriptor() ([]byte, []int) {
	return file_curp_command_proto_rawDescGZIP(), []int{2}
}

type LeaderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLeader uint32 `protobuf:"varint,1,opt,name=IsLeader,proto3" json:"IsLeader,omitempty"`
}

func (x *LeaderReply) Reset() {
	*x = LeaderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_curp_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderReply) ProtoMessage() {}

func (x *LeaderReply) ProtoReflect() protoreflect.Message {
	mi := &file_curp_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderReply.ProtoReflect.Descriptor instead.
func (*LeaderReply) Descriptor() ([]byte, []int) {
	return file_curp_command_proto_rawDescGZIP(), []int{3}
}

func (x *LeaderReply) GetIsLeader() uint32 {
	if x != nil {
		return x.IsLeader
	}
	return 0
}

type ProposeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId uint64 `protobuf:"varint,1,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	SeqId    uint64 `protobuf:"varint,2,opt,name=SeqId,proto3" json:"SeqId,omitempty"`
}

func (x *ProposeId) Reset() {
	*x = ProposeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_curp_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeId) ProtoMessage() {}

func (x *ProposeId) ProtoReflect() protoreflect.Message {
	mi := &file_curp_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeId.ProtoReflect.Descriptor instead.
func (*ProposeId) Descriptor() ([]byte, []int) {
	return file_curp_command_proto_rawDescGZIP(), []int{4}
}

func (x *ProposeId) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ProposeId) GetSeqId() uint64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

var File_curp_command_proto protoreflect.FileDescriptor

var file_curp_command_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x75, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x75, 0x72, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x91, 0x01, 0x0a, 0x11, 0x43, 0x75, 0x72, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x71, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x53, 0x65, 0x71, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x53, 0x79, 0x6e, 0x63, 0x22, 0x5b, 0x0a, 0x09, 0x43, 0x75, 0x72, 0x70, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53,
	0x65, 0x71, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x53, 0x65, 0x71, 0x49,
	0x64, 0x22, 0x0b, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x63, 0x6b, 0x22, 0x29,
	0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x71, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x53, 0x65, 0x71, 0x49, 0x64, 0x32, 0xbf, 0x01, 0x0a, 0x04, 0x43, 0x75, 0x72,
	0x70, 0x12, 0x3f, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x63,
	0x75, 0x72, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x70, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x15, 0x2e, 0x63, 0x75,
	0x72, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x70, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x57, 0x61, 0x69, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64,
	0x12, 0x15, 0x2e, 0x63, 0x75, 0x72, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x63, 0x75, 0x72, 0x70, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a,
	0x0a, 0x08, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x63, 0x75, 0x72,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x63,
	0x6b, 0x1a, 0x17, 0x2e, 0x63, 0x75, 0x72, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x63,
	0x75, 0x72, 0x70, 0x2f, 0x63, 0x75, 0x72, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_curp_command_proto_rawDescOnce sync.Once
	file_curp_command_proto_rawDescData = file_curp_command_proto_rawDesc
)

func file_curp_command_proto_rawDescGZIP() []byte {
	file_curp_command_proto_rawDescOnce.Do(func() {
		file_curp_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_curp_command_proto_rawDescData)
	})
	return file_curp_command_proto_rawDescData
}

var file_curp_command_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_curp_command_proto_goTypes = []interface{}{
	(*CurpClientCommand)(nil), // 0: curp_proto.CurpClientCommand
	(*CurpReply)(nil),         // 1: curp_proto.CurpReply
	(*LeaderAck)(nil),         // 2: curp_proto.LeaderAck
	(*LeaderReply)(nil),       // 3: curp_proto.LeaderReply
	(*ProposeId)(nil),         // 4: curp_proto.ProposeId
}
var file_curp_command_proto_depIdxs = []int32{
	0, // 0: curp_proto.Curp.Propose:input_type -> curp_proto.CurpClientCommand
	4, // 1: curp_proto.Curp.WaitSynced:input_type -> curp_proto.ProposeId
	2, // 2: curp_proto.Curp.IsLeader:input_type -> curp_proto.LeaderAck
	1, // 3: curp_proto.Curp.Propose:output_type -> curp_proto.CurpReply
	1, // 4: curp_proto.Curp.WaitSynced:output_type -> curp_proto.CurpReply
	3, // 5: curp_proto.Curp.IsLeader:output_type -> curp_proto.LeaderReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_curp_command_proto_init() }
func file_curp_command_proto_init() {
	if File_curp_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_curp_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurpClientCommand); i {
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
		file_curp_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurpReply); i {
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
		file_curp_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderAck); i {
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
		file_curp_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderReply); i {
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
		file_curp_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeId); i {
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
			RawDescriptor: file_curp_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_curp_command_proto_goTypes,
		DependencyIndexes: file_curp_command_proto_depIdxs,
		MessageInfos:      file_curp_command_proto_msgTypes,
	}.Build()
	File_curp_command_proto = out.File
	file_curp_command_proto_rawDesc = nil
	file_curp_command_proto_goTypes = nil
	file_curp_command_proto_depIdxs = nil
}