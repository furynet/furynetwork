// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/strategicreserve/exposure.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// OrderBookOddsExposure represents the exposures taken on odds.
type OrderBookOddsExposure struct {
	// order_book_uid is the universally unique identifier corresponding to the order book.
	OrderBookUID string `protobuf:"bytes,1,opt,name=order_book_uid,proto3" json:"order_book_uid"`
	// odds_uid is the universally unique identifier of the odds.
	OddsUID string `protobuf:"bytes,2,opt,name=odds_uid,proto3" json:"odds_uid"`
	// fulfillment_queue is the slice of indices of participations to be
	// fulfilled.
	FulfillmentQueue []uint64 `protobuf:"varint,3,rep,packed,name=fulfillment_queue,json=fulfillmentQueue,proto3" json:"fulfillment_queue,omitempty" yaml:"fulfillment_queue"`
}

func (m *OrderBookOddsExposure) Reset()      { *m = OrderBookOddsExposure{} }
func (*OrderBookOddsExposure) ProtoMessage() {}
func (*OrderBookOddsExposure) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d2f27f5a9b2729, []int{0}
}
func (m *OrderBookOddsExposure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderBookOddsExposure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBookOddsExposure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderBookOddsExposure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBookOddsExposure.Merge(m, src)
}
func (m *OrderBookOddsExposure) XXX_Size() int {
	return m.Size()
}
func (m *OrderBookOddsExposure) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBookOddsExposure.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBookOddsExposure proto.InternalMessageInfo

// ParticipationExposure represents the exposures taken on odds by
// participations.
type ParticipationExposure struct {
	// order_book_uid is the universally unique identifier of the order book that the
	// exposure is being set.
	OrderBookUID string `protobuf:"bytes,1,opt,name=order_book_uid,proto3" json:"order_book_uid"`
	// odds_uid is the odds universal unique identifier that the exposure is being
	// set.
	OddsUID string `protobuf:"bytes,2,opt,name=odds_uid,proto3" json:"odds_uid"`
	// participation_index is the index of participation in the queue.
	ParticipationIndex uint64 `protobuf:"varint,3,opt,name=participation_index,json=participationIndex,proto3" json:"participation_index,omitempty" yaml:"participation_index"`
	// exposure is the total exposure taken on given odds.
	Exposure github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=exposure,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"exposure" yaml:"exposure"`
	// bet_amount is the total bet amount corresponding to the exposure.
	BetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=bet_amount,json=betAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_amount" yaml:"bet_amount"`
	// is_fulfilled represents if the participation exposure is fulfilled or not.
	IsFulfilled bool `protobuf:"varint,6,opt,name=is_fulfilled,json=isFulfilled,proto3" json:"is_fulfilled,omitempty" yaml:"is_fulfilled"`
	// round is the current round number in the queue.
	Round uint64 `protobuf:"varint,7,opt,name=round,proto3" json:"round,omitempty" yaml:"round"`
}

func (m *ParticipationExposure) Reset()      { *m = ParticipationExposure{} }
func (*ParticipationExposure) ProtoMessage() {}
func (*ParticipationExposure) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d2f27f5a9b2729, []int{1}
}
func (m *ParticipationExposure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParticipationExposure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParticipationExposure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParticipationExposure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParticipationExposure.Merge(m, src)
}
func (m *ParticipationExposure) XXX_Size() int {
	return m.Size()
}
func (m *ParticipationExposure) XXX_DiscardUnknown() {
	xxx_messageInfo_ParticipationExposure.DiscardUnknown(m)
}

var xxx_messageInfo_ParticipationExposure proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OrderBookOddsExposure)(nil), "furynet.furynetwork.strategicreserve.OrderBookOddsExposure")
	proto.RegisterType((*ParticipationExposure)(nil), "furynet.furynetwork.strategicreserve.ParticipationExposure")
}

func init() {
	proto.RegisterFile("fury/strategicreserve/exposure.proto", fileDescriptor_90d2f27f5a9b2729)
}

var fileDescriptor_90d2f27f5a9b2729 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x49, 0xda, 0xa6, 0x47, 0x04, 0xad, 0x4b, 0x55, 0xab, 0x42, 0xbe, 0xe8, 0x90, 0xaa,
	0x2c, 0x8d, 0x87, 0x6c, 0xd9, 0x6a, 0x7e, 0x48, 0x41, 0x42, 0x01, 0x4b, 0x2c, 0x48, 0xc8, 0x8a,
	0x73, 0xaf, 0xe6, 0x94, 0xc4, 0x17, 0xee, 0xce, 0x90, 0xfe, 0x07, 0x8c, 0x8c, 0x4c, 0x28, 0x7f,
	0x4e, 0xc7, 0x8e, 0x88, 0xe1, 0x84, 0x92, 0x05, 0xb1, 0x91, 0xbf, 0x00, 0xf9, 0x9c, 0x84, 0x34,
	0x65, 0x61, 0xeb, 0xe4, 0x77, 0xdf, 0xfb, 0xde, 0xf7, 0xdd, 0xf3, 0xbd, 0x87, 0x1e, 0xc9, 0x04,
	0x7c, 0xa9, 0x44, 0x57, 0x41, 0xc2, 0x7a, 0x02, 0x24, 0x88, 0x0f, 0xe0, 0xc3, 0x78, 0xc4, 0x65,
	0x26, 0xa0, 0x31, 0x12, 0x5c, 0x71, 0x07, 0xcb, 0x04, 0x52, 0x50, 0x1f, 0xb9, 0xe8, 0x37, 0x64,
	0x02, 0x8d, 0x4d, 0xfe, 0xf1, 0x83, 0x84, 0x27, 0xdc, 0x70, 0xfd, 0x3c, 0x2a, 0xca, 0xc8, 0x6f,
	0x1b, 0x1d, 0x76, 0x04, 0x05, 0x11, 0x70, 0xde, 0xef, 0x50, 0x2a, 0x9f, 0x2e, 0x64, 0x9d, 0xe7,
	0xe8, 0x1e, 0xcf, 0x13, 0x51, 0xcc, 0x79, 0x3f, 0xca, 0x18, 0x75, 0xed, 0x9a, 0x5d, 0xdf, 0x0d,
	0xc8, 0x54, 0xe3, 0xea, 0xaa, 0xe4, 0x75, 0xfb, 0xc9, 0x2f, 0x8d, 0x37, 0x98, 0xe1, 0xc6, 0xd9,
	0x69, 0xa2, 0x0a, 0xa7, 0x54, 0x1a, 0x95, 0x3b, 0x46, 0xe5, 0x68, 0xaa, 0xf1, 0x4e, 0xee, 0x57,
	0x08, 0xac, 0xd2, 0xe1, 0x2a, 0x72, 0xda, 0x68, 0xff, 0x3c, 0x1b, 0x9c, 0xb3, 0xc1, 0x60, 0x08,
	0xa9, 0x8a, 0xde, 0x67, 0x90, 0x81, 0x5b, 0xaa, 0x95, 0xea, 0xe5, 0xe0, 0xe1, 0x5c, 0x63, 0xf7,
	0xa2, 0x3b, 0x1c, 0xb4, 0xc8, 0x0d, 0x0a, 0x09, 0xf7, 0xd6, 0xb0, 0x57, 0x39, 0xd4, 0xaa, 0x7e,
	0x9a, 0x60, 0xeb, 0xcb, 0x04, 0x5b, 0x3f, 0x27, 0xd8, 0x22, 0x5f, 0xcb, 0xe8, 0xf0, 0x65, 0x57,
	0x28, 0xd6, 0x63, 0xa3, 0xae, 0x62, 0x3c, 0xbd, 0x3d, 0x3d, 0x77, 0xd0, 0xc1, 0x68, 0xfd, 0x66,
	0x11, 0x4b, 0x29, 0x8c, 0xdd, 0x52, 0xcd, 0xae, 0x97, 0x03, 0x6f, 0xae, 0xf1, 0x71, 0xd1, 0xf5,
	0x3f, 0x48, 0x24, 0x74, 0xae, 0xa1, 0xed, 0x1c, 0x74, 0xde, 0xa2, 0xca, 0x72, 0x50, 0xdc, 0xb2,
	0xb9, 0xc5, 0xd9, 0xa5, 0xc6, 0xd6, 0x77, 0x8d, 0x4f, 0x12, 0xa6, 0xde, 0x65, 0x71, 0xa3, 0xc7,
	0x87, 0x7e, 0x8f, 0xcb, 0x21, 0x97, 0x8b, 0xcf, 0xa9, 0xa4, 0x7d, 0x5f, 0x5d, 0x8c, 0x40, 0x36,
	0xda, 0xa9, 0x9a, 0x6b, 0x7c, 0xbf, 0xf0, 0x5c, 0xea, 0x90, 0x70, 0x25, 0xe9, 0xc4, 0x08, 0xc5,
	0xa0, 0xa2, 0xee, 0x90, 0x67, 0xa9, 0x72, 0xb7, 0x8c, 0xc1, 0xe3, 0xff, 0x36, 0xd8, 0x2f, 0x0c,
	0xfe, 0x2a, 0x91, 0x70, 0x37, 0x06, 0x75, 0x66, 0x62, 0xa7, 0x85, 0xaa, 0x4c, 0x46, 0x8b, 0x37,
	0x05, 0xea, 0x6e, 0xd7, 0xec, 0x7a, 0x25, 0x38, 0x9a, 0x6b, 0x7c, 0x50, 0xd4, 0xad, 0x67, 0x49,
	0x78, 0x97, 0xc9, 0x67, 0xcb, 0x93, 0x73, 0x82, 0xb6, 0x04, 0xcf, 0x52, 0xea, 0xee, 0x98, 0x3f,
	0xb8, 0x37, 0xd7, 0xb8, 0x5a, 0x14, 0x19, 0x98, 0x84, 0x45, 0xfa, 0xfa, 0x80, 0x04, 0x2f, 0x2e,
	0xa7, 0x9e, 0x7d, 0x35, 0xf5, 0xec, 0x1f, 0x53, 0xcf, 0xfe, 0x3c, 0xf3, 0xac, 0xab, 0x99, 0x67,
	0x7d, 0x9b, 0x79, 0xd6, 0x9b, 0xe6, 0x5a, 0x4f, 0x32, 0x81, 0xd3, 0xc5, 0xc6, 0xe5, 0xb1, 0x3f,
	0xbe, 0xb9, 0xa3, 0xa6, 0xc9, 0x78, 0xdb, 0xac, 0x5a, 0xf3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x16, 0x65, 0xe3, 0x51, 0xc8, 0x03, 0x00, 0x00,
}

func (m *OrderBookOddsExposure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBookOddsExposure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBookOddsExposure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FulfillmentQueue) > 0 {
		dAtA2 := make([]byte, len(m.FulfillmentQueue)*10)
		var j1 int
		for _, num := range m.FulfillmentQueue {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintExposure(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OddsUID) > 0 {
		i -= len(m.OddsUID)
		copy(dAtA[i:], m.OddsUID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.OddsUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OrderBookUID) > 0 {
		i -= len(m.OrderBookUID)
		copy(dAtA[i:], m.OrderBookUID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.OrderBookUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ParticipationExposure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParticipationExposure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParticipationExposure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Round != 0 {
		i = encodeVarintExposure(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x38
	}
	if m.IsFulfilled {
		i--
		if m.IsFulfilled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.BetAmount.Size()
		i -= size
		if _, err := m.BetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExposure(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Exposure.Size()
		i -= size
		if _, err := m.Exposure.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExposure(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.ParticipationIndex != 0 {
		i = encodeVarintExposure(dAtA, i, uint64(m.ParticipationIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OddsUID) > 0 {
		i -= len(m.OddsUID)
		copy(dAtA[i:], m.OddsUID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.OddsUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OrderBookUID) > 0 {
		i -= len(m.OrderBookUID)
		copy(dAtA[i:], m.OrderBookUID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.OrderBookUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExposure(dAtA []byte, offset int, v uint64) int {
	offset -= sovExposure(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderBookOddsExposure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OrderBookUID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	l = len(m.OddsUID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	if len(m.FulfillmentQueue) > 0 {
		l = 0
		for _, e := range m.FulfillmentQueue {
			l += sovExposure(uint64(e))
		}
		n += 1 + sovExposure(uint64(l)) + l
	}
	return n
}

func (m *ParticipationExposure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OrderBookUID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	l = len(m.OddsUID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	if m.ParticipationIndex != 0 {
		n += 1 + sovExposure(uint64(m.ParticipationIndex))
	}
	l = m.Exposure.Size()
	n += 1 + l + sovExposure(uint64(l))
	l = m.BetAmount.Size()
	n += 1 + l + sovExposure(uint64(l))
	if m.IsFulfilled {
		n += 2
	}
	if m.Round != 0 {
		n += 1 + sovExposure(uint64(m.Round))
	}
	return n
}

func sovExposure(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExposure(x uint64) (n int) {
	return sovExposure(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderBookOddsExposure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExposure
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
			return fmt.Errorf("proto: OrderBookOddsExposure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBookOddsExposure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OddsUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExposure
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FulfillmentQueue = append(m.FulfillmentQueue, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExposure
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthExposure
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthExposure
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.FulfillmentQueue) == 0 {
					m.FulfillmentQueue = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExposure
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FulfillmentQueue = append(m.FulfillmentQueue, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FulfillmentQueue", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExposure(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExposure
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
func (m *ParticipationExposure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExposure
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
			return fmt.Errorf("proto: ParticipationExposure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParticipationExposure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OddsUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationIndex", wireType)
			}
			m.ParticipationIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exposure", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Exposure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsFulfilled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
			m.IsFulfilled = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExposure(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExposure
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
func skipExposure(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExposure
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
					return 0, ErrIntOverflowExposure
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
					return 0, ErrIntOverflowExposure
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
				return 0, ErrInvalidLengthExposure
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExposure
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExposure
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExposure        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExposure          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExposure = fmt.Errorf("proto: unexpected end of group")
)
