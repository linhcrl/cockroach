// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/execinfrapb/processors_changefeeds.proto

// Beware! This package name must not be changed, even though it doesn't match
// the Go package name, because it defines the Protobuf message names which
// can't be changed without breaking backward compatibility.

package execinfrapb

import (
	fmt "fmt"
	github_com_cockroachdb_cockroach_pkg_jobs_jobspb "github.com/cockroachdb/cockroach/pkg/jobs/jobspb"
	jobspb "github.com/cockroachdb/cockroach/pkg/jobs/jobspb"
	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	github_com_cockroachdb_cockroach_pkg_security "github.com/cockroachdb/cockroach/pkg/security"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
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

// ChangeAggregatorSpec is the specification for a processor that watches for
// changes in a set of spans. Each span may cross multiple ranges.
type ChangeAggregatorSpec struct {
	Watches []ChangeAggregatorSpec_Watch `protobuf:"bytes,1,rep,name=watches" json:"watches"`
	// Feed is the specification for this changefeed.
	Feed jobspb.ChangefeedDetails `protobuf:"bytes,2,opt,name=feed" json:"feed"`
	// User who initiated the changefeed. This is used to check access privileges
	// when using FileTable ExternalStorage.
	UserProto github_com_cockroachdb_cockroach_pkg_security.SQLUsernameProto `protobuf:"bytes,3,opt,name=user_proto,json=userProto,casttype=github.com/cockroachdb/cockroach/pkg/security.SQLUsernameProto" json:"user_proto"`
	// JobID is the id of this changefeed in the system jobs.
	JobID github_com_cockroachdb_cockroach_pkg_jobs_jobspb.JobID `protobuf:"varint,4,opt,name=job_id,json=jobId,casttype=github.com/cockroachdb/cockroach/pkg/jobs/jobspb.JobID" json:"job_id"`
}

func (m *ChangeAggregatorSpec) Reset()         { *m = ChangeAggregatorSpec{} }
func (m *ChangeAggregatorSpec) String() string { return proto.CompactTextString(m) }
func (*ChangeAggregatorSpec) ProtoMessage()    {}
func (*ChangeAggregatorSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cd2be2dc9fdf3b5, []int{0}
}
func (m *ChangeAggregatorSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeAggregatorSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ChangeAggregatorSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAggregatorSpec.Merge(m, src)
}
func (m *ChangeAggregatorSpec) XXX_Size() int {
	return m.Size()
}
func (m *ChangeAggregatorSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAggregatorSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAggregatorSpec proto.InternalMessageInfo

type ChangeAggregatorSpec_Watch struct {
	InitialResolved hlc.Timestamp `protobuf:"bytes,1,opt,name=initial_resolved,json=initialResolved" json:"initial_resolved"`
	Span            roachpb.Span  `protobuf:"bytes,2,opt,name=span" json:"span"`
}

func (m *ChangeAggregatorSpec_Watch) Reset()         { *m = ChangeAggregatorSpec_Watch{} }
func (m *ChangeAggregatorSpec_Watch) String() string { return proto.CompactTextString(m) }
func (*ChangeAggregatorSpec_Watch) ProtoMessage()    {}
func (*ChangeAggregatorSpec_Watch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cd2be2dc9fdf3b5, []int{0, 0}
}
func (m *ChangeAggregatorSpec_Watch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeAggregatorSpec_Watch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ChangeAggregatorSpec_Watch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAggregatorSpec_Watch.Merge(m, src)
}
func (m *ChangeAggregatorSpec_Watch) XXX_Size() int {
	return m.Size()
}
func (m *ChangeAggregatorSpec_Watch) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAggregatorSpec_Watch.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAggregatorSpec_Watch proto.InternalMessageInfo

// ChangeFrontierSpec is the specification for a processor that receives
// span-level resolved timestamps, track them, and emits the changefeed-level
// resolved timestamp whenever it changes.
type ChangeFrontierSpec struct {
	// TrackedSpans is the entire span set being watched. Once all these spans
	// have been resolved at a certain timestamp, then it's safe to resolve the
	// changefeed at that timestamp.
	TrackedSpans []roachpb.Span `protobuf:"bytes,1,rep,name=tracked_spans,json=trackedSpans" json:"tracked_spans"`
	// Feed is the specification for this changefeed.
	Feed jobspb.ChangefeedDetails `protobuf:"bytes,2,opt,name=feed" json:"feed"`
	// JobID is the id of this changefeed in the system jobs.
	JobID github_com_cockroachdb_cockroach_pkg_jobs_jobspb.JobID `protobuf:"varint,3,opt,name=job_id,json=jobId,casttype=github.com/cockroachdb/cockroach/pkg/jobs/jobspb.JobID" json:"job_id"`
	// User who initiated the changefeed. This is used to check access privileges
	// when using FileTable ExternalStorage.
	UserProto github_com_cockroachdb_cockroach_pkg_security.SQLUsernameProto `protobuf:"bytes,4,opt,name=user_proto,json=userProto,casttype=github.com/cockroachdb/cockroach/pkg/security.SQLUsernameProto" json:"user_proto"`
}

func (m *ChangeFrontierSpec) Reset()         { *m = ChangeFrontierSpec{} }
func (m *ChangeFrontierSpec) String() string { return proto.CompactTextString(m) }
func (*ChangeFrontierSpec) ProtoMessage()    {}
func (*ChangeFrontierSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cd2be2dc9fdf3b5, []int{1}
}
func (m *ChangeFrontierSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeFrontierSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ChangeFrontierSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeFrontierSpec.Merge(m, src)
}
func (m *ChangeFrontierSpec) XXX_Size() int {
	return m.Size()
}
func (m *ChangeFrontierSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeFrontierSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeFrontierSpec proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChangeAggregatorSpec)(nil), "cockroach.sql.distsqlrun.ChangeAggregatorSpec")
	proto.RegisterType((*ChangeAggregatorSpec_Watch)(nil), "cockroach.sql.distsqlrun.ChangeAggregatorSpec.Watch")
	proto.RegisterType((*ChangeFrontierSpec)(nil), "cockroach.sql.distsqlrun.ChangeFrontierSpec")
}

func init() {
	proto.RegisterFile("sql/execinfrapb/processors_changefeeds.proto", fileDescriptor_0cd2be2dc9fdf3b5)
}

var fileDescriptor_0cd2be2dc9fdf3b5 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x36, 0x1b, 0xa5, 0x13, 0x45, 0x59, 0x8a, 0x2e, 0x01, 0x37, 0xa1, 0xa7, 0x1c, 0xea,
	0x2c, 0x16, 0xf1, 0x28, 0x18, 0x4b, 0xa4, 0x22, 0xa2, 0x49, 0x45, 0xf0, 0x60, 0x98, 0x9d, 0x7d,
	0xdd, 0x9d, 0x64, 0xb3, 0x33, 0x99, 0x99, 0xf8, 0xe3, 0x3f, 0x10, 0x4f, 0xfe, 0x59, 0x39, 0xf6,
	0xd8, 0x53, 0xd0, 0xe4, 0xe0, 0xff, 0xd0, 0x93, 0xcc, 0x64, 0x62, 0x43, 0x51, 0xf4, 0xd0, 0x5e,
	0x92, 0xc7, 0x9b, 0xf7, 0x7d, 0x33, 0xef, 0xfb, 0x3e, 0x16, 0xed, 0xa9, 0x49, 0x11, 0xc3, 0x27,
	0xa0, 0xac, 0x3c, 0x96, 0x44, 0x24, 0xb1, 0x90, 0x9c, 0x82, 0x52, 0x5c, 0xaa, 0x01, 0xcd, 0x49,
	0x99, 0xc1, 0x31, 0x40, 0xaa, 0xb0, 0x90, 0x5c, 0xf3, 0x20, 0xa4, 0x9c, 0x8e, 0x24, 0x27, 0x34,
	0xc7, 0x6a, 0x52, 0xe0, 0x94, 0x29, 0xad, 0x26, 0x85, 0x9c, 0x96, 0x8d, 0x3b, 0x43, 0x9e, 0xa8,
	0xd8, 0xfc, 0x88, 0xc4, 0xfe, 0xad, 0x10, 0x8d, 0xc0, 0x4e, 0x8b, 0x24, 0x4e, 0x89, 0x26, 0xae,
	0x17, 0x4e, 0x35, 0x2b, 0xe2, 0xbc, 0xa0, 0xb1, 0x66, 0x63, 0x50, 0x9a, 0x8c, 0x85, 0x3b, 0xd9,
	0xc9, 0x78, 0xc6, 0x6d, 0x19, 0x9b, 0x6a, 0xd5, 0xdd, 0xfd, 0xe2, 0xa3, 0x9d, 0xa7, 0xf6, 0x2d,
	0x4f, 0xb2, 0x4c, 0x42, 0x46, 0x34, 0x97, 0x7d, 0x01, 0x34, 0x38, 0x42, 0xd7, 0x3f, 0x12, 0x4d,
	0x73, 0x50, 0xa1, 0xd7, 0xaa, 0xb6, 0xeb, 0xfb, 0x0f, 0xf1, 0xdf, 0x1e, 0x88, 0xff, 0x44, 0x80,
	0xdf, 0x1a, 0x74, 0xc7, 0x9f, 0xcd, 0x9b, 0x95, 0xde, 0x9a, 0x2a, 0xe8, 0x22, 0xdf, 0xec, 0x1c,
	0x6e, 0xb5, 0xbc, 0x76, 0x7d, 0x7f, 0xef, 0x02, 0xa5, 0xdd, 0x6d, 0xb5, 0xa7, 0xe3, 0x34, 0xc3,
	0x07, 0xa0, 0x09, 0x2b, 0x94, 0xa3, 0xb2, 0xf8, 0x00, 0x10, 0x9a, 0x2a, 0x90, 0x03, 0xbb, 0x44,
	0x58, 0x6d, 0x79, 0xed, 0xed, 0x4e, 0xd7, 0x9c, 0x9f, 0xcd, 0x9b, 0x8f, 0x33, 0xa6, 0xf3, 0x69,
	0x82, 0x29, 0x1f, 0xc7, 0xbf, 0xf9, 0xd3, 0xe4, 0xbc, 0x8e, 0xc5, 0x28, 0x8b, 0x15, 0xd0, 0xa9,
	0x64, 0xfa, 0x33, 0xee, 0xbf, 0x7e, 0xf1, 0x46, 0x81, 0x2c, 0xc9, 0x18, 0x5e, 0x19, 0xb6, 0xde,
	0xb6, 0x61, 0xb6, 0x65, 0xf0, 0x1e, 0x5d, 0x1b, 0xf2, 0x64, 0xc0, 0xd2, 0xd0, 0x6f, 0x79, 0xed,
	0x6a, 0xe7, 0x99, 0xb9, 0x62, 0x31, 0x6f, 0xd6, 0x9e, 0xf3, 0xe4, 0xf0, 0xe0, 0x6c, 0xde, 0x7c,
	0xf4, 0x5f, 0x77, 0x6d, 0xf8, 0x87, 0x2d, 0xb2, 0x57, 0x1b, 0xf2, 0xe4, 0x30, 0x6d, 0x7c, 0xf5,
	0x50, 0xcd, 0xea, 0x14, 0xbc, 0x44, 0xb7, 0x59, 0xc9, 0x34, 0x23, 0xc5, 0x40, 0x82, 0xe2, 0xc5,
	0x07, 0x48, 0x43, 0xcf, 0x8a, 0x74, 0x6f, 0x43, 0x24, 0x63, 0x2e, 0xce, 0x0b, 0x8a, 0x8f, 0xd6,
	0xe6, 0x3a, 0x55, 0x6e, 0x39, 0x70, 0xcf, 0x61, 0x83, 0x07, 0xc8, 0x57, 0x82, 0x94, 0x4e, 0xe8,
	0xbb, 0x1b, 0x1c, 0x2e, 0x34, 0xb8, 0x2f, 0x48, 0xb9, 0xd6, 0xd4, 0x8c, 0xee, 0xfe, 0xdc, 0x42,
	0xc1, 0x4a, 0xf5, 0xae, 0xe4, 0xa5, 0x66, 0xb0, 0x0a, 0x42, 0x07, 0xdd, 0xd4, 0x92, 0xd0, 0x11,
	0xa4, 0x03, 0x33, 0xb6, 0x8e, 0xc3, 0x3f, 0x28, 0x6f, 0x38, 0x8c, 0x69, 0x5d, 0x9e, 0xed, 0xe7,
	0x7e, 0x54, 0xaf, 0xc2, 0x8f, 0x0b, 0xb1, 0xf2, 0xaf, 0x28, 0x56, 0x9d, 0xfb, 0xb3, 0x1f, 0x51,
	0x65, 0xb6, 0x88, 0xbc, 0x93, 0x45, 0xe4, 0x9d, 0x2e, 0x22, 0xef, 0xfb, 0x22, 0xf2, 0xbe, 0x2d,
	0xa3, 0xca, 0xc9, 0x32, 0xaa, 0x9c, 0x2e, 0xa3, 0xca, 0xbb, 0xfa, 0xc6, 0x67, 0xe3, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x49, 0x5b, 0x15, 0x71, 0x48, 0x04, 0x00, 0x00,
}

func (m *ChangeAggregatorSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeAggregatorSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeAggregatorSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(m.JobID))
	i--
	dAtA[i] = 0x20
	i -= len(m.UserProto)
	copy(dAtA[i:], m.UserProto)
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(len(m.UserProto)))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Feed.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Watches) > 0 {
		for iNdEx := len(m.Watches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Watches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChangeAggregatorSpec_Watch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeAggregatorSpec_Watch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeAggregatorSpec_Watch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Span.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.InitialResolved.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ChangeFrontierSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeFrontierSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeFrontierSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.UserProto)
	copy(dAtA[i:], m.UserProto)
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(len(m.UserProto)))
	i--
	dAtA[i] = 0x22
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(m.JobID))
	i--
	dAtA[i] = 0x18
	{
		size, err := m.Feed.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.TrackedSpans) > 0 {
		for iNdEx := len(m.TrackedSpans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TrackedSpans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProcessorsChangefeeds(dAtA []byte, offset int, v uint64) int {
	offset -= sovProcessorsChangefeeds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChangeAggregatorSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Watches) > 0 {
		for _, e := range m.Watches {
			l = e.Size()
			n += 1 + l + sovProcessorsChangefeeds(uint64(l))
		}
	}
	l = m.Feed.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	l = len(m.UserProto)
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	n += 1 + sovProcessorsChangefeeds(uint64(m.JobID))
	return n
}

func (m *ChangeAggregatorSpec_Watch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InitialResolved.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	l = m.Span.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	return n
}

func (m *ChangeFrontierSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TrackedSpans) > 0 {
		for _, e := range m.TrackedSpans {
			l = e.Size()
			n += 1 + l + sovProcessorsChangefeeds(uint64(l))
		}
	}
	l = m.Feed.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	n += 1 + sovProcessorsChangefeeds(uint64(m.JobID))
	l = len(m.UserProto)
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	return n
}

func sovProcessorsChangefeeds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProcessorsChangefeeds(x uint64) (n int) {
	return sovProcessorsChangefeeds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChangeAggregatorSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsChangefeeds
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
			return fmt.Errorf("proto: ChangeAggregatorSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeAggregatorSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Watches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Watches = append(m.Watches, ChangeAggregatorSpec_Watch{})
			if err := m.Watches[len(m.Watches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Feed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserProto", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserProto = github_com_cockroachdb_cockroach_pkg_security.SQLUsernameProto(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			m.JobID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobID |= github_com_cockroachdb_cockroach_pkg_jobs_jobspb.JobID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsChangefeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
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
func (m *ChangeAggregatorSpec_Watch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsChangefeeds
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
			return fmt.Errorf("proto: Watch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Watch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialResolved", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialResolved.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Span", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Span.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsChangefeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
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
func (m *ChangeFrontierSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsChangefeeds
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
			return fmt.Errorf("proto: ChangeFrontierSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeFrontierSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackedSpans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrackedSpans = append(m.TrackedSpans, roachpb.Span{})
			if err := m.TrackedSpans[len(m.TrackedSpans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Feed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			m.JobID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobID |= github_com_cockroachdb_cockroach_pkg_jobs_jobspb.JobID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserProto", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserProto = github_com_cockroachdb_cockroach_pkg_security.SQLUsernameProto(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsChangefeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
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
func skipProcessorsChangefeeds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProcessorsChangefeeds
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
					return 0, ErrIntOverflowProcessorsChangefeeds
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
					return 0, ErrIntOverflowProcessorsChangefeeds
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
				return 0, ErrInvalidLengthProcessorsChangefeeds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProcessorsChangefeeds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProcessorsChangefeeds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProcessorsChangefeeds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProcessorsChangefeeds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProcessorsChangefeeds = fmt.Errorf("proto: unexpected end of group")
)
