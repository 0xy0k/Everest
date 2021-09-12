// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/gov/network_properties.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NetworkProperty int32

const (
	MinTxFee                    NetworkProperty = 0
	MaxTxFee                    NetworkProperty = 1
	VoteQuorum                  NetworkProperty = 2
	ProposalEndTime             NetworkProperty = 3
	ProposalEnactmentTime       NetworkProperty = 4
	MinProposalEndBlocks        NetworkProperty = 5
	MinProposalEnactmentBlocks  NetworkProperty = 6
	EnableForeignFeePayments    NetworkProperty = 7
	MischanceRankDecreaseAmount NetworkProperty = 8
	MaxMischance                NetworkProperty = 9
	MischanceConfidence         NetworkProperty = 10
	InactiveRankDecreasePercent NetworkProperty = 11
	PoorNetworkMaxBankSend      NetworkProperty = 12
	MinValidators               NetworkProperty = 13
	JailMaxTime                 NetworkProperty = 14
	EnableTokenWhitelist        NetworkProperty = 15
	EnableTokenBlacklist        NetworkProperty = 16
	MinIdentityApprovalTip      NetworkProperty = 17
)

var NetworkProperty_name = map[int32]string{
	0:  "MIN_TX_FEE",
	1:  "MAX_TX_FEE",
	2:  "VOTE_QUORUM",
	3:  "PROPOSAL_END_TIME",
	4:  "PROPOSAL_ENACTMENT_TIME",
	5:  "MIN_PROPOSAL_END_BLOCKS",
	6:  "MIN_PROPOSAL_ENACTMENT_BLOCKS",
	7:  "ENABLE_FOREIGN_FEE_PAYMENTS",
	8:  "MISCHANCE_RANK_DECREASE_AMOUNT",
	9:  "MAX_MISCHANCE",
	10: "MISCHANCE_CONFIDENCE",
	11: "INACTIVE_RANK_DECREASE_PERCENT",
	12: "POOR_NETWORK_MAX_BANK_SEND",
	13: "MIN_VALIDATORS",
	14: "JAIL_MAX_TIME",
	15: "ENABLE_TOKEN_WHITELIST",
	16: "ENABLE_TOKEN_BLACKLIST",
	17: "MIN_IDENTITY_APPROVAL_TIP",
}

var NetworkProperty_value = map[string]int32{
	"MIN_TX_FEE":                     0,
	"MAX_TX_FEE":                     1,
	"VOTE_QUORUM":                    2,
	"PROPOSAL_END_TIME":              3,
	"PROPOSAL_ENACTMENT_TIME":        4,
	"MIN_PROPOSAL_END_BLOCKS":        5,
	"MIN_PROPOSAL_ENACTMENT_BLOCKS":  6,
	"ENABLE_FOREIGN_FEE_PAYMENTS":    7,
	"MISCHANCE_RANK_DECREASE_AMOUNT": 8,
	"MAX_MISCHANCE":                  9,
	"MISCHANCE_CONFIDENCE":           10,
	"INACTIVE_RANK_DECREASE_PERCENT": 11,
	"POOR_NETWORK_MAX_BANK_SEND":     12,
	"MIN_VALIDATORS":                 13,
	"JAIL_MAX_TIME":                  14,
	"ENABLE_TOKEN_WHITELIST":         15,
	"ENABLE_TOKEN_BLACKLIST":         16,
	"MIN_IDENTITY_APPROVAL_TIP":      17,
}

func (x NetworkProperty) String() string {
	return proto.EnumName(NetworkProperty_name, int32(x))
}

func (NetworkProperty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}

type MsgSetNetworkProperties struct {
	NetworkProperties *NetworkProperties                            `protobuf:"bytes,1,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	Proposer          github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
}

func (m *MsgSetNetworkProperties) Reset()         { *m = MsgSetNetworkProperties{} }
func (m *MsgSetNetworkProperties) String() string { return proto.CompactTextString(m) }
func (*MsgSetNetworkProperties) ProtoMessage()    {}
func (*MsgSetNetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}
func (m *MsgSetNetworkProperties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetNetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetNetworkProperties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetNetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNetworkProperties.Merge(m, src)
}
func (m *MsgSetNetworkProperties) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetNetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNetworkProperties proto.InternalMessageInfo

func (m *MsgSetNetworkProperties) GetNetworkProperties() *NetworkProperties {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *MsgSetNetworkProperties) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type NetworkProperties struct {
	MinTxFee                    uint64 `protobuf:"varint,1,opt,name=min_tx_fee,json=minTxFee,proto3" json:"min_tx_fee,omitempty"`
	MaxTxFee                    uint64 `protobuf:"varint,2,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee,omitempty"`
	VoteQuorum                  uint64 `protobuf:"varint,3,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	ProposalEndTime             uint64 `protobuf:"varint,4,opt,name=proposal_end_time,json=proposalEndTime,proto3" json:"proposal_end_time,omitempty"`
	ProposalEnactmentTime       uint64 `protobuf:"varint,5,opt,name=proposal_enactment_time,json=proposalEnactmentTime,proto3" json:"proposal_enactment_time,omitempty"`
	MinProposalEndBlocks        uint64 `protobuf:"varint,6,opt,name=min_proposal_end_blocks,json=minProposalEndBlocks,proto3" json:"min_proposal_end_blocks,omitempty"`
	MinProposalEnactmentBlocks  uint64 `protobuf:"varint,7,opt,name=min_proposal_enactment_blocks,json=minProposalEnactmentBlocks,proto3" json:"min_proposal_enactment_blocks,omitempty"`
	EnableForeignFeePayments    bool   `protobuf:"varint,8,opt,name=enable_foreign_fee_payments,json=enableForeignFeePayments,proto3" json:"enable_foreign_fee_payments,omitempty"`
	MischanceRankDecreaseAmount uint64 `protobuf:"varint,9,opt,name=mischance_rank_decrease_amount,json=mischanceRankDecreaseAmount,proto3" json:"mischance_rank_decrease_amount,omitempty"`
	MaxMischance                uint64 `protobuf:"varint,10,opt,name=max_mischance,json=maxMischance,proto3" json:"max_mischance,omitempty"`
	MischanceConfidence         uint64 `protobuf:"varint,11,opt,name=mischance_confidence,json=mischanceConfidence,proto3" json:"mischance_confidence,omitempty"`
	InactiveRankDecreasePercent uint64 `protobuf:"varint,12,opt,name=inactive_rank_decrease_percent,json=inactiveRankDecreasePercent,proto3" json:"inactive_rank_decrease_percent,omitempty"`
	MinValidators               uint64 `protobuf:"varint,13,opt,name=min_validators,json=minValidators,proto3" json:"min_validators,omitempty"`
	PoorNetworkMaxBankSend      uint64 `protobuf:"varint,14,opt,name=poor_network_max_bank_send,json=poorNetworkMaxBankSend,proto3" json:"poor_network_max_bank_send,omitempty"`
	JailMaxTime                 uint64 `protobuf:"varint,15,opt,name=jail_max_time,json=jailMaxTime,proto3" json:"jail_max_time,omitempty"`
	EnableTokenWhitelist        bool   `protobuf:"varint,16,opt,name=enable_token_whitelist,json=enableTokenWhitelist,proto3" json:"enable_token_whitelist,omitempty"`
	EnableTokenBlacklist        bool   `protobuf:"varint,17,opt,name=enable_token_blacklist,json=enableTokenBlacklist,proto3" json:"enable_token_blacklist,omitempty"`
	MinIdentityApprovalTip      uint64 `protobuf:"varint,18,opt,name=min_identity_approval_tip,json=minIdentityApprovalTip,proto3" json:"min_identity_approval_tip,omitempty"`
}

func (m *NetworkProperties) Reset()         { *m = NetworkProperties{} }
func (m *NetworkProperties) String() string { return proto.CompactTextString(m) }
func (*NetworkProperties) ProtoMessage()    {}
func (*NetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{1}
}
func (m *NetworkProperties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkProperties.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProperties.Merge(m, src)
}
func (m *NetworkProperties) XXX_Size() int {
	return m.Size()
}
func (m *NetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProperties proto.InternalMessageInfo

func (m *NetworkProperties) GetMinTxFee() uint64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

func (m *NetworkProperties) GetMaxTxFee() uint64 {
	if m != nil {
		return m.MaxTxFee
	}
	return 0
}

func (m *NetworkProperties) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *NetworkProperties) GetProposalEndTime() uint64 {
	if m != nil {
		return m.ProposalEndTime
	}
	return 0
}

func (m *NetworkProperties) GetProposalEnactmentTime() uint64 {
	if m != nil {
		return m.ProposalEnactmentTime
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEndBlocks() uint64 {
	if m != nil {
		return m.MinProposalEndBlocks
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEnactmentBlocks() uint64 {
	if m != nil {
		return m.MinProposalEnactmentBlocks
	}
	return 0
}

func (m *NetworkProperties) GetEnableForeignFeePayments() bool {
	if m != nil {
		return m.EnableForeignFeePayments
	}
	return false
}

func (m *NetworkProperties) GetMischanceRankDecreaseAmount() uint64 {
	if m != nil {
		return m.MischanceRankDecreaseAmount
	}
	return 0
}

func (m *NetworkProperties) GetMaxMischance() uint64 {
	if m != nil {
		return m.MaxMischance
	}
	return 0
}

func (m *NetworkProperties) GetMischanceConfidence() uint64 {
	if m != nil {
		return m.MischanceConfidence
	}
	return 0
}

func (m *NetworkProperties) GetInactiveRankDecreasePercent() uint64 {
	if m != nil {
		return m.InactiveRankDecreasePercent
	}
	return 0
}

func (m *NetworkProperties) GetMinValidators() uint64 {
	if m != nil {
		return m.MinValidators
	}
	return 0
}

func (m *NetworkProperties) GetPoorNetworkMaxBankSend() uint64 {
	if m != nil {
		return m.PoorNetworkMaxBankSend
	}
	return 0
}

func (m *NetworkProperties) GetJailMaxTime() uint64 {
	if m != nil {
		return m.JailMaxTime
	}
	return 0
}

func (m *NetworkProperties) GetEnableTokenWhitelist() bool {
	if m != nil {
		return m.EnableTokenWhitelist
	}
	return false
}

func (m *NetworkProperties) GetEnableTokenBlacklist() bool {
	if m != nil {
		return m.EnableTokenBlacklist
	}
	return false
}

func (m *NetworkProperties) GetMinIdentityApprovalTip() uint64 {
	if m != nil {
		return m.MinIdentityApprovalTip
	}
	return 0
}

func init() {
	proto.RegisterEnum("tsuki.gov.NetworkProperty", NetworkProperty_name, NetworkProperty_value)
	proto.RegisterType((*MsgSetNetworkProperties)(nil), "tsuki.gov.MsgSetNetworkProperties")
	proto.RegisterType((*NetworkProperties)(nil), "tsuki.gov.NetworkProperties")
}

func init() { proto.RegisterFile("tsuki/gov/network_properties.proto", fileDescriptor_98011a6048e5dde3) }

var fileDescriptor_98011a6048e5dde3 = []byte{
	// 1109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xdd, 0x8e, 0xd3, 0x46,
	0x14, 0xde, 0xc0, 0xc2, 0x86, 0xd9, 0xcd, 0xc6, 0x31, 0x0b, 0x6b, 0xbc, 0x34, 0xb8, 0x5b, 0x21,
	0x21, 0x24, 0x12, 0xd1, 0x3f, 0xa9, 0x48, 0x55, 0xe5, 0x24, 0x4e, 0x31, 0x89, 0x7f, 0x70, 0x4c,
	0x80, 0xde, 0x8c, 0x26, 0xce, 0x10, 0xa6, 0x89, 0x3d, 0xa9, 0xed, 0x0d, 0xd9, 0x37, 0xa8, 0x7c,
	0xd5, 0x17, 0xf0, 0x45, 0xd5, 0x57, 0xe8, 0x43, 0xf4, 0x92, 0xcb, 0x5e, 0x55, 0x15, 0xbc, 0x45,
	0xaf, 0xaa, 0x19, 0x3b, 0xd9, 0x64, 0xc3, 0xee, 0xd5, 0xae, 0xe6, 0xfb, 0x39, 0xe7, 0xcc, 0x99,
	0x4f, 0x31, 0xf8, 0x7c, 0x4c, 0x42, 0x54, 0x1f, 0xd1, 0x59, 0x3d, 0xc0, 0xf1, 0x3b, 0x1a, 0x8e,
	0xe1, 0x34, 0xa4, 0x53, 0x1c, 0xc6, 0x04, 0x47, 0xb5, 0x69, 0x48, 0x63, 0x2a, 0x16, 0x19, 0xa5,
	0x36, 0xa2, 0x33, 0xf9, 0x60, 0x44, 0x47, 0x94, 0x1f, 0xd6, 0xd9, 0x7f, 0x19, 0x7e, 0xfc, 0x67,
	0x01, 0x1c, 0x1a, 0xd1, 0xa8, 0x87, 0x63, 0x33, 0xb3, 0xb0, 0x97, 0x0e, 0xe2, 0x33, 0x20, 0x6e,
	0xfa, 0x4a, 0x05, 0xa5, 0xf0, 0x60, 0xf7, 0xcb, 0xa3, 0xda, 0xc2, 0xb8, 0xb6, 0x21, 0x74, 0x2a,
	0xc1, 0x86, 0x97, 0x01, 0x8a, 0xcc, 0x83, 0x46, 0x38, 0x94, 0xae, 0x28, 0x85, 0x07, 0x7b, 0x8d,
	0xc7, 0xff, 0xfd, 0x73, 0xef, 0xd1, 0x88, 0xc4, 0x6f, 0x4f, 0x06, 0x35, 0x8f, 0xfa, 0x75, 0x8f,
	0x46, 0x3e, 0x8d, 0xf2, 0x3f, 0x8f, 0xa2, 0xe1, 0xb8, 0x1e, 0x9f, 0x4e, 0x71, 0x54, 0x53, 0x3d,
	0x4f, 0x1d, 0x0e, 0x43, 0x1c, 0x45, 0xce, 0xd2, 0xe2, 0x38, 0xdd, 0x01, 0x95, 0xcd, 0x86, 0xef,
	0x02, 0xe0, 0x93, 0x00, 0xc6, 0x73, 0xf8, 0x06, 0x63, 0xde, 0xe8, 0xb6, 0x53, 0xf4, 0x49, 0xe0,
	0xce, 0xdb, 0x18, 0x73, 0x14, 0xcd, 0x17, 0xe8, 0x95, 0x1c, 0x45, 0xf3, 0x0c, 0xbd, 0x07, 0x76,
	0x67, 0x34, 0xc6, 0xf0, 0x97, 0x13, 0x1a, 0x9e, 0xf8, 0xd2, 0x55, 0x0e, 0x03, 0x76, 0xf4, 0x9c,
	0x9f, 0x88, 0x0f, 0x41, 0x25, 0x2b, 0x8f, 0x26, 0x10, 0x07, 0x43, 0x18, 0x13, 0x1f, 0x4b, 0xdb,
	0x9c, 0x56, 0x5e, 0x00, 0x5a, 0x30, 0x74, 0x89, 0x8f, 0xc5, 0x6f, 0xc1, 0xe1, 0x0a, 0x17, 0x79,
	0xb1, 0x8f, 0x83, 0x38, 0x53, 0x5c, 0xe3, 0x8a, 0x5b, 0x67, 0x8a, 0x1c, 0xe5, 0xba, 0x6f, 0xc0,
	0x21, 0x1b, 0x60, 0xad, 0xce, 0x60, 0x42, 0xbd, 0x71, 0x24, 0x5d, 0xe7, 0xba, 0x03, 0x9f, 0x04,
	0xf6, 0x59, 0xb1, 0x06, 0xc7, 0x44, 0x15, 0x7c, 0x76, 0x4e, 0xb6, 0x28, 0x99, 0x8b, 0x77, 0xb8,
	0x58, 0x5e, 0x13, 0xe7, 0x94, 0xdc, 0xe2, 0x7b, 0x70, 0x84, 0x03, 0x34, 0x98, 0x60, 0xf8, 0x86,
	0x86, 0x98, 0x8c, 0x02, 0x76, 0x49, 0x70, 0x8a, 0x4e, 0x19, 0x27, 0x92, 0x8a, 0x4a, 0xe1, 0x41,
	0xd1, 0x91, 0x32, 0x4a, 0x3b, 0x63, 0xb4, 0x31, 0xb6, 0x73, 0x5c, 0x6c, 0x82, 0xaa, 0x4f, 0x22,
	0xef, 0x2d, 0x0a, 0x3c, 0x0c, 0x43, 0x14, 0x8c, 0xe1, 0x10, 0x7b, 0x21, 0x46, 0x11, 0x86, 0xc8,
	0xa7, 0x27, 0x41, 0x2c, 0xdd, 0xe0, 0x2d, 0x1c, 0x2d, 0x59, 0x0e, 0x0a, 0xc6, 0xad, 0x9c, 0xa3,
	0x72, 0x8a, 0xf8, 0x05, 0x28, 0xb1, 0x05, 0x2d, 0x29, 0x12, 0xe0, 0x9a, 0x3d, 0x1f, 0xcd, 0x8d,
	0xc5, 0x99, 0xf8, 0x18, 0x1c, 0x9c, 0x55, 0xf2, 0x68, 0xf0, 0x86, 0x0c, 0x31, 0xe3, 0xee, 0x72,
	0xee, 0xcd, 0x25, 0xd6, 0x5c, 0x42, 0xac, 0x39, 0xc2, 0xc6, 0x25, 0xb3, 0xf3, 0xbd, 0x4d, 0x71,
	0xe8, 0xe1, 0x20, 0x96, 0xf6, 0xb2, 0xe6, 0x16, 0xac, 0xd5, 0xde, 0xec, 0x8c, 0x22, 0xde, 0x07,
	0xfb, 0xec, 0x8e, 0x67, 0x68, 0x42, 0x86, 0x28, 0xa6, 0x61, 0x24, 0x95, 0xb8, 0xa8, 0xe4, 0x93,
	0xa0, 0xbf, 0x3c, 0x14, 0x9f, 0x00, 0x79, 0x4a, 0x69, 0x08, 0x17, 0xc1, 0x61, 0x03, 0x0d, 0x58,
	0xcd, 0x08, 0x07, 0x43, 0x69, 0x9f, 0x4b, 0x6e, 0x33, 0x46, 0xfe, 0x7a, 0x0d, 0x34, 0x6f, 0xa0,
	0x60, 0xdc, 0xc3, 0xc1, 0x50, 0x3c, 0x06, 0xa5, 0x9f, 0x11, 0x99, 0x70, 0x0d, 0x7f, 0x2b, 0x65,
	0x4e, 0xdf, 0x65, 0x87, 0x06, 0x9a, 0xf3, 0x17, 0xf2, 0x35, 0xb8, 0x9d, 0xef, 0x29, 0xa6, 0x63,
	0x1c, 0xc0, 0x77, 0x6f, 0x49, 0x8c, 0x27, 0x24, 0x8a, 0x25, 0x81, 0xaf, 0xe8, 0x20, 0x43, 0x5d,
	0x06, 0xbe, 0x5c, 0x60, 0x1b, 0xaa, 0xc1, 0x04, 0x79, 0x63, 0xae, 0xaa, 0x6c, 0xa8, 0x1a, 0x0b,
	0x4c, 0xfc, 0x0e, 0xdc, 0x61, 0x23, 0xb3, 0x5b, 0x8c, 0x49, 0x7c, 0x0a, 0xd1, 0x74, 0x1a, 0xd2,
	0x19, 0x9a, 0xc0, 0x98, 0x4c, 0x25, 0x31, 0x1b, 0xc5, 0x27, 0x81, 0x9e, 0xe3, 0x6a, 0x0e, 0xbb,
	0x64, 0xfa, 0xf0, 0xf7, 0x1d, 0x50, 0x5e, 0xcf, 0xe7, 0x29, 0xcb, 0x9f, 0xa1, 0x9b, 0xd0, 0x7d,
	0x05, 0xdb, 0x9a, 0x26, 0x6c, 0xc9, 0x7b, 0x49, 0xaa, 0x14, 0x8d, 0x95, 0x74, 0x1a, 0xea, 0xab,
	0x05, 0x5a, 0xc8, 0xd1, 0x95, 0x74, 0xf6, 0x2d, 0x57, 0x83, 0xcf, 0x5f, 0x58, 0xce, 0x0b, 0x43,
	0xb8, 0x22, 0xef, 0x27, 0xa9, 0x02, 0xfa, 0x6b, 0xe9, 0xb4, 0x1d, 0xcb, 0xb6, 0x7a, 0x6a, 0x17,
	0x6a, 0x66, 0x0b, 0xba, 0xba, 0xa1, 0x09, 0x57, 0xe5, 0x9b, 0x49, 0xaa, 0x94, 0xed, 0xcd, 0x74,
	0xae, 0x70, 0xd5, 0xa6, 0x6b, 0x68, 0xa6, 0x9b, 0x29, 0xb6, 0xe5, 0x3b, 0x49, 0xaa, 0xdc, 0xb2,
	0x2f, 0x4a, 0x27, 0x1b, 0x60, 0xad, 0x4e, 0xa3, 0x6b, 0x35, 0x3b, 0x3d, 0xe1, 0x9a, 0x2c, 0x25,
	0xa9, 0x72, 0x60, 0x5c, 0x90, 0xce, 0x73, 0xb2, 0x45, 0xc9, 0x5c, 0x7c, 0x5d, 0xae, 0x26, 0xa9,
	0x22, 0x1b, 0x97, 0xa6, 0x53, 0x33, 0xd5, 0x46, 0x57, 0x83, 0x6d, 0xcb, 0xd1, 0xf4, 0x1f, 0x4d,
	0x76, 0x49, 0xd0, 0x56, 0x5f, 0x33, 0x9b, 0x9e, 0xb0, 0x23, 0xdf, 0x4d, 0x52, 0x45, 0xd2, 0x2e,
	0x49, 0xa7, 0xa1, 0xf7, 0x9a, 0x4f, 0x55, 0xb3, 0xa9, 0x41, 0x47, 0x35, 0x3b, 0xb0, 0xa5, 0x35,
	0x1d, 0x4d, 0xed, 0x69, 0x50, 0x35, 0xac, 0x17, 0xa6, 0x2b, 0x14, 0xe5, 0x7b, 0x49, 0xaa, 0x1c,
	0x19, 0x97, 0xa7, 0x93, 0x2d, 0x68, 0x69, 0x24, 0xdc, 0x90, 0x85, 0x24, 0x55, 0xf6, 0x8c, 0x73,
	0xe9, 0x3c, 0xab, 0xd4, 0xb4, 0xcc, 0xb6, 0xde, 0xd2, 0x18, 0x17, 0xc8, 0x87, 0x49, 0xaa, 0xdc,
	0x34, 0x3e, 0x9d, 0x4e, 0x9d, 0xdd, 0x88, 0xde, 0x3f, 0xdf, 0x9b, 0xad, 0x39, 0x4d, 0xcd, 0x74,
	0x85, 0xdd, 0xac, 0x39, 0xfd, 0x92, 0x74, 0x3e, 0x01, 0xb2, 0x6d, 0x59, 0x0e, 0x34, 0x35, 0xf7,
	0xa5, 0xe5, 0x74, 0x20, 0xeb, 0xb4, 0xc1, 0xcc, 0x7a, 0x9a, 0xd9, 0x12, 0xf6, 0x64, 0x39, 0x49,
	0x95, 0xdb, 0xf6, 0xa7, 0x63, 0x77, 0x1f, 0xec, 0xb3, 0xfd, 0xf4, 0xd5, 0xae, 0xde, 0x52, 0x5d,
	0xcb, 0xe9, 0x09, 0x25, 0xb9, 0x92, 0xa4, 0x4a, 0xc9, 0x58, 0x4b, 0xf6, 0x31, 0x28, 0x3d, 0x53,
	0xf5, 0x2e, 0xb7, 0xe6, 0x6f, 0x65, 0x5f, 0x2e, 0x27, 0xa9, 0xb2, 0xfb, 0x6c, 0x3d, 0x9d, 0xf9,
	0x9e, 0x5c, 0xab, 0xa3, 0x99, 0xf0, 0xe5, 0x53, 0xdd, 0xd5, 0xba, 0x7a, 0xcf, 0x15, 0xca, 0xd9,
	0x03, 0xd1, 0x2e, 0x48, 0xe7, 0x9a, 0xaa, 0xd1, 0x55, 0x9b, 0x1d, 0xae, 0x12, 0x36, 0x54, 0x6b,
	0xe9, 0x64, 0x6d, 0xb3, 0x0b, 0x76, 0x75, 0xf7, 0x35, 0x54, 0x6d, 0xdb, 0xb1, 0xfa, 0x6a, 0x17,
	0xba, 0xba, 0x2d, 0x54, 0xb2, 0x89, 0x8d, 0x4f, 0xa6, 0x53, 0xde, 0xfe, 0xf5, 0x8f, 0xea, 0x56,
	0xe3, 0x87, 0xbf, 0x3e, 0x54, 0x0b, 0xef, 0x3f, 0x54, 0x0b, 0xff, 0x7e, 0xa8, 0x16, 0x7e, 0xfb,
	0x58, 0xdd, 0x7a, 0xff, 0xb1, 0xba, 0xf5, 0xf7, 0xc7, 0xea, 0xd6, 0x4f, 0xf7, 0x57, 0x7e, 0x96,
	0x3b, 0x24, 0x44, 0x4d, 0x1a, 0xe2, 0x7a, 0x84, 0xc7, 0x88, 0xd4, 0xe7, 0xfc, 0x73, 0x83, 0xff,
	0x32, 0x0f, 0xae, 0xf3, 0x4f, 0x88, 0xaf, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x9b, 0x66,
	0x09, 0x87, 0x08, 0x00, 0x00,
}

func (m *MsgSetNetworkProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetNetworkProperties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetNetworkProperties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintNetworkProperties(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if m.NetworkProperties != nil {
		{
			size, err := m.NetworkProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNetworkProperties(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NetworkProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkProperties) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkProperties) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinIdentityApprovalTip != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinIdentityApprovalTip))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.EnableTokenBlacklist {
		i--
		if m.EnableTokenBlacklist {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.EnableTokenWhitelist {
		i--
		if m.EnableTokenWhitelist {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.JailMaxTime != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.JailMaxTime))
		i--
		dAtA[i] = 0x78
	}
	if m.PoorNetworkMaxBankSend != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.PoorNetworkMaxBankSend))
		i--
		dAtA[i] = 0x70
	}
	if m.MinValidators != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinValidators))
		i--
		dAtA[i] = 0x68
	}
	if m.InactiveRankDecreasePercent != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.InactiveRankDecreasePercent))
		i--
		dAtA[i] = 0x60
	}
	if m.MischanceConfidence != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MischanceConfidence))
		i--
		dAtA[i] = 0x58
	}
	if m.MaxMischance != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MaxMischance))
		i--
		dAtA[i] = 0x50
	}
	if m.MischanceRankDecreaseAmount != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MischanceRankDecreaseAmount))
		i--
		dAtA[i] = 0x48
	}
	if m.EnableForeignFeePayments {
		i--
		if m.EnableForeignFeePayments {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.MinProposalEnactmentBlocks != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinProposalEnactmentBlocks))
		i--
		dAtA[i] = 0x38
	}
	if m.MinProposalEndBlocks != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinProposalEndBlocks))
		i--
		dAtA[i] = 0x30
	}
	if m.ProposalEnactmentTime != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.ProposalEnactmentTime))
		i--
		dAtA[i] = 0x28
	}
	if m.ProposalEndTime != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.ProposalEndTime))
		i--
		dAtA[i] = 0x20
	}
	if m.VoteQuorum != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.VoteQuorum))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxTxFee != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MaxTxFee))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTxFee != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinTxFee))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetworkProperties(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetworkProperties(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetNetworkProperties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NetworkProperties != nil {
		l = m.NetworkProperties.Size()
		n += 1 + l + sovNetworkProperties(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovNetworkProperties(uint64(l))
	}
	return n
}

func (m *NetworkProperties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTxFee != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MinTxFee))
	}
	if m.MaxTxFee != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MaxTxFee))
	}
	if m.VoteQuorum != 0 {
		n += 1 + sovNetworkProperties(uint64(m.VoteQuorum))
	}
	if m.ProposalEndTime != 0 {
		n += 1 + sovNetworkProperties(uint64(m.ProposalEndTime))
	}
	if m.ProposalEnactmentTime != 0 {
		n += 1 + sovNetworkProperties(uint64(m.ProposalEnactmentTime))
	}
	if m.MinProposalEndBlocks != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MinProposalEndBlocks))
	}
	if m.MinProposalEnactmentBlocks != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MinProposalEnactmentBlocks))
	}
	if m.EnableForeignFeePayments {
		n += 2
	}
	if m.MischanceRankDecreaseAmount != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MischanceRankDecreaseAmount))
	}
	if m.MaxMischance != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MaxMischance))
	}
	if m.MischanceConfidence != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MischanceConfidence))
	}
	if m.InactiveRankDecreasePercent != 0 {
		n += 1 + sovNetworkProperties(uint64(m.InactiveRankDecreasePercent))
	}
	if m.MinValidators != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MinValidators))
	}
	if m.PoorNetworkMaxBankSend != 0 {
		n += 1 + sovNetworkProperties(uint64(m.PoorNetworkMaxBankSend))
	}
	if m.JailMaxTime != 0 {
		n += 1 + sovNetworkProperties(uint64(m.JailMaxTime))
	}
	if m.EnableTokenWhitelist {
		n += 3
	}
	if m.EnableTokenBlacklist {
		n += 3
	}
	if m.MinIdentityApprovalTip != 0 {
		n += 2 + sovNetworkProperties(uint64(m.MinIdentityApprovalTip))
	}
	return n
}

func sovNetworkProperties(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetworkProperties(x uint64) (n int) {
	return sovNetworkProperties(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetNetworkProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkProperties
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetNetworkProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetNetworkProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkProperties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NetworkProperties == nil {
				m.NetworkProperties = &NetworkProperties{}
			}
			if err := m.NetworkProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NetworkProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkProperties
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTxFee", wireType)
			}
			m.MinTxFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTxFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTxFee", wireType)
			}
			m.MaxTxFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTxFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteQuorum", wireType)
			}
			m.VoteQuorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteQuorum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalEndTime", wireType)
			}
			m.ProposalEndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalEndTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalEnactmentTime", wireType)
			}
			m.ProposalEnactmentTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalEnactmentTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinProposalEndBlocks", wireType)
			}
			m.MinProposalEndBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinProposalEndBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinProposalEnactmentBlocks", wireType)
			}
			m.MinProposalEnactmentBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinProposalEnactmentBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableForeignFeePayments", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableForeignFeePayments = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MischanceRankDecreaseAmount", wireType)
			}
			m.MischanceRankDecreaseAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MischanceRankDecreaseAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMischance", wireType)
			}
			m.MaxMischance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMischance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MischanceConfidence", wireType)
			}
			m.MischanceConfidence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MischanceConfidence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveRankDecreasePercent", wireType)
			}
			m.InactiveRankDecreasePercent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InactiveRankDecreasePercent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinValidators", wireType)
			}
			m.MinValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinValidators |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoorNetworkMaxBankSend", wireType)
			}
			m.PoorNetworkMaxBankSend = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoorNetworkMaxBankSend |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailMaxTime", wireType)
			}
			m.JailMaxTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailMaxTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableTokenWhitelist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableTokenWhitelist = bool(v != 0)
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableTokenBlacklist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableTokenBlacklist = bool(v != 0)
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinIdentityApprovalTip", wireType)
			}
			m.MinIdentityApprovalTip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinIdentityApprovalTip |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkProperties(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetworkProperties
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetworkProperties(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkProperties
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetworkProperties
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNetworkProperties
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetworkProperties
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetworkProperties
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetworkProperties        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkProperties          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetworkProperties = fmt.Errorf("proto: unexpected end of group")
)
