// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: coursemanagement.proto

package course_management

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Professor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfessorName string `protobuf:"bytes,1,opt,name=professor_name,json=professorName,proto3" json:"professor_name,omitempty"`
}

func (x *Professor) Reset() {
	*x = Professor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coursemanagement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Professor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Professor) ProtoMessage() {}

func (x *Professor) ProtoReflect() protoreflect.Message {
	mi := &file_coursemanagement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Professor.ProtoReflect.Descriptor instead.
func (*Professor) Descriptor() ([]byte, []int) {
	return file_coursemanagement_proto_rawDescGZIP(), []int{0}
}

func (x *Professor) GetProfessorName() string {
	if x != nil {
		return x.ProfessorName
	}
	return ""
}

type Prerequisite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId string `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *Prerequisite) Reset() {
	*x = Prerequisite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coursemanagement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prerequisite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prerequisite) ProtoMessage() {}

func (x *Prerequisite) ProtoReflect() protoreflect.Message {
	mi := &file_coursemanagement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prerequisite.ProtoReflect.Descriptor instead.
func (*Prerequisite) Descriptor() ([]byte, []int) {
	return file_coursemanagement_proto_rawDescGZIP(), []int{1}
}

func (x *Prerequisite) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coursemanagement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_coursemanagement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_coursemanagement_proto_rawDescGZIP(), []int{2}
}

type CourseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId          string          `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`                            // course id Ex. 2110415
	CourseName        string          `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`                      // course name Ex. Software Architecture
	CourseDescription string          `protobuf:"bytes,3,opt,name=course_description,json=courseDescription,proto3" json:"course_description,omitempty"` // course contents Ex. [1. Introduction to Software Architecture]
	FacultyDepartment string          `protobuf:"bytes,4,opt,name=faculty_department,json=facultyDepartment,proto3" json:"faculty_department,omitempty"` // department name Ex. Computer Engineering
	AcademicTerm      string          `protobuf:"bytes,5,opt,name=academic_term,json=academicTerm,proto3" json:"academic_term,omitempty"`                // term name Ex. 1 , 2 , summer
	AcademicYear      int32           `protobuf:"varint,6,opt,name=academic_year,json=academicYear,proto3" json:"academic_year,omitempty"`               // year of term Ex. 2563
	Professors        []*Professor    `protobuf:"bytes,7,rep,name=professors,proto3" json:"professors,omitempty"`                                        // professors name Ex. [somchai jaidee]
	Prerequisites     []*Prerequisite `protobuf:"bytes,8,rep,name=prerequisites,proto3" json:"prerequisites,omitempty"`                                  // course id Ex. [2110416]
	Status            string          `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`                                                // course status Ex. open , close
	CurriculumName    string          `protobuf:"bytes,10,opt,name=curriculum_name,json=curriculumName,proto3" json:"curriculum_name,omitempty"`         // computer engineering
	DegreeLevel       string          `protobuf:"bytes,11,opt,name=degree_level,json=degreeLevel,proto3" json:"degree_level,omitempty"`                  // bachelor , master , doctor
	TeachingHours     int32           `protobuf:"varint,12,opt,name=teaching_hours,json=teachingHours,proto3" json:"teaching_hours,omitempty"`           // teaching hours Ex. 3
}

func (x *CourseItem) Reset() {
	*x = CourseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coursemanagement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseItem) ProtoMessage() {}

func (x *CourseItem) ProtoReflect() protoreflect.Message {
	mi := &file_coursemanagement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseItem.ProtoReflect.Descriptor instead.
func (*CourseItem) Descriptor() ([]byte, []int) {
	return file_coursemanagement_proto_rawDescGZIP(), []int{3}
}

func (x *CourseItem) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *CourseItem) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *CourseItem) GetCourseDescription() string {
	if x != nil {
		return x.CourseDescription
	}
	return ""
}

func (x *CourseItem) GetFacultyDepartment() string {
	if x != nil {
		return x.FacultyDepartment
	}
	return ""
}

func (x *CourseItem) GetAcademicTerm() string {
	if x != nil {
		return x.AcademicTerm
	}
	return ""
}

func (x *CourseItem) GetAcademicYear() int32 {
	if x != nil {
		return x.AcademicYear
	}
	return 0
}

func (x *CourseItem) GetProfessors() []*Professor {
	if x != nil {
		return x.Professors
	}
	return nil
}

func (x *CourseItem) GetPrerequisites() []*Prerequisite {
	if x != nil {
		return x.Prerequisites
	}
	return nil
}

func (x *CourseItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CourseItem) GetCurriculumName() string {
	if x != nil {
		return x.CurriculumName
	}
	return ""
}

func (x *CourseItem) GetDegreeLevel() string {
	if x != nil {
		return x.DegreeLevel
	}
	return ""
}

func (x *CourseItem) GetTeachingHours() int32 {
	if x != nil {
		return x.TeachingHours
	}
	return 0
}

type CourseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course []*CourseItem `protobuf:"bytes,1,rep,name=course,proto3" json:"course,omitempty"` // course list
}

func (x *CourseList) Reset() {
	*x = CourseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coursemanagement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseList) ProtoMessage() {}

func (x *CourseList) ProtoReflect() protoreflect.Message {
	mi := &file_coursemanagement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseList.ProtoReflect.Descriptor instead.
func (*CourseList) Descriptor() ([]byte, []int) {
	return file_coursemanagement_proto_rawDescGZIP(), []int{4}
}

func (x *CourseList) GetCourse() []*CourseItem {
	if x != nil {
		return x.Course
	}
	return nil
}

type CourseId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId string `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"` // course id Ex. 2110415
}

func (x *CourseId) Reset() {
	*x = CourseId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coursemanagement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseId) ProtoMessage() {}

func (x *CourseId) ProtoReflect() protoreflect.Message {
	mi := &file_coursemanagement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseId.ProtoReflect.Descriptor instead.
func (*CourseId) Descriptor() ([]byte, []int) {
	return file_coursemanagement_proto_rawDescGZIP(), []int{5}
}

func (x *CourseId) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

var File_coursemanagement_proto protoreflect.FileDescriptor

var file_coursemanagement_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a,
	0x0c, 0x50, 0x72, 0x65, 0x72, 0x65, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x82, 0x04, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x12, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x61, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x54,
	0x65, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x63, 0x61, 0x64,
	0x65, 0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x72, 0x65, 0x72, 0x65, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x65, 0x52, 0x0d,
	0x70, 0x72, 0x65, 0x72, 0x65, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75,
	0x6c, 0x75, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x75, 0x72, 0x72, 0x69, 0x63, 0x75, 0x6c, 0x75, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x6f,
	0x75, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x27, 0x0a,
	0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x32, 0xc2, 0x05, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12,
	0x86, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x37, 0x12, 0x35, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x82, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x64,
	0x4e, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22,
	0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x94, 0x01,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x1a, 0x35, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x84, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x37, 0x2a, 0x35, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x2d, 0x5a, 0x2b, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_coursemanagement_proto_rawDescOnce sync.Once
	file_coursemanagement_proto_rawDescData = file_coursemanagement_proto_rawDesc
)

func file_coursemanagement_proto_rawDescGZIP() []byte {
	file_coursemanagement_proto_rawDescOnce.Do(func() {
		file_coursemanagement_proto_rawDescData = protoimpl.X.CompressGZIP(file_coursemanagement_proto_rawDescData)
	})
	return file_coursemanagement_proto_rawDescData
}

var file_coursemanagement_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_coursemanagement_proto_goTypes = []interface{}{
	(*Professor)(nil),    // 0: course_management.Professor
	(*Prerequisite)(nil), // 1: course_management.Prerequisite
	(*Empty)(nil),        // 2: course_management.Empty
	(*CourseItem)(nil),   // 3: course_management.CourseItem
	(*CourseList)(nil),   // 4: course_management.CourseList
	(*CourseId)(nil),     // 5: course_management.CourseId
}
var file_coursemanagement_proto_depIdxs = []int32{
	0, // 0: course_management.CourseItem.professors:type_name -> course_management.Professor
	1, // 1: course_management.CourseItem.prerequisites:type_name -> course_management.Prerequisite
	3, // 2: course_management.CourseList.course:type_name -> course_management.CourseItem
	2, // 3: course_management.CourseManagementService.GetAllCourses:input_type -> course_management.Empty
	5, // 4: course_management.CourseManagementService.GetCourse:input_type -> course_management.CourseId
	3, // 5: course_management.CourseManagementService.AddNewCourse:input_type -> course_management.CourseItem
	3, // 6: course_management.CourseManagementService.UpdateCourseDetail:input_type -> course_management.CourseItem
	5, // 7: course_management.CourseManagementService.DeleteCourse:input_type -> course_management.CourseId
	4, // 8: course_management.CourseManagementService.GetAllCourses:output_type -> course_management.CourseList
	3, // 9: course_management.CourseManagementService.GetCourse:output_type -> course_management.CourseItem
	3, // 10: course_management.CourseManagementService.AddNewCourse:output_type -> course_management.CourseItem
	3, // 11: course_management.CourseManagementService.UpdateCourseDetail:output_type -> course_management.CourseItem
	2, // 12: course_management.CourseManagementService.DeleteCourse:output_type -> course_management.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_coursemanagement_proto_init() }
func file_coursemanagement_proto_init() {
	if File_coursemanagement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coursemanagement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Professor); i {
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
		file_coursemanagement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prerequisite); i {
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
		file_coursemanagement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_coursemanagement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseItem); i {
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
		file_coursemanagement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseList); i {
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
		file_coursemanagement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseId); i {
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
			RawDescriptor: file_coursemanagement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coursemanagement_proto_goTypes,
		DependencyIndexes: file_coursemanagement_proto_depIdxs,
		MessageInfos:      file_coursemanagement_proto_msgTypes,
	}.Build()
	File_coursemanagement_proto = out.File
	file_coursemanagement_proto_rawDesc = nil
	file_coursemanagement_proto_goTypes = nil
	file_coursemanagement_proto_depIdxs = nil
}
