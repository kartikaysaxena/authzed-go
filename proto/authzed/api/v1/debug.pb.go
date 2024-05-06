// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: authzed/api/v1/debug.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckDebugTrace_PermissionType int32

const (
	CheckDebugTrace_PERMISSION_TYPE_UNSPECIFIED CheckDebugTrace_PermissionType = 0
	CheckDebugTrace_PERMISSION_TYPE_RELATION    CheckDebugTrace_PermissionType = 1
	CheckDebugTrace_PERMISSION_TYPE_PERMISSION  CheckDebugTrace_PermissionType = 2
)

// Enum value maps for CheckDebugTrace_PermissionType.
var (
	CheckDebugTrace_PermissionType_name = map[int32]string{
		0: "PERMISSION_TYPE_UNSPECIFIED",
		1: "PERMISSION_TYPE_RELATION",
		2: "PERMISSION_TYPE_PERMISSION",
	}
	CheckDebugTrace_PermissionType_value = map[string]int32{
		"PERMISSION_TYPE_UNSPECIFIED": 0,
		"PERMISSION_TYPE_RELATION":    1,
		"PERMISSION_TYPE_PERMISSION":  2,
	}
)

func (x CheckDebugTrace_PermissionType) Enum() *CheckDebugTrace_PermissionType {
	p := new(CheckDebugTrace_PermissionType)
	*p = x
	return p
}

func (x CheckDebugTrace_PermissionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckDebugTrace_PermissionType) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_debug_proto_enumTypes[0].Descriptor()
}

func (CheckDebugTrace_PermissionType) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_debug_proto_enumTypes[0]
}

func (x CheckDebugTrace_PermissionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckDebugTrace_PermissionType.Descriptor instead.
func (CheckDebugTrace_PermissionType) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1, 0}
}

type CheckDebugTrace_Permissionship int32

const (
	CheckDebugTrace_PERMISSIONSHIP_UNSPECIFIED            CheckDebugTrace_Permissionship = 0
	CheckDebugTrace_PERMISSIONSHIP_NO_PERMISSION          CheckDebugTrace_Permissionship = 1
	CheckDebugTrace_PERMISSIONSHIP_HAS_PERMISSION         CheckDebugTrace_Permissionship = 2
	CheckDebugTrace_PERMISSIONSHIP_CONDITIONAL_PERMISSION CheckDebugTrace_Permissionship = 3
)

// Enum value maps for CheckDebugTrace_Permissionship.
var (
	CheckDebugTrace_Permissionship_name = map[int32]string{
		0: "PERMISSIONSHIP_UNSPECIFIED",
		1: "PERMISSIONSHIP_NO_PERMISSION",
		2: "PERMISSIONSHIP_HAS_PERMISSION",
		3: "PERMISSIONSHIP_CONDITIONAL_PERMISSION",
	}
	CheckDebugTrace_Permissionship_value = map[string]int32{
		"PERMISSIONSHIP_UNSPECIFIED":            0,
		"PERMISSIONSHIP_NO_PERMISSION":          1,
		"PERMISSIONSHIP_HAS_PERMISSION":         2,
		"PERMISSIONSHIP_CONDITIONAL_PERMISSION": 3,
	}
)

func (x CheckDebugTrace_Permissionship) Enum() *CheckDebugTrace_Permissionship {
	p := new(CheckDebugTrace_Permissionship)
	*p = x
	return p
}

func (x CheckDebugTrace_Permissionship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckDebugTrace_Permissionship) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_debug_proto_enumTypes[1].Descriptor()
}

func (CheckDebugTrace_Permissionship) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_debug_proto_enumTypes[1]
}

func (x CheckDebugTrace_Permissionship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckDebugTrace_Permissionship.Descriptor instead.
func (CheckDebugTrace_Permissionship) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1, 1}
}

type CaveatEvalInfo_Result int32

const (
	CaveatEvalInfo_RESULT_UNSPECIFIED          CaveatEvalInfo_Result = 0
	CaveatEvalInfo_RESULT_UNEVALUATED          CaveatEvalInfo_Result = 1
	CaveatEvalInfo_RESULT_FALSE                CaveatEvalInfo_Result = 2
	CaveatEvalInfo_RESULT_TRUE                 CaveatEvalInfo_Result = 3
	CaveatEvalInfo_RESULT_MISSING_SOME_CONTEXT CaveatEvalInfo_Result = 4
)

// Enum value maps for CaveatEvalInfo_Result.
var (
	CaveatEvalInfo_Result_name = map[int32]string{
		0: "RESULT_UNSPECIFIED",
		1: "RESULT_UNEVALUATED",
		2: "RESULT_FALSE",
		3: "RESULT_TRUE",
		4: "RESULT_MISSING_SOME_CONTEXT",
	}
	CaveatEvalInfo_Result_value = map[string]int32{
		"RESULT_UNSPECIFIED":          0,
		"RESULT_UNEVALUATED":          1,
		"RESULT_FALSE":                2,
		"RESULT_TRUE":                 3,
		"RESULT_MISSING_SOME_CONTEXT": 4,
	}
)

func (x CaveatEvalInfo_Result) Enum() *CaveatEvalInfo_Result {
	p := new(CaveatEvalInfo_Result)
	*p = x
	return p
}

func (x CaveatEvalInfo_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaveatEvalInfo_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_authzed_api_v1_debug_proto_enumTypes[2].Descriptor()
}

func (CaveatEvalInfo_Result) Type() protoreflect.EnumType {
	return &file_authzed_api_v1_debug_proto_enumTypes[2]
}

func (x CaveatEvalInfo_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaveatEvalInfo_Result.Descriptor instead.
func (CaveatEvalInfo_Result) EnumDescriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{2, 0}
}

// DebugInformation defines debug information returned by an API call in a footer when
// requested with a specific debugging header.
//
// The specific debug information returned will depend on the type of the API call made.
//
// See the github.com/authzed/authzed-go project for the specific header and footer names.
type DebugInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check holds debug information about a check request.
	Check *CheckDebugTrace `protobuf:"bytes,1,opt,name=check,proto3" json:"check,omitempty"`
	// schema_used holds the schema used for the request.
	SchemaUsed string `protobuf:"bytes,2,opt,name=schema_used,json=schemaUsed,proto3" json:"schema_used,omitempty"`
}

func (x *DebugInformation) Reset() {
	*x = DebugInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInformation) ProtoMessage() {}

func (x *DebugInformation) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInformation.ProtoReflect.Descriptor instead.
func (*DebugInformation) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{0}
}

func (x *DebugInformation) GetCheck() *CheckDebugTrace {
	if x != nil {
		return x.Check
	}
	return nil
}

func (x *DebugInformation) GetSchemaUsed() string {
	if x != nil {
		return x.SchemaUsed
	}
	return ""
}

// CheckDebugTrace is a recursive trace of the requests made for resolving a CheckPermission
// API call.
type CheckDebugTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource holds the resource on which the Check was performed.
	Resource *ObjectReference `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// permission holds the name of the permission or relation on which the Check was performed.
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// permission_type holds information indicating whether it was a permission or relation.
	PermissionType CheckDebugTrace_PermissionType `protobuf:"varint,3,opt,name=permission_type,json=permissionType,proto3,enum=authzed.api.v1.CheckDebugTrace_PermissionType" json:"permission_type,omitempty"`
	// subject holds the subject on which the Check was performed. This will be static across all calls within
	// the same Check tree.
	Subject *SubjectReference `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	// result holds the result of the Check call.
	Result CheckDebugTrace_Permissionship `protobuf:"varint,5,opt,name=result,proto3,enum=authzed.api.v1.CheckDebugTrace_Permissionship" json:"result,omitempty"`
	// caveat_evaluation_info holds information about the caveat evaluated for this step of the trace.
	CaveatEvaluationInfo *CaveatEvalInfo `protobuf:"bytes,8,opt,name=caveat_evaluation_info,json=caveatEvaluationInfo,proto3" json:"caveat_evaluation_info,omitempty"`
	// duration holds the time spent executing this Check operation.
	Duration *durationpb.Duration `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// resolution holds information about how the problem was resolved.
	//
	// Types that are assignable to Resolution:
	//
	//	*CheckDebugTrace_WasCachedResult
	//	*CheckDebugTrace_SubProblems_
	Resolution isCheckDebugTrace_Resolution `protobuf_oneof:"resolution"`
}

func (x *CheckDebugTrace) Reset() {
	*x = CheckDebugTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckDebugTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDebugTrace) ProtoMessage() {}

func (x *CheckDebugTrace) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDebugTrace.ProtoReflect.Descriptor instead.
func (*CheckDebugTrace) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1}
}

func (x *CheckDebugTrace) GetResource() *ObjectReference {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *CheckDebugTrace) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *CheckDebugTrace) GetPermissionType() CheckDebugTrace_PermissionType {
	if x != nil {
		return x.PermissionType
	}
	return CheckDebugTrace_PERMISSION_TYPE_UNSPECIFIED
}

func (x *CheckDebugTrace) GetSubject() *SubjectReference {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *CheckDebugTrace) GetResult() CheckDebugTrace_Permissionship {
	if x != nil {
		return x.Result
	}
	return CheckDebugTrace_PERMISSIONSHIP_UNSPECIFIED
}

func (x *CheckDebugTrace) GetCaveatEvaluationInfo() *CaveatEvalInfo {
	if x != nil {
		return x.CaveatEvaluationInfo
	}
	return nil
}

func (x *CheckDebugTrace) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (m *CheckDebugTrace) GetResolution() isCheckDebugTrace_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (x *CheckDebugTrace) GetWasCachedResult() bool {
	if x, ok := x.GetResolution().(*CheckDebugTrace_WasCachedResult); ok {
		return x.WasCachedResult
	}
	return false
}

func (x *CheckDebugTrace) GetSubProblems() *CheckDebugTrace_SubProblems {
	if x, ok := x.GetResolution().(*CheckDebugTrace_SubProblems_); ok {
		return x.SubProblems
	}
	return nil
}

type isCheckDebugTrace_Resolution interface {
	isCheckDebugTrace_Resolution()
}

type CheckDebugTrace_WasCachedResult struct {
	// was_cached_result, if true, indicates that the result was found in the cache and returned directly.
	WasCachedResult bool `protobuf:"varint,6,opt,name=was_cached_result,json=wasCachedResult,proto3,oneof"`
}

type CheckDebugTrace_SubProblems_ struct {
	// sub_problems holds the sub problems that were executed to resolve the answer to this Check. An empty list
	// and a permissionship of PERMISSIONSHIP_HAS_PERMISSION indicates the subject was found within this relation.
	SubProblems *CheckDebugTrace_SubProblems `protobuf:"bytes,7,opt,name=sub_problems,json=subProblems,proto3,oneof"`
}

func (*CheckDebugTrace_WasCachedResult) isCheckDebugTrace_Resolution() {}

func (*CheckDebugTrace_SubProblems_) isCheckDebugTrace_Resolution() {}

// CaveatEvalInfo holds information about a caveat expression that was evaluated.
type CaveatEvalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// expression is the expression that was evaluated.
	Expression string `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	// result is the result of the evaluation.
	Result CaveatEvalInfo_Result `protobuf:"varint,2,opt,name=result,proto3,enum=authzed.api.v1.CaveatEvalInfo_Result" json:"result,omitempty"`
	// context consists of any named values that were used for evaluating the caveat expression.
	Context *structpb.Struct `protobuf:"bytes,3,opt,name=context,proto3" json:"context,omitempty"`
	// partial_caveat_info holds information of a partially-evaluated caveated response, if applicable.
	PartialCaveatInfo *PartialCaveatInfo `protobuf:"bytes,4,opt,name=partial_caveat_info,json=partialCaveatInfo,proto3" json:"partial_caveat_info,omitempty"`
	// caveat_name is the name of the caveat that was executed, if applicable.
	CaveatName string `protobuf:"bytes,5,opt,name=caveat_name,json=caveatName,proto3" json:"caveat_name,omitempty"`
}

func (x *CaveatEvalInfo) Reset() {
	*x = CaveatEvalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaveatEvalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaveatEvalInfo) ProtoMessage() {}

func (x *CaveatEvalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaveatEvalInfo.ProtoReflect.Descriptor instead.
func (*CaveatEvalInfo) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{2}
}

func (x *CaveatEvalInfo) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *CaveatEvalInfo) GetResult() CaveatEvalInfo_Result {
	if x != nil {
		return x.Result
	}
	return CaveatEvalInfo_RESULT_UNSPECIFIED
}

func (x *CaveatEvalInfo) GetContext() *structpb.Struct {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *CaveatEvalInfo) GetPartialCaveatInfo() *PartialCaveatInfo {
	if x != nil {
		return x.PartialCaveatInfo
	}
	return nil
}

func (x *CaveatEvalInfo) GetCaveatName() string {
	if x != nil {
		return x.CaveatName
	}
	return ""
}

type CheckDebugTrace_SubProblems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Traces []*CheckDebugTrace `protobuf:"bytes,1,rep,name=traces,proto3" json:"traces,omitempty"`
}

func (x *CheckDebugTrace_SubProblems) Reset() {
	*x = CheckDebugTrace_SubProblems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authzed_api_v1_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckDebugTrace_SubProblems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckDebugTrace_SubProblems) ProtoMessage() {}

func (x *CheckDebugTrace_SubProblems) ProtoReflect() protoreflect.Message {
	mi := &file_authzed_api_v1_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckDebugTrace_SubProblems.ProtoReflect.Descriptor instead.
func (*CheckDebugTrace_SubProblems) Descriptor() ([]byte, []int) {
	return file_authzed_api_v1_debug_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CheckDebugTrace_SubProblems) GetTraces() []*CheckDebugTrace {
	if x != nil {
		return x.Traces
	}
	return nil
}

var File_authzed_api_v1_debug_proto protoreflect.FileDescriptor

var file_authzed_api_v1_debug_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a,
	0x0a, 0x10, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x62, 0x75, 0x67, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x73, 0x65, 0x64, 0x22, 0xf3, 0x07, 0x0a, 0x0f, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x62, 0x75, 0x67, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x63, 0x0a, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x62, 0x75, 0x67, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x52, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x62, 0x75, 0x67, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x54, 0x0a, 0x16, 0x63, 0x61, 0x76, 0x65, 0x61, 0x74, 0x5f, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x76, 0x65, 0x61, 0x74, 0x45, 0x76, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x63, 0x61, 0x76, 0x65, 0x61, 0x74, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x77, 0x61, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f,
	0x77, 0x61, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x50, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x73, 0x1a, 0x46, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73,
	0x12, 0x37, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x62, 0x75, 0x67, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x22, 0x6f, 0x0a, 0x0e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x1e, 0x0a,
	0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a,
	0x1c, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x4e, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12,
	0x21, 0x0a, 0x1d, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49,
	0x50, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x53, 0x48, 0x49, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c,
	0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x11, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0x94, 0x03, 0x0a, 0x0e, 0x43, 0x61, 0x76, 0x65, 0x61, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x76, 0x65, 0x61, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x51, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x63, 0x61, 0x76, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x61, 0x76, 0x65, 0x61,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x61,
	0x76, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x76, 0x65,
	0x61, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x61, 0x76, 0x65, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41,
	0x4c, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x54, 0x52, 0x55, 0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4f, 0x4d, 0x45, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x58, 0x54, 0x10, 0x04, 0x42, 0x4a, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authzed_api_v1_debug_proto_rawDescOnce sync.Once
	file_authzed_api_v1_debug_proto_rawDescData = file_authzed_api_v1_debug_proto_rawDesc
)

func file_authzed_api_v1_debug_proto_rawDescGZIP() []byte {
	file_authzed_api_v1_debug_proto_rawDescOnce.Do(func() {
		file_authzed_api_v1_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_authzed_api_v1_debug_proto_rawDescData)
	})
	return file_authzed_api_v1_debug_proto_rawDescData
}

var file_authzed_api_v1_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_authzed_api_v1_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_authzed_api_v1_debug_proto_goTypes = []interface{}{
	(CheckDebugTrace_PermissionType)(0), // 0: authzed.api.v1.CheckDebugTrace.PermissionType
	(CheckDebugTrace_Permissionship)(0), // 1: authzed.api.v1.CheckDebugTrace.Permissionship
	(CaveatEvalInfo_Result)(0),          // 2: authzed.api.v1.CaveatEvalInfo.Result
	(*DebugInformation)(nil),            // 3: authzed.api.v1.DebugInformation
	(*CheckDebugTrace)(nil),             // 4: authzed.api.v1.CheckDebugTrace
	(*CaveatEvalInfo)(nil),              // 5: authzed.api.v1.CaveatEvalInfo
	(*CheckDebugTrace_SubProblems)(nil), // 6: authzed.api.v1.CheckDebugTrace.SubProblems
	(*ObjectReference)(nil),             // 7: authzed.api.v1.ObjectReference
	(*SubjectReference)(nil),            // 8: authzed.api.v1.SubjectReference
	(*durationpb.Duration)(nil),         // 9: google.protobuf.Duration
	(*structpb.Struct)(nil),             // 10: google.protobuf.Struct
	(*PartialCaveatInfo)(nil),           // 11: authzed.api.v1.PartialCaveatInfo
}
var file_authzed_api_v1_debug_proto_depIdxs = []int32{
	4,  // 0: authzed.api.v1.DebugInformation.check:type_name -> authzed.api.v1.CheckDebugTrace
	7,  // 1: authzed.api.v1.CheckDebugTrace.resource:type_name -> authzed.api.v1.ObjectReference
	0,  // 2: authzed.api.v1.CheckDebugTrace.permission_type:type_name -> authzed.api.v1.CheckDebugTrace.PermissionType
	8,  // 3: authzed.api.v1.CheckDebugTrace.subject:type_name -> authzed.api.v1.SubjectReference
	1,  // 4: authzed.api.v1.CheckDebugTrace.result:type_name -> authzed.api.v1.CheckDebugTrace.Permissionship
	5,  // 5: authzed.api.v1.CheckDebugTrace.caveat_evaluation_info:type_name -> authzed.api.v1.CaveatEvalInfo
	9,  // 6: authzed.api.v1.CheckDebugTrace.duration:type_name -> google.protobuf.Duration
	6,  // 7: authzed.api.v1.CheckDebugTrace.sub_problems:type_name -> authzed.api.v1.CheckDebugTrace.SubProblems
	2,  // 8: authzed.api.v1.CaveatEvalInfo.result:type_name -> authzed.api.v1.CaveatEvalInfo.Result
	10, // 9: authzed.api.v1.CaveatEvalInfo.context:type_name -> google.protobuf.Struct
	11, // 10: authzed.api.v1.CaveatEvalInfo.partial_caveat_info:type_name -> authzed.api.v1.PartialCaveatInfo
	4,  // 11: authzed.api.v1.CheckDebugTrace.SubProblems.traces:type_name -> authzed.api.v1.CheckDebugTrace
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_authzed_api_v1_debug_proto_init() }
func file_authzed_api_v1_debug_proto_init() {
	if File_authzed_api_v1_debug_proto != nil {
		return
	}
	file_authzed_api_v1_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_authzed_api_v1_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInformation); i {
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
		file_authzed_api_v1_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckDebugTrace); i {
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
		file_authzed_api_v1_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaveatEvalInfo); i {
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
		file_authzed_api_v1_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckDebugTrace_SubProblems); i {
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
	file_authzed_api_v1_debug_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CheckDebugTrace_WasCachedResult)(nil),
		(*CheckDebugTrace_SubProblems_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authzed_api_v1_debug_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authzed_api_v1_debug_proto_goTypes,
		DependencyIndexes: file_authzed_api_v1_debug_proto_depIdxs,
		EnumInfos:         file_authzed_api_v1_debug_proto_enumTypes,
		MessageInfos:      file_authzed_api_v1_debug_proto_msgTypes,
	}.Build()
	File_authzed_api_v1_debug_proto = out.File
	file_authzed_api_v1_debug_proto_rawDesc = nil
	file_authzed_api_v1_debug_proto_goTypes = nil
	file_authzed_api_v1_debug_proto_depIdxs = nil
}
