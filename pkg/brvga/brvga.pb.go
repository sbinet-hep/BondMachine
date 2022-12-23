// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: brvga.proto

package brvga

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

type Textmemupdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpid uint32                   `protobuf:"varint,1,opt,name=cpid,proto3" json:"cpid,omitempty"`
	Seq  []*Textmemupdate_Byteseq `protobuf:"bytes,2,rep,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Textmemupdate) Reset() {
	*x = Textmemupdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brvga_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Textmemupdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Textmemupdate) ProtoMessage() {}

func (x *Textmemupdate) ProtoReflect() protoreflect.Message {
	mi := &file_brvga_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Textmemupdate.ProtoReflect.Descriptor instead.
func (*Textmemupdate) Descriptor() ([]byte, []int) {
	return file_brvga_proto_rawDescGZIP(), []int{0}
}

func (x *Textmemupdate) GetCpid() uint32 {
	if x != nil {
		return x.Cpid
	}
	return 0
}

func (x *Textmemupdate) GetSeq() []*Textmemupdate_Byteseq {
	if x != nil {
		return x.Seq
	}
	return nil
}

type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updates []*Textmemupdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brvga_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_brvga_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update.ProtoReflect.Descriptor instead.
func (*Update) Descriptor() ([]byte, []int) {
	return file_brvga_proto_rawDescGZIP(), []int{1}
}

func (x *Update) GetUpdates() []*Textmemupdate {
	if x != nil {
		return x.Updates
	}
	return nil
}

type Textmemupdate_Byteseq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos     uint32 `protobuf:"varint,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Textmemupdate_Byteseq) Reset() {
	*x = Textmemupdate_Byteseq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brvga_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Textmemupdate_Byteseq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Textmemupdate_Byteseq) ProtoMessage() {}

func (x *Textmemupdate_Byteseq) ProtoReflect() protoreflect.Message {
	mi := &file_brvga_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Textmemupdate_Byteseq.ProtoReflect.Descriptor instead.
func (*Textmemupdate_Byteseq) Descriptor() ([]byte, []int) {
	return file_brvga_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Textmemupdate_Byteseq) GetPos() uint32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *Textmemupdate_Byteseq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_brvga_proto protoreflect.FileDescriptor

var file_brvga_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x72, 0x76, 0x67, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x72, 0x76, 0x67, 0x61, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x78, 0x74, 0x6d, 0x65, 0x6d,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x70, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x03, 0x73, 0x65,
	0x71, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x72, 0x76, 0x67, 0x61, 0x2e,
	0x54, 0x65, 0x78, 0x74, 0x6d, 0x65, 0x6d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x65, 0x71, 0x52, 0x03, 0x73, 0x65, 0x71, 0x1a, 0x35, 0x0a, 0x07, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62,
	0x72, 0x76, 0x67, 0x61, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x6d, 0x65, 0x6d, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x62, 0x72, 0x76, 0x67, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_brvga_proto_rawDescOnce sync.Once
	file_brvga_proto_rawDescData = file_brvga_proto_rawDesc
)

func file_brvga_proto_rawDescGZIP() []byte {
	file_brvga_proto_rawDescOnce.Do(func() {
		file_brvga_proto_rawDescData = protoimpl.X.CompressGZIP(file_brvga_proto_rawDescData)
	})
	return file_brvga_proto_rawDescData
}

var file_brvga_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_brvga_proto_goTypes = []interface{}{
	(*Textmemupdate)(nil),         // 0: brvga.Textmemupdate
	(*Update)(nil),                // 1: brvga.Update
	(*Textmemupdate_Byteseq)(nil), // 2: brvga.Textmemupdate.Byteseq
}
var file_brvga_proto_depIdxs = []int32{
	2, // 0: brvga.Textmemupdate.seq:type_name -> brvga.Textmemupdate.Byteseq
	0, // 1: brvga.Update.updates:type_name -> brvga.Textmemupdate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_brvga_proto_init() }
func file_brvga_proto_init() {
	if File_brvga_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_brvga_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Textmemupdate); i {
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
		file_brvga_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update); i {
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
		file_brvga_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Textmemupdate_Byteseq); i {
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
			RawDescriptor: file_brvga_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_brvga_proto_goTypes,
		DependencyIndexes: file_brvga_proto_depIdxs,
		MessageInfos:      file_brvga_proto_msgTypes,
	}.Build()
	File_brvga_proto = out.File
	file_brvga_proto_rawDesc = nil
	file_brvga_proto_goTypes = nil
	file_brvga_proto_depIdxs = nil
}
