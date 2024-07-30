// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btcstaking/v1/events.proto

package types

import (
	fmt "fmt"
	github_com_babylonlabs_io_babylon_types "github.com/babylonlabs-io/babylon/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// EventNewFinalityProvider is the event emitted when a finality provider is created
type EventNewFinalityProvider struct {
	Fp *FinalityProvider `protobuf:"bytes,1,opt,name=fp,proto3" json:"fp,omitempty"`
}

func (m *EventNewFinalityProvider) Reset()         { *m = EventNewFinalityProvider{} }
func (m *EventNewFinalityProvider) String() string { return proto.CompactTextString(m) }
func (*EventNewFinalityProvider) ProtoMessage()    {}
func (*EventNewFinalityProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{0}
}
func (m *EventNewFinalityProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventNewFinalityProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventNewFinalityProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventNewFinalityProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNewFinalityProvider.Merge(m, src)
}
func (m *EventNewFinalityProvider) XXX_Size() int {
	return m.Size()
}
func (m *EventNewFinalityProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNewFinalityProvider.DiscardUnknown(m)
}

var xxx_messageInfo_EventNewFinalityProvider proto.InternalMessageInfo

func (m *EventNewFinalityProvider) GetFp() *FinalityProvider {
	if m != nil {
		return m.Fp
	}
	return nil
}

// EventBTCDelegationStateUpdate is the event emitted when a BTC delegation's state is
// updated. There are the following possible state transitions:
// - non-existing -> pending, which happens upon `MsgCreateBTCDelegation`
// - pending -> active, which happens upon `MsgAddCovenantSigs`
// - active -> unbonded, which happens upon `MsgBTCUndelegate` or upon staking tx timelock expires
type EventBTCDelegationStateUpdate struct {
	// staking_tx_hash is the hash of the staking tx.
	// It uniquely identifies a BTC delegation
	StakingTxHash string `protobuf:"bytes,1,opt,name=staking_tx_hash,json=stakingTxHash,proto3" json:"staking_tx_hash,omitempty"`
	// new_state is the new state of this BTC delegation
	NewState BTCDelegationStatus `protobuf:"varint,2,opt,name=new_state,json=newState,proto3,enum=babylon.btcstaking.v1.BTCDelegationStatus" json:"new_state,omitempty"`
}

func (m *EventBTCDelegationStateUpdate) Reset()         { *m = EventBTCDelegationStateUpdate{} }
func (m *EventBTCDelegationStateUpdate) String() string { return proto.CompactTextString(m) }
func (*EventBTCDelegationStateUpdate) ProtoMessage()    {}
func (*EventBTCDelegationStateUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{1}
}
func (m *EventBTCDelegationStateUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBTCDelegationStateUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBTCDelegationStateUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBTCDelegationStateUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBTCDelegationStateUpdate.Merge(m, src)
}
func (m *EventBTCDelegationStateUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventBTCDelegationStateUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBTCDelegationStateUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventBTCDelegationStateUpdate proto.InternalMessageInfo

func (m *EventBTCDelegationStateUpdate) GetStakingTxHash() string {
	if m != nil {
		return m.StakingTxHash
	}
	return ""
}

func (m *EventBTCDelegationStateUpdate) GetNewState() BTCDelegationStatus {
	if m != nil {
		return m.NewState
	}
	return BTCDelegationStatus_PENDING
}

// EventSelectiveSlashing is the event emitted when an adversarial
// finality provider selectively slashes a BTC delegation. This will
// result in slashing of all BTC delegations under this finality provider.
type EventSelectiveSlashing struct {
	// evidence is the evidence of selective slashing
	Evidence *SelectiveSlashingEvidence `protobuf:"bytes,1,opt,name=evidence,proto3" json:"evidence,omitempty"`
}

func (m *EventSelectiveSlashing) Reset()         { *m = EventSelectiveSlashing{} }
func (m *EventSelectiveSlashing) String() string { return proto.CompactTextString(m) }
func (*EventSelectiveSlashing) ProtoMessage()    {}
func (*EventSelectiveSlashing) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{2}
}
func (m *EventSelectiveSlashing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSelectiveSlashing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSelectiveSlashing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSelectiveSlashing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSelectiveSlashing.Merge(m, src)
}
func (m *EventSelectiveSlashing) XXX_Size() int {
	return m.Size()
}
func (m *EventSelectiveSlashing) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSelectiveSlashing.DiscardUnknown(m)
}

var xxx_messageInfo_EventSelectiveSlashing proto.InternalMessageInfo

func (m *EventSelectiveSlashing) GetEvidence() *SelectiveSlashingEvidence {
	if m != nil {
		return m.Evidence
	}
	return nil
}

// EventPowerDistUpdate is an event that affects voting power distirbution
// of BTC staking protocol
type EventPowerDistUpdate struct {
	// ev is the event that affects voting power distribution
	//
	// Types that are valid to be assigned to Ev:
	//	*EventPowerDistUpdate_SlashedFp
	//	*EventPowerDistUpdate_BtcDelStateUpdate
	Ev isEventPowerDistUpdate_Ev `protobuf_oneof:"ev"`
}

func (m *EventPowerDistUpdate) Reset()         { *m = EventPowerDistUpdate{} }
func (m *EventPowerDistUpdate) String() string { return proto.CompactTextString(m) }
func (*EventPowerDistUpdate) ProtoMessage()    {}
func (*EventPowerDistUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{3}
}
func (m *EventPowerDistUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPowerDistUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPowerDistUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPowerDistUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPowerDistUpdate.Merge(m, src)
}
func (m *EventPowerDistUpdate) XXX_Size() int {
	return m.Size()
}
func (m *EventPowerDistUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPowerDistUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EventPowerDistUpdate proto.InternalMessageInfo

type isEventPowerDistUpdate_Ev interface {
	isEventPowerDistUpdate_Ev()
	MarshalTo([]byte) (int, error)
	Size() int
}

type EventPowerDistUpdate_SlashedFp struct {
	SlashedFp *EventPowerDistUpdate_EventSlashedFinalityProvider `protobuf:"bytes,1,opt,name=slashed_fp,json=slashedFp,proto3,oneof" json:"slashed_fp,omitempty"`
}
type EventPowerDistUpdate_BtcDelStateUpdate struct {
	BtcDelStateUpdate *EventBTCDelegationStateUpdate `protobuf:"bytes,2,opt,name=btc_del_state_update,json=btcDelStateUpdate,proto3,oneof" json:"btc_del_state_update,omitempty"`
}

func (*EventPowerDistUpdate_SlashedFp) isEventPowerDistUpdate_Ev()         {}
func (*EventPowerDistUpdate_BtcDelStateUpdate) isEventPowerDistUpdate_Ev() {}

func (m *EventPowerDistUpdate) GetEv() isEventPowerDistUpdate_Ev {
	if m != nil {
		return m.Ev
	}
	return nil
}

func (m *EventPowerDistUpdate) GetSlashedFp() *EventPowerDistUpdate_EventSlashedFinalityProvider {
	if x, ok := m.GetEv().(*EventPowerDistUpdate_SlashedFp); ok {
		return x.SlashedFp
	}
	return nil
}

func (m *EventPowerDistUpdate) GetBtcDelStateUpdate() *EventBTCDelegationStateUpdate {
	if x, ok := m.GetEv().(*EventPowerDistUpdate_BtcDelStateUpdate); ok {
		return x.BtcDelStateUpdate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventPowerDistUpdate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventPowerDistUpdate_SlashedFp)(nil),
		(*EventPowerDistUpdate_BtcDelStateUpdate)(nil),
	}
}

// EventSlashedFinalityProvider defines an event that a finality provider
// is slashed
// TODO: unify with existing slashing events
type EventPowerDistUpdate_EventSlashedFinalityProvider struct {
	Pk *github_com_babylonlabs_io_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=pk,proto3,customtype=github.com/babylonlabs-io/babylon/types.BIP340PubKey" json:"pk,omitempty"`
}

func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) Reset() {
	*m = EventPowerDistUpdate_EventSlashedFinalityProvider{}
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) String() string {
	return proto.CompactTextString(m)
}
func (*EventPowerDistUpdate_EventSlashedFinalityProvider) ProtoMessage() {}
func (*EventPowerDistUpdate_EventSlashedFinalityProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{3, 0}
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPowerDistUpdate_EventSlashedFinalityProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPowerDistUpdate_EventSlashedFinalityProvider.Merge(m, src)
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) XXX_Size() int {
	return m.Size()
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPowerDistUpdate_EventSlashedFinalityProvider.DiscardUnknown(m)
}

var xxx_messageInfo_EventPowerDistUpdate_EventSlashedFinalityProvider proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventNewFinalityProvider)(nil), "babylon.btcstaking.v1.EventNewFinalityProvider")
	proto.RegisterType((*EventBTCDelegationStateUpdate)(nil), "babylon.btcstaking.v1.EventBTCDelegationStateUpdate")
	proto.RegisterType((*EventSelectiveSlashing)(nil), "babylon.btcstaking.v1.EventSelectiveSlashing")
	proto.RegisterType((*EventPowerDistUpdate)(nil), "babylon.btcstaking.v1.EventPowerDistUpdate")
	proto.RegisterType((*EventPowerDistUpdate_EventSlashedFinalityProvider)(nil), "babylon.btcstaking.v1.EventPowerDistUpdate.EventSlashedFinalityProvider")
}

func init() {
	proto.RegisterFile("babylon/btcstaking/v1/events.proto", fileDescriptor_74118427820fff75)
}

var fileDescriptor_74118427820fff75 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x6d, 0x0b, 0xa1, 0x66, 0xca, 0x8f, 0xb0, 0x02, 0x8a, 0x22, 0x30, 0x95, 0x17, 0xa5,
	0x42, 0xc2, 0x6e, 0xd3, 0x48, 0xb0, 0x36, 0x69, 0x31, 0x02, 0x55, 0x91, 0x5d, 0x36, 0x6c, 0xac,
	0x19, 0xe7, 0xc6, 0x1e, 0xc5, 0xcc, 0x58, 0x99, 0x89, 0x93, 0xbc, 0x45, 0x1f, 0x8b, 0x65, 0x97,
	0xa8, 0x0b, 0x84, 0x92, 0x17, 0x41, 0x9e, 0x0c, 0x25, 0x6a, 0x93, 0xaa, 0xbb, 0xe4, 0xea, 0x9c,
	0xf3, 0x9d, 0x7b, 0xad, 0x41, 0x2e, 0xc1, 0x64, 0x5e, 0x70, 0xe6, 0x13, 0x99, 0x0a, 0x89, 0x47,
	0x94, 0x65, 0x7e, 0x75, 0xe4, 0x43, 0x05, 0x4c, 0x0a, 0xaf, 0x1c, 0x73, 0xc9, 0xed, 0xe7, 0x5a,
	0xe3, 0xfd, 0xd7, 0x78, 0xd5, 0x51, 0xbb, 0x99, 0xf1, 0x8c, 0x2b, 0x85, 0x5f, 0xff, 0x5a, 0x89,
	0xdb, 0xfb, 0x9b, 0x03, 0xd7, 0xac, 0x4a, 0xe7, 0xc6, 0xa8, 0x75, 0x52, 0x43, 0xce, 0x60, 0x7a,
	0x4a, 0x19, 0x2e, 0xa8, 0x9c, 0xf7, 0xc7, 0xbc, 0xa2, 0x03, 0x18, 0xdb, 0xef, 0x91, 0x35, 0x2c,
	0x5b, 0xe6, 0x9e, 0x79, 0xb0, 0xdb, 0x79, 0xe3, 0x6d, 0xa4, 0x7b, 0x37, 0x4d, 0x91, 0x35, 0x2c,
	0xdd, 0x0b, 0x13, 0xbd, 0x52, 0xa9, 0xc1, 0xf9, 0xc7, 0x1e, 0x14, 0x90, 0x61, 0x49, 0x39, 0x8b,
	0x25, 0x96, 0xf0, 0xad, 0x1c, 0x60, 0x09, 0xf6, 0x3e, 0x7a, 0xaa, 0x43, 0x12, 0x39, 0x4b, 0x72,
	0x2c, 0x72, 0xc5, 0x69, 0x44, 0x8f, 0xf5, 0xf8, 0x7c, 0x16, 0x62, 0x91, 0xdb, 0x9f, 0x50, 0x83,
	0xc1, 0x34, 0x11, 0xb5, 0xb5, 0x65, 0xed, 0x99, 0x07, 0x4f, 0x3a, 0x6f, 0xb7, 0x34, 0xb9, 0xc5,
	0x9a, 0x88, 0x68, 0x87, 0xc1, 0x54, 0x61, 0xdd, 0x21, 0x7a, 0xa1, 0x1a, 0xc5, 0x50, 0x40, 0x2a,
	0x69, 0x05, 0x71, 0x81, 0x45, 0x4e, 0x59, 0x66, 0x7f, 0x45, 0x3b, 0x50, 0x57, 0x67, 0x29, 0xe8,
	0x5d, 0x0f, 0xb7, 0x10, 0x6e, 0x79, 0x4f, 0xb4, 0x2f, 0xba, 0x4e, 0x70, 0xaf, 0x2c, 0xd4, 0x54,
	0xa0, 0x3e, 0x9f, 0xc2, 0xb8, 0x47, 0x85, 0xd4, 0x1b, 0x53, 0x84, 0x44, 0x6d, 0x83, 0x41, 0x72,
	0x7d, 0xd4, 0x70, 0x0b, 0x68, 0x53, 0xc0, 0x6a, 0x18, 0xaf, 0x22, 0x6e, 0x5e, 0x3d, 0x34, 0xa2,
	0x86, 0x4e, 0x3f, 0x2d, 0xed, 0x0c, 0x35, 0x89, 0x4c, 0x93, 0x01, 0x14, 0xab, 0xc3, 0x25, 0x13,
	0x95, 0xa0, 0xee, 0xb7, 0xdb, 0xe9, 0xde, 0x05, 0xdd, 0xf6, 0xc1, 0x42, 0x23, 0x7a, 0x46, 0x64,
	0xda, 0x83, 0x62, 0x6d, 0xd8, 0xce, 0xd1, 0xcb, 0xbb, 0x5a, 0xd9, 0x21, 0xb2, 0xca, 0x91, 0xda,
	0xf5, 0x51, 0xf0, 0xe1, 0xea, 0xf7, 0xeb, 0x6e, 0x46, 0x65, 0x3e, 0x21, 0x5e, 0xca, 0x7f, 0xf8,
	0xba, 0x44, 0x81, 0x89, 0x78, 0x47, 0xf9, 0xbf, 0xbf, 0xbe, 0x9c, 0x97, 0x20, 0xbc, 0xe0, 0x73,
	0xff, 0xb8, 0x7b, 0xd8, 0x9f, 0x90, 0x2f, 0x30, 0x8f, 0xac, 0x72, 0x14, 0x3c, 0x40, 0x16, 0x54,
	0xc1, 0xd9, 0xcf, 0x85, 0x63, 0x5e, 0x2e, 0x1c, 0xf3, 0xcf, 0xc2, 0x31, 0x2f, 0x96, 0x8e, 0x71,
	0xb9, 0x74, 0x8c, 0x5f, 0x4b, 0xc7, 0xf8, 0x7e, 0x8f, 0xe4, 0xd9, 0xfa, 0x53, 0x50, 0x18, 0xf2,
	0x50, 0xbd, 0x81, 0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x5a, 0xe3, 0x63, 0x7e, 0x03,
	0x00, 0x00,
}

func (m *EventNewFinalityProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventNewFinalityProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventNewFinalityProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fp != nil {
		{
			size, err := m.Fp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBTCDelegationStateUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBTCDelegationStateUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBTCDelegationStateUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewState != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.NewState))
		i--
		dAtA[i] = 0x10
	}
	if len(m.StakingTxHash) > 0 {
		i -= len(m.StakingTxHash)
		copy(dAtA[i:], m.StakingTxHash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.StakingTxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSelectiveSlashing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSelectiveSlashing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSelectiveSlashing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Evidence != nil {
		{
			size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventPowerDistUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPowerDistUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPowerDistUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ev != nil {
		{
			size := m.Ev.Size()
			i -= size
			if _, err := m.Ev.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *EventPowerDistUpdate_SlashedFp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPowerDistUpdate_SlashedFp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SlashedFp != nil {
		{
			size, err := m.SlashedFp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *EventPowerDistUpdate_BtcDelStateUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPowerDistUpdate_BtcDelStateUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BtcDelStateUpdate != nil {
		{
			size, err := m.BtcDelStateUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pk != nil {
		{
			size := m.Pk.Size()
			i -= size
			if _, err := m.Pk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventNewFinalityProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fp != nil {
		l = m.Fp.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventBTCDelegationStateUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StakingTxHash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.NewState != 0 {
		n += 1 + sovEvents(uint64(m.NewState))
	}
	return n
}

func (m *EventSelectiveSlashing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Evidence != nil {
		l = m.Evidence.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventPowerDistUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ev != nil {
		n += m.Ev.Size()
	}
	return n
}

func (m *EventPowerDistUpdate_SlashedFp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SlashedFp != nil {
		l = m.SlashedFp.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}
func (m *EventPowerDistUpdate_BtcDelStateUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcDelStateUpdate != nil {
		l = m.BtcDelStateUpdate.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pk != nil {
		l = m.Pk.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventNewFinalityProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventNewFinalityProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventNewFinalityProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fp == nil {
				m.Fp = &FinalityProvider{}
			}
			if err := m.Fp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBTCDelegationStateUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBTCDelegationStateUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBTCDelegationStateUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakingTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewState", wireType)
			}
			m.NewState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewState |= BTCDelegationStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSelectiveSlashing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSelectiveSlashing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSelectiveSlashing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evidence == nil {
				m.Evidence = &SelectiveSlashingEvidence{}
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventPowerDistUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventPowerDistUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPowerDistUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashedFp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EventPowerDistUpdate_EventSlashedFinalityProvider{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Ev = &EventPowerDistUpdate_SlashedFp{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcDelStateUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EventBTCDelegationStateUpdate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Ev = &EventPowerDistUpdate_BtcDelStateUpdate{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventPowerDistUpdate_EventSlashedFinalityProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSlashedFinalityProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSlashedFinalityProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonlabs_io_babylon_types.BIP340PubKey
			m.Pk = &v
			if err := m.Pk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
