// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/ubi/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
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
	Amount uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
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

func (m *UpsertUBIProposal) GetAmount() uint64 {
	if m != nil {
		return m.Amount
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
	proto.RegisterType((*UpsertUBIProposal)(nil), "tsuki.ubi.UpsertUBIProposal")
	proto.RegisterType((*RemoveUBIProposal)(nil), "tsuki.ubi.RemoveUBIProposal")
}

func init() { proto.RegisterFile("tsuki/ubi/proposal.proto", fileDescriptor_8d565c6e8c34ce59) }

var fileDescriptor_8d565c6e8c34ce59 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x6a, 0x5a, 0x41,
	0x14, 0xc6, 0x9d, 0xd6, 0xaa, 0x9d, 0x8d, 0x75, 0x28, 0xed, 0xd5, 0xc5, 0x55, 0x84, 0x82, 0x5d,
	0xe8, 0x2c, 0xba, 0xeb, 0x46, 0x50, 0xba, 0x28, 0x85, 0x52, 0x2c, 0x6e, 0xba, 0x91, 0x19, 0xef,
	0xf4, 0x66, 0xd0, 0x99, 0x33, 0xcc, 0x9f, 0x10, 0xdf, 0x22, 0x8f, 0x90, 0x87, 0xc8, 0x43, 0x84,
	0xac, 0x5c, 0x85, 0x2c, 0x83, 0x6e, 0xf2, 0x18, 0xe1, 0xfe, 0x49, 0xb8, 0x49, 0x76, 0xe7, 0xfb,
	0x7e, 0x1f, 0x87, 0xef, 0x1c, 0xfc, 0x79, 0x23, 0x2d, 0xa3, 0x81, 0x4b, 0x6a, 0x2c, 0x18, 0x70,
	0x6c, 0x3b, 0x31, 0x16, 0x3c, 0x90, 0x56, 0x06, 0x26, 0x81, 0xcb, 0xde, 0xc7, 0x14, 0x52, 0xc8,
	0x4d, 0x9a, 0x4d, 0x05, 0xef, 0xf5, 0x53, 0x80, 0x74, 0x2b, 0x68, 0xae, 0x78, 0xf8, 0x4f, 0xbd,
	0x54, 0xc2, 0x79, 0xa6, 0x4c, 0x19, 0xe8, 0xbe, 0x0c, 0x30, 0xbd, 0x7b, 0x44, 0x6b, 0x70, 0x0a,
	0xdc, 0xaa, 0x58, 0x5a, 0x88, 0x12, 0x91, 0xa7, 0x3e, 0x81, 0xcb, 0xc2, 0x1b, 0xde, 0x20, 0xdc,
	0x59, 0x1a, 0x27, 0xac, 0x5f, 0xce, 0x7e, 0xfe, 0x29, 0x6b, 0x12, 0x82, 0xeb, 0x9a, 0x29, 0x11,
	0xa1, 0x01, 0x1a, 0xbd, 0x5f, 0xe4, 0x33, 0x19, 0x63, 0x92, 0x48, 0xe7, 0xad, 0xe4, 0xc1, 0x4b,
	0xd0, 0x2b, 0xe7, 0x99, 0xf5, 0xd1, 0x9b, 0x01, 0x1a, 0xd5, 0x17, 0x9d, 0x2a, 0xf9, 0x9b, 0x01,
	0xf2, 0x15, 0x7f, 0x78, 0x16, 0x17, 0x3a, 0x89, 0xde, 0xe6, 0xe1, 0x76, 0xd5, 0xff, 0xa1, 0x13,
	0xf2, 0x09, 0x37, 0x98, 0x82, 0xa0, 0x7d, 0x54, 0xcf, 0x03, 0xa5, 0xca, 0x7c, 0x23, 0xac, 0x84,
	0x24, 0x7a, 0x57, 0xf8, 0x85, 0xca, 0xda, 0x19, 0x80, 0x6d, 0xd4, 0x28, 0xda, 0x65, 0xf3, 0xf7,
	0xf6, 0xfd, 0x45, 0x1f, 0x5d, 0x5f, 0x8e, 0x9b, 0x73, 0xd0, 0x5e, 0x68, 0x3f, 0x9c, 0xe2, 0xce,
	0x42, 0x28, 0x38, 0x15, 0xd5, 0xbb, 0xba, 0xb8, 0x15, 0xb8, 0x5c, 0x55, 0x6e, 0x6b, 0x06, 0x2e,
	0x7f, 0x33, 0x25, 0x5e, 0x2d, 0x98, 0x4d, 0xaf, 0x0e, 0x31, 0xda, 0x1f, 0x62, 0x74, 0x77, 0x88,
	0xd1, 0xf9, 0x31, 0xae, 0xed, 0x8f, 0x71, 0xed, 0xf6, 0x18, 0xd7, 0xfe, 0x7d, 0x49, 0xa5, 0x3f,
	0x09, 0x7c, 0xb2, 0x06, 0x45, 0x7f, 0x49, 0xcb, 0xe6, 0x60, 0x05, 0x75, 0x62, 0xc3, 0x24, 0x3d,
	0xcb, 0xdf, 0xeb, 0x77, 0x46, 0x38, 0xde, 0xc8, 0x3f, 0xfc, 0xed, 0x21, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x7b, 0x22, 0x10, 0x07, 0x02, 0x00, 0x00,
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
	if this.Amount != that1.Amount {
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
	if m.Amount != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
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
	if m.Amount != 0 {
		n += 1 + sovProposal(uint64(m.Amount))
	}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
