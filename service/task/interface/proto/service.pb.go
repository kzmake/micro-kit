// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: service/task/interface/proto/service.proto

package proto

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description *wrappers.StringValue `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_task_interface_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_task_interface_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_service_task_interface_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetDescription() *wrappers.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Task `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_task_interface_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_task_interface_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_service_task_interface_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetResult() *Task {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_task_interface_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_task_interface_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_service_task_interface_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetId() *wrappers.StringValue {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Task `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_task_interface_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_task_interface_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_service_task_interface_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetResult() *Task {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_service_task_interface_proto_service_proto protoreflect.FileDescriptor

var file_service_task_interface_proto_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6b, 0x7a,
	0x6d, 0x61, 0x6b, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x00, 0x18, 0xff, 0x01, 0x28, 0xff, 0x01, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6b, 0x7a, 0x6d, 0x61, 0x6b, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xfa, 0x42,
	0x08, 0x72, 0x06, 0x10, 0x1a, 0x18, 0x1a, 0x28, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b,
	0x7a, 0x6d, 0x61, 0x6b, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0xba, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e,
	0x6b, 0x7a, 0x6d, 0x61, 0x6b, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x7a, 0x6d, 0x61, 0x6b, 0x65, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x6b, 0x7a, 0x6d, 0x61, 0x6b, 0x65, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6b, 0x7a, 0x6d,
	0x61, 0x6b, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6b, 0x69, 0x74, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x7a, 0x6d, 0x61, 0x6b, 0x65, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_task_interface_proto_service_proto_rawDescOnce sync.Once
	file_service_task_interface_proto_service_proto_rawDescData = file_service_task_interface_proto_service_proto_rawDesc
)

func file_service_task_interface_proto_service_proto_rawDescGZIP() []byte {
	file_service_task_interface_proto_service_proto_rawDescOnce.Do(func() {
		file_service_task_interface_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_task_interface_proto_service_proto_rawDescData)
	})
	return file_service_task_interface_proto_service_proto_rawDescData
}

var file_service_task_interface_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_task_interface_proto_service_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),        // 0: kzmake.microkit.task.v1.CreateRequest
	(*CreateResponse)(nil),       // 1: kzmake.microkit.task.v1.CreateResponse
	(*GetRequest)(nil),           // 2: kzmake.microkit.task.v1.GetRequest
	(*GetResponse)(nil),          // 3: kzmake.microkit.task.v1.GetResponse
	(*wrappers.StringValue)(nil), // 4: google.protobuf.StringValue
	(*Task)(nil),                 // 5: kzmake.microkit.task.v1.Task
}
var file_service_task_interface_proto_service_proto_depIdxs = []int32{
	4, // 0: kzmake.microkit.task.v1.CreateRequest.description:type_name -> google.protobuf.StringValue
	5, // 1: kzmake.microkit.task.v1.CreateResponse.result:type_name -> kzmake.microkit.task.v1.Task
	4, // 2: kzmake.microkit.task.v1.GetRequest.id:type_name -> google.protobuf.StringValue
	5, // 3: kzmake.microkit.task.v1.GetResponse.result:type_name -> kzmake.microkit.task.v1.Task
	0, // 4: kzmake.microkit.task.v1.TaskService.Create:input_type -> kzmake.microkit.task.v1.CreateRequest
	2, // 5: kzmake.microkit.task.v1.TaskService.Get:input_type -> kzmake.microkit.task.v1.GetRequest
	1, // 6: kzmake.microkit.task.v1.TaskService.Create:output_type -> kzmake.microkit.task.v1.CreateResponse
	3, // 7: kzmake.microkit.task.v1.TaskService.Get:output_type -> kzmake.microkit.task.v1.GetResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_task_interface_proto_service_proto_init() }
func file_service_task_interface_proto_service_proto_init() {
	if File_service_task_interface_proto_service_proto != nil {
		return
	}
	file_service_task_interface_proto_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_task_interface_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_service_task_interface_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_service_task_interface_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_service_task_interface_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
			RawDescriptor: file_service_task_interface_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_task_interface_proto_service_proto_goTypes,
		DependencyIndexes: file_service_task_interface_proto_service_proto_depIdxs,
		MessageInfos:      file_service_task_interface_proto_service_proto_msgTypes,
	}.Build()
	File_service_task_interface_proto_service_proto = out.File
	file_service_task_interface_proto_service_proto_rawDesc = nil
	file_service_task_interface_proto_service_proto_goTypes = nil
	file_service_task_interface_proto_service_proto_depIdxs = nil
}
