// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: IDLs/order.proto

package order

import (
	context "context"
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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int64  `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=StatusMessage,proto3" json:"StatusMessage,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`                 //销售记录ID
	UserId     int64  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`         //买票人
	ScheduleId int64  `protobuf:"varint,3,opt,name=ScheduleId,proto3" json:"ScheduleId,omitempty"` //演出计划ID
	SeatRow    int32  `protobuf:"varint,4,opt,name=SeatRow,proto3" json:"SeatRow,omitempty"`       //座位行数
	SeatCol    int32  `protobuf:"varint,5,opt,name=SeatCol,proto3" json:"SeatCol,omitempty"`       //座位列数
	Date       string `protobuf:"bytes,6,opt,name=Date,proto3" json:"Date,omitempty"`              //订单处理时间 按照示例格式：2006-01-02 15:04:05（\d{4}-\d{2}-\d{2}[ T]\d{2}:\d{2}:\d{2}）
	Value      int32  `protobuf:"varint,7,opt,name=Value,proto3" json:"Value,omitempty"`           //票价
	Type       int32  `protobuf:"varint,8,opt,name=Type,proto3" json:"Type,omitempty"`             //交易类型，1-买票，-1-退票
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Order) GetScheduleId() int64 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *Order) GetSeatRow() int32 {
	if x != nil {
		return x.SeatRow
	}
	return 0
}

func (x *Order) GetSeatCol() int32 {
	if x != nil {
		return x.SeatCol
	}
	return 0
}

func (x *Order) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Order) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Order) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

// 用户抢到了票，创建订单，默认是买票行为
//
//	message AddOrderRequest{
//	 int64 UserId = 1;
//	 int64 ScheduleId = 2;//演出计划ID
//	 int32 SeatRow = 3;//座位行数
//	 int32 SeatCol = 4;//座位列数
//	 string Data = 5;//下单时间
//	 int32 Value = 6;//票价
//	}
//
//	message AddOrderResponse{
//	 BaseResp BaseResp = 1;
//	}
//
// 进行支付，确认订单
type CommitOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ScheduleId int64 `protobuf:"varint,2,opt,name=ScheduleId,proto3" json:"ScheduleId,omitempty"` //演出计划ID
	SeatRow    int32 `protobuf:"varint,3,opt,name=SeatRow,proto3" json:"SeatRow,omitempty"`       //座位行数
	SeatCol    int32 `protobuf:"varint,4,opt,name=SeatCol,proto3" json:"SeatCol,omitempty"`       //座位列数
}

func (x *CommitOrderRequest) Reset() {
	*x = CommitOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitOrderRequest) ProtoMessage() {}

func (x *CommitOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitOrderRequest.ProtoReflect.Descriptor instead.
func (*CommitOrderRequest) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{2}
}

func (x *CommitOrderRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommitOrderRequest) GetScheduleId() int64 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *CommitOrderRequest) GetSeatRow() int32 {
	if x != nil {
		return x.SeatRow
	}
	return 0
}

func (x *CommitOrderRequest) GetSeatCol() int32 {
	if x != nil {
		return x.SeatCol
	}
	return 0
}

type CommitOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
}

func (x *CommitOrderResponse) Reset() {
	*x = CommitOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitOrderResponse) ProtoMessage() {}

func (x *CommitOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitOrderResponse.ProtoReflect.Descriptor instead.
func (*CommitOrderResponse) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{3}
}

func (x *CommitOrderResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// [给用户使用]根据用户ID 获取全部订单信息
type GetAllOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetAllOrderRequest) Reset() {
	*x = GetAllOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrderRequest) ProtoMessage() {}

func (x *GetAllOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrderRequest.ProtoReflect.Descriptor instead.
func (*GetAllOrderRequest) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllOrderRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAllOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
	List     []*Order  `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *GetAllOrderResponse) Reset() {
	*x = GetAllOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrderResponse) ProtoMessage() {}

func (x *GetAllOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrderResponse.ProtoReflect.Descriptor instead.
func (*GetAllOrderResponse) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllOrderResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetAllOrderResponse) GetList() []*Order {
	if x != nil {
		return x.List
	}
	return nil
}

// 票房
type OrderAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayId       int64  `protobuf:"varint,1,opt,name=PlayId,proto3" json:"PlayId,omitempty"`
	PlayName     string `protobuf:"bytes,2,opt,name=PlayName,proto3" json:"PlayName,omitempty"`
	PlayArea     string `protobuf:"bytes,3,opt,name=PlayArea,proto3" json:"PlayArea,omitempty"`
	PlayDuration string `protobuf:"bytes,4,opt,name=PlayDuration,proto3" json:"PlayDuration,omitempty"`
	StartData    string `protobuf:"bytes,5,opt,name=StartData,proto3" json:"StartData,omitempty"`      //剧目开始时间 按照示例格式：2006-01-02 15:04:05（\d{4}-\d{2}-\d{2}[ T]\d{2}:\d{2}:\d{2}）
	EndData      string `protobuf:"bytes,6,opt,name=EndData,proto3" json:"EndData,omitempty"`          //剧目结束时间 按照示例格式：2006-01-02 15:04:05（\d{4}-\d{2}-\d{2}[ T]\d{2}:\d{2}:\d{2}）
	TotalTicket  int64  `protobuf:"varint,7,opt,name=TotalTicket,proto3" json:"TotalTicket,omitempty"` //总卖票数
	Sales        int64  `protobuf:"varint,8,opt,name=Sales,proto3" json:"Sales,omitempty"`             //销售额
	Price        int32  `protobuf:"varint,9,opt,name=Price,proto3" json:"Price,omitempty"`             //票价
}

func (x *OrderAnalysis) Reset() {
	*x = OrderAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAnalysis) ProtoMessage() {}

func (x *OrderAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAnalysis.ProtoReflect.Descriptor instead.
func (*OrderAnalysis) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderAnalysis) GetPlayId() int64 {
	if x != nil {
		return x.PlayId
	}
	return 0
}

func (x *OrderAnalysis) GetPlayName() string {
	if x != nil {
		return x.PlayName
	}
	return ""
}

func (x *OrderAnalysis) GetPlayArea() string {
	if x != nil {
		return x.PlayArea
	}
	return ""
}

func (x *OrderAnalysis) GetPlayDuration() string {
	if x != nil {
		return x.PlayDuration
	}
	return ""
}

func (x *OrderAnalysis) GetStartData() string {
	if x != nil {
		return x.StartData
	}
	return ""
}

func (x *OrderAnalysis) GetEndData() string {
	if x != nil {
		return x.EndData
	}
	return ""
}

func (x *OrderAnalysis) GetTotalTicket() int64 {
	if x != nil {
		return x.TotalTicket
	}
	return 0
}

func (x *OrderAnalysis) GetSales() int64 {
	if x != nil {
		return x.Sales
	}
	return 0
}

func (x *OrderAnalysis) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetOrderAnalysisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayId int64 `protobuf:"varint,1,opt,name=PlayId,proto3" json:"PlayId,omitempty"`
}

func (x *GetOrderAnalysisRequest) Reset() {
	*x = GetOrderAnalysisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderAnalysisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderAnalysisRequest) ProtoMessage() {}

func (x *GetOrderAnalysisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderAnalysisRequest.ProtoReflect.Descriptor instead.
func (*GetOrderAnalysisRequest) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{7}
}

func (x *GetOrderAnalysisRequest) GetPlayId() int64 {
	if x != nil {
		return x.PlayId
	}
	return 0
}

type GetOrderAnalysisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp      *BaseResp      `protobuf:"bytes,1,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
	OrderAnalysis *OrderAnalysis `protobuf:"bytes,2,opt,name=OrderAnalysis,proto3" json:"OrderAnalysis,omitempty"`
}

func (x *GetOrderAnalysisResponse) Reset() {
	*x = GetOrderAnalysisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IDLs_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderAnalysisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderAnalysisResponse) ProtoMessage() {}

func (x *GetOrderAnalysisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_IDLs_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderAnalysisResponse.ProtoReflect.Descriptor instead.
func (*GetOrderAnalysisResponse) Descriptor() ([]byte, []int) {
	return file_IDLs_order_proto_rawDescGZIP(), []int{8}
}

func (x *GetOrderAnalysisResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetOrderAnalysisResponse) GetOrderAnalysis() *OrderAnalysis {
	if x != nil {
		return x.OrderAnalysis
	}
	return nil
}

var File_IDLs_order_proto protoreflect.FileDescriptor

var file_IDLs_order_proto_rawDesc = []byte{
	0x0a, 0x10, 0x49, 0x44, 0x4c, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x08, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x53, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x74, 0x43,
	0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x61, 0x74, 0x43, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x80, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x53, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x61, 0x74,
	0x43, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x61, 0x74, 0x43,
	0x6f, 0x6c, 0x22, 0x42, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x89, 0x02, 0x0a, 0x0d, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6c,
	0x61, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x72, 0x65, 0x61, 0x12, 0x22, 0x0a, 0x0c,
	0x50, 0x6c, 0x61, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x52, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x32,
	0xf5, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x1e, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x54, 0x54, 0x4d, 0x53, 0x2f,
	0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IDLs_order_proto_rawDescOnce sync.Once
	file_IDLs_order_proto_rawDescData = file_IDLs_order_proto_rawDesc
)

func file_IDLs_order_proto_rawDescGZIP() []byte {
	file_IDLs_order_proto_rawDescOnce.Do(func() {
		file_IDLs_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_IDLs_order_proto_rawDescData)
	})
	return file_IDLs_order_proto_rawDescData
}

var file_IDLs_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_IDLs_order_proto_goTypes = []interface{}{
	(*BaseResp)(nil),                 // 0: order.BaseResp
	(*Order)(nil),                    // 1: order.Order
	(*CommitOrderRequest)(nil),       // 2: order.CommitOrderRequest
	(*CommitOrderResponse)(nil),      // 3: order.CommitOrderResponse
	(*GetAllOrderRequest)(nil),       // 4: order.GetAllOrderRequest
	(*GetAllOrderResponse)(nil),      // 5: order.GetAllOrderResponse
	(*OrderAnalysis)(nil),            // 6: order.OrderAnalysis
	(*GetOrderAnalysisRequest)(nil),  // 7: order.GetOrderAnalysisRequest
	(*GetOrderAnalysisResponse)(nil), // 8: order.GetOrderAnalysisResponse
}
var file_IDLs_order_proto_depIdxs = []int32{
	0, // 0: order.CommitOrderResponse.BaseResp:type_name -> order.BaseResp
	0, // 1: order.GetAllOrderResponse.BaseResp:type_name -> order.BaseResp
	1, // 2: order.GetAllOrderResponse.List:type_name -> order.Order
	0, // 3: order.GetOrderAnalysisResponse.BaseResp:type_name -> order.BaseResp
	6, // 4: order.GetOrderAnalysisResponse.OrderAnalysis:type_name -> order.OrderAnalysis
	4, // 5: order.OrderService.GetAllOrder:input_type -> order.GetAllOrderRequest
	7, // 6: order.OrderService.GetOrderAnalysis:input_type -> order.GetOrderAnalysisRequest
	2, // 7: order.OrderService.CommitOrder:input_type -> order.CommitOrderRequest
	5, // 8: order.OrderService.GetAllOrder:output_type -> order.GetAllOrderResponse
	8, // 9: order.OrderService.GetOrderAnalysis:output_type -> order.GetOrderAnalysisResponse
	3, // 10: order.OrderService.CommitOrder:output_type -> order.CommitOrderResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_IDLs_order_proto_init() }
func file_IDLs_order_proto_init() {
	if File_IDLs_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_IDLs_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_IDLs_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_IDLs_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitOrderRequest); i {
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
		file_IDLs_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitOrderResponse); i {
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
		file_IDLs_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrderRequest); i {
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
		file_IDLs_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrderResponse); i {
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
		file_IDLs_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAnalysis); i {
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
		file_IDLs_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderAnalysisRequest); i {
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
		file_IDLs_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderAnalysisResponse); i {
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
			RawDescriptor: file_IDLs_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_IDLs_order_proto_goTypes,
		DependencyIndexes: file_IDLs_order_proto_depIdxs,
		MessageInfos:      file_IDLs_order_proto_msgTypes,
	}.Build()
	File_IDLs_order_proto = out.File
	file_IDLs_order_proto_rawDesc = nil
	file_IDLs_order_proto_goTypes = nil
	file_IDLs_order_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type OrderService interface {
	GetAllOrder(ctx context.Context, req *GetAllOrderRequest) (res *GetAllOrderResponse, err error)
	GetOrderAnalysis(ctx context.Context, req *GetOrderAnalysisRequest) (res *GetOrderAnalysisResponse, err error)
	CommitOrder(ctx context.Context, req *CommitOrderRequest) (res *CommitOrderResponse, err error)
}
