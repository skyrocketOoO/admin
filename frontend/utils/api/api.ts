import { MainClient, ListAccountReq, ListAccountResp } from '@/proto/main';
import * as grpc from '@grpc/grpc-js';

const client = new MainClient("localhost:50051", grpc.credentials.createInsecure());

export async function listAccount(listReq: ListAccountReq): Promise<ListAccountResp> {
  return new Promise((resolve, reject) => {
    client.listAccount(listReq, (error, response) => {
      if (error) {
        return reject(error); 
      }
      resolve(response);
    });
  });
}