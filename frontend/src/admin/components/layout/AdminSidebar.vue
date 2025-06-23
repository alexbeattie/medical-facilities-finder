<template>
  <div class="flex h-full flex-col overflow-y-auto bg-white">
    <!-- Logo and branding -->
    <div
      class="flex h-16 flex-shrink-0 items-center px-4 border-b border-gray-200"
    >
      <img class="h-8 w-auto" src="/logo.png" alt="Medical Facilities" />
      <span class="ml-2 text-lg font-semibold text-gray-900">Admin Panel</span>

      <!-- Close button for mobile -->
      <button
        v-if="$props.showCloseButton"
        type="button"
        class="ml-auto lg:hidden -mr-2 h-8 w-8 inline-flex items-center justify-center rounded-md text-gray-500 hover:text-gray-900"
        @click="$emit('close')"
      >
        <span class="sr-only">Close sidebar</span>
        <XMarkIcon class="h-6 w-6" aria-hidden="true" />
      </button>
    </div>

    <!-- Navigation -->
    <nav class="mt-6 flex-1 space-y-1 px-3">
      <!-- Dashboard -->
      <router-link
        to="/admin"
        :class="navLinkClass('/admin')"
        exact-active-class="bg-blue-50 border-blue-500 text-blue-700"
      >
        <HomeIcon class="mr-3 h-5 w-5 flex-shrink-0" aria-hidden="true" />
        Dashboard
      </router-link>

      <!-- Submissions Management -->
      <div class="space-y-1">
        <div
          class="text-xs font-medium text-gray-500 uppercase tracking-wider px-3 py-2"
        >
          Submissions
        </div>

        <router-link
          to="/admin/submissions/pending"
          :class="navLinkClass('/admin/submissions')"
        >
          <InboxIcon class="mr-3 h-5 w-5 flex-shrink-0" aria-hidden="true" />
          Pending Review
          <span
            v-if="pendingCount > 0"
            class="ml-auto inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800"
          >
            {{ pendingCount }}
          </span>
        </router-link>

        <router-link
          to="/admin/submissions/all"
          :class="navLinkClass('/admin/submissions/all')"
        >
          <DocumentTextIcon
            class="mr-3 h-5 w-5 flex-shrink-0"
            aria-hidden="true"
          />
          All Submissions
        </router-link>
      </div>

      <!-- Data Management -->
      <div class="space-y-1">
        <div
          class="text-xs font-medium text-gray-500 uppercase tracking-wider px-3 py-2"
        >
          Data Management
        </div>

        <router-link
          v-if="canManageFacilities"
          to="/admin/facilities/aba-centers"
          :class="navLinkClass('/admin/facilities/aba-centers')"
        >
          <BuildingOfficeIcon
            class="mr-3 h-5 w-5 flex-shrink-0"
            aria-hidden="true"
          />
          ABA Centers
        </router-link>

        <router-link
          v-if="canManageFacilities"
          to="/admin/facilities/regional-centers"
          :class="navLinkClass('/admin/facilities/regional-centers')"
        >
          <MapPinIcon class="mr-3 h-5 w-5 flex-shrink-0" aria-hidden="true" />
          Regional Centers
        </router-link>

        <router-link
          v-if="canManageFacilities"
          to="/admin/facilities/resources"
          :class="navLinkClass('/admin/facilities/resources')"
        >
          <BookOpenIcon class="mr-3 h-5 w-5 flex-shrink-0" aria-hidden="true" />
          Resources
        </router-link>

        <router-link
          v-if="canManageFacilities"
          to="/admin/facilities/providers"
          :class="navLinkClass('/admin/facilities/providers')"
        >
          <UserGroupIcon
            class="mr-3 h-5 w-5 flex-shrink-0"
            aria-hidden="true"
          />
          Providers
        </router-link>
      </div>

      <!-- Analytics -->
      <div v-if="canViewAnalytics" class="space-y-1">
        <div
          class="text-xs font-medium text-gray-500 uppercase tracking-wider px-3 py-2"
        >
          Analytics
        </div>

        <router-link
          to="/admin/analytics/overview"
          :class="navLinkClass('/admin/analytics')"
        >
          <ChartBarIcon class="mr-3 h-5 w-5 flex-shrink-0" aria-hidden="true" />
          Overview
        </router-link>

        <router-link
          to="/admin/analytics/reports"
          :class="navLinkClass('/admin/analytics/reports')"
        >
          <DocumentChartBarIcon
            class="mr-3 h-5 w-5 flex-shrink-0"
            aria-hidden="true"
          />
          Reports
        </router-link>
      </div>

      <!-- System Management -->
      <div v-if="canManageSystem" class="space-y-1">
        <div
          class="text-xs font-medium text-gray-500 uppercase tracking-wider px-3 py-2"
        >
          System
        </div>

        <router-link to="/admin/users" :class="navLinkClass('/admin/users')">
          <UsersIcon class="mr-3 h-5 w-5 flex-shrink-0" aria-hidden="true" />
          User Management
        </router-link>

        <router-link
          to="/admin/settings"
          :class="navLinkClass('/admin/settings')"
        >
          <CogIcon class="mr-3 h-5 w-5 flex-shrink-0" aria-hidden="true" />
          Settings
        </router-link>
      </div>
    </nav>

    <!-- Footer with version info -->
    <div class="flex-shrink-0 p-4 border-t border-gray-200">
      <div class="text-xs text-gray-500 text-center">Admin Panel v1.0.0</div>
    </div>
  </div>
</template>

<script>
import { computed, ref, onMounted } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { useRoute } from "vue-router";
import {
  HomeIcon,
  InboxIcon,
  DocumentTextIcon,
  BuildingOfficeIcon,
  MapPinIcon,
  BookOpenIcon,
  UserGroupIcon,
  ChartBarIcon,
  DocumentChartBarIcon,
  UsersIcon,
  CogIcon,
  XMarkIcon,
} from "@heroicons/vue/24/outline";
import { hasAnyPermission, hasAnyRole, ROLES, PERMISSIONS } from "@/auth/auth0";

export default {
  name: "AdminSidebar",
  components: {
    HomeIcon,
    InboxIcon,
    DocumentTextIcon,
    BuildingOfficeIcon,
    MapPinIcon,
    BookOpenIcon,
    UserGroupIcon,
    ChartBarIcon,
    DocumentChartBarIcon,
    UsersIcon,
    CogIcon,
    XMarkIcon,
  },
  props: {
    showCloseButton: {
      type: Boolean,
      default: true,
    },
  },
  emits: ["close"],
  setup() {
    const { user } = useAuth0();
    const route = useRoute();
    const pendingCount = ref(0);

    // Permission checks
    const canManageFacilities = computed(() => {
      return hasAnyPermission(user.value, [
        PERMISSIONS.WRITE_FACILITIES,
        PERMISSIONS.EDIT_FACILITIES,
        PERMISSIONS.WRITE_ALL,
      ]);
    });

    const canViewAnalytics = computed(() => {
      return hasAnyPermission(user.value, [
        PERMISSIONS.VIEW_ANALYTICS,
        PERMISSIONS.READ_ALL,
      ]);
    });

    const canManageSystem = computed(() => {
      return (
        hasAnyRole(user.value, [ROLES.SUPER_ADMIN]) ||
        hasAnyPermission(user.value, [PERMISSIONS.MANAGE_USERS])
      );
    });

    // Navigation link styling
    const navLinkClass = (path) => {
      const isActive =
        route.path.startsWith(path) && path !== "/admin"
          ? true
          : route.path === path;
      const baseClasses =
        "group flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors duration-200";

      if (isActive) {
        return `${baseClasses} bg-blue-50 border-r-2 border-blue-500 text-blue-700`;
      }

      return `${baseClasses} text-gray-700 hover:bg-gray-50 hover:text-gray-900`;
    };

    // Load pending submissions count
    const loadPendingCount = async () => {
      try {
        // This will be implemented when we add the API endpoints
        // const response = await adminAPI.getPendingSubmissionsCount()
        // pendingCount.value = response.count
        pendingCount.value = 5; // Placeholder
      } catch (error) {
        console.error("Failed to load pending count:", error);
      }
    };

    onMounted(() => {
      loadPendingCount();
    });

    return {
      pendingCount,
      canManageFacilities,
      canViewAnalytics,
      canManageSystem,
      navLinkClass,
    };
  },
};
</script>

<style scoped>
/* Ensure icons are properly aligned */
.router-link-active {
  @apply bg-blue-50 text-blue-700;
}

/* Custom hover effects */
.group:hover .text-gray-400 {
  @apply text-gray-500;
}
</style>
