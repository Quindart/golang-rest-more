import { lazy, Suspense } from 'react'
import { Link, Outlet, useLocation, useNavigate } from 'react-router-dom'
import { BookOpen, Boxes, ChevronDown, ClipboardList, LogOut, Package, ShieldCheck, ShoppingBag, Users, UserRound } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Breadcrumb, BreadcrumbItem, BreadcrumbList, BreadcrumbPage } from '@/components/ui/breadcrumb'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { Sidebar, SidebarContent, SidebarGroup, SidebarGroupLabel, SidebarHeader, SidebarMenu, SidebarMenuButton, SidebarMenuItem, SidebarProvider } from '@/components/ui/sidebar'
import { logout } from '@/lib/auth'

const DashboardHomePage = lazy(() => import('@/routes/dashboard-home-page').then((module) => ({ default: module.DashboardHomePage })))

const navigation = [
  { label: 'Blogs', path: '/blogs', icon: BookOpen },
  { label: 'Project', path: '/projects', icon: Boxes },
  { label: 'Store', path: '/stores', icon: ShoppingBag },
  { label: 'Product', path: '/products', icon: Package },
  { label: 'User', path: '/users', icon: Users },
  { label: 'Customers', path: '/customers', icon: UserRound },
  { label: 'Order', path: '/orders', icon: ClipboardList },
]

export function DashboardPage() {
  const location = useLocation()
  const navigate = useNavigate()
  const activeItem = navigation.find((item) => item.path === location.pathname)
  const pageLabel = activeItem?.label ?? 'Dashboard'

  function handleLogout() {
    logout()
    navigate('/login', { replace: true })
  }

  return (
    <SidebarProvider className="bg-stone-50 text-stone-900">
      <Sidebar aria-label="Main navigation">
        <SidebarHeader>
          <Link to="/dashboard" className="flex items-center gap-3">
            <span className="grid size-9 place-items-center rounded-lg bg-stone-900 text-white"><ShieldCheck className="size-5" /></span>
            <span><span className="block text-sm font-semibold tracking-tight">Northstar Admin</span><span className="block text-xs text-muted-foreground">Workspace</span></span>
          </Link>
        </SidebarHeader>
        <SidebarContent>
          <SidebarGroup>
            <SidebarGroupLabel>Management</SidebarGroupLabel>
            <SidebarMenu>
              {navigation.map(({ label, path, icon: Icon }) => (
                <SidebarMenuItem  key={path}>
                  <SidebarMenuButton asChild isActive={location.pathname === path}>
                    <Link className="w-full items-center text-md flex gap-2" to={path}><Icon className="size-5" /><span>{label}</span></Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroup>
        </SidebarContent>
        <div className="border-t p-4 text-xs text-muted-foreground">Admin workspace</div>
      </Sidebar>

      <div className="flex min-w-0 flex-1 flex-col">
        <header className="sticky top-0 z-10 flex h-16 items-center justify-between border-b bg-white/90 px-4 backdrop-blur-xl sm:px-8">
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem><Link to="/dashboard" className="hover:text-foreground">Workspace</Link></BreadcrumbItem>
              <BreadcrumbItem><span aria-hidden="true">/</span></BreadcrumbItem>
              <BreadcrumbItem><BreadcrumbPage>{pageLabel}</BreadcrumbPage></BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
          <Popover>
            <PopoverTrigger asChild>
              <Button aria-label="User menu" variant="ghost" className="gap-2 rounded-full px-2">
                <span className="grid size-8 place-items-center rounded-full bg-stone-100 text-stone-700"><Users className="size-4" /></span>
                <span className="hidden text-sm font-medium sm:inline">Admin user</span>
                <ChevronDown className="size-4 text-muted-foreground" />
              </Button>
            </PopoverTrigger>
            <PopoverContent align="end" className="w-52">
              <div className="border-b px-3 py-2"><p className="text-sm font-medium">Admin user</p><p className="text-xs text-muted-foreground">Workspace account</p></div>
              <Button role="menuitem" variant="ghost" className="mt-1 w-full justify-start gap-2 text-destructive hover:text-destructive" onClick={handleLogout}>
                <LogOut className="size-4" />Log out
              </Button>
            </PopoverContent>
          </Popover>
        </header>
        <nav aria-label="Main navigation" className="flex gap-1 overflow-x-auto border-b bg-white px-3 py-2 md:hidden">
          {navigation.map(({ label, path, icon: Icon }) => (
            <Link key={path} to={path} aria-current={location.pathname === path ? 'page' : undefined} className={`flex shrink-0 items-center gap-2 rounded-md px-3 py-2 text-xs font-medium ${location.pathname === path ? 'bg-stone-100 text-stone-900' : 'text-stone-500'}`}>
              <Icon className="size-4" /><span>{label}</span>
            </Link>
          ))}
        </nav>
        <main className="flex-1 p-5 sm:p-8">
          {location.pathname === '/dashboard' ? (
            <Suspense fallback={<div className="py-12 text-center text-sm text-muted-foreground">Loading page…</div>}>
              <DashboardHomePage />
            </Suspense>
          ) : (
            <Suspense fallback={<div className="py-12 text-center text-sm text-muted-foreground">Loading page…</div>}>
              <Outlet />
            </Suspense>
          )}
        </main>
      </div>
    </SidebarProvider>
  )
}

