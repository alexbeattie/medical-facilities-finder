<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Mobile sidebar overlay -->
    <div
      v-if="sidebarOpen"
      class="fixed inset-0 z-40 lg:hidden"
      @click="sidebarOpen = false"
    >
      <div class="fixed inset-0 bg-gray-600 bg-opacity-75"></div>
    </div>

    <!-- Mobile sidebar -->
    <div
      :class="[
        'fixed inset-y-0 left-0 z-50 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out lg:hidden',
        sidebarOpen ? 'translate-x-0' : '-translate-x-full',
      ]"
    >
      <AdminSidebar @close="sidebarOpen = false" />
    </div>

    <!-- Desktop sidebar -->
    <div
      class="hidden lg:fixed lg:inset-y-0 lg:left-0 lg:z-50 lg:block lg:w-64 lg:bg-white lg:shadow-lg"
    >
      <AdminSidebar />
    </div>

    <!-- Main content -->
    <div class="lg:pl-64">
      <!-- Header -->
      <AdminHeader @toggle-sidebar="sidebarOpen = !sidebarOpen" />

      <!-- Page content -->
      <main class="py-6">
        <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <!-- Loading spinner -->
          <div
            v-if="$auth0.isLoading"
            class="flex justify-center items-center h-64"
          >
            <div
              class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"
            ></div>
          </div>

          <!-- Error message -->
          <div
            v-else-if="error"
            class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6"
          >
            <div class="flex">
              <div class="flex-shrink-0">
                <svg
                  class="h-5 w-5 text-red-400"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
              <div class="ml-3">
                <h3 class="text-sm font-medium text-red-800">Error</h3>
                <div class="mt-2 text-sm text-red-700">
                  {{ error }}
                </div>
              </div>
            </div>
          </div>

          <!-- Main content area -->
          <div v-else>
            <router-view />
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { canAccessAdmin } from "@/auth/auth0";
import AdminHeader from "./AdminHeader.vue";
import AdminSidebar from "./AdminSidebar.vue";

export default {
  name: "AdminLayout",
  components: {
    AdminHeader,
    AdminSidebar,
  },
  setup() {
    const { user, isAuthenticated } = useAuth0();
    const sidebarOpen = ref(false);
    const error = ref("");

    onMounted(() => {
      // Check if user has admin access
      if (isAuthenticated.value && !canAccessAdmin(user.value)) {
        error.value = "You do not have permission to access the admin panel.";
      }
    });

    return {
      sidebarOpen,
      error,
    };
  },
};
</script>

<style scoped>
/* Custom scrollbar for sidebar */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style>
