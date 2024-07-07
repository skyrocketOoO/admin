import * as grpc from '@grpc/grpc-js';
import { 
  CreateAccountRequest,
  ListRoleRequest ,
  ListAccountRequest,
  ListAccountResponse, 
  MainClient } from './proto/main';
import { Empty } from './proto/common';


const client = new MainClient("localhost:50051", grpc.credentials.createInsecure());

const createAccountReq: CreateAccountRequest = { 
  UserName: "admin", 
  Password: "admin", 
  DisplayName: "admin"
};
client.createAccount(createAccountReq, (error: grpc.ServiceError | null, response: Empty) => {
  if (error) {
    console.error('Error:', error.message);
  } else {
    // console.log('createAccount:', response);
  }
});

const listAccountReq: ListAccountRequest = {};
client.listAccount(listAccountReq, (error: grpc.ServiceError | null, response: ListAccountResponse) => {
  if (error) {
    console.error('Error:', error.message);
  } else {
    console.log('ListAccountResponse:', response);
  }
});

const UpdateAccountRequest = {

};
