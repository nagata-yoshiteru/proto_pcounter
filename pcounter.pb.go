// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pcounter.proto

package proto_pcounter // import "github.com/synerex/proto_pcounter"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// People Counter Service message
type PCounter struct {
	DeviceId             string               `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Hostname             string               `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Location             string               `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Timezone             string               `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Mac                  string               `protobuf:"bytes,5,opt,name=mac,proto3" json:"mac,omitempty"`
	Hardware             string               `protobuf:"bytes,6,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Protocol             string               `protobuf:"bytes,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip                   string               `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip,omitempty"`
	IpVpn                string               `protobuf:"bytes,9,opt,name=ip_vpn,json=ipVpn,proto3" json:"ip_vpn,omitempty"`
	Ts                   *timestamp.Timestamp `protobuf:"bytes,10,opt,name=ts,proto3" json:"ts,omitempty"`
	Data                 []*PEvent            `protobuf:"bytes,11,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PCounter) Reset()         { *m = PCounter{} }
func (m *PCounter) String() string { return proto.CompactTextString(m) }
func (*PCounter) ProtoMessage()    {}
func (*PCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_pcounter_2415586b03415292, []int{0}
}
func (m *PCounter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PCounter.Unmarshal(m, b)
}
func (m *PCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PCounter.Marshal(b, m, deterministic)
}
func (dst *PCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PCounter.Merge(dst, src)
}
func (m *PCounter) XXX_Size() int {
	return xxx_messageInfo_PCounter.Size(m)
}
func (m *PCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_PCounter.DiscardUnknown(m)
}

var xxx_messageInfo_PCounter proto.InternalMessageInfo

func (m *PCounter) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *PCounter) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *PCounter) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *PCounter) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *PCounter) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *PCounter) GetHardware() string {
	if m != nil {
		return m.Hardware
	}
	return ""
}

func (m *PCounter) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *PCounter) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *PCounter) GetIpVpn() string {
	if m != nil {
		return m.IpVpn
	}
	return ""
}

func (m *PCounter) GetTs() *timestamp.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *PCounter) GetData() []*PEvent {
	if m != nil {
		return m.Data
	}
	return nil
}

type PEvent struct {
	Typ                  string               `protobuf:"bytes,1,opt,name=typ,proto3" json:"typ,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Seq                  uint32               `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Height               uint32               `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Dir                  string               `protobuf:"bytes,5,opt,name=dir,proto3" json:"dir,omitempty"`
	Ts                   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
	TsExit               *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ts_exit,json=tsExit,proto3" json:"ts_exit,omitempty"`
	FillLevel            uint32               `protobuf:"varint,8,opt,name=fillLevel,proto3" json:"fillLevel,omitempty"`
	DwellTime            float32              `protobuf:"fixed32,9,opt,name=dwellTime,proto3" json:"dwellTime,omitempty"`
	ExpDwellTime         float32              `protobuf:"fixed32,10,opt,name=expDwellTime,proto3" json:"expDwellTime,omitempty"`
	ObjectId             uint32               `protobuf:"varint,11,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PEvent) Reset()         { *m = PEvent{} }
func (m *PEvent) String() string { return proto.CompactTextString(m) }
func (*PEvent) ProtoMessage()    {}
func (*PEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_pcounter_2415586b03415292, []int{1}
}
func (m *PEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PEvent.Unmarshal(m, b)
}
func (m *PEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PEvent.Marshal(b, m, deterministic)
}
func (dst *PEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PEvent.Merge(dst, src)
}
func (m *PEvent) XXX_Size() int {
	return xxx_messageInfo_PEvent.Size(m)
}
func (m *PEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PEvent proto.InternalMessageInfo

func (m *PEvent) GetTyp() string {
	if m != nil {
		return m.Typ
	}
	return ""
}

func (m *PEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PEvent) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PEvent) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PEvent) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *PEvent) GetTs() *timestamp.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *PEvent) GetTsExit() *timestamp.Timestamp {
	if m != nil {
		return m.TsExit
	}
	return nil
}

func (m *PEvent) GetFillLevel() uint32 {
	if m != nil {
		return m.FillLevel
	}
	return 0
}

func (m *PEvent) GetDwellTime() float32 {
	if m != nil {
		return m.DwellTime
	}
	return 0
}

func (m *PEvent) GetExpDwellTime() float32 {
	if m != nil {
		return m.ExpDwellTime
	}
	return 0
}

func (m *PEvent) GetObjectId() uint32 {
	if m != nil {
		return m.ObjectId
	}
	return 0
}

func init() {
	proto.RegisterType((*PCounter)(nil), "pcounter.PCounter")
	proto.RegisterType((*PEvent)(nil), "pcounter.PEvent")
}

func init() { proto.RegisterFile("pcounter.proto", fileDescriptor_pcounter_2415586b03415292) }

var fileDescriptor_pcounter_2415586b03415292 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x55, 0x77, 0xcb, 0x52, 0x97, 0x4e, 0x93, 0x25, 0x90, 0x55, 0x90, 0x28, 0x85, 0x43,
	0xc5, 0x21, 0x95, 0xb6, 0xff, 0x00, 0xd8, 0x61, 0x12, 0x87, 0x29, 0x42, 0x1c, 0xb8, 0x54, 0x69,
	0xfc, 0xd6, 0x3e, 0x94, 0xc4, 0x26, 0x79, 0xed, 0x32, 0xfe, 0x38, 0x4e, 0xfc, 0x61, 0xc8, 0xcf,
	0xf1, 0x26, 0x2e, 0x88, 0x9b, 0xbf, 0x3f, 0x1c, 0xbf, 0x7c, 0xf4, 0xe4, 0xb9, 0x2b, 0xed, 0xa1,
	0x21, 0x68, 0x33, 0xd7, 0x5a, 0xb2, 0x2a, 0x8d, 0x7a, 0xfe, 0x7a, 0x67, 0xed, 0xae, 0x82, 0x35,
	0xfb, 0xdb, 0xc3, 0xdd, 0x9a, 0xb0, 0x86, 0x8e, 0x8a, 0xda, 0x85, 0xea, 0xf2, 0x97, 0x90, 0xe9,
	0xed, 0xc7, 0xd0, 0x56, 0x2f, 0xe5, 0xc4, 0xc0, 0x11, 0x4b, 0xd8, 0xa0, 0xd1, 0xa3, 0xc5, 0x68,
	0x35, 0xc9, 0xd3, 0x60, 0xdc, 0x18, 0x35, 0x97, 0xe9, 0xde, 0x76, 0xd4, 0x14, 0x35, 0x68, 0x11,
	0xb2, 0xa8, 0x7d, 0x56, 0xd9, 0xb2, 0x20, 0xb4, 0x8d, 0x1e, 0x87, 0x2c, 0x6a, 0x9f, 0xf9, 0x47,
	0x7f, 0xda, 0x06, 0xf4, 0x49, 0xc8, 0xa2, 0x56, 0x17, 0x72, 0x5c, 0x17, 0xa5, 0x3e, 0x65, 0xdb,
	0x1f, 0xf9, 0x95, 0xa2, 0x35, 0xf7, 0x45, 0x0b, 0x3a, 0x19, 0x5e, 0x19, 0xb4, 0xcf, 0x78, 0xe8,
	0xd2, 0x56, 0xfa, 0x2c, 0x64, 0x51, 0xab, 0x73, 0x29, 0xd0, 0xe9, 0x94, 0x5d, 0x81, 0x4e, 0x3d,
	0x97, 0x09, 0xba, 0xcd, 0xd1, 0x35, 0x7a, 0xc2, 0xde, 0x29, 0xba, 0xaf, 0xae, 0x51, 0xef, 0xa5,
	0xa0, 0x4e, 0xcb, 0xc5, 0x68, 0x35, 0xbd, 0x9c, 0x67, 0x01, 0x4e, 0x16, 0xe1, 0x64, 0x5f, 0x22,
	0x9c, 0x5c, 0x50, 0xa7, 0xde, 0xc9, 0x13, 0x53, 0x50, 0xa1, 0xa7, 0x8b, 0xf1, 0x6a, 0x7a, 0x79,
	0x91, 0x3d, 0x42, 0xbe, 0xbd, 0x3e, 0x42, 0x43, 0x39, 0xa7, 0xcb, 0xdf, 0x42, 0x26, 0xc1, 0xf0,
	0x7f, 0x43, 0x0f, 0x6e, 0x00, 0xe7, 0x8f, 0x3c, 0x95, 0x19, 0x68, 0x09, 0x34, 0xbe, 0xd1, 0xc1,
	0x0f, 0x46, 0x34, 0xcb, 0xfd, 0x51, 0xbd, 0x90, 0xc9, 0x1e, 0x70, 0xb7, 0x27, 0x66, 0x33, 0xcb,
	0x07, 0xe5, 0x9b, 0x06, 0xdb, 0x48, 0xc6, 0x60, 0x3b, 0x8c, 0x9e, 0xfc, 0xd7, 0xe8, 0x57, 0xf2,
	0x8c, 0xba, 0x0d, 0xf4, 0x48, 0x0c, 0xea, 0xdf, 0x17, 0x12, 0xea, 0xae, 0x7b, 0x24, 0xf5, 0x4a,
	0x4e, 0xee, 0xb0, 0xaa, 0x3e, 0xc3, 0x11, 0x2a, 0x26, 0x39, 0xcb, 0x9f, 0x0c, 0x9f, 0x9a, 0x7b,
	0xa8, 0x2a, 0x7f, 0x8f, 0x99, 0x8a, 0xfc, 0xc9, 0x50, 0x4b, 0xf9, 0x0c, 0x7a, 0xf7, 0xe9, 0xb1,
	0x20, 0xb9, 0xf0, 0x97, 0xe7, 0xb7, 0xcb, 0x6e, 0xbf, 0x43, 0x49, 0x7e, 0xbb, 0xa6, 0xfc, 0xfd,
	0x34, 0x18, 0x37, 0xe6, 0xc3, 0xdb, 0x6f, 0x6f, 0x76, 0x48, 0xfb, 0xc3, 0x36, 0x2b, 0x6d, 0xbd,
	0xee, 0x1e, 0x1a, 0x68, 0xa1, 0x0f, 0x6b, 0xbb, 0x89, 0xe0, 0xb7, 0x09, 0xeb, 0xab, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x43, 0x7e, 0x88, 0x75, 0xf0, 0x02, 0x00, 0x00,
}
