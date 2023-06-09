// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/strategicreserve/orderbook.proto

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

// OrderBookStatus is the enum type for the status of the order book.
type OrderBookStatus int32

const (
	// invalid
	OrderBookStatus_ORDER_BOOK_STATUS_UNSPECIFIED OrderBookStatus = 0
	// active
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_ACTIVE OrderBookStatus = 1
	// resolved not settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_RESOLVED OrderBookStatus = 2
	// resolved and settled
	OrderBookStatus_ORDER_BOOK_STATUS_STATUS_SETTLED OrderBookStatus = 3
)

var OrderBookStatus_name = map[int32]string{
	0: "ORDER_BOOK_STATUS_UNSPECIFIED",
	1: "ORDER_BOOK_STATUS_STATUS_ACTIVE",
	2: "ORDER_BOOK_STATUS_STATUS_RESOLVED",
	3: "ORDER_BOOK_STATUS_STATUS_SETTLED",
}

var OrderBookStatus_value = map[string]int32{
	"ORDER_BOOK_STATUS_UNSPECIFIED":     0,
	"ORDER_BOOK_STATUS_STATUS_ACTIVE":   1,
	"ORDER_BOOK_STATUS_STATUS_RESOLVED": 2,
	"ORDER_BOOK_STATUS_STATUS_SETTLED":  3,
}

func (x OrderBookStatus) String() string {
	return proto.EnumName(OrderBookStatus_name, int32(x))
}

func (OrderBookStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dab097ebfa6cba07, []int{0}
}

// OrderBook represents the order book maintained against a market.
type OrderBook struct {
	// uid is the universal unique identifier of the order book.
	UID string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	// participation_count is the count of participations in the order book.
	ParticipationCount uint64 `protobuf:"varint,2,opt,name=participation_count,json=participationCount,proto3" json:"participation_count,omitempty" yaml:"participation_count"`
	// odds_count is the count of the odds in the order book.
	OddsCount uint64 `protobuf:"varint,3,opt,name=odds_count,json=oddsCount,proto3" json:"odds_count,omitempty" yaml:"odds_count"`
	// status represents the status of the order book.
	Status OrderBookStatus `protobuf:"varint,4,opt,name=status,proto3,enum=furynet.furynetwork.strategicreserve.OrderBookStatus" json:"status,omitempty"`
}

func (m *OrderBook) Reset()      { *m = OrderBook{} }
func (*OrderBook) ProtoMessage() {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab097ebfa6cba07, []int{0}
}
func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(m, src)
}
func (m *OrderBook) XXX_Size() int {
	return m.Size()
}
func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("furynet.furynetwork.strategicreserve.OrderBookStatus", OrderBookStatus_name, OrderBookStatus_value)
	proto.RegisterType((*OrderBook)(nil), "furynet.furynetwork.strategicreserve.OrderBook")
}

func init() {
	proto.RegisterFile("fury/strategicreserve/orderbook.proto", fileDescriptor_dab097ebfa6cba07)
}

var fileDescriptor_dab097ebfa6cba07 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x2b, 0x2d, 0xaa,
	0xd4, 0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x49, 0x4d, 0xcf, 0x4c, 0x2e, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a,
	0x4b, 0xd5, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x4a, 0xca, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x52, 0x01, 0x29, 0xcb, 0x4b, 0x2d, 0xd1, 0x83, 0xd2, 0xe5, 0xf9, 0x45, 0xd9, 0x7a,
	0xe8, 0xba, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x1a, 0xf4, 0x41, 0x2c, 0x88, 0x5e, 0xa5,
	0x89, 0x4c, 0x5c, 0x9c, 0xfe, 0x20, 0xf3, 0x9c, 0xf2, 0xf3, 0xb3, 0x85, 0x14, 0xb8, 0x98, 0x4b,
	0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xf8, 0x1e, 0xdd, 0x93, 0x67, 0x0e, 0xf5,
	0x74, 0x79, 0x75, 0x4f, 0x1e, 0x24, 0x1a, 0x04, 0x22, 0x84, 0xfc, 0xb9, 0x84, 0x0b, 0x12, 0x8b,
	0x4a, 0x32, 0x93, 0x33, 0x0b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xe2, 0x93, 0xf3, 0x4b, 0xf3, 0x4a,
	0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x9c, 0xe4, 0x3e, 0xdd, 0x93, 0x97, 0xaa, 0x4c, 0xcc, 0xcd,
	0xb1, 0x52, 0xc2, 0xa2, 0x48, 0x29, 0x48, 0x08, 0x45, 0xd4, 0x19, 0x24, 0x28, 0x64, 0xc2, 0xc5,
	0x95, 0x9f, 0x92, 0x52, 0x0c, 0x35, 0x87, 0x19, 0x6c, 0x8e, 0xe8, 0xa7, 0x7b, 0xf2, 0x82, 0x10,
	0x73, 0x10, 0x72, 0x4a, 0x41, 0x9c, 0x20, 0x0e, 0x44, 0x97, 0x2f, 0x17, 0x5b, 0x71, 0x49, 0x62,
	0x49, 0x69, 0xb1, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0xa9, 0x1e, 0x31, 0x61, 0xa0, 0x07,
	0xf7, 0x69, 0x30, 0x58, 0x73, 0x10, 0xd4, 0x10, 0x2b, 0x9e, 0x8e, 0x05, 0xf2, 0x0c, 0x33, 0x16,
	0xc8, 0x33, 0xbc, 0x58, 0x20, 0xcf, 0xa0, 0xb5, 0x8c, 0x91, 0x8b, 0x1f, 0x4d, 0xa5, 0x90, 0x22,
	0x97, 0xac, 0x7f, 0x90, 0x8b, 0x6b, 0x50, 0xbc, 0x93, 0xbf, 0xbf, 0x77, 0x7c, 0x70, 0x88, 0x63,
	0x48, 0x68, 0x70, 0x7c, 0xa8, 0x5f, 0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00,
	0x83, 0x90, 0x32, 0x97, 0x3c, 0xa6, 0x12, 0x28, 0xe5, 0xe8, 0x1c, 0xe2, 0x19, 0xe6, 0x2a, 0xc0,
	0x28, 0xa4, 0xca, 0xa5, 0x88, 0x53, 0x51, 0x90, 0x6b, 0xb0, 0xbf, 0x4f, 0x98, 0xab, 0x8b, 0x00,
	0x93, 0x90, 0x0a, 0x97, 0x02, 0x4e, 0x65, 0xc1, 0xae, 0x21, 0x21, 0x3e, 0xae, 0x2e, 0x02, 0xcc,
	0x4e, 0x81, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9e, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x0d, 0x11, 0x7d, 0xa4, 0x90, 0xd1, 0xaf,
	0xc0, 0x4c, 0x55, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x64, 0x61, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xf2, 0x9a, 0xc7, 0xb8, 0x7b, 0x02, 0x00, 0x00,
}

func (m *OrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.OddsCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.OddsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.ParticipationCount != 0 {
		i = encodeVarintOrderbook(dAtA, i, uint64(m.ParticipationCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintOrderbook(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrderbook(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderbook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovOrderbook(uint64(l))
	}
	if m.ParticipationCount != 0 {
		n += 1 + sovOrderbook(uint64(m.ParticipationCount))
	}
	if m.OddsCount != 0 {
		n += 1 + sovOrderbook(uint64(m.OddsCount))
	}
	if m.Status != 0 {
		n += 1 + sovOrderbook(uint64(m.Status))
	}
	return n
}

func sovOrderbook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderbook(x uint64) (n int) {
	return sovOrderbook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderbook
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
			return fmt.Errorf("proto: OrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
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
				return ErrInvalidLengthOrderbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationCount", wireType)
			}
			m.ParticipationCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsCount", wireType)
			}
			m.OddsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OddsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OrderBookStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrderbook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderbook
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
func skipOrderbook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
					return 0, ErrIntOverflowOrderbook
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
				return 0, ErrInvalidLengthOrderbook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderbook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderbook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderbook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderbook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderbook = fmt.Errorf("proto: unexpected end of group")
)
