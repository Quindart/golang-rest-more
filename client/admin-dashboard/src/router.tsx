import { BrowserRouter, Navigate, Outlet, Route, Routes } from 'react-router-dom'
import { isAuthenticated } from '@/lib/auth'
import { DashboardPage } from '@/routes/dashboard-page'
import { LoginPage } from '@/routes/login-page'

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
      <Routes>
        <Route path="/" element={<RootRedirect />} />
        <Route path="/login" element={<LoginRoute />} />
        <Route element={<ProtectedRoutes />}>
          <Route path="/dashboard" element={<DashboardPage />} />
        </Route>
        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
    </BrowserRouter>
  )
}
