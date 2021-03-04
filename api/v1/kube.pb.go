// 定义项目 API 的 proto 文件 可以同时描述 gRPC 和 HTTP API
// protobuf 文件参考:
//  - https://developers.google.com/protocol-buffers/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: kube.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/duration"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                       string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GenerateName               string                `protobuf:"bytes,2,opt,name=generateName,proto3" json:"generateName,omitempty"`
	Namespace                  string                `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	SelfLink                   string                `protobuf:"bytes,4,opt,name=selfLink,proto3" json:"selfLink,omitempty"`
	Uid                        string                `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty"`
	ResourceVersion            string                `protobuf:"bytes,6,opt,name=resourceVersion,proto3" json:"resourceVersion,omitempty"`
	Generation                 int64                 `protobuf:"varint,7,opt,name=generation,proto3" json:"generation,omitempty"`
	CreationTimestamp          string                `protobuf:"bytes,8,opt,name=creationTimestamp,proto3" json:"creationTimestamp,omitempty"`
	DeletionTimestamp          string                `protobuf:"bytes,9,opt,name=deletionTimestamp,proto3" json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSeconds int64                 `protobuf:"varint,10,opt,name=deletionGracePeriodSeconds,proto3" json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     map[string]string     `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations                map[string]string     `protobuf:"bytes,12,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OwnerReferences            []*OwnerReference     `protobuf:"bytes,13,rep,name=ownerReferences,proto3" json:"ownerReferences,omitempty"`
	Finalizers                 []string              `protobuf:"bytes,14,rep,name=finalizers,proto3" json:"finalizers,omitempty"`
	ClusterName                string                `protobuf:"bytes,15,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	ManagedFields              []*ManagedFieldsEntry `protobuf:"bytes,17,rep,name=managedFields,proto3" json:"managedFields,omitempty"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_kube_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_kube_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectMeta) GetGenerateName() string {
	if x != nil {
		return x.GenerateName
	}
	return ""
}

func (x *ObjectMeta) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ObjectMeta) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

func (x *ObjectMeta) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ObjectMeta) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *ObjectMeta) GetGeneration() int64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *ObjectMeta) GetCreationTimestamp() string {
	if x != nil {
		return x.CreationTimestamp
	}
	return ""
}

func (x *ObjectMeta) GetDeletionTimestamp() string {
	if x != nil {
		return x.DeletionTimestamp
	}
	return ""
}

func (x *ObjectMeta) GetDeletionGracePeriodSeconds() int64 {
	if x != nil {
		return x.DeletionGracePeriodSeconds
	}
	return 0
}

func (x *ObjectMeta) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ObjectMeta) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *ObjectMeta) GetOwnerReferences() []*OwnerReference {
	if x != nil {
		return x.OwnerReferences
	}
	return nil
}

func (x *ObjectMeta) GetFinalizers() []string {
	if x != nil {
		return x.Finalizers
	}
	return nil
}

func (x *ObjectMeta) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ObjectMeta) GetManagedFields() []*ManagedFieldsEntry {
	if x != nil {
		return x.ManagedFields
	}
	return nil
}

type OwnerReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion         string `protobuf:"bytes,5,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind               string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Name               string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Uid                string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Controller         bool   `protobuf:"varint,6,opt,name=controller,proto3" json:"controller,omitempty"`
	BlockOwnerDeletion bool   `protobuf:"varint,7,opt,name=blockOwnerDeletion,proto3" json:"blockOwnerDeletion,omitempty"`
}

func (x *OwnerReference) Reset() {
	*x = OwnerReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerReference) ProtoMessage() {}

func (x *OwnerReference) ProtoReflect() protoreflect.Message {
	mi := &file_kube_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerReference.ProtoReflect.Descriptor instead.
func (*OwnerReference) Descriptor() ([]byte, []int) {
	return file_kube_proto_rawDescGZIP(), []int{1}
}

func (x *OwnerReference) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *OwnerReference) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *OwnerReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OwnerReference) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *OwnerReference) GetController() bool {
	if x != nil {
		return x.Controller
	}
	return false
}

func (x *OwnerReference) GetBlockOwnerDeletion() bool {
	if x != nil {
		return x.BlockOwnerDeletion
	}
	return false
}

type FieldsV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw []byte `protobuf:"bytes,1,opt,name=Raw,proto3" json:"Raw,omitempty"`
}

func (x *FieldsV1) Reset() {
	*x = FieldsV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldsV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldsV1) ProtoMessage() {}

func (x *FieldsV1) ProtoReflect() protoreflect.Message {
	mi := &file_kube_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldsV1.ProtoReflect.Descriptor instead.
func (*FieldsV1) Descriptor() ([]byte, []int) {
	return file_kube_proto_rawDescGZIP(), []int{2}
}

func (x *FieldsV1) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

type ManagedFieldsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manager    string    `protobuf:"bytes,1,opt,name=manager,proto3" json:"manager,omitempty"`
	Operation  string    `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	ApiVersion string    `protobuf:"bytes,3,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Time       string    `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	FieldsType string    `protobuf:"bytes,6,opt,name=fieldsType,proto3" json:"fieldsType,omitempty"`
	FieldsV1   *FieldsV1 `protobuf:"bytes,7,opt,name=fieldsV1,proto3" json:"fieldsV1,omitempty"`
}

func (x *ManagedFieldsEntry) Reset() {
	*x = ManagedFieldsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagedFieldsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagedFieldsEntry) ProtoMessage() {}

func (x *ManagedFieldsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_kube_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagedFieldsEntry.ProtoReflect.Descriptor instead.
func (*ManagedFieldsEntry) Descriptor() ([]byte, []int) {
	return file_kube_proto_rawDescGZIP(), []int{3}
}

func (x *ManagedFieldsEntry) GetManager() string {
	if x != nil {
		return x.Manager
	}
	return ""
}

func (x *ManagedFieldsEntry) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *ManagedFieldsEntry) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ManagedFieldsEntry) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ManagedFieldsEntry) GetFieldsType() string {
	if x != nil {
		return x.FieldsType
	}
	return ""
}

func (x *ManagedFieldsEntry) GetFieldsV1() *FieldsV1 {
	if x != nil {
		return x.FieldsV1
	}
	return nil
}

var File_kube_proto protoreflect.FileDescriptor

var file_kube_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x06, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x3e, 0x0a, 0x1a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x63, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x1a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61,
	0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x3d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4c,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x0f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xba, 0x01, 0x0a, 0x0e,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1c, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x56, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x52, 0x61, 0x77, 0x22, 0xd5, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x56, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x56, 0x31, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x56, 0x31, 0x42, 0x4a,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4b, 0x75, 0x62, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6f, 0x74, 0x7a, 0x2f, 0x63, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kube_proto_rawDescOnce sync.Once
	file_kube_proto_rawDescData = file_kube_proto_rawDesc
)

func file_kube_proto_rawDescGZIP() []byte {
	file_kube_proto_rawDescOnce.Do(func() {
		file_kube_proto_rawDescData = protoimpl.X.CompressGZIP(file_kube_proto_rawDescData)
	})
	return file_kube_proto_rawDescData
}

var file_kube_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kube_proto_goTypes = []interface{}{
	(*ObjectMeta)(nil),         // 0: cratos.api.v1.ObjectMeta
	(*OwnerReference)(nil),     // 1: cratos.api.v1.OwnerReference
	(*FieldsV1)(nil),           // 2: cratos.api.v1.FieldsV1
	(*ManagedFieldsEntry)(nil), // 3: cratos.api.v1.ManagedFieldsEntry
	nil,                        // 4: cratos.api.v1.ObjectMeta.LabelsEntry
	nil,                        // 5: cratos.api.v1.ObjectMeta.AnnotationsEntry
}
var file_kube_proto_depIdxs = []int32{
	4, // 0: cratos.api.v1.ObjectMeta.labels:type_name -> cratos.api.v1.ObjectMeta.LabelsEntry
	5, // 1: cratos.api.v1.ObjectMeta.annotations:type_name -> cratos.api.v1.ObjectMeta.AnnotationsEntry
	1, // 2: cratos.api.v1.ObjectMeta.ownerReferences:type_name -> cratos.api.v1.OwnerReference
	3, // 3: cratos.api.v1.ObjectMeta.managedFields:type_name -> cratos.api.v1.ManagedFieldsEntry
	2, // 4: cratos.api.v1.ManagedFieldsEntry.fieldsV1:type_name -> cratos.api.v1.FieldsV1
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kube_proto_init() }
func file_kube_proto_init() {
	if File_kube_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kube_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
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
		file_kube_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerReference); i {
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
		file_kube_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldsV1); i {
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
		file_kube_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagedFieldsEntry); i {
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
			RawDescriptor: file_kube_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kube_proto_goTypes,
		DependencyIndexes: file_kube_proto_depIdxs,
		MessageInfos:      file_kube_proto_msgTypes,
	}.Build()
	File_kube_proto = out.File
	file_kube_proto_rawDesc = nil
	file_kube_proto_goTypes = nil
	file_kube_proto_depIdxs = nil
}
