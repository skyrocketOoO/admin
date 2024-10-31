import { columns } from "./columns";
import { DataTable } from "./data-table";
import { ListAccountReq } from "@/proto/main_pb";
import { listAccount, serverSideClient } from "@/utils/api/api";

export default async function Page() {
  const page = 1,
        size = 10,
        sortField = "UserName",
        sortOrder = "asc",
        query = "UserName";
  
  const listReq: ListAccountReq = {
    Option: {
      Pager: { Number: page, Size: size },
      Sorter: { Asc: sortOrder === 'asc', Field: sortField },
      Query: query ? [{ Fuzzy: true, Fields: ['UserName'], Value: query }] : [],
    },
  };

  try {
    const listAccountResp = await listAccount(serverSideClient, listReq);
    return (
      <div className="container mx-auto py-10">
        <DataTable columns={columns} data={listAccountResp.List} />
      </div>
    );
  } catch (err) {
    console.error("Error fetching account list:", err);
    return <div>Error loading data. Please try again later.</div>;
  }
}

import { MainClient } from '@/proto/main';
import { grpc } from '@improbable-eng/grpc-web';

const client = new MainClient("localhost:50051", grpc);

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
