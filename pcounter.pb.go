// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: pcounter.proto

package proto_pcounter

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// People Counter Service message
type PCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string               `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Hostname string               `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Location string               `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Timezone string               `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Mac      string               `protobuf:"bytes,5,opt,name=mac,proto3" json:"mac,omitempty"`
	Hardware string               `protobuf:"bytes,6,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Protocol string               `protobuf:"bytes,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip       string               `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip,omitempty"`
	IpVpn    string               `protobuf:"bytes,9,opt,name=ip_vpn,json=ipVpn,proto3" json:"ip_vpn,omitempty"`
	Ts       *timestamp.Timestamp `protobuf:"bytes,10,opt,name=ts,proto3" json:"ts,omitempty"`
	Data     []*PEvent            `protobuf:"bytes,11,rep,name=data,proto3" json:"data,omitempty"` // message has several event data.
}

func (x *PCounter) Reset() {
	*x = PCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcounter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCounter) ProtoMessage() {}

func (x *PCounter) ProtoReflect() protoreflect.Message {
	mi := &file_pcounter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCounter.ProtoReflect.Descriptor instead.
func (*PCounter) Descriptor() ([]byte, []int) {
	return file_pcounter_proto_rawDescGZIP(), []int{0}
}

func (x *PCounter) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *PCounter) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *PCounter) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *PCounter) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *PCounter) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *PCounter) GetHardware() string {
	if x != nil {
		return x.Hardware
	}
	return ""
}

func (x *PCounter) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *PCounter) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PCounter) GetIpVpn() string {
	if x != nil {
		return x.IpVpn
	}
	return ""
}

func (x *PCounter) GetTs() *timestamp.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *PCounter) GetData() []*PEvent {
	if x != nil {
		return x.Data
	}
	return nil
}

type PEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Typ          string               `protobuf:"bytes,1,opt,name=typ,proto3" json:"typ,omitempty"` // one of "counter, fillLevel, dwellTime"
	Id           string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Seq          uint32               `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"` // sequence number (onr)
	Height       float32              `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
	Dir          string               `protobuf:"bytes,5,opt,name=dir,proto3" json:"dir,omitempty"` // direction
	Ts           *timestamp.Timestamp `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
	TsExit       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ts_exit,json=tsExit,proto3" json:"ts_exit,omitempty"`
	FillLevel    uint32               `protobuf:"varint,8,opt,name=fillLevel,proto3" json:"fillLevel,omitempty"`
	DwellTime    float32              `protobuf:"fixed32,9,opt,name=dwellTime,proto3" json:"dwellTime,omitempty"`
	ExpDwellTime float32              `protobuf:"fixed32,10,opt,name=expDwellTime,proto3" json:"expDwellTime,omitempty"`
	ObjectId     uint32               `protobuf:"varint,11,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	WalkSpeed    float32              `protobuf:"fixed32,12,opt,name=walk_speed,json=walkSpeed,proto3" json:"walk_speed,omitempty"`
}

func (x *PEvent) Reset() {
	*x = PEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcounter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PEvent) ProtoMessage() {}

func (x *PEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pcounter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PEvent.ProtoReflect.Descriptor instead.
func (*PEvent) Descriptor() ([]byte, []int) {
	return file_pcounter_proto_rawDescGZIP(), []int{1}
}

func (x *PEvent) GetTyp() string {
	if x != nil {
		return x.Typ
	}
	return ""
}

func (x *PEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PEvent) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PEvent) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *PEvent) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *PEvent) GetTs() *timestamp.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *PEvent) GetTsExit() *timestamp.Timestamp {
	if x != nil {
		return x.TsExit
	}
	return nil
}

func (x *PEvent) GetFillLevel() uint32 {
	if x != nil {
		return x.FillLevel
	}
	return 0
}

func (x *PEvent) GetDwellTime() float32 {
	if x != nil {
		return x.DwellTime
	}
	return 0
}

func (x *PEvent) GetExpDwellTime() float32 {
	if x != nil {
		return x.ExpDwellTime
	}
	return 0
}

func (x *PEvent) GetObjectId() uint32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *PEvent) GetWalkSpeed() float32 {
	if x != nil {
		return x.WalkSpeed
	}
	return 0
}

type PCounters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pcs []*PCounter `protobuf:"bytes,1,rep,name=pcs,proto3" json:"pcs,omitempty"` // multiple counter data for first transfer
}

func (x *PCounters) Reset() {
	*x = PCounters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcounter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCounters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCounters) ProtoMessage() {}

func (x *PCounters) ProtoReflect() protoreflect.Message {
	mi := &file_pcounter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCounters.ProtoReflect.Descriptor instead.
func (*PCounters) Descriptor() ([]byte, []int) {
	return file_pcounter_proto_rawDescGZIP(), []int{2}
}

func (x *PCounters) GetPcs() []*PCounter {
	if x != nil {
		return x.Pcs
	}
	return nil
}

type Retrieve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     *timestamp.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Duration *duration.Duration   `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Result   *PCounters           `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Retrieve) Reset() {
	*x = Retrieve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pcounter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retrieve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retrieve) ProtoMessage() {}

func (x *Retrieve) ProtoReflect() protoreflect.Message {
	mi := &file_pcounter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retrieve.ProtoReflect.Descriptor instead.
func (*Retrieve) Descriptor() ([]byte, []int) {
	return file_pcounter_proto_rawDescGZIP(), []int{3}
}

func (x *Retrieve) GetFrom() *timestamp.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Retrieve) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Retrieve) GetResult() *PCounters {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_pcounter_proto protoreflect.FileDescriptor

var file_pcounter_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x08,
	0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x70, 0x5f, 0x76, 0x70, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x70, 0x56, 0x70, 0x6e, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x50, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe3, 0x02, 0x0a,
	0x06, 0x50, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x79, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x73, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x74, 0x73, 0x45, 0x78, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x6c, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x6c, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x77, 0x65, 0x6c, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x64, 0x77, 0x65, 0x6c, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x44, 0x77, 0x65, 0x6c, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x44, 0x77, 0x65,
	0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x77, 0x61, 0x6c, 0x6b, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x22, 0x31, 0x0a, 0x09, 0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x24, 0x0a, 0x03, 0x70, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x03, 0x70, 0x63, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x65, 0x78, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x70, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pcounter_proto_rawDescOnce sync.Once
	file_pcounter_proto_rawDescData = file_pcounter_proto_rawDesc
)

func file_pcounter_proto_rawDescGZIP() []byte {
	file_pcounter_proto_rawDescOnce.Do(func() {
		file_pcounter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pcounter_proto_rawDescData)
	})
	return file_pcounter_proto_rawDescData
}

var file_pcounter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pcounter_proto_goTypes = []interface{}{
	(*PCounter)(nil),            // 0: pcounter.PCounter
	(*PEvent)(nil),              // 1: pcounter.PEvent
	(*PCounters)(nil),           // 2: pcounter.PCounters
	(*Retrieve)(nil),            // 3: pcounter.Retrieve
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 5: google.protobuf.Duration
}
var file_pcounter_proto_depIdxs = []int32{
	4, // 0: pcounter.PCounter.ts:type_name -> google.protobuf.Timestamp
	1, // 1: pcounter.PCounter.data:type_name -> pcounter.PEvent
	4, // 2: pcounter.PEvent.ts:type_name -> google.protobuf.Timestamp
	4, // 3: pcounter.PEvent.ts_exit:type_name -> google.protobuf.Timestamp
	0, // 4: pcounter.PCounters.pcs:type_name -> pcounter.PCounter
	4, // 5: pcounter.Retrieve.from:type_name -> google.protobuf.Timestamp
	5, // 6: pcounter.Retrieve.duration:type_name -> google.protobuf.Duration
	2, // 7: pcounter.Retrieve.result:type_name -> pcounter.PCounters
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pcounter_proto_init() }
func file_pcounter_proto_init() {
	if File_pcounter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pcounter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCounter); i {
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
		file_pcounter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PEvent); i {
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
		file_pcounter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCounters); i {
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
		file_pcounter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retrieve); i {
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
			RawDescriptor: file_pcounter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pcounter_proto_goTypes,
		DependencyIndexes: file_pcounter_proto_depIdxs,
		MessageInfos:      file_pcounter_proto_msgTypes,
	}.Build()
	File_pcounter_proto = out.File
	file_pcounter_proto_rawDesc = nil
	file_pcounter_proto_goTypes = nil
	file_pcounter_proto_depIdxs = nil
}
