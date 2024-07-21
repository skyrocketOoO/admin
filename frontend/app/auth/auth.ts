"use server"
import { z } from 'zod';
import { MainClient, LoginReq, LoginResp } from '@/proto/main';
import * as grpc from '@grpc/grpc-js';

async function Login(username: string, password: string): Promise<LoginResp | undefined> {
  const client = new MainClient("localhost:50051", grpc.credentials.createInsecure());
  const loginReq: LoginReq = {
    UserName: username,
    Password: password,
  };
  try {
    return new Promise((resolve, reject) => {
      client.login(loginReq, (error: grpc.ServiceError | null, resp: LoginResp) => {
        if (error) {
          console.error('Error:', error.message);
          reject(new Error('Failed to login.'));
        } else {
          resolve(resp);
        }
      });
    });
  }catch (err) {
    console.error('Failed to fetch user:', err);
    throw new Error('Failed to fetch user.');
  }
}

export async function signIn(username: string, password: string) {

  console.log('==================---')
  const resp = await Login(username, password);
  if (!resp) return null;
  return resp;

}