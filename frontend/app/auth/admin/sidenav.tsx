'use client'
import NavLinks from './nav-links';
import { PowerIcon } from '@heroicons/react/24/outline';
import { lusitana } from '@/global/fonts';
import { useRouter } from 'next/navigation';

export default function SideNav() {
  const router = useRouter();
  return (
    <div className="flex h-full flex-col px-3 py-4 md:px-2 bg-gray-50 dark:bg-gray-800">
      <div className="w-full text-white">
        <div
          className={`${lusitana.className} items-center leading-none text-black`}
        >
          <p className="text-[26px]">Admin</p>
        </div>
      </div>
      <div className="flex-grow px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
        <NavLinks />
      </div>
      <div className="h-auto w-full rounded-md bg-gray-50">
        <form
          action={async () => {
            localStorage.removeItem('session');
            router.push('/auth');
          }}
        >
          <button className="flex h-[48px] w-full items-center justify-center gap-2 rounded-md bg-gray-50 p-3 text-sm font-medium hover:bg-sky-100 hover:text-blue-600">
            <PowerIcon className="w-6" />
            <div className="hidden md:block">Sign Out</div>
          </button>
        </form>
      </div>
    </div>
  );
}
