<template>
  <header class="bg-white shadow-sm border-b border-gray-200">
    <div class="flex h-16 justify-between items-center px-4 sm:px-6 lg:px-8">
      <!-- Mobile menu button -->
      <button
        type="button"
        class="lg:hidden -ml-0.5 -mt-0.5 h-12 w-12 inline-flex items-center justify-center rounded-md text-gray-500 hover:text-gray-900 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-blue-500"
        @click="$emit('toggle-sidebar')"
      >
        <span class="sr-only">Open sidebar</span>
        <Bars3Icon class="h-6 w-6" aria-hidden="true" />
      </button>

      <!-- Page title -->
      <div class="flex-1 lg:flex-none">
        <h1 class="text-xl font-semibold text-gray-900 lg:text-2xl">
          {{ pageTitle }}
        </h1>
      </div>

      <!-- Right side - User menu -->
      <div class="flex items-center space-x-4">
        <!-- Notifications (placeholder) -->
        <button
          type="button"
          class="rounded-full bg-white p-1 text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          <span class="sr-only">View notifications</span>
          <BellIcon class="h-6 w-6" aria-hidden="true" />
        </button>

        <!-- User dropdown -->
        <div class="relative">
          <button
            type="button"
            class="flex max-w-xs items-center rounded-full bg-white text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            @click="userMenuOpen = !userMenuOpen"
          >
            <span class="sr-only">Open user menu</span>
            <img
              class="h-8 w-8 rounded-full"
              :src="user?.picture || '/default-avatar.png'"
              :alt="user?.name || 'User'"
            />
          </button>

          <!-- User dropdown menu -->
          <div
            v-if="userMenuOpen"
            class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
          >
            <!-- User info -->
            <div class="px-4 py-2 border-b border-gray-200">
              <p class="text-sm font-medium text-gray-900">{{ user?.name }}</p>
              <p class="text-sm text-gray-500">{{ userRole }}</p>
            </div>

            <!-- Menu items -->
            <router-link
              to="/admin/profile"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              @click="userMenuOpen = false"
            >
              Your Profile
            </router-link>

            <router-link
              to="/admin/settings"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              @click="userMenuOpen = false"
            >
              Settings
            </router-link>

            <button
              @click="logout"
              class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
            >
              Sign out
            </button>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import { ref, computed, onMounted, onUnmounted } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { useRoute } from "vue-router";
import { Bars3Icon, BellIcon } from "@heroicons/vue/24/outline";
import { getUserRoleLevel } from "@/auth/auth0";

export default {
  name: "AdminHeader",
  components: {
    Bars3Icon,
    BellIcon,
  },
  emits: ["toggle-sidebar"],
  setup() {
    const { user, logout: auth0Logout } = useAuth0();
    const route = useRoute();
    const userMenuOpen = ref(false);

    // Development mode mock user
    const isDevelopment = process.env.NODE_ENV === "development";
    const bypassAuth =
      isDevelopment &&
      (!process.env.VUE_APP_AUTH0_DOMAIN ||
        process.env.VUE_APP_AUTH0_DOMAIN === "your-tenant.auth0.com");
    const mockUser = {
      name: "Development User",
      email: "dev@example.com",
      picture: "https://via.placeholder.com/32",
    };

    // Get page title from route meta or default
    const pageTitle = computed(() => {
      return route.meta?.title || "Admin Dashboard";
    });

    // Get user's role for display
    const userRole = computed(() => {
      const currentUser = bypassAuth ? mockUser : user.value;
      return bypassAuth ? "Developer" : getUserRoleLevel(currentUser);
    });

    // Get current user (mock or real)
    const currentUser = computed(() => {
      return bypassAuth ? mockUser : user.value;
    });

    const logout = () => {
      userMenuOpen.value = false;
      if (bypassAuth) {
        // In development mode, just redirect
        window.location.href = "/admin/login";
      } else {
        auth0Logout({
          returnTo: `${window.location.origin}/admin/login`,
        });
      }
    };

    // Handle click outside to close dropdown
    const handleClickOutside = (event) => {
      const dropdown = event.target.closest(".relative");
      if (!dropdown && userMenuOpen.value) {
        userMenuOpen.value = false;
      }
    };

    onMounted(() => {
      document.addEventListener("click", handleClickOutside);
    });

    onUnmounted(() => {
      document.removeEventListener("click", handleClickOutside);
    });

    return {
      user: currentUser,
      userMenuOpen,
      pageTitle,
      userRole,
      logout,
    };
  },
};
</script>

<style scoped>
/* Click away directive alternative using @click.away */
.click-away {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}
</style>
