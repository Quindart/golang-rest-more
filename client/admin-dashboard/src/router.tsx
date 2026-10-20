import { lazy, Suspense } from 'react'
import { BrowserRouter, Navigate, Outlet, Route, Routes } from 'react-router-dom'
import { isAuthenticated } from '@/lib/auth'

const DashboardPage = lazy(() => import('@/routes/dashboard-page').then((module) => ({ default: module.DashboardPage })))
const BlogPage = lazy(() => import('@/routes/blog-page').then((module) => ({ default: module.BlogPage })))
const CustomersPage = lazy(() => import('@/routes/customers-page').then((module) => ({ default: module.CustomersPage })))
const UsersPage = lazy(() => import('@/routes/users-page').then((module) => ({ default: module.UsersPage })))
const LoginPage = lazy(() => import('@/routes/login-page').then((module) => ({ default: module.LoginPage })))
const DashboardSectionPage = lazy(() => import('@/routes/section-page').then((module) => ({ default: module.DashboardSectionPage })))

function RootRedirect() {
  return <Navigate to={isAuthenticated() ? '/dashboard' : '/login'} replace />
}

function LoginRoute() {
  return isAuthenticated() ? <Navigate to="/dashboard" replace /> : <LoginPage />
}

function ProtectedRoutes() {
  return isAuthenticated() ? <Outlet /> : <Navigate to="/login" replace />
}

export function AppRouter() {
  return (
    <BrowserRouter>
      <Suspense fallback={<div className="grid min-h-svh place-items-center bg-stone-50 text-sm text-muted-foreground">Loading…</div>}>
        <Routes>
          <Route path="/" element={<RootRedirect />} />
          <Route path="/login" element={<LoginRoute />} />
          <Route element={<ProtectedRoutes />}>
            <Route element={<DashboardPage />}>
              <Route path="/dashboard" element={<DashboardSectionPage />} />
              <Route path="/blogs" element={<BlogPage />} />
              <Route path="/customers" element={<CustomersPage />} />
              <Route path="/projects" element={<DashboardSectionPage />} />
              <Route path="/stores" element={<DashboardSectionPage />} />
              <Route path="/products" element={<DashboardSectionPage />} />
              <Route path="/users" element={<UsersPage />} />
              <Route path="/orders" element={<DashboardSectionPage />} />
            </Route>
          </Route>
          <Route path="*" element={<Navigate to="/login" replace />} />
        </Routes>
      </Suspense>
    </BrowserRouter>
  )
}
