'use client';
import {
  UserGroupIcon,
  HomeIcon,
  DocumentDuplicateIcon,
} from '@heroicons/react/24/outline';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import clsx from 'clsx';
import { useState } from 'react';


const links = [
  { name: 'Home', href: '/dashboard', icon: HomeIcon, count: 1 },
  {
    name: 'Invoices',
    href: '/dashboard/invoices',
    icon: DocumentDuplicateIcon,
    count: 2, 
    subLinks: [
      { name: 'Pending', href: '/dashboard/invoices/pending', count: 1 },
      { name: 'Paid', href: '/dashboard/invoices/paid', count: 1 },
    ],
  },
  { name: 'Customers', href: '/dashboard/customers', icon: UserGroupIcon, count: 0 },
];

export default function NavLinks() {
  const pathname = usePathname();
  const [openDropdown, setOpenDropdown] = useState<string | null>(null);

  const handleDropdown = (name: string) => {
    setOpenDropdown(openDropdown === name ? null : name);
  };

  return (
    <nav className="space-y-1 ">
      {links.map((link) => {
        const LinkIcon = link.icon;
        const isActive = openDropdown === link.name;

        return (
          <div key={link.name} className="relative">
            <button
              type="button"
              onClick={() => handleDropdown(link.name)}
              className={clsx(
                'flex items-center justify-between w-full p-2 text-base text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700',
                {
                  'bg-gray-100 dark:bg-gray-700': isActive,
                }
              )}
            >
              <div className="flex items-center">
                <LinkIcon className="w-6 mr-3" />
                {link.name}
              </div>
              {link.count > 0 && (
                <span
                  className={clsx(
                    'flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300'
                  )}
                >
                  {link.count}
                </span>
              )}
              {link.subLinks && (
                <span className="">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className={`h-4 w-4 transform ${isActive ? 'rotate-180' : 'rotate-0'}`}
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fillRule="evenodd"
                      d="M10 12a1 1 0 0 1-.707-.293l-4-4a1 1 0 1 1 1.414-1.414L10 9.586l3.293-3.293a1 1 0 0 1 1.414 1.414l-4 4A1 1 0 0 1 10 12z"
                    />
                  </svg>
                </span>
              )}
            </button>
            {link.subLinks && isActive && (
              <div className="py-1 space-y-2">
                {link.subLinks.map((subLink) => (
                  <Link
                    key={subLink.name}
                    href={subLink.href}
                    className={clsx(
                      'pl-8 block p-2 text-base text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700',
                      {
                        'bg-gray-100 dark:bg-gray-700': pathname === subLink.href,
                      }
                    )}
                  >
                    {subLink.name}
                    {link.count > 0 && (
                      <span
                        className={clsx(
                          'inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300'
                        )}
                      >
                        {link.count}
                      </span>
                    )}
                  </Link>
                ))}
              </div>
            )}
          </div>
        );
      })}
    </nav>
  );
}
