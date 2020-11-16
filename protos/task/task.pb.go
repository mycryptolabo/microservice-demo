// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: protos/task/task.proto

package task

import (
	project "github.com/Joker666/microservice-demo/protos/project"
	user "github.com/Joker666/microservice-demo/protos/user"
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

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProjectId      string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	TagId          string `protobuf:"bytes,4,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	AssignedUserId string `protobuf:"bytes,5,opt,name=assigned_user_id,json=assignedUserId,proto3" json:"assigned_user_id,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_protos_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTaskRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateTaskRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *CreateTaskRequest) GetAssignedUserId() string {
	if x != nil {
		return x.AssignedUserId
	}
	return ""
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string                   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name         string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Project      *project.ProjectResponse `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Tag          *project.TagResponse     `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	AssignedUser *user.UserResponse       `protobuf:"bytes,5,opt,name=assigned_user,json=assignedUser,proto3" json:"assigned_user,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_protos_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TaskResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskResponse) GetProject() *project.ProjectResponse {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *TaskResponse) GetTag() *project.TagResponse {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *TaskResponse) GetAssignedUser() *user.UserResponse {
	if x != nil {
		return x.AssignedUser
	}
	return nil
}

var File_protos_task_task_proto protoreflect.FileDescriptor

var file_protos_task_task_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x61, 0x67, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xdf,
	0x01, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x3c, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x32, 0x4e, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x76, 0x63, 0x12, 0x43, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a,
	0x6f, 0x6b, 0x65, 0x72, 0x36, 0x36, 0x36, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_task_task_proto_rawDescOnce sync.Once
	file_protos_task_task_proto_rawDescData = file_protos_task_task_proto_rawDesc
)

func file_protos_task_task_proto_rawDescGZIP() []byte {
	file_protos_task_task_proto_rawDescOnce.Do(func() {
		file_protos_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_task_task_proto_rawDescData)
	})
	return file_protos_task_task_proto_rawDescData
}

var file_protos_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_task_task_proto_goTypes = []interface{}{
	(*CreateTaskRequest)(nil),       // 0: demo_task.CreateTaskRequest
	(*TaskResponse)(nil),            // 1: demo_task.TaskResponse
	(*project.ProjectResponse)(nil), // 2: demo_project.ProjectResponse
	(*project.TagResponse)(nil),     // 3: demo_project.TagResponse
	(*user.UserResponse)(nil),       // 4: demo_user.UserResponse
}
var file_protos_task_task_proto_depIdxs = []int32{
	2, // 0: demo_task.TaskResponse.project:type_name -> demo_project.ProjectResponse
	3, // 1: demo_task.TaskResponse.tag:type_name -> demo_project.TagResponse
	4, // 2: demo_task.TaskResponse.assigned_user:type_name -> demo_user.UserResponse
	0, // 3: demo_task.TaskSvc.createTask:input_type -> demo_task.CreateTaskRequest
	1, // 4: demo_task.TaskSvc.createTask:output_type -> demo_task.TaskResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_task_task_proto_init() }
func file_protos_task_task_proto_init() {
	if File_protos_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_task_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_protos_task_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
			RawDescriptor: file_protos_task_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_task_task_proto_goTypes,
		DependencyIndexes: file_protos_task_task_proto_depIdxs,
		MessageInfos:      file_protos_task_task_proto_msgTypes,
	}.Build()
	File_protos_task_task_proto = out.File
	file_protos_task_task_proto_rawDesc = nil
	file_protos_task_task_proto_goTypes = nil
	file_protos_task_task_proto_depIdxs = nil
}
