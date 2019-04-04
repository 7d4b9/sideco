// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package scoring

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetScoreRequest struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Tag                  []string `protobuf:"bytes,2,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetScoreRequest) Reset()         { *m = GetScoreRequest{} }
func (m *GetScoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetScoreRequest) ProtoMessage()    {}
func (*GetScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *GetScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetScoreRequest.Unmarshal(m, b)
}
func (m *GetScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetScoreRequest.Marshal(b, m, deterministic)
}
func (m *GetScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetScoreRequest.Merge(m, src)
}
func (m *GetScoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetScoreRequest.Size(m)
}
func (m *GetScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetScoreRequest proto.InternalMessageInfo

func (m *GetScoreRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *GetScoreRequest) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

type GetScoreResponse struct {
	Score                int32    `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	MatchingTag          []string `protobuf:"bytes,2,rep,name=matchingTag,proto3" json:"matchingTag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetScoreResponse) Reset()         { *m = GetScoreResponse{} }
func (m *GetScoreResponse) String() string { return proto.CompactTextString(m) }
func (*GetScoreResponse) ProtoMessage()    {}
func (*GetScoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *GetScoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetScoreResponse.Unmarshal(m, b)
}
func (m *GetScoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetScoreResponse.Marshal(b, m, deterministic)
}
func (m *GetScoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetScoreResponse.Merge(m, src)
}
func (m *GetScoreResponse) XXX_Size() int {
	return xxx_messageInfo_GetScoreResponse.Size(m)
}
func (m *GetScoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetScoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetScoreResponse proto.InternalMessageInfo

func (m *GetScoreResponse) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *GetScoreResponse) GetMatchingTag() []string {
	if m != nil {
		return m.MatchingTag
	}
	return nil
}

func init() {
	proto.RegisterType((*GetScoreRequest)(nil), "scoring.GetScoreRequest")
	proto.RegisterType((*GetScoreResponse)(nil), "scoring.GetScoreResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x2f, 0xca, 0xcc, 0x4b, 0x57, 0xb2,
	0xe6, 0xe2, 0x77, 0x4f, 0x2d, 0x09, 0x4e, 0xce, 0x2f, 0x4a, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x49, 0x2c, 0xce, 0xf6, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x12, 0xd3, 0x25, 0x98, 0x14, 0x98, 0x35,
	0x38, 0x83, 0x40, 0x4c, 0x25, 0x2f, 0x2e, 0x01, 0x84, 0xe6, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x11, 0x2e, 0x56, 0x90, 0xd9, 0xa9, 0x60, 0xcd, 0xac, 0x41, 0x10, 0x8e, 0x90, 0x02, 0x17,
	0x77, 0x6e, 0x62, 0x49, 0x72, 0x46, 0x66, 0x5e, 0x7a, 0x08, 0xdc, 0x0c, 0x64, 0x21, 0x23, 0x1f,
	0x2e, 0xf6, 0x60, 0x88, 0x9b, 0x84, 0x1c, 0xb9, 0x38, 0x60, 0xc6, 0x0a, 0x49, 0xe8, 0x41, 0x5d,
	0xaa, 0x87, 0xe6, 0x4c, 0x29, 0x49, 0x2c, 0x32, 0x10, 0x37, 0x28, 0x31, 0x24, 0xb1, 0x81, 0xbd,
	0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x26, 0xb5, 0xf5, 0x1d, 0xf3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScoringClient is the client API for Scoring service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScoringClient interface {
	// A simple RPC.
	//
	// GetScore returns the list of tags for a particular taskthe matching between a given set of applicants and a given mission.
	GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error)
}

type scoringClient struct {
	cc *grpc.ClientConn
}

func NewScoringClient(cc *grpc.ClientConn) ScoringClient {
	return &scoringClient{cc}
}

func (c *scoringClient) GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error) {
	out := new(GetScoreResponse)
	err := c.cc.Invoke(ctx, "/scoring.Scoring/GetScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoringServer is the server API for Scoring service.
type ScoringServer interface {
	// A simple RPC.
	//
	// GetScore returns the list of tags for a particular taskthe matching between a given set of applicants and a given mission.
	GetScore(context.Context, *GetScoreRequest) (*GetScoreResponse, error)
}

func RegisterScoringServer(s *grpc.Server, srv ScoringServer) {
	s.RegisterService(&_Scoring_serviceDesc, srv)
}

func _Scoring_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring.Scoring/GetScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServer).GetScore(ctx, req.(*GetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scoring_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scoring.Scoring",
	HandlerType: (*ScoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScore",
			Handler:    _Scoring_GetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}