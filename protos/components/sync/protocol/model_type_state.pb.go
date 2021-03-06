// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/model_type_state.proto

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

// Sync proto to store data type global metadata in model type storage.
type ModelTypeState struct {
	// The latest progress markers received from the server.
	ProgressMarker *DataTypeProgressMarker `protobuf:"bytes,1,opt,name=progress_marker,json=progressMarker" json:"progress_marker,omitempty"`
	// A data type context.  Sent to the server in every commit or update
	// request.  May be updated by either responses from the server or requests
	// made on the model thread.  The interpretation of this value may be
	// data-type specific.  Many data types ignore it.
	TypeContext *DataTypeContext `protobuf:"bytes,2,opt,name=type_context,json=typeContext" json:"type_context,omitempty"`
	// This value is set if this type's data should be encrypted on the server.
	// If this key changes, the client will need to re-commit all of its local
	// data to the server using the new encryption key.
	EncryptionKeyName *string `protobuf:"bytes,3,opt,name=encryption_key_name,json=encryptionKeyName" json:"encryption_key_name,omitempty"`
	// This flag is set to true when the first download cycle is complete.  The
	// ModelTypeProcessor should not attempt to commit any items until this
	// flag is set.
	InitialSyncDone *bool `protobuf:"varint,4,opt,name=initial_sync_done,json=initialSyncDone" json:"initial_sync_done,omitempty"`
	// A GUID that identifies the committing sync client. It's persisted within
	// the sync metadata and should be used to check the integrity of the
	// metadata. Mismatches with the guid of the running client indicates invalid
	// persisted sync metadata, because cache_guid is reset when sync is disabled,
	// and disabling sync is supposed to clear sync metadata.
	CacheGuid *string `protobuf:"bytes,5,opt,name=cache_guid,json=cacheGuid" json:"cache_guid,omitempty"`
	// Syncing account ID, representing the user.
	AuthenticatedAccountId *string  `protobuf:"bytes,6,opt,name=authenticated_account_id,json=authenticatedAccountId" json:"authenticated_account_id,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ModelTypeState) Reset()         { *m = ModelTypeState{} }
func (m *ModelTypeState) String() string { return proto.CompactTextString(m) }
func (*ModelTypeState) ProtoMessage()    {}
func (*ModelTypeState) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_type_state_4132d339a5c855d4, []int{0}
}
func (m *ModelTypeState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelTypeState.Unmarshal(m, b)
}
func (m *ModelTypeState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelTypeState.Marshal(b, m, deterministic)
}
func (dst *ModelTypeState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelTypeState.Merge(dst, src)
}
func (m *ModelTypeState) XXX_Size() int {
	return xxx_messageInfo_ModelTypeState.Size(m)
}
func (m *ModelTypeState) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelTypeState.DiscardUnknown(m)
}

var xxx_messageInfo_ModelTypeState proto.InternalMessageInfo

func (m *ModelTypeState) GetProgressMarker() *DataTypeProgressMarker {
	if m != nil {
		return m.ProgressMarker
	}
	return nil
}

func (m *ModelTypeState) GetTypeContext() *DataTypeContext {
	if m != nil {
		return m.TypeContext
	}
	return nil
}

func (m *ModelTypeState) GetEncryptionKeyName() string {
	if m != nil && m.EncryptionKeyName != nil {
		return *m.EncryptionKeyName
	}
	return ""
}

func (m *ModelTypeState) GetInitialSyncDone() bool {
	if m != nil && m.InitialSyncDone != nil {
		return *m.InitialSyncDone
	}
	return false
}

func (m *ModelTypeState) GetCacheGuid() string {
	if m != nil && m.CacheGuid != nil {
		return *m.CacheGuid
	}
	return ""
}

func (m *ModelTypeState) GetAuthenticatedAccountId() string {
	if m != nil && m.AuthenticatedAccountId != nil {
		return *m.AuthenticatedAccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*ModelTypeState)(nil), "sync_pb.ModelTypeState")
}

func init() {
	proto.RegisterFile("components/sync/protocol/model_type_state.proto", fileDescriptor_model_type_state_4132d339a5c855d4)
}

var fileDescriptor_model_type_state_4132d339a5c855d4 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x40, 0xe9, 0xf6, 0x7d, 0xea, 0x32, 0xd9, 0x58, 0x05, 0x09, 0x82, 0x38, 0x14, 0x61, 0x28,
	0xb4, 0xe0, 0x93, 0xe0, 0x93, 0x3a, 0x70, 0x22, 0x93, 0xd1, 0xf9, 0x1e, 0x62, 0x7a, 0xd9, 0xc2,
	0xda, 0xdc, 0x90, 0xde, 0x82, 0xf9, 0x8f, 0xfe, 0x28, 0x69, 0x36, 0x28, 0x43, 0x7c, 0xbc, 0x39,
	0xe7, 0x1e, 0x42, 0xc2, 0x52, 0x85, 0xa5, 0x45, 0x03, 0x86, 0xaa, 0xb4, 0xf2, 0x46, 0xa5, 0xd6,
	0x21, 0xa1, 0xc2, 0x22, 0x2d, 0x31, 0x87, 0x42, 0x90, 0xb7, 0x20, 0x2a, 0x92, 0x04, 0x49, 0x20,
	0xf1, 0x61, 0x63, 0x09, 0xfb, 0x79, 0x76, 0xf5, 0xe7, 0x66, 0x33, 0x6d, 0xed, 0xcb, 0xef, 0x0e,
	0x1b, 0xcc, 0x9b, 0xd0, 0x87, 0xb7, 0xb0, 0x6c, 0x32, 0xf1, 0x8c, 0x0d, 0xad, 0xc3, 0x95, 0x83,
	0xaa, 0x12, 0xa5, 0x74, 0x1b, 0x70, 0x3c, 0x1a, 0x47, 0x93, 0xfe, 0xdd, 0x45, 0xb2, 0x4b, 0x27,
	0x53, 0x49, 0xb2, 0x59, 0x58, 0xec, 0xbc, 0x79, 0xd0, 0xb2, 0x81, 0xdd, 0x9b, 0xe3, 0x07, 0x76,
	0x1c, 0xae, 0xa7, 0xd0, 0x10, 0x7c, 0x11, 0xef, 0x84, 0x0c, 0xff, 0x95, 0x79, 0xde, 0xf2, 0xac,
	0x4f, 0xed, 0x10, 0x27, 0xec, 0x04, 0x8c, 0x72, 0xde, 0x92, 0x46, 0x23, 0x36, 0xe0, 0x85, 0x91,
	0x25, 0xf0, 0xee, 0x38, 0x9a, 0xf4, 0xb2, 0x51, 0x8b, 0xde, 0xc0, 0xbf, 0xcb, 0x12, 0xe2, 0x1b,
	0x36, 0xd2, 0x46, 0x93, 0x96, 0x85, 0x08, 0xfd, 0x1c, 0x0d, 0xf0, 0x7f, 0xe3, 0x68, 0x72, 0x94,
	0x0d, 0x77, 0x60, 0xe9, 0x8d, 0x9a, 0xa2, 0x81, 0xf8, 0x9c, 0x31, 0x25, 0xd5, 0x1a, 0xc4, 0xaa,
	0xd6, 0x39, 0xff, 0x1f, 0x92, 0xbd, 0x70, 0xf2, 0x52, 0xeb, 0x3c, 0xbe, 0x67, 0x5c, 0xd6, 0xb4,
	0x06, 0x43, 0x5a, 0x49, 0x82, 0x5c, 0x48, 0xa5, 0xb0, 0x36, 0x24, 0x74, 0xce, 0x0f, 0x82, 0x7c,
	0xba, 0xc7, 0x1f, 0xb7, 0xf8, 0x35, 0x7f, 0xba, 0x65, 0xd7, 0xe8, 0x56, 0x89, 0x5a, 0x3b, 0x2c,
	0x75, 0x5d, 0x26, 0xed, 0x17, 0x24, 0xed, 0xa3, 0x2b, 0x2c, 0x66, 0xdd, 0x45, 0xf4, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x9d, 0x29, 0xd5, 0x34, 0xdb, 0x01, 0x00, 0x00,
}
