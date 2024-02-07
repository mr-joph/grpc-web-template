// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/auth.proto

package pb

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

type LoginIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string  `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Password   string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role       *string `protobuf:"bytes,3,opt,name=role,proto3,oneof" json:"role,omitempty"`
}

func (x *LoginIn) Reset() {
	*x = LoginIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginIn) ProtoMessage() {}

func (x *LoginIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginIn.ProtoReflect.Descriptor instead.
func (*LoginIn) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginIn) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *LoginIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginIn) GetRole() string {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return ""
}

type LoginOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginOut) Reset() {
	*x = LoginOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginOut) ProtoMessage() {}

func (x *LoginOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginOut.ProtoReflect.Descriptor instead.
func (*LoginOut) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginOut) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckEmailIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CheckEmailIn) Reset() {
	*x = CheckEmailIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEmailIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEmailIn) ProtoMessage() {}

func (x *CheckEmailIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEmailIn.ProtoReflect.Descriptor instead.
func (*CheckEmailIn) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CheckEmailIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CheckEmailOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid  bool   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *CheckEmailOut) Reset() {
	*x = CheckEmailOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckEmailOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckEmailOut) ProtoMessage() {}

func (x *CheckEmailOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckEmailOut.ProtoReflect.Descriptor instead.
func (*CheckEmailOut) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CheckEmailOut) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *CheckEmailOut) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type VerifyUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyUserIn) Reset() {
	*x = VerifyUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserIn) ProtoMessage() {}

func (x *VerifyUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserIn.ProtoReflect.Descriptor instead.
func (*VerifyUserIn) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyUserIn) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyUserIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyUserOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *VerifyUserOut) Reset() {
	*x = VerifyUserOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserOut) ProtoMessage() {}

func (x *VerifyUserOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserOut.ProtoReflect.Descriptor instead.
func (*VerifyUserOut) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyUserOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type TicketResetPasswordIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string  `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	CallbackURL *string `protobuf:"bytes,2,opt,name=callbackURL,proto3,oneof" json:"callbackURL,omitempty"`
}

func (x *TicketResetPasswordIn) Reset() {
	*x = TicketResetPasswordIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketResetPasswordIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketResetPasswordIn) ProtoMessage() {}

func (x *TicketResetPasswordIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketResetPasswordIn.ProtoReflect.Descriptor instead.
func (*TicketResetPasswordIn) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *TicketResetPasswordIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *TicketResetPasswordIn) GetCallbackURL() string {
	if x != nil && x.CallbackURL != nil {
		return *x.CallbackURL
	}
	return ""
}

type TicketResetPasswordOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TicketResetPasswordOut) Reset() {
	*x = TicketResetPasswordOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketResetPasswordOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketResetPasswordOut) ProtoMessage() {}

func (x *TicketResetPasswordOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketResetPasswordOut.ProtoReflect.Descriptor instead.
func (*TicketResetPasswordOut) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{7}
}

func (x *TicketResetPasswordOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ResetPasswordIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Ticket          string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	// optional field
	CallbackURL *string `protobuf:"bytes,5,opt,name=callbackURL,proto3,oneof" json:"callbackURL,omitempty"`
}

func (x *ResetPasswordIn) Reset() {
	*x = ResetPasswordIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordIn) ProtoMessage() {}

func (x *ResetPasswordIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordIn.ProtoReflect.Descriptor instead.
func (*ResetPasswordIn) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{8}
}

func (x *ResetPasswordIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResetPasswordIn) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (x *ResetPasswordIn) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ResetPasswordIn) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

func (x *ResetPasswordIn) GetCallbackURL() string {
	if x != nil && x.CallbackURL != nil {
		return *x.CallbackURL
	}
	return ""
}

type ResendEmailIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// optional field
	CallbackURL *string `protobuf:"bytes,4,opt,name=callbackURL,proto3,oneof" json:"callbackURL,omitempty"`
}

func (x *ResendEmailIn) Reset() {
	*x = ResendEmailIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendEmailIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendEmailIn) ProtoMessage() {}

func (x *ResendEmailIn) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendEmailIn.ProtoReflect.Descriptor instead.
func (*ResendEmailIn) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{9}
}

func (x *ResendEmailIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResendEmailIn) GetCallbackURL() string {
	if x != nil && x.CallbackURL != nil {
		return *x.CallbackURL
	}
	return ""
}

type ResetPasswordOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ResetPasswordOut) Reset() {
	*x = ResetPasswordOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordOut) ProtoMessage() {}

func (x *ResetPasswordOut) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordOut.ProtoReflect.Descriptor instead.
func (*ResetPasswordOut) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{10}
}

func (x *ResetPasswordOut) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_auth_proto protoreflect.FileDescriptor

var file_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4f,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3d,
	0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x38, 0x0a,
	0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x29, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x64, 0x0a, 0x15, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52, 0x4c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x55, 0x52, 0x4c, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52, 0x4c, 0x22, 0x32, 0x0a, 0x16, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4f,
	0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xc2, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52, 0x4c, 0x88, 0x01,
	0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52,
	0x4c, 0x22, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x55, 0x52, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52, 0x4c, 0x88, 0x01, 0x01, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x52, 0x4c, 0x22,
	0x2c, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xf0, 0x02,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x35, 0x0a, 0x0a,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x13,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x12, 0x50, 0x0a, 0x13, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x1a, 0x16, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x39, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData = file_proto_auth_proto_rawDesc
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_proto_rawDescData)
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_auth_proto_goTypes = []interface{}{
	(*LoginIn)(nil),                // 0: core.LoginIn
	(*LoginOut)(nil),               // 1: core.LoginOut
	(*CheckEmailIn)(nil),           // 2: core.CheckEmailIn
	(*CheckEmailOut)(nil),          // 3: core.CheckEmailOut
	(*VerifyUserIn)(nil),           // 4: core.VerifyUserIn
	(*VerifyUserOut)(nil),          // 5: core.VerifyUserOut
	(*TicketResetPasswordIn)(nil),  // 6: core.TicketResetPasswordIn
	(*TicketResetPasswordOut)(nil), // 7: core.TicketResetPasswordOut
	(*ResetPasswordIn)(nil),        // 8: core.ResetPasswordIn
	(*ResendEmailIn)(nil),          // 9: core.ResendEmailIn
	(*ResetPasswordOut)(nil),       // 10: core.ResetPasswordOut
	(*Empty)(nil),                  // 11: core.Empty
}
var file_proto_auth_proto_depIdxs = []int32{
	0,  // 0: core.AuthService.Login:input_type -> core.LoginIn
	2,  // 1: core.AuthService.CheckEmail:input_type -> core.CheckEmailIn
	4,  // 2: core.AuthService.VerifyUser:input_type -> core.VerifyUserIn
	6,  // 3: core.AuthService.TicketResetPassword:input_type -> core.TicketResetPasswordIn
	8,  // 4: core.AuthService.ResetPassword:input_type -> core.ResetPasswordIn
	9,  // 5: core.AuthService.ResendEmailToValidate:input_type -> core.ResendEmailIn
	1,  // 6: core.AuthService.Login:output_type -> core.LoginOut
	3,  // 7: core.AuthService.CheckEmail:output_type -> core.CheckEmailOut
	5,  // 8: core.AuthService.VerifyUser:output_type -> core.VerifyUserOut
	7,  // 9: core.AuthService.TicketResetPassword:output_type -> core.TicketResetPasswordOut
	10, // 10: core.AuthService.ResetPassword:output_type -> core.ResetPasswordOut
	11, // 11: core.AuthService.ResendEmailToValidate:output_type -> core.Empty
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	file_proto_schemas_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginIn); i {
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
		file_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginOut); i {
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
		file_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEmailIn); i {
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
		file_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckEmailOut); i {
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
		file_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserIn); i {
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
		file_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserOut); i {
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
		file_proto_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketResetPasswordIn); i {
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
		file_proto_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketResetPasswordOut); i {
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
		file_proto_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordIn); i {
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
		file_proto_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendEmailIn); i {
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
		file_proto_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordOut); i {
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
	file_proto_auth_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_auth_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_proto_auth_proto_msgTypes[8].OneofWrappers = []interface{}{}
	file_proto_auth_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_rawDesc = nil
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}
