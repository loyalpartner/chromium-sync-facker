// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/history_delete_directive_specifics.proto

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

// Properties of history delete directive sync objects.
type HistoryDeleteDirectiveSpecifics struct {
	// Exactly one of the fields below must be filled in.  Otherwise, this
	// delete directive must be ignored.
	GlobalIdDirective    *GlobalIdDirective  `protobuf:"bytes,1,opt,name=global_id_directive,json=globalIdDirective" json:"global_id_directive,omitempty"`
	TimeRangeDirective   *TimeRangeDirective `protobuf:"bytes,2,opt,name=time_range_directive,json=timeRangeDirective" json:"time_range_directive,omitempty"`
	UrlDirective         *UrlDirective       `protobuf:"bytes,3,opt,name=url_directive,json=urlDirective" json:"url_directive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HistoryDeleteDirectiveSpecifics) Reset()         { *m = HistoryDeleteDirectiveSpecifics{} }
func (m *HistoryDeleteDirectiveSpecifics) String() string { return proto.CompactTextString(m) }
func (*HistoryDeleteDirectiveSpecifics) ProtoMessage()    {}
func (*HistoryDeleteDirectiveSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_history_delete_directive_specifics_e483bb43413eb85b, []int{0}
}
func (m *HistoryDeleteDirectiveSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryDeleteDirectiveSpecifics.Unmarshal(m, b)
}
func (m *HistoryDeleteDirectiveSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryDeleteDirectiveSpecifics.Marshal(b, m, deterministic)
}
func (dst *HistoryDeleteDirectiveSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryDeleteDirectiveSpecifics.Merge(dst, src)
}
func (m *HistoryDeleteDirectiveSpecifics) XXX_Size() int {
	return xxx_messageInfo_HistoryDeleteDirectiveSpecifics.Size(m)
}
func (m *HistoryDeleteDirectiveSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryDeleteDirectiveSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryDeleteDirectiveSpecifics proto.InternalMessageInfo

func (m *HistoryDeleteDirectiveSpecifics) GetGlobalIdDirective() *GlobalIdDirective {
	if m != nil {
		return m.GlobalIdDirective
	}
	return nil
}

func (m *HistoryDeleteDirectiveSpecifics) GetTimeRangeDirective() *TimeRangeDirective {
	if m != nil {
		return m.TimeRangeDirective
	}
	return nil
}

func (m *HistoryDeleteDirectiveSpecifics) GetUrlDirective() *UrlDirective {
	if m != nil {
		return m.UrlDirective
	}
	return nil
}

type GlobalIdDirective struct {
	// The global IDs of the navigations to delete.
	GlobalId []int64 `protobuf:"varint,1,rep,name=global_id,json=globalId" json:"global_id,omitempty"`
	// Time range for searching for navigations to delete. Client should delete
	// all navigations to a URL between [start_time_usec, end_time_usec]
	// if one of them matches a |global_id|.
	StartTimeUsec        *int64   `protobuf:"varint,2,opt,name=start_time_usec,json=startTimeUsec" json:"start_time_usec,omitempty"`
	EndTimeUsec          *int64   `protobuf:"varint,3,opt,name=end_time_usec,json=endTimeUsec" json:"end_time_usec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalIdDirective) Reset()         { *m = GlobalIdDirective{} }
func (m *GlobalIdDirective) String() string { return proto.CompactTextString(m) }
func (*GlobalIdDirective) ProtoMessage()    {}
func (*GlobalIdDirective) Descriptor() ([]byte, []int) {
	return fileDescriptor_history_delete_directive_specifics_e483bb43413eb85b, []int{1}
}
func (m *GlobalIdDirective) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalIdDirective.Unmarshal(m, b)
}
func (m *GlobalIdDirective) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalIdDirective.Marshal(b, m, deterministic)
}
func (dst *GlobalIdDirective) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalIdDirective.Merge(dst, src)
}
func (m *GlobalIdDirective) XXX_Size() int {
	return xxx_messageInfo_GlobalIdDirective.Size(m)
}
func (m *GlobalIdDirective) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalIdDirective.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalIdDirective proto.InternalMessageInfo

func (m *GlobalIdDirective) GetGlobalId() []int64 {
	if m != nil {
		return m.GlobalId
	}
	return nil
}

func (m *GlobalIdDirective) GetStartTimeUsec() int64 {
	if m != nil && m.StartTimeUsec != nil {
		return *m.StartTimeUsec
	}
	return 0
}

func (m *GlobalIdDirective) GetEndTimeUsec() int64 {
	if m != nil && m.EndTimeUsec != nil {
		return *m.EndTimeUsec
	}
	return 0
}

type TimeRangeDirective struct {
	// The time on or after which entries must be deleted.
	StartTimeUsec *int64 `protobuf:"varint,1,opt,name=start_time_usec,json=startTimeUsec" json:"start_time_usec,omitempty"`
	// The time on or before which entries must be deleted.
	EndTimeUsec          *int64   `protobuf:"varint,2,opt,name=end_time_usec,json=endTimeUsec" json:"end_time_usec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRangeDirective) Reset()         { *m = TimeRangeDirective{} }
func (m *TimeRangeDirective) String() string { return proto.CompactTextString(m) }
func (*TimeRangeDirective) ProtoMessage()    {}
func (*TimeRangeDirective) Descriptor() ([]byte, []int) {
	return fileDescriptor_history_delete_directive_specifics_e483bb43413eb85b, []int{2}
}
func (m *TimeRangeDirective) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRangeDirective.Unmarshal(m, b)
}
func (m *TimeRangeDirective) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRangeDirective.Marshal(b, m, deterministic)
}
func (dst *TimeRangeDirective) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRangeDirective.Merge(dst, src)
}
func (m *TimeRangeDirective) XXX_Size() int {
	return xxx_messageInfo_TimeRangeDirective.Size(m)
}
func (m *TimeRangeDirective) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRangeDirective.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRangeDirective proto.InternalMessageInfo

func (m *TimeRangeDirective) GetStartTimeUsec() int64 {
	if m != nil && m.StartTimeUsec != nil {
		return *m.StartTimeUsec
	}
	return 0
}

func (m *TimeRangeDirective) GetEndTimeUsec() int64 {
	if m != nil && m.EndTimeUsec != nil {
		return *m.EndTimeUsec
	}
	return 0
}

type UrlDirective struct {
	// The URL that should be removed from history.
	Url *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// The time on or before which entries must be deleted.
	// In microseconds since the Unix epoch.
	EndTimeUsec          *int64   `protobuf:"varint,2,opt,name=end_time_usec,json=endTimeUsec" json:"end_time_usec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UrlDirective) Reset()         { *m = UrlDirective{} }
func (m *UrlDirective) String() string { return proto.CompactTextString(m) }
func (*UrlDirective) ProtoMessage()    {}
func (*UrlDirective) Descriptor() ([]byte, []int) {
	return fileDescriptor_history_delete_directive_specifics_e483bb43413eb85b, []int{3}
}
func (m *UrlDirective) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlDirective.Unmarshal(m, b)
}
func (m *UrlDirective) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlDirective.Marshal(b, m, deterministic)
}
func (dst *UrlDirective) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlDirective.Merge(dst, src)
}
func (m *UrlDirective) XXX_Size() int {
	return xxx_messageInfo_UrlDirective.Size(m)
}
func (m *UrlDirective) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlDirective.DiscardUnknown(m)
}

var xxx_messageInfo_UrlDirective proto.InternalMessageInfo

func (m *UrlDirective) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *UrlDirective) GetEndTimeUsec() int64 {
	if m != nil && m.EndTimeUsec != nil {
		return *m.EndTimeUsec
	}
	return 0
}

func init() {
	proto.RegisterType((*HistoryDeleteDirectiveSpecifics)(nil), "sync_pb.HistoryDeleteDirectiveSpecifics")
	proto.RegisterType((*GlobalIdDirective)(nil), "sync_pb.GlobalIdDirective")
	proto.RegisterType((*TimeRangeDirective)(nil), "sync_pb.TimeRangeDirective")
	proto.RegisterType((*UrlDirective)(nil), "sync_pb.UrlDirective")
}

func init() {
	proto.RegisterFile("components/sync/protocol/history_delete_directive_specifics.proto", fileDescriptor_history_delete_directive_specifics_e483bb43413eb85b)
}

var fileDescriptor_history_delete_directive_specifics_e483bb43413eb85b = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x49, 0xf7, 0xf0, 0xfd, 0x76, 0xda, 0xa2, 0x5d, 0x15, 0x8a, 0x3d, 0x58, 0x02, 0x4a,
	0x41, 0x48, 0xc1, 0xa3, 0x37, 0xa5, 0x60, 0x15, 0x04, 0x89, 0xf6, 0xbc, 0xa6, 0x9b, 0x31, 0x5d,
	0xd8, 0x64, 0xc3, 0xee, 0x46, 0x28, 0xf8, 0x87, 0x7b, 0x94, 0xac, 0xe6, 0x07, 0xc6, 0x43, 0x6f,
	0xc9, 0xf0, 0xde, 0x67, 0xde, 0x5b, 0x06, 0x6e, 0xb8, 0x4a, 0x73, 0x95, 0x61, 0x66, 0xcd, 0xc2,
	0xec, 0x32, 0xbe, 0xc8, 0xb5, 0xb2, 0x8a, 0x2b, 0xb9, 0xd8, 0x0a, 0x63, 0x95, 0xde, 0xb1, 0x18,
	0x25, 0x5a, 0x64, 0xb1, 0xd0, 0xc8, 0xad, 0x78, 0x47, 0x66, 0x72, 0xe4, 0xe2, 0x4d, 0x70, 0x13,
	0x38, 0x2d, 0xfd, 0x57, 0xfa, 0x58, 0xbe, 0xf1, 0x3f, 0x3d, 0x38, 0x5b, 0x7d, 0xbb, 0x96, 0xce,
	0xb4, 0xac, 0x3c, 0xcf, 0x95, 0x85, 0x3e, 0xc0, 0x51, 0x22, 0xd5, 0x26, 0x92, 0x4c, 0xc4, 0x0d,
	0x73, 0xe2, 0xcd, 0xbc, 0xf9, 0xe0, 0xea, 0x34, 0xf8, 0x41, 0x05, 0x77, 0x4e, 0x73, 0x1f, 0xd7,
	0x84, 0x70, 0x9c, 0xfc, 0x1e, 0xd1, 0x47, 0x38, 0xb6, 0x22, 0x45, 0xa6, 0xa3, 0x2c, 0x69, 0x05,
	0x9c, 0xf4, 0x1c, 0x6c, 0x5a, 0xc3, 0x5e, 0x44, 0x8a, 0x61, 0xa9, 0x69, 0x68, 0xd4, 0x76, 0x66,
	0xf4, 0x1a, 0x46, 0x85, 0x96, 0x2d, 0x0e, 0x71, 0x9c, 0x93, 0x9a, 0xb3, 0xd6, 0xb2, 0x21, 0x0c,
	0x8b, 0xd6, 0x9f, 0xff, 0x01, 0xe3, 0x4e, 0x64, 0x3a, 0x85, 0x7e, 0xdd, 0x75, 0xe2, 0xcd, 0xc8,
	0x9c, 0x84, 0xff, 0xab, 0x16, 0xf4, 0x02, 0x0e, 0x8c, 0x8d, 0xb4, 0x65, 0xae, 0x42, 0x61, 0x90,
	0xbb, 0xdc, 0x24, 0x1c, 0xb9, 0x71, 0x99, 0x79, 0x6d, 0x90, 0x53, 0x1f, 0x46, 0x98, 0xc5, 0x2d,
	0x15, 0x71, 0xaa, 0x01, 0x66, 0x71, 0xa5, 0xf1, 0x5f, 0x81, 0x76, 0x3b, 0xfe, 0xb5, 0xc1, 0xdb,
	0x6b, 0x43, 0xaf, 0xbb, 0x61, 0x09, 0xc3, 0x76, 0x7b, 0x7a, 0x08, 0xa4, 0xd0, 0xd2, 0xf1, 0xfa,
	0x61, 0xf9, 0xb9, 0x0f, 0xe5, 0xf6, 0x12, 0xce, 0x95, 0x4e, 0x02, 0xbe, 0xd5, 0x2a, 0x15, 0x45,
	0x1a, 0x34, 0xb7, 0xe7, 0xde, 0x38, 0xa8, 0x6e, 0x6f, 0x45, 0x9e, 0xbc, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x49, 0x97, 0x89, 0x4b, 0x9a, 0x02, 0x00, 0x00,
}
