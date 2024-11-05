import { ListOption, Pager, Sorter, ConditionGroup, Concator, Condition} from "@/proto/common_pb";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { ListAccountReq } from "@/proto/main_pb";
import { serverSideClient } from "@/utils/proto/client";

export default async function Page() {

  // Initialize ListAccountReq with nested constructors
  const listReq = new ListAccountReq();

  try {
    // Execute gRPC request
    const listAccountResp = await serverSideClient.listAccount(listReq);
    const accountData = listAccountResp.List.map(account => ({
      ID: account.ID,
      UserName: account.UserName,
      DisplayName: account.DisplayName,
      Email: account.Email,
      State: account.State,
    }));

    // Render data table with results
    return (
      <div className="container mx-auto py-10">
        <DataTable columns={columns} data={accountData} />
      </div>
    );
  } catch (err) {
    console.error("Error fetching account list:", err);
    return <div>Error loading data. Please try again later.</div>;
  }
}
