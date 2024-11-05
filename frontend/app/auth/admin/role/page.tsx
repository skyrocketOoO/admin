'use client';
import { useEffect, useState } from "react";
import { cn } from "@/lib/utils";
import { serverSideClient } from "@/utils/proto/client";
import { ListAccountReq } from "@/proto/main_pb";
import { ListOption, Pager, Sorter, ConditionGroup, Condition } from "@/proto/common_pb";

// Define the type for account data
interface AccountData {
  ID: number;
  UserName: string;
  DisplayName: string;
  Email: string;
  State: number;
}


export default function Page() {
  const columns = ["ID", "UserName", "DisplayName", "Email", "State"];
  const pageSize = 10;

  // Specify the type for accountData as AccountData[]
  const [accountData, setAccountData] = useState<AccountData[]>([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [filter, setFilter] = useState("");

  useEffect(() => {
    const fetchAccountData = async () => {
      setLoading(true);

      const
        sortField = "UserName",
        sortOrder = "asc",
        query = "UserName";

      const listReq = new ListAccountReq({
        Option: new ListOption({
          pager: new Pager({ number: page, size: pageSize }),
          sorters: [new Sorter({ ascending: sortOrder === "asc", field: sortField })],
          // conditionGroup: 
        })
      });

      try {
        const listAccountResp = await serverSideClient.listAccount(listReq);
        const formattedData: AccountData[] = listAccountResp.List.map(account => ({
          ID: account.ID,
          UserName: account.UserName,
          DisplayName: account.DisplayName,
          Email: account.Email,
          State: account.State,
        }));

        setAccountData(formattedData);
        setTotalPages(Math.ceil(Number(listAccountResp.Total)/ pageSize));
      } catch (error) {
        console.error("Error fetching account data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchAccountData();
  }, [page, filter]);

  const handleNextPage = () => {
    if (page < totalPages) {
      setPage(page + 1);
    }
  };

  const handlePreviousPage = () => {
    if (page > 1) {
      setPage(page - 1);
    }
  };

  return (
    <div>
      <div className="flex items-center py-4">
        <input 
          className="max-w-sm" 
          placeholder="Filter all..." 
          value={filter} 
          onChange={(e) => setFilter(e.target.value)} 
        />
      </div>
      <div className="rounded-md border">
        {loading ? (
          <p>Loading...</p>
        ) : (
          <table className={cn("w-full caption-bottom text-sm")}>
            <thead className={cn("[&_tr]:border-b")}>
              <tr className={cn("border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted")}>
                {columns.map((column) => (
                  <th
                    key={column}
                    className={cn("h-10 px-2 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]")}
                  >
                    {column}
                  </th>
                ))}
              </tr>
            </thead>
            <tbody className={cn("[&_tr:last-child]:border-0")}>
              {accountData.map((row) => (
                <tr
                  key={row.ID}
                  className={cn("border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted")}
                >
                  <td className={cn("p-2 align-middle")}>{row.ID}</td>
                  <td className={cn("p-2 align-middle")}>{row.UserName}</td>
                  <td className={cn("p-2 align-middle")}>{row.DisplayName}</td>
                  <td className={cn("p-2 align-middle")}>{row.Email}</td>
                  <td className={cn("p-2 align-middle")}>{row.State}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      <div className="flex items-center justify-center space-x-2 py-4">
        <button
          onClick={handlePreviousPage}
          disabled={page === 1}
          className="px-4 py-2 bg-gray-200 rounded disabled:opacity-50"
        >
          Previous
        </button>
        <span>
          Page {page} of {totalPages}
        </span>
        <button
          onClick={handleNextPage}
          disabled={page === totalPages}
          className="px-4 py-2 bg-gray-200 rounded disabled:opacity-50"
        >
          Next
        </button>
      </div>
    </div>
  );
}
