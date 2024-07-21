'use client';
import Image from 'next/image';
import homeImg from '@/public/home.png'; // Adjust the import path if necessary

export default function Home() {
  return (
    <div className="relative w-full h-screen overflow-hidden">
      <Image
        src={homeImg}
        alt="Home"
        fill
        style={{ objectFit: 'cover' }} // Ensure the image covers the container
        className="absolute inset-0"
      />
      <main className="flex items-center justify-center relative z-10">
        {/* Your content here */}
      </main>
    </div>
  );
}
