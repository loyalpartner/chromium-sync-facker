// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/entity_metadata.proto

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

// Sync proto to store entity metadata in model type storage.
type EntityMetadata struct {
	// A hash based on the client tag and model type.
	// Used for various map lookups. Should always be available.
	// Sent to the server as SyncEntity::client_defined_unique_tag.
	ClientTagHash *string `protobuf:"bytes,1,opt,name=client_tag_hash,json=clientTagHash" json:"client_tag_hash,omitempty"`
	// The entity's server-assigned ID.
	//
	// Prior to the item's first commit, we leave this value as an empty string.
	// The initial ID for a newly created item has to meet certain uniqueness
	// requirements, and we handle those on the sync thread.
	ServerId *string `protobuf:"bytes,2,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	// Whether or not the entity is deleted.
	IsDeleted *bool `protobuf:"varint,3,opt,name=is_deleted,json=isDeleted" json:"is_deleted,omitempty"`
	// A version number used to track in-progress commits. Each local change
	// increments this number.
	SequenceNumber *int64 `protobuf:"varint,4,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// The sequence number of the last item known to be successfully committed.
	AckedSequenceNumber *int64 `protobuf:"varint,5,opt,name=acked_sequence_number,json=ackedSequenceNumber" json:"acked_sequence_number,omitempty"`
	// The server version on which this item is based.
	//
	// If there are no local changes, this is the version of the entity as we see
	// it here.
	//
	// If there are local changes, this is the version of the entity on which
	// those changes are based.
	ServerVersion *int64 `protobuf:"varint,6,opt,name=server_version,json=serverVersion,def=-1" json:"server_version,omitempty"`
	// Entity creation and modification timestamps. Assigned by the client and
	// synced by the server, though the server usually doesn't bother to inspect
	// their values. They are encoded as milliseconds since the Unix epoch.
	CreationTime     *int64 `protobuf:"varint,7,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	ModificationTime *int64 `protobuf:"varint,8,opt,name=modification_time,json=modificationTime" json:"modification_time,omitempty"`
	// A hash of the current entity specifics value. Used to detect whether
	// entity's specifics value has changed without having to keep specifics in
	// memory.
	SpecificsHash *string `protobuf:"bytes,9,opt,name=specifics_hash,json=specificsHash" json:"specifics_hash,omitempty"`
	// A hash of the last specifics known by both the client and server. Used to
	// detect when local commits and remote updates are just for encryption. This
	// value will be the empty string only in the following cases: the entity is
	// in sync with the server, has never been synced, or is deleted.
	BaseSpecificsHash *string `protobuf:"bytes,10,opt,name=base_specifics_hash,json=baseSpecificsHash" json:"base_specifics_hash,omitempty"`
	// Used for positioning entities among their siblings. Relevant only for data
	// types that support positions (e.g bookmarks). Refer to its definition in
	// unique_position.proto for more information about its internal
	// representation.
	UniquePosition *UniquePosition `protobuf:"bytes,11,opt,name=unique_position,json=uniquePosition" json:"unique_position,omitempty"`
	// Used only for bookmarks. It's analogous to |specifics_hash| but it
	// exclusively hashes the content of the favicon image, as represented in
	// proto field BookmarkSpecifics.favicon, using base::PersistentHash().
	BookmarkFaviconHash  *uint32  `protobuf:"fixed32,12,opt,name=bookmark_favicon_hash,json=bookmarkFaviconHash" json:"bookmark_favicon_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityMetadata) Reset()         { *m = EntityMetadata{} }
func (m *EntityMetadata) String() string { return proto.CompactTextString(m) }
func (*EntityMetadata) ProtoMessage()    {}
func (*EntityMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_metadata_a18bc036f040841a, []int{0}
}
func (m *EntityMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityMetadata.Unmarshal(m, b)
}
func (m *EntityMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityMetadata.Marshal(b, m, deterministic)
}
func (dst *EntityMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityMetadata.Merge(dst, src)
}
func (m *EntityMetadata) XXX_Size() int {
	return xxx_messageInfo_EntityMetadata.Size(m)
}
func (m *EntityMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_EntityMetadata proto.InternalMessageInfo

const Default_EntityMetadata_ServerVersion int64 = -1

func (m *EntityMetadata) GetClientTagHash() string {
	if m != nil && m.ClientTagHash != nil {
		return *m.ClientTagHash
	}
	return ""
}

func (m *EntityMetadata) GetServerId() string {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return ""
}

func (m *EntityMetadata) GetIsDeleted() bool {
	if m != nil && m.IsDeleted != nil {
		return *m.IsDeleted
	}
	return false
}

func (m *EntityMetadata) GetSequenceNumber() int64 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

func (m *EntityMetadata) GetAckedSequenceNumber() int64 {
	if m != nil && m.AckedSequenceNumber != nil {
		return *m.AckedSequenceNumber
	}
	return 0
}

func (m *EntityMetadata) GetServerVersion() int64 {
	if m != nil && m.ServerVersion != nil {
		return *m.ServerVersion
	}
	return Default_EntityMetadata_ServerVersion
}

func (m *EntityMetadata) GetCreationTime() int64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *EntityMetadata) GetModificationTime() int64 {
	if m != nil && m.ModificationTime != nil {
		return *m.ModificationTime
	}
	return 0
}

func (m *EntityMetadata) GetSpecificsHash() string {
	if m != nil && m.SpecificsHash != nil {
		return *m.SpecificsHash
	}
	return ""
}

func (m *EntityMetadata) GetBaseSpecificsHash() string {
	if m != nil && m.BaseSpecificsHash != nil {
		return *m.BaseSpecificsHash
	}
	return ""
}

func (m *EntityMetadata) GetUniquePosition() *UniquePosition {
	if m != nil {
		return m.UniquePosition
	}
	return nil
}

func (m *EntityMetadata) GetBookmarkFaviconHash() uint32 {
	if m != nil && m.BookmarkFaviconHash != nil {
		return *m.BookmarkFaviconHash
	}
	return 0
}

func init() {
	proto.RegisterType((*EntityMetadata)(nil), "sync_pb.EntityMetadata")
}

func init() {
	proto.RegisterFile("components/sync/protocol/entity_metadata.proto", fileDescriptor_entity_metadata_a18bc036f040841a)
}

var fileDescriptor_entity_metadata_a18bc036f040841a = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0x75, 0xd0, 0xf6, 0x6c, 0x4d, 0x99, 0x2b, 0x44, 0x04, 0x42, 0x8a, 0x40, 0x83,
	0xa0, 0x89, 0x54, 0xec, 0x92, 0x2b, 0x84, 0x00, 0x8d, 0x0b, 0xd0, 0x94, 0x0d, 0x6e, 0x2d, 0xc7,
	0x39, 0x6b, 0xac, 0xd6, 0x76, 0x66, 0x3b, 0x95, 0xf6, 0x2c, 0xbc, 0x2c, 0x8a, 0x9d, 0xa0, 0xb6,
	0x12, 0x97, 0xf9, 0xbf, 0xcf, 0xe7, 0x44, 0xbf, 0x0d, 0x39, 0xd7, 0xb2, 0xd1, 0x0a, 0x95, 0xb3,
	0x4b, 0xfb, 0xa0, 0xf8, 0xb2, 0x31, 0xda, 0x69, 0xae, 0x37, 0x4b, 0x54, 0x4e, 0xb8, 0x07, 0x2a,
	0xd1, 0xb1, 0x8a, 0x39, 0x96, 0x7b, 0x40, 0xc6, 0x9d, 0x44, 0x9b, 0xf2, 0xf9, 0xff, 0x0f, 0xb6,
	0x4a, 0xdc, 0xb7, 0x48, 0x1b, 0x6d, 0x85, 0x13, 0x5a, 0x85, 0x83, 0xaf, 0xfe, 0x1c, 0x43, 0xfc,
	0xd5, 0x8f, 0xfc, 0xd1, 0x4f, 0x24, 0x6f, 0x60, 0xce, 0x37, 0x02, 0x95, 0xa3, 0x8e, 0xad, 0x68,
	0xcd, 0x6c, 0x9d, 0x44, 0x69, 0x94, 0x4d, 0x8b, 0x59, 0x88, 0x6f, 0xd9, 0xea, 0x8a, 0xd9, 0x9a,
	0xbc, 0x80, 0xa9, 0x45, 0xb3, 0x45, 0x43, 0x45, 0x95, 0x1c, 0x79, 0x63, 0x12, 0x82, 0xef, 0x15,
	0x79, 0x09, 0x20, 0x2c, 0xad, 0x70, 0x83, 0x0e, 0xab, 0x64, 0x94, 0x46, 0xd9, 0xa4, 0x98, 0x0a,
	0xfb, 0x25, 0x04, 0xe4, 0x2d, 0xcc, 0x2d, 0xde, 0xb7, 0xa8, 0x38, 0x52, 0xd5, 0xca, 0x12, 0x4d,
	0x72, 0x9c, 0x46, 0xd9, 0xa8, 0x88, 0x87, 0xf8, 0xa7, 0x4f, 0xc9, 0x25, 0x3c, 0x65, 0x7c, 0x8d,
	0x15, 0x3d, 0xd4, 0x1f, 0x79, 0x7d, 0xe1, 0xe1, 0xcd, 0xfe, 0x99, 0x77, 0x10, 0xf7, 0x3f, 0xb6,
	0x45, 0x63, 0x85, 0x56, 0xc9, 0xe3, 0x4e, 0xfe, 0x78, 0xf4, 0xfe, 0x43, 0x31, 0x0b, 0xe4, 0x77,
	0x00, 0xe4, 0x35, 0xcc, 0xb8, 0x41, 0xd6, 0x15, 0x42, 0x9d, 0x90, 0x98, 0x8c, 0xfd, 0xd8, 0xd3,
	0x21, 0xbc, 0x15, 0x12, 0xc9, 0x05, 0x9c, 0x49, 0x5d, 0x89, 0x3b, 0xc1, 0x77, 0xc4, 0x89, 0x17,
	0x9f, 0xec, 0x02, 0x2f, 0x9f, 0x43, 0x6c, 0x1b, 0xe4, 0x5d, 0x68, 0x43, 0x79, 0xd3, 0x50, 0xde,
	0xbf, 0xd4, 0x97, 0x97, 0xc3, 0xa2, 0x64, 0x16, 0xe9, 0x81, 0x0b, 0xde, 0x3d, 0xeb, 0xd0, 0xcd,
	0x9e, 0xff, 0x09, 0xe6, 0x07, 0x17, 0x98, 0x9c, 0xa4, 0x51, 0x76, 0x72, 0xf9, 0x2c, 0xef, 0xaf,
	0x3e, 0xff, 0xe5, 0xf9, 0x75, 0x8f, 0x8b, 0xb8, 0xdd, 0xfb, 0xee, 0x9a, 0x2c, 0xb5, 0x5e, 0x4b,
	0x66, 0xd6, 0xf4, 0x8e, 0x6d, 0x05, 0xd7, 0x2a, 0xec, 0x3c, 0x4d, 0xa3, 0x6c, 0x5c, 0x2c, 0x06,
	0xf8, 0x2d, 0xb0, 0x6e, 0xeb, 0xe7, 0x0b, 0x38, 0xd7, 0x66, 0x95, 0xf3, 0xda, 0x68, 0x29, 0x5a,
	0xb9, 0xf3, 0xb8, 0xfc, 0xd6, 0x7c, 0x78, 0x5c, 0x57, 0xa3, 0xeb, 0xe8, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0x0f, 0x4d, 0xf6, 0xb4, 0x02, 0x00, 0x00,
}
