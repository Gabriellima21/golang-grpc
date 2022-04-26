// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: publicacao.proto

package publicacao

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

type UpvoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpvoteRequest) Reset() {
	*x = UpvoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicacao_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpvoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpvoteRequest) ProtoMessage() {}

func (x *UpvoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publicacao_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpvoteRequest.ProtoReflect.Descriptor instead.
func (*UpvoteRequest) Descriptor() ([]byte, []int) {
	return file_publicacao_proto_rawDescGZIP(), []int{0}
}

func (x *UpvoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PublicacaoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PublicacaoRequest) Reset() {
	*x = PublicacaoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicacao_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicacaoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicacaoRequest) ProtoMessage() {}

func (x *PublicacaoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publicacao_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicacaoRequest.ProtoReflect.Descriptor instead.
func (*PublicacaoRequest) Descriptor() ([]byte, []int) {
	return file_publicacao_proto_rawDescGZIP(), []int{1}
}

func (x *PublicacaoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublicacaoRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PublicacaoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Likes   int64  `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *PublicacaoResponse) Reset() {
	*x = PublicacaoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publicacao_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicacaoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicacaoResponse) ProtoMessage() {}

func (x *PublicacaoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publicacao_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicacaoResponse.ProtoReflect.Descriptor instead.
func (*PublicacaoResponse) Descriptor() ([]byte, []int) {
	return file_publicacao_proto_rawDescGZIP(), []int{2}
}

func (x *PublicacaoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PublicacaoResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublicacaoResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PublicacaoResponse) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

var File_publicacao_proto protoreflect.FileDescriptor

var file_publicacao_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x22, 0x1f,
	0x0a, 0x0d, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x43, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x6a, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63,
	0x61, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x32, 0xb9, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x12, 0x19, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63,
	0x61, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61,
	0x63, 0x61, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63,
	0x61, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17,
	0x2e, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x61, 0x63, 0x61, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publicacao_proto_rawDescOnce sync.Once
	file_publicacao_proto_rawDescData = file_publicacao_proto_rawDesc
)

func file_publicacao_proto_rawDescGZIP() []byte {
	file_publicacao_proto_rawDescOnce.Do(func() {
		file_publicacao_proto_rawDescData = protoimpl.X.CompressGZIP(file_publicacao_proto_rawDescData)
	})
	return file_publicacao_proto_rawDescData
}

var file_publicacao_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_publicacao_proto_goTypes = []interface{}{
	(*UpvoteRequest)(nil),      // 0: publicacao.UpvoteRequest
	(*PublicacaoRequest)(nil),  // 1: publicacao.PublicacaoRequest
	(*PublicacaoResponse)(nil), // 2: publicacao.PublicacaoResponse
}
var file_publicacao_proto_depIdxs = []int32{
	0, // 0: publicacao.PublicacaoService.UpvotePublicacao:input_type -> publicacao.UpvoteRequest
	1, // 1: publicacao.PublicacaoService.CreatePublicacao:input_type -> publicacao.PublicacaoRequest
	2, // 2: publicacao.PublicacaoService.UpvotePublicacao:output_type -> publicacao.PublicacaoResponse
	2, // 3: publicacao.PublicacaoService.CreatePublicacao:output_type -> publicacao.PublicacaoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_publicacao_proto_init() }
func file_publicacao_proto_init() {
	if File_publicacao_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publicacao_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpvoteRequest); i {
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
		file_publicacao_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicacaoRequest); i {
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
		file_publicacao_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicacaoResponse); i {
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
			RawDescriptor: file_publicacao_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publicacao_proto_goTypes,
		DependencyIndexes: file_publicacao_proto_depIdxs,
		MessageInfos:      file_publicacao_proto_msgTypes,
	}.Build()
	File_publicacao_proto = out.File
	file_publicacao_proto_rawDesc = nil
	file_publicacao_proto_goTypes = nil
	file_publicacao_proto_depIdxs = nil
}
