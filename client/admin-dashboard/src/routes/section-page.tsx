import { Link, useLocation } from 'react-router-dom'
import { BookOpen, Boxes, CircleHelp, ClipboardList, Package, ShoppingBag, Users } from 'lucide-react'

const navigation = [
  { label: 'Blogs', path: '/blogs', icon: BookOpen },
  { label: 'Project', path: '/projects', icon: Boxes },
  { label: 'Store', path: '/stores', icon: ShoppingBag },
  { label: 'Product', path: '/products', icon: Package },
  { label: 'User', path: '/users', icon: Users },
  { label: 'Order', path: '/orders', icon: ClipboardList },
]

export function DashboardSectionPage() {
  const location = useLocation()
  const selected = navigation.find((item) => item.path === location.pathname)
  const Icon = selected?.icon ?? CircleHelp
  return (
    <section className="mx-auto max-w-4xl rounded-xl border bg-white p-8 shadow-sm">
      <span className="grid size-11 place-items-center rounded-lg bg-stone-100"><Icon className="size-5 text-stone-700" /></span>
      <h1 className="mt-5 text-2xl font-semibold tracking-tight">{selected?.label ?? 'Dashboard'}</h1>
      <p className="mt-2 text-sm leading-6 text-muted-foreground">This section is ready for its management tools.</p>
      <Link to="/dashboard" className="mt-5 inline-flex text-sm font-medium underline underline-offset-4">Back to dashboard</Link>
    </section>
  )
}
