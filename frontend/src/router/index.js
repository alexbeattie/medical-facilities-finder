import { createRouter, createWebHistory } from 'vue-router'

// Public components
import PublicApp from '@/PublicApp.vue'

// Admin routes
import adminRoutes from '@/admin/router/admin'

const routes = [
  // Public routes (existing functionality)
  {
    path: '/',
    name: 'Home',
    component: PublicApp,
    meta: { title: 'Medical Facilities Finder' }
  },

  // Import all admin routes
  ...adminRoutes.options.routes,

  // Catch-all route
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
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
    document.title = to.meta.title.includes('Admin')
      ? to.meta.title
      : `${to.meta.title} - Medical Facilities`
  } else {
    document.title = 'Medical Facilities'
  }

  next()
})

export default router 