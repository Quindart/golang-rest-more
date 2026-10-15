# Admin Dashboard Login Design

## Goal

Set up the empty `admin-dashboard` package as a small React admin app with a working demo login flow and browser tests.

## Current Context

The directory currently contains only `package.json`. It has no app source, build setup, tests, or Git repository. The package declares pnpm as its package manager.

## Scope

- Scaffold a React + TypeScript single-page app with Vite.
- Configure Tailwind CSS and shadcn/ui, adding only the UI components needed by the login screen.
- Add React Router routes for `/login` and `/dashboard`.
- Build an accessible, responsive login form using React Hook Form.
- Accept the demo credentials `admin` / `admin123456Aa@`.
- Keep demo authentication state in `sessionStorage` and protect `/dashboard`.
- Use notistack snackbars for successful and failed login feedback.
- Add Playwright end-to-end tests for validation, invalid credentials, successful login, and route protection.

## User Flow

Opening `/` or `/login` shows the login page unless the current tab has an authenticated demo session, in which case it redirects to `/dashboard`. Submitting with a missing username or password shows inline required-field feedback. Submitting credentials that do not match the demo account shows an error snackbar and keeps the user on the login page. A successful submission stores the demo session in `sessionStorage`, displays a success snackbar, and navigates to `/dashboard`. Visiting `/dashboard` without a session redirects to `/login`. The dashboard includes a sign-out action that clears the session and returns to `/login`.

## Architecture

- Vite owns local development and production builds.
- React Router owns route matching and navigation.
- A small authentication module owns credential comparison and session storage; it is explicitly demo-only and has no server-side security guarantees.
- A route guard reads the authentication module's session state before rendering the dashboard.
- The login page uses React Hook Form for field registration and validation, shadcn/ui primitives for controls, and notistack for submission results.
- Playwright exercises the app through the browser using its local web server configuration.

## Visual and Accessibility Requirements

- Use a restrained admin sign-in layout with a clear page title, labeled username and password fields, and a prominent submit button.
- Keep the form usable at narrow mobile widths and desktop widths.
- Associate labels and validation messages with their inputs, preserve keyboard navigation, and expose asynchronous snackbar feedback to assistive technology through notistack's defaults.
- Do not expose the demo password in visible application copy.

## Validation Rules

- Username is required.
- Password is required.
- Authentication succeeds only when both values exactly match `admin` and `admin123456Aa@`.
- Failed credentials produce a generic error message that does not reveal which field was incorrect.

## Testing and Acceptance

Playwright must verify:

1. Empty submission reports required-field errors.
2. Incorrect credentials show an error snackbar and remain on `/login`.
3. Correct credentials show success feedback and navigate to `/dashboard`.
4. Direct access to `/dashboard` without a session redirects to `/login`.
5. Signing out removes access to `/dashboard`.

The project must provide scripts for development, production build, Playwright tests, and Playwright browser installation if needed. The production build and requested Playwright checks must pass.

## Out of Scope

- A real backend, secure authorization, account management, password reset, or production authentication.
- Additional admin dashboard modules beyond a minimal protected landing page and sign-out control.
- Persistent login across browser sessions.
