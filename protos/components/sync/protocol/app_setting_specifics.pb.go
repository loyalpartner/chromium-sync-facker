// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/app_setting_specifics.proto

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

// Properties of app setting sync objects; just an extension setting.
type AppSettingSpecifics struct {
	ExtensionSetting     *ExtensionSettingSpecifics `protobuf:"bytes,1,opt,name=extension_setting,json=extensionSetting" json:"extension_setting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AppSettingSpecifics) Reset()         { *m = AppSettingSpecifics{} }
func (m *AppSettingSpecifics) String() string { return proto.CompactTextString(m) }
func (*AppSettingSpecifics) ProtoMessage()    {}
func (*AppSettingSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_setting_specifics_9514f65bb72ca56a, []int{0}
}
func (m *AppSettingSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppSettingSpecifics.Unmarshal(m, b)
}
func (m *AppSettingSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppSettingSpecifics.Marshal(b, m, deterministic)
}
func (dst *AppSettingSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppSettingSpecifics.Merge(dst, src)
}
func (m *AppSettingSpecifics) XXX_Size() int {
	return xxx_messageInfo_AppSettingSpecifics.Size(m)
}
func (m *AppSettingSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_AppSettingSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_AppSettingSpecifics proto.InternalMessageInfo

func (m *AppSettingSpecifics) GetExtensionSetting() *ExtensionSettingSpecifics {
	if m != nil {
		return m.ExtensionSetting
	}
	return nil
}

func init() {
	proto.RegisterType((*AppSettingSpecifics)(nil), "sync_pb.AppSettingSpecifics")
}

func init() {
	proto.RegisterFile("components/sync/protocol/app_setting_specifics.proto", fileDescriptor_app_setting_specifics_9514f65bb72ca56a)
}

var fileDescriptor_app_setting_specifics_9514f65bb72ca56a = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xce, 0xcf, 0x2d,
	0xc8, 0xcf, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2f, 0xae, 0xcc, 0x4b, 0xd6, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x2c, 0x28, 0x88, 0x2f, 0x4e, 0x2d, 0x29, 0xc9, 0xcc, 0x4b,
	0x8f, 0x2f, 0x2e, 0x48, 0x4d, 0xce, 0x4c, 0xcb, 0x4c, 0x2e, 0xd6, 0x03, 0x4b, 0x0b, 0xb1, 0x83,
	0x94, 0xc6, 0x17, 0x24, 0x49, 0x59, 0xe1, 0xd4, 0x9e, 0x5a, 0x51, 0x92, 0x9a, 0x57, 0x9c, 0x99,
	0x9f, 0x87, 0xcb, 0x10, 0xa5, 0x34, 0x2e, 0x61, 0xc7, 0x82, 0x82, 0x60, 0x88, 0x6c, 0x30, 0x4c,
	0x52, 0xc8, 0x9f, 0x4b, 0x10, 0x43, 0xaf, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x92, 0x1e,
	0xd4, 0x5e, 0x3d, 0x57, 0x98, 0x0a, 0x74, 0xed, 0x41, 0x02, 0xa9, 0x68, 0x52, 0x4e, 0xda, 0x5c,
	0xaa, 0xf9, 0x45, 0xe9, 0x7a, 0xc9, 0x19, 0x45, 0xf9, 0xb9, 0x99, 0xa5, 0xb9, 0x7a, 0x08, 0x27,
	0x83, 0x8d, 0xd3, 0x83, 0x39, 0xd9, 0x83, 0x39, 0x80, 0x11, 0x10, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x84, 0x79, 0x5a, 0x10, 0x01, 0x00, 0x00,
}
