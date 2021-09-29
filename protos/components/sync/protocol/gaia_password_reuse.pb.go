// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/gaia_password_reuse.proto

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

// If SafeBrowsing is enabled, is the user opted-in to extended
// reporting or Scout?
type GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation int32

const (
	GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_REPORTING_POPULATION_UNSPECIFIED GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation = 0
	GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_NONE                             GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation = 1
	GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_EXTENDED_REPORTING               GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation = 2
	GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_SCOUT                            GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation = 3
)

var GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation_name = map[int32]string{
	0: "REPORTING_POPULATION_UNSPECIFIED",
	1: "NONE",
	2: "EXTENDED_REPORTING",
	3: "SCOUT",
}
var GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation_value = map[string]int32{
	"REPORTING_POPULATION_UNSPECIFIED": 0,
	"NONE":                             1,
	"EXTENDED_REPORTING":               2,
	"SCOUT":                            3,
}

func (x GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation) Enum() *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation {
	p := new(GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation)
	*p = x
	return p
}
func (x GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation) String() string {
	return proto.EnumName(GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation_name, int32(x))
}
func (x *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation_value, data, "GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation")
	if err != nil {
		return err
	}
	*x = GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation(value)
	return nil
}
func (GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 0, 0, 0}
}

type GaiaPasswordReuse_PasswordReuseLookup_LookupResult int32

const (
	GaiaPasswordReuse_PasswordReuseLookup_UNSPECIFIED GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 0
	// URL did match the password reuse whitelist.
	// No further action required related to this re-use event.
	GaiaPasswordReuse_PasswordReuseLookup_WHITELIST_HIT GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 1
	// The URL exists in the client’s cache.
	// No further action required related to this re-use event.
	// This event also logs the ReputationVerdict.
	GaiaPasswordReuse_PasswordReuseLookup_CACHE_HIT GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 2
	// A valid response received from the SafeBrowsing service.
	// This event also logs the ReputationVerdict.
	GaiaPasswordReuse_PasswordReuseLookup_REQUEST_SUCCESS GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 3
	// Unable to get a valid response from the SafeBrowsing service.
	GaiaPasswordReuse_PasswordReuseLookup_REQUEST_FAILURE GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 4
	// We won't be able to compute reputation for the URL e.g. local IP
	// address, localhost, not-yet-assigned by ICANN gTLD, etc.
	GaiaPasswordReuse_PasswordReuseLookup_URL_UNSUPPORTED GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 5
	// URL did match enterprise whitelist.
	// No further action required related to this re-use event.
	GaiaPasswordReuse_PasswordReuseLookup_ENTERPRISE_WHITELIST_HIT GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 6
	// Password reuse lookup is turned off by enterprise policy.
	// No further action required related to this re-use event.
	GaiaPasswordReuse_PasswordReuseLookup_TURNED_OFF_BY_POLICY GaiaPasswordReuse_PasswordReuseLookup_LookupResult = 7
)

var GaiaPasswordReuse_PasswordReuseLookup_LookupResult_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "WHITELIST_HIT",
	2: "CACHE_HIT",
	3: "REQUEST_SUCCESS",
	4: "REQUEST_FAILURE",
	5: "URL_UNSUPPORTED",
	6: "ENTERPRISE_WHITELIST_HIT",
	7: "TURNED_OFF_BY_POLICY",
}
var GaiaPasswordReuse_PasswordReuseLookup_LookupResult_value = map[string]int32{
	"UNSPECIFIED":              0,
	"WHITELIST_HIT":            1,
	"CACHE_HIT":                2,
	"REQUEST_SUCCESS":          3,
	"REQUEST_FAILURE":          4,
	"URL_UNSUPPORTED":          5,
	"ENTERPRISE_WHITELIST_HIT": 6,
	"TURNED_OFF_BY_POLICY":     7,
}

func (x GaiaPasswordReuse_PasswordReuseLookup_LookupResult) Enum() *GaiaPasswordReuse_PasswordReuseLookup_LookupResult {
	p := new(GaiaPasswordReuse_PasswordReuseLookup_LookupResult)
	*p = x
	return p
}
func (x GaiaPasswordReuse_PasswordReuseLookup_LookupResult) String() string {
	return proto.EnumName(GaiaPasswordReuse_PasswordReuseLookup_LookupResult_name, int32(x))
}
func (x *GaiaPasswordReuse_PasswordReuseLookup_LookupResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GaiaPasswordReuse_PasswordReuseLookup_LookupResult_value, data, "GaiaPasswordReuse_PasswordReuseLookup_LookupResult")
	if err != nil {
		return err
	}
	*x = GaiaPasswordReuse_PasswordReuseLookup_LookupResult(value)
	return nil
}
func (GaiaPasswordReuse_PasswordReuseLookup_LookupResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 1, 0}
}

// The following two are only present for CACHE_HIT and REQUEST_SUCCESS.
// The verdict received from the Reputation service. This is set only
// if the user has SafeBrowsing enabled and we fetch the verdict from the
// cache or by sending a verdict request.
type GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict int32

const (
	GaiaPasswordReuse_PasswordReuseLookup_VERDICT_UNSPECIFIED GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict = 0
	GaiaPasswordReuse_PasswordReuseLookup_SAFE                GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict = 1
	GaiaPasswordReuse_PasswordReuseLookup_LOW_REPUTATION      GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict = 2
	GaiaPasswordReuse_PasswordReuseLookup_PHISHING            GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict = 3
)

var GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict_name = map[int32]string{
	0: "VERDICT_UNSPECIFIED",
	1: "SAFE",
	2: "LOW_REPUTATION",
	3: "PHISHING",
}
var GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict_value = map[string]int32{
	"VERDICT_UNSPECIFIED": 0,
	"SAFE":                1,
	"LOW_REPUTATION":      2,
	"PHISHING":            3,
}

func (x GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict) Enum() *GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict {
	p := new(GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict)
	*p = x
	return p
}
func (x GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict) String() string {
	return proto.EnumName(GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict_name, int32(x))
}
func (x *GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict_value, data, "GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict")
	if err != nil {
		return err
	}
	*x = GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict(value)
	return nil
}
func (GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 1, 1}
}

type GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult int32

const (
	GaiaPasswordReuse_PasswordReuseDialogInteraction_UNSPECIFIED GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult = 0
	// The user took the action suggested by the warning prompt.
	GaiaPasswordReuse_PasswordReuseDialogInteraction_WARNING_ACTION_TAKEN GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult = 1
	// The user clicked ignore in the warning prompt.
	GaiaPasswordReuse_PasswordReuseDialogInteraction_WARNING_ACTION_IGNORED GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult = 2
	// The warning UI was ignored, i.e. not interacted with by the user.
	// This could happen if the user navigates away from the page.
	GaiaPasswordReuse_PasswordReuseDialogInteraction_WARNING_UI_IGNORED GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult = 3
	// The user clicked "Change Password" on chrome://settings page.
	GaiaPasswordReuse_PasswordReuseDialogInteraction_WARNING_ACTION_TAKEN_ON_SETTINGS GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult = 4
)

var GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "WARNING_ACTION_TAKEN",
	2: "WARNING_ACTION_IGNORED",
	3: "WARNING_UI_IGNORED",
	4: "WARNING_ACTION_TAKEN_ON_SETTINGS",
}
var GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult_value = map[string]int32{
	"UNSPECIFIED":                      0,
	"WARNING_ACTION_TAKEN":             1,
	"WARNING_ACTION_IGNORED":           2,
	"WARNING_UI_IGNORED":               3,
	"WARNING_ACTION_TAKEN_ON_SETTINGS": 4,
}

func (x GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult) Enum() *GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult {
	p := new(GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult)
	*p = x
	return p
}
func (x GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult) String() string {
	return proto.EnumName(GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult_name, int32(x))
}
func (x *GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult_value, data, "GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult")
	if err != nil {
		return err
	}
	*x = GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult(value)
	return nil
}
func (GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 2, 0}
}

type GaiaPasswordReuse_PasswordCaptured_EventTrigger int32

const (
	GaiaPasswordReuse_PasswordCaptured_UNSPECIFIED GaiaPasswordReuse_PasswordCaptured_EventTrigger = 0
	// Event added because user logged in.
	GaiaPasswordReuse_PasswordCaptured_USER_LOGGED_IN GaiaPasswordReuse_PasswordCaptured_EventTrigger = 1
	// Event added because 28d timer fired.
	GaiaPasswordReuse_PasswordCaptured_EXPIRED_28D_TIMER GaiaPasswordReuse_PasswordCaptured_EventTrigger = 2
)

var GaiaPasswordReuse_PasswordCaptured_EventTrigger_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "USER_LOGGED_IN",
	2: "EXPIRED_28D_TIMER",
}
var GaiaPasswordReuse_PasswordCaptured_EventTrigger_value = map[string]int32{
	"UNSPECIFIED":       0,
	"USER_LOGGED_IN":    1,
	"EXPIRED_28D_TIMER": 2,
}

func (x GaiaPasswordReuse_PasswordCaptured_EventTrigger) Enum() *GaiaPasswordReuse_PasswordCaptured_EventTrigger {
	p := new(GaiaPasswordReuse_PasswordCaptured_EventTrigger)
	*p = x
	return p
}
func (x GaiaPasswordReuse_PasswordCaptured_EventTrigger) String() string {
	return proto.EnumName(GaiaPasswordReuse_PasswordCaptured_EventTrigger_name, int32(x))
}
func (x *GaiaPasswordReuse_PasswordCaptured_EventTrigger) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GaiaPasswordReuse_PasswordCaptured_EventTrigger_value, data, "GaiaPasswordReuse_PasswordCaptured_EventTrigger")
	if err != nil {
		return err
	}
	*x = GaiaPasswordReuse_PasswordCaptured_EventTrigger(value)
	return nil
}
func (GaiaPasswordReuse_PasswordCaptured_EventTrigger) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 3, 0}
}

// User reused their GAIA password on another website.
type GaiaPasswordReuse struct {
	ReuseDetected *GaiaPasswordReuse_PasswordReuseDetected `protobuf:"bytes,1,opt,name=reuse_detected,json=reuseDetected" json:"reuse_detected,omitempty"`
	// Logged when we try to detect whether the password was reused on a
	// Phishing or a Low-reputation site.
	ReuseLookup          *GaiaPasswordReuse_PasswordReuseLookup            `protobuf:"bytes,2,opt,name=reuse_lookup,json=reuseLookup" json:"reuse_lookup,omitempty"`
	DialogInteraction    *GaiaPasswordReuse_PasswordReuseDialogInteraction `protobuf:"bytes,3,opt,name=dialog_interaction,json=dialogInteraction" json:"dialog_interaction,omitempty"`
	PasswordCaptured     *GaiaPasswordReuse_PasswordCaptured               `protobuf:"bytes,4,opt,name=password_captured,json=passwordCaptured" json:"password_captured,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *GaiaPasswordReuse) Reset()         { *m = GaiaPasswordReuse{} }
func (m *GaiaPasswordReuse) String() string { return proto.CompactTextString(m) }
func (*GaiaPasswordReuse) ProtoMessage()    {}
func (*GaiaPasswordReuse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0}
}
func (m *GaiaPasswordReuse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaiaPasswordReuse.Unmarshal(m, b)
}
func (m *GaiaPasswordReuse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaiaPasswordReuse.Marshal(b, m, deterministic)
}
func (dst *GaiaPasswordReuse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaiaPasswordReuse.Merge(dst, src)
}
func (m *GaiaPasswordReuse) XXX_Size() int {
	return xxx_messageInfo_GaiaPasswordReuse.Size(m)
}
func (m *GaiaPasswordReuse) XXX_DiscardUnknown() {
	xxx_messageInfo_GaiaPasswordReuse.DiscardUnknown(m)
}

var xxx_messageInfo_GaiaPasswordReuse proto.InternalMessageInfo

func (m *GaiaPasswordReuse) GetReuseDetected() *GaiaPasswordReuse_PasswordReuseDetected {
	if m != nil {
		return m.ReuseDetected
	}
	return nil
}

func (m *GaiaPasswordReuse) GetReuseLookup() *GaiaPasswordReuse_PasswordReuseLookup {
	if m != nil {
		return m.ReuseLookup
	}
	return nil
}

func (m *GaiaPasswordReuse) GetDialogInteraction() *GaiaPasswordReuse_PasswordReuseDialogInteraction {
	if m != nil {
		return m.DialogInteraction
	}
	return nil
}

// Deprecated: Do not use.
func (m *GaiaPasswordReuse) GetPasswordCaptured() *GaiaPasswordReuse_PasswordCaptured {
	if m != nil {
		return m.PasswordCaptured
	}
	return nil
}

// Logged when we detect a password re-use event on a non-GAIA site.
// If the user hasn’t enabled SafeBrowsing, this will be the last event.
type GaiaPasswordReuse_PasswordReuseDetected struct {
	Status               *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *GaiaPasswordReuse_PasswordReuseDetected) Reset() {
	*m = GaiaPasswordReuse_PasswordReuseDetected{}
}
func (m *GaiaPasswordReuse_PasswordReuseDetected) String() string { return proto.CompactTextString(m) }
func (*GaiaPasswordReuse_PasswordReuseDetected) ProtoMessage()    {}
func (*GaiaPasswordReuse_PasswordReuseDetected) Descriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 0}
}
func (m *GaiaPasswordReuse_PasswordReuseDetected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected.Unmarshal(m, b)
}
func (m *GaiaPasswordReuse_PasswordReuseDetected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected.Marshal(b, m, deterministic)
}
func (dst *GaiaPasswordReuse_PasswordReuseDetected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected.Merge(dst, src)
}
func (m *GaiaPasswordReuse_PasswordReuseDetected) XXX_Size() int {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected.Size(m)
}
func (m *GaiaPasswordReuse_PasswordReuseDetected) XXX_DiscardUnknown() {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected.DiscardUnknown(m)
}

var xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected proto.InternalMessageInfo

func (m *GaiaPasswordReuse_PasswordReuseDetected) GetStatus() *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus struct {
	// Is SafeBrowsing enabled?
	Enabled                         *bool                                                                           `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	SafeBrowsingReportingPopulation *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation `protobuf:"varint,2,opt,name=safe_browsing_reporting_population,json=safeBrowsingReportingPopulation,enum=sync_pb.GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation" json:"safe_browsing_reporting_population,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                                                                        `json:"-"`
	XXX_unrecognized                []byte                                                                          `json:"-"`
	XXX_sizecache                   int32                                                                           `json:"-"`
}

func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) Reset() {
	*m = GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus{}
}
func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) String() string {
	return proto.CompactTextString(m)
}
func (*GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) ProtoMessage() {}
func (*GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 0, 0}
}
func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus.Unmarshal(m, b)
}
func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus.Marshal(b, m, deterministic)
}
func (dst *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus.Merge(dst, src)
}
func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) XXX_Size() int {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus.Size(m)
}
func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus.DiscardUnknown(m)
}

var xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus proto.InternalMessageInfo

func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

func (m *GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus) GetSafeBrowsingReportingPopulation() GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation {
	if m != nil && m.SafeBrowsingReportingPopulation != nil {
		return *m.SafeBrowsingReportingPopulation
	}
	return GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_REPORTING_POPULATION_UNSPECIFIED
}

type GaiaPasswordReuse_PasswordReuseLookup struct {
	LookupResult *GaiaPasswordReuse_PasswordReuseLookup_LookupResult      `protobuf:"varint,1,opt,name=lookup_result,json=lookupResult,enum=sync_pb.GaiaPasswordReuse_PasswordReuseLookup_LookupResult" json:"lookup_result,omitempty"`
	Verdict      *GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict `protobuf:"varint,2,opt,name=verdict,enum=sync_pb.GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict" json:"verdict,omitempty"`
	// PhishGuard token that identifies the verdict on the server.
	VerdictToken         []byte   `protobuf:"bytes,3,opt,name=verdict_token,json=verdictToken" json:"verdict_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GaiaPasswordReuse_PasswordReuseLookup) Reset()         { *m = GaiaPasswordReuse_PasswordReuseLookup{} }
func (m *GaiaPasswordReuse_PasswordReuseLookup) String() string { return proto.CompactTextString(m) }
func (*GaiaPasswordReuse_PasswordReuseLookup) ProtoMessage()    {}
func (*GaiaPasswordReuse_PasswordReuseLookup) Descriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 1}
}
func (m *GaiaPasswordReuse_PasswordReuseLookup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseLookup.Unmarshal(m, b)
}
func (m *GaiaPasswordReuse_PasswordReuseLookup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseLookup.Marshal(b, m, deterministic)
}
func (dst *GaiaPasswordReuse_PasswordReuseLookup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseLookup.Merge(dst, src)
}
func (m *GaiaPasswordReuse_PasswordReuseLookup) XXX_Size() int {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseLookup.Size(m)
}
func (m *GaiaPasswordReuse_PasswordReuseLookup) XXX_DiscardUnknown() {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseLookup.DiscardUnknown(m)
}

var xxx_messageInfo_GaiaPasswordReuse_PasswordReuseLookup proto.InternalMessageInfo

func (m *GaiaPasswordReuse_PasswordReuseLookup) GetLookupResult() GaiaPasswordReuse_PasswordReuseLookup_LookupResult {
	if m != nil && m.LookupResult != nil {
		return *m.LookupResult
	}
	return GaiaPasswordReuse_PasswordReuseLookup_UNSPECIFIED
}

func (m *GaiaPasswordReuse_PasswordReuseLookup) GetVerdict() GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict {
	if m != nil && m.Verdict != nil {
		return *m.Verdict
	}
	return GaiaPasswordReuse_PasswordReuseLookup_VERDICT_UNSPECIFIED
}

func (m *GaiaPasswordReuse_PasswordReuseLookup) GetVerdictToken() []byte {
	if m != nil {
		return m.VerdictToken
	}
	return nil
}

// Logged when the user interacts with the warning UI shown to encourage
// password change if the site is Phishing or Low-reputation.
type GaiaPasswordReuse_PasswordReuseDialogInteraction struct {
	InteractionResult    *GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult `protobuf:"varint,1,opt,name=interaction_result,json=interactionResult,enum=sync_pb.GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult" json:"interaction_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                            `json:"-"`
	XXX_unrecognized     []byte                                                              `json:"-"`
	XXX_sizecache        int32                                                               `json:"-"`
}

func (m *GaiaPasswordReuse_PasswordReuseDialogInteraction) Reset() {
	*m = GaiaPasswordReuse_PasswordReuseDialogInteraction{}
}
func (m *GaiaPasswordReuse_PasswordReuseDialogInteraction) String() string {
	return proto.CompactTextString(m)
}
func (*GaiaPasswordReuse_PasswordReuseDialogInteraction) ProtoMessage() {}
func (*GaiaPasswordReuse_PasswordReuseDialogInteraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 2}
}
func (m *GaiaPasswordReuse_PasswordReuseDialogInteraction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDialogInteraction.Unmarshal(m, b)
}
func (m *GaiaPasswordReuse_PasswordReuseDialogInteraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDialogInteraction.Marshal(b, m, deterministic)
}
func (dst *GaiaPasswordReuse_PasswordReuseDialogInteraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDialogInteraction.Merge(dst, src)
}
func (m *GaiaPasswordReuse_PasswordReuseDialogInteraction) XXX_Size() int {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDialogInteraction.Size(m)
}
func (m *GaiaPasswordReuse_PasswordReuseDialogInteraction) XXX_DiscardUnknown() {
	xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDialogInteraction.DiscardUnknown(m)
}

var xxx_messageInfo_GaiaPasswordReuse_PasswordReuseDialogInteraction proto.InternalMessageInfo

func (m *GaiaPasswordReuse_PasswordReuseDialogInteraction) GetInteractionResult() GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult {
	if m != nil && m.InteractionResult != nil {
		return *m.InteractionResult
	}
	return GaiaPasswordReuse_PasswordReuseDialogInteraction_UNSPECIFIED
}

// TODO(markusheintz): Remove
// DEPRECATED: DO NOT USE!
// Logged when the user logs into Google, and at least once per 28d.
type GaiaPasswordReuse_PasswordCaptured struct {
	EventTrigger         *GaiaPasswordReuse_PasswordCaptured_EventTrigger `protobuf:"varint,1,opt,name=event_trigger,json=eventTrigger,enum=sync_pb.GaiaPasswordReuse_PasswordCaptured_EventTrigger" json:"event_trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *GaiaPasswordReuse_PasswordCaptured) Reset()         { *m = GaiaPasswordReuse_PasswordCaptured{} }
func (m *GaiaPasswordReuse_PasswordCaptured) String() string { return proto.CompactTextString(m) }
func (*GaiaPasswordReuse_PasswordCaptured) ProtoMessage()    {}
func (*GaiaPasswordReuse_PasswordCaptured) Descriptor() ([]byte, []int) {
	return fileDescriptor_gaia_password_reuse_428ce931026fa3f6, []int{0, 3}
}
func (m *GaiaPasswordReuse_PasswordCaptured) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordCaptured.Unmarshal(m, b)
}
func (m *GaiaPasswordReuse_PasswordCaptured) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordCaptured.Marshal(b, m, deterministic)
}
func (dst *GaiaPasswordReuse_PasswordCaptured) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaiaPasswordReuse_PasswordCaptured.Merge(dst, src)
}
func (m *GaiaPasswordReuse_PasswordCaptured) XXX_Size() int {
	return xxx_messageInfo_GaiaPasswordReuse_PasswordCaptured.Size(m)
}
func (m *GaiaPasswordReuse_PasswordCaptured) XXX_DiscardUnknown() {
	xxx_messageInfo_GaiaPasswordReuse_PasswordCaptured.DiscardUnknown(m)
}

var xxx_messageInfo_GaiaPasswordReuse_PasswordCaptured proto.InternalMessageInfo

func (m *GaiaPasswordReuse_PasswordCaptured) GetEventTrigger() GaiaPasswordReuse_PasswordCaptured_EventTrigger {
	if m != nil && m.EventTrigger != nil {
		return *m.EventTrigger
	}
	return GaiaPasswordReuse_PasswordCaptured_UNSPECIFIED
}

func init() {
	proto.RegisterType((*GaiaPasswordReuse)(nil), "sync_pb.GaiaPasswordReuse")
	proto.RegisterType((*GaiaPasswordReuse_PasswordReuseDetected)(nil), "sync_pb.GaiaPasswordReuse.PasswordReuseDetected")
	proto.RegisterType((*GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus)(nil), "sync_pb.GaiaPasswordReuse.PasswordReuseDetected.SafeBrowsingStatus")
	proto.RegisterType((*GaiaPasswordReuse_PasswordReuseLookup)(nil), "sync_pb.GaiaPasswordReuse.PasswordReuseLookup")
	proto.RegisterType((*GaiaPasswordReuse_PasswordReuseDialogInteraction)(nil), "sync_pb.GaiaPasswordReuse.PasswordReuseDialogInteraction")
	proto.RegisterType((*GaiaPasswordReuse_PasswordCaptured)(nil), "sync_pb.GaiaPasswordReuse.PasswordCaptured")
	proto.RegisterEnum("sync_pb.GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation", GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation_name, GaiaPasswordReuse_PasswordReuseDetected_SafeBrowsingStatus_ReportingPopulation_value)
	proto.RegisterEnum("sync_pb.GaiaPasswordReuse_PasswordReuseLookup_LookupResult", GaiaPasswordReuse_PasswordReuseLookup_LookupResult_name, GaiaPasswordReuse_PasswordReuseLookup_LookupResult_value)
	proto.RegisterEnum("sync_pb.GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict", GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict_name, GaiaPasswordReuse_PasswordReuseLookup_ReputationVerdict_value)
	proto.RegisterEnum("sync_pb.GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult", GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult_name, GaiaPasswordReuse_PasswordReuseDialogInteraction_InteractionResult_value)
	proto.RegisterEnum("sync_pb.GaiaPasswordReuse_PasswordCaptured_EventTrigger", GaiaPasswordReuse_PasswordCaptured_EventTrigger_name, GaiaPasswordReuse_PasswordCaptured_EventTrigger_value)
}

func init() {
	proto.RegisterFile("components/sync/protocol/gaia_password_reuse.proto", fileDescriptor_gaia_password_reuse_428ce931026fa3f6)
}

var fileDescriptor_gaia_password_reuse_428ce931026fa3f6 = []byte{
	// 852 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0x4e, 0xda, 0xdd, 0x9e, 0x4d, 0x52, 0x67, 0xb6, 0x2d, 0x56, 0x84, 0xa0, 0x0a, 0x20,
	0x21, 0x55, 0xf2, 0xa2, 0x5c, 0x15, 0x71, 0x43, 0xd6, 0x9e, 0x24, 0x43, 0x83, 0xed, 0x8e, 0xc7,
	0xcd, 0x96, 0x1f, 0x0d, 0xde, 0x64, 0x9a, 0xb5, 0x9a, 0x8d, 0x2d, 0xff, 0xb4, 0x82, 0x77, 0xe0,
	0x9e, 0x07, 0xe0, 0x19, 0x80, 0x1b, 0xc4, 0x3d, 0x4f, 0x85, 0xc6, 0x71, 0x76, 0xbd, 0x49, 0x84,
	0x36, 0xa8, 0x77, 0x3e, 0x7f, 0xdf, 0x8c, 0xbf, 0xf3, 0x9d, 0x33, 0xd0, 0x9b, 0x46, 0x97, 0x71,
	0xb4, 0x14, 0xcb, 0x2c, 0x3d, 0x49, 0x7f, 0x5a, 0x4e, 0x4f, 0xe2, 0x24, 0xca, 0xa2, 0x69, 0xb4,
	0x38, 0x99, 0x07, 0x61, 0xc0, 0xe3, 0x20, 0x4d, 0xdf, 0x46, 0xc9, 0x8c, 0x27, 0x22, 0x4f, 0x85,
	0x51, 0x04, 0xd1, 0x81, 0x4c, 0xe4, 0xf1, 0x79, 0xf7, 0x9f, 0x16, 0xb4, 0x87, 0x41, 0x18, 0xb8,
	0x65, 0x16, 0x95, 0x49, 0x68, 0x02, 0xad, 0x22, 0x9b, 0xcf, 0x44, 0x26, 0xa6, 0x99, 0x98, 0xe9,
	0xca, 0x63, 0xe5, 0xb3, 0xa3, 0xde, 0xe7, 0x46, 0x59, 0x67, 0x6c, 0xd5, 0x18, 0x37, 0x2c, 0xab,
	0xac, 0xa3, 0xcd, 0xa4, 0x6a, 0xa2, 0xe7, 0xd0, 0x58, 0x01, 0x2f, 0xa2, 0xe8, 0x75, 0x1e, 0xeb,
	0x6a, 0x01, 0x6b, 0xdc, 0x16, 0x76, 0x5c, 0x54, 0xd1, 0xa3, 0xe4, 0xda, 0x40, 0x17, 0x80, 0x66,
	0x61, 0xb0, 0x88, 0xe6, 0x3c, 0x5c, 0x66, 0x22, 0x09, 0xa6, 0x59, 0x18, 0x2d, 0xf5, 0x5a, 0x01,
	0xfc, 0xc5, 0xad, 0xef, 0x5b, 0x20, 0x90, 0x6b, 0x00, 0xda, 0x9e, 0x6d, 0xba, 0xd0, 0xf7, 0xd0,
	0xbe, 0x22, 0x73, 0x1a, 0xc4, 0x59, 0x9e, 0x88, 0x99, 0x5e, 0x2f, 0x0e, 0x7a, 0x72, 0x8b, 0x83,
	0xcc, 0xb2, 0xe4, 0x54, 0xd5, 0x15, 0xaa, 0xc5, 0x1b, 0xde, 0xce, 0xef, 0x35, 0x78, 0xb8, 0x93,
	0x43, 0xf4, 0x1d, 0xdc, 0x4d, 0xb3, 0x20, 0xcb, 0xd3, 0xb2, 0x0b, 0xe6, 0xbe, 0x5d, 0x30, 0xbc,
	0xe0, 0x95, 0x38, 0x4d, 0xa2, 0xb7, 0x69, 0xb8, 0x9c, 0x7b, 0x05, 0x14, 0x2d, 0x21, 0x3b, 0x7f,
	0xab, 0x80, 0xb6, 0xc3, 0x48, 0x87, 0x03, 0xb1, 0x0c, 0xce, 0x17, 0x65, 0xeb, 0x0f, 0xe9, 0xda,
	0x44, 0xbf, 0x29, 0xd0, 0x4d, 0x83, 0x57, 0x82, 0x9f, 0x97, 0x15, 0x3c, 0x11, 0x71, 0x94, 0x64,
	0xf2, 0x2b, 0x8e, 0xe2, 0x7c, 0x11, 0x14, 0x0d, 0x90, 0x9d, 0x6d, 0xf5, 0x26, 0xef, 0xe0, 0xaa,
	0x06, 0x5d, 0xe3, 0xbb, 0x57, 0xf0, 0xf4, 0xa3, 0xb4, 0x92, 0xb7, 0x23, 0xa1, 0x7b, 0x01, 0xc7,
	0x3b, 0xdc, 0xe8, 0x13, 0x78, 0x4c, 0xb1, 0xeb, 0x50, 0x46, 0xec, 0x21, 0x77, 0x1d, 0xd7, 0x1f,
	0xf7, 0x19, 0x71, 0x6c, 0xee, 0xdb, 0x9e, 0x8b, 0x4d, 0x32, 0x20, 0xd8, 0xd2, 0xde, 0x43, 0x87,
	0x50, 0xb7, 0x1d, 0x1b, 0x6b, 0x0a, 0x7a, 0x04, 0x08, 0x9f, 0x31, 0x6c, 0x5b, 0xd8, 0xe2, 0x57,
	0x85, 0x9a, 0x8a, 0xee, 0xc1, 0x1d, 0xcf, 0x74, 0x7c, 0xa6, 0xd5, 0x3a, 0xbf, 0xd4, 0xe1, 0x78,
	0x87, 0x4a, 0xd1, 0x8f, 0xd0, 0x5c, 0xa9, 0x9c, 0x27, 0x22, 0xcd, 0x17, 0x59, 0x41, 0x64, 0xab,
	0xf7, 0xe5, 0x7e, 0x62, 0x37, 0x4a, 0xcd, 0x17, 0x10, 0xb4, 0xb1, 0xa8, 0x58, 0xe8, 0x5b, 0x38,
	0x78, 0x23, 0x92, 0x59, 0x38, 0xcd, 0x4a, 0xba, 0xbf, 0xda, 0x13, 0x9b, 0x8a, 0x38, 0xcf, 0x0a,
	0x62, 0x5e, 0xac, 0x70, 0xe8, 0x1a, 0x10, 0x7d, 0x0c, 0xcd, 0xf2, 0x93, 0x67, 0xd1, 0x6b, 0xb1,
	0x9a, 0xa8, 0x06, 0x6d, 0x94, 0x4e, 0x26, 0x7d, 0xdd, 0x3f, 0x15, 0x68, 0x54, 0xef, 0x87, 0xee,
	0xc3, 0xd1, 0x4d, 0x26, 0xdb, 0xd0, 0x9c, 0x8c, 0x08, 0xc3, 0x63, 0xe2, 0x31, 0x3e, 0x22, 0x4c,
	0x53, 0x50, 0x13, 0xee, 0x99, 0x7d, 0x73, 0x84, 0x0b, 0x53, 0x45, 0xc7, 0x70, 0x9f, 0xe2, 0xe7,
	0x3e, 0xf6, 0x18, 0xf7, 0x7c, 0xd3, 0xc4, 0x9e, 0xa7, 0xd5, 0xaa, 0xce, 0x41, 0x9f, 0x8c, 0x7d,
	0x8a, 0xb5, 0xba, 0x74, 0xfa, 0x74, 0x2c, 0x5b, 0xe5, 0xbb, 0xb2, 0x13, 0xd8, 0xd2, 0xee, 0xa0,
	0x0f, 0x40, 0xc7, 0x36, 0xc3, 0xd4, 0xa5, 0xc4, 0xc3, 0xfc, 0xe6, 0x59, 0x77, 0x91, 0x0e, 0x0f,
	0x98, 0x4f, 0x6d, 0x6c, 0x71, 0x67, 0x30, 0xe0, 0xa7, 0x2f, 0xb9, 0xeb, 0x8c, 0x89, 0xf9, 0x52,
	0x3b, 0xe8, 0x9e, 0x41, 0x7b, 0xeb, 0xef, 0xd1, 0xfb, 0x70, 0xfc, 0x02, 0x53, 0x8b, 0x98, 0x6c,
	0x5b, 0x10, 0x5e, 0x7f, 0x20, 0x05, 0x81, 0xa0, 0x35, 0x76, 0x26, 0x52, 0x0b, 0x3e, 0x2b, 0xa4,
	0xa3, 0xa9, 0xa8, 0x01, 0x87, 0xee, 0x88, 0x78, 0x23, 0x29, 0x8d, 0x5a, 0xe7, 0x0f, 0x15, 0x3e,
	0xfc, 0xef, 0xe5, 0x82, 0x7e, 0x06, 0x54, 0x59, 0x56, 0x37, 0xf5, 0xf1, 0xec, 0x7f, 0xef, 0x2c,
	0xa3, 0xba, 0xbf, 0x56, 0x7a, 0x69, 0x87, 0x9b, 0xae, 0xee, 0xaf, 0x0a, 0xb4, 0xb7, 0x12, 0xb7,
	0x1b, 0xa7, 0xc3, 0x83, 0x49, 0x9f, 0xda, 0x72, 0x4c, 0xfa, 0x66, 0x31, 0x22, 0xac, 0xff, 0x0c,
	0xdb, 0x9a, 0x82, 0x3a, 0xf0, 0x68, 0x23, 0x42, 0x86, 0xb6, 0x43, 0xb1, 0xa5, 0xa9, 0x72, 0x5c,
	0xd6, 0x31, 0x9f, 0x5c, 0xf9, 0x6b, 0x72, 0xec, 0x76, 0xa1, 0x71, 0xc7, 0xe6, 0x1e, 0x66, 0x72,
	0xa6, 0x3c, 0xad, 0xde, 0xf9, 0x4b, 0x01, 0x6d, 0x73, 0x5b, 0xa2, 0x1f, 0xa0, 0x29, 0xde, 0x88,
	0x65, 0xc6, 0xb3, 0x24, 0x9c, 0xcf, 0x45, 0x52, 0xd2, 0xf4, 0x74, 0x8f, 0x8d, 0x6b, 0x60, 0x09,
	0xc0, 0x56, 0xf5, 0xb4, 0x21, 0x2a, 0x56, 0xf7, 0x6b, 0x68, 0x54, 0xa3, 0xdb, 0x44, 0x20, 0x68,
	0xf9, 0x1e, 0xa6, 0x7c, 0xec, 0x0c, 0x87, 0xd8, 0xe2, 0x44, 0x52, 0xf0, 0x10, 0xda, 0xf8, 0xcc,
	0x25, 0x14, 0x5b, 0xbc, 0xf7, 0xd4, 0xe2, 0x8c, 0x7c, 0x83, 0xa9, 0xa6, 0x9e, 0x3e, 0x81, 0x4f,
	0xa3, 0x64, 0x6e, 0x4c, 0x2f, 0x92, 0xe8, 0x32, 0xcc, 0x2f, 0x8d, 0xeb, 0x87, 0xb9, 0xb8, 0xac,
	0xb1, 0x7e, 0x98, 0x47, 0x35, 0x57, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xab, 0xba, 0x25, 0x2c,
	0xb7, 0x07, 0x00, 0x00,
}
