// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type MsgProposalResetWholeValidatorRank struct {
	Proposer    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
	Description string                                        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *MsgProposalResetWholeValidatorRank) Reset()         { *m = MsgProposalResetWholeValidatorRank{} }
func (m *MsgProposalResetWholeValidatorRank) String() string { return proto.CompactTextString(m) }
func (*MsgProposalResetWholeValidatorRank) ProtoMessage()    {}
func (*MsgProposalResetWholeValidatorRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{0}
}
func (m *MsgProposalResetWholeValidatorRank) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalResetWholeValidatorRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalResetWholeValidatorRank.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalResetWholeValidatorRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalResetWholeValidatorRank.Merge(m, src)
}
func (m *MsgProposalResetWholeValidatorRank) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalResetWholeValidatorRank) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalResetWholeValidatorRank.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalResetWholeValidatorRank proto.InternalMessageInfo

func (m *MsgProposalResetWholeValidatorRank) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgProposalResetWholeValidatorRank) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// MsgProposalResetWholeValidatorRankResponse defines the Msg/ProposalResetWholeValidatorRank response type.
type MsgProposalResetWholeValidatorRankResponse struct {
	ProposalID uint64 `protobuf:"varint,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
}

func (m *MsgProposalResetWholeValidatorRankResponse) Reset() {
	*m = MsgProposalResetWholeValidatorRankResponse{}
}
func (m *MsgProposalResetWholeValidatorRankResponse) String() string {
	return proto.CompactTextString(m)
}
func (*MsgProposalResetWholeValidatorRankResponse) ProtoMessage() {}
func (*MsgProposalResetWholeValidatorRankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{1}
}
func (m *MsgProposalResetWholeValidatorRankResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalResetWholeValidatorRankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalResetWholeValidatorRankResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalResetWholeValidatorRankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalResetWholeValidatorRankResponse.Merge(m, src)
}
func (m *MsgProposalResetWholeValidatorRankResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalResetWholeValidatorRankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalResetWholeValidatorRankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalResetWholeValidatorRankResponse proto.InternalMessageInfo

func (m *MsgProposalResetWholeValidatorRankResponse) GetProposalID() uint64 {
	if m != nil {
		return m.ProposalID
	}
	return 0
}

type ProposalResetWholeValidatorRank struct {
	Proposer    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty"`
	Description string                                        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *ProposalResetWholeValidatorRank) Reset()         { *m = ProposalResetWholeValidatorRank{} }
func (m *ProposalResetWholeValidatorRank) String() string { return proto.CompactTextString(m) }
func (*ProposalResetWholeValidatorRank) ProtoMessage()    {}
func (*ProposalResetWholeValidatorRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{2}
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

func init() {
	proto.RegisterType((*MsgProposalResetWholeValidatorRank)(nil), "tsuki.slashing.MsgProposalResetWholeValidatorRank")
	proto.RegisterType((*MsgProposalResetWholeValidatorRankResponse)(nil), "tsuki.slashing.MsgProposalResetWholeValidatorRankResponse")
	proto.RegisterType((*ProposalResetWholeValidatorRank)(nil), "tsuki.slashing.ProposalResetWholeValidatorRank")
}

func init() { proto.RegisterFile("proposal.proto", fileDescriptor_c3ac5ce23bf32d05) }

var fileDescriptor_c3ac5ce23bf32d05 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x51, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0x5d, 0x44, 0x44, 0xe3, 0x9f, 0x43, 0xf1, 0x30, 0x77, 0xc8, 0x4a, 0x4f, 0x63, 0xb0, 0x16,
	0xf1, 0x13, 0x6c, 0xee, 0x22, 0x3a, 0x90, 0x1c, 0x14, 0xbc, 0x48, 0xd6, 0x86, 0x36, 0xb4, 0xeb,
	0x2f, 0xe4, 0x17, 0x41, 0xbf, 0x85, 0x08, 0x7e, 0x27, 0x8f, 0x3b, 0x7a, 0x12, 0x69, 0xbf, 0x85,
	0x27, 0xb1, 0xe9, 0xa4, 0x37, 0xaf, 0x9e, 0x92, 0xbc, 0x17, 0xde, 0x7b, 0xbf, 0xdf, 0xa3, 0x47,
	0xda, 0x80, 0x06, 0x14, 0x45, 0xa8, 0x0d, 0x58, 0xf0, 0x0e, 0x73, 0x65, 0x44, 0x88, 0x85, 0xc0,
	0x4c, 0x95, 0xe9, 0xe0, 0x38, 0x85, 0x14, 0x1a, 0x26, 0xfa, 0xb9, 0xb9, 0x4f, 0x83, 0x93, 0x18,
	0x70, 0x05, 0x78, 0xef, 0x08, 0xf7, 0x70, 0x54, 0xf0, 0x4a, 0x68, 0xb0, 0xc0, 0xf4, 0xba, 0x55,
	0xe5, 0x12, 0xa5, 0xbd, 0xcd, 0xa0, 0x90, 0x37, 0xa2, 0x50, 0x89, 0xb0, 0x60, 0xb8, 0x28, 0x73,
	0x6f, 0x41, 0x77, 0x9d, 0xb1, 0x34, 0x7d, 0xe2, 0x93, 0xd1, 0xc1, 0xec, 0xf4, 0xeb, 0x63, 0x38,
	0x49, 0x95, 0xcd, 0x1e, 0x96, 0x61, 0x0c, 0xab, 0x56, 0xb5, 0x3d, 0x26, 0x98, 0xe4, 0x91, 0x7d,
	0xd2, 0x12, 0xc3, 0x69, 0x1c, 0x4f, 0x93, 0xc4, 0x48, 0x44, 0xfe, 0x2b, 0xe1, 0xf9, 0x74, 0x3f,
	0x91, 0x18, 0x1b, 0xa5, 0xad, 0x82, 0xb2, 0xbf, 0xe5, 0x93, 0xd1, 0x1e, 0xef, 0x42, 0xc1, 0x15,
	0x1d, 0xff, 0x1d, 0x8b, 0x4b, 0xd4, 0x50, 0xa2, 0xf4, 0x18, 0xa5, 0x9b, 0xbd, 0x5c, 0xcc, 0x9b,
	0x80, 0xdb, 0xbc, 0x83, 0x04, 0x2f, 0x84, 0x0e, 0xff, 0xdb, 0x88, 0xb3, 0xf9, 0x5b, 0xc5, 0xc8,
	0xba, 0x62, 0xe4, 0xb3, 0x62, 0xe4, 0xb9, 0x66, 0xbd, 0x75, 0xcd, 0x7a, 0xef, 0x35, 0xeb, 0xdd,
	0x8d, 0x3b, 0xa6, 0x97, 0xca, 0x88, 0x73, 0x30, 0x32, 0x42, 0x99, 0x0b, 0x15, 0x3d, 0x46, 0x9b,
	0xae, 0x9d, 0xf9, 0x72, 0xa7, 0xe9, 0xf1, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x83, 0x7d, 0x5b,
	0x90, 0x19, 0x02, 0x00, 0x00,
}

func (m *MsgProposalResetWholeValidatorRank) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalResetWholeValidatorRank) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalResetWholeValidatorRank) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgProposalResetWholeValidatorRankResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalResetWholeValidatorRankResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalResetWholeValidatorRankResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalID != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
func (m *MsgProposalResetWholeValidatorRank) Size() (n int) {
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

func (m *MsgProposalResetWholeValidatorRankResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovProposal(uint64(m.ProposalID))
	}
	return n
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

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgProposalResetWholeValidatorRank) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgProposalResetWholeValidatorRank: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalResetWholeValidatorRank: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgProposalResetWholeValidatorRankResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgProposalResetWholeValidatorRankResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalResetWholeValidatorRankResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
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
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
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
