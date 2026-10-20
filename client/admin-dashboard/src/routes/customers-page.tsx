import { useEffect, useMemo, useState } from 'react'
import { Filter, Search, Users } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'

type Customer = {
  id: string
  userId: string
  fullName: string
  phoneNumber: string
  email: string
}

type FilterMode = 'all' | 'email' | 'phone'

export function CustomersPage() {
  const [customers, setCustomers] = useState<Customer[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [search, setSearch] = useState('')
  const [filterMode, setFilterMode] = useState<FilterMode>('all')
  const [filterOpen, setFilterOpen] = useState(false)

  useEffect(() => {
    const controller = new AbortController()

    async function fetchCustomers() {
      try {
        setLoading(true)
        setError('')
        const response = await fetch('http://localhost:8083/api/v1/customers', { signal: controller.signal })
        if (!response.ok) throw new Error(`Failed to fetch customers (${response.status})`)
        const result: unknown = await response.json()
        if (!Array.isArray(result)) throw new Error('Invalid customers response format')
        setCustomers(result as Customer[])
      } catch (err) {
        if (err instanceof Error && err.name !== 'AbortError') setError(err.message || 'Unable to load customers.')
      } finally {
        if (!controller.signal.aborted) setLoading(false)
      }
    }

    fetchCustomers()
    return () => controller.abort()
  }, [])

  const filteredCustomers = useMemo(() => {
    const query = search.trim().toLocaleLowerCase()
    return customers.filter((customer) => {
      if (!query) return true
      if (filterMode === 'email') return customer.email.toLocaleLowerCase().includes(query)
      if (filterMode === 'phone') return customer.phoneNumber.toLocaleLowerCase().includes(query)
      return [customer.fullName, customer.email, customer.phoneNumber, customer.id]
        .some((value) => value.toLocaleLowerCase().includes(query))
    })
  }, [customers, filterMode, search])

  const activeFilterLabel = filterMode === 'all' ? 'All fields' : filterMode === 'email' ? 'Email' : 'Phone number'

  return (
    <div className="mx-auto max-w-7xl space-y-7">
      <section className="flex flex-col justify-between gap-5 sm:flex-row sm:items-end">
        <div>
          <p className="text-sm font-medium text-stone-500">Workspace / Management</p>
          <h1 className="mt-2 text-3xl font-semibold tracking-tight">Customers</h1>
          <p className="mt-2 text-sm text-stone-600">View and manage your customer directory.</p>
        </div>
        <div className="flex items-center gap-3 self-start rounded-2xl border border-stone-200 bg-white px-4 py-3 shadow-sm sm:self-auto">
          <span className="grid size-10 place-items-center rounded-xl bg-stone-100"><Users className="size-5 text-stone-700" /></span>
          <div><p className="text-2xl font-semibold tabular-nums">{loading ? '—' : customers.length}</p><p className="text-xs text-stone-500">Total customers</p></div>
        </div>
      </section>

      <section className="overflow-hidden rounded-2xl border border-stone-200 bg-white shadow-sm">
        <div className="flex flex-col gap-4 border-b border-stone-100 p-4 sm:flex-row sm:items-center sm:justify-between sm:px-6">
          <div>
            <h2 className="font-semibold tracking-tight">Customer directory</h2>
            <p className="mt-1 text-sm text-stone-500">{loading ? 'Loading customers…' : `${filteredCustomers.length} ${filteredCustomers.length === 1 ? 'customer' : 'customers'} shown`}</p>
          </div>
          <div className="flex flex-col gap-2 sm:flex-row">
            <label className="relative min-w-0 sm:w-64">
              <Search className="pointer-events-none absolute left-3 top-1/2 size-4 -translate-y-1/2 text-stone-400" />
              <Input aria-label="Search customers" placeholder="Search customers…" value={search} onChange={(event) => setSearch(event.target.value)} className="h-10 bg-white pl-9" />
            </label>
            <div className="relative">
              <Button type="button" variant="outline" className="h-10 w-full gap-2 sm:w-auto" aria-expanded={filterOpen} onClick={() => setFilterOpen((open) => !open)}>
                <Filter className="size-4" />Filter<span className="text-stone-400">·</span>{activeFilterLabel}
              </Button>
              {filterOpen && <div className="absolute right-0 z-20 mt-2 w-44 rounded-xl border border-stone-200 bg-white p-1.5 shadow-lg">
                {([{ value: 'all', label: 'All fields' }, { value: 'email', label: 'Email' }, { value: 'phone', label: 'Phone number' }] as const).map((option) => <button key={option.value} type="button" className={`block w-full rounded-lg px-3 py-2 text-left text-sm hover:bg-stone-50 ${filterMode === option.value ? 'font-medium text-stone-900' : 'text-stone-600'}`} onClick={() => { setFilterMode(option.value); setFilterOpen(false) }}>{option.label}</button>)}
              </div>}
            </div>
          </div>
        </div>

        {loading ? <div className="space-y-3 p-6">{Array.from({ length: 4 }).map((_, index) => <div key={index} className="h-14 animate-pulse rounded-lg bg-stone-100" />)}</div>
          : error ? <div className="px-6 py-16 text-center"><p className="font-medium text-red-700">Unable to load customers</p><p className="mt-2 text-sm text-stone-500">{error}</p><Button variant="outline" className="mt-5" onClick={() => window.location.reload()}>Reload page</Button></div>
            : filteredCustomers.length === 0 ? <div className="px-6 py-16 text-center"><span className="mx-auto grid size-11 place-items-center rounded-xl bg-stone-100"><Users className="size-5 text-stone-500" /></span><h3 className="mt-4 font-semibold">{search ? 'No matching customers' : 'No customers yet'}</h3><p className="mt-2 text-sm text-stone-500">{search ? 'Try another search or change the filter.' : 'Customers will appear here when they are available.'}</p></div>
              : <div className="overflow-x-auto"><table className="w-full min-w-[720px] text-left text-sm"><thead className="bg-stone-50/80 text-xs font-medium uppercase tracking-wide text-stone-500"><tr><th className="px-6 py-3.5">Customer</th><th className="px-6 py-3.5">Phone number</th><th className="px-6 py-3.5">Email</th><th className="px-6 py-3.5">Customer ID</th></tr></thead><tbody className="divide-y divide-stone-100">{filteredCustomers.map((customer) => <tr key={customer.id} className="transition hover:bg-stone-50/70"><td className="px-6 py-4"><div className="flex items-center gap-3"><span className="grid size-10 shrink-0 place-items-center rounded-full bg-stone-100 text-sm font-semibold text-stone-700">{customer.fullName.split(/\s+/).map((part) => part[0]).slice(0, 2).join('').toUpperCase()}</span><div><p className="font-medium text-stone-900">{customer.fullName}</p><p className="mt-0.5 text-xs text-stone-500">User ID: {customer.userId}</p></div></div></td><td className="px-6 py-4 text-stone-600">{customer.phoneNumber}</td><td className="px-6 py-4 text-stone-600">{customer.email}</td><td className="px-6 py-4 font-mono text-xs text-stone-500">{customer.id}</td></tr>)}</tbody></table></div>}
      </section>

    </div>
  )
}
