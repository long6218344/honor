// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: honor.proto

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

type GetHonorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHonorRequest) Reset() {
	*x = GetHonorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHonorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHonorRequest) ProtoMessage() {}

func (x *GetHonorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_honor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHonorRequest.ProtoReflect.Descriptor instead.
func (*GetHonorRequest) Descriptor() ([]byte, []int) {
	return file_honor_proto_rawDescGZIP(), []int{0}
}

func (x *GetHonorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetHonorReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetHonorReply) Reset() {
	*x = GetHonorReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHonorReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHonorReply) ProtoMessage() {}

func (x *GetHonorReply) ProtoReflect() protoreflect.Message {
	mi := &file_honor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHonorReply.ProtoReflect.Descriptor instead.
func (*GetHonorReply) Descriptor() ([]byte, []int) {
	return file_honor_proto_rawDescGZIP(), []int{1}
}

func (x *GetHonorReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_honor_proto protoreflect.FileDescriptor

var file_honor_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x6f, 0x6e, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6e, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6e, 0x6f, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0x3b, 0x0a, 0x05, 0x48,
	0x6f, 0x6e, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6e, 0x6f, 0x72,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6e, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f,
	0x6e, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x6f, 0x75, 0x73, 0x68, 0x75, 0x67, 0x75,
	0x61, 0x6e, 0x67, 0x2f, 0x68, 0x6f, 0x6e, 0x6f, 0x72, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_honor_proto_rawDescOnce sync.Once
	file_honor_proto_rawDescData = file_honor_proto_rawDesc
)

func file_honor_proto_rawDescGZIP() []byte {
	file_honor_proto_rawDescOnce.Do(func() {
		file_honor_proto_rawDescData = protoimpl.X.CompressGZIP(file_honor_proto_rawDescData)
	})
	return file_honor_proto_rawDescData
}

var file_honor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_honor_proto_goTypes = []interface{}{
	(*GetHonorRequest)(nil), // 0: pb.GetHonorRequest
	(*GetHonorReply)(nil),   // 1: pb.GetHonorReply
}
var file_honor_proto_depIdxs = []int32{
	0, // 0: pb.Honor.GetHonor:input_type -> pb.GetHonorRequest
	1, // 1: pb.Honor.GetHonor:output_type -> pb.GetHonorReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_honor_proto_init() }
func file_honor_proto_init() {
	if File_honor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_honor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHonorRequest); i {
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
		file_honor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHonorReply); i {
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
			RawDescriptor: file_honor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_honor_proto_goTypes,
		DependencyIndexes: file_honor_proto_depIdxs,
		MessageInfos:      file_honor_proto_msgTypes,
	}.Build()
	File_honor_proto = out.File
	file_honor_proto_rawDesc = nil
	file_honor_proto_goTypes = nil
	file_honor_proto_depIdxs = nil
}
