# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: classification/classification.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#classification/classification.proto\"\x1e\n\x0fRequestClassImg\x12\x0b\n\x03img\x18\x01 \x01(\x0c\"\x1e\n\x10ResponseClassImg\x12\n\n\x02id\x18\x01 \x01(\r2B\n\x0f\x43lassImgService\x12/\n\x08\x43lassImg\x12\x10.RequestClassImg\x1a\x11.ResponseClassImgB\x11Z\x0f\x63lassification/b\x06proto3')



_REQUESTCLASSIMG = DESCRIPTOR.message_types_by_name['RequestClassImg']
_RESPONSECLASSIMG = DESCRIPTOR.message_types_by_name['ResponseClassImg']
RequestClassImg = _reflection.GeneratedProtocolMessageType('RequestClassImg', (_message.Message,), {
  'DESCRIPTOR' : _REQUESTCLASSIMG,
  '__module__' : 'classification.classification_pb2'
  # @@protoc_insertion_point(class_scope:RequestClassImg)
  })
_sym_db.RegisterMessage(RequestClassImg)

ResponseClassImg = _reflection.GeneratedProtocolMessageType('ResponseClassImg', (_message.Message,), {
  'DESCRIPTOR' : _RESPONSECLASSIMG,
  '__module__' : 'classification.classification_pb2'
  # @@protoc_insertion_point(class_scope:ResponseClassImg)
  })
_sym_db.RegisterMessage(ResponseClassImg)

_CLASSIMGSERVICE = DESCRIPTOR.services_by_name['ClassImgService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\017classification/'
  _REQUESTCLASSIMG._serialized_start=39
  _REQUESTCLASSIMG._serialized_end=69
  _RESPONSECLASSIMG._serialized_start=71
  _RESPONSECLASSIMG._serialized_end=101
  _CLASSIMGSERVICE._serialized_start=103
  _CLASSIMGSERVICE._serialized_end=169
# @@protoc_insertion_point(module_scope)
