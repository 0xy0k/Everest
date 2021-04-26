// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tsuki/gov/actor.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActorStatus int32

const (
	// Undefined status
	ActorStatus_UNDEFINED ActorStatus = 0
	// Unclaimed status
	ActorStatus_UNCLAIMED ActorStatus = 1
	// Active status
	ActorStatus_ACTIVE ActorStatus = 2
	// Paused status
	ActorStatus_PAUSED ActorStatus = 3
	// Inactive status
	ActorStatus_INACTIVE ActorStatus = 4
	// Jailed status
	ActorStatus_JAILED ActorStatus = 5
	// Removed status
	ActorStatus_REMOVED ActorStatus = 6
)

var ActorStatus_name = map[int32]string{
	0: "UNDEFINED",
	1: "UNCLAIMED",
	2: "ACTIVE",
	3: "PAUSED",
	4: "INACTIVE",
	5: "JAILED",
	6: "REMOVED",
}

var ActorStatus_value = map[string]int32{
	"UNDEFINED": 0,
	"UNCLAIMED": 1,
	"ACTIVE":    2,
	"PAUSED":    3,
	"INACTIVE":  4,
	"JAILED":    5,
	"REMOVED":   6,
}

func (x ActorStatus) String() string {
	return proto.EnumName(ActorStatus_name, int32(x))
}

func (ActorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{0}
}

type Permissions struct {
	Blacklist            []uint32 `protobuf:"varint,1,rep,packed,name=blacklist,proto3" json:"blacklist,omitempty"`
	Whitelist            []uint32 `protobuf:"varint,2,rep,packed,name=whitelist,proto3" json:"whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{0}
}

func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetBlacklist() []uint32 {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *Permissions) GetWhitelist() []uint32 {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

type NetworkActor struct {
	Address              []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Roles                []uint64     `protobuf:"varint,2,rep,packed,name=roles,proto3" json:"roles,omitempty"`
	Status               ActorStatus  `protobuf:"varint,3,opt,name=status,proto3,enum=tsuki.gov.ActorStatus" json:"status,omitempty"`
	Votes                []VoteOption `protobuf:"varint,4,rep,packed,name=votes,proto3,enum=tsuki.gov.VoteOption" json:"votes,omitempty"`
	Permissions          *Permissions `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Skin                 uint64       `protobuf:"varint,6,opt,name=skin,proto3" json:"skin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NetworkActor) Reset()         { *m = NetworkActor{} }
func (m *NetworkActor) String() string { return proto.CompactTextString(m) }
func (*NetworkActor) ProtoMessage()    {}
func (*NetworkActor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{1}
}

func (m *NetworkActor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkActor.Unmarshal(m, b)
}
func (m *NetworkActor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkActor.Marshal(b, m, deterministic)
}
func (m *NetworkActor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkActor.Merge(m, src)
}
func (m *NetworkActor) XXX_Size() int {
	return xxx_messageInfo_NetworkActor.Size(m)
}
func (m *NetworkActor) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkActor.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkActor proto.InternalMessageInfo

func (m *NetworkActor) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NetworkActor) GetRoles() []uint64 {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *NetworkActor) GetStatus() ActorStatus {
	if m != nil {
		return m.Status
	}
	return ActorStatus_UNDEFINED
}

func (m *NetworkActor) GetVotes() []VoteOption {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *NetworkActor) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *NetworkActor) GetSkin() uint64 {
	if m != nil {
		return m.Skin
	}
	return 0
}

type MsgWhitelistPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgWhitelistPermissions) Reset()         { *m = MsgWhitelistPermissions{} }
func (m *MsgWhitelistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgWhitelistPermissions) ProtoMessage()    {}
func (*MsgWhitelistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{2}
}

func (m *MsgWhitelistPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgWhitelistPermissions.Unmarshal(m, b)
}
func (m *MsgWhitelistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgWhitelistPermissions.Marshal(b, m, deterministic)
}
func (m *MsgWhitelistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWhitelistPermissions.Merge(m, src)
}
func (m *MsgWhitelistPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgWhitelistPermissions.Size(m)
}
func (m *MsgWhitelistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWhitelistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWhitelistPermissions proto.InternalMessageInfo

func (m *MsgWhitelistPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgWhitelistPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgWhitelistPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type MsgBlacklistPermissions struct {
	Proposer             []byte   `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Permission           uint32   `protobuf:"varint,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBlacklistPermissions) Reset()         { *m = MsgBlacklistPermissions{} }
func (m *MsgBlacklistPermissions) String() string { return proto.CompactTextString(m) }
func (*MsgBlacklistPermissions) ProtoMessage()    {}
func (*MsgBlacklistPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4584595efc3386e, []int{3}
}

func (m *MsgBlacklistPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBlacklistPermissions.Unmarshal(m, b)
}
func (m *MsgBlacklistPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBlacklistPermissions.Marshal(b, m, deterministic)
}
func (m *MsgBlacklistPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBlacklistPermissions.Merge(m, src)
}
func (m *MsgBlacklistPermissions) XXX_Size() int {
	return xxx_messageInfo_MsgBlacklistPermissions.Size(m)
}
func (m *MsgBlacklistPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBlacklistPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBlacklistPermissions proto.InternalMessageInfo

func (m *MsgBlacklistPermissions) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func (m *MsgBlacklistPermissions) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgBlacklistPermissions) GetPermission() uint32 {
	if m != nil {
		return m.Permission
	}
	return 0
}

func init() {
	proto.RegisterEnum("tsuki.gov.ActorStatus", ActorStatus_name, ActorStatus_value)
	proto.RegisterType((*Permissions)(nil), "tsuki.gov.Permissions")
	proto.RegisterType((*NetworkActor)(nil), "tsuki.gov.NetworkActor")
	proto.RegisterType((*MsgWhitelistPermissions)(nil), "tsuki.gov.MsgWhitelistPermissions")
	proto.RegisterType((*MsgBlacklistPermissions)(nil), "tsuki.gov.MsgBlacklistPermissions")
}

func init() {
	proto.RegisterFile("tsuki/gov/actor.proto", fileDescriptor_c4584595efc3386e)
}

var fileDescriptor_c4584595efc3386e = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x8d, 0x6c, 0x59, 0x71, 0xc6, 0x49, 0x10, 0x83, 0xdb, 0x08, 0x51, 0x12, 0xa1, 0x45, 0x11,
	0x81, 0x48, 0x34, 0x5d, 0x14, 0xb2, 0x53, 0x2c, 0x15, 0x94, 0xc6, 0x4e, 0x50, 0x1e, 0x2d, 0x85,
	0x2e, 0xc6, 0xd2, 0x54, 0x19, 0xf4, 0x18, 0xa1, 0x19, 0x3b, 0xe4, 0x0f, 0x8a, 0xff, 0xc1, 0x50,
	0xe8, 0xb2, 0xdf, 0xd2, 0x5f, 0xc8, 0xb6, 0xdd, 0x74, 0xd3, 0x65, 0x57, 0x45, 0x92, 0x1f, 0xda,
	0x76, 0xd1, 0x4d, 0x57, 0x33, 0x3a, 0xe7, 0x9e, 0xcb, 0xbd, 0xe7, 0x5e, 0x0d, 0xe8, 0xc7, 0xa4,
	0x40, 0x56, 0x44, 0xa7, 0x16, 0x0a, 0x38, 0x2d, 0xcc, 0xbc, 0xa0, 0x9c, 0xc2, 0x6e, 0x89, 0x9a,
	0x11, 0x9d, 0xaa, 0xfd, 0x88, 0x46, 0xb4, 0x02, 0xad, 0xf2, 0x56, 0xf3, 0xea, 0xde, 0x4a, 0x95,
	0x17, 0x34, 0xa7, 0x0c, 0x25, 0x35, 0xa1, 0x7b, 0xa0, 0x77, 0x89, 0x8b, 0x94, 0x30, 0x46, 0x68,
	0xc6, 0xe0, 0x33, 0xb0, 0x35, 0x4e, 0x50, 0x10, 0x27, 0x84, 0x71, 0x45, 0xd0, 0xda, 0xc6, 0x8e,
	0xbf, 0x06, 0x4a, 0xf6, 0xfe, 0x8e, 0x70, 0x5c, 0xb1, 0xad, 0x9a, 0x5d, 0x01, 0xfa, 0xd7, 0x16,
	0xd8, 0x1e, 0x61, 0x7e, 0x4f, 0x8b, 0xd8, 0x2e, 0x4b, 0x83, 0x1f, 0xc0, 0x26, 0x0a, 0xc3, 0x02,
	0x33, 0xa6, 0x08, 0x9a, 0x60, 0x6c, 0x9f, 0x0e, 0x7e, 0x3d, 0x1e, 0xec, 0x3e, 0xa0, 0x34, 0x39,
	0xd1, 0x17, 0x84, 0xfe, 0xfb, 0xf1, 0xe0, 0x28, 0x22, 0xfc, 0x6e, 0x32, 0x36, 0x03, 0x9a, 0x5a,
	0x01, 0x65, 0x29, 0x65, 0x8b, 0xe3, 0x88, 0x85, 0xb1, 0xc5, 0x1f, 0x72, 0xcc, 0x4c, 0x3b, 0x08,
	0xec, 0x5a, 0xe1, 0x2f, 0x73, 0xc2, 0x3e, 0xe8, 0x14, 0x34, 0xc1, 0xac, 0xaa, 0x44, 0xf4, 0xeb,
	0x0f, 0x78, 0x04, 0x24, 0xc6, 0x11, 0x9f, 0x30, 0xa5, 0xad, 0x09, 0xc6, 0xee, 0xf1, 0x13, 0x73,
	0x69, 0x8d, 0x59, 0x55, 0x75, 0x55, 0x91, 0xfe, 0x22, 0x08, 0x1e, 0x82, 0xce, 0x94, 0x72, 0xcc,
	0x14, 0x51, 0x6b, 0x1b, 0xbb, 0xc7, 0xfd, 0x75, 0xf4, 0x2d, 0xe5, 0xf8, 0x22, 0xe7, 0x84, 0x66,
	0x7e, 0x1d, 0x02, 0x5f, 0x81, 0x5e, 0xbe, 0xf6, 0x4a, 0xe9, 0x68, 0x82, 0xd1, 0x6b, 0xe6, 0x6f,
	0x18, 0xe9, 0x37, 0x23, 0x21, 0x04, 0x22, 0x8b, 0x49, 0xa6, 0x48, 0x9a, 0x60, 0x88, 0x7e, 0x75,
	0xd7, 0x7f, 0x0a, 0x60, 0x6f, 0xc8, 0xa2, 0xb7, 0x4b, 0xfb, 0x9a, 0x53, 0x18, 0x82, 0x6e, 0x3d,
	0x26, 0x5c, 0x2c, 0x9c, 0x7b, 0xf1, 0xf7, 0x3e, 0xad, 0x52, 0x34, 0xe7, 0xd0, 0xfa, 0x07, 0x73,
	0xd8, 0x07, 0x60, 0xdd, 0x6c, 0xe5, 0xfa, 0x8e, 0xdf, 0x40, 0x4e, 0xc4, 0x1f, 0x9f, 0x0f, 0x04,
	0xfd, 0x7b, 0xdd, 0xef, 0xe9, 0x72, 0x99, 0xfe, 0xdb, 0x7e, 0x0f, 0xbf, 0x09, 0xa0, 0xd7, 0x58,
	0xb5, 0xf2, 0xaf, 0xb9, 0x19, 0x39, 0xee, 0x6b, 0x6f, 0xe4, 0x3a, 0xf2, 0x86, 0xba, 0x33, 0x9b,
	0x6b, 0x5b, 0x37, 0x59, 0x88, 0x3f, 0x92, 0x0c, 0x87, 0x35, 0x3b, 0x38, 0xb7, 0xbd, 0xa1, 0xeb,
	0xc8, 0xc2, 0x92, 0x0d, 0x12, 0x44, 0x52, 0x1c, 0xc2, 0xa7, 0x40, 0xb2, 0x07, 0xd7, 0xde, 0xad,
	0x2b, 0xb7, 0x54, 0x30, 0x9b, 0x6b, 0x92, 0x1d, 0x70, 0x32, 0xc5, 0x25, 0x7e, 0x69, 0xdf, 0x5c,
	0xb9, 0x8e, 0xdc, 0xae, 0xf1, 0x4b, 0x34, 0x61, 0x38, 0x84, 0x2a, 0xe8, 0x7a, 0xa3, 0x85, 0x42,
	0x54, 0xb7, 0x67, 0x73, 0xad, 0xeb, 0x65, 0x68, 0xa5, 0x39, 0xb3, 0xbd, 0x73, 0xd7, 0x91, 0x3b,
	0xb5, 0xe6, 0x0c, 0x91, 0x04, 0x87, 0x50, 0x01, 0x9b, 0xbe, 0x3b, 0xbc, 0xb8, 0x75, 0x1d, 0x59,
	0x52, 0x7b, 0xb3, 0xb9, 0xb6, 0xe9, 0xe3, 0x94, 0x4e, 0x71, 0xa8, 0x8a, 0x9f, 0xbe, 0xec, 0x6f,
	0x9c, 0x1a, 0xef, 0x9f, 0x37, 0x9c, 0x7a, 0x43, 0x0a, 0x34, 0xa0, 0x05, 0xb6, 0x18, 0x8e, 0x11,
	0xb1, 0xbc, 0xd1, 0xb5, 0xeb, 0xbf, 0xb3, 0xaa, 0xc7, 0x64, 0x2c, 0x55, 0xc7, 0xcb, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x01, 0xda, 0xf3, 0xd6, 0xa4, 0x04, 0x00, 0x00,
}
