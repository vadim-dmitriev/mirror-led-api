// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: led-service.proto

package led_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type LightLEDRequest_LightMode int32

const (
	LightLEDRequest_NO_MODE  LightLEDRequest_LightMode = 0
	LightLEDRequest_KEY_WORD LightLEDRequest_LightMode = 1
)

// Enum value maps for LightLEDRequest_LightMode.
var (
	LightLEDRequest_LightMode_name = map[int32]string{
		0: "NO_MODE",
		1: "KEY_WORD",
	}
	LightLEDRequest_LightMode_value = map[string]int32{
		"NO_MODE":  0,
		"KEY_WORD": 1,
	}
)

func (x LightLEDRequest_LightMode) Enum() *LightLEDRequest_LightMode {
	p := new(LightLEDRequest_LightMode)
	*p = x
	return p
}

func (x LightLEDRequest_LightMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LightLEDRequest_LightMode) Descriptor() protoreflect.EnumDescriptor {
	return file_led_service_proto_enumTypes[0].Descriptor()
}

func (LightLEDRequest_LightMode) Type() protoreflect.EnumType {
	return &file_led_service_proto_enumTypes[0]
}

func (x LightLEDRequest_LightMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LightLEDRequest_LightMode.Descriptor instead.
func (LightLEDRequest_LightMode) EnumDescriptor() ([]byte, []int) {
	return file_led_service_proto_rawDescGZIP(), []int{0, 0}
}

type LightLEDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode LightLEDRequest_LightMode `protobuf:"varint,1,opt,name=mode,proto3,enum=led_service.LightLEDRequest_LightMode" json:"mode,omitempty"`
}

func (x *LightLEDRequest) Reset() {
	*x = LightLEDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_led_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightLEDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightLEDRequest) ProtoMessage() {}

func (x *LightLEDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_led_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightLEDRequest.ProtoReflect.Descriptor instead.
func (*LightLEDRequest) Descriptor() ([]byte, []int) {
	return file_led_service_proto_rawDescGZIP(), []int{0}
}

func (x *LightLEDRequest) GetMode() LightLEDRequest_LightMode {
	if x != nil {
		return x.Mode
	}
	return LightLEDRequest_NO_MODE
}

type LightLEDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LightLEDResponse) Reset() {
	*x = LightLEDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_led_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightLEDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightLEDResponse) ProtoMessage() {}

func (x *LightLEDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_led_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightLEDResponse.ProtoReflect.Descriptor instead.
func (*LightLEDResponse) Descriptor() ([]byte, []int) {
	return file_led_service_proto_rawDescGZIP(), []int{1}
}

func (x *LightLEDResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_led_service_proto protoreflect.FileDescriptor

var file_led_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x65, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a,
	0x0f, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x45, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x6c, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x4c, 0x45, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x26, 0x0a, 0x09,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x59, 0x5f, 0x57, 0x4f,
	0x52, 0x44, 0x10, 0x01, 0x22, 0x2c, 0x0a, 0x10, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x45, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0x96, 0x01, 0x0a, 0x0a, 0x4c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x08, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x45, 0x44, 0x12, 0x1c, 0x2e,
	0x6c, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x67, 0x68,
	0x74, 0x4c, 0x45, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x4c,
	0x45, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09,
	0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4c, 0x45, 0x44, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70,
	0x6b, 0x67, 0x2f, 0x6c, 0x65, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_led_service_proto_rawDescOnce sync.Once
	file_led_service_proto_rawDescData = file_led_service_proto_rawDesc
)

func file_led_service_proto_rawDescGZIP() []byte {
	file_led_service_proto_rawDescOnce.Do(func() {
		file_led_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_led_service_proto_rawDescData)
	})
	return file_led_service_proto_rawDescData
}

var file_led_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_led_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_led_service_proto_goTypes = []interface{}{
	(LightLEDRequest_LightMode)(0), // 0: led_service.LightLEDRequest.LightMode
	(*LightLEDRequest)(nil),        // 1: led_service.LightLEDRequest
	(*LightLEDResponse)(nil),       // 2: led_service.LightLEDResponse
	(*empty.Empty)(nil),            // 3: google.protobuf.Empty
}
var file_led_service_proto_depIdxs = []int32{
	0, // 0: led_service.LightLEDRequest.mode:type_name -> led_service.LightLEDRequest.LightMode
	1, // 1: led_service.LedService.LightLED:input_type -> led_service.LightLEDRequest
	3, // 2: led_service.LedService.SwitchLED:input_type -> google.protobuf.Empty
	2, // 3: led_service.LedService.LightLED:output_type -> led_service.LightLEDResponse
	3, // 4: led_service.LedService.SwitchLED:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_led_service_proto_init() }
func file_led_service_proto_init() {
	if File_led_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_led_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightLEDRequest); i {
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
		file_led_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightLEDResponse); i {
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
			RawDescriptor: file_led_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_led_service_proto_goTypes,
		DependencyIndexes: file_led_service_proto_depIdxs,
		EnumInfos:         file_led_service_proto_enumTypes,
		MessageInfos:      file_led_service_proto_msgTypes,
	}.Build()
	File_led_service_proto = out.File
	file_led_service_proto_rawDesc = nil
	file_led_service_proto_goTypes = nil
	file_led_service_proto_depIdxs = nil
}
