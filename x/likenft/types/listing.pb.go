// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/v1/listing.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Listing struct {
	ClassId          string    `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	NftId            string    `protobuf:"bytes,2,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	Seller           string    `protobuf:"bytes,3,opt,name=seller,proto3" json:"seller,omitempty"`
	Price            uint64    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Expiration       time.Time `protobuf:"bytes,5,opt,name=expiration,proto3,stdtime" json:"expiration"`
	FullPayToRoyalty bool      `protobuf:"varint,6,opt,name=full_pay_to_royalty,json=fullPayToRoyalty,proto3" json:"full_pay_to_royalty,omitempty"`
}

func (m *Listing) Reset()         { *m = Listing{} }
func (m *Listing) String() string { return proto.CompactTextString(m) }
func (*Listing) ProtoMessage()    {}
func (*Listing) Descriptor() ([]byte, []int) {
	return fileDescriptor_592867f987c9f178, []int{0}
}
func (m *Listing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Listing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Listing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Listing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listing.Merge(m, src)
}
func (m *Listing) XXX_Size() int {
	return m.Size()
}
func (m *Listing) XXX_DiscardUnknown() {
	xxx_messageInfo_Listing.DiscardUnknown(m)
}

var xxx_messageInfo_Listing proto.InternalMessageInfo

func (m *Listing) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *Listing) GetNftId() string {
	if m != nil {
		return m.NftId
	}
	return ""
}

func (m *Listing) GetSeller() string {
	if m != nil {
		return m.Seller
	}
	return ""
}

func (m *Listing) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Listing) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

func (m *Listing) GetFullPayToRoyalty() bool {
	if m != nil {
		return m.FullPayToRoyalty
	}
	return false
}

type ListingStoreRecord struct {
	ClassId          string                                        `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	NftId            string                                        `protobuf:"bytes,2,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	Seller           github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=seller,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"seller,omitempty"`
	Price            uint64                                        `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Expiration       time.Time                                     `protobuf:"bytes,5,opt,name=expiration,proto3,stdtime" json:"expiration"`
	FullPayToRoyalty bool                                          `protobuf:"varint,6,opt,name=full_pay_to_royalty,json=fullPayToRoyalty,proto3" json:"full_pay_to_royalty,omitempty"`
}

func (m *ListingStoreRecord) Reset()         { *m = ListingStoreRecord{} }
func (m *ListingStoreRecord) String() string { return proto.CompactTextString(m) }
func (*ListingStoreRecord) ProtoMessage()    {}
func (*ListingStoreRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_592867f987c9f178, []int{1}
}
func (m *ListingStoreRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListingStoreRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListingStoreRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListingStoreRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListingStoreRecord.Merge(m, src)
}
func (m *ListingStoreRecord) XXX_Size() int {
	return m.Size()
}
func (m *ListingStoreRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ListingStoreRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ListingStoreRecord proto.InternalMessageInfo

func (m *ListingStoreRecord) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *ListingStoreRecord) GetNftId() string {
	if m != nil {
		return m.NftId
	}
	return ""
}

func (m *ListingStoreRecord) GetSeller() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Seller
	}
	return nil
}

func (m *ListingStoreRecord) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ListingStoreRecord) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

func (m *ListingStoreRecord) GetFullPayToRoyalty() bool {
	if m != nil {
		return m.FullPayToRoyalty
	}
	return false
}

func init() {
	proto.RegisterType((*Listing)(nil), "likechain.likenft.v1.Listing")
	proto.RegisterType((*ListingStoreRecord)(nil), "likechain.likenft.v1.ListingStoreRecord")
}

func init() {
	proto.RegisterFile("likechain/likenft/v1/listing.proto", fileDescriptor_592867f987c9f178)
}

var fileDescriptor_592867f987c9f178 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0xb1, 0x6e, 0xdb, 0x30,
	0x14, 0x34, 0xd3, 0xd8, 0x71, 0xd9, 0x0e, 0x05, 0xeb, 0x16, 0xaa, 0x07, 0x49, 0xf0, 0xa4, 0x45,
	0x24, 0xdc, 0x22, 0x1f, 0x10, 0xa3, 0x8b, 0x81, 0x02, 0x2d, 0xd4, 0x4c, 0x5d, 0x04, 0x99, 0xa2,
	0x14, 0x22, 0x94, 0x9e, 0x40, 0xd2, 0x46, 0xf4, 0x17, 0xf9, 0x80, 0x7e, 0x50, 0xc6, 0x8c, 0x9d,
	0xdc, 0xc2, 0xfe, 0x8b, 0x4e, 0x85, 0x28, 0x25, 0x48, 0xd7, 0x4e, 0x9d, 0xf8, 0xee, 0xdd, 0x91,
	0xb8, 0x23, 0x0e, 0x2f, 0x94, 0xbc, 0x16, 0xfc, 0x2a, 0x93, 0x35, 0xeb, 0xa6, 0xba, 0xb0, 0x6c,
	0xb7, 0x64, 0x4a, 0x1a, 0x2b, 0xeb, 0x92, 0x36, 0x1a, 0x2c, 0x90, 0xd9, 0xa3, 0x86, 0x0e, 0x1a,
	0xba, 0x5b, 0xce, 0x67, 0x25, 0x94, 0xe0, 0x04, 0xac, 0x9b, 0x7a, 0xed, 0x3c, 0x28, 0x01, 0x4a,
	0x25, 0x98, 0x43, 0x9b, 0x6d, 0xc1, 0xac, 0xac, 0x84, 0xb1, 0x59, 0xd5, 0xf4, 0x82, 0xc5, 0x1e,
	0xe1, 0xb3, 0x4f, 0xfd, 0xf3, 0xe4, 0x1d, 0x9e, 0x72, 0x95, 0x19, 0x93, 0xca, 0xdc, 0x43, 0x21,
	0x8a, 0x9e, 0x27, 0x67, 0x0e, 0xaf, 0x73, 0xf2, 0x06, 0x4f, 0xea, 0xc2, 0x76, 0xc4, 0x89, 0x23,
	0xc6, 0x75, 0x61, 0xd7, 0x39, 0x79, 0x8b, 0x27, 0x46, 0x28, 0x25, 0xb4, 0xf7, 0xcc, 0xad, 0x07,
	0x44, 0x66, 0x78, 0xdc, 0x68, 0xc9, 0x85, 0x77, 0x1a, 0xa2, 0xe8, 0x34, 0xe9, 0x01, 0xf9, 0x88,
	0xb1, 0xb8, 0x69, 0xa4, 0xce, 0xac, 0x84, 0xda, 0x1b, 0x87, 0x28, 0x7a, 0xf1, 0x7e, 0x4e, 0x7b,
	0x87, 0xf4, 0xc1, 0x21, 0xbd, 0x7c, 0x70, 0xb8, 0x9a, 0xde, 0xed, 0x83, 0xd1, 0xed, 0xcf, 0x00,
	0x25, 0x4f, 0xee, 0x91, 0x18, 0xbf, 0x2e, 0xb6, 0x4a, 0xa5, 0x4d, 0xd6, 0xa6, 0x16, 0x52, 0x0d,
	0x6d, 0xa6, 0x6c, 0xeb, 0x4d, 0x42, 0x14, 0x4d, 0x93, 0x57, 0x1d, 0xf5, 0x25, 0x6b, 0x2f, 0x21,
	0xe9, 0xf7, 0x8b, 0xef, 0x27, 0x98, 0x0c, 0x01, 0xbf, 0x5a, 0xd0, 0x22, 0x11, 0x1c, 0x74, 0xfe,
	0x0f, 0x59, 0xd7, 0x7f, 0x65, 0x7d, 0xb9, 0x5a, 0xfe, 0xde, 0x07, 0x71, 0x29, 0xed, 0xd5, 0x76,
	0x43, 0x39, 0x54, 0x8c, 0x83, 0xa9, 0xc0, 0x0c, 0x47, 0x6c, 0xf2, 0x6b, 0x66, 0xdb, 0x46, 0x18,
	0x7a, 0xc1, 0xf9, 0x45, 0x9e, 0x6b, 0x61, 0xcc, 0x7f, 0xf8, 0x3d, 0xab, 0xcf, 0x77, 0x07, 0x1f,
	0xdd, 0x1f, 0x7c, 0xf4, 0xeb, 0xe0, 0xa3, 0xdb, 0xa3, 0x3f, 0xba, 0x3f, 0xfa, 0xa3, 0x1f, 0x47,
	0x7f, 0xf4, 0xed, 0xfc, 0x49, 0x36, 0xd7, 0x38, 0x18, 0x4a, 0xd9, 0x0d, 0x71, 0xdf, 0xd1, 0xdd,
	0x39, 0xbb, 0x79, 0x2c, 0xaa, 0x8b, 0xbb, 0x99, 0x38, 0xa7, 0x1f, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x66, 0x05, 0x76, 0xc2, 0xca, 0x02, 0x00, 0x00,
}

func (m *Listing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Listing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FullPayToRoyalty {
		i--
		if m.FullPayToRoyalty {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintListing(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.Price != 0 {
		i = encodeVarintListing(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintListing(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintListing(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListingStoreRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListingStoreRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListingStoreRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FullPayToRoyalty {
		i--
		if m.FullPayToRoyalty {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintListing(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if m.Price != 0 {
		i = encodeVarintListing(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintListing(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintListing(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintListing(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintListing(dAtA []byte, offset int, v uint64) int {
	offset -= sovListing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Listing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovListing(uint64(m.Price))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovListing(uint64(l))
	if m.FullPayToRoyalty {
		n += 2
	}
	return n
}

func (m *ListingStoreRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovListing(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovListing(uint64(m.Price))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovListing(uint64(l))
	if m.FullPayToRoyalty {
		n += 2
	}
	return n
}

func sovListing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListing(x uint64) (n int) {
	return sovListing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Listing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListing
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
			return fmt.Errorf("proto: Listing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Listing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullPayToRoyalty", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FullPayToRoyalty = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipListing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListing
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
func (m *ListingStoreRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListing
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
			return fmt.Errorf("proto: ListingStoreRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListingStoreRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = append(m.Seller[:0], dAtA[iNdEx:postIndex]...)
			if m.Seller == nil {
				m.Seller = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
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
				return ErrInvalidLengthListing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullPayToRoyalty", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FullPayToRoyalty = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipListing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListing
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
func skipListing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListing
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
					return 0, ErrIntOverflowListing
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
					return 0, ErrIntOverflowListing
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
				return 0, ErrInvalidLengthListing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListing = fmt.Errorf("proto: unexpected end of group")
)
