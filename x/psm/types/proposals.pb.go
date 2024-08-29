// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onomyprotocol/psm/v1/proposals.proto

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

type AddStableCoinProposal struct {
	Title       string                                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Denom       string                                 `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	LimitTotal  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=limit_total,json=limitTotal,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"limit_total" yaml:"pool_rate"`
	Price       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"pool_rate"`
	FeeIn       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=fee_in,json=feeIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_in" yaml:"pool_rate"`
	FeeOut      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=fee_out,json=feeOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_out" yaml:"pool_rate"`
}

func (m *AddStableCoinProposal) Reset()         { *m = AddStableCoinProposal{} }
func (m *AddStableCoinProposal) String() string { return proto.CompactTextString(m) }
func (*AddStableCoinProposal) ProtoMessage()    {}
func (*AddStableCoinProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f562e7796f334250, []int{0}
}
func (m *AddStableCoinProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddStableCoinProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddStableCoinProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddStableCoinProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStableCoinProposal.Merge(m, src)
}
func (m *AddStableCoinProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddStableCoinProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStableCoinProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddStableCoinProposal proto.InternalMessageInfo

func (m *AddStableCoinProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddStableCoinProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddStableCoinProposal) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type UpdatesStableCoinProposal struct {
	Title             string                                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string                                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Denom             string                                 `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	UpdatesLimitTotal github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=updates_limit_total,json=updatesLimitTotal,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"updates_limit_total" yaml:"pool_rate"`
	Price             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"pool_rate"`
	FeeIn             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=fee_in,json=feeIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_in" yaml:"pool_rate"`
	FeeOut            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=fee_out,json=feeOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_out" yaml:"pool_rate"`
}

func (m *UpdatesStableCoinProposal) Reset()         { *m = UpdatesStableCoinProposal{} }
func (m *UpdatesStableCoinProposal) String() string { return proto.CompactTextString(m) }
func (*UpdatesStableCoinProposal) ProtoMessage()    {}
func (*UpdatesStableCoinProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f562e7796f334250, []int{1}
}
func (m *UpdatesStableCoinProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatesStableCoinProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatesStableCoinProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatesStableCoinProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatesStableCoinProposal.Merge(m, src)
}
func (m *UpdatesStableCoinProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdatesStableCoinProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatesStableCoinProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatesStableCoinProposal proto.InternalMessageInfo

func (m *UpdatesStableCoinProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdatesStableCoinProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdatesStableCoinProposal) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*AddStableCoinProposal)(nil), "onomyprotocol.psm.v1.AddStableCoinProposal")
	proto.RegisterType((*UpdatesStableCoinProposal)(nil), "onomyprotocol.psm.v1.UpdatesStableCoinProposal")
}

func init() {
	proto.RegisterFile("onomyprotocol/psm/v1/proposals.proto", fileDescriptor_f562e7796f334250)
}

var fileDescriptor_f562e7796f334250 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x13, 0x6b, 0x5b, 0x9c, 0xba, 0xd0, 0x58, 0x21, 0xba, 0x48, 0x4b, 0x10, 0xe9, 0xc6,
	0x84, 0xe2, 0xce, 0x9d, 0xd1, 0x4d, 0x41, 0x50, 0xaa, 0x82, 0x97, 0x45, 0x48, 0x93, 0x69, 0x1d,
	0x9c, 0xc9, 0x19, 0x32, 0x93, 0x62, 0xdf, 0xc2, 0xc7, 0xea, 0xb2, 0x4b, 0x11, 0x29, 0xd2, 0x3e,
	0x80, 0xe0, 0x13, 0x48, 0x26, 0x5d, 0x28, 0xa8, 0x20, 0xb4, 0x3b, 0x57, 0x33, 0xe7, 0xf6, 0x9d,
	0x03, 0x3f, 0xfc, 0x68, 0x07, 0x62, 0x60, 0x03, 0x9e, 0x80, 0x84, 0x10, 0xa8, 0xcb, 0x05, 0x73,
	0xfb, 0x4d, 0x97, 0x27, 0xc0, 0x41, 0x04, 0x54, 0x38, 0xaa, 0x62, 0x54, 0xbf, 0x74, 0x39, 0x5c,
	0x30, 0xa7, 0xdf, 0xdc, 0xae, 0xf6, 0xa0, 0x07, 0x2a, 0xe9, 0x66, 0xbf, 0xbc, 0xd7, 0x7e, 0x29,
	0xa0, 0xcd, 0xc3, 0x28, 0x3a, 0x97, 0x41, 0x87, 0xe2, 0x23, 0x20, 0xf1, 0xd9, 0x0c, 0x66, 0x54,
	0x51, 0x51, 0x12, 0x49, 0xb1, 0xa9, 0xd7, 0xf5, 0xc6, 0x4a, 0x3b, 0x0f, 0x8c, 0x3a, 0xaa, 0x44,
	0x58, 0x84, 0x09, 0xe1, 0x92, 0x40, 0x6c, 0x2e, 0xa9, 0xda, 0xe7, 0x54, 0x36, 0x17, 0xe1, 0x18,
	0x98, 0x59, 0xc8, 0xe7, 0x54, 0x60, 0x84, 0xa8, 0x42, 0x09, 0x23, 0xd2, 0x97, 0x20, 0x03, 0x6a,
	0x2e, 0xd7, 0xf5, 0xc6, 0xaa, 0xe7, 0x0d, 0xc7, 0x35, 0xed, 0x79, 0x5c, 0xdb, 0xed, 0x11, 0x79,
	0x97, 0x76, 0x9c, 0x10, 0x98, 0x1b, 0x82, 0x60, 0x20, 0x66, 0xcf, 0x9e, 0x88, 0xee, 0x5d, 0x39,
	0xe0, 0x58, 0x38, 0xad, 0x58, 0xbe, 0x8f, 0x6b, 0x6b, 0x83, 0x80, 0xd1, 0x03, 0x9b, 0x03, 0x50,
	0x3f, 0x09, 0x24, 0xb6, 0xdb, 0x48, 0x61, 0x2f, 0x32, 0xaa, 0x71, 0x85, 0x8a, 0x3c, 0x21, 0x21,
	0x36, 0x8b, 0x7f, 0xc6, 0x1f, 0xe3, 0xf0, 0x5b, 0x7c, 0x0e, 0x34, 0xae, 0x51, 0xa9, 0x8b, 0xb1,
	0x4f, 0x62, 0xb3, 0x34, 0x3f, 0x74, 0x17, 0xe3, 0x56, 0x6c, 0xdc, 0xa2, 0x72, 0x86, 0x86, 0x54,
	0x9a, 0xe5, 0xb9, 0xb1, 0xb3, 0x6b, 0x4f, 0x53, 0x69, 0xbf, 0x15, 0xd0, 0xd6, 0x25, 0x8f, 0x02,
	0x89, 0xc5, 0xc2, 0x25, 0x4e, 0xd0, 0x46, 0x9a, 0xaf, 0xf2, 0x17, 0x23, 0xf5, 0xfa, 0x0c, 0x7f,
	0xf2, 0xaf, 0xf8, 0x8f, 0x8a, 0x7b, 0xde, 0x70, 0x62, 0xe9, 0xa3, 0x89, 0xa5, 0xbf, 0x4e, 0x2c,
	0xfd, 0x71, 0x6a, 0x69, 0xa3, 0xa9, 0xa5, 0x3d, 0x4d, 0x2d, 0xed, 0xa6, 0xf1, 0x2b, 0xfd, 0x41,
	0xb9, 0x89, 0xda, 0xd1, 0x29, 0x29, 0x6f, 0xd8, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xe5,
	0x77, 0x96, 0x6f, 0x04, 0x00, 0x00,
}

func (m *AddStableCoinProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddStableCoinProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddStableCoinProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeOut.Size()
		i -= size
		if _, err := m.FeeOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.FeeIn.Size()
		i -= size
		if _, err := m.FeeIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.LimitTotal.Size()
		i -= size
		if _, err := m.LimitTotal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdatesStableCoinProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatesStableCoinProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdatesStableCoinProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeOut.Size()
		i -= size
		if _, err := m.FeeOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.FeeIn.Size()
		i -= size
		if _, err := m.FeeIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.UpdatesLimitTotal.Size()
		i -= size
		if _, err := m.UpdatesLimitTotal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposals(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposals(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddStableCoinProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.LimitTotal.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.FeeIn.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.FeeOut.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func (m *UpdatesStableCoinProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.UpdatesLimitTotal.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.FeeIn.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.FeeOut.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func sovProposals(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposals(x uint64) (n int) {
	return sovProposals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddStableCoinProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: AddStableCoinProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddStableCoinProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitTotal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LimitTotal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeIn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeOut", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *UpdatesStableCoinProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: UpdatesStableCoinProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatesStableCoinProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatesLimitTotal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UpdatesLimitTotal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeIn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeOut", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func skipProposals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
				return 0, ErrInvalidLengthProposals
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposals
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposals
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposals        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposals          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposals = fmt.Errorf("proto: unexpected end of group")
)