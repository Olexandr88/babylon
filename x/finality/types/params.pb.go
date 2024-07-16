// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/finality/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// Params defines the parameters for the module.
type Params struct {
	// signed_blocks_window defines the size of the sliding window for tracking finality provider liveness
	SignedBlocksWindow int64 `protobuf:"varint,1,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty"`
	// finality_sig_timeout defines how much time (in terms of blocks) finality providers have to cast a finality
	// vote before being judged as missing their voting turn on the given block
	FinalitySigTimeout int64 `protobuf:"varint,2,opt,name=finality_sig_timeout,json=finalitySigTimeout,proto3" json:"finality_sig_timeout,omitempty"`
	// min_signed_per_window defines the minimum number of blocks that a finality provider is required to sign
	// within the sliding window to avoid being detected as inactive
	MinSignedPerWindow cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=min_signed_per_window,json=minSignedPerWindow,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_signed_per_window"`
	// min_pub_rand is the minimum number of public randomness each
	// message should commit
	MinPubRand uint64 `protobuf:"varint,4,opt,name=min_pub_rand,json=minPubRand,proto3" json:"min_pub_rand,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_25539c9a61c72ee9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSignedBlocksWindow() int64 {
	if m != nil {
		return m.SignedBlocksWindow
	}
	return 0
}

func (m *Params) GetFinalitySigTimeout() int64 {
	if m != nil {
		return m.FinalitySigTimeout
	}
	return 0
}

func (m *Params) GetMinPubRand() uint64 {
	if m != nil {
		return m.MinPubRand
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "babylon.finality.v1.Params")
}

func init() { proto.RegisterFile("babylon/finality/v1/params.proto", fileDescriptor_25539c9a61c72ee9) }

var fileDescriptor_25539c9a61c72ee9 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x73, 0xff, 0x96, 0x0e, 0xa1, 0xcb, 0x3f, 0x56, 0xa8, 0x15, 0xd2, 0xe0, 0x54, 0x04,
	0x73, 0x2d, 0x82, 0x83, 0x63, 0xe9, 0x24, 0x0e, 0x25, 0x15, 0x04, 0x97, 0x70, 0x49, 0xce, 0xf4,
	0xa5, 0xbd, 0xbb, 0x90, 0x4b, 0x5a, 0xf3, 0x2d, 0x1c, 0x1d, 0x1d, 0x1d, 0x1d, 0xfc, 0x10, 0x1d,
	0x8b, 0x93, 0x38, 0x14, 0x69, 0x07, 0x3f, 0x86, 0xd2, 0xbb, 0x04, 0x97, 0xe3, 0xde, 0xfb, 0x3d,
	0xef, 0x3d, 0xcf, 0xbd, 0x67, 0x3a, 0x01, 0x09, 0x8a, 0xb9, 0xe0, 0xf8, 0x1e, 0x38, 0x99, 0x43,
	0x56, 0xe0, 0xc5, 0x00, 0x27, 0x24, 0x25, 0x4c, 0xba, 0x49, 0x2a, 0x32, 0x61, 0x1d, 0x94, 0x0a,
	0xb7, 0x52, 0xb8, 0x8b, 0x41, 0xa7, 0x15, 0x8b, 0x58, 0x28, 0x8e, 0xf7, 0x3b, 0x2d, 0xed, 0xfc,
	0x27, 0x0c, 0xb8, 0xc0, 0x6a, 0x2d, 0x8f, 0x8e, 0x42, 0x21, 0x99, 0x90, 0xbe, 0xd6, 0xea, 0x42,
	0xa3, 0x93, 0x1f, 0x64, 0x36, 0xc6, 0xca, 0xc9, 0xea, 0x9b, 0x2d, 0x09, 0x31, 0xa7, 0x91, 0x1f,
	0xcc, 0x45, 0x38, 0x93, 0xfe, 0x12, 0x78, 0x24, 0x96, 0x6d, 0xe4, 0xa0, 0x5e, 0xcd, 0xb3, 0x34,
	0x1b, 0x2a, 0x74, 0xab, 0xc8, 0xbe, 0xa3, 0xca, 0xe3, 0x4b, 0x88, 0xfd, 0x0c, 0x18, 0x15, 0x79,
	0xd6, 0xfe, 0xa7, 0x3b, 0x2a, 0x36, 0x81, 0xf8, 0x46, 0x13, 0x0b, 0xcc, 0x43, 0x06, 0xdc, 0x2f,
	0x7d, 0x12, 0x9a, 0x56, 0x26, 0x35, 0x07, 0xf5, 0x9a, 0xc3, 0x8b, 0xd5, 0xa6, 0x6b, 0x7c, 0x6e,
	0xba, 0xc7, 0x3a, 0xa3, 0x8c, 0x66, 0x2e, 0x08, 0xcc, 0x48, 0x36, 0x75, 0xaf, 0x69, 0x4c, 0xc2,
	0x62, 0x44, 0xc3, 0xf7, 0xb7, 0x33, 0xb3, 0x7c, 0xc2, 0x88, 0x86, 0x2f, 0xdf, 0xaf, 0xa7, 0xc8,
	0xb3, 0x18, 0xf0, 0x89, 0xba, 0x73, 0x4c, 0xd3, 0x32, 0x9c, 0x63, 0x36, 0xf7, 0x56, 0x49, 0x1e,
	0xf8, 0x29, 0xe1, 0x51, 0xbb, 0xee, 0xa0, 0x5e, 0xdd, 0x33, 0x19, 0xf0, 0x71, 0x1e, 0x78, 0x84,
	0x47, 0x97, 0xf5, 0xa7, 0xe7, 0xae, 0x31, 0xbc, 0x5a, 0x6d, 0x6d, 0xb4, 0xde, 0xda, 0xe8, 0x6b,
	0x6b, 0xa3, 0xc7, 0x9d, 0x6d, 0xac, 0x77, 0xb6, 0xf1, 0xb1, 0xb3, 0x8d, 0xbb, 0x7e, 0x0c, 0xd9,
	0x34, 0x0f, 0xdc, 0x50, 0x30, 0x5c, 0xce, 0x3f, 0x9c, 0x12, 0xe0, 0x55, 0x81, 0x1f, 0xfe, 0x3e,
	0x2c, 0x2b, 0x12, 0x2a, 0x83, 0x86, 0x1a, 0xea, 0xf9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0,
	0x39, 0x84, 0xbb, 0xd1, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinPubRand != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinPubRand))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.MinSignedPerWindow.Size()
		i -= size
		if _, err := m.MinSignedPerWindow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.FinalitySigTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FinalitySigTimeout))
		i--
		dAtA[i] = 0x10
	}
	if m.SignedBlocksWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SignedBlocksWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedBlocksWindow != 0 {
		n += 1 + sovParams(uint64(m.SignedBlocksWindow))
	}
	if m.FinalitySigTimeout != 0 {
		n += 1 + sovParams(uint64(m.FinalitySigTimeout))
	}
	l = m.MinSignedPerWindow.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinPubRand != 0 {
		n += 1 + sovParams(uint64(m.MinPubRand))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
			}
			m.SignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBlocksWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalitySigTimeout", wireType)
			}
			m.FinalitySigTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalitySigTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSignedPerWindow", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSignedPerWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPubRand", wireType)
			}
			m.MinPubRand = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinPubRand |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
