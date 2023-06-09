// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/strategicreserve/participation.proto

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

// OrderBookParticipation represents the participants of an order book.
type OrderBookParticipation struct {
	// index is the index of the participation in the participation queue.
	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	// order_book_uid is the unique identifier corresponding to the order book.
	OrderBookUID string `protobuf:"bytes,2,opt,name=order_book_uid,proto3" json:"order_book_uid"`
	// participant_address is the bech32-encoded address of the participant.
	ParticipantAddress string `protobuf:"bytes,3,opt,name=participant_address,json=participantAddress,proto3" json:"participant_address,omitempty" yaml:"participant_address"`
	// is_module_account represents if the participant is a module account or not.
	IsModuleAccount bool `protobuf:"varint,4,opt,name=is_module_account,json=isModuleAccount,proto3" json:"is_module_account,omitempty" yaml:"is_module_account"`
	// liquidity is the total initial liquidity provided.
	Liquidity github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=liquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity" yaml:"liquidity"`
	// current_round_liquidity is the liquidity provided for the current round.
	CurrentRoundLiquidity github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=current_round_liquidity,json=currentRoundLiquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_round_liquidity" yaml:"current_round_liquidity"`
	// exposures_not_filled reresents if all of the exposures of the participation
	// are filled or not.
	ExposuresNotFilled uint64 `protobuf:"varint,7,opt,name=exposures_not_filled,json=exposuresNotFilled,proto3" json:"exposures_not_filled,omitempty" yaml:"exposures_not_filled"`
	// total_bet_amount is the total bet amount corresponding to all exposures.
	TotalBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=total_bet_amount,json=totalBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_bet_amount" yaml:"total_bet_amount"`
	// current_round_total_bet_amount is the total bet amount corresponding to all
	// exposures in the current round.
	CurrentRoundTotalBetAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=current_round_total_bet_amount,json=currentRoundTotalBetAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_round_total_bet_amount" yaml:"current_round_total_bet_amount"`
	// max_loss is the total bet amount corresponding to all exposure.
	MaxLoss github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=max_loss,json=maxLoss,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_loss" yaml:"max_loss"`
	// current_round_max_loss is the current round max loss.
	CurrentRoundMaxLoss github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=current_round_max_loss,json=currentRoundMaxLoss,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_round_max_loss" yaml:"current_round_max_loss"`
	// current_round_max_loss_odds_uid is the total max loss corresponding to
	// all exposures.
	CurrentRoundMaxLossOddsUID string `protobuf:"bytes,12,opt,name=current_round_max_loss_odds_uid,proto3" json:"current_round_max_loss_odds_uid" yaml:"current_round_max_loss_odds_uid"`
	// actual_profit is the actual profit of the participation fulfillment.
	ActualProfit github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=actual_profit,json=actualProfit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"actual_profit" yaml:"actual_profit"`
	// is_settled represents if the participation is settled or not.
	IsSettled bool `protobuf:"varint,14,opt,name=is_settled,json=isSettled,proto3" json:"is_settled,omitempty" yaml:"is_settled"`
}

func (m *OrderBookParticipation) Reset()      { *m = OrderBookParticipation{} }
func (*OrderBookParticipation) ProtoMessage() {}
func (*OrderBookParticipation) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a5eb982384ce13, []int{0}
}
func (m *OrderBookParticipation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderBookParticipation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBookParticipation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderBookParticipation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBookParticipation.Merge(m, src)
}
func (m *OrderBookParticipation) XXX_Size() int {
	return m.Size()
}
func (m *OrderBookParticipation) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBookParticipation.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBookParticipation proto.InternalMessageInfo

// ParticipationBetPair represents the book participation and bet bond.
type ParticipationBetPair struct {
	// order_book_uid is the unique identifier corresponding to the order book
	OrderBookUID string `protobuf:"bytes,1,opt,name=order_book_uid,proto3" json:"order_book_uid"`
	// participation_index is the index of participation corresponding to the bet
	// fulfillment.
	ParticipationIndex uint64 `protobuf:"varint,2,opt,name=participation_index,json=participationIndex,proto3" json:"participation_index,omitempty" yaml:"participation_index"`
	// bet_uid is the bet universal unique identifier of the bet that is
	// fulfilled.
	BetUID string `protobuf:"bytes,3,opt,name=bet_uid,proto3" json:"bet_uid"`
}

func (m *ParticipationBetPair) Reset()         { *m = ParticipationBetPair{} }
func (m *ParticipationBetPair) String() string { return proto.CompactTextString(m) }
func (*ParticipationBetPair) ProtoMessage()    {}
func (*ParticipationBetPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a5eb982384ce13, []int{1}
}
func (m *ParticipationBetPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParticipationBetPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParticipationBetPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParticipationBetPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParticipationBetPair.Merge(m, src)
}
func (m *ParticipationBetPair) XXX_Size() int {
	return m.Size()
}
func (m *ParticipationBetPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ParticipationBetPair.DiscardUnknown(m)
}

var xxx_messageInfo_ParticipationBetPair proto.InternalMessageInfo

func (m *ParticipationBetPair) GetOrderBookUID() string {
	if m != nil {
		return m.OrderBookUID
	}
	return ""
}

func (m *ParticipationBetPair) GetParticipationIndex() uint64 {
	if m != nil {
		return m.ParticipationIndex
	}
	return 0
}

func (m *ParticipationBetPair) GetBetUID() string {
	if m != nil {
		return m.BetUID
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderBookParticipation)(nil), "furynet.furynetwork.strategicreserve.OrderBookParticipation")
	proto.RegisterType((*ParticipationBetPair)(nil), "furynet.furynetwork.strategicreserve.ParticipationBetPair")
}

func init() {
	proto.RegisterFile("fury/strategicreserve/participation.proto", fileDescriptor_50a5eb982384ce13)
}

var fileDescriptor_50a5eb982384ce13 = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6e, 0xe3, 0x36,
	0x14, 0xb5, 0xd2, 0xbc, 0xcc, 0x38, 0x2f, 0xc5, 0x49, 0x04, 0xb7, 0x15, 0x03, 0xa1, 0x0d, 0xd2,
	0x45, 0xad, 0x45, 0x0b, 0x14, 0xc8, 0xce, 0x6a, 0x11, 0xd4, 0x45, 0x52, 0x3b, 0x6a, 0x8b, 0x02,
	0x45, 0x01, 0x56, 0x96, 0x18, 0x97, 0xb0, 0x2c, 0xba, 0x24, 0xd5, 0xda, 0xfb, 0x2e, 0xb2, 0x9c,
	0xe5, 0x00, 0xb3, 0xc9, 0x4f, 0xcc, 0x3f, 0x64, 0x99, 0xe5, 0x60, 0x16, 0xc2, 0xc0, 0x99, 0xc5,
	0x60, 0x96, 0xfe, 0x82, 0x81, 0x48, 0x47, 0x7e, 0xce, 0x04, 0x06, 0x66, 0x65, 0xf2, 0xde, 0xcb,
	0x73, 0xce, 0x25, 0xad, 0x73, 0xc1, 0x57, 0x57, 0x31, 0xeb, 0xd9, 0x5c, 0x30, 0x4f, 0xe0, 0x26,
	0xf1, 0x19, 0xe6, 0x98, 0xfd, 0x8b, 0xed, 0x8e, 0xc7, 0x04, 0xf1, 0x49, 0xc7, 0x13, 0x84, 0x46,
	0xe5, 0x0e, 0xa3, 0x82, 0xea, 0x5f, 0xa4, 0xa5, 0x11, 0x16, 0xe5, 0xe1, 0xef, 0x7f, 0x94, 0xb5,
	0xca, 0xd3, 0x27, 0x4b, 0xc5, 0x26, 0x6d, 0x52, 0x79, 0xc0, 0x4e, 0x57, 0xea, 0xac, 0x75, 0xbb,
	0x01, 0x0e, 0x6a, 0x2c, 0xc0, 0xcc, 0xa1, 0xb4, 0x55, 0x1f, 0x07, 0xd7, 0x8f, 0xc1, 0x0a, 0x89,
	0x02, 0xdc, 0x35, 0xb4, 0x23, 0xed, 0x64, 0xd9, 0xd9, 0x19, 0x24, 0xb0, 0xd0, 0xf3, 0xda, 0xe1,
	0xa9, 0x25, 0xc3, 0x96, 0xab, 0xd2, 0xfa, 0x4f, 0x60, 0x8b, 0xa6, 0x08, 0xa8, 0x41, 0x69, 0x0b,
	0xc5, 0x24, 0x30, 0x96, 0x8e, 0xb4, 0x93, 0xbc, 0x63, 0xf5, 0x13, 0x58, 0xc8, 0xb0, 0x7f, 0xab,
	0xfe, 0xf0, 0x36, 0x81, 0x53, 0x95, 0xee, 0xd4, 0x5e, 0xaf, 0x81, 0xbd, 0xac, 0xc3, 0x48, 0x20,
	0x2f, 0x08, 0x18, 0xe6, 0xdc, 0xf8, 0x44, 0x02, 0x9a, 0x83, 0x04, 0x96, 0x94, 0x82, 0x39, 0x45,
	0x96, 0xab, 0x8f, 0x45, 0x2b, 0x2a, 0xa8, 0xff, 0x08, 0x76, 0x09, 0x47, 0x6d, 0x1a, 0xc4, 0x21,
	0x46, 0x9e, 0xef, 0xd3, 0x38, 0x12, 0xc6, 0xf2, 0x91, 0x76, 0xb2, 0xee, 0x7c, 0x36, 0x48, 0xa0,
	0x31, 0x6c, 0x68, 0xba, 0xc4, 0x72, 0xb7, 0x09, 0xbf, 0x90, 0xa1, 0x8a, 0x8a, 0xe8, 0x7f, 0x81,
	0x7c, 0x48, 0xfe, 0x89, 0x49, 0x40, 0x44, 0xcf, 0x58, 0x91, 0x82, 0x9c, 0xdb, 0x04, 0xe6, 0x5e,
	0x26, 0xf0, 0xb8, 0x49, 0xc4, 0xdf, 0x71, 0xa3, 0xec, 0xd3, 0xb6, 0xed, 0x53, 0xde, 0xa6, 0x7c,
	0xf8, 0xf3, 0x35, 0x0f, 0x5a, 0xb6, 0xe8, 0x75, 0x30, 0x2f, 0x57, 0x23, 0x31, 0x48, 0xe0, 0x8e,
	0xe2, 0xcb, 0x80, 0x2c, 0x77, 0x04, 0xaa, 0x5f, 0x6b, 0xe0, 0xd0, 0x8f, 0x19, 0xc3, 0x91, 0x40,
	0x8c, 0xc6, 0x51, 0x80, 0x46, 0x84, 0xab, 0x92, 0xb0, 0xbe, 0x30, 0xa1, 0xa9, 0x08, 0xdf, 0x03,
	0x6b, 0xb9, 0xfb, 0xc3, 0x8c, 0x9b, 0x26, 0xce, 0x33, 0x29, 0x97, 0xa0, 0x88, 0xbb, 0x1d, 0xca,
	0x63, 0x86, 0x39, 0x8a, 0xa8, 0x40, 0x57, 0x24, 0x0c, 0x71, 0x60, 0xac, 0xc9, 0xbf, 0x02, 0x1c,
	0x24, 0xf0, 0x53, 0x05, 0x3c, 0xaf, 0xca, 0x72, 0xf5, 0x2c, 0xfc, 0x33, 0x15, 0x67, 0x32, 0xa8,
	0x73, 0xb0, 0x23, 0xa8, 0xf0, 0x42, 0xd4, 0xc0, 0x02, 0x79, 0x6d, 0xf9, 0x10, 0xeb, 0xb2, 0xab,
	0xea, 0xc2, 0x5d, 0x1d, 0x2a, 0xf2, 0x69, 0x3c, 0xcb, 0xdd, 0x92, 0x21, 0x07, 0x8b, 0x8a, 0x0c,
	0xe8, 0xcf, 0x34, 0x60, 0x4e, 0xf6, 0x3e, 0xa3, 0x21, 0x2f, 0x35, 0xfc, 0xbe, 0xb0, 0x86, 0x2f,
	0xe7, 0xdd, 0xec, 0xac, 0xa2, 0xd2, 0xf8, 0x05, 0xff, 0x3a, 0xa9, 0xee, 0x4f, 0xb0, 0xde, 0xf6,
	0xba, 0x28, 0xa4, 0x9c, 0x1b, 0x40, 0xca, 0xa8, 0x2c, 0x2c, 0x63, 0x5b, 0xc9, 0x78, 0xc0, 0xb1,
	0xdc, 0xb5, 0xb6, 0xd7, 0x3d, 0xa7, 0x9c, 0xeb, 0xff, 0x6b, 0xe0, 0x60, 0x52, 0x5d, 0x46, 0xb6,
	0x21, 0xc9, 0x6a, 0x0b, 0x93, 0x7d, 0x3e, 0xaf, 0xe7, 0x11, 0xf5, 0xde, 0x78, 0xaf, 0x17, 0x43,
	0x19, 0xcf, 0x35, 0x00, 0xe7, 0x1f, 0x40, 0x34, 0x08, 0xb8, 0x34, 0x8c, 0x82, 0xd4, 0xd3, 0xea,
	0x27, 0xb0, 0xf4, 0xfd, 0x2c, 0x44, 0x2d, 0x08, 0xb8, 0xb2, 0x8f, 0xc7, 0x80, 0x06, 0x09, 0x3c,
	0xfe, 0x90, 0xc4, 0xac, 0xd0, 0x72, 0x1f, 0x83, 0xd2, 0x5b, 0x60, 0xd3, 0xf3, 0x45, 0xec, 0x85,
	0xa8, 0xc3, 0xe8, 0x15, 0x11, 0xc6, 0xa6, 0x14, 0x79, 0xb6, 0xf0, 0xa5, 0x15, 0x95, 0xa2, 0x09,
	0x30, 0xcb, 0x2d, 0xa8, 0x7d, 0x5d, 0x6e, 0xf5, 0x6f, 0x01, 0x20, 0x1c, 0x71, 0x2c, 0x44, 0xfa,
	0x95, 0x6d, 0x49, 0x7f, 0xda, 0x1f, 0x24, 0x70, 0x37, 0xf3, 0xa7, 0x61, 0xce, 0x72, 0xf3, 0x84,
	0xff, 0xa2, 0xd6, 0xa7, 0x85, 0xeb, 0x1b, 0x98, 0x7b, 0x7a, 0x03, 0x73, 0x6f, 0x6e, 0x60, 0xce,
	0x7a, 0xad, 0x81, 0xe2, 0x84, 0x83, 0x3b, 0x58, 0xd4, 0x3d, 0xc2, 0xe6, 0x18, 0xb4, 0xf6, 0x51,
	0x0c, 0x3a, 0xe5, 0x40, 0x6a, 0x44, 0x2c, 0x49, 0x5f, 0x98, 0x67, 0xd0, 0xa3, 0xa2, 0x71, 0x83,
	0x4e, 0xa3, 0x55, 0x39, 0x3d, 0x6c, 0xb0, 0x96, 0x7e, 0x2e, 0xa9, 0x2a, 0xe5, 0xf2, 0xfb, 0xfd,
	0x04, 0xae, 0x3a, 0x58, 0x28, 0x3d, 0x0f, 0x49, 0xf7, 0x61, 0xe1, 0x5c, 0xde, 0xf6, 0x4d, 0xed,
	0xae, 0x6f, 0x6a, 0xaf, 0xfa, 0xa6, 0xf6, 0xe4, 0xde, 0xcc, 0xdd, 0xdd, 0x9b, 0xb9, 0x17, 0xf7,
	0x66, 0xee, 0x8f, 0xef, 0xc6, 0x9e, 0x64, 0x38, 0x0a, 0xed, 0xb1, 0x91, 0x68, 0x77, 0x67, 0xc7,
	0xa9, 0x7c, 0xa7, 0xc6, 0xaa, 0x9c, 0x85, 0xdf, 0xbc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x0d,
	0x97, 0x9f, 0x74, 0x07, 0x00, 0x00,
}

func (m *OrderBookParticipation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBookParticipation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBookParticipation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsSettled {
		i--
		if m.IsSettled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	{
		size := m.ActualProfit.Size()
		i -= size
		if _, err := m.ActualProfit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if len(m.CurrentRoundMaxLossOddsUID) > 0 {
		i -= len(m.CurrentRoundMaxLossOddsUID)
		copy(dAtA[i:], m.CurrentRoundMaxLossOddsUID)
		i = encodeVarintParticipation(dAtA, i, uint64(len(m.CurrentRoundMaxLossOddsUID)))
		i--
		dAtA[i] = 0x62
	}
	{
		size := m.CurrentRoundMaxLoss.Size()
		i -= size
		if _, err := m.CurrentRoundMaxLoss.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.MaxLoss.Size()
		i -= size
		if _, err := m.MaxLoss.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.CurrentRoundTotalBetAmount.Size()
		i -= size
		if _, err := m.CurrentRoundTotalBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.TotalBetAmount.Size()
		i -= size
		if _, err := m.TotalBetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.ExposuresNotFilled != 0 {
		i = encodeVarintParticipation(dAtA, i, uint64(m.ExposuresNotFilled))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.CurrentRoundLiquidity.Size()
		i -= size
		if _, err := m.CurrentRoundLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Liquidity.Size()
		i -= size
		if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParticipation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.IsModuleAccount {
		i--
		if m.IsModuleAccount {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.ParticipantAddress) > 0 {
		i -= len(m.ParticipantAddress)
		copy(dAtA[i:], m.ParticipantAddress)
		i = encodeVarintParticipation(dAtA, i, uint64(len(m.ParticipantAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OrderBookUID) > 0 {
		i -= len(m.OrderBookUID)
		copy(dAtA[i:], m.OrderBookUID)
		i = encodeVarintParticipation(dAtA, i, uint64(len(m.OrderBookUID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintParticipation(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ParticipationBetPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParticipationBetPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParticipationBetPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BetUID) > 0 {
		i -= len(m.BetUID)
		copy(dAtA[i:], m.BetUID)
		i = encodeVarintParticipation(dAtA, i, uint64(len(m.BetUID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ParticipationIndex != 0 {
		i = encodeVarintParticipation(dAtA, i, uint64(m.ParticipationIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.OrderBookUID) > 0 {
		i -= len(m.OrderBookUID)
		copy(dAtA[i:], m.OrderBookUID)
		i = encodeVarintParticipation(dAtA, i, uint64(len(m.OrderBookUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParticipation(dAtA []byte, offset int, v uint64) int {
	offset -= sovParticipation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderBookParticipation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovParticipation(uint64(m.Index))
	}
	l = len(m.OrderBookUID)
	if l > 0 {
		n += 1 + l + sovParticipation(uint64(l))
	}
	l = len(m.ParticipantAddress)
	if l > 0 {
		n += 1 + l + sovParticipation(uint64(l))
	}
	if m.IsModuleAccount {
		n += 2
	}
	l = m.Liquidity.Size()
	n += 1 + l + sovParticipation(uint64(l))
	l = m.CurrentRoundLiquidity.Size()
	n += 1 + l + sovParticipation(uint64(l))
	if m.ExposuresNotFilled != 0 {
		n += 1 + sovParticipation(uint64(m.ExposuresNotFilled))
	}
	l = m.TotalBetAmount.Size()
	n += 1 + l + sovParticipation(uint64(l))
	l = m.CurrentRoundTotalBetAmount.Size()
	n += 1 + l + sovParticipation(uint64(l))
	l = m.MaxLoss.Size()
	n += 1 + l + sovParticipation(uint64(l))
	l = m.CurrentRoundMaxLoss.Size()
	n += 1 + l + sovParticipation(uint64(l))
	l = len(m.CurrentRoundMaxLossOddsUID)
	if l > 0 {
		n += 1 + l + sovParticipation(uint64(l))
	}
	l = m.ActualProfit.Size()
	n += 1 + l + sovParticipation(uint64(l))
	if m.IsSettled {
		n += 2
	}
	return n
}

func (m *ParticipationBetPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OrderBookUID)
	if l > 0 {
		n += 1 + l + sovParticipation(uint64(l))
	}
	if m.ParticipationIndex != 0 {
		n += 1 + sovParticipation(uint64(m.ParticipationIndex))
	}
	l = len(m.BetUID)
	if l > 0 {
		n += 1 + l + sovParticipation(uint64(l))
	}
	return n
}

func sovParticipation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParticipation(x uint64) (n int) {
	return sovParticipation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderBookParticipation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParticipation
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
			return fmt.Errorf("proto: OrderBookParticipation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBookParticipation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipantAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipantAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsModuleAccount", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
			m.IsModuleAccount = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRoundLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentRoundLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExposuresNotFilled", wireType)
			}
			m.ExposuresNotFilled = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExposuresNotFilled |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRoundTotalBetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentRoundTotalBetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLoss", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLoss.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRoundMaxLoss", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentRoundMaxLoss.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRoundMaxLossOddsUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentRoundMaxLossOddsUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualProfit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ActualProfit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSettled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
			m.IsSettled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParticipation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParticipation
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
func (m *ParticipationBetPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParticipation
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
			return fmt.Errorf("proto: ParticipationBetPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParticipationBetPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationIndex", wireType)
			}
			m.ParticipationIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipation
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
				return ErrInvalidLengthParticipation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BetUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParticipation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParticipation
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
func skipParticipation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParticipation
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
					return 0, ErrIntOverflowParticipation
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
					return 0, ErrIntOverflowParticipation
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
				return 0, ErrInvalidLengthParticipation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParticipation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParticipation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParticipation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParticipation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParticipation = fmt.Errorf("proto: unexpected end of group")
)
