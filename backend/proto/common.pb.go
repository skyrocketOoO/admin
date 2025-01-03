// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: common.proto

package proto

import (
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

type Concator int32

const (
	Concator_AND Concator = 0
	Concator_OR  Concator = 1
)

// Enum value maps for Concator.
var (
	Concator_name = map[int32]string{
		0: "AND",
		1: "OR",
	}
	Concator_value = map[string]int32{
		"AND": 0,
		"OR":  1,
	}
)

func (x Concator) Enum() *Concator {
	p := new(Concator)
	*p = x
	return p
}

func (x Concator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Concator) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Concator) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Concator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Concator.Descriptor instead.
func (Concator) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type ListOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pager          *Pager          `protobuf:"bytes,1,opt,name=pager,proto3,oneof" json:"pager,omitempty"`
	Sorters        []*Sorter       `protobuf:"bytes,2,rep,name=sorters,proto3" json:"sorters,omitempty"`
	ConditionGroup *ConditionGroup `protobuf:"bytes,3,opt,name=conditionGroup,proto3,oneof" json:"conditionGroup,omitempty"` // Root of filter expressions.
}

func (x *ListOption) Reset() {
	*x = ListOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOption) ProtoMessage() {}

func (x *ListOption) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOption.ProtoReflect.Descriptor instead.
func (*ListOption) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *ListOption) GetPager() *Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListOption) GetSorters() []*Sorter {
	if x != nil {
		return x.Sorters
	}
	return nil
}

func (x *ListOption) GetConditionGroup() *ConditionGroup {
	if x != nil {
		return x.ConditionGroup
	}
	return nil
}

type Pager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Size   int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Pager) Reset() {
	*x = Pager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pager) ProtoMessage() {}

func (x *Pager) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pager.ProtoReflect.Descriptor instead.
func (*Pager) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Pager) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Pager) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Sorter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ascending bool   `protobuf:"varint,1,opt,name=ascending,proto3" json:"ascending,omitempty"`
	Field     string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *Sorter) Reset() {
	*x = Sorter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sorter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sorter) ProtoMessage() {}

func (x *Sorter) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sorter.ProtoReflect.Descriptor instead.
func (*Sorter) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Sorter) GetAscending() bool {
	if x != nil {
		return x.Ascending
	}
	return false
}

func (x *Sorter) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

// NOT SUPPORTED: "BETWEEN" operator.
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Operator string `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"` // same as sql operator
	Value    string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *Condition) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Condition) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *Condition) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Group of filters with a logical operator.
type ConditionGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concator        Concator          `protobuf:"varint,1,opt,name=concator,proto3,enum=proto.Concator" json:"concator,omitempty"` // Logical operator (AND/OR) to combine filters.
	Conditions      []*Condition      `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`                  // List of individual filters.
	ConditionGroups []*ConditionGroup `protobuf:"bytes,3,rep,name=conditionGroups,proto3" json:"conditionGroups,omitempty"`        // Nested filter groups for recursive logic.
}

func (x *ConditionGroup) Reset() {
	*x = ConditionGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionGroup) ProtoMessage() {}

func (x *ConditionGroup) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionGroup.ProtoReflect.Descriptor instead.
func (*ConditionGroup) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *ConditionGroup) GetConcator() Concator {
	if x != nil {
		return x.Concator
	}
	return Concator_AND
}

func (x *ConditionGroup) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *ConditionGroup) GetConditionGroups() []*ConditionGroup {
	if x != nil {
		return x.ConditionGroups
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xbf,
	0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x42, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x01,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x33, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3c, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x22, 0x53, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2b, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2a, 0x1b, 0x0a, 0x08, 0x43,
	0x6f, 0x6e, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x44, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x59, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca,
	0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []any{
	(Concator)(0),          // 0: proto.Concator
	(*Empty)(nil),          // 1: proto.Empty
	(*ListOption)(nil),     // 2: proto.ListOption
	(*Pager)(nil),          // 3: proto.Pager
	(*Sorter)(nil),         // 4: proto.Sorter
	(*Condition)(nil),      // 5: proto.Condition
	(*ConditionGroup)(nil), // 6: proto.ConditionGroup
}
var file_common_proto_depIdxs = []int32{
	3, // 0: proto.ListOption.pager:type_name -> proto.Pager
	4, // 1: proto.ListOption.sorters:type_name -> proto.Sorter
	6, // 2: proto.ListOption.conditionGroup:type_name -> proto.ConditionGroup
	0, // 3: proto.ConditionGroup.concator:type_name -> proto.Concator
	5, // 4: proto.ConditionGroup.conditions:type_name -> proto.Condition
	6, // 5: proto.ConditionGroup.conditionGroups:type_name -> proto.ConditionGroup
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListOption); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Pager); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Sorter); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Condition); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ConditionGroup); i {
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
	file_common_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
