// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/directoryd.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TableID int32

const (
	TableID_IMSI_TO_HWID     TableID = 0
	TableID_HWID_TO_HOSTNAME TableID = 1
)

var TableID_name = map[int32]string{
	0: "IMSI_TO_HWID",
	1: "HWID_TO_HOSTNAME",
}

var TableID_value = map[string]int32{
	"IMSI_TO_HWID":     0,
	"HWID_TO_HOSTNAME": 1,
}

func (x TableID) String() string {
	return proto.EnumName(TableID_name, int32(x))
}

func (TableID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{0}
}

type GetLocationRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Table                TableID  `protobuf:"varint,2,opt,name=table,proto3,enum=magma.orc8r.TableID" json:"table,omitempty"`
	NetworkID            string   `protobuf:"bytes,3,opt,name=networkID,proto3" json:"networkID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLocationRequest) Reset()         { *m = GetLocationRequest{} }
func (m *GetLocationRequest) String() string { return proto.CompactTextString(m) }
func (*GetLocationRequest) ProtoMessage()    {}
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{0}
}

func (m *GetLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLocationRequest.Unmarshal(m, b)
}
func (m *GetLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLocationRequest.Marshal(b, m, deterministic)
}
func (m *GetLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLocationRequest.Merge(m, src)
}
func (m *GetLocationRequest) XXX_Size() int {
	return xxx_messageInfo_GetLocationRequest.Size(m)
}
func (m *GetLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLocationRequest proto.InternalMessageInfo

func (m *GetLocationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetLocationRequest) GetTable() TableID {
	if m != nil {
		return m.Table
	}
	return TableID_IMSI_TO_HWID
}

func (m *GetLocationRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

type DeleteLocationRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Table                TableID  `protobuf:"varint,2,opt,name=table,proto3,enum=magma.orc8r.TableID" json:"table,omitempty"`
	NetworkID            string   `protobuf:"bytes,3,opt,name=networkID,proto3" json:"networkID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLocationRequest) Reset()         { *m = DeleteLocationRequest{} }
func (m *DeleteLocationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLocationRequest) ProtoMessage()    {}
func (*DeleteLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{1}
}

func (m *DeleteLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLocationRequest.Unmarshal(m, b)
}
func (m *DeleteLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLocationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLocationRequest.Merge(m, src)
}
func (m *DeleteLocationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteLocationRequest.Size(m)
}
func (m *DeleteLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLocationRequest proto.InternalMessageInfo

func (m *DeleteLocationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteLocationRequest) GetTable() TableID {
	if m != nil {
		return m.Table
	}
	return TableID_IMSI_TO_HWID
}

func (m *DeleteLocationRequest) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

type LocationRecord struct {
	// Object location. Client is responsible for location encoding and decoding
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationRecord) Reset()         { *m = LocationRecord{} }
func (m *LocationRecord) String() string { return proto.CompactTextString(m) }
func (*LocationRecord) ProtoMessage()    {}
func (*LocationRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{2}
}

func (m *LocationRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationRecord.Unmarshal(m, b)
}
func (m *LocationRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationRecord.Marshal(b, m, deterministic)
}
func (m *LocationRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationRecord.Merge(m, src)
}
func (m *LocationRecord) XXX_Size() int {
	return xxx_messageInfo_LocationRecord.Size(m)
}
func (m *LocationRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LocationRecord proto.InternalMessageInfo

func (m *LocationRecord) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type UpdateDirectoryLocationRequest struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Record               *LocationRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	Table                TableID         `protobuf:"varint,3,opt,name=table,proto3,enum=magma.orc8r.TableID" json:"table,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateDirectoryLocationRequest) Reset()         { *m = UpdateDirectoryLocationRequest{} }
func (m *UpdateDirectoryLocationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDirectoryLocationRequest) ProtoMessage()    {}
func (*UpdateDirectoryLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{3}
}

func (m *UpdateDirectoryLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDirectoryLocationRequest.Unmarshal(m, b)
}
func (m *UpdateDirectoryLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDirectoryLocationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDirectoryLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDirectoryLocationRequest.Merge(m, src)
}
func (m *UpdateDirectoryLocationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDirectoryLocationRequest.Size(m)
}
func (m *UpdateDirectoryLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDirectoryLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDirectoryLocationRequest proto.InternalMessageInfo

func (m *UpdateDirectoryLocationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateDirectoryLocationRequest) GetRecord() *LocationRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *UpdateDirectoryLocationRequest) GetTable() TableID {
	if m != nil {
		return m.Table
	}
	return TableID_IMSI_TO_HWID
}

type UpdateRecordRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location             string            `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Fields               map[string]string `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateRecordRequest) Reset()         { *m = UpdateRecordRequest{} }
func (m *UpdateRecordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRecordRequest) ProtoMessage()    {}
func (*UpdateRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{4}
}

func (m *UpdateRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRecordRequest.Unmarshal(m, b)
}
func (m *UpdateRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRecordRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRecordRequest.Merge(m, src)
}
func (m *UpdateRecordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRecordRequest.Size(m)
}
func (m *UpdateRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRecordRequest proto.InternalMessageInfo

func (m *UpdateRecordRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRecordRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *UpdateRecordRequest) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type DirectoryField struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectoryField) Reset()         { *m = DirectoryField{} }
func (m *DirectoryField) String() string { return proto.CompactTextString(m) }
func (*DirectoryField) ProtoMessage()    {}
func (*DirectoryField) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{5}
}

func (m *DirectoryField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectoryField.Unmarshal(m, b)
}
func (m *DirectoryField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectoryField.Marshal(b, m, deterministic)
}
func (m *DirectoryField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryField.Merge(m, src)
}
func (m *DirectoryField) XXX_Size() int {
	return xxx_messageInfo_DirectoryField.Size(m)
}
func (m *DirectoryField) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryField.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryField proto.InternalMessageInfo

func (m *DirectoryField) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DirectoryField) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DeleteRecordRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRecordRequest) Reset()         { *m = DeleteRecordRequest{} }
func (m *DeleteRecordRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRecordRequest) ProtoMessage()    {}
func (*DeleteRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{6}
}

func (m *DeleteRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRecordRequest.Unmarshal(m, b)
}
func (m *DeleteRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRecordRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRecordRequest.Merge(m, src)
}
func (m *DeleteRecordRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRecordRequest.Size(m)
}
func (m *DeleteRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRecordRequest proto.InternalMessageInfo

func (m *DeleteRecordRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDirectoryFieldRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FieldKey             string   `protobuf:"bytes,2,opt,name=field_key,json=fieldKey,proto3" json:"field_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDirectoryFieldRequest) Reset()         { *m = GetDirectoryFieldRequest{} }
func (m *GetDirectoryFieldRequest) String() string { return proto.CompactTextString(m) }
func (*GetDirectoryFieldRequest) ProtoMessage()    {}
func (*GetDirectoryFieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{7}
}

func (m *GetDirectoryFieldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDirectoryFieldRequest.Unmarshal(m, b)
}
func (m *GetDirectoryFieldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDirectoryFieldRequest.Marshal(b, m, deterministic)
}
func (m *GetDirectoryFieldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDirectoryFieldRequest.Merge(m, src)
}
func (m *GetDirectoryFieldRequest) XXX_Size() int {
	return xxx_messageInfo_GetDirectoryFieldRequest.Size(m)
}
func (m *GetDirectoryFieldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDirectoryFieldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDirectoryFieldRequest proto.InternalMessageInfo

func (m *GetDirectoryFieldRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetDirectoryFieldRequest) GetFieldKey() string {
	if m != nil {
		return m.FieldKey
	}
	return ""
}

type DirectoryRecord struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationHistory      []string          `protobuf:"bytes,2,rep,name=location_history,json=locationHistory,proto3" json:"location_history,omitempty"`
	Fields               map[string]string `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DirectoryRecord) Reset()         { *m = DirectoryRecord{} }
func (m *DirectoryRecord) String() string { return proto.CompactTextString(m) }
func (*DirectoryRecord) ProtoMessage()    {}
func (*DirectoryRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{8}
}

func (m *DirectoryRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectoryRecord.Unmarshal(m, b)
}
func (m *DirectoryRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectoryRecord.Marshal(b, m, deterministic)
}
func (m *DirectoryRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryRecord.Merge(m, src)
}
func (m *DirectoryRecord) XXX_Size() int {
	return xxx_messageInfo_DirectoryRecord.Size(m)
}
func (m *DirectoryRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryRecord proto.InternalMessageInfo

func (m *DirectoryRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DirectoryRecord) GetLocationHistory() []string {
	if m != nil {
		return m.LocationHistory
	}
	return nil
}

func (m *DirectoryRecord) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type AllDirectoryRecords struct {
	Records              []*DirectoryRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AllDirectoryRecords) Reset()         { *m = AllDirectoryRecords{} }
func (m *AllDirectoryRecords) String() string { return proto.CompactTextString(m) }
func (*AllDirectoryRecords) ProtoMessage()    {}
func (*AllDirectoryRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_f02336ef077163fd, []int{9}
}

func (m *AllDirectoryRecords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllDirectoryRecords.Unmarshal(m, b)
}
func (m *AllDirectoryRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllDirectoryRecords.Marshal(b, m, deterministic)
}
func (m *AllDirectoryRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllDirectoryRecords.Merge(m, src)
}
func (m *AllDirectoryRecords) XXX_Size() int {
	return xxx_messageInfo_AllDirectoryRecords.Size(m)
}
func (m *AllDirectoryRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_AllDirectoryRecords.DiscardUnknown(m)
}

var xxx_messageInfo_AllDirectoryRecords proto.InternalMessageInfo

func (m *AllDirectoryRecords) GetRecords() []*DirectoryRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterEnum("magma.orc8r.TableID", TableID_name, TableID_value)
	proto.RegisterType((*GetLocationRequest)(nil), "magma.orc8r.GetLocationRequest")
	proto.RegisterType((*DeleteLocationRequest)(nil), "magma.orc8r.DeleteLocationRequest")
	proto.RegisterType((*LocationRecord)(nil), "magma.orc8r.LocationRecord")
	proto.RegisterType((*UpdateDirectoryLocationRequest)(nil), "magma.orc8r.UpdateDirectoryLocationRequest")
	proto.RegisterType((*UpdateRecordRequest)(nil), "magma.orc8r.UpdateRecordRequest")
	proto.RegisterMapType((map[string]string)(nil), "magma.orc8r.UpdateRecordRequest.FieldsEntry")
	proto.RegisterType((*DirectoryField)(nil), "magma.orc8r.DirectoryField")
	proto.RegisterType((*DeleteRecordRequest)(nil), "magma.orc8r.DeleteRecordRequest")
	proto.RegisterType((*GetDirectoryFieldRequest)(nil), "magma.orc8r.GetDirectoryFieldRequest")
	proto.RegisterType((*DirectoryRecord)(nil), "magma.orc8r.DirectoryRecord")
	proto.RegisterMapType((map[string]string)(nil), "magma.orc8r.DirectoryRecord.FieldsEntry")
	proto.RegisterType((*AllDirectoryRecords)(nil), "magma.orc8r.AllDirectoryRecords")
}

func init() { proto.RegisterFile("orc8r/protos/directoryd.proto", fileDescriptor_f02336ef077163fd) }

var fileDescriptor_f02336ef077163fd = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x8d, 0x6d, 0xb5, 0xfd, 0x32, 0xa9, 0x5c, 0x77, 0xdb, 0x0f, 0x8c, 0xdb, 0x42, 0xb4, 0x52,
	0xa5, 0x50, 0xaa, 0x44, 0xa4, 0x12, 0x2a, 0x9c, 0x28, 0x72, 0x48, 0x2d, 0x08, 0x95, 0xdc, 0x00,
	0x82, 0x4b, 0xe4, 0xda, 0x4b, 0xb1, 0xea, 0x64, 0x5b, 0x7b, 0xd3, 0x2a, 0x37, 0xfe, 0x02, 0x7f,
	0x89, 0x1b, 0x67, 0xfe, 0x10, 0xf2, 0xda, 0x49, 0xbc, 0x8e, 0x93, 0x70, 0x40, 0x9c, 0x62, 0x4f,
	0x9e, 0xdf, 0x7b, 0xf3, 0x3c, 0xb3, 0x86, 0x3d, 0x1a, 0xba, 0xc7, 0x61, 0xe3, 0x3a, 0xa4, 0x8c,
	0x46, 0x0d, 0xcf, 0x0f, 0x89, 0xcb, 0x68, 0x38, 0xf2, 0xea, 0xbc, 0x82, 0x2a, 0x7d, 0xe7, 0xb2,
	0xef, 0xd4, 0x39, 0xc8, 0x78, 0x20, 0x60, 0x5d, 0xda, 0xef, 0xd3, 0x41, 0x82, 0xc3, 0x03, 0x40,
	0x6d, 0xc2, 0xde, 0x52, 0xd7, 0x61, 0x3e, 0x1d, 0xd8, 0xe4, 0x66, 0x48, 0x22, 0x86, 0x54, 0x90,
	0x7d, 0x4f, 0x97, 0xaa, 0x52, 0xad, 0x6c, 0xcb, 0xbe, 0x87, 0x0e, 0x60, 0x85, 0x39, 0x17, 0x01,
	0xd1, 0xe5, 0xaa, 0x54, 0x53, 0x9b, 0xdb, 0xf5, 0x0c, 0x7b, 0xbd, 0x1b, 0xff, 0x63, 0x99, 0x76,
	0x02, 0x41, 0xbb, 0x50, 0x1e, 0x10, 0x76, 0x47, 0xc3, 0x2b, 0xcb, 0xd4, 0x15, 0x4e, 0x31, 0x2d,
	0xe0, 0x1b, 0xf8, 0xdf, 0x24, 0x01, 0x61, 0xe4, 0xdf, 0x49, 0x1e, 0x82, 0x3a, 0x15, 0x73, 0x69,
	0xe8, 0x21, 0x03, 0xfe, 0x0b, 0xd2, 0x4a, 0xaa, 0x38, 0xb9, 0xc7, 0xdf, 0x25, 0x78, 0xf8, 0xfe,
	0xda, 0x73, 0x18, 0x31, 0xc7, 0x99, 0x2e, 0xb3, 0x7a, 0x04, 0xab, 0x21, 0x27, 0xe6, 0x5e, 0x2b,
	0xcd, 0x1d, 0xc1, 0xab, 0xa8, 0x6d, 0xa7, 0xd0, 0x69, 0x7f, 0xca, 0xd2, 0xfe, 0xf0, 0x0f, 0x09,
	0xb6, 0x12, 0x4f, 0x29, 0xc9, 0x1c, 0x23, 0xd9, 0xbe, 0x64, 0xb1, 0x2f, 0x64, 0xc2, 0xea, 0x17,
	0x9f, 0x04, 0x5e, 0xa4, 0x2b, 0x55, 0xa5, 0x56, 0x69, 0x1e, 0x0a, 0x82, 0x05, 0xec, 0xf5, 0xd7,
	0x1c, 0xde, 0x1a, 0xb0, 0x70, 0x64, 0xa7, 0xcf, 0x1a, 0xcf, 0xa1, 0x92, 0x29, 0x23, 0x0d, 0x94,
	0x2b, 0x32, 0x4a, 0x1d, 0xc4, 0x97, 0x68, 0x1b, 0x56, 0x6e, 0x9d, 0x60, 0x48, 0x52, 0xfd, 0xe4,
	0xe6, 0x85, 0x7c, 0x2c, 0xe1, 0x63, 0x50, 0x27, 0x89, 0x72, 0x8e, 0x3f, 0x7d, 0x1a, 0xef, 0xc3,
	0x56, 0x32, 0x33, 0x0b, 0xbb, 0xc7, 0x6d, 0xd0, 0xdb, 0x84, 0x89, 0x1a, 0xf3, 0x92, 0xda, 0x81,
	0x32, 0xef, 0xa8, 0x17, 0x1b, 0x48, 0xa3, 0xe2, 0x85, 0x37, 0x64, 0x84, 0x7f, 0x4a, 0xb0, 0x31,
	0xa1, 0x49, 0x47, 0x26, 0x4f, 0xf0, 0x18, 0xb4, 0x71, 0xb4, 0xbd, 0xaf, 0x7e, 0x14, 0x23, 0x75,
	0xb9, 0xaa, 0xd4, 0xca, 0xf6, 0xc6, 0xb8, 0x7e, 0x9a, 0x94, 0xd1, 0xcb, 0x5c, 0xf2, 0x35, 0x21,
	0xf9, 0x9c, 0xd0, 0xdf, 0x4e, 0xbd, 0x03, 0x5b, 0x27, 0x41, 0x90, 0x13, 0x89, 0xd0, 0x33, 0x58,
	0x4b, 0xe6, 0x30, 0xd2, 0x25, 0x6e, 0x6a, 0x77, 0x91, 0x29, 0x7b, 0x0c, 0x3e, 0x78, 0x0a, 0x6b,
	0xe9, 0x6c, 0x22, 0x0d, 0xd6, 0xad, 0xce, 0xb9, 0xd5, 0xeb, 0x9e, 0xf5, 0x4e, 0x3f, 0x5a, 0xa6,
	0x56, 0x42, 0xdb, 0xa0, 0xc5, 0x57, 0xbc, 0x72, 0x76, 0xde, 0x7d, 0x77, 0xd2, 0x69, 0x69, 0x52,
	0xf3, 0x9b, 0x0c, 0xda, 0x84, 0xef, 0x9c, 0x84, 0xb7, 0xbe, 0x4b, 0x50, 0x07, 0x2a, 0x99, 0x63,
	0x07, 0x3d, 0x12, 0xd4, 0x67, 0x0f, 0x24, 0x63, 0xd1, 0x4a, 0xe1, 0x12, 0xb2, 0x41, 0x4d, 0x26,
	0x78, 0xc2, 0xf8, 0xa4, 0x60, 0xbc, 0xe7, 0x2d, 0xb4, 0xb1, 0x29, 0x80, 0x3f, 0x50, 0x3f, 0xe6,
	0xb4, 0x40, 0x15, 0x4f, 0x2a, 0x84, 0xc5, 0x8c, 0x8a, 0x8e, 0xb1, 0x42, 0xaa, 0xe6, 0x2f, 0x19,
	0xee, 0xb7, 0x1d, 0x46, 0xee, 0x9c, 0xd1, 0x4c, 0x12, 0x2d, 0x58, 0xcf, 0x2e, 0x1f, 0xaa, 0x2e,
	0xdb, 0xcb, 0x62, 0xb7, 0x2d, 0x58, 0xcf, 0xee, 0x48, 0x8e, 0xa6, 0x60, 0x7d, 0x8a, 0x69, 0x3e,
	0xc1, 0xe6, 0xcc, 0x0e, 0xa1, 0xfd, 0xfc, 0xdb, 0x29, 0xdc, 0xb1, 0xdc, 0x3b, 0x12, 0x31, 0xb8,
	0x84, 0xce, 0xe0, 0x5e, 0x9b, 0xb0, 0xa2, 0x61, 0x9c, 0x75, 0x62, 0x88, 0xf6, 0x0b, 0x1e, 0xc2,
	0xa5, 0x57, 0x7b, 0x9f, 0x77, 0x38, 0xa8, 0x91, 0x7c, 0xdd, 0xdc, 0x80, 0x0e, 0xbd, 0xc6, 0x25,
	0x4d, 0x3f, 0x73, 0x17, 0xab, 0xfc, 0xf7, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x94,
	0x46, 0x01, 0x29, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DirectoryServiceClient is the client API for DirectoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DirectoryServiceClient interface {
	// Get location of an object from the directory service
	// Throws UNKNOWN if object ID does not exist
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*LocationRecord, error)
	// Update the location record of an object in the directory service
	UpdateLocation(ctx context.Context, in *UpdateDirectoryLocationRequest, opts ...grpc.CallOption) (*Void, error)
	// Delete location record of an object from the directory service
	// Throws UNKNOWN if object ID does not exist
	DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*Void, error)
}

type directoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDirectoryServiceClient(cc *grpc.ClientConn) DirectoryServiceClient {
	return &directoryServiceClient{cc}
}

func (c *directoryServiceClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*LocationRecord, error) {
	out := new(LocationRecord)
	err := c.cc.Invoke(ctx, "/magma.orc8r.DirectoryService/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) UpdateLocation(ctx context.Context, in *UpdateDirectoryLocationRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.DirectoryService/UpdateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.DirectoryService/DeleteLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryServiceServer is the server API for DirectoryService service.
type DirectoryServiceServer interface {
	// Get location of an object from the directory service
	// Throws UNKNOWN if object ID does not exist
	GetLocation(context.Context, *GetLocationRequest) (*LocationRecord, error)
	// Update the location record of an object in the directory service
	UpdateLocation(context.Context, *UpdateDirectoryLocationRequest) (*Void, error)
	// Delete location record of an object from the directory service
	// Throws UNKNOWN if object ID does not exist
	DeleteLocation(context.Context, *DeleteLocationRequest) (*Void, error)
}

// UnimplementedDirectoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDirectoryServiceServer struct {
}

func (*UnimplementedDirectoryServiceServer) GetLocation(ctx context.Context, req *GetLocationRequest) (*LocationRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (*UnimplementedDirectoryServiceServer) UpdateLocation(ctx context.Context, req *UpdateDirectoryLocationRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (*UnimplementedDirectoryServiceServer) DeleteLocation(ctx context.Context, req *DeleteLocationRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocation not implemented")
}

func RegisterDirectoryServiceServer(s *grpc.Server, srv DirectoryServiceServer) {
	s.RegisterService(&_DirectoryService_serviceDesc, srv)
}

func _DirectoryService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.DirectoryService/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDirectoryLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.DirectoryService/UpdateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).UpdateLocation(ctx, req.(*UpdateDirectoryLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_DeleteLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).DeleteLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.DirectoryService/DeleteLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).DeleteLocation(ctx, req.(*DeleteLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DirectoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.DirectoryService",
	HandlerType: (*DirectoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocation",
			Handler:    _DirectoryService_GetLocation_Handler,
		},
		{
			MethodName: "UpdateLocation",
			Handler:    _DirectoryService_UpdateLocation_Handler,
		},
		{
			MethodName: "DeleteLocation",
			Handler:    _DirectoryService_DeleteLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/directoryd.proto",
}

// GatewayDirectoryServiceClient is the client API for GatewayDirectoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayDirectoryServiceClient interface {
	// Update the directory record of an object in the directory service
	UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*Void, error)
	// Delete directory record of an object from the directory service
	// Throws UNKNOWN if object ID does not exist
	DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*Void, error)
	// Get directory field for a given id and key
	GetDirectoryField(ctx context.Context, in *GetDirectoryFieldRequest, opts ...grpc.CallOption) (*DirectoryField, error)
	// Get all directory records
	GetAllDirectoryRecords(ctx context.Context, in *Void, opts ...grpc.CallOption) (*AllDirectoryRecords, error)
}

type gatewayDirectoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayDirectoryServiceClient(cc *grpc.ClientConn) GatewayDirectoryServiceClient {
	return &gatewayDirectoryServiceClient{cc}
}

func (c *gatewayDirectoryServiceClient) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.GatewayDirectoryService/UpdateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDirectoryServiceClient) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.GatewayDirectoryService/DeleteRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDirectoryServiceClient) GetDirectoryField(ctx context.Context, in *GetDirectoryFieldRequest, opts ...grpc.CallOption) (*DirectoryField, error) {
	out := new(DirectoryField)
	err := c.cc.Invoke(ctx, "/magma.orc8r.GatewayDirectoryService/GetDirectoryField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDirectoryServiceClient) GetAllDirectoryRecords(ctx context.Context, in *Void, opts ...grpc.CallOption) (*AllDirectoryRecords, error) {
	out := new(AllDirectoryRecords)
	err := c.cc.Invoke(ctx, "/magma.orc8r.GatewayDirectoryService/GetAllDirectoryRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayDirectoryServiceServer is the server API for GatewayDirectoryService service.
type GatewayDirectoryServiceServer interface {
	// Update the directory record of an object in the directory service
	UpdateRecord(context.Context, *UpdateRecordRequest) (*Void, error)
	// Delete directory record of an object from the directory service
	// Throws UNKNOWN if object ID does not exist
	DeleteRecord(context.Context, *DeleteRecordRequest) (*Void, error)
	// Get directory field for a given id and key
	GetDirectoryField(context.Context, *GetDirectoryFieldRequest) (*DirectoryField, error)
	// Get all directory records
	GetAllDirectoryRecords(context.Context, *Void) (*AllDirectoryRecords, error)
}

// UnimplementedGatewayDirectoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayDirectoryServiceServer struct {
}

func (*UnimplementedGatewayDirectoryServiceServer) UpdateRecord(ctx context.Context, req *UpdateRecordRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}
func (*UnimplementedGatewayDirectoryServiceServer) DeleteRecord(ctx context.Context, req *DeleteRecordRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (*UnimplementedGatewayDirectoryServiceServer) GetDirectoryField(ctx context.Context, req *GetDirectoryFieldRequest) (*DirectoryField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDirectoryField not implemented")
}
func (*UnimplementedGatewayDirectoryServiceServer) GetAllDirectoryRecords(ctx context.Context, req *Void) (*AllDirectoryRecords, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDirectoryRecords not implemented")
}

func RegisterGatewayDirectoryServiceServer(s *grpc.Server, srv GatewayDirectoryServiceServer) {
	s.RegisterService(&_GatewayDirectoryService_serviceDesc, srv)
}

func _GatewayDirectoryService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDirectoryServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.GatewayDirectoryService/UpdateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDirectoryServiceServer).UpdateRecord(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDirectoryService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDirectoryServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.GatewayDirectoryService/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDirectoryServiceServer).DeleteRecord(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDirectoryService_GetDirectoryField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDirectoryFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDirectoryServiceServer).GetDirectoryField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.GatewayDirectoryService/GetDirectoryField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDirectoryServiceServer).GetDirectoryField(ctx, req.(*GetDirectoryFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDirectoryService_GetAllDirectoryRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDirectoryServiceServer).GetAllDirectoryRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.GatewayDirectoryService/GetAllDirectoryRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDirectoryServiceServer).GetAllDirectoryRecords(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayDirectoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.GatewayDirectoryService",
	HandlerType: (*GatewayDirectoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateRecord",
			Handler:    _GatewayDirectoryService_UpdateRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _GatewayDirectoryService_DeleteRecord_Handler,
		},
		{
			MethodName: "GetDirectoryField",
			Handler:    _GatewayDirectoryService_GetDirectoryField_Handler,
		},
		{
			MethodName: "GetAllDirectoryRecords",
			Handler:    _GatewayDirectoryService_GetAllDirectoryRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/directoryd.proto",
}