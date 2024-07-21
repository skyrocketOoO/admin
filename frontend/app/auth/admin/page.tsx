'use client'
import { lusitana } from '@/global/fonts';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';
export default async function Page() {
  const router = useRouter();

  useEffect(() => {
    if (!localStorage.getItem('session')) {
      router.push('/auth');
    }
  });

  return (
    <main>
      <h1 className={`${lusitana.className} mb-4 text-xl md:text-2xl`}>
        Dashboard
      </h1>
    </main>
  );
}