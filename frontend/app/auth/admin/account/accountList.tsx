// components/AccountList.tsx
'use client';
import { useEffect, useState } from 'react';
import * as grpc from '@grpc/grpc-js';
import { MainClient, AccountData, ListAccountReq, ListAccountResp } from '@/proto/main';
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
    <div>
      <form onSubmit={handleQuery}>
        <input type="text" name="query" placeholder="Search..." />
        <button type="submit">Search</button>
      </form>

      <table>
        <thead>
          <tr>
            <th onClick={() => handleSort('UserName')}>Username</th>
            <th onClick={() => handleSort('DisplayName')}>Display Name</th>
            <th onClick={() => handleSort('Email')}>Email</th>
            <th onClick={() => handleSort('State')}>State</th>
          </tr>
        </thead>
        <tbody>
          {accounts.map((account) => (
            <tr key={account.UserName}>
              <td>{account.UserName}</td>
              <td>{account.DisplayName}</td>
              <td>{account.Email}</td>
              <td>{account.State}</td>
            </tr>
          ))}
        </tbody>
      </table>

      <div>
        <button disabled={page === 1} onClick={() => setPage(page - 1)}>
          Previous
        </button>
        <span>
          Page {page} of {Math.ceil(total / size)}
        </span>
        <button disabled={page * size >= total} onClick={() => setPage(page + 1)}>
          Next
        </button>
      </div>
    </div>
  );
};

export default AccountList;
