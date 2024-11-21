import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { compareAsc } from 'date-fns'

export const requireAuth = (
  _to: RouteLocationNormalized,
  _from: RouteLocationNormalized,
  next: NavigationGuardNext
) => {
  const token = localStorage.getItem('authToken')
  const expiry = localStorage.getItem('authTokenExpiry') || ''

  if (!token || !expiry) {
    next({ name: 'Login' })
  }

  const isValidDate = compareAsc(expiry, new Date()) === 1
  if (!!token && isValidDate) {
    next()
  } else {
    localStorage.removeItem('authToken')
    localStorage.removeItem('authTokenExpiry')
    next({ name: 'Login' })
  }
}
