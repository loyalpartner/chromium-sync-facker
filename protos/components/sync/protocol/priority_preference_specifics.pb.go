// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/priority_preference_specifics.proto

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

// Properties of a synced priority preference.
type PriorityPreferenceSpecifics struct {
	Preference           *PreferenceSpecifics `protobuf:"bytes,1,opt,name=preference" json:"preference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PriorityPreferenceSpecifics) Reset()         { *m = PriorityPreferenceSpecifics{} }
func (m *PriorityPreferenceSpecifics) String() string { return proto.CompactTextString(m) }
func (*PriorityPreferenceSpecifics) ProtoMessage()    {}
func (*PriorityPreferenceSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_priority_preference_specifics_10d3b648752a5e6b, []int{0}
}
func (m *PriorityPreferenceSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriorityPreferenceSpecifics.Unmarshal(m, b)
}
func (m *PriorityPreferenceSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriorityPreferenceSpecifics.Marshal(b, m, deterministic)
}
func (dst *PriorityPreferenceSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriorityPreferenceSpecifics.Merge(dst, src)
}
func (m *PriorityPreferenceSpecifics) XXX_Size() int {
	return xxx_messageInfo_PriorityPreferenceSpecifics.Size(m)
}
func (m *PriorityPreferenceSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_PriorityPreferenceSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_PriorityPreferenceSpecifics proto.InternalMessageInfo

func (m *PriorityPreferenceSpecifics) GetPreference() *PreferenceSpecifics {
	if m != nil {
		return m.Preference
	}
	return nil
}

func init() {
	proto.RegisterType((*PriorityPreferenceSpecifics)(nil), "sync_pb.PriorityPreferenceSpecifics")
}

func init() {
	proto.RegisterFile("components/sync/protocol/priority_preference_specifics.proto", fileDescriptor_priority_preference_specifics_10d3b648752a5e6b)
}

var fileDescriptor_priority_preference_specifics_10d3b648752a5e6b = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2f, 0xae, 0xcc, 0x4b, 0xd6, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x28, 0xca, 0xcc, 0x2f, 0xca, 0x2c, 0xa9, 0x8c, 0x2f, 0x28,
	0x4a, 0x4d, 0x4b, 0x2d, 0x4a, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x2e, 0x48, 0x4d, 0xce, 0x4c, 0xcb,
	0x4c, 0x2e, 0xd6, 0x03, 0x2b, 0x13, 0x62, 0x07, 0x69, 0x89, 0x2f, 0x48, 0x92, 0x32, 0xc6, 0x63,
	0x0c, 0x2e, 0xdd, 0x4a, 0xd1, 0x5c, 0xd2, 0x01, 0x50, 0x4b, 0x02, 0xe0, 0xaa, 0x82, 0x61, 0x8a,
	0x84, 0x6c, 0xb8, 0xb8, 0x10, 0x9a, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x64, 0xf4, 0xa0,
	0x36, 0xea, 0x61, 0xd1, 0x11, 0x84, 0xa4, 0xde, 0x49, 0x9b, 0x4b, 0x35, 0xbf, 0x28, 0x5d, 0x2f,
	0x39, 0xa3, 0x28, 0x3f, 0x37, 0xb3, 0x34, 0x57, 0x0f, 0xe1, 0x40, 0xb0, 0x11, 0x7a, 0x30, 0x07,
	0x7a, 0x30, 0x07, 0x30, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xed, 0xc4, 0xf2, 0xad, 0x06, 0x01,
	0x00, 0x00,
}
