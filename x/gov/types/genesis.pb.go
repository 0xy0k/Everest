// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: genesis.proto

package types

import (
	fmt "fmt"
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

type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	// permissions is the roles that are active from genesis.
	Permissions map[uint64]*Permissions `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// NetworkActors are the actors that are saved from genesis.
	NetworkActors               []*NetworkActor               `protobuf:"bytes,3,rep,name=network_actors,json=networkActors,proto3" json:"network_actors,omitempty"`
	NetworkProperties           *NetworkProperties            `protobuf:"bytes,4,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	ExecutionFees               []*ExecutionFee               `protobuf:"bytes,5,rep,name=execution_fees,json=executionFees,proto3" json:"execution_fees,omitempty"`
	PoorNetworkMessages         *AllowedMessages              `protobuf:"bytes,6,opt,name=poor_network_messages,json=poorNetworkMessages,proto3" json:"poor_network_messages,omitempty"`
	Proposals                   []Proposal                    `protobuf:"bytes,7,rep,name=proposals,proto3" json:"proposals"`
	Votes                       []Vote                        `protobuf:"bytes,8,rep,name=votes,proto3" json:"votes"`
	DataRegistry                map[string]*DataRegistryEntry `protobuf:"bytes,9,rep,name=data_registry,json=dataRegistry,proto3" json:"data_registry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastIdentityRecordId        uint64                        `protobuf:"varint,10,opt,name=last_identity_record_id,json=lastIdentityRecordId,proto3" json:"last_identity_record_id,omitempty"`
	LastIdRecordVerifyRequestId uint64                        `protobuf:"varint,11,opt,name=last_id_record_verify_request_id,json=lastIdRecordVerifyRequestId,proto3" json:"last_id_record_verify_request_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_14205810582f3203, []int{0}
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

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetPermissions() map[uint64]*Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *GenesisState) GetNetworkActors() []*NetworkActor {
	if m != nil {
		return m.NetworkActors
	}
	return nil
}

func (m *GenesisState) GetNetworkProperties() *NetworkProperties {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *GenesisState) GetExecutionFees() []*ExecutionFee {
	if m != nil {
		return m.ExecutionFees
	}
	return nil
}

func (m *GenesisState) GetPoorNetworkMessages() *AllowedMessages {
	if m != nil {
		return m.PoorNetworkMessages
	}
	return nil
}

func (m *GenesisState) GetProposals() []Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetVotes() []Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetDataRegistry() map[string]*DataRegistryEntry {
	if m != nil {
		return m.DataRegistry
	}
	return nil
}

func (m *GenesisState) GetLastIdentityRecordId() uint64 {
	if m != nil {
		return m.LastIdentityRecordId
	}
	return 0
}

func (m *GenesisState) GetLastIdRecordVerifyRequestId() uint64 {
	if m != nil {
		return m.LastIdRecordVerifyRequestId
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tsuki.gov.GenesisState")
	proto.RegisterMapType((map[string]*DataRegistryEntry)(nil), "tsuki.gov.GenesisState.DataRegistryEntry")
	proto.RegisterMapType((map[uint64]*Permissions)(nil), "tsuki.gov.GenesisState.PermissionsEntry")
}

func init() { proto.RegisterFile("genesis.proto", fileDescriptor_14205810582f3203) }

var fileDescriptor_14205810582f3203 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xce, 0xf6, 0xbf, 0x93, 0xa6, 0xb4, 0xd3, 0x1f, 0xc7, 0x14, 0xd6, 0x20, 0x88, 0x41, 0x61,
	0x57, 0x2b, 0x8a, 0x08, 0x22, 0xad, 0x56, 0x89, 0x52, 0x29, 0x2b, 0xf6, 0x42, 0x84, 0x65, 0x9a,
	0x3d, 0x5d, 0x87, 0x6c, 0x77, 0xd6, 0x99, 0x49, 0xda, 0xbc, 0x82, 0x57, 0x3e, 0x56, 0x2f, 0x7b,
	0xe9, 0x95, 0x48, 0xf2, 0x22, 0xb2, 0xb3, 0x33, 0xee, 0x26, 0xc1, 0xbb, 0x3d, 0xe7, 0xfb, 0xe1,
	0x9b, 0x73, 0xce, 0xa2, 0x46, 0x0c, 0x29, 0x48, 0x26, 0xbd, 0x4c, 0x70, 0xc5, 0xf1, 0x4a, 0x8f,
	0x09, 0xea, 0xc5, 0x7c, 0xd0, 0xdc, 0x8e, 0x79, 0xcc, 0x75, 0xd3, 0xcf, 0xbf, 0x0a, 0xbc, 0x59,
	0xa7, 0x5d, 0xc5, 0x85, 0x29, 0x90, 0xe0, 0x09, 0x98, 0xef, 0xf5, 0x4c, 0xf0, 0x8c, 0x4b, 0x9a,
	0x98, 0x7a, 0x2b, 0xa2, 0x8a, 0x86, 0x02, 0x62, 0x26, 0x95, 0x18, 0xda, 0x26, 0x5c, 0x41, 0xb7,
	0xaf, 0x18, 0x4f, 0xc3, 0x73, 0xb0, 0x4a, 0x92, 0x82, 0xba, 0xe4, 0xa2, 0x17, 0xe6, 0x0e, 0x20,
	0x14, 0x03, 0x13, 0xa6, 0xb9, 0x4b, 0x93, 0x84, 0x5f, 0x42, 0x14, 0x5e, 0x80, 0x94, 0x34, 0xb6,
	0xfd, 0xbb, 0x3f, 0x96, 0xd1, 0xda, 0xbb, 0x22, 0xf6, 0x27, 0x45, 0x15, 0xe0, 0x47, 0x68, 0x5b,
	0x2a, 0x2a, 0x14, 0x4b, 0xe3, 0xd0, 0xe6, 0x08, 0x59, 0x44, 0x9c, 0x96, 0xd3, 0x5e, 0x08, 0xb0,
	0xc5, 0x4e, 0x0c, 0xd4, 0x89, 0x70, 0x07, 0xd5, 0x33, 0x10, 0x17, 0x4c, 0x4a, 0xc6, 0x53, 0x49,
	0xe6, 0x5a, 0xf3, 0xed, 0xfa, 0xfe, 0x7d, 0xcf, 0xbe, 0xde, 0xab, 0xda, 0x7b, 0x27, 0x25, 0xf3,
	0x28, 0x55, 0x62, 0x18, 0x54, 0xb5, 0xf8, 0x25, 0x5a, 0xb7, 0x2f, 0xd0, 0xc3, 0x91, 0x64, 0x5e,
	0xbb, 0xed, 0x96, 0x6e, 0x1f, 0x0b, 0xfc, 0x20, 0x87, 0x83, 0x46, 0x5a, 0xa9, 0x24, 0x7e, 0x8f,
	0xf0, 0xec, 0x00, 0xc8, 0x42, 0xcb, 0x69, 0xd7, 0xf7, 0xf7, 0x66, 0x2c, 0x4e, 0xfe, 0x51, 0x82,
	0xcd, 0x74, 0xba, 0x95, 0x47, 0x99, 0x98, 0xb0, 0x24, 0x8b, 0xd3, 0x51, 0x8e, 0x2c, 0xfe, 0x16,
	0x20, 0x68, 0x40, 0xa5, 0x92, 0xf8, 0x18, 0xed, 0x64, 0x9c, 0x8b, 0xd0, 0xe6, 0xb1, 0x63, 0x27,
	0x4b, 0x3a, 0xcd, 0xed, 0xd2, 0xe5, 0xa0, 0x58, 0xcc, 0xb1, 0x21, 0x04, 0x5b, 0xb9, 0xce, 0x44,
	0xb4, 0x4d, 0xfc, 0x0c, 0xad, 0xda, 0x65, 0x48, 0xb2, 0xac, 0x83, 0xe0, 0xd2, 0xc2, 0x2e, 0xe3,
	0x70, 0xe1, 0xfa, 0xf7, 0x9d, 0x5a, 0x50, 0x52, 0xf1, 0x03, 0xb4, 0x38, 0xe0, 0x0a, 0x24, 0x59,
	0xd1, 0x9a, 0xf5, 0x52, 0x73, 0xca, 0x15, 0x18, 0x7e, 0x41, 0xc1, 0xc7, 0xa8, 0x31, 0x71, 0x68,
	0x64, 0x55, 0x6b, 0xda, 0xff, 0xd9, 0xe4, 0x1b, 0xaa, 0x68, 0x60, 0xa8, 0xc5, 0x2a, 0xd7, 0xa2,
	0x4a, 0x0b, 0x3f, 0x45, 0xb7, 0x12, 0x2a, 0x55, 0xc8, 0x22, 0x48, 0x15, 0x53, 0xc3, 0x50, 0x40,
	0x97, 0x8b, 0x28, 0xbf, 0x25, 0xa4, 0x6f, 0x69, 0x3b, 0x87, 0x3b, 0x06, 0x0d, 0x34, 0xd8, 0x89,
	0xf0, 0x11, 0x6a, 0x19, 0x99, 0x15, 0x0c, 0x40, 0xb0, 0xf3, 0x5c, 0xfe, 0xbd, 0x0f, 0x1a, 0x20,
	0x75, 0xad, 0xdf, 0x2b, 0xf4, 0x85, 0xf2, 0x54, 0x93, 0x82, 0x82, 0xd3, 0x89, 0x9a, 0x9f, 0xd1,
	0xc6, 0xf4, 0xa9, 0xe1, 0x0d, 0x34, 0xdf, 0x83, 0xa1, 0xb9, 0xe4, 0xfc, 0x13, 0x3f, 0x44, 0x8b,
	0x03, 0x9a, 0xf4, 0x81, 0xcc, 0xe9, 0xad, 0xec, 0x54, 0x46, 0x5a, 0x8a, 0x83, 0x82, 0xf3, 0x62,
	0xee, 0xb9, 0xd3, 0xfc, 0x8a, 0x36, 0x67, 0xde, 0x5d, 0xf5, 0x5d, 0x2d, 0x7c, 0x1f, 0x4f, 0xfa,
	0x56, 0x6e, 0x6f, 0x76, 0x6a, 0xa5, 0xfb, 0xe1, 0xab, 0xeb, 0x91, 0xeb, 0xdc, 0x8c, 0x5c, 0xe7,
	0xcf, 0xc8, 0x75, 0x7e, 0x8e, 0xdd, 0xda, 0xcd, 0xd8, 0xad, 0xfd, 0x1a, 0xbb, 0xb5, 0x2f, 0xf7,
	0x62, 0xa6, 0xbe, 0xf5, 0xcf, 0xbc, 0x2e, 0xbf, 0xf0, 0x3f, 0x30, 0x41, 0x5f, 0x73, 0x01, 0xbe,
	0x84, 0x1e, 0x65, 0xfe, 0x95, 0x1f, 0xf3, 0x81, 0xaf, 0x86, 0x19, 0xc8, 0xb3, 0x25, 0xfd, 0x53,
	0x3f, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x63, 0xc4, 0xa9, 0x8a, 0x04, 0x00, 0x00,
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
	if m.LastIdRecordVerifyRequestId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastIdRecordVerifyRequestId))
		i--
		dAtA[i] = 0x58
	}
	if m.LastIdentityRecordId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastIdentityRecordId))
		i--
		dAtA[i] = 0x50
	}
	if len(m.DataRegistry) > 0 {
		for k := range m.DataRegistry {
			v := m.DataRegistry[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintGenesis(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.PoorNetworkMessages != nil {
		{
			size, err := m.PoorNetworkMessages.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ExecutionFees) > 0 {
		for iNdEx := len(m.ExecutionFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExecutionFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.NetworkProperties != nil {
		{
			size, err := m.NetworkProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.NetworkActors) > 0 {
		for iNdEx := len(m.NetworkActors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NetworkActors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Permissions) > 0 {
		for k := range m.Permissions {
			v := m.Permissions[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintGenesis(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintGenesis(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Permissions) > 0 {
		for k, v := range m.Permissions {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGenesis(uint64(l))
			}
			mapEntrySize := 1 + sovGenesis(uint64(k)) + l
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if len(m.NetworkActors) > 0 {
		for _, e := range m.NetworkActors {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NetworkProperties != nil {
		l = m.NetworkProperties.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ExecutionFees) > 0 {
		for _, e := range m.ExecutionFees {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PoorNetworkMessages != nil {
		l = m.PoorNetworkMessages.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DataRegistry) > 0 {
		for k, v := range m.DataRegistry {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGenesis(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if m.LastIdentityRecordId != 0 {
		n += 1 + sovGenesis(uint64(m.LastIdentityRecordId))
	}
	if m.LastIdRecordVerifyRequestId != 0 {
		n += 1 + sovGenesis(uint64(m.LastIdRecordVerifyRequestId))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
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
			if m.Permissions == nil {
				m.Permissions = make(map[uint64]*Permissions)
			}
			var mapkey uint64
			var mapvalue *Permissions
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Permissions{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Permissions[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkActors", wireType)
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
			m.NetworkActors = append(m.NetworkActors, &NetworkActor{})
			if err := m.NetworkActors[len(m.NetworkActors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkProperties", wireType)
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
			if m.NetworkProperties == nil {
				m.NetworkProperties = &NetworkProperties{}
			}
			if err := m.NetworkProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionFees", wireType)
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
			m.ExecutionFees = append(m.ExecutionFees, &ExecutionFee{})
			if err := m.ExecutionFees[len(m.ExecutionFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoorNetworkMessages", wireType)
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
			if m.PoorNetworkMessages == nil {
				m.PoorNetworkMessages = &AllowedMessages{}
			}
			if err := m.PoorNetworkMessages.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRegistry", wireType)
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
			if m.DataRegistry == nil {
				m.DataRegistry = make(map[string]*DataRegistryEntry)
			}
			var mapkey string
			var mapvalue *DataRegistryEntry
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &DataRegistryEntry{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.DataRegistry[mapkey] = mapvalue
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIdentityRecordId", wireType)
			}
			m.LastIdentityRecordId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIdentityRecordId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIdRecordVerifyRequestId", wireType)
			}
			m.LastIdRecordVerifyRequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIdRecordVerifyRequestId |= uint64(b&0x7F) << shift
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
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
