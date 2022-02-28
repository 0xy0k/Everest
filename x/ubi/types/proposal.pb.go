// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/ubi/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// UpsertUBIProposal a proposal to create a new or modify existing UBI registry
// record. The proposal should fail if sum of all ((float)amount / period) *
// 31556952 for all UBI records is greater than ubi-hard-cap.
type UpsertUBIProposal struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// distribution-start & distribution-end - defines the exact time period (Unix
	// timestamps) between which tokens will be distributed to the pool, allowing
	// for a precise funds spending.
	DistributionStart uint64 `protobuf:"varint,2,opt,name=distribution_start,json=distributionStart,proto3" json:"distribution_start,omitempty"`
	DistributionEnd   uint64 `protobuf:"varint,3,opt,name=distribution_end,json=distributionEnd,proto3" json:"distribution_end,omitempty"`
	// amount - the amount of KEX tokens to be minted and distributed every period
	// number of seconds into pool
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	// period - time in seconds that must elapse sincedistribution-last
	// for the funds to be distributed automatically into pool
	Period uint64 `protobuf:"varint,5,opt,name=period,proto3" json:"period,omitempty"`
	// pool - spending pool name where exact amount of KEX should be deposited
	Pool string `protobuf:"bytes,6,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (m *UpsertUBIProposal) Reset()         { *m = UpsertUBIProposal{} }
func (m *UpsertUBIProposal) String() string { return proto.CompactTextString(m) }
func (*UpsertUBIProposal) ProtoMessage()    {}
func (*UpsertUBIProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d565c6e8c34ce59, []int{0}
}
func (m *UpsertUBIProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpsertUBIProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpsertUBIProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpsertUBIProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertUBIProposal.Merge(m, src)
}
func (m *UpsertUBIProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpsertUBIProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertUBIProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertUBIProposal proto.InternalMessageInfo

func (m *UpsertUBIProposal) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpsertUBIProposal) GetDistributionStart() uint64 {
	if m != nil {
		return m.DistributionStart
	}
	return 0
}

func (m *UpsertUBIProposal) GetDistributionEnd() uint64 {
	if m != nil {
		return m.DistributionEnd
	}
	return 0
}

func (m *UpsertUBIProposal) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *UpsertUBIProposal) GetPool() string {
	if m != nil {
		return m.Pool
	}
	return ""
}

// RemoveUBIProposal - a proposal to delete UBI record
type RemoveUBIProposal struct {
	UbiName string `protobuf:"bytes,1,opt,name=ubi_name,json=ubiName,proto3" json:"ubi_name,omitempty"`
}

func (m *RemoveUBIProposal) Reset()         { *m = RemoveUBIProposal{} }
func (m *RemoveUBIProposal) String() string { return proto.CompactTextString(m) }
func (*RemoveUBIProposal) ProtoMessage()    {}
func (*RemoveUBIProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d565c6e8c34ce59, []int{1}
}
func (m *RemoveUBIProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveUBIProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveUBIProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveUBIProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveUBIProposal.Merge(m, src)
}
func (m *RemoveUBIProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveUBIProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveUBIProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveUBIProposal proto.InternalMessageInfo

func (m *RemoveUBIProposal) GetUbiName() string {
	if m != nil {
		return m.UbiName
	}
	return ""
}

func init() {
	proto.RegisterType((*UpsertUBIProposal)(nil), "tsuki.gov.UpsertUBIProposal")
	proto.RegisterType((*RemoveUBIProposal)(nil), "tsuki.gov.RemoveUBIProposal")
}

func init() { proto.RegisterFile("tsuki/ubi/proposal.proto", fileDescriptor_8d565c6e8c34ce59) }

var fileDescriptor_8d565c6e8c34ce59 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x3d, 0x8b, 0x14, 0x41,
	0x10, 0xdd, 0x39, 0xd7, 0xb9, 0xb3, 0x41, 0xce, 0x6d, 0x44, 0xe7, 0x2e, 0x98, 0x39, 0x06, 0x94,
	0x33, 0xd8, 0xe9, 0xc0, 0xec, 0x92, 0x85, 0x3d, 0x0c, 0x0e, 0x41, 0x64, 0xe4, 0x10, 0x4c, 0x96,
	0xee, 0x9b, 0x76, 0x6c, 0x76, 0xba, 0xab, 0xe9, 0x8f, 0xc3, 0xfd, 0x17, 0x66, 0xa6, 0xfe, 0x08,
	0x7f, 0xc4, 0x85, 0x8b, 0x91, 0x18, 0x2c, 0xb2, 0x9b, 0x18, 0xfb, 0x0b, 0x64, 0x7a, 0x46, 0x19,
	0x35, 0xea, 0xaa, 0xf7, 0x1e, 0xaf, 0xeb, 0x55, 0xa1, 0x87, 0x4b, 0x61, 0x28, 0xf1, 0x4c, 0x10,
	0x6d, 0x40, 0x83, 0xa5, 0x4d, 0xa1, 0x0d, 0x38, 0xc0, 0x07, 0x2d, 0x51, 0xd4, 0x70, 0x7d, 0x7c,
	0xbf, 0x86, 0x1a, 0x02, 0x48, 0xda, 0xaa, 0xe3, 0x8f, 0xb3, 0x1a, 0xa0, 0x6e, 0x38, 0x09, 0x1d,
	0xf3, 0x6f, 0x89, 0x13, 0x92, 0x5b, 0x47, 0xa5, 0xee, 0x05, 0x47, 0xff, 0x0a, 0xa8, 0x5a, 0xfd,
	0xa6, 0xae, 0xc0, 0x4a, 0xb0, 0x8b, 0xce, 0xb4, 0x6b, 0x7a, 0x0a, 0xff, 0x99, 0xc7, 0x33, 0xd1,
	0x61, 0xf9, 0xc7, 0x3d, 0x34, 0xb9, 0xd4, 0x96, 0x1b, 0x77, 0x39, 0xbf, 0x78, 0xd9, 0x8f, 0x89,
	0x31, 0x1a, 0x2b, 0x2a, 0x79, 0x12, 0x9d, 0x44, 0xa7, 0x77, 0xca, 0x50, 0xe3, 0x29, 0xc2, 0x95,
	0xb0, 0xce, 0x08, 0xe6, 0x9d, 0x00, 0xb5, 0xb0, 0x8e, 0x1a, 0x97, 0xec, 0x9d, 0x44, 0xa7, 0xe3,
	0x72, 0x32, 0x64, 0x5e, 0xb5, 0x04, 0x7e, 0x82, 0xee, 0xfd, 0x25, 0xe7, 0xaa, 0x4a, 0x6e, 0x05,
	0xf1, 0xe1, 0x10, 0x7f, 0xa6, 0x2a, 0xfc, 0x1a, 0xc5, 0x54, 0x82, 0x57, 0x2e, 0x19, 0xb7, 0xff,
	0xcd, 0x67, 0x37, 0x9b, 0x6c, 0xf4, 0x6d, 0x93, 0x3d, 0xae, 0x85, 0x7b, 0xe7, 0x59, 0x71, 0x05,
	0xb2, 0x0f, 0xd2, 0x3f, 0x53, 0x5b, 0x2d, 0x89, 0x5b, 0x69, 0x6e, 0x8b, 0x0b, 0xe5, 0x7e, 0x6e,
	0xb2, 0xbb, 0x2b, 0x2a, 0x9b, 0xb3, 0xbc, 0x73, 0xc9, 0xcb, 0xde, 0x0e, 0x3f, 0x40, 0xb1, 0xe6,
	0x46, 0x40, 0x95, 0xdc, 0x0e, 0x3f, 0xf7, 0x5d, 0x1b, 0x4f, 0x03, 0x34, 0x49, 0xdc, 0xc5, 0x6b,
	0xeb, 0xb3, 0xc3, 0x1f, 0x9f, 0xb2, 0xe8, 0xcb, 0xe7, 0xe9, 0xfe, 0x39, 0x28, 0xc7, 0x95, 0xcb,
	0x67, 0x68, 0x52, 0x72, 0x09, 0xd7, 0x7c, 0xb8, 0x98, 0x23, 0x74, 0xe0, 0x99, 0x58, 0x0c, 0x96,
	0xb3, 0xef, 0x99, 0x78, 0x41, 0x25, 0xff, 0xcf, 0x60, 0x3e, 0xbb, 0xd9, 0xa6, 0xd1, 0x7a, 0x9b,
	0x46, 0xdf, 0xb7, 0x69, 0xf4, 0x61, 0x97, 0x8e, 0xd6, 0xbb, 0x74, 0xf4, 0x75, 0x97, 0x8e, 0xde,
	0x3c, 0x1a, 0x04, 0x7b, 0x2e, 0x0c, 0x3d, 0x07, 0xc3, 0x89, 0xe5, 0x4b, 0x2a, 0xc8, 0xfb, 0x70,
	0x9f, 0x90, 0x8d, 0xc5, 0xe1, 0x44, 0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x63, 0x96, 0xca,
	0x1e, 0x48, 0x02, 0x00, 0x00,
}

func (this *UpsertUBIProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpsertUBIProposal)
	if !ok {
		that2, ok := that.(UpsertUBIProposal)
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
	if this.DistributionStart != that1.DistributionStart {
		return false
	}
	if this.DistributionEnd != that1.DistributionEnd {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	if this.Period != that1.Period {
		return false
	}
	if this.Pool != that1.Pool {
		return false
	}
	return true
}
func (this *RemoveUBIProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RemoveUBIProposal)
	if !ok {
		that2, ok := that.(RemoveUBIProposal)
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
	if this.UbiName != that1.UbiName {
		return false
	}
	return true
}
func (m *UpsertUBIProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpsertUBIProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpsertUBIProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pool) > 0 {
		i -= len(m.Pool)
		copy(dAtA[i:], m.Pool)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Pool)))
		i--
		dAtA[i] = 0x32
	}
	if m.Period != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.DistributionEnd != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.DistributionEnd))
		i--
		dAtA[i] = 0x18
	}
	if m.DistributionStart != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.DistributionStart))
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

func (m *RemoveUBIProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveUBIProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveUBIProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UbiName) > 0 {
		i -= len(m.UbiName)
		copy(dAtA[i:], m.UbiName)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.UbiName)))
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
func (m *UpsertUBIProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.DistributionStart != 0 {
		n += 1 + sovProposal(uint64(m.DistributionStart))
	}
	if m.DistributionEnd != 0 {
		n += 1 + sovProposal(uint64(m.DistributionEnd))
	}
	l = m.Amount.Size()
	n += 1 + l + sovProposal(uint64(l))
	if m.Period != 0 {
		n += 1 + sovProposal(uint64(m.Period))
	}
	l = len(m.Pool)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *RemoveUBIProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UbiName)
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
func (m *UpsertUBIProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpsertUBIProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpsertUBIProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionStart", wireType)
			}
			m.DistributionStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionEnd", wireType)
			}
			m.DistributionEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistributionEnd |= uint64(b&0x7F) << shift
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
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
			m.Pool = string(dAtA[iNdEx:postIndex])
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
func (m *RemoveUBIProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RemoveUBIProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveUBIProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UbiName", wireType)
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
			m.UbiName = string(dAtA[iNdEx:postIndex])
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
