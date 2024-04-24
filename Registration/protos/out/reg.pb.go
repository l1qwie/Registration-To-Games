// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: reg.proto

package out

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

type ScheduleServRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gameid   int64  `protobuf:"varint,1,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Sport    string `protobuf:"bytes,2,opt,name=sport,proto3" json:"sport,omitempty"`
	Date     string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Time     string `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Seats    int64  `protobuf:"varint,5,opt,name=seats,proto3" json:"seats,omitempty"`
	Price    int64  `protobuf:"varint,9,opt,name=price,proto3" json:"price,omitempty"`
	Currency string `protobuf:"bytes,10,opt,name=currency,proto3" json:"currency,omitempty"`
	Action   string `protobuf:"bytes,12,opt,name=action,proto3" json:"action,omitempty"` //new, del or change
}

func (x *ScheduleServRequest) Reset() {
	*x = ScheduleServRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleServRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleServRequest) ProtoMessage() {}

func (x *ScheduleServRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleServRequest.ProtoReflect.Descriptor instead.
func (*ScheduleServRequest) Descriptor() ([]byte, []int) {
	return file_reg_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleServRequest) GetGameid() int64 {
	if x != nil {
		return x.Gameid
	}
	return 0
}

func (x *ScheduleServRequest) GetSport() string {
	if x != nil {
		return x.Sport
	}
	return ""
}

func (x *ScheduleServRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ScheduleServRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ScheduleServRequest) GetSeats() int64 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *ScheduleServRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ScheduleServRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ScheduleServRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type MediaServRequestSch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gameid int64  `protobuf:"varint,1,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Sport  string `protobuf:"bytes,2,opt,name=sport,proto3" json:"sport,omitempty"`
	Date   int64  `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	Time   int64  `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"` //new, del or change
	Status int64  `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MediaServRequestSch) Reset() {
	*x = MediaServRequestSch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaServRequestSch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaServRequestSch) ProtoMessage() {}

func (x *MediaServRequestSch) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaServRequestSch.ProtoReflect.Descriptor instead.
func (*MediaServRequestSch) Descriptor() ([]byte, []int) {
	return file_reg_proto_rawDescGZIP(), []int{1}
}

func (x *MediaServRequestSch) GetGameid() int64 {
	if x != nil {
		return x.Gameid
	}
	return 0
}

func (x *MediaServRequestSch) GetSport() string {
	if x != nil {
		return x.Sport
	}
	return ""
}

func (x *MediaServRequestSch) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *MediaServRequestSch) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *MediaServRequestSch) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *MediaServRequestSch) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SettingServRequestSch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gameid   int64  `protobuf:"varint,1,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Sport    string `protobuf:"bytes,2,opt,name=sport,proto3" json:"sport,omitempty"`
	Date     int64  `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	Time     int64  `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Price    int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Currency string `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Seats    int64  `protobuf:"varint,7,opt,name=seats,proto3" json:"seats,omitempty"`
	Status   int64  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Action   string `protobuf:"bytes,9,opt,name=action,proto3" json:"action,omitempty"` //new, del or change
}

func (x *SettingServRequestSch) Reset() {
	*x = SettingServRequestSch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingServRequestSch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingServRequestSch) ProtoMessage() {}

func (x *SettingServRequestSch) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingServRequestSch.ProtoReflect.Descriptor instead.
func (*SettingServRequestSch) Descriptor() ([]byte, []int) {
	return file_reg_proto_rawDescGZIP(), []int{2}
}

func (x *SettingServRequestSch) GetGameid() int64 {
	if x != nil {
		return x.Gameid
	}
	return 0
}

func (x *SettingServRequestSch) GetSport() string {
	if x != nil {
		return x.Sport
	}
	return ""
}

func (x *SettingServRequestSch) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *SettingServRequestSch) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *SettingServRequestSch) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SettingServRequestSch) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SettingServRequestSch) GetSeats() int64 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *SettingServRequestSch) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SettingServRequestSch) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type SettingServRequestGWU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid  int64  `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Gameid  int64  `protobuf:"varint,3,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Seats   int64  `protobuf:"varint,4,opt,name=seats,proto3" json:"seats,omitempty"`
	Payment string `protobuf:"bytes,5,opt,name=payment,proto3" json:"payment,omitempty"`
	Statpay int64  `protobuf:"varint,6,opt,name=statpay,proto3" json:"statpay,omitempty"`
	Status  int64  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Action  string `protobuf:"bytes,8,opt,name=action,proto3" json:"action,omitempty"` //new, del or change
}

func (x *SettingServRequestGWU) Reset() {
	*x = SettingServRequestGWU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingServRequestGWU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingServRequestGWU) ProtoMessage() {}

func (x *SettingServRequestGWU) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingServRequestGWU.ProtoReflect.Descriptor instead.
func (*SettingServRequestGWU) Descriptor() ([]byte, []int) {
	return file_reg_proto_rawDescGZIP(), []int{3}
}

func (x *SettingServRequestGWU) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SettingServRequestGWU) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *SettingServRequestGWU) GetGameid() int64 {
	if x != nil {
		return x.Gameid
	}
	return 0
}

func (x *SettingServRequestGWU) GetSeats() int64 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *SettingServRequestGWU) GetPayment() string {
	if x != nil {
		return x.Payment
	}
	return ""
}

func (x *SettingServRequestGWU) GetStatpay() int64 {
	if x != nil {
		return x.Statpay
	}
	return 0
}

func (x *SettingServRequestGWU) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SettingServRequestGWU) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type RegServRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gameid   int64  `protobuf:"varint,1,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Sport    string `protobuf:"bytes,2,opt,name=sport,proto3" json:"sport,omitempty"`
	Date     int64  `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	Time     int64  `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Seats    int64  `protobuf:"varint,5,opt,name=seats,proto3" json:"seats,omitempty"`
	Price    int64  `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	Currency string `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	Status   int64  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Action   string `protobuf:"bytes,9,opt,name=action,proto3" json:"action,omitempty"` //new, del or change
}

func (x *RegServRequest) Reset() {
	*x = RegServRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegServRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegServRequest) ProtoMessage() {}

func (x *RegServRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegServRequest.ProtoReflect.Descriptor instead.
func (*RegServRequest) Descriptor() ([]byte, []int) {
	return file_reg_proto_rawDescGZIP(), []int{4}
}

func (x *RegServRequest) GetGameid() int64 {
	if x != nil {
		return x.Gameid
	}
	return 0
}

func (x *RegServRequest) GetSport() string {
	if x != nil {
		return x.Sport
	}
	return ""
}

func (x *RegServRequest) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *RegServRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RegServRequest) GetSeats() int64 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *RegServRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *RegServRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *RegServRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RegServRequest) GetAction() string {
	if x != nil {
		return x.Action
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
		mi := &file_reg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reg_proto_msgTypes[5]
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
	return file_reg_proto_rawDescGZIP(), []int{5}
}

var File_reg_proto protoreflect.FileDescriptor

var file_reg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x65, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x22, 0xcb, 0x01, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x9b, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xe5, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd1, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47,
	0x57, 0x55, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x70, 0x61, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x70, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xde, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x67, 0x53, 0x65, 0x72, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe5,
	0x02, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x43, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1c,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x53, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x63,
	0x68, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x47, 0x57, 0x55, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x57, 0x55, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x52, 0x65, 0x67, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e,
	0x52, 0x65, 0x67, 0x53, 0x65, 0x72, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6f, 0x75, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reg_proto_rawDescOnce sync.Once
	file_reg_proto_rawDescData = file_reg_proto_rawDesc
)

func file_reg_proto_rawDescGZIP() []byte {
	file_reg_proto_rawDescOnce.Do(func() {
		file_reg_proto_rawDescData = protoimpl.X.CompressGZIP(file_reg_proto_rawDescData)
	})
	return file_reg_proto_rawDescData
}

var file_reg_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_reg_proto_goTypes = []interface{}{
	(*ScheduleServRequest)(nil),   // 0: registr.ScheduleServRequest
	(*MediaServRequestSch)(nil),   // 1: registr.MediaServRequestSch
	(*SettingServRequestSch)(nil), // 2: registr.SettingServRequestSch
	(*SettingServRequestGWU)(nil), // 3: registr.SettingServRequestGWU
	(*RegServRequest)(nil),        // 4: registr.RegServRequest
	(*EmptyResponse)(nil),         // 5: registr.EmptyResponse
}
var file_reg_proto_depIdxs = []int32{
	0, // 0: registr.Registration.UpdSchedule:input_type -> registr.ScheduleServRequest
	1, // 1: registr.Registration.UpdMediaSch:input_type -> registr.MediaServRequestSch
	2, // 2: registr.Registration.UpdSettingSch:input_type -> registr.SettingServRequestSch
	3, // 3: registr.Registration.UpdSettingGWU:input_type -> registr.SettingServRequestGWU
	4, // 4: registr.Registration.UpdReg:input_type -> registr.RegServRequest
	5, // 5: registr.Registration.UpdSchedule:output_type -> registr.EmptyResponse
	5, // 6: registr.Registration.UpdMediaSch:output_type -> registr.EmptyResponse
	5, // 7: registr.Registration.UpdSettingSch:output_type -> registr.EmptyResponse
	5, // 8: registr.Registration.UpdSettingGWU:output_type -> registr.EmptyResponse
	5, // 9: registr.Registration.UpdReg:output_type -> registr.EmptyResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reg_proto_init() }
func file_reg_proto_init() {
	if File_reg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleServRequest); i {
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
		file_reg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaServRequestSch); i {
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
		file_reg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingServRequestSch); i {
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
		file_reg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingServRequestGWU); i {
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
		file_reg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegServRequest); i {
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
		file_reg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_reg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reg_proto_goTypes,
		DependencyIndexes: file_reg_proto_depIdxs,
		MessageInfos:      file_reg_proto_msgTypes,
	}.Build()
	File_reg_proto = out.File
	file_reg_proto_rawDesc = nil
	file_reg_proto_goTypes = nil
	file_reg_proto_depIdxs = nil
}
