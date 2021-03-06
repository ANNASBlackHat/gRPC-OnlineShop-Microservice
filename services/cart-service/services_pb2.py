# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='services.proto',
  package='proto',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0eservices.proto\x12\x05proto\x1a\x1bgoogle/protobuf/empty.proto\"\x1c\n\x0eProductRequest\x12\n\n\x02id\x18\x01 \x01(\x05\"2\n\x07Product\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05price\x18\x03 \x01(\x05\".\n\x0bProductList\x12\x1f\n\x07product\x18\x01 \x03(\x0b\x32\x0e.proto.Product\"F\n\x0e\x41\x64\x64\x43\x61rtRequest\x12\x19\n\x04user\x18\x01 \x01(\x0b\x32\x0b.proto.User\x12\x19\n\x04\x63\x61rt\x18\x02 \x01(\x0b\x32\x0b.proto.Cart\"+\n\x08Response\x12\x0e\n\x06status\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t\"\x12\n\x04User\x12\n\n\x02id\x18\x01 \x01(\x05\"!\n\x04\x43\x61rt\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0b\n\x03qty\x18\x02 \x01(\x05\"%\n\x08\x43\x61rtList\x12\x19\n\x04\x63\x61rt\x18\x01 \x03(\x0b\x32\x0b.proto.Cart\"u\n\x10\x43heckoutResponse\x12\x0e\n\x06status\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\x19\n\x04\x63\x61rt\x18\x03 \x03(\x0b\x32\x0b.proto.Cart\x12%\n\x08shipping\x18\x04 \x01(\x0b\x32\x13.proto.ShippingCost\"=\n\x10ShippingCostList\x12)\n\x0cshippingCost\x18\x01 \x01(\x0b\x32\x13.proto.ShippingCost\"7\n\x0cShippingCost\x12\x0c\n\x04\x66rom\x18\x01 \x01(\t\x12\n\n\x02to\x18\x02 \x01(\t\x12\r\n\x05price\x18\x03 \x01(\x05\x32\x87\x01\n\x0eProductService\x12\x39\n\x0eGetProductById\x12\x15.proto.ProductRequest\x1a\x0e.proto.Product\"\x00\x12:\n\nGetProduct\x12\x16.google.protobuf.Empty\x1a\x12.proto.ProductList\"\x00\x32\x9b\x01\n\x0b\x43\x61rtService\x12\x33\n\x07\x41\x64\x64\x43\x61rt\x12\x15.proto.AddCartRequest\x1a\x0f.proto.Response\"\x00\x12)\n\x07GetCart\x12\x0b.proto.User\x1a\x0f.proto.CartList\"\x00\x12,\n\nRemoveCart\x12\x0b.proto.User\x1a\x0f.proto.Response\"\x00\x32\x8b\x01\n\x0f\x43heckoutService\x12\x32\n\x08\x43heckout\x12\x0b.proto.User\x1a\x17.proto.CheckoutResponse\"\x00\x12\x44\n\x0fGetShippingCost\x12\x16.google.protobuf.Empty\x1a\x17.proto.ShippingCostList\"\x00\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])




_PRODUCTREQUEST = _descriptor.Descriptor(
  name='ProductRequest',
  full_name='proto.ProductRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='proto.ProductRequest.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=54,
  serialized_end=82,
)


_PRODUCT = _descriptor.Descriptor(
  name='Product',
  full_name='proto.Product',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='proto.Product.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='proto.Product.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='proto.Product.price', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=84,
  serialized_end=134,
)


_PRODUCTLIST = _descriptor.Descriptor(
  name='ProductList',
  full_name='proto.ProductList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='product', full_name='proto.ProductList.product', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=136,
  serialized_end=182,
)


_ADDCARTREQUEST = _descriptor.Descriptor(
  name='AddCartRequest',
  full_name='proto.AddCartRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='proto.AddCartRequest.user', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cart', full_name='proto.AddCartRequest.cart', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=184,
  serialized_end=254,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='proto.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='proto.Response.status', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='proto.Response.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=256,
  serialized_end=299,
)


_USER = _descriptor.Descriptor(
  name='User',
  full_name='proto.User',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='proto.User.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=301,
  serialized_end=319,
)


_CART = _descriptor.Descriptor(
  name='Cart',
  full_name='proto.Cart',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='proto.Cart.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='qty', full_name='proto.Cart.qty', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=321,
  serialized_end=354,
)


_CARTLIST = _descriptor.Descriptor(
  name='CartList',
  full_name='proto.CartList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cart', full_name='proto.CartList.cart', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=356,
  serialized_end=393,
)


_CHECKOUTRESPONSE = _descriptor.Descriptor(
  name='CheckoutResponse',
  full_name='proto.CheckoutResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='proto.CheckoutResponse.status', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='proto.CheckoutResponse.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cart', full_name='proto.CheckoutResponse.cart', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='shipping', full_name='proto.CheckoutResponse.shipping', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=395,
  serialized_end=512,
)


_SHIPPINGCOSTLIST = _descriptor.Descriptor(
  name='ShippingCostList',
  full_name='proto.ShippingCostList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='shippingCost', full_name='proto.ShippingCostList.shippingCost', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=514,
  serialized_end=575,
)


_SHIPPINGCOST = _descriptor.Descriptor(
  name='ShippingCost',
  full_name='proto.ShippingCost',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='from', full_name='proto.ShippingCost.from', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='to', full_name='proto.ShippingCost.to', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='price', full_name='proto.ShippingCost.price', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=577,
  serialized_end=632,
)

_PRODUCTLIST.fields_by_name['product'].message_type = _PRODUCT
_ADDCARTREQUEST.fields_by_name['user'].message_type = _USER
_ADDCARTREQUEST.fields_by_name['cart'].message_type = _CART
_CARTLIST.fields_by_name['cart'].message_type = _CART
_CHECKOUTRESPONSE.fields_by_name['cart'].message_type = _CART
_CHECKOUTRESPONSE.fields_by_name['shipping'].message_type = _SHIPPINGCOST
_SHIPPINGCOSTLIST.fields_by_name['shippingCost'].message_type = _SHIPPINGCOST
DESCRIPTOR.message_types_by_name['ProductRequest'] = _PRODUCTREQUEST
DESCRIPTOR.message_types_by_name['Product'] = _PRODUCT
DESCRIPTOR.message_types_by_name['ProductList'] = _PRODUCTLIST
DESCRIPTOR.message_types_by_name['AddCartRequest'] = _ADDCARTREQUEST
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
DESCRIPTOR.message_types_by_name['User'] = _USER
DESCRIPTOR.message_types_by_name['Cart'] = _CART
DESCRIPTOR.message_types_by_name['CartList'] = _CARTLIST
DESCRIPTOR.message_types_by_name['CheckoutResponse'] = _CHECKOUTRESPONSE
DESCRIPTOR.message_types_by_name['ShippingCostList'] = _SHIPPINGCOSTLIST
DESCRIPTOR.message_types_by_name['ShippingCost'] = _SHIPPINGCOST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProductRequest = _reflection.GeneratedProtocolMessageType('ProductRequest', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTREQUEST,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.ProductRequest)
  })
_sym_db.RegisterMessage(ProductRequest)

Product = _reflection.GeneratedProtocolMessageType('Product', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCT,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.Product)
  })
_sym_db.RegisterMessage(Product)

ProductList = _reflection.GeneratedProtocolMessageType('ProductList', (_message.Message,), {
  'DESCRIPTOR' : _PRODUCTLIST,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.ProductList)
  })
_sym_db.RegisterMessage(ProductList)

AddCartRequest = _reflection.GeneratedProtocolMessageType('AddCartRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDCARTREQUEST,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.AddCartRequest)
  })
_sym_db.RegisterMessage(AddCartRequest)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
  'DESCRIPTOR' : _RESPONSE,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.Response)
  })
_sym_db.RegisterMessage(Response)

User = _reflection.GeneratedProtocolMessageType('User', (_message.Message,), {
  'DESCRIPTOR' : _USER,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.User)
  })
_sym_db.RegisterMessage(User)

Cart = _reflection.GeneratedProtocolMessageType('Cart', (_message.Message,), {
  'DESCRIPTOR' : _CART,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.Cart)
  })
_sym_db.RegisterMessage(Cart)

CartList = _reflection.GeneratedProtocolMessageType('CartList', (_message.Message,), {
  'DESCRIPTOR' : _CARTLIST,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.CartList)
  })
_sym_db.RegisterMessage(CartList)

CheckoutResponse = _reflection.GeneratedProtocolMessageType('CheckoutResponse', (_message.Message,), {
  'DESCRIPTOR' : _CHECKOUTRESPONSE,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.CheckoutResponse)
  })
_sym_db.RegisterMessage(CheckoutResponse)

ShippingCostList = _reflection.GeneratedProtocolMessageType('ShippingCostList', (_message.Message,), {
  'DESCRIPTOR' : _SHIPPINGCOSTLIST,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.ShippingCostList)
  })
_sym_db.RegisterMessage(ShippingCostList)

ShippingCost = _reflection.GeneratedProtocolMessageType('ShippingCost', (_message.Message,), {
  'DESCRIPTOR' : _SHIPPINGCOST,
  '__module__' : 'services_pb2'
  # @@protoc_insertion_point(class_scope:proto.ShippingCost)
  })
_sym_db.RegisterMessage(ShippingCost)



_PRODUCTSERVICE = _descriptor.ServiceDescriptor(
  name='ProductService',
  full_name='proto.ProductService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=635,
  serialized_end=770,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetProductById',
    full_name='proto.ProductService.GetProductById',
    index=0,
    containing_service=None,
    input_type=_PRODUCTREQUEST,
    output_type=_PRODUCT,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetProduct',
    full_name='proto.ProductService.GetProduct',
    index=1,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_PRODUCTLIST,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PRODUCTSERVICE)

DESCRIPTOR.services_by_name['ProductService'] = _PRODUCTSERVICE


_CARTSERVICE = _descriptor.ServiceDescriptor(
  name='CartService',
  full_name='proto.CartService',
  file=DESCRIPTOR,
  index=1,
  serialized_options=None,
  serialized_start=773,
  serialized_end=928,
  methods=[
  _descriptor.MethodDescriptor(
    name='AddCart',
    full_name='proto.CartService.AddCart',
    index=0,
    containing_service=None,
    input_type=_ADDCARTREQUEST,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetCart',
    full_name='proto.CartService.GetCart',
    index=1,
    containing_service=None,
    input_type=_USER,
    output_type=_CARTLIST,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RemoveCart',
    full_name='proto.CartService.RemoveCart',
    index=2,
    containing_service=None,
    input_type=_USER,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CARTSERVICE)

DESCRIPTOR.services_by_name['CartService'] = _CARTSERVICE


_CHECKOUTSERVICE = _descriptor.ServiceDescriptor(
  name='CheckoutService',
  full_name='proto.CheckoutService',
  file=DESCRIPTOR,
  index=2,
  serialized_options=None,
  serialized_start=931,
  serialized_end=1070,
  methods=[
  _descriptor.MethodDescriptor(
    name='Checkout',
    full_name='proto.CheckoutService.Checkout',
    index=0,
    containing_service=None,
    input_type=_USER,
    output_type=_CHECKOUTRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetShippingCost',
    full_name='proto.CheckoutService.GetShippingCost',
    index=1,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_SHIPPINGCOSTLIST,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CHECKOUTSERVICE)

DESCRIPTOR.services_by_name['CheckoutService'] = _CHECKOUTSERVICE

# @@protoc_insertion_point(module_scope)
