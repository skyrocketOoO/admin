import { columns } from "./columns"
import { DataTable } from "./data-table"
import { ListAccountReq, ListAccountResp } from '@/proto/main'; // Adjust the import according to your project structure
import { MainClient } from '@/proto/main';
import * as grpc from '@grpc/grpc-js';

async function getData(listReq: ListAccountReq): Promise<ListAccountResp> {
  const client = new MainClient('localhost:50051', grpc.credentials.createInsecure());

  // Simulate fetching account data from an API
  return new Promise((resolve, reject) => {
    client.listAccount(listReq, (error, response) => {
      if (error) {
        return reject(error);  // Reject the promise if an error occurs
      }
      resolve(response);  // Resolve the promise with the response data
    });
  });
}


export default async function DemoPage() {
  const  page = 1, size = 10, sortField = "UserName", sortOrder = "asc", query = "UserName";
  const listReq: ListAccountReq = {
    Option: {
      Pager: { Number: Number(page), Size: Number(size) },
      Sorter: { Asc: sortOrder === 'asc', Field: sortField as string },
      Query: query ? [{ Fuzzy: true, Fields: ['UserName'], Value: query as string }] : [],
    },
  };
  const listAccountResp = await getData(listReq);
  console.log(listAccountResp)
  return (
    <div className="container mx-auto py-10">
      <DataTable columns={columns} data={listAccountResp.List} />
    </div>
  )
}
