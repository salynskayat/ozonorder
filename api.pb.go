// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/api.proto

package ozonorder

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64   `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Quantity  int32   `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items            []*Item                `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	CustomerId       int64                  `protobuf:"varint,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	PaymentMethodId  int64                  `protobuf:"varint,3,opt,name=PaymentMethodId,proto3" json:"PaymentMethodId,omitempty"`
	DeliveryMethodId int64                  `protobuf:"varint,4,opt,name=DeliveryMethodId,proto3" json:"DeliveryMethodId,omitempty"`
	Operation        string                 `protobuf:"bytes,5,opt,name=Operation,proto3" json:"Operation,omitempty"`
	OperationDate    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=OperationDate,proto3" json:"OperationDate,omitempty"`
	Ts               int64                  `protobuf:"varint,7,opt,name=Ts,proto3" json:"Ts,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
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
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Order) GetPaymentMethodId() int64 {
	if x != nil {
		return x.PaymentMethodId
	}
	return 0
}

func (x *Order) GetDeliveryMethodId() int64 {
	if x != nil {
		return x.DeliveryMethodId
	}
	return 0
}

func (x *Order) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Order) GetOperationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OperationDate
	}
	return nil
}

func (x *Order) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Ts int64 `protobuf:"varint,2,opt,name=Ts,proto3" json:"Ts,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *OrderResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderResponse) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x8e, 0x02,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0d,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x54, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x54, 0x73, 0x22, 0x2f,
	0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x54, 0x73, 0x32,
	0x3c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x12, 0x0a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x79,
	0x6e, 0x73, 0x6b, 0x61, 0x79, 0x61, 0x74, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_api_proto_goTypes = []interface{}{
	(*Item)(nil),                  // 0: api.Item
	(*Order)(nil),                 // 1: api.Order
	(*OrderResponse)(nil),         // 2: api.OrderResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_api_proto_depIdxs = []int32{
	0, // 0: api.Order.items:type_name -> api.Item
	3, // 1: api.Order.OperationDate:type_name -> google.protobuf.Timestamp
	1, // 2: api.OrderService.Interact:input_type -> api.Order
	2, // 3: api.OrderService.Interact:output_type -> api.OrderResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
