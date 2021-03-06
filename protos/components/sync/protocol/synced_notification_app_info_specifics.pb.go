// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/synced_notification_app_info_specifics.proto

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

// This message is kept around for backwards compatibility sake.
type SyncedNotificationAppInfoSpecifics struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncedNotificationAppInfoSpecifics) Reset()         { *m = SyncedNotificationAppInfoSpecifics{} }
func (m *SyncedNotificationAppInfoSpecifics) String() string { return proto.CompactTextString(m) }
func (*SyncedNotificationAppInfoSpecifics) ProtoMessage()    {}
func (*SyncedNotificationAppInfoSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_synced_notification_app_info_specifics_79ae92be70e83fb1, []int{0}
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Unmarshal(m, b)
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Marshal(b, m, deterministic)
}
func (dst *SyncedNotificationAppInfoSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Merge(dst, src)
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_Size() int {
	return xxx_messageInfo_SyncedNotificationAppInfoSpecifics.Size(m)
}
func (m *SyncedNotificationAppInfoSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncedNotificationAppInfoSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_SyncedNotificationAppInfoSpecifics proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SyncedNotificationAppInfoSpecifics)(nil), "sync_pb.SyncedNotificationAppInfoSpecifics")
}

func init() {
	proto.RegisterFile("components/sync/protocol/synced_notification_app_info_specifics.proto", fileDescriptor_synced_notification_app_info_specifics_79ae92be70e83fb1)
}

var fileDescriptor_synced_notification_app_info_specifics_79ae92be70e83fb1 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2f, 0xae, 0xcc, 0x4b, 0xd6, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0x01, 0xf3, 0x52, 0x53, 0xe2, 0xf3, 0xf2, 0x4b, 0x32, 0xd3, 0x32, 0x93,
	0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xe2, 0x13, 0x0b, 0x0a, 0xe2, 0x33, 0xf3, 0xd2, 0xf2, 0xe3, 0x8b,
	0x0b, 0x52, 0x93, 0x41, 0xc2, 0xc5, 0x7a, 0x60, 0xf5, 0x42, 0xec, 0x20, 0xd5, 0xf1, 0x05, 0x49,
	0x4a, 0x2a, 0x5c, 0x4a, 0xc1, 0x60, 0x8d, 0x7e, 0x48, 0xfa, 0x1c, 0x0b, 0x0a, 0x3c, 0xf3, 0xd2,
	0xf2, 0x83, 0x61, 0x9a, 0x9c, 0xb4, 0xb9, 0x54, 0xf3, 0x8b, 0xd2, 0xf5, 0x92, 0x33, 0x8a, 0xf2,
	0x73, 0x33, 0x4b, 0x73, 0xf5, 0x10, 0x8e, 0xd0, 0x03, 0x19, 0xa4, 0x07, 0x73, 0x84, 0x07, 0x73,
	0x00, 0x23, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x9f, 0x9e, 0x44, 0xa3, 0x00, 0x00, 0x00,
}
