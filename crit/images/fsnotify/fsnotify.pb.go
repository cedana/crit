// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: fsnotify.proto

package fsnotify

import (
	fh "github.com/cedana/go-criu/v7/crit/images/fh"
	fown "github.com/cedana/go-criu/v7/crit/images/fown"
	_ "github.com/cedana/go-criu/v7/crit/images/opts"
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

type MarkType int32

const (
	MarkType_INODE MarkType = 1
	MarkType_MOUNT MarkType = 2
)

// Enum value maps for MarkType.
var (
	MarkType_name = map[int32]string{
		1: "INODE",
		2: "MOUNT",
	}
	MarkType_value = map[string]int32{
		"INODE": 1,
		"MOUNT": 2,
	}
)

func (x MarkType) Enum() *MarkType {
	p := new(MarkType)
	*p = x
	return p
}

func (x MarkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MarkType) Descriptor() protoreflect.EnumDescriptor {
	return file_fsnotify_proto_enumTypes[0].Descriptor()
}

func (MarkType) Type() protoreflect.EnumType {
	return &file_fsnotify_proto_enumTypes[0]
}

func (x MarkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MarkType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MarkType(num)
	return nil
}

// Deprecated: Use MarkType.Descriptor instead.
func (MarkType) EnumDescriptor() ([]byte, []int) {
	return file_fsnotify_proto_rawDescGZIP(), []int{0}
}

type InotifyWdEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	IIno          *uint64                `protobuf:"varint,2,req,name=i_ino,json=iIno" json:"i_ino,omitempty"`
	Mask          *uint32                `protobuf:"varint,3,req,name=mask" json:"mask,omitempty"`
	IgnoredMask   *uint32                `protobuf:"varint,4,req,name=ignored_mask,json=ignoredMask" json:"ignored_mask,omitempty"`
	SDev          *uint32                `protobuf:"varint,5,req,name=s_dev,json=sDev" json:"s_dev,omitempty"`
	Wd            *uint32                `protobuf:"varint,6,req,name=wd" json:"wd,omitempty"`
	FHandle       *fh.FhEntry            `protobuf:"bytes,7,req,name=f_handle,json=fHandle" json:"f_handle,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InotifyWdEntry) Reset() {
	*x = InotifyWdEntry{}
	mi := &file_fsnotify_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InotifyWdEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InotifyWdEntry) ProtoMessage() {}

func (x *InotifyWdEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fsnotify_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InotifyWdEntry.ProtoReflect.Descriptor instead.
func (*InotifyWdEntry) Descriptor() ([]byte, []int) {
	return file_fsnotify_proto_rawDescGZIP(), []int{0}
}

func (x *InotifyWdEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *InotifyWdEntry) GetIIno() uint64 {
	if x != nil && x.IIno != nil {
		return *x.IIno
	}
	return 0
}

func (x *InotifyWdEntry) GetMask() uint32 {
	if x != nil && x.Mask != nil {
		return *x.Mask
	}
	return 0
}

func (x *InotifyWdEntry) GetIgnoredMask() uint32 {
	if x != nil && x.IgnoredMask != nil {
		return *x.IgnoredMask
	}
	return 0
}

func (x *InotifyWdEntry) GetSDev() uint32 {
	if x != nil && x.SDev != nil {
		return *x.SDev
	}
	return 0
}

func (x *InotifyWdEntry) GetWd() uint32 {
	if x != nil && x.Wd != nil {
		return *x.Wd
	}
	return 0
}

func (x *InotifyWdEntry) GetFHandle() *fh.FhEntry {
	if x != nil {
		return x.FHandle
	}
	return nil
}

type InotifyFileEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Flags         *uint32                `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Fown          *fown.FownEntry        `protobuf:"bytes,4,req,name=fown" json:"fown,omitempty"`
	Wd            []*InotifyWdEntry      `protobuf:"bytes,5,rep,name=wd" json:"wd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InotifyFileEntry) Reset() {
	*x = InotifyFileEntry{}
	mi := &file_fsnotify_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InotifyFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InotifyFileEntry) ProtoMessage() {}

func (x *InotifyFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fsnotify_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InotifyFileEntry.ProtoReflect.Descriptor instead.
func (*InotifyFileEntry) Descriptor() ([]byte, []int) {
	return file_fsnotify_proto_rawDescGZIP(), []int{1}
}

func (x *InotifyFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *InotifyFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *InotifyFileEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *InotifyFileEntry) GetWd() []*InotifyWdEntry {
	if x != nil {
		return x.Wd
	}
	return nil
}

type FanotifyInodeMarkEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IIno          *uint64                `protobuf:"varint,1,req,name=i_ino,json=iIno" json:"i_ino,omitempty"`
	FHandle       *fh.FhEntry            `protobuf:"bytes,2,req,name=f_handle,json=fHandle" json:"f_handle,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FanotifyInodeMarkEntry) Reset() {
	*x = FanotifyInodeMarkEntry{}
	mi := &file_fsnotify_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FanotifyInodeMarkEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanotifyInodeMarkEntry) ProtoMessage() {}

func (x *FanotifyInodeMarkEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fsnotify_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanotifyInodeMarkEntry.ProtoReflect.Descriptor instead.
func (*FanotifyInodeMarkEntry) Descriptor() ([]byte, []int) {
	return file_fsnotify_proto_rawDescGZIP(), []int{2}
}

func (x *FanotifyInodeMarkEntry) GetIIno() uint64 {
	if x != nil && x.IIno != nil {
		return *x.IIno
	}
	return 0
}

func (x *FanotifyInodeMarkEntry) GetFHandle() *fh.FhEntry {
	if x != nil {
		return x.FHandle
	}
	return nil
}

type FanotifyMountMarkEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MntId         *uint32                `protobuf:"varint,1,req,name=mnt_id,json=mntId" json:"mnt_id,omitempty"`
	Path          *string                `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FanotifyMountMarkEntry) Reset() {
	*x = FanotifyMountMarkEntry{}
	mi := &file_fsnotify_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FanotifyMountMarkEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanotifyMountMarkEntry) ProtoMessage() {}

func (x *FanotifyMountMarkEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fsnotify_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanotifyMountMarkEntry.ProtoReflect.Descriptor instead.
func (*FanotifyMountMarkEntry) Descriptor() ([]byte, []int) {
	return file_fsnotify_proto_rawDescGZIP(), []int{3}
}

func (x *FanotifyMountMarkEntry) GetMntId() uint32 {
	if x != nil && x.MntId != nil {
		return *x.MntId
	}
	return 0
}

func (x *FanotifyMountMarkEntry) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

type FanotifyMarkEntry struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            *uint32                 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Type          *MarkType               `protobuf:"varint,2,req,name=type,enum=MarkType" json:"type,omitempty"`
	Mflags        *uint32                 `protobuf:"varint,3,req,name=mflags" json:"mflags,omitempty"`
	Mask          *uint32                 `protobuf:"varint,4,req,name=mask" json:"mask,omitempty"`
	IgnoredMask   *uint32                 `protobuf:"varint,5,req,name=ignored_mask,json=ignoredMask" json:"ignored_mask,omitempty"`
	SDev          *uint32                 `protobuf:"varint,6,req,name=s_dev,json=sDev" json:"s_dev,omitempty"`
	Ie            *FanotifyInodeMarkEntry `protobuf:"bytes,7,opt,name=ie" json:"ie,omitempty"`
	Me            *FanotifyMountMarkEntry `protobuf:"bytes,8,opt,name=me" json:"me,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FanotifyMarkEntry) Reset() {
	*x = FanotifyMarkEntry{}
	mi := &file_fsnotify_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FanotifyMarkEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanotifyMarkEntry) ProtoMessage() {}

func (x *FanotifyMarkEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fsnotify_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanotifyMarkEntry.ProtoReflect.Descriptor instead.
func (*FanotifyMarkEntry) Descriptor() ([]byte, []int) {
	return file_fsnotify_proto_rawDescGZIP(), []int{4}
}

func (x *FanotifyMarkEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *FanotifyMarkEntry) GetType() MarkType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return MarkType_INODE
}

func (x *FanotifyMarkEntry) GetMflags() uint32 {
	if x != nil && x.Mflags != nil {
		return *x.Mflags
	}
	return 0
}

func (x *FanotifyMarkEntry) GetMask() uint32 {
	if x != nil && x.Mask != nil {
		return *x.Mask
	}
	return 0
}

func (x *FanotifyMarkEntry) GetIgnoredMask() uint32 {
	if x != nil && x.IgnoredMask != nil {
		return *x.IgnoredMask
	}
	return 0
}

func (x *FanotifyMarkEntry) GetSDev() uint32 {
	if x != nil && x.SDev != nil {
		return *x.SDev
	}
	return 0
}

func (x *FanotifyMarkEntry) GetIe() *FanotifyInodeMarkEntry {
	if x != nil {
		return x.Ie
	}
	return nil
}

func (x *FanotifyMarkEntry) GetMe() *FanotifyMountMarkEntry {
	if x != nil {
		return x.Me
	}
	return nil
}

type FanotifyFileEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Flags         *uint32                `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Fown          *fown.FownEntry        `protobuf:"bytes,3,req,name=fown" json:"fown,omitempty"`
	Faflags       *uint32                `protobuf:"varint,4,req,name=faflags" json:"faflags,omitempty"`
	Evflags       *uint32                `protobuf:"varint,5,req,name=evflags" json:"evflags,omitempty"`
	Mark          []*FanotifyMarkEntry   `protobuf:"bytes,6,rep,name=mark" json:"mark,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FanotifyFileEntry) Reset() {
	*x = FanotifyFileEntry{}
	mi := &file_fsnotify_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FanotifyFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanotifyFileEntry) ProtoMessage() {}

func (x *FanotifyFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fsnotify_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanotifyFileEntry.ProtoReflect.Descriptor instead.
func (*FanotifyFileEntry) Descriptor() ([]byte, []int) {
	return file_fsnotify_proto_rawDescGZIP(), []int{5}
}

func (x *FanotifyFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *FanotifyFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *FanotifyFileEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *FanotifyFileEntry) GetFaflags() uint32 {
	if x != nil && x.Faflags != nil {
		return *x.Faflags
	}
	return 0
}

func (x *FanotifyFileEntry) GetEvflags() uint32 {
	if x != nil && x.Evflags != nil {
		return *x.Evflags
	}
	return 0
}

func (x *FanotifyFileEntry) GetMark() []*FanotifyMarkEntry {
	if x != nil {
		return x.Mark
	}
	return nil
}

var File_fsnotify_proto protoreflect.FileDescriptor

const file_fsnotify_proto_rawDesc = "" +
	"\n" +
	"\x0efsnotify.proto\x1a\n" +
	"opts.proto\x1a\bfh.proto\x1a\n" +
	"fown.proto\"\xce\x01\n" +
	"\x10inotify_wd_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x13\n" +
	"\x05i_ino\x18\x02 \x02(\x04R\x04iIno\x12\x19\n" +
	"\x04mask\x18\x03 \x02(\rB\x05\x92~\x02\b\x01R\x04mask\x12(\n" +
	"\fignored_mask\x18\x04 \x02(\rB\x05\x92~\x02\b\x01R\vignoredMask\x12\x1a\n" +
	"\x05s_dev\x18\x05 \x02(\rB\x05\x92~\x02 \x01R\x04sDev\x12\x0e\n" +
	"\x02wd\x18\x06 \x02(\rR\x02wd\x12$\n" +
	"\bf_handle\x18\a \x02(\v2\t.fh_entryR\afHandle\"\x85\x01\n" +
	"\x12inotify_file_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x1b\n" +
	"\x05flags\x18\x02 \x02(\rB\x05\x92~\x02\b\x01R\x05flags\x12\x1f\n" +
	"\x04fown\x18\x04 \x02(\v2\v.fown_entryR\x04fown\x12!\n" +
	"\x02wd\x18\x05 \x03(\v2\x11.inotify_wd_entryR\x02wd\"V\n" +
	"\x19fanotify_inode_mark_entry\x12\x13\n" +
	"\x05i_ino\x18\x01 \x02(\x04R\x04iIno\x12$\n" +
	"\bf_handle\x18\x02 \x02(\v2\t.fh_entryR\afHandle\"F\n" +
	"\x19fanotify_mount_mark_entry\x12\x15\n" +
	"\x06mnt_id\x18\x01 \x02(\rR\x05mntId\x12\x12\n" +
	"\x04path\x18\x02 \x01(\tR\x04path\"\x9d\x02\n" +
	"\x13fanotify_mark_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x1e\n" +
	"\x04type\x18\x02 \x02(\x0e2\n" +
	".mark_typeR\x04type\x12\x1d\n" +
	"\x06mflags\x18\x03 \x02(\rB\x05\x92~\x02\b\x01R\x06mflags\x12\x19\n" +
	"\x04mask\x18\x04 \x02(\rB\x05\x92~\x02\b\x01R\x04mask\x12(\n" +
	"\fignored_mask\x18\x05 \x02(\rB\x05\x92~\x02\b\x01R\vignoredMask\x12\x1a\n" +
	"\x05s_dev\x18\x06 \x02(\rB\x05\x92~\x02 \x01R\x04sDev\x12*\n" +
	"\x02ie\x18\a \x01(\v2\x1a.fanotify_inode_mark_entryR\x02ie\x12*\n" +
	"\x02me\x18\b \x01(\v2\x1a.fanotify_mount_mark_entryR\x02me\"\xcf\x01\n" +
	"\x13fanotify_file_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x1b\n" +
	"\x05flags\x18\x02 \x02(\rB\x05\x92~\x02\b\x01R\x05flags\x12\x1f\n" +
	"\x04fown\x18\x03 \x02(\v2\v.fown_entryR\x04fown\x12\x1f\n" +
	"\afaflags\x18\x04 \x02(\rB\x05\x92~\x02\b\x01R\afaflags\x12\x1f\n" +
	"\aevflags\x18\x05 \x02(\rB\x05\x92~\x02\b\x01R\aevflags\x12(\n" +
	"\x04mark\x18\x06 \x03(\v2\x14.fanotify_mark_entryR\x04mark*!\n" +
	"\tmark_type\x12\t\n" +
	"\x05INODE\x10\x01\x12\t\n" +
	"\x05MOUNT\x10\x02"

var (
	file_fsnotify_proto_rawDescOnce sync.Once
	file_fsnotify_proto_rawDescData []byte
)

func file_fsnotify_proto_rawDescGZIP() []byte {
	file_fsnotify_proto_rawDescOnce.Do(func() {
		file_fsnotify_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_fsnotify_proto_rawDesc), len(file_fsnotify_proto_rawDesc)))
	})
	return file_fsnotify_proto_rawDescData
}

var file_fsnotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fsnotify_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_fsnotify_proto_goTypes = []any{
	(MarkType)(0),                  // 0: mark_type
	(*InotifyWdEntry)(nil),         // 1: inotify_wd_entry
	(*InotifyFileEntry)(nil),       // 2: inotify_file_entry
	(*FanotifyInodeMarkEntry)(nil), // 3: fanotify_inode_mark_entry
	(*FanotifyMountMarkEntry)(nil), // 4: fanotify_mount_mark_entry
	(*FanotifyMarkEntry)(nil),      // 5: fanotify_mark_entry
	(*FanotifyFileEntry)(nil),      // 6: fanotify_file_entry
	(*fh.FhEntry)(nil),             // 7: fh_entry
	(*fown.FownEntry)(nil),         // 8: fown_entry
}
var file_fsnotify_proto_depIdxs = []int32{
	7, // 0: inotify_wd_entry.f_handle:type_name -> fh_entry
	8, // 1: inotify_file_entry.fown:type_name -> fown_entry
	1, // 2: inotify_file_entry.wd:type_name -> inotify_wd_entry
	7, // 3: fanotify_inode_mark_entry.f_handle:type_name -> fh_entry
	0, // 4: fanotify_mark_entry.type:type_name -> mark_type
	3, // 5: fanotify_mark_entry.ie:type_name -> fanotify_inode_mark_entry
	4, // 6: fanotify_mark_entry.me:type_name -> fanotify_mount_mark_entry
	8, // 7: fanotify_file_entry.fown:type_name -> fown_entry
	5, // 8: fanotify_file_entry.mark:type_name -> fanotify_mark_entry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_fsnotify_proto_init() }
func file_fsnotify_proto_init() {
	if File_fsnotify_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fsnotify_proto_rawDesc), len(file_fsnotify_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fsnotify_proto_goTypes,
		DependencyIndexes: file_fsnotify_proto_depIdxs,
		EnumInfos:         file_fsnotify_proto_enumTypes,
		MessageInfos:      file_fsnotify_proto_msgTypes,
	}.Build()
	File_fsnotify_proto = out.File
	file_fsnotify_proto_goTypes = nil
	file_fsnotify_proto_depIdxs = nil
}
