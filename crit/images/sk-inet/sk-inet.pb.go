// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: sk-inet.proto

package sk_inet

import (
	fown "github.com/cedana/go-criu/v7/crit/images/fown"
	_ "github.com/cedana/go-criu/v7/crit/images/opts"
	sk_opts "github.com/cedana/go-criu/v7/crit/images/sk-opts"
	tcp_stream "github.com/cedana/go-criu/v7/crit/images/tcp-stream"
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

type IpOptsRawEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Hdrincl       *bool                  `protobuf:"varint,1,opt,name=hdrincl" json:"hdrincl,omitempty"`
	Nodefrag      *bool                  `protobuf:"varint,2,opt,name=nodefrag" json:"nodefrag,omitempty"`
	Checksum      *bool                  `protobuf:"varint,3,opt,name=checksum" json:"checksum,omitempty"`
	IcmpvFilter   []uint32               `protobuf:"varint,4,rep,name=icmpv_filter,json=icmpvFilter" json:"icmpv_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IpOptsRawEntry) Reset() {
	*x = IpOptsRawEntry{}
	mi := &file_sk_inet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpOptsRawEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpOptsRawEntry) ProtoMessage() {}

func (x *IpOptsRawEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_inet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpOptsRawEntry.ProtoReflect.Descriptor instead.
func (*IpOptsRawEntry) Descriptor() ([]byte, []int) {
	return file_sk_inet_proto_rawDescGZIP(), []int{0}
}

func (x *IpOptsRawEntry) GetHdrincl() bool {
	if x != nil && x.Hdrincl != nil {
		return *x.Hdrincl
	}
	return false
}

func (x *IpOptsRawEntry) GetNodefrag() bool {
	if x != nil && x.Nodefrag != nil {
		return *x.Nodefrag
	}
	return false
}

func (x *IpOptsRawEntry) GetChecksum() bool {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return false
}

func (x *IpOptsRawEntry) GetIcmpvFilter() []uint32 {
	if x != nil {
		return x.IcmpvFilter
	}
	return nil
}

type IpOptsEntry struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Freebind *bool                  `protobuf:"varint,1,opt,name=freebind" json:"freebind,omitempty"`
	// Fields 2 and 3 are reserved for vz7 use
	Raw           *IpOptsRawEntry `protobuf:"bytes,4,opt,name=raw" json:"raw,omitempty"`
	Pktinfo       *bool           `protobuf:"varint,5,opt,name=pktinfo" json:"pktinfo,omitempty"`
	Tos           *uint32         `protobuf:"varint,6,opt,name=tos" json:"tos,omitempty"`
	Ttl           *uint32         `protobuf:"varint,7,opt,name=ttl" json:"ttl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IpOptsEntry) Reset() {
	*x = IpOptsEntry{}
	mi := &file_sk_inet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpOptsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpOptsEntry) ProtoMessage() {}

func (x *IpOptsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_inet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpOptsEntry.ProtoReflect.Descriptor instead.
func (*IpOptsEntry) Descriptor() ([]byte, []int) {
	return file_sk_inet_proto_rawDescGZIP(), []int{1}
}

func (x *IpOptsEntry) GetFreebind() bool {
	if x != nil && x.Freebind != nil {
		return *x.Freebind
	}
	return false
}

func (x *IpOptsEntry) GetRaw() *IpOptsRawEntry {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *IpOptsEntry) GetPktinfo() bool {
	if x != nil && x.Pktinfo != nil {
		return *x.Pktinfo
	}
	return false
}

func (x *IpOptsEntry) GetTos() uint32 {
	if x != nil && x.Tos != nil {
		return *x.Tos
	}
	return 0
}

func (x *IpOptsEntry) GetTtl() uint32 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

type InetSkEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// We have two IDs here -- id and ino. The first one
	// is used when restoring socket behind a file descriprot.
	// The fdinfo image's id is it. The second one is used
	// in sk-inet.c internally, in particular we identify
	// a TCP stream to restore into this socket using the
	// ino value.
	Id      *uint32              `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Ino     *uint32              `protobuf:"varint,2,req,name=ino" json:"ino,omitempty"`
	Family  *uint32              `protobuf:"varint,3,req,name=family" json:"family,omitempty"`
	Type    *uint32              `protobuf:"varint,4,req,name=type" json:"type,omitempty"`
	Proto   *uint32              `protobuf:"varint,5,req,name=proto" json:"proto,omitempty"`
	State   *uint32              `protobuf:"varint,6,req,name=state" json:"state,omitempty"`
	SrcPort *uint32              `protobuf:"varint,7,req,name=src_port,json=srcPort" json:"src_port,omitempty"`
	DstPort *uint32              `protobuf:"varint,8,req,name=dst_port,json=dstPort" json:"dst_port,omitempty"`
	Flags   *uint32              `protobuf:"varint,9,req,name=flags" json:"flags,omitempty"`
	Backlog *uint32              `protobuf:"varint,10,req,name=backlog" json:"backlog,omitempty"`
	SrcAddr []uint32             `protobuf:"varint,11,rep,name=src_addr,json=srcAddr" json:"src_addr,omitempty"`
	DstAddr []uint32             `protobuf:"varint,12,rep,name=dst_addr,json=dstAddr" json:"dst_addr,omitempty"`
	Fown    *fown.FownEntry      `protobuf:"bytes,13,req,name=fown" json:"fown,omitempty"`
	Opts    *sk_opts.SkOptsEntry `protobuf:"bytes,14,req,name=opts" json:"opts,omitempty"`
	V6Only  *bool                `protobuf:"varint,15,opt,name=v6only" json:"v6only,omitempty"`
	IpOpts  *IpOptsEntry         `protobuf:"bytes,16,opt,name=ip_opts,json=ipOpts" json:"ip_opts,omitempty"`
	// for ipv6, we need to send the ifindex to bind(); we keep the ifname
	// here and convert it on restore
	Ifname        *string                  `protobuf:"bytes,17,opt,name=ifname" json:"ifname,omitempty"`
	NsId          *uint32                  `protobuf:"varint,18,opt,name=ns_id,json=nsId" json:"ns_id,omitempty"`
	Shutdown      *sk_opts.SkShutdown      `protobuf:"varint,19,opt,name=shutdown,enum=SkShutdown" json:"shutdown,omitempty"`
	TcpOpts       *tcp_stream.TcpOptsEntry `protobuf:"bytes,20,opt,name=tcp_opts,json=tcpOpts" json:"tcp_opts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InetSkEntry) Reset() {
	*x = InetSkEntry{}
	mi := &file_sk_inet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InetSkEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InetSkEntry) ProtoMessage() {}

func (x *InetSkEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_inet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InetSkEntry.ProtoReflect.Descriptor instead.
func (*InetSkEntry) Descriptor() ([]byte, []int) {
	return file_sk_inet_proto_rawDescGZIP(), []int{2}
}

func (x *InetSkEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *InetSkEntry) GetIno() uint32 {
	if x != nil && x.Ino != nil {
		return *x.Ino
	}
	return 0
}

func (x *InetSkEntry) GetFamily() uint32 {
	if x != nil && x.Family != nil {
		return *x.Family
	}
	return 0
}

func (x *InetSkEntry) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *InetSkEntry) GetProto() uint32 {
	if x != nil && x.Proto != nil {
		return *x.Proto
	}
	return 0
}

func (x *InetSkEntry) GetState() uint32 {
	if x != nil && x.State != nil {
		return *x.State
	}
	return 0
}

func (x *InetSkEntry) GetSrcPort() uint32 {
	if x != nil && x.SrcPort != nil {
		return *x.SrcPort
	}
	return 0
}

func (x *InetSkEntry) GetDstPort() uint32 {
	if x != nil && x.DstPort != nil {
		return *x.DstPort
	}
	return 0
}

func (x *InetSkEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *InetSkEntry) GetBacklog() uint32 {
	if x != nil && x.Backlog != nil {
		return *x.Backlog
	}
	return 0
}

func (x *InetSkEntry) GetSrcAddr() []uint32 {
	if x != nil {
		return x.SrcAddr
	}
	return nil
}

func (x *InetSkEntry) GetDstAddr() []uint32 {
	if x != nil {
		return x.DstAddr
	}
	return nil
}

func (x *InetSkEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *InetSkEntry) GetOpts() *sk_opts.SkOptsEntry {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *InetSkEntry) GetV6Only() bool {
	if x != nil && x.V6Only != nil {
		return *x.V6Only
	}
	return false
}

func (x *InetSkEntry) GetIpOpts() *IpOptsEntry {
	if x != nil {
		return x.IpOpts
	}
	return nil
}

func (x *InetSkEntry) GetIfname() string {
	if x != nil && x.Ifname != nil {
		return *x.Ifname
	}
	return ""
}

func (x *InetSkEntry) GetNsId() uint32 {
	if x != nil && x.NsId != nil {
		return *x.NsId
	}
	return 0
}

func (x *InetSkEntry) GetShutdown() sk_opts.SkShutdown {
	if x != nil && x.Shutdown != nil {
		return *x.Shutdown
	}
	return sk_opts.SkShutdown(0)
}

func (x *InetSkEntry) GetTcpOpts() *tcp_stream.TcpOptsEntry {
	if x != nil {
		return x.TcpOpts
	}
	return nil
}

var File_sk_inet_proto protoreflect.FileDescriptor

const file_sk_inet_proto_rawDesc = "" +
	"\n" +
	"\rsk-inet.proto\x1a\n" +
	"opts.proto\x1a\n" +
	"fown.proto\x1a\rsk-opts.proto\x1a\x10tcp-stream.proto\"\x88\x01\n" +
	"\x11ip_opts_raw_entry\x12\x18\n" +
	"\ahdrincl\x18\x01 \x01(\bR\ahdrincl\x12\x1a\n" +
	"\bnodefrag\x18\x02 \x01(\bR\bnodefrag\x12\x1a\n" +
	"\bchecksum\x18\x03 \x01(\bR\bchecksum\x12!\n" +
	"\ficmpv_filter\x18\x04 \x03(\rR\vicmpvFilter\"\x8f\x01\n" +
	"\rip_opts_entry\x12\x1a\n" +
	"\bfreebind\x18\x01 \x01(\bR\bfreebind\x12$\n" +
	"\x03raw\x18\x04 \x01(\v2\x12.ip_opts_raw_entryR\x03raw\x12\x18\n" +
	"\apktinfo\x18\x05 \x01(\bR\apktinfo\x12\x10\n" +
	"\x03tos\x18\x06 \x01(\rR\x03tos\x12\x10\n" +
	"\x03ttl\x18\a \x01(\rR\x03ttl\"\xe7\x04\n" +
	"\rinet_sk_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x10\n" +
	"\x03ino\x18\x02 \x02(\rR\x03ino\x12\x1f\n" +
	"\x06family\x18\x03 \x02(\rB\a\x92~\x042\x02skR\x06family\x12\x1b\n" +
	"\x04type\x18\x04 \x02(\rB\a\x92~\x042\x02skR\x04type\x12\x1d\n" +
	"\x05proto\x18\x05 \x02(\rB\a\x92~\x042\x02skR\x05proto\x12\x1d\n" +
	"\x05state\x18\x06 \x02(\rB\a\x92~\x042\x02skR\x05state\x12\x19\n" +
	"\bsrc_port\x18\a \x02(\rR\asrcPort\x12\x19\n" +
	"\bdst_port\x18\b \x02(\rR\adstPort\x12\x1b\n" +
	"\x05flags\x18\t \x02(\rB\x05\x92~\x02\b\x01R\x05flags\x12\x18\n" +
	"\abacklog\x18\n" +
	" \x02(\rR\abacklog\x12 \n" +
	"\bsrc_addr\x18\v \x03(\rB\x05\x92~\x02\x10\x01R\asrcAddr\x12 \n" +
	"\bdst_addr\x18\f \x03(\rB\x05\x92~\x02\x10\x01R\adstAddr\x12\x1f\n" +
	"\x04fown\x18\r \x02(\v2\v.fown_entryR\x04fown\x12\"\n" +
	"\x04opts\x18\x0e \x02(\v2\x0e.sk_opts_entryR\x04opts\x12\x16\n" +
	"\x06v6only\x18\x0f \x01(\bR\x06v6only\x12'\n" +
	"\aip_opts\x18\x10 \x01(\v2\x0e.ip_opts_entryR\x06ipOpts\x12\x16\n" +
	"\x06ifname\x18\x11 \x01(\tR\x06ifname\x12\x13\n" +
	"\x05ns_id\x18\x12 \x01(\rR\x04nsId\x12(\n" +
	"\bshutdown\x18\x13 \x01(\x0e2\f.sk_shutdownR\bshutdown\x12*\n" +
	"\btcp_opts\x18\x14 \x01(\v2\x0f.tcp_opts_entryR\atcpOpts"

var (
	file_sk_inet_proto_rawDescOnce sync.Once
	file_sk_inet_proto_rawDescData []byte
)

func file_sk_inet_proto_rawDescGZIP() []byte {
	file_sk_inet_proto_rawDescOnce.Do(func() {
		file_sk_inet_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sk_inet_proto_rawDesc), len(file_sk_inet_proto_rawDesc)))
	})
	return file_sk_inet_proto_rawDescData
}

var file_sk_inet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sk_inet_proto_goTypes = []any{
	(*IpOptsRawEntry)(nil),          // 0: ip_opts_raw_entry
	(*IpOptsEntry)(nil),             // 1: ip_opts_entry
	(*InetSkEntry)(nil),             // 2: inet_sk_entry
	(*fown.FownEntry)(nil),          // 3: fown_entry
	(*sk_opts.SkOptsEntry)(nil),     // 4: sk_opts_entry
	(sk_opts.SkShutdown)(0),         // 5: sk_shutdown
	(*tcp_stream.TcpOptsEntry)(nil), // 6: tcp_opts_entry
}
var file_sk_inet_proto_depIdxs = []int32{
	0, // 0: ip_opts_entry.raw:type_name -> ip_opts_raw_entry
	3, // 1: inet_sk_entry.fown:type_name -> fown_entry
	4, // 2: inet_sk_entry.opts:type_name -> sk_opts_entry
	1, // 3: inet_sk_entry.ip_opts:type_name -> ip_opts_entry
	5, // 4: inet_sk_entry.shutdown:type_name -> sk_shutdown
	6, // 5: inet_sk_entry.tcp_opts:type_name -> tcp_opts_entry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_sk_inet_proto_init() }
func file_sk_inet_proto_init() {
	if File_sk_inet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sk_inet_proto_rawDesc), len(file_sk_inet_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sk_inet_proto_goTypes,
		DependencyIndexes: file_sk_inet_proto_depIdxs,
		MessageInfos:      file_sk_inet_proto_msgTypes,
	}.Build()
	File_sk_inet_proto = out.File
	file_sk_inet_proto_goTypes = nil
	file_sk_inet_proto_depIdxs = nil
}
