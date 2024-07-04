// Import necessary modules
import * as grpcWeb from 'grpc-web';
import { AccountClient } from './AccountServiceClientPb';  // Replace with the correct path to your generated client file
import { CreateAccountRequest, Empty } from './account_pb';    // Replace with the correct path to your generated message file

// Define your gRPC-Web client
const client = new AccountClient('http://localhost:50051', null, null);

// Prepare the request message
const request = new CreateAccountRequest();
request.setUsername('username');
request.setPassword('password');
request.setDisplayname('Display Name');

// Make the gRPC call
client.createAccount(request, {}, (err: grpcWeb.RpcError, response: Empty) => {
  if (err) {
    console.error('Error:', err.message);
    // Handle error
  } else {
    console.log('Response:', response.toObject());
    // Handle success
  }
});
