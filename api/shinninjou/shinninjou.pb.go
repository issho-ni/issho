// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shinninjou/shinninjou.proto

package shinninjou

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/issho-ni/issho/api/common"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CredentialType int32

const (
	CredentialType_PASSWORD CredentialType = 0
)

var CredentialType_name = map[int32]string{
	0: "PASSWORD",
}

var CredentialType_value = map[string]int32{
	"PASSWORD": 0,
}

func (x CredentialType) String() string {
	return proto.EnumName(CredentialType_name, int32(x))
}

func (CredentialType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6fa052e50045eab0, []int{0}
}

type Credential struct {
	UserID         *common.UUID   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	CredentialType CredentialType `protobuf:"varint,2,opt,name=credentialType,proto3,enum=shinninjou.CredentialType" json:"credentialType,omitempty"`
	Credential     []byte         `protobuf:"bytes,3,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa052e50045eab0, []int{0}
}
func (m *Credential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return m.Size()
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

func (m *Credential) GetUserID() *common.UUID {
	if m != nil {
		return m.UserID
	}
	return nil
}

func (m *Credential) GetCredentialType() CredentialType {
	if m != nil {
		return m.CredentialType
	}
	return CredentialType_PASSWORD
}

func (m *Credential) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

type CredentialResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *CredentialResponse) Reset()         { *m = CredentialResponse{} }
func (m *CredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialResponse) ProtoMessage()    {}
func (*CredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa052e50045eab0, []int{1}
}
func (m *CredentialResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialResponse.Merge(m, src)
}
func (m *CredentialResponse) XXX_Size() int {
	return m.Size()
}
func (m *CredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialResponse proto.InternalMessageInfo

func (m *CredentialResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("shinninjou.CredentialType", CredentialType_name, CredentialType_value)
	proto.RegisterType((*Credential)(nil), "shinninjou.Credential")
	proto.RegisterType((*CredentialResponse)(nil), "shinninjou.CredentialResponse")
}

func init() { proto.RegisterFile("shinninjou/shinninjou.proto", fileDescriptor_6fa052e50045eab0) }

var fileDescriptor_6fa052e50045eab0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xce, 0xc8, 0xcc,
	0xcb, 0xcb, 0xcc, 0xcb, 0xca, 0x2f, 0xd5, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x52, 0xc2, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x10, 0x0a, 0xa2, 0x40,
	0x69, 0x1a, 0x23, 0x17, 0x97, 0x73, 0x51, 0x6a, 0x4a, 0x6a, 0x5e, 0x49, 0x66, 0x62, 0x8e, 0x90,
	0x0a, 0x17, 0x5b, 0x69, 0x71, 0x6a, 0x91, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11,
	0x8f, 0x1e, 0x54, 0x75, 0x68, 0xa8, 0xa7, 0x4b, 0x10, 0x54, 0x4e, 0xc8, 0x89, 0x8b, 0x2f, 0x19,
	0xae, 0x27, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x4a, 0x0f, 0xc9,
	0x01, 0xce, 0x28, 0x2a, 0x82, 0xd0, 0x74, 0x08, 0xc9, 0x71, 0x71, 0x21, 0x44, 0x24, 0x98, 0x15,
	0x18, 0x35, 0x78, 0x82, 0x90, 0x44, 0x94, 0xf4, 0xb8, 0x84, 0x10, 0x26, 0x04, 0xa5, 0x16, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0x17, 0x97, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x83,
	0x1d, 0xc8, 0x11, 0x04, 0xe3, 0x6a, 0xc9, 0x71, 0xf1, 0xa1, 0xda, 0x28, 0xc4, 0xc3, 0xc5, 0x11,
	0xe0, 0x18, 0x1c, 0x1c, 0xee, 0x1f, 0xe4, 0x22, 0xc0, 0x60, 0xb4, 0x8a, 0x91, 0x8b, 0x2b, 0x18,
	0xee, 0x3a, 0x21, 0x1f, 0x2e, 0x01, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0x24, 0xcf, 0x8b, 0x61,
	0x77, 0xbe, 0x94, 0x1c, 0x76, 0x71, 0x98, 0xa3, 0x94, 0x18, 0x84, 0xfc, 0xb8, 0x84, 0xc2, 0x12,
	0x73, 0x32, 0x53, 0xa8, 0x64, 0x9e, 0x93, 0xd3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0x69, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0x81, 0xa2, 0x45, 0x3f, 0xb3, 0xb8, 0x38,
	0x23, 0x5f, 0x37, 0x2f, 0x13, 0xc2, 0xd0, 0x4f, 0x2c, 0xc8, 0x44, 0x4a, 0x00, 0x49, 0x6c, 0xe0,
	0x08, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x9b, 0x9c, 0xdc, 0x20, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShinninjouClient is the client API for Shinninjou service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShinninjouClient interface {
	CreateCredential(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error)
	ValidateCredential(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error)
}

type shinninjouClient struct {
	cc *grpc.ClientConn
}

func NewShinninjouClient(cc *grpc.ClientConn) ShinninjouClient {
	return &shinninjouClient{cc}
}

func (c *shinninjouClient) CreateCredential(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error) {
	out := new(CredentialResponse)
	err := c.cc.Invoke(ctx, "/shinninjou.Shinninjou/CreateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shinninjouClient) ValidateCredential(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*CredentialResponse, error) {
	out := new(CredentialResponse)
	err := c.cc.Invoke(ctx, "/shinninjou.Shinninjou/ValidateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShinninjouServer is the server API for Shinninjou service.
type ShinninjouServer interface {
	CreateCredential(context.Context, *Credential) (*CredentialResponse, error)
	ValidateCredential(context.Context, *Credential) (*CredentialResponse, error)
}

func RegisterShinninjouServer(s *grpc.Server, srv ShinninjouServer) {
	s.RegisterService(&_Shinninjou_serviceDesc, srv)
}

func _Shinninjou_CreateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShinninjouServer).CreateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shinninjou.Shinninjou/CreateCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShinninjouServer).CreateCredential(ctx, req.(*Credential))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shinninjou_ValidateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShinninjouServer).ValidateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shinninjou.Shinninjou/ValidateCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShinninjouServer).ValidateCredential(ctx, req.(*Credential))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shinninjou_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shinninjou.Shinninjou",
	HandlerType: (*ShinninjouServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCredential",
			Handler:    _Shinninjou_CreateCredential_Handler,
		},
		{
			MethodName: "ValidateCredential",
			Handler:    _Shinninjou_ValidateCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shinninjou/shinninjou.proto",
}

func (m *Credential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credential) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserID != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintShinninjou(dAtA, i, uint64(m.UserID.Size()))
		n1, err := m.UserID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.CredentialType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShinninjou(dAtA, i, uint64(m.CredentialType))
	}
	if len(m.Credential) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintShinninjou(dAtA, i, uint64(len(m.Credential)))
		i += copy(dAtA[i:], m.Credential)
	}
	return i, nil
}

func (m *CredentialResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintShinninjou(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Credential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserID != nil {
		l = m.UserID.Size()
		n += 1 + l + sovShinninjou(uint64(l))
	}
	if m.CredentialType != 0 {
		n += 1 + sovShinninjou(uint64(m.CredentialType))
	}
	l = len(m.Credential)
	if l > 0 {
		n += 1 + l + sovShinninjou(uint64(l))
	}
	return n
}

func (m *CredentialResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovShinninjou(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozShinninjou(x uint64) (n int) {
	return sovShinninjou(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShinninjou
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Credential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShinninjou
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShinninjou
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UserID == nil {
				m.UserID = &common.UUID{}
			}
			if err := m.UserID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialType", wireType)
			}
			m.CredentialType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShinninjou
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CredentialType |= (CredentialType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credential", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShinninjou
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthShinninjou
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credential = append(m.Credential[:0], dAtA[iNdEx:postIndex]...)
			if m.Credential == nil {
				m.Credential = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShinninjou(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShinninjou
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
func (m *CredentialResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShinninjou
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CredentialResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShinninjou
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipShinninjou(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShinninjou
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
func skipShinninjou(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShinninjou
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
					return 0, ErrIntOverflowShinninjou
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShinninjou
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthShinninjou
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowShinninjou
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipShinninjou(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthShinninjou = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShinninjou   = fmt.Errorf("proto: integer overflow")
)
