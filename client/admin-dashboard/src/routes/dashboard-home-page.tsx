import { BookOpen } from 'lucide-react'

export function DashboardHomePage() {
  return (
    <section className="mx-auto max-w-4xl rounded-xl border bg-white p-8 shadow-sm">
      <span className="grid size-11 place-items-center rounded-lg bg-stone-100"><BookOpen className="size-5 text-stone-700" /></span>
      <h1 className="mt-5 text-2xl font-semibold tracking-tight">Dashboard</h1>
      <p className="mt-2 text-sm leading-6 text-muted-foreground">Welcome to your admin workspace. Choose a section from the navigation to get started.</p>
      <div className="mt-6 border-t pt-6">
        <h2 className="text-sm font-semibold">Latest articles</h2>
        <p className="mt-1 text-sm text-muted-foreground">Open the blogs section to browse your content library.</p>
      </div>
    </section>
  )
}
