// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetArtifactsRequest struct {
	// Deprecated
	IncludeEventArtifacts  bool     `protobuf:"varint,1,opt,name=include_event_artifacts,json=includeEventArtifacts,proto3" json:"include_event_artifacts,omitempty"`
	IncludeServerArtifacts bool     `protobuf:"varint,2,opt,name=include_server_artifacts,json=includeServerArtifacts,proto3" json:"include_server_artifacts,omitempty"`
	SearchTerm             string   `protobuf:"bytes,3,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	NumberOfResults        uint64   `protobuf:"varint,4,opt,name=number_of_results,json=numberOfResults,proto3" json:"number_of_results,omitempty"`
	Type                   string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Names                  []string `protobuf:"bytes,6,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetArtifactsRequest) Reset()         { *m = GetArtifactsRequest{} }
func (m *GetArtifactsRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactsRequest) ProtoMessage()    {}
func (*GetArtifactsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{0}
}
func (m *GetArtifactsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactsRequest.Unmarshal(m, b)
}
func (m *GetArtifactsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactsRequest.Marshal(b, m, deterministic)
}
func (dst *GetArtifactsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactsRequest.Merge(dst, src)
}
func (m *GetArtifactsRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactsRequest.Size(m)
}
func (m *GetArtifactsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactsRequest proto.InternalMessageInfo

func (m *GetArtifactsRequest) GetIncludeEventArtifacts() bool {
	if m != nil {
		return m.IncludeEventArtifacts
	}
	return false
}

func (m *GetArtifactsRequest) GetIncludeServerArtifacts() bool {
	if m != nil {
		return m.IncludeServerArtifacts
	}
	return false
}

func (m *GetArtifactsRequest) GetSearchTerm() string {
	if m != nil {
		return m.SearchTerm
	}
	return ""
}

func (m *GetArtifactsRequest) GetNumberOfResults() uint64 {
	if m != nil {
		return m.NumberOfResults
	}
	return 0
}

func (m *GetArtifactsRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetArtifactsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type GetArtifactRequest struct {
	// Deprecated.
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactRequest) Reset()         { *m = GetArtifactRequest{} }
func (m *GetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactRequest) ProtoMessage()    {}
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{1}
}
func (m *GetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactRequest.Unmarshal(m, b)
}
func (m *GetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *GetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactRequest.Merge(dst, src)
}
func (m *GetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactRequest.Size(m)
}
func (m *GetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactRequest proto.InternalMessageInfo

func (m *GetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

type GetArtifactResponse struct {
	Artifact             string   `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactResponse) Reset()         { *m = GetArtifactResponse{} }
func (m *GetArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactResponse) ProtoMessage()    {}
func (*GetArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{2}
}
func (m *GetArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactResponse.Unmarshal(m, b)
}
func (m *GetArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactResponse.Marshal(b, m, deterministic)
}
func (dst *GetArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactResponse.Merge(dst, src)
}
func (m *GetArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactResponse.Size(m)
}
func (m *GetArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactResponse proto.InternalMessageInfo

func (m *GetArtifactResponse) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type SetArtifactRequest struct {
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Artifact             string   `protobuf:"bytes,2,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetArtifactRequest) Reset()         { *m = SetArtifactRequest{} }
func (m *SetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*SetArtifactRequest) ProtoMessage()    {}
func (*SetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{3}
}
func (m *SetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetArtifactRequest.Unmarshal(m, b)
}
func (m *SetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *SetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetArtifactRequest.Merge(dst, src)
}
func (m *SetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_SetArtifactRequest.Size(m)
}
func (m *SetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetArtifactRequest proto.InternalMessageInfo

func (m *SetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *SetArtifactRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type APIResponse struct {
	Error                bool     `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIResponse) Reset()         { *m = APIResponse{} }
func (m *APIResponse) String() string { return proto.CompactTextString(m) }
func (*APIResponse) ProtoMessage()    {}
func (*APIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{4}
}
func (m *APIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIResponse.Unmarshal(m, b)
}
func (m *APIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIResponse.Marshal(b, m, deterministic)
}
func (dst *APIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIResponse.Merge(dst, src)
}
func (m *APIResponse) XXX_Size() int {
	return xxx_messageInfo_APIResponse.Size(m)
}
func (m *APIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_APIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_APIResponse proto.InternalMessageInfo

func (m *APIResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *APIResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type GetReportRequest struct {
	Artifact string `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Format   string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	// Common parameters
	ClientId string `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Parameters for MONITORING_DAILY
	DayName   string `protobuf:"bytes,6,opt,name=day_name,json=dayName,proto3" json:"day_name,omitempty"`
	StartTime uint64 `protobuf:"varint,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   uint64 `protobuf:"varint,9,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Parameters for CLIENT
	FlowId               string           `protobuf:"bytes,7,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Parameters           []*proto1.VQLEnv `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetReportRequest) Reset()         { *m = GetReportRequest{} }
func (m *GetReportRequest) String() string { return proto.CompactTextString(m) }
func (*GetReportRequest) ProtoMessage()    {}
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{5}
}
func (m *GetReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportRequest.Unmarshal(m, b)
}
func (m *GetReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportRequest.Marshal(b, m, deterministic)
}
func (dst *GetReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportRequest.Merge(dst, src)
}
func (m *GetReportRequest) XXX_Size() int {
	return xxx_messageInfo_GetReportRequest.Size(m)
}
func (m *GetReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportRequest proto.InternalMessageInfo

func (m *GetReportRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

func (m *GetReportRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetReportRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *GetReportRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetReportRequest) GetDayName() string {
	if m != nil {
		return m.DayName
	}
	return ""
}

func (m *GetReportRequest) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *GetReportRequest) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *GetReportRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *GetReportRequest) GetParameters() []*proto1.VQLEnv {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// This presents the report in a form that can be rendered in the
// GUI. The data is presented in two parts - first "data" contains a
// json encoded object, then "template" is an angular template to be
// evaluated with the data.
type GetReportResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Template             string   `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Messages             []string `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReportResponse) Reset()         { *m = GetReportResponse{} }
func (m *GetReportResponse) String() string { return proto.CompactTextString(m) }
func (*GetReportResponse) ProtoMessage()    {}
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{6}
}
func (m *GetReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportResponse.Unmarshal(m, b)
}
func (m *GetReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportResponse.Marshal(b, m, deterministic)
}
func (dst *GetReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportResponse.Merge(dst, src)
}
func (m *GetReportResponse) XXX_Size() int {
	return xxx_messageInfo_GetReportResponse.Size(m)
}
func (m *GetReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportResponse proto.InternalMessageInfo

func (m *GetReportResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *GetReportResponse) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *GetReportResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Deprecated.
type ArtifactCompressionDict struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactCompressionDict) Reset()         { *m = ArtifactCompressionDict{} }
func (m *ArtifactCompressionDict) String() string { return proto.CompactTextString(m) }
func (*ArtifactCompressionDict) ProtoMessage()    {}
func (*ArtifactCompressionDict) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_602307f40f681219, []int{7}
}
func (m *ArtifactCompressionDict) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCompressionDict.Unmarshal(m, b)
}
func (m *ArtifactCompressionDict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCompressionDict.Marshal(b, m, deterministic)
}
func (dst *ArtifactCompressionDict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCompressionDict.Merge(dst, src)
}
func (m *ArtifactCompressionDict) XXX_Size() int {
	return xxx_messageInfo_ArtifactCompressionDict.Size(m)
}
func (m *ArtifactCompressionDict) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCompressionDict.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCompressionDict proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetArtifactsRequest)(nil), "proto.GetArtifactsRequest")
	proto.RegisterType((*GetArtifactRequest)(nil), "proto.GetArtifactRequest")
	proto.RegisterType((*GetArtifactResponse)(nil), "proto.GetArtifactResponse")
	proto.RegisterType((*SetArtifactRequest)(nil), "proto.SetArtifactRequest")
	proto.RegisterType((*APIResponse)(nil), "proto.APIResponse")
	proto.RegisterType((*GetReportRequest)(nil), "proto.GetReportRequest")
	proto.RegisterType((*GetReportResponse)(nil), "proto.GetReportResponse")
	proto.RegisterType((*ArtifactCompressionDict)(nil), "proto.ArtifactCompressionDict")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_artifacts_602307f40f681219) }

var fileDescriptor_artifacts_602307f40f681219 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0xe3, 0xc4, 0x3f, 0x13, 0xaa, 0xb6, 0x83, 0x20, 0xdb, 0x50, 0x60, 0x64, 0x10, 0x32,
	0x08, 0xad, 0xf9, 0x91, 0xa0, 0x8a, 0x00, 0x61, 0x93, 0x10, 0x59, 0xb4, 0x49, 0x99, 0x5a, 0x48,
	0xa8, 0x17, 0xab, 0xe9, 0xee, 0x59, 0xef, 0x48, 0xbb, 0x33, 0x9b, 0x99, 0xe3, 0x75, 0x73, 0xcd,
	0xab, 0xf0, 0x0e, 0xbc, 0x00, 0xaf, 0xc0, 0x0b, 0xc0, 0x6b, 0x70, 0x81, 0x66, 0x66, 0xed, 0x3a,
	0x77, 0x11, 0x17, 0xbd, 0xda, 0xd9, 0xf3, 0x9d, 0xef, 0x3b, 0x67, 0xce, 0xcf, 0x90, 0xbb, 0xc2,
	0xa0, 0xcc, 0x45, 0x8a, 0x36, 0xae, 0x8d, 0x46, 0x4d, 0x0f, 0xfc, 0xe7, 0xf8, 0x64, 0xbd, 0x5e,
	0xc7, 0x0d, 0x94, 0x3a, 0x95, 0x19, 0xbc, 0x8c, 0x53, 0x5d, 0x4d, 0x96, 0xba, 0x14, 0x6a, 0x39,
	0x09, 0x46, 0x23, 0x6a, 0xd4, 0x66, 0xe2, 0x9d, 0x27, 0x16, 0x2a, 0xa1, 0x50, 0xa6, 0x41, 0xe2,
	0xf8, 0xdb, 0xdb, 0x71, 0x45, 0x8a, 0x52, 0x2b, 0xdb, 0x6a, 0x34, 0x57, 0x65, 0xa0, 0x8f, 0x7e,
	0xdf, 0x23, 0x6f, 0x9e, 0x03, 0x4e, 0x37, 0x89, 0x71, 0xb8, 0x5a, 0x81, 0x45, 0xfa, 0x15, 0x39,
	0x92, 0x2a, 0x2d, 0x57, 0x19, 0x24, 0xd0, 0x80, 0xc2, 0x64, 0x9b, 0x7a, 0xd4, 0x61, 0x9d, 0xf1,
	0x80, 0xbf, 0xd5, 0xc2, 0x67, 0x0e, 0xdd, 0xd2, 0xe9, 0x23, 0x12, 0x6d, 0x78, 0x16, 0x4c, 0x03,
	0x66, 0x87, 0xb8, 0xe7, 0x89, 0x6f, 0xb7, 0xf8, 0x33, 0x0f, 0xbf, 0x62, 0xbe, 0x4f, 0x0e, 0x2d,
	0x08, 0x93, 0x16, 0x09, 0x82, 0xa9, 0xa2, 0x2e, 0xeb, 0x8c, 0x87, 0x9c, 0x04, 0xd3, 0x02, 0x4c,
	0x45, 0x3f, 0x21, 0xf7, 0xd5, 0xaa, 0x7a, 0x01, 0x26, 0xd1, 0x79, 0x62, 0xc0, 0xae, 0x4a, 0xb4,
	0xd1, 0x3e, 0xeb, 0x8c, 0xf7, 0xf9, 0xdd, 0x00, 0x5c, 0xe6, 0x3c, 0x98, 0x29, 0x25, 0xfb, 0x78,
	0x5d, 0x43, 0x74, 0xe0, 0x55, 0xfc, 0x99, 0x9e, 0x90, 0x03, 0x25, 0x2a, 0xb0, 0x51, 0x8f, 0x75,
	0xc7, 0xc3, 0xd9, 0x87, 0x7f, 0xff, 0xfb, 0xcf, 0x9f, 0x9d, 0xf7, 0xe8, 0xc3, 0x29, 0x2b, 0xa5,
	0x45, 0xa6, 0x73, 0xb6, 0x4d, 0x95, 0xa1, 0x66, 0x39, 0x60, 0x5a, 0xf0, 0x40, 0x19, 0x5d, 0x11,
	0xba, 0x53, 0xa5, 0x4d, 0x91, 0x9e, 0x93, 0x41, 0x93, 0xdb, 0xa4, 0x16, 0x58, 0xf8, 0xaa, 0x0c,
	0x67, 0xdf, 0x7b, 0xd1, 0x13, 0xfa, 0x68, 0x51, 0x00, 0x6b, 0x72, 0xcb, 0x1c, 0xc6, 0x0c, 0x94,
	0x02, 0x65, 0x03, 0x4e, 0x16, 0x0b, 0xd8, 0x89, 0x93, 0x41, 0x2e, 0x95, 0x74, 0xad, 0x61, 0x16,
	0xb5, 0x81, 0x98, 0xf7, 0x9b, 0xdc, 0x3e, 0x15, 0x58, 0x8c, 0x9e, 0xdf, 0x68, 0x0c, 0x07, 0x5b,
	0x6b, 0x65, 0x81, 0x9e, 0x92, 0xc1, 0x86, 0xde, 0xc6, 0x1c, 0xfb, 0x98, 0x23, 0xca, 0x16, 0x3b,
	0xd2, 0x2c, 0x13, 0x28, 0x3e, 0x65, 0xda, 0x30, 0xe1, 0x82, 0x88, 0x55, 0x89, 0x31, 0xdf, 0x32,
	0x47, 0x7f, 0x74, 0x08, 0x7d, 0xf6, 0x7a, 0x2f, 0x74, 0x23, 0xf3, 0xbd, 0xff, 0x9d, 0xf9, 0x4b,
	0x72, 0x38, 0x7d, 0x3a, 0xdf, 0x29, 0xc7, 0x01, 0x18, 0xa3, 0x4d, 0x98, 0xca, 0x59, 0xec, 0x15,
	0xc7, 0xf4, 0xa3, 0xa9, 0x62, 0xde, 0xce, 0x74, 0x9a, 0xae, 0x0c, 0x64, 0xcc, 0x02, 0xa2, 0x54,
	0xcb, 0x1b, 0xe9, 0xc6, 0x3c, 0x90, 0xe9, 0x07, 0xe4, 0x8e, 0x3f, 0x24, 0x15, 0x58, 0x2b, 0x96,
	0x10, 0xf2, 0xe3, 0x6f, 0x78, 0xe3, 0x93, 0x60, 0x1b, 0xfd, 0xd5, 0x25, 0xf7, 0xce, 0x01, 0x39,
	0xd4, 0xda, 0x6c, 0x2b, 0x76, 0xdb, 0x76, 0xe4, 0xda, 0xb0, 0x75, 0x21, 0xd3, 0x82, 0xad, 0x81,
	0x99, 0x20, 0xb1, 0x65, 0xd2, 0x9f, 0xda, 0x71, 0x0d, 0x65, 0xf9, 0xda, 0x2b, 0x7c, 0x4e, 0x27,
	0x4e, 0x21, 0xf8, 0x32, 0x07, 0x3b, 0xaa, 0x02, 0xc8, 0xd8, 0x18, 0xe2, 0x65, 0xcc, 0x9e, 0x5c,
	0x5e, 0xcc, 0x17, 0x97, 0x7c, 0x7e, 0x71, 0x9e, 0x9c, 0x4e, 0xe7, 0x8f, 0x7f, 0xfd, 0xb8, 0x9d,
	0xf3, 0xcf, 0x48, 0x2f, 0xd7, 0xa6, 0x12, 0x18, 0x76, 0x68, 0x16, 0x79, 0x39, 0x4a, 0xef, 0xfd,
	0xe8, 0xad, 0xcc, 0x93, 0x0b, 0xac, 0x4a, 0xde, 0xfa, 0xd1, 0x77, 0xc8, 0x30, 0x2d, 0xa5, 0xdb,
	0x72, 0x99, 0xb5, 0x2b, 0x33, 0x08, 0x86, 0x79, 0x46, 0x1f, 0x90, 0x41, 0x26, 0xae, 0x13, 0xb7,
	0x07, 0x51, 0xcf, 0x63, 0xfd, 0x4c, 0x5c, 0x5f, 0x88, 0x0a, 0xe8, 0xbb, 0x84, 0x58, 0x14, 0x06,
	0x13, 0x94, 0x15, 0x44, 0x03, 0xbf, 0x8a, 0x43, 0x6f, 0x59, 0xc8, 0x0a, 0x1c, 0x13, 0x54, 0x16,
	0xc0, 0xa1, 0x07, 0xfb, 0xa0, 0x32, 0x0f, 0x1d, 0x91, 0x7e, 0x5e, 0xea, 0xb5, 0x8b, 0xd7, 0xf7,
	0x9a, 0x3d, 0xf7, 0x3b, 0xcf, 0xe8, 0x15, 0x21, 0xb5, 0x30, 0xa2, 0x02, 0x04, 0xe3, 0xb6, 0xbb,
	0x3b, 0x3e, 0xfc, 0xe2, 0x4e, 0x78, 0xab, 0xe2, 0x5f, 0x7e, 0x7e, 0x7c, 0xa6, 0x9a, 0xd9, 0xcc,
	0xdf, 0xe7, 0x1b, 0x7a, 0x12, 0x3a, 0xc1, 0x5e, 0xf9, 0xc7, 0xae, 0xe6, 0x16, 0x58, 0x06, 0x35,
	0xa8, 0x8c, 0x69, 0xe5, 0xbb, 0xed, 0x4b, 0xa7, 0x73, 0x7f, 0x0e, 0xd5, 0x8c, 0xf9, 0x4e, 0x90,
	0xd1, 0x6f, 0x1d, 0x72, 0x7f, 0xa7, 0xaf, 0xed, 0x60, 0x51, 0xb2, 0xef, 0x06, 0x31, 0x34, 0x95,
	0xfb, 0x33, 0x3d, 0x26, 0x03, 0x84, 0xaa, 0x2e, 0x05, 0x6e, 0x26, 0x64, 0xfb, 0x4f, 0xbf, 0x23,
	0x83, 0x76, 0x78, 0x6c, 0xd4, 0xf5, 0x0f, 0xcc, 0xc8, 0xe7, 0xf9, 0x90, 0x1e, 0x9f, 0x85, 0x41,
	0x34, 0x6c, 0x2d, 0x8c, 0x72, 0x33, 0xb8, 0x71, 0x8c, 0xf9, 0x96, 0x33, 0x7a, 0x40, 0x8e, 0x36,
	0xdb, 0xf8, 0x83, 0xae, 0x6a, 0x03, 0xd6, 0x4a, 0xad, 0x4e, 0x65, 0x8a, 0x2f, 0x7a, 0xfe, 0xfa,
	0x5f, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x01, 0x78, 0x26, 0x9c, 0x3f, 0x06, 0x00, 0x00,
}
