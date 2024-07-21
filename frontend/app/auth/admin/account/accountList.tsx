// components/AccountList.tsx
'use client';
import { useEffect, useState } from 'react';
import { AccountData, ListAccountReq, ListAccountResp } from '@/proto/main';
import { listAccount } from '@/utils/api/api';

const AccountList = () => {
  const [accounts, setAccounts] = useState<AccountData[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [page, setPage] = useState<number>(1);
  const [size, setSize] = useState<number>(10);
  const [sortField, setSortField] = useState<string>('UserName');
  const [sortOrder, setSortOrder] = useState<string>('asc');
  const [query, setQuery] = useState<string>('');

  const fetchData = async () => {
    const req: ListAccountReq = {
      Option: {
        Pager: { Number: page, Size: size },
        Sorter: { Asc: sortOrder === 'asc', Field: sortField },
        Query: query ? [{ Fuzzy: true, Field: 'UserName', Value: query }] : [],
      },
    };

    const resp: ListAccountResp | null = await listAccount(req);

    if (resp) {
      setAccounts(resp.List);
      setTotal(resp.Total);
    }
  };

  useEffect(() => {
    fetchData();
  }, [page, size, sortField, sortOrder, query]);

  const handleSort = (field: string) => {
    const order = sortField === field && sortOrder === 'asc' ? 'desc' : 'asc';
    setSortField(field);
    setSortOrder(order);
  };

  const handleQuery = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    setQuery(formData.get('query') as string);
  };

  return (
    <div className="p-4">
      <form onSubmit={handleQuery} className="mb-4">
        <input
          type="text"
          name="query"
          placeholder="Search..."
          className="p-2 border border-gray-300 rounded-md"
        />
        <button type="submit" className="ml-2 p-2 bg-blue-500 text-white rounded-md">
          Search
        </button>
      </form>

      <table className="min-w-full divide-y divide-gray-200 border">
        <thead className="bg-gray-50">
          <tr>
            <th
              onClick={() => handleSort('UserName')}
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer"
            >
              Username
            </th>
            <th
              onClick={() => handleSort('DisplayName')}
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer"
            >
              Display Name
            </th>
            <th
              onClick={() => handleSort('Email')}
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer"
            >
              Email
            </th>
            <th
              onClick={() => handleSort('State')}
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer"
            >
              State
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {accounts.map((account) => (
            <tr key={account.UserName}>
              <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{account.UserName}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{account.DisplayName}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{account.Email}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{account.State}</td>
            </tr>
          ))}
        </tbody>
      </table>

      <div className="flex justify-between items-center mt-4">
        <button
          disabled={page === 1}
          onClick={() => setPage(page - 1)}
          className="px-4 py-2 bg-gray-300 text-gray-700 rounded-md disabled:opacity-50"
        >
          Previous
        </button>
        <span className="text-sm text-gray-700">
          Page {page} of {Math.ceil(total / size)}
        </span>
        <button
          disabled={page * size >= total}
          onClick={() => setPage(page + 1)}
          className="px-4 py-2 bg-gray-300 text-gray-700 rounded-md disabled:opacity-50"
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default AccountList;
