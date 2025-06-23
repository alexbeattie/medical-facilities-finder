import { canAccessAdmin, hasAnyRole, hasAnyPermission } from './auth0'

// Development mode bypass (remove in production)
const isDevelopment = process.env.NODE_ENV === 'development'
const bypassAuth = isDevelopment && (!process.env.VUE_APP_AUTH0_DOMAIN || process.env.VUE_APP_AUTH0_DOMAIN === 'your-tenant.auth0.com')

// Authentication guard - checks if user is logged in
export const authGuard = (to, from, next) => {
  if (bypassAuth) {
    console.log('🚧 Development mode: Bypassing authentication')
    next()
    return
  }
  const auth0 = window.Vue?.config?.globalProperties?.$auth0

  if (!auth0) {
    console.error('Auth0 instance not found')
    next('/login')
    return
  }

  if (auth0.isLoading) {
    // Wait for auth to finish loading (Vue 3 compatible)
    const checkAuthStatus = () => {
      setTimeout(() => {
        if (!auth0.isLoading) {
          checkAuth()
        } else {
          checkAuthStatus()
        }
      }, 100)
    }
    checkAuthStatus()
  } else {
    checkAuth()
  }

  function checkAuth() {
    if (!auth0.isAuthenticated) {
      // Redirect to login
      auth0.loginWithRedirect({
        redirectUri: `${window.location.origin}/admin/callback`,
        returnTo: to.fullPath
      })
    } else {
      next()
    }
  }
}

// Admin access guard - checks if user can access admin panel
export const adminGuard = (to, from, next) => {
  if (bypassAuth) {
    console.log('🚧 Development mode: Bypassing admin guard')
    next()
    return
  }

  const auth0 = window.Vue?.config?.globalProperties?.$auth0

  if (!auth0 || !auth0.isAuthenticated) {
    next('/admin/login')
    return
  }

  const user = auth0.user
  if (!canAccessAdmin(user)) {
    // User doesn't have admin access
    next('/admin/unauthorized')
    return
  }

  next()
}

// Role-based guard factory
export const createRoleGuard = (requiredRoles) => {
  return (to, from, next) => {
    const auth0 = window.Vue?.config?.globalProperties?.$auth0

    if (!auth0 || !auth0.isAuthenticated) {
      next('/admin/login')
      return
    }

    const user = auth0.user
    if (!hasAnyRole(user, requiredRoles)) {
      // User doesn't have required role
      next('/admin/forbidden')
      return
    }

    next()
  }
}

// Permission-based guard factory
export const createPermissionGuard = (requiredPermissions) => {
  return (to, from, next) => {
    const auth0 = window.Vue?.config?.globalProperties?.$auth0

    if (!auth0 || !auth0.isAuthenticated) {
      next('/admin/login')
      return
    }

    const user = auth0.user
    if (!hasAnyPermission(user, requiredPermissions)) {
      // User doesn't have required permission
      next('/admin/forbidden')
      return
    }

    next()
  }
}

// Combined auth and admin guard for convenience
export const adminAuthGuard = (to, from, next) => {
  if (bypassAuth) {
    console.log('🚧 Development mode: Bypassing admin auth guard')
    next()
    return
  }

  authGuard(to, from, (path) => {
    if (typeof path === 'string') {
      next(path)
    } else {
      adminGuard(to, from, next)
    }
  })
} 