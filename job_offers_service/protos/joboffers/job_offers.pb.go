// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protos/job_offers.proto

package job_offers

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

type JobOffersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JobOffersRequest) Reset() {
	*x = JobOffersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_job_offers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffersRequest) ProtoMessage() {}

func (x *JobOffersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_job_offers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffersRequest.ProtoReflect.Descriptor instead.
func (*JobOffersRequest) Descriptor() ([]byte, []int) {
	return file_protos_job_offers_proto_rawDescGZIP(), []int{0}
}

type JobOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Position     string `protobuf:"bytes,2,opt,name=Position,proto3" json:"Position,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Requirements string `protobuf:"bytes,4,opt,name=Requirements,proto3" json:"Requirements,omitempty"`
}

func (x *JobOffer) Reset() {
	*x = JobOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_job_offers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffer) ProtoMessage() {}

func (x *JobOffer) ProtoReflect() protoreflect.Message {
	mi := &file_protos_job_offers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffer.ProtoReflect.Descriptor instead.
func (*JobOffer) Descriptor() ([]byte, []int) {
	return file_protos_job_offers_proto_rawDescGZIP(), []int{1}
}

func (x *JobOffer) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *JobOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *JobOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *JobOffer) GetRequirements() string {
	if x != nil {
		return x.Requirements
	}
	return ""
}

type JobOffersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*JobOffer `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *JobOffersResponse) Reset() {
	*x = JobOffersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_job_offers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobOffersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOffersResponse) ProtoMessage() {}

func (x *JobOffersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_job_offers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOffersResponse.ProtoReflect.Descriptor instead.
func (*JobOffersResponse) Descriptor() ([]byte, []int) {
	return file_protos_job_offers_proto_rawDescGZIP(), []int{2}
}

func (x *JobOffersResponse) GetResults() []*JobOffer {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_protos_job_offers_proto protoreflect.FileDescriptor

var file_protos_job_offers_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x4a, 0x6f, 0x62,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7c, 0x0a,
	0x08, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x4a,
	0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0xa3, 0x01, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x1a, 0x12, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x09, 0x2e,
	0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x4a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x2e,
	0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_job_offers_proto_rawDescOnce sync.Once
	file_protos_job_offers_proto_rawDescData = file_protos_job_offers_proto_rawDesc
)

func file_protos_job_offers_proto_rawDescGZIP() []byte {
	file_protos_job_offers_proto_rawDescOnce.Do(func() {
		file_protos_job_offers_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_job_offers_proto_rawDescData)
	})
	return file_protos_job_offers_proto_rawDescData
}

var file_protos_job_offers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_job_offers_proto_goTypes = []interface{}{
	(*JobOffersRequest)(nil),  // 0: JobOffersRequest
	(*JobOffer)(nil),          // 1: JobOffer
	(*JobOffersResponse)(nil), // 2: JobOffersResponse
}
var file_protos_job_offers_proto_depIdxs = []int32{
	1, // 0: JobOffersResponse.results:type_name -> JobOffer
	1, // 1: JobOffers.CreateJobOffer:input_type -> JobOffer
	0, // 2: JobOffers.GetJobOffers:input_type -> JobOffersRequest
	1, // 3: JobOffers.RemoveJobOffer:input_type -> JobOffer
	2, // 4: JobOffers.CreateJobOffer:output_type -> JobOffersResponse
	2, // 5: JobOffers.GetJobOffers:output_type -> JobOffersResponse
	0, // 6: JobOffers.RemoveJobOffer:output_type -> JobOffersRequest
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_job_offers_proto_init() }
func file_protos_job_offers_proto_init() {
	if File_protos_job_offers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_job_offers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOffersRequest); i {
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
		file_protos_job_offers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOffer); i {
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
		file_protos_job_offers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobOffersResponse); i {
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
			RawDescriptor: file_protos_job_offers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_job_offers_proto_goTypes,
		DependencyIndexes: file_protos_job_offers_proto_depIdxs,
		MessageInfos:      file_protos_job_offers_proto_msgTypes,
	}.Build()
	File_protos_job_offers_proto = out.File
	file_protos_job_offers_proto_rawDesc = nil
	file_protos_job_offers_proto_goTypes = nil
	file_protos_job_offers_proto_depIdxs = nil
}
