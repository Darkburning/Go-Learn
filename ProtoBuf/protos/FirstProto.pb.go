// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: ProtoBuf/FirstProto.proto

package protos

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

type PersonPhoneType int32

const (
	Person_MOBILE PersonPhoneType = 0
	Person_HOME   PersonPhoneType = 1
	Person_WORK   PersonPhoneType = 2
)

// Enum value maps for PersonPhoneType.
var (
	PersonPhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	PersonPhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x PersonPhoneType) Enum() *PersonPhoneType {
	p := new(PersonPhoneType)
	*p = x
	return p
}

func (x PersonPhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PersonPhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_ProtoBuf_FirstProto_proto_enumTypes[0].Descriptor()
}

func (PersonPhoneType) Type() protoreflect.EnumType {
	return &file_ProtoBuf_FirstProto_proto_enumTypes[0]
}

func (x PersonPhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PersonPhoneType.Descriptor instead.
func (PersonPhoneType) EnumDescriptor() ([]byte, []int) {
	return file_ProtoBuf_FirstProto_proto_rawDescGZIP(), []int{0, 0}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id     int32                `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email  string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones []*PersonPhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBuf_FirstProto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBuf_FirstProto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_ProtoBuf_FirstProto_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Person) GetPhones() []*PersonPhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

type AddressBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *AddressBook) Reset() {
	*x = AddressBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBuf_FirstProto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBook) ProtoMessage() {}

func (x *AddressBook) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBuf_FirstProto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBook.ProtoReflect.Descriptor instead.
func (*AddressBook) Descriptor() ([]byte, []int) {
	return file_ProtoBuf_FirstProto_proto_rawDescGZIP(), []int{1}
}

func (x *AddressBook) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

type PersonPhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string          `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   PersonPhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=PersonPhoneType" json:"type,omitempty"`
}

func (x *PersonPhoneNumber) Reset() {
	*x = PersonPhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBuf_FirstProto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonPhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonPhoneNumber) ProtoMessage() {}

func (x *PersonPhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBuf_FirstProto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonPhoneNumber.ProtoReflect.Descriptor instead.
func (*PersonPhoneNumber) Descriptor() ([]byte, []int) {
	return file_ProtoBuf_FirstProto_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PersonPhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PersonPhoneNumber) GetType() PersonPhoneType {
	if x != nil {
		return x.Type
	}
	return Person_MOBILE
}

var File_ProtoBuf_FirstProto_proto protoreflect.FileDescriptor

var file_ProtoBuf_FirstProto_proto_rawDesc = []byte{
	0x0a, 0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x06,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x2b, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x1a, 0x4c, 0x0a,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x09, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49,
	0x4c, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x22, 0x2e, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x2f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProtoBuf_FirstProto_proto_rawDescOnce sync.Once
	file_ProtoBuf_FirstProto_proto_rawDescData = file_ProtoBuf_FirstProto_proto_rawDesc
)

func file_ProtoBuf_FirstProto_proto_rawDescGZIP() []byte {
	file_ProtoBuf_FirstProto_proto_rawDescOnce.Do(func() {
		file_ProtoBuf_FirstProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtoBuf_FirstProto_proto_rawDescData)
	})
	return file_ProtoBuf_FirstProto_proto_rawDescData
}

var file_ProtoBuf_FirstProto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ProtoBuf_FirstProto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ProtoBuf_FirstProto_proto_goTypes = []interface{}{
	(PersonPhoneType)(0),      // 0: Person.phoneType
	(*Person)(nil),            // 1: Person
	(*AddressBook)(nil),       // 2: AddressBook
	(*PersonPhoneNumber)(nil), // 3: Person.phoneNumber
}
var file_ProtoBuf_FirstProto_proto_depIdxs = []int32{
	3, // 0: Person.phones:type_name -> Person.phoneNumber
	1, // 1: AddressBook.people:type_name -> Person
	0, // 2: Person.phoneNumber.type:type_name -> Person.phoneType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ProtoBuf_FirstProto_proto_init() }
func file_ProtoBuf_FirstProto_proto_init() {
	if File_ProtoBuf_FirstProto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProtoBuf_FirstProto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_ProtoBuf_FirstProto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBook); i {
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
		file_ProtoBuf_FirstProto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonPhoneNumber); i {
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
			RawDescriptor: file_ProtoBuf_FirstProto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProtoBuf_FirstProto_proto_goTypes,
		DependencyIndexes: file_ProtoBuf_FirstProto_proto_depIdxs,
		EnumInfos:         file_ProtoBuf_FirstProto_proto_enumTypes,
		MessageInfos:      file_ProtoBuf_FirstProto_proto_msgTypes,
	}.Build()
	File_ProtoBuf_FirstProto_proto = out.File
	file_ProtoBuf_FirstProto_proto_rawDesc = nil
	file_ProtoBuf_FirstProto_proto_goTypes = nil
	file_ProtoBuf_FirstProto_proto_depIdxs = nil
}