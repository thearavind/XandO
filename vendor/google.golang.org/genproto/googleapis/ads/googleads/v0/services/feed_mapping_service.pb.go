// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/feed_mapping_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [FeedMappingService.GetFeedMapping][google.ads.googleads.v0.services.FeedMappingService.GetFeedMapping].
type GetFeedMappingRequest struct {
	// The resource name of the feed mapping to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedMappingRequest) Reset()         { *m = GetFeedMappingRequest{} }
func (m *GetFeedMappingRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedMappingRequest) ProtoMessage()    {}
func (*GetFeedMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_service_0f3de8fce9e30788, []int{0}
}
func (m *GetFeedMappingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedMappingRequest.Unmarshal(m, b)
}
func (m *GetFeedMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedMappingRequest.Marshal(b, m, deterministic)
}
func (dst *GetFeedMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedMappingRequest.Merge(dst, src)
}
func (m *GetFeedMappingRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedMappingRequest.Size(m)
}
func (m *GetFeedMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedMappingRequest proto.InternalMessageInfo

func (m *GetFeedMappingRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedMappingService.MutateFeedMappings][google.ads.googleads.v0.services.FeedMappingService.MutateFeedMappings].
type MutateFeedMappingsRequest struct {
	// The ID of the customer whose feed mappings are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feed mappings.
	Operations []*FeedMappingOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedMappingsRequest) Reset()         { *m = MutateFeedMappingsRequest{} }
func (m *MutateFeedMappingsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingsRequest) ProtoMessage()    {}
func (*MutateFeedMappingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_service_0f3de8fce9e30788, []int{1}
}
func (m *MutateFeedMappingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingsRequest.Unmarshal(m, b)
}
func (m *MutateFeedMappingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateFeedMappingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingsRequest.Merge(dst, src)
}
func (m *MutateFeedMappingsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingsRequest.Size(m)
}
func (m *MutateFeedMappingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingsRequest proto.InternalMessageInfo

func (m *MutateFeedMappingsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedMappingsRequest) GetOperations() []*FeedMappingOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedMappingsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedMappingsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a feed mapping.
type FeedMappingOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedMappingOperation_Create
	//	*FeedMappingOperation_Remove
	Operation            isFeedMappingOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *FeedMappingOperation) Reset()         { *m = FeedMappingOperation{} }
func (m *FeedMappingOperation) String() string { return proto.CompactTextString(m) }
func (*FeedMappingOperation) ProtoMessage()    {}
func (*FeedMappingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_service_0f3de8fce9e30788, []int{2}
}
func (m *FeedMappingOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingOperation.Unmarshal(m, b)
}
func (m *FeedMappingOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingOperation.Marshal(b, m, deterministic)
}
func (dst *FeedMappingOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingOperation.Merge(dst, src)
}
func (m *FeedMappingOperation) XXX_Size() int {
	return xxx_messageInfo_FeedMappingOperation.Size(m)
}
func (m *FeedMappingOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingOperation proto.InternalMessageInfo

type isFeedMappingOperation_Operation interface {
	isFeedMappingOperation_Operation()
}

type FeedMappingOperation_Create struct {
	Create *resources.FeedMapping `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedMappingOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedMappingOperation_Create) isFeedMappingOperation_Operation() {}

func (*FeedMappingOperation_Remove) isFeedMappingOperation_Operation() {}

func (m *FeedMappingOperation) GetOperation() isFeedMappingOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedMappingOperation) GetCreate() *resources.FeedMapping {
	if x, ok := m.GetOperation().(*FeedMappingOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedMappingOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedMappingOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FeedMappingOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FeedMappingOperation_OneofMarshaler, _FeedMappingOperation_OneofUnmarshaler, _FeedMappingOperation_OneofSizer, []interface{}{
		(*FeedMappingOperation_Create)(nil),
		(*FeedMappingOperation_Remove)(nil),
	}
}

func _FeedMappingOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FeedMappingOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedMappingOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *FeedMappingOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("FeedMappingOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _FeedMappingOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FeedMappingOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.FeedMapping)
		err := b.DecodeMessage(msg)
		m.Operation = &FeedMappingOperation_Create{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &FeedMappingOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _FeedMappingOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FeedMappingOperation)
	// operation
	switch x := m.Operation.(type) {
	case *FeedMappingOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *FeedMappingOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for a feed mapping mutate.
type MutateFeedMappingsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedMappingResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MutateFeedMappingsResponse) Reset()         { *m = MutateFeedMappingsResponse{} }
func (m *MutateFeedMappingsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingsResponse) ProtoMessage()    {}
func (*MutateFeedMappingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_service_0f3de8fce9e30788, []int{3}
}
func (m *MutateFeedMappingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingsResponse.Unmarshal(m, b)
}
func (m *MutateFeedMappingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateFeedMappingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingsResponse.Merge(dst, src)
}
func (m *MutateFeedMappingsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingsResponse.Size(m)
}
func (m *MutateFeedMappingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingsResponse proto.InternalMessageInfo

func (m *MutateFeedMappingsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedMappingsResponse) GetResults() []*MutateFeedMappingResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mapping mutate.
type MutateFeedMappingResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedMappingResult) Reset()         { *m = MutateFeedMappingResult{} }
func (m *MutateFeedMappingResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedMappingResult) ProtoMessage()    {}
func (*MutateFeedMappingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_service_0f3de8fce9e30788, []int{4}
}
func (m *MutateFeedMappingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedMappingResult.Unmarshal(m, b)
}
func (m *MutateFeedMappingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedMappingResult.Marshal(b, m, deterministic)
}
func (dst *MutateFeedMappingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedMappingResult.Merge(dst, src)
}
func (m *MutateFeedMappingResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedMappingResult.Size(m)
}
func (m *MutateFeedMappingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedMappingResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedMappingResult proto.InternalMessageInfo

func (m *MutateFeedMappingResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedMappingRequest)(nil), "google.ads.googleads.v0.services.GetFeedMappingRequest")
	proto.RegisterType((*MutateFeedMappingsRequest)(nil), "google.ads.googleads.v0.services.MutateFeedMappingsRequest")
	proto.RegisterType((*FeedMappingOperation)(nil), "google.ads.googleads.v0.services.FeedMappingOperation")
	proto.RegisterType((*MutateFeedMappingsResponse)(nil), "google.ads.googleads.v0.services.MutateFeedMappingsResponse")
	proto.RegisterType((*MutateFeedMappingResult)(nil), "google.ads.googleads.v0.services.MutateFeedMappingResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedMappingServiceClient is the client API for FeedMappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedMappingServiceClient interface {
	// Returns the requested feed mapping in full detail.
	GetFeedMapping(ctx context.Context, in *GetFeedMappingRequest, opts ...grpc.CallOption) (*resources.FeedMapping, error)
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error)
}

type feedMappingServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedMappingServiceClient(cc *grpc.ClientConn) FeedMappingServiceClient {
	return &feedMappingServiceClient{cc}
}

func (c *feedMappingServiceClient) GetFeedMapping(ctx context.Context, in *GetFeedMappingRequest, opts ...grpc.CallOption) (*resources.FeedMapping, error) {
	out := new(resources.FeedMapping)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedMappingService/GetFeedMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedMappingServiceClient) MutateFeedMappings(ctx context.Context, in *MutateFeedMappingsRequest, opts ...grpc.CallOption) (*MutateFeedMappingsResponse, error) {
	out := new(MutateFeedMappingsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedMappingService/MutateFeedMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedMappingServiceServer is the server API for FeedMappingService service.
type FeedMappingServiceServer interface {
	// Returns the requested feed mapping in full detail.
	GetFeedMapping(context.Context, *GetFeedMappingRequest) (*resources.FeedMapping, error)
	// Creates or removes feed mappings. Operation statuses are
	// returned.
	MutateFeedMappings(context.Context, *MutateFeedMappingsRequest) (*MutateFeedMappingsResponse, error)
}

func RegisterFeedMappingServiceServer(s *grpc.Server, srv FeedMappingServiceServer) {
	s.RegisterService(&_FeedMappingService_serviceDesc, srv)
}

func _FeedMappingService_GetFeedMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).GetFeedMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedMappingService/GetFeedMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).GetFeedMapping(ctx, req.(*GetFeedMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedMappingService_MutateFeedMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedMappingService/MutateFeedMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedMappingServiceServer).MutateFeedMappings(ctx, req.(*MutateFeedMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedMappingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.FeedMappingService",
	HandlerType: (*FeedMappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedMapping",
			Handler:    _FeedMappingService_GetFeedMapping_Handler,
		},
		{
			MethodName: "MutateFeedMappings",
			Handler:    _FeedMappingService_MutateFeedMappings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/feed_mapping_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/feed_mapping_service.proto", fileDescriptor_feed_mapping_service_0f3de8fce9e30788)
}

var fileDescriptor_feed_mapping_service_0f3de8fce9e30788 = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0xfe, 0x6d, 0xf2, 0xa3, 0xda, 0x49, 0xad, 0x30, 0x5a, 0x1a, 0x83, 0x68, 0x58, 0x0b, 0x96,
	0x1c, 0x66, 0x63, 0x94, 0x8a, 0xdb, 0x56, 0x48, 0xc1, 0xb6, 0x1e, 0x6a, 0xcb, 0x16, 0x72, 0x90,
	0xc0, 0x32, 0xcd, 0xbe, 0x5d, 0x16, 0x76, 0x77, 0xd6, 0x99, 0xd9, 0x48, 0x29, 0xbd, 0x78, 0xf1,
	0x03, 0x78, 0xf1, 0xec, 0xd1, 0x9b, 0x47, 0xbf, 0x82, 0xe0, 0xc9, 0x6f, 0x20, 0x9e, 0xfc, 0x10,
	0x22, 0xbb, 0xb3, 0x13, 0x37, 0x6d, 0x42, 0xb4, 0xb7, 0x99, 0xf7, 0x7d, 0x9f, 0xe7, 0x7d, 0xde,
	0x3f, 0x33, 0x68, 0xdd, 0x67, 0xcc, 0x0f, 0xc1, 0xa2, 0x9e, 0xb0, 0xd4, 0x31, 0x3b, 0x0d, 0xdb,
	0x96, 0x00, 0x3e, 0x0c, 0x06, 0x20, 0xac, 0x63, 0x00, 0xcf, 0x8d, 0x68, 0x92, 0x04, 0xb1, 0xef,
	0x16, 0x56, 0x92, 0x70, 0x26, 0x19, 0x6e, 0x2a, 0x04, 0xa1, 0x9e, 0x20, 0x23, 0x30, 0x19, 0xb6,
	0x89, 0x06, 0x37, 0x1e, 0x4d, 0xa3, 0xe7, 0x20, 0x58, 0xca, 0xcf, 0xf3, 0x2b, 0xde, 0xc6, 0x6d,
	0x8d, 0x4a, 0x02, 0x8b, 0xc6, 0x31, 0x93, 0x54, 0x06, 0x2c, 0x16, 0x85, 0xf7, 0x4e, 0xe1, 0xcd,
	0x6f, 0x47, 0xe9, 0xb1, 0xf5, 0x9a, 0xd3, 0x24, 0x01, 0xae, 0xfd, 0xcb, 0x85, 0x9f, 0x27, 0x03,
	0x4b, 0x48, 0x2a, 0xd3, 0xc2, 0x61, 0x6e, 0xa0, 0xa5, 0x1d, 0x90, 0xdb, 0x00, 0xde, 0x9e, 0x4a,
	0xe7, 0xc0, 0xab, 0x14, 0x84, 0xc4, 0xf7, 0xd0, 0x35, 0xad, 0xc7, 0x8d, 0x69, 0x04, 0x75, 0xa3,
	0x69, 0xac, 0xce, 0x3b, 0x0b, 0xda, 0xf8, 0x82, 0x46, 0x60, 0x7e, 0x37, 0xd0, 0xad, 0xbd, 0x54,
	0x52, 0x09, 0x25, 0x06, 0xa1, 0x29, 0xee, 0xa2, 0xda, 0x20, 0x15, 0x92, 0x45, 0xc0, 0xdd, 0xc0,
	0x2b, 0x08, 0x90, 0x36, 0x3d, 0xf7, 0x70, 0x0f, 0x21, 0x96, 0x00, 0x57, 0x95, 0xd4, 0x2b, 0xcd,
	0xea, 0x6a, 0xad, 0xb3, 0x46, 0x66, 0x35, 0x90, 0x94, 0x72, 0xed, 0x6b, 0xb8, 0x53, 0x62, 0xc2,
	0xf7, 0xd1, 0xf5, 0x84, 0x72, 0x19, 0xd0, 0xd0, 0x3d, 0xa6, 0x41, 0x98, 0x72, 0xa8, 0x57, 0x9b,
	0xc6, 0xea, 0x55, 0x67, 0xb1, 0x30, 0x6f, 0x2b, 0x6b, 0x56, 0xe4, 0x90, 0x86, 0x81, 0x47, 0x25,
	0xb8, 0x2c, 0x0e, 0x4f, 0xea, 0xff, 0xe7, 0x61, 0x0b, 0xda, 0xb8, 0x1f, 0x87, 0x27, 0xe6, 0x5b,
	0x03, 0xdd, 0x9c, 0x94, 0x12, 0xef, 0xa2, 0xb9, 0x01, 0x07, 0x2a, 0x55, 0x6f, 0x6a, 0x1d, 0x32,
	0x55, 0xfa, 0x68, 0xb2, 0x65, 0xed, 0xbb, 0xff, 0x39, 0x05, 0x1e, 0xd7, 0xd1, 0x1c, 0x87, 0x88,
	0x0d, 0x95, 0xce, 0xf9, 0xcc, 0xa3, 0xee, 0x5b, 0x35, 0x34, 0x3f, 0x2a, 0xcc, 0xfc, 0x6c, 0xa0,
	0xc6, 0xa4, 0x76, 0x8b, 0x84, 0xc5, 0x02, 0xf0, 0x36, 0x5a, 0x3a, 0x57, 0xb6, 0x0b, 0x9c, 0x33,
	0x9e, 0x93, 0xd6, 0x3a, 0x58, 0xcb, 0xe3, 0xc9, 0x80, 0x1c, 0xe6, 0x4b, 0xe0, 0xdc, 0x18, 0x6f,
	0xc8, 0xb3, 0x2c, 0x1c, 0x1f, 0xa2, 0x2b, 0x1c, 0x44, 0x1a, 0x4a, 0x3d, 0x93, 0x27, 0xb3, 0x67,
	0x72, 0x41, 0x96, 0x93, 0x33, 0x38, 0x9a, 0xc9, 0x7c, 0x8a, 0x96, 0xa7, 0xc4, 0xfc, 0xd5, 0xaa,
	0x75, 0xde, 0x57, 0x11, 0x2e, 0x41, 0x0f, 0x55, 0x62, 0xfc, 0xc9, 0x40, 0x8b, 0xe3, 0x0b, 0x8c,
	0x1f, 0xcf, 0x56, 0x3b, 0x71, 0xe5, 0x1b, 0xff, 0x38, 0x3f, 0x73, 0xed, 0xcd, 0xb7, 0x1f, 0xef,
	0x2a, 0x6d, 0x4c, 0xb2, 0xc7, 0x7b, 0x3a, 0x56, 0xc2, 0xa6, 0xde, 0x72, 0x61, 0xb5, 0xf2, 0xd7,
	0xac, 0x87, 0x65, 0xb5, 0xce, 0xf0, 0x57, 0x03, 0xe1, 0x8b, 0x63, 0xc4, 0xeb, 0x97, 0xe8, 0xb2,
	0x7e, 0x6b, 0x8d, 0x8d, 0xcb, 0x81, 0xd5, 0xe6, 0x98, 0x1b, 0x79, 0x25, 0x6b, 0xe6, 0x83, 0xac,
	0x92, 0x3f, 0xd2, 0x4f, 0x4b, 0xcf, 0x77, 0xb3, 0x75, 0x36, 0x56, 0x88, 0x1d, 0xe5, 0x74, 0xb6,
	0xd1, 0xda, 0xfa, 0x65, 0xa0, 0x95, 0x01, 0x8b, 0x66, 0x2a, 0xd8, 0x5a, 0xbe, 0x38, 0xc0, 0x83,
	0xec, 0x17, 0x3a, 0x30, 0x5e, 0xee, 0x16, 0x60, 0x9f, 0x85, 0x34, 0xf6, 0x09, 0xe3, 0xbe, 0xe5,
	0x43, 0x9c, 0xff, 0x51, 0xfa, 0x93, 0x4c, 0x02, 0x31, 0xfd, 0x4b, 0x5e, 0xd7, 0x87, 0x0f, 0x95,
	0xea, 0x4e, 0xb7, 0xfb, 0xb1, 0xd2, 0xdc, 0x51, 0x84, 0x5d, 0x4f, 0x10, 0x75, 0xcc, 0x4e, 0xbd,
	0x36, 0x29, 0x12, 0x8b, 0x2f, 0x3a, 0xa4, 0xdf, 0xf5, 0x44, 0x7f, 0x14, 0xd2, 0xef, 0xb5, 0xfb,
	0x3a, 0xe4, 0x67, 0x65, 0x45, 0xd9, 0x6d, 0xbb, 0xeb, 0x09, 0xdb, 0x1e, 0x05, 0xd9, 0x76, 0xaf,
	0x6d, 0xdb, 0x3a, 0xec, 0x68, 0x2e, 0xd7, 0xf9, 0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0x40, 0x3b, 0xda, 0x39, 0x06, 0x00, 0x00,
}
