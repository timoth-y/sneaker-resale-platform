// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: interact.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	EntityID string `protobuf:"bytes,2,opt,name=entityID,proto3" json:"entityID,omitempty"`
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_interact_proto_rawDescGZIP(), []int{0}
}

func (x *LikeRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *LikeRequest) GetEntityID() string {
	if x != nil {
		return x.EntityID
	}
	return ""
}

type LikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LikeResponse) Reset() {
	*x = LikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeResponse) ProtoMessage() {}

func (x *LikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeResponse.ProtoReflect.Descriptor instead.
func (*LikeResponse) Descriptor() ([]byte, []int) {
	return file_interact_proto_rawDescGZIP(), []int{1}
}

func (x *LikeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LikeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_interact_proto protoreflect.FileDescriptor

var file_interact_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x22, 0x3e, 0x0a, 0x0c, 0x4c, 0x69,
	0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x75, 0x0a, 0x0f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69,
	0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x05, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interact_proto_rawDescOnce sync.Once
	file_interact_proto_rawDescData = file_interact_proto_rawDesc
)

func file_interact_proto_rawDescGZIP() []byte {
	file_interact_proto_rawDescOnce.Do(func() {
		file_interact_proto_rawDescData = protoimpl.X.CompressGZIP(file_interact_proto_rawDescData)
	})
	return file_interact_proto_rawDescData
}

var file_interact_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interact_proto_goTypes = []interface{}{
	(*LikeRequest)(nil),  // 0: proto.LikeRequest
	(*LikeResponse)(nil), // 1: proto.LikeResponse
}
var file_interact_proto_depIdxs = []int32{
	0, // 0: proto.InteractService.Like:input_type -> proto.LikeRequest
	0, // 1: proto.InteractService.Unlike:input_type -> proto.LikeRequest
	1, // 2: proto.InteractService.Like:output_type -> proto.LikeResponse
	1, // 3: proto.InteractService.Unlike:output_type -> proto.LikeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interact_proto_init() }
func file_interact_proto_init() {
	if File_interact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interact_proto_goTypes,
		DependencyIndexes: file_interact_proto_depIdxs,
		MessageInfos:      file_interact_proto_msgTypes,
	}.Build()
	File_interact_proto = out.File
	file_interact_proto_rawDesc = nil
	file_interact_proto_goTypes = nil
	file_interact_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InteractServiceClient is the client API for InteractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InteractServiceClient interface {
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
	Unlike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
}

type interactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractServiceClient(cc grpc.ClientConnInterface) InteractServiceClient {
	return &interactServiceClient{cc}
}

func (c *interactServiceClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	out := new(LikeResponse)
	err := c.cc.Invoke(ctx, "/proto.InteractService/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactServiceClient) Unlike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	out := new(LikeResponse)
	err := c.cc.Invoke(ctx, "/proto.InteractService/Unlike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractServiceServer is the server API for InteractService service.
type InteractServiceServer interface {
	Like(context.Context, *LikeRequest) (*LikeResponse, error)
	Unlike(context.Context, *LikeRequest) (*LikeResponse, error)
}

// UnimplementedInteractServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInteractServiceServer struct {
}

func (*UnimplementedInteractServiceServer) Like(context.Context, *LikeRequest) (*LikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (*UnimplementedInteractServiceServer) Unlike(context.Context, *LikeRequest) (*LikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlike not implemented")
}

func RegisterInteractServiceServer(s *grpc.Server, srv InteractServiceServer) {
	s.RegisterService(&_InteractService_serviceDesc, srv)
}

func _InteractService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.InteractService/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractService_Unlike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractServiceServer).Unlike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.InteractService/Unlike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractServiceServer).Unlike(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InteractService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.InteractService",
	HandlerType: (*InteractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Like",
			Handler:    _InteractService_Like_Handler,
		},
		{
			MethodName: "Unlike",
			Handler:    _InteractService_Unlike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interact.proto",
}