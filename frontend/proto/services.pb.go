// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProductRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{0}
}

func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (m *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(m, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Product struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{1}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ProductList struct {
	Product              []*Product `protobuf:"bytes,1,rep,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProductList) Reset()         { *m = ProductList{} }
func (m *ProductList) String() string { return proto.CompactTextString(m) }
func (*ProductList) ProtoMessage()    {}
func (*ProductList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{2}
}

func (m *ProductList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductList.Unmarshal(m, b)
}
func (m *ProductList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductList.Marshal(b, m, deterministic)
}
func (m *ProductList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductList.Merge(m, src)
}
func (m *ProductList) XXX_Size() int {
	return xxx_messageInfo_ProductList.Size(m)
}
func (m *ProductList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductList.DiscardUnknown(m)
}

var xxx_messageInfo_ProductList proto.InternalMessageInfo

func (m *ProductList) GetProduct() []*Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type AddCartRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Cart                 *Cart    `protobuf:"bytes,2,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddCartRequest) Reset()         { *m = AddCartRequest{} }
func (m *AddCartRequest) String() string { return proto.CompactTextString(m) }
func (*AddCartRequest) ProtoMessage()    {}
func (*AddCartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{3}
}

func (m *AddCartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCartRequest.Unmarshal(m, b)
}
func (m *AddCartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCartRequest.Marshal(b, m, deterministic)
}
func (m *AddCartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCartRequest.Merge(m, src)
}
func (m *AddCartRequest) XXX_Size() int {
	return xxx_messageInfo_AddCartRequest.Size(m)
}
func (m *AddCartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddCartRequest proto.InternalMessageInfo

func (m *AddCartRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AddCartRequest) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type Response struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{5}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Cart struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Qty                  int32    `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{6}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cart) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

type CartList struct {
	Cart                 []*Cart  `protobuf:"bytes,1,rep,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartList) Reset()         { *m = CartList{} }
func (m *CartList) String() string { return proto.CompactTextString(m) }
func (*CartList) ProtoMessage()    {}
func (*CartList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{7}
}

func (m *CartList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartList.Unmarshal(m, b)
}
func (m *CartList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartList.Marshal(b, m, deterministic)
}
func (m *CartList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartList.Merge(m, src)
}
func (m *CartList) XXX_Size() int {
	return xxx_messageInfo_CartList.Size(m)
}
func (m *CartList) XXX_DiscardUnknown() {
	xxx_messageInfo_CartList.DiscardUnknown(m)
}

var xxx_messageInfo_CartList proto.InternalMessageInfo

func (m *CartList) GetCart() []*Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type CheckoutResponse struct {
	Status               bool          `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Cart                 []*Cart       `protobuf:"bytes,3,rep,name=cart,proto3" json:"cart,omitempty"`
	Shipping             *ShippingCost `protobuf:"bytes,4,opt,name=shipping,proto3" json:"shipping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CheckoutResponse) Reset()         { *m = CheckoutResponse{} }
func (m *CheckoutResponse) String() string { return proto.CompactTextString(m) }
func (*CheckoutResponse) ProtoMessage()    {}
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{8}
}

func (m *CheckoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckoutResponse.Unmarshal(m, b)
}
func (m *CheckoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckoutResponse.Marshal(b, m, deterministic)
}
func (m *CheckoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckoutResponse.Merge(m, src)
}
func (m *CheckoutResponse) XXX_Size() int {
	return xxx_messageInfo_CheckoutResponse.Size(m)
}
func (m *CheckoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckoutResponse proto.InternalMessageInfo

func (m *CheckoutResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *CheckoutResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CheckoutResponse) GetCart() []*Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

func (m *CheckoutResponse) GetShipping() *ShippingCost {
	if m != nil {
		return m.Shipping
	}
	return nil
}

type ShippingCostList struct {
	ShippingCost         *ShippingCost `protobuf:"bytes,1,opt,name=shippingCost,proto3" json:"shippingCost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ShippingCostList) Reset()         { *m = ShippingCostList{} }
func (m *ShippingCostList) String() string { return proto.CompactTextString(m) }
func (*ShippingCostList) ProtoMessage()    {}
func (*ShippingCostList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{9}
}

func (m *ShippingCostList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShippingCostList.Unmarshal(m, b)
}
func (m *ShippingCostList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShippingCostList.Marshal(b, m, deterministic)
}
func (m *ShippingCostList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShippingCostList.Merge(m, src)
}
func (m *ShippingCostList) XXX_Size() int {
	return xxx_messageInfo_ShippingCostList.Size(m)
}
func (m *ShippingCostList) XXX_DiscardUnknown() {
	xxx_messageInfo_ShippingCostList.DiscardUnknown(m)
}

var xxx_messageInfo_ShippingCostList proto.InternalMessageInfo

func (m *ShippingCostList) GetShippingCost() *ShippingCost {
	if m != nil {
		return m.ShippingCost
	}
	return nil
}

type ShippingCost struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShippingCost) Reset()         { *m = ShippingCost{} }
func (m *ShippingCost) String() string { return proto.CompactTextString(m) }
func (*ShippingCost) ProtoMessage()    {}
func (*ShippingCost) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{10}
}

func (m *ShippingCost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShippingCost.Unmarshal(m, b)
}
func (m *ShippingCost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShippingCost.Marshal(b, m, deterministic)
}
func (m *ShippingCost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShippingCost.Merge(m, src)
}
func (m *ShippingCost) XXX_Size() int {
	return xxx_messageInfo_ShippingCost.Size(m)
}
func (m *ShippingCost) XXX_DiscardUnknown() {
	xxx_messageInfo_ShippingCost.DiscardUnknown(m)
}

var xxx_messageInfo_ShippingCost proto.InternalMessageInfo

func (m *ShippingCost) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ShippingCost) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ShippingCost) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "proto.ProductRequest")
	proto.RegisterType((*Product)(nil), "proto.Product")
	proto.RegisterType((*ProductList)(nil), "proto.ProductList")
	proto.RegisterType((*AddCartRequest)(nil), "proto.AddCartRequest")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Cart)(nil), "proto.Cart")
	proto.RegisterType((*CartList)(nil), "proto.CartList")
	proto.RegisterType((*CheckoutResponse)(nil), "proto.CheckoutResponse")
	proto.RegisterType((*ShippingCostList)(nil), "proto.ShippingCostList")
	proto.RegisterType((*ShippingCost)(nil), "proto.ShippingCost")
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x4d, 0xda, 0x74, 0x13, 0x26, 0xab, 0x6c, 0x65, 0xa0, 0x1b, 0x95, 0x03, 0x95, 0x4f, 0x45,
	0xac, 0x52, 0x29, 0x7b, 0x58, 0x81, 0xb8, 0x40, 0x41, 0x05, 0xc1, 0x01, 0x79, 0xc5, 0x07, 0x64,
	0x13, 0x6f, 0x37, 0x82, 0xd4, 0xd9, 0xd8, 0x59, 0xa9, 0x5f, 0xc0, 0x81, 0x3b, 0x17, 0x7e, 0x16,
	0xd9, 0xb1, 0xd3, 0xa4, 0xb4, 0x17, 0x4e, 0xc9, 0xcc, 0xbc, 0x99, 0x79, 0x7e, 0x7e, 0x86, 0x80,
	0xd3, 0xea, 0x21, 0x4f, 0x29, 0x8f, 0xca, 0x8a, 0x09, 0x86, 0x46, 0xea, 0x33, 0x7d, 0xb6, 0x66,
	0x6c, 0xfd, 0x83, 0x2e, 0x54, 0x74, 0x53, 0xdf, 0x2e, 0x68, 0x51, 0x8a, 0x6d, 0x83, 0xc1, 0x33,
	0x08, 0xbe, 0x56, 0x2c, 0xab, 0x53, 0x41, 0xe8, 0x7d, 0x4d, 0xb9, 0x40, 0x01, 0x0c, 0xf2, 0x2c,
	0xb4, 0x67, 0xf6, 0x7c, 0x44, 0x06, 0x79, 0x86, 0x97, 0xe0, 0x6a, 0xc4, 0x7e, 0x09, 0x21, 0x70,
	0x36, 0x49, 0x41, 0xc3, 0xc1, 0xcc, 0x9e, 0x3f, 0x22, 0xea, 0x1f, 0x3d, 0x81, 0x51, 0x59, 0xe5,
	0x29, 0x0d, 0x87, 0x0a, 0xd6, 0x04, 0xf8, 0x0a, 0x7c, 0x3d, 0xe4, 0x4b, 0xce, 0x05, 0x9a, 0x83,
	0x5b, 0x36, 0x61, 0x68, 0xcf, 0x86, 0x73, 0x3f, 0x0e, 0x1a, 0x3a, 0x91, 0xe1, 0x62, 0xca, 0x98,
	0x40, 0xf0, 0x36, 0xcb, 0x96, 0x49, 0xd5, 0xf2, 0x7b, 0x0e, 0x4e, 0xcd, 0x69, 0xa5, 0x68, 0xf8,
	0xb1, 0xaf, 0x1b, 0xbf, 0x71, 0x5a, 0x11, 0x55, 0x90, 0x80, 0x34, 0xa9, 0x84, 0x62, 0xb5, 0x03,
	0xa8, 0x11, 0xaa, 0x80, 0xdf, 0x80, 0x47, 0x28, 0x2f, 0xd9, 0x86, 0x53, 0x34, 0x81, 0x13, 0x2e,
	0x12, 0x51, 0x73, 0x35, 0xcf, 0x23, 0x3a, 0x42, 0x21, 0xb8, 0x05, 0xe5, 0x3c, 0x59, 0x9b, 0xd3,
	0x99, 0x10, 0x4f, 0xc0, 0x91, 0xcb, 0xfe, 0xd1, 0xe9, 0x02, 0x1c, 0xb9, 0xa3, 0x15, 0xc5, 0xee,
	0x88, 0x32, 0x86, 0xe1, 0xbd, 0xd8, 0xaa, 0x49, 0x23, 0x22, 0x7f, 0xf1, 0x4b, 0xf0, 0x24, 0x5a,
	0xa9, 0x61, 0x08, 0x37, 0x52, 0x1c, 0x20, 0xfc, 0xdb, 0x86, 0xf1, 0xf2, 0x8e, 0xa6, 0xdf, 0x59,
	0x2d, 0xfe, 0x9f, 0x79, 0xbb, 0x67, 0x78, 0x64, 0x0f, 0x5a, 0x80, 0xc7, 0xef, 0xf2, 0xb2, 0xcc,
	0x37, 0xeb, 0xd0, 0x51, 0xea, 0x3d, 0xd6, 0xa0, 0x6b, 0x9d, 0x5e, 0x32, 0x2e, 0x48, 0x0b, 0xc2,
	0x9f, 0x61, 0xdc, 0xad, 0xa8, 0xd3, 0x5c, 0xc1, 0x29, 0xef, 0xe4, 0xf4, 0x3d, 0x1d, 0x1c, 0xd4,
	0x03, 0xe2, 0x8f, 0x70, 0xda, 0xad, 0x4a, 0x21, 0x6f, 0x2b, 0x56, 0x18, 0x21, 0xe5, 0xbf, 0x14,
	0x5d, 0x30, 0x7d, 0xae, 0x81, 0x60, 0x87, 0xdd, 0x16, 0xff, 0xb4, 0x5b, 0x57, 0x5f, 0x37, 0x4f,
	0x02, 0xbd, 0x82, 0x60, 0x45, 0x85, 0x4e, 0xbe, 0xdb, 0x7e, 0xca, 0xd0, 0xd3, 0x3d, 0xcb, 0x35,
	0xf6, 0x9a, 0xee, 0x39, 0x11, 0x5b, 0xe8, 0x35, 0xc0, 0xae, 0x15, 0x4d, 0xa2, 0xe6, 0x39, 0x45,
	0xe6, 0x39, 0x45, 0x1f, 0xe4, 0x73, 0x9a, 0xa2, 0x7e, 0x9f, 0x94, 0x02, 0x5b, 0xf1, 0x1f, 0x1b,
	0x7c, 0x29, 0xb0, 0xa1, 0x71, 0x09, 0xae, 0xb6, 0x73, 0xbb, 0xbf, 0x6f, 0xef, 0xe9, 0x99, 0x4e,
	0x9b, 0x7b, 0xc6, 0x16, 0x7a, 0x01, 0xee, 0x8a, 0x0a, 0xd5, 0xd4, 0xb5, 0x7b, 0x0b, 0x35, 0x46,
	0xc2, 0x16, 0xba, 0x00, 0x20, 0xb4, 0x60, 0x0f, 0xf4, 0x38, 0x7a, 0x37, 0x38, 0xfe, 0x65, 0xc3,
	0x99, 0xf1, 0x95, 0x61, 0x18, 0x83, 0x67, 0x52, 0xfd, 0xfe, 0x73, 0xb3, 0x6d, 0xcf, 0x88, 0xd8,
	0x42, 0xef, 0xe1, 0x6c, 0x45, 0x45, 0xef, 0xf2, 0x8e, 0xc9, 0x74, 0x7e, 0xc0, 0x07, 0x0d, 0xf7,
	0x9b, 0x13, 0x55, 0xb9, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x56, 0xb2, 0xa2, 0x77, 0xc7, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProductById(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error)
	GetProduct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProductList, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProductById(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/proto.ProductService/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/proto.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProductById(context.Context, *ProductRequest) (*Product, error)
	GetProduct(context.Context, *empty.Empty) (*ProductList, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProductById(ctx context.Context, req *ProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (*UnimplementedProductServiceServer) GetProduct(ctx context.Context, req *empty.Empty) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductById(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductById",
			Handler:    _ProductService_GetProductById_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartServiceClient interface {
	AddCart(ctx context.Context, in *AddCartRequest, opts ...grpc.CallOption) (*Response, error)
	GetCart(ctx context.Context, in *User, opts ...grpc.CallOption) (*CartList, error)
	RemoveCart(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
}

type cartServiceClient struct {
	cc *grpc.ClientConn
}

func NewCartServiceClient(cc *grpc.ClientConn) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddCart(ctx context.Context, in *AddCartRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CartService/AddCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *User, opts ...grpc.CallOption) (*CartList, error) {
	out := new(CartList)
	err := c.cc.Invoke(ctx, "/proto.CartService/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveCart(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CartService/RemoveCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
type CartServiceServer interface {
	AddCart(context.Context, *AddCartRequest) (*Response, error)
	GetCart(context.Context, *User) (*CartList, error)
	RemoveCart(context.Context, *User) (*Response, error)
}

// UnimplementedCartServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (*UnimplementedCartServiceServer) AddCart(ctx context.Context, req *AddCartRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCart not implemented")
}
func (*UnimplementedCartServiceServer) GetCart(ctx context.Context, req *User) (*CartList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (*UnimplementedCartServiceServer) RemoveCart(ctx context.Context, req *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCart not implemented")
}

func RegisterCartServiceServer(s *grpc.Server, srv CartServiceServer) {
	s.RegisterService(&_CartService_serviceDesc, srv)
}

func _CartService_AddCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/AddCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddCart(ctx, req.(*AddCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CartService/RemoveCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveCart(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCart",
			Handler:    _CartService_AddCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "RemoveCart",
			Handler:    _CartService_RemoveCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// CheckoutServiceClient is the client API for CheckoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckoutServiceClient interface {
	Checkout(ctx context.Context, in *User, opts ...grpc.CallOption) (*CheckoutResponse, error)
	GetShippingCost(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShippingCostList, error)
}

type checkoutServiceClient struct {
	cc *grpc.ClientConn
}

func NewCheckoutServiceClient(cc *grpc.ClientConn) CheckoutServiceClient {
	return &checkoutServiceClient{cc}
}

func (c *checkoutServiceClient) Checkout(ctx context.Context, in *User, opts ...grpc.CallOption) (*CheckoutResponse, error) {
	out := new(CheckoutResponse)
	err := c.cc.Invoke(ctx, "/proto.CheckoutService/Checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) GetShippingCost(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShippingCostList, error) {
	out := new(ShippingCostList)
	err := c.cc.Invoke(ctx, "/proto.CheckoutService/GetShippingCost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckoutServiceServer is the server API for CheckoutService service.
type CheckoutServiceServer interface {
	Checkout(context.Context, *User) (*CheckoutResponse, error)
	GetShippingCost(context.Context, *empty.Empty) (*ShippingCostList, error)
}

// UnimplementedCheckoutServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCheckoutServiceServer struct {
}

func (*UnimplementedCheckoutServiceServer) Checkout(ctx context.Context, req *User) (*CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (*UnimplementedCheckoutServiceServer) GetShippingCost(ctx context.Context, req *empty.Empty) (*ShippingCostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShippingCost not implemented")
}

func RegisterCheckoutServiceServer(s *grpc.Server, srv CheckoutServiceServer) {
	s.RegisterService(&_CheckoutService_serviceDesc, srv)
}

func _CheckoutService_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CheckoutService/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).Checkout(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_GetShippingCost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).GetShippingCost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CheckoutService/GetShippingCost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).GetShippingCost(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckoutService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CheckoutService",
	HandlerType: (*CheckoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Checkout",
			Handler:    _CheckoutService_Checkout_Handler,
		},
		{
			MethodName: "GetShippingCost",
			Handler:    _CheckoutService_GetShippingCost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}