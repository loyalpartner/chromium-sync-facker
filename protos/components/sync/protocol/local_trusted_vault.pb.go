// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/local_trusted_vault.proto

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

type LocalTrustedVaultKey struct {
	// The actual key.
	KeyMaterial          []byte   `protobuf:"bytes,1,opt,name=key_material,json=keyMaterial" json:"key_material,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalTrustedVaultKey) Reset()         { *m = LocalTrustedVaultKey{} }
func (m *LocalTrustedVaultKey) String() string { return proto.CompactTextString(m) }
func (*LocalTrustedVaultKey) ProtoMessage()    {}
func (*LocalTrustedVaultKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_local_trusted_vault_05f18c20ca51ed55, []int{0}
}
func (m *LocalTrustedVaultKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTrustedVaultKey.Unmarshal(m, b)
}
func (m *LocalTrustedVaultKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTrustedVaultKey.Marshal(b, m, deterministic)
}
func (dst *LocalTrustedVaultKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTrustedVaultKey.Merge(dst, src)
}
func (m *LocalTrustedVaultKey) XXX_Size() int {
	return xxx_messageInfo_LocalTrustedVaultKey.Size(m)
}
func (m *LocalTrustedVaultKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTrustedVaultKey.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTrustedVaultKey proto.InternalMessageInfo

func (m *LocalTrustedVaultKey) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

type LocalDeviceRegistrationInfo struct {
	// Private SecureBox key.
	PrivateKeyMaterial []byte `protobuf:"bytes,1,opt,name=private_key_material,json=privateKeyMaterial" json:"private_key_material,omitempty"`
	// Indicates whether device is registered, i.e. whether its public key is
	// successfully submitted to the server.
	DeviceRegistered     *bool    `protobuf:"varint,2,opt,name=device_registered,json=deviceRegistered" json:"device_registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalDeviceRegistrationInfo) Reset()         { *m = LocalDeviceRegistrationInfo{} }
func (m *LocalDeviceRegistrationInfo) String() string { return proto.CompactTextString(m) }
func (*LocalDeviceRegistrationInfo) ProtoMessage()    {}
func (*LocalDeviceRegistrationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_local_trusted_vault_05f18c20ca51ed55, []int{1}
}
func (m *LocalDeviceRegistrationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDeviceRegistrationInfo.Unmarshal(m, b)
}
func (m *LocalDeviceRegistrationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDeviceRegistrationInfo.Marshal(b, m, deterministic)
}
func (dst *LocalDeviceRegistrationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDeviceRegistrationInfo.Merge(dst, src)
}
func (m *LocalDeviceRegistrationInfo) XXX_Size() int {
	return xxx_messageInfo_LocalDeviceRegistrationInfo.Size(m)
}
func (m *LocalDeviceRegistrationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDeviceRegistrationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDeviceRegistrationInfo proto.InternalMessageInfo

func (m *LocalDeviceRegistrationInfo) GetPrivateKeyMaterial() []byte {
	if m != nil {
		return m.PrivateKeyMaterial
	}
	return nil
}

func (m *LocalDeviceRegistrationInfo) GetDeviceRegistered() bool {
	if m != nil && m.DeviceRegistered != nil {
		return *m.DeviceRegistered
	}
	return false
}

type LocalTrustedVaultPerUser struct {
	// User identifier.
	GaiaId []byte `protobuf:"bytes,1,opt,name=gaia_id,json=gaiaId" json:"gaia_id,omitempty"`
	// All keys known for a user.
	VaultKey []*LocalTrustedVaultKey `protobuf:"bytes,2,rep,name=vault_key,json=vaultKey" json:"vault_key,omitempty"`
	// The version corresponding to the last element in |vault_key|.
	LastVaultKeyVersion *int32 `protobuf:"varint,3,opt,name=last_vault_key_version,json=lastVaultKeyVersion" json:"last_vault_key_version,omitempty"`
	// Indicates whether |vault_key| is stale, i.e. that the latest locally
	// available key isn't the latest key in the vault. New keys need to be
	// fetched through key retrieval procedure or by following key rotation.
	KeysAreStale *bool `protobuf:"varint,4,opt,name=keys_are_stale,json=keysAreStale" json:"keys_are_stale,omitempty"`
	// Device key and corresponding registration metadata.
	LocalDeviceRegistrationInfo *LocalDeviceRegistrationInfo `protobuf:"bytes,5,opt,name=local_device_registration_info,json=localDeviceRegistrationInfo" json:"local_device_registration_info,omitempty"`
	// The time (in milliseconds since UNIX epoch) at which last unsuccessful (due
	// to transient errors) request was sent to the vault service. Used for
	// throttling requests to the server.
	LastFailedRequestMillisSinceUnixEpoch *int64 `protobuf:"varint,6,opt,name=last_failed_request_millis_since_unix_epoch,json=lastFailedRequestMillisSinceUnixEpoch" json:"last_failed_request_millis_since_unix_epoch,omitempty"`
	// Whether keys relevant for the user should be deleted when account becomes
	// non-primary.
	ShouldDeleteKeysWhenNonPrimary *bool    `protobuf:"varint,7,opt,name=should_delete_keys_when_non_primary,json=shouldDeleteKeysWhenNonPrimary" json:"should_delete_keys_when_non_primary,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *LocalTrustedVaultPerUser) Reset()         { *m = LocalTrustedVaultPerUser{} }
func (m *LocalTrustedVaultPerUser) String() string { return proto.CompactTextString(m) }
func (*LocalTrustedVaultPerUser) ProtoMessage()    {}
func (*LocalTrustedVaultPerUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_local_trusted_vault_05f18c20ca51ed55, []int{2}
}
func (m *LocalTrustedVaultPerUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTrustedVaultPerUser.Unmarshal(m, b)
}
func (m *LocalTrustedVaultPerUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTrustedVaultPerUser.Marshal(b, m, deterministic)
}
func (dst *LocalTrustedVaultPerUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTrustedVaultPerUser.Merge(dst, src)
}
func (m *LocalTrustedVaultPerUser) XXX_Size() int {
	return xxx_messageInfo_LocalTrustedVaultPerUser.Size(m)
}
func (m *LocalTrustedVaultPerUser) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTrustedVaultPerUser.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTrustedVaultPerUser proto.InternalMessageInfo

func (m *LocalTrustedVaultPerUser) GetGaiaId() []byte {
	if m != nil {
		return m.GaiaId
	}
	return nil
}

func (m *LocalTrustedVaultPerUser) GetVaultKey() []*LocalTrustedVaultKey {
	if m != nil {
		return m.VaultKey
	}
	return nil
}

func (m *LocalTrustedVaultPerUser) GetLastVaultKeyVersion() int32 {
	if m != nil && m.LastVaultKeyVersion != nil {
		return *m.LastVaultKeyVersion
	}
	return 0
}

func (m *LocalTrustedVaultPerUser) GetKeysAreStale() bool {
	if m != nil && m.KeysAreStale != nil {
		return *m.KeysAreStale
	}
	return false
}

func (m *LocalTrustedVaultPerUser) GetLocalDeviceRegistrationInfo() *LocalDeviceRegistrationInfo {
	if m != nil {
		return m.LocalDeviceRegistrationInfo
	}
	return nil
}

func (m *LocalTrustedVaultPerUser) GetLastFailedRequestMillisSinceUnixEpoch() int64 {
	if m != nil && m.LastFailedRequestMillisSinceUnixEpoch != nil {
		return *m.LastFailedRequestMillisSinceUnixEpoch
	}
	return 0
}

func (m *LocalTrustedVaultPerUser) GetShouldDeleteKeysWhenNonPrimary() bool {
	if m != nil && m.ShouldDeleteKeysWhenNonPrimary != nil {
		return *m.ShouldDeleteKeysWhenNonPrimary
	}
	return false
}

type LocalTrustedVault struct {
	User                 []*LocalTrustedVaultPerUser `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LocalTrustedVault) Reset()         { *m = LocalTrustedVault{} }
func (m *LocalTrustedVault) String() string { return proto.CompactTextString(m) }
func (*LocalTrustedVault) ProtoMessage()    {}
func (*LocalTrustedVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_local_trusted_vault_05f18c20ca51ed55, []int{3}
}
func (m *LocalTrustedVault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalTrustedVault.Unmarshal(m, b)
}
func (m *LocalTrustedVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalTrustedVault.Marshal(b, m, deterministic)
}
func (dst *LocalTrustedVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalTrustedVault.Merge(dst, src)
}
func (m *LocalTrustedVault) XXX_Size() int {
	return xxx_messageInfo_LocalTrustedVault.Size(m)
}
func (m *LocalTrustedVault) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalTrustedVault.DiscardUnknown(m)
}

var xxx_messageInfo_LocalTrustedVault proto.InternalMessageInfo

func (m *LocalTrustedVault) GetUser() []*LocalTrustedVaultPerUser {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalTrustedVaultKey)(nil), "sync_pb.LocalTrustedVaultKey")
	proto.RegisterType((*LocalDeviceRegistrationInfo)(nil), "sync_pb.LocalDeviceRegistrationInfo")
	proto.RegisterType((*LocalTrustedVaultPerUser)(nil), "sync_pb.LocalTrustedVaultPerUser")
	proto.RegisterType((*LocalTrustedVault)(nil), "sync_pb.LocalTrustedVault")
}

func init() {
	proto.RegisterFile("components/sync/protocol/local_trusted_vault.proto", fileDescriptor_local_trusted_vault_05f18c20ca51ed55)
}

var fileDescriptor_local_trusted_vault_05f18c20ca51ed55 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6b, 0xd4, 0x40,
	0x10, 0xc6, 0x49, 0xd3, 0xf6, 0xea, 0x5e, 0x11, 0xbb, 0x16, 0x5d, 0x28, 0x96, 0xf4, 0xac, 0x10,
	0x28, 0xdc, 0xc9, 0x89, 0x0f, 0xfa, 0xa6, 0x54, 0xb1, 0x9e, 0x95, 0x92, 0xda, 0x0a, 0xbe, 0x0c,
	0x6b, 0x32, 0xd7, 0x5b, 0x6e, 0x6f, 0x37, 0xee, 0x6e, 0x62, 0x03, 0xfe, 0xad, 0xfe, 0x2d, 0xb2,
	0x9b, 0x9c, 0x5a, 0x7a, 0xfa, 0x98, 0x99, 0xdf, 0x7c, 0x99, 0xf9, 0xbe, 0x25, 0xe3, 0x5c, 0x2f,
	0x4a, 0xad, 0x50, 0x39, 0x3b, 0xb2, 0x8d, 0xca, 0x47, 0xa5, 0xd1, 0x4e, 0xe7, 0x5a, 0x8e, 0xa4,
	0xce, 0xb9, 0x04, 0x67, 0x2a, 0xeb, 0xb0, 0x80, 0x9a, 0x57, 0xd2, 0x0d, 0x43, 0x93, 0xf6, 0x3c,
	0x08, 0xe5, 0xd7, 0xc1, 0x0b, 0xb2, 0xfb, 0xc1, 0x53, 0x9f, 0x5a, 0xe8, 0xd2, 0x33, 0x13, 0x6c,
	0xe8, 0x01, 0xd9, 0x9e, 0x63, 0x03, 0x0b, 0xee, 0xd0, 0x08, 0x2e, 0x59, 0x94, 0x44, 0xe9, 0x76,
	0xd6, 0x9f, 0x63, 0x73, 0xda, 0x95, 0x06, 0x3f, 0xc8, 0x5e, 0x18, 0x3d, 0xc6, 0x5a, 0xe4, 0x98,
	0xe1, 0x95, 0xb0, 0xce, 0x70, 0x27, 0xb4, 0x3a, 0x51, 0x53, 0x4d, 0x9f, 0x92, 0xdd, 0xd2, 0x88,
	0x9a, 0x3b, 0x84, 0x15, 0x4a, 0xb4, 0xeb, 0x4d, 0xfe, 0x08, 0xd2, 0x23, 0xb2, 0x53, 0x04, 0x2d,
	0x30, 0x41, 0x0c, 0x0d, 0x16, 0x6c, 0x2d, 0x89, 0xd2, 0xad, 0xec, 0x5e, 0xf1, 0xd7, 0x4f, 0x7c,
	0x7d, 0xf0, 0x33, 0x26, 0xec, 0xd6, 0xe6, 0x67, 0x68, 0x2e, 0x2c, 0x1a, 0xfa, 0x90, 0xf4, 0xae,
	0xb8, 0xe0, 0x20, 0x8a, 0xee, 0x77, 0x9b, 0xfe, 0xf3, 0xa4, 0xa0, 0x2f, 0xc9, 0x9d, 0x60, 0x83,
	0x5f, 0x89, 0xad, 0x25, 0x71, 0xda, 0x1f, 0x3f, 0x1a, 0x76, 0x5e, 0x0c, 0x57, 0x19, 0x91, 0x6d,
	0xd5, 0x4b, 0x4b, 0x9e, 0x91, 0x07, 0x92, 0x5b, 0x07, 0xbf, 0x05, 0xa0, 0x46, 0x63, 0x85, 0x56,
	0x2c, 0x4e, 0xa2, 0x74, 0x23, 0xbb, 0xef, 0xbb, 0xcb, 0xb9, 0xcb, 0xb6, 0x45, 0x0f, 0xc9, 0xdd,
	0x39, 0x36, 0x16, 0xb8, 0x41, 0xb0, 0x8e, 0x4b, 0x64, 0xeb, 0xe1, 0x20, 0xef, 0xae, 0x7d, 0x65,
	0xf0, 0xdc, 0xd7, 0xa8, 0x20, 0xfb, 0x6d, 0x56, 0x37, 0xee, 0x6f, 0xcd, 0x04, 0xa1, 0xa6, 0x9a,
	0x6d, 0x24, 0x51, 0xda, 0x1f, 0x1f, 0xde, 0xdc, 0x75, 0xb5, 0xf3, 0xd9, 0x9e, 0xfc, 0x4f, 0x2c,
	0x5f, 0xc8, 0x51, 0xb8, 0x62, 0xca, 0x85, 0xc4, 0x02, 0x0c, 0x7e, 0xab, 0xd0, 0x3a, 0x58, 0x08,
	0x29, 0x85, 0x05, 0x2b, 0x54, 0x8e, 0x50, 0x29, 0x71, 0x0d, 0x58, 0xea, 0x7c, 0xc6, 0x36, 0x93,
	0x28, 0x8d, 0xb3, 0x27, 0x7e, 0xe4, 0x6d, 0x98, 0xc8, 0xda, 0x81, 0xd3, 0xc0, 0x9f, 0x7b, 0xfc,
	0x42, 0x89, 0xeb, 0x37, 0x1e, 0xa6, 0x13, 0xf2, 0xd8, 0xce, 0x74, 0x25, 0x0b, 0x28, 0x50, 0x62,
	0x1b, 0xbc, 0x85, 0xef, 0x33, 0x54, 0xa0, 0xb4, 0x82, 0xd2, 0x88, 0x05, 0x37, 0x0d, 0xeb, 0x05,
	0x07, 0xf6, 0x5b, 0xf4, 0x38, 0x90, 0x13, 0x6c, 0xec, 0xe7, 0x19, 0xaa, 0x8f, 0x5a, 0x9d, 0xb5,
	0xd4, 0xe0, 0x3d, 0xd9, 0xb9, 0x15, 0x08, 0x7d, 0x4e, 0xd6, 0x2b, 0x8b, 0x86, 0x45, 0x21, 0xba,
	0x83, 0x7f, 0x47, 0xd7, 0xbd, 0x84, 0x2c, 0xe0, 0xaf, 0xfb, 0x64, 0xf9, 0xe0, 0xdf, 0xc5, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x71, 0xb9, 0xe3, 0x30, 0x03, 0x00, 0x00,
}
