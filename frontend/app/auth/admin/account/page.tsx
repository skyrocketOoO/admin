import { ListOption, Pager, Sorter, Query } from "@/proto/common_pb";
import { columns } from "./columns";
import { DataTable } from "./data-table";
import { ListAccountReq } from "@/proto/main_pb";
import { serverSideClient } from "@/utils/proto/client";

export default async function Page() {
  // Set up the parameters
  const page = 1,
        size = 10,
        sortField = "UserName",
        sortOrder = "asc",
        query = "UserName";

  // Initialize ListAccountReq with nested constructors
  const listReq = new ListAccountReq({
    Option: new ListOption({
      Pager: new Pager({ Number: page, Size: size }),
      Sorter: new Sorter({ Asc: sortOrder === "asc", Field: sortField }),
      Query: query ? [new Query({ Fuzzy: true, Fields: ["UserName"], Value: query })] : []
    })
  });

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
