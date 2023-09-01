// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: api/api.proto

package api

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

type RequestObserved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalAddr  string `protobuf:"bytes,1,opt,name=local_addr,json=localAddr,proto3" json:"local_addr,omitempty"`
	RemoteAddr string `protobuf:"bytes,2,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	L4Protocol string `protobuf:"bytes,3,opt,name=l4_protocol,json=l4Protocol,proto3" json:"l4_protocol,omitempty"`
	L7Protocol string `protobuf:"bytes,4,opt,name=l7_protocol,json=l7Protocol,proto3" json:"l7_protocol,omitempty"`
	Request    []byte `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	Response   []byte `protobuf:"bytes,6,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RequestObserved) Reset() {
	*x = RequestObserved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestObserved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestObserved) ProtoMessage() {}

func (x *RequestObserved) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestObserved.ProtoReflect.Descriptor instead.
func (*RequestObserved) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *RequestObserved) GetLocalAddr() string {
	if x != nil {
		return x.LocalAddr
	}
	return ""
}

func (x *RequestObserved) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *RequestObserved) GetL4Protocol() string {
	if x != nil {
		return x.L4Protocol
	}
	return ""
}

func (x *RequestObserved) GetL7Protocol() string {
	if x != nil {
		return x.L7Protocol
	}
	return ""
}

func (x *RequestObserved) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *RequestObserved) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AgentStarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AgentStarted) Reset() {
	*x = AgentStarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentStarted) ProtoMessage() {}

func (x *AgentStarted) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentStarted.ProtoReflect.Descriptor instead.
func (*AgentStarted) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0xc9, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x34, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c,
	0x34, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x37, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x37, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x32, 0x80, 0x01, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x44, 0x6f, 0x67, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x76, 0x61, 0x6e, 0x72, 0x6f, 0x6c, 0x66, 0x65, 0x2f, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_api_proto_goTypes = []interface{}{
	(*RequestObserved)(nil), // 0: api.RequestObserved
	(*Reply)(nil),           // 1: api.Reply
	(*AgentStarted)(nil),    // 2: api.AgentStarted
}
var file_api_api_proto_depIdxs = []int32{
	0, // 0: api.DockerDogAgent.SendRequestObserved:input_type -> api.RequestObserved
	2, // 1: api.DockerDogAgent.SendAgentStarted:input_type -> api.AgentStarted
	1, // 2: api.DockerDogAgent.SendRequestObserved:output_type -> api.Reply
	1, // 3: api.DockerDogAgent.SendAgentStarted:output_type -> api.Reply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestObserved); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentStarted); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}