'use server'
import { MainClient, AccountData, ListAccountReq, ListAccountResp } from '@/proto/main';
import * as grpc from '@grpc/grpc-js';

const client = new MainClient("localhost:50051", grpc.credentials.createInsecure());

export async function listAccount(req: ListAccountReq): Promise<ListAccountResp | null> {
  try {
    const resp = await new Promise<ListAccountResp>((resolve, reject) => {
      client.listAccount(req, (error: grpc.ServiceError | null, resp: ListAccountResp) => {
        if (error) {
          reject(new Error('Failed to fetch accounts.'));
        } else {
          resolve(resp);
        }
      });
    });
    return resp;
  } catch (err) {
    console.error('Error fetching accounts:', err);
    return null;
  }
}