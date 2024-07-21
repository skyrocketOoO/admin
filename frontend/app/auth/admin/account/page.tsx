import { lusitana } from '@/global/fonts';
import  AccountList from './accountList';

export default function Page() {
  return (
    <main>
      <h1 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        Account List
      </h1>
      <AccountList />
    </main>
  );
}
