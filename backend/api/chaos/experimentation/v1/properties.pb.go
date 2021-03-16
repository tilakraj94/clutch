// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: chaos/experimentation/v1/properties.proto

package experimentationv1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type PropertiesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ordered list of items.
	Items []*Property `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PropertiesList) Reset() {
	*x = PropertiesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_experimentation_v1_properties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertiesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertiesList) ProtoMessage() {}

func (x *PropertiesList) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_experimentation_v1_properties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertiesList.ProtoReflect.Descriptor instead.
func (*PropertiesList) Descriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_properties_proto_rawDescGZIP(), []int{0}
}

func (x *PropertiesList) GetItems() []*Property {
	if x != nil {
		return x.Items
	}
	return nil
}

type PropertiesMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unordered list of items that provides a fast random access to a property with a given identifier.
	Items map[string]*Property `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PropertiesMap) Reset() {
	*x = PropertiesMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_experimentation_v1_properties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertiesMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertiesMap) ProtoMessage() {}

func (x *PropertiesMap) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_experimentation_v1_properties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertiesMap.ProtoReflect.Descriptor instead.
func (*PropertiesMap) Descriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_properties_proto_rawDescGZIP(), []int{1}
}

func (x *PropertiesMap) GetItems() map[string]*Property {
	if x != nil {
		return x.Items
	}
	return nil
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the property.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The human readable name of the property.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// The human readable value of the property. If it's not provided it's up to the caller to determine how
	// to display the content of the `value` field of the property.
	DisplayValue *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=display_value,json=displayValue,proto3" json:"display_value,omitempty"`
	// The value of the property.
	//
	// Types that are assignable to Value:
	//	*Property_DateValue
	//	*Property_StringValue
	//	*Property_IntValue
	//	*Property_UrlValue
	Value isProperty_Value `protobuf_oneof:"value"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_experimentation_v1_properties_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_experimentation_v1_properties_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_properties_proto_rawDescGZIP(), []int{2}
}

func (x *Property) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Property) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Property) GetDisplayValue() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayValue
	}
	return nil
}

func (m *Property) GetValue() isProperty_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Property) GetDateValue() *timestamppb.Timestamp {
	if x, ok := x.GetValue().(*Property_DateValue); ok {
		return x.DateValue
	}
	return nil
}

func (x *Property) GetStringValue() string {
	if x, ok := x.GetValue().(*Property_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Property) GetIntValue() int64 {
	if x, ok := x.GetValue().(*Property_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *Property) GetUrlValue() string {
	if x, ok := x.GetValue().(*Property_UrlValue); ok {
		return x.UrlValue
	}
	return ""
}

type isProperty_Value interface {
	isProperty_Value()
}

type Property_DateValue struct {
	DateValue *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date_value,json=dateValue,proto3,oneof"`
}

type Property_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Property_IntValue struct {
	IntValue int64 `protobuf:"varint,6,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Property_UrlValue struct {
	UrlValue string `protobuf:"bytes,7,opt,name=url_value,json=urlValue,proto3,oneof"`
}

func (*Property_DateValue) isProperty_Value() {}

func (*Property_StringValue) isProperty_Value() {}

func (*Property_IntValue) isProperty_Value() {}

func (*Property_UrlValue) isProperty_Value() {}

var File_chaos_experimentation_v1_properties_proto protoreflect.FileDescriptor

var file_chaos_experimentation_v1_properties_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0xc5, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x4d,
	0x61, 0x70, 0x12, 0x4f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f, 0x73,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x4d, 0x61,
	0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x1a, 0x63, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x6f,
	0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9c, 0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x41, 0x0a, 0x0d, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b,
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x75, 0x72, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x72, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68,
	0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaos_experimentation_v1_properties_proto_rawDescOnce sync.Once
	file_chaos_experimentation_v1_properties_proto_rawDescData = file_chaos_experimentation_v1_properties_proto_rawDesc
)

func file_chaos_experimentation_v1_properties_proto_rawDescGZIP() []byte {
	file_chaos_experimentation_v1_properties_proto_rawDescOnce.Do(func() {
		file_chaos_experimentation_v1_properties_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaos_experimentation_v1_properties_proto_rawDescData)
	})
	return file_chaos_experimentation_v1_properties_proto_rawDescData
}

var file_chaos_experimentation_v1_properties_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chaos_experimentation_v1_properties_proto_goTypes = []interface{}{
	(*PropertiesList)(nil),         // 0: clutch.chaos.experimentation.v1.PropertiesList
	(*PropertiesMap)(nil),          // 1: clutch.chaos.experimentation.v1.PropertiesMap
	(*Property)(nil),               // 2: clutch.chaos.experimentation.v1.Property
	nil,                            // 3: clutch.chaos.experimentation.v1.PropertiesMap.ItemsEntry
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_chaos_experimentation_v1_properties_proto_depIdxs = []int32{
	2, // 0: clutch.chaos.experimentation.v1.PropertiesList.items:type_name -> clutch.chaos.experimentation.v1.Property
	3, // 1: clutch.chaos.experimentation.v1.PropertiesMap.items:type_name -> clutch.chaos.experimentation.v1.PropertiesMap.ItemsEntry
	4, // 2: clutch.chaos.experimentation.v1.Property.display_value:type_name -> google.protobuf.StringValue
	5, // 3: clutch.chaos.experimentation.v1.Property.date_value:type_name -> google.protobuf.Timestamp
	2, // 4: clutch.chaos.experimentation.v1.PropertiesMap.ItemsEntry.value:type_name -> clutch.chaos.experimentation.v1.Property
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_chaos_experimentation_v1_properties_proto_init() }
func file_chaos_experimentation_v1_properties_proto_init() {
	if File_chaos_experimentation_v1_properties_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaos_experimentation_v1_properties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertiesList); i {
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
		file_chaos_experimentation_v1_properties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertiesMap); i {
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
		file_chaos_experimentation_v1_properties_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
	file_chaos_experimentation_v1_properties_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Property_DateValue)(nil),
		(*Property_StringValue)(nil),
		(*Property_IntValue)(nil),
		(*Property_UrlValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chaos_experimentation_v1_properties_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaos_experimentation_v1_properties_proto_goTypes,
		DependencyIndexes: file_chaos_experimentation_v1_properties_proto_depIdxs,
		MessageInfos:      file_chaos_experimentation_v1_properties_proto_msgTypes,
	}.Build()
	File_chaos_experimentation_v1_properties_proto = out.File
	file_chaos_experimentation_v1_properties_proto_rawDesc = nil
	file_chaos_experimentation_v1_properties_proto_goTypes = nil
	file_chaos_experimentation_v1_properties_proto_depIdxs = nil
}
