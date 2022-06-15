// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/multistaking/multistaking.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type StakingPool struct {
	Id                 uint64                                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Validator          string                                    `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator,omitempty"`
	Enabled            bool                                      `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	TotalStakingTokens []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,rep,name=total_staking_tokens,json=totalStakingTokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"total_staking_tokens"`
	TotalShareTokens   []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,rep,name=total_share_tokens,json=totalShareTokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"total_share_tokens"`
	TotalRewards       []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,rep,name=total_rewards,json=totalRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"total_rewards"`
}

func (m *StakingPool) Reset()         { *m = StakingPool{} }
func (m *StakingPool) String() string { return proto.CompactTextString(m) }
func (*StakingPool) ProtoMessage()    {}
func (*StakingPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_a141b2df58d204b4, []int{0}
}
func (m *StakingPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakingPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakingPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakingPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingPool.Merge(m, src)
}
func (m *StakingPool) XXX_Size() int {
	return m.Size()
}
func (m *StakingPool) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingPool.DiscardUnknown(m)
}

var xxx_messageInfo_StakingPool proto.InternalMessageInfo

func (m *StakingPool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StakingPool) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *StakingPool) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type Undelegation struct {
	Id      uint64                                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string                                    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Expiry  uint64                                    `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Amount  []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,rep,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
}

func (m *Undelegation) Reset()         { *m = Undelegation{} }
func (m *Undelegation) String() string { return proto.CompactTextString(m) }
func (*Undelegation) ProtoMessage()    {}
func (*Undelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a141b2df58d204b4, []int{1}
}
func (m *Undelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Undelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Undelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Undelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Undelegation.Merge(m, src)
}
func (m *Undelegation) XXX_Size() int {
	return m.Size()
}
func (m *Undelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_Undelegation.DiscardUnknown(m)
}

var xxx_messageInfo_Undelegation proto.InternalMessageInfo

func (m *Undelegation) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Undelegation) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Undelegation) GetExpiry() uint64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type Rewards struct {
	Delegator string                                    `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
	Rewards   []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,rep,name=rewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"rewards"`
}

func (m *Rewards) Reset()         { *m = Rewards{} }
func (m *Rewards) String() string { return proto.CompactTextString(m) }
func (*Rewards) ProtoMessage()    {}
func (*Rewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_a141b2df58d204b4, []int{2}
}
func (m *Rewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rewards.Merge(m, src)
}
func (m *Rewards) XXX_Size() int {
	return m.Size()
}
func (m *Rewards) XXX_DiscardUnknown() {
	xxx_messageInfo_Rewards.DiscardUnknown(m)
}

var xxx_messageInfo_Rewards proto.InternalMessageInfo

func (m *Rewards) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func init() {
	proto.RegisterType((*StakingPool)(nil), "tsuki.multistaking.StakingPool")
	proto.RegisterType((*Undelegation)(nil), "tsuki.multistaking.Undelegation")
	proto.RegisterType((*Rewards)(nil), "tsuki.multistaking.Rewards")
}

func init() {
	proto.RegisterFile("tsuki/multistaking/multistaking.proto", fileDescriptor_a141b2df58d204b4)
}

var fileDescriptor_a141b2df58d204b4 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbf, 0x6e, 0xd5, 0x30,
	0x18, 0xc5, 0xaf, 0xd3, 0xcb, 0x0d, 0x31, 0x05, 0x81, 0x55, 0xa1, 0xa8, 0x42, 0xb9, 0xd1, 0x15,
	0x12, 0x59, 0x88, 0x07, 0xde, 0xa0, 0x1d, 0x10, 0x82, 0x01, 0x85, 0xb2, 0x20, 0xa1, 0xca, 0xa9,
	0x4d, 0x6a, 0xe5, 0xcf, 0x17, 0xd9, 0x0e, 0xb4, 0x6f, 0xc1, 0xca, 0x1b, 0xf0, 0x28, 0x1d, 0x3b,
	0x22, 0x86, 0x0a, 0xe5, 0xce, 0xbc, 0x03, 0x8a, 0xe3, 0xa8, 0xbd, 0x62, 0xcb, 0x94, 0x1c, 0xfb,
	0xf3, 0xef, 0x1c, 0x9d, 0xc4, 0xf8, 0x79, 0x29, 0x15, 0xa3, 0x75, 0x57, 0x19, 0xa9, 0x0d, 0x2b,
	0x65, 0x53, 0xec, 0x88, 0xb4, 0x55, 0x60, 0x80, 0x3c, 0x19, 0xa6, 0xd2, 0xbb, 0x1b, 0x87, 0x07,
	0x05, 0x14, 0x60, 0x77, 0xe9, 0xf0, 0x36, 0x0e, 0x1e, 0xae, 0x0b, 0x80, 0xa2, 0x12, 0xd4, 0xaa,
	0xbc, 0xfb, 0x42, 0x8d, 0xac, 0x85, 0x36, 0xac, 0x6e, 0xc7, 0x81, 0xcd, 0x5f, 0x0f, 0x3f, 0xf8,
	0x30, 0x22, 0xde, 0x03, 0x54, 0xe4, 0x11, 0xf6, 0x24, 0x0f, 0x51, 0x8c, 0x92, 0x65, 0xe6, 0x49,
	0x4e, 0x9e, 0xe1, 0xe0, 0x2b, 0xab, 0x24, 0x67, 0x06, 0x54, 0xe8, 0xc5, 0x28, 0x09, 0xb2, 0xdb,
	0x05, 0x12, 0x62, 0x5f, 0x34, 0x2c, 0xaf, 0x04, 0x0f, 0xf7, 0x62, 0x94, 0xdc, 0xcf, 0x26, 0x49,
	0x18, 0x3e, 0x30, 0x60, 0x58, 0x75, 0xea, 0xf2, 0x9d, 0x1a, 0x28, 0x45, 0xa3, 0xc3, 0x65, 0xbc,
	0x97, 0x04, 0x47, 0xf4, 0xea, 0x66, 0xbd, 0xf8, 0x7d, 0xb3, 0x7e, 0x51, 0x48, 0x73, 0xde, 0xe5,
	0xe9, 0x19, 0xd4, 0xf4, 0x0c, 0x74, 0x0d, 0xda, 0x3d, 0x5e, 0x6a, 0x5e, 0x52, 0x73, 0xd9, 0x0a,
	0x9d, 0x1e, 0x83, 0x6c, 0x32, 0x62, 0x61, 0x2e, 0xe8, 0x89, 0x45, 0x91, 0xcf, 0x98, 0x38, 0x8b,
	0x73, 0xa6, 0xc4, 0x64, 0x70, 0x6f, 0x9e, 0xc1, 0xe3, 0xd1, 0x60, 0x20, 0x39, 0xfc, 0x09, 0x7e,
	0x38, 0xe2, 0x95, 0xf8, 0xc6, 0x14, 0xd7, 0xe1, 0x6a, 0x1e, 0x79, 0xdf, 0x52, 0xb2, 0x11, 0xb2,
	0xf9, 0x81, 0xf0, 0xfe, 0xc7, 0x86, 0x8b, 0x4a, 0x14, 0xcc, 0x48, 0x68, 0xfe, 0x2b, 0x3c, 0xc4,
	0x3e, 0xe3, 0x5c, 0x09, 0xad, 0x5d, 0xdd, 0x93, 0x24, 0x4f, 0xf1, 0x4a, 0x5c, 0xb4, 0x52, 0x5d,
	0xda, 0xae, 0x97, 0x99, 0x53, 0xe4, 0x35, 0x5e, 0xb1, 0x1a, 0xba, 0xc6, 0xcc, 0x2d, 0xd7, 0x1d,
	0xdf, 0x28, 0xec, 0xbb, 0x98, 0xc3, 0x67, 0x77, 0x19, 0x41, 0xd9, 0x70, 0x41, 0x76, 0xbb, 0x40,
	0xde, 0x60, 0x7f, 0x2a, 0xc5, 0x9b, 0x67, 0x39, 0x9d, 0x3f, 0x7a, 0xf7, 0xb3, 0x8f, 0xd0, 0x55,
	0x1f, 0xa1, 0xeb, 0x3e, 0x42, 0x7f, 0xfa, 0x08, 0x7d, 0xdf, 0x46, 0x8b, 0xeb, 0x6d, 0xb4, 0xf8,
	0xb5, 0x8d, 0x16, 0x9f, 0xd2, 0x3b, 0xbc, 0xb7, 0x52, 0xb1, 0x63, 0x50, 0x82, 0x6a, 0x51, 0x32,
	0x49, 0x2f, 0x76, 0x2f, 0x89, 0x65, 0xe7, 0x2b, 0xfb, 0x53, 0xbf, 0xfa, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xbd, 0x49, 0xd9, 0x52, 0x46, 0x03, 0x00, 0x00,
}

func (this *StakingPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StakingPool)
	if !ok {
		that2, ok := that.(StakingPool)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Validator != that1.Validator {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if len(this.TotalStakingTokens) != len(that1.TotalStakingTokens) {
		return false
	}
	for i := range this.TotalStakingTokens {
		if !this.TotalStakingTokens[i].Equal(that1.TotalStakingTokens[i]) {
			return false
		}
	}
	if len(this.TotalShareTokens) != len(that1.TotalShareTokens) {
		return false
	}
	for i := range this.TotalShareTokens {
		if !this.TotalShareTokens[i].Equal(that1.TotalShareTokens[i]) {
			return false
		}
	}
	if len(this.TotalRewards) != len(that1.TotalRewards) {
		return false
	}
	for i := range this.TotalRewards {
		if !this.TotalRewards[i].Equal(that1.TotalRewards[i]) {
			return false
		}
	}
	return true
}
func (this *Undelegation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Undelegation)
	if !ok {
		that2, ok := that.(Undelegation)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Expiry != that1.Expiry {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(that1.Amount[i]) {
			return false
		}
	}
	return true
}
func (this *Rewards) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Rewards)
	if !ok {
		that2, ok := that.(Rewards)
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
	if this.Delegator != that1.Delegator {
		return false
	}
	if len(this.Rewards) != len(that1.Rewards) {
		return false
	}
	for i := range this.Rewards {
		if !this.Rewards[i].Equal(that1.Rewards[i]) {
			return false
		}
	}
	return true
}
func (m *StakingPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakingPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakingPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalRewards) > 0 {
		for iNdEx := len(m.TotalRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.TotalRewards[iNdEx].Size()
				i -= size
				if _, err := m.TotalRewards[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintMultistaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TotalShareTokens) > 0 {
		for iNdEx := len(m.TotalShareTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.TotalShareTokens[iNdEx].Size()
				i -= size
				if _, err := m.TotalShareTokens[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintMultistaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TotalStakingTokens) > 0 {
		for iNdEx := len(m.TotalStakingTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.TotalStakingTokens[iNdEx].Size()
				i -= size
				if _, err := m.TotalStakingTokens[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintMultistaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintMultistaking(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMultistaking(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Undelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Undelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Undelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Amount[iNdEx].Size()
				i -= size
				if _, err := m.Amount[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintMultistaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Expiry != 0 {
		i = encodeVarintMultistaking(dAtA, i, uint64(m.Expiry))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMultistaking(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMultistaking(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Rewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Rewards[iNdEx].Size()
				i -= size
				if _, err := m.Rewards[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintMultistaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintMultistaking(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMultistaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovMultistaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakingPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMultistaking(uint64(m.Id))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovMultistaking(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.TotalStakingTokens) > 0 {
		for _, e := range m.TotalStakingTokens {
			l = e.Size()
			n += 1 + l + sovMultistaking(uint64(l))
		}
	}
	if len(m.TotalShareTokens) > 0 {
		for _, e := range m.TotalShareTokens {
			l = e.Size()
			n += 1 + l + sovMultistaking(uint64(l))
		}
	}
	if len(m.TotalRewards) > 0 {
		for _, e := range m.TotalRewards {
			l = e.Size()
			n += 1 + l + sovMultistaking(uint64(l))
		}
	}
	return n
}

func (m *Undelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMultistaking(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMultistaking(uint64(l))
	}
	if m.Expiry != 0 {
		n += 1 + sovMultistaking(uint64(m.Expiry))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovMultistaking(uint64(l))
		}
	}
	return n
}

func (m *Rewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovMultistaking(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovMultistaking(uint64(l))
		}
	}
	return n
}

func sovMultistaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMultistaking(x uint64) (n int) {
	return sovMultistaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakingPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultistaking
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
			return fmt.Errorf("proto: StakingPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakingPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStakingTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.TotalStakingTokens = append(m.TotalStakingTokens, v)
			if err := m.TotalStakingTokens[len(m.TotalStakingTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShareTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.TotalShareTokens = append(m.TotalShareTokens, v)
			if err := m.TotalShareTokens[len(m.TotalShareTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.TotalRewards = append(m.TotalRewards, v)
			if err := m.TotalRewards[len(m.TotalRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultistaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultistaking
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
func (m *Undelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultistaking
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
			return fmt.Errorf("proto: Undelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Undelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiry |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Amount = append(m.Amount, v)
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultistaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultistaking
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
func (m *Rewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultistaking
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
			return fmt.Errorf("proto: Rewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultistaking
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
				return ErrInvalidLengthMultistaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultistaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Rewards = append(m.Rewards, v)
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultistaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultistaking
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
func skipMultistaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMultistaking
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
					return 0, ErrIntOverflowMultistaking
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
					return 0, ErrIntOverflowMultistaking
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
				return 0, ErrInvalidLengthMultistaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMultistaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMultistaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMultistaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMultistaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMultistaking = fmt.Errorf("proto: unexpected end of group")
)
