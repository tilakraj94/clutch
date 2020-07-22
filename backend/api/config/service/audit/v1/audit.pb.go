// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/service/audit/v1/audit.proto

package auditv1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// Switch to control what field the message
type EventFilter_FilterType int32

const (
	// Amounts to a no-op filter.
	EventFilter_UNSPECIFIED EventFilter_FilterType = 0
	// Compare to the service performing the operation.
	EventFilter_SERVICE EventFilter_FilterType = 1
	// Compare to the method being called.
	EventFilter_METHOD EventFilter_FilterType = 2
	// Compare against the action type of the event.
	EventFilter_TYPE EventFilter_FilterType = 3
)

var EventFilter_FilterType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "SERVICE",
	2: "METHOD",
	3: "TYPE",
}

var EventFilter_FilterType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"SERVICE":     1,
	"METHOD":      2,
	"TYPE":        3,
}

func (x EventFilter_FilterType) String() string {
	return proto.EnumName(EventFilter_FilterType_name, int32(x))
}

func (EventFilter_FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a5e61bfed6edc4f, []int{0, 0}
}

type EventFilter struct {
	Field EventFilter_FilterType `protobuf:"varint,1,opt,name=field,proto3,enum=clutch.config.service.audit.v1.EventFilter_FilterType" json:"field,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*EventFilter_Text
	Value                isEventFilter_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a5e61bfed6edc4f, []int{0}
}

func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (m *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(m, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetField() EventFilter_FilterType {
	if m != nil {
		return m.Field
	}
	return EventFilter_UNSPECIFIED
}

type isEventFilter_Value interface {
	isEventFilter_Value()
}

type EventFilter_Text struct {
	Text string `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

func (*EventFilter_Text) isEventFilter_Value() {}

func (m *EventFilter) GetValue() isEventFilter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *EventFilter) GetText() string {
	if x, ok := m.GetValue().(*EventFilter_Text); ok {
		return x.Text
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventFilter_Text)(nil),
	}
}

type Filter struct {
	// Whether to treat the list as a allowlist (default) or a denylist.
	Denylist bool `protobuf:"varint,1,opt,name=denylist,proto3" json:"denylist,omitempty"`
	// The filter rules to apply against messages.
	Rules                []*EventFilter `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a5e61bfed6edc4f, []int{1}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetDenylist() bool {
	if m != nil {
		return m.Denylist
	}
	return false
}

func (m *Filter) GetRules() []*EventFilter {
	if m != nil {
		return m.Rules
	}
	return nil
}

type SinkConfig struct {
	// The rule(s) to filter events between audit source and sink emission.
	Filter               *Filter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SinkConfig) Reset()         { *m = SinkConfig{} }
func (m *SinkConfig) String() string { return proto.CompactTextString(m) }
func (*SinkConfig) ProtoMessage()    {}
func (*SinkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a5e61bfed6edc4f, []int{2}
}

func (m *SinkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SinkConfig.Unmarshal(m, b)
}
func (m *SinkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SinkConfig.Marshal(b, m, deterministic)
}
func (m *SinkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SinkConfig.Merge(m, src)
}
func (m *SinkConfig) XXX_Size() int {
	return xxx_messageInfo_SinkConfig.Size(m)
}
func (m *SinkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SinkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SinkConfig proto.InternalMessageInfo

func (m *SinkConfig) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type Config struct {
	// The name of the service where the auditor will persist events.
	DbProvider string `protobuf:"bytes,1,opt,name=db_provider,json=dbProvider,proto3" json:"db_provider,omitempty"`
	// The rule to apply before between request ingress and the database.
	Filter *Filter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// The registered name of sinks to fan-out events to.
	Sinks                []string `protobuf:"bytes,3,rep,name=sinks,proto3" json:"sinks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a5e61bfed6edc4f, []int{3}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetDbProvider() string {
	if m != nil {
		return m.DbProvider
	}
	return ""
}

func (m *Config) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Config) GetSinks() []string {
	if m != nil {
		return m.Sinks
	}
	return nil
}

func init() {
	proto.RegisterEnum("clutch.config.service.audit.v1.EventFilter_FilterType", EventFilter_FilterType_name, EventFilter_FilterType_value)
	proto.RegisterType((*EventFilter)(nil), "clutch.config.service.audit.v1.EventFilter")
	proto.RegisterType((*Filter)(nil), "clutch.config.service.audit.v1.Filter")
	proto.RegisterType((*SinkConfig)(nil), "clutch.config.service.audit.v1.SinkConfig")
	proto.RegisterType((*Config)(nil), "clutch.config.service.audit.v1.Config")
}

func init() {
	proto.RegisterFile("config/service/audit/v1/audit.proto", fileDescriptor_9a5e61bfed6edc4f)
}

var fileDescriptor_9a5e61bfed6edc4f = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6b, 0xa3, 0x40,
	0x14, 0xc7, 0xa3, 0x46, 0x8d, 0x4f, 0xd8, 0x95, 0x21, 0xb0, 0x92, 0xc3, 0x22, 0x2e, 0x2c, 0x42,
	0x41, 0x49, 0x0a, 0x3d, 0x96, 0xd6, 0xc4, 0x90, 0x40, 0xda, 0x06, 0x93, 0x16, 0xda, 0x4b, 0x31,
	0x3a, 0x49, 0x87, 0x88, 0x06, 0x1d, 0x87, 0xe6, 0x1b, 0xf4, 0x43, 0xf5, 0x93, 0xf5, 0x54, 0x70,
	0x6c, 0x9b, 0x53, 0x4b, 0x7b, 0xf2, 0x3d, 0xde, 0xff, 0xf7, 0x7b, 0x4f, 0x18, 0xf8, 0x17, 0xe7,
	0xd9, 0x9a, 0x6c, 0xbc, 0x12, 0x17, 0x8c, 0xc4, 0xd8, 0x8b, 0xaa, 0x84, 0x50, 0x8f, 0xf5, 0x79,
	0xe1, 0xee, 0x8a, 0x9c, 0xe6, 0xe8, 0x6f, 0x9c, 0x56, 0x34, 0x7e, 0x70, 0x79, 0xd6, 0x6d, 0xb2,
	0x2e, 0x8f, 0xb0, 0x7e, 0xef, 0x0f, 0x8b, 0x52, 0x92, 0x44, 0x14, 0x7b, 0x6f, 0x05, 0x07, 0xed,
	0x67, 0x01, 0xf4, 0x80, 0xe1, 0x8c, 0x8e, 0x49, 0x4a, 0x71, 0x81, 0x66, 0x20, 0xaf, 0x09, 0x4e,
	0x13, 0x53, 0xb0, 0x04, 0xe7, 0xd7, 0xe0, 0xc4, 0xfd, 0x5c, 0xec, 0x1e, 0xb0, 0x2e, 0xff, 0x2c,
	0xf7, 0x3b, 0x1c, 0x72, 0x09, 0xea, 0x42, 0x9b, 0xe2, 0x47, 0x6a, 0x8a, 0x96, 0xe0, 0x68, 0x93,
	0x56, 0x58, 0x77, 0xf6, 0x19, 0xc0, 0x47, 0x14, 0xfd, 0x06, 0xfd, 0xfa, 0x72, 0x31, 0x0f, 0x86,
	0xd3, 0xf1, 0x34, 0x18, 0x19, 0x2d, 0xa4, 0x83, 0xba, 0x08, 0xc2, 0x9b, 0xe9, 0x30, 0x30, 0x04,
	0x04, 0xa0, 0x5c, 0x04, 0xcb, 0xc9, 0xd5, 0xc8, 0x10, 0x51, 0x07, 0xda, 0xcb, 0xdb, 0x79, 0x60,
	0x48, 0xbe, 0x0a, 0x32, 0x8b, 0xd2, 0x0a, 0xdb, 0x1b, 0x50, 0x9a, 0xc3, 0x7b, 0xd0, 0x49, 0x70,
	0xb6, 0x4f, 0x49, 0x49, 0xeb, 0xdb, 0x3b, 0xe1, 0x7b, 0x8f, 0xce, 0x41, 0x2e, 0xaa, 0x14, 0x97,
	0xa6, 0x68, 0x49, 0x8e, 0x3e, 0x38, 0xfa, 0xc6, 0x4f, 0x85, 0x9c, 0xb4, 0x67, 0x00, 0x0b, 0x92,
	0x6d, 0x87, 0x35, 0x81, 0x4e, 0x41, 0x59, 0xd7, 0xe3, 0x7a, 0x95, 0x3e, 0xf8, 0xff, 0x95, 0xb1,
	0x91, 0x35, 0x94, 0xfd, 0x24, 0x80, 0xd2, 0xa8, 0x1c, 0xd0, 0x93, 0xd5, 0xfd, 0xae, 0xc8, 0x19,
	0x49, 0x1a, 0x9f, 0xe6, 0xab, 0x2f, 0x7e, 0xbb, 0x10, 0x2d, 0x21, 0x84, 0x64, 0x35, 0x6f, 0x46,
	0x07, 0x4b, 0xc5, 0x9f, 0x2c, 0x45, 0x5d, 0x90, 0x4b, 0x92, 0x6d, 0x4b, 0x53, 0xb2, 0x24, 0x47,
	0x0b, 0x79, 0xe3, 0x6b, 0x77, 0x6a, 0x0d, 0xb0, 0xfe, 0x4a, 0xa9, 0x9f, 0xc4, 0xf1, 0x6b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x79, 0x42, 0x07, 0xb5, 0x72, 0x02, 0x00, 0x00,
}