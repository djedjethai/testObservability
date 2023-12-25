// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: service/proto/v1/Operand.proto

package operand

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operand []float32 `protobuf:"fixed32,1,rep,packed,name=Operand,proto3" json:"Operand,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_Operand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_Operand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_Operand_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetOperand() []float32 {
	if x != nil {
		return x.Operand
	}
	return nil
}

var File_service_proto_v1_Operand_proto protoreflect.FileDescriptor

var file_service_proto_v1_Operand_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x76, 0x31, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x6c, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x07, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x32,
	0x53, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64,
	0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_proto_v1_Operand_proto_rawDescOnce sync.Once
	file_service_proto_v1_Operand_proto_rawDescData = file_service_proto_v1_Operand_proto_rawDesc
)

func file_service_proto_v1_Operand_proto_rawDescGZIP() []byte {
	file_service_proto_v1_Operand_proto_rawDescOnce.Do(func() {
		file_service_proto_v1_Operand_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_v1_Operand_proto_rawDescData)
	})
	return file_service_proto_v1_Operand_proto_rawDescData
}

var file_service_proto_v1_Operand_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_proto_v1_Operand_proto_goTypes = []interface{}{
	(*Data)(nil),                   // 0: v1_operandl.Data
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_service_proto_v1_Operand_proto_depIdxs = []int32{
	0, // 0: v1_operandl.OperandManagement.SendOperand:input_type -> v1_operandl.Data
	1, // 1: v1_operandl.OperandManagement.SendOperand:output_type -> google.protobuf.StringValue
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_v1_Operand_proto_init() }
func file_service_proto_v1_Operand_proto_init() {
	if File_service_proto_v1_Operand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_v1_Operand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_service_proto_v1_Operand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_v1_Operand_proto_goTypes,
		DependencyIndexes: file_service_proto_v1_Operand_proto_depIdxs,
		MessageInfos:      file_service_proto_v1_Operand_proto_msgTypes,
	}.Build()
	File_service_proto_v1_Operand_proto = out.File
	file_service_proto_v1_Operand_proto_rawDesc = nil
	file_service_proto_v1_Operand_proto_goTypes = nil
	file_service_proto_v1_Operand_proto_depIdxs = nil
}
