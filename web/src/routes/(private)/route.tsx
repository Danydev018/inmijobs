import { Outlet, createFileRoute, redirect } from '@tanstack/react-router'
import { authClient } from '@/lib/auth'
import { Navbar } from '@/components/Navbar'

export const Route = createFileRoute('/(private)')({
  component: RouteComponent,
  beforeLoad: async () => {
    const session = await authClient.getSession();
    if (!session.data) {
      throw redirect({ to: "/signin" });
    }
  }
})

function RouteComponent() {
  return (
    <div className="min-h-screen bg-[#F9FAFB]">
      <Navbar />
      <main>
        <Outlet />
      </main>
    </div>
  )
}