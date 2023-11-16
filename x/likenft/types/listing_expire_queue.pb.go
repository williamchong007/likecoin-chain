// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/v1/listing_expire_queue.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListingExpireQueueEntry struct {
	ExpireTime time.Time `protobuf:"bytes,1,opt,name=expire_time,json=expireTime,proto3,stdtime" json:"expire_time"`
	ListingKey []byte    `protobuf:"bytes,2,opt,name=listing_key,json=listingKey,proto3" json:"listing_key,omitempty"`
}

func (m *ListingExpireQueueEntry) Reset()         { *m = ListingExpireQueueEntry{} }
func (m *ListingExpireQueueEntry) String() string { return proto.CompactTextString(m) }
func (*ListingExpireQueueEntry) ProtoMessage()    {}
func (*ListingExpireQueueEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e368762f8e5e688e, []int{0}
}
func (m *ListingExpireQueueEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListingExpireQueueEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListingExpireQueueEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListingExpireQueueEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListingExpireQueueEntry.Merge(m, src)
}
func (m *ListingExpireQueueEntry) XXX_Size() int {
	return m.Size()
}
func (m *ListingExpireQueueEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ListingExpireQueueEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ListingExpireQueueEntry proto.InternalMessageInfo

func (m *ListingExpireQueueEntry) GetExpireTime() time.Time {
	if m != nil {
		return m.ExpireTime
	}
	return time.Time{}
}

func (m *ListingExpireQueueEntry) GetListingKey() []byte {
	if m != nil {
		return m.ListingKey
	}
	return nil
}

func init() {
	proto.RegisterType((*ListingExpireQueueEntry)(nil), "likechain.likenft.v1.ListingExpireQueueEntry")
}

func init() {
	proto.RegisterFile("likechain/likenft/v1/listing_expire_queue.proto", fileDescriptor_e368762f8e5e688e)
}

var fileDescriptor_e368762f8e5e688e = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x50, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x8d, 0x19, 0x10, 0x72, 0x98, 0xa2, 0x4a, 0x54, 0x19, 0x9c, 0x8a, 0xa9, 0x0b, 0xb6, 0x0a,
	0xea, 0x0f, 0x54, 0xca, 0x04, 0x12, 0xa2, 0x62, 0x62, 0xa9, 0x9a, 0xea, 0xd6, 0xb5, 0x9a, 0xc4,
	0x21, 0x71, 0xa2, 0x66, 0xe4, 0x0f, 0xfa, 0x59, 0x1d, 0x3b, 0x32, 0x01, 0x4a, 0x7e, 0x04, 0xf9,
	0x91, 0x6e, 0xc7, 0xbe, 0xe7, 0x71, 0xef, 0xc1, 0x2c, 0x15, 0x7b, 0xd8, 0xec, 0xd6, 0x22, 0x37,
	0x28, 0xdf, 0x2a, 0xd6, 0xcc, 0x58, 0x2a, 0x2a, 0x25, 0x72, 0xbe, 0x82, 0x43, 0x21, 0x4a, 0x58,
	0x7d, 0xd6, 0x50, 0x03, 0x2d, 0x4a, 0xa9, 0x64, 0x30, 0xba, 0x08, 0xa8, 0x13, 0xd0, 0x66, 0x16,
	0x8e, 0xb8, 0xe4, 0xd2, 0x10, 0x98, 0x46, 0x96, 0x1b, 0x46, 0x5c, 0x4a, 0x9e, 0x02, 0x33, 0xaf,
	0xa4, 0xde, 0x32, 0x25, 0x32, 0xa8, 0xd4, 0x3a, 0x2b, 0x2c, 0xe1, 0xfe, 0x0b, 0xe1, 0xbb, 0x17,
	0x9b, 0x15, 0x9b, 0xa8, 0x37, 0x9d, 0x14, 0xe7, 0xaa, 0x6c, 0x83, 0x18, 0xfb, 0x2e, 0x5e, 0xab,
	0xc6, 0x68, 0x82, 0xa6, 0xfe, 0x63, 0x48, 0xad, 0x25, 0x1d, 0x2c, 0xe9, 0xfb, 0x60, 0xb9, 0xb8,
	0x39, 0xfd, 0x44, 0xde, 0xf1, 0x37, 0x42, 0x4b, 0x6c, 0x85, 0x7a, 0x14, 0x44, 0xd8, 0x1f, 0xae,
	0xd9, 0x43, 0x3b, 0xbe, 0x9a, 0xa0, 0xe9, 0xed, 0x12, 0xbb, 0xaf, 0x67, 0x68, 0x17, 0xaf, 0xa7,
	0x8e, 0xa0, 0x73, 0x47, 0xd0, 0x5f, 0x47, 0xd0, 0xb1, 0x27, 0xde, 0xb9, 0x27, 0xde, 0x77, 0x4f,
	0xbc, 0x8f, 0x39, 0x17, 0x6a, 0x57, 0x27, 0x74, 0x23, 0x33, 0x5b, 0x93, 0x74, 0x2d, 0x69, 0xf0,
	0x60, 0x4b, 0x6b, 0xe6, 0xec, 0x70, 0x69, 0x4e, 0xb5, 0x05, 0x54, 0xc9, 0xb5, 0xd9, 0xed, 0xe9,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0xac, 0x79, 0x14, 0xc8, 0x5b, 0x01, 0x00, 0x00,
}

func (m *ListingExpireQueueEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListingExpireQueueEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListingExpireQueueEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ListingKey) > 0 {
		i -= len(m.ListingKey)
		copy(dAtA[i:], m.ListingKey)
		i = encodeVarintListingExpireQueue(dAtA, i, uint64(len(m.ListingKey)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpireTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintListingExpireQueue(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintListingExpireQueue(dAtA []byte, offset int, v uint64) int {
	offset -= sovListingExpireQueue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListingExpireQueueEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireTime)
	n += 1 + l + sovListingExpireQueue(uint64(l))
	l = len(m.ListingKey)
	if l > 0 {
		n += 1 + l + sovListingExpireQueue(uint64(l))
	}
	return n
}

func sovListingExpireQueue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListingExpireQueue(x uint64) (n int) {
	return sovListingExpireQueue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListingExpireQueueEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListingExpireQueue
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
			return fmt.Errorf("proto: ListingExpireQueueEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListingExpireQueueEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListingExpireQueue
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
				return ErrInvalidLengthListingExpireQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListingExpireQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpireTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListingKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListingExpireQueue
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
				return ErrInvalidLengthListingExpireQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListingExpireQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListingKey = append(m.ListingKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ListingKey == nil {
				m.ListingKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListingExpireQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListingExpireQueue
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
func skipListingExpireQueue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListingExpireQueue
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
					return 0, ErrIntOverflowListingExpireQueue
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
					return 0, ErrIntOverflowListingExpireQueue
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
				return 0, ErrInvalidLengthListingExpireQueue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListingExpireQueue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListingExpireQueue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListingExpireQueue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListingExpireQueue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListingExpireQueue = fmt.Errorf("proto: unexpected end of group")
)
