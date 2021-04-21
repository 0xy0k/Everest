// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: tsuki/gov/network_properties.proto

package proto

import (
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NetworkProperty int32

const (
	NetworkProperty_MIN_TX_FEE                     NetworkProperty = 0
	NetworkProperty_MAX_TX_FEE                     NetworkProperty = 1
	NetworkProperty_VOTE_QUORUM                    NetworkProperty = 2
	NetworkProperty_PROPOSAL_END_TIME              NetworkProperty = 3
	NetworkProperty_PROPOSAL_ENACTMENT_TIME        NetworkProperty = 4
	NetworkProperty_MIN_PROPOSAL_END_BLOCKS        NetworkProperty = 5
	NetworkProperty_MIN_PROPOSAL_ENACTMENT_BLOCKS  NetworkProperty = 6
	NetworkProperty_ENABLE_FOREIGN_FEE_PAYMENTS    NetworkProperty = 7
	NetworkProperty_MISCHANCE_RANK_DECREASE_AMOUNT NetworkProperty = 8
	NetworkProperty_INACTIVE_RANK_DECREASE_PERCENT NetworkProperty = 9
	NetworkProperty_POOR_NETWORK_MAX_BANK_SEND     NetworkProperty = 10
	NetworkProperty_MIN_VALIDATORS                 NetworkProperty = 11
	NetworkProperty_JAIL_MAX_TIME                  NetworkProperty = 12
	NetworkProperty_ENABLE_TOKEN_WHITELIST         NetworkProperty = 13
	NetworkProperty_ENABLE_TOKEN_BLACKLIST         NetworkProperty = 14
)

// Enum value maps for NetworkProperty.
var (
	NetworkProperty_name = map[int32]string{
		0:  "MIN_TX_FEE",
		1:  "MAX_TX_FEE",
		2:  "VOTE_QUORUM",
		3:  "PROPOSAL_END_TIME",
		4:  "PROPOSAL_ENACTMENT_TIME",
		5:  "MIN_PROPOSAL_END_BLOCKS",
		6:  "MIN_PROPOSAL_ENACTMENT_BLOCKS",
		7:  "ENABLE_FOREIGN_FEE_PAYMENTS",
		8:  "MISCHANCE_RANK_DECREASE_AMOUNT",
		9:  "INACTIVE_RANK_DECREASE_PERCENT",
		10: "POOR_NETWORK_MAX_BANK_SEND",
		11: "MIN_VALIDATORS",
		12: "JAIL_MAX_TIME",
		13: "ENABLE_TOKEN_WHITELIST",
		14: "ENABLE_TOKEN_BLACKLIST",
	}
	NetworkProperty_value = map[string]int32{
		"MIN_TX_FEE":                     0,
		"MAX_TX_FEE":                     1,
		"VOTE_QUORUM":                    2,
		"PROPOSAL_END_TIME":              3,
		"PROPOSAL_ENACTMENT_TIME":        4,
		"MIN_PROPOSAL_END_BLOCKS":        5,
		"MIN_PROPOSAL_ENACTMENT_BLOCKS":  6,
		"ENABLE_FOREIGN_FEE_PAYMENTS":    7,
		"MISCHANCE_RANK_DECREASE_AMOUNT": 8,
		"INACTIVE_RANK_DECREASE_PERCENT": 9,
		"POOR_NETWORK_MAX_BANK_SEND":     10,
		"MIN_VALIDATORS":                 11,
		"JAIL_MAX_TIME":                  12,
		"ENABLE_TOKEN_WHITELIST":         13,
		"ENABLE_TOKEN_BLACKLIST":         14,
	}
)

func (x NetworkProperty) Enum() *NetworkProperty {
	p := new(NetworkProperty)
	*p = x
	return p
}

func (x NetworkProperty) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkProperty) Descriptor() protoreflect.EnumDescriptor {
	return file_tsuki_gov_network_properties_proto_enumTypes[0].Descriptor()
}

func (NetworkProperty) Type() protoreflect.EnumType {
	return &file_tsuki_gov_network_properties_proto_enumTypes[0]
}

func (x NetworkProperty) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkProperty.Descriptor instead.
func (NetworkProperty) EnumDescriptor() ([]byte, []int) {
	return file_tsuki_gov_network_properties_proto_rawDescGZIP(), []int{0}
}

type MsgSetNetworkProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkProperties *NetworkProperties `protobuf:"bytes,1,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	Proposer          []byte             `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
}

func (x *MsgSetNetworkProperties) Reset() {
	*x = MsgSetNetworkProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tsuki_gov_network_properties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSetNetworkProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSetNetworkProperties) ProtoMessage() {}

func (x *MsgSetNetworkProperties) ProtoReflect() protoreflect.Message {
	mi := &file_tsuki_gov_network_properties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSetNetworkProperties.ProtoReflect.Descriptor instead.
func (*MsgSetNetworkProperties) Descriptor() ([]byte, []int) {
	return file_tsuki_gov_network_properties_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSetNetworkProperties) GetNetworkProperties() *NetworkProperties {
	if x != nil {
		return x.NetworkProperties
	}
	return nil
}

func (x *MsgSetNetworkProperties) GetProposer() []byte {
	if x != nil {
		return x.Proposer
	}
	return nil
}

type NetworkProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinTxFee                   uint64 `protobuf:"varint,1,opt,name=min_tx_fee,json=minTxFee,proto3" json:"min_tx_fee,omitempty"`
	MaxTxFee                   uint64 `protobuf:"varint,2,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee,omitempty"`
	VoteQuorum                 uint64 `protobuf:"varint,3,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	ProposalEndTime            uint64 `protobuf:"varint,4,opt,name=proposal_end_time,json=proposalEndTime,proto3" json:"proposal_end_time,omitempty"`
	ProposalEnactmentTime      uint64 `protobuf:"varint,5,opt,name=proposal_enactment_time,json=proposalEnactmentTime,proto3" json:"proposal_enactment_time,omitempty"`
	MinProposalEndBlocks       uint64 `protobuf:"varint,6,opt,name=min_proposal_end_blocks,json=minProposalEndBlocks,proto3" json:"min_proposal_end_blocks,omitempty"`                   // minimum blocks required for proposal voting
	MinProposalEnactmentBlocks uint64 `protobuf:"varint,7,opt,name=min_proposal_enactment_blocks,json=minProposalEnactmentBlocks,proto3" json:"min_proposal_enactment_blocks,omitempty"` // min blocks required for proposal enactment
	EnableForeignFeePayments   bool   `protobuf:"varint,8,opt,name=enable_foreign_fee_payments,json=enableForeignFeePayments,proto3" json:"enable_foreign_fee_payments,omitempty"`
	// The rank property is a long term statistics implying the "longest" streak that validator ever achieved,
	// it can be expressed as rank = MAX(rank, streak).
	// Under certain circumstances we should however decrease the rank of the validator.
	// If the mischance property is incremented, the rank should be decremented by X (default 10), that is rank = MAX(rank - X, 0).
	// Every time node status changes to inactive the rank should be divided by 2, that is rank = FLOOR(rank / 2)
	// The streak and rank will enable governance to judge real life performance of validators on the mainnet or testnet, and potentially propose eviction of the weakest and least reliable operators.
	MischanceRankDecreaseAmount uint64 `protobuf:"varint,9,opt,name=mischance_rank_decrease_amount,json=mischanceRankDecreaseAmount,proto3" json:"mischance_rank_decrease_amount,omitempty"`  // X (default 10)
	InactiveRankDecreasePercent uint64 `protobuf:"varint,10,opt,name=inactive_rank_decrease_percent,json=inactiveRankDecreasePercent,proto3" json:"inactive_rank_decrease_percent,omitempty"` // Y (default 50%)
	MinValidators               uint64 `protobuf:"varint,11,opt,name=min_validators,json=minValidators,proto3" json:"min_validators,omitempty"`
	PoorNetworkMaxBankSend      uint64 `protobuf:"varint,12,opt,name=poor_network_max_bank_send,json=poorNetworkMaxBankSend,proto3" json:"poor_network_max_bank_send,omitempty"` // default 10000 ukex
	JailMaxTime                 uint64 `protobuf:"varint,13,opt,name=jail_max_time,json=jailMaxTime,proto3" json:"jail_max_time,omitempty"`                                      // Jailing validator maximum time in minutes
	EnableTokenWhitelist        bool   `protobuf:"varint,14,opt,name=enable_token_whitelist,json=enableTokenWhitelist,proto3" json:"enable_token_whitelist,omitempty"`           // TokenWhitelist is valid when this param is set
	EnableTokenBlacklist        bool   `protobuf:"varint,15,opt,name=enable_token_blacklist,json=enableTokenBlacklist,proto3" json:"enable_token_blacklist,omitempty"`           // TokenBlacklist is valid when this param is set
}

func (x *NetworkProperties) Reset() {
	*x = NetworkProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tsuki_gov_network_properties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkProperties) ProtoMessage() {}

func (x *NetworkProperties) ProtoReflect() protoreflect.Message {
	mi := &file_tsuki_gov_network_properties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkProperties.ProtoReflect.Descriptor instead.
func (*NetworkProperties) Descriptor() ([]byte, []int) {
	return file_tsuki_gov_network_properties_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkProperties) GetMinTxFee() uint64 {
	if x != nil {
		return x.MinTxFee
	}
	return 0
}

func (x *NetworkProperties) GetMaxTxFee() uint64 {
	if x != nil {
		return x.MaxTxFee
	}
	return 0
}

func (x *NetworkProperties) GetVoteQuorum() uint64 {
	if x != nil {
		return x.VoteQuorum
	}
	return 0
}

func (x *NetworkProperties) GetProposalEndTime() uint64 {
	if x != nil {
		return x.ProposalEndTime
	}
	return 0
}

func (x *NetworkProperties) GetProposalEnactmentTime() uint64 {
	if x != nil {
		return x.ProposalEnactmentTime
	}
	return 0
}

func (x *NetworkProperties) GetMinProposalEndBlocks() uint64 {
	if x != nil {
		return x.MinProposalEndBlocks
	}
	return 0
}

func (x *NetworkProperties) GetMinProposalEnactmentBlocks() uint64 {
	if x != nil {
		return x.MinProposalEnactmentBlocks
	}
	return 0
}

func (x *NetworkProperties) GetEnableForeignFeePayments() bool {
	if x != nil {
		return x.EnableForeignFeePayments
	}
	return false
}

func (x *NetworkProperties) GetMischanceRankDecreaseAmount() uint64 {
	if x != nil {
		return x.MischanceRankDecreaseAmount
	}
	return 0
}

func (x *NetworkProperties) GetInactiveRankDecreasePercent() uint64 {
	if x != nil {
		return x.InactiveRankDecreasePercent
	}
	return 0
}

func (x *NetworkProperties) GetMinValidators() uint64 {
	if x != nil {
		return x.MinValidators
	}
	return 0
}

func (x *NetworkProperties) GetPoorNetworkMaxBankSend() uint64 {
	if x != nil {
		return x.PoorNetworkMaxBankSend
	}
	return 0
}

func (x *NetworkProperties) GetJailMaxTime() uint64 {
	if x != nil {
		return x.JailMaxTime
	}
	return 0
}

func (x *NetworkProperties) GetEnableTokenWhitelist() bool {
	if x != nil {
		return x.EnableTokenWhitelist
	}
	return false
}

func (x *NetworkProperties) GetEnableTokenBlacklist() bool {
	if x != nil {
		return x.EnableTokenBlacklist
	}
	return false
}

var File_tsuki_gov_network_properties_proto protoreflect.FileDescriptor

var file_tsuki_gov_network_properties_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6b, 0x69, 0x72, 0x61, 0x2f, 0x67, 0x6f, 0x76, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x69, 0x72, 0x61, 0x2e, 0x67, 0x6f, 0x76, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x4a, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69,
	0x72, 0x61, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x31, 0xfa,
	0xde, 0x1f, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x22, 0x8a, 0x06, 0x0a, 0x11, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x54, 0x78, 0x46, 0x65, 0x65, 0x12, 0x1c,
	0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x54, 0x78, 0x46, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x76, 0x6f, 0x74, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x12, 0x2a, 0x0a,
	0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x61, 0x63, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x45, 0x6e, 0x61, 0x63, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x14, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x6d, 0x69, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x61, 0x63, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x1a, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45, 0x6e, 0x61, 0x63,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x66, 0x65,
	0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x46,
	0x65, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x6d, 0x69,
	0x73, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x64, 0x65, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x1b, 0x6d, 0x69, 0x73, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x61, 0x6e,
	0x6b, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x43, 0x0a, 0x1e, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x5f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1b, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x52, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6d, 0x69,
	0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x1a, 0x70,
	0x6f, 0x6f, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x78, 0x5f,
	0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x16, 0x70, 0x6f, 0x6f, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x78, 0x42,
	0x61, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6a, 0x61, 0x69, 0x6c, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x6a, 0x61, 0x69, 0x6c, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x77, 0x68, 0x69, 0x74,
	0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2a, 0x8e, 0x06, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x0a, 0x4d,
	0x49, 0x4e, 0x5f, 0x54, 0x58, 0x5f, 0x46, 0x45, 0x45, 0x10, 0x00, 0x1a, 0x0c, 0x8a, 0x9d, 0x20,
	0x08, 0x4d, 0x69, 0x6e, 0x54, 0x78, 0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x4d, 0x41, 0x58,
	0x5f, 0x54, 0x58, 0x5f, 0x46, 0x45, 0x45, 0x10, 0x01, 0x1a, 0x0c, 0x8a, 0x9d, 0x20, 0x08, 0x4d,
	0x61, 0x78, 0x54, 0x78, 0x46, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x56, 0x4f, 0x54, 0x45, 0x5f,
	0x51, 0x55, 0x4f, 0x52, 0x55, 0x4d, 0x10, 0x02, 0x1a, 0x0e, 0x8a, 0x9d, 0x20, 0x0a, 0x56, 0x6f,
	0x74, 0x65, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x50,
	0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x1a,
	0x13, 0x8a, 0x9d, 0x20, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c,
	0x5f, 0x45, 0x4e, 0x41, 0x43, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10,
	0x04, 0x1a, 0x19, 0x8a, 0x9d, 0x20, 0x15, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45,
	0x6e, 0x61, 0x63, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x17,
	0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x45, 0x4e, 0x44,
	0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x53, 0x10, 0x05, 0x1a, 0x18, 0x8a, 0x9d, 0x20, 0x14, 0x4d,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f,
	0x53, 0x41, 0x4c, 0x5f, 0x45, 0x4e, 0x41, 0x43, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x53, 0x10, 0x06, 0x1a, 0x1e, 0x8a, 0x9d, 0x20, 0x1a, 0x4d, 0x69, 0x6e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x45, 0x6e, 0x61, 0x63, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x46, 0x4f, 0x52, 0x45, 0x49, 0x47, 0x4e, 0x5f, 0x46, 0x45, 0x45, 0x5f, 0x50, 0x41, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x07, 0x1a, 0x1c, 0x8a, 0x9d, 0x20, 0x18, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x46, 0x65, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x4d, 0x49, 0x53, 0x43, 0x48, 0x41, 0x4e,
	0x43, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x44, 0x45, 0x43, 0x52, 0x45, 0x41, 0x53, 0x45,
	0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x08, 0x1a, 0x1f, 0x8a, 0x9d, 0x20, 0x1b, 0x4d,
	0x69, 0x73, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x63, 0x72,
	0x65, 0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x1e, 0x49, 0x4e,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x44, 0x45, 0x43, 0x52,
	0x45, 0x41, 0x53, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x09, 0x1a, 0x1f,
	0x8a, 0x9d, 0x20, 0x1b, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x6e, 0x6b,
	0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12,
	0x3a, 0x0a, 0x1a, 0x50, 0x4f, 0x4f, 0x52, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x4d, 0x41, 0x58, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x10, 0x0a, 0x1a,
	0x1a, 0x8a, 0x9d, 0x20, 0x16, 0x50, 0x6f, 0x6f, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x4d,
	0x49, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x4f, 0x52, 0x53, 0x10, 0x0b, 0x1a,
	0x11, 0x8a, 0x9d, 0x20, 0x0d, 0x4d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x4a, 0x41, 0x49, 0x4c, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x10, 0x0c, 0x1a, 0x0f, 0x8a, 0x9d, 0x20, 0x0b, 0x4a, 0x61, 0x69, 0x6c, 0x4d,
	0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x57, 0x48, 0x49, 0x54, 0x45, 0x4c, 0x49, 0x53, 0x54,
	0x10, 0x0d, 0x1a, 0x18, 0x8a, 0x9d, 0x20, 0x14, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16,
	0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x42, 0x4c, 0x41,
	0x43, 0x4b, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0e, 0x1a, 0x18, 0x8a, 0x9d, 0x20, 0x14, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x1a, 0x04, 0x88, 0xa3, 0x1e, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x72, 0x61, 0x43, 0x6f, 0x72, 0x65, 0x2f,
	0x73, 0x65, 0x6b, 0x61, 0x69, 0x2f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x58, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tsuki_gov_network_properties_proto_rawDescOnce sync.Once
	file_tsuki_gov_network_properties_proto_rawDescData = file_tsuki_gov_network_properties_proto_rawDesc
)

func file_tsuki_gov_network_properties_proto_rawDescGZIP() []byte {
	file_tsuki_gov_network_properties_proto_rawDescOnce.Do(func() {
		file_tsuki_gov_network_properties_proto_rawDescData = protoimpl.X.CompressGZIP(file_tsuki_gov_network_properties_proto_rawDescData)
	})
	return file_tsuki_gov_network_properties_proto_rawDescData
}

var file_tsuki_gov_network_properties_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tsuki_gov_network_properties_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tsuki_gov_network_properties_proto_goTypes = []interface{}{
	(NetworkProperty)(0),            // 0: tsuki.gov.NetworkProperty
	(*MsgSetNetworkProperties)(nil), // 1: tsuki.gov.MsgSetNetworkProperties
	(*NetworkProperties)(nil),       // 2: tsuki.gov.NetworkProperties
}
var file_tsuki_gov_network_properties_proto_depIdxs = []int32{
	2, // 0: tsuki.gov.MsgSetNetworkProperties.network_properties:type_name -> tsuki.gov.NetworkProperties
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tsuki_gov_network_properties_proto_init() }
func file_tsuki_gov_network_properties_proto_init() {
	if File_tsuki_gov_network_properties_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tsuki_gov_network_properties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSetNetworkProperties); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tsuki_gov_network_properties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkProperties); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tsuki_gov_network_properties_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tsuki_gov_network_properties_proto_goTypes,
		DependencyIndexes: file_tsuki_gov_network_properties_proto_depIdxs,
		EnumInfos:         file_tsuki_gov_network_properties_proto_enumTypes,
		MessageInfos:      file_tsuki_gov_network_properties_proto_msgTypes,
	}.Build()
	File_tsuki_gov_network_properties_proto = out.File
	file_tsuki_gov_network_properties_proto_rawDesc = nil
	file_tsuki_gov_network_properties_proto_goTypes = nil
	file_tsuki_gov_network_properties_proto_depIdxs = nil
}
