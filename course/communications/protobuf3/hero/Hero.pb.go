// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: Hero.proto

package hero

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

type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alias  string             `protobuf:"bytes,2,opt,name=Alias,proto3" json:"Alias,omitempty"`
	Skills []*Hero_HeroSkills `protobuf:"bytes,3,rep,name=Skills,proto3" json:"Skills,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_Hero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_Hero_proto_rawDescGZIP(), []int{0}
}

func (x *Hero) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hero) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Hero) GetSkills() []*Hero_HeroSkills {
	if x != nil {
		return x.Skills
	}
	return nil
}

type Hero_HeroSkills struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind   string  `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Damage float32 `protobuf:"fixed32,2,opt,name=damage,proto3" json:"damage,omitempty"`
	Energy int32   `protobuf:"varint,3,opt,name=energy,proto3" json:"energy,omitempty"`
}

func (x *Hero_HeroSkills) Reset() {
	*x = Hero_HeroSkills{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Hero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero_HeroSkills) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero_HeroSkills) ProtoMessage() {}

func (x *Hero_HeroSkills) ProtoReflect() protoreflect.Message {
	mi := &file_Hero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero_HeroSkills.ProtoReflect.Descriptor instead.
func (*Hero_HeroSkills) Descriptor() ([]byte, []int) {
	return file_Hero_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Hero_HeroSkills) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Hero_HeroSkills) GetDamage() float32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *Hero_HeroSkills) GetEnergy() int32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

var File_Hero_proto protoreflect.FileDescriptor

var file_Hero_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x48, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x65,
	0x72, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x48, 0x65, 0x72,
	0x6f, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x06, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x1a, 0x50, 0x0a, 0x0a, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Hero_proto_rawDescOnce sync.Once
	file_Hero_proto_rawDescData = file_Hero_proto_rawDesc
)

func file_Hero_proto_rawDescGZIP() []byte {
	file_Hero_proto_rawDescOnce.Do(func() {
		file_Hero_proto_rawDescData = protoimpl.X.CompressGZIP(file_Hero_proto_rawDescData)
	})
	return file_Hero_proto_rawDescData
}

var file_Hero_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Hero_proto_goTypes = []interface{}{
	(*Hero)(nil),            // 0: hero.Hero
	(*Hero_HeroSkills)(nil), // 1: hero.Hero.HeroSkills
}
var file_Hero_proto_depIdxs = []int32{
	1, // 0: hero.Hero.Skills:type_name -> hero.Hero.HeroSkills
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Hero_proto_init() }
func file_Hero_proto_init() {
	if File_Hero_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Hero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
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
		file_Hero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero_HeroSkills); i {
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
			RawDescriptor: file_Hero_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Hero_proto_goTypes,
		DependencyIndexes: file_Hero_proto_depIdxs,
		MessageInfos:      file_Hero_proto_msgTypes,
	}.Build()
	File_Hero_proto = out.File
	file_Hero_proto_rawDesc = nil
	file_Hero_proto_goTypes = nil
	file_Hero_proto_depIdxs = nil
}