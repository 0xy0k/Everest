// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/spending/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// proposal-spending-pool-update - a function to create a proposal allowing
// modification of the existing spending pool, adding owners, beneficiaries,
// or otherwise editing any of the existing properties.
type UpdateSpendingPoolProposal struct {
	Name          string                                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClaimStart    uint64                                       `protobuf:"varint,2,opt,name=claim_start,json=claimStart,proto3" json:"claim_start,omitempty"`
	ClaimEnd      uint64                                       `protobuf:"varint,3,opt,name=claim_end,json=claimEnd,proto3" json:"claim_end,omitempty"`
	Token         string                                       `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Rates         []github_com_cosmos_cosmos_sdk_types.DecCoin `protobuf:"bytes,5,rep,name=rates,proto3,customtype=github.com/cosmos/cosmos-sdk/types.DecCoin" json:"rates" yaml:"rates"`
	VoteQuorum    uint64                                       `protobuf:"varint,6,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	VotePeriod    uint64                                       `protobuf:"varint,7,opt,name=vote_period,json=votePeriod,proto3" json:"vote_period,omitempty"`
	VoteEnactment uint64                                       `protobuf:"varint,8,opt,name=vote_enactment,json=voteEnactment,proto3" json:"vote_enactment,omitempty"`
	Owners        PermInfo                                     `protobuf:"bytes,9,opt,name=owners,proto3" json:"owners"`
	Beneficiaries WeightedPermInfo                             `protobuf:"bytes,10,opt,name=beneficiaries,proto3" json:"beneficiaries"`
}

func (m *UpdateSpendingPoolProposal) Reset()         { *m = UpdateSpendingPoolProposal{} }
func (m *UpdateSpendingPoolProposal) String() string { return proto.CompactTextString(m) }
func (*UpdateSpendingPoolProposal) ProtoMessage()    {}
func (*UpdateSpendingPoolProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e006ef21562b5bc9, []int{0}
}
func (m *UpdateSpendingPoolProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateSpendingPoolProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateSpendingPoolProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateSpendingPoolProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSpendingPoolProposal.Merge(m, src)
}
func (m *UpdateSpendingPoolProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateSpendingPoolProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSpendingPoolProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSpendingPoolProposal proto.InternalMessageInfo

func (m *UpdateSpendingPoolProposal) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateSpendingPoolProposal) GetClaimStart() uint64 {
	if m != nil {
		return m.ClaimStart
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetClaimEnd() uint64 {
	if m != nil {
		return m.ClaimEnd
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateSpendingPoolProposal) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetVotePeriod() uint64 {
	if m != nil {
		return m.VotePeriod
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetVoteEnactment() uint64 {
	if m != nil {
		return m.VoteEnactment
	}
	return 0
}

func (m *UpdateSpendingPoolProposal) GetOwners() PermInfo {
	if m != nil {
		return m.Owners
	}
	return PermInfo{}
}

func (m *UpdateSpendingPoolProposal) GetBeneficiaries() WeightedPermInfo {
	if m != nil {
		return m.Beneficiaries
	}
	return WeightedPermInfo{}
}

// SpendingPoolDistributionProposal - force distribution of tokens to all
// beneficiaries registered in the claims array (this function should be
// automatically triggered before upgrades are executed)
type SpendingPoolDistributionProposal struct {
	PoolName string `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
}

func (m *SpendingPoolDistributionProposal) Reset()         { *m = SpendingPoolDistributionProposal{} }
func (m *SpendingPoolDistributionProposal) String() string { return proto.CompactTextString(m) }
func (*SpendingPoolDistributionProposal) ProtoMessage()    {}
func (*SpendingPoolDistributionProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e006ef21562b5bc9, []int{1}
}
func (m *SpendingPoolDistributionProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpendingPoolDistributionProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpendingPoolDistributionProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpendingPoolDistributionProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendingPoolDistributionProposal.Merge(m, src)
}
func (m *SpendingPoolDistributionProposal) XXX_Size() int {
	return m.Size()
}
func (m *SpendingPoolDistributionProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendingPoolDistributionProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SpendingPoolDistributionProposal proto.InternalMessageInfo

func (m *SpendingPoolDistributionProposal) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

// SpendingPoolWithdrawProposal - proposal allowing withdrawal of funds
// from the pool to one or many specified accounts. Withdrawal should only
// be possible if the receiving account/s are on the list of registered
// beneficiaries.
type SpendingPoolWithdrawProposal struct {
	PoolName      string                                    `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	Beneficiaries []string                                  `protobuf:"bytes,2,rep,name=beneficiaries,proto3" json:"beneficiaries,omitempty"`
	Amounts       []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,rep,name=amounts,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amounts"`
}

func (m *SpendingPoolWithdrawProposal) Reset()         { *m = SpendingPoolWithdrawProposal{} }
func (m *SpendingPoolWithdrawProposal) String() string { return proto.CompactTextString(m) }
func (*SpendingPoolWithdrawProposal) ProtoMessage()    {}
func (*SpendingPoolWithdrawProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e006ef21562b5bc9, []int{2}
}
func (m *SpendingPoolWithdrawProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpendingPoolWithdrawProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpendingPoolWithdrawProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpendingPoolWithdrawProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendingPoolWithdrawProposal.Merge(m, src)
}
func (m *SpendingPoolWithdrawProposal) XXX_Size() int {
	return m.Size()
}
func (m *SpendingPoolWithdrawProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendingPoolWithdrawProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SpendingPoolWithdrawProposal proto.InternalMessageInfo

func (m *SpendingPoolWithdrawProposal) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *SpendingPoolWithdrawProposal) GetBeneficiaries() []string {
	if m != nil {
		return m.Beneficiaries
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateSpendingPoolProposal)(nil), "tsuki.gov.UpdateSpendingPoolProposal")
	proto.RegisterType((*SpendingPoolDistributionProposal)(nil), "tsuki.gov.SpendingPoolDistributionProposal")
	proto.RegisterType((*SpendingPoolWithdrawProposal)(nil), "tsuki.gov.SpendingPoolWithdrawProposal")
}

func init() { proto.RegisterFile("tsuki/spending/proposal.proto", fileDescriptor_e006ef21562b5bc9) }

var fileDescriptor_e006ef21562b5bc9 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x4f, 0x13, 0x41,
	0x14, 0xee, 0xd2, 0x02, 0xed, 0x20, 0x9a, 0x4c, 0x48, 0x1c, 0x81, 0x74, 0x9b, 0x46, 0x63, 0x63,
	0x42, 0x37, 0xd1, 0x78, 0xe1, 0x58, 0xe0, 0x40, 0x48, 0x4c, 0x2d, 0x31, 0x18, 0x2f, 0xcd, 0x74,
	0xf7, 0xb1, 0x4c, 0xda, 0x99, 0xb7, 0xce, 0xcc, 0x82, 0xfc, 0x0b, 0x7f, 0x82, 0x3f, 0xc2, 0xab,
	0x77, 0x8e, 0xc4, 0x93, 0x31, 0x86, 0x18, 0x7a, 0xf1, 0xec, 0x2f, 0x30, 0x3b, 0xbb, 0x85, 0xba,
	0x5c, 0x38, 0x75, 0xe6, 0x7b, 0xdf, 0xfb, 0xfa, 0xbd, 0x7d, 0xdf, 0x90, 0xcd, 0xb1, 0xd0, 0x3c,
	0x30, 0x09, 0xa8, 0x48, 0xa8, 0x38, 0x48, 0x34, 0x26, 0x68, 0xf8, 0xa4, 0x9b, 0x68, 0xb4, 0x48,
	0xeb, 0x59, 0xb5, 0x1b, 0xe3, 0xe9, 0xfa, 0x5a, 0x8c, 0x31, 0x3a, 0x30, 0xc8, 0x4e, 0x79, 0x7d,
	0xdd, 0x8f, 0x11, 0xe3, 0x09, 0x04, 0xee, 0x36, 0x4a, 0x8f, 0x03, 0x2b, 0x24, 0x18, 0xcb, 0x65,
	0x52, 0x10, 0x9e, 0x94, 0x09, 0x5c, 0x9d, 0xcf, 0x4a, 0x21, 0x1a, 0x89, 0x66, 0x98, 0x8b, 0xe6,
	0x97, 0xa2, 0xc4, 0x4a, 0xa6, 0x10, 0x0b, 0x43, 0xed, 0x5f, 0x55, 0xb2, 0xfe, 0x2e, 0x89, 0xb8,
	0x85, 0xc3, 0xa2, 0xda, 0x47, 0x9c, 0xf4, 0x0b, 0xd7, 0x94, 0x92, 0x9a, 0xe2, 0x12, 0x98, 0xd7,
	0xf2, 0x3a, 0x8d, 0x81, 0x3b, 0x53, 0x9f, 0xac, 0x84, 0x13, 0x2e, 0xe4, 0xd0, 0x58, 0xae, 0x2d,
	0x5b, 0x68, 0x79, 0x9d, 0xda, 0x80, 0x38, 0xe8, 0x30, 0x43, 0xe8, 0x06, 0x69, 0xe4, 0x04, 0x50,
	0x11, 0xab, 0xba, 0x72, 0xdd, 0x01, 0x7b, 0x2a, 0xa2, 0x6b, 0x64, 0xd1, 0xe2, 0x18, 0x14, 0xab,
	0x39, 0xc9, 0xfc, 0x42, 0xdf, 0x93, 0x45, 0xcd, 0x2d, 0x18, 0xb6, 0xd8, 0xaa, 0x76, 0x1a, 0xbd,
	0xde, 0xc5, 0x95, 0x5f, 0xf9, 0x79, 0xe5, 0xbf, 0x88, 0x85, 0x3d, 0x49, 0x47, 0xdd, 0x10, 0x65,
	0x31, 0x50, 0xf1, 0xb3, 0x65, 0xa2, 0x71, 0x60, 0xcf, 0x13, 0x30, 0xdd, 0x5d, 0x08, 0x77, 0x50,
	0xa8, 0xbf, 0x57, 0xfe, 0x83, 0x73, 0x2e, 0x27, 0xdb, 0x6d, 0x27, 0xd4, 0x1e, 0xe4, 0x82, 0x99,
	0xdb, 0x53, 0xb4, 0x30, 0xfc, 0x98, 0xa2, 0x4e, 0x25, 0x5b, 0xca, 0xdd, 0x66, 0xd0, 0x5b, 0x87,
	0xdc, 0x10, 0x12, 0xd0, 0x02, 0x23, 0xb6, 0x7c, 0x4b, 0xe8, 0x3b, 0x84, 0x3e, 0x23, 0x0f, 0x1d,
	0x01, 0x14, 0x0f, 0xad, 0x04, 0x65, 0x59, 0xdd, 0x71, 0x56, 0x33, 0x74, 0x6f, 0x06, 0xd2, 0xd7,
	0x64, 0x09, 0xcf, 0x14, 0x68, 0xc3, 0x1a, 0x2d, 0xaf, 0xb3, 0xf2, 0xf2, 0x71, 0xd7, 0xed, 0x7a,
	0xf6, 0xd1, 0xbb, 0x7d, 0xd0, 0x72, 0x5f, 0x1d, 0x63, 0xaf, 0x96, 0x0d, 0x37, 0x28, 0xc8, 0xf4,
	0x80, 0xac, 0x8e, 0x40, 0xc1, 0xb1, 0x08, 0x05, 0xd7, 0x02, 0x0c, 0x23, 0xae, 0xdb, 0x2f, 0x75,
	0x1f, 0x81, 0x88, 0x4f, 0x2c, 0x44, 0x25, 0x95, 0xff, 0x7b, 0xb7, 0x1f, 0xfd, 0xf9, 0xe2, 0x7b,
	0xdf, 0xbf, 0x6e, 0x2d, 0xef, 0xa0, 0xb2, 0xa0, 0x6c, 0xbb, 0x4f, 0x5a, 0xf3, 0x7b, 0xdd, 0x15,
	0xc6, 0x6a, 0x31, 0x4a, 0xad, 0x40, 0x75, 0xb3, 0xe3, 0x0d, 0xd2, 0xc8, 0x02, 0x31, 0x9c, 0x5b,
	0x74, 0x3d, 0x03, 0xde, 0x70, 0x09, 0x77, 0x15, 0xbf, 0x79, 0x64, 0x73, 0x5e, 0xf2, 0x48, 0xd8,
	0x93, 0x48, 0xf3, 0xb3, 0x7b, 0xc9, 0xd1, 0xa7, 0xe5, 0x69, 0x17, 0xb2, 0x7d, 0x97, 0xc6, 0xa0,
	0xfb, 0x64, 0x99, 0x4b, 0x4c, 0x95, 0x35, 0xac, 0xea, 0xf2, 0x10, 0x14, 0x79, 0x78, 0x7e, 0x8f,
	0x3c, 0x64, 0x61, 0x18, 0xcc, 0xfa, 0xef, 0xf8, 0xef, 0xed, 0x5e, 0x5c, 0x37, 0xbd, 0xcb, 0xeb,
	0xa6, 0xf7, 0xfb, 0xba, 0xe9, 0x7d, 0x9e, 0x36, 0x2b, 0x97, 0xd3, 0x66, 0xe5, 0xc7, 0xb4, 0x59,
	0xf9, 0x30, 0x1f, 0xb6, 0x03, 0xa1, 0xf9, 0x0e, 0x6a, 0x08, 0x0c, 0x8c, 0xb9, 0x08, 0x3e, 0xdd,
	0xbe, 0x1d, 0xf7, 0x27, 0xa3, 0x25, 0xf7, 0x7a, 0x5e, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x12,
	0x98, 0xf3, 0xc9, 0xee, 0x03, 0x00, 0x00,
}

func (this *UpdateSpendingPoolProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateSpendingPoolProposal)
	if !ok {
		that2, ok := that.(UpdateSpendingPoolProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.ClaimStart != that1.ClaimStart {
		return false
	}
	if this.ClaimEnd != that1.ClaimEnd {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	if len(this.Rates) != len(that1.Rates) {
		return false
	}
	for i := range this.Rates {
		if !this.Rates[i].Equal(that1.Rates[i]) {
			return false
		}
	}
	if this.VoteQuorum != that1.VoteQuorum {
		return false
	}
	if this.VotePeriod != that1.VotePeriod {
		return false
	}
	if this.VoteEnactment != that1.VoteEnactment {
		return false
	}
	if !this.Owners.Equal(&that1.Owners) {
		return false
	}
	if !this.Beneficiaries.Equal(&that1.Beneficiaries) {
		return false
	}
	return true
}
func (this *SpendingPoolDistributionProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpendingPoolDistributionProposal)
	if !ok {
		that2, ok := that.(SpendingPoolDistributionProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolName != that1.PoolName {
		return false
	}
	return true
}
func (this *SpendingPoolWithdrawProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpendingPoolWithdrawProposal)
	if !ok {
		that2, ok := that.(SpendingPoolWithdrawProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolName != that1.PoolName {
		return false
	}
	if len(this.Beneficiaries) != len(that1.Beneficiaries) {
		return false
	}
	for i := range this.Beneficiaries {
		if this.Beneficiaries[i] != that1.Beneficiaries[i] {
			return false
		}
	}
	if len(this.Amounts) != len(that1.Amounts) {
		return false
	}
	for i := range this.Amounts {
		if !this.Amounts[i].Equal(that1.Amounts[i]) {
			return false
		}
	}
	return true
}
func (m *UpdateSpendingPoolProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateSpendingPoolProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSpendingPoolProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Beneficiaries.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size, err := m.Owners.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.VoteEnactment != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VoteEnactment))
		i--
		dAtA[i] = 0x40
	}
	if m.VotePeriod != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VotePeriod))
		i--
		dAtA[i] = 0x38
	}
	if m.VoteQuorum != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VoteQuorum))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Rates) > 0 {
		for iNdEx := len(m.Rates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Rates[iNdEx].Size()
				i -= size
				if _, err := m.Rates[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x22
	}
	if m.ClaimEnd != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ClaimEnd))
		i--
		dAtA[i] = 0x18
	}
	if m.ClaimStart != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ClaimStart))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpendingPoolDistributionProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendingPoolDistributionProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpendingPoolDistributionProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpendingPoolWithdrawProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpendingPoolWithdrawProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpendingPoolWithdrawProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amounts) > 0 {
		for iNdEx := len(m.Amounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Amounts[iNdEx].Size()
				i -= size
				if _, err := m.Amounts[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Beneficiaries) > 0 {
		for iNdEx := len(m.Beneficiaries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Beneficiaries[iNdEx])
			copy(dAtA[i:], m.Beneficiaries[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Beneficiaries[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateSpendingPoolProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.ClaimStart != 0 {
		n += 1 + sovProposal(uint64(m.ClaimStart))
	}
	if m.ClaimEnd != 0 {
		n += 1 + sovProposal(uint64(m.ClaimEnd))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Rates) > 0 {
		for _, e := range m.Rates {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if m.VoteQuorum != 0 {
		n += 1 + sovProposal(uint64(m.VoteQuorum))
	}
	if m.VotePeriod != 0 {
		n += 1 + sovProposal(uint64(m.VotePeriod))
	}
	if m.VoteEnactment != 0 {
		n += 1 + sovProposal(uint64(m.VoteEnactment))
	}
	l = m.Owners.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.Beneficiaries.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func (m *SpendingPoolDistributionProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *SpendingPoolWithdrawProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Beneficiaries) > 0 {
		for _, s := range m.Beneficiaries {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.Amounts) > 0 {
		for _, e := range m.Amounts {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateSpendingPoolProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: UpdateSpendingPoolProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateSpendingPoolProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimStart", wireType)
			}
			m.ClaimStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimEnd", wireType)
			}
			m.ClaimEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimEnd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.DecCoin
			m.Rates = append(m.Rates, v)
			if err := m.Rates[len(m.Rates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteQuorum", wireType)
			}
			m.VoteQuorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotePeriod", wireType)
			}
			m.VotePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteEnactment", wireType)
			}
			m.VoteEnactment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteEnactment |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owners.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiaries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Beneficiaries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *SpendingPoolDistributionProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: SpendingPoolDistributionProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendingPoolDistributionProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *SpendingPoolWithdrawProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: SpendingPoolWithdrawProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpendingPoolWithdrawProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiaries", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beneficiaries = append(m.Beneficiaries, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Amounts = append(m.Amounts, v)
			if err := m.Amounts[len(m.Amounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
