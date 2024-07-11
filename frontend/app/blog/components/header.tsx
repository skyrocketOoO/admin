import Link from 'next/link';

const Header = () => {
  return (
    <header className="bg-gray-800 text-slate-300 p-4 h-10 flex items-center justify-between">
      <Link href="/blog" className="text-lg font-light">
        Blog
      </Link>
    </header>
  );
};

export default Header;
