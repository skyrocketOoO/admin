// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var api_pb = require('./api_pb.js');

function serialize_api_HelloRequest(arg) {
  if (!(arg instanceof api_pb.HelloRequest)) {
    throw new Error('Expected argument of type api.HelloRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_api_HelloRequest(buffer_arg) {
  return api_pb.HelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_api_HelloResponse(arg) {
  if (!(arg instanceof api_pb.HelloResponse)) {
    throw new Error('Expected argument of type api.HelloResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_api_HelloResponse(buffer_arg) {
  return api_pb.HelloResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var MyServiceService = exports.MyServiceService = {
  sayHello: {
    path: '/api.MyService/SayHello',
    requestStream: false,
    responseStream: false,
    requestType: api_pb.HelloRequest,
    responseType: api_pb.HelloResponse,
    requestSerialize: serialize_api_HelloRequest,
    requestDeserialize: deserialize_api_HelloRequest,
    responseSerialize: serialize_api_HelloResponse,
    responseDeserialize: deserialize_api_HelloResponse,
  },
};

exports.MyServiceClient = grpc.makeGenericClientConstructor(MyServiceService);
