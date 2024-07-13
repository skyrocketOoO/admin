'use client';
import Header from './blog/components/header';

export default function Home() {
  const handleBlogLinkClick = () => {
  };
  return (
    <div>
    <Header onBlogLinkClick={handleBlogLinkClick}/>
    <
      main className="flex items-center justify-center md:h-screen"
      style={{ backgroundImage: "url('/home.png')" }}
    >
    
    </main>
    </div>
  );
}
