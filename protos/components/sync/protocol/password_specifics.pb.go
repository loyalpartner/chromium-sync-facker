// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/password_specifics.proto

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

// All the strings are encoded with UTF-8. URLs are encoded in Punycode.
type PasswordSpecificsData struct {
	// SCHEME_HTML(0), the credential represents either a parsed HTML form, or an
	// android credential or a password saved through Credential Manager API
	// (https://w3c.github.io/webappsec/specs/credentialmanagement/).
	// SCHEME_BASIC(1), basic access http authentication.
	// SCHEME_DIGEST(2), digest access authentication.
	// SCHEME_OTHER(3), another access authentication.
	Scheme *int32 `protobuf:"varint,1,opt,name=scheme" json:"scheme,omitempty"`
	// For parsed web forms and normal passwords saved through Credential Manager
	// API: url-scheme://url-host[:url-port]/
	// For Android apps (local + federated):
	//     "android://<hash of cert>@<package name>/"
	// where the hash is base64 encoded SHA512 of the app's public certificate.
	// For federated credentials:
	//     "federation://" + origin_host + "/" + federation_host
	// For proxy auth: proxy-host/auth-realm
	// For HTTP auth: url-scheme://url-host[:url-port]/auth-realm
	SignonRealm *string `protobuf:"bytes,2,opt,name=signon_realm,json=signonRealm" json:"signon_realm,omitempty"`
	// For parsed web forms and Credential Manager API:
	//     url-scheme://url-host[:url-port]/path
	// For Android: "android://<hash of cert>@<package name>/"
	// For proxy/HTTP auth: url-scheme://url-host[:url-port]/path
	Origin *string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	// Only for web-parsed forms - the action target of the form:
	//     url-scheme://url-host[:url-port]/path
	Action *string `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	// Only for web-parsed forms - the name of the element containing username.
	UsernameElement *string `protobuf:"bytes,5,opt,name=username_element,json=usernameElement" json:"username_element,omitempty"`
	// For all: the username.
	// For blacklisted forms: <empty>.
	UsernameValue *string `protobuf:"bytes,6,opt,name=username_value,json=usernameValue" json:"username_value,omitempty"`
	// Only for web-parsed forms - the name of the element containing password.
	PasswordElement *string `protobuf:"bytes,7,opt,name=password_element,json=passwordElement" json:"password_element,omitempty"`
	// For all: the password.
	// For federated logins and blacklisted forms: <empty>
	PasswordValue *string `protobuf:"bytes,8,opt,name=password_value,json=passwordValue" json:"password_value,omitempty"`
	// Deprecated: http://crbug.com/413020
	// True if the credential was saved for a HTTPS session with a valid SSL cert.
	// Ignored for Android apps.
	SslValid *bool `protobuf:"varint,9,opt,name=ssl_valid,json=sslValid" json:"ssl_valid,omitempty"` // Deprecated: Do not use.
	// True for the last credential used for logging in on a given site.
	// Deprecated in M81.
	Preferred *bool `protobuf:"varint,10,opt,name=preferred" json:"preferred,omitempty"` // Deprecated: Do not use.
	// Time when the credential was created. Amount of microseconds since 1601.
	DateCreated *int64 `protobuf:"varint,11,opt,name=date_created,json=dateCreated" json:"date_created,omitempty"`
	// True, if user chose permanently not to save the credentials for the form.
	Blacklisted *bool `protobuf:"varint,12,opt,name=blacklisted" json:"blacklisted,omitempty"`
	// TYPE_MANUAL(0), user manually filled the username and the password.
	// TYPE_GENERATED(1), the credential was auto generated.
	Type *int32 `protobuf:"varint,13,opt,name=type" json:"type,omitempty"`
	// Number of times this login was used for logging in. Chrome uses this field
	// to distinguish log-in and sign-up forms.
	TimesUsed *int32 `protobuf:"varint,14,opt,name=times_used,json=timesUsed" json:"times_used,omitempty"`
	// A human readable name of the account holder. Set by CredentialManager API
	// and Android.
	DisplayName *string `protobuf:"bytes,15,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// A URL of the avatar for the credential. Set by CredentialManager API and
	// Android.
	AvatarUrl *string `protobuf:"bytes,16,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
	// A URL of the IdP used to verify the credential. Set by Credential Manager
	// API and Android.
	FederationUrl *string `protobuf:"bytes,17,opt,name=federation_url,json=federationUrl" json:"federation_url,omitempty"`
	// Time when the credential was last used. Amount of microseconds since 1601.
	DateLastUsed *int64 `protobuf:"varint,18,opt,name=date_last_used,json=dateLastUsed" json:"date_last_used,omitempty"`
	// Set if an issue was detected that puts this password at risk. All the
	// clients are expected to clear the field when the password value is updated.
	// 'reused' part can be additionally reset when the analysis on the entire
	// password store is completed.
	PasswordIssues       *PasswordSpecificsData_PasswordIssues `protobuf:"bytes,19,opt,name=password_issues,json=passwordIssues" json:"password_issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PasswordSpecificsData) Reset()         { *m = PasswordSpecificsData{} }
func (m *PasswordSpecificsData) String() string { return proto.CompactTextString(m) }
func (*PasswordSpecificsData) ProtoMessage()    {}
func (*PasswordSpecificsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_password_specifics_f60ffcfc0f5fdcd2, []int{0}
}
func (m *PasswordSpecificsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordSpecificsData.Unmarshal(m, b)
}
func (m *PasswordSpecificsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordSpecificsData.Marshal(b, m, deterministic)
}
func (dst *PasswordSpecificsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordSpecificsData.Merge(dst, src)
}
func (m *PasswordSpecificsData) XXX_Size() int {
	return xxx_messageInfo_PasswordSpecificsData.Size(m)
}
func (m *PasswordSpecificsData) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordSpecificsData.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordSpecificsData proto.InternalMessageInfo

func (m *PasswordSpecificsData) GetScheme() int32 {
	if m != nil && m.Scheme != nil {
		return *m.Scheme
	}
	return 0
}

func (m *PasswordSpecificsData) GetSignonRealm() string {
	if m != nil && m.SignonRealm != nil {
		return *m.SignonRealm
	}
	return ""
}

func (m *PasswordSpecificsData) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *PasswordSpecificsData) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *PasswordSpecificsData) GetUsernameElement() string {
	if m != nil && m.UsernameElement != nil {
		return *m.UsernameElement
	}
	return ""
}

func (m *PasswordSpecificsData) GetUsernameValue() string {
	if m != nil && m.UsernameValue != nil {
		return *m.UsernameValue
	}
	return ""
}

func (m *PasswordSpecificsData) GetPasswordElement() string {
	if m != nil && m.PasswordElement != nil {
		return *m.PasswordElement
	}
	return ""
}

func (m *PasswordSpecificsData) GetPasswordValue() string {
	if m != nil && m.PasswordValue != nil {
		return *m.PasswordValue
	}
	return ""
}

// Deprecated: Do not use.
func (m *PasswordSpecificsData) GetSslValid() bool {
	if m != nil && m.SslValid != nil {
		return *m.SslValid
	}
	return false
}

// Deprecated: Do not use.
func (m *PasswordSpecificsData) GetPreferred() bool {
	if m != nil && m.Preferred != nil {
		return *m.Preferred
	}
	return false
}

func (m *PasswordSpecificsData) GetDateCreated() int64 {
	if m != nil && m.DateCreated != nil {
		return *m.DateCreated
	}
	return 0
}

func (m *PasswordSpecificsData) GetBlacklisted() bool {
	if m != nil && m.Blacklisted != nil {
		return *m.Blacklisted
	}
	return false
}

func (m *PasswordSpecificsData) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *PasswordSpecificsData) GetTimesUsed() int32 {
	if m != nil && m.TimesUsed != nil {
		return *m.TimesUsed
	}
	return 0
}

func (m *PasswordSpecificsData) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *PasswordSpecificsData) GetAvatarUrl() string {
	if m != nil && m.AvatarUrl != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *PasswordSpecificsData) GetFederationUrl() string {
	if m != nil && m.FederationUrl != nil {
		return *m.FederationUrl
	}
	return ""
}

func (m *PasswordSpecificsData) GetDateLastUsed() int64 {
	if m != nil && m.DateLastUsed != nil {
		return *m.DateLastUsed
	}
	return 0
}

func (m *PasswordSpecificsData) GetPasswordIssues() *PasswordSpecificsData_PasswordIssues {
	if m != nil {
		return m.PasswordIssues
	}
	return nil
}

type PasswordSpecificsData_PasswordIssues struct {
	LeakedPasswordIssue  *PasswordSpecificsData_PasswordIssues_PasswordIssue `protobuf:"bytes,1,opt,name=leaked_password_issue,json=leakedPasswordIssue" json:"leaked_password_issue,omitempty"`
	ReusedPasswordIssue  *PasswordSpecificsData_PasswordIssues_PasswordIssue `protobuf:"bytes,2,opt,name=reused_password_issue,json=reusedPasswordIssue" json:"reused_password_issue,omitempty"`
	WeakPasswordIssue    *PasswordSpecificsData_PasswordIssues_PasswordIssue `protobuf:"bytes,3,opt,name=weak_password_issue,json=weakPasswordIssue" json:"weak_password_issue,omitempty"`
	PhishedPasswordIssue *PasswordSpecificsData_PasswordIssues_PasswordIssue `protobuf:"bytes,4,opt,name=phished_password_issue,json=phishedPasswordIssue" json:"phished_password_issue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *PasswordSpecificsData_PasswordIssues) Reset()         { *m = PasswordSpecificsData_PasswordIssues{} }
func (m *PasswordSpecificsData_PasswordIssues) String() string { return proto.CompactTextString(m) }
func (*PasswordSpecificsData_PasswordIssues) ProtoMessage()    {}
func (*PasswordSpecificsData_PasswordIssues) Descriptor() ([]byte, []int) {
	return fileDescriptor_password_specifics_f60ffcfc0f5fdcd2, []int{0, 0}
}
func (m *PasswordSpecificsData_PasswordIssues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordSpecificsData_PasswordIssues.Unmarshal(m, b)
}
func (m *PasswordSpecificsData_PasswordIssues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordSpecificsData_PasswordIssues.Marshal(b, m, deterministic)
}
func (dst *PasswordSpecificsData_PasswordIssues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordSpecificsData_PasswordIssues.Merge(dst, src)
}
func (m *PasswordSpecificsData_PasswordIssues) XXX_Size() int {
	return xxx_messageInfo_PasswordSpecificsData_PasswordIssues.Size(m)
}
func (m *PasswordSpecificsData_PasswordIssues) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordSpecificsData_PasswordIssues.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordSpecificsData_PasswordIssues proto.InternalMessageInfo

func (m *PasswordSpecificsData_PasswordIssues) GetLeakedPasswordIssue() *PasswordSpecificsData_PasswordIssues_PasswordIssue {
	if m != nil {
		return m.LeakedPasswordIssue
	}
	return nil
}

func (m *PasswordSpecificsData_PasswordIssues) GetReusedPasswordIssue() *PasswordSpecificsData_PasswordIssues_PasswordIssue {
	if m != nil {
		return m.ReusedPasswordIssue
	}
	return nil
}

func (m *PasswordSpecificsData_PasswordIssues) GetWeakPasswordIssue() *PasswordSpecificsData_PasswordIssues_PasswordIssue {
	if m != nil {
		return m.WeakPasswordIssue
	}
	return nil
}

func (m *PasswordSpecificsData_PasswordIssues) GetPhishedPasswordIssue() *PasswordSpecificsData_PasswordIssues_PasswordIssue {
	if m != nil {
		return m.PhishedPasswordIssue
	}
	return nil
}

type PasswordSpecificsData_PasswordIssues_PasswordIssue struct {
	// Timestamp set by a client detecting the issue for the first time.
	// Number of microseconds since Windows epoch (1601).
	DateFirstDetectionMicroseconds *uint64 `protobuf:"varint,1,opt,name=date_first_detection_microseconds,json=dateFirstDetectionMicroseconds" json:"date_first_detection_microseconds,omitempty"`
	// Whether the issue was muted by user.
	IsMuted              *bool    `protobuf:"varint,2,opt,name=is_muted,json=isMuted" json:"is_muted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) Reset() {
	*m = PasswordSpecificsData_PasswordIssues_PasswordIssue{}
}
func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) String() string {
	return proto.CompactTextString(m)
}
func (*PasswordSpecificsData_PasswordIssues_PasswordIssue) ProtoMessage() {}
func (*PasswordSpecificsData_PasswordIssues_PasswordIssue) Descriptor() ([]byte, []int) {
	return fileDescriptor_password_specifics_f60ffcfc0f5fdcd2, []int{0, 0, 0}
}
func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordSpecificsData_PasswordIssues_PasswordIssue.Unmarshal(m, b)
}
func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordSpecificsData_PasswordIssues_PasswordIssue.Marshal(b, m, deterministic)
}
func (dst *PasswordSpecificsData_PasswordIssues_PasswordIssue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordSpecificsData_PasswordIssues_PasswordIssue.Merge(dst, src)
}
func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) XXX_Size() int {
	return xxx_messageInfo_PasswordSpecificsData_PasswordIssues_PasswordIssue.Size(m)
}
func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordSpecificsData_PasswordIssues_PasswordIssue.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordSpecificsData_PasswordIssues_PasswordIssue proto.InternalMessageInfo

func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) GetDateFirstDetectionMicroseconds() uint64 {
	if m != nil && m.DateFirstDetectionMicroseconds != nil {
		return *m.DateFirstDetectionMicroseconds
	}
	return 0
}

func (m *PasswordSpecificsData_PasswordIssues_PasswordIssue) GetIsMuted() bool {
	if m != nil && m.IsMuted != nil {
		return *m.IsMuted
	}
	return false
}

// Contains the password specifics metadata which simplifies its lookup.
type PasswordSpecificsMetadata struct {
	Url *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// True, if user chose permanently not to save the credentials for the form.
	// Introduced in M82.
	Blacklisted          *bool    `protobuf:"varint,2,opt,name=blacklisted" json:"blacklisted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordSpecificsMetadata) Reset()         { *m = PasswordSpecificsMetadata{} }
func (m *PasswordSpecificsMetadata) String() string { return proto.CompactTextString(m) }
func (*PasswordSpecificsMetadata) ProtoMessage()    {}
func (*PasswordSpecificsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_password_specifics_f60ffcfc0f5fdcd2, []int{1}
}
func (m *PasswordSpecificsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordSpecificsMetadata.Unmarshal(m, b)
}
func (m *PasswordSpecificsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordSpecificsMetadata.Marshal(b, m, deterministic)
}
func (dst *PasswordSpecificsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordSpecificsMetadata.Merge(dst, src)
}
func (m *PasswordSpecificsMetadata) XXX_Size() int {
	return xxx_messageInfo_PasswordSpecificsMetadata.Size(m)
}
func (m *PasswordSpecificsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordSpecificsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordSpecificsMetadata proto.InternalMessageInfo

func (m *PasswordSpecificsMetadata) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *PasswordSpecificsMetadata) GetBlacklisted() bool {
	if m != nil && m.Blacklisted != nil {
		return *m.Blacklisted
	}
	return false
}

// Properties of password sync objects.
type PasswordSpecifics struct {
	// The actual password data. Contains an encrypted PasswordSpecificsData
	// message.
	Encrypted *EncryptedData `protobuf:"bytes,1,opt,name=encrypted" json:"encrypted,omitempty"`
	// An unsynced field for use internally on the client. This field should
	// never be set in any network-based communications because it contains
	// unencrypted material.
	ClientOnlyEncryptedData *PasswordSpecificsData `protobuf:"bytes,2,opt,name=client_only_encrypted_data,json=clientOnlyEncryptedData" json:"client_only_encrypted_data,omitempty"`
	// Password related metadata, which is sent to the server side. The field
	// should never be set for full encryption users. If encryption is enabled,
	// this field must be cleared.
	UnencryptedMetadata  *PasswordSpecificsMetadata `protobuf:"bytes,3,opt,name=unencrypted_metadata,json=unencryptedMetadata" json:"unencrypted_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PasswordSpecifics) Reset()         { *m = PasswordSpecifics{} }
func (m *PasswordSpecifics) String() string { return proto.CompactTextString(m) }
func (*PasswordSpecifics) ProtoMessage()    {}
func (*PasswordSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_password_specifics_f60ffcfc0f5fdcd2, []int{2}
}
func (m *PasswordSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordSpecifics.Unmarshal(m, b)
}
func (m *PasswordSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordSpecifics.Marshal(b, m, deterministic)
}
func (dst *PasswordSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordSpecifics.Merge(dst, src)
}
func (m *PasswordSpecifics) XXX_Size() int {
	return xxx_messageInfo_PasswordSpecifics.Size(m)
}
func (m *PasswordSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordSpecifics proto.InternalMessageInfo

func (m *PasswordSpecifics) GetEncrypted() *EncryptedData {
	if m != nil {
		return m.Encrypted
	}
	return nil
}

func (m *PasswordSpecifics) GetClientOnlyEncryptedData() *PasswordSpecificsData {
	if m != nil {
		return m.ClientOnlyEncryptedData
	}
	return nil
}

func (m *PasswordSpecifics) GetUnencryptedMetadata() *PasswordSpecificsMetadata {
	if m != nil {
		return m.UnencryptedMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*PasswordSpecificsData)(nil), "sync_pb.PasswordSpecificsData")
	proto.RegisterType((*PasswordSpecificsData_PasswordIssues)(nil), "sync_pb.PasswordSpecificsData.PasswordIssues")
	proto.RegisterType((*PasswordSpecificsData_PasswordIssues_PasswordIssue)(nil), "sync_pb.PasswordSpecificsData.PasswordIssues.PasswordIssue")
	proto.RegisterType((*PasswordSpecificsMetadata)(nil), "sync_pb.PasswordSpecificsMetadata")
	proto.RegisterType((*PasswordSpecifics)(nil), "sync_pb.PasswordSpecifics")
}

func init() {
	proto.RegisterFile("components/sync/protocol/password_specifics.proto", fileDescriptor_password_specifics_f60ffcfc0f5fdcd2)
}

var fileDescriptor_password_specifics_f60ffcfc0f5fdcd2 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x6f, 0x4f, 0xdb, 0x38,
	0x1c, 0x56, 0xda, 0x02, 0xad, 0x0b, 0x05, 0x5c, 0xe0, 0x42, 0xa5, 0xe3, 0x4a, 0x75, 0x48, 0x45,
	0xa7, 0x2b, 0x3a, 0x74, 0xef, 0xee, 0x1d, 0x07, 0xd3, 0x90, 0xc6, 0x40, 0x99, 0xe0, 0xcd, 0x5e,
	0x58, 0x26, 0xf9, 0x95, 0x5a, 0x75, 0xe2, 0xcc, 0x76, 0x40, 0xfd, 0x1e, 0xfb, 0x52, 0xfb, 0x50,
	0x93, 0x26, 0xdb, 0x49, 0xda, 0xd0, 0x31, 0x69, 0x62, 0xef, 0xe2, 0xc7, 0x8f, 0x9f, 0xc7, 0xbf,
	0x7f, 0x0e, 0xfa, 0x27, 0x14, 0x71, 0x2a, 0x12, 0x48, 0xb4, 0x3a, 0x51, 0xb3, 0x24, 0x3c, 0x49,
	0xa5, 0xd0, 0x22, 0x14, 0xfc, 0x24, 0xa5, 0x4a, 0x3d, 0x09, 0x19, 0x11, 0x95, 0x42, 0xc8, 0xc6,
	0x2c, 0x54, 0x23, 0xbb, 0x87, 0xd7, 0x0c, 0x8f, 0xa4, 0xf7, 0xbd, 0xe3, 0x17, 0xcf, 0x42, 0x12,
	0xca, 0x59, 0xaa, 0x99, 0x48, 0xdc, 0x99, 0xc1, 0x67, 0x84, 0x76, 0x6f, 0x72, 0xc1, 0x0f, 0x85,
	0xde, 0x39, 0xd5, 0x14, 0xef, 0xa1, 0x55, 0x15, 0x4e, 0x20, 0x06, 0xdf, 0xeb, 0x7b, 0xc3, 0x95,
	0x20, 0x5f, 0xe1, 0x43, 0xb4, 0xae, 0xd8, 0x43, 0x22, 0x12, 0x22, 0x81, 0xf2, 0xd8, 0xaf, 0xf5,
	0xbd, 0x61, 0x2b, 0x68, 0x3b, 0x2c, 0x30, 0x90, 0x39, 0x2a, 0x24, 0x7b, 0x60, 0x89, 0x5f, 0xb7,
	0x9b, 0xf9, 0xca, 0xe0, 0x34, 0x34, 0xe6, 0x7e, 0xc3, 0xe1, 0x6e, 0x85, 0x8f, 0xd1, 0x56, 0xa6,
	0x40, 0x26, 0x34, 0x06, 0x02, 0x1c, 0x62, 0x48, 0xb4, 0xbf, 0x62, 0x19, 0x9b, 0x05, 0x7e, 0xe1,
	0x60, 0x7c, 0x84, 0x3a, 0x25, 0xf5, 0x91, 0xf2, 0x0c, 0xfc, 0x55, 0x4b, 0xdc, 0x28, 0xd0, 0x3b,
	0x03, 0x1a, 0xc5, 0x32, 0x4d, 0x85, 0xe2, 0x9a, 0x53, 0x2c, 0xf0, 0x05, 0xc5, 0x92, 0xea, 0x14,
	0x9b, 0x4e, 0xb1, 0x40, 0x9d, 0xe2, 0x1f, 0xa8, 0xa5, 0x14, 0x37, 0x0c, 0x16, 0xf9, 0xad, 0xbe,
	0x37, 0x6c, 0x9e, 0xd5, 0x7c, 0x2f, 0x68, 0x2a, 0xc5, 0xef, 0x0c, 0x86, 0xfb, 0xa8, 0x95, 0x4a,
	0x18, 0x83, 0x94, 0x10, 0xf9, 0xa8, 0x24, 0xcc, 0x41, 0x93, 0xb9, 0x88, 0x6a, 0x20, 0xa1, 0x04,
	0xaa, 0x21, 0xf2, 0xdb, 0x7d, 0x6f, 0x58, 0x0f, 0xda, 0x06, 0xfb, 0xdf, 0x41, 0xb8, 0x8f, 0xda,
	0xf7, 0x9c, 0x86, 0x53, 0xce, 0x94, 0x61, 0xac, 0x1b, 0x99, 0x60, 0x11, 0xc2, 0x18, 0x35, 0xf4,
	0x2c, 0x05, 0x7f, 0xc3, 0x16, 0xc5, 0x7e, 0xe3, 0xdf, 0x11, 0xd2, 0x2c, 0x06, 0x45, 0x32, 0x05,
	0x91, 0xdf, 0xb1, 0x3b, 0x2d, 0x8b, 0xdc, 0xaa, 0xdc, 0x97, 0xa9, 0x94, 0xd3, 0x19, 0x31, 0x19,
	0xf2, 0x37, 0x5d, 0xc5, 0x72, 0xec, 0x3d, 0x8d, 0xad, 0x02, 0x7d, 0xa4, 0x9a, 0x4a, 0x92, 0x49,
	0xee, 0x6f, 0x59, 0x42, 0xcb, 0x21, 0xb7, 0x92, 0x9b, 0x1c, 0x8d, 0x21, 0x02, 0x49, 0x4d, 0xb9,
	0x2c, 0x65, 0xdb, 0xe5, 0x68, 0x8e, 0x1a, 0xda, 0x9f, 0xa8, 0x63, 0x03, 0xe4, 0x54, 0x69, 0x77,
	0x17, 0x6c, 0x43, 0xb4, 0x61, 0xbf, 0xa3, 0x4a, 0xdb, 0xeb, 0xdc, 0xa1, 0xb2, 0x06, 0x84, 0x29,
	0x95, 0x81, 0xf2, 0xbb, 0x7d, 0x6f, 0xd8, 0x3e, 0xfd, 0x7b, 0x94, 0x37, 0xf0, 0xe8, 0xbb, 0x1d,
	0x59, 0xa2, 0x97, 0xf6, 0x50, 0x50, 0x96, 0xcd, 0xad, 0x7b, 0x5f, 0x1a, 0xa8, 0x53, 0xa5, 0x60,
	0x81, 0x76, 0x39, 0xd0, 0x29, 0x44, 0xa4, 0xea, 0x68, 0x5b, 0xba, 0x7d, 0xfa, 0xdf, 0x4f, 0x19,
	0x56, 0x97, 0x41, 0xd7, 0x29, 0x57, 0x40, 0x63, 0x28, 0xc1, 0x44, 0xfe, 0xdc, 0xb0, 0xf6, 0x0b,
	0x0c, 0x9d, 0x72, 0xd5, 0x70, 0x8a, 0xba, 0x4f, 0x40, 0xa7, 0xcf, 0xed, 0xea, 0xaf, 0xb7, 0xdb,
	0x36, 0xba, 0x55, 0xb3, 0x4f, 0x68, 0x2f, 0x9d, 0x30, 0x35, 0x59, 0x0e, 0xaf, 0xf1, 0x7a, 0xbf,
	0x9d, 0x5c, 0xba, 0x82, 0xf6, 0x32, 0xb4, 0x51, 0xbd, 0xc3, 0x25, 0x3a, 0xb4, 0x3d, 0x36, 0x66,
	0x52, 0x69, 0x12, 0x81, 0x06, 0xfb, 0x86, 0x90, 0x98, 0x85, 0x52, 0x28, 0x08, 0x45, 0x12, 0x29,
	0x5b, 0xde, 0x46, 0x70, 0x60, 0x88, 0x6f, 0x0c, 0xef, 0xbc, 0xa0, 0x5d, 0x2d, 0xb0, 0xf0, 0x3e,
	0x6a, 0x32, 0x45, 0xe2, 0xcc, 0x4c, 0x5a, 0xcd, 0x4e, 0xda, 0x1a, 0x53, 0x57, 0x66, 0x39, 0xb8,
	0x46, 0xfb, 0x4b, 0x21, 0x5c, 0x81, 0xa6, 0x91, 0x79, 0x19, 0xb7, 0x50, 0xdd, 0x8c, 0x80, 0x67,
	0x47, 0xc0, 0x7c, 0x3e, 0x1f, 0xdb, 0xda, 0xd2, 0xd8, 0x0e, 0xbe, 0x7a, 0x68, 0x7b, 0x49, 0x11,
	0xff, 0x8b, 0x5a, 0xf9, 0x8b, 0x0c, 0x51, 0xde, 0x93, 0x7b, 0x65, 0x0e, 0x2f, 0x8a, 0x1d, 0x93,
	0xbb, 0x60, 0x4e, 0xc4, 0x1f, 0x51, 0x2f, 0xe4, 0x0c, 0x12, 0x4d, 0x44, 0xc2, 0x67, 0xa4, 0xdc,
	0x20, 0xe6, 0x76, 0x79, 0xa7, 0x1d, 0xfc, 0xb8, 0x14, 0xc1, 0x6f, 0x4e, 0xe1, 0x3a, 0xe1, 0xb3,
	0x8a, 0x0f, 0xbe, 0x45, 0x3b, 0x59, 0x32, 0x97, 0x8c, 0xf3, 0xa0, 0xf3, 0x8e, 0x1a, 0xbc, 0x2c,
	0x5b, 0xa4, 0x27, 0xe8, 0x2e, 0x9c, 0x2f, 0xc0, 0xb3, 0xbf, 0xd0, 0x91, 0x90, 0x0f, 0xa3, 0x70,
	0x22, 0x45, 0xcc, 0xb2, 0x78, 0x34, 0xff, 0x43, 0x59, 0xc5, 0x51, 0xf1, 0x87, 0x7a, 0x5b, 0xbf,
	0xf1, 0xbe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x71, 0xb1, 0x68, 0xfc, 0x06, 0x00, 0x00,
}
