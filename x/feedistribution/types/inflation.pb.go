// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/feedistribute/inflation.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
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

// InflationDistribution defines the distribution in which inflation is
// allocated through minting on each epoch (staking, incentives, community). It
// excludes the team vesting distribution, as this is minted once at genesis.
// The initial InflationDistribution can be calculated from the Evmos Token
// Model like this:
// mintDistribution1 = distribution1 / (1 - teamVestingDistribution)
// 0.5333333         = 40%           / (1 - 25%)
type InflationDistribution struct {
	// staking_rewards defines the proportion of the minted minted_denom that is
	// to be allocated as staking rewards
	StakingRewards cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=staking_rewards,json=stakingRewards,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"staking_rewards"`
	// Deprecated: usage_incentives defines the proportion of the minted minted_denom that is
	// to be allocated to the incentives module address
	UsageIncentives cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=usage_incentives,json=usageIncentives,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"usage_incentives"` // Deprecated: Do not use.
	// community_pool defines the proportion of the minted minted_denom that is to
	// be allocated to the community pool
	CommunityPool cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=community_pool,json=communityPool,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"community_pool"`
}

func (m *InflationDistribution) Reset()         { *m = InflationDistribution{} }
func (m *InflationDistribution) String() string { return proto.CompactTextString(m) }
func (*InflationDistribution) ProtoMessage()    {}
func (*InflationDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcb490784dc43978, []int{0}
}
func (m *InflationDistribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationDistribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationDistribution.Merge(m, src)
}
func (m *InflationDistribution) XXX_Size() int {
	return m.Size()
}
func (m *InflationDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_InflationDistribution proto.InternalMessageInfo

// ExponentialCalculation holds factors to calculate exponential inflation on
// each period. Calculation reference:
// periodProvision = exponentialDecay       *  bondingIncentive
// f(x)            = (a * (1 - r) ^ x + c)  *  (1 + max_variance - bondedRatio *
// (max_variance / bonding_target))
type ExponentialCalculation struct {
	// a defines the initial value
	A cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=a,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"a"`
	// r defines the reduction factor
	R cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=r,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"r"`
	// c defines the parameter for long term inflation
	C cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=c,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"c"`
	// bonding_target
	BondingTarget cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=bonding_target,json=bondingTarget,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"bonding_target"`
	// max_variance
	MaxVariance cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=max_variance,json=maxVariance,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_variance"`
}

func (m *ExponentialCalculation) Reset()         { *m = ExponentialCalculation{} }
func (m *ExponentialCalculation) String() string { return proto.CompactTextString(m) }
func (*ExponentialCalculation) ProtoMessage()    {}
func (*ExponentialCalculation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcb490784dc43978, []int{1}
}
func (m *ExponentialCalculation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExponentialCalculation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExponentialCalculation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExponentialCalculation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExponentialCalculation.Merge(m, src)
}
func (m *ExponentialCalculation) XXX_Size() int {
	return m.Size()
}
func (m *ExponentialCalculation) XXX_DiscardUnknown() {
	xxx_messageInfo_ExponentialCalculation.DiscardUnknown(m)
}

var xxx_messageInfo_ExponentialCalculation proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InflationDistribution)(nil), "exocore.feedistribute.InflationDistribution")
	proto.RegisterType((*ExponentialCalculation)(nil), "exocore.feedistribute.ExponentialCalculation")
}

func init() {
	proto.RegisterFile("exocore/feedistribute/inflation.proto", fileDescriptor_fcb490784dc43978)
}

var fileDescriptor_fcb490784dc43978 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0xf5, 0x0f, 0x38, 0xea, 0xae, 0x06, 0x57, 0xc2, 0x0a, 0x59, 0x59, 0x59, 0xf0,
	0x94, 0x41, 0x3c, 0x79, 0x5d, 0x77, 0x85, 0x95, 0x45, 0x96, 0x2a, 0x1e, 0xbc, 0x84, 0x37, 0x93,
	0x69, 0x3a, 0x34, 0x33, 0x6f, 0x98, 0x99, 0xb4, 0xe9, 0xd9, 0x2f, 0xe0, 0x97, 0x12, 0x7a, 0xec,
	0x51, 0x3c, 0x14, 0x69, 0xbf, 0x88, 0xe4, 0x4f, 0x0b, 0xbd, 0xa5, 0x97, 0xe1, 0x65, 0x78, 0x7e,
	0x0f, 0xef, 0x33, 0xf3, 0xd0, 0x0b, 0x51, 0x21, 0x47, 0x23, 0xd8, 0x48, 0x88, 0x54, 0x5a, 0x67,
	0x64, 0x52, 0x3a, 0xc1, 0xa4, 0x1e, 0xe5, 0xe0, 0x24, 0xea, 0xa8, 0x30, 0xe8, 0xd0, 0x3f, 0xe9,
	0x64, 0xd1, 0x9e, 0xec, 0xf4, 0x39, 0x28, 0xa9, 0x91, 0x35, 0x67, 0xab, 0x3c, 0x7d, 0x91, 0x61,
	0x86, 0xcd, 0xc8, 0xea, 0xa9, 0xbd, 0x3d, 0xff, 0x39, 0xa0, 0x27, 0x37, 0x5b, 0xcf, 0xab, 0xad,
	0x81, 0x44, 0xed, 0xdf, 0xd2, 0x63, 0xeb, 0x60, 0x22, 0x75, 0x16, 0x1b, 0x31, 0x03, 0x93, 0xda,
	0x80, 0xbc, 0x26, 0x6f, 0x1f, 0x5d, 0xbe, 0x59, 0xac, 0xce, 0xbc, 0xbf, 0xab, 0xb3, 0x57, 0x1c,
	0xad, 0x42, 0x6b, 0xd3, 0x49, 0x24, 0x91, 0x29, 0x70, 0xe3, 0xe8, 0x56, 0x64, 0xc0, 0xe7, 0x57,
	0x82, 0x0f, 0x8f, 0x3a, 0x76, 0xd8, 0xa2, 0xfe, 0x1d, 0x7d, 0x56, 0x5a, 0xc8, 0x44, 0x2c, 0x35,
	0x17, 0xda, 0xc9, 0xa9, 0xb0, 0xc1, 0xa0, 0xb1, 0xbb, 0xe8, 0x61, 0x17, 0x90, 0xe1, 0x71, 0x83,
	0xdf, 0xec, 0x68, 0xff, 0x33, 0x3d, 0xe2, 0xa8, 0x54, 0xa9, 0xa5, 0x9b, 0xc7, 0x05, 0x62, 0x1e,
	0xdc, 0xeb, 0xbf, 0xde, 0xd3, 0x1d, 0x7a, 0x87, 0x98, 0x9f, 0xff, 0x1e, 0xd0, 0x97, 0xd7, 0x55,
	0x81, 0xba, 0x36, 0x87, 0xfc, 0x23, 0xe4, 0xbc, 0x6c, 0x9f, 0xc4, 0x7f, 0x47, 0x09, 0x1c, 0x12,
	0x9c, 0x40, 0x8d, 0x98, 0x2e, 0x5c, 0x3f, 0xc4, 0xd4, 0x08, 0x3f, 0x64, 0x7f, 0xc2, 0xeb, 0xfc,
	0x09, 0xea, 0xb4, 0xfe, 0x1f, 0x07, 0x26, 0x13, 0x2e, 0xb8, 0x7f, 0x40, 0xfe, 0x0e, 0xfd, 0xd6,
	0x90, 0xfe, 0x27, 0xfa, 0x44, 0x41, 0x15, 0x4f, 0xc1, 0x48, 0xd0, 0x5c, 0x04, 0x0f, 0xfa, 0x3b,
	0x3d, 0x56, 0x50, 0x7d, 0xef, 0xb8, 0xcb, 0xaf, 0x8b, 0x75, 0x48, 0x96, 0xeb, 0x90, 0xfc, 0x5b,
	0x87, 0xe4, 0xd7, 0x26, 0xf4, 0x96, 0x9b, 0xd0, 0xfb, 0xb3, 0x09, 0xbd, 0x1f, 0x1f, 0x32, 0xe9,
	0xc6, 0x65, 0x12, 0x71, 0x54, 0xec, 0xba, 0xad, 0xec, 0x17, 0xe1, 0x66, 0x68, 0x26, 0x6c, 0x5b,
	0xf4, 0x6a, 0xaf, 0xea, 0x12, 0x35, 0x73, 0xf3, 0x42, 0xd8, 0xe4, 0x61, 0xd3, 0xd4, 0xf7, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x1a, 0xda, 0xd5, 0x12, 0x03, 0x00, 0x00,
}

func (m *InflationDistribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationDistribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationDistribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.UsageIncentives.Size()
		i -= size
		if _, err := m.UsageIncentives.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.StakingRewards.Size()
		i -= size
		if _, err := m.StakingRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ExponentialCalculation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExponentialCalculation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExponentialCalculation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxVariance.Size()
		i -= size
		if _, err := m.MaxVariance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.BondingTarget.Size()
		i -= size
		if _, err := m.BondingTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.C.Size()
		i -= size
		if _, err := m.C.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.R.Size()
		i -= size
		if _, err := m.R.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.A.Size()
		i -= size
		if _, err := m.A.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintInflation(dAtA []byte, offset int, v uint64) int {
	offset -= sovInflation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InflationDistribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StakingRewards.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.UsageIncentives.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func (m *ExponentialCalculation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.A.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.R.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.C.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.BondingTarget.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.MaxVariance.Size()
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func sovInflation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInflation(x uint64) (n int) {
	return sovInflation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InflationDistribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: InflationDistribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationDistribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsageIncentives", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UsageIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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
func (m *ExponentialCalculation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: ExponentialCalculation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExponentialCalculation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.R.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.C.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondingTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondingTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVariance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxVariance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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
func skipInflation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
				return 0, ErrInvalidLengthInflation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInflation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInflation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInflation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInflation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInflation = fmt.Errorf("proto: unexpected end of group")
)
