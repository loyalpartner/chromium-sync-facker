// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/client_debug_info.proto

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

// Per-type hint information.
type TypeHint struct {
	// The data type this hint applied to.
	DataTypeId *int32 `protobuf:"varint,1,opt,name=data_type_id,json=dataTypeId" json:"data_type_id,omitempty"`
	// Whether or not a valid hint is provided.
	HasValidHint         *bool    `protobuf:"varint,2,opt,name=has_valid_hint,json=hasValidHint" json:"has_valid_hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeHint) Reset()         { *m = TypeHint{} }
func (m *TypeHint) String() string { return proto.CompactTextString(m) }
func (*TypeHint) ProtoMessage()    {}
func (*TypeHint) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_debug_info_1cd2ab3c25597fab, []int{0}
}
func (m *TypeHint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeHint.Unmarshal(m, b)
}
func (m *TypeHint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeHint.Marshal(b, m, deterministic)
}
func (dst *TypeHint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeHint.Merge(dst, src)
}
func (m *TypeHint) XXX_Size() int {
	return xxx_messageInfo_TypeHint.Size(m)
}
func (m *TypeHint) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeHint.DiscardUnknown(m)
}

var xxx_messageInfo_TypeHint proto.InternalMessageInfo

func (m *TypeHint) GetDataTypeId() int32 {
	if m != nil && m.DataTypeId != nil {
		return *m.DataTypeId
	}
	return 0
}

func (m *TypeHint) GetHasValidHint() bool {
	if m != nil && m.HasValidHint != nil {
		return *m.HasValidHint
	}
	return false
}

// The additional info here is from the StatusController. They get sent when
// the event SYNC_CYCLE_COMPLETED  is sent.
type SyncCycleCompletedEventInfo struct {
	// These new conflict counters replace the ones above.
	NumEncryptionConflicts *int32 `protobuf:"varint,4,opt,name=num_encryption_conflicts,json=numEncryptionConflicts" json:"num_encryption_conflicts,omitempty"`
	NumHierarchyConflicts  *int32 `protobuf:"varint,5,opt,name=num_hierarchy_conflicts,json=numHierarchyConflicts" json:"num_hierarchy_conflicts,omitempty"`
	NumSimpleConflicts     *int32 `protobuf:"varint,6,opt,name=num_simple_conflicts,json=numSimpleConflicts" json:"num_simple_conflicts,omitempty"`
	NumServerConflicts     *int32 `protobuf:"varint,7,opt,name=num_server_conflicts,json=numServerConflicts" json:"num_server_conflicts,omitempty"`
	// Counts to track the effective usefulness of our GetUpdate requests.
	NumUpdatesDownloaded          *int32 `protobuf:"varint,8,opt,name=num_updates_downloaded,json=numUpdatesDownloaded" json:"num_updates_downloaded,omitempty"`
	NumReflectedUpdatesDownloaded *int32 `protobuf:"varint,9,opt,name=num_reflected_updates_downloaded,json=numReflectedUpdatesDownloaded" json:"num_reflected_updates_downloaded,omitempty"`
	// |caller_info| was mostly replaced by |get_updates_origin|; now it only
	// contains the |notifications_enabled| flag.
	CallerInfo           *GetUpdatesCallerInfo       `protobuf:"bytes,10,opt,name=caller_info,json=callerInfo" json:"caller_info,omitempty"`
	GetUpdatesOrigin     *SyncEnums_GetUpdatesOrigin `protobuf:"varint,12,opt,name=get_updates_origin,json=getUpdatesOrigin,enum=sync_pb.SyncEnums_GetUpdatesOrigin" json:"get_updates_origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SyncCycleCompletedEventInfo) Reset()         { *m = SyncCycleCompletedEventInfo{} }
func (m *SyncCycleCompletedEventInfo) String() string { return proto.CompactTextString(m) }
func (*SyncCycleCompletedEventInfo) ProtoMessage()    {}
func (*SyncCycleCompletedEventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_debug_info_1cd2ab3c25597fab, []int{1}
}
func (m *SyncCycleCompletedEventInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncCycleCompletedEventInfo.Unmarshal(m, b)
}
func (m *SyncCycleCompletedEventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncCycleCompletedEventInfo.Marshal(b, m, deterministic)
}
func (dst *SyncCycleCompletedEventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncCycleCompletedEventInfo.Merge(dst, src)
}
func (m *SyncCycleCompletedEventInfo) XXX_Size() int {
	return xxx_messageInfo_SyncCycleCompletedEventInfo.Size(m)
}
func (m *SyncCycleCompletedEventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncCycleCompletedEventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SyncCycleCompletedEventInfo proto.InternalMessageInfo

func (m *SyncCycleCompletedEventInfo) GetNumEncryptionConflicts() int32 {
	if m != nil && m.NumEncryptionConflicts != nil {
		return *m.NumEncryptionConflicts
	}
	return 0
}

func (m *SyncCycleCompletedEventInfo) GetNumHierarchyConflicts() int32 {
	if m != nil && m.NumHierarchyConflicts != nil {
		return *m.NumHierarchyConflicts
	}
	return 0
}

func (m *SyncCycleCompletedEventInfo) GetNumSimpleConflicts() int32 {
	if m != nil && m.NumSimpleConflicts != nil {
		return *m.NumSimpleConflicts
	}
	return 0
}

func (m *SyncCycleCompletedEventInfo) GetNumServerConflicts() int32 {
	if m != nil && m.NumServerConflicts != nil {
		return *m.NumServerConflicts
	}
	return 0
}

func (m *SyncCycleCompletedEventInfo) GetNumUpdatesDownloaded() int32 {
	if m != nil && m.NumUpdatesDownloaded != nil {
		return *m.NumUpdatesDownloaded
	}
	return 0
}

func (m *SyncCycleCompletedEventInfo) GetNumReflectedUpdatesDownloaded() int32 {
	if m != nil && m.NumReflectedUpdatesDownloaded != nil {
		return *m.NumReflectedUpdatesDownloaded
	}
	return 0
}

func (m *SyncCycleCompletedEventInfo) GetCallerInfo() *GetUpdatesCallerInfo {
	if m != nil {
		return m.CallerInfo
	}
	return nil
}

func (m *SyncCycleCompletedEventInfo) GetGetUpdatesOrigin() SyncEnums_GetUpdatesOrigin {
	if m != nil && m.GetUpdatesOrigin != nil {
		return *m.GetUpdatesOrigin
	}
	return SyncEnums_UNKNOWN_ORIGIN
}

// Datatype specifics statistics gathered at association time.
type DatatypeAssociationStats struct {
	// The datatype that was associated.
	DataTypeId *int32 `protobuf:"varint,1,opt,name=data_type_id,json=dataTypeId" json:"data_type_id,omitempty"`
	// Waiting time before downloading starts. This measures the time between
	// receiving configuration request for a set of data types to starting
	// downloading data of this type.
	DownloadWaitTimeUs *int64 `protobuf:"varint,15,opt,name=download_wait_time_us,json=downloadWaitTimeUs" json:"download_wait_time_us,omitempty"`
	// Time spent on downloading sync data for first time sync.
	// Note: This measures the time between asking backend to download data to
	//       being notified of download-ready by backend. So it consists of
	//       time on data downloading and processing at sync backend. But
	//       downloading time should dominate. It's also the total time spent on
	//       downloading data of all types in the priority group of
	//       |data_type_id| instead of just one data type.
	DownloadTimeUs *int64 `protobuf:"varint,13,opt,name=download_time_us,json=downloadTimeUs" json:"download_time_us,omitempty"`
	// Higher priority type that's configured before this type.
	HighPriorityTypeConfiguredBefore []int32 `protobuf:"varint,18,rep,name=high_priority_type_configured_before,json=highPriorityTypeConfiguredBefore" json:"high_priority_type_configured_before,omitempty"`
	// Same priority type that's configured before this type.
	SamePriorityTypeConfiguredBefore []int32  `protobuf:"varint,19,rep,name=same_priority_type_configured_before,json=samePriorityTypeConfiguredBefore" json:"same_priority_type_configured_before,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *DatatypeAssociationStats) Reset()         { *m = DatatypeAssociationStats{} }
func (m *DatatypeAssociationStats) String() string { return proto.CompactTextString(m) }
func (*DatatypeAssociationStats) ProtoMessage()    {}
func (*DatatypeAssociationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_debug_info_1cd2ab3c25597fab, []int{2}
}
func (m *DatatypeAssociationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatatypeAssociationStats.Unmarshal(m, b)
}
func (m *DatatypeAssociationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatatypeAssociationStats.Marshal(b, m, deterministic)
}
func (dst *DatatypeAssociationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatatypeAssociationStats.Merge(dst, src)
}
func (m *DatatypeAssociationStats) XXX_Size() int {
	return xxx_messageInfo_DatatypeAssociationStats.Size(m)
}
func (m *DatatypeAssociationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_DatatypeAssociationStats.DiscardUnknown(m)
}

var xxx_messageInfo_DatatypeAssociationStats proto.InternalMessageInfo

func (m *DatatypeAssociationStats) GetDataTypeId() int32 {
	if m != nil && m.DataTypeId != nil {
		return *m.DataTypeId
	}
	return 0
}

func (m *DatatypeAssociationStats) GetDownloadWaitTimeUs() int64 {
	if m != nil && m.DownloadWaitTimeUs != nil {
		return *m.DownloadWaitTimeUs
	}
	return 0
}

func (m *DatatypeAssociationStats) GetDownloadTimeUs() int64 {
	if m != nil && m.DownloadTimeUs != nil {
		return *m.DownloadTimeUs
	}
	return 0
}

func (m *DatatypeAssociationStats) GetHighPriorityTypeConfiguredBefore() []int32 {
	if m != nil {
		return m.HighPriorityTypeConfiguredBefore
	}
	return nil
}

func (m *DatatypeAssociationStats) GetSamePriorityTypeConfiguredBefore() []int32 {
	if m != nil {
		return m.SamePriorityTypeConfiguredBefore
	}
	return nil
}

type DebugEventInfo struct {
	// Each of the following fields correspond to different kinds of events. as
	// a result, only one is set during any single DebugEventInfo.
	// A singleton event. See enum definition.
	SingletonEvent *SyncEnums_SingletonDebugEventType `protobuf:"varint,1,opt,name=singleton_event,json=singletonEvent,enum=sync_pb.SyncEnums_SingletonDebugEventType" json:"singleton_event,omitempty"`
	// A sync cycle completed.
	SyncCycleCompletedEventInfo *SyncCycleCompletedEventInfo `protobuf:"bytes,2,opt,name=sync_cycle_completed_event_info,json=syncCycleCompletedEventInfo" json:"sync_cycle_completed_event_info,omitempty"`
	// A datatype triggered a nudge.
	NudgingDatatype *int32 `protobuf:"varint,3,opt,name=nudging_datatype,json=nudgingDatatype" json:"nudging_datatype,omitempty"`
	// A notification triggered a nudge.
	DatatypesNotifiedFromServer []int32 `protobuf:"varint,4,rep,name=datatypes_notified_from_server,json=datatypesNotifiedFromServer" json:"datatypes_notified_from_server,omitempty"`
	// A datatype finished model association.
	DatatypeAssociationStats *DatatypeAssociationStats `protobuf:"bytes,5,opt,name=datatype_association_stats,json=datatypeAssociationStats" json:"datatype_association_stats,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *DebugEventInfo) Reset()         { *m = DebugEventInfo{} }
func (m *DebugEventInfo) String() string { return proto.CompactTextString(m) }
func (*DebugEventInfo) ProtoMessage()    {}
func (*DebugEventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_debug_info_1cd2ab3c25597fab, []int{3}
}
func (m *DebugEventInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugEventInfo.Unmarshal(m, b)
}
func (m *DebugEventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugEventInfo.Marshal(b, m, deterministic)
}
func (dst *DebugEventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugEventInfo.Merge(dst, src)
}
func (m *DebugEventInfo) XXX_Size() int {
	return xxx_messageInfo_DebugEventInfo.Size(m)
}
func (m *DebugEventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugEventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DebugEventInfo proto.InternalMessageInfo

func (m *DebugEventInfo) GetSingletonEvent() SyncEnums_SingletonDebugEventType {
	if m != nil && m.SingletonEvent != nil {
		return *m.SingletonEvent
	}
	return SyncEnums_CONNECTION_STATUS_CHANGE
}

func (m *DebugEventInfo) GetSyncCycleCompletedEventInfo() *SyncCycleCompletedEventInfo {
	if m != nil {
		return m.SyncCycleCompletedEventInfo
	}
	return nil
}

func (m *DebugEventInfo) GetNudgingDatatype() int32 {
	if m != nil && m.NudgingDatatype != nil {
		return *m.NudgingDatatype
	}
	return 0
}

func (m *DebugEventInfo) GetDatatypesNotifiedFromServer() []int32 {
	if m != nil {
		return m.DatatypesNotifiedFromServer
	}
	return nil
}

func (m *DebugEventInfo) GetDatatypeAssociationStats() *DatatypeAssociationStats {
	if m != nil {
		return m.DatatypeAssociationStats
	}
	return nil
}

type DebugInfo struct {
	Events []*DebugEventInfo `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	// Whether cryptographer is ready to encrypt and decrypt data.
	CryptographerReady *bool `protobuf:"varint,2,opt,name=cryptographer_ready,json=cryptographerReady" json:"cryptographer_ready,omitempty"`
	// Cryptographer has pending keys which indicates the correct passphrase
	// has not been provided yet.
	CryptographerHasPendingKeys *bool `protobuf:"varint,3,opt,name=cryptographer_has_pending_keys,json=cryptographerHasPendingKeys" json:"cryptographer_has_pending_keys,omitempty"`
	// Indicates client has dropped some events to save bandwidth.
	EventsDropped        *bool    `protobuf:"varint,4,opt,name=events_dropped,json=eventsDropped" json:"events_dropped,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugInfo) Reset()         { *m = DebugInfo{} }
func (m *DebugInfo) String() string { return proto.CompactTextString(m) }
func (*DebugInfo) ProtoMessage()    {}
func (*DebugInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_debug_info_1cd2ab3c25597fab, []int{4}
}
func (m *DebugInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugInfo.Unmarshal(m, b)
}
func (m *DebugInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugInfo.Marshal(b, m, deterministic)
}
func (dst *DebugInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugInfo.Merge(dst, src)
}
func (m *DebugInfo) XXX_Size() int {
	return xxx_messageInfo_DebugInfo.Size(m)
}
func (m *DebugInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DebugInfo proto.InternalMessageInfo

func (m *DebugInfo) GetEvents() []*DebugEventInfo {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *DebugInfo) GetCryptographerReady() bool {
	if m != nil && m.CryptographerReady != nil {
		return *m.CryptographerReady
	}
	return false
}

func (m *DebugInfo) GetCryptographerHasPendingKeys() bool {
	if m != nil && m.CryptographerHasPendingKeys != nil {
		return *m.CryptographerHasPendingKeys
	}
	return false
}

func (m *DebugInfo) GetEventsDropped() bool {
	if m != nil && m.EventsDropped != nil {
		return *m.EventsDropped
	}
	return false
}

func init() {
	proto.RegisterType((*TypeHint)(nil), "sync_pb.TypeHint")
	proto.RegisterType((*SyncCycleCompletedEventInfo)(nil), "sync_pb.SyncCycleCompletedEventInfo")
	proto.RegisterType((*DatatypeAssociationStats)(nil), "sync_pb.DatatypeAssociationStats")
	proto.RegisterType((*DebugEventInfo)(nil), "sync_pb.DebugEventInfo")
	proto.RegisterType((*DebugInfo)(nil), "sync_pb.DebugInfo")
}

func init() {
	proto.RegisterFile("components/sync/protocol/client_debug_info.proto", fileDescriptor_client_debug_info_1cd2ab3c25597fab)
}

var fileDescriptor_client_debug_info_1cd2ab3c25597fab = []byte{
	// 1069 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xef, 0x6e, 0xdb, 0x36,
	0x10, 0x87, 0x63, 0x25, 0x61, 0x68, 0xc7, 0x51, 0x95, 0x7f, 0x42, 0xb2, 0x6c, 0xae, 0x97, 0x02,
	0x6e, 0x07, 0x38, 0x5d, 0x30, 0x14, 0xfb, 0x34, 0x60, 0x75, 0xb2, 0xa6, 0x1a, 0xd0, 0x65, 0x4c,
	0xbb, 0x7d, 0x24, 0x18, 0x91, 0xb6, 0xb9, 0x48, 0xa4, 0x40, 0x52, 0x29, 0xfc, 0x3a, 0x7b, 0xab,
	0x3d, 0xc3, 0xf6, 0x10, 0x03, 0x29, 0xc9, 0x96, 0x9d, 0x38, 0xdb, 0x37, 0xe3, 0x7e, 0x7f, 0x44,
	0xde, 0x9d, 0xef, 0x08, 0x5f, 0xc7, 0x32, 0xcd, 0xa4, 0x60, 0xc2, 0xe8, 0x33, 0x3d, 0x15, 0xf1,
	0x59, 0xa6, 0xa4, 0x91, 0xb1, 0x4c, 0xce, 0xe2, 0x84, 0x33, 0x61, 0x30, 0x65, 0xb7, 0xf9, 0x18,
	0x73, 0x31, 0x92, 0x03, 0x07, 0x05, 0x9b, 0x96, 0x86, 0xb3, 0xdb, 0xa3, 0x37, 0x2b, 0xa5, 0x63,
	0x66, 0x70, 0x9e, 0x51, 0x62, 0x98, 0xc6, 0x31, 0x49, 0x12, 0xa6, 0x6a, 0x06, 0x47, 0x2f, 0x57,
	0xea, 0x9c, 0x33, 0x13, 0x79, 0xaa, 0x0b, 0x6a, 0x0f, 0x41, 0xf0, 0x71, 0x9a, 0xb1, 0x2b, 0x2e,
	0x4c, 0xd0, 0x85, 0x6d, 0x4a, 0x0c, 0xc1, 0x66, 0x9a, 0x31, 0xcc, 0x69, 0xd8, 0xe8, 0x36, 0xfa,
	0xeb, 0x08, 0xda, 0x98, 0xe5, 0xbc, 0xa7, 0xc1, 0x29, 0xec, 0x4c, 0x88, 0xc6, 0xf7, 0x24, 0xe1,
	0x14, 0x4f, 0xb8, 0x30, 0xe1, 0x5a, 0xb7, 0xd1, 0x07, 0xa8, 0x3d, 0x21, 0xfa, 0x37, 0x1b, 0xb4,
	0x3e, 0xbd, 0x7f, 0x3c, 0x78, 0x7c, 0x33, 0x15, 0xf1, 0x70, 0x1a, 0x27, 0x6c, 0x28, 0xd3, 0x2c,
	0x61, 0x86, 0xd1, 0xcb, 0x7b, 0x26, 0xcc, 0x7b, 0x31, 0x92, 0xc1, 0xf7, 0x30, 0x14, 0x79, 0x8a,
	0x99, 0x88, 0xd5, 0x34, 0x33, 0x5c, 0x0a, 0x1c, 0x4b, 0x31, 0x4a, 0x78, 0x6c, 0x74, 0xe8, 0xb9,
	0x6f, 0x1e, 0x88, 0x3c, 0xbd, 0x9c, 0xc1, 0xc3, 0x0a, 0x0d, 0xde, 0xc0, 0x43, 0xab, 0x9c, 0x70,
	0xa6, 0x88, 0x8a, 0x27, 0xd3, 0x9a, 0x70, 0xdd, 0x09, 0xf7, 0x45, 0x9e, 0x5e, 0x55, 0xe8, 0x5c,
	0xf7, 0x1a, 0xee, 0x59, 0x9d, 0xe6, 0xf6, 0x2c, 0x35, 0xd1, 0x86, 0x13, 0x05, 0x22, 0x4f, 0x6f,
	0x1c, 0xf4, 0x50, 0xc1, 0xd4, 0x3d, 0x53, 0x35, 0xc5, 0xe6, 0x5c, 0xe1, 0xa0, 0xb9, 0xe2, 0x3b,
	0x68, 0x4f, 0x3d, 0xab, 0x0a, 0x95, 0x9f, 0x45, 0x22, 0x09, 0x65, 0x34, 0x04, 0x4e, 0x63, 0xfd,
	0x3e, 0x15, 0xe0, 0xc5, 0x0c, 0x0b, 0xde, 0xc1, 0xae, 0x55, 0x29, 0x36, 0x4a, 0x58, 0x6c, 0x18,
	0x7d, 0x4c, 0xbf, 0xe5, 0xf4, 0x27, 0x22, 0x4f, 0x51, 0x45, 0x7b, 0x68, 0xf4, 0x03, 0x6c, 0xd5,
	0x1a, 0x21, 0x84, 0xdd, 0x46, 0xbf, 0x75, 0x7e, 0x32, 0x28, 0x5b, 0x69, 0xf0, 0x8e, 0x99, 0x52,
	0x33, 0x74, 0x2c, 0x5b, 0x08, 0x04, 0xe3, 0xd9, 0xef, 0xe0, 0x57, 0x18, 0xd4, 0x9b, 0x4a, 0x2a,
	0x3e, 0xe6, 0x22, 0x6c, 0x77, 0x1b, 0xfd, 0xce, 0xf9, 0xd7, 0x33, 0x1b, 0x5b, 0xd6, 0x4b, 0xd7,
	0x3e, 0x73, 0xc3, 0x5f, 0x1c, 0x15, 0xf9, 0xe3, 0xa5, 0x48, 0xe4, 0x81, 0x86, 0xbf, 0x16, 0x79,
	0x60, 0xcd, 0x6f, 0x46, 0x1e, 0x68, 0xfa, 0x5e, 0xe4, 0x81, 0x96, 0xdf, 0x46, 0x6d, 0xeb, 0xc6,
	0x14, 0xd6, 0x26, 0x8f, 0xef, 0x5c, 0xad, 0xf1, 0x6d, 0x22, 0xe3, 0x3b, 0x2e, 0xc6, 0xf3, 0x4c,
	0xa3, 0x23, 0x1b, 0x17, 0x52, 0x3c, 0x86, 0xb5, 0xb4, 0xcc, 0x55, 0xcc, 0xdc, 0x55, 0x7b, 0x7f,
	0x6f, 0xc0, 0xf0, 0x82, 0x18, 0x62, 0xdb, 0xf6, 0x47, 0xad, 0x65, 0xcc, 0x89, 0xed, 0x9a, 0x1b,
	0x43, 0x8c, 0xfe, 0x1f, 0x3d, 0xfd, 0x2d, 0xdc, 0xaf, 0x72, 0x8d, 0x3f, 0x13, 0x6e, 0xb0, 0xe1,
	0x29, 0xc3, 0xb9, 0x0e, 0x77, 0xba, 0x8d, 0x7e, 0x13, 0x05, 0x15, 0xf8, 0x3b, 0xe1, 0xe6, 0x23,
	0x4f, 0xd9, 0x27, 0x1d, 0xf4, 0xa1, 0x3f, 0x93, 0x54, 0xec, 0x6d, 0xc7, 0xee, 0x54, 0xf1, 0x92,
	0xf9, 0x01, 0x9e, 0x4e, 0xf8, 0x78, 0x82, 0x33, 0xc5, 0xa5, 0xe2, 0x66, 0x5a, 0x9c, 0xc3, 0x5e,
	0x84, 0x8f, 0x73, 0xc5, 0x28, 0xbe, 0x65, 0x23, 0xa9, 0x58, 0x18, 0x74, 0x9b, 0xfd, 0x75, 0xd4,
	0xb5, 0xdc, 0xeb, 0x92, 0x6a, 0x8f, 0x37, 0x9c, 0x11, 0xdf, 0x3a, 0x9e, 0xf5, 0xd3, 0x24, 0x65,
	0xff, 0xe9, 0xb7, 0x5b, 0xf8, 0x59, 0xee, 0x53, 0x7e, 0x4b, 0xc5, 0xf1, 0xfc, 0xf5, 0xc8, 0x03,
	0xeb, 0xfe, 0x46, 0xe4, 0x81, 0x0d, 0x7f, 0x33, 0xf2, 0xc0, 0xa6, 0x0f, 0x22, 0x0f, 0x00, 0x7f,
	0x2b, 0xf2, 0xc0, 0x96, 0x0f, 0x23, 0x0f, 0x40, 0xbf, 0x55, 0x94, 0x31, 0xf2, 0x40, 0xdb, 0xdf,
	0x8e, 0x3c, 0xd0, 0xf1, 0x77, 0x22, 0x0f, 0xf8, 0xfe, 0xb3, 0xc8, 0x03, 0xcf, 0xfc, 0x20, 0xf2,
	0xc0, 0x9e, 0xbf, 0x1f, 0x79, 0x60, 0xdf, 0x3f, 0x40, 0x3d, 0x5b, 0xc4, 0x44, 0xc6, 0x24, 0xc1,
	0xdc, 0xb0, 0x54, 0x97, 0xe7, 0xc3, 0x64, 0x5e, 0x27, 0xf4, 0xdc, 0xfd, 0xd1, 0x6c, 0x83, 0x3d,
	0x4d, 0xa9, 0xdb, 0x90, 0x91, 0x61, 0x6a, 0x81, 0xd2, 0x5d, 0x72, 0x79, 0xc8, 0xd8, 0x7f, 0x60,
	0x42, 0x29, 0xa3, 0xe8, 0x70, 0x39, 0x4c, 0x99, 0x9b, 0x57, 0x28, 0x5c, 0x06, 0x52, 0x49, 0xf9,
	0x88, 0x33, 0x8a, 0xf6, 0x96, 0xbf, 0xe5, 0x8c, 0x0e, 0x96, 0xa2, 0x95, 0xcf, 0xe1, 0x52, 0x7c,
	0x66, 0xb3, 0x35, 0x21, 0x14, 0x33, 0xa5, 0xa4, 0x42, 0xaf, 0x6a, 0x07, 0xad, 0x75, 0xe1, 0x48,
	0x2a, 0xbc, 0x58, 0xef, 0x5c, 0x3f, 0xc5, 0x5d, 0xec, 0xb5, 0x5c, 0xa3, 0xdd, 0x3a, 0xb7, 0x6c,
	0x55, 0x74, 0x52, 0x5c, 0xea, 0x9e, 0x29, 0x6d, 0xc3, 0xd9, 0x52, 0xb2, 0xbf, 0x70, 0x67, 0x5d,
	0x81, 0xf6, 0xfe, 0x6c, 0xc2, 0xce, 0x85, 0xdd, 0x58, 0xf3, 0x79, 0x7e, 0x03, 0x77, 0x34, 0x17,
	0xe3, 0x84, 0x19, 0x29, 0x30, 0xb3, 0x61, 0xf7, 0x37, 0xeb, 0x9c, 0xbf, 0x7a, 0x64, 0x6e, 0xdc,
	0x54, 0xcc, 0xb9, 0x89, 0xed, 0x4c, 0xd4, 0x99, 0x59, 0xb8, 0x58, 0xf0, 0x07, 0xfc, 0xca, 0x89,
	0x63, 0xbb, 0x44, 0x70, 0x5c, 0x6d, 0x91, 0xc2, 0xbf, 0x98, 0x71, 0x6b, 0x6e, 0xc6, 0x9d, 0x2e,
	0x7c, 0x64, 0xc5, 0xce, 0x41, 0xc7, 0xfa, 0x89, 0x85, 0xf4, 0x12, 0xfa, 0x22, 0xa7, 0x63, 0x3b,
	0x63, 0x68, 0x39, 0x48, 0xc2, 0xa6, 0x1b, 0x14, 0x3b, 0x65, 0xbc, 0x9a, 0x2f, 0xc1, 0x10, 0x7e,
	0x59, 0x51, 0x34, 0x16, 0xd2, 0xb8, 0x4a, 0xe2, 0x91, 0x92, 0xd5, 0xae, 0x08, 0x3d, 0xf7, 0xd7,
	0x3b, 0x9e, 0xb1, 0x3e, 0x94, 0xa4, 0x9f, 0x94, 0x2c, 0x77, 0x46, 0x80, 0xe1, 0x51, 0x05, 0xd7,
	0x73, 0x8b, 0xb5, 0x1d, 0x59, 0x6e, 0x93, 0xb5, 0xce, 0x9f, 0xcf, 0xae, 0xb5, 0x6a, 0xb6, 0xa1,
	0x90, 0xae, 0x40, 0x7a, 0x7f, 0x35, 0xe0, 0x96, 0xcb, 0xaf, 0xbb, 0xde, 0x19, 0xdc, 0x70, 0x59,
	0xd3, 0x61, 0xa3, 0xdb, 0xec, 0xb7, 0xce, 0x0f, 0xe7, 0xd6, 0x0b, 0x85, 0x44, 0x25, 0x2d, 0x38,
	0x83, 0xbb, 0x6e, 0xf7, 0xca, 0xb1, 0x22, 0xd9, 0x84, 0x29, 0xac, 0x18, 0xa1, 0xd3, 0x72, 0xd7,
	0x07, 0x0b, 0x10, 0xb2, 0x88, 0xcd, 0xca, 0xa2, 0xc0, 0xbe, 0x12, 0x32, 0x26, 0xa8, 0x4d, 0xe9,
	0x1d, 0x9b, 0x6a, 0x97, 0x4e, 0x80, 0x8e, 0x17, 0x58, 0x57, 0x44, 0x5f, 0x17, 0x9c, 0x9f, 0xd9,
	0x54, 0x07, 0x2f, 0x60, 0xa7, 0xf8, 0x3e, 0xa6, 0x4a, 0x66, 0x19, 0xa3, 0xee, 0x31, 0x00, 0xd0,
	0x76, 0x11, 0xbd, 0x28, 0x82, 0x6f, 0xbf, 0x81, 0x2f, 0xa4, 0x1a, 0x0f, 0xe2, 0x89, 0x92, 0x29,
	0xcf, 0xd3, 0xc1, 0xfc, 0xad, 0xe3, 0xae, 0x35, 0xa8, 0xde, 0x3a, 0x57, 0xcd, 0xeb, 0xc6, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x54, 0x74, 0xf7, 0x46, 0x7d, 0x09, 0x00, 0x00,
}
