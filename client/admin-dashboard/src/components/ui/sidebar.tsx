import { createContext, useContext } from 'react'
import { cn } from '@/lib/utils'

const SidebarContext = createContext(true)

function SidebarProvider({ children, className }: React.ComponentProps<'div'>) {
  return (
    <SidebarContext.Provider value>
      <div data-slot="sidebar-provider" className={cn('flex min-h-svh w-full', className)}>{children}</div>
    </SidebarContext.Provider>
  )
}

function Sidebar({ children, className, ...props }: React.ComponentProps<'aside'>) {
  const expanded = useContext(SidebarContext)
  return expanded ? <aside data-slot="sidebar" className={cn('hidden w-64 shrink-0 flex-col border-r bg-sidebar text-sidebar-foreground md:flex', className)} {...props}>{children}</aside> : null
}

function SidebarHeader({ children, className }: React.ComponentProps<'div'>) {
  return <div data-slot="sidebar-header" className={cn('flex h-16 items-center border-b px-5', className)}>{children}</div>
}

function SidebarContent({ children, className }: React.ComponentProps<'div'>) {
  return <div data-slot="sidebar-content" className={cn('flex-1 overflow-auto p-3', className)}>{children}</div>
}

function SidebarGroup({ children, className }: React.ComponentProps<'div'>) {
  return <div data-slot="sidebar-group" className={cn('space-y-1', className)}>{children}</div>
}

function SidebarGroupLabel({ children, className }: React.ComponentProps<'div'>) {
  return <div data-slot="sidebar-group-label" className={cn('px-3 pb-2 pt-3 text-xs font-medium text-muted-foreground', className)}>{children}</div>
}

function SidebarMenu({ children, className }: React.ComponentProps<'ul'>) {
  return <ul data-slot="sidebar-menu" className={cn('space-y-1', className)}>{children}</ul>
}

function SidebarMenuItem({ children, className }: React.ComponentProps<'li'>) {
  return <li data-slot="sidebar-menu-item" className={cn('', className)}>{children}</li>
}

function SidebarMenuButton({ children, className, isActive, asChild = false, ...props }: React.ComponentProps<'a'> & { isActive?: boolean; asChild?: boolean }) {
  const Comp = asChild ? 'span' : 'a'
  return <Comp data-slot="sidebar-menu-button" aria-current={isActive ? 'page' : undefined} className={cn('flex h-9 items-center gap-3 rounded-md px-3 text-sm text-muted-foreground transition-colors hover:bg-sidebar-accent hover:text-sidebar-accent-foreground', isActive && 'bg-sidebar-accent font-medium text-sidebar-accent-foreground', className)} {...props}>{children}</Comp>
}

export { Sidebar, SidebarContent, SidebarGroup, SidebarGroupLabel, SidebarHeader, SidebarMenu, SidebarMenuButton, SidebarMenuItem, SidebarProvider }
