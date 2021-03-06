// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/loopback_server.proto

package sync_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Entity type mapping to one of the subclasses of LoopbackServerEntity.
type LoopbackServerEntity_Type int32

const (
	LoopbackServerEntity_UNKNOWN   LoopbackServerEntity_Type = 0
	LoopbackServerEntity_BOOKMARK  LoopbackServerEntity_Type = 1
	LoopbackServerEntity_PERMANENT LoopbackServerEntity_Type = 2
	LoopbackServerEntity_TOMBSTONE LoopbackServerEntity_Type = 3
	LoopbackServerEntity_UNIQUE    LoopbackServerEntity_Type = 4
)

var LoopbackServerEntity_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "BOOKMARK",
	2: "PERMANENT",
	3: "TOMBSTONE",
	4: "UNIQUE",
}
var LoopbackServerEntity_Type_value = map[string]int32{
	"UNKNOWN":   0,
	"BOOKMARK":  1,
	"PERMANENT": 2,
	"TOMBSTONE": 3,
	"UNIQUE":    4,
}

func (x LoopbackServerEntity_Type) Enum() *LoopbackServerEntity_Type {
	p := new(LoopbackServerEntity_Type)
	*p = x
	return p
}
func (x LoopbackServerEntity_Type) String() string {
	return proto.EnumName(LoopbackServerEntity_Type_name, int32(x))
}
func (x *LoopbackServerEntity_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LoopbackServerEntity_Type_value, data, "LoopbackServerEntity_Type")
	if err != nil {
		return err
	}
	*x = LoopbackServerEntity_Type(value)
	return nil
}
func (LoopbackServerEntity_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_loopback_server_49b14bbbe2270758, []int{0, 0}
}

// Serialization of the LoopbackServerEntity and its ancestors.
type LoopbackServerEntity struct {
	Type                 *LoopbackServerEntity_Type `protobuf:"varint,1,opt,name=type,enum=sync_pb.LoopbackServerEntity_Type" json:"type,omitempty"`
	Entity               *SyncEntity                `protobuf:"bytes,2,opt,name=entity" json:"entity,omitempty"`
	ModelType            *int64                     `protobuf:"varint,3,opt,name=model_type,json=modelType" json:"model_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *LoopbackServerEntity) Reset()         { *m = LoopbackServerEntity{} }
func (m *LoopbackServerEntity) String() string { return proto.CompactTextString(m) }
func (*LoopbackServerEntity) ProtoMessage()    {}
func (*LoopbackServerEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_loopback_server_49b14bbbe2270758, []int{0}
}
func (m *LoopbackServerEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoopbackServerEntity.Unmarshal(m, b)
}
func (m *LoopbackServerEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoopbackServerEntity.Marshal(b, m, deterministic)
}
func (dst *LoopbackServerEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoopbackServerEntity.Merge(dst, src)
}
func (m *LoopbackServerEntity) XXX_Size() int {
	return xxx_messageInfo_LoopbackServerEntity.Size(m)
}
func (m *LoopbackServerEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_LoopbackServerEntity.DiscardUnknown(m)
}

var xxx_messageInfo_LoopbackServerEntity proto.InternalMessageInfo

func (m *LoopbackServerEntity) GetType() LoopbackServerEntity_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return LoopbackServerEntity_UNKNOWN
}

func (m *LoopbackServerEntity) GetEntity() *SyncEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *LoopbackServerEntity) GetModelType() int64 {
	if m != nil && m.ModelType != nil {
		return *m.ModelType
	}
	return 0
}

// Contains the loopback server state.
type LoopbackServerProto struct {
	// The protocol buffer format version.
	Version       *int64                  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	StoreBirthday *int64                  `protobuf:"varint,2,opt,name=store_birthday,json=storeBirthday" json:"store_birthday,omitempty"`
	Entities      []*LoopbackServerEntity `protobuf:"bytes,3,rep,name=entities" json:"entities,omitempty"`
	// All Keystore keys known to the server.
	KeystoreKeys [][]byte `protobuf:"bytes,4,rep,name=keystore_keys,json=keystoreKeys" json:"keystore_keys,omitempty"`
	// The last entity ID that was assigned to an entity.
	LastVersionAssigned  *int64   `protobuf:"varint,5,opt,name=last_version_assigned,json=lastVersionAssigned" json:"last_version_assigned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoopbackServerProto) Reset()         { *m = LoopbackServerProto{} }
func (m *LoopbackServerProto) String() string { return proto.CompactTextString(m) }
func (*LoopbackServerProto) ProtoMessage()    {}
func (*LoopbackServerProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_loopback_server_49b14bbbe2270758, []int{1}
}
func (m *LoopbackServerProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoopbackServerProto.Unmarshal(m, b)
}
func (m *LoopbackServerProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoopbackServerProto.Marshal(b, m, deterministic)
}
func (dst *LoopbackServerProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoopbackServerProto.Merge(dst, src)
}
func (m *LoopbackServerProto) XXX_Size() int {
	return xxx_messageInfo_LoopbackServerProto.Size(m)
}
func (m *LoopbackServerProto) XXX_DiscardUnknown() {
	xxx_messageInfo_LoopbackServerProto.DiscardUnknown(m)
}

var xxx_messageInfo_LoopbackServerProto proto.InternalMessageInfo

func (m *LoopbackServerProto) GetVersion() int64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *LoopbackServerProto) GetStoreBirthday() int64 {
	if m != nil && m.StoreBirthday != nil {
		return *m.StoreBirthday
	}
	return 0
}

func (m *LoopbackServerProto) GetEntities() []*LoopbackServerEntity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *LoopbackServerProto) GetKeystoreKeys() [][]byte {
	if m != nil {
		return m.KeystoreKeys
	}
	return nil
}

func (m *LoopbackServerProto) GetLastVersionAssigned() int64 {
	if m != nil && m.LastVersionAssigned != nil {
		return *m.LastVersionAssigned
	}
	return 0
}

func init() {
	proto.RegisterType((*LoopbackServerEntity)(nil), "sync_pb.LoopbackServerEntity")
	proto.RegisterType((*LoopbackServerProto)(nil), "sync_pb.LoopbackServerProto")
	proto.RegisterEnum("sync_pb.LoopbackServerEntity_Type", LoopbackServerEntity_Type_name, LoopbackServerEntity_Type_value)
}

func init() {
	proto.RegisterFile("components/sync/protocol/loopback_server.proto", fileDescriptor_loopback_server_49b14bbbe2270758)
}

var fileDescriptor_loopback_server_49b14bbbe2270758 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x9d, 0x2a, 0xaf, 0x69, 0x6f, 0x92, 0x62, 0x94, 0x0d, 0xcc, 0xa0, 0x60, 0x5c, 0x0a, 0x86,
	0x82, 0x03, 0x7e, 0x18, 0xec, 0x31, 0x06, 0xc3, 0x86, 0x17, 0x3b, 0x53, 0x9c, 0xed, 0xd1, 0x38,
	0x8e, 0x68, 0x4d, 0x6d, 0xcb, 0x48, 0x5a, 0xc0, 0xbf, 0x78, 0xb0, 0x5f, 0x31, 0x2c, 0x3b, 0xd9,
	0x06, 0x5b, 0x9f, 0xa4, 0x7b, 0xce, 0xb9, 0xe7, 0x7e, 0x48, 0xe0, 0x15, 0xbc, 0x6e, 0x79, 0xc3,
	0x1a, 0x25, 0x97, 0xb2, 0x6b, 0x8a, 0x65, 0x2b, 0xb8, 0xe2, 0x05, 0xaf, 0x96, 0x15, 0xe7, 0xed,
	0x3e, 0x2f, 0x9e, 0x33, 0xc9, 0xc4, 0x91, 0x09, 0x4f, 0x13, 0x64, 0xd2, 0x8b, 0xb2, 0x76, 0xff,
	0xee, 0xee, 0xbf, 0x89, 0x7d, 0x34, 0xa8, 0x9d, 0x1f, 0x08, 0xde, 0x7c, 0x1e, 0x7d, 0xb6, 0xda,
	0x26, 0x6c, 0x54, 0xa9, 0x3a, 0xf2, 0x1e, 0x0c, 0xd5, 0xb5, 0xcc, 0x42, 0x36, 0x72, 0x6f, 0x7c,
	0xc7, 0x1b, 0x5d, 0xbd, 0x7f, 0x89, 0xbd, 0xb4, 0x6b, 0x19, 0xd5, 0x7a, 0xf2, 0x00, 0x97, 0x4c,
	0x83, 0xd6, 0x85, 0x8d, 0xdc, 0xa9, 0xbf, 0x38, 0x67, 0x6e, 0xbb, 0xa6, 0x18, 0xf4, 0x74, 0x94,
	0x90, 0x5b, 0x80, 0x9a, 0x1f, 0x58, 0x95, 0xe9, 0x52, 0xd8, 0x46, 0x2e, 0xa6, 0xd7, 0x1a, 0xe9,
	0x1d, 0x9d, 0x08, 0x8c, 0xfe, 0x24, 0x53, 0x98, 0xec, 0xe2, 0x28, 0x4e, 0xbe, 0xc5, 0xe6, 0x2b,
	0x32, 0x83, 0xab, 0x20, 0x49, 0xa2, 0xf5, 0x8a, 0x46, 0x26, 0x22, 0x73, 0xb8, 0xde, 0x84, 0x74,
	0xbd, 0x8a, 0xc3, 0x38, 0x35, 0x2f, 0xfa, 0x30, 0x4d, 0xd6, 0xc1, 0x36, 0x4d, 0xe2, 0xd0, 0xc4,
	0x04, 0xe0, 0x72, 0x17, 0x7f, 0xfa, 0xb2, 0x0b, 0x4d, 0xc3, 0xf9, 0x89, 0x60, 0xf1, 0x77, 0xf3,
	0x1b, 0xbd, 0x2f, 0x0b, 0x26, 0x47, 0x26, 0x64, 0xc9, 0x1b, 0x3d, 0x2b, 0xa6, 0xa7, 0x90, 0xdc,
	0xc3, 0x8d, 0x54, 0x5c, 0xb0, 0x6c, 0x5f, 0x0a, 0xf5, 0x74, 0xc8, 0x87, 0x91, 0x30, 0x9d, 0x6b,
	0x34, 0x18, 0x41, 0xf2, 0x01, 0xae, 0xf4, 0x38, 0x25, 0x93, 0x16, 0xb6, 0xb1, 0x3b, 0xf5, 0x6f,
	0x5f, 0xdc, 0x16, 0x3d, 0xcb, 0xc9, 0x1d, 0xcc, 0x9f, 0x59, 0x37, 0x14, 0xe9, 0x2f, 0x96, 0x61,
	0x63, 0x77, 0x46, 0x67, 0x27, 0x30, 0x62, 0x9d, 0x24, 0x3e, 0xbc, 0xad, 0x72, 0xa9, 0xb2, 0xb1,
	0xad, 0x2c, 0x97, 0xb2, 0x7c, 0x6c, 0xd8, 0xc1, 0x7a, 0xad, 0xbb, 0x59, 0xf4, 0xe4, 0xd7, 0x81,
	0x5b, 0x8d, 0x54, 0xf0, 0x00, 0xf7, 0x5c, 0x3c, 0x7a, 0xc5, 0x93, 0xe0, 0x75, 0xf9, 0xbd, 0xfe,
	0xe3, 0x0f, 0x79, 0xbf, 0x1f, 0xbf, 0xe0, 0xd5, 0x47, 0xbc, 0x41, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x96, 0x3e, 0x7d, 0x62, 0x02, 0x00, 0x00,
}
