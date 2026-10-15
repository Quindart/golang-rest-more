# Admin Dashboard Login Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Turn the empty package into a React admin app with demo login, a protected dashboard route, and Playwright browser coverage.

**Architecture:** Vite serves a React + TypeScript SPA. React Router handles `/login` and `/dashboard`; a small auth module compares the requested demo credentials and stores a tab-scoped session. React Hook Form handles the login form, shadcn/ui supplies its controls, and notistack provides login feedback.

**Tech Stack:** React, TypeScript, Vite, Tailwind CSS v4, shadcn/ui, React Router, React Hook Form, notistack, Playwright.

**Spec:** `docs/superpowers/specs/2026-10-15-admin-dashboard-login-design.md`

## Global Constraints

- Demo username is `admin` and demo password is `admin123456Aa@`.
- Authentication state uses `sessionStorage`, not persistent storage.
- The dashboard remains protected unless the demo session is present.
- Login errors use generic feedback and do not distinguish username from password.
- No backend or production authentication is part of this work.
- Use pnpm as the package manager.

## Review Focus

- Empty credentials: both fields show required feedback before authentication.
- Incorrect credentials: show one generic error and remain on `/login`.
- Correct credentials: write session state, show success feedback, and navigate to `/dashboard`.
- Direct dashboard navigation without a session: redirect to `/login`.
- Sign-out after a successful login: clear session and prevent dashboard access after navigation or refresh.

---

### Task 1: Scaffold the React and UI toolchain

**Files:**
- Modify: `package.json`
- Create: `index.html`
- Create: `vite.config.ts`
- Create: `tsconfig.json`
- Create: `tsconfig.app.json`
- Create: `tsconfig.node.json`
- Create: `src/main.tsx`
- Create: `src/index.css`
- Create: `components.json`
- Generate: `src/components/ui/*` for the button, card, input, and label primitives
- Generate: `pnpm-lock.yaml`

**Interfaces:**
- Produces: Vite React TypeScript app shell; `@/*` resolves to `src/*`; Tailwind v4 is loaded through the Vite plugin; shadcn/ui CLI can add project-local components.

- [ ] **Step 1: Replace the placeholder package scripts and establish the Vite entry files.** Configure `dev`, `build`, `preview`, `test:e2e`, and `test:e2e:install` scripts. Set up React and Vite TypeScript entry points without adding login behavior.
- [ ] **Step 2: Install runtime and development dependencies.** Add React, React DOM, React Router, React Hook Form, notistack, Vite, the React Vite plugin, TypeScript, Tailwind CSS v4 with `@tailwindcss/vite`, Playwright, and shadcn/ui prerequisites. Generate the pnpm lockfile.
- [ ] **Step 3: Configure Tailwind and shadcn/ui.** Use the Tailwind Vite plugin and `@import "tailwindcss"`; configure neutral shadcn/ui defaults, the `@/*` alias, and add button, card, input, and label primitives.
- [ ] **Step 4: Verify the empty app shell builds.** Run `pnpm build`. Expected: Vite emits a production bundle without TypeScript or CSS configuration errors.

### Task 2: Specify the login behavior in Playwright

**Files:**
- Create: `playwright.config.ts`
- Create: `e2e/login.spec.ts`
- Modify: `package.json`

**Interfaces:**
- Consumes: the Vite development server started by `pnpm dev`.
- Produces: browser tests using accessible labels `Username` and `Password`, a `Sign in` button, a dashboard heading, and a `Sign out` button.

- [ ] **Step 1: Configure Playwright.** Set the test directory to `e2e`, use the Chromium project, and configure `webServer` to start Vite on a fixed local port with reuse of an already-running server.
- [ ] **Step 2: Write the five acceptance tests before implementing auth or pages.** Cover empty submission, incorrect credentials, successful credentials, direct protected-route access, and sign-out followed by blocked access. Use role/label locators and assert the URL, visible feedback, and protected content.
- [ ] **Step 3: Run the Playwright tests to verify RED.** Run `pnpm exec playwright test`. Expected: tests launch the app but fail because the login form and route behavior do not exist yet.

### Task 3: Implement demo authentication, routing, and the login screen

**Files:**
- Create: `src/lib/auth.ts`
- Create: `src/router.tsx`
- Create: `src/routes/login-page.tsx`
- Create: `src/routes/dashboard-page.tsx`
- Modify: `src/main.tsx`
- Modify: `src/index.css`

**Interfaces:**
- `src/lib/auth.ts` exports `isAuthenticated(): boolean`, `login(username: string, password: string): boolean`, and `logout(): void`.
- Router redirects `/` to `/dashboard` for an authenticated session and to `/login` otherwise; unauthenticated `/dashboard` redirects to `/login`.
- Login page uses React Hook Form, shadcn/ui controls, and notistack feedback.
- Dashboard provides a heading and sign-out control.

- [ ] **Step 1: Implement the auth module.** Compare exact credentials `admin` and `admin123456Aa@`; persist only a boolean session marker under a namespaced `sessionStorage` key; clear it on logout.
- [ ] **Step 2: Implement route behavior.** Add `/login`, a protected `/dashboard`, root redirects, and a catch-all redirect to `/login`. Ensure the route guard checks the current session when a route is entered.
- [ ] **Step 3: Implement the login form.** Add required validation, accessible labels and inline errors, disable submission while processing, show generic error/success snackbars, then navigate to `/dashboard` after successful login.
- [ ] **Step 4: Implement the protected landing page and sign-out.** Show a simple dashboard heading; logout clears the session and navigates to `/login`.
- [ ] **Step 5: Run the Playwright suite to verify GREEN.** Run `pnpm exec playwright test`. Expected: all five acceptance tests pass.
- [ ] **Step 6: Verify production output.** Run `pnpm build`. Expected: production bundle succeeds.

## Execution Notes

- The workspace is not a Git repository, so implementation commits are not available unless the user separately asks to initialize Git.
- Official setup references: [shadcn/ui with Vite](https://ui.shadcn.com/docs/installation/vite), [Tailwind CSS with Vite](https://tailwindcss.com/docs/installation/using-vite), [React Router routing](https://reactrouter.com/start/data/routing).
