// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/v1/genesis.proto

package types

import (
	fmt "fmt"
	ed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
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

// GenesisState defines the checkpointing module's genesis state.
type GenesisState struct {
	// genesis_keys defines the public keys for the genesis validators
	GenesisKeys []*GenesisKey `protobuf:"bytes,1,rep,name=genesis_keys,json=genesisKeys,proto3" json:"genesis_keys,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c524ebc9800de, []int{0}
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

func (m *GenesisState) GetGenesisKeys() []*GenesisKey {
	if m != nil {
		return m.GenesisKeys
	}
	return nil
}

// GenesisKey defines public key information about the genesis validators
type GenesisKey struct {
	// validator_address is the address corresponding to a validator
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// bls_key defines the BLS key of the validator at genesis
	BlsKey *BlsKey `protobuf:"bytes,2,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	// val_pubkey defines the ed25519 public key of the validator at genesis
	ValPubkey *ed25519.PubKey `protobuf:"bytes,3,opt,name=val_pubkey,json=valPubkey,proto3" json:"val_pubkey,omitempty"`
}

func (m *GenesisKey) Reset()         { *m = GenesisKey{} }
func (m *GenesisKey) String() string { return proto.CompactTextString(m) }
func (*GenesisKey) ProtoMessage()    {}
func (*GenesisKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c524ebc9800de, []int{1}
}
func (m *GenesisKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisKey.Merge(m, src)
}
func (m *GenesisKey) XXX_Size() int {
	return m.Size()
}
func (m *GenesisKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisKey.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisKey proto.InternalMessageInfo

func (m *GenesisKey) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *GenesisKey) GetBlsKey() *BlsKey {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *GenesisKey) GetValPubkey() *ed25519.PubKey {
	if m != nil {
		return m.ValPubkey
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "babylon.checkpointing.v1.GenesisState")
	proto.RegisterType((*GenesisKey)(nil), "babylon.checkpointing.v1.GenesisKey")
}

func init() {
	proto.RegisterFile("babylon/checkpointing/v1/genesis.proto", fileDescriptor_bf2c524ebc9800de)
}

var fileDescriptor_bf2c524ebc9800de = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xfb, 0x30,
	0x1c, 0xc6, 0x97, 0xdf, 0x60, 0x3f, 0x96, 0xed, 0xa0, 0x3d, 0x8d, 0x81, 0xa1, 0x0c, 0x91, 0x81,
	0x98, 0xb0, 0xc9, 0x90, 0x81, 0x17, 0x77, 0xd9, 0xc1, 0xcb, 0x9c, 0x07, 0xc1, 0xcb, 0x48, 0xda,
	0xd0, 0x85, 0x65, 0x4d, 0x69, 0xb2, 0x62, 0xde, 0x85, 0xaf, 0xc5, 0x57, 0xe1, 0x71, 0x47, 0x8f,
	0xb2, 0xbe, 0x11, 0x69, 0x1b, 0x27, 0x0a, 0x3d, 0xe5, 0xcf, 0xf7, 0x79, 0xbe, 0xcf, 0xc3, 0x07,
	0x5e, 0x30, 0xca, 0xac, 0x54, 0x31, 0x09, 0xd6, 0x3c, 0xd8, 0x24, 0x4a, 0xc4, 0x46, 0xc4, 0x11,
	0xc9, 0x46, 0x24, 0xe2, 0x31, 0xd7, 0x42, 0xe3, 0x24, 0x55, 0x46, 0x79, 0x3d, 0xa7, 0xc3, 0xbf,
	0x74, 0x38, 0x1b, 0xf5, 0xfd, 0x40, 0xe9, 0xad, 0xd2, 0x24, 0x48, 0x6d, 0x62, 0x14, 0xe1, 0xe1,
	0x78, 0x32, 0x19, 0x4d, 0xc9, 0x86, 0x5b, 0xe7, 0xed, 0xd7, 0x67, 0x30, 0xa9, 0x57, 0x1b, 0x6e,
	0x2b, 0xdd, 0xe0, 0x09, 0x76, 0xe7, 0x55, 0xe8, 0xa3, 0xa1, 0x86, 0x7b, 0x73, 0xd8, 0x75, 0x25,
	0x0a, 0x91, 0xee, 0x01, 0xbf, 0x39, 0xec, 0x8c, 0xcf, 0x71, 0x5d, 0x15, 0xec, 0xdc, 0xf7, 0xdc,
	0x2e, 0x3b, 0xd1, 0xf1, 0xae, 0x07, 0x6f, 0x00, 0xc2, 0x9f, 0x99, 0x77, 0x09, 0x4f, 0x33, 0x2a,
	0x45, 0x48, 0x8d, 0x4a, 0x57, 0x34, 0x0c, 0x53, 0xae, 0x8b, 0xe5, 0x60, 0xd8, 0x5e, 0x9e, 0x1c,
	0x07, 0x77, 0xd5, 0xbf, 0x37, 0x85, 0xff, 0x5d, 0xcb, 0xde, 0x3f, 0x1f, 0x0c, 0x3b, 0x63, 0xbf,
	0x3e, 0x7f, 0x26, 0xcb, 0xec, 0x16, 0x2b, 0x4f, 0xef, 0x16, 0xc2, 0x8c, 0xca, 0x55, 0xb2, 0x63,
	0x85, 0xbb, 0x59, 0xba, 0xcf, 0x70, 0x85, 0x0b, 0x57, 0xb8, 0xb0, 0xc3, 0x85, 0x17, 0x3b, 0x56,
	0x58, 0xdb, 0x19, 0x95, 0x8b, 0x52, 0x3f, 0x7b, 0x78, 0x3f, 0x20, 0xb0, 0x3f, 0x20, 0xf0, 0x79,
	0x40, 0xe0, 0x35, 0x47, 0x8d, 0x7d, 0x8e, 0x1a, 0x1f, 0x39, 0x6a, 0x3c, 0xdf, 0x44, 0xc2, 0xac,
	0x77, 0x0c, 0x07, 0x6a, 0x4b, 0x5c, 0x17, 0x49, 0x99, 0xbe, 0x12, 0xea, 0xfb, 0x49, 0x5e, 0xfe,
	0xb0, 0x36, 0x36, 0xe1, 0x9a, 0xb5, 0x4a, 0xce, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75,
	0xb3, 0x01, 0x5c, 0xf5, 0x01, 0x00, 0x00,
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
	if len(m.GenesisKeys) > 0 {
		for iNdEx := len(m.GenesisKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValPubkey != nil {
		{
			size, err := m.ValPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlsKey != nil {
		{
			size, err := m.BlsKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
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
	if len(m.GenesisKeys) > 0 {
		for _, e := range m.GenesisKeys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.BlsKey != nil {
		l = m.BlsKey.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ValPubkey != nil {
		l = m.ValPubkey.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisKeys", wireType)
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
			m.GenesisKeys = append(m.GenesisKeys, &GenesisKey{})
			if err := m.GenesisKeys[len(m.GenesisKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsKey", wireType)
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
			if m.BlsKey == nil {
				m.BlsKey = &BlsKey{}
			}
			if err := m.BlsKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValPubkey", wireType)
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
			if m.ValPubkey == nil {
				m.ValPubkey = &ed25519.PubKey{}
			}
			if err := m.ValPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
