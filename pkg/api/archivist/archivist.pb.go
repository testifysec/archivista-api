// Copyright 2022 The Archivist Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: archivist.proto

package archivist

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

type GetBySubjectDigestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm      string `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Value          string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	CollectionName string `protobuf:"bytes,3,opt,name=collectionName,proto3" json:"collectionName,omitempty"`
}

func (x *GetBySubjectDigestRequest) Reset() {
	*x = GetBySubjectDigestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archivist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBySubjectDigestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBySubjectDigestRequest) ProtoMessage() {}

func (x *GetBySubjectDigestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_archivist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBySubjectDigestRequest.ProtoReflect.Descriptor instead.
func (*GetBySubjectDigestRequest) Descriptor() ([]byte, []int) {
	return file_archivist_proto_rawDescGZIP(), []int{0}
}

func (x *GetBySubjectDigestRequest) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *GetBySubjectDigestRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetBySubjectDigestRequest) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

type GetBySubjectDigestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gitoid         string   `protobuf:"bytes,1,opt,name=gitoid,proto3" json:"gitoid,omitempty"`
	CollectionName string   `protobuf:"bytes,2,opt,name=collectionName,proto3" json:"collectionName,omitempty"`
	Attestations   []string `protobuf:"bytes,3,rep,name=attestations,proto3" json:"attestations,omitempty"`
}

func (x *GetBySubjectDigestResponse) Reset() {
	*x = GetBySubjectDigestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archivist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBySubjectDigestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBySubjectDigestResponse) ProtoMessage() {}

func (x *GetBySubjectDigestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_archivist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBySubjectDigestResponse.ProtoReflect.Descriptor instead.
func (*GetBySubjectDigestResponse) Descriptor() ([]byte, []int) {
	return file_archivist_proto_rawDescGZIP(), []int{1}
}

func (x *GetBySubjectDigestResponse) GetGitoid() string {
	if x != nil {
		return x.Gitoid
	}
	return ""
}

func (x *GetBySubjectDigestResponse) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *GetBySubjectDigestResponse) GetAttestations() []string {
	if x != nil {
		return x.Attestations
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archivist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_archivist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_archivist_proto_rawDescGZIP(), []int{2}
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gitoid string `protobuf:"bytes,1,opt,name=gitoid,proto3" json:"gitoid,omitempty"`
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archivist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_archivist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_archivist_proto_rawDescGZIP(), []int{3}
}

func (x *StoreResponse) GetGitoid() string {
	if x != nil {
		return x.Gitoid
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gitoid string `protobuf:"bytes,1,opt,name=gitoid,proto3" json:"gitoid,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archivist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_archivist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_archivist_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetGitoid() string {
	if x != nil {
		return x.Gitoid
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gitoid string `protobuf:"bytes,1,opt,name=gitoid,proto3" json:"gitoid,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_archivist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_archivist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_archivist_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetGitoid() string {
	if x != nil {
		return x.Gitoid
	}
	return ""
}

func (x *GetResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

var File_archivist_proto protoreflect.FileDescriptor

var file_archivist_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x6f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x6f, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1d, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x6f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x6f, 0x69, 0x64,
	0x22, 0x24, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x69, 0x74, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x69, 0x74, 0x6f, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x6f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x6f, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x32, 0x70, 0x0a, 0x09, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69,
	0x73, 0x74, 0x12, 0x63, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x32, 0x74, 0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x10, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a,
	0x18, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x30, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x30, 0x01, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x66, 0x79, 0x73, 0x65, 0x63, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_archivist_proto_rawDescOnce sync.Once
	file_archivist_proto_rawDescData = file_archivist_proto_rawDesc
)

func file_archivist_proto_rawDescGZIP() []byte {
	file_archivist_proto_rawDescOnce.Do(func() {
		file_archivist_proto_rawDescData = protoimpl.X.CompressGZIP(file_archivist_proto_rawDescData)
	})
	return file_archivist_proto_rawDescData
}

var file_archivist_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_archivist_proto_goTypes = []interface{}{
	(*GetBySubjectDigestRequest)(nil),  // 0: archivist.GetBySubjectDigestRequest
	(*GetBySubjectDigestResponse)(nil), // 1: archivist.GetBySubjectDigestResponse
	(*Chunk)(nil),                      // 2: archivist.Chunk
	(*StoreResponse)(nil),              // 3: archivist.StoreResponse
	(*GetRequest)(nil),                 // 4: archivist.GetRequest
	(*GetResponse)(nil),                // 5: archivist.GetResponse
}
var file_archivist_proto_depIdxs = []int32{
	0, // 0: archivist.Archivist.GetBySubjectDigest:input_type -> archivist.GetBySubjectDigestRequest
	2, // 1: archivist.Collector.Store:input_type -> archivist.Chunk
	4, // 2: archivist.Collector.Get:input_type -> archivist.GetRequest
	1, // 3: archivist.Archivist.GetBySubjectDigest:output_type -> archivist.GetBySubjectDigestResponse
	3, // 4: archivist.Collector.Store:output_type -> archivist.StoreResponse
	2, // 5: archivist.Collector.Get:output_type -> archivist.Chunk
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_archivist_proto_init() }
func file_archivist_proto_init() {
	if File_archivist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_archivist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBySubjectDigestRequest); i {
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
		file_archivist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBySubjectDigestResponse); i {
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
		file_archivist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_archivist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResponse); i {
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
		file_archivist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_archivist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
			RawDescriptor: file_archivist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_archivist_proto_goTypes,
		DependencyIndexes: file_archivist_proto_depIdxs,
		MessageInfos:      file_archivist_proto_msgTypes,
	}.Build()
	File_archivist_proto = out.File
	file_archivist_proto_rawDesc = nil
	file_archivist_proto_goTypes = nil
	file_archivist_proto_depIdxs = nil
}
