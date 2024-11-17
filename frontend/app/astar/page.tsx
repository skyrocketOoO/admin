'use client';

import { Html, Head, Main, NextScript } from 'next/document';

export default function Home() {
  return (
    <div>
      <h1>Video Streaming Example</h1>
      <video controls width="800">
        <source src="http://localhost:8080/stream" type="video/mp4" />
        Your browser does not support the video tag.
      </video>
    </div>
  );
}
