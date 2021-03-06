// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/sync/protocol/autofill_offer_specifics.proto

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

// Next tag: 11
type AutofillOfferSpecifics struct {
	// The id for this offer data. Will be used as the client tag.
	Id *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The link leading to the offer details page on Gpay app. Will be populated
	// on Android only.
	OfferDetailsUrl *string `protobuf:"bytes,2,opt,name=offer_details_url,json=offerDetailsUrl" json:"offer_details_url,omitempty"`
	// Merchant domain and merchant app package name refers to the merchant this
	// offer is applied to.
	MerchantDomain     []string `protobuf:"bytes,3,rep,name=merchant_domain,json=merchantDomain" json:"merchant_domain,omitempty"`
	MerchantAppPackage []string `protobuf:"bytes,4,rep,name=merchant_app_package,json=merchantAppPackage" json:"merchant_app_package,omitempty"`
	// The expiry of this offer. Will be represented in the form of unix epoch
	// time in seconds. Once the offer is expired it will not be shown in the
	// client.
	OfferExpiryDate *int64 `protobuf:"varint,5,opt,name=offer_expiry_date,json=offerExpiryDate" json:"offer_expiry_date,omitempty"`
	// The unique offer data for different offer types.
	//
	// Types that are valid to be assigned to TypeSpecificOfferData:
	//	*AutofillOfferSpecifics_CardLinkedOfferData_
	//	*AutofillOfferSpecifics_PromoCodeOfferData_
	TypeSpecificOfferData isAutofillOfferSpecifics_TypeSpecificOfferData `protobuf_oneof:"type_specific_offer_data"`
	DisplayStrings        *AutofillOfferSpecifics_DisplayStrings         `protobuf:"bytes,10,opt,name=display_strings,json=displayStrings" json:"display_strings,omitempty"`
	// The reward type of the offer. Will be used to generate the display text in
	// the UI. Each type has its own client side text template.
	//
	// Types that are valid to be assigned to RewardType:
	//	*AutofillOfferSpecifics_PercentageReward_
	//	*AutofillOfferSpecifics_FixedAmountReward_
	RewardType           isAutofillOfferSpecifics_RewardType `protobuf_oneof:"reward_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *AutofillOfferSpecifics) Reset()         { *m = AutofillOfferSpecifics{} }
func (m *AutofillOfferSpecifics) String() string { return proto.CompactTextString(m) }
func (*AutofillOfferSpecifics) ProtoMessage()    {}
func (*AutofillOfferSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_autofill_offer_specifics_d8a25529d974b566, []int{0}
}
func (m *AutofillOfferSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutofillOfferSpecifics.Unmarshal(m, b)
}
func (m *AutofillOfferSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutofillOfferSpecifics.Marshal(b, m, deterministic)
}
func (dst *AutofillOfferSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutofillOfferSpecifics.Merge(dst, src)
}
func (m *AutofillOfferSpecifics) XXX_Size() int {
	return xxx_messageInfo_AutofillOfferSpecifics.Size(m)
}
func (m *AutofillOfferSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_AutofillOfferSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_AutofillOfferSpecifics proto.InternalMessageInfo

type isAutofillOfferSpecifics_TypeSpecificOfferData interface {
	isAutofillOfferSpecifics_TypeSpecificOfferData()
}
type isAutofillOfferSpecifics_RewardType interface {
	isAutofillOfferSpecifics_RewardType()
}

type AutofillOfferSpecifics_CardLinkedOfferData_ struct {
	CardLinkedOfferData *AutofillOfferSpecifics_CardLinkedOfferData `protobuf:"bytes,6,opt,name=card_linked_offer_data,json=cardLinkedOfferData,oneof"`
}
type AutofillOfferSpecifics_PromoCodeOfferData_ struct {
	PromoCodeOfferData *AutofillOfferSpecifics_PromoCodeOfferData `protobuf:"bytes,9,opt,name=promo_code_offer_data,json=promoCodeOfferData,oneof"`
}
type AutofillOfferSpecifics_PercentageReward_ struct {
	PercentageReward *AutofillOfferSpecifics_PercentageReward `protobuf:"bytes,7,opt,name=percentage_reward,json=percentageReward,oneof"`
}
type AutofillOfferSpecifics_FixedAmountReward_ struct {
	FixedAmountReward *AutofillOfferSpecifics_FixedAmountReward `protobuf:"bytes,8,opt,name=fixed_amount_reward,json=fixedAmountReward,oneof"`
}

func (*AutofillOfferSpecifics_CardLinkedOfferData_) isAutofillOfferSpecifics_TypeSpecificOfferData() {
}
func (*AutofillOfferSpecifics_PromoCodeOfferData_) isAutofillOfferSpecifics_TypeSpecificOfferData() {}
func (*AutofillOfferSpecifics_PercentageReward_) isAutofillOfferSpecifics_RewardType()              {}
func (*AutofillOfferSpecifics_FixedAmountReward_) isAutofillOfferSpecifics_RewardType()             {}

func (m *AutofillOfferSpecifics) GetTypeSpecificOfferData() isAutofillOfferSpecifics_TypeSpecificOfferData {
	if m != nil {
		return m.TypeSpecificOfferData
	}
	return nil
}
func (m *AutofillOfferSpecifics) GetRewardType() isAutofillOfferSpecifics_RewardType {
	if m != nil {
		return m.RewardType
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *AutofillOfferSpecifics) GetOfferDetailsUrl() string {
	if m != nil && m.OfferDetailsUrl != nil {
		return *m.OfferDetailsUrl
	}
	return ""
}

func (m *AutofillOfferSpecifics) GetMerchantDomain() []string {
	if m != nil {
		return m.MerchantDomain
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetMerchantAppPackage() []string {
	if m != nil {
		return m.MerchantAppPackage
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetOfferExpiryDate() int64 {
	if m != nil && m.OfferExpiryDate != nil {
		return *m.OfferExpiryDate
	}
	return 0
}

func (m *AutofillOfferSpecifics) GetCardLinkedOfferData() *AutofillOfferSpecifics_CardLinkedOfferData {
	if x, ok := m.GetTypeSpecificOfferData().(*AutofillOfferSpecifics_CardLinkedOfferData_); ok {
		return x.CardLinkedOfferData
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetPromoCodeOfferData() *AutofillOfferSpecifics_PromoCodeOfferData {
	if x, ok := m.GetTypeSpecificOfferData().(*AutofillOfferSpecifics_PromoCodeOfferData_); ok {
		return x.PromoCodeOfferData
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetDisplayStrings() *AutofillOfferSpecifics_DisplayStrings {
	if m != nil {
		return m.DisplayStrings
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetPercentageReward() *AutofillOfferSpecifics_PercentageReward {
	if x, ok := m.GetRewardType().(*AutofillOfferSpecifics_PercentageReward_); ok {
		return x.PercentageReward
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetFixedAmountReward() *AutofillOfferSpecifics_FixedAmountReward {
	if x, ok := m.GetRewardType().(*AutofillOfferSpecifics_FixedAmountReward_); ok {
		return x.FixedAmountReward
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AutofillOfferSpecifics) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AutofillOfferSpecifics_OneofMarshaler, _AutofillOfferSpecifics_OneofUnmarshaler, _AutofillOfferSpecifics_OneofSizer, []interface{}{
		(*AutofillOfferSpecifics_CardLinkedOfferData_)(nil),
		(*AutofillOfferSpecifics_PromoCodeOfferData_)(nil),
		(*AutofillOfferSpecifics_PercentageReward_)(nil),
		(*AutofillOfferSpecifics_FixedAmountReward_)(nil),
	}
}

func _AutofillOfferSpecifics_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AutofillOfferSpecifics)
	// type_specific_offer_data
	switch x := m.TypeSpecificOfferData.(type) {
	case *AutofillOfferSpecifics_CardLinkedOfferData_:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CardLinkedOfferData); err != nil {
			return err
		}
	case *AutofillOfferSpecifics_PromoCodeOfferData_:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PromoCodeOfferData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AutofillOfferSpecifics.TypeSpecificOfferData has unexpected type %T", x)
	}
	// reward_type
	switch x := m.RewardType.(type) {
	case *AutofillOfferSpecifics_PercentageReward_:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PercentageReward); err != nil {
			return err
		}
	case *AutofillOfferSpecifics_FixedAmountReward_:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedAmountReward); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AutofillOfferSpecifics.RewardType has unexpected type %T", x)
	}
	return nil
}

func _AutofillOfferSpecifics_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AutofillOfferSpecifics)
	switch tag {
	case 6: // type_specific_offer_data.card_linked_offer_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AutofillOfferSpecifics_CardLinkedOfferData)
		err := b.DecodeMessage(msg)
		m.TypeSpecificOfferData = &AutofillOfferSpecifics_CardLinkedOfferData_{msg}
		return true, err
	case 9: // type_specific_offer_data.promo_code_offer_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AutofillOfferSpecifics_PromoCodeOfferData)
		err := b.DecodeMessage(msg)
		m.TypeSpecificOfferData = &AutofillOfferSpecifics_PromoCodeOfferData_{msg}
		return true, err
	case 7: // reward_type.percentage_reward
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AutofillOfferSpecifics_PercentageReward)
		err := b.DecodeMessage(msg)
		m.RewardType = &AutofillOfferSpecifics_PercentageReward_{msg}
		return true, err
	case 8: // reward_type.fixed_amount_reward
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AutofillOfferSpecifics_FixedAmountReward)
		err := b.DecodeMessage(msg)
		m.RewardType = &AutofillOfferSpecifics_FixedAmountReward_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AutofillOfferSpecifics_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AutofillOfferSpecifics)
	// type_specific_offer_data
	switch x := m.TypeSpecificOfferData.(type) {
	case *AutofillOfferSpecifics_CardLinkedOfferData_:
		s := proto.Size(x.CardLinkedOfferData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutofillOfferSpecifics_PromoCodeOfferData_:
		s := proto.Size(x.PromoCodeOfferData)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// reward_type
	switch x := m.RewardType.(type) {
	case *AutofillOfferSpecifics_PercentageReward_:
		s := proto.Size(x.PercentageReward)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutofillOfferSpecifics_FixedAmountReward_:
		s := proto.Size(x.FixedAmountReward)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Proto containing data specific to a card-linked offer.
type AutofillOfferSpecifics_CardLinkedOfferData struct {
	// The server id of the card to which the offer is linked.
	InstrumentId         []int64  `protobuf:"varint,3,rep,name=instrument_id,json=instrumentId" json:"instrument_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutofillOfferSpecifics_CardLinkedOfferData) Reset() {
	*m = AutofillOfferSpecifics_CardLinkedOfferData{}
}
func (m *AutofillOfferSpecifics_CardLinkedOfferData) String() string {
	return proto.CompactTextString(m)
}
func (*AutofillOfferSpecifics_CardLinkedOfferData) ProtoMessage() {}
func (*AutofillOfferSpecifics_CardLinkedOfferData) Descriptor() ([]byte, []int) {
	return fileDescriptor_autofill_offer_specifics_d8a25529d974b566, []int{0, 0}
}
func (m *AutofillOfferSpecifics_CardLinkedOfferData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutofillOfferSpecifics_CardLinkedOfferData.Unmarshal(m, b)
}
func (m *AutofillOfferSpecifics_CardLinkedOfferData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutofillOfferSpecifics_CardLinkedOfferData.Marshal(b, m, deterministic)
}
func (dst *AutofillOfferSpecifics_CardLinkedOfferData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutofillOfferSpecifics_CardLinkedOfferData.Merge(dst, src)
}
func (m *AutofillOfferSpecifics_CardLinkedOfferData) XXX_Size() int {
	return xxx_messageInfo_AutofillOfferSpecifics_CardLinkedOfferData.Size(m)
}
func (m *AutofillOfferSpecifics_CardLinkedOfferData) XXX_DiscardUnknown() {
	xxx_messageInfo_AutofillOfferSpecifics_CardLinkedOfferData.DiscardUnknown(m)
}

var xxx_messageInfo_AutofillOfferSpecifics_CardLinkedOfferData proto.InternalMessageInfo

func (m *AutofillOfferSpecifics_CardLinkedOfferData) GetInstrumentId() []int64 {
	if m != nil {
		return m.InstrumentId
	}
	return nil
}

// Proto containing data specific to a promo code offer.
type AutofillOfferSpecifics_PromoCodeOfferData struct {
	// The actual promo code which can be applied at checkout.
	PromoCode            *string  `protobuf:"bytes,1,opt,name=promo_code,json=promoCode" json:"promo_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutofillOfferSpecifics_PromoCodeOfferData) Reset() {
	*m = AutofillOfferSpecifics_PromoCodeOfferData{}
}
func (m *AutofillOfferSpecifics_PromoCodeOfferData) String() string {
	return proto.CompactTextString(m)
}
func (*AutofillOfferSpecifics_PromoCodeOfferData) ProtoMessage() {}
func (*AutofillOfferSpecifics_PromoCodeOfferData) Descriptor() ([]byte, []int) {
	return fileDescriptor_autofill_offer_specifics_d8a25529d974b566, []int{0, 1}
}
func (m *AutofillOfferSpecifics_PromoCodeOfferData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutofillOfferSpecifics_PromoCodeOfferData.Unmarshal(m, b)
}
func (m *AutofillOfferSpecifics_PromoCodeOfferData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutofillOfferSpecifics_PromoCodeOfferData.Marshal(b, m, deterministic)
}
func (dst *AutofillOfferSpecifics_PromoCodeOfferData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutofillOfferSpecifics_PromoCodeOfferData.Merge(dst, src)
}
func (m *AutofillOfferSpecifics_PromoCodeOfferData) XXX_Size() int {
	return xxx_messageInfo_AutofillOfferSpecifics_PromoCodeOfferData.Size(m)
}
func (m *AutofillOfferSpecifics_PromoCodeOfferData) XXX_DiscardUnknown() {
	xxx_messageInfo_AutofillOfferSpecifics_PromoCodeOfferData.DiscardUnknown(m)
}

var xxx_messageInfo_AutofillOfferSpecifics_PromoCodeOfferData proto.InternalMessageInfo

func (m *AutofillOfferSpecifics_PromoCodeOfferData) GetPromoCode() string {
	if m != nil && m.PromoCode != nil {
		return *m.PromoCode
	}
	return ""
}

// Strings to be shown in client UI, based on the offer type and details.
type AutofillOfferSpecifics_DisplayStrings struct {
	// A message translated in the user's GPay app locale, explaining the value
	// of the offer. For example, a promo code offer might display
	// "$5 off on shoes, up to $50."
	ValuePropText *string `protobuf:"bytes,1,opt,name=value_prop_text,json=valuePropText" json:"value_prop_text,omitempty"`
	// A message translated in the user's GPay app locale and shown on mobile as
	// a link, prompting the user to click it to learn more about the offer.
	// Generally, "See details".
	SeeDetailsTextMobile *string `protobuf:"bytes,2,opt,name=see_details_text_mobile,json=seeDetailsTextMobile" json:"see_details_text_mobile,omitempty"`
	// A message translated in the user's GPay app locale and shown on desktop
	// (not as a link), informing the user that exclusions and restrictions may
	// apply to the value prop text. Generally, "Terms apply."
	SeeDetailsTextDesktop *string `protobuf:"bytes,3,opt,name=see_details_text_desktop,json=seeDetailsTextDesktop" json:"see_details_text_desktop,omitempty"`
	// A message translated in the user's GPay app locale and shown on mobile,
	// instructing them on how to redeem the offer. For example, a promo code
	// offer might display "Tap the promo code field at checkout to autofill
	// it."
	UsageInstructionsTextMobile *string `protobuf:"bytes,4,opt,name=usage_instructions_text_mobile,json=usageInstructionsTextMobile" json:"usage_instructions_text_mobile,omitempty"`
	// A message translated in the user's GPay app locale and shown on desktop,
	// instructing them on how to redeem the offer. For example, a promo code
	// offer might display "Click the promo code field at checkout to autofill
	// it."
	UsageInstructionsTextDesktop *string  `protobuf:"bytes,5,opt,name=usage_instructions_text_desktop,json=usageInstructionsTextDesktop" json:"usage_instructions_text_desktop,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *AutofillOfferSpecifics_DisplayStrings) Reset()         { *m = AutofillOfferSpecifics_DisplayStrings{} }
func (m *AutofillOfferSpecifics_DisplayStrings) String() string { return proto.CompactTextString(m) }
func (*AutofillOfferSpecifics_DisplayStrings) ProtoMessage()    {}
func (*AutofillOfferSpecifics_DisplayStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_autofill_offer_specifics_d8a25529d974b566, []int{0, 2}
}
func (m *AutofillOfferSpecifics_DisplayStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutofillOfferSpecifics_DisplayStrings.Unmarshal(m, b)
}
func (m *AutofillOfferSpecifics_DisplayStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutofillOfferSpecifics_DisplayStrings.Marshal(b, m, deterministic)
}
func (dst *AutofillOfferSpecifics_DisplayStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutofillOfferSpecifics_DisplayStrings.Merge(dst, src)
}
func (m *AutofillOfferSpecifics_DisplayStrings) XXX_Size() int {
	return xxx_messageInfo_AutofillOfferSpecifics_DisplayStrings.Size(m)
}
func (m *AutofillOfferSpecifics_DisplayStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_AutofillOfferSpecifics_DisplayStrings.DiscardUnknown(m)
}

var xxx_messageInfo_AutofillOfferSpecifics_DisplayStrings proto.InternalMessageInfo

func (m *AutofillOfferSpecifics_DisplayStrings) GetValuePropText() string {
	if m != nil && m.ValuePropText != nil {
		return *m.ValuePropText
	}
	return ""
}

func (m *AutofillOfferSpecifics_DisplayStrings) GetSeeDetailsTextMobile() string {
	if m != nil && m.SeeDetailsTextMobile != nil {
		return *m.SeeDetailsTextMobile
	}
	return ""
}

func (m *AutofillOfferSpecifics_DisplayStrings) GetSeeDetailsTextDesktop() string {
	if m != nil && m.SeeDetailsTextDesktop != nil {
		return *m.SeeDetailsTextDesktop
	}
	return ""
}

func (m *AutofillOfferSpecifics_DisplayStrings) GetUsageInstructionsTextMobile() string {
	if m != nil && m.UsageInstructionsTextMobile != nil {
		return *m.UsageInstructionsTextMobile
	}
	return ""
}

func (m *AutofillOfferSpecifics_DisplayStrings) GetUsageInstructionsTextDesktop() string {
	if m != nil && m.UsageInstructionsTextDesktop != nil {
		return *m.UsageInstructionsTextDesktop
	}
	return ""
}

// This value will be shown in the offer text template as "XXX% cashback".
// Percentage has a range of (0, 100].
type AutofillOfferSpecifics_PercentageReward struct {
	// The string contains a number and a percent sign.
	Percentage           *string  `protobuf:"bytes,1,opt,name=percentage" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutofillOfferSpecifics_PercentageReward) Reset() {
	*m = AutofillOfferSpecifics_PercentageReward{}
}
func (m *AutofillOfferSpecifics_PercentageReward) String() string { return proto.CompactTextString(m) }
func (*AutofillOfferSpecifics_PercentageReward) ProtoMessage()    {}
func (*AutofillOfferSpecifics_PercentageReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_autofill_offer_specifics_d8a25529d974b566, []int{0, 3}
}
func (m *AutofillOfferSpecifics_PercentageReward) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutofillOfferSpecifics_PercentageReward.Unmarshal(m, b)
}
func (m *AutofillOfferSpecifics_PercentageReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutofillOfferSpecifics_PercentageReward.Marshal(b, m, deterministic)
}
func (dst *AutofillOfferSpecifics_PercentageReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutofillOfferSpecifics_PercentageReward.Merge(dst, src)
}
func (m *AutofillOfferSpecifics_PercentageReward) XXX_Size() int {
	return xxx_messageInfo_AutofillOfferSpecifics_PercentageReward.Size(m)
}
func (m *AutofillOfferSpecifics_PercentageReward) XXX_DiscardUnknown() {
	xxx_messageInfo_AutofillOfferSpecifics_PercentageReward.DiscardUnknown(m)
}

var xxx_messageInfo_AutofillOfferSpecifics_PercentageReward proto.InternalMessageInfo

func (m *AutofillOfferSpecifics_PercentageReward) GetPercentage() string {
	if m != nil && m.Percentage != nil {
		return *m.Percentage
	}
	return ""
}

// This value will be shown in the offer text template as "XXX$ off".
type AutofillOfferSpecifics_FixedAmountReward struct {
	// The string contains a number and a currency sign.
	Amount               *string  `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutofillOfferSpecifics_FixedAmountReward) Reset() {
	*m = AutofillOfferSpecifics_FixedAmountReward{}
}
func (m *AutofillOfferSpecifics_FixedAmountReward) String() string { return proto.CompactTextString(m) }
func (*AutofillOfferSpecifics_FixedAmountReward) ProtoMessage()    {}
func (*AutofillOfferSpecifics_FixedAmountReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_autofill_offer_specifics_d8a25529d974b566, []int{0, 4}
}
func (m *AutofillOfferSpecifics_FixedAmountReward) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutofillOfferSpecifics_FixedAmountReward.Unmarshal(m, b)
}
func (m *AutofillOfferSpecifics_FixedAmountReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutofillOfferSpecifics_FixedAmountReward.Marshal(b, m, deterministic)
}
func (dst *AutofillOfferSpecifics_FixedAmountReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutofillOfferSpecifics_FixedAmountReward.Merge(dst, src)
}
func (m *AutofillOfferSpecifics_FixedAmountReward) XXX_Size() int {
	return xxx_messageInfo_AutofillOfferSpecifics_FixedAmountReward.Size(m)
}
func (m *AutofillOfferSpecifics_FixedAmountReward) XXX_DiscardUnknown() {
	xxx_messageInfo_AutofillOfferSpecifics_FixedAmountReward.DiscardUnknown(m)
}

var xxx_messageInfo_AutofillOfferSpecifics_FixedAmountReward proto.InternalMessageInfo

func (m *AutofillOfferSpecifics_FixedAmountReward) GetAmount() string {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*AutofillOfferSpecifics)(nil), "sync_pb.AutofillOfferSpecifics")
	proto.RegisterType((*AutofillOfferSpecifics_CardLinkedOfferData)(nil), "sync_pb.AutofillOfferSpecifics.CardLinkedOfferData")
	proto.RegisterType((*AutofillOfferSpecifics_PromoCodeOfferData)(nil), "sync_pb.AutofillOfferSpecifics.PromoCodeOfferData")
	proto.RegisterType((*AutofillOfferSpecifics_DisplayStrings)(nil), "sync_pb.AutofillOfferSpecifics.DisplayStrings")
	proto.RegisterType((*AutofillOfferSpecifics_PercentageReward)(nil), "sync_pb.AutofillOfferSpecifics.PercentageReward")
	proto.RegisterType((*AutofillOfferSpecifics_FixedAmountReward)(nil), "sync_pb.AutofillOfferSpecifics.FixedAmountReward")
}

func init() {
	proto.RegisterFile("components/sync/protocol/autofill_offer_specifics.proto", fileDescriptor_autofill_offer_specifics_d8a25529d974b566)
}

var fileDescriptor_autofill_offer_specifics_d8a25529d974b566 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x69, 0x4f, 0x1b, 0x3b,
	0x14, 0x7d, 0x59, 0x58, 0x72, 0x79, 0x64, 0x31, 0xcb, 0x1b, 0xe5, 0xb5, 0x34, 0x6a, 0xd5, 0x36,
	0x2a, 0xd2, 0x40, 0x41, 0x15, 0x9f, 0x81, 0x80, 0x00, 0xb5, 0x6a, 0x34, 0xb4, 0xea, 0x47, 0xcb,
	0xd8, 0x37, 0xc1, 0x65, 0x66, 0x6c, 0x79, 0x9c, 0x36, 0xf9, 0x15, 0xfd, 0x33, 0xfd, 0x81, 0xd5,
	0x78, 0x26, 0x1b, 0xa1, 0xca, 0x47, 0x9f, 0x7b, 0xce, 0x3d, 0x57, 0x77, 0x31, 0x9c, 0x70, 0x15,
	0x69, 0x15, 0x63, 0x6c, 0x93, 0x83, 0x64, 0x14, 0xf3, 0x03, 0x6d, 0x94, 0x55, 0x5c, 0x85, 0x07,
	0x6c, 0x60, 0x55, 0x4f, 0x86, 0x21, 0x55, 0xbd, 0x1e, 0x1a, 0x9a, 0x68, 0xe4, 0xb2, 0x27, 0x79,
	0xe2, 0x3b, 0x06, 0x59, 0x4b, 0xd9, 0x54, 0xdf, 0xbd, 0xfc, 0x05, 0xb0, 0x7b, 0x9a, 0x73, 0x3f,
	0xa7, 0xd4, 0xdb, 0x31, 0x93, 0x54, 0xa1, 0x28, 0x85, 0x57, 0x68, 0x15, 0xda, 0xa5, 0xa0, 0x28,
	0x05, 0x79, 0x07, 0x8d, 0x2c, 0x99, 0x40, 0xcb, 0x64, 0x98, 0xd0, 0x81, 0x09, 0xbd, 0x62, 0xab,
	0xd0, 0xae, 0x04, 0x35, 0x17, 0xe8, 0x64, 0xf8, 0x57, 0x13, 0x92, 0xb7, 0x50, 0x8b, 0xd0, 0xf0,
	0x7b, 0x16, 0x5b, 0x2a, 0x54, 0xc4, 0x64, 0xec, 0x95, 0x5a, 0xa5, 0x76, 0x25, 0xa8, 0x8e, 0xe1,
	0x8e, 0x43, 0xc9, 0x21, 0x6c, 0x4f, 0x88, 0x4c, 0x6b, 0xaa, 0x19, 0x7f, 0x60, 0x7d, 0xf4, 0xca,
	0x8e, 0x4d, 0xc6, 0xb1, 0x53, 0xad, 0xbb, 0x59, 0x64, 0x5a, 0x06, 0x0e, 0xb5, 0x34, 0x23, 0x2a,
	0x98, 0x45, 0x6f, 0xc5, 0x55, 0x99, 0x95, 0x71, 0xe1, 0xf0, 0x0e, 0xb3, 0x48, 0xbe, 0xc3, 0x2e,
	0x67, 0x46, 0xd0, 0x50, 0xc6, 0x0f, 0x28, 0xf2, 0x5e, 0x08, 0x66, 0x99, 0xb7, 0xda, 0x2a, 0xb4,
	0x37, 0x8e, 0x8e, 0xfd, 0xbc, 0x0f, 0xfe, 0xd3, 0x3d, 0xf0, 0xcf, 0x99, 0x11, 0x1f, 0x9d, 0xd8,
	0x05, 0x3a, 0xcc, 0xb2, 0xab, 0x7f, 0x82, 0x2d, 0xbe, 0x08, 0x93, 0x3e, 0xec, 0x68, 0xa3, 0x22,
	0x45, 0xb9, 0x12, 0x38, 0x6b, 0x55, 0x71, 0x56, 0x47, 0xcb, 0xac, 0xba, 0xa9, 0xf8, 0x5c, 0x09,
	0x9c, 0x75, 0x22, 0x7a, 0x01, 0x25, 0xdf, 0xa0, 0x26, 0x64, 0xa2, 0x43, 0x36, 0xa2, 0x89, 0x35,
	0x32, 0xee, 0x27, 0x1e, 0x38, 0x0b, 0x7f, 0x99, 0x45, 0x27, 0x93, 0xdd, 0x66, 0xaa, 0xa0, 0x2a,
	0xe6, 0xde, 0x84, 0x42, 0x43, 0xa3, 0xe1, 0x18, 0x5b, 0xd6, 0x47, 0x6a, 0xf0, 0x27, 0x33, 0xc2,
	0x5b, 0x73, 0xa9, 0x0f, 0x97, 0x56, 0x3f, 0x11, 0x06, 0x4e, 0x77, 0x55, 0x08, 0xea, 0xfa, 0x11,
	0x46, 0x38, 0x6c, 0xf5, 0xe4, 0x10, 0x05, 0x65, 0x91, 0x1a, 0xc4, 0x76, 0x6c, 0xb1, 0xee, 0x2c,
	0xde, 0x2f, 0xb3, 0xb8, 0x4c, 0xa5, 0xa7, 0x4e, 0x39, 0xf1, 0x68, 0xf4, 0x1e, 0x83, 0xcd, 0x4b,
	0xd8, 0x7a, 0x62, 0x6a, 0xe4, 0x15, 0x6c, 0xca, 0x38, 0xb1, 0x66, 0x10, 0x61, 0x6c, 0xa9, 0x14,
	0x6e, 0x1f, 0x4b, 0xc1, 0xbf, 0x53, 0xf0, 0x5a, 0xdc, 0x94, 0xd7, 0x0b, 0xf5, 0xe2, 0x4d, 0x79,
	0xbd, 0x58, 0x2f, 0x35, 0x8f, 0x81, 0x2c, 0x8e, 0x84, 0x3c, 0x07, 0x98, 0x4e, 0xd9, 0x1d, 0x47,
	0x25, 0xa8, 0x4c, 0x86, 0xd4, 0xfc, 0x5d, 0x84, 0xea, 0x7c, 0x97, 0xc9, 0x1b, 0xa8, 0xfd, 0x60,
	0xe1, 0x00, 0xa9, 0x36, 0x4a, 0x53, 0x8b, 0x43, 0x9b, 0xcb, 0x36, 0x1d, 0xdc, 0x35, 0x4a, 0x7f,
	0xc1, 0xa1, 0x25, 0x1f, 0xe0, 0xbf, 0x04, 0x71, 0x72, 0x5c, 0x29, 0x91, 0x46, 0xea, 0x4e, 0x86,
	0x98, 0x1f, 0xd9, 0x76, 0x82, 0x98, 0x9f, 0x58, 0x2a, 0xf8, 0xe4, 0x62, 0xe4, 0x04, 0xbc, 0x05,
	0x99, 0xc0, 0xe4, 0xc1, 0x2a, 0xed, 0x95, 0x9c, 0x6e, 0x67, 0x5e, 0xd7, 0xc9, 0x82, 0xe4, 0x1c,
	0xf6, 0x06, 0x49, 0x3a, 0xe8, 0xac, 0x03, 0xdc, 0x4a, 0x15, 0xcf, 0xdb, 0x96, 0x9d, 0xfc, 0x7f,
	0xc7, 0xba, 0x9e, 0x21, 0xcd, 0xb8, 0x5f, 0xc0, 0x8b, 0xbf, 0x25, 0x19, 0x17, 0xb1, 0xe2, 0xb2,
	0x3c, 0x7b, 0x32, 0x4b, 0x5e, 0x4b, 0xf3, 0x08, 0xea, 0x8f, 0x17, 0x88, 0xec, 0x01, 0x4c, 0x17,
	0x28, 0x6f, 0xd9, 0x0c, 0xd2, 0xdc, 0x87, 0xc6, 0xc2, 0x46, 0x90, 0x5d, 0x58, 0xcd, 0x76, 0x2b,
	0x17, 0xe4, 0xaf, 0xb3, 0x26, 0x78, 0x76, 0xa4, 0x71, 0xf2, 0x0f, 0xce, 0xdc, 0xe7, 0xd9, 0x26,
	0x6c, 0x64, 0x8b, 0x48, 0x53, 0xca, 0xd9, 0x3e, 0xbc, 0x56, 0xa6, 0xef, 0xf3, 0x7b, 0xa3, 0x22,
	0x39, 0x88, 0xfc, 0xe9, 0x17, 0xeb, 0x16, 0xd4, 0x1f, 0x7f, 0xb1, 0x57, 0xa5, 0x6e, 0xe1, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0xd0, 0x9b, 0xb9, 0x81, 0x05, 0x00, 0x00,
}
