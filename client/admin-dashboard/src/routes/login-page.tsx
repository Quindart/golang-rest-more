import { useForm } from 'react-hook-form'
import { useNavigate } from 'react-router-dom'
import { useSnackbar } from 'notistack'
import { ArrowRight, ShieldCheck } from 'lucide-react'
import { login } from '@/lib/auth'
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

interface LoginFormValues {
  username: string
  password: string
}

export function LoginPage() {
  const navigate = useNavigate()
  const { enqueueSnackbar } = useSnackbar()
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginFormValues>()

  const onSubmit = handleSubmit(async ({ username, password }) => {
    if (!login(username, password)) {
      enqueueSnackbar('Invalid username or password.', { variant: 'error' })
      return
    }

    enqueueSnackbar('Welcome back!', { variant: 'success' })
    navigate('/dashboard', { replace: true })
  })

  return (
    <main className="grid min-h-svh place-items-center bg-stone-100 px-4 py-10 text-stone-900 sm:px-6">
      <div className="w-full max-w-md">
        <div className="mb-8 flex items-center justify-center gap-3">
          <span className="grid size-10 place-items-center rounded-xl bg-stone-900 text-white shadow-sm">
            <ShieldCheck aria-hidden="true" className="size-5" />
          </span>
          <span className="text-lg font-semibold tracking-tight">Northstar Admin</span>
        </div>

        <Card className="border-stone-200/80 bg-white shadow-xl shadow-stone-900/5">
          <CardHeader className="gap-2 px-6 pt-7 sm:px-8 sm:pt-8">
            <CardTitle>
              <h1 className="text-2xl font-semibold tracking-tight">Welcome back</h1>
            </CardTitle>
            <CardDescription className="text-sm leading-relaxed">
              Sign in to continue to your admin workspace.
            </CardDescription>
          </CardHeader>
          <CardContent className="px-6 pb-7 sm:px-8 sm:pb-8">
            <form className="space-y-5" noValidate onSubmit={onSubmit}>
              <div className="space-y-2">
                <Label htmlFor="username">Username</Label>
                <Input
                  autoComplete="username"
                  id="username"
                  aria-describedby={errors.username ? 'username-error' : undefined}
                  aria-invalid={Boolean(errors.username)}
                  placeholder="Enter your username"
                  {...register('username', { required: 'Username is required.' })}
                />
                {errors.username && (
                  <p id="username-error" className="text-sm text-destructive" role="alert">
                    {errors.username.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input
                  autoComplete="current-password"
                  id="password"
                  aria-describedby={errors.password ? 'password-error' : undefined}
                  aria-invalid={Boolean(errors.password)}
                  placeholder="Enter your password"
                  type="password"
                  {...register('password', { required: 'Password is required.' })}
                />
                {errors.password && (
                  <p id="password-error" className="text-sm text-destructive" role="alert">
                    {errors.password.message}
                  </p>
                )}
              </div>

              <Button className="h-11 w-full justify-between px-4" disabled={isSubmitting} type="submit">
                <span>{isSubmitting ? 'Signing in…' : 'Sign in'}</span>
                <ArrowRight aria-hidden="true" className="size-4" />
              </Button>
            </form>
          </CardContent>
        </Card>

        <p className="mt-6 text-center text-xs text-stone-500">
          Demo workspace · Access is limited to this browser tab
        </p>
      </div>
    </main>
  )
}
