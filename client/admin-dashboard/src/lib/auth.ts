const SESSION_KEY = 'admin-dashboard:demo-session'
const DEMO_USERNAME = 'admin'
const DEMO_PASSWORD = 'admin123456Aa@'

export function isAuthenticated(): boolean {
  return sessionStorage.getItem(SESSION_KEY) === 'true'
}

export function login(username: string, password: string): boolean {
  if (username !== DEMO_USERNAME || password !== DEMO_PASSWORD) {
    return false
  }

  sessionStorage.setItem(SESSION_KEY, 'true')
  return true
}

export function logout(): void {
  sessionStorage.removeItem(SESSION_KEY)
}
