// pages/api/accounts.ts
import { NextApiRequest, NextApiResponse } from 'next';
import { ListAccountReq, ListAccountResp } from '@/proto/main'; // Adjust the import according to your project structure
import { MainClient } from '@/proto/main';
import * as grpc from '@grpc/grpc-js';

const client = new MainClient('localhost:50051', grpc.credentials.createInsecure());

export default async (req: NextApiRequest, res: NextApiResponse) => {
  const { page = 1, size = 10, sortField, sortOrder, query } = req.query;

  const listReq: ListAccountReq = {
    Option: {
      Pager: { Number: Number(page), Size: Number(size) },
      Sorter: { Asc: sortOrder === 'asc', Field: sortField as string },
      Query: query ? [{ Fuzzy: true, Field: 'UserName', Value: query as string }] : [],
    },
  };

  try {
    const resp = await new Promise<ListAccountResp>((resolve, reject) => {
      client.listAccount(listReq, (error, response) => {
        if (error) {
          return reject(error);
        }
        resolve(response);
      });
    });

    res.status(200).json(resp);
  } catch (error) {
    res.status(500).json({ error: 'Failed to fetch accounts' });
  }
};
