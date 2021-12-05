// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: fulcrum.proto

package grpc_fulcrum

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

type F_FromLeia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FPlaneta *F_Planeta `protobuf:"bytes,1,opt,name=f_planeta,json=fPlaneta,proto3" json:"f_planeta,omitempty"`
	FCiudad  *F_Ciudad  `protobuf:"bytes,2,opt,name=f_ciudad,json=fCiudad,proto3" json:"f_ciudad,omitempty"`
}

func (x *F_FromLeia) Reset() {
	*x = F_FromLeia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_FromLeia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_FromLeia) ProtoMessage() {}

func (x *F_FromLeia) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_FromLeia.ProtoReflect.Descriptor instead.
func (*F_FromLeia) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{0}
}

func (x *F_FromLeia) GetFPlaneta() *F_Planeta {
	if x != nil {
		return x.FPlaneta
	}
	return nil
}

func (x *F_FromLeia) GetFCiudad() *F_Ciudad {
	if x != nil {
		return x.FCiudad
	}
	return nil
}

type F_From_Informante struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FCommand *F_Command `protobuf:"bytes,1,opt,name=f_command,json=fCommand,proto3" json:"f_command,omitempty"`
	FReloj   *F_Reloj   `protobuf:"bytes,2,opt,name=f_reloj,json=fReloj,proto3" json:"f_reloj,omitempty"`
}

func (x *F_From_Informante) Reset() {
	*x = F_From_Informante{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_From_Informante) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_From_Informante) ProtoMessage() {}

func (x *F_From_Informante) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_From_Informante.ProtoReflect.Descriptor instead.
func (*F_From_Informante) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{1}
}

func (x *F_From_Informante) GetFCommand() *F_Command {
	if x != nil {
		return x.FCommand
	}
	return nil
}

func (x *F_From_Informante) GetFReloj() *F_Reloj {
	if x != nil {
		return x.FReloj
	}
	return nil
}

type F_To_Informante struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FReloj *F_Reloj `protobuf:"bytes,1,opt,name=f_reloj,json=fReloj,proto3" json:"f_reloj,omitempty"`
	FLog   *F_Log   `protobuf:"bytes,2,opt,name=f_log,json=fLog,proto3" json:"f_log,omitempty"`
}

func (x *F_To_Informante) Reset() {
	*x = F_To_Informante{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_To_Informante) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_To_Informante) ProtoMessage() {}

func (x *F_To_Informante) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_To_Informante.ProtoReflect.Descriptor instead.
func (*F_To_Informante) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{2}
}

func (x *F_To_Informante) GetFReloj() *F_Reloj {
	if x != nil {
		return x.FReloj
	}
	return nil
}

func (x *F_To_Informante) GetFLog() *F_Log {
	if x != nil {
		return x.FLog
	}
	return nil
}

type F_Merge_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FReloj *F_Reloj `protobuf:"bytes,1,opt,name=f_reloj,json=fReloj,proto3" json:"f_reloj,omitempty"`
	FLog   *F_Log   `protobuf:"bytes,2,opt,name=f_log,json=fLog,proto3" json:"f_log,omitempty"`
}

func (x *F_Merge_Data) Reset() {
	*x = F_Merge_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_Merge_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_Merge_Data) ProtoMessage() {}

func (x *F_Merge_Data) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_Merge_Data.ProtoReflect.Descriptor instead.
func (*F_Merge_Data) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{3}
}

func (x *F_Merge_Data) GetFReloj() *F_Reloj {
	if x != nil {
		return x.FReloj
	}
	return nil
}

func (x *F_Merge_Data) GetFLog() *F_Log {
	if x != nil {
		return x.FLog
	}
	return nil
}

type F_ToLeia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FRebeldes *F_Rebeldes `protobuf:"bytes,1,opt,name=f_rebeldes,json=fRebeldes,proto3" json:"f_rebeldes,omitempty"`
	FReloj    *F_Reloj    `protobuf:"bytes,2,opt,name=f_reloj,json=fReloj,proto3" json:"f_reloj,omitempty"`
}

func (x *F_ToLeia) Reset() {
	*x = F_ToLeia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_ToLeia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_ToLeia) ProtoMessage() {}

func (x *F_ToLeia) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_ToLeia.ProtoReflect.Descriptor instead.
func (*F_ToLeia) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{4}
}

func (x *F_ToLeia) GetFRebeldes() *F_Rebeldes {
	if x != nil {
		return x.FRebeldes
	}
	return nil
}

func (x *F_ToLeia) GetFReloj() *F_Reloj {
	if x != nil {
		return x.FReloj
	}
	return nil
}

type F_Planeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
}

func (x *F_Planeta) Reset() {
	*x = F_Planeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_Planeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_Planeta) ProtoMessage() {}

func (x *F_Planeta) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_Planeta.ProtoReflect.Descriptor instead.
func (*F_Planeta) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{5}
}

func (x *F_Planeta) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

type F_Ciudad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ciudad string `protobuf:"bytes,1,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
}

func (x *F_Ciudad) Reset() {
	*x = F_Ciudad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_Ciudad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_Ciudad) ProtoMessage() {}

func (x *F_Ciudad) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_Ciudad.ProtoReflect.Descriptor instead.
func (*F_Ciudad) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{6}
}

func (x *F_Ciudad) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

type F_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *F_Log) Reset() {
	*x = F_Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_Log) ProtoMessage() {}

func (x *F_Log) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_Log.ProtoReflect.Descriptor instead.
func (*F_Log) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{7}
}

func (x *F_Log) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

type F_Rebeldes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cantidad int64 `protobuf:"varint,1,opt,name=cantidad,proto3" json:"cantidad,omitempty"`
}

func (x *F_Rebeldes) Reset() {
	*x = F_Rebeldes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_Rebeldes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_Rebeldes) ProtoMessage() {}

func (x *F_Rebeldes) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_Rebeldes.ProtoReflect.Descriptor instead.
func (*F_Rebeldes) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{8}
}

func (x *F_Rebeldes) GetCantidad() int64 {
	if x != nil {
		return x.Cantidad
	}
	return 0
}

type F_Reloj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z int64 `protobuf:"varint,3,opt,name=Z,proto3" json:"Z,omitempty"`
}

func (x *F_Reloj) Reset() {
	*x = F_Reloj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_Reloj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_Reloj) ProtoMessage() {}

func (x *F_Reloj) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_Reloj.ProtoReflect.Descriptor instead.
func (*F_Reloj) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{9}
}

func (x *F_Reloj) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *F_Reloj) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *F_Reloj) GetZ() int64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type F_Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *F_Command) Reset() {
	*x = F_Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *F_Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*F_Command) ProtoMessage() {}

func (x *F_Command) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use F_Command.ProtoReflect.Descriptor instead.
func (*F_Command) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{10}
}

func (x *F_Command) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type Fantasma struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fantasma string `protobuf:"bytes,1,opt,name=fantasma,proto3" json:"fantasma,omitempty"`
}

func (x *Fantasma) Reset() {
	*x = Fantasma{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcrum_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fantasma) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fantasma) ProtoMessage() {}

func (x *Fantasma) ProtoReflect() protoreflect.Message {
	mi := &file_fulcrum_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fantasma.ProtoReflect.Descriptor instead.
func (*Fantasma) Descriptor() ([]byte, []int) {
	return file_fulcrum_proto_rawDescGZIP(), []int{11}
}

func (x *Fantasma) GetFantasma() string {
	if x != nil {
		return x.Fantasma
	}
	return ""
}

var File_fulcrum_proto protoreflect.FileDescriptor

var file_fulcrum_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5b, 0x0a, 0x0a, 0x46, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x65, 0x69, 0x61, 0x12, 0x27, 0x0a,
	0x09, 0x66, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x46, 0x5f, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x52, 0x08, 0x66, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x08, 0x66, 0x5f, 0x63, 0x69, 0x75, 0x64,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x5f, 0x43, 0x69, 0x75,
	0x64, 0x61, 0x64, 0x52, 0x07, 0x66, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x22, 0x5f, 0x0a, 0x11,
	0x46, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74,
	0x65, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x46, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x08, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x66, 0x5f,
	0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x46, 0x5f,
	0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x52, 0x06, 0x66, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x22, 0x51, 0x0a,
	0x0f, 0x46, 0x5f, 0x54, 0x6f, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65,
	0x12, 0x21, 0x0a, 0x07, 0x66, 0x5f, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x46, 0x5f, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x52, 0x06, 0x66, 0x52, 0x65,
	0x6c, 0x6f, 0x6a, 0x12, 0x1b, 0x0a, 0x05, 0x66, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x46, 0x5f, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x66, 0x4c, 0x6f, 0x67,
	0x22, 0x4e, 0x0a, 0x0c, 0x46, 0x5f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x21, 0x0a, 0x07, 0x66, 0x5f, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x46, 0x5f, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x52, 0x06, 0x66, 0x52, 0x65,
	0x6c, 0x6f, 0x6a, 0x12, 0x1b, 0x0a, 0x05, 0x66, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x46, 0x5f, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x66, 0x4c, 0x6f, 0x67,
	0x22, 0x59, 0x0a, 0x08, 0x46, 0x5f, 0x54, 0x6f, 0x4c, 0x65, 0x69, 0x61, 0x12, 0x2a, 0x0a, 0x0a,
	0x66, 0x5f, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x46, 0x5f, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x52, 0x09, 0x66,
	0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x66, 0x5f, 0x72, 0x65,
	0x6c, 0x6f, 0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x46, 0x5f, 0x52, 0x65,
	0x6c, 0x6f, 0x6a, 0x52, 0x06, 0x66, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x22, 0x25, 0x0a, 0x09, 0x46,
	0x5f, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x61, 0x22, 0x22, 0x0a, 0x08, 0x46, 0x5f, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x22, 0x19, 0x0a, 0x05, 0x46, 0x5f, 0x4c, 0x6f, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f,
	0x67, 0x22, 0x28, 0x0a, 0x0a, 0x46, 0x5f, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x22, 0x33, 0x0a, 0x07, 0x46,
	0x5f, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x5a,
	0x22, 0x25, 0x0a, 0x09, 0x46, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x26, 0x0a, 0x08, 0x46, 0x61, 0x6e, 0x74, 0x61,
	0x73, 0x6d, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x6d, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x6d, 0x61, 0x32,
	0xc1, 0x01, 0x0a, 0x07, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x2d, 0x0a, 0x11, 0x46,
	0x5f, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x0b, 0x2e, 0x46, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x65, 0x69, 0x61, 0x1a, 0x09, 0x2e,
	0x46, 0x5f, 0x54, 0x6f, 0x4c, 0x65, 0x69, 0x61, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x46, 0x5f,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x2e, 0x46, 0x5f,
	0x46, 0x72, 0x6f, 0x6d, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x1a,
	0x10, 0x2e, 0x46, 0x5f, 0x54, 0x6f, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74,
	0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x09, 0x46, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x09, 0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x6d, 0x61, 0x1a, 0x0d, 0x2e, 0x46, 0x5f,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07,
	0x46, 0x5f, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x46, 0x5f, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x09, 0x2e, 0x46, 0x61, 0x6e, 0x74, 0x61, 0x73, 0x6d,
	0x61, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fulcrum_proto_rawDescOnce sync.Once
	file_fulcrum_proto_rawDescData = file_fulcrum_proto_rawDesc
)

func file_fulcrum_proto_rawDescGZIP() []byte {
	file_fulcrum_proto_rawDescOnce.Do(func() {
		file_fulcrum_proto_rawDescData = protoimpl.X.CompressGZIP(file_fulcrum_proto_rawDescData)
	})
	return file_fulcrum_proto_rawDescData
}

var file_fulcrum_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_fulcrum_proto_goTypes = []interface{}{
	(*F_FromLeia)(nil),        // 0: F_FromLeia
	(*F_From_Informante)(nil), // 1: F_From_Informante
	(*F_To_Informante)(nil),   // 2: F_To_Informante
	(*F_Merge_Data)(nil),      // 3: F_Merge_Data
	(*F_ToLeia)(nil),          // 4: F_ToLeia
	(*F_Planeta)(nil),         // 5: F_Planeta
	(*F_Ciudad)(nil),          // 6: F_Ciudad
	(*F_Log)(nil),             // 7: F_Log
	(*F_Rebeldes)(nil),        // 8: F_Rebeldes
	(*F_Reloj)(nil),           // 9: F_Reloj
	(*F_Command)(nil),         // 10: F_Command
	(*Fantasma)(nil),          // 11: Fantasma
}
var file_fulcrum_proto_depIdxs = []int32{
	5,  // 0: F_FromLeia.f_planeta:type_name -> F_Planeta
	6,  // 1: F_FromLeia.f_ciudad:type_name -> F_Ciudad
	10, // 2: F_From_Informante.f_command:type_name -> F_Command
	9,  // 3: F_From_Informante.f_reloj:type_name -> F_Reloj
	9,  // 4: F_To_Informante.f_reloj:type_name -> F_Reloj
	7,  // 5: F_To_Informante.f_log:type_name -> F_Log
	9,  // 6: F_Merge_Data.f_reloj:type_name -> F_Reloj
	7,  // 7: F_Merge_Data.f_log:type_name -> F_Log
	8,  // 8: F_ToLeia.f_rebeldes:type_name -> F_Rebeldes
	9,  // 9: F_ToLeia.f_reloj:type_name -> F_Reloj
	0,  // 10: Fulcrum.F_GetNumberRebels:input_type -> F_FromLeia
	1,  // 11: Fulcrum.F_SendCommand:input_type -> F_From_Informante
	11, // 12: Fulcrum.F_Request:input_type -> Fantasma
	3,  // 13: Fulcrum.F_Merge:input_type -> F_Merge_Data
	4,  // 14: Fulcrum.F_GetNumberRebels:output_type -> F_ToLeia
	2,  // 15: Fulcrum.F_SendCommand:output_type -> F_To_Informante
	3,  // 16: Fulcrum.F_Request:output_type -> F_Merge_Data
	11, // 17: Fulcrum.F_Merge:output_type -> Fantasma
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_fulcrum_proto_init() }
func file_fulcrum_proto_init() {
	if File_fulcrum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fulcrum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_FromLeia); i {
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
		file_fulcrum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_From_Informante); i {
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
		file_fulcrum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_To_Informante); i {
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
		file_fulcrum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_Merge_Data); i {
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
		file_fulcrum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_ToLeia); i {
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
		file_fulcrum_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_Planeta); i {
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
		file_fulcrum_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_Ciudad); i {
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
		file_fulcrum_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_Log); i {
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
		file_fulcrum_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_Rebeldes); i {
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
		file_fulcrum_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_Reloj); i {
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
		file_fulcrum_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*F_Command); i {
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
		file_fulcrum_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fantasma); i {
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
			RawDescriptor: file_fulcrum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fulcrum_proto_goTypes,
		DependencyIndexes: file_fulcrum_proto_depIdxs,
		MessageInfos:      file_fulcrum_proto_msgTypes,
	}.Build()
	File_fulcrum_proto = out.File
	file_fulcrum_proto_rawDesc = nil
	file_fulcrum_proto_goTypes = nil
	file_fulcrum_proto_depIdxs = nil
}
