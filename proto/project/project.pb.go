// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/project/project.proto

package project

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProjectRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	mi := &file_proto_project_project_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_project_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_project_project_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateProjectResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProjectResponse) Reset() {
	*x = CreateProjectResponse{}
	mi := &file_proto_project_project_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectResponse) ProtoMessage() {}

func (x *CreateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_project_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectResponse.ProtoReflect.Descriptor instead.
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_proto_project_project_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProjectResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetProjectRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProjectRequest) Reset() {
	*x = GetProjectRequest{}
	mi := &file_proto_project_project_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectRequest) ProtoMessage() {}

func (x *GetProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_project_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectRequest.ProtoReflect.Descriptor instead.
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_project_project_proto_rawDescGZIP(), []int{2}
}

func (x *GetProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProjectResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProjectResponse) Reset() {
	*x = GetProjectResponse{}
	mi := &file_proto_project_project_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectResponse) ProtoMessage() {}

func (x *GetProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_project_project_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectResponse.ProtoReflect.Descriptor instead.
func (*GetProjectResponse) Descriptor() ([]byte, []int) {
	return file_proto_project_project_proto_rawDescGZIP(), []int{3}
}

func (x *GetProjectResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProjectResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProjectResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_proto_project_project_proto protoreflect.FileDescriptor

const file_proto_project_project_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/project/project.proto\x12\aproject\"L\n" +
	"\x14CreateProjectRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\"]\n" +
	"\x15CreateProjectResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"#\n" +
	"\x11GetProjectRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"Z\n" +
	"\x12GetProjectResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription2\xa7\x01\n" +
	"\x0eProjectService\x12N\n" +
	"\rCreateProject\x12\x1d.project.CreateProjectRequest\x1a\x1e.project.CreateProjectResponse\x12E\n" +
	"\n" +
	"GetProject\x12\x1a.project.GetProjectRequest\x1a\x1b.project.GetProjectResponseB\x0fZ\rproto/projectb\x06proto3"

var (
	file_proto_project_project_proto_rawDescOnce sync.Once
	file_proto_project_project_proto_rawDescData []byte
)

func file_proto_project_project_proto_rawDescGZIP() []byte {
	file_proto_project_project_proto_rawDescOnce.Do(func() {
		file_proto_project_project_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_project_project_proto_rawDesc), len(file_proto_project_project_proto_rawDesc)))
	})
	return file_proto_project_project_proto_rawDescData
}

var file_proto_project_project_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_project_project_proto_goTypes = []any{
	(*CreateProjectRequest)(nil),  // 0: project.CreateProjectRequest
	(*CreateProjectResponse)(nil), // 1: project.CreateProjectResponse
	(*GetProjectRequest)(nil),     // 2: project.GetProjectRequest
	(*GetProjectResponse)(nil),    // 3: project.GetProjectResponse
}
var file_proto_project_project_proto_depIdxs = []int32{
	0, // 0: project.ProjectService.CreateProject:input_type -> project.CreateProjectRequest
	2, // 1: project.ProjectService.GetProject:input_type -> project.GetProjectRequest
	1, // 2: project.ProjectService.CreateProject:output_type -> project.CreateProjectResponse
	3, // 3: project.ProjectService.GetProject:output_type -> project.GetProjectResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_project_project_proto_init() }
func file_proto_project_project_proto_init() {
	if File_proto_project_project_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_project_project_proto_rawDesc), len(file_proto_project_project_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_project_project_proto_goTypes,
		DependencyIndexes: file_proto_project_project_proto_depIdxs,
		MessageInfos:      file_proto_project_project_proto_msgTypes,
	}.Build()
	File_proto_project_project_proto = out.File
	file_proto_project_project_proto_goTypes = nil
	file_proto_project_project_proto_depIdxs = nil
}
