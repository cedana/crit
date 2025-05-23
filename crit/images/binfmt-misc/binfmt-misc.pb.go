// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: binfmt-misc.proto

package binfmt_misc

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

type BinfmtMiscEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *string                `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Enabled       *bool                  `protobuf:"varint,2,req,name=enabled" json:"enabled,omitempty"`
	Interpreter   *string                `protobuf:"bytes,3,req,name=interpreter" json:"interpreter,omitempty"`
	Flags         *string                `protobuf:"bytes,4,opt,name=flags" json:"flags,omitempty"`
	Extension     *string                `protobuf:"bytes,5,opt,name=extension" json:"extension,omitempty"`
	Magic         *string                `protobuf:"bytes,6,opt,name=magic" json:"magic,omitempty"`
	Mask          *string                `protobuf:"bytes,7,opt,name=mask" json:"mask,omitempty"`
	Offset        *int32                 `protobuf:"varint,8,opt,name=offset" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BinfmtMiscEntry) Reset() {
	*x = BinfmtMiscEntry{}
	mi := &file_binfmt_misc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinfmtMiscEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinfmtMiscEntry) ProtoMessage() {}

func (x *BinfmtMiscEntry) ProtoReflect() protoreflect.Message {
	mi := &file_binfmt_misc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinfmtMiscEntry.ProtoReflect.Descriptor instead.
func (*BinfmtMiscEntry) Descriptor() ([]byte, []int) {
	return file_binfmt_misc_proto_rawDescGZIP(), []int{0}
}

func (x *BinfmtMiscEntry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BinfmtMiscEntry) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *BinfmtMiscEntry) GetInterpreter() string {
	if x != nil && x.Interpreter != nil {
		return *x.Interpreter
	}
	return ""
}

func (x *BinfmtMiscEntry) GetFlags() string {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return ""
}

func (x *BinfmtMiscEntry) GetExtension() string {
	if x != nil && x.Extension != nil {
		return *x.Extension
	}
	return ""
}

func (x *BinfmtMiscEntry) GetMagic() string {
	if x != nil && x.Magic != nil {
		return *x.Magic
	}
	return ""
}

func (x *BinfmtMiscEntry) GetMask() string {
	if x != nil && x.Mask != nil {
		return *x.Mask
	}
	return ""
}

func (x *BinfmtMiscEntry) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

var File_binfmt_misc_proto protoreflect.FileDescriptor

const file_binfmt_misc_proto_rawDesc = "" +
	"\n" +
	"\x11binfmt-misc.proto\"\xd9\x01\n" +
	"\x11binfmt_misc_entry\x12\x12\n" +
	"\x04name\x18\x01 \x02(\tR\x04name\x12\x18\n" +
	"\aenabled\x18\x02 \x02(\bR\aenabled\x12 \n" +
	"\vinterpreter\x18\x03 \x02(\tR\vinterpreter\x12\x14\n" +
	"\x05flags\x18\x04 \x01(\tR\x05flags\x12\x1c\n" +
	"\textension\x18\x05 \x01(\tR\textension\x12\x14\n" +
	"\x05magic\x18\x06 \x01(\tR\x05magic\x12\x12\n" +
	"\x04mask\x18\a \x01(\tR\x04mask\x12\x16\n" +
	"\x06offset\x18\b \x01(\x05R\x06offset"

var (
	file_binfmt_misc_proto_rawDescOnce sync.Once
	file_binfmt_misc_proto_rawDescData []byte
)

func file_binfmt_misc_proto_rawDescGZIP() []byte {
	file_binfmt_misc_proto_rawDescOnce.Do(func() {
		file_binfmt_misc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_binfmt_misc_proto_rawDesc), len(file_binfmt_misc_proto_rawDesc)))
	})
	return file_binfmt_misc_proto_rawDescData
}

var file_binfmt_misc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_binfmt_misc_proto_goTypes = []any{
	(*BinfmtMiscEntry)(nil), // 0: binfmt_misc_entry
}
var file_binfmt_misc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_binfmt_misc_proto_init() }
func file_binfmt_misc_proto_init() {
	if File_binfmt_misc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_binfmt_misc_proto_rawDesc), len(file_binfmt_misc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_binfmt_misc_proto_goTypes,
		DependencyIndexes: file_binfmt_misc_proto_depIdxs,
		MessageInfos:      file_binfmt_misc_proto_msgTypes,
	}.Build()
	File_binfmt_misc_proto = out.File
	file_binfmt_misc_proto_goTypes = nil
	file_binfmt_misc_proto_depIdxs = nil
}
