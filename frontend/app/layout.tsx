import '@/app/ui/global.css';
import { inter } from '@/app/ui/fonts';
import { Metadata } from 'next';
import Header from '@/app/components/header';
import LinkBar from '@/app/components/LinkBar';

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
        <div className="flex">
          <main className="flex-1 pr-14"> {/* Add padding to the right to make space for the LinkBar */}
            {children}
          </main>
          <LinkBar />
        </div>
      </body>
    </html>
  );
}
