import { createRouter, createWebHistory } from 'vue-router'
import { adminAuthGuard, createRoleGuard, createPermissionGuard } from '@/auth/authGuard'
import { ROLES, PERMISSIONS } from '@/auth/auth0'

// Layout
import AdminLayout from '@/admin/components/layout/AdminLayout.vue'

// Views
import AdminDashboard from '@/admin/views/AdminDashboard.vue'

// Lazy-loaded admin views
const AdminLogin = () => import('@/admin/views/AdminLogin.vue')
const AdminCallback = () => import('@/admin/views/AdminCallback.vue')
const AdminUnauthorized = () => import('@/admin/views/AdminUnauthorized.vue')
const AdminForbidden = () => import('@/admin/views/AdminForbidden.vue')

// Submissions
const PendingSubmissions = () => import('@/admin/views/PendingSubmissions.vue')
const AllSubmissions = () => import('@/admin/views/AllSubmissions.vue')

// Data Management
const ABAcentersManager = () => import('@/admin/views/ABACentersManager.vue')
const RegionalCentersManager = () => import('@/admin/views/RegionalCentersManager.vue')
const ResourcesManager = () => import('@/admin/views/ResourcesManager.vue')
const ProvidersManager = () => import('@/admin/views/ProvidersManager.vue')

// Analytics
const AnalyticsOverview = () => import('@/admin/views/AnalyticsOverview.vue')
const ReportsManager = () => import('@/admin/views/ReportsManager.vue')

// System Management
const UserManagement = () => import('@/admin/views/UserManagement.vue')
const SystemSettings = () => import('@/admin/views/SystemSettings.vue')

// Profile
const AdminProfile = () => import('@/admin/views/AdminProfile.vue')

const routes = [
  // Public admin routes (no auth required)
  {
    path: '/admin/login',
    name: 'AdminLogin',
    component: AdminLogin,
    meta: { title: 'Admin Login', requiresAuth: false }
  },
  {
    path: '/admin/callback',
    name: 'AdminCallback',
    component: AdminCallback,
    meta: { title: 'Processing Login...', requiresAuth: false }
  },
  {
    path: '/admin/unauthorized',
    name: 'AdminUnauthorized',
    component: AdminUnauthorized,
    meta: { title: 'Unauthorized', requiresAuth: false }
  },
  {
    path: '/admin/forbidden',
    name: 'AdminForbidden',
    component: AdminForbidden,
    meta: { title: 'Access Forbidden', requiresAuth: false }
  },

  // Protected admin routes
  {
    path: '/admin',
    component: AdminLayout,
    beforeEnter: adminAuthGuard,
    children: [
      // Dashboard
      {
        path: '',
        name: 'AdminDashboard',
        component: AdminDashboard,
        meta: {
          title: 'Dashboard',
          requiresAuth: true,
          permissions: [PERMISSIONS.READ_SUBMISSIONS]
        }
      },

      // Submissions Management
      {
        path: 'submissions',
        meta: {
          title: 'Submissions',
          requiresAuth: true
        },
        children: [
          {
            path: 'pending',
            name: 'PendingSubmissions',
            component: PendingSubmissions,
            beforeEnter: createPermissionGuard([PERMISSIONS.READ_SUBMISSIONS, PERMISSIONS.APPROVE_SUBMISSIONS]),
            meta: {
              title: 'Pending Submissions',
              permissions: [PERMISSIONS.READ_SUBMISSIONS]
            }
          },
          {
            path: 'all',
            name: 'AllSubmissions',
            component: AllSubmissions,
            beforeEnter: createPermissionGuard([PERMISSIONS.READ_SUBMISSIONS]),
            meta: {
              title: 'All Submissions',
              permissions: [PERMISSIONS.READ_SUBMISSIONS]
            }
          },
          {
            path: '',
            redirect: 'pending'
          }
        ]
      },

      // Data Management
      {
        path: 'facilities',
        meta: {
          title: 'Facility Management',
          requiresAuth: true
        },
        beforeEnter: createPermissionGuard([PERMISSIONS.WRITE_FACILITIES, PERMISSIONS.EDIT_FACILITIES]),
        children: [
          {
            path: 'aba-centers',
            name: 'ABACentersManager',
            component: ABAcentersManager,
            meta: {
              title: 'ABA Centers',
              permissions: [PERMISSIONS.WRITE_FACILITIES]
            }
          },
          {
            path: 'aba-centers/new',
            name: 'NewABACenter',
            component: () => import('@/admin/views/NewABACenter.vue'),
            meta: {
              title: 'Add ABA Center',
              permissions: [PERMISSIONS.WRITE_FACILITIES]
            }
          },
          {
            path: 'regional-centers',
            name: 'RegionalCentersManager',
            component: RegionalCentersManager,
            meta: {
              title: 'Regional Centers',
              permissions: [PERMISSIONS.EDIT_FACILITIES]
            }
          },
          {
            path: 'resources',
            name: 'ResourcesManager',
            component: ResourcesManager,
            meta: {
              title: 'Resources',
              permissions: [PERMISSIONS.WRITE_FACILITIES]
            }
          },
          {
            path: 'providers',
            name: 'ProvidersManager',
            component: ProvidersManager,
            meta: {
              title: 'Providers',
              permissions: [PERMISSIONS.WRITE_FACILITIES]
            }
          },
          {
            path: '',
            redirect: 'aba-centers'
          }
        ]
      },

      // Analytics (Nurse Admin+ only)
      {
        path: 'analytics',
        meta: {
          title: 'Analytics',
          requiresAuth: true
        },
        beforeEnter: createPermissionGuard([PERMISSIONS.VIEW_ANALYTICS]),
        children: [
          {
            path: 'overview',
            name: 'AnalyticsOverview',
            component: AnalyticsOverview,
            meta: {
              title: 'Analytics Overview',
              permissions: [PERMISSIONS.VIEW_ANALYTICS]
            }
          },
          {
            path: 'reports',
            name: 'ReportsManager',
            component: ReportsManager,
            meta: {
              title: 'Reports',
              permissions: [PERMISSIONS.VIEW_ANALYTICS]
            }
          },
          {
            path: '',
            redirect: 'overview'
          }
        ]
      },

      // System Management (Super Admin only)
      {
        path: 'users',
        name: 'UserManagement',
        component: UserManagement,
        beforeEnter: createRoleGuard([ROLES.SUPER_ADMIN]),
        meta: {
          title: 'User Management',
          roles: [ROLES.SUPER_ADMIN]
        }
      },
      {
        path: 'settings',
        name: 'SystemSettings',
        component: SystemSettings,
        beforeEnter: createRoleGuard([ROLES.SUPER_ADMIN]),
        meta: {
          title: 'System Settings',
          roles: [ROLES.SUPER_ADMIN]
        }
      },

      // Profile (all admin users)
      {
        path: 'profile',
        name: 'AdminProfile',
        component: AdminProfile,
        meta: {
          title: 'Profile',
          requiresAuth: true
        }
      }
    ]
  },

  // Redirect /admin/* to /admin for any unmatched routes
  {
    path: '/admin/:pathMatch(.*)*',
    redirect: '/admin'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Global navigation guard for title updates
router.beforeEach((to, from, next) => {
  // Update document title
  if (to.meta?.title) {
    document.title = `${to.meta.title} - Medical Facilities Admin`
  } else {
    document.title = 'Medical Facilities Admin'
  }

  next()
})

// Global error handling
router.onError((error) => {
  console.error('Router error:', error)

  // Handle chunk loading errors (lazy loading failures)
  if (error.message.includes('Loading chunk')) {
    // Reload the page to get the latest chunks
    window.location.reload()
  }
})

export default router 