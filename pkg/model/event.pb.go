// Copyright 2022 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/model/event.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type EventStatus int32

const (
	EventStatus_EVENT_NOT_HANDLED EventStatus = 0
	EventStatus_EVENT_SUCCESS     EventStatus = 1
	EventStatus_EVENT_FAILURE     EventStatus = 2
	EventStatus_EVENT_OUTDATED    EventStatus = 3
)

// Enum value maps for EventStatus.
var (
	EventStatus_name = map[int32]string{
		0: "EVENT_NOT_HANDLED",
		1: "EVENT_SUCCESS",
		2: "EVENT_FAILURE",
		3: "EVENT_OUTDATED",
	}
	EventStatus_value = map[string]int32{
		"EVENT_NOT_HANDLED": 0,
		"EVENT_SUCCESS":     1,
		"EVENT_FAILURE":     2,
		"EVENT_OUTDATED":    3,
	}
)

func (x EventStatus) Enum() *EventStatus {
	p := new(EventStatus)
	*p = x
	return p
}

func (x EventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_model_event_proto_enumTypes[0].Descriptor()
}

func (EventStatus) Type() protoreflect.EnumType {
	return &file_pkg_model_event_proto_enumTypes[0]
}

func (x EventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventStatus.Descriptor instead.
func (EventStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_model_event_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The generated unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of event.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The data of event.
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// The ID of the project this event belongs to.
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The key/value pairs that are attached to event.
	// This is intended to be used to specify additional attributes of event.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A fixed-length identifier consists of its own name and labels.
	EventKey string `protobuf:"bytes,6,opt,name=event_key,json=eventKey,proto3" json:"event_key,omitempty"`
	// The handle status of event.
	Status            EventStatus `protobuf:"varint,8,opt,name=status,proto3,enum=model.EventStatus" json:"status,omitempty"`
	StatusDescription string      `protobuf:"bytes,9,opt,name=status_description,json=statusDescription,proto3" json:"status_description,omitempty"`
	// Unix time when the event was handled.
	HandledAt int64 `protobuf:"varint,13,opt,name=handled_at,json=handledAt,proto3" json:"handled_at,omitempty"`
	// Unix time when the event was created.
	CreatedAt int64 `protobuf:"varint,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Unix time of the last time when the event was updated.
	UpdatedAt int64 `protobuf:"varint,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_pkg_model_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Event) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Event) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Event) GetEventKey() string {
	if x != nil {
		return x.EventKey
	}
	return ""
}

func (x *Event) GetStatus() EventStatus {
	if x != nil {
		return x.Status
	}
	return EventStatus_EVENT_NOT_HANDLED
}

func (x *Event) GetStatusDescription() string {
	if x != nil {
		return x.StatusDescription
	}
	return ""
}

func (x *Event) GetHandledAt() int64 {
	if x != nil {
		return x.HandledAt
	}
	return 0
}

func (x *Event) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Event) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_pkg_model_event_proto protoreflect.FileDescriptor

var file_pkg_model_event_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x03, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x24,
	0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x26, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x2a, 0x5e, 0x0a, 0x0b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4f,
	0x55, 0x54, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x64, 0x2f,
	0x70, 0x69, 0x70, 0x65, 0x63, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_model_event_proto_rawDescOnce sync.Once
	file_pkg_model_event_proto_rawDescData = file_pkg_model_event_proto_rawDesc
)

func file_pkg_model_event_proto_rawDescGZIP() []byte {
	file_pkg_model_event_proto_rawDescOnce.Do(func() {
		file_pkg_model_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_model_event_proto_rawDescData)
	})
	return file_pkg_model_event_proto_rawDescData
}

var file_pkg_model_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_model_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_model_event_proto_goTypes = []interface{}{
	(EventStatus)(0), // 0: model.EventStatus
	(*Event)(nil),    // 1: model.Event
	nil,              // 2: model.Event.LabelsEntry
}
var file_pkg_model_event_proto_depIdxs = []int32{
	2, // 0: model.Event.labels:type_name -> model.Event.LabelsEntry
	0, // 1: model.Event.status:type_name -> model.EventStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_model_event_proto_init() }
func file_pkg_model_event_proto_init() {
	if File_pkg_model_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_model_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_pkg_model_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_model_event_proto_goTypes,
		DependencyIndexes: file_pkg_model_event_proto_depIdxs,
		EnumInfos:         file_pkg_model_event_proto_enumTypes,
		MessageInfos:      file_pkg_model_event_proto_msgTypes,
	}.Build()
	File_pkg_model_event_proto = out.File
	file_pkg_model_event_proto_rawDesc = nil
	file_pkg_model_event_proto_goTypes = nil
	file_pkg_model_event_proto_depIdxs = nil
}
