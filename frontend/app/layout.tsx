import '@/app/ui/global.css';
import { inter } from '@/app/ui/fonts';
import { Metadata } from 'next';
import Header from '@/app/header';
 
export const metadata: Metadata = {
  title: {
    template: '%s | SkyrocketOoO playground',
    default: 'SkyrocketOoO playground',
  },
  description: 'SkyrocketOoO playground.',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${inter.className} antialiased`}>
        <Header />
        {children}
      </body>
    </html>
  );
}
