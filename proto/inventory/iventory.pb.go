// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/inventory/iventory.proto

package inventory

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

type ProductID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProductID) Reset() {
	*x = ProductID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_iventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductID) ProtoMessage() {}

func (x *ProductID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_iventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductID.ProtoReflect.Descriptor instead.
func (*ProductID) Descriptor() ([]byte, []int) {
	return file_proto_inventory_iventory_proto_rawDescGZIP(), []int{0}
}

func (x *ProductID) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Inventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Stock     int32 `protobuf:"varint,2,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *Inventory) Reset() {
	*x = Inventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_inventory_iventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inventory) ProtoMessage() {}

func (x *Inventory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_iventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inventory.ProtoReflect.Descriptor instead.
func (*Inventory) Descriptor() ([]byte, []int) {
	return file_proto_inventory_iventory_proto_rawDescGZIP(), []int{1}
}

func (x *Inventory) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Inventory) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

var File_proto_inventory_iventory_proto protoreflect.FileDescriptor

var file_proto_inventory_iventory_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x69, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x21, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40,
	0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x32, 0x83, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x14,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_inventory_iventory_proto_rawDescOnce sync.Once
	file_proto_inventory_iventory_proto_rawDescData = file_proto_inventory_iventory_proto_rawDesc
)

func file_proto_inventory_iventory_proto_rawDescGZIP() []byte {
	file_proto_inventory_iventory_proto_rawDescOnce.Do(func() {
		file_proto_inventory_iventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_inventory_iventory_proto_rawDescData)
	})
	return file_proto_inventory_iventory_proto_rawDescData
}

var file_proto_inventory_iventory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_inventory_iventory_proto_goTypes = []interface{}{
	(*ProductID)(nil), // 0: inventory.ProductID
	(*Inventory)(nil), // 1: inventory.Inventory
}
var file_proto_inventory_iventory_proto_depIdxs = []int32{
	0, // 0: inventory.InventoryInfo.GetStock:input_type -> inventory.ProductID
	1, // 1: inventory.InventoryInfo.AddStock:input_type -> inventory.Inventory
	1, // 2: inventory.InventoryInfo.GetStock:output_type -> inventory.Inventory
	1, // 3: inventory.InventoryInfo.AddStock:output_type -> inventory.Inventory
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_inventory_iventory_proto_init() }
func file_proto_inventory_iventory_proto_init() {
	if File_proto_inventory_iventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_inventory_iventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductID); i {
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
		file_proto_inventory_iventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inventory); i {
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
			RawDescriptor: file_proto_inventory_iventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inventory_iventory_proto_goTypes,
		DependencyIndexes: file_proto_inventory_iventory_proto_depIdxs,
		MessageInfos:      file_proto_inventory_iventory_proto_msgTypes,
	}.Build()
	File_proto_inventory_iventory_proto = out.File
	file_proto_inventory_iventory_proto_rawDesc = nil
	file_proto_inventory_iventory_proto_goTypes = nil
	file_proto_inventory_iventory_proto_depIdxs = nil
}
