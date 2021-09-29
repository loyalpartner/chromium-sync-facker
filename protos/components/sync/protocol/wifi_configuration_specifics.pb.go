// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/wifi_configuration_specifics.proto

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

type WifiConfigurationSpecifics_SecurityType int32

const (
	WifiConfigurationSpecifics_SECURITY_TYPE_UNSPECIFIED WifiConfigurationSpecifics_SecurityType = 0
	WifiConfigurationSpecifics_SECURITY_TYPE_NONE        WifiConfigurationSpecifics_SecurityType = 1
	WifiConfigurationSpecifics_SECURITY_TYPE_WEP         WifiConfigurationSpecifics_SecurityType = 2
	WifiConfigurationSpecifics_SECURITY_TYPE_PSK         WifiConfigurationSpecifics_SecurityType = 3
)

var WifiConfigurationSpecifics_SecurityType_name = map[int32]string{
	0: "SECURITY_TYPE_UNSPECIFIED",
	1: "SECURITY_TYPE_NONE",
	2: "SECURITY_TYPE_WEP",
	3: "SECURITY_TYPE_PSK",
}
var WifiConfigurationSpecifics_SecurityType_value = map[string]int32{
	"SECURITY_TYPE_UNSPECIFIED": 0,
	"SECURITY_TYPE_NONE":        1,
	"SECURITY_TYPE_WEP":         2,
	"SECURITY_TYPE_PSK":         3,
}

func (x WifiConfigurationSpecifics_SecurityType) Enum() *WifiConfigurationSpecifics_SecurityType {
	p := new(WifiConfigurationSpecifics_SecurityType)
	*p = x
	return p
}
func (x WifiConfigurationSpecifics_SecurityType) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_SecurityType_name, int32(x))
}
func (x *WifiConfigurationSpecifics_SecurityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_SecurityType_value, data, "WifiConfigurationSpecifics_SecurityType")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_SecurityType(value)
	return nil
}
func (WifiConfigurationSpecifics_SecurityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 0}
}

type WifiConfigurationSpecifics_AutomaticallyConnectOption int32

const (
	WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_UNSPECIFIED WifiConfigurationSpecifics_AutomaticallyConnectOption = 0
	WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_DISABLED    WifiConfigurationSpecifics_AutomaticallyConnectOption = 1
	WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_ENABLED     WifiConfigurationSpecifics_AutomaticallyConnectOption = 2
)

var WifiConfigurationSpecifics_AutomaticallyConnectOption_name = map[int32]string{
	0: "AUTOMATICALLY_CONNECT_UNSPECIFIED",
	1: "AUTOMATICALLY_CONNECT_DISABLED",
	2: "AUTOMATICALLY_CONNECT_ENABLED",
}
var WifiConfigurationSpecifics_AutomaticallyConnectOption_value = map[string]int32{
	"AUTOMATICALLY_CONNECT_UNSPECIFIED": 0,
	"AUTOMATICALLY_CONNECT_DISABLED":    1,
	"AUTOMATICALLY_CONNECT_ENABLED":     2,
}

func (x WifiConfigurationSpecifics_AutomaticallyConnectOption) Enum() *WifiConfigurationSpecifics_AutomaticallyConnectOption {
	p := new(WifiConfigurationSpecifics_AutomaticallyConnectOption)
	*p = x
	return p
}
func (x WifiConfigurationSpecifics_AutomaticallyConnectOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_AutomaticallyConnectOption_name, int32(x))
}
func (x *WifiConfigurationSpecifics_AutomaticallyConnectOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_AutomaticallyConnectOption_value, data, "WifiConfigurationSpecifics_AutomaticallyConnectOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_AutomaticallyConnectOption(value)
	return nil
}
func (WifiConfigurationSpecifics_AutomaticallyConnectOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 1}
}

type WifiConfigurationSpecifics_IsPreferredOption int32

const (
	WifiConfigurationSpecifics_IS_PREFERRED_UNSPECIFIED WifiConfigurationSpecifics_IsPreferredOption = 0
	WifiConfigurationSpecifics_IS_PREFERRED_DISABLED    WifiConfigurationSpecifics_IsPreferredOption = 1
	WifiConfigurationSpecifics_IS_PREFERRED_ENABLED     WifiConfigurationSpecifics_IsPreferredOption = 2
)

var WifiConfigurationSpecifics_IsPreferredOption_name = map[int32]string{
	0: "IS_PREFERRED_UNSPECIFIED",
	1: "IS_PREFERRED_DISABLED",
	2: "IS_PREFERRED_ENABLED",
}
var WifiConfigurationSpecifics_IsPreferredOption_value = map[string]int32{
	"IS_PREFERRED_UNSPECIFIED": 0,
	"IS_PREFERRED_DISABLED":    1,
	"IS_PREFERRED_ENABLED":     2,
}

func (x WifiConfigurationSpecifics_IsPreferredOption) Enum() *WifiConfigurationSpecifics_IsPreferredOption {
	p := new(WifiConfigurationSpecifics_IsPreferredOption)
	*p = x
	return p
}
func (x WifiConfigurationSpecifics_IsPreferredOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_IsPreferredOption_name, int32(x))
}
func (x *WifiConfigurationSpecifics_IsPreferredOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_IsPreferredOption_value, data, "WifiConfigurationSpecifics_IsPreferredOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_IsPreferredOption(value)
	return nil
}
func (WifiConfigurationSpecifics_IsPreferredOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 2}
}

type WifiConfigurationSpecifics_MeteredOption int32

const (
	WifiConfigurationSpecifics_METERED_OPTION_UNSPECIFIED WifiConfigurationSpecifics_MeteredOption = 0
	WifiConfigurationSpecifics_METERED_OPTION_NO          WifiConfigurationSpecifics_MeteredOption = 1
	WifiConfigurationSpecifics_METERED_OPTION_YES         WifiConfigurationSpecifics_MeteredOption = 2
	// Allows the device to use heuristics to determine if network is metered.
	WifiConfigurationSpecifics_METERED_OPTION_AUTO WifiConfigurationSpecifics_MeteredOption = 3
)

var WifiConfigurationSpecifics_MeteredOption_name = map[int32]string{
	0: "METERED_OPTION_UNSPECIFIED",
	1: "METERED_OPTION_NO",
	2: "METERED_OPTION_YES",
	3: "METERED_OPTION_AUTO",
}
var WifiConfigurationSpecifics_MeteredOption_value = map[string]int32{
	"METERED_OPTION_UNSPECIFIED": 0,
	"METERED_OPTION_NO":          1,
	"METERED_OPTION_YES":         2,
	"METERED_OPTION_AUTO":        3,
}

func (x WifiConfigurationSpecifics_MeteredOption) Enum() *WifiConfigurationSpecifics_MeteredOption {
	p := new(WifiConfigurationSpecifics_MeteredOption)
	*p = x
	return p
}
func (x WifiConfigurationSpecifics_MeteredOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_MeteredOption_name, int32(x))
}
func (x *WifiConfigurationSpecifics_MeteredOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_MeteredOption_value, data, "WifiConfigurationSpecifics_MeteredOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_MeteredOption(value)
	return nil
}
func (WifiConfigurationSpecifics_MeteredOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 3}
}

type WifiConfigurationSpecifics_DnsOption int32

const (
	WifiConfigurationSpecifics_DNS_OPTION_UNSPECIFIED  WifiConfigurationSpecifics_DnsOption = 0
	WifiConfigurationSpecifics_DNS_OPTION_DEFAULT_DHCP WifiConfigurationSpecifics_DnsOption = 1
	WifiConfigurationSpecifics_DNS_OPTION_CUSTOM       WifiConfigurationSpecifics_DnsOption = 2
)

var WifiConfigurationSpecifics_DnsOption_name = map[int32]string{
	0: "DNS_OPTION_UNSPECIFIED",
	1: "DNS_OPTION_DEFAULT_DHCP",
	2: "DNS_OPTION_CUSTOM",
}
var WifiConfigurationSpecifics_DnsOption_value = map[string]int32{
	"DNS_OPTION_UNSPECIFIED":  0,
	"DNS_OPTION_DEFAULT_DHCP": 1,
	"DNS_OPTION_CUSTOM":       2,
}

func (x WifiConfigurationSpecifics_DnsOption) Enum() *WifiConfigurationSpecifics_DnsOption {
	p := new(WifiConfigurationSpecifics_DnsOption)
	*p = x
	return p
}
func (x WifiConfigurationSpecifics_DnsOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_DnsOption_name, int32(x))
}
func (x *WifiConfigurationSpecifics_DnsOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_DnsOption_value, data, "WifiConfigurationSpecifics_DnsOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_DnsOption(value)
	return nil
}
func (WifiConfigurationSpecifics_DnsOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 4}
}

type WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption int32

const (
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_UNSPECIFIED WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 0
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_DISABLED    WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 1
	// Use a Proxy Auto-config(PAC) Url, set in proxy_url
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_AUTOMATIC WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 2
	// Uses Web Proxy Auto-Discovery Protocol (WPAD) to discover the proxy
	// settings using DHCP/DNS.
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_AUTODISCOVERY WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 3
	// User sets details in manual_proxy_configuration.
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_MANUAL WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 4
)

var WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_name = map[int32]string{
	0: "PROXY_OPTION_UNSPECIFIED",
	1: "PROXY_OPTION_DISABLED",
	2: "PROXY_OPTION_AUTOMATIC",
	3: "PROXY_OPTION_AUTODISCOVERY",
	4: "PROXY_OPTION_MANUAL",
}
var WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_value = map[string]int32{
	"PROXY_OPTION_UNSPECIFIED":   0,
	"PROXY_OPTION_DISABLED":      1,
	"PROXY_OPTION_AUTOMATIC":     2,
	"PROXY_OPTION_AUTODISCOVERY": 3,
	"PROXY_OPTION_MANUAL":        4,
}

func (x WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) Enum() *WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption {
	p := new(WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption)
	*p = x
	return p
}
func (x WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_name, int32(x))
}
func (x *WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_value, data, "WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption(value)
	return nil
}
func (WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 0, 0}
}

type WifiConfigurationSpecifics struct {
	// SSID encoded to hex, letters should be upper case and 0x prefix should be
	// omitted. For example, ssid "network" would be provided as "6E6574776F726B".
	HexSsid      []byte                                   `protobuf:"bytes,1,opt,name=hex_ssid,json=hexSsid" json:"hex_ssid,omitempty"`
	SecurityType *WifiConfigurationSpecifics_SecurityType `protobuf:"varint,2,opt,name=security_type,json=securityType,enum=sync_pb.WifiConfigurationSpecifics_SecurityType" json:"security_type,omitempty"`
	// The passphrase can be ASCII, UTF-8, or a string of hex digits.
	Passphrase           []byte                                                 `protobuf:"bytes,3,opt,name=passphrase" json:"passphrase,omitempty"`
	AutomaticallyConnect *WifiConfigurationSpecifics_AutomaticallyConnectOption `protobuf:"varint,4,opt,name=automatically_connect,json=automaticallyConnect,enum=sync_pb.WifiConfigurationSpecifics_AutomaticallyConnectOption" json:"automatically_connect,omitempty"`
	IsPreferred          *WifiConfigurationSpecifics_IsPreferredOption          `protobuf:"varint,5,opt,name=is_preferred,json=isPreferred,enum=sync_pb.WifiConfigurationSpecifics_IsPreferredOption" json:"is_preferred,omitempty"`
	Metered              *WifiConfigurationSpecifics_MeteredOption              `protobuf:"varint,6,opt,name=metered,enum=sync_pb.WifiConfigurationSpecifics_MeteredOption" json:"metered,omitempty"`
	ProxyConfiguration   *WifiConfigurationSpecifics_ProxyConfiguration         `protobuf:"bytes,7,opt,name=proxy_configuration,json=proxyConfiguration" json:"proxy_configuration,omitempty"`
	DnsOption            *WifiConfigurationSpecifics_DnsOption                  `protobuf:"varint,10,opt,name=dns_option,json=dnsOption,enum=sync_pb.WifiConfigurationSpecifics_DnsOption" json:"dns_option,omitempty"`
	// List of DNS servers to be used when set to DNS_OPTION_CUSTOM.  Up to 4.
	CustomDns []string `protobuf:"bytes,8,rep,name=custom_dns,json=customDns" json:"custom_dns,omitempty"`
	// The last time this configuration was connected to before being synced.  It
	// will only be updated when the configuration is changed. This is represented
	// with the UNIX timestamp, ms since epoch.
	LastConnectedTimestamp *int64   `protobuf:"varint,9,opt,name=last_connected_timestamp,json=lastConnectedTimestamp" json:"last_connected_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *WifiConfigurationSpecifics) Reset()         { *m = WifiConfigurationSpecifics{} }
func (m *WifiConfigurationSpecifics) String() string { return proto.CompactTextString(m) }
func (*WifiConfigurationSpecifics) ProtoMessage()    {}
func (*WifiConfigurationSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0}
}
func (m *WifiConfigurationSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecifics.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecifics.Marshal(b, m, deterministic)
}
func (dst *WifiConfigurationSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecifics.Merge(dst, src)
}
func (m *WifiConfigurationSpecifics) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecifics.Size(m)
}
func (m *WifiConfigurationSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecifics proto.InternalMessageInfo

func (m *WifiConfigurationSpecifics) GetHexSsid() []byte {
	if m != nil {
		return m.HexSsid
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetSecurityType() WifiConfigurationSpecifics_SecurityType {
	if m != nil && m.SecurityType != nil {
		return *m.SecurityType
	}
	return WifiConfigurationSpecifics_SECURITY_TYPE_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetPassphrase() []byte {
	if m != nil {
		return m.Passphrase
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetAutomaticallyConnect() WifiConfigurationSpecifics_AutomaticallyConnectOption {
	if m != nil && m.AutomaticallyConnect != nil {
		return *m.AutomaticallyConnect
	}
	return WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetIsPreferred() WifiConfigurationSpecifics_IsPreferredOption {
	if m != nil && m.IsPreferred != nil {
		return *m.IsPreferred
	}
	return WifiConfigurationSpecifics_IS_PREFERRED_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetMetered() WifiConfigurationSpecifics_MeteredOption {
	if m != nil && m.Metered != nil {
		return *m.Metered
	}
	return WifiConfigurationSpecifics_METERED_OPTION_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetProxyConfiguration() *WifiConfigurationSpecifics_ProxyConfiguration {
	if m != nil {
		return m.ProxyConfiguration
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetDnsOption() WifiConfigurationSpecifics_DnsOption {
	if m != nil && m.DnsOption != nil {
		return *m.DnsOption
	}
	return WifiConfigurationSpecifics_DNS_OPTION_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetCustomDns() []string {
	if m != nil {
		return m.CustomDns
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetLastConnectedTimestamp() int64 {
	if m != nil && m.LastConnectedTimestamp != nil {
		return *m.LastConnectedTimestamp
	}
	return 0
}

type WifiConfigurationSpecifics_ProxyConfiguration struct {
	ProxyOption *WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption `protobuf:"varint,1,opt,name=proxy_option,json=proxyOption,enum=sync_pb.WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption" json:"proxy_option,omitempty"`
	// Only set if PROXY_OPTION_AUTOMATIC.
	AutoconfigurationUrl *string `protobuf:"bytes,2,opt,name=autoconfiguration_url,json=autoconfigurationUrl" json:"autoconfiguration_url,omitempty"`
	// Only set if PROXY_OPTION_MANUAL.
	ManualProxyConfiguration *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration `protobuf:"bytes,3,opt,name=manual_proxy_configuration,json=manualProxyConfiguration" json:"manual_proxy_configuration,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                                                                `json:"-"`
	XXX_unrecognized         []byte                                                                  `json:"-"`
	XXX_sizecache            int32                                                                   `json:"-"`
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) Reset() {
	*m = WifiConfigurationSpecifics_ProxyConfiguration{}
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*WifiConfigurationSpecifics_ProxyConfiguration) ProtoMessage() {}
func (*WifiConfigurationSpecifics_ProxyConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 0}
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Marshal(b, m, deterministic)
}
func (dst *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Merge(dst, src)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Size(m)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration proto.InternalMessageInfo

func (m *WifiConfigurationSpecifics_ProxyConfiguration) GetProxyOption() WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption {
	if m != nil && m.ProxyOption != nil {
		return *m.ProxyOption
	}
	return WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) GetAutoconfigurationUrl() string {
	if m != nil && m.AutoconfigurationUrl != nil {
		return *m.AutoconfigurationUrl
	}
	return ""
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) GetManualProxyConfiguration() *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration {
	if m != nil {
		return m.ManualProxyConfiguration
	}
	return nil
}

type WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration struct {
	HttpProxyUrl         *string  `protobuf:"bytes,1,opt,name=http_proxy_url,json=httpProxyUrl" json:"http_proxy_url,omitempty"`
	HttpProxyPort        *int32   `protobuf:"varint,2,opt,name=http_proxy_port,json=httpProxyPort" json:"http_proxy_port,omitempty"`
	SecureHttpProxyUrl   *string  `protobuf:"bytes,3,opt,name=secure_http_proxy_url,json=secureHttpProxyUrl" json:"secure_http_proxy_url,omitempty"`
	SecureHttpProxyPort  *int32   `protobuf:"varint,4,opt,name=secure_http_proxy_port,json=secureHttpProxyPort" json:"secure_http_proxy_port,omitempty"`
	SocksHostUrl         *string  `protobuf:"bytes,5,opt,name=socks_host_url,json=socksHostUrl" json:"socks_host_url,omitempty"`
	SocksHostPort        *int32   `protobuf:"varint,6,opt,name=socks_host_port,json=socksHostPort" json:"socks_host_port,omitempty"`
	ExcludedDomains      []string `protobuf:"bytes,7,rep,name=excluded_domains,json=excludedDomains" json:"excluded_domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) Reset() {
	*m = WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration{}
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) ProtoMessage() {}
func (*WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_wifi_configuration_specifics_825016093e423dc6, []int{0, 0, 0}
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration.Marshal(b, m, deterministic)
}
func (dst *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration.Merge(dst, src)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration.Size(m)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration proto.InternalMessageInfo

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) GetHttpProxyUrl() string {
	if m != nil && m.HttpProxyUrl != nil {
		return *m.HttpProxyUrl
	}
	return ""
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) GetHttpProxyPort() int32 {
	if m != nil && m.HttpProxyPort != nil {
		return *m.HttpProxyPort
	}
	return 0
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) GetSecureHttpProxyUrl() string {
	if m != nil && m.SecureHttpProxyUrl != nil {
		return *m.SecureHttpProxyUrl
	}
	return ""
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) GetSecureHttpProxyPort() int32 {
	if m != nil && m.SecureHttpProxyPort != nil {
		return *m.SecureHttpProxyPort
	}
	return 0
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) GetSocksHostUrl() string {
	if m != nil && m.SocksHostUrl != nil {
		return *m.SocksHostUrl
	}
	return ""
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) GetSocksHostPort() int32 {
	if m != nil && m.SocksHostPort != nil {
		return *m.SocksHostPort
	}
	return 0
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration) GetExcludedDomains() []string {
	if m != nil {
		return m.ExcludedDomains
	}
	return nil
}

func init() {
	proto.RegisterType((*WifiConfigurationSpecifics)(nil), "sync_pb.WifiConfigurationSpecifics")
	proto.RegisterType((*WifiConfigurationSpecifics_ProxyConfiguration)(nil), "sync_pb.WifiConfigurationSpecifics.ProxyConfiguration")
	proto.RegisterType((*WifiConfigurationSpecifics_ProxyConfiguration_ManualProxyConfiguration)(nil), "sync_pb.WifiConfigurationSpecifics.ProxyConfiguration.ManualProxyConfiguration")
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_SecurityType", WifiConfigurationSpecifics_SecurityType_name, WifiConfigurationSpecifics_SecurityType_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_AutomaticallyConnectOption", WifiConfigurationSpecifics_AutomaticallyConnectOption_name, WifiConfigurationSpecifics_AutomaticallyConnectOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_IsPreferredOption", WifiConfigurationSpecifics_IsPreferredOption_name, WifiConfigurationSpecifics_IsPreferredOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_MeteredOption", WifiConfigurationSpecifics_MeteredOption_name, WifiConfigurationSpecifics_MeteredOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_DnsOption", WifiConfigurationSpecifics_DnsOption_name, WifiConfigurationSpecifics_DnsOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption", WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_name, WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_value)
}

func init() {
	proto.RegisterFile("components/sync/protocol/wifi_configuration_specifics.proto", fileDescriptor_wifi_configuration_specifics_825016093e423dc6)
}

var fileDescriptor_wifi_configuration_specifics_825016093e423dc6 = []byte{
	// 921 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xe1, 0x6e, 0x9b, 0x56,
	0x14, 0x1e, 0x76, 0x13, 0xd7, 0x27, 0x4e, 0x4a, 0x6e, 0xea, 0x94, 0x78, 0x4b, 0xe4, 0x5a, 0xed,
	0xe4, 0x69, 0x9a, 0xb3, 0xb6, 0xda, 0x34, 0x69, 0xd2, 0x24, 0x07, 0x88, 0x62, 0xd5, 0x06, 0x04,
	0x78, 0xad, 0xb5, 0x1f, 0x88, 0xc1, 0x75, 0x8c, 0x66, 0xb8, 0x88, 0x7b, 0xad, 0xc6, 0xaf, 0x30,
	0x69, 0x4f, 0xb0, 0x37, 0xd8, 0x8b, 0x4c, 0x7b, 0xab, 0x89, 0x8b, 0x49, 0xc0, 0x76, 0x24, 0xab,
	0xff, 0xe0, 0x7c, 0xdf, 0xf9, 0xbe, 0x73, 0xee, 0xb9, 0x1c, 0x01, 0x3f, 0x7b, 0x24, 0x8c, 0x49,
	0x84, 0x23, 0x46, 0x2f, 0xe9, 0x32, 0xf2, 0x2e, 0xe3, 0x84, 0x30, 0xe2, 0x91, 0xf9, 0xe5, 0xa7,
	0x60, 0x1a, 0x38, 0x1e, 0x89, 0xa6, 0xc1, 0xed, 0x22, 0x71, 0x59, 0x40, 0x22, 0x87, 0xc6, 0xd8,
	0x0b, 0xa6, 0x81, 0x47, 0x7b, 0x9c, 0x85, 0x6a, 0x69, 0x86, 0x13, 0xff, 0xde, 0xf9, 0x57, 0x84,
	0xd6, 0x87, 0x60, 0x1a, 0xc8, 0x45, 0xba, 0x95, 0xb3, 0xd1, 0x19, 0x3c, 0x9d, 0xe1, 0x3b, 0x87,
	0xd2, 0xc0, 0x97, 0x84, 0xb6, 0xd0, 0x6d, 0x98, 0xb5, 0x19, 0xbe, 0xb3, 0x68, 0xe0, 0xa3, 0x31,
	0x1c, 0x52, 0xec, 0x2d, 0x92, 0x80, 0x2d, 0x1d, 0xb6, 0x8c, 0xb1, 0x54, 0x69, 0x0b, 0xdd, 0xa3,
	0xb7, 0xdf, 0xf7, 0x56, 0xd2, 0xbd, 0xc7, 0x65, 0x7b, 0xd6, 0x2a, 0xd1, 0x5e, 0xc6, 0xd8, 0x6c,
	0xd0, 0xc2, 0x1b, 0xba, 0x00, 0x88, 0x5d, 0x4a, 0xe3, 0x59, 0xe2, 0x52, 0x2c, 0x55, 0xb9, 0x67,
	0x21, 0x82, 0x28, 0x34, 0xdd, 0x05, 0x23, 0xa1, 0xcb, 0x02, 0xcf, 0x9d, 0xcf, 0x97, 0x69, 0xa3,
	0x11, 0xf6, 0x98, 0xf4, 0x84, 0xdb, 0xff, 0xb2, 0x8b, 0x7d, 0xbf, 0x28, 0x20, 0x67, 0xf9, 0x7a,
	0x9c, 0x72, 0xcc, 0xe7, 0xee, 0x16, 0x0c, 0x7d, 0x84, 0x46, 0x40, 0x9d, 0x38, 0xc1, 0x53, 0x9c,
	0x24, 0xd8, 0x97, 0xf6, 0xb8, 0xd7, 0x0f, 0xbb, 0x78, 0x0d, 0xa8, 0x91, 0xa7, 0xad, 0x2c, 0x0e,
	0x82, 0x87, 0x10, 0x7a, 0x0f, 0xb5, 0x10, 0x33, 0x9c, 0x8a, 0xee, 0x73, 0xd1, 0x37, 0xbb, 0x88,
	0x8e, 0xb2, 0x94, 0x95, 0x60, 0xae, 0x80, 0x6e, 0xe1, 0x24, 0x4e, 0xc8, 0xdd, 0xb2, 0x3c, 0x7c,
	0xa9, 0xd6, 0x16, 0xba, 0x07, 0x6f, 0x7f, 0xdc, 0x45, 0xd8, 0x48, 0xd3, 0x4b, 0x98, 0x89, 0xe2,
	0x8d, 0x18, 0x1a, 0x02, 0xf8, 0x11, 0x75, 0x08, 0xf7, 0x97, 0x80, 0x17, 0xfe, 0xdd, 0x2e, 0xfa,
	0x4a, 0x44, 0x57, 0x45, 0xd7, 0xfd, 0xfc, 0x11, 0x9d, 0x03, 0x78, 0x0b, 0xca, 0x48, 0xe8, 0xf8,
	0x11, 0x95, 0x9e, 0xb6, 0xab, 0xdd, 0xba, 0x59, 0xcf, 0x22, 0x4a, 0x44, 0xd1, 0x4f, 0x20, 0xcd,
	0x5d, 0xca, 0xf2, 0x41, 0x63, 0xdf, 0x61, 0x41, 0x88, 0x29, 0x73, 0xc3, 0x58, 0xaa, 0xb7, 0x85,
	0x6e, 0xd5, 0x3c, 0x4d, 0x71, 0x39, 0x87, 0xed, 0x1c, 0x6d, 0xfd, 0xb3, 0x0f, 0x68, 0xb3, 0x23,
	0x84, 0xa1, 0x91, 0x1d, 0xd3, 0xaa, 0x7e, 0x81, 0xd7, 0x7f, 0xf5, 0x79, 0xe7, 0x93, 0x85, 0xf2,
	0xd1, 0xc6, 0x0f, 0x2f, 0xe8, 0x5d, 0x76, 0x53, 0xcb, 0x1f, 0xe2, 0x22, 0x99, 0xf3, 0x0f, 0xa5,
	0x9e, 0xdd, 0xb4, 0x12, 0x38, 0x4e, 0xe6, 0xe8, 0x2f, 0x01, 0x5a, 0xa1, 0x1b, 0x2d, 0xdc, 0xb9,
	0xb3, 0x6d, 0x94, 0x55, 0x3e, 0x4a, 0xfd, 0x33, 0x4b, 0x1d, 0x71, 0xe1, 0x2d, 0x33, 0x96, 0xc2,
	0x47, 0x90, 0xd6, 0x7f, 0x15, 0x90, 0x1e, 0x4b, 0x43, 0xaf, 0xe0, 0x68, 0xc6, 0x58, 0xbc, 0xaa,
	0x34, 0x6d, 0x4d, 0xe0, 0xad, 0x35, 0xd2, 0x28, 0xe7, 0xa7, 0x2d, 0x7d, 0x0d, 0xcf, 0x0a, 0xac,
	0x98, 0x24, 0x8c, 0x9f, 0xc0, 0x9e, 0x79, 0x78, 0x4f, 0x33, 0x48, 0xc2, 0xd0, 0x1b, 0x68, 0xf2,
	0x4d, 0x80, 0x9d, 0x35, 0xd1, 0x2a, 0x17, 0x45, 0x19, 0x78, 0x53, 0x94, 0x7e, 0x07, 0xa7, 0x9b,
	0x29, 0xdc, 0xe1, 0x09, 0x77, 0x38, 0x59, 0xcb, 0xe1, 0x3e, 0xaf, 0xe0, 0x88, 0x12, 0xef, 0x0f,
	0xea, 0xcc, 0x08, 0x65, 0xdc, 0x60, 0x2f, 0xab, 0x9a, 0x47, 0x6f, 0x08, 0x65, 0xab, 0xaa, 0x0b,
	0x2c, 0xae, 0xb9, 0x9f, 0x55, 0x7d, 0x4f, 0xe3, 0x6a, 0xdf, 0x80, 0x88, 0xef, 0xbc, 0xf9, 0xc2,
	0xc7, 0xbe, 0xe3, 0x93, 0xd0, 0x0d, 0x22, 0x2a, 0xd5, 0xf8, 0x15, 0x7e, 0x96, 0xc7, 0x95, 0x2c,
	0xdc, 0xf9, 0x5b, 0x80, 0x83, 0xc2, 0x6d, 0x41, 0x5f, 0x81, 0x64, 0x98, 0xfa, 0xc7, 0x89, 0xa3,
	0x1b, 0xf6, 0x40, 0xd7, 0x9c, 0xb1, 0x66, 0x19, 0xaa, 0x3c, 0xb8, 0x1e, 0xa8, 0x8a, 0xf8, 0x05,
	0x3a, 0x83, 0x66, 0x09, 0x55, 0x06, 0x56, 0xff, 0x6a, 0xa8, 0x2a, 0xa2, 0x80, 0x5a, 0x70, 0x5a,
	0x82, 0xfa, 0x63, 0x5b, 0x1f, 0xf5, 0xed, 0x81, 0x2c, 0x56, 0xd0, 0x05, 0xb4, 0x36, 0x30, 0x65,
	0x60, 0xc9, 0xfa, 0xaf, 0xaa, 0x39, 0x11, 0xab, 0xe8, 0x05, 0x9c, 0x94, 0xf0, 0x51, 0x5f, 0x1b,
	0xf7, 0x87, 0xe2, 0x93, 0x0e, 0x85, 0x46, 0x71, 0x2d, 0xa3, 0x73, 0x38, 0xb3, 0x54, 0x79, 0x6c,
	0x0e, 0xec, 0x89, 0x63, 0x4f, 0x0c, 0x75, 0xad, 0xbc, 0x53, 0x40, 0x65, 0x58, 0xd3, 0x35, 0x55,
	0x14, 0x50, 0x13, 0x8e, 0xcb, 0xf1, 0x0f, 0xaa, 0x21, 0x56, 0x36, 0xc3, 0x86, 0xf5, 0x5e, 0xac,
	0x76, 0xfe, 0x14, 0xa0, 0xf5, 0xf8, 0x36, 0x46, 0xaf, 0xe1, 0xe5, 0x7d, 0x6f, 0xfd, 0xe1, 0x70,
	0xe2, 0xc8, 0xba, 0xa6, 0xa9, 0xb2, 0xbd, 0x56, 0x4b, 0x07, 0x2e, 0xb6, 0xd3, 0x0a, 0x67, 0xf6,
	0x12, 0xce, 0xb7, 0x73, 0x54, 0x2d, 0xa3, 0x54, 0x3a, 0x53, 0x38, 0xde, 0xd8, 0xd6, 0xe9, 0x90,
	0x06, 0x96, 0x63, 0x98, 0xea, 0xb5, 0x6a, 0x9a, 0xaa, 0xb2, 0x39, 0xa4, 0x12, 0x5a, 0x30, 0x94,
	0xe0, 0x79, 0x09, 0x7a, 0xf0, 0xf9, 0x04, 0x87, 0xa5, 0x05, 0x9e, 0xce, 0x6c, 0xa4, 0xda, 0x6a,
	0xca, 0xda, 0x7a, 0x15, 0x9a, 0x70, 0xbc, 0x86, 0x6b, 0xba, 0x28, 0xa4, 0x23, 0x58, 0x0b, 0x4f,
	0x54, 0x4b, 0xac, 0xa4, 0x23, 0x5e, 0x8b, 0xa7, 0x9d, 0x8b, 0xd5, 0xce, 0x6f, 0x50, 0xbf, 0x5f,
	0xc0, 0xe9, 0x25, 0x52, 0x34, 0x6b, 0xbb, 0xe1, 0x97, 0xf0, 0xa2, 0x80, 0x29, 0xea, 0x75, 0x7f,
	0x3c, 0xb4, 0x1d, 0xe5, 0x46, 0x36, 0xb2, 0x09, 0x17, 0x40, 0x79, 0x6c, 0xd9, 0xfa, 0x48, 0xac,
	0x5c, 0x7d, 0x0b, 0xaf, 0x49, 0x72, 0xdb, 0xf3, 0x66, 0x09, 0x09, 0x83, 0x45, 0xd8, 0x7b, 0xf8,
	0x3d, 0xe1, 0xdb, 0xaa, 0x97, 0xff, 0x9e, 0xdc, 0x54, 0x0d, 0xe1, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xea, 0x79, 0xc6, 0xc2, 0xbd, 0x08, 0x00, 0x00,
}