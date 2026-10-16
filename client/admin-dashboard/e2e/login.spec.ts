import { expect, test } from '@playwright/test'

test('empty submission reports required fields', async ({ page }) => {
  await page.goto('/login')
  await page.getByRole('button', { name: 'Sign in' }).click()

  await expect(page.getByText('Username is required.')).toBeVisible()
  await expect(page.getByText('Password is required.')).toBeVisible()
  await expect(page).toHaveURL(/\/login$/)
})

test('incorrect credentials show a generic error and stay on login', async ({ page }) => {
  await page.goto('/login')
  await page.getByLabel('Username').fill('admin')
  await page.getByLabel('Password').fill('incorrect-password')
  await page.getByRole('button', { name: 'Sign in' }).click()

  await expect(page.getByText('Invalid username or password.')).toBeVisible()
  await expect(page).toHaveURL(/\/login$/)
})

test('correct credentials sign in and open the dashboard', async ({ page }) => {
  await page.goto('/login')
  await page.getByLabel('Username').fill('admin')
  await page.getByLabel('Password').fill('admin123456Aa@')
  await page.getByRole('button', { name: 'Sign in' }).click()

  await expect(page.getByText('Welcome back!')).toBeVisible()
  await expect(page).toHaveURL(/\/dashboard$/)
  await expect(page.getByRole('heading', { name: 'Dashboard' })).toBeVisible()
})

test('direct dashboard access redirects to login without a session', async ({ page }) => {
  await page.goto('/dashboard')

  await expect(page).toHaveURL(/\/login$/)
  await expect(page.getByLabel('Username')).toBeVisible()
})

test('sign out clears access to the dashboard', async ({ page }) => {
  await page.goto('/login')
  await page.getByLabel('Username').fill('admin')
  await page.getByLabel('Password').fill('admin123456Aa@')
  await page.getByRole('button', { name: 'Sign in' }).click()
  await expect(page).toHaveURL(/\/dashboard$/)

  await page.getByRole('button', { name: 'Sign out' }).click()
  await expect(page).toHaveURL(/\/login$/)
  await page.goto('/dashboard')

  await expect(page).toHaveURL(/\/login$/)
  await expect(page.getByLabel('Username')).toBeVisible()
})
