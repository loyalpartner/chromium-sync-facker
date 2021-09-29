// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/security_event_specifics.proto

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

type SecurityEventSpecifics struct {
	// The specific security event to record.
	//
	// Types that are valid to be assigned to Event:
	//	*SecurityEventSpecifics_GaiaPasswordReuseEvent
	Event isSecurityEventSpecifics_Event `protobuf_oneof:"event"`
	// Time of event, as measured by client in microseconds since Windows epoch.
	EventTimeUsec        *int64   `protobuf:"varint,2,opt,name=event_time_usec,json=eventTimeUsec" json:"event_time_usec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecurityEventSpecifics) Reset()         { *m = SecurityEventSpecifics{} }
func (m *SecurityEventSpecifics) String() string { return proto.CompactTextString(m) }
func (*SecurityEventSpecifics) ProtoMessage()    {}
func (*SecurityEventSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_security_event_specifics_8e91775c7c4c49ff, []int{0}
}
func (m *SecurityEventSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityEventSpecifics.Unmarshal(m, b)
}
func (m *SecurityEventSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityEventSpecifics.Marshal(b, m, deterministic)
}
func (dst *SecurityEventSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityEventSpecifics.Merge(dst, src)
}
func (m *SecurityEventSpecifics) XXX_Size() int {
	return xxx_messageInfo_SecurityEventSpecifics.Size(m)
}
func (m *SecurityEventSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityEventSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityEventSpecifics proto.InternalMessageInfo

type isSecurityEventSpecifics_Event interface {
	isSecurityEventSpecifics_Event()
}

type SecurityEventSpecifics_GaiaPasswordReuseEvent struct {
	GaiaPasswordReuseEvent *GaiaPasswordReuse `protobuf:"bytes,1,opt,name=gaia_password_reuse_event,json=gaiaPasswordReuseEvent,oneof"`
}

func (*SecurityEventSpecifics_GaiaPasswordReuseEvent) isSecurityEventSpecifics_Event() {}

func (m *SecurityEventSpecifics) GetEvent() isSecurityEventSpecifics_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *SecurityEventSpecifics) GetGaiaPasswordReuseEvent() *GaiaPasswordReuse {
	if x, ok := m.GetEvent().(*SecurityEventSpecifics_GaiaPasswordReuseEvent); ok {
		return x.GaiaPasswordReuseEvent
	}
	return nil
}

func (m *SecurityEventSpecifics) GetEventTimeUsec() int64 {
	if m != nil && m.EventTimeUsec != nil {
		return *m.EventTimeUsec
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SecurityEventSpecifics) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SecurityEventSpecifics_OneofMarshaler, _SecurityEventSpecifics_OneofUnmarshaler, _SecurityEventSpecifics_OneofSizer, []interface{}{
		(*SecurityEventSpecifics_GaiaPasswordReuseEvent)(nil),
	}
}

func _SecurityEventSpecifics_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SecurityEventSpecifics)
	// event
	switch x := m.Event.(type) {
	case *SecurityEventSpecifics_GaiaPasswordReuseEvent:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GaiaPasswordReuseEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SecurityEventSpecifics.Event has unexpected type %T", x)
	}
	return nil
}

func _SecurityEventSpecifics_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SecurityEventSpecifics)
	switch tag {
	case 1: // event.gaia_password_reuse_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GaiaPasswordReuse)
		err := b.DecodeMessage(msg)
		m.Event = &SecurityEventSpecifics_GaiaPasswordReuseEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SecurityEventSpecifics_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SecurityEventSpecifics)
	// event
	switch x := m.Event.(type) {
	case *SecurityEventSpecifics_GaiaPasswordReuseEvent:
		s := proto.Size(x.GaiaPasswordReuseEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*SecurityEventSpecifics)(nil), "sync_pb.SecurityEventSpecifics")
}

func init() {
	proto.RegisterFile("components/sync/protocol/security_event_specifics.proto", fileDescriptor_security_event_specifics_8e91775c7c4c49ff)
}

var fileDescriptor_security_event_specifics_8e91775c7c4c49ff = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x5d, 0x8b, 0x14, 0x56, 0x44, 0xc8, 0xa1, 0xd4, 0x9e, 0x8a, 0xa0, 0x14, 0x84, 0x0d,
	0xf4, 0xe2, 0xbd, 0x20, 0xf6, 0x58, 0x52, 0xc5, 0xe3, 0xb2, 0x8e, 0x63, 0x1c, 0x70, 0x33, 0xcb,
	0xce, 0x46, 0xe9, 0xdf, 0xf1, 0x97, 0x4a, 0x36, 0x09, 0x82, 0xe8, 0xf5, 0xf1, 0xcd, 0xf7, 0xe6,
	0xe9, 0x5b, 0x60, 0x1f, 0xb8, 0xc1, 0x26, 0x49, 0x29, 0x87, 0x06, 0xca, 0x10, 0x39, 0x31, 0xf0,
	0x7b, 0x29, 0x08, 0x6d, 0xa4, 0x74, 0xb0, 0xf8, 0x81, 0x4d, 0xb2, 0x12, 0x10, 0xe8, 0x95, 0x40,
	0x4c, 0x26, 0x8a, 0x69, 0x47, 0xdb, 0xf0, 0xbc, 0x58, 0xff, 0x6b, 0xa8, 0x1d, 0x39, 0x1b, 0x9c,
	0xc8, 0x27, 0xc7, 0x17, 0x1b, 0xb1, 0x15, 0xec, 0x8f, 0x2f, 0xbf, 0x94, 0x9e, 0xed, 0x07, 0xff,
	0x5d, 0xa7, 0xdf, 0x8f, 0xf6, 0xe2, 0x49, 0x5f, 0xfc, 0x71, 0xd7, 0x3f, 0x31, 0x57, 0x4b, 0xb5,
	0x3a, 0x5d, 0x2f, 0xcc, 0xd0, 0x6d, 0xee, 0x1d, 0xb9, 0xdd, 0x00, 0x56, 0x1d, 0xb7, 0x3d, 0xaa,
	0x66, 0xf5, 0xef, 0x30, 0x37, 0x14, 0xd7, 0xfa, 0xbc, 0x5f, 0x92, 0xc8, 0xa3, 0x6d, 0x05, 0x61,
	0x7e, 0xbc, 0x54, 0xab, 0x49, 0x75, 0x96, 0xe3, 0x07, 0xf2, 0xf8, 0x28, 0x08, 0x9b, 0xa9, 0x3e,
	0xc9, 0xc1, 0xe6, 0x46, 0x5f, 0x71, 0xac, 0x0d, 0xbc, 0x45, 0xf6, 0xd4, 0x7a, 0xf3, 0xb3, 0x33,
	0xf7, 0x9b, 0x71, 0xe7, 0x76, 0xb2, 0x53, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0xaa, 0xd2,
	0x6c, 0x48, 0x01, 0x00, 0x00,
}