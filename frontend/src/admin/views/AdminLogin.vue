<template>
  <div
    class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8"
  >
    <div class="max-w-md w-full space-y-8">
      <!-- Header -->
      <div>
        <div class="mx-auto h-12 w-auto flex justify-center">
          <img class="h-12 w-auto" src="/logo.png" alt="Medical Facilities" />
        </div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Admin Panel Access
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          Sign in to manage medical facilities data
        </p>
      </div>

      <!-- Login form -->
      <div class="mt-8 space-y-6">
        <div class="bg-white shadow-lg rounded-lg px-6 py-8">
          <!-- Auth0 login button -->
          <div class="space-y-4">
            <button
              @click="login"
              :disabled="loading"
              class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-200"
            >
              <span class="absolute left-0 inset-y-0 flex items-center pl-3">
                <LockClosedIcon
                  v-if="!loading"
                  class="h-5 w-5 text-blue-500 group-hover:text-blue-400"
                  aria-hidden="true"
                />
                <div
                  v-else
                  class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"
                ></div>
              </span>
              {{ loading ? "Signing in..." : "Sign in with Auth0" }}
            </button>

            <!-- Error message -->
            <div
              v-if="error"
              class="bg-red-50 border border-red-200 rounded-md p-4"
            >
              <div class="flex">
                <div class="flex-shrink-0">
                  <XCircleIcon
                    class="h-5 w-5 text-red-400"
                    aria-hidden="true"
                  />
                </div>
                <div class="ml-3">
                  <h3 class="text-sm font-medium text-red-800">
                    Authentication Error
                  </h3>
                  <div class="mt-2 text-sm text-red-700">
                    {{ error }}
                  </div>
                </div>
              </div>
            </div>

            <!-- Info message -->
            <div class="bg-blue-50 border border-blue-200 rounded-md p-4">
              <div class="flex">
                <div class="flex-shrink-0">
                  <InformationCircleIcon
                    class="h-5 w-5 text-blue-400"
                    aria-hidden="true"
                  />
                </div>
                <div class="ml-3">
                  <h3 class="text-sm font-medium text-blue-800">
                    Admin Access Required
                  </h3>
                  <div class="mt-2 text-sm text-blue-700">
                    <p>
                      You need admin privileges to access this panel. Contact
                      your system administrator if you believe you should have
                      access.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer links -->
        <div class="text-center space-y-2">
          <router-link to="/" class="text-sm text-blue-600 hover:text-blue-500">
            ← Back to public site
          </router-link>
        </div>
      </div>

      <!-- Help section -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Need Help?</h3>
        <div class="space-y-3 text-sm text-gray-600">
          <div class="flex items-start">
            <div class="flex-shrink-0">
              <QuestionMarkCircleIcon class="h-5 w-5 text-gray-400 mt-0.5" />
            </div>
            <div class="ml-3">
              <p>
                <strong>First time logging in?</strong>
                <br />
                Use your assigned admin credentials provided by IT.
              </p>
            </div>
          </div>

          <div class="flex items-start">
            <div class="flex-shrink-0">
              <ExclamationTriangleIcon class="h-5 w-5 text-yellow-400 mt-0.5" />
            </div>
            <div class="ml-3">
              <p>
                <strong>Access Issues?</strong>
                <br />
                Contact the system administrator or IT support.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { useRouter, useRoute } from "vue-router";
import {
  LockClosedIcon,
  XCircleIcon,
  InformationCircleIcon,
  QuestionMarkCircleIcon,
  ExclamationTriangleIcon,
} from "@heroicons/vue/24/outline";

export default {
  name: "AdminLogin",
  components: {
    LockClosedIcon,
    XCircleIcon,
    InformationCircleIcon,
    QuestionMarkCircleIcon,
    ExclamationTriangleIcon,
  },
  setup() {
    const { loginWithRedirect, isAuthenticated, isLoading } = useAuth0();
    const router = useRouter();
    const route = useRoute();
    const loading = ref(false);
    const error = ref("");

    const login = async () => {
      loading.value = true;
      error.value = "";

      try {
        const returnTo = route.query.returnTo || "/admin";

        await loginWithRedirect({
          redirectUri: `${window.location.origin}/admin/callback`,
          returnTo,
          // Add admin-specific scope and audience
          scope: "openid profile email read:admin write:admin",
          audience: process.env.VUE_APP_AUTH0_AUDIENCE,
        });
      } catch (err) {
        console.error("Login error:", err);
        error.value =
          err.message || "Failed to initiate login. Please try again.";
        loading.value = false;
      }
    };

    onMounted(() => {
      // If already authenticated, redirect to admin dashboard
      if (isAuthenticated.value) {
        const returnTo = route.query.returnTo || "/admin";
        router.push(returnTo);
      }
    });

    return {
      loading: loading.value || isLoading,
      error,
      login,
    };
  },
};
</script>

<style scoped>
/* Additional custom styles can be added here */
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
