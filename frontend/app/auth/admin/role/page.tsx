'use client';
import { useEffect, useState } from "react";
import { cn } from "@/lib/utils";
import { serverSideClient } from "@/utils/proto/client";
import { ListAccountReq } from "@/proto/main_pb";
import { ListOption, Pager, Sorter, ConditionGroup, Condition, Concator } from "@/proto/common_pb";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuGroup,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Filter } from "lucide-react";


interface AccountData {
  ID: number;
  UserName: string;
  DisplayName: string;
  State: number;
}

interface Filter {
  
  column: string;
  value: string;
}

export default function Page() {
  const columns = ["ID", "UserName", "DisplayName", "State"];
  const pageSize = 10;

  const [accountData, setAccountData] = useState<AccountData[]>([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  // Sorting state
  const [sortedData, setSortedData] = useState<AccountData[]>([]);
  const [sortColumn, setSortColumn] = useState<keyof AccountData | null>(null);
  const [sortDirection, setSortDirection] = useState<"asc" | "desc">("asc");
  const [filterCount, setFilterCount] = useState(3);
  const [filters, setFilters] = useState<Filter[]>([
    {
      column: "Any column",
      value: "",
    },
    {
      column: "Any column",
      value: "",
    },
    {
      column: "Any column",
      value: "",
    },
  ]);
  const [debouncedFilters, setDebouncedFilters] = useState(filters);
  // debounce
  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedFilters(filters);
    }, 800);
    return () => clearTimeout(handler);
  }, [filters]);

  useEffect(() => {
    const fetchAccountData = async () => {
      setLoading(true);

      const listReq = new ListAccountReq({
        Option: new ListOption({
          pager: new Pager({ number: page, size: pageSize }),
        })
      });

      if (debouncedFilters.length > 0){
        listReq.Option = listReq.Option || new ListOption();
        listReq.Option.conditionGroup = new ConditionGroup({
          concator: Concator.AND,
          conditionGroups: []
        });   
        for (const filter of filters){
          if (filter.value == ""){
            continue;
          }
          
          let condGroup = new ConditionGroup({
            concator: Concator.OR,
            conditions: columns.map((column) => new Condition({
              field: column,
              operator: "LIKE",
              value: "%"+filter.value+"%"
            })),
          })
          if (filter.column == "Any column"){
            condGroup.conditions = columns.map((column) => new Condition({
              field: column,
              operator: "LIKE",
              value: "%"+filter.value+"%"
            }))
          }else{
            condGroup.conditions = [new Condition({
              field: filter.column,
              operator: "LIKE",
              value: "%"+filter.value+"%"
            })]
          }

          listReq.Option.conditionGroup.conditionGroups.push(condGroup);
        }     
      }
      console.log(listReq);
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
  }, [page, debouncedFilters]);

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
      setSortDirection((prev) => (
        prev === "asc" ? "desc" : "asc"
      ));
    } else {
      setSortColumn(column);
      setSortDirection("asc");
    }
  };

  const addFilter = () => {
    let c = filterCount + 1;
    setFilterCount(c);
    let newFilters = filters;
    newFilters.push({
      column: "Any column",
      value: "",
    });
    setFilters(newFilters);
  };

  // Function to remove the last filter input
  const removeFilter = () => {
    let c = filterCount - 1;
    setFilterCount(c);
    let newFilters = filters;
    newFilters.pop();
    setFilters(newFilters);
  };

  return (
    <div>
      <div className="flex items-center py-4">
        {Array.from({ length: filterCount }, (_, index) => (
          <div>
            <div className="flex items-center space-x-4">
              <DropdownMenu>
                <DropdownMenuTrigger>{filters[index].column}</DropdownMenuTrigger>
                <DropdownMenuContent>
                  <DropdownMenuItem
                  onClick={() => {
                    const newFilters = [...filters];
                    newFilters[index] = newFilters[index] || { column: "", value: "" }; // Ensure the object exists
                    newFilters[index].column = "Any column";
                    setFilters(newFilters);
                  }}
                  >Any column</DropdownMenuItem>
                  {columns.map((column) => (
                    <DropdownMenuItem 
                    onClick={() => {
                      const newFilters = [...filters];
                      newFilters[index] = newFilters[index] || { column: "", value: "" }; // Ensure the object exists
                      newFilters[index].column = column;
                      setFilters(newFilters);
                    }}
                    >{column}</DropdownMenuItem>
                  ))}
                </DropdownMenuContent>
              </DropdownMenu>
              <button>abc</button>
            </div>
            <input
              key={index}
              className="max-w-sm"
              placeholder={`Filter ${index + 1}...`}
              value={filters[index].value || ""}
              onChange={(e) => {
                const newFilters = [...filters];
                newFilters[index].value = e.target.value;
                setFilters(newFilters);
              }}
            />
          </div>
        ))}
        <div>
          <button className="mr-3" onClick={addFilter}>+</button>
          <button onClick={removeFilter} disabled={filterCount == 0}>-</button>
        </div>
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
                    {column} {sortColumn === column ? (sortDirection === "asc" ? "▲" : "▼") : "▲▼"}
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
