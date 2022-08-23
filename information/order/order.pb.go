// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: information/order.proto

package order

import (
	order "github.com/modular-project/protobuffers/order/order"
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

type ValidateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *order.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *ValidateOrderRequest) Reset() {
	*x = ValidateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_information_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateOrderRequest) ProtoMessage() {}

func (x *ValidateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_information_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateOrderRequest.ProtoReflect.Descriptor instead.
func (*ValidateOrderRequest) Descriptor() ([]byte, []int) {
	return file_information_order_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateOrderRequest) GetOrder() *order.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type ValidateProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderProducts *order.OrderProduct `protobuf:"bytes,1,opt,name=order_products,json=orderProducts,proto3" json:"order_products,omitempty"`
}

func (x *ValidateProductsRequest) Reset() {
	*x = ValidateProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_information_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateProductsRequest) ProtoMessage() {}

func (x *ValidateProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_information_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateProductsRequest.ProtoReflect.Descriptor instead.
func (*ValidateProductsRequest) Descriptor() ([]byte, []int) {
	return file_information_order_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateProductsRequest) GetOrderProducts() *order.OrderProduct {
	if x != nil {
		return x.OrderProducts
	}
	return nil
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total float32 `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_information_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_information_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_information_order_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateResponse) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_information_order_proto protoreflect.FileDescriptor

var file_information_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x1a, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x61, 0x0a,
	0x17, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x22, 0x28, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xeb, 0x01, 0x0a, 0x0d, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x0d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x30, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_information_order_proto_rawDescOnce sync.Once
	file_information_order_proto_rawDescData = file_information_order_proto_rawDesc
)

func file_information_order_proto_rawDescGZIP() []byte {
	file_information_order_proto_rawDescOnce.Do(func() {
		file_information_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_information_order_proto_rawDescData)
	})
	return file_information_order_proto_rawDescData
}

var file_information_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_information_order_proto_goTypes = []interface{}{
	(*ValidateOrderRequest)(nil),    // 0: proto.information.order.ValidateOrderRequest
	(*ValidateProductsRequest)(nil), // 1: proto.information.order.ValidateProductsRequest
	(*ValidateResponse)(nil),        // 2: proto.information.order.ValidateResponse
	(*order.Order)(nil),             // 3: proto.order.order.Order
	(*order.OrderProduct)(nil),      // 4: proto.order.order.OrderProduct
}
var file_information_order_proto_depIdxs = []int32{
	3, // 0: proto.information.order.ValidateOrderRequest.order:type_name -> proto.order.order.Order
	4, // 1: proto.information.order.ValidateProductsRequest.order_products:type_name -> proto.order.order.OrderProduct
	0, // 2: proto.information.order.ValidateOrder.ValidateOrder:input_type -> proto.information.order.ValidateOrderRequest
	1, // 3: proto.information.order.ValidateOrder.ValidateProducts:input_type -> proto.information.order.ValidateProductsRequest
	2, // 4: proto.information.order.ValidateOrder.ValidateOrder:output_type -> proto.information.order.ValidateResponse
	2, // 5: proto.information.order.ValidateOrder.ValidateProducts:output_type -> proto.information.order.ValidateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_information_order_proto_init() }
func file_information_order_proto_init() {
	if File_information_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_information_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateOrderRequest); i {
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
		file_information_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateProductsRequest); i {
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
		file_information_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
			RawDescriptor: file_information_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_information_order_proto_goTypes,
		DependencyIndexes: file_information_order_proto_depIdxs,
		MessageInfos:      file_information_order_proto_msgTypes,
	}.Build()
	File_information_order_proto = out.File
	file_information_order_proto_rawDesc = nil
	file_information_order_proto_goTypes = nil
	file_information_order_proto_depIdxs = nil
}
