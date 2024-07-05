// package: api
// file: api.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as api_pb from "./api_pb";

interface IMyServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    sayHello: IMyServiceService_ISayHello;
}

interface IMyServiceService_ISayHello extends grpc.MethodDefinition<api_pb.HelloRequest, api_pb.HelloResponse> {
    path: "/api.MyService/SayHello";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<api_pb.HelloRequest>;
    requestDeserialize: grpc.deserialize<api_pb.HelloRequest>;
    responseSerialize: grpc.serialize<api_pb.HelloResponse>;
    responseDeserialize: grpc.deserialize<api_pb.HelloResponse>;
}

export const MyServiceService: IMyServiceService;

export interface IMyServiceServer extends grpc.UntypedServiceImplementation {
    sayHello: grpc.handleUnaryCall<api_pb.HelloRequest, api_pb.HelloResponse>;
}

export interface IMyServiceClient {
    sayHello(request: api_pb.HelloRequest, callback: (error: grpc.ServiceError | null, response: api_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    sayHello(request: api_pb.HelloRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: api_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    sayHello(request: api_pb.HelloRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: api_pb.HelloResponse) => void): grpc.ClientUnaryCall;
}

export class MyServiceClient extends grpc.Client implements IMyServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public sayHello(request: api_pb.HelloRequest, callback: (error: grpc.ServiceError | null, response: api_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    public sayHello(request: api_pb.HelloRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: api_pb.HelloResponse) => void): grpc.ClientUnaryCall;
    public sayHello(request: api_pb.HelloRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: api_pb.HelloResponse) => void): grpc.ClientUnaryCall;
}
