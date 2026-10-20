import { cn } from '@/lib/utils'

function Breadcrumb({ className, ...props }: React.ComponentProps<'nav'>) {
  return <nav aria-label="Breadcrumb" data-slot="breadcrumb" className={cn('', className)} {...props} />
}

function BreadcrumbList({ className, ...props }: React.ComponentProps<'ol'>) {
  return <ol data-slot="breadcrumb-list" className={cn('flex items-center gap-2 text-sm text-muted-foreground', className)} {...props} />
}

function BreadcrumbItem({ className, ...props }: React.ComponentProps<'li'>) {
  return <li data-slot="breadcrumb-item" className={cn('inline-flex items-center gap-2', className)} {...props} />
}

function BreadcrumbPage({ className, ...props }: React.ComponentProps<'span'>) {
  return <span aria-current="page" data-slot="breadcrumb-page" className={cn('font-medium text-foreground', className)} {...props} />
}

export { Breadcrumb, BreadcrumbItem, BreadcrumbList, BreadcrumbPage }
