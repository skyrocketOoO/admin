'use client';
import { useEffect, useState } from "react";
import { cn } from "@/lib/utils";
import { serverSideClient } from "@/utils/proto/client";
import { ListAccountReq } from "@/proto/main_pb";
import { ListOption, Pager, Sorter, ConditionGroup, Condition, Concator } from "@/proto/common_pb";

interface AccountData {
  ID: number;
  UserName: string;
  DisplayName: string;
  State: number;
}

export default function Page() {
  const columns = ["ID", "UserName", "DisplayName", "State"];
  const pageSize = 10;

  const [accountData, setAccountData] = useState<AccountData[]>([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [filter, setFilter] = useState("");
  const [debouncedFilter, setDebouncedFilter] = useState(filter);
  // Sorting state
  const [sortedData, setSortedData] = useState<AccountData[]>([]);
  const [sortColumn, setSortColumn] = useState<keyof AccountData | null>(null);
  const [sortDirection, setSortDirection] = useState<"asc" | "desc">("asc");

  // debounce
  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedFilter(filter);
    }, 800);
    return () => clearTimeout(handler);
  }, [filter]);

  useEffect(() => {
    const fetchAccountData = async () => {
      setLoading(true);

      const listReq = new ListAccountReq({
        Option: new ListOption({
          pager: new Pager({ number: page, size: pageSize }),
        })
      });

      if (debouncedFilter != ""){
        listReq.Option = listReq.Option || new ListOption();
        listReq.Option.conditionGroup = new ConditionGroup({
          concator: Concator.OR,
          conditions: columns.map((column) => new Condition({
            field: column,
            operator: "LIKE",
            value: "%"+debouncedFilter+"%"
          })),
        });        
      }
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
        setSortedData(formattedData);
        setTotalPages(Math.ceil(Number(listAccountResp.Total)/ pageSize));
      } catch (error) {
        console.error("Error fetching account data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchAccountData();
  }, [page, debouncedFilter]);

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

  // Sort data based on current column and direction
  useEffect(() => {
    const sorted = [...accountData].sort((a, b) => {
      if (!sortColumn) return 0;
      const valueA = a[sortColumn];
      const valueB = b[sortColumn];

      if (valueA < valueB) return sortDirection === "asc" ? -1 : 1;
      if (valueA > valueB) return sortDirection === "asc" ? 1 : -1;
      return 0;
    });
    setSortedData(sorted);
  }, [sortColumn, sortDirection, accountData]);

  // Handle sort toggle
  const handleSort = (column: keyof AccountData) => {
    if (sortColumn === column) {
      setSortDirection((prev) => (prev === "asc" ? "desc" : "asc"));
    } else {
      setSortColumn(column);
      setSortDirection("asc");
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
                    onClick={() => handleSort(column as keyof AccountData)}
                    className={cn("h-10 px-2 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]")}
                  >
                    {column}
                  </th>
                ))}
              </tr>
            </thead>
            <tbody className={cn("[&_tr:last-child]:border-0")}>
              {sortedData.map((row) => (
                <tr
                  key={row.ID}
                  className={cn("border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted")}
                >
                  <td className={cn("p-2 align-middle")}>{row.ID}</td>
                  <td className={cn("p-2 align-middle")}>{row.UserName}</td>
                  <td className={cn("p-2 align-middle")}>{row.DisplayName}</td>
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
