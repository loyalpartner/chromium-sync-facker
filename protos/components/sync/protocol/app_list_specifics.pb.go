// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/app_list_specifics.proto

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

// What type of item this is.
type AppListSpecifics_AppListItemType int32

const (
	// An app, whether a web app, Android app, etc.
	//
	// For bookmark apps (URL shortcuts), additional information such as their
	// URLs are kept in the AppSpecifics.bookmark_app_foobar fields.
	AppListSpecifics_TYPE_APP AppListSpecifics_AppListItemType = 1
	// A request to remove any matching default installed apps.
	AppListSpecifics_TYPE_REMOVE_DEFAULT_APP AppListSpecifics_AppListItemType = 2
	// A folder containing entries whose |parent_id| matches |item_id|.
	AppListSpecifics_TYPE_FOLDER AppListSpecifics_AppListItemType = 3
	// Obsolete type, intended for URL shortcuts, that was never implemented.
	AppListSpecifics_TYPE_OBSOLETE_URL AppListSpecifics_AppListItemType = 4
	// A "page break" item (Indicate creation of a new page in app list).
	AppListSpecifics_TYPE_PAGE_BREAK AppListSpecifics_AppListItemType = 5
)

var AppListSpecifics_AppListItemType_name = map[int32]string{
	1: "TYPE_APP",
	2: "TYPE_REMOVE_DEFAULT_APP",
	3: "TYPE_FOLDER",
	4: "TYPE_OBSOLETE_URL",
	5: "TYPE_PAGE_BREAK",
}
var AppListSpecifics_AppListItemType_value = map[string]int32{
	"TYPE_APP":                1,
	"TYPE_REMOVE_DEFAULT_APP": 2,
	"TYPE_FOLDER":             3,
	"TYPE_OBSOLETE_URL":       4,
	"TYPE_PAGE_BREAK":         5,
}

func (x AppListSpecifics_AppListItemType) Enum() *AppListSpecifics_AppListItemType {
	p := new(AppListSpecifics_AppListItemType)
	*p = x
	return p
}
func (x AppListSpecifics_AppListItemType) String() string {
	return proto.EnumName(AppListSpecifics_AppListItemType_name, int32(x))
}
func (x *AppListSpecifics_AppListItemType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AppListSpecifics_AppListItemType_value, data, "AppListSpecifics_AppListItemType")
	if err != nil {
		return err
	}
	*x = AppListSpecifics_AppListItemType(value)
	return nil
}
func (AppListSpecifics_AppListItemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_app_list_specifics_cf6ebc6332efdaf1, []int{0, 0}
}

// Properties of app list objects.
type AppListSpecifics struct {
	// Unique identifier for the item:
	// * TYPE_FOLDER: Folder id (generated)
	// * TYPE_APP: App Id
	ItemId   *string                           `protobuf:"bytes,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	ItemType *AppListSpecifics_AppListItemType `protobuf:"varint,2,opt,name=item_type,json=itemType,enum=sync_pb.AppListSpecifics_AppListItemType" json:"item_type,omitempty"`
	// Item name (FOLDER).
	ItemName *string `protobuf:"bytes,3,opt,name=item_name,json=itemName" json:"item_name,omitempty"`
	// Id of the parent (folder) item.
	ParentId *string `protobuf:"bytes,4,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	// Marked OBSOLETE because this is unused for the app list.
	// Which page this item will appear on in the app list.
	OBSOLETEPageOrdinal *string `protobuf:"bytes,5,opt,name=OBSOLETE_page_ordinal,json=OBSOLETEPageOrdinal" json:"OBSOLETE_page_ordinal,omitempty"` // Deprecated: Do not use.
	// Where on a page this item will appear.
	ItemOrdinal *string `protobuf:"bytes,6,opt,name=item_ordinal,json=itemOrdinal" json:"item_ordinal,omitempty"`
	// Where on a shelf this item will appear.
	ItemPinOrdinal       *string  `protobuf:"bytes,7,opt,name=item_pin_ordinal,json=itemPinOrdinal" json:"item_pin_ordinal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppListSpecifics) Reset()         { *m = AppListSpecifics{} }
func (m *AppListSpecifics) String() string { return proto.CompactTextString(m) }
func (*AppListSpecifics) ProtoMessage()    {}
func (*AppListSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_list_specifics_cf6ebc6332efdaf1, []int{0}
}
func (m *AppListSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppListSpecifics.Unmarshal(m, b)
}
func (m *AppListSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppListSpecifics.Marshal(b, m, deterministic)
}
func (dst *AppListSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppListSpecifics.Merge(dst, src)
}
func (m *AppListSpecifics) XXX_Size() int {
	return xxx_messageInfo_AppListSpecifics.Size(m)
}
func (m *AppListSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_AppListSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_AppListSpecifics proto.InternalMessageInfo

func (m *AppListSpecifics) GetItemId() string {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return ""
}

func (m *AppListSpecifics) GetItemType() AppListSpecifics_AppListItemType {
	if m != nil && m.ItemType != nil {
		return *m.ItemType
	}
	return AppListSpecifics_TYPE_APP
}

func (m *AppListSpecifics) GetItemName() string {
	if m != nil && m.ItemName != nil {
		return *m.ItemName
	}
	return ""
}

func (m *AppListSpecifics) GetParentId() string {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return ""
}

// Deprecated: Do not use.
func (m *AppListSpecifics) GetOBSOLETEPageOrdinal() string {
	if m != nil && m.OBSOLETEPageOrdinal != nil {
		return *m.OBSOLETEPageOrdinal
	}
	return ""
}

func (m *AppListSpecifics) GetItemOrdinal() string {
	if m != nil && m.ItemOrdinal != nil {
		return *m.ItemOrdinal
	}
	return ""
}

func (m *AppListSpecifics) GetItemPinOrdinal() string {
	if m != nil && m.ItemPinOrdinal != nil {
		return *m.ItemPinOrdinal
	}
	return ""
}

func init() {
	proto.RegisterType((*AppListSpecifics)(nil), "sync_pb.AppListSpecifics")
	proto.RegisterEnum("sync_pb.AppListSpecifics_AppListItemType", AppListSpecifics_AppListItemType_name, AppListSpecifics_AppListItemType_value)
}

func init() {
	proto.RegisterFile("components/sync/protocol/app_list_specifics.proto", fileDescriptor_app_list_specifics_cf6ebc6332efdaf1)
}

var fileDescriptor_app_list_specifics_cf6ebc6332efdaf1 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x6b, 0xc2, 0x30,
	0x18, 0xc6, 0xa9, 0xf5, 0x6f, 0x14, 0xed, 0x22, 0x62, 0xc1, 0x8b, 0x13, 0x06, 0x8e, 0x41, 0x65,
	0x3b, 0xec, 0xde, 0x62, 0xdc, 0x64, 0xdd, 0x5a, 0x6a, 0x1d, 0xec, 0x14, 0xba, 0x36, 0x73, 0x01,
	0x9b, 0x84, 0x36, 0x3b, 0xf4, 0x43, 0xed, 0x3b, 0x8e, 0xc6, 0xd6, 0x81, 0xb7, 0xe4, 0xf9, 0xfd,
	0x78, 0x5e, 0x5e, 0x5e, 0x70, 0x1f, 0xf3, 0x54, 0x70, 0x46, 0x98, 0xcc, 0x57, 0x79, 0xc1, 0xe2,
	0x95, 0xc8, 0xb8, 0xe4, 0x31, 0x3f, 0xae, 0x22, 0x21, 0xf0, 0x91, 0xe6, 0x12, 0xe7, 0x82, 0xc4,
	0xf4, 0x8b, 0xc6, 0xb9, 0xa5, 0x18, 0xec, 0x94, 0x1e, 0x16, 0x9f, 0x8b, 0x5f, 0x1d, 0x18, 0xb6,
	0x10, 0x2e, 0xcd, 0xe5, 0xae, 0x76, 0xe0, 0x14, 0x74, 0xa8, 0x24, 0x29, 0xa6, 0x89, 0xa9, 0xcd,
	0xb5, 0x65, 0x2f, 0x68, 0x97, 0xdf, 0x6d, 0x02, 0x37, 0xa0, 0xa7, 0x80, 0x2c, 0x04, 0x31, 0x1b,
	0x73, 0x6d, 0x39, 0x7c, 0xb8, 0xb5, 0xaa, 0x2a, 0xeb, 0xb2, 0xa6, 0x0e, 0xb6, 0x92, 0xa4, 0x61,
	0x21, 0x48, 0xd0, 0xa5, 0xd5, 0x0b, 0xce, 0xaa, 0x1e, 0x16, 0xa5, 0xc4, 0xd4, 0xd5, 0x08, 0x05,
	0xdf, 0xa2, 0x54, 0x41, 0x11, 0x65, 0x84, 0xc9, 0x72, 0x7e, 0xf3, 0x04, 0x4f, 0xc1, 0x36, 0x81,
	0x8f, 0x60, 0xe2, 0x39, 0x3b, 0xcf, 0x45, 0x21, 0xc2, 0x22, 0x3a, 0x10, 0xcc, 0xb3, 0x84, 0xb2,
	0xe8, 0x68, 0xb6, 0x4a, 0xd1, 0x69, 0x98, 0x5a, 0x30, 0xae, 0x05, 0x3f, 0x3a, 0x10, 0xef, 0x84,
	0xe1, 0x35, 0x18, 0xa8, 0x89, 0xb5, 0xde, 0x56, 0xbd, 0xfd, 0x32, 0xab, 0x95, 0x25, 0x30, 0x94,
	0x22, 0x28, 0x3b, 0x6b, 0x1d, 0xa5, 0x0d, 0xcb, 0xdc, 0xa7, 0xac, 0x32, 0x17, 0x05, 0x18, 0x5d,
	0xec, 0x06, 0x07, 0xa0, 0x1b, 0x7e, 0xf8, 0x08, 0xdb, 0xbe, 0x6f, 0x68, 0x70, 0x06, 0xa6, 0xea,
	0x17, 0xa0, 0x57, 0xef, 0x1d, 0xe1, 0x35, 0xda, 0xd8, 0x7b, 0x37, 0x54, 0xb0, 0x01, 0x47, 0xa0,
	0xaf, 0xe0, 0xc6, 0x73, 0xd7, 0x28, 0x30, 0x74, 0x38, 0x01, 0x57, 0x2a, 0x38, 0x2f, 0xb6, 0x0f,
	0x5c, 0xa3, 0x09, 0xc7, 0x60, 0xa4, 0x62, 0xdf, 0x7e, 0x42, 0xd8, 0x09, 0x90, 0xfd, 0x62, 0xb4,
	0x9c, 0x3b, 0x70, 0xc3, 0xb3, 0x83, 0x15, 0x7f, 0x67, 0x3c, 0xa5, 0x3f, 0xa9, 0xf5, 0x7f, 0x7a,
	0x75, 0x07, 0xab, 0x3e, 0xfd, 0xb3, 0xee, 0x6b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x9b,
	0x7d, 0xf2, 0x19, 0x02, 0x00, 0x00,
}
