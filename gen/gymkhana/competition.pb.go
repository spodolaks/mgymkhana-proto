// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: gymkhana/competition.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Competition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	LocationId  string                 `protobuf:"bytes,15,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Start       *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=start,proto3" json:"start,omitempty"`
	End         *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=end,proto3" json:"end,omitempty"`
	ParentId    string                 `protobuf:"bytes,30,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Course      string                 `protobuf:"bytes,35,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *Competition) Reset() {
	*x = Competition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Competition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Competition) ProtoMessage() {}

func (x *Competition) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Competition.ProtoReflect.Descriptor instead.
func (*Competition) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{0}
}

func (x *Competition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Competition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Competition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Competition) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Competition) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Competition) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *Competition) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Competition) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

type CompetitionCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Competition *Competition `protobuf:"bytes,1,opt,name=competition,proto3" json:"competition,omitempty"`
}

func (x *CompetitionCreateRequest) Reset() {
	*x = CompetitionCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionCreateRequest) ProtoMessage() {}

func (x *CompetitionCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionCreateRequest.ProtoReflect.Descriptor instead.
func (*CompetitionCreateRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{1}
}

func (x *CompetitionCreateRequest) GetCompetition() *Competition {
	if x != nil {
		return x.Competition
	}
	return nil
}

type CompetitionUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Competition *Competition `protobuf:"bytes,1,opt,name=competition,proto3" json:"competition,omitempty"`
}

func (x *CompetitionUpdateRequest) Reset() {
	*x = CompetitionUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionUpdateRequest) ProtoMessage() {}

func (x *CompetitionUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionUpdateRequest.ProtoReflect.Descriptor instead.
func (*CompetitionUpdateRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{2}
}

func (x *CompetitionUpdateRequest) GetCompetition() *Competition {
	if x != nil {
		return x.Competition
	}
	return nil
}

type CompetitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Competition *Competition `protobuf:"bytes,1,opt,name=competition,proto3" json:"competition,omitempty"`
}

func (x *CompetitionResponse) Reset() {
	*x = CompetitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionResponse) ProtoMessage() {}

func (x *CompetitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionResponse.ProtoReflect.Descriptor instead.
func (*CompetitionResponse) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{3}
}

func (x *CompetitionResponse) GetCompetition() *Competition {
	if x != nil {
		return x.Competition
	}
	return nil
}

type CompetitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CompetitionRequest) Reset() {
	*x = CompetitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionRequest) ProtoMessage() {}

func (x *CompetitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionRequest.ProtoReflect.Descriptor instead.
func (*CompetitionRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{4}
}

func (x *CompetitionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CompetitionStages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CompetitionStages) Reset() {
	*x = CompetitionStages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionStages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionStages) ProtoMessage() {}

func (x *CompetitionStages) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionStages.ProtoReflect.Descriptor instead.
func (*CompetitionStages) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{5}
}

func (x *CompetitionStages) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CompetitionStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Competition *Competition `protobuf:"bytes,1,opt,name=competition,proto3" json:"competition,omitempty"`
}

func (x *CompetitionStage) Reset() {
	*x = CompetitionStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionStage) ProtoMessage() {}

func (x *CompetitionStage) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionStage.ProtoReflect.Descriptor instead.
func (*CompetitionStage) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{6}
}

func (x *CompetitionStage) GetCompetition() *Competition {
	if x != nil {
		return x.Competition
	}
	return nil
}

type CompetitionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *CompetitionListRequest) Reset() {
	*x = CompetitionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_competition_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompetitionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompetitionListRequest) ProtoMessage() {}

func (x *CompetitionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_competition_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompetitionListRequest.ProtoReflect.Descriptor instead.
func (*CompetitionListRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_competition_proto_rawDescGZIP(), []int{7}
}

func (x *CompetitionListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

var File_gymkhana_competition_proto protoreflect.FileDescriptor

var file_gymkhana_competition_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x79,
	0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x79, 0x6d, 0x6b,
	0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x13,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68,
	0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x10,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30,
	0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x32, 0xd2, 0x04, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x65,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e,
	0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x79, 0x6d, 0x6b,
	0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5b, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61,
	0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x61, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x25, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61,
	0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x62, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68,
	0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67,
	0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x6f, 0x64, 0x6f, 0x6c, 0x61, 0x6b, 0x73, 0x2f, 0x6d, 0x67,
	0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gymkhana_competition_proto_rawDescOnce sync.Once
	file_gymkhana_competition_proto_rawDescData = file_gymkhana_competition_proto_rawDesc
)

func file_gymkhana_competition_proto_rawDescGZIP() []byte {
	file_gymkhana_competition_proto_rawDescOnce.Do(func() {
		file_gymkhana_competition_proto_rawDescData = protoimpl.X.CompressGZIP(file_gymkhana_competition_proto_rawDescData)
	})
	return file_gymkhana_competition_proto_rawDescData
}

var file_gymkhana_competition_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gymkhana_competition_proto_goTypes = []interface{}{
	(*Competition)(nil),              // 0: gymkhanaproto.Competition
	(*CompetitionCreateRequest)(nil), // 1: gymkhanaproto.CompetitionCreateRequest
	(*CompetitionUpdateRequest)(nil), // 2: gymkhanaproto.CompetitionUpdateRequest
	(*CompetitionResponse)(nil),      // 3: gymkhanaproto.CompetitionResponse
	(*CompetitionRequest)(nil),       // 4: gymkhanaproto.CompetitionRequest
	(*CompetitionStages)(nil),        // 5: gymkhanaproto.CompetitionStages
	(*CompetitionStage)(nil),         // 6: gymkhanaproto.CompetitionStage
	(*CompetitionListRequest)(nil),   // 7: gymkhanaproto.CompetitionListRequest
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_gymkhana_competition_proto_depIdxs = []int32{
	8,  // 0: gymkhanaproto.Competition.start:type_name -> google.protobuf.Timestamp
	8,  // 1: gymkhanaproto.Competition.end:type_name -> google.protobuf.Timestamp
	0,  // 2: gymkhanaproto.CompetitionCreateRequest.competition:type_name -> gymkhanaproto.Competition
	0,  // 3: gymkhanaproto.CompetitionUpdateRequest.competition:type_name -> gymkhanaproto.Competition
	0,  // 4: gymkhanaproto.CompetitionResponse.competition:type_name -> gymkhanaproto.Competition
	0,  // 5: gymkhanaproto.CompetitionStage.competition:type_name -> gymkhanaproto.Competition
	4,  // 6: gymkhanaproto.CompetitionService.Competition:input_type -> gymkhanaproto.CompetitionRequest
	4,  // 7: gymkhanaproto.CompetitionService.CompetitionStages:input_type -> gymkhanaproto.CompetitionRequest
	7,  // 8: gymkhanaproto.CompetitionService.ListCompetitions:input_type -> gymkhanaproto.CompetitionListRequest
	1,  // 9: gymkhanaproto.CompetitionService.CreateCompetition:input_type -> gymkhanaproto.CompetitionCreateRequest
	2,  // 10: gymkhanaproto.CompetitionService.UpdateCompetition:input_type -> gymkhanaproto.CompetitionUpdateRequest
	4,  // 11: gymkhanaproto.CompetitionService.DeleteCompetition:input_type -> gymkhanaproto.CompetitionRequest
	3,  // 12: gymkhanaproto.CompetitionService.Competition:output_type -> gymkhanaproto.CompetitionResponse
	6,  // 13: gymkhanaproto.CompetitionService.CompetitionStages:output_type -> gymkhanaproto.CompetitionStage
	3,  // 14: gymkhanaproto.CompetitionService.ListCompetitions:output_type -> gymkhanaproto.CompetitionResponse
	3,  // 15: gymkhanaproto.CompetitionService.CreateCompetition:output_type -> gymkhanaproto.CompetitionResponse
	3,  // 16: gymkhanaproto.CompetitionService.UpdateCompetition:output_type -> gymkhanaproto.CompetitionResponse
	3,  // 17: gymkhanaproto.CompetitionService.DeleteCompetition:output_type -> gymkhanaproto.CompetitionResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_gymkhana_competition_proto_init() }
func file_gymkhana_competition_proto_init() {
	if File_gymkhana_competition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gymkhana_competition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Competition); i {
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
		file_gymkhana_competition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionCreateRequest); i {
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
		file_gymkhana_competition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionUpdateRequest); i {
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
		file_gymkhana_competition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionResponse); i {
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
		file_gymkhana_competition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionRequest); i {
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
		file_gymkhana_competition_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionStages); i {
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
		file_gymkhana_competition_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionStage); i {
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
		file_gymkhana_competition_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompetitionListRequest); i {
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
			RawDescriptor: file_gymkhana_competition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gymkhana_competition_proto_goTypes,
		DependencyIndexes: file_gymkhana_competition_proto_depIdxs,
		MessageInfos:      file_gymkhana_competition_proto_msgTypes,
	}.Build()
	File_gymkhana_competition_proto = out.File
	file_gymkhana_competition_proto_rawDesc = nil
	file_gymkhana_competition_proto_goTypes = nil
	file_gymkhana_competition_proto_depIdxs = nil
}
