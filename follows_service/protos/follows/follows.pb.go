// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protos/follows.proto

package follows

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

type Follow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowID string `protobuf:"bytes,1,opt,name=FollowID,proto3" json:"FollowID,omitempty" bson:"_id,omitempty"`
	Follower string `protobuf:"bytes,2,opt,name=Follower,proto3" json:"Follower,omitempty"`
	Followee string `protobuf:"bytes,3,opt,name=Followee,proto3" json:"Followee,omitempty"`
}

func (x *Follow) Reset() {
	*x = Follow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_follows_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Follow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Follow) ProtoMessage() {}

func (x *Follow) ProtoReflect() protoreflect.Message {
	mi := &file_protos_follows_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Follow.ProtoReflect.Descriptor instead.
func (*Follow) Descriptor() ([]byte, []int) {
	return file_protos_follows_proto_rawDescGZIP(), []int{0}
}

func (x *Follow) GetFollowID() string {
	if x != nil {
		return x.FollowID
	}
	return ""
}

func (x *Follow) GetFollower() string {
	if x != nil {
		return x.Follower
	}
	return ""
}

func (x *Follow) GetFollowee() string {
	if x != nil {
		return x.Followee
	}
	return ""
}

type GetFollowRRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *GetFollowRRequest) Reset() {
	*x = GetFollowRRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_follows_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowRRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowRRequest) ProtoMessage() {}

func (x *GetFollowRRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_follows_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowRRequest.ProtoReflect.Descriptor instead.
func (*GetFollowRRequest) Descriptor() ([]byte, []int) {
	return file_protos_follows_proto_rawDescGZIP(), []int{1}
}

func (x *GetFollowRRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestID string `protobuf:"bytes,1,opt,name=RequestID,proto3" json:"RequestID,omitempty" bson:"_id,omitempty"`
	Requester string `protobuf:"bytes,2,opt,name=Requester,proto3" json:"Requester,omitempty"`
	Requestee string `protobuf:"bytes,3,opt,name=Requestee,proto3" json:"Requestee,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_follows_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_follows_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_protos_follows_proto_rawDescGZIP(), []int{2}
}

func (x *FollowRequest) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *FollowRequest) GetRequester() string {
	if x != nil {
		return x.Requester
	}
	return ""
}

func (x *FollowRequest) GetRequestee() string {
	if x != nil {
		return x.Requestee
	}
	return ""
}

type FollowRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*FollowRequest `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *FollowRequests) Reset() {
	*x = FollowRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_follows_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequests) ProtoMessage() {}

func (x *FollowRequests) ProtoReflect() protoreflect.Message {
	mi := &file_protos_follows_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequests.ProtoReflect.Descriptor instead.
func (*FollowRequests) Descriptor() ([]byte, []int) {
	return file_protos_follows_proto_rawDescGZIP(), []int{3}
}

func (x *FollowRequests) GetResults() []*FollowRequest {
	if x != nil {
		return x.Results
	}
	return nil
}

type DeleteFollowRRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowRequestID string `protobuf:"bytes,1,opt,name=FollowRequestID,proto3" json:"FollowRequestID,omitempty"`
}

func (x *DeleteFollowRRequest) Reset() {
	*x = DeleteFollowRRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_follows_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFollowRRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFollowRRequest) ProtoMessage() {}

func (x *DeleteFollowRRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_follows_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFollowRRequest.ProtoReflect.Descriptor instead.
func (*DeleteFollowRRequest) Descriptor() ([]byte, []int) {
	return file_protos_follows_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFollowRRequest) GetFollowRequestID() string {
	if x != nil {
		return x.FollowRequestID
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_follows_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_follows_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_protos_follows_proto_rawDescGZIP(), []int{5}
}

var File_protos_follows_proto protoreflect.FileDescriptor

var file_protos_follows_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x65,
	0x22, 0x3a, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x22, 0x0f,
	0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xb6, 0x02, 0x0a, 0x07, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x6f, 0x44, 0x42, 0x12, 0x07, 0x2e, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x1a, 0x0e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x44, 0x42, 0x12, 0x0e, 0x2e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x12, 0x07, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x1a, 0x07, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x32, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_follows_proto_rawDescOnce sync.Once
	file_protos_follows_proto_rawDescData = file_protos_follows_proto_rawDesc
)

func file_protos_follows_proto_rawDescGZIP() []byte {
	file_protos_follows_proto_rawDescOnce.Do(func() {
		file_protos_follows_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_follows_proto_rawDescData)
	})
	return file_protos_follows_proto_rawDescData
}

var file_protos_follows_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_follows_proto_goTypes = []interface{}{
	(*Follow)(nil),               // 0: Follow
	(*GetFollowRRequest)(nil),    // 1: GetFollowRRequest
	(*FollowRequest)(nil),        // 2: FollowRequest
	(*FollowRequests)(nil),       // 3: FollowRequests
	(*DeleteFollowRRequest)(nil), // 4: DeleteFollowRRequest
	(*EmptyResponse)(nil),        // 5: EmptyResponse
}
var file_protos_follows_proto_depIdxs = []int32{
	2, // 0: FollowRequests.results:type_name -> FollowRequest
	0, // 1: Follows.AddFollowToDB:input_type -> Follow
	2, // 2: Follows.AddFollowRequestToDB:input_type -> FollowRequest
	1, // 3: Follows.GetFollowRequests:input_type -> GetFollowRRequest
	4, // 4: Follows.DeleteFollowRequest:input_type -> DeleteFollowRRequest
	0, // 5: Follows.GetFollow:input_type -> Follow
	2, // 6: Follows.GetFollowRequest:input_type -> FollowRequest
	5, // 7: Follows.AddFollowToDB:output_type -> EmptyResponse
	5, // 8: Follows.AddFollowRequestToDB:output_type -> EmptyResponse
	3, // 9: Follows.GetFollowRequests:output_type -> FollowRequests
	5, // 10: Follows.DeleteFollowRequest:output_type -> EmptyResponse
	0, // 11: Follows.GetFollow:output_type -> Follow
	2, // 12: Follows.GetFollowRequest:output_type -> FollowRequest
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_follows_proto_init() }
func file_protos_follows_proto_init() {
	if File_protos_follows_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_follows_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Follow); i {
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
		file_protos_follows_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowRRequest); i {
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
		file_protos_follows_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_protos_follows_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequests); i {
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
		file_protos_follows_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFollowRRequest); i {
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
		file_protos_follows_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_protos_follows_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_follows_proto_goTypes,
		DependencyIndexes: file_protos_follows_proto_depIdxs,
		MessageInfos:      file_protos_follows_proto_msgTypes,
	}.Build()
	File_protos_follows_proto = out.File
	file_protos_follows_proto_rawDesc = nil
	file_protos_follows_proto_goTypes = nil
	file_protos_follows_proto_depIdxs = nil
}