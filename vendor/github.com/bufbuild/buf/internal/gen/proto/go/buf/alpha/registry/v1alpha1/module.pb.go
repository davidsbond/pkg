// Copyright 2020-2021 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: buf/alpha/registry/v1alpha1/module.proto

package registryv1alpha1

import (
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

// LocalModuleReference is a local module reference.
//
// It does not include a remote.
type LocalModuleReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner      string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	// Types that are assignable to Reference:
	//	*LocalModuleReference_Branch
	//	*LocalModuleReference_Commit
	Reference isLocalModuleReference_Reference `protobuf_oneof:"reference"`
}

func (x *LocalModuleReference) Reset() {
	*x = LocalModuleReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalModuleReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalModuleReference) ProtoMessage() {}

func (x *LocalModuleReference) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalModuleReference.ProtoReflect.Descriptor instead.
func (*LocalModuleReference) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_module_proto_rawDescGZIP(), []int{0}
}

func (x *LocalModuleReference) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *LocalModuleReference) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (m *LocalModuleReference) GetReference() isLocalModuleReference_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (x *LocalModuleReference) GetBranch() string {
	if x, ok := x.GetReference().(*LocalModuleReference_Branch); ok {
		return x.Branch
	}
	return ""
}

func (x *LocalModuleReference) GetCommit() string {
	if x, ok := x.GetReference().(*LocalModuleReference_Commit); ok {
		return x.Commit
	}
	return ""
}

type isLocalModuleReference_Reference interface {
	isLocalModuleReference_Reference()
}

type LocalModuleReference_Branch struct {
	Branch string `protobuf:"bytes,3,opt,name=branch,proto3,oneof"`
}

type LocalModuleReference_Commit struct {
	Commit string `protobuf:"bytes,4,opt,name=commit,proto3,oneof"`
}

func (*LocalModuleReference_Branch) isLocalModuleReference_Reference() {}

func (*LocalModuleReference_Commit) isLocalModuleReference_Reference() {}

// LocalModulePin is a local module pin.
//
// It does not include a remote.
type LocalModulePin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner      string               `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repository string               `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch     string               `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	Commit     string               `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	Digest     string               `protobuf:"bytes,5,opt,name=digest,proto3" json:"digest,omitempty"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *LocalModulePin) Reset() {
	*x = LocalModulePin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_registry_v1alpha1_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalModulePin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalModulePin) ProtoMessage() {}

func (x *LocalModulePin) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalModulePin.ProtoReflect.Descriptor instead.
func (*LocalModulePin) Descriptor() ([]byte, []int) {
	return file_buf_alpha_registry_v1alpha1_module_proto_rawDescGZIP(), []int{1}
}

func (x *LocalModulePin) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *LocalModulePin) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *LocalModulePin) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *LocalModulePin) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *LocalModulePin) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *LocalModulePin) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

var File_buf_alpha_registry_v1alpha1_module_proto protoreflect.FileDescriptor

var file_buf_alpha_registry_v1alpha1_module_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x14, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x18, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x5c, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75,
	0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_registry_v1alpha1_module_proto_rawDescOnce sync.Once
	file_buf_alpha_registry_v1alpha1_module_proto_rawDescData = file_buf_alpha_registry_v1alpha1_module_proto_rawDesc
)

func file_buf_alpha_registry_v1alpha1_module_proto_rawDescGZIP() []byte {
	file_buf_alpha_registry_v1alpha1_module_proto_rawDescOnce.Do(func() {
		file_buf_alpha_registry_v1alpha1_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_registry_v1alpha1_module_proto_rawDescData)
	})
	return file_buf_alpha_registry_v1alpha1_module_proto_rawDescData
}

var file_buf_alpha_registry_v1alpha1_module_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_alpha_registry_v1alpha1_module_proto_goTypes = []interface{}{
	(*LocalModuleReference)(nil), // 0: buf.alpha.registry.v1alpha1.LocalModuleReference
	(*LocalModulePin)(nil),       // 1: buf.alpha.registry.v1alpha1.LocalModulePin
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_buf_alpha_registry_v1alpha1_module_proto_depIdxs = []int32{
	2, // 0: buf.alpha.registry.v1alpha1.LocalModulePin.create_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_alpha_registry_v1alpha1_module_proto_init() }
func file_buf_alpha_registry_v1alpha1_module_proto_init() {
	if File_buf_alpha_registry_v1alpha1_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_registry_v1alpha1_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalModuleReference); i {
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
		file_buf_alpha_registry_v1alpha1_module_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalModulePin); i {
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
	file_buf_alpha_registry_v1alpha1_module_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LocalModuleReference_Branch)(nil),
		(*LocalModuleReference_Commit)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_registry_v1alpha1_module_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_registry_v1alpha1_module_proto_goTypes,
		DependencyIndexes: file_buf_alpha_registry_v1alpha1_module_proto_depIdxs,
		MessageInfos:      file_buf_alpha_registry_v1alpha1_module_proto_msgTypes,
	}.Build()
	File_buf_alpha_registry_v1alpha1_module_proto = out.File
	file_buf_alpha_registry_v1alpha1_module_proto_rawDesc = nil
	file_buf_alpha_registry_v1alpha1_module_proto_goTypes = nil
	file_buf_alpha_registry_v1alpha1_module_proto_depIdxs = nil
}
