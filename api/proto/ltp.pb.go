// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/ltp.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{2, 0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{1}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{2}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

type VersionResponse struct {
	VersionString        string   `protobuf:"bytes,1,opt,name=VersionString,proto3" json:"VersionString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{3}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersionString() string {
	if m != nil {
		return m.VersionString
	}
	return ""
}

type Item struct {
	IRI                  string       `protobuf:"bytes,1,opt,name=IRI,proto3" json:"IRI,omitempty"`
	ItemTypes            []string     `protobuf:"bytes,2,rep,name=ItemTypes,proto3" json:"ItemTypes,omitempty"`
	Statements           []*Statement `protobuf:"bytes,3,rep,name=Statements,proto3" json:"Statements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{4}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetIRI() string {
	if m != nil {
		return m.IRI
	}
	return ""
}

func (m *Item) GetItemTypes() []string {
	if m != nil {
		return m.ItemTypes
	}
	return nil
}

func (m *Item) GetStatements() []*Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

// A semantic 'statement' about the world. Can generally be viewed as
// an RDF triple, with an additional 'Scope', used for provenance.
// This can be implemented as a named graph, where the graph name (or label)
// is used for this purpose.
//
// See also:
// <http://patterns.dataincubator.org/book/named-graphs.html>
type Statement struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate            string   `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object               string   `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	Scope                *Scope   `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{5}
}

func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Statement) GetPredicate() string {
	if m != nil {
		return m.Predicate
	}
	return ""
}

func (m *Statement) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *Statement) GetScope() *Scope {
	if m != nil {
		return m.Scope
	}
	return nil
}

// The scope bounds the set of statements or assertions being made by
// an agent.
// Example: scope := &Scope{Time.now(), "ltp_client.shawnlower.net", nil}
type Scope struct {
	Agent                string               `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	AssertionTime        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=assertionTime,proto3" json:"assertionTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Scope) Reset()         { *m = Scope{} }
func (m *Scope) String() string { return proto.CompactTextString(m) }
func (*Scope) ProtoMessage()    {}
func (*Scope) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{6}
}

func (m *Scope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scope.Unmarshal(m, b)
}
func (m *Scope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scope.Marshal(b, m, deterministic)
}
func (m *Scope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scope.Merge(m, src)
}
func (m *Scope) XXX_Size() int {
	return xxx_messageInfo_Scope.Size(m)
}
func (m *Scope) XXX_DiscardUnknown() {
	xxx_messageInfo_Scope.DiscardUnknown(m)
}

var xxx_messageInfo_Scope proto.InternalMessageInfo

func (m *Scope) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

func (m *Scope) GetAssertionTime() *timestamp.Timestamp {
	if m != nil {
		return m.AssertionTime
	}
	return nil
}

type CreateItemRequest struct {
	ItemTypes            []string     `protobuf:"bytes,1,rep,name=ItemTypes,proto3" json:"ItemTypes,omitempty"`
	Statements           []*Statement `protobuf:"bytes,2,rep,name=Statements,proto3" json:"Statements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateItemRequest) Reset()         { *m = CreateItemRequest{} }
func (m *CreateItemRequest) String() string { return proto.CompactTextString(m) }
func (*CreateItemRequest) ProtoMessage()    {}
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{7}
}

func (m *CreateItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateItemRequest.Unmarshal(m, b)
}
func (m *CreateItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateItemRequest.Marshal(b, m, deterministic)
}
func (m *CreateItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateItemRequest.Merge(m, src)
}
func (m *CreateItemRequest) XXX_Size() int {
	return xxx_messageInfo_CreateItemRequest.Size(m)
}
func (m *CreateItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateItemRequest proto.InternalMessageInfo

func (m *CreateItemRequest) GetItemTypes() []string {
	if m != nil {
		return m.ItemTypes
	}
	return nil
}

func (m *CreateItemRequest) GetStatements() []*Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

type CreateItemResponse struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateItemResponse) Reset()         { *m = CreateItemResponse{} }
func (m *CreateItemResponse) String() string { return proto.CompactTextString(m) }
func (*CreateItemResponse) ProtoMessage()    {}
func (*CreateItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{8}
}

func (m *CreateItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateItemResponse.Unmarshal(m, b)
}
func (m *CreateItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateItemResponse.Marshal(b, m, deterministic)
}
func (m *CreateItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateItemResponse.Merge(m, src)
}
func (m *CreateItemResponse) XXX_Size() int {
	return xxx_messageInfo_CreateItemResponse.Size(m)
}
func (m *CreateItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateItemResponse proto.InternalMessageInfo

func (m *CreateItemResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*HealthCheckRequest)(nil), "proto.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "proto.HealthCheckResponse")
	proto.RegisterType((*VersionResponse)(nil), "proto.VersionResponse")
	proto.RegisterType((*Item)(nil), "proto.Item")
	proto.RegisterType((*Statement)(nil), "proto.Statement")
	proto.RegisterType((*Scope)(nil), "proto.Scope")
	proto.RegisterType((*CreateItemRequest)(nil), "proto.CreateItemRequest")
	proto.RegisterType((*CreateItemResponse)(nil), "proto.CreateItemResponse")
}

func init() { proto.RegisterFile("api/proto/ltp.proto", fileDescriptor_83f39c07d8bb6265) }

var fileDescriptor_83f39c07d8bb6265 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0x4d, 0xd3, 0xd2, 0x1b, 0xeb, 0xd6, 0xbb, 0xb2, 0x64, 0x83, 0xb0, 0x65, 0xf0, 0xa1,
	0x4f, 0xa9, 0x44, 0x44, 0xf0, 0x41, 0x76, 0x59, 0x96, 0xdd, 0x22, 0x64, 0x25, 0xad, 0xeb, 0xe3,
	0x92, 0xc6, 0x6b, 0x1a, 0x6d, 0x3e, 0xcc, 0x4c, 0x85, 0x05, 0xc1, 0x1f, 0xe1, 0x1f, 0x96, 0xcc,
	0x4c, 0xda, 0x8d, 0x56, 0xf1, 0xa9, 0x73, 0xee, 0x39, 0x77, 0xee, 0x9d, 0xd3, 0x13, 0x38, 0x8c,
	0xca, 0x74, 0x5a, 0x56, 0x85, 0x28, 0xa6, 0x6b, 0x51, 0x7a, 0xf2, 0x84, 0x96, 0xfc, 0x71, 0x4f,
	0x92, 0xa2, 0x48, 0xd6, 0xa4, 0xe8, 0xe5, 0xe6, 0xd3, 0x54, 0xa4, 0x19, 0x71, 0x11, 0x65, 0x5a,
	0xc7, 0xfa, 0x60, 0x5d, 0x64, 0xa5, 0xb8, 0x63, 0x1e, 0xe0, 0x15, 0x45, 0x6b, 0xb1, 0x3a, 0x5f,
	0x51, 0xfc, 0x25, 0xa4, 0xaf, 0x1b, 0xe2, 0x02, 0x1d, 0xe8, 0x73, 0xaa, 0xbe, 0xa5, 0x31, 0x39,
	0xc6, 0xd8, 0x98, 0x0c, 0xc2, 0x06, 0xb2, 0x9f, 0x06, 0x1c, 0xb6, 0x1a, 0x78, 0x59, 0xe4, 0x9c,
	0xf0, 0x14, 0x7a, 0x5c, 0x44, 0x62, 0xc3, 0x65, 0xc3, 0x23, 0x7f, 0xa2, 0x06, 0x79, 0x7b, 0xb4,
	0xde, 0xbc, 0xbe, 0x2b, 0x4f, 0xe6, 0x52, 0x1f, 0xea, 0x3e, 0xf6, 0x1a, 0x86, 0x2d, 0x02, 0x6d,
	0xe8, 0xbf, 0x0f, 0xde, 0x06, 0xd7, 0x1f, 0x82, 0xd1, 0x83, 0x1a, 0xcc, 0x2f, 0xc2, 0x9b, 0x59,
	0x70, 0x39, 0x32, 0xf0, 0x00, 0xec, 0xe0, 0x7a, 0x71, 0xdb, 0x14, 0x3a, 0xec, 0x15, 0x1c, 0xdc,
	0x50, 0xc5, 0xd3, 0x22, 0xdf, 0x2e, 0xf4, 0x0c, 0x86, 0xba, 0x34, 0x17, 0x55, 0x9a, 0x27, 0xfa,
	0x21, 0xed, 0x22, 0x5b, 0x41, 0x77, 0x26, 0x28, 0xc3, 0x11, 0x98, 0xb3, 0x70, 0xa6, 0x35, 0xf5,
	0x11, 0x9f, 0xc2, 0xa0, 0x66, 0x16, 0x77, 0x25, 0x71, 0xa7, 0x33, 0x36, 0x27, 0x83, 0x70, 0x57,
	0xc0, 0xe7, 0x00, 0xf5, 0x96, 0x94, 0x51, 0x2e, 0xb8, 0x63, 0x8e, 0xcd, 0x89, 0xed, 0x8f, 0xf4,
	0x93, 0xb7, 0x44, 0x78, 0x4f, 0xc3, 0x7e, 0xc0, 0x60, 0x8b, 0xa4, 0xbf, 0x9b, 0xe5, 0x67, 0x8a,
	0xc5, 0xd6, 0x5f, 0x05, 0xeb, 0xb1, 0x65, 0x45, 0x1f, 0xd3, 0x38, 0x12, 0xe4, 0x74, 0x24, 0xb7,
	0x2b, 0xe0, 0x11, 0xf4, 0x0a, 0xd5, 0x66, 0x4a, 0x4a, 0x23, 0x64, 0x60, 0xf1, 0xb8, 0x28, 0xc9,
	0xe9, 0x8e, 0x8d, 0x89, 0xed, 0x3f, 0x6c, 0x36, 0xa9, 0x6b, 0xa1, 0xa2, 0xd8, 0x2d, 0x58, 0x12,
	0xe3, 0x13, 0xb0, 0xa2, 0x84, 0xf2, 0x66, 0xb4, 0x02, 0x78, 0x0a, 0xc3, 0x88, 0x73, 0xaa, 0x44,
	0x5a, 0xe4, 0x8b, 0x34, 0x53, 0xc3, 0x6d, 0xdf, 0xf5, 0x54, 0x94, 0xbc, 0x26, 0x4a, 0xde, 0xa2,
	0x89, 0x52, 0xd8, 0x6e, 0x60, 0x31, 0x3c, 0x3e, 0xaf, 0x28, 0x12, 0x54, 0xdb, 0xd4, 0x24, 0xa9,
	0x65, 0xa3, 0xf1, 0x6f, 0x1b, 0x3b, 0xff, 0x61, 0xe3, 0x4b, 0xc0, 0xfb, 0x43, 0xf4, 0x9f, 0x7d,
	0x02, 0xdd, 0x54, 0x50, 0x26, 0x5f, 0x64, 0xfb, 0xb6, 0xbe, 0x41, 0x4a, 0x24, 0xe1, 0x5f, 0x41,
	0x4f, 0x25, 0x11, 0xdf, 0x80, 0x25, 0xd3, 0x88, 0xc7, 0xfb, 0x12, 0x2a, 0x97, 0x76, 0xdd, 0xbf,
	0x87, 0xd7, 0xff, 0x0e, 0xe6, 0xd9, 0xbb, 0x19, 0xfa, 0x00, 0x97, 0x24, 0x74, 0x98, 0xb0, 0x31,
	0x5c, 0x7e, 0x53, 0xee, 0x91, 0x46, 0xbf, 0x47, 0xf2, 0x0c, 0x60, 0xb7, 0x3b, 0x3a, 0x5a, 0xf5,
	0x87, 0x67, 0xee, 0xf1, 0x1e, 0x46, 0x5d, 0xb1, 0xec, 0x49, 0xe6, 0xc5, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe3, 0x5a, 0xe1, 0x5e, 0xfd, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/proto.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/ltp.proto",
}

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	GetVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetVersion(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/proto.API/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, "/proto.API/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	GetVersion(context.Context, *Empty) (*VersionResponse, error)
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetVersion(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _API_GetVersion_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _API_CreateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/ltp.proto",
}