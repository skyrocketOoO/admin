'use client';
import Header from './blog/components/header';
import { useEffect, useState } from 'react';
import { AddVisitCount } from '@/utils/visit';

export default function Home() {
  useEffect(() => {
    AddVisitCount();
  }, []);


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
