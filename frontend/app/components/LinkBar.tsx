import React from 'react';
import Image from 'next/image';
import Link from 'next/link';

const LinkBar = () => {
  return (
    <div className="linkbar bg-violet-50 w-14 p-2 border border-t-0 items-center h-auto">
      <Link href="https://github.com/skyrocketOoO" target="_blank" rel="noopener noreferrer">
        <div className="w-10 h-10 relative ml-auto mb-8 mt-4 hover:bg-violet-200 rounded-full">
          <Image src='/github.png' alt="Head" fill className="object-contain " />
        </div>
      </Link>
      <Link href="https://www.instagram.com/jimmyjenny0801" target="_blank" rel="noopener noreferrer">
        <div className="w-10 h-10 relative ml-auto mb-8 hover:bg-violet-200 rounded-full">
          <Image src='/instagram.png' alt="Head" fill className="object-contain" />
        </div>
      </Link>
      <Link href="https://www.linkedin.com/in/jimmy-huang-07aa4722a/" target="_blank" rel="noopener noreferrer">
        <div className="w-10 h-10 relative ml-auto mb-8 hover:bg-violet-200 rounded-full">
          <Image src='/linkedin.png' alt="Head" fill className="object-contain" />
        </div>
      </Link>
      <Link href="https://leetcode.com/u/rivendinner/" target="_blank" rel="noopener noreferrer">
        <div className="w-10 h-10 relative ml-auto mb-8 hover:bg-violet-200 rounded-full">
          <Image src='/leetcode.svg' alt="Head" fill className="object-contain" />
        </div>
      </Link>
      <Link href="https://www.hackerrank.com/profile/skyrocketqy81" target="_blank" rel="noopener noreferrer">
        <div className="w-10 h-10 relative ml-auto mb-8 hover:bg-violet-200 rounded-full">
          <Image src='/hackerrank.png' alt="Head" fill className="object-contain" />
        </div>
      </Link>
    </div>
  );
};

export default LinkBar;
