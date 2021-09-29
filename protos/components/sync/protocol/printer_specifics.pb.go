// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/printer_specifics.proto

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

// User PPD configuration
type PrinterPPDReference struct {
	// Url for user provided file.  Overrides other fields.
	UserSuppliedPpdUrl *string `protobuf:"bytes,1,opt,name=user_supplied_ppd_url,json=userSuppliedPpdUrl" json:"user_supplied_ppd_url,omitempty"`
	// Retired fields
	EffectiveManufacturer *string `protobuf:"bytes,2,opt,name=effective_manufacturer,json=effectiveManufacturer" json:"effective_manufacturer,omitempty"` // Deprecated: Do not use.
	EffectiveModel        *string `protobuf:"bytes,3,opt,name=effective_model,json=effectiveModel" json:"effective_model,omitempty"`                      // Deprecated: Do not use.
	// String identifying the type of printer, used to look up a ppd to drive the
	// printer.
	EffectiveMakeAndModel *string `protobuf:"bytes,4,opt,name=effective_make_and_model,json=effectiveMakeAndModel" json:"effective_make_and_model,omitempty"`
	// True if the printer should be automatically configured, false otherwise.
	Autoconf             *bool    `protobuf:"varint,5,opt,name=autoconf,def=0" json:"autoconf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrinterPPDReference) Reset()         { *m = PrinterPPDReference{} }
func (m *PrinterPPDReference) String() string { return proto.CompactTextString(m) }
func (*PrinterPPDReference) ProtoMessage()    {}
func (*PrinterPPDReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_printer_specifics_9b925cb5523b69c9, []int{0}
}
func (m *PrinterPPDReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrinterPPDReference.Unmarshal(m, b)
}
func (m *PrinterPPDReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrinterPPDReference.Marshal(b, m, deterministic)
}
func (dst *PrinterPPDReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrinterPPDReference.Merge(dst, src)
}
func (m *PrinterPPDReference) XXX_Size() int {
	return xxx_messageInfo_PrinterPPDReference.Size(m)
}
func (m *PrinterPPDReference) XXX_DiscardUnknown() {
	xxx_messageInfo_PrinterPPDReference.DiscardUnknown(m)
}

var xxx_messageInfo_PrinterPPDReference proto.InternalMessageInfo

const Default_PrinterPPDReference_Autoconf bool = false

func (m *PrinterPPDReference) GetUserSuppliedPpdUrl() string {
	if m != nil && m.UserSuppliedPpdUrl != nil {
		return *m.UserSuppliedPpdUrl
	}
	return ""
}

// Deprecated: Do not use.
func (m *PrinterPPDReference) GetEffectiveManufacturer() string {
	if m != nil && m.EffectiveManufacturer != nil {
		return *m.EffectiveManufacturer
	}
	return ""
}

// Deprecated: Do not use.
func (m *PrinterPPDReference) GetEffectiveModel() string {
	if m != nil && m.EffectiveModel != nil {
		return *m.EffectiveModel
	}
	return ""
}

func (m *PrinterPPDReference) GetEffectiveMakeAndModel() string {
	if m != nil && m.EffectiveMakeAndModel != nil {
		return *m.EffectiveMakeAndModel
	}
	return ""
}

func (m *PrinterPPDReference) GetAutoconf() bool {
	if m != nil && m.Autoconf != nil {
		return *m.Autoconf
	}
	return Default_PrinterPPDReference_Autoconf
}

type PrinterSpecifics struct {
	// Printer record GUID
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// User visible name.  Any string.
	DisplayName *string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// User visible description.  Any string.
	Description *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Printer manufacturer.  Should be a known manufacturuer.
	// Deprecated in favor of make_and_model.
	Manufacturer *string `protobuf:"bytes,4,opt,name=manufacturer" json:"manufacturer,omitempty"` // Deprecated: Do not use.
	// Printer model.  Should match a known model for the manufacturer.
	// Deprecated in favor of make_and_model.
	Model *string `protobuf:"bytes,5,opt,name=model" json:"model,omitempty"` // Deprecated: Do not use.
	// Universal Resource Identifier for the printer on the network.  usb:// will
	// be the scheme for USB printers.  Example
	// ipp://address.example:port/queue/queue/queue.
	Uri *string `protobuf:"bytes,6,opt,name=uri" json:"uri,omitempty"`
	// Universally Unique Identifier provided by the printer.  Used for unique
	// identification of printers in a zeroconf environment.
	Uuid *string `protobuf:"bytes,7,opt,name=uuid" json:"uuid,omitempty"`
	// PPDData was deprecated in favor of PPDReference format.
	Ppd []byte `protobuf:"bytes,8,opt,name=ppd" json:"ppd,omitempty"` // Deprecated: Do not use.
	// Structure representing the user's ppd configuration.
	PpdReference *PrinterPPDReference `protobuf:"bytes,9,opt,name=ppd_reference,json=ppdReference" json:"ppd_reference,omitempty"`
	// Timestamp when printer was last updated.
	UpdatedTimestamp *int64 `protobuf:"varint,10,opt,name=updated_timestamp,json=updatedTimestamp" json:"updated_timestamp,omitempty"`
	// The make and model of the printer in one string.  The typical arrangement
	// for this is '<make> <model>'.  This aligns with the typical formatting of
	// the IPP attribute printer-make-and-model.
	MakeAndModel *string `protobuf:"bytes,11,opt,name=make_and_model,json=makeAndModel" json:"make_and_model,omitempty"`
	// Universal Resource Identifier for the print server on the network.  This
	// will only be populated if the printer is from a print server.  Example
	// ipp://address.example:port/
	PrintServerUri       *string  `protobuf:"bytes,12,opt,name=print_server_uri,json=printServerUri" json:"print_server_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrinterSpecifics) Reset()         { *m = PrinterSpecifics{} }
func (m *PrinterSpecifics) String() string { return proto.CompactTextString(m) }
func (*PrinterSpecifics) ProtoMessage()    {}
func (*PrinterSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_printer_specifics_9b925cb5523b69c9, []int{1}
}
func (m *PrinterSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrinterSpecifics.Unmarshal(m, b)
}
func (m *PrinterSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrinterSpecifics.Marshal(b, m, deterministic)
}
func (dst *PrinterSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrinterSpecifics.Merge(dst, src)
}
func (m *PrinterSpecifics) XXX_Size() int {
	return xxx_messageInfo_PrinterSpecifics.Size(m)
}
func (m *PrinterSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_PrinterSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_PrinterSpecifics proto.InternalMessageInfo

func (m *PrinterSpecifics) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *PrinterSpecifics) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *PrinterSpecifics) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

// Deprecated: Do not use.
func (m *PrinterSpecifics) GetManufacturer() string {
	if m != nil && m.Manufacturer != nil {
		return *m.Manufacturer
	}
	return ""
}

// Deprecated: Do not use.
func (m *PrinterSpecifics) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *PrinterSpecifics) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *PrinterSpecifics) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

// Deprecated: Do not use.
func (m *PrinterSpecifics) GetPpd() []byte {
	if m != nil {
		return m.Ppd
	}
	return nil
}

func (m *PrinterSpecifics) GetPpdReference() *PrinterPPDReference {
	if m != nil {
		return m.PpdReference
	}
	return nil
}

func (m *PrinterSpecifics) GetUpdatedTimestamp() int64 {
	if m != nil && m.UpdatedTimestamp != nil {
		return *m.UpdatedTimestamp
	}
	return 0
}

func (m *PrinterSpecifics) GetMakeAndModel() string {
	if m != nil && m.MakeAndModel != nil {
		return *m.MakeAndModel
	}
	return ""
}

func (m *PrinterSpecifics) GetPrintServerUri() string {
	if m != nil && m.PrintServerUri != nil {
		return *m.PrintServerUri
	}
	return ""
}

func init() {
	proto.RegisterType((*PrinterPPDReference)(nil), "sync_pb.PrinterPPDReference")
	proto.RegisterType((*PrinterSpecifics)(nil), "sync_pb.PrinterSpecifics")
}

func init() {
	proto.RegisterFile("components/sync/protocol/printer_specifics.proto", fileDescriptor_printer_specifics_9b925cb5523b69c9)
}

var fileDescriptor_printer_specifics_9b925cb5523b69c9 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x56, 0x9a, 0x95, 0x75, 0x6f, 0x43, 0x29, 0x86, 0x21, 0x1f, 0x38, 0x64, 0x13, 0xa0, 0x48,
	0x95, 0x52, 0xe0, 0x82, 0xe0, 0xb6, 0x89, 0x03, 0x17, 0x50, 0x95, 0xb2, 0x73, 0x64, 0xec, 0x37,
	0x60, 0x2d, 0x71, 0x2c, 0x3b, 0x9e, 0xb4, 0xff, 0xc0, 0xdf, 0xe0, 0x7f, 0x22, 0x3b, 0x69, 0x9a,
	0x21, 0x6e, 0xce, 0xf3, 0x11, 0xfb, 0x79, 0x3f, 0xe0, 0x2d, 0x6f, 0x1b, 0xdd, 0x2a, 0x54, 0x9d,
	0xdd, 0xda, 0x7b, 0xc5, 0xb7, 0xda, 0xb4, 0x5d, 0xcb, 0xdb, 0x7a, 0xab, 0x8d, 0x54, 0x1d, 0x9a,
	0xd2, 0x6a, 0xe4, 0xb2, 0x92, 0xdc, 0xe6, 0x81, 0x22, 0xa7, 0x5e, 0x56, 0xea, 0x1f, 0x97, 0xbf,
	0x67, 0xf0, 0x6c, 0xd7, 0x8b, 0x76, 0xbb, 0xcf, 0x05, 0x56, 0x68, 0x50, 0x71, 0x24, 0xef, 0xe0,
	0xdc, 0x59, 0x6f, 0x74, 0x5a, 0xd7, 0x12, 0x45, 0xa9, 0xb5, 0x28, 0x9d, 0xa9, 0x69, 0x94, 0x46,
	0xd9, 0x59, 0x41, 0x3c, 0xb9, 0x1f, 0xb8, 0x9d, 0x16, 0x37, 0xa6, 0x26, 0x1f, 0xe1, 0x05, 0x56,
	0x15, 0xf2, 0x4e, 0xde, 0x61, 0xd9, 0x30, 0xe5, 0x2a, 0xc6, 0x3b, 0x67, 0xd0, 0xd0, 0x99, 0xf7,
	0x5c, 0xcf, 0x68, 0x54, 0x9c, 0x8f, 0x8a, 0xaf, 0x13, 0x01, 0xd9, 0xc0, 0x93, 0x89, 0xb5, 0x15,
	0x58, 0xd3, 0x78, 0xf4, 0xac, 0x8e, 0x1e, 0xcf, 0x90, 0x0f, 0x40, 0xa7, 0xf7, 0xdc, 0x62, 0xc9,
	0x94, 0x18, 0x5c, 0x27, 0xe1, 0x75, 0xd3, 0x5b, 0x6e, 0xf1, 0x4a, 0x89, 0xde, 0x78, 0x01, 0x0b,
	0xe6, 0x7c, 0x5d, 0x54, 0x45, 0xe7, 0x69, 0x94, 0x2d, 0x3e, 0xcd, 0x2b, 0x56, 0x5b, 0x2c, 0x46,
	0xf8, 0xf2, 0x4f, 0x0c, 0xeb, 0xa1, 0x1c, 0xfb, 0x43, 0xc9, 0xc8, 0x0a, 0x66, 0x52, 0x0c, 0xc1,
	0x67, 0x52, 0x90, 0x0b, 0x48, 0x84, 0xb4, 0xba, 0x66, 0xf7, 0xa5, 0x62, 0x0d, 0xf6, 0xf1, 0x8a,
	0xe5, 0x80, 0x7d, 0x63, 0x0d, 0x92, 0x14, 0x96, 0x02, 0x2d, 0x37, 0x52, 0x77, 0xb2, 0x55, 0x7d,
	0x98, 0x62, 0x0a, 0x91, 0x37, 0x90, 0x3c, 0xa8, 0xd1, 0xc9, 0x98, 0xf7, 0x01, 0x4e, 0x28, 0xcc,
	0xfb, 0x68, 0xf3, 0x51, 0xd0, 0x03, 0x64, 0x0d, 0xb1, 0x33, 0x92, 0x3e, 0x0a, 0xff, 0xf6, 0x47,
	0x42, 0xe0, 0xc4, 0x39, 0x29, 0xe8, 0x69, 0x80, 0xc2, 0x99, 0x3c, 0x87, 0x58, 0x6b, 0x41, 0x17,
	0x69, 0x94, 0x25, 0xc1, 0xed, 0x3f, 0xc9, 0x15, 0x3c, 0xf6, 0x0d, 0x35, 0x87, 0x7e, 0xd3, 0xb3,
	0x34, 0xca, 0x96, 0xef, 0x5f, 0xe6, 0xc3, 0x5c, 0xe4, 0xff, 0x99, 0x89, 0x22, 0xd1, 0x5a, 0x1c,
	0x27, 0x64, 0x03, 0x4f, 0x9d, 0x16, 0xac, 0x43, 0x51, 0x76, 0xb2, 0x41, 0xdb, 0xb1, 0x46, 0x53,
	0x48, 0xa3, 0x2c, 0x2e, 0xd6, 0x03, 0xf1, 0xfd, 0x80, 0x93, 0x57, 0xb0, 0xfa, 0xa7, 0x53, 0xcb,
	0xf0, 0xc6, 0xa4, 0x99, 0x36, 0x28, 0x83, 0x75, 0x18, 0xd8, 0xd2, 0xa2, 0xb9, 0x43, 0x53, 0xfa,
	0x78, 0x49, 0xd0, 0xad, 0x02, 0xbe, 0x0f, 0xf0, 0x8d, 0x91, 0xd7, 0x1b, 0x78, 0xdd, 0x9a, 0x9f,
	0x39, 0xff, 0x65, 0xda, 0x46, 0xba, 0x26, 0x3f, 0x2e, 0x40, 0x48, 0x90, 0x1f, 0x16, 0xe0, 0x4b,
	0xbc, 0x8b, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x88, 0xf9, 0xb2, 0x95, 0x1f, 0x03, 0x00, 0x00,
}
