// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tsuki/staking/tx.proto

package types

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgClaimValidator struct {
	Moniker string                                        `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	ValKey  github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,2,opt,name=val_key,json=valKey,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"val_key,omitempty" yaml:"val_key"`
	PubKey  *types.Any                                    `protobuf:"bytes,3,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty" yaml:"pub_key"`
}

func (m *MsgClaimValidator) Reset()         { *m = MsgClaimValidator{} }
func (m *MsgClaimValidator) String() string { return proto.CompactTextString(m) }
func (*MsgClaimValidator) ProtoMessage()    {}
func (*MsgClaimValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3ae268f0016acc, []int{0}
}
func (m *MsgClaimValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimValidator.Merge(m, src)
}
func (m *MsgClaimValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimValidator proto.InternalMessageInfo

func (m *MsgClaimValidator) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *MsgClaimValidator) GetValKey() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValKey
	}
	return nil
}

func (m *MsgClaimValidator) GetPubKey() *types.Any {
	if m != nil {
		return m.PubKey
	}
	return nil
}

// MsgClaimValidatorResponse defines the Msg/ClaimValidator response type.
type MsgClaimValidatorResponse struct {
}

func (m *MsgClaimValidatorResponse) Reset()         { *m = MsgClaimValidatorResponse{} }
func (m *MsgClaimValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimValidatorResponse) ProtoMessage()    {}
func (*MsgClaimValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d3ae268f0016acc, []int{1}
}
func (m *MsgClaimValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimValidatorResponse.Merge(m, src)
}
func (m *MsgClaimValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimValidatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgClaimValidator)(nil), "tsuki.staking.MsgClaimValidator")
	proto.RegisterType((*MsgClaimValidatorResponse)(nil), "tsuki.staking.MsgClaimValidatorResponse")
}

func init() { proto.RegisterFile("tsuki/staking/tx.proto", fileDescriptor_7d3ae268f0016acc) }

var fileDescriptor_7d3ae268f0016acc = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6a, 0x1b, 0x31,
	0x10, 0xb6, 0x9a, 0x62, 0x53, 0x35, 0x04, 0xba, 0xb8, 0xe0, 0x38, 0xb0, 0x6b, 0xf6, 0x52, 0xb7,
	0x60, 0x09, 0xd2, 0x5b, 0x6e, 0xb1, 0x8f, 0x26, 0x50, 0x7c, 0x08, 0x34, 0x14, 0x82, 0xd6, 0x56,
	0x55, 0xb1, 0xda, 0x1d, 0x21, 0x69, 0x43, 0xf4, 0x16, 0x7d, 0x84, 0x3e, 0x44, 0x1f, 0xa2, 0xf4,
	0x94, 0x63, 0x4f, 0xa1, 0xd8, 0x97, 0x1e, 0x7a, 0xea, 0xb1, 0xa7, 0xe2, 0x5d, 0x2d, 0xf9, 0xf1,
	0xa1, 0xa7, 0xd1, 0xcc, 0xf7, 0xcd, 0x8c, 0xbe, 0x4f, 0xc2, 0x2f, 0x73, 0x69, 0x18, 0xb5, 0x8e,
	0xe5, 0xb2, 0x14, 0xd4, 0x5d, 0x13, 0x6d, 0xc0, 0x41, 0xb4, 0xbf, 0x2d, 0x93, 0x50, 0x1e, 0xf6,
	0x05, 0x08, 0xa8, 0x01, 0xba, 0x3d, 0x35, 0x9c, 0xe1, 0xa1, 0x00, 0x10, 0x8a, 0xd3, 0x3a, 0xcb,
	0xaa, 0x8f, 0x94, 0x95, 0xbe, 0x85, 0x96, 0x60, 0x0b, 0xb0, 0x97, 0x4d, 0x4f, 0x93, 0x04, 0x28,
	0x79, 0xdc, 0xe5, 0x64, 0xc1, 0xad, 0x63, 0x85, 0x0e, 0x84, 0xa3, 0x07, 0x37, 0xd2, 0x06, 0x34,
	0x58, 0xa6, 0x02, 0x38, 0x7c, 0x00, 0x86, 0xd8, 0x60, 0xe9, 0x6f, 0x84, 0x5f, 0x9c, 0x59, 0x31,
	0x53, 0x4c, 0x16, 0xe7, 0x4c, 0xc9, 0x15, 0x73, 0x60, 0xa2, 0x01, 0xee, 0x15, 0x50, 0xca, 0x9c,
	0x9b, 0x01, 0x1a, 0xa1, 0xf1, 0xb3, 0x45, 0x9b, 0x46, 0x1f, 0x70, 0xef, 0x8a, 0xa9, 0xcb, 0x9c,
	0xfb, 0xc1, 0x93, 0x11, 0x1a, 0xef, 0x4f, 0x67, 0x7f, 0x6e, 0x93, 0x03, 0xcf, 0x0a, 0x75, 0x92,
	0x06, 0x20, 0xfd, 0x7b, 0x9b, 0x4c, 0x84, 0x74, 0x9f, 0xaa, 0x8c, 0x2c, 0xa1, 0x08, 0x4a, 0x42,
	0x98, 0xd8, 0x55, 0x4e, 0x9d, 0xd7, 0xdc, 0x92, 0x73, 0xa6, 0x4e, 0x57, 0x2b, 0xc3, 0xad, 0x5d,
	0x74, 0xaf, 0x98, 0x9a, 0x73, 0x1f, 0xbd, 0xc7, 0x3d, 0x5d, 0x65, 0xf5, 0xf4, 0xbd, 0x11, 0x1a,
	0x3f, 0x3f, 0xee, 0x93, 0x46, 0x39, 0x69, 0x95, 0x93, 0xd3, 0xd2, 0x4f, 0xdf, 0xdc, 0xed, 0x0c,
	0xf4, 0xf4, 0xfb, 0xd7, 0x49, 0x3f, 0x58, 0xb6, 0x34, 0x5e, 0x3b, 0x20, 0xef, 0xaa, 0x6c, 0xce,
	0xfd, 0xa2, 0xab, 0xeb, 0x78, 0xf2, 0xf4, 0xd7, 0x97, 0x04, 0xa5, 0x47, 0xf8, 0x70, 0x47, 0xed,
	0x82, 0x5b, 0x0d, 0xa5, 0xe5, 0xc7, 0x0c, 0xef, 0x9d, 0x59, 0x11, 0x5d, 0xe0, 0x83, 0x47, 0x76,
	0x24, 0xe4, 0xfe, 0xcb, 0x92, 0x9d, 0x09, 0xc3, 0x57, 0xff, 0x21, 0xb4, 0x2b, 0xa6, 0xb3, 0x6f,
	0xeb, 0x18, 0xdd, 0xac, 0x63, 0xf4, 0x73, 0x1d, 0xa3, 0xcf, 0x9b, 0xb8, 0x73, 0xb3, 0x89, 0x3b,
	0x3f, 0x36, 0x71, 0xe7, 0xe2, 0xf5, 0x3d, 0xc7, 0xe6, 0xd2, 0xb0, 0x19, 0x18, 0x4e, 0x2d, 0xcf,
	0x99, 0xa4, 0xd7, 0x77, 0x5f, 0x6d, 0x6b, 0x5c, 0xd6, 0xad, 0xcd, 0x78, 0xfb, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0xcb, 0x01, 0x30, 0x87, 0x02, 0x00, 0x00,
}

func (this *MsgClaimValidator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimValidator)
	if !ok {
		that2, ok := that.(MsgClaimValidator)
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
	if this.Moniker != that1.Moniker {
		return false
	}
	if !bytes.Equal(this.ValKey, that1.ValKey) {
		return false
	}
	if !this.PubKey.Equal(that1.PubKey) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ClaimValidator defines a method for claiming a new validator.
	ClaimValidator(ctx context.Context, in *MsgClaimValidator, opts ...grpc.CallOption) (*MsgClaimValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ClaimValidator(ctx context.Context, in *MsgClaimValidator, opts ...grpc.CallOption) (*MsgClaimValidatorResponse, error) {
	out := new(MsgClaimValidatorResponse)
	err := c.cc.Invoke(ctx, "/tsuki.staking.Msg/ClaimValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ClaimValidator defines a method for claiming a new validator.
	ClaimValidator(context.Context, *MsgClaimValidator) (*MsgClaimValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ClaimValidator(ctx context.Context, req *MsgClaimValidator) (*MsgClaimValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ClaimValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tsuki.staking.Msg/ClaimValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimValidator(ctx, req.(*MsgClaimValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tsuki.staking.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimValidator",
			Handler:    _Msg_ClaimValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tsuki/staking/tx.proto",
}

func (m *MsgClaimValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValKey) > 0 {
		i -= len(m.ValKey)
		copy(dAtA[i:], m.ValKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgClaimValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgClaimValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValKey = append(m.ValKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ValKey == nil {
				m.ValKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
