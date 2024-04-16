// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/protobuf.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Nationality          string   `protobuf:"bytes,3,opt,name=nationality,proto3" json:"nationality,omitempty"`
	Zip                  int32    `protobuf:"varint,4,opt,name=zip,proto3" json:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *User) GetZip() int32 {
	if m != nil {
		return m.Zip
	}
	return 0
}

type FetchUserRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchUserRequest) Reset()         { *m = FetchUserRequest{} }
func (m *FetchUserRequest) String() string { return proto.CompactTextString(m) }
func (*FetchUserRequest) ProtoMessage()    {}
func (*FetchUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{1}
}
func (m *FetchUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchUserRequest.Unmarshal(m, b)
}
func (m *FetchUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchUserRequest.Marshal(b, m, deterministic)
}
func (m *FetchUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchUserRequest.Merge(m, src)
}
func (m *FetchUserRequest) XXX_Size() int {
	return xxx_messageInfo_FetchUserRequest.Size(m)
}
func (m *FetchUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchUserRequest proto.InternalMessageInfo

func (m *FetchUserRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type FetchUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchUserResponse) Reset()         { *m = FetchUserResponse{} }
func (m *FetchUserResponse) String() string { return proto.CompactTextString(m) }
func (*FetchUserResponse) ProtoMessage()    {}
func (*FetchUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{2}
}
func (m *FetchUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchUserResponse.Unmarshal(m, b)
}
func (m *FetchUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchUserResponse.Marshal(b, m, deterministic)
}
func (m *FetchUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchUserResponse.Merge(m, src)
}
func (m *FetchUserResponse) XXX_Size() int {
	return xxx_messageInfo_FetchUserResponse.Size(m)
}
func (m *FetchUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchUserResponse proto.InternalMessageInfo

func (m *FetchUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{3}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{4}
}
func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{5}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{6}
}
func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{7}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type DeleteUserResponse struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac732507fd34a3d2, []int{8}
}
func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "protobuf.User")
	proto.RegisterType((*FetchUserRequest)(nil), "protobuf.FetchUserRequest")
	proto.RegisterType((*FetchUserResponse)(nil), "protobuf.FetchUserResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "protobuf.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "protobuf.CreateUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "protobuf.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "protobuf.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "protobuf.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "protobuf.DeleteUserResponse")
}

func init() { proto.RegisterFile("protobuf/protobuf.proto", fileDescriptor_ac732507fd34a3d2) }

var fileDescriptor_ac732507fd34a3d2 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x53, 0xa8, 0x46, 0xa6, 0x89, 0x81, 0xb9, 0xd8, 0x54, 0x0f, 0x4d, 0xa3, 0x86, 0x13,
	0x26, 0x78, 0xd0, 0xbb, 0x44, 0xee, 0x35, 0xdc, 0x2d, 0x30, 0xc6, 0x4d, 0xb0, 0xad, 0xbb, 0x5b,
	0x13, 0xf9, 0x51, 0xfe, 0x46, 0xb3, 0x2b, 0xb0, 0xa3, 0xdb, 0x83, 0x70, 0x9b, 0xf6, 0xbd, 0x79,
	0xaf, 0xfd, 0x06, 0xce, 0x6a, 0x59, 0xe9, 0x6a, 0xde, 0xbc, 0xdc, 0x6c, 0x87, 0x91, 0x1d, 0xf0,
	0x64, 0xfb, 0x9c, 0x3d, 0x43, 0x38, 0x53, 0x24, 0xb1, 0x0f, 0xdd, 0x46, 0x2c, 0xe3, 0x20, 0x0d,
	0x86, 0x47, 0xb9, 0x19, 0x11, 0x21, 0x2c, 0x8b, 0x37, 0x8a, 0x3b, 0x69, 0x30, 0xec, 0xe5, 0x76,
	0xc6, 0x14, 0xa2, 0xb2, 0xd0, 0xa2, 0x2a, 0x8b, 0x95, 0xd0, 0x9f, 0x71, 0xd7, 0x4a, 0xfc, 0x95,
	0xc9, 0x59, 0x8b, 0x3a, 0x0e, 0x7f, 0x72, 0xd6, 0xa2, 0xce, 0x2e, 0xa1, 0xff, 0x48, 0x7a, 0xf1,
	0x6a, 0x6a, 0x72, 0x7a, 0x6f, 0x48, 0x69, 0xbf, 0x2d, 0xbb, 0x83, 0x01, 0x73, 0xa9, 0xba, 0x2a,
	0x15, 0x61, 0x06, 0x61, 0xa3, 0x48, 0x5a, 0x5f, 0x34, 0x3e, 0x1d, 0xed, 0xfe, 0xc2, 0xba, 0xac,
	0x66, 0x16, 0x1f, 0x24, 0x15, 0x9a, 0x78, 0xfe, 0x7f, 0x16, 0xef, 0x01, 0xf9, 0xe2, 0x7e, 0x95,
	0xb3, 0x7a, 0x79, 0x58, 0x25, 0x5f, 0xdc, 0xa3, 0xf2, 0x0a, 0x06, 0x13, 0x5a, 0xd1, 0xef, 0x4a,
	0x9f, 0xe2, 0x35, 0x20, 0xb7, 0x6d, 0x0a, 0x3c, 0xdf, 0xf8, 0xab, 0x03, 0x91, 0xb1, 0x3c, 0x91,
	0xfc, 0x10, 0x0b, 0xc2, 0x09, 0xf4, 0x76, 0xf4, 0x31, 0x71, 0x5f, 0xf0, 0xf7, 0x70, 0xc9, 0x79,
	0xab, 0xb6, 0xe9, 0x99, 0x02, 0x38, 0xa2, 0xc8, 0xac, 0xde, 0x81, 0x92, 0x8b, 0x76, 0xd1, 0x05,
	0x39, 0x4e, 0x3c, 0xc8, 0xc3, 0xce, 0x83, 0x5a, 0xd0, 0x4e, 0x01, 0x1c, 0x0f, 0x1e, 0xe4, 0xc1,
	0xe4, 0x41, 0x3e, 0xc2, 0xf9, 0xb1, 0x15, 0x6f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x74,
	0x65, 0x8c, 0x52, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	FetchUser(ctx context.Context, in *FetchUserRequest, opts ...grpc.CallOption) (*FetchUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FetchUser(ctx context.Context, in *FetchUserRequest, opts ...grpc.CallOption) (*FetchUserResponse, error) {
	out := new(FetchUserResponse)
	err := c.cc.Invoke(ctx, "/protobuf.UserService/FetchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/protobuf.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/protobuf.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/protobuf.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	FetchUser(context.Context, *FetchUserRequest) (*FetchUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) FetchUser(ctx context.Context, req *FetchUserRequest) (*FetchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUser not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_FetchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FetchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.UserService/FetchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FetchUser(ctx, req.(*FetchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchUser",
			Handler:    _UserService_FetchUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/protobuf.proto",
}
