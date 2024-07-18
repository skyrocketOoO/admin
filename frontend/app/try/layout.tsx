export default function DashboardLayout({
  children, // will be a page or nested layout
}: {
  children: React.ReactNode
}) {
  return (
    <section>
      <header>gg</header>
      <nav></nav>
 
      {children}
    </section>
  )
}