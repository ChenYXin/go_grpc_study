// 指定proto版本

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: stream.proto

// 指定默认包名

package stream_proto

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

// 请求参数
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 文件回调参数
type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 字节数据类型，即文件内容、数据
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *FileResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x56, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x45, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0f, 0x5a,
	0x0d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stream_proto_goTypes = []interface{}{
	(*Request)(nil),      // 0: stream_proto.Request
	(*FileResponse)(nil), // 1: stream_proto.FileResponse
}
var file_stream_proto_depIdxs = []int32{
	0, // 0: stream_proto.ServiceStream.DownloadFile:input_type -> stream_proto.Request
	1, // 1: stream_proto.ServiceStream.DownloadFile:output_type -> stream_proto.FileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}
