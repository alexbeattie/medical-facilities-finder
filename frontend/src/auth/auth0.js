import { createAuth0 } from '@auth0/auth0-vue'

// Auth0 configuration
export const auth0Config = {
  domain: process.env.VUE_APP_AUTH0_DOMAIN || 'your-tenant.auth0.com',
  clientId: process.env.VUE_APP_AUTH0_CLIENT_ID || 'your-client-id',
  audience: process.env.VUE_APP_AUTH0_AUDIENCE || 'medical-facilities-api',
  redirectUri: `${window.location.origin}/admin/callback`,
  cacheLocation: 'localstorage',
  useRefreshTokens: true,
  scope: 'openid profile email read:admin write:admin'
}

// Create Auth0 instance
export const auth0 = createAuth0(auth0Config)

// Role-based permissions
export const ROLES = {
  SUPER_ADMIN: 'super-admin',
  NURSE_ADMIN: 'nurse-admin',
  REVIEWER: 'reviewer'
}

export const PERMISSIONS = {
  READ_ALL: 'read:all',
  WRITE_ALL: 'write:all',
  DELETE_ALL: 'delete:all',
  MANAGE_USERS: 'manage:users',
  READ_SUBMISSIONS: 'read:submissions',
  WRITE_FACILITIES: 'write:facilities',
  APPROVE_SUBMISSIONS: 'approve:submissions',
  EDIT_FACILITIES: 'edit:facilities',
  VIEW_ANALYTICS: 'view:analytics',
  COMMENT_SUBMISSIONS: 'comment:submissions'
}

// Helper functions
export const getUserRoles = (user) => {
  if (!user) return []
  return user['https://medicalfacilities.com/roles'] || []
}

export const getUserPermissions = (user) => {
  if (!user) return []
  return user['https://medicalfacilities.com/permissions'] || []
}

export const hasRole = (user, role) => {
  const roles = getUserRoles(user)
  return roles.includes(role)
}

export const hasPermission = (user, permission) => {
  const permissions = getUserPermissions(user)
  return permissions.includes(permission)
}

export const hasAnyRole = (user, roles) => {
  const userRoles = getUserRoles(user)
  return roles.some(role => userRoles.includes(role))
}

export const hasAnyPermission = (user, permissions) => {
  const userPermissions = getUserPermissions(user)
  return permissions.some(permission => userPermissions.includes(permission))
}

// Check if user can access admin panel
export const canAccessAdmin = (user) => {
  return hasAnyRole(user, [ROLES.SUPER_ADMIN, ROLES.NURSE_ADMIN, ROLES.REVIEWER])
}

// Get user's highest role level (for UI display)
export const getUserRoleLevel = (user) => {
  const roles = getUserRoles(user)
  if (roles.includes(ROLES.SUPER_ADMIN)) return 'Super Admin'
  if (roles.includes(ROLES.NURSE_ADMIN)) return 'Nurse Admin'
  if (roles.includes(ROLES.REVIEWER)) return 'Reviewer'
  return 'No Admin Access'
} 