// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: core.proto

package core

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
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
	return file_core_proto_rawDescGZIP(), []int{0}
}

type CfgDataApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DSN string `protobuf:"bytes,1,opt,name=DSN,proto3" json:"DSN,omitempty"`
}

func (x *CfgDataApp) Reset() {
	*x = CfgDataApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgDataApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgDataApp) ProtoMessage() {}

func (x *CfgDataApp) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgDataApp.ProtoReflect.Descriptor instead.
func (*CfgDataApp) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *CfgDataApp) GetDSN() string {
	if x != nil {
		return x.DSN
	}
	return ""
}

type CfgDataAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DSN string `protobuf:"bytes,1,opt,name=DSN,proto3" json:"DSN,omitempty"`
}

func (x *CfgDataAuth) Reset() {
	*x = CfgDataAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgDataAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgDataAuth) ProtoMessage() {}

func (x *CfgDataAuth) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgDataAuth.ProtoReflect.Descriptor instead.
func (*CfgDataAuth) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{2}
}

func (x *CfgDataAuth) GetDSN() string {
	if x != nil {
		return x.DSN
	}
	return ""
}

type CfgDataGate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DSN string `protobuf:"bytes,1,opt,name=DSN,proto3" json:"DSN,omitempty"`
}

func (x *CfgDataGate) Reset() {
	*x = CfgDataGate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgDataGate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgDataGate) ProtoMessage() {}

func (x *CfgDataGate) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgDataGate.ProtoReflect.Descriptor instead.
func (*CfgDataGate) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{3}
}

func (x *CfgDataGate) GetDSN() string {
	if x != nil {
		return x.DSN
	}
	return ""
}

type CfgApi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CfgApi) Reset() {
	*x = CfgApi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgApi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgApi) ProtoMessage() {}

func (x *CfgApi) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgApi.ProtoReflect.Descriptor instead.
func (*CfgApi) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{4}
}

type CfgApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CfgApp) Reset() {
	*x = CfgApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgApp) ProtoMessage() {}

func (x *CfgApp) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgApp.ProtoReflect.Descriptor instead.
func (*CfgApp) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{5}
}

type CfgAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashCost int32  `protobuf:"varint,1,opt,name=HashCost,proto3" json:"HashCost,omitempty"`
	Salt     string `protobuf:"bytes,2,opt,name=Salt,proto3" json:"Salt,omitempty"`
}

func (x *CfgAuth) Reset() {
	*x = CfgAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgAuth) ProtoMessage() {}

func (x *CfgAuth) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgAuth.ProtoReflect.Descriptor instead.
func (*CfgAuth) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{6}
}

func (x *CfgAuth) GetHashCost() int32 {
	if x != nil {
		return x.HashCost
	}
	return 0
}

func (x *CfgAuth) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

type CfgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App  *CfgDataApp  `protobuf:"bytes,1,opt,name=App,proto3" json:"App,omitempty"`
	Auth *CfgDataAuth `protobuf:"bytes,2,opt,name=Auth,proto3" json:"Auth,omitempty"`
	Gate *CfgDataGate `protobuf:"bytes,3,opt,name=Gate,proto3" json:"Gate,omitempty"`
}

func (x *CfgData) Reset() {
	*x = CfgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgData) ProtoMessage() {}

func (x *CfgData) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgData.ProtoReflect.Descriptor instead.
func (*CfgData) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{7}
}

func (x *CfgData) GetApp() *CfgDataApp {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *CfgData) GetAuth() *CfgDataAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *CfgData) GetGate() *CfgDataGate {
	if x != nil {
		return x.Gate
	}
	return nil
}

type CfgGate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CfgGate) Reset() {
	*x = CfgGate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgGate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgGate) ProtoMessage() {}

func (x *CfgGate) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgGate.ProtoReflect.Descriptor instead.
func (*CfgGate) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{8}
}

type CfgWeb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertFile string `protobuf:"bytes,1,opt,name=certFile,proto3" json:"certFile,omitempty"`
	KeyFile  string `protobuf:"bytes,2,opt,name=keyFile,proto3" json:"keyFile,omitempty"`
}

func (x *CfgWeb) Reset() {
	*x = CfgWeb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CfgWeb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CfgWeb) ProtoMessage() {}

func (x *CfgWeb) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CfgWeb.ProtoReflect.Descriptor instead.
func (*CfgWeb) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{9}
}

func (x *CfgWeb) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *CfgWeb) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a,
	0x0a, 0x43, 0x66, 0x67, 0x44, 0x61, 0x74, 0x61, 0x41, 0x70, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x44,
	0x53, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x53, 0x4e, 0x22, 0x1f, 0x0a,
	0x0b, 0x43, 0x66, 0x67, 0x44, 0x61, 0x74, 0x61, 0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x44, 0x53, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x53, 0x4e, 0x22, 0x1f,
	0x0a, 0x0b, 0x43, 0x66, 0x67, 0x44, 0x61, 0x74, 0x61, 0x47, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x44, 0x53, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x53, 0x4e, 0x22,
	0x08, 0x0a, 0x06, 0x43, 0x66, 0x67, 0x41, 0x70, 0x69, 0x22, 0x08, 0x0a, 0x06, 0x43, 0x66, 0x67,
	0x41, 0x70, 0x70, 0x22, 0x39, 0x0a, 0x07, 0x43, 0x66, 0x67, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x61,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x61, 0x6c, 0x74, 0x22, 0x7b,
	0x0a, 0x07, 0x43, 0x66, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x03, 0x41, 0x70, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x66,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x41, 0x70, 0x70, 0x52, 0x03, 0x41, 0x70, 0x70, 0x12, 0x25, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x43, 0x66, 0x67, 0x44, 0x61, 0x74, 0x61, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x04, 0x47, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x66, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x47, 0x61, 0x74, 0x65, 0x52, 0x04, 0x47, 0x61, 0x74, 0x65, 0x22, 0x09, 0x0a, 0x07, 0x43,
	0x66, 0x67, 0x47, 0x61, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x43, 0x66, 0x67, 0x57, 0x65, 0x62,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x61, 0x6b, 0x69, 0x72, 0x69, 0x2f, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_core_proto_goTypes = []interface{}{
	(*Request)(nil),     // 0: core.Request
	(*CfgDataApp)(nil),  // 1: core.CfgDataApp
	(*CfgDataAuth)(nil), // 2: core.CfgDataAuth
	(*CfgDataGate)(nil), // 3: core.CfgDataGate
	(*CfgApi)(nil),      // 4: core.CfgApi
	(*CfgApp)(nil),      // 5: core.CfgApp
	(*CfgAuth)(nil),     // 6: core.CfgAuth
	(*CfgData)(nil),     // 7: core.CfgData
	(*CfgGate)(nil),     // 8: core.CfgGate
	(*CfgWeb)(nil),      // 9: core.CfgWeb
}
var file_core_proto_depIdxs = []int32{
	1, // 0: core.CfgData.App:type_name -> core.CfgDataApp
	2, // 1: core.CfgData.Auth:type_name -> core.CfgDataAuth
	3, // 2: core.CfgData.Gate:type_name -> core.CfgDataGate
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgDataApp); i {
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
		file_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgDataAuth); i {
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
		file_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgDataGate); i {
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
		file_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgApi); i {
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
		file_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgApp); i {
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
		file_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgAuth); i {
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
		file_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgData); i {
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
		file_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgGate); i {
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
		file_core_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CfgWeb); i {
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
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}