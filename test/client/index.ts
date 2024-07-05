import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import { HelloRequest, HelloResponse } from './api_pb'; // Adjust path as necessary

const PROTO_PATH = __dirname + '/../proto/api.proto';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
});

const grpcObject = grpc.loadPackageDefinition(packageDefinition) as any; // Use 'any' if types are not directly available

const client = new grpcObject.api.MyService('localhost:50051', grpc.credentials.createInsecure());

const request = new HelloRequest();
request.setName('World');

client.sayHello(request, (error: grpc.ServiceError | null, response: HelloResponse) => {
  if (error) {
    console.error('Error:', error.message);
  } else {
    console.log('Greeting:', response.getMessage());
  }
});
