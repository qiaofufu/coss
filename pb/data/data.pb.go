// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pb/data/data.proto

package data

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string `protobuf:"bytes,1,opt,name=TableName,proto3" json:"TableName,omitempty"`
	Size      int64  `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	BlockID   int64  `protobuf:"varint,3,opt,name=BlockID,proto3" json:"BlockID,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_pb_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Block) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Block) GetBlockID() int64 {
	if x != nil {
		return x.BlockID
	}
	return 0
}

type DataBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
}

func (x *DataBlocks) Reset() {
	*x = DataBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBlocks) ProtoMessage() {}

func (x *DataBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBlocks.ProtoReflect.Descriptor instead.
func (*DataBlocks) Descriptor() ([]byte, []int) {
	return file_pb_data_data_proto_rawDescGZIP(), []int{1}
}

func (x *DataBlocks) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type SaveBlockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *SaveBlockReq) Reset() {
	*x = SaveBlockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBlockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBlockReq) ProtoMessage() {}

func (x *SaveBlockReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBlockReq.ProtoReflect.Descriptor instead.
func (*SaveBlockReq) Descriptor() ([]byte, []int) {
	return file_pb_data_data_proto_rawDescGZIP(), []int{2}
}

func (x *SaveBlockReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SaveBlockResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectID uint64 `protobuf:"varint,1,opt,name=ObjectID,proto3" json:"ObjectID,omitempty"`
}

func (x *SaveBlockResp) Reset() {
	*x = SaveBlockResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBlockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBlockResp) ProtoMessage() {}

func (x *SaveBlockResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBlockResp.ProtoReflect.Descriptor instead.
func (*SaveBlockResp) Descriptor() ([]byte, []int) {
	return file_pb_data_data_proto_rawDescGZIP(), []int{3}
}

func (x *SaveBlockResp) GetObjectID() uint64 {
	if x != nil {
		return x.ObjectID
	}
	return 0
}

type DownloadObjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectID uint64 `protobuf:"varint,1,opt,name=ObjectID,proto3" json:"ObjectID,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit    int64  `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *DownloadObjectReq) Reset() {
	*x = DownloadObjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadObjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadObjectReq) ProtoMessage() {}

func (x *DownloadObjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadObjectReq.ProtoReflect.Descriptor instead.
func (*DownloadObjectReq) Descriptor() ([]byte, []int) {
	return file_pb_data_data_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadObjectReq) GetObjectID() uint64 {
	if x != nil {
		return x.ObjectID
	}
	return 0
}

func (x *DownloadObjectReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DownloadObjectReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type DownloadObjectResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DownloadObjectResp) Reset() {
	*x = DownloadObjectResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadObjectResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadObjectResp) ProtoMessage() {}

func (x *DownloadObjectResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadObjectResp.ProtoReflect.Descriptor instead.
func (*DownloadObjectResp) Descriptor() ([]byte, []int) {
	return file_pb_data_data_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadObjectResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pb_data_data_proto protoreflect.FileDescriptor

var file_pb_data_data_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x53, 0x0a, 0x05, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x22,
	0x31, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x23, 0x0a,
	0x06, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x22, 0x5d, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x28, 0x0a, 0x12, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x85, 0x01, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x12, 0x45, 0x0a,
	0x0e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_data_data_proto_rawDescOnce sync.Once
	file_pb_data_data_proto_rawDescData = file_pb_data_data_proto_rawDesc
)

func file_pb_data_data_proto_rawDescGZIP() []byte {
	file_pb_data_data_proto_rawDescOnce.Do(func() {
		file_pb_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_data_data_proto_rawDescData)
	})
	return file_pb_data_data_proto_rawDescData
}

var file_pb_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_data_data_proto_goTypes = []interface{}{
	(*Block)(nil),              // 0: data.Block
	(*DataBlocks)(nil),         // 1: data.DataBlocks
	(*SaveBlockReq)(nil),       // 2: data.SaveBlockReq
	(*SaveBlockResp)(nil),      // 3: data.SaveBlockResp
	(*DownloadObjectReq)(nil),  // 4: data.DownloadObjectReq
	(*DownloadObjectResp)(nil), // 5: data.DownloadObjectResp
}
var file_pb_data_data_proto_depIdxs = []int32{
	0, // 0: data.DataBlocks.Blocks:type_name -> data.Block
	2, // 1: data.Data.SaveBlock:input_type -> data.SaveBlockReq
	4, // 2: data.Data.DownloadObject:input_type -> data.DownloadObjectReq
	3, // 3: data.Data.SaveBlock:output_type -> data.SaveBlockResp
	5, // 4: data.Data.DownloadObject:output_type -> data.DownloadObjectResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_data_data_proto_init() }
func file_pb_data_data_proto_init() {
	if File_pb_data_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_data_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_pb_data_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBlocks); i {
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
		file_pb_data_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBlockReq); i {
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
		file_pb_data_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBlockResp); i {
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
		file_pb_data_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadObjectReq); i {
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
		file_pb_data_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadObjectResp); i {
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
			RawDescriptor: file_pb_data_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_data_data_proto_goTypes,
		DependencyIndexes: file_pb_data_data_proto_depIdxs,
		MessageInfos:      file_pb_data_data_proto_msgTypes,
	}.Build()
	File_pb_data_data_proto = out.File
	file_pb_data_data_proto_rawDesc = nil
	file_pb_data_data_proto_goTypes = nil
	file_pb_data_data_proto_depIdxs = nil
}
