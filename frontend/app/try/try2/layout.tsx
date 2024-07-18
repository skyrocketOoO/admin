

export default function DashboardLayout({
  children, // will be a page or nested layout
}: {
  children: React.ReactNode
}) {
  return (
    <section>
      <header>gg2</header>
      <nav></nav>
 
      {children}
    </section>
  )
}