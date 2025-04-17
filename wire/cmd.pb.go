// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: protos/cmd.proto

package wire

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Command struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cmd           string                 `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Args          []string               `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Command) Reset() {
	*x = Command{}
	mi := &file_protos_cmd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cmd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_protos_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *Command) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_protos_cmd_proto protoreflect.FileDescriptor

var file_protos_cmd_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x77, 0x69, 0x72, 0x65, 0x22, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x77,
	0x69, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protos_cmd_proto_rawDescOnce sync.Once
	file_protos_cmd_proto_rawDescData []byte
)

func file_protos_cmd_proto_rawDescGZIP() []byte {
	file_protos_cmd_proto_rawDescOnce.Do(func() {
		file_protos_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_cmd_proto_rawDesc), len(file_protos_cmd_proto_rawDesc)))
	})
	return file_protos_cmd_proto_rawDescData
}

var file_protos_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_cmd_proto_goTypes = []any{
	(*Command)(nil), // 0: wire.Command
}
var file_protos_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_cmd_proto_init() }
func file_protos_cmd_proto_init() {
	if File_protos_cmd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_cmd_proto_rawDesc), len(file_protos_cmd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_cmd_proto_goTypes,
		DependencyIndexes: file_protos_cmd_proto_depIdxs,
		MessageInfos:      file_protos_cmd_proto_msgTypes,
	}.Build()
	File_protos_cmd_proto = out.File
	file_protos_cmd_proto_goTypes = nil
	file_protos_cmd_proto_depIdxs = nil
}
