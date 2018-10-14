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
	Subject              string   `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Predicate            string   `protobuf:"bytes,2,opt,name=Predicate,proto3" json:"Predicate,omitempty"`
	Object               string   `protobuf:"bytes,3,opt,name=Object,proto3" json:"Object,omitempty"`
	Label                string   `protobuf:"bytes,4,opt,name=Label,proto3" json:"Label,omitempty"`
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

func (m *Statement) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
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

type ServerInfoResponse struct {
	InfoItems            map[string]string `protobuf:"bytes,1,rep,name=InfoItems,proto3" json:"InfoItems,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServerInfoResponse) Reset()         { *m = ServerInfoResponse{} }
func (m *ServerInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ServerInfoResponse) ProtoMessage()    {}
func (*ServerInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f39c07d8bb6265, []int{9}
}

func (m *ServerInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfoResponse.Unmarshal(m, b)
}
func (m *ServerInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfoResponse.Marshal(b, m, deterministic)
}
func (m *ServerInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfoResponse.Merge(m, src)
}
func (m *ServerInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ServerInfoResponse.Size(m)
}
func (m *ServerInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfoResponse proto.InternalMessageInfo

func (m *ServerInfoResponse) GetInfoItems() map[string]string {
	if m != nil {
		return m.InfoItems
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
	proto.RegisterType((*ServerInfoResponse)(nil), "proto.ServerInfoResponse")
	proto.RegisterMapType((map[string]string)(nil), "proto.ServerInfoResponse.InfoItemsEntry")
}

func init() { proto.RegisterFile("api/proto/ltp.proto", fileDescriptor_83f39c07d8bb6265) }

var fileDescriptor_83f39c07d8bb6265 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0xfd, 0x64, 0xf9, 0x07, 0x5f, 0x7d, 0x4e, 0xdc, 0x49, 0x09, 0x8a, 0x28, 0xc4, 0x88, 0x2e,
	0xbc, 0x92, 0x8b, 0x4a, 0x69, 0x08, 0xa5, 0x24, 0x04, 0xd7, 0x11, 0x2d, 0x76, 0x90, 0xdc, 0x74,
	0x19, 0x64, 0xf5, 0xc6, 0x56, 0x6d, 0xfd, 0x44, 0x33, 0x0e, 0xf8, 0x39, 0xba, 0xeb, 0x83, 0xf4,
	0xf9, 0xca, 0x8c, 0x46, 0x72, 0x14, 0x27, 0xa5, 0x2b, 0xe9, 0xdc, 0x73, 0xee, 0xfc, 0x9c, 0x39,
	0x17, 0x0e, 0xfc, 0x34, 0x1c, 0xa4, 0x59, 0xc2, 0x92, 0xc1, 0x8a, 0xa5, 0x96, 0xf8, 0x23, 0x0d,
	0xf1, 0x31, 0x8e, 0xe7, 0x49, 0x32, 0x5f, 0x61, 0x4e, 0xcf, 0xd6, 0xb7, 0x03, 0x16, 0x46, 0x48,
	0x99, 0x1f, 0x49, 0x9d, 0xd9, 0x82, 0xc6, 0x30, 0x4a, 0xd9, 0xc6, 0xb4, 0x80, 0x5c, 0xa2, 0xbf,
	0x62, 0x8b, 0x8b, 0x05, 0x06, 0x4b, 0x17, 0xef, 0xd6, 0x48, 0x19, 0xd1, 0xa1, 0x45, 0x31, 0xbb,
	0x0f, 0x03, 0xd4, 0x95, 0x9e, 0xd2, 0x6f, 0xbb, 0x05, 0x34, 0x7f, 0x2a, 0x70, 0x50, 0x69, 0xa0,
	0x69, 0x12, 0x53, 0x24, 0x67, 0xd0, 0xa4, 0xcc, 0x67, 0x6b, 0x2a, 0x1a, 0xf6, 0xec, 0x7e, 0xbe,
	0x91, 0xf5, 0x84, 0xd6, 0xf2, 0xf8, 0x5a, 0xf1, 0xdc, 0x13, 0x7a, 0x57, 0xf6, 0x99, 0xa7, 0xd0,
	0xa9, 0x10, 0x44, 0x83, 0xd6, 0xd7, 0xf1, 0xe7, 0xf1, 0xe4, 0xdb, 0xb8, 0xfb, 0x1f, 0x07, 0xde,
	0xd0, 0xbd, 0x76, 0xc6, 0xa3, 0xae, 0x42, 0xf6, 0x41, 0x1b, 0x4f, 0xa6, 0x37, 0x45, 0xa1, 0x66,
	0xbe, 0x87, 0xfd, 0x6b, 0xcc, 0x68, 0x98, 0xc4, 0xe5, 0x81, 0x5e, 0x43, 0x47, 0x96, 0x3c, 0x96,
	0x85, 0xf1, 0x5c, 0x5e, 0xa4, 0x5a, 0x34, 0x17, 0x50, 0x77, 0x18, 0x46, 0xa4, 0x0b, 0xaa, 0xe3,
	0x3a, 0x52, 0xc3, 0x7f, 0xc9, 0x2b, 0x68, 0x73, 0x66, 0xba, 0x49, 0x91, 0xea, 0xb5, 0x9e, 0xda,
	0x6f, 0xbb, 0xdb, 0x02, 0x79, 0x03, 0xc0, 0x4f, 0x89, 0x11, 0xc6, 0x8c, 0xea, 0x6a, 0x4f, 0xed,
	0x6b, 0x76, 0x57, 0x5e, 0xb9, 0x24, 0xdc, 0x07, 0x1a, 0xf3, 0x0e, 0xda, 0x25, 0xe2, 0xfe, 0x7a,
	0xeb, 0xd9, 0x0f, 0x0c, 0x58, 0xe1, 0xaf, 0x84, 0x7c, 0xdb, 0xab, 0x0c, 0xbf, 0x87, 0x81, 0xcf,
	0x50, 0xaf, 0x09, 0x6e, 0x5b, 0x20, 0x87, 0xd0, 0x9c, 0xe4, 0x6d, 0xaa, 0xa0, 0x24, 0x22, 0x2f,
	0xa1, 0xf1, 0xc5, 0x9f, 0xe1, 0x4a, 0xaf, 0x8b, 0x72, 0x0e, 0xcc, 0x1b, 0x68, 0x78, 0x41, 0x92,
	0x22, 0xa7, 0xfd, 0x39, 0xc6, 0xc5, 0x66, 0x39, 0x20, 0x67, 0xd0, 0xf1, 0x29, 0xc5, 0x8c, 0x85,
	0x49, 0x3c, 0x0d, 0xa3, 0x7c, 0x3b, 0xcd, 0x36, 0xac, 0x3c, 0x3c, 0x56, 0x11, 0x1e, 0x6b, 0x5a,
	0x84, 0xc7, 0xad, 0x36, 0x98, 0x01, 0xbc, 0xb8, 0xc8, 0xd0, 0x67, 0xc8, 0x8d, 0x29, 0xb2, 0x53,
	0x31, 0x4e, 0xf9, 0xbb, 0x71, 0xb5, 0x7f, 0x30, 0xee, 0x1d, 0x90, 0x87, 0x9b, 0xc8, 0xe7, 0x3d,
	0x86, 0x7a, 0xc8, 0x30, 0x12, 0x37, 0xd2, 0x6c, 0x4d, 0xae, 0x20, 0x24, 0x82, 0x30, 0x7f, 0x29,
	0x40, 0x78, 0x9e, 0x30, 0x73, 0xe2, 0xdb, 0xa4, 0xec, 0xfb, 0x04, 0x6d, 0x8e, 0xb9, 0x30, 0x3f,
	0x9d, 0x56, 0x46, 0x75, 0x57, 0x6d, 0x95, 0xd2, 0x61, 0xcc, 0xb2, 0x8d, 0xbb, 0x6d, 0x35, 0x3e,
	0xc0, 0x5e, 0x95, 0xe4, 0x11, 0x5a, 0xe2, 0xa6, 0x88, 0xd0, 0x12, 0x37, 0xdc, 0xf6, 0x7b, 0x7f,
	0xb5, 0x2e, 0xde, 0x31, 0x07, 0xa7, 0xb5, 0x13, 0xc5, 0xbe, 0x84, 0x66, 0x3e, 0x18, 0xe4, 0x23,
	0x34, 0xc4, 0x70, 0x90, 0xa3, 0xa7, 0x06, 0x46, 0x38, 0x6a, 0x18, 0xcf, 0xcf, 0x92, 0xfd, 0x5b,
	0x01, 0xf5, 0xfc, 0xca, 0x21, 0x36, 0xc0, 0x08, 0x99, 0x0c, 0x37, 0xf9, 0x5f, 0x76, 0x88, 0x19,
	0x37, 0x0e, 0x25, 0x7a, 0x3c, 0x22, 0x27, 0xd0, 0x19, 0x21, 0xdb, 0x5e, 0xfb, 0x51, 0xdb, 0xd1,
	0xb3, 0xbe, 0x90, 0x73, 0x80, 0xed, 0x9b, 0x10, 0x5d, 0x0a, 0x77, 0xb2, 0x50, 0x2e, 0xb1, 0xfb,
	0x80, 0xb3, 0xa6, 0x60, 0xde, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x1e, 0xe4, 0xf4, 0xc7,
	0x04, 0x00, 0x00,
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
	GetServerInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerInfoResponse, error)
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

func (c *aPIClient) GetServerInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerInfoResponse, error) {
	out := new(ServerInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.API/GetServerInfo", in, out, opts...)
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
	GetServerInfo(context.Context, *Empty) (*ServerInfoResponse, error)
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

func _API_GetServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/GetServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetServerInfo(ctx, req.(*Empty))
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
			MethodName: "GetServerInfo",
			Handler:    _API_GetServerInfo_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _API_CreateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/ltp.proto",
}
