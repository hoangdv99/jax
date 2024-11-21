import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'

export const requireAuth = (
  _to: RouteLocationNormalized,
  _from: RouteLocationNormalized,
  next: NavigationGuardNext
) => {
  const token = localStorage.getItem('authToken')
  if (token) {
    next()
  } else {
    next({ name: 'Login' })
  }
}
