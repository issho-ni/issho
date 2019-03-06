// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: youji.proto

package youji

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	ninshou "github.com/issho-ni/issho/api/ninshou"
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

type NewTodo struct {
	Text   string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (m *NewTodo) Reset()         { *m = NewTodo{} }
func (m *NewTodo) String() string { return proto.CompactTextString(m) }
func (*NewTodo) ProtoMessage()    {}
func (*NewTodo) Descriptor() ([]byte, []int) {
	return fileDescriptor_544de1e744607021, []int{0}
}
func (m *NewTodo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NewTodo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NewTodo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NewTodo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTodo.Merge(m, src)
}
func (m *NewTodo) XXX_Size() int {
	return m.Size()
}
func (m *NewTodo) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTodo.DiscardUnknown(m)
}

var xxx_messageInfo_NewTodo proto.InternalMessageInfo

func (m *NewTodo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *NewTodo) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type Todo struct {
	Id   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string        `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Done bool          `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	User *ninshou.User `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_544de1e744607021, []int{1}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return m.Size()
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Todo) GetUser() *ninshou.User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*NewTodo)(nil), "youji.NewTodo")
	proto.RegisterType((*Todo)(nil), "youji.Todo")
}

func init() { proto.RegisterFile("youji.proto", fileDescriptor_544de1e744607021) }

var fileDescriptor_544de1e744607021 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xae, 0xcc, 0x2f, 0xcd,
	0xca, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x78, 0xf3, 0x32, 0xf3,
	0x8a, 0x33, 0xf2, 0x4b, 0x21, 0xa2, 0x4a, 0xa6, 0x5c, 0xec, 0x7e, 0xa9, 0xe5, 0x21, 0xf9, 0x29,
	0xf9, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x90, 0x18, 0x17, 0x5b, 0x69, 0x71, 0x6a, 0x91, 0xa7, 0x8b, 0x04, 0x13, 0x58, 0x14,
	0xca, 0x53, 0x4a, 0xe4, 0x62, 0x01, 0xeb, 0xe1, 0xe3, 0x62, 0xca, 0x4c, 0x81, 0xea, 0x60, 0xca,
	0x4c, 0x81, 0x9b, 0xc1, 0x84, 0x64, 0x86, 0x10, 0x17, 0x4b, 0x4a, 0x7e, 0x5e, 0xaa, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0x47, 0x10, 0x98, 0x2d, 0xa4, 0xc8, 0xc5, 0x02, 0x32, 0x49, 0x82, 0x45, 0x81,
	0x51, 0x83, 0xdb, 0x88, 0x57, 0x0f, 0xe6, 0xa8, 0xd0, 0xe2, 0xd4, 0xa2, 0x20, 0xb0, 0x94, 0x91,
	0x09, 0x17, 0x6b, 0x24, 0xc8, 0xc5, 0x42, 0xda, 0x5c, 0x5c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9,
	0x10, 0x1b, 0xf5, 0x20, 0x9e, 0x82, 0xba, 0x5a, 0x8a, 0x1b, 0xca, 0x07, 0x71, 0x94, 0x18, 0x9c,
	0x6c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x39, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xb3, 0xb8, 0x38, 0x23, 0x5f, 0x37, 0x2f, 0x13,
	0xc2, 0xd0, 0x4f, 0x2c, 0xc8, 0xd4, 0x07, 0x9b, 0x92, 0xc4, 0x06, 0x0e, 0x15, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0xb2, 0x0d, 0xfb, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// YoujiClient is the client API for Youji service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YoujiClient interface {
	CreateTodo(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todo, error)
}

type youjiClient struct {
	cc *grpc.ClientConn
}

func NewYoujiClient(cc *grpc.ClientConn) YoujiClient {
	return &youjiClient{cc}
}

func (c *youjiClient) CreateTodo(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/youji.Youji/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoujiServer is the server API for Youji service.
type YoujiServer interface {
	CreateTodo(context.Context, *NewTodo) (*Todo, error)
}

func RegisterYoujiServer(s *grpc.Server, srv YoujiServer) {
	s.RegisterService(&_Youji_serviceDesc, srv)
}

func _Youji_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTodo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoujiServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/youji.Youji/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoujiServer).CreateTodo(ctx, req.(*NewTodo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Youji_serviceDesc = grpc.ServiceDesc{
	ServiceName: "youji.Youji",
	HandlerType: (*YoujiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _Youji_CreateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "youji.proto",
}

func (m *NewTodo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewTodo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintYouji(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	if len(m.UserID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintYouji(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	return i, nil
}

func (m *Todo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Todo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintYouji(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Text) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintYouji(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	if m.Done {
		dAtA[i] = 0x18
		i++
		if m.Done {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.User != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintYouji(dAtA, i, uint64(m.User.Size()))
		n1, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintYouji(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NewTodo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovYouji(uint64(l))
	}
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovYouji(uint64(l))
	}
	return n
}

func (m *Todo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovYouji(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovYouji(uint64(l))
	}
	if m.Done {
		n += 2
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovYouji(uint64(l))
	}
	return n
}

func sovYouji(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozYouji(x uint64) (n int) {
	return sovYouji(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NewTodo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYouji
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
			return fmt.Errorf("proto: NewTodo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewTodo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYouji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthYouji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYouji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthYouji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipYouji(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYouji
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
func (m *Todo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYouji
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
			return fmt.Errorf("proto: Todo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Todo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYouji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthYouji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYouji
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthYouji
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Done", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYouji
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
			m.Done = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYouji
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
				return ErrInvalidLengthYouji
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &ninshou.User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipYouji(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYouji
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
func skipYouji(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowYouji
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
					return 0, ErrIntOverflowYouji
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
					return 0, ErrIntOverflowYouji
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
				return 0, ErrInvalidLengthYouji
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowYouji
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
				next, err := skipYouji(dAtA[start:])
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
	ErrInvalidLengthYouji = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowYouji   = fmt.Errorf("proto: integer overflow")
)
