// Code generated by protoc-gen-go. DO NOT EDIT.
// source: docs.proto

package documents

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SignedEnvelope struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=Signature,proto3" json:"Signature,omitempty"`
	SignerCID            string   `protobuf:"bytes,2,opt,name=SignerCID,proto3" json:"SignerCID,omitempty"`
	Message              []byte   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedEnvelope) Reset()         { *m = SignedEnvelope{} }
func (m *SignedEnvelope) String() string { return proto.CompactTextString(m) }
func (*SignedEnvelope) ProtoMessage()    {}
func (*SignedEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{0}
}

func (m *SignedEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedEnvelope.Unmarshal(m, b)
}
func (m *SignedEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedEnvelope.Marshal(b, m, deterministic)
}
func (m *SignedEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedEnvelope.Merge(m, src)
}
func (m *SignedEnvelope) XXX_Size() int {
	return xxx_messageInfo_SignedEnvelope.Size(m)
}
func (m *SignedEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_SignedEnvelope proto.InternalMessageInfo

func (m *SignedEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedEnvelope) GetSignerCID() string {
	if m != nil {
		return m.SignerCID
	}
	return ""
}

func (m *SignedEnvelope) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type Envelope struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	EncryptedBody        []byte   `protobuf:"bytes,3,opt,name=EncryptedBody,proto3" json:"EncryptedBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{1}
}

func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Envelope) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Envelope) GetEncryptedBody() []byte {
	if m != nil {
		return m.EncryptedBody
	}
	return nil
}

type Header struct {
	IPFSID                string       `protobuf:"bytes,1,opt,name=IPFSID,proto3" json:"IPFSID,omitempty"`
	Version               float32      `protobuf:"fixed32,2,opt,name=Version,proto3" json:"Version,omitempty"`
	DateTime              int64        `protobuf:"varint,3,opt,name=DateTime,proto3" json:"DateTime,omitempty"`
	PreviousCID           string       `protobuf:"bytes,4,opt,name=PreviousCID,proto3" json:"PreviousCID,omitempty"`
	BodyTypeCode          float32      `protobuf:"fixed32,5,opt,name=BodyTypeCode,proto3" json:"BodyTypeCode,omitempty"`
	BodyVersion           float32      `protobuf:"fixed32,6,opt,name=BodyVersion,proto3" json:"BodyVersion,omitempty"`
	EncryptedBodyTypeCode float32      `protobuf:"fixed32,7,opt,name=EncryptedBodyTypeCode,proto3" json:"EncryptedBodyTypeCode,omitempty"`
	EncryptedBodyVersion  float32      `protobuf:"fixed32,8,opt,name=EncryptedBodyVersion,proto3" json:"EncryptedBodyVersion,omitempty"`
	EncryptedBodyIV       []byte       `protobuf:"bytes,9,opt,name=EncryptedBodyIV,proto3" json:"EncryptedBodyIV,omitempty"`
	Recipients            []*Recipient `protobuf:"bytes,10,rep,name=Recipients,proto3" json:"Recipients,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}     `json:"-"`
	XXX_unrecognized      []byte       `json:"-"`
	XXX_sizecache         int32        `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{2}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetIPFSID() string {
	if m != nil {
		return m.IPFSID
	}
	return ""
}

func (m *Header) GetVersion() float32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetDateTime() int64 {
	if m != nil {
		return m.DateTime
	}
	return 0
}

func (m *Header) GetPreviousCID() string {
	if m != nil {
		return m.PreviousCID
	}
	return ""
}

func (m *Header) GetBodyTypeCode() float32 {
	if m != nil {
		return m.BodyTypeCode
	}
	return 0
}

func (m *Header) GetBodyVersion() float32 {
	if m != nil {
		return m.BodyVersion
	}
	return 0
}

func (m *Header) GetEncryptedBodyTypeCode() float32 {
	if m != nil {
		return m.EncryptedBodyTypeCode
	}
	return 0
}

func (m *Header) GetEncryptedBodyVersion() float32 {
	if m != nil {
		return m.EncryptedBodyVersion
	}
	return 0
}

func (m *Header) GetEncryptedBodyIV() []byte {
	if m != nil {
		return m.EncryptedBodyIV
	}
	return nil
}

func (m *Header) GetRecipients() []*Recipient {
	if m != nil {
		return m.Recipients
	}
	return nil
}

type Recipient struct {
	Version              float32  `protobuf:"fixed32,1,opt,name=Version,proto3" json:"Version,omitempty"`
	CID                  string   `protobuf:"bytes,2,opt,name=CID,proto3" json:"CID,omitempty"`
	EncapsulatedKey      []byte   `protobuf:"bytes,3,opt,name=EncapsulatedKey,proto3" json:"EncapsulatedKey,omitempty"`
	CipherText           []byte   `protobuf:"bytes,4,opt,name=CipherText,proto3" json:"CipherText,omitempty"`
	IV                   []byte   `protobuf:"bytes,5,opt,name=IV,proto3" json:"IV,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recipient) Reset()         { *m = Recipient{} }
func (m *Recipient) String() string { return proto.CompactTextString(m) }
func (*Recipient) ProtoMessage()    {}
func (*Recipient) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{3}
}

func (m *Recipient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipient.Unmarshal(m, b)
}
func (m *Recipient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipient.Marshal(b, m, deterministic)
}
func (m *Recipient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipient.Merge(m, src)
}
func (m *Recipient) XXX_Size() int {
	return xxx_messageInfo_Recipient.Size(m)
}
func (m *Recipient) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipient.DiscardUnknown(m)
}

var xxx_messageInfo_Recipient proto.InternalMessageInfo

func (m *Recipient) GetVersion() float32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Recipient) GetCID() string {
	if m != nil {
		return m.CID
	}
	return ""
}

func (m *Recipient) GetEncapsulatedKey() []byte {
	if m != nil {
		return m.EncapsulatedKey
	}
	return nil
}

func (m *Recipient) GetCipherText() []byte {
	if m != nil {
		return m.CipherText
	}
	return nil
}

func (m *Recipient) GetIV() []byte {
	if m != nil {
		return m.IV
	}
	return nil
}

type IDDocument struct {
	AuthenticationReference string   `protobuf:"bytes,1,opt,name=AuthenticationReference,proto3" json:"AuthenticationReference,omitempty"`
	BeneficiaryECPublicKey  []byte   `protobuf:"bytes,2,opt,name=BeneficiaryECPublicKey,proto3" json:"BeneficiaryECPublicKey,omitempty"`
	SikePublicKey           []byte   `protobuf:"bytes,3,opt,name=SikePublicKey,proto3" json:"SikePublicKey,omitempty"`
	BLSPublicKey            []byte   `protobuf:"bytes,4,opt,name=BLSPublicKey,proto3" json:"BLSPublicKey,omitempty"`
	Timestamp               int64    `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *IDDocument) Reset()         { *m = IDDocument{} }
func (m *IDDocument) String() string { return proto.CompactTextString(m) }
func (*IDDocument) ProtoMessage()    {}
func (*IDDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{4}
}

func (m *IDDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDDocument.Unmarshal(m, b)
}
func (m *IDDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDDocument.Marshal(b, m, deterministic)
}
func (m *IDDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDDocument.Merge(m, src)
}
func (m *IDDocument) XXX_Size() int {
	return xxx_messageInfo_IDDocument.Size(m)
}
func (m *IDDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_IDDocument.DiscardUnknown(m)
}

var xxx_messageInfo_IDDocument proto.InternalMessageInfo

func (m *IDDocument) GetAuthenticationReference() string {
	if m != nil {
		return m.AuthenticationReference
	}
	return ""
}

func (m *IDDocument) GetBeneficiaryECPublicKey() []byte {
	if m != nil {
		return m.BeneficiaryECPublicKey
	}
	return nil
}

func (m *IDDocument) GetSikePublicKey() []byte {
	if m != nil {
		return m.SikePublicKey
	}
	return nil
}

func (m *IDDocument) GetBLSPublicKey() []byte {
	if m != nil {
		return m.BLSPublicKey
	}
	return nil
}

func (m *IDDocument) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type OrderDocument struct {
	Type                 string      `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Coin                 int64       `protobuf:"varint,2,opt,name=Coin,proto3" json:"Coin,omitempty"`
	PrincipalCID         string      `protobuf:"bytes,3,opt,name=PrincipalCID,proto3" json:"PrincipalCID,omitempty"`
	BeneficiaryCID       string      `protobuf:"bytes,4,opt,name=BeneficiaryCID,proto3" json:"BeneficiaryCID,omitempty"`
	Reference            string      `protobuf:"bytes,5,opt,name=Reference,proto3" json:"Reference,omitempty"`
	Timestamp            int64       `protobuf:"varint,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	OrderPart2           *OrderPart2 `protobuf:"bytes,7,opt,name=OrderPart2,proto3" json:"OrderPart2,omitempty"`
	OrderPart3           *OrderPart3 `protobuf:"bytes,8,opt,name=OrderPart3,proto3" json:"OrderPart3,omitempty"`
	OrderPart4           *OrderPart4 `protobuf:"bytes,9,opt,name=OrderPart4,proto3" json:"OrderPart4,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OrderDocument) Reset()         { *m = OrderDocument{} }
func (m *OrderDocument) String() string { return proto.CompactTextString(m) }
func (*OrderDocument) ProtoMessage()    {}
func (*OrderDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{5}
}

func (m *OrderDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderDocument.Unmarshal(m, b)
}
func (m *OrderDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderDocument.Marshal(b, m, deterministic)
}
func (m *OrderDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderDocument.Merge(m, src)
}
func (m *OrderDocument) XXX_Size() int {
	return xxx_messageInfo_OrderDocument.Size(m)
}
func (m *OrderDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderDocument.DiscardUnknown(m)
}

var xxx_messageInfo_OrderDocument proto.InternalMessageInfo

func (m *OrderDocument) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OrderDocument) GetCoin() int64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

func (m *OrderDocument) GetPrincipalCID() string {
	if m != nil {
		return m.PrincipalCID
	}
	return ""
}

func (m *OrderDocument) GetBeneficiaryCID() string {
	if m != nil {
		return m.BeneficiaryCID
	}
	return ""
}

func (m *OrderDocument) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *OrderDocument) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *OrderDocument) GetOrderPart2() *OrderPart2 {
	if m != nil {
		return m.OrderPart2
	}
	return nil
}

func (m *OrderDocument) GetOrderPart3() *OrderPart3 {
	if m != nil {
		return m.OrderPart3
	}
	return nil
}

func (m *OrderDocument) GetOrderPart4() *OrderPart4 {
	if m != nil {
		return m.OrderPart4
	}
	return nil
}

type OrderPart2 struct {
	CommitmentPublicKey  string   `protobuf:"bytes,1,opt,name=CommitmentPublicKey,proto3" json:"CommitmentPublicKey,omitempty"`
	PreviousOrderCID     string   `protobuf:"bytes,2,opt,name=PreviousOrderCID,proto3" json:"PreviousOrderCID,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderPart2) Reset()         { *m = OrderPart2{} }
func (m *OrderPart2) String() string { return proto.CompactTextString(m) }
func (*OrderPart2) ProtoMessage()    {}
func (*OrderPart2) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{6}
}

func (m *OrderPart2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPart2.Unmarshal(m, b)
}
func (m *OrderPart2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPart2.Marshal(b, m, deterministic)
}
func (m *OrderPart2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPart2.Merge(m, src)
}
func (m *OrderPart2) XXX_Size() int {
	return xxx_messageInfo_OrderPart2.Size(m)
}
func (m *OrderPart2) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPart2.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPart2 proto.InternalMessageInfo

func (m *OrderPart2) GetCommitmentPublicKey() string {
	if m != nil {
		return m.CommitmentPublicKey
	}
	return ""
}

func (m *OrderPart2) GetPreviousOrderCID() string {
	if m != nil {
		return m.PreviousOrderCID
	}
	return ""
}

func (m *OrderPart2) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type OrderPart3 struct {
	Redemption               string   `protobuf:"bytes,1,opt,name=Redemption,proto3" json:"Redemption,omitempty"`
	PreviousOrderCID         string   `protobuf:"bytes,2,opt,name=PreviousOrderCID,proto3" json:"PreviousOrderCID,omitempty"`
	BeneficiaryEncryptedData []byte   `protobuf:"bytes,3,opt,name=BeneficiaryEncryptedData,proto3" json:"BeneficiaryEncryptedData,omitempty"`
	Timestamp                int64    `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *OrderPart3) Reset()         { *m = OrderPart3{} }
func (m *OrderPart3) String() string { return proto.CompactTextString(m) }
func (*OrderPart3) ProtoMessage()    {}
func (*OrderPart3) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{7}
}

func (m *OrderPart3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPart3.Unmarshal(m, b)
}
func (m *OrderPart3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPart3.Marshal(b, m, deterministic)
}
func (m *OrderPart3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPart3.Merge(m, src)
}
func (m *OrderPart3) XXX_Size() int {
	return xxx_messageInfo_OrderPart3.Size(m)
}
func (m *OrderPart3) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPart3.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPart3 proto.InternalMessageInfo

func (m *OrderPart3) GetRedemption() string {
	if m != nil {
		return m.Redemption
	}
	return ""
}

func (m *OrderPart3) GetPreviousOrderCID() string {
	if m != nil {
		return m.PreviousOrderCID
	}
	return ""
}

func (m *OrderPart3) GetBeneficiaryEncryptedData() []byte {
	if m != nil {
		return m.BeneficiaryEncryptedData
	}
	return nil
}

func (m *OrderPart3) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type OrderPart4 struct {
	Secret               string   `protobuf:"bytes,1,opt,name=Secret,proto3" json:"Secret,omitempty"`
	PreviousOrderCID     string   `protobuf:"bytes,2,opt,name=PreviousOrderCID,proto3" json:"PreviousOrderCID,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderPart4) Reset()         { *m = OrderPart4{} }
func (m *OrderPart4) String() string { return proto.CompactTextString(m) }
func (*OrderPart4) ProtoMessage()    {}
func (*OrderPart4) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{8}
}

func (m *OrderPart4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderPart4.Unmarshal(m, b)
}
func (m *OrderPart4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderPart4.Marshal(b, m, deterministic)
}
func (m *OrderPart4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderPart4.Merge(m, src)
}
func (m *OrderPart4) XXX_Size() int {
	return xxx_messageInfo_OrderPart4.Size(m)
}
func (m *OrderPart4) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderPart4.DiscardUnknown(m)
}

var xxx_messageInfo_OrderPart4 proto.InternalMessageInfo

func (m *OrderPart4) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *OrderPart4) GetPreviousOrderCID() string {
	if m != nil {
		return m.PreviousOrderCID
	}
	return ""
}

func (m *OrderPart4) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Policy struct {
	Version              float32  `protobuf:"fixed32,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{9}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetVersion() float32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PlainTestMessage1 struct {
	Nametest1            string   `protobuf:"bytes,1,opt,name=Nametest1,proto3" json:"Nametest1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlainTestMessage1) Reset()         { *m = PlainTestMessage1{} }
func (m *PlainTestMessage1) String() string { return proto.CompactTextString(m) }
func (*PlainTestMessage1) ProtoMessage()    {}
func (*PlainTestMessage1) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{10}
}

func (m *PlainTestMessage1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainTestMessage1.Unmarshal(m, b)
}
func (m *PlainTestMessage1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainTestMessage1.Marshal(b, m, deterministic)
}
func (m *PlainTestMessage1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTestMessage1.Merge(m, src)
}
func (m *PlainTestMessage1) XXX_Size() int {
	return xxx_messageInfo_PlainTestMessage1.Size(m)
}
func (m *PlainTestMessage1) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTestMessage1.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTestMessage1 proto.InternalMessageInfo

func (m *PlainTestMessage1) GetNametest1() string {
	if m != nil {
		return m.Nametest1
	}
	return ""
}

type EncryptTestMessage1 struct {
	Nametest2            string   `protobuf:"bytes,1,opt,name=Nametest2,proto3" json:"Nametest2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptTestMessage1) Reset()         { *m = EncryptTestMessage1{} }
func (m *EncryptTestMessage1) String() string { return proto.CompactTextString(m) }
func (*EncryptTestMessage1) ProtoMessage()    {}
func (*EncryptTestMessage1) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{11}
}

func (m *EncryptTestMessage1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptTestMessage1.Unmarshal(m, b)
}
func (m *EncryptTestMessage1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptTestMessage1.Marshal(b, m, deterministic)
}
func (m *EncryptTestMessage1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptTestMessage1.Merge(m, src)
}
func (m *EncryptTestMessage1) XXX_Size() int {
	return xxx_messageInfo_EncryptTestMessage1.Size(m)
}
func (m *EncryptTestMessage1) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptTestMessage1.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptTestMessage1 proto.InternalMessageInfo

func (m *EncryptTestMessage1) GetNametest2() string {
	if m != nil {
		return m.Nametest2
	}
	return ""
}

type SimpleString struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleString) Reset()         { *m = SimpleString{} }
func (m *SimpleString) String() string { return proto.CompactTextString(m) }
func (*SimpleString) ProtoMessage()    {}
func (*SimpleString) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a25dace11219bce, []int{12}
}

func (m *SimpleString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleString.Unmarshal(m, b)
}
func (m *SimpleString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleString.Marshal(b, m, deterministic)
}
func (m *SimpleString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleString.Merge(m, src)
}
func (m *SimpleString) XXX_Size() int {
	return xxx_messageInfo_SimpleString.Size(m)
}
func (m *SimpleString) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleString.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleString proto.InternalMessageInfo

func (m *SimpleString) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*SignedEnvelope)(nil), "documents.SignedEnvelope")
	proto.RegisterType((*Envelope)(nil), "documents.Envelope")
	proto.RegisterType((*Header)(nil), "documents.Header")
	proto.RegisterType((*Recipient)(nil), "documents.Recipient")
	proto.RegisterType((*IDDocument)(nil), "documents.IDDocument")
	proto.RegisterType((*OrderDocument)(nil), "documents.OrderDocument")
	proto.RegisterType((*OrderPart2)(nil), "documents.OrderPart2")
	proto.RegisterType((*OrderPart3)(nil), "documents.OrderPart3")
	proto.RegisterType((*OrderPart4)(nil), "documents.OrderPart4")
	proto.RegisterType((*Policy)(nil), "documents.Policy")
	proto.RegisterType((*PlainTestMessage1)(nil), "documents.PlainTestMessage1")
	proto.RegisterType((*EncryptTestMessage1)(nil), "documents.EncryptTestMessage1")
	proto.RegisterType((*SimpleString)(nil), "documents.SimpleString")
}

func init() { proto.RegisterFile("docs.proto", fileDescriptor_2a25dace11219bce) }

var fileDescriptor_2a25dace11219bce = []byte{
	// 956 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0xda, 0x8e, 0x13, 0xbf, 0xb8, 0x21, 0x9d, 0x26, 0xed, 0xaa, 0x20, 0x12, 0xad, 0x22,
	0x64, 0x24, 0xe2, 0x10, 0xdb, 0x8d, 0x20, 0x42, 0x42, 0xd8, 0x0e, 0x60, 0xf1, 0xcf, 0x8c, 0xa3,
	0x08, 0xa9, 0x6a, 0xa5, 0xc9, 0xee, 0x8b, 0x33, 0xaa, 0x77, 0x67, 0xb5, 0x3b, 0x4e, 0xb1, 0x80,
	0x33, 0x07, 0x8e, 0x9c, 0xb9, 0x20, 0xf5, 0xca, 0x11, 0xbe, 0x00, 0x47, 0xbe, 0x43, 0x44, 0x0e,
	0x20, 0xbe, 0x01, 0x37, 0xd0, 0xcc, 0xfe, 0x0f, 0x71, 0x93, 0x03, 0x52, 0x7d, 0x9a, 0xf7, 0x7e,
	0xef, 0xf7, 0xe6, 0xfd, 0x1d, 0x2f, 0x80, 0x23, 0xec, 0xb0, 0xe9, 0x07, 0x42, 0x0a, 0x52, 0x73,
	0x84, 0x3d, 0x75, 0xd1, 0x93, 0xe1, 0xfd, 0xbd, 0x31, 0x97, 0xa7, 0xd3, 0xe3, 0xa6, 0x2d, 0xdc,
	0x1d, 0xf7, 0x29, 0x97, 0x4f, 0xc4, 0xd3, 0x9d, 0xb1, 0xd8, 0xd6, 0x76, 0xdb, 0x67, 0x6c, 0xc2,
	0x1d, 0x26, 0x45, 0x10, 0xee, 0xa4, 0xc7, 0xc8, 0xc5, 0xfd, 0xed, 0x1c, 0x6f, 0x2c, 0xc6, 0x62,
	0x47, 0xab, 0x8f, 0xa7, 0x27, 0x5a, 0xd2, 0x82, 0x3e, 0x45, 0xe6, 0xd6, 0x77, 0x06, 0xac, 0x8c,
	0xf8, 0xd8, 0x43, 0xe7, 0xc0, 0x3b, 0xc3, 0x89, 0xf0, 0x91, 0x6c, 0x41, 0x4d, 0x69, 0x98, 0x9c,
	0x06, 0x68, 0x1a, 0x9b, 0x46, 0xa3, 0xde, 0xad, 0x5e, 0x9c, 0x6f, 0x94, 0xfc, 0x35, 0x9a, 0x01,
	0xe4, 0xed, 0xc8, 0x0a, 0x83, 0xde, 0xa0, 0x6f, 0x96, 0x36, 0x8d, 0x46, 0xad, 0xfb, 0xf2, 0xc5,
	0xf9, 0xc6, 0x3d, 0x58, 0x7f, 0xfc, 0xf9, 0xc3, 0x87, 0xfb, 0x6c, 0xe2, 0x4d, 0xdd, 0xfd, 0x47,
	0x8f, 0xbe, 0xea, 0x3c, 0xf8, 0x66, 0xeb, 0xeb, 0xc7, 0x5b, 0x34, 0xb3, 0x26, 0x26, 0x2c, 0x7e,
	0x82, 0x61, 0xc8, 0xc6, 0x68, 0x96, 0x95, 0x7b, 0x9a, 0x88, 0x96, 0x80, 0xa5, 0x34, 0x8c, 0xd7,
	0xa1, 0xfa, 0x21, 0x32, 0x07, 0x03, 0x1d, 0xc3, 0x72, 0xeb, 0x76, 0x33, 0x2d, 0x4e, 0x33, 0x02,
	0x68, 0x6c, 0x40, 0x08, 0x54, 0xba, 0xc2, 0x99, 0xe9, 0x30, 0xea, 0x54, 0x9f, 0xc9, 0x16, 0xdc,
	0x3a, 0xf0, 0xec, 0x60, 0xe6, 0x4b, 0x74, 0x34, 0x18, 0x5d, 0x55, 0x54, 0x5a, 0x3f, 0x96, 0x93,
	0x5b, 0xc8, 0x5d, 0xa8, 0x0e, 0x86, 0xef, 0x8f, 0x06, 0x7d, 0x7d, 0x5f, 0x8d, 0xc6, 0x92, 0x8a,
	0xf6, 0x08, 0x83, 0x90, 0x0b, 0x4f, 0xfb, 0x2f, 0xd1, 0x44, 0x24, 0x6f, 0xc0, 0x52, 0x9f, 0x49,
	0x3c, 0xe4, 0x6e, 0x94, 0x48, 0xb9, 0xbb, 0x7a, 0x71, 0xbe, 0x51, 0x5f, 0x7d, 0xf6, 0xed, 0x1f,
	0x7f, 0x2d, 0x98, 0xcf, 0x7e, 0xfd, 0xf9, 0xfb, 0x19, 0x4d, 0x2d, 0xc8, 0x26, 0x2c, 0x0f, 0x03,
	0x3c, 0xe3, 0x62, 0x1a, 0xaa, 0x92, 0x55, 0xf4, 0x25, 0x79, 0x15, 0xb1, 0xa0, 0xae, 0x82, 0x3a,
	0x9c, 0xf9, 0xd8, 0x13, 0x0e, 0x9a, 0x0b, 0xfa, 0xba, 0x82, 0x4e, 0x79, 0x51, 0x72, 0x12, 0x51,
	0x55, 0x9b, 0xe4, 0x55, 0xa4, 0x03, 0xeb, 0x85, 0x1c, 0x53, 0x77, 0x8b, 0xda, 0xf6, 0x6a, 0x90,
	0xb4, 0x60, 0xad, 0x00, 0x24, 0x17, 0x2c, 0x69, 0xd2, 0x95, 0x18, 0x69, 0xc0, 0x4b, 0x05, 0xfd,
	0xe0, 0xc8, 0xac, 0xe9, 0x22, 0x5f, 0x56, 0x93, 0x77, 0x00, 0x28, 0xda, 0xdc, 0xe7, 0xaa, 0x7b,
	0x26, 0x6c, 0x96, 0x1b, 0xcb, 0xad, 0xb5, 0x5c, 0x3f, 0x53, 0x30, 0x9a, 0xb4, 0xd3, 0x35, 0x9a,
	0xb3, 0xb7, 0x7e, 0x32, 0xa0, 0x96, 0x8a, 0xf9, 0x7e, 0x18, 0xc5, 0x7e, 0x6c, 0x43, 0xf9, 0x86,
	0xc3, 0xa8, 0xec, 0xe2, 0xf0, 0x99, 0x1f, 0x4e, 0x27, 0x4c, 0xa2, 0xf3, 0x11, 0x26, 0x33, 0x72,
	0x59, 0x4d, 0x5e, 0x05, 0xe8, 0x71, 0xff, 0x14, 0x83, 0x43, 0xfc, 0x52, 0xea, 0xce, 0xd5, 0x69,
	0x4e, 0x43, 0x56, 0xa0, 0x34, 0x38, 0xd2, 0xed, 0xaa, 0xd3, 0xd2, 0xe0, 0xc8, 0xfa, 0xdb, 0x00,
	0x18, 0xf4, 0xfb, 0x71, 0x7a, 0xe4, 0x2d, 0xb8, 0xf7, 0xde, 0x54, 0x9e, 0xa2, 0x27, 0xb9, 0xcd,
	0x24, 0x17, 0x1e, 0xc5, 0x13, 0x0c, 0xd0, 0xb3, 0x31, 0x1e, 0xb5, 0x79, 0x30, 0xd9, 0x83, 0xbb,
	0x5d, 0xf4, 0xf0, 0x84, 0xdb, 0x9c, 0x05, 0xb3, 0x83, 0xde, 0x70, 0x7a, 0x3c, 0xe1, 0xb6, 0x8a,
	0x34, 0x1a, 0xf5, 0x39, 0xa8, 0x1a, 0xfe, 0x11, 0x7f, 0x82, 0x99, 0x79, 0x3c, 0xfc, 0x05, 0xa5,
	0x9e, 0xb7, 0x8f, 0x47, 0x99, 0x51, 0x94, 0x58, 0x41, 0x47, 0x9a, 0x50, 0x53, 0xd3, 0x1b, 0x4a,
	0xe6, 0xfa, 0x3a, 0xc3, 0xab, 0x86, 0x3c, 0x33, 0xb1, 0x7e, 0x2b, 0xc3, 0xad, 0xcf, 0x02, 0x07,
	0x83, 0x34, 0x7b, 0x02, 0x15, 0x35, 0x65, 0x71, 0xaa, 0xfa, 0x4c, 0x5e, 0x83, 0x4a, 0x4f, 0xf0,
	0x68, 0xa1, 0xca, 0x5d, 0x72, 0x71, 0xbe, 0xb1, 0xb2, 0xfa, 0x4f, 0xf2, 0x33, 0xcc, 0x3f, 0x17,
	0xa9, 0xc6, 0xc9, 0xbb, 0x50, 0x1f, 0x06, 0xdc, 0xb3, 0xb9, 0xcf, 0x26, 0xaa, 0xb5, 0xe5, 0xeb,
	0x5b, 0x5b, 0x20, 0x90, 0x1e, 0xac, 0xe4, 0x4a, 0x94, 0xee, 0xdd, 0xf3, 0x5d, 0x5c, 0xa2, 0xa8,
	0x07, 0x31, 0xeb, 0xd8, 0x82, 0xe6, 0xeb, 0x31, 0xfd, 0xc2, 0xa0, 0x19, 0x50, 0xac, 0x54, 0xf5,
	0xda, 0x4a, 0x91, 0x07, 0x00, 0xba, 0x50, 0x43, 0x16, 0xc8, 0x96, 0x5e, 0xce, 0xe5, 0xd6, 0x7a,
	0x6e, 0x27, 0x32, 0x90, 0xe6, 0x0c, 0x0b, 0xb4, 0xb6, 0x5e, 0xcf, 0x39, 0xb4, 0x76, 0x8e, 0xd6,
	0x2e, 0xd0, 0x3a, 0x7a, 0x4d, 0xe7, 0xd0, 0x3a, 0x39, 0x5a, 0xc7, 0xfa, 0xc5, 0xc8, 0x47, 0x49,
	0xde, 0x84, 0x3b, 0x3d, 0xe1, 0xba, 0x5c, 0x2a, 0x52, 0x36, 0x38, 0x51, 0x6b, 0xaf, 0x82, 0xc8,
	0x07, 0xb0, 0x9a, 0x3c, 0x71, 0xda, 0xcf, 0x0d, 0x17, 0xf4, 0x3f, 0xa4, 0x62, 0x79, 0xcb, 0xd7,
	0x0f, 0xe2, 0xef, 0xf9, 0xc8, 0xdb, 0x6a, 0x85, 0x29, 0x3a, 0xe8, 0xfa, 0x32, 0x79, 0x38, 0x6a,
	0x34, 0xa7, 0xf9, 0xff, 0xe2, 0xdc, 0x07, 0x33, 0xbf, 0x94, 0xc9, 0x43, 0xd8, 0x67, 0x92, 0xc5,
	0x5b, 0x38, 0x17, 0x2f, 0xe6, 0x58, 0xb9, 0x3e, 0xc7, 0x1f, 0xf2, 0x39, 0x76, 0xd4, 0x3f, 0xd8,
	0x08, 0xed, 0x00, 0x65, 0xf2, 0x0f, 0x16, 0x49, 0x2f, 0xae, 0x07, 0x7b, 0x50, 0x1d, 0x8a, 0x09,
	0xb7, 0x67, 0xcf, 0x79, 0xb4, 0x09, 0x54, 0x3e, 0x65, 0x2e, 0x46, 0x01, 0x51, 0x7d, 0xb6, 0x76,
	0xe1, 0xf6, 0x70, 0xc2, 0xb8, 0x77, 0x88, 0xa1, 0x8c, 0x3f, 0x0d, 0x76, 0xc9, 0x2b, 0x50, 0x53,
	0xa0, 0xc4, 0x50, 0xee, 0xc6, 0x09, 0x66, 0x0a, 0xab, 0x0d, 0x77, 0xe2, 0x5a, 0xce, 0x23, 0xb5,
	0x2e, 0x93, 0x5a, 0x56, 0x03, 0xea, 0x23, 0xee, 0xfa, 0x13, 0x1c, 0xc9, 0x80, 0x7b, 0x63, 0x15,
	0x65, 0x4f, 0x78, 0x12, 0xbd, 0xa4, 0x82, 0x89, 0x78, 0x5c, 0xd5, 0x5f, 0x4b, 0xed, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x54, 0x80, 0x91, 0x77, 0xad, 0x09, 0x00, 0x00,
}
