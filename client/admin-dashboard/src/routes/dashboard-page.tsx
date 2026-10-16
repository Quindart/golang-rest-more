import { useNavigate } from 'react-router-dom'
import { LogOut, ShieldCheck } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { logout } from '@/lib/auth'

export function DashboardPage() {
  const navigate = useNavigate()

  function handleSignOut() {
    logout()
    navigate('/login', { replace: true })
  }

  return (
    <main className="min-h-svh bg-stone-100 text-stone-900">
      <header className="border-b border-stone-200 bg-white">
        <div className="mx-auto flex max-w-6xl items-center justify-between px-5 py-4 sm:px-8">
          <div className="flex items-center gap-3">
            <span className="grid size-9 place-items-center rounded-lg bg-stone-900 text-white">
              <ShieldCheck aria-hidden="true" className="size-5" />
            </span>
            <span className="font-semibold tracking-tight">Northstar Admin</span>
          </div>
          <Button variant="outline" onClick={handleSignOut}>
            <LogOut aria-hidden="true" />
            Sign out
          </Button>
        </div>
      </header>

      <section className="mx-auto max-w-6xl px-5 py-12 sm:px-8 sm:py-16">
        <p className="text-sm font-medium text-stone-500">Overview</p>
        <h1 className="mt-2 text-3xl font-semibold tracking-tight">Dashboard</h1>
        <p className="mt-3 max-w-xl text-stone-600">
          You are signed in to the demo admin workspace.
        </p>
      </section>
    </main>
  )
}
