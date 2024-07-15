import Link from 'next/link';
import Image from 'next/image';
import head from '@/public/head.png';
import { useEffect, useState } from 'react';
import { loadVisitCount } from '@/utils/visit';

interface HeaderProps {
  onBlogLinkClick: () => void; // Callback function to handle blog link click
}

const Header: React.FC<HeaderProps> = ({ onBlogLinkClick }) => {
  const [visitCount, setVisitCount] = useState(0);

  useEffect(() => {
    const fetchVisitCount = async () => {
      try {
        const count = await loadVisitCount();
        setVisitCount(count);
      } catch (error) {
        console.error('Error loading visit count:', error);
      }
    }
    fetchVisitCount();
  }, []);

  return (
    <header className="bg-violet-50 font-semibold text-lg border text-gray-600 pl-6 pr-2 h-12 flex items-center">
      <div className='w-1/8'></div>
      <Link href="/" className="w-32 h-full font-mono  flex items-center hover:border-b-4 justify-center hover:border-black hover:text-gray-800">Home</Link>
      <div className='p-2'></div>
      <Link href="/blog" onClick={onBlogLinkClick} className="w-32 h-full  font-mono flex items-center justify-center hover:border-b-4 hover:border-black hover:text-gray-800">Blog</Link>
      <div className='p-2'></div>
      <Link href="/admin" className="w-32 h-full font-mono  flex items-center hover:border-b-4 justify-center hover:border-black hover:text-gray-800">Admin</Link>
      <div className="flex items-center ml-auto">
        <span className="text-gray-700 h-8 text-sm p-4 rounded-md bg-cyan-100 flex items-center justify-center mr-4">{visitCount}</span>
        <div className="w-9 h-9 relative mr-2">
          <Image src={head} alt="Head" layout="fill" className="object-contain" />
        </div>
      </div>
    </header>
  );
};

export default Header;
