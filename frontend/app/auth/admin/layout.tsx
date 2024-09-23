import SideNav from './sidenav';

export const experimental_ppr = true;
 
export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex h-full">
      <div className="w-1/4">
        <SideNav />
      </div>
      <div className='w-full'>
        <section>{children}</section>
      </div>
    </div>
  );
}