// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: chatCliente/chatCliente.proto

package chatCliente

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MensajeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje1 string `protobuf:"bytes,1,opt,name=mensaje1,proto3" json:"mensaje1,omitempty"`
}

func (x *MensajeRequest) Reset() {
	*x = MensajeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatCliente_chatCliente_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MensajeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MensajeRequest) ProtoMessage() {}

func (x *MensajeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatCliente_chatCliente_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MensajeRequest.ProtoReflect.Descriptor instead.
func (*MensajeRequest) Descriptor() ([]byte, []int) {
	return file_chatCliente_chatCliente_proto_rawDescGZIP(), []int{0}
}

func (x *MensajeRequest) GetMensaje1() string {
	if x != nil {
		return x.Mensaje1
	}
	return ""
}

type MensajeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Respuesta1 string `protobuf:"bytes,1,opt,name=respuesta1,proto3" json:"respuesta1,omitempty"`
}

func (x *MensajeReply) Reset() {
	*x = MensajeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatCliente_chatCliente_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MensajeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MensajeReply) ProtoMessage() {}

func (x *MensajeReply) ProtoReflect() protoreflect.Message {
	mi := &file_chatCliente_chatCliente_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MensajeReply.ProtoReflect.Descriptor instead.
func (*MensajeReply) Descriptor() ([]byte, []int) {
	return file_chatCliente_chatCliente_proto_rawDescGZIP(), []int{1}
}

func (x *MensajeReply) GetRespuesta1() string {
	if x != nil {
		return x.Respuesta1
	}
	return ""
}

var File_chatCliente_chatCliente_proto protoreflect.FileDescriptor

var file_chatCliente_chatCliente_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x68, 0x61, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x22, 0x2c, 0x0a, 0x0e,
	0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x31, 0x22, 0x2e, 0x0a, 0x0c, 0x4d, 0x65,
	0x6e, 0x73, 0x61, 0x6a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x31, 0x32, 0x5a, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x12, 0x47, 0x0a,
	0x0d, 0x46, 0x75, 0x6e, 0x63, 0x48, 0x6f, 0x6c, 0x61, 0x4d, 0x55, 0x6e, 0x64, 0x6f, 0x12, 0x1b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x6e,
	0x73, 0x61, 0x6a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatCliente_chatCliente_proto_rawDescOnce sync.Once
	file_chatCliente_chatCliente_proto_rawDescData = file_chatCliente_chatCliente_proto_rawDesc
)

func file_chatCliente_chatCliente_proto_rawDescGZIP() []byte {
	file_chatCliente_chatCliente_proto_rawDescOnce.Do(func() {
		file_chatCliente_chatCliente_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatCliente_chatCliente_proto_rawDescData)
	})
	return file_chatCliente_chatCliente_proto_rawDescData
}

var file_chatCliente_chatCliente_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chatCliente_chatCliente_proto_goTypes = []interface{}{
	(*MensajeRequest)(nil), // 0: chatCliente.MensajeRequest
	(*MensajeReply)(nil),   // 1: chatCliente.MensajeReply
}
var file_chatCliente_chatCliente_proto_depIdxs = []int32{
	0, // 0: chatCliente.ServicioCliente.FuncHolaMUndo:input_type -> chatCliente.MensajeRequest
	1, // 1: chatCliente.ServicioCliente.FuncHolaMUndo:output_type -> chatCliente.MensajeReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chatCliente_chatCliente_proto_init() }
func file_chatCliente_chatCliente_proto_init() {
	if File_chatCliente_chatCliente_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatCliente_chatCliente_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MensajeRequest); i {
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
		file_chatCliente_chatCliente_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MensajeReply); i {
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
			RawDescriptor: file_chatCliente_chatCliente_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatCliente_chatCliente_proto_goTypes,
		DependencyIndexes: file_chatCliente_chatCliente_proto_depIdxs,
		MessageInfos:      file_chatCliente_chatCliente_proto_msgTypes,
	}.Build()
	File_chatCliente_chatCliente_proto = out.File
	file_chatCliente_chatCliente_proto_rawDesc = nil
	file_chatCliente_chatCliente_proto_goTypes = nil
	file_chatCliente_chatCliente_proto_depIdxs = nil
}
