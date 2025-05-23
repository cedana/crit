// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: eventfd.proto

package eventfd

import (
	fown "github.com/cedana/go-criu/v7/crit/images/fown"
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

type EventfdFileEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Flags         *uint32                `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Fown          *fown.FownEntry        `protobuf:"bytes,3,req,name=fown" json:"fown,omitempty"`
	Counter       *uint64                `protobuf:"varint,4,req,name=counter" json:"counter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventfdFileEntry) Reset() {
	*x = EventfdFileEntry{}
	mi := &file_eventfd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventfdFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventfdFileEntry) ProtoMessage() {}

func (x *EventfdFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_eventfd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventfdFileEntry.ProtoReflect.Descriptor instead.
func (*EventfdFileEntry) Descriptor() ([]byte, []int) {
	return file_eventfd_proto_rawDescGZIP(), []int{0}
}

func (x *EventfdFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *EventfdFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *EventfdFileEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *EventfdFileEntry) GetCounter() uint64 {
	if x != nil && x.Counter != nil {
		return *x.Counter
	}
	return 0
}

var File_eventfd_proto protoreflect.FileDescriptor

const file_eventfd_proto_rawDesc = "" +
	"\n" +
	"\reventfd.proto\x1a\n" +
	"fown.proto\"u\n" +
	"\x12eventfd_file_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x14\n" +
	"\x05flags\x18\x02 \x02(\rR\x05flags\x12\x1f\n" +
	"\x04fown\x18\x03 \x02(\v2\v.fown_entryR\x04fown\x12\x18\n" +
	"\acounter\x18\x04 \x02(\x04R\acounter"

var (
	file_eventfd_proto_rawDescOnce sync.Once
	file_eventfd_proto_rawDescData []byte
)

func file_eventfd_proto_rawDescGZIP() []byte {
	file_eventfd_proto_rawDescOnce.Do(func() {
		file_eventfd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eventfd_proto_rawDesc), len(file_eventfd_proto_rawDesc)))
	})
	return file_eventfd_proto_rawDescData
}

var file_eventfd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eventfd_proto_goTypes = []any{
	(*EventfdFileEntry)(nil), // 0: eventfd_file_entry
	(*fown.FownEntry)(nil),   // 1: fown_entry
}
var file_eventfd_proto_depIdxs = []int32{
	1, // 0: eventfd_file_entry.fown:type_name -> fown_entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eventfd_proto_init() }
func file_eventfd_proto_init() {
	if File_eventfd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eventfd_proto_rawDesc), len(file_eventfd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventfd_proto_goTypes,
		DependencyIndexes: file_eventfd_proto_depIdxs,
		MessageInfos:      file_eventfd_proto_msgTypes,
	}.Build()
	File_eventfd_proto = out.File
	file_eventfd_proto_goTypes = nil
	file_eventfd_proto_depIdxs = nil
}
