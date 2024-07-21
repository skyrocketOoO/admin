'use client';
import Image from 'next/image';
import homeImg from '@/public/home.png'; // Adjust the import path if necessary

export default function Home() {
  return (
    <div className="relative w-full h-full">
      <Image
        src={homeImg}
        alt="Home"
        fill
        className="object-fill"
      />
    </div>
  );
}
