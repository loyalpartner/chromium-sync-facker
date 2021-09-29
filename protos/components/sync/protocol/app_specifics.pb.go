// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/app_specifics.proto

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

// The possible launch types for an app.
// This enum should be kept in sync with extensions::LaunchType.
type AppSpecifics_LaunchType int32

const (
	AppSpecifics_PINNED     AppSpecifics_LaunchType = 0
	AppSpecifics_REGULAR    AppSpecifics_LaunchType = 1
	AppSpecifics_FULLSCREEN AppSpecifics_LaunchType = 2
	AppSpecifics_WINDOW     AppSpecifics_LaunchType = 3
)

var AppSpecifics_LaunchType_name = map[int32]string{
	0: "PINNED",
	1: "REGULAR",
	2: "FULLSCREEN",
	3: "WINDOW",
}
var AppSpecifics_LaunchType_value = map[string]int32{
	"PINNED":     0,
	"REGULAR":    1,
	"FULLSCREEN": 2,
	"WINDOW":     3,
}

func (x AppSpecifics_LaunchType) Enum() *AppSpecifics_LaunchType {
	p := new(AppSpecifics_LaunchType)
	*p = x
	return p
}
func (x AppSpecifics_LaunchType) String() string {
	return proto.EnumName(AppSpecifics_LaunchType_name, int32(x))
}
func (x *AppSpecifics_LaunchType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AppSpecifics_LaunchType_value, data, "AppSpecifics_LaunchType")
	if err != nil {
		return err
	}
	*x = AppSpecifics_LaunchType(value)
	return nil
}
func (AppSpecifics_LaunchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_app_specifics_adb82cf65975d053, []int{2, 0}
}

// Settings related to push notifications for apps.
type AppNotificationSettings struct {
	// DEPRECATED: Use oauth_client_id below.
	// Whether or not the user has setup notifications at least once.
	// The value for this field will start out false and will be set
	// to true when the user accepts receiving notifications for the
	// first time and then it will always remain true.
	InitialSetupDone *bool `protobuf:"varint,1,opt,name=initial_setup_done,json=initialSetupDone" json:"initial_setup_done,omitempty"` // Deprecated: Do not use.
	// Whether or not the user has disabled notifications.
	Disabled *bool `protobuf:"varint,2,opt,name=disabled" json:"disabled,omitempty"`
	// OAuth2 client id to which the user granted the notification permission.
	// This field will start out empty.
	// It will be set when the user accepts receiving notifications.
	// This field is used when the user revokes the notifications permission.
	// Note that it is never cleared after it was set once. Hence, the presence
	// of this field can be used to determine if the user has setup notifications
	// at least once for the given app.
	OauthClientId        *string  `protobuf:"bytes,3,opt,name=oauth_client_id,json=oauthClientId" json:"oauth_client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppNotificationSettings) Reset()         { *m = AppNotificationSettings{} }
func (m *AppNotificationSettings) String() string { return proto.CompactTextString(m) }
func (*AppNotificationSettings) ProtoMessage()    {}
func (*AppNotificationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_specifics_adb82cf65975d053, []int{0}
}
func (m *AppNotificationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppNotificationSettings.Unmarshal(m, b)
}
func (m *AppNotificationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppNotificationSettings.Marshal(b, m, deterministic)
}
func (dst *AppNotificationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppNotificationSettings.Merge(dst, src)
}
func (m *AppNotificationSettings) XXX_Size() int {
	return xxx_messageInfo_AppNotificationSettings.Size(m)
}
func (m *AppNotificationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AppNotificationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AppNotificationSettings proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *AppNotificationSettings) GetInitialSetupDone() bool {
	if m != nil && m.InitialSetupDone != nil {
		return *m.InitialSetupDone
	}
	return false
}

func (m *AppNotificationSettings) GetDisabled() bool {
	if m != nil && m.Disabled != nil {
		return *m.Disabled
	}
	return false
}

func (m *AppNotificationSettings) GetOauthClientId() string {
	if m != nil && m.OauthClientId != nil {
		return *m.OauthClientId
	}
	return ""
}

// Information about a linked app icon.
type LinkedAppIconInfo struct {
	// The URL of the app icon.
	Url *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// The size of the app icon in DIPs.
	Size                 *uint32  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkedAppIconInfo) Reset()         { *m = LinkedAppIconInfo{} }
func (m *LinkedAppIconInfo) String() string { return proto.CompactTextString(m) }
func (*LinkedAppIconInfo) ProtoMessage()    {}
func (*LinkedAppIconInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_specifics_adb82cf65975d053, []int{1}
}
func (m *LinkedAppIconInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkedAppIconInfo.Unmarshal(m, b)
}
func (m *LinkedAppIconInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkedAppIconInfo.Marshal(b, m, deterministic)
}
func (dst *LinkedAppIconInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkedAppIconInfo.Merge(dst, src)
}
func (m *LinkedAppIconInfo) XXX_Size() int {
	return xxx_messageInfo_LinkedAppIconInfo.Size(m)
}
func (m *LinkedAppIconInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkedAppIconInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LinkedAppIconInfo proto.InternalMessageInfo

func (m *LinkedAppIconInfo) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *LinkedAppIconInfo) GetSize() uint32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

// Properties of app sync objects.
//
// For now, an app is just an extension.  We keep the two data types
// separate for future-proofing purposes.
type AppSpecifics struct {
	// Extension data.
	Extension *ExtensionSpecifics `protobuf:"bytes,1,opt,name=extension" json:"extension,omitempty"`
	// Notification settings.
	NotificationSettings *AppNotificationSettings `protobuf:"bytes,2,opt,name=notification_settings,json=notificationSettings" json:"notification_settings,omitempty"`
	// This controls where on a page this application icon will appear.
	AppLaunchOrdinal *string `protobuf:"bytes,3,opt,name=app_launch_ordinal,json=appLaunchOrdinal" json:"app_launch_ordinal,omitempty"`
	// This specifics which page the application icon will appear on in the NTP.
	// This values only provide the order within the application pages, not within
	// all of the panels in the NTP.
	PageOrdinal *string `protobuf:"bytes,4,opt,name=page_ordinal,json=pageOrdinal" json:"page_ordinal,omitempty"`
	// This describes how the extension should be launched.
	LaunchType *AppSpecifics_LaunchType `protobuf:"varint,5,opt,name=launch_type,json=launchType,enum=sync_pb.AppSpecifics_LaunchType" json:"launch_type,omitempty"`
	// This is the url of a bookmark app. If this exists, the app is a bookmark
	// app.
	BookmarkAppUrl *string `protobuf:"bytes,6,opt,name=bookmark_app_url,json=bookmarkAppUrl" json:"bookmark_app_url,omitempty"`
	// This is the description of a bookmark app.
	BookmarkAppDescription *string `protobuf:"bytes,7,opt,name=bookmark_app_description,json=bookmarkAppDescription" json:"bookmark_app_description,omitempty"`
	// This is the color to use when generating bookmark app icons. The string is
	// in #rrggbb or #rgb syntax, e.g. #d8d8d8.
	BookmarkAppIconColor *string `protobuf:"bytes,8,opt,name=bookmark_app_icon_color,json=bookmarkAppIconColor" json:"bookmark_app_icon_color,omitempty"`
	// This is information about linked icons (that is, icons that are downloaded
	// from outside the app's bundle of files.
	LinkedAppIcons []*LinkedAppIconInfo `protobuf:"bytes,9,rep,name=linked_app_icons,json=linkedAppIcons" json:"linked_app_icons,omitempty"`
	// This is the scope of the bookmark app.
	BookmarkAppScope *string `protobuf:"bytes,10,opt,name=bookmark_app_scope,json=bookmarkAppScope" json:"bookmark_app_scope,omitempty"`
	// This is the SkColor used for the browser frame for this bookmark app.
	BookmarkAppThemeColor *uint32  `protobuf:"varint,11,opt,name=bookmark_app_theme_color,json=bookmarkAppThemeColor" json:"bookmark_app_theme_color,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *AppSpecifics) Reset()         { *m = AppSpecifics{} }
func (m *AppSpecifics) String() string { return proto.CompactTextString(m) }
func (*AppSpecifics) ProtoMessage()    {}
func (*AppSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_app_specifics_adb82cf65975d053, []int{2}
}
func (m *AppSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppSpecifics.Unmarshal(m, b)
}
func (m *AppSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppSpecifics.Marshal(b, m, deterministic)
}
func (dst *AppSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppSpecifics.Merge(dst, src)
}
func (m *AppSpecifics) XXX_Size() int {
	return xxx_messageInfo_AppSpecifics.Size(m)
}
func (m *AppSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_AppSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_AppSpecifics proto.InternalMessageInfo

func (m *AppSpecifics) GetExtension() *ExtensionSpecifics {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *AppSpecifics) GetNotificationSettings() *AppNotificationSettings {
	if m != nil {
		return m.NotificationSettings
	}
	return nil
}

func (m *AppSpecifics) GetAppLaunchOrdinal() string {
	if m != nil && m.AppLaunchOrdinal != nil {
		return *m.AppLaunchOrdinal
	}
	return ""
}

func (m *AppSpecifics) GetPageOrdinal() string {
	if m != nil && m.PageOrdinal != nil {
		return *m.PageOrdinal
	}
	return ""
}

func (m *AppSpecifics) GetLaunchType() AppSpecifics_LaunchType {
	if m != nil && m.LaunchType != nil {
		return *m.LaunchType
	}
	return AppSpecifics_PINNED
}

func (m *AppSpecifics) GetBookmarkAppUrl() string {
	if m != nil && m.BookmarkAppUrl != nil {
		return *m.BookmarkAppUrl
	}
	return ""
}

func (m *AppSpecifics) GetBookmarkAppDescription() string {
	if m != nil && m.BookmarkAppDescription != nil {
		return *m.BookmarkAppDescription
	}
	return ""
}

func (m *AppSpecifics) GetBookmarkAppIconColor() string {
	if m != nil && m.BookmarkAppIconColor != nil {
		return *m.BookmarkAppIconColor
	}
	return ""
}

func (m *AppSpecifics) GetLinkedAppIcons() []*LinkedAppIconInfo {
	if m != nil {
		return m.LinkedAppIcons
	}
	return nil
}

func (m *AppSpecifics) GetBookmarkAppScope() string {
	if m != nil && m.BookmarkAppScope != nil {
		return *m.BookmarkAppScope
	}
	return ""
}

func (m *AppSpecifics) GetBookmarkAppThemeColor() uint32 {
	if m != nil && m.BookmarkAppThemeColor != nil {
		return *m.BookmarkAppThemeColor
	}
	return 0
}

func init() {
	proto.RegisterType((*AppNotificationSettings)(nil), "sync_pb.AppNotificationSettings")
	proto.RegisterType((*LinkedAppIconInfo)(nil), "sync_pb.LinkedAppIconInfo")
	proto.RegisterType((*AppSpecifics)(nil), "sync_pb.AppSpecifics")
	proto.RegisterEnum("sync_pb.AppSpecifics_LaunchType", AppSpecifics_LaunchType_name, AppSpecifics_LaunchType_value)
}

func init() {
	proto.RegisterFile("components/sync/protocol/app_specifics.proto", fileDescriptor_app_specifics_adb82cf65975d053)
}

var fileDescriptor_app_specifics_adb82cf65975d053 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xdd, 0x6e, 0xda, 0x3c,
	0x18, 0xfe, 0x52, 0xfa, 0xc7, 0x9b, 0x96, 0x2f, 0xb3, 0xda, 0x35, 0xea, 0x4e, 0x18, 0xd2, 0x26,
	0xa4, 0x55, 0xe9, 0x84, 0x34, 0x6d, 0x3d, 0x4c, 0x81, 0x6d, 0x48, 0x88, 0x56, 0xa6, 0xa8, 0x87,
	0x51, 0x70, 0x5c, 0xb0, 0x30, 0xb6, 0x15, 0x1b, 0x69, 0xec, 0x26, 0x76, 0x49, 0xbb, 0xb5, 0xc9,
	0x06, 0x42, 0xe8, 0xd6, 0x33, 0xe7, 0xf9, 0xf1, 0xeb, 0xe7, 0x89, 0x0d, 0x57, 0x44, 0xce, 0x95,
	0x14, 0x54, 0x18, 0x7d, 0xad, 0x97, 0x82, 0x5c, 0xab, 0x5c, 0x1a, 0x49, 0x24, 0xbf, 0x4e, 0x95,
	0x4a, 0xb4, 0xa2, 0x84, 0x3d, 0x31, 0xa2, 0x23, 0x07, 0xa3, 0x23, 0x2b, 0x49, 0xd4, 0xf8, 0xb2,
	0xf5, 0xa2, 0x8d, 0xfe, 0x30, 0x54, 0x68, 0x26, 0xc5, 0x73, 0x73, 0xe3, 0x97, 0x07, 0x17, 0xb1,
	0x52, 0x03, 0x69, 0x2c, 0x9a, 0x1a, 0x26, 0xc5, 0x90, 0x1a, 0xc3, 0xc4, 0x44, 0xa3, 0x8f, 0x80,
	0x98, 0x60, 0x86, 0xa5, 0x3c, 0xd1, 0xd4, 0x2c, 0x54, 0x92, 0x49, 0x41, 0x43, 0xaf, 0xee, 0x35,
	0x8f, 0x6f, 0xf7, 0x42, 0x0f, 0x07, 0x6b, 0x76, 0x68, 0xc9, 0x8e, 0x14, 0x14, 0x5d, 0xc2, 0x71,
	0xc6, 0x74, 0x3a, 0xe6, 0x34, 0x0b, 0xf7, 0xac, 0x0e, 0x17, 0xdf, 0xe8, 0x3d, 0xfc, 0x2f, 0xd3,
	0x85, 0x99, 0x26, 0x84, 0x33, 0x2a, 0x4c, 0xc2, 0xb2, 0xb0, 0x52, 0xf7, 0x9a, 0x55, 0x7c, 0xea,
	0xe0, 0xb6, 0x43, 0x7b, 0x59, 0xe3, 0x06, 0x5e, 0xf5, 0x99, 0x98, 0xd1, 0x2c, 0x56, 0xaa, 0x47,
	0xa4, 0xe8, 0x89, 0x27, 0x89, 0x02, 0xa8, 0x2c, 0x72, 0xee, 0x66, 0x57, 0xb1, 0x5d, 0x22, 0x04,
	0xfb, 0x9a, 0xfd, 0xa4, 0x6e, 0xcc, 0x29, 0x76, 0xeb, 0xc6, 0xef, 0x03, 0x38, 0x89, 0x95, 0x1a,
	0x6e, 0x32, 0xa2, 0x1b, 0xa8, 0x16, 0xd1, 0x9d, 0xd9, 0x6f, 0xbd, 0x89, 0xd6, 0x75, 0x45, 0xdd,
	0x0d, 0x53, 0xe8, 0xf1, 0x56, 0x8d, 0x46, 0x70, 0x2e, 0x4a, 0xa5, 0xd8, 0x06, 0x5c, 0x2b, 0x6e,
	0xa0, 0xdf, 0xaa, 0x17, 0xdb, 0xbc, 0xd0, 0x1e, 0x3e, 0x13, 0xff, 0xea, 0xf4, 0x0a, 0x90, 0xfd,
	0x87, 0x3c, 0x5d, 0x08, 0x32, 0x4d, 0x64, 0x9e, 0x31, 0x91, 0xf2, 0x75, 0x11, 0x41, 0xaa, 0x54,
	0xdf, 0x11, 0x77, 0x2b, 0x1c, 0xbd, 0x85, 0x13, 0x95, 0x4e, 0x68, 0xa1, 0xdb, 0x77, 0x3a, 0xdf,
	0x62, 0x1b, 0x49, 0x0c, 0xfe, 0x7a, 0x33, 0xb3, 0x54, 0x34, 0x3c, 0xa8, 0x7b, 0xcd, 0xda, 0xee,
	0xe9, 0x8a, 0x78, 0xd1, 0x6a, 0xf3, 0x87, 0xa5, 0xa2, 0x18, 0x78, 0xb1, 0x46, 0x4d, 0x08, 0xc6,
	0x52, 0xce, 0xe6, 0x69, 0x3e, 0x4b, 0xec, 0xe1, 0x6c, 0xd3, 0x87, 0x6e, 0x52, 0x6d, 0x83, 0xc7,
	0x4a, 0x8d, 0x72, 0x8e, 0xbe, 0x40, 0xb8, 0xa3, 0xcc, 0xa8, 0x26, 0x39, 0x53, 0x36, 0x61, 0x78,
	0xe4, 0x1c, 0xaf, 0x4b, 0x8e, 0xce, 0x96, 0x45, 0x9f, 0xe0, 0x62, 0xc7, 0xc9, 0x88, 0x14, 0x09,
	0x91, 0x5c, 0xe6, 0xe1, 0xb1, 0x33, 0x9e, 0x95, 0x8c, 0xf6, 0xb7, 0xb7, 0x2d, 0x87, 0x3a, 0x10,
	0x70, 0x77, 0x19, 0x0a, 0x93, 0x0e, 0xab, 0xf5, 0x4a, 0xd3, 0x6f, 0x5d, 0x16, 0x11, 0xff, 0xba,
	0x2d, 0xb8, 0xc6, 0xcb, 0x90, 0x2b, 0x7d, 0x67, 0xb8, 0x26, 0x52, 0xd1, 0x10, 0x56, 0xa5, 0x97,
	0xe6, 0x0e, 0x2d, 0x8e, 0x3e, 0x3f, 0x0b, 0x69, 0xa6, 0x74, 0x4e, 0xd7, 0x67, 0xf5, 0xdd, 0x6d,
	0x3b, 0x2f, 0x79, 0x1e, 0x2c, 0xeb, 0x0e, 0xdb, 0x88, 0x01, 0xb6, 0x0d, 0x23, 0x80, 0xc3, 0xfb,
	0xde, 0x60, 0xd0, 0xed, 0x04, 0xff, 0x21, 0x1f, 0x8e, 0x70, 0xf7, 0xdb, 0xa8, 0x1f, 0xe3, 0xc0,
	0x43, 0x35, 0x80, 0xaf, 0xa3, 0x7e, 0x7f, 0xd8, 0xc6, 0xdd, 0xee, 0x20, 0xd8, 0xb3, 0xc2, 0xc7,
	0xde, 0xa0, 0x73, 0xf7, 0x18, 0x54, 0x6e, 0x3f, 0xc0, 0x3b, 0x99, 0x4f, 0x22, 0x32, 0xcd, 0xe5,
	0x9c, 0x2d, 0xe6, 0xd1, 0xf6, 0x45, 0xbb, 0xb8, 0xd1, 0xe6, 0x45, 0x7f, 0xaf, 0xdc, 0x7b, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x07, 0xe4, 0x2a, 0x27, 0x04, 0x00, 0x00,
}