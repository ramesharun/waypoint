// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package hashicorp_devflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Supported component types, the values here MUST match the enum values
// in the Go sdk/component package exactly. A test in internal/server
// validates this.
type Component_Type int32

const (
	Component_UNKNOWN  Component_Type = 0
	Component_BUILDER  Component_Type = 1
	Component_REGISTRY Component_Type = 2
)

var Component_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "BUILDER",
	2: "REGISTRY",
}
var Component_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"BUILDER":  1,
	"REGISTRY": 2,
}

func (x Component_Type) String() string {
	return proto.EnumName(Component_Type_name, int32(x))
}
func (Component_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{0, 0}
}

type Status_State int32

const (
	Status_UNKNOWN Status_State = 0
	Status_RUNNING Status_State = 1
	Status_SUCCESS Status_State = 2
	Status_ERROR   Status_State = 3
)

var Status_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "RUNNING",
	2: "SUCCESS",
	3: "ERROR",
}
var Status_State_value = map[string]int32{
	"UNKNOWN": 0,
	"RUNNING": 1,
	"SUCCESS": 2,
	"ERROR":   3,
}

func (x Status_State) String() string {
	return proto.EnumName(Status_State_name, int32(x))
}
func (Status_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{1, 0}
}

// Component represents metadata about a component. A component is the
// generic name for a builder, registry, platform, etc.
type Component struct {
	// type of the component
	Type Component_Type `protobuf:"varint,1,opt,name=type,proto3,enum=hashicorp.devflow.Component_Type" json:"type,omitempty"`
	// name of the component
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Component) Reset()         { *m = Component{} }
func (m *Component) String() string { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()    {}
func (*Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{0}
}
func (m *Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Component.Unmarshal(m, b)
}
func (m *Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Component.Marshal(b, m, deterministic)
}
func (dst *Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Component.Merge(dst, src)
}
func (m *Component) XXX_Size() int {
	return xxx_messageInfo_Component.Size(m)
}
func (m *Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Component proto.InternalMessageInfo

func (m *Component) GetType() Component_Type {
	if m != nil {
		return m.Type
	}
	return Component_UNKNOWN
}

func (m *Component) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Status represents the status of an async operation.
type Status struct {
	// state is the state of this operation.
	State Status_State `protobuf:"varint,1,opt,name=state,proto3,enum=hashicorp.devflow.Status_State" json:"state,omitempty"`
	// details may be non-empty to provide human-friendly information
	// about the current status. This may change between status updates
	// for the same state to provide updated details about the state.
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	// error is set if the state == ERROR with the error that occurred.
	Error *status.Status `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// start_time is the time the operation was started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// complete_time is the time the operation completed (success or fail).
	CompleteTime         *timestamp.Timestamp `protobuf:"bytes,5,opt,name=complete_time,json=completeTime,proto3" json:"complete_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{1}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetState() Status_State {
	if m != nil {
		return m.State
	}
	return Status_UNKNOWN
}

func (m *Status) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Status) GetError() *status.Status {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Status) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Status) GetCompleteTime() *timestamp.Timestamp {
	if m != nil {
		return m.CompleteTime
	}
	return nil
}

type CreateBuildRequest struct {
	// component type being used for the build
	Component            *Component `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateBuildRequest) Reset()         { *m = CreateBuildRequest{} }
func (m *CreateBuildRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBuildRequest) ProtoMessage()    {}
func (*CreateBuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{2}
}
func (m *CreateBuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBuildRequest.Unmarshal(m, b)
}
func (m *CreateBuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBuildRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBuildRequest.Merge(dst, src)
}
func (m *CreateBuildRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBuildRequest.Size(m)
}
func (m *CreateBuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBuildRequest proto.InternalMessageInfo

func (m *CreateBuildRequest) GetComponent() *Component {
	if m != nil {
		return m.Component
	}
	return nil
}

type CreateBuildResponse struct {
	// id is the unique ID for this build.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBuildResponse) Reset()         { *m = CreateBuildResponse{} }
func (m *CreateBuildResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBuildResponse) ProtoMessage()    {}
func (*CreateBuildResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{3}
}
func (m *CreateBuildResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBuildResponse.Unmarshal(m, b)
}
func (m *CreateBuildResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBuildResponse.Marshal(b, m, deterministic)
}
func (dst *CreateBuildResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBuildResponse.Merge(dst, src)
}
func (m *CreateBuildResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBuildResponse.Size(m)
}
func (m *CreateBuildResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBuildResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBuildResponse proto.InternalMessageInfo

func (m *CreateBuildResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CompleteBuildRequest struct {
	// id is the unique ID of the build to complete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*CompleteBuildRequest_Artifact
	//	*CompleteBuildRequest_Error
	Result               isCompleteBuildRequest_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CompleteBuildRequest) Reset()         { *m = CompleteBuildRequest{} }
func (m *CompleteBuildRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteBuildRequest) ProtoMessage()    {}
func (*CompleteBuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{4}
}
func (m *CompleteBuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteBuildRequest.Unmarshal(m, b)
}
func (m *CompleteBuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteBuildRequest.Marshal(b, m, deterministic)
}
func (dst *CompleteBuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteBuildRequest.Merge(dst, src)
}
func (m *CompleteBuildRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteBuildRequest.Size(m)
}
func (m *CompleteBuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteBuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteBuildRequest proto.InternalMessageInfo

func (m *CompleteBuildRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isCompleteBuildRequest_Result interface {
	isCompleteBuildRequest_Result()
}

type CompleteBuildRequest_Artifact struct {
	Artifact *Artifact `protobuf:"bytes,2,opt,name=artifact,proto3,oneof"`
}

type CompleteBuildRequest_Error struct {
	Error *status.Status `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*CompleteBuildRequest_Artifact) isCompleteBuildRequest_Result() {}

func (*CompleteBuildRequest_Error) isCompleteBuildRequest_Result() {}

func (m *CompleteBuildRequest) GetResult() isCompleteBuildRequest_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CompleteBuildRequest) GetArtifact() *Artifact {
	if x, ok := m.GetResult().(*CompleteBuildRequest_Artifact); ok {
		return x.Artifact
	}
	return nil
}

func (m *CompleteBuildRequest) GetError() *status.Status {
	if x, ok := m.GetResult().(*CompleteBuildRequest_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CompleteBuildRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CompleteBuildRequest_OneofMarshaler, _CompleteBuildRequest_OneofUnmarshaler, _CompleteBuildRequest_OneofSizer, []interface{}{
		(*CompleteBuildRequest_Artifact)(nil),
		(*CompleteBuildRequest_Error)(nil),
	}
}

func _CompleteBuildRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CompleteBuildRequest)
	// result
	switch x := m.Result.(type) {
	case *CompleteBuildRequest_Artifact:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Artifact); err != nil {
			return err
		}
	case *CompleteBuildRequest_Error:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CompleteBuildRequest.Result has unexpected type %T", x)
	}
	return nil
}

func _CompleteBuildRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CompleteBuildRequest)
	switch tag {
	case 2: // result.artifact
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Artifact)
		err := b.DecodeMessage(msg)
		m.Result = &CompleteBuildRequest_Artifact{msg}
		return true, err
	case 3: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(status.Status)
		err := b.DecodeMessage(msg)
		m.Result = &CompleteBuildRequest_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CompleteBuildRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CompleteBuildRequest)
	// result
	switch x := m.Result.(type) {
	case *CompleteBuildRequest_Artifact:
		s := proto.Size(x.Artifact)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CompleteBuildRequest_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ListBuildsResponse struct {
	// builds is the list of builds.
	Builds               []*Build `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBuildsResponse) Reset()         { *m = ListBuildsResponse{} }
func (m *ListBuildsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBuildsResponse) ProtoMessage()    {}
func (*ListBuildsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{5}
}
func (m *ListBuildsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBuildsResponse.Unmarshal(m, b)
}
func (m *ListBuildsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBuildsResponse.Marshal(b, m, deterministic)
}
func (dst *ListBuildsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBuildsResponse.Merge(dst, src)
}
func (m *ListBuildsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBuildsResponse.Size(m)
}
func (m *ListBuildsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBuildsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBuildsResponse proto.InternalMessageInfo

func (m *ListBuildsResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

// Build represents a process of creating an artifact that can be in any state,
// such as complete. A successful complete build produces an artifact.
type Build struct {
	// id is the unique ID of the build
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// status of the build
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// component is the component that was used for this build
	Component *Component `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	// artifact is the result of the build if the status.state == SUCCESS
	Artifact             *Artifact `protobuf:"bytes,4,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{6}
}
func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (dst *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(dst, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Build) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Build) GetComponent() *Component {
	if m != nil {
		return m.Component
	}
	return nil
}

func (m *Build) GetArtifact() *Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

// Artifact is the result of a build or registry. This is the metadata only.
// The binary contents of an artifact are expected to be stored in a registry.
type Artifact struct {
	// artifact is the full artifact encoded directly from the component plugin.
	// The receiving end must have access to the component proto files to
	// know how to decode this.
	Artifact             *any.Any `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{7}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (dst *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(dst, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetArtifact() *any.Any {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type UpsertPushedArtifactRequest struct {
	// artifact to upsert. If the id in the artifact is empty, then this
	// will be an insert. Otherwise, this will be an update and if the ID
	// isn't found, it will be an error.
	Artifact             *PushedArtifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpsertPushedArtifactRequest) Reset()         { *m = UpsertPushedArtifactRequest{} }
func (m *UpsertPushedArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertPushedArtifactRequest) ProtoMessage()    {}
func (*UpsertPushedArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{8}
}
func (m *UpsertPushedArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertPushedArtifactRequest.Unmarshal(m, b)
}
func (m *UpsertPushedArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertPushedArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *UpsertPushedArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertPushedArtifactRequest.Merge(dst, src)
}
func (m *UpsertPushedArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertPushedArtifactRequest.Size(m)
}
func (m *UpsertPushedArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertPushedArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertPushedArtifactRequest proto.InternalMessageInfo

func (m *UpsertPushedArtifactRequest) GetArtifact() *PushedArtifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type UpsertPushedArtifactResponse struct {
	// resulting push object, you should replace this with what was sent
	// since the update operation may touch up the input data (i.e. update
	// timestamps)
	Artifact             *PushedArtifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpsertPushedArtifactResponse) Reset()         { *m = UpsertPushedArtifactResponse{} }
func (m *UpsertPushedArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*UpsertPushedArtifactResponse) ProtoMessage()    {}
func (*UpsertPushedArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{9}
}
func (m *UpsertPushedArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertPushedArtifactResponse.Unmarshal(m, b)
}
func (m *UpsertPushedArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertPushedArtifactResponse.Marshal(b, m, deterministic)
}
func (dst *UpsertPushedArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertPushedArtifactResponse.Merge(dst, src)
}
func (m *UpsertPushedArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_UpsertPushedArtifactResponse.Size(m)
}
func (m *UpsertPushedArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertPushedArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertPushedArtifactResponse proto.InternalMessageInfo

func (m *UpsertPushedArtifactResponse) GetArtifact() *PushedArtifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type PushedArtifact struct {
	// id is a unique ID for this push
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// status of the push operation
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// component that pushed this artifact
	Component *Component `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	// artifact is the artifact that was a result from the push.
	Artifact             *Artifact `protobuf:"bytes,4,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PushedArtifact) Reset()         { *m = PushedArtifact{} }
func (m *PushedArtifact) String() string { return proto.CompactTextString(m) }
func (*PushedArtifact) ProtoMessage()    {}
func (*PushedArtifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_6854e6c6e112be6a, []int{10}
}
func (m *PushedArtifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushedArtifact.Unmarshal(m, b)
}
func (m *PushedArtifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushedArtifact.Marshal(b, m, deterministic)
}
func (dst *PushedArtifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushedArtifact.Merge(dst, src)
}
func (m *PushedArtifact) XXX_Size() int {
	return xxx_messageInfo_PushedArtifact.Size(m)
}
func (m *PushedArtifact) XXX_DiscardUnknown() {
	xxx_messageInfo_PushedArtifact.DiscardUnknown(m)
}

var xxx_messageInfo_PushedArtifact proto.InternalMessageInfo

func (m *PushedArtifact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PushedArtifact) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PushedArtifact) GetComponent() *Component {
	if m != nil {
		return m.Component
	}
	return nil
}

func (m *PushedArtifact) GetArtifact() *Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

func init() {
	proto.RegisterType((*Component)(nil), "hashicorp.devflow.Component")
	proto.RegisterType((*Status)(nil), "hashicorp.devflow.Status")
	proto.RegisterType((*CreateBuildRequest)(nil), "hashicorp.devflow.CreateBuildRequest")
	proto.RegisterType((*CreateBuildResponse)(nil), "hashicorp.devflow.CreateBuildResponse")
	proto.RegisterType((*CompleteBuildRequest)(nil), "hashicorp.devflow.CompleteBuildRequest")
	proto.RegisterType((*ListBuildsResponse)(nil), "hashicorp.devflow.ListBuildsResponse")
	proto.RegisterType((*Build)(nil), "hashicorp.devflow.Build")
	proto.RegisterType((*Artifact)(nil), "hashicorp.devflow.Artifact")
	proto.RegisterType((*UpsertPushedArtifactRequest)(nil), "hashicorp.devflow.UpsertPushedArtifactRequest")
	proto.RegisterType((*UpsertPushedArtifactResponse)(nil), "hashicorp.devflow.UpsertPushedArtifactResponse")
	proto.RegisterType((*PushedArtifact)(nil), "hashicorp.devflow.PushedArtifact")
	proto.RegisterEnum("hashicorp.devflow.Component_Type", Component_Type_name, Component_Type_value)
	proto.RegisterEnum("hashicorp.devflow.Status_State", Status_State_name, Status_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DevflowClient is the client API for Devflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevflowClient interface {
	// ListBuilds returns the builds.
	ListBuilds(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListBuildsResponse, error)
	// CreateBuild creates the metadata associated with an artifact build.
	CreateBuild(ctx context.Context, in *CreateBuildRequest, opts ...grpc.CallOption) (*CreateBuildResponse, error)
	// CompleteBuild is called when a build is completed (for a local run).
	// This updates the final status of the build, including information about
	// the generated artifact.
	CompleteBuild(ctx context.Context, in *CompleteBuildRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// UpsertPushedArtifact updates or inserts a pushed artifact. This is
	// useful for local operations to work on a pushed artifact.
	UpsertPushedArtifact(ctx context.Context, in *UpsertPushedArtifactRequest, opts ...grpc.CallOption) (*UpsertPushedArtifactResponse, error)
}

type devflowClient struct {
	cc *grpc.ClientConn
}

func NewDevflowClient(cc *grpc.ClientConn) DevflowClient {
	return &devflowClient{cc}
}

func (c *devflowClient) ListBuilds(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListBuildsResponse, error) {
	out := new(ListBuildsResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.devflow.Devflow/ListBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devflowClient) CreateBuild(ctx context.Context, in *CreateBuildRequest, opts ...grpc.CallOption) (*CreateBuildResponse, error) {
	out := new(CreateBuildResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.devflow.Devflow/CreateBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devflowClient) CompleteBuild(ctx context.Context, in *CompleteBuildRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/hashicorp.devflow.Devflow/CompleteBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devflowClient) UpsertPushedArtifact(ctx context.Context, in *UpsertPushedArtifactRequest, opts ...grpc.CallOption) (*UpsertPushedArtifactResponse, error) {
	out := new(UpsertPushedArtifactResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.devflow.Devflow/UpsertPushedArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevflowServer is the server API for Devflow service.
type DevflowServer interface {
	// ListBuilds returns the builds.
	ListBuilds(context.Context, *empty.Empty) (*ListBuildsResponse, error)
	// CreateBuild creates the metadata associated with an artifact build.
	CreateBuild(context.Context, *CreateBuildRequest) (*CreateBuildResponse, error)
	// CompleteBuild is called when a build is completed (for a local run).
	// This updates the final status of the build, including information about
	// the generated artifact.
	CompleteBuild(context.Context, *CompleteBuildRequest) (*empty.Empty, error)
	// UpsertPushedArtifact updates or inserts a pushed artifact. This is
	// useful for local operations to work on a pushed artifact.
	UpsertPushedArtifact(context.Context, *UpsertPushedArtifactRequest) (*UpsertPushedArtifactResponse, error)
}

func RegisterDevflowServer(s *grpc.Server, srv DevflowServer) {
	s.RegisterService(&_Devflow_serviceDesc, srv)
}

func _Devflow_ListBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevflowServer).ListBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.devflow.Devflow/ListBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevflowServer).ListBuilds(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devflow_CreateBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevflowServer).CreateBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.devflow.Devflow/CreateBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevflowServer).CreateBuild(ctx, req.(*CreateBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devflow_CompleteBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevflowServer).CompleteBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.devflow.Devflow/CompleteBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevflowServer).CompleteBuild(ctx, req.(*CompleteBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devflow_UpsertPushedArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPushedArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevflowServer).UpsertPushedArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.devflow.Devflow/UpsertPushedArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevflowServer).UpsertPushedArtifact(ctx, req.(*UpsertPushedArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Devflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.devflow.Devflow",
	HandlerType: (*DevflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBuilds",
			Handler:    _Devflow_ListBuilds_Handler,
		},
		{
			MethodName: "CreateBuild",
			Handler:    _Devflow_CreateBuild_Handler,
		},
		{
			MethodName: "CompleteBuild",
			Handler:    _Devflow_CompleteBuild_Handler,
		},
		{
			MethodName: "UpsertPushedArtifact",
			Handler:    _Devflow_UpsertPushedArtifact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_server_6854e6c6e112be6a) }

var fileDescriptor_server_6854e6c6e112be6a = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xc7, 0xd9, 0x7e, 0xd1, 0x9e, 0x02, 0xa9, 0x47, 0xa2, 0xcb, 0x42, 0x42, 0xdd, 0x04, 0x6d,
	0xbc, 0xd8, 0x62, 0x0d, 0x31, 0x18, 0x8d, 0x81, 0x52, 0x81, 0x40, 0x0a, 0x99, 0xd2, 0x18, 0x13,
	0x8c, 0x59, 0xda, 0x01, 0x36, 0x69, 0xbb, 0xeb, 0xcc, 0x14, 0xd2, 0x17, 0xf0, 0x15, 0x7c, 0x17,
	0x6f, 0xbc, 0xf5, 0xad, 0x34, 0x3b, 0x3b, 0xdb, 0xaf, 0x5d, 0xa8, 0xc6, 0x2b, 0xaf, 0xda, 0x99,
	0xf9, 0x9d, 0x7f, 0xcf, 0x39, 0xff, 0x33, 0x53, 0x58, 0xe0, 0x94, 0xdd, 0x50, 0x66, 0x79, 0xcc,
	0x15, 0x2e, 0x3e, 0xb8, 0xb6, 0xf9, 0xb5, 0xd3, 0x72, 0x99, 0x67, 0xb5, 0xe9, 0xcd, 0x65, 0xc7,
	0xbd, 0x35, 0x56, 0xae, 0x5c, 0xf7, 0xaa, 0x43, 0xcb, 0x12, 0xb8, 0xe8, 0x5f, 0x96, 0xed, 0xde,
	0x20, 0xa0, 0x8d, 0xd5, 0xe9, 0x23, 0xda, 0xf5, 0x44, 0x78, 0xb8, 0x3e, 0x7d, 0x28, 0x9c, 0x2e,
	0xe5, 0xc2, 0xee, 0x7a, 0x0a, 0x78, 0xac, 0x00, 0xe6, 0xb5, 0xca, 0x5c, 0xd8, 0xa2, 0xcf, 0x83,
	0x03, 0xf3, 0xab, 0x06, 0xb9, 0xaa, 0xdb, 0xf5, 0xdc, 0x1e, 0xed, 0x09, 0xdc, 0x82, 0x94, 0x18,
	0x78, 0x54, 0xd7, 0x8a, 0x5a, 0x69, 0xa9, 0xf2, 0xc4, 0x8a, 0x64, 0x68, 0x0d, 0x59, 0xeb, 0x6c,
	0xe0, 0x51, 0x22, 0x71, 0x44, 0x48, 0xf5, 0xec, 0x2e, 0xd5, 0x13, 0x45, 0xad, 0x94, 0x23, 0xf2,
	0xbb, 0x69, 0x41, 0xca, 0x27, 0x30, 0x0f, 0xf3, 0xcd, 0xfa, 0x51, 0xfd, 0xe4, 0x43, 0xbd, 0x30,
	0xe7, 0x2f, 0x76, 0x9b, 0x87, 0xc7, 0x7b, 0x35, 0x52, 0xd0, 0x70, 0x01, 0xb2, 0xa4, 0xb6, 0x7f,
	0xd8, 0x38, 0x23, 0x1f, 0x0b, 0x09, 0xf3, 0x7b, 0x02, 0x32, 0x0d, 0x99, 0x19, 0x6e, 0x41, 0xda,
	0xcf, 0x31, 0x4c, 0x63, 0x3d, 0x26, 0x8d, 0x80, 0x94, 0x1f, 0x94, 0x04, 0x34, 0xea, 0x30, 0xdf,
	0xa6, 0xc2, 0x76, 0x3a, 0x5c, 0x25, 0x12, 0x2e, 0xb1, 0x04, 0x69, 0xca, 0x98, 0xcb, 0xf4, 0x64,
	0x51, 0x2b, 0xe5, 0x2b, 0x68, 0x05, 0xdd, 0xb0, 0x98, 0xd7, 0x52, 0x4a, 0x24, 0x00, 0x70, 0x1b,
	0x80, 0x0b, 0x9b, 0x89, 0xcf, 0x7e, 0x03, 0xf5, 0x94, 0xc4, 0x8d, 0x10, 0x0f, 0xbb, 0x6b, 0x9d,
	0x85, 0xdd, 0x25, 0x39, 0x49, 0xfb, 0x6b, 0x7c, 0x07, 0x8b, 0x2d, 0xb7, 0xeb, 0x75, 0xa8, 0xa0,
	0x41, 0x74, 0x7a, 0x66, 0xf4, 0x42, 0x18, 0xe0, 0x6f, 0x99, 0xdb, 0x90, 0x96, 0xf5, 0x44, 0x5a,
	0x46, 0x9a, 0xf5, 0xfa, 0x61, 0x7d, 0xbf, 0xa0, 0xf9, 0x8b, 0x46, 0xb3, 0x5a, 0xad, 0x35, 0x1a,
	0x85, 0x04, 0xe6, 0x20, 0x5d, 0x23, 0xe4, 0x84, 0x14, 0x92, 0xe6, 0x29, 0x60, 0x95, 0x51, 0x5b,
	0xd0, 0xdd, 0xbe, 0xd3, 0x69, 0x13, 0xfa, 0xa5, 0x4f, 0xb9, 0xc0, 0xd7, 0x90, 0x6b, 0x85, 0x76,
	0xc9, 0x5e, 0xe6, 0x2b, 0x6b, 0xf7, 0x59, 0x4a, 0x46, 0xb8, 0xb9, 0x01, 0x0f, 0x27, 0x14, 0xb9,
	0xe7, 0xf6, 0x38, 0xc5, 0x25, 0x48, 0x38, 0x6d, 0xa9, 0x95, 0x23, 0x09, 0xa7, 0x6d, 0x7e, 0xd3,
	0x60, 0xb9, 0xaa, 0x8a, 0x98, 0xf8, 0xed, 0x29, 0x10, 0xb7, 0x21, 0x6b, 0x33, 0xe1, 0x5c, 0xda,
	0x2d, 0x21, 0xdd, 0xc9, 0x57, 0x56, 0x63, 0x52, 0xd9, 0x51, 0xc8, 0xc1, 0x1c, 0x19, 0xe2, 0xf8,
	0x7c, 0xa6, 0x7b, 0x07, 0x73, 0xca, 0xbf, 0xdd, 0x2c, 0x64, 0x18, 0xe5, 0xfd, 0x8e, 0x30, 0xdf,
	0x03, 0x1e, 0x3b, 0x5c, 0xc8, 0xa4, 0xf8, 0x30, 0xff, 0x4d, 0xc8, 0x5c, 0xc8, 0x1d, 0x5d, 0x2b,
	0x26, 0x4b, 0xf9, 0x8a, 0x1e, 0x93, 0x44, 0x50, 0x87, 0xe2, 0xcc, 0x1f, 0x1a, 0xa4, 0xe5, 0x4e,
	0xa4, 0xa4, 0x17, 0x90, 0x09, 0xae, 0x92, 0x2a, 0x68, 0xe5, 0xce, 0x39, 0x25, 0x0a, 0x9c, 0x74,
	0x24, 0xf9, 0x57, 0x8e, 0xe0, 0xab, 0xb1, 0x0e, 0xa6, 0x66, 0x76, 0x70, 0xd4, 0x3f, 0xf3, 0x0d,
	0x64, 0xc3, 0x5d, 0xdc, 0x1c, 0x13, 0x09, 0x26, 0x62, 0x39, 0x32, 0x9f, 0x3b, 0xbd, 0xc1, 0x58,
	0xf4, 0x39, 0xac, 0x36, 0x3d, 0x4e, 0x99, 0x38, 0xed, 0xf3, 0x6b, 0xda, 0x1e, 0xea, 0x2b, 0x9f,
	0xdf, 0x46, 0x04, 0xe3, 0x5e, 0x8d, 0xa9, 0xd8, 0x91, 0xfa, 0x27, 0x58, 0x8b, 0x57, 0x57, 0x7e,
	0xfd, 0xa3, 0xfc, 0x4f, 0x0d, 0x96, 0x26, 0x0f, 0xff, 0x57, 0x17, 0x2b, 0xbf, 0x12, 0x30, 0xbf,
	0x17, 0x1c, 0xe3, 0x11, 0xc0, 0x68, 0xb6, 0xf1, 0x51, 0xc4, 0xc1, 0x9a, 0xff, 0xd7, 0x60, 0x6c,
	0xc4, 0x08, 0xc7, 0x5c, 0x89, 0x73, 0xc8, 0x8f, 0xdd, 0x74, 0x8c, 0x8b, 0x8a, 0xbe, 0x2d, 0xc6,
	0xd3, 0x59, 0x98, 0x52, 0x3f, 0x85, 0xc5, 0x89, 0xf7, 0x01, 0x9f, 0xdd, 0xd1, 0xa9, 0xe9, 0x17,
	0xc4, 0xb8, 0xa3, 0x2c, 0xbc, 0x85, 0xe5, 0xb8, 0x91, 0x41, 0x2b, 0x46, 0xf8, 0x9e, 0xc9, 0x35,
	0xca, 0x7f, 0xcc, 0x07, 0xa5, 0x5c, 0x64, 0x64, 0x22, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x50, 0xff, 0x8f, 0x61, 0xc6, 0x07, 0x00, 0x00,
}