// Code generated by protoc-gen-go.
// source: msgpb.proto
// DO NOT EDIT!

/*
Package msgpb is a generated protocol buffer package.

It is generated from these files:
	msgpb.proto

It has these top-level messages:
	Message
*/
package msgpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import raft_cmdpb "github.com/pingcap/kvproto/pkg/raft_cmdpb"
import raft_serverpb "github.com/pingcap/kvproto/pkg/raft_serverpb"
import kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"
import coprocessor "github.com/pingcap/kvproto/pkg/coprocessor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MessageType int32

const (
	MessageType_None    MessageType = 0
	MessageType_Cmd     MessageType = 1
	MessageType_CmdResp MessageType = 2
	MessageType_Raft    MessageType = 3
	MessageType_KvReq   MessageType = 4
	MessageType_KvResp  MessageType = 5
	MessageType_CopReq  MessageType = 6
	MessageType_CopResp MessageType = 7
)

var MessageType_name = map[int32]string{
	0: "None",
	1: "Cmd",
	2: "CmdResp",
	3: "Raft",
	4: "KvReq",
	5: "KvResp",
	6: "CopReq",
	7: "CopResp",
}
var MessageType_value = map[string]int32{
	"None":    0,
	"Cmd":     1,
	"CmdResp": 2,
	"Raft":    3,
	"KvReq":   4,
	"KvResp":  5,
	"CopReq":  6,
	"CopResp": 7,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Message holds all messages communicating with TiKV.
type Message struct {
	MsgType          *MessageType                `protobuf:"varint,1,opt,name=msg_type,enum=msgpb.MessageType" json:"msg_type,omitempty"`
	CmdReq           *raft_cmdpb.RaftCmdRequest  `protobuf:"bytes,2,opt,name=cmd_req" json:"cmd_req,omitempty"`
	CmdResp          *raft_cmdpb.RaftCmdResponse `protobuf:"bytes,3,opt,name=cmd_resp" json:"cmd_resp,omitempty"`
	Raft             *raft_serverpb.RaftMessage  `protobuf:"bytes,4,opt,name=raft" json:"raft,omitempty"`
	KvReq            *kvrpcpb.Request            `protobuf:"bytes,5,opt,name=kv_req" json:"kv_req,omitempty"`
	KvResp           *kvrpcpb.Response           `protobuf:"bytes,6,opt,name=kv_resp" json:"kv_resp,omitempty"`
	CopReq           *coprocessor.Request        `protobuf:"bytes,7,opt,name=cop_req" json:"cop_req,omitempty"`
	CopResp          *coprocessor.Response       `protobuf:"bytes,8,opt,name=cop_resp" json:"cop_resp,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetMsgType() MessageType {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return MessageType_None
}

func (m *Message) GetCmdReq() *raft_cmdpb.RaftCmdRequest {
	if m != nil {
		return m.CmdReq
	}
	return nil
}

func (m *Message) GetCmdResp() *raft_cmdpb.RaftCmdResponse {
	if m != nil {
		return m.CmdResp
	}
	return nil
}

func (m *Message) GetRaft() *raft_serverpb.RaftMessage {
	if m != nil {
		return m.Raft
	}
	return nil
}

func (m *Message) GetKvReq() *kvrpcpb.Request {
	if m != nil {
		return m.KvReq
	}
	return nil
}

func (m *Message) GetKvResp() *kvrpcpb.Response {
	if m != nil {
		return m.KvResp
	}
	return nil
}

func (m *Message) GetCopReq() *coprocessor.Request {
	if m != nil {
		return m.CopReq
	}
	return nil
}

func (m *Message) GetCopResp() *coprocessor.Response {
	if m != nil {
		return m.CopResp
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "msgpb.Message")
	proto.RegisterEnum("msgpb.MessageType", MessageType_name, MessageType_value)
}

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0x69, 0xfe, 0x1c, 0x36, 0x02, 0xb9, 0x06, 0xa4, 0x28, 0x5c, 0xaa, 0x0a, 0x44, 0x05,
	0x22, 0x87, 0xbe, 0x42, 0x8f, 0x08, 0x0e, 0x15, 0xf7, 0xa8, 0xa4, 0x26, 0x87, 0x2a, 0xb5, 0x6b,
	0x87, 0x48, 0x3c, 0x23, 0x2f, 0xc5, 0x7a, 0xf3, 0xa3, 0x08, 0x71, 0xdb, 0xec, 0x7c, 0x99, 0xd9,
	0x31, 0x24, 0xb5, 0xad, 0xf4, 0x47, 0xae, 0x8d, 0x6a, 0x94, 0x08, 0xe9, 0x23, 0xe3, 0x66, 0xf7,
	0xd9, 0x14, 0x65, 0xbd, 0x1f, 0x84, 0xec, 0x8a, 0x36, 0x56, 0x9a, 0x56, 0x9a, 0x71, 0x79, 0x71,
	0x68, 0x8d, 0x2e, 0xc7, 0xcf, 0x79, 0xa9, 0x70, 0x28, 0xa5, 0xb5, 0xca, 0x74, 0xab, 0xe5, 0x8f,
	0x07, 0xec, 0x15, 0x17, 0xbb, 0x4a, 0x8a, 0x3b, 0x88, 0xd1, 0xbd, 0x68, 0xbe, 0xb5, 0x4c, 0x67,
	0x8b, 0xd9, 0xea, 0x72, 0x2d, 0xf2, 0x2e, 0xbb, 0x27, 0xde, 0x51, 0x11, 0x4f, 0xc0, 0x30, 0xb7,
	0x30, 0xf2, 0x94, 0x7a, 0x08, 0x25, 0xeb, 0x2c, 0x9f, 0x1c, 0xb3, 0xc5, 0x71, 0x53, 0xef, 0xb7,
	0xf2, 0xf4, 0x25, 0x6d, 0x23, 0x9e, 0x21, 0xee, 0x60, 0xab, 0x53, 0x9f, 0xe8, 0xdb, 0x7f, 0x69,
	0xab, 0xd5, 0xd1, 0x4a, 0xb1, 0x82, 0xc0, 0xa9, 0x69, 0x30, 0x35, 0x1e, 0x3b, 0x39, 0x7a, 0xb8,
	0x75, 0x01, 0xd1, 0xa1, 0xa5, 0x23, 0x42, 0x62, 0x79, 0x3e, 0x54, 0x1d, 0xa2, 0x97, 0xc0, 0x88,
	0xc0, 0xe4, 0x88, 0x90, 0xf9, 0x04, 0xe9, 0xf3, 0xee, 0xb1, 0x8b, 0xd2, 0x64, 0xc3, 0x88, 0xb9,
	0xce, 0xa7, 0x4f, 0x34, 0x58, 0x3d, 0x60, 0x0b, 0xc2, 0xd0, 0x2b, 0x26, 0xee, 0xe6, 0x0f, 0xd7,
	0xf9, 0x3d, 0x56, 0x90, 0x4c, 0x9f, 0x2a, 0x86, 0xe0, 0x4d, 0x1d, 0x25, 0x3f, 0x13, 0x0c, 0x7c,
	0xec, 0xc9, 0x67, 0x22, 0x01, 0xd6, 0x17, 0xe6, 0x9e, 0xd3, 0x5d, 0x27, 0xee, 0x8b, 0x73, 0x08,
	0x5f, 0x5a, 0x8c, 0xe3, 0x81, 0x00, 0x88, 0xdc, 0x88, 0x40, 0xe8, 0xe6, 0x8d, 0xd2, 0x6e, 0x1f,
	0xd1, 0x9f, 0x6e, 0x46, 0x81, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xd5, 0x92, 0x3a, 0x14,
	0x02, 0x00, 0x00,
}
