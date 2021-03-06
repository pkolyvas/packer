// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mongodb/v1/resource_preset.proto

package mongodb

import (
	fmt "fmt"
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

// A ResourcePreset resource for describing hardware configuration presets.
type ResourcePreset struct {
	// ID of the ResourcePreset resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IDs of availability zones where the resource preset is available.
	ZoneIds []string `protobuf:"bytes,2,rep,name=zone_ids,json=zoneIds,proto3" json:"zone_ids,omitempty"`
	// Number of CPU cores for a MongoDB host created with the preset.
	Cores int64 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	// RAM volume for a MongoDB host created with the preset, in bytes.
	Memory               int64    `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcePreset) Reset()         { *m = ResourcePreset{} }
func (m *ResourcePreset) String() string { return proto.CompactTextString(m) }
func (*ResourcePreset) ProtoMessage()    {}
func (*ResourcePreset) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c48b84d9988201, []int{0}
}

func (m *ResourcePreset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourcePreset.Unmarshal(m, b)
}
func (m *ResourcePreset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourcePreset.Marshal(b, m, deterministic)
}
func (m *ResourcePreset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcePreset.Merge(m, src)
}
func (m *ResourcePreset) XXX_Size() int {
	return xxx_messageInfo_ResourcePreset.Size(m)
}
func (m *ResourcePreset) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcePreset.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcePreset proto.InternalMessageInfo

func (m *ResourcePreset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourcePreset) GetZoneIds() []string {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func (m *ResourcePreset) GetCores() int64 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *ResourcePreset) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourcePreset)(nil), "yandex.cloud.mdb.mongodb.v1.ResourcePreset")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mongodb/v1/resource_preset.proto", fileDescriptor_07c48b84d9988201)
}

var fileDescriptor_07c48b84d9988201 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xd9, 0x5d, 0x3d, 0xbd, 0x14, 0x57, 0x04, 0x91, 0x15, 0x9b, 0xc5, 0x6a, 0x9b, 0x4b,
	0x58, 0x2c, 0xed, 0x6c, 0xe4, 0x3a, 0x49, 0x69, 0x73, 0x98, 0xcc, 0x10, 0x03, 0x26, 0x73, 0x24,
	0xbb, 0x87, 0xeb, 0xaf, 0x17, 0x93, 0xd4, 0xd7, 0xcd, 0x37, 0xbc, 0x0f, 0xde, 0x63, 0xd3, 0xfa,
	0x19, 0x00, 0x7f, 0xa4, 0xf9, 0xa6, 0x05, 0xa4, 0x07, 0x2d, 0x3d, 0x05, 0x4b, 0xa0, 0xe5, 0x79,
	0x92, 0x11, 0x13, 0x2d, 0xd1, 0xe0, 0xf1, 0x14, 0x31, 0xe1, 0x2c, 0x4e, 0x91, 0x66, 0xe2, 0x8f,
	0x45, 0x11, 0x59, 0x11, 0x1e, 0xb4, 0xa8, 0x8a, 0x38, 0x4f, 0x4f, 0x8e, 0xed, 0x54, 0xb5, 0xde,
	0xb3, 0xc4, 0x77, 0xac, 0x75, 0xd0, 0x37, 0x43, 0x33, 0x6e, 0x55, 0xeb, 0x80, 0x3f, 0xb0, 0xdb,
	0x5f, 0x0a, 0x78, 0x74, 0x90, 0xfa, 0x76, 0xe8, 0xc6, 0xad, 0xba, 0xf9, 0xe7, 0x03, 0x24, 0x7e,
	0xc7, 0xae, 0x0d, 0x45, 0x4c, 0x7d, 0x37, 0x34, 0x63, 0xa7, 0x0a, 0xf0, 0x7b, 0xb6, 0xf1, 0xe8,
	0x29, 0xae, 0xfd, 0x55, 0x7e, 0x57, 0x7a, 0x3d, 0x7c, 0xbc, 0x59, 0x37, 0x7f, 0x2d, 0x5a, 0x18,
	0xf2, 0xb2, 0x94, 0xda, 0x97, 0x1d, 0x96, 0xf6, 0x16, 0x43, 0xae, 0x2b, 0x2f, 0x0c, 0x7c, 0xa9,
	0xa7, 0xde, 0xe4, 0xe8, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0xbd, 0x87, 0x16, 0x0e,
	0x01, 0x00, 0x00,
}
