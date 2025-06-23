<template>
  <div>
    <!-- Page header -->
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-gray-900">Dashboard</h1>
      <p class="mt-1 text-sm text-gray-500">
        Welcome back, {{ user?.name }}! Here's what's happening with your
        medical facilities.
      </p>
    </div>

    <!-- Stats overview -->
    <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4 mb-8">
      <!-- Pending submissions -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <InboxIcon class="h-6 w-6 text-gray-400" aria-hidden="true" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Pending Reviews
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ stats.pendingSubmissions }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <router-link
              to="/admin/submissions/pending"
              class="font-medium text-blue-700 hover:text-blue-900"
            >
              View all
            </router-link>
          </div>
        </div>
      </div>

      <!-- Total ABA centers -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <BuildingOfficeIcon
                class="h-6 w-6 text-gray-400"
                aria-hidden="true"
              />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  ABA Centers
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ stats.abaCenters }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <router-link
              to="/admin/facilities/aba-centers"
              class="font-medium text-blue-700 hover:text-blue-900"
            >
              Manage
            </router-link>
          </div>
        </div>
      </div>

      <!-- Regional centers -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <MapPinIcon class="h-6 w-6 text-gray-400" aria-hidden="true" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Regional Centers
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ stats.regionalCenters }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <router-link
              to="/admin/facilities/regional-centers"
              class="font-medium text-blue-700 hover:text-blue-900"
            >
              Manage
            </router-link>
          </div>
        </div>
      </div>

      <!-- This month's submissions -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <ChartBarIcon class="h-6 w-6 text-gray-400" aria-hidden="true" />
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  This Month
                </dt>
                <dd class="text-lg font-medium text-gray-900">
                  {{ stats.monthlySubmissions }}
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-5 py-3">
          <div class="text-sm">
            <router-link
              to="/admin/analytics/overview"
              class="font-medium text-blue-700 hover:text-blue-900"
            >
              View analytics
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <!-- Quick actions and recent activity -->
    <div class="grid grid-cols-1 gap-6 lg:grid-cols-2">
      <!-- Quick actions -->
      <div class="bg-white shadow rounded-lg">
        <div class="p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Quick Actions</h3>
          <div class="space-y-3">
            <router-link
              to="/admin/submissions/pending"
              class="block p-3 border border-gray-200 rounded-md hover:bg-gray-50 transition-colors"
            >
              <div class="flex items-center">
                <InboxIcon class="h-5 w-5 text-blue-500 mr-3" />
                <div>
                  <div class="text-sm font-medium text-gray-900">
                    Review Pending Submissions
                  </div>
                  <div class="text-sm text-gray-500">
                    {{ stats.pendingSubmissions }} items waiting for review
                  </div>
                </div>
              </div>
            </router-link>

            <router-link
              v-if="canManageFacilities"
              to="/admin/facilities/aba-centers/new"
              class="block p-3 border border-gray-200 rounded-md hover:bg-gray-50 transition-colors"
            >
              <div class="flex items-center">
                <PlusIcon class="h-5 w-5 text-green-500 mr-3" />
                <div>
                  <div class="text-sm font-medium text-gray-900">
                    Add New ABA Center
                  </div>
                  <div class="text-sm text-gray-500">
                    Manually add a new facility
                  </div>
                </div>
              </div>
            </router-link>

            <router-link
              v-if="canViewAnalytics"
              to="/admin/analytics/reports"
              class="block p-3 border border-gray-200 rounded-md hover:bg-gray-50 transition-colors"
            >
              <div class="flex items-center">
                <DocumentChartBarIcon class="h-5 w-5 text-purple-500 mr-3" />
                <div>
                  <div class="text-sm font-medium text-gray-900">
                    Generate Report
                  </div>
                  <div class="text-sm text-gray-500">
                    Create custom analytics reports
                  </div>
                </div>
              </div>
            </router-link>
          </div>
        </div>
      </div>

      <!-- Recent activity -->
      <div class="bg-white shadow rounded-lg">
        <div class="p-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            Recent Activity
          </h3>
          <div class="space-y-3">
            <div
              v-for="activity in recentActivity"
              :key="activity.id"
              class="flex items-start space-x-3"
            >
              <div class="flex-shrink-0">
                <div
                  :class="[
                    'h-8 w-8 rounded-full flex items-center justify-center',
                    activity.type === 'approved'
                      ? 'bg-green-100'
                      : activity.type === 'rejected'
                      ? 'bg-red-100'
                      : 'bg-blue-100',
                  ]"
                >
                  <CheckIcon
                    v-if="activity.type === 'approved'"
                    class="h-4 w-4 text-green-600"
                  />
                  <XMarkIcon
                    v-else-if="activity.type === 'rejected'"
                    class="h-4 w-4 text-red-600"
                  />
                  <DocumentTextIcon v-else class="h-4 w-4 text-blue-600" />
                </div>
              </div>
              <div class="min-w-0 flex-1">
                <p class="text-sm text-gray-900">
                  {{ activity.description }}
                </p>
                <p class="text-xs text-gray-500">
                  {{ formatTimeAgo(activity.timestamp) }}
                </p>
              </div>
            </div>
          </div>

          <div class="mt-4 pt-4 border-t border-gray-200">
            <router-link
              to="/admin/activity"
              class="text-sm font-medium text-blue-700 hover:text-blue-900"
            >
              View all activity →
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import {
  InboxIcon,
  BuildingOfficeIcon,
  MapPinIcon,
  ChartBarIcon,
  PlusIcon,
  DocumentChartBarIcon,
  CheckIcon,
  XMarkIcon,
  DocumentTextIcon,
} from "@heroicons/vue/24/outline";
import { hasAnyPermission, PERMISSIONS } from "@/auth/auth0";

export default {
  name: "AdminDashboard",
  components: {
    InboxIcon,
    BuildingOfficeIcon,
    MapPinIcon,
    ChartBarIcon,
    PlusIcon,
    DocumentChartBarIcon,
    CheckIcon,
    XMarkIcon,
    DocumentTextIcon,
  },
  setup() {
    const { user } = useAuth0();
    const stats = ref({
      pendingSubmissions: 0,
      abaCenters: 0,
      regionalCenters: 0,
      monthlySubmissions: 0,
    });
    const recentActivity = ref([]);

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

    // Load dashboard data
    const loadDashboardData = async () => {
      try {
        // TODO: Replace with actual API calls
        stats.value = {
          pendingSubmissions: 5,
          abaCenters: 42,
          regionalCenters: 21,
          monthlySubmissions: 18,
        };

        recentActivity.value = [
          {
            id: 1,
            type: "approved",
            description:
              'Approved ABA center submission for "Harmony Therapy Center"',
            timestamp: new Date(Date.now() - 2 * 60 * 60 * 1000), // 2 hours ago
          },
          {
            id: 2,
            type: "submitted",
            description:
              'New regional center update submitted for "San Diego Regional Center"',
            timestamp: new Date(Date.now() - 4 * 60 * 60 * 1000), // 4 hours ago
          },
          {
            id: 3,
            type: "rejected",
            description:
              "Rejected ABA center submission - incomplete information",
            timestamp: new Date(Date.now() - 6 * 60 * 60 * 1000), // 6 hours ago
          },
          {
            id: 4,
            type: "approved",
            description: "Approved regional center contact update",
            timestamp: new Date(Date.now() - 24 * 60 * 60 * 1000), // 1 day ago
          },
        ];
      } catch (error) {
        console.error("Failed to load dashboard data:", error);
      }
    };

    // Format timestamp for display
    const formatTimeAgo = (timestamp) => {
      const now = new Date();
      const diff = now - timestamp;
      const hours = Math.floor(diff / (1000 * 60 * 60));
      const days = Math.floor(hours / 24);

      if (days > 0) {
        return `${days} day${days > 1 ? "s" : ""} ago`;
      } else if (hours > 0) {
        return `${hours} hour${hours > 1 ? "s" : ""} ago`;
      } else {
        const minutes = Math.floor(diff / (1000 * 60));
        return `${minutes} minute${minutes > 1 ? "s" : ""} ago`;
      }
    };

    onMounted(() => {
      loadDashboardData();
    });

    return {
      user,
      stats,
      recentActivity,
      canManageFacilities,
      canViewAnalytics,
      formatTimeAgo,
    };
  },
};
</script>

<style scoped>
/* Add any additional styles here */
</style>
