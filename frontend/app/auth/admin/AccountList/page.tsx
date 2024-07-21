import { lusitana } from '@/global/fonts';

export default async function Page() {
  return (
    <main>
      <h1 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        AccountList
      </h1>
    </main>
  );
}