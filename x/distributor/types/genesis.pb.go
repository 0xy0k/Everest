// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/distributor/genesis.proto

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

type ValidatorVote struct {
	ConsAddr string `protobuf:"bytes,1,opt,name=cons_addr,json=consAddr,proto3" json:"cons_addr,omitempty"`
	Height   int64  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *ValidatorVote) Reset()         { *m = ValidatorVote{} }
func (m *ValidatorVote) String() string { return proto.CompactTextString(m) }
func (*ValidatorVote) ProtoMessage()    {}
func (*ValidatorVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_e815530f0f0e0b78, []int{0}
}
func (m *ValidatorVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorVote.Merge(m, src)
}
func (m *ValidatorVote) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorVote) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorVote.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorVote proto.InternalMessageInfo

func (m *ValidatorVote) GetConsAddr() string {
	if m != nil {
		return m.ConsAddr
	}
	return ""
}

func (m *ValidatorVote) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type SupplySnapshot struct {
	SnapshotTime   int64                                  `protobuf:"varint,1,opt,name=snapshot_time,json=snapshotTime,proto3" json:"snapshot_time,omitempty"`
	SnapshotAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=snapshot_amount,json=snapshotAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"snapshot_amount"`
}

func (m *SupplySnapshot) Reset()         { *m = SupplySnapshot{} }
func (m *SupplySnapshot) String() string { return proto.CompactTextString(m) }
func (*SupplySnapshot) ProtoMessage()    {}
func (*SupplySnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_e815530f0f0e0b78, []int{1}
}
func (m *SupplySnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplySnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupplySnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupplySnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplySnapshot.Merge(m, src)
}
func (m *SupplySnapshot) XXX_Size() int {
	return m.Size()
}
func (m *SupplySnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplySnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_SupplySnapshot proto.InternalMessageInfo

func (m *SupplySnapshot) GetSnapshotTime() int64 {
	if m != nil {
		return m.SnapshotTime
	}
	return 0
}

// GenesisState defines the distributor module's genesis state.
type GenesisState struct {
	// fees that are kept in treasury that is not distributed yet - community pool
	FeesTreasury []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,1,rep,name=fees_treasury,json=feesTreasury,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"fees_treasury"`
	// fees collected from genesis
	FeesCollected []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,rep,name=fees_collected,json=feesCollected,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"fees_collected"`
	// number of blocks considered for reward calculation
	SnapPeriod int64 `protobuf:"varint,3,opt,name=snap_period,json=snapPeriod,proto3" json:"snap_period,omitempty"`
	// validator historical votes
	ValidatorVotes []ValidatorVote `protobuf:"bytes,4,rep,name=validator_votes,json=validatorVotes,proto3" json:"validator_votes"`
	// previous proposer
	PreviousProposer string `protobuf:"bytes,5,opt,name=previous_proposer,json=previousProposer,proto3" json:"previous_proposer,omitempty"`
	// year start snapshot
	YearStartSnapshot SupplySnapshot `protobuf:"bytes,6,opt,name=year_start_snapshot,json=yearStartSnapshot,proto3" json:"year_start_snapshot"`
	// period start snapshot
	PeriodicSnapshot SupplySnapshot `protobuf:"bytes,7,opt,name=periodic_snapshot,json=periodicSnapshot,proto3" json:"periodic_snapshot"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e815530f0f0e0b78, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetSnapPeriod() int64 {
	if m != nil {
		return m.SnapPeriod
	}
	return 0
}

func (m *GenesisState) GetValidatorVotes() []ValidatorVote {
	if m != nil {
		return m.ValidatorVotes
	}
	return nil
}

func (m *GenesisState) GetPreviousProposer() string {
	if m != nil {
		return m.PreviousProposer
	}
	return ""
}

func (m *GenesisState) GetYearStartSnapshot() SupplySnapshot {
	if m != nil {
		return m.YearStartSnapshot
	}
	return SupplySnapshot{}
}

func (m *GenesisState) GetPeriodicSnapshot() SupplySnapshot {
	if m != nil {
		return m.PeriodicSnapshot
	}
	return SupplySnapshot{}
}

func init() {
	proto.RegisterType((*ValidatorVote)(nil), "tsuki.distributor.ValidatorVote")
	proto.RegisterType((*SupplySnapshot)(nil), "tsuki.distributor.SupplySnapshot")
	proto.RegisterType((*GenesisState)(nil), "tsuki.distributor.GenesisState")
}

func init() { proto.RegisterFile("tsuki/distributor/genesis.proto", fileDescriptor_e815530f0f0e0b78) }

var fileDescriptor_e815530f0f0e0b78 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x9b, 0xed, 0x5a, 0xed, 0xf4, 0xcf, 0xb6, 0xa3, 0x48, 0x50, 0x48, 0x43, 0x05, 0x2d,
	0xc8, 0x26, 0xb0, 0x3e, 0xc1, 0xb6, 0xc2, 0x22, 0x82, 0x2c, 0x69, 0xa9, 0xe0, 0x4d, 0x98, 0x26,
	0xc7, 0x74, 0x68, 0x93, 0x13, 0x66, 0x26, 0xc5, 0x3e, 0x81, 0x77, 0xe2, 0x63, 0xed, 0xe5, 0x5e,
	0x8a, 0x17, 0x8b, 0xb4, 0x2f, 0x22, 0x99, 0x26, 0xb5, 0xd5, 0x1b, 0xf5, 0x2a, 0x93, 0xef, 0x0c,
	0xbf, 0xf9, 0xe6, 0x7c, 0x73, 0x88, 0xb5, 0xe0, 0x82, 0xb9, 0x21, 0x97, 0x4a, 0xf0, 0x59, 0xa6,
	0x50, 0xb8, 0x11, 0x24, 0x20, 0xb9, 0x74, 0x52, 0x81, 0x0a, 0x69, 0x27, 0xaf, 0x3b, 0x07, 0xf5,
	0x27, 0x8f, 0x22, 0x8c, 0x50, 0x17, 0xdd, 0x7c, 0xb5, 0xdb, 0xd7, 0x7f, 0x4d, 0x5a, 0x53, 0xb6,
	0xe4, 0x21, 0x53, 0x28, 0xa6, 0xa8, 0x80, 0x3e, 0x25, 0xf5, 0x00, 0x13, 0xe9, 0xb3, 0x30, 0x14,
	0xa6, 0x61, 0x1b, 0x83, 0xba, 0xf7, 0x20, 0x17, 0x2e, 0xc3, 0x50, 0xd0, 0xc7, 0xa4, 0x36, 0x07,
	0x1e, 0xcd, 0x95, 0x79, 0x62, 0x1b, 0x83, 0xaa, 0x57, 0xfc, 0xf5, 0xbf, 0x18, 0xa4, 0x3d, 0xce,
	0xd2, 0x74, 0xb9, 0x1e, 0x27, 0x2c, 0x95, 0x73, 0x54, 0xf4, 0x19, 0x69, 0xc9, 0x62, 0xed, 0x2b,
	0x1e, 0x83, 0x66, 0x55, 0xbd, 0x66, 0x29, 0x4e, 0x78, 0x0c, 0xf4, 0x3d, 0x39, 0xdb, 0x6f, 0x62,
	0x31, 0x66, 0xc9, 0x0e, 0x5c, 0x1f, 0x3a, 0x37, 0x77, 0xbd, 0xca, 0xf7, 0xbb, 0xde, 0xf3, 0x88,
	0xab, 0x79, 0x36, 0x73, 0x02, 0x8c, 0xdd, 0x00, 0x65, 0x8c, 0xb2, 0xf8, 0x9c, 0xcb, 0x70, 0xe1,
	0xaa, 0x75, 0x0a, 0xd2, 0x79, 0x93, 0x28, 0xaf, 0x5d, 0x62, 0x2e, 0x35, 0xa5, 0xff, 0xf9, 0x94,
	0x34, 0xaf, 0x76, 0x0d, 0x19, 0x2b, 0xa6, 0x80, 0x4e, 0x48, 0xeb, 0x23, 0x80, 0xf4, 0x95, 0x00,
	0x26, 0x33, 0xb1, 0x36, 0x0d, 0xbb, 0x3a, 0xa8, 0x0f, 0xdd, 0xe2, 0x9c, 0x17, 0x7f, 0x71, 0xce,
	0x08, 0x79, 0xe2, 0x35, 0x73, 0xca, 0xa4, 0x80, 0xd0, 0x29, 0x69, 0x6b, 0x6a, 0x80, 0xcb, 0x25,
	0x04, 0x0a, 0x42, 0xf3, 0xe4, 0xff, 0xb0, 0xda, 0xdc, 0xa8, 0xa4, 0xd0, 0x1e, 0x69, 0xe4, 0x17,
	0xf2, 0x53, 0x10, 0x1c, 0x43, 0xb3, 0xaa, 0x5b, 0x47, 0x72, 0xe9, 0x5a, 0x2b, 0xf4, 0x1d, 0x39,
	0x5b, 0x95, 0xb1, 0xf9, 0x2b, 0x54, 0x20, 0xcd, 0x53, 0xbb, 0x3a, 0x68, 0x5c, 0xf4, 0x9c, 0xdf,
	0x83, 0x77, 0x8e, 0xf2, 0x1d, 0x9e, 0xe6, 0xd6, 0xbc, 0xf6, 0xea, 0x50, 0x94, 0xf4, 0x25, 0xe9,
	0xa6, 0x02, 0x56, 0x1c, 0x33, 0xe9, 0xa7, 0x02, 0x53, 0x94, 0x20, 0xcc, 0x7b, 0x3a, 0xfd, 0x4e,
	0x59, 0xb8, 0x2e, 0x74, 0x3a, 0x25, 0x0f, 0xd7, 0xc0, 0x84, 0x2f, 0x15, 0x13, 0xca, 0x2f, 0x3b,
	0x6f, 0xd6, 0x6c, 0x63, 0xd0, 0xb8, 0xb0, 0xff, 0x34, 0x70, 0xfc, 0x32, 0x0a, 0x07, 0xdd, 0x1c,
	0x31, 0xce, 0x09, 0xfb, 0x27, 0x33, 0x26, 0xdd, 0xdd, 0x85, 0x79, 0xf0, 0x8b, 0x7a, 0xff, 0x9f,
	0xa8, 0x9d, 0x12, 0xb0, 0xd7, 0xaf, 0x6e, 0x36, 0x96, 0x71, 0xbb, 0xb1, 0x8c, 0x1f, 0x1b, 0xcb,
	0xf8, 0xba, 0xb5, 0x2a, 0xb7, 0x5b, 0xab, 0xf2, 0x6d, 0x6b, 0x55, 0x3e, 0x9c, 0x1f, 0x84, 0xf3,
	0x96, 0x0b, 0x36, 0x42, 0x01, 0xae, 0x84, 0x05, 0xe3, 0xee, 0xa7, 0xa3, 0xc9, 0xd2, 0x39, 0xcd,
	0x6a, 0x7a, 0x60, 0x5e, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x94, 0xfd, 0x69, 0x7a, 0x03,
	0x00, 0x00,
}

func (m *ValidatorVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ConsAddr) > 0 {
		i -= len(m.ConsAddr)
		copy(dAtA[i:], m.ConsAddr)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConsAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SupplySnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplySnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplySnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SnapshotAmount.Size()
		i -= size
		if _, err := m.SnapshotAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SnapshotTime != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SnapshotTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PeriodicSnapshot.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.YearStartSnapshot.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.PreviousProposer) > 0 {
		i -= len(m.PreviousProposer)
		copy(dAtA[i:], m.PreviousProposer)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PreviousProposer)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ValidatorVotes) > 0 {
		for iNdEx := len(m.ValidatorVotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorVotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.SnapPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SnapPeriod))
		i--
		dAtA[i] = 0x18
	}
	if len(m.FeesCollected) > 0 {
		for iNdEx := len(m.FeesCollected) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.FeesCollected[iNdEx].Size()
				i -= size
				if _, err := m.FeesCollected[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.FeesTreasury) > 0 {
		for iNdEx := len(m.FeesTreasury) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.FeesTreasury[iNdEx].Size()
				i -= size
				if _, err := m.FeesTreasury[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsAddr)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovGenesis(uint64(m.Height))
	}
	return n
}

func (m *SupplySnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SnapshotTime != 0 {
		n += 1 + sovGenesis(uint64(m.SnapshotTime))
	}
	l = m.SnapshotAmount.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeesTreasury) > 0 {
		for _, e := range m.FeesTreasury {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FeesCollected) > 0 {
		for _, e := range m.FeesCollected {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SnapPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.SnapPeriod))
	}
	if len(m.ValidatorVotes) > 0 {
		for _, e := range m.ValidatorVotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.PreviousProposer)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.YearStartSnapshot.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.PeriodicSnapshot.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *SupplySnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: SupplySnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplySnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotTime", wireType)
			}
			m.SnapshotTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapshotTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SnapshotAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeesTreasury", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.FeesTreasury = append(m.FeesTreasury, v)
			if err := m.FeesTreasury[len(m.FeesTreasury)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeesCollected", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.FeesCollected = append(m.FeesCollected, v)
			if err := m.FeesCollected[len(m.FeesCollected)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapPeriod", wireType)
			}
			m.SnapPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorVotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorVotes = append(m.ValidatorVotes, ValidatorVote{})
			if err := m.ValidatorVotes[len(m.ValidatorVotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousProposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousProposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field YearStartSnapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.YearStartSnapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodicSnapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PeriodicSnapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
