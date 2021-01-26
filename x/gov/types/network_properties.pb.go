// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network_properties.proto

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
	EnableForeignFeePayments    NetworkProperty = 5
	MischanceRankDecreaseAmount NetworkProperty = 6
	InactiveRankDecreasePercent NetworkProperty = 7
	PoorNetworkMaxBankSend      NetworkProperty = 8
	MinValidators               NetworkProperty = 9
)

var NetworkProperty_name = map[int32]string{
	0: "MIN_TX_FEE",
	1: "MAX_TX_FEE",
	2: "VOTE_QUORUM",
	3: "PROPOSAL_END_TIME",
	4: "PROPOSAL_ENACTMENT_TIME",
	5: "ENABLE_FOREIGN_TX_FEE_PAYMENTS",
	6: "MISCHANCE_RANK_DECREASE_AMOUNT",
	7: "INACTIVE_RANK_DECREASE_PERCENT",
	8: "POOR_NETWORK_MAX_BANK_SEND",
	9: "MIN_VALIDATORS",
}

var NetworkProperty_value = map[string]int32{
	"MIN_TX_FEE":                     0,
	"MAX_TX_FEE":                     1,
	"VOTE_QUORUM":                    2,
	"PROPOSAL_END_TIME":              3,
	"PROPOSAL_ENACTMENT_TIME":        4,
	"ENABLE_FOREIGN_TX_FEE_PAYMENTS": 5,
	"MISCHANCE_RANK_DECREASE_AMOUNT": 6,
	"INACTIVE_RANK_DECREASE_PERCENT": 7,
	"POOR_NETWORK_MAX_BANK_SEND":     8,
	"MIN_VALIDATORS":                 9,
}

func (x NetworkProperty) String() string {
	return proto.EnumName(NetworkProperty_name, int32(x))
}

func (NetworkProperty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_afa35a4ab1e9e2c2, []int{0}
}

type MsgSetNetworkProperties struct {
	NetworkProperties *NetworkProperties                            `protobuf:"bytes,1,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	Proposer          github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
}

func (m *MsgSetNetworkProperties) Reset()         { *m = MsgSetNetworkProperties{} }
func (m *MsgSetNetworkProperties) String() string { return proto.CompactTextString(m) }
func (*MsgSetNetworkProperties) ProtoMessage()    {}
func (*MsgSetNetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa35a4ab1e9e2c2, []int{0}
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
	MinTxFee                 uint64 `protobuf:"varint,1,opt,name=min_tx_fee,json=minTxFee,proto3" json:"min_tx_fee,omitempty"`
	MaxTxFee                 uint64 `protobuf:"varint,2,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee,omitempty"`
	VoteQuorum               uint64 `protobuf:"varint,3,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	ProposalEndTime          uint64 `protobuf:"varint,4,opt,name=proposal_end_time,json=proposalEndTime,proto3" json:"proposal_end_time,omitempty"`
	ProposalEnactmentTime    uint64 `protobuf:"varint,5,opt,name=proposal_enactment_time,json=proposalEnactmentTime,proto3" json:"proposal_enactment_time,omitempty"`
	EnableForeignFeePayments bool   `protobuf:"varint,6,opt,name=enable_foreign_fee_payments,json=enableForeignFeePayments,proto3" json:"enable_foreign_fee_payments,omitempty"`
	// The rank property is a long term statistics implying the "longest" streak that validator ever achieved,
	// it can be expressed as rank = MAX(rank, streak).
	// Under certain circumstances we should however decrease the rank of the validator.
	// If the mischance property is incremented, the rank should be decremented by X (default 10), that is rank = MAX(rank - X, 0).
	// Every time node status changes to inactive the rank should be divided by 2, that is rank = FLOOR(rank / 2)
	// The streak and rank will enable governance to judge real life performance of validators on the mainnet or testnet, and potentially propose eviction of the weakest and least reliable operators.
	MischanceRankDecreaseAmount uint64 `protobuf:"varint,7,opt,name=mischance_rank_decrease_amount,json=mischanceRankDecreaseAmount,proto3" json:"mischance_rank_decrease_amount,omitempty"`
	InactiveRankDecreasePercent uint64 `protobuf:"varint,8,opt,name=inactive_rank_decrease_percent,json=inactiveRankDecreasePercent,proto3" json:"inactive_rank_decrease_percent,omitempty"`
	MinValidators               uint64 `protobuf:"varint,9,opt,name=min_validators,json=minValidators,proto3" json:"min_validators,omitempty"`
	PoorNetworkMaxBankSend      uint64 `protobuf:"varint,10,opt,name=poor_network_max_bank_send,json=poorNetworkMaxBankSend,proto3" json:"poor_network_max_bank_send,omitempty"`
}

func (m *NetworkProperties) Reset()         { *m = NetworkProperties{} }
func (m *NetworkProperties) String() string { return proto.CompactTextString(m) }
func (*NetworkProperties) ProtoMessage()    {}
func (*NetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa35a4ab1e9e2c2, []int{1}
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

func init() {
	proto.RegisterEnum("tsuki.gov.NetworkProperty", NetworkProperty_name, NetworkProperty_value)
	proto.RegisterType((*MsgSetNetworkProperties)(nil), "tsuki.gov.MsgSetNetworkProperties")
	proto.RegisterType((*NetworkProperties)(nil), "tsuki.gov.NetworkProperties")
}

func init() { proto.RegisterFile("network_properties.proto", fileDescriptor_afa35a4ab1e9e2c2) }

var fileDescriptor_afa35a4ab1e9e2c2 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x4d, 0x8f, 0xdb, 0x44,
	0x1c, 0xc6, 0xe3, 0x6d, 0xba, 0x4d, 0xa7, 0x2f, 0xeb, 0x35, 0xb4, 0x35, 0xde, 0xca, 0x6b, 0x55,
	0xaa, 0xb4, 0xaa, 0xd4, 0x44, 0x80, 0xc4, 0xa1, 0x12, 0x02, 0x27, 0x99, 0x80, 0xd9, 0xf5, 0x4b,
	0x6d, 0x6f, 0x28, 0x5c, 0x46, 0x13, 0x7b, 0x9a, 0x5a, 0x59, 0xcf, 0x84, 0xb1, 0x13, 0xb2, 0x37,
	0x6e, 0x20, 0x9f, 0xf8, 0x02, 0x3e, 0xf1, 0x15, 0xf8, 0x10, 0x1c, 0x7b, 0xe4, 0x84, 0xd0, 0xee,
	0xb7, 0xe0, 0x84, 0xc6, 0x4e, 0xd2, 0xb4, 0x59, 0xed, 0x29, 0xd1, 0x3c, 0xcf, 0xf3, 0x9b, 0x7f,
	0xf2, 0x7f, 0x34, 0x40, 0xa5, 0x24, 0xff, 0x99, 0xf1, 0x09, 0x9a, 0x72, 0x36, 0x25, 0x3c, 0x4f,
	0x48, 0xd6, 0x9e, 0x72, 0x96, 0x33, 0xa5, 0x35, 0x49, 0x38, 0x6e, 0x8f, 0xd9, 0x5c, 0xfb, 0x78,
	0xcc, 0xc6, 0xac, 0x3a, 0xec, 0x88, 0x6f, 0xb5, 0xfe, 0xe4, 0x4f, 0x09, 0x3c, 0xb2, 0xb3, 0x71,
	0x40, 0x72, 0xa7, 0x46, 0x78, 0x6b, 0x82, 0xf2, 0x1d, 0x50, 0xb6, 0xb9, 0xaa, 0x64, 0x48, 0x47,
	0x77, 0x3e, 0x3b, 0x68, 0xaf, 0xc0, 0xed, 0xad, 0xa0, 0xbf, 0x4f, 0xb7, 0x58, 0x36, 0x68, 0x09,
	0x06, 0xcb, 0x08, 0x57, 0x77, 0x0c, 0xe9, 0xe8, 0x6e, 0xf7, 0xd3, 0xff, 0xfe, 0x39, 0x7c, 0x3e,
	0x4e, 0xf2, 0x37, 0xb3, 0x51, 0x3b, 0x62, 0x69, 0x27, 0x62, 0x59, 0xca, 0xb2, 0xe5, 0xc7, 0xf3,
	0x2c, 0x9e, 0x74, 0xf2, 0xf3, 0x29, 0xc9, 0xda, 0x66, 0x14, 0x99, 0x71, 0xcc, 0x49, 0x96, 0xf9,
	0x6b, 0xc4, 0x93, 0x5f, 0x9a, 0x60, 0x7f, 0x7b, 0xe0, 0xc7, 0x00, 0xa4, 0x09, 0x45, 0xf9, 0x02,
	0xbd, 0x26, 0xa4, 0x1a, 0xb4, 0xe9, 0xb7, 0xd2, 0x84, 0x86, 0x8b, 0x01, 0x21, 0x95, 0x8a, 0x17,
	0x2b, 0x75, 0x67, 0xa9, 0xe2, 0x45, 0xad, 0x1e, 0x82, 0x3b, 0x73, 0x96, 0x13, 0xf4, 0xd3, 0x8c,
	0xf1, 0x59, 0xaa, 0xde, 0xa8, 0x64, 0x20, 0x8e, 0x5e, 0x56, 0x27, 0xca, 0x33, 0xb0, 0x5f, 0x5f,
	0x8f, 0xcf, 0x10, 0xa1, 0x31, 0xca, 0x93, 0x94, 0xa8, 0xcd, 0xca, 0xb6, 0xb7, 0x12, 0x20, 0x8d,
	0xc3, 0x24, 0x25, 0xca, 0x17, 0xe0, 0xd1, 0x86, 0x17, 0x47, 0x79, 0x4a, 0x68, 0x5e, 0x27, 0x6e,
	0x56, 0x89, 0x07, 0xef, 0x12, 0x4b, 0xb5, 0xca, 0x7d, 0x09, 0x0e, 0x08, 0xc5, 0xa3, 0x33, 0x82,
	0x5e, 0x33, 0x4e, 0x92, 0x31, 0x15, 0xa3, 0xa2, 0x29, 0x3e, 0x17, 0x8e, 0x4c, 0xdd, 0x35, 0xa4,
	0xa3, 0x96, 0xaf, 0xd6, 0x96, 0x41, 0xed, 0x18, 0x10, 0xe2, 0x2d, 0x75, 0xa5, 0x07, 0xf4, 0x34,
	0xc9, 0xa2, 0x37, 0x98, 0x46, 0x04, 0x71, 0x4c, 0x27, 0x28, 0x26, 0x11, 0x27, 0x38, 0x23, 0x08,
	0xa7, 0x6c, 0x46, 0x73, 0xf5, 0x56, 0x75, 0xfb, 0xc1, 0xda, 0xe5, 0x63, 0x3a, 0xe9, 0x2f, 0x3d,
	0x66, 0x65, 0x11, 0x90, 0x44, 0x0c, 0x95, 0xcc, 0x3f, 0x64, 0x4c, 0x09, 0x8f, 0x08, 0xcd, 0xd5,
	0x56, 0x0d, 0x59, 0xb9, 0x36, 0x19, 0x5e, 0x6d, 0x51, 0x9e, 0x82, 0xfb, 0x62, 0x13, 0x73, 0x7c,
	0x96, 0xc4, 0x38, 0x67, 0x3c, 0x53, 0x6f, 0x57, 0xa1, 0x7b, 0x69, 0x42, 0x87, 0xeb, 0x43, 0xe5,
	0x05, 0xd0, 0xa6, 0x8c, 0x71, 0xb4, 0xaa, 0x99, 0xd8, 0xcf, 0x48, 0xdc, 0x99, 0x11, 0x1a, 0xab,
	0xa0, 0x8a, 0x3c, 0x14, 0x8e, 0xe5, 0xae, 0x6d, 0xbc, 0xe8, 0x62, 0x3a, 0x09, 0x08, 0x8d, 0x9f,
	0xfd, 0xda, 0x04, 0x7b, 0xef, 0x57, 0xe0, 0x5c, 0xac, 0xd8, 0xb6, 0x1c, 0x14, 0xbe, 0x42, 0x03,
	0x08, 0xe5, 0x86, 0x76, 0xb7, 0x28, 0x8d, 0x96, 0xbd, 0x51, 0x00, 0xdb, 0x7c, 0xb5, 0x52, 0xa5,
	0xa5, 0xba, 0x51, 0x80, 0xa1, 0x1b, 0x42, 0xf4, 0xf2, 0xd4, 0xf5, 0x4f, 0x6d, 0x79, 0x47, 0xbb,
	0x5f, 0x94, 0x06, 0x18, 0xbe, 0x57, 0x00, 0xcf, 0x77, 0x3d, 0x37, 0x30, 0x4f, 0x10, 0x74, 0xfa,
	0x28, 0xb4, 0x6c, 0x28, 0xdf, 0xd0, 0x3e, 0x2a, 0x4a, 0x63, 0xcf, 0xdb, 0x2e, 0xc0, 0x86, 0xd7,
	0xec, 0x85, 0x36, 0x74, 0xc2, 0x3a, 0xd1, 0xd4, 0x3e, 0x29, 0x4a, 0xe3, 0x81, 0x77, 0x65, 0x01,
	0xbe, 0x06, 0x3a, 0x74, 0xcc, 0xee, 0x09, 0x44, 0x03, 0xd7, 0x87, 0xd6, 0x37, 0xab, 0xdf, 0x82,
	0x3c, 0xf3, 0x07, 0x81, 0x08, 0xe4, 0x9b, 0xda, 0xe3, 0xa2, 0x34, 0x54, 0x78, 0x4d, 0x07, 0x6c,
	0x2b, 0xe8, 0x7d, 0x6b, 0x3a, 0x3d, 0x88, 0x7c, 0xd3, 0x39, 0x46, 0x7d, 0xd8, 0xf3, 0xa1, 0x19,
	0x40, 0x64, 0xda, 0xee, 0xa9, 0x13, 0xca, 0xbb, 0xda, 0x61, 0x51, 0x1a, 0x07, 0xf6, 0xf5, 0x1d,
	0xb0, 0xc4, 0xd4, 0xd6, 0xf0, 0x43, 0x86, 0x07, 0xfd, 0x1e, 0x74, 0x42, 0xf9, 0x56, 0x0d, 0xb1,
	0xae, 0xe9, 0xc0, 0x0b, 0xa0, 0x79, 0xae, 0xeb, 0x23, 0x07, 0x86, 0xdf, 0xbb, 0xfe, 0x31, 0x12,
	0xff, 0x7d, 0x57, 0xc0, 0x02, 0xe8, 0xf4, 0xe5, 0x96, 0xa6, 0x15, 0xa5, 0xf1, 0xd0, 0xbb, 0x72,
	0xb9, 0xa2, 0x3f, 0x62, 0x91, 0x43, 0xf3, 0xc4, 0xea, 0x9b, 0xa1, 0xeb, 0x07, 0xf2, 0x6d, 0x6d,
	0xbf, 0x28, 0x8d, 0x7b, 0xf6, 0x66, 0x7f, 0xb4, 0xe6, 0x6f, 0x7f, 0xe8, 0x8d, 0xee, 0x57, 0x7f,
	0x5d, 0xe8, 0xd2, 0xdb, 0x0b, 0x5d, 0xfa, 0xf7, 0x42, 0x97, 0x7e, 0xbf, 0xd4, 0x1b, 0x6f, 0x2f,
	0xf5, 0xc6, 0xdf, 0x97, 0x7a, 0xe3, 0xc7, 0xa7, 0x1b, 0xef, 0xcb, 0x71, 0xc2, 0x71, 0x8f, 0x71,
	0xd2, 0xc9, 0xc8, 0x04, 0x27, 0x9d, 0x45, 0x67, 0xcc, 0xe6, 0xf5, 0x13, 0x33, 0xda, 0xad, 0xde,
	0xc2, 0xcf, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x31, 0x5b, 0x22, 0xec, 0x47, 0x05, 0x00, 0x00,
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
	if m.PoorNetworkMaxBankSend != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.PoorNetworkMaxBankSend))
		i--
		dAtA[i] = 0x50
	}
	if m.MinValidators != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MinValidators))
		i--
		dAtA[i] = 0x48
	}
	if m.InactiveRankDecreasePercent != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.InactiveRankDecreasePercent))
		i--
		dAtA[i] = 0x40
	}
	if m.MischanceRankDecreaseAmount != 0 {
		i = encodeVarintNetworkProperties(dAtA, i, uint64(m.MischanceRankDecreaseAmount))
		i--
		dAtA[i] = 0x38
	}
	if m.EnableForeignFeePayments {
		i--
		if m.EnableForeignFeePayments {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
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
	if m.EnableForeignFeePayments {
		n += 2
	}
	if m.MischanceRankDecreaseAmount != 0 {
		n += 1 + sovNetworkProperties(uint64(m.MischanceRankDecreaseAmount))
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
		case 7:
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
		case 8:
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
		case 9:
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
		case 10:
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
