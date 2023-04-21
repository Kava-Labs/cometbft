// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/consensus/v2/types.proto

package v2

import (
	fmt "fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/consensus/v1"
	v3 "github.com/cometbft/cometbft/api/cometbft/types/v3"
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

// Vote is sent when voting for a proposal (or lack thereof).
type Vote struct {
	Vote *v3.Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbfa7f975842dd1, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetVote() *v3.Vote {
	if m != nil {
		return m.Vote
	}
	return nil
}

type Message struct {
	// Types that are valid to be assigned to Sum:
	//	*Message_NewRoundStep
	//	*Message_NewValidBlock
	//	*Message_Proposal
	//	*Message_ProposalPol
	//	*Message_BlockPart
	//	*Message_Vote
	//	*Message_HasVote
	//	*Message_VoteSetMaj23
	//	*Message_VoteSetBits
	Sum isMessage_Sum `protobuf_oneof:"sum"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fbfa7f975842dd1, []int{1}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Sum interface {
	isMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Message_NewRoundStep struct {
	NewRoundStep *v1.NewRoundStep `protobuf:"bytes,1,opt,name=new_round_step,json=newRoundStep,proto3,oneof" json:"new_round_step,omitempty"`
}
type Message_NewValidBlock struct {
	NewValidBlock *v1.NewValidBlock `protobuf:"bytes,2,opt,name=new_valid_block,json=newValidBlock,proto3,oneof" json:"new_valid_block,omitempty"`
}
type Message_Proposal struct {
	Proposal *v1.Proposal `protobuf:"bytes,3,opt,name=proposal,proto3,oneof" json:"proposal,omitempty"`
}
type Message_ProposalPol struct {
	ProposalPol *v1.ProposalPOL `protobuf:"bytes,4,opt,name=proposal_pol,json=proposalPol,proto3,oneof" json:"proposal_pol,omitempty"`
}
type Message_BlockPart struct {
	BlockPart *v1.BlockPart `protobuf:"bytes,5,opt,name=block_part,json=blockPart,proto3,oneof" json:"block_part,omitempty"`
}
type Message_Vote struct {
	Vote *Vote `protobuf:"bytes,6,opt,name=vote,proto3,oneof" json:"vote,omitempty"`
}
type Message_HasVote struct {
	HasVote *v1.HasVote `protobuf:"bytes,7,opt,name=has_vote,json=hasVote,proto3,oneof" json:"has_vote,omitempty"`
}
type Message_VoteSetMaj23 struct {
	VoteSetMaj23 *v1.VoteSetMaj23 `protobuf:"bytes,8,opt,name=vote_set_maj23,json=voteSetMaj23,proto3,oneof" json:"vote_set_maj23,omitempty"`
}
type Message_VoteSetBits struct {
	VoteSetBits *v1.VoteSetBits `protobuf:"bytes,9,opt,name=vote_set_bits,json=voteSetBits,proto3,oneof" json:"vote_set_bits,omitempty"`
}

func (*Message_NewRoundStep) isMessage_Sum()  {}
func (*Message_NewValidBlock) isMessage_Sum() {}
func (*Message_Proposal) isMessage_Sum()      {}
func (*Message_ProposalPol) isMessage_Sum()   {}
func (*Message_BlockPart) isMessage_Sum()     {}
func (*Message_Vote) isMessage_Sum()          {}
func (*Message_HasVote) isMessage_Sum()       {}
func (*Message_VoteSetMaj23) isMessage_Sum()  {}
func (*Message_VoteSetBits) isMessage_Sum()   {}

func (m *Message) GetSum() isMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Message) GetNewRoundStep() *v1.NewRoundStep {
	if x, ok := m.GetSum().(*Message_NewRoundStep); ok {
		return x.NewRoundStep
	}
	return nil
}

func (m *Message) GetNewValidBlock() *v1.NewValidBlock {
	if x, ok := m.GetSum().(*Message_NewValidBlock); ok {
		return x.NewValidBlock
	}
	return nil
}

func (m *Message) GetProposal() *v1.Proposal {
	if x, ok := m.GetSum().(*Message_Proposal); ok {
		return x.Proposal
	}
	return nil
}

func (m *Message) GetProposalPol() *v1.ProposalPOL {
	if x, ok := m.GetSum().(*Message_ProposalPol); ok {
		return x.ProposalPol
	}
	return nil
}

func (m *Message) GetBlockPart() *v1.BlockPart {
	if x, ok := m.GetSum().(*Message_BlockPart); ok {
		return x.BlockPart
	}
	return nil
}

func (m *Message) GetVote() *Vote {
	if x, ok := m.GetSum().(*Message_Vote); ok {
		return x.Vote
	}
	return nil
}

func (m *Message) GetHasVote() *v1.HasVote {
	if x, ok := m.GetSum().(*Message_HasVote); ok {
		return x.HasVote
	}
	return nil
}

func (m *Message) GetVoteSetMaj23() *v1.VoteSetMaj23 {
	if x, ok := m.GetSum().(*Message_VoteSetMaj23); ok {
		return x.VoteSetMaj23
	}
	return nil
}

func (m *Message) GetVoteSetBits() *v1.VoteSetBits {
	if x, ok := m.GetSum().(*Message_VoteSetBits); ok {
		return x.VoteSetBits
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_NewRoundStep)(nil),
		(*Message_NewValidBlock)(nil),
		(*Message_Proposal)(nil),
		(*Message_ProposalPol)(nil),
		(*Message_BlockPart)(nil),
		(*Message_Vote)(nil),
		(*Message_HasVote)(nil),
		(*Message_VoteSetMaj23)(nil),
		(*Message_VoteSetBits)(nil),
	}
}

func init() {
	proto.RegisterType((*Vote)(nil), "cometbft.consensus.v2.Vote")
	proto.RegisterType((*Message)(nil), "cometbft.consensus.v2.Message")
}

func init() { proto.RegisterFile("cometbft/consensus/v2/types.proto", fileDescriptor_1fbfa7f975842dd1) }

var fileDescriptor_1fbfa7f975842dd1 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x1d, 0xd6, 0xad, 0x9d, 0xb7, 0x81, 0x64, 0x09, 0x11, 0x0d, 0x11, 0x46, 0xe1, 0x80,
	0x84, 0x94, 0xa8, 0x89, 0xc4, 0x05, 0x71, 0xa0, 0x17, 0x2c, 0xc1, 0x46, 0x94, 0x49, 0x3b, 0x70,
	0xb1, 0x9c, 0xce, 0xac, 0x81, 0x34, 0xb6, 0x62, 0xc7, 0x13, 0x6f, 0xc1, 0x63, 0x71, 0xdc, 0x91,
	0x23, 0x6a, 0x1f, 0x04, 0x64, 0x37, 0x09, 0x39, 0x24, 0xe5, 0xe6, 0xcf, 0xf9, 0xff, 0x7e, 0x76,
	0xfc, 0xe9, 0x83, 0xcf, 0x16, 0x7c, 0xc5, 0x54, 0xfa, 0x45, 0x05, 0x0b, 0x5e, 0x48, 0x56, 0xc8,
	0x4a, 0x06, 0x3a, 0x0c, 0xd4, 0x77, 0xc1, 0xa4, 0x2f, 0x4a, 0xae, 0x38, 0x7a, 0xd8, 0x44, 0xfc,
	0x36, 0xe2, 0xeb, 0xf0, 0xb4, 0x97, 0x9c, 0x75, 0xc9, 0xd3, 0x27, 0x6d, 0xc4, 0xee, 0x06, 0x3a,
	0xea, 0x7e, 0x9e, 0x46, 0x70, 0x74, 0xc5, 0x15, 0x43, 0xaf, 0xe0, 0x48, 0x73, 0xc5, 0x5c, 0xe7,
	0xcc, 0x79, 0x79, 0x14, 0x3e, 0xf2, 0xdb, 0xf3, 0xb6, 0x61, 0x1d, 0xf9, 0x26, 0x96, 0xd8, 0xd0,
	0xf4, 0xcf, 0x08, 0x8e, 0xcf, 0x99, 0x94, 0xf4, 0x86, 0xa1, 0x0f, 0xf0, 0x7e, 0xc1, 0x6e, 0x49,
	0xc9, 0xab, 0xe2, 0x9a, 0x48, 0xc5, 0x44, 0xad, 0x78, 0xee, 0xf7, 0x5d, 0x79, 0xe6, 0x5f, 0xb0,
	0xdb, 0xc4, 0x64, 0x2f, 0x15, 0x13, 0x18, 0x24, 0xc7, 0x45, 0xa7, 0x46, 0x17, 0xf0, 0x81, 0x91,
	0x69, 0x9a, 0x67, 0xd7, 0x24, 0xcd, 0xf9, 0xe2, 0x9b, 0x7b, 0xcf, 0xda, 0x5e, 0x0c, 0xdb, 0xae,
	0x4c, 0x78, 0x6e, 0xb2, 0x18, 0x24, 0x27, 0x45, 0x77, 0x03, 0xbd, 0x85, 0x13, 0x51, 0x72, 0xc1,
	0x25, 0xcd, 0xdd, 0x3d, 0x2b, 0x7a, 0x3a, 0x20, 0x8a, 0xeb, 0x18, 0x06, 0x49, 0x8b, 0xa0, 0xf7,
	0xf0, 0xb8, 0x59, 0x13, 0xc1, 0x73, 0x77, 0x64, 0x15, 0xd3, 0xff, 0x28, 0xe2, 0x4f, 0x1f, 0x31,
	0x48, 0x8e, 0x1a, 0x32, 0xe6, 0x39, 0x7a, 0x07, 0xa1, 0xfd, 0x1b, 0x22, 0x68, 0xa9, 0xdc, 0x7d,
	0xab, 0x39, 0x1b, 0xd0, 0xd8, 0x9b, 0xc7, 0xb4, 0x54, 0x18, 0x24, 0x87, 0x69, 0x53, 0xa0, 0x59,
	0xdd, 0xa0, 0x03, 0x0b, 0x3f, 0xee, 0x85, 0x43, 0xdb, 0x24, 0x0c, 0xb6, 0x6d, 0x42, 0x6f, 0xe0,
	0x64, 0x49, 0x25, 0xb1, 0xd8, 0xd8, 0x62, 0xde, 0xc0, 0x99, 0x98, 0xca, 0x9a, 0x1c, 0x2f, 0xb7,
	0x4b, 0xd3, 0x57, 0x03, 0x12, 0xc9, 0x14, 0x59, 0xd1, 0xaf, 0x61, 0xe4, 0x4e, 0x76, 0xf6, 0xd5,
	0x40, 0x97, 0x4c, 0x9d, 0x9b, 0xa8, 0xe9, 0xab, 0xee, 0xd4, 0x08, 0xc3, 0x93, 0x56, 0x96, 0x66,
	0x4a, 0xba, 0x87, 0x3b, 0x5f, 0xb2, 0x76, 0xcd, 0x33, 0x25, 0xcd, 0x4b, 0xea, 0x7f, 0xe5, 0x7c,
	0x1f, 0xee, 0xc9, 0x6a, 0x35, 0x8f, 0x7f, 0xae, 0x3d, 0xe7, 0x6e, 0xed, 0x39, 0xbf, 0xd7, 0x9e,
	0xf3, 0x63, 0xe3, 0x81, 0xbb, 0x8d, 0x07, 0x7e, 0x6d, 0x3c, 0xf0, 0xf9, 0xf5, 0x4d, 0xa6, 0x96,
	0x55, 0x6a, 0xcc, 0x41, 0x67, 0x3a, 0xea, 0x05, 0x15, 0x59, 0xd0, 0x3b, 0x6d, 0xe9, 0x81, 0x9d,
	0x87, 0xe8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xa1, 0x1b, 0xee, 0x8d, 0x03, 0x00, 0x00,
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != nil {
		{
			size, err := m.Vote.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message_NewRoundStep) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_NewRoundStep) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NewRoundStep != nil {
		{
			size, err := m.NewRoundStep.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Message_NewValidBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_NewValidBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NewValidBlock != nil {
		{
			size, err := m.NewValidBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Message_Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Proposal != nil {
		{
			size, err := m.Proposal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Message_ProposalPol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_ProposalPol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProposalPol != nil {
		{
			size, err := m.ProposalPol.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Message_BlockPart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_BlockPart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BlockPart != nil {
		{
			size, err := m.BlockPart.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Message_Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Vote != nil {
		{
			size, err := m.Vote.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Message_HasVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_HasVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HasVote != nil {
		{
			size, err := m.HasVote.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Message_VoteSetMaj23) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_VoteSetMaj23) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.VoteSetMaj23 != nil {
		{
			size, err := m.VoteSetMaj23.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Message_VoteSetBits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_VoteSetBits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.VoteSetBits != nil {
		{
			size, err := m.VoteSetBits.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vote != nil {
		l = m.Vote.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Message_NewRoundStep) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NewRoundStep != nil {
		l = m.NewRoundStep.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_NewValidBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NewValidBlock != nil {
		l = m.NewValidBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proposal != nil {
		l = m.Proposal.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_ProposalPol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalPol != nil {
		l = m.ProposalPol.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_BlockPart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockPart != nil {
		l = m.BlockPart.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vote != nil {
		l = m.Vote.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_HasVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HasVote != nil {
		l = m.HasVote.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_VoteSetMaj23) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VoteSetMaj23 != nil {
		l = m.VoteSetMaj23.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_VoteSetBits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VoteSetBits != nil {
		l = m.VoteSetBits.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Vote == nil {
				m.Vote = &v3.Vote{}
			}
			if err := m.Vote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewRoundStep", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.NewRoundStep{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_NewRoundStep{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewValidBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.NewValidBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_NewValidBlock{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.Proposal{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_Proposal{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalPol", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.ProposalPOL{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_ProposalPol{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockPart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.BlockPart{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_BlockPart{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Vote{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_Vote{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasVote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.HasVote{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_HasVote{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteSetMaj23", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.VoteSetMaj23{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_VoteSetMaj23{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteSetBits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.VoteSetBits{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_VoteSetBits{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
