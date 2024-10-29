// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: store/collection.proto

package store

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

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId   int32      `protobuf:"varint,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatedTs   int64      `protobuf:"varint,3,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`
	UpdatedTs   int64      `protobuf:"varint,4,opt,name=updated_ts,json=updatedTs,proto3" json:"updated_ts,omitempty"`
	Name        string     `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Title       string     `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	ShortcutIds []int32    `protobuf:"varint,9,rep,packed,name=shortcut_ids,json=shortcutIds,proto3" json:"shortcut_ids,omitempty"`
	Visibility  Visibility `protobuf:"varint,10,opt,name=visibility,proto3,enum=slash.store.Visibility" json:"visibility,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	mi := &file_store_collection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_store_collection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_store_collection_proto_rawDescGZIP(), []int{0}
}

func (x *Collection) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Collection) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Collection) GetCreatedTs() int64 {
	if x != nil {
		return x.CreatedTs
	}
	return 0
}

func (x *Collection) GetUpdatedTs() int64 {
	if x != nil {
		return x.UpdatedTs
	}
	return 0
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Collection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Collection) GetShortcutIds() []int32 {
	if x != nil {
		return x.ShortcutIds
	}
	return nil
}

func (x *Collection) GetVisibility() Visibility {
	if x != nil {
		return x.Visibility
	}
	return Visibility_VISIBILITY_UNSPECIFIED
}

var File_store_collection_proto protoreflect.FileDescriptor

var file_store_collection_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0a, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75,
	0x74, 0x49, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0xa0, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x42, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6c, 0x66, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x2f,
	0x73, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0b, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_collection_proto_rawDescOnce sync.Once
	file_store_collection_proto_rawDescData = file_store_collection_proto_rawDesc
)

func file_store_collection_proto_rawDescGZIP() []byte {
	file_store_collection_proto_rawDescOnce.Do(func() {
		file_store_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_collection_proto_rawDescData)
	})
	return file_store_collection_proto_rawDescData
}

var file_store_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_collection_proto_goTypes = []any{
	(*Collection)(nil), // 0: slash.store.Collection
	(Visibility)(0),    // 1: slash.store.Visibility
}
var file_store_collection_proto_depIdxs = []int32{
	1, // 0: slash.store.Collection.visibility:type_name -> slash.store.Visibility
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_store_collection_proto_init() }
func file_store_collection_proto_init() {
	if File_store_collection_proto != nil {
		return
	}
	file_store_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_collection_proto_goTypes,
		DependencyIndexes: file_store_collection_proto_depIdxs,
		MessageInfos:      file_store_collection_proto_msgTypes,
	}.Build()
	File_store_collection_proto = out.File
	file_store_collection_proto_rawDesc = nil
	file_store_collection_proto_goTypes = nil
	file_store_collection_proto_depIdxs = nil
}
