import { useEffect, useMemo, useState } from 'react'
import type { FormEvent } from 'react'
import { Eye, EyeOff, Filter, Plus, Search, UserRound, Users, X } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'

type User = {
  id: string
  full_name: string
  user_name: string
  date_of_birth: string
  email?: string
  phone_number?: string
}

type CreateUserResponse = { message?: string; user?: User }

type FilterMode = 'all' | 'name' | 'email' | 'id'

export function UsersPage() {
  const [users, setUsers] = useState<User[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [search, setSearch] = useState('')
  const [filterMode, setFilterMode] = useState<FilterMode>('all')
  const [filterOpen, setFilterOpen] = useState(false)
  const [addOpen, setAddOpen] = useState(false)
  const [showPassword, setShowPassword] = useState(false)
  const [creating, setCreating] = useState(false)
  const [createError, setCreateError] = useState('')

  useEffect(() => {
    const controller = new AbortController()
    async function fetchUsers() {
      try {
        setLoading(true)
        setError('')
        const response = await fetch('http://localhost:8082/api/v1/users', { signal: controller.signal })
        if (!response.ok) throw new Error(`Failed to fetch users (${response.status})`)
        const result: unknown = await response.json()
        if (typeof result !== 'object' || result === null || !('users' in result) || !Array.isArray(result.users)) {
          throw new Error('Invalid users response format')
        }
        setUsers(result.users as User[])
      } catch (err) {
        if (err instanceof Error && err.name !== 'AbortError') setError(err.message || 'Unable to load users.')
      } finally {
        if (!controller.signal.aborted) setLoading(false)
      }
    }
    fetchUsers()
    return () => controller.abort()
  }, [])

  const filteredUsers = useMemo(() => {
    const query = search.trim().toLocaleLowerCase()
    return users.filter((user) => {
      if (!query) return true
      if (filterMode === 'name') return user.full_name.toLocaleLowerCase().includes(query)
      if (filterMode === 'email') return (user.email ?? '').toLocaleLowerCase().includes(query)
      if (filterMode === 'id') return user.id.toLocaleLowerCase().includes(query)
      return [user.full_name, user.user_name, user.email ?? '', user.phone_number ?? '', user.id].some((value) => value.toLocaleLowerCase().includes(query))
    })
  }, [filterMode, search, users])

  const activeFilterLabel = ({ all: 'All fields', name: 'Name', email: 'Email', id: 'User ID' })[filterMode]

  async function createUser(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setCreating(true)
    setCreateError('')
    const formData = new FormData(event.currentTarget)
    const payload = {
      full_name: String(formData.get('full_name') ?? '').trim(),
      user_name: String(formData.get('user_name') ?? '').trim(),
      password: String(formData.get('password') ?? ''),
      date_of_birth: String(formData.get('date_of_birth') ?? ''),
      phone_number: String(formData.get('phone_number') ?? '').trim(),
      email: String(formData.get('email') ?? '').trim(),
    }

    try {
      const response = await fetch('http://localhost:8082/api/v1/users', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', Accept: 'application/json' },
        body: JSON.stringify(payload),
      })
      const result: CreateUserResponse = await response.json().catch(() => ({}))
      if (!response.ok) throw new Error(result.message || `Failed to create user (${response.status})`)
      if (!result.user) throw new Error('User was created, but the API response was incomplete.')
      setUsers((current) => [result.user!, ...current])
      setAddOpen(false)
      setShowPassword(false)
      event.currentTarget.reset()
    } catch (err) {
      setCreateError(err instanceof Error ? err.message : 'Unable to create user.')
    } finally {
      setCreating(false)
    }
  }

  return (
    <div className="mx-auto max-w-7xl space-y-7">
      <section className="flex flex-col justify-between gap-5 sm:flex-row sm:items-end">
        <div><p className="text-sm font-medium text-stone-500">Workspace / Management</p><h1 className="mt-2 text-3xl font-semibold tracking-tight">Users</h1><p className="mt-2 text-sm text-stone-600">View and manage your user directory.</p></div>
        <div className="flex items-center gap-3 self-start rounded-2xl border border-stone-200 bg-white px-4 py-3 shadow-sm sm:self-auto"><span className="grid size-10 place-items-center rounded-xl bg-stone-100"><Users className="size-5 text-stone-700" /></span><div><p className="text-2xl font-semibold tabular-nums">{loading ? '—' : users.length}</p><p className="text-xs text-stone-500">Total users</p></div></div>
      </section>

      <section className="overflow-hidden rounded-2xl border border-stone-200 bg-white shadow-sm">
        <div className="flex flex-col gap-4 border-b border-stone-100 p-4 sm:flex-row sm:items-center sm:justify-between sm:px-6">
          <div><h2 className="font-semibold tracking-tight">User directory</h2><p className="mt-1 text-sm text-stone-500">{loading ? 'Loading users…' : `${filteredUsers.length} ${filteredUsers.length === 1 ? 'user' : 'users'} shown`}</p></div>
          <div className="flex flex-col gap-2 sm:flex-row">
            <label className="relative min-w-0 sm:w-64"><Search className="pointer-events-none absolute left-3 top-1/2 size-4 -translate-y-1/2 text-stone-400" /><Input aria-label="Search users" placeholder="Search users…" value={search} onChange={(event) => setSearch(event.target.value)} className="h-10 bg-white pl-9" /></label>
            <div className="relative"><Button type="button" variant="outline" className="h-10 w-full gap-2 sm:w-auto" aria-expanded={filterOpen} onClick={() => setFilterOpen((open) => !open)}><Filter className="size-4" />Filter<span className="text-stone-400">·</span>{activeFilterLabel}</Button>
              {filterOpen && <div className="absolute right-0 z-20 mt-2 w-44 rounded-xl border border-stone-200 bg-white p-1.5 shadow-lg">{([{ value: 'all', label: 'All fields' }, { value: 'name', label: 'Name' }, { value: 'email', label: 'Email' }, { value: 'id', label: 'User ID' }] as const).map((option) => <button key={option.value} type="button" className={`block w-full rounded-lg px-3 py-2 text-left text-sm hover:bg-stone-50 ${filterMode === option.value ? 'font-medium text-stone-900' : 'text-stone-600'}`} onClick={() => { setFilterMode(option.value); setFilterOpen(false) }}>{option.label}</button>)}</div>}
            </div>
            <Button type="button" className="h-10 gap-2" onClick={() => { setCreateError(''); setAddOpen(true) }}><Plus className="size-4" />Add user</Button>
          </div>
        </div>
        {loading ? <div className="space-y-3 p-6">{Array.from({ length: 4 }).map((_, index) => <div key={index} className="h-14 animate-pulse rounded-lg bg-stone-100" />)}</div>
          : error ? <div className="px-6 py-16 text-center"><p className="font-medium text-red-700">Unable to load users</p><p className="mt-2 text-sm text-stone-500">{error}</p><Button variant="outline" className="mt-5" onClick={() => window.location.reload()}>Reload page</Button></div>
            : filteredUsers.length === 0 ? <div className="px-6 py-16 text-center"><span className="mx-auto grid size-11 place-items-center rounded-xl bg-stone-100"><Users className="size-5 text-stone-500" /></span><h3 className="mt-4 font-semibold">{search ? 'No matching users' : 'No users yet'}</h3><p className="mt-2 text-sm text-stone-500">{search ? 'Try another search or change the filter.' : 'Users will appear here when they are available.'}</p></div>
              : <div className="overflow-x-auto"><table className="w-full min-w-[680px] text-left text-sm"><thead className="bg-stone-50/80 text-xs font-medium uppercase tracking-wide text-stone-500"><tr><th className="px-6 py-3.5">User</th><th className="px-6 py-3.5">Username</th><th className="px-6 py-3.5">Date of birth</th><th className="px-6 py-3.5">User ID</th></tr></thead><tbody className="divide-y divide-stone-100">{filteredUsers.map((user) => <tr key={user.id} className="transition hover:bg-stone-50/70"><td className="px-6 py-4"><div className="flex items-center gap-3"><span className="grid size-10 shrink-0 place-items-center rounded-full bg-stone-100 text-sm font-semibold text-stone-700">{user.full_name.split(/\s+/).map((part) => part[0]).slice(0, 2).join('').toUpperCase()}</span><span className="font-medium text-stone-900">{user.full_name}</span></div></td><td className="px-6 py-4 text-stone-600">{user.user_name}</td><td className="px-6 py-4 text-stone-600">{user.date_of_birth || '—'}</td><td className="px-6 py-4 font-mono text-xs text-stone-500">{user.id}</td></tr>)}</tbody></table></div>}
      </section>

      {addOpen && <div className="fixed inset-0 z-50 grid place-items-center overflow-y-auto bg-stone-950/30 p-4" onMouseDown={(event) => { if (event.target === event.currentTarget && !creating) setAddOpen(false) }}>
        <form onSubmit={createUser} className="my-auto w-full max-w-md rounded-2xl border border-stone-200 bg-white p-6 shadow-xl">
          <div className="flex items-start justify-between"><div><h2 id="add-user-title" className="text-lg font-semibold">Add user</h2><p className="mt-1 text-sm text-stone-500">Enter account and contact details.</p></div><Button type="button" variant="ghost" size="icon" aria-label="Close" disabled={creating} onClick={() => setAddOpen(false)}><X className="size-4" /></Button></div>
          <div className="mt-6 space-y-4"><label className="block space-y-1.5 text-sm font-medium">Full name<Input name="full_name" required placeholder="e.g. Demo User" className="h-10" /></label><label className="block space-y-1.5 text-sm font-medium">User name<Input name="user_name" required placeholder="e.g. demo.user" className="h-10" /></label><label className="block space-y-1.5 text-sm font-medium">Password<span className="relative block"><Input name="password" type={showPassword ? 'text' : 'password'} minLength={8} required placeholder="At least 8 characters" className="h-10 pr-10" /><button type="button" aria-label={showPassword ? 'Hide password' : 'Show password'} aria-pressed={showPassword} onClick={() => setShowPassword((visible) => !visible)} className="absolute inset-y-0 right-0 grid w-10 place-items-center text-stone-500 hover:text-stone-900">{showPassword ? <EyeOff className="size-4" /> : <Eye className="size-4" />}</button></span></label><label className="block space-y-1.5 text-sm font-medium">Date of birth<Input name="date_of_birth" type="date" className="h-10" /></label><label className="block space-y-1.5 text-sm font-medium">Phone number<Input name="phone_number" type="tel" required placeholder="e.g. 0901234567" className="h-10" /></label><label className="block space-y-1.5 text-sm font-medium">Email<Input name="email" type="email" required placeholder="name@example.com" className="h-10" /></label></div>
          {createError && <p role="alert" className="mt-4 rounded-lg bg-red-50 px-3 py-2 text-sm text-red-700">{createError}</p>}
          <div className="mt-6 flex justify-end gap-2"><Button type="button" variant="outline" disabled={creating} onClick={() => setAddOpen(false)}>Cancel</Button><Button type="submit" disabled={creating}><UserRound className="size-4" />{creating ? 'Creating…' : 'Create user'}</Button></div>
        </form>
      </div>}
    </div>
  )
}
