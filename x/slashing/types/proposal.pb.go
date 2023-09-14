// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/slashing/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProposalResetWholeValidatorRank struct {
	Proposer    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
	Description string                                        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *ProposalResetWholeValidatorRank) Reset()         { *m = ProposalResetWholeValidatorRank{} }
func (m *ProposalResetWholeValidatorRank) String() string { return proto.CompactTextString(m) }
func (*ProposalResetWholeValidatorRank) ProtoMessage()    {}
func (*ProposalResetWholeValidatorRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_5754bfe5de951bba, []int{0}
}
func (m *ProposalResetWholeValidatorRank) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalResetWholeValidatorRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalResetWholeValidatorRank.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalResetWholeValidatorRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResetWholeValidatorRank.Merge(m, src)
}
func (m *ProposalResetWholeValidatorRank) XXX_Size() int {
	return m.Size()
}
func (m *ProposalResetWholeValidatorRank) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResetWholeValidatorRank.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResetWholeValidatorRank proto.InternalMessageInfo

func (m *ProposalResetWholeValidatorRank) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *ProposalResetWholeValidatorRank) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ProposalSlashValidator struct {
	Offender         string    `protobuf:"bytes,1,opt,name=offender,proto3" json:"offender,omitempty"`
	StakingPoolId    uint64    `protobuf:"varint,2,opt,name=staking_pool_id,json=stakingPoolId,proto3" json:"staking_pool_id,omitempty"`
	MisbehaviourTime time.Time `protobuf:"bytes,3,opt,name=misbehaviour_time,json=misbehaviourTime,proto3,stdtime" json:"misbehaviour_time"`
	MisbehaviourType string    `protobuf:"bytes,4,opt,name=misbehaviour_type,json=misbehaviourType,proto3" json:"misbehaviour_type,omitempty"`
	JailPercentage   uint64    `protobuf:"varint,5,opt,name=jail_percentage,json=jailPercentage,proto3" json:"jail_percentage,omitempty"`
	Colluders        []string  `protobuf:"bytes,6,rep,name=colluders,proto3" json:"colluders,omitempty"`
	Refutation       string    `protobuf:"bytes,7,opt,name=refutation,proto3" json:"refutation,omitempty"`
}

func (m *ProposalSlashValidator) Reset()         { *m = ProposalSlashValidator{} }
func (m *ProposalSlashValidator) String() string { return proto.CompactTextString(m) }
func (*ProposalSlashValidator) ProtoMessage()    {}
func (*ProposalSlashValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_5754bfe5de951bba, []int{1}
}
func (m *ProposalSlashValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalSlashValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalSlashValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalSlashValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalSlashValidator.Merge(m, src)
}
func (m *ProposalSlashValidator) XXX_Size() int {
	return m.Size()
}
func (m *ProposalSlashValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalSlashValidator.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalSlashValidator proto.InternalMessageInfo

func (m *ProposalSlashValidator) GetOffender() string {
	if m != nil {
		return m.Offender
	}
	return ""
}

func (m *ProposalSlashValidator) GetStakingPoolId() uint64 {
	if m != nil {
		return m.StakingPoolId
	}
	return 0
}

func (m *ProposalSlashValidator) GetMisbehaviourTime() time.Time {
	if m != nil {
		return m.MisbehaviourTime
	}
	return time.Time{}
}

func (m *ProposalSlashValidator) GetMisbehaviourType() string {
	if m != nil {
		return m.MisbehaviourType
	}
	return ""
}

func (m *ProposalSlashValidator) GetJailPercentage() uint64 {
	if m != nil {
		return m.JailPercentage
	}
	return 0
}

func (m *ProposalSlashValidator) GetColluders() []string {
	if m != nil {
		return m.Colluders
	}
	return nil
}

func (m *ProposalSlashValidator) GetRefutation() string {
	if m != nil {
		return m.Refutation
	}
	return ""
}

func init() {
	proto.RegisterType((*ProposalResetWholeValidatorRank)(nil), "tsuki.slashing.ProposalResetWholeValidatorRank")
	proto.RegisterType((*ProposalSlashValidator)(nil), "tsuki.slashing.ProposalSlashValidator")
}

func init() {
	proto.RegisterFile("tsuki/slashing/v1beta1/proposal.proto", fileDescriptor_5754bfe5de951bba)
}

var fileDescriptor_5754bfe5de951bba = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xdb, 0x52, 0x92, 0x2d, 0xa5, 0xb0, 0x42, 0xc8, 0x44, 0xc8, 0xb6, 0x2a, 0x04, 0x11,
	0xa8, 0xb6, 0x02, 0xbf, 0xa0, 0x81, 0x0b, 0x42, 0x48, 0xc1, 0x20, 0x90, 0xb8, 0x58, 0x6b, 0x7b,
	0xe2, 0x2c, 0x5e, 0x7b, 0x56, 0xbb, 0xeb, 0x8a, 0xfc, 0x8b, 0x8a, 0xbf, 0xc4, 0xa5, 0xc7, 0x1e,
	0x39, 0x15, 0x94, 0xfc, 0x0b, 0x4e, 0xc8, 0x1f, 0x49, 0x43, 0x4f, 0xf6, 0xbc, 0x79, 0x3b, 0xef,
	0xed, 0x9b, 0x25, 0x4f, 0x72, 0xae, 0x58, 0xa0, 0x05, 0xd3, 0x73, 0x5e, 0x66, 0xc1, 0xd9, 0x38,
	0x06, 0xc3, 0xc6, 0x81, 0x54, 0x28, 0x51, 0x33, 0xe1, 0x4b, 0x85, 0x06, 0xe9, 0x61, 0xcd, 0xf2,
	0xd7, 0xac, 0xe1, 0x83, 0x0c, 0x33, 0x6c, 0x3a, 0x41, 0xfd, 0xd7, 0x92, 0x86, 0x8f, 0x12, 0xd4,
	0x05, 0xea, 0xa8, 0x6d, 0xb4, 0x45, 0xd7, 0x72, 0x33, 0xc4, 0x4c, 0x40, 0xd0, 0x54, 0x71, 0x35,
	0x0b, 0x0c, 0x2f, 0x40, 0x1b, 0x56, 0xc8, 0x96, 0x70, 0xfc, 0xc3, 0x22, 0xee, 0xb4, 0xd3, 0x0c,
	0x41, 0x83, 0xf9, 0x32, 0x47, 0x01, 0x9f, 0x99, 0xe0, 0x29, 0x33, 0xa8, 0x42, 0x56, 0xe6, 0xf4,
	0x3d, 0xe9, 0xb7, 0xb6, 0x40, 0xd9, 0x96, 0x67, 0x8d, 0xee, 0x4c, 0xc6, 0x7f, 0xaf, 0xdc, 0x93,
	0x8c, 0x9b, 0x79, 0x15, 0xfb, 0x09, 0x16, 0x9d, 0x66, 0xf7, 0x39, 0xd1, 0x69, 0x1e, 0x98, 0x85,
	0x04, 0xed, 0x9f, 0x26, 0xc9, 0x69, 0x9a, 0x2a, 0xd0, 0x3a, 0xdc, 0x8c, 0xa0, 0x1e, 0x39, 0x48,
	0x41, 0x27, 0x8a, 0x4b, 0xc3, 0xb1, 0xb4, 0x77, 0x3c, 0x6b, 0x34, 0x08, 0xb7, 0xa1, 0xe3, 0x9f,
	0x3b, 0xe4, 0xe1, 0xda, 0xd4, 0xc7, 0xfa, 0xee, 0x1b, 0x3f, 0x74, 0x48, 0xfa, 0x38, 0x9b, 0x41,
	0x99, 0x76, 0x5e, 0x06, 0xe1, 0xa6, 0xa6, 0x4f, 0xc9, 0x91, 0x36, 0x2c, 0xe7, 0x65, 0x16, 0x49,
	0x44, 0x11, 0xf1, 0xb4, 0x19, 0xbe, 0x17, 0x1e, 0x76, 0xf0, 0x14, 0x51, 0xbc, 0x4d, 0xe9, 0x07,
	0x72, 0xbf, 0xe0, 0x3a, 0x86, 0x39, 0x3b, 0xe3, 0x58, 0xa9, 0xa8, 0xce, 0xc4, 0xde, 0xf5, 0xac,
	0xd1, 0xc1, 0xcb, 0xa1, 0xdf, 0x06, 0xe6, 0xaf, 0x03, 0xf3, 0x3f, 0xad, 0x03, 0x9b, 0xf4, 0x2f,
	0xae, 0xdc, 0xde, 0xf9, 0x6f, 0xd7, 0x0a, 0xef, 0x6d, 0x1f, 0xaf, 0x09, 0xf4, 0xc5, 0xcd, 0x91,
	0x0b, 0x09, 0xf6, 0x5e, 0xe3, 0xef, 0x7f, 0xf2, 0x42, 0x02, 0x7d, 0x46, 0x8e, 0xbe, 0x31, 0x2e,
	0x22, 0x09, 0x2a, 0x81, 0xd2, 0xb0, 0x0c, 0xec, 0x5b, 0x8d, 0xcf, 0xbb, 0x35, 0x3c, 0xdd, 0xa0,
	0xf4, 0x31, 0x19, 0x24, 0x28, 0x44, 0x95, 0x82, 0xd2, 0xf6, 0xbe, 0xb7, 0x3b, 0x1a, 0x84, 0xd7,
	0x00, 0x75, 0x08, 0x51, 0x30, 0xab, 0x0c, 0x6b, 0x62, 0xbc, 0xdd, 0x88, 0x6d, 0x21, 0x93, 0x37,
	0x17, 0x4b, 0xc7, 0xba, 0x5c, 0x3a, 0xd6, 0x9f, 0xa5, 0x63, 0x9d, 0xaf, 0x9c, 0xde, 0xe5, 0xca,
	0xe9, 0xfd, 0x5a, 0x39, 0xbd, 0xaf, 0xcf, 0xb7, 0x56, 0xf7, 0x8e, 0x2b, 0xf6, 0x1a, 0x15, 0x04,
	0x1a, 0x72, 0xc6, 0x83, 0xef, 0xd7, 0x4f, 0xb2, 0x59, 0x61, 0xbc, 0xdf, 0x24, 0xf1, 0xea, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xb0, 0xbc, 0xcb, 0xb0, 0x02, 0x00, 0x00,
}

func (m *ProposalResetWholeValidatorRank) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalResetWholeValidatorRank) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalResetWholeValidatorRank) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalSlashValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalSlashValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalSlashValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Refutation) > 0 {
		i -= len(m.Refutation)
		copy(dAtA[i:], m.Refutation)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Refutation)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Colluders) > 0 {
		for iNdEx := len(m.Colluders) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Colluders[iNdEx])
			copy(dAtA[i:], m.Colluders[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Colluders[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.JailPercentage != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.JailPercentage))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MisbehaviourType) > 0 {
		i -= len(m.MisbehaviourType)
		copy(dAtA[i:], m.MisbehaviourType)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.MisbehaviourType)))
		i--
		dAtA[i] = 0x22
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.MisbehaviourTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.MisbehaviourTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProposal(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if m.StakingPoolId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.StakingPoolId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Offender) > 0 {
		i -= len(m.Offender)
		copy(dAtA[i:], m.Offender)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Offender)))
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
func (m *ProposalResetWholeValidatorRank) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *ProposalSlashValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Offender)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.StakingPoolId != 0 {
		n += 1 + sovProposal(uint64(m.StakingPoolId))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.MisbehaviourTime)
	n += 1 + l + sovProposal(uint64(l))
	l = len(m.MisbehaviourType)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.JailPercentage != 0 {
		n += 1 + sovProposal(uint64(m.JailPercentage))
	}
	if len(m.Colluders) > 0 {
		for _, s := range m.Colluders {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	l = len(m.Refutation)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalResetWholeValidatorRank) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProposalResetWholeValidatorRank: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalResetWholeValidatorRank: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
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
func (m *ProposalSlashValidator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProposalSlashValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalSlashValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offender", wireType)
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
			m.Offender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingPoolId", wireType)
			}
			m.StakingPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakingPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MisbehaviourTime", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.MisbehaviourTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MisbehaviourType", wireType)
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
			m.MisbehaviourType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailPercentage", wireType)
			}
			m.JailPercentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailPercentage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Colluders", wireType)
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
			m.Colluders = append(m.Colluders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refutation", wireType)
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
			m.Refutation = string(dAtA[iNdEx:postIndex])
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
