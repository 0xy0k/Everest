// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/upgrade/upgrade.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

type ResourceV03123 struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Git      string `protobuf:"bytes,2,opt,name=git,proto3" json:"git,omitempty"`
	Checkout string `protobuf:"bytes,3,opt,name=checkout,proto3" json:"checkout,omitempty"`
	Checksum string `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *ResourceV03123) Reset()         { *m = ResourceV03123{} }
func (m *ResourceV03123) String() string { return proto.CompactTextString(m) }
func (*ResourceV03123) ProtoMessage()    {}
func (*ResourceV03123) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfa79131e4b330bc, []int{2}
}
func (m *ResourceV03123) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceV03123) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceV03123) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *ResourceV03123) XXX_Size() int {
	return m.Size()
}
func (m *ResourceV03123) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ResourceV03123)(nil), "tsuki.upgrade.ResourceV03123")
}

func init() { proto.RegisterFile("tsuki/upgrade/upgrade.proto", fileDescriptor_cfa79131e4b330bc) }

var fileDescriptor_cfa79131e4b330bc = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xb6, 0x8c, 0xd6, 0xad, 0xb6, 0x61, 0x4d, 0xc3, 0x54, 0x22, 0x2b, 0xbb, 0xac,
	0x08, 0xa9, 0x91, 0x80, 0x0b, 0x3b, 0xae, 0x20, 0x40, 0x5c, 0x50, 0x80, 0x0b, 0x97, 0xc8, 0x4d,
	0xde, 0x52, 0xd3, 0xd8, 0xaf, 0x38, 0x8e, 0xda, 0x7d, 0x0b, 0x3e, 0x02, 0x9f, 0x85, 0x53, 0x8f,
	0x3b, 0x72, 0x42, 0xd0, 0x5e, 0xf8, 0x18, 0x28, 0x8e, 0xd3, 0xed, 0xc8, 0x29, 0x2f, 0xbf, 0xff,
	0x2f, 0xb1, 0xfd, 0x9e, 0xc9, 0x60, 0x2e, 0x34, 0x0f, 0x8a, 0x45, 0xaa, 0x79, 0x02, 0xf5, 0x73,
	0xbc, 0xd0, 0x68, 0x90, 0xf6, 0xcb, 0x6c, 0xec, 0xd8, 0xe0, 0x41, 0x8a, 0x98, 0x66, 0x10, 0xd8,
	0x6c, 0x5a, 0x5c, 0x06, 0x5c, 0x5d, 0x55, 0xe2, 0xe0, 0x28, 0xc5, 0x14, 0x6d, 0x19, 0x94, 0x55,
	0x45, 0x4f, 0x7f, 0xb4, 0xc8, 0xfd, 0xf7, 0x1a, 0x17, 0x98, 0xf3, 0xec, 0x03, 0x5e, 0x9a, 0x25,
	0xd7, 0xf0, 0xa9, 0xfa, 0x19, 0xa5, 0xa4, 0xad, 0xb8, 0x04, 0xe6, 0x0d, 0xbd, 0x51, 0x37, 0xb4,
	0x35, 0x3d, 0x27, 0x5d, 0x0d, 0x39, 0x16, 0x3a, 0x86, 0x9c, 0x35, 0x87, 0xad, 0x51, 0xef, 0xe9,
	0xf1, 0xf8, 0xf6, 0x16, 0xc6, 0xa1, 0x8b, 0x2f, 0xda, 0xeb, 0x5f, 0x27, 0x8d, 0xf0, 0x46, 0xa7,
	0x8f, 0x48, 0xdf, 0x49, 0x91, 0x11, 0x12, 0x58, 0x6b, 0xe8, 0x8d, 0x5a, 0x61, 0xcf, 0xb1, 0x8f,
	0x42, 0x02, 0x1d, 0x92, 0x3e, 0x66, 0x49, 0x14, 0xcf, 0xb8, 0x50, 0x91, 0x48, 0x58, 0xdb, 0x2e,
	0x4d, 0x30, 0x4b, 0x26, 0x25, 0x7a, 0x9b, 0x94, 0x86, 0x82, 0xe5, 0x8d, 0x71, 0xa7, 0x32, 0x14,
	0x2c, 0x6b, 0xe3, 0x09, 0xb9, 0xa7, 0x31, 0xcb, 0xa6, 0x3c, 0x9e, 0x47, 0xf1, 0x0c, 0xe2, 0x79,
	0x5e, 0x48, 0xb6, 0x67, 0xb5, 0xc3, 0x3a, 0x98, 0x38, 0x4e, 0x9f, 0x93, 0x63, 0xc9, 0x57, 0x11,
	0x28, 0x8d, 0x99, 0x04, 0x65, 0xa2, 0xa4, 0xd0, 0xdc, 0x08, 0x54, 0xec, 0xae, 0xdd, 0xdd, 0x91,
	0xe4, 0xab, 0x57, 0x75, 0xf8, 0xd2, 0x65, 0x65, 0x67, 0x24, 0x48, 0x64, 0x9d, 0xaa, 0x33, 0x65,
	0x4d, 0xcf, 0xc8, 0x81, 0x50, 0xb9, 0xe1, 0x06, 0x22, 0x77, 0x22, 0xd6, 0x1d, 0x7a, 0xa3, 0x4e,
	0xb8, 0xef, 0x70, 0xdd, 0xd6, 0x33, 0x72, 0xa0, 0x61, 0x8a, 0x68, 0x22, 0x0d, 0x5f, 0x0b, 0xa1,
	0x21, 0x61, 0xa4, 0x12, 0x2b, 0x1c, 0x3a, 0x5a, 0xf6, 0x2b, 0x9f, 0x8b, 0x45, 0x34, 0xe3, 0x2a,
	0xc9, 0x40, 0xb3, 0x9e, 0xb5, 0x7a, 0x25, 0x7b, 0x53, 0xa1, 0xf3, 0xf6, 0xdf, 0xef, 0x27, 0xde,
	0xe9, 0x0b, 0xf2, 0xb0, 0x9e, 0xe1, 0x84, 0xab, 0x18, 0xfe, 0x67, 0x92, 0xee, 0xd3, 0x2f, 0xa4,
	0x53, 0x0f, 0x8c, 0xee, 0x93, 0xa6, 0x48, 0x9c, 0xd3, 0x14, 0x09, 0x3d, 0x24, 0xad, 0x54, 0x18,
	0xd6, 0xb4, 0xa0, 0x2c, 0xe9, 0x80, 0x74, 0x6c, 0x47, 0xb1, 0x30, 0x76, 0x7a, 0xdd, 0x70, 0xf7,
	0xbe, 0xcb, 0xca, 0x6e, 0xb7, 0x6f, 0x65, 0x79, 0x21, 0xab, 0xb5, 0x2e, 0x5e, 0xaf, 0xff, 0xf8,
	0x8d, 0xf5, 0xc6, 0xf7, 0xae, 0x37, 0xbe, 0xf7, 0x7b, 0xe3, 0x7b, 0xdf, 0xb6, 0x7e, 0xe3, 0x7a,
	0xeb, 0x37, 0x7e, 0x6e, 0xfd, 0xc6, 0xe7, 0xc7, 0xa9, 0x30, 0xb3, 0x62, 0x3a, 0x8e, 0x51, 0x06,
	0xef, 0x84, 0xe6, 0x13, 0xd4, 0x10, 0xe4, 0x30, 0xe7, 0x22, 0x58, 0xed, 0xee, 0xbe, 0xb9, 0x5a,
	0x40, 0x3e, 0xdd, 0xb3, 0x77, 0xf7, 0xd9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x1a, 0x6a,
	0x56, 0x18, 0x03, 0x00, 0x00,
}

func (this *ResourceV03123) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResourceV03123)
	if !ok {
		that2, ok := that.(ResourceV03123)
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
	if this.Git != that1.Git {
		return false
	}
	if this.Checkout != that1.Checkout {
		return false
	}
	if this.Checksum != that1.Checksum {
		return false
	}
	return true
}

func (m *ResourceV03123) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceV03123) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceV03123) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Checkout) > 0 {
		i -= len(m.Checkout)
		copy(dAtA[i:], m.Checkout)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Checkout)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Git) > 0 {
		i -= len(m.Git)
		copy(dAtA[i:], m.Git)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Git)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpgrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpgrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *ResourceV03123) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = len(m.Git)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = len(m.Checkout)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func sovUpgrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpgrade(x uint64) (n int) {
	return sovUpgrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *ResourceV03123) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Git", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Git = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checkout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func skipUpgrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
				return 0, ErrInvalidLengthUpgrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpgrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpgrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpgrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpgrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpgrade = fmt.Errorf("proto: unexpected end of group")
)
