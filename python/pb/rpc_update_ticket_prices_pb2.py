# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: rpc_update_ticket_prices.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'rpc_update_ticket_prices.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1erpc_update_ticket_prices.proto\x12\nprivate_pb\"-\n\x19UpdateTicketPricesRequest\x12\x10\n\x08match_id\x18\x01 \x01(\x05\"=\n\x1aUpdateTicketPricesResponse\x12\x0e\n\x06status\x18\x01 \x01(\x08\x12\x0f\n\x07message\x18\x02 \x01(\t2\x82\x01\n\x1bPrivatePremierLeagueBooking\x12\x63\n\x12UpdateTicketPrices\x12%.private_pb.UpdateTicketPricesRequest\x1a&.private_pb.UpdateTicketPricesResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'rpc_update_ticket_prices_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_UPDATETICKETPRICESREQUEST']._serialized_start=46
  _globals['_UPDATETICKETPRICESREQUEST']._serialized_end=91
  _globals['_UPDATETICKETPRICESRESPONSE']._serialized_start=93
  _globals['_UPDATETICKETPRICESRESPONSE']._serialized_end=154
  _globals['_PRIVATEPREMIERLEAGUEBOOKING']._serialized_start=157
  _globals['_PRIVATEPREMIERLEAGUEBOOKING']._serialized_end=287
# @@protoc_insertion_point(module_scope)
