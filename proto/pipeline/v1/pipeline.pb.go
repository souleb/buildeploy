// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/pipeline/v1/pipeline.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreatePipelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePipelineResponse) Reset() {
	*x = CreatePipelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePipelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePipelineResponse) ProtoMessage() {}

func (x *CreatePipelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePipelineResponse.ProtoReflect.Descriptor instead.
func (*CreatePipelineResponse) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePipelineResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreatePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Pipeline `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreatePipelineRequest) Reset() {
	*x = CreatePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePipelineRequest) ProtoMessage() {}

func (x *CreatePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePipelineRequest.ProtoReflect.Descriptor instead.
func (*CreatePipelineRequest) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePipelineRequest) GetItem() *Pipeline {
	if x != nil {
		return x.Item
	}
	return nil
}

type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Workflows []*Workflow `protobuf:"bytes,3,rep,name=workflows,proto3" json:"workflows,omitempty"`
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *Pipeline) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pipeline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pipeline) GetWorkflows() []*Workflow {
	if x != nil {
		return x.Workflows
	}
	return nil
}

type Workflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Jobs []*Job `protobuf:"bytes,3,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{3}
}

func (x *Workflow) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Workflow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workflow) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Runner   *Job_Runner `protobuf:"bytes,3,opt,name=runner,proto3" json:"runner,omitempty"`
	Steps    *Steps      `protobuf:"bytes,4,opt,name=steps,proto3" json:"steps,omitempty"`
	Env      string      `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
	Branches string      `protobuf:"bytes,6,opt,name=branches,proto3" json:"branches,omitempty"`
	Needs    []string    `protobuf:"bytes,7,rep,name=needs,proto3" json:"needs,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{4}
}

func (x *Job) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetRunner() *Job_Runner {
	if x != nil {
		return x.Runner
	}
	return nil
}

func (x *Job) GetSteps() *Steps {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Job) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Job) GetBranches() string {
	if x != nil {
		return x.Branches
	}
	return ""
}

func (x *Job) GetNeeds() []string {
	if x != nil {
		return x.Needs
	}
	return nil
}

type Docker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Tags  string `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Docker) Reset() {
	*x = Docker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Docker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Docker) ProtoMessage() {}

func (x *Docker) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Docker.ProtoReflect.Descriptor instead.
func (*Docker) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{5}
}

func (x *Docker) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Docker) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os     string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Cpus   string `protobuf:"bytes,2,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Memory string `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{6}
}

func (x *Machine) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *Machine) GetCpus() string {
	if x != nil {
		return x.Cpus
	}
	return ""
}

func (x *Machine) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

type Steps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command []string `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
}

func (x *Steps) Reset() {
	*x = Steps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Steps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Steps) ProtoMessage() {}

func (x *Steps) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Steps.ProtoReflect.Descriptor instead.
func (*Steps) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{7}
}

func (x *Steps) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

type Job_Runner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Job_Runner_Docker
	//	*Job_Runner_Machine
	Type isJob_Runner_Type `protobuf_oneof:"type"`
}

func (x *Job_Runner) Reset() {
	*x = Job_Runner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job_Runner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job_Runner) ProtoMessage() {}

func (x *Job_Runner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_v1_pipeline_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job_Runner.ProtoReflect.Descriptor instead.
func (*Job_Runner) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{4, 0}
}

func (m *Job_Runner) GetType() isJob_Runner_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Job_Runner) GetDocker() *Docker {
	if x, ok := x.GetType().(*Job_Runner_Docker); ok {
		return x.Docker
	}
	return nil
}

func (x *Job_Runner) GetMachine() *Machine {
	if x, ok := x.GetType().(*Job_Runner_Machine); ok {
		return x.Machine
	}
	return nil
}

type isJob_Runner_Type interface {
	isJob_Runner_Type()
}

type Job_Runner_Docker struct {
	Docker *Docker `protobuf:"bytes,1,opt,name=docker,proto3,oneof"`
}

type Job_Runner_Machine struct {
	Machine *Machine `protobuf:"bytes,2,opt,name=machine,proto3,oneof"`
}

func (*Job_Runner_Docker) isJob_Runner_Type() {}

func (*Job_Runner_Machine) isJob_Runner_Type() {}

var File_proto_pipeline_v1_pipeline_proto protoreflect.FileDescriptor

var file_proto_pipeline_v1_pipeline_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x63, 0x0a, 0x08, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x22, 0x54, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62,
	0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xbb, 0x02, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x72, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x65, 0x70, 0x73, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x65, 0x65, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x65, 0x64,
	0x73, 0x1a, 0x71, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x32, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x70, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x70, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22,
	0x21, 0x0a, 0x05, 0x53, 0x74, 0x65, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x32, 0x88, 0x01, 0x0a, 0x0f, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x6c,
	0x65, 0x62, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pipeline_v1_pipeline_proto_rawDescOnce sync.Once
	file_proto_pipeline_v1_pipeline_proto_rawDescData = file_proto_pipeline_v1_pipeline_proto_rawDesc
)

func file_proto_pipeline_v1_pipeline_proto_rawDescGZIP() []byte {
	file_proto_pipeline_v1_pipeline_proto_rawDescOnce.Do(func() {
		file_proto_pipeline_v1_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pipeline_v1_pipeline_proto_rawDescData)
	})
	return file_proto_pipeline_v1_pipeline_proto_rawDescData
}

var file_proto_pipeline_v1_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_pipeline_v1_pipeline_proto_goTypes = []interface{}{
	(*CreatePipelineResponse)(nil), // 0: pipeline.v1.CreatePipelineResponse
	(*CreatePipelineRequest)(nil),  // 1: pipeline.v1.CreatePipelineRequest
	(*Pipeline)(nil),               // 2: pipeline.v1.Pipeline
	(*Workflow)(nil),               // 3: pipeline.v1.Workflow
	(*Job)(nil),                    // 4: pipeline.v1.Job
	(*Docker)(nil),                 // 5: pipeline.v1.Docker
	(*Machine)(nil),                // 6: pipeline.v1.Machine
	(*Steps)(nil),                  // 7: pipeline.v1.Steps
	(*Job_Runner)(nil),             // 8: pipeline.v1.Job.Runner
}
var file_proto_pipeline_v1_pipeline_proto_depIdxs = []int32{
	2, // 0: pipeline.v1.CreatePipelineRequest.item:type_name -> pipeline.v1.Pipeline
	3, // 1: pipeline.v1.Pipeline.workflows:type_name -> pipeline.v1.Workflow
	4, // 2: pipeline.v1.Workflow.jobs:type_name -> pipeline.v1.Job
	8, // 3: pipeline.v1.Job.runner:type_name -> pipeline.v1.Job.Runner
	7, // 4: pipeline.v1.Job.steps:type_name -> pipeline.v1.Steps
	5, // 5: pipeline.v1.Job.Runner.docker:type_name -> pipeline.v1.Docker
	6, // 6: pipeline.v1.Job.Runner.machine:type_name -> pipeline.v1.Machine
	1, // 7: pipeline.v1.PipelineService.CreatePipeline:input_type -> pipeline.v1.CreatePipelineRequest
	0, // 8: pipeline.v1.PipelineService.CreatePipeline:output_type -> pipeline.v1.CreatePipelineResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_pipeline_v1_pipeline_proto_init() }
func file_proto_pipeline_v1_pipeline_proto_init() {
	if File_proto_pipeline_v1_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pipeline_v1_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePipelineResponse); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePipelineRequest); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workflow); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Docker); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Steps); i {
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
		file_proto_pipeline_v1_pipeline_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job_Runner); i {
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
	file_proto_pipeline_v1_pipeline_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Job_Runner_Docker)(nil),
		(*Job_Runner_Machine)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_pipeline_v1_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pipeline_v1_pipeline_proto_goTypes,
		DependencyIndexes: file_proto_pipeline_v1_pipeline_proto_depIdxs,
		MessageInfos:      file_proto_pipeline_v1_pipeline_proto_msgTypes,
	}.Build()
	File_proto_pipeline_v1_pipeline_proto = out.File
	file_proto_pipeline_v1_pipeline_proto_rawDesc = nil
	file_proto_pipeline_v1_pipeline_proto_goTypes = nil
	file_proto_pipeline_v1_pipeline_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PipelineServiceClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error)
}

type pipelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineServiceClient(cc grpc.ClientConnInterface) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error) {
	out := new(CreatePipelineResponse)
	err := c.cc.Invoke(ctx, "/pipeline.v1.PipelineService/CreatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineServiceServer is the server API for PipelineService service.
type PipelineServiceServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error)
}

// UnimplementedPipelineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPipelineServiceServer struct {
}

func (*UnimplementedPipelineServiceServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}

func RegisterPipelineServiceServer(s *grpc.Server, srv PipelineServiceServer) {
	s.RegisterService(&_PipelineService_serviceDesc, srv)
}

func _PipelineService_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipeline.v1.PipelineService/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PipelineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pipeline.v1.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineService_CreatePipeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pipeline/v1/pipeline.proto",
}
