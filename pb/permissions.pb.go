// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pb/permissions.proto

package pb

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

type CheckDirectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Entity     string `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Permission string `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	Object     string `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *CheckDirectReq) Reset() {
	*x = CheckDirectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckDirectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDirectReq) ProtoMessage() {}

func (x *CheckDirectReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDirectReq.ProtoReflect.Descriptor instead.
func (*CheckDirectReq) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *CheckDirectReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CheckDirectReq) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *CheckDirectReq) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *CheckDirectReq) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type CheckDirectRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Valid bool   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *CheckDirectRes) Reset() {
	*x = CheckDirectRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckDirectRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDirectRes) ProtoMessage() {}

func (x *CheckDirectRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDirectRes.ProtoReflect.Descriptor instead.
func (*CheckDirectRes) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *CheckDirectRes) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CheckDirectRes) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type ListEntityRelationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Entity     string `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Permission string `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *ListEntityRelationsReq) Reset() {
	*x = ListEntityRelationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntityRelationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntityRelationsReq) ProtoMessage() {}

func (x *ListEntityRelationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntityRelationsReq.ProtoReflect.Descriptor instead.
func (*ListEntityRelationsReq) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *ListEntityRelationsReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListEntityRelationsReq) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *ListEntityRelationsReq) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type ListObjectRelationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Object     string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Permission string `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *ListObjectRelationsReq) Reset() {
	*x = ListObjectRelationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectRelationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectRelationsReq) ProtoMessage() {}

func (x *ListObjectRelationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectRelationsReq.ProtoReflect.Descriptor instead.
func (*ListObjectRelationsReq) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{3}
}

func (x *ListObjectRelationsReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListObjectRelationsReq) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ListObjectRelationsReq) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type RelationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relations []*Relation `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations,omitempty"`
}

func (x *RelationsResponse) Reset() {
	*x = RelationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationsResponse) ProtoMessage() {}

func (x *RelationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationsResponse.ProtoReflect.Descriptor instead.
func (*RelationsResponse) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{4}
}

func (x *RelationsResponse) GetRelations() []*Relation {
	if x != nil {
		return x.Relations
	}
	return nil
}

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity     string `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	Object     string `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_permissions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_permissions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_pb_permissions_proto_rawDescGZIP(), []int{5}
}

func (x *Relation) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *Relation) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *Relation) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

var File_pb_permissions_proto protoreflect.FileDescriptor

var file_pb_permissions_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x11,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5a, 0x0a, 0x08, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_permissions_proto_rawDescOnce sync.Once
	file_pb_permissions_proto_rawDescData = file_pb_permissions_proto_rawDesc
)

func file_pb_permissions_proto_rawDescGZIP() []byte {
	file_pb_permissions_proto_rawDescOnce.Do(func() {
		file_pb_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_permissions_proto_rawDescData)
	})
	return file_pb_permissions_proto_rawDescData
}

var file_pb_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_permissions_proto_goTypes = []interface{}{
	(*CheckDirectReq)(nil),         // 0: CheckDirectReq
	(*CheckDirectRes)(nil),         // 1: CheckDirectRes
	(*ListEntityRelationsReq)(nil), // 2: ListEntityRelationsReq
	(*ListObjectRelationsReq)(nil), // 3: ListObjectRelationsReq
	(*RelationsResponse)(nil),      // 4: RelationsResponse
	(*Relation)(nil),               // 5: Relation
}
var file_pb_permissions_proto_depIdxs = []int32{
	5, // 0: RelationsResponse.relations:type_name -> Relation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_permissions_proto_init() }
func file_pb_permissions_proto_init() {
	if File_pb_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckDirectReq); i {
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
		file_pb_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckDirectRes); i {
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
		file_pb_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntityRelationsReq); i {
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
		file_pb_permissions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectRelationsReq); i {
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
		file_pb_permissions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationsResponse); i {
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
		file_pb_permissions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relation); i {
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
			RawDescriptor: file_pb_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_permissions_proto_goTypes,
		DependencyIndexes: file_pb_permissions_proto_depIdxs,
		MessageInfos:      file_pb_permissions_proto_msgTypes,
	}.Build()
	File_pb_permissions_proto = out.File
	file_pb_permissions_proto_rawDesc = nil
	file_pb_permissions_proto_goTypes = nil
	file_pb_permissions_proto_depIdxs = nil
}
