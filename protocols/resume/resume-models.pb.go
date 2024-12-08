// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.0
// source: resume-models.proto

package resume

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

type Resume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeUuid []byte      `protobuf:"bytes,1,opt,name=resume_uuid,json=resumeUuid,proto3" json:"resume_uuid,omitempty"`
	UserUuid   []byte      `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Name       string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ResumeDesc *ResumeDesc `protobuf:"bytes,4,opt,name=ResumeDesc,proto3" json:"ResumeDesc,omitempty"`
}

func (x *Resume) Reset() {
	*x = Resume{}
	mi := &file_resume_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resume) ProtoMessage() {}

func (x *Resume) ProtoReflect() protoreflect.Message {
	mi := &file_resume_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resume.ProtoReflect.Descriptor instead.
func (*Resume) Descriptor() ([]byte, []int) {
	return file_resume_models_proto_rawDescGZIP(), []int{0}
}

func (x *Resume) GetResumeUuid() []byte {
	if x != nil {
		return x.ResumeUuid
	}
	return nil
}

func (x *Resume) GetUserUuid() []byte {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

func (x *Resume) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resume) GetResumeDesc() *ResumeDesc {
	if x != nil {
		return x.ResumeDesc
	}
	return nil
}

type ResumeDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobExps   []*JobExp  `protobuf:"bytes,1,rep,name=JobExps,proto3" json:"JobExps,omitempty"`
	Candidate *Candidate `protobuf:"bytes,2,opt,name=Candidate,proto3" json:"Candidate,omitempty"`
	Skills    []string   `protobuf:"bytes,3,rep,name=skills,proto3" json:"skills,omitempty"`
}

func (x *ResumeDesc) Reset() {
	*x = ResumeDesc{}
	mi := &file_resume_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResumeDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeDesc) ProtoMessage() {}

func (x *ResumeDesc) ProtoReflect() protoreflect.Message {
	mi := &file_resume_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeDesc.ProtoReflect.Descriptor instead.
func (*ResumeDesc) Descriptor() ([]byte, []int) {
	return file_resume_models_proto_rawDescGZIP(), []int{1}
}

func (x *ResumeDesc) GetJobExps() []*JobExp {
	if x != nil {
		return x.JobExps
	}
	return nil
}

func (x *ResumeDesc) GetCandidate() *Candidate {
	if x != nil {
		return x.Candidate
	}
	return nil
}

func (x *ResumeDesc) GetSkills() []string {
	if x != nil {
		return x.Skills
	}
	return nil
}

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmploymentType int64         `protobuf:"varint,1,opt,name=employment_type,json=employmentType,proto3" json:"employment_type,omitempty"`
	WorkingHours   int64         `protobuf:"varint,2,opt,name=working_hours,json=workingHours,proto3" json:"working_hours,omitempty"`
	WishSalaryFrom int64         `protobuf:"varint,3,opt,name=wishSalary_from,json=wishSalaryFrom,proto3" json:"wishSalary_from,omitempty"`
	WishSalaryTo   int64         `protobuf:"varint,4,opt,name=wishSalary_to,json=wishSalaryTo,proto3" json:"wishSalary_to,omitempty"`
	Specialization string        `protobuf:"bytes,5,opt,name=specialization,proto3" json:"specialization,omitempty"`
	CandidateBio   *CandidateBio `protobuf:"bytes,6,opt,name=CandidateBio,proto3" json:"CandidateBio,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	mi := &file_resume_models_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_resume_models_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_resume_models_proto_rawDescGZIP(), []int{2}
}

func (x *Candidate) GetEmploymentType() int64 {
	if x != nil {
		return x.EmploymentType
	}
	return 0
}

func (x *Candidate) GetWorkingHours() int64 {
	if x != nil {
		return x.WorkingHours
	}
	return 0
}

func (x *Candidate) GetWishSalaryFrom() int64 {
	if x != nil {
		return x.WishSalaryFrom
	}
	return 0
}

func (x *Candidate) GetWishSalaryTo() int64 {
	if x != nil {
		return x.WishSalaryTo
	}
	return 0
}

func (x *Candidate) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

func (x *Candidate) GetCandidateBio() *CandidateBio {
	if x != nil {
		return x.CandidateBio
	}
	return nil
}

type CandidateBio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gender  int64  `protobuf:"varint,1,opt,name=gender,proto3" json:"gender,omitempty"`
	Age     int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CandidateBio) Reset() {
	*x = CandidateBio{}
	mi := &file_resume_models_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CandidateBio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateBio) ProtoMessage() {}

func (x *CandidateBio) ProtoReflect() protoreflect.Message {
	mi := &file_resume_models_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateBio.ProtoReflect.Descriptor instead.
func (*CandidateBio) Descriptor() ([]byte, []int) {
	return file_resume_models_proto_rawDescGZIP(), []int{3}
}

func (x *CandidateBio) GetGender() int64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *CandidateBio) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CandidateBio) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type JobExp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateFrom int64    `protobuf:"varint,1,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo   int64    `protobuf:"varint,2,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
	JobName  string   `protobuf:"bytes,3,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	JobDesc  *JobDesc `protobuf:"bytes,4,opt,name=JobDesc,proto3" json:"JobDesc,omitempty"`
	JobPos   string   `protobuf:"bytes,5,opt,name=job_pos,json=jobPos,proto3" json:"job_pos,omitempty"`
	JobSpec  string   `protobuf:"bytes,6,opt,name=job_spec,json=jobSpec,proto3" json:"job_spec,omitempty"`
}

func (x *JobExp) Reset() {
	*x = JobExp{}
	mi := &file_resume_models_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobExp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobExp) ProtoMessage() {}

func (x *JobExp) ProtoReflect() protoreflect.Message {
	mi := &file_resume_models_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobExp.ProtoReflect.Descriptor instead.
func (*JobExp) Descriptor() ([]byte, []int) {
	return file_resume_models_proto_rawDescGZIP(), []int{4}
}

func (x *JobExp) GetDateFrom() int64 {
	if x != nil {
		return x.DateFrom
	}
	return 0
}

func (x *JobExp) GetDateTo() int64 {
	if x != nil {
		return x.DateTo
	}
	return 0
}

func (x *JobExp) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *JobExp) GetJobDesc() *JobDesc {
	if x != nil {
		return x.JobDesc
	}
	return nil
}

func (x *JobExp) GetJobPos() string {
	if x != nil {
		return x.JobPos
	}
	return ""
}

func (x *JobExp) GetJobSpec() string {
	if x != nil {
		return x.JobSpec
	}
	return ""
}

type JobDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Website  string `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *JobDesc) Reset() {
	*x = JobDesc{}
	mi := &file_resume_models_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDesc) ProtoMessage() {}

func (x *JobDesc) ProtoReflect() protoreflect.Message {
	mi := &file_resume_models_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDesc.ProtoReflect.Descriptor instead.
func (*JobDesc) Descriptor() ([]byte, []int) {
	return file_resume_models_proto_rawDescGZIP(), []int{5}
}

func (x *JobDesc) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *JobDesc) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

var File_resume_models_proto protoreflect.FileDescriptor

var file_resume_models_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x8e, 0x01,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x65, 0x73, 0x63, 0x22, 0x7f,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x28, 0x0a, 0x07,
	0x4a, 0x6f, 0x62, 0x45, 0x78, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x70, 0x52, 0x07, 0x4a,
	0x6f, 0x62, 0x45, 0x78, 0x70, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x09, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x22,
	0x89, 0x02, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x77,
	0x69, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x77, 0x69, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x69, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x77, 0x69, 0x73,
	0x68, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x54, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6f, 0x52, 0x0c, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6f, 0x22, 0x52, 0x0a, 0x0c, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xb8, 0x01, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x45, 0x78, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x4a,
	0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x52, 0x07, 0x4a,
	0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x5f, 0x70, 0x6f,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x53, 0x70, 0x65, 0x63, 0x22, 0x3f, 0x0a, 0x07, 0x4a, 0x6f,
	0x62, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resume_models_proto_rawDescOnce sync.Once
	file_resume_models_proto_rawDescData = file_resume_models_proto_rawDesc
)

func file_resume_models_proto_rawDescGZIP() []byte {
	file_resume_models_proto_rawDescOnce.Do(func() {
		file_resume_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_resume_models_proto_rawDescData)
	})
	return file_resume_models_proto_rawDescData
}

var file_resume_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_resume_models_proto_goTypes = []any{
	(*Resume)(nil),       // 0: models.Resume
	(*ResumeDesc)(nil),   // 1: models.ResumeDesc
	(*Candidate)(nil),    // 2: models.Candidate
	(*CandidateBio)(nil), // 3: models.CandidateBio
	(*JobExp)(nil),       // 4: models.JobExp
	(*JobDesc)(nil),      // 5: models.JobDesc
}
var file_resume_models_proto_depIdxs = []int32{
	1, // 0: models.Resume.ResumeDesc:type_name -> models.ResumeDesc
	4, // 1: models.ResumeDesc.JobExps:type_name -> models.JobExp
	2, // 2: models.ResumeDesc.Candidate:type_name -> models.Candidate
	3, // 3: models.Candidate.CandidateBio:type_name -> models.CandidateBio
	5, // 4: models.JobExp.JobDesc:type_name -> models.JobDesc
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_resume_models_proto_init() }
func file_resume_models_proto_init() {
	if File_resume_models_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resume_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resume_models_proto_goTypes,
		DependencyIndexes: file_resume_models_proto_depIdxs,
		MessageInfos:      file_resume_models_proto_msgTypes,
	}.Build()
	File_resume_models_proto = out.File
	file_resume_models_proto_rawDesc = nil
	file_resume_models_proto_goTypes = nil
	file_resume_models_proto_depIdxs = nil
}