// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: random.proto

package protocols

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// The request message containing the user's name.
type ProfileRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ProfileRequest) Reset()      { *m = ProfileRequest{} }
func (*ProfileRequest) ProtoMessage() {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee7c418402e21a3f, []int{0}
}
func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(m, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type ProfileReply struct {
	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (m *ProfileReply) Reset()      { *m = ProfileReply{} }
func (*ProfileReply) ProtoMessage() {}
func (*ProfileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee7c418402e21a3f, []int{1}
}
func (m *ProfileReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProfileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProfileReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProfileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileReply.Merge(m, src)
}
func (m *ProfileReply) XXX_Size() int {
	return m.Size()
}
func (m *ProfileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileReply proto.InternalMessageInfo

func (m *ProfileReply) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*ProfileRequest)(nil), "protocols.ProfileRequest")
	proto.RegisterType((*ProfileReply)(nil), "protocols.ProfileReply")
}

func init() { proto.RegisterFile("random.proto", fileDescriptor_ee7c418402e21a3f) }

var fileDescriptor_ee7c418402e21a3f = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0xc9, 0xf9, 0x39, 0xc5,
	0x52, 0xc2, 0x10, 0x89, 0x82, 0xa2, 0xfc, 0xb4, 0xcc, 0x9c, 0x54, 0x88, 0xbc, 0x92, 0x0a, 0x17,
	0x5f, 0x00, 0x44, 0x20, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f,
	0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x72, 0xe0, 0xe2, 0x81,
	0xab, 0x2a, 0xc8, 0xa9, 0x14, 0x32, 0xe0, 0x62, 0x87, 0x1a, 0x03, 0x56, 0xc6, 0x6d, 0x24, 0xa6,
	0x87, 0x6a, 0x38, 0x4c, 0x35, 0x4c, 0x99, 0x51, 0x00, 0x17, 0x57, 0x10, 0x58, 0x45, 0x66, 0x55,
	0x6a, 0x91, 0x90, 0x13, 0x17, 0x97, 0x7b, 0x6a, 0x09, 0x54, 0x91, 0x90, 0xa4, 0x1e, 0xdc, 0x91,
	0x7a, 0xa8, 0x8e, 0x91, 0x12, 0xc7, 0x26, 0x55, 0x90, 0x53, 0xa9, 0xc4, 0xe0, 0xe4, 0x79, 0xe1,
	0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c,
	0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x18, 0x6e, 0x9e, 0x35, 0x9c, 0x95, 0xc4, 0x06, 0x66,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43, 0xa3, 0xd0, 0xc4, 0x3b, 0x01, 0x00, 0x00,
}

func (this *ProfileRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProfileRequest)
	if !ok {
		that2, ok := that.(ProfileRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *ProfileReply) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProfileReply)
	if !ok {
		that2, ok := that.(ProfileReply)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Profile.Equal(that1.Profile) {
		return false
	}
	return true
}
func (this *ProfileRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&protocols.ProfileRequest{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ProfileReply) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&protocols.ProfileReply{")
	if this.Profile != nil {
		s = append(s, "Profile: "+fmt.Sprintf("%#v", this.Profile)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRandom(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RandomizerClient is the client API for Randomizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RandomizerClient interface {
	GetProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
}

type randomizerClient struct {
	cc *grpc.ClientConn
}

func NewRandomizerClient(cc *grpc.ClientConn) RandomizerClient {
	return &randomizerClient{cc}
}

func (c *randomizerClient) GetProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/protocols.Randomizer/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomizerServer is the server API for Randomizer service.
type RandomizerServer interface {
	GetProfile(context.Context, *ProfileRequest) (*ProfileReply, error)
}

// UnimplementedRandomizerServer can be embedded to have forward compatible implementations.
type UnimplementedRandomizerServer struct {
}

func (*UnimplementedRandomizerServer) GetProfile(ctx context.Context, req *ProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}

func RegisterRandomizerServer(s *grpc.Server, srv RandomizerServer) {
	s.RegisterService(&_Randomizer_serviceDesc, srv)
}

func _Randomizer_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomizerServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocols.Randomizer/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomizerServer).GetProfile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Randomizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocols.Randomizer",
	HandlerType: (*RandomizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _Randomizer_GetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "random.proto",
}

func (m *ProfileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProfileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProfileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRandom(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProfileReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProfileReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProfileReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Profile != nil {
		{
			size, err := m.Profile.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRandom(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRandom(dAtA []byte, offset int, v uint64) int {
	offset -= sovRandom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProfileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRandom(uint64(l))
	}
	return n
}

func (m *ProfileReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Profile != nil {
		l = m.Profile.Size()
		n += 1 + l + sovRandom(uint64(l))
	}
	return n
}

func sovRandom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRandom(x uint64) (n int) {
	return sovRandom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ProfileRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProfileRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProfileReply) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProfileReply{`,
		`Profile:` + strings.Replace(fmt.Sprintf("%v", this.Profile), "Profile", "Profile", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRandom(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ProfileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandom
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
			return fmt.Errorf("proto: ProfileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProfileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRandom
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRandom
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
func (m *ProfileReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandom
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
			return fmt.Errorf("proto: ProfileReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProfileReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandom
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
				return ErrInvalidLengthRandom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRandom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Profile == nil {
				m.Profile = &Profile{}
			}
			if err := m.Profile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRandom
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRandom
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
func skipRandom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRandom
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
					return 0, ErrIntOverflowRandom
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
					return 0, ErrIntOverflowRandom
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
				return 0, ErrInvalidLengthRandom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRandom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRandom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRandom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRandom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRandom = fmt.Errorf("proto: unexpected end of group")
)