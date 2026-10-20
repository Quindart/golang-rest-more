import { useEffect, useState } from 'react'
import { ArrowUpRight, CalendarDays, UserRound, BookOpen, RefreshCw } from 'lucide-react'

import { Button } from '@/components/ui/button'

type Blog = {
  id: string
  author_name: string
  category_name: string
  category_path: string
  content: string
  created_at: string
  img_url: string
  title: string
}

export function BlogPage() {
  const [data, setData] = useState<Blog[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    const controller = new AbortController()
    async function fetchBlogs() {
      try {
        setLoading(true)
        setError('')
        const response = await fetch(
          'http://localhost:8082/api/v1/blogs',
          { signal: controller.signal },
        )
        if (!response.ok) {
          throw new Error(`Failed to fetch blogs (${response.status})`)
        }
        const result = await response.json()
        const blogs = Array.isArray(result) ? result : result.blogs
        if (!Array.isArray(blogs)) {
          throw new Error('Invalid blogs response format')
        }
        setData(blogs)
      } catch (err) {
        if (err instanceof Error && err.name !== 'AbortError') {
          setError(err.message || 'Unable to load blogs.')
        }
      } finally {
        if (!controller.signal.aborted) {
          setLoading(false)
        }
      }
    }
    fetchBlogs()
    return () => controller.abort()
  }, [])

  const formatDate = (date: string) => {
    const parsedDate = new Date(date)

    if (Number.isNaN(parsedDate.getTime())) return 'Unknown date'

    return new Intl.DateTimeFormat('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    }).format(parsedDate)
  }

  return (
    <main className="min-h-svh bg-stone-50 text-stone-900">
      <section className="mx-auto max-w-7xl px-5 pb-8 pt-10 sm:px-8 sm:pt-14">
        <div className="flex flex-col justify-between gap-5 sm:flex-row sm:items-end">
          <div>
            <p className="text-sm font-medium text-stone-500">
              Workspace / Overview
            </p>

            <h1 className="mt-2 text-3xl font-semibold tracking-tight sm:text-4xl">
              Blog dashboard
            </h1>

            <p className="mt-3 max-w-xl text-sm leading-6 text-stone-600 sm:text-base">
              Manage your stories, explore categories, and keep track of your
              published content in one place.
            </p>
          </div>

          <div className="flex w-fit items-center gap-3 rounded-2xl border border-stone-200 bg-white px-4 py-3 shadow-sm">
            <span className="grid size-10 place-items-center rounded-xl bg-stone-100">
              <BookOpen className="size-5 text-stone-700" />
            </span>

            <div>
              <p className="text-2xl font-semibold tabular-nums">
                {loading ? '—' : data.length}
              </p>
              <p className="text-xs text-stone-500">Total articles</p>
            </div>
          </div>
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-5 pb-16 sm:px-8">
        <div className="mb-6 flex items-center justify-between">
          <div>
            <h2 className="text-lg font-semibold tracking-tight">
              Latest articles
            </h2>
            <p className="mt-1 text-sm text-stone-500">
              Browse your content library
            </p>
          </div>
        </div>

        {loading ? (
          <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {Array.from({ length: 6 }).map((_, index) => (
              <div
                key={index}
                className="animate-pulse overflow-hidden rounded-2xl border border-stone-200 bg-white"
              >
                <div className="aspect-[16/9] bg-stone-200" />
                <div className="space-y-4 p-5">
                  <div className="h-4 w-24 rounded bg-stone-200" />
                  <div className="h-6 w-4/5 rounded bg-stone-200" />
                  <div className="h-4 w-full rounded bg-stone-100" />
                  <div className="h-4 w-2/3 rounded bg-stone-100" />
                </div>
              </div>
            ))}
          </div>
        ) : error ? (
          <div className="rounded-2xl border border-red-200 bg-white px-6 py-12 text-center">
            <p className="font-medium text-red-700">Unable to load articles</p>
            <p className="mt-2 text-sm text-stone-500">{error}</p>
            <Button
              variant="outline"
              className="mt-5 gap-2"
              onClick={() => window.location.reload()}
            >
              <RefreshCw className="size-4" />
              Reload page
            </Button>
          </div>
        ) : data.length === 0 ? (
          <div className="rounded-2xl border border-dashed border-stone-300 bg-white px-6 py-16 text-center">
            <BookOpen className="mx-auto size-10 text-stone-400" />
            <h3 className="mt-4 font-semibold">No articles yet</h3>
            <p className="mt-2 text-sm text-stone-500">
              Your articles will appear here once they are available.
            </p>
          </div>
        ) : (
          <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {data.map((item) => (
              <article
                key={item.id}
                className="group flex h-full flex-col overflow-hidden rounded-2xl border border-stone-200/80 bg-white shadow-sm transition duration-300 hover:-translate-y-1 hover:border-stone-300 hover:shadow-xl hover:shadow-stone-200/50"
              >
                <div className="relative aspect-[16/9] overflow-hidden bg-stone-100">
                  <img
                    src={item.img_url}
                    alt={item.title}
                    loading="lazy"
                    className="size-full object-cover transition duration-500 group-hover:scale-105"
                    onError={(event) => {
                      event.currentTarget.style.display = 'none'
                    }}
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/35 via-transparent to-transparent" />
                  <span className="absolute left-4 top-4 rounded-full border border-white/50 bg-white/90 px-3 py-1 text-xs font-semibold text-stone-800 shadow-sm backdrop-blur">
                    {item.category_name || 'Uncategorized'}
                  </span>
                  <span className="absolute bottom-4 right-4 rounded-full bg-black/40 p-2 text-white opacity-0 backdrop-blur transition group-hover:opacity-100">
                    <ArrowUpRight className="size-4" />
                  </span>
                </div>
                <div className="flex flex-1 flex-col p-5 sm:p-6">
                  <div className="flex items-center gap-2 text-xs text-stone-500">
                    <CalendarDays className="size-3.5 shrink-0" />
                    <time dateTime={item.created_at}>
                      {formatDate(item.created_at)}
                    </time>
                    <span className="text-stone-300">/</span>
                    <span className="truncate">
                      {item.category_path || 'general'}
                    </span>
                  </div>

                  <h3 className="mt-4 line-clamp-2 text-xl font-semibold leading-snug tracking-tight text-stone-900 transition group-hover:text-stone-600">
                    {item.title}
                  </h3>

                  <p className="mt-3 line-clamp-3 flex-1 text-sm leading-6 text-stone-600">
                    {item.content}
                  </p>

                  <div className="mt-6 flex items-center justify-between gap-3 border-t border-stone-100 pt-4">
                    <div className="flex min-w-0 items-center gap-3">
                      <span className="grid size-9 shrink-0 place-items-center rounded-full bg-stone-100 text-stone-700">
                        <UserRound className="size-4" />
                      </span>

                      <div className="min-w-0">
                        <p className="text-[11px] text-stone-400">Written by</p>
                        <p className="truncate text-sm font-medium text-stone-800">
                          {item.author_name || 'Unknown author'}
                        </p>
                      </div>
                    </div>
                    <span className="shrink-0 text-xs text-stone-400">
                      Article
                    </span>
                  </div>
                </div>
              </article>
            ))}
          </div>
        )}
      </section>
    </main>
  )
}
