<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <div class="max-w-md w-full space-y-8">
      <div class="text-center">
        <!-- Loading animation -->
        <div class="mx-auto h-16 w-16 relative">
          <div
            class="absolute inset-0 rounded-full border-4 border-blue-200"
          ></div>
          <div
            class="absolute inset-0 rounded-full border-4 border-blue-600 border-t-transparent animate-spin"
          ></div>
        </div>

        <h2 class="mt-6 text-center text-2xl font-bold text-gray-900">
          {{ status }}
        </h2>

        <p class="mt-2 text-center text-sm text-gray-600">
          {{ message }}
        </p>

        <!-- Error display -->
        <div
          v-if="error"
          class="mt-6 bg-red-50 border border-red-200 rounded-md p-4"
        >
          <div class="flex">
            <div class="flex-shrink-0">
              <XCircleIcon class="h-5 w-5 text-red-400" aria-hidden="true" />
            </div>
            <div class="ml-3">
              <h3 class="text-sm font-medium text-red-800">
                Authentication Failed
              </h3>
              <div class="mt-2 text-sm text-red-700">
                {{ error }}
              </div>
              <div class="mt-4">
                <button
                  @click="goToLogin"
                  class="text-sm font-medium text-red-600 hover:text-red-500"
                >
                  Try again
                </button>
              </div>
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
import { XCircleIcon } from "@heroicons/vue/24/outline";
import { canAccessAdmin } from "@/auth/auth0";

export default {
  name: "AdminCallback",
  components: {
    XCircleIcon,
  },
  setup() {
    const { isLoading, isAuthenticated, user, error: auth0Error } = useAuth0();
    const router = useRouter();
    const route = useRoute();

    const status = ref("Processing login...");
    const message = ref("Please wait while we verify your credentials.");
    const error = ref("");

    const handleAuthResult = async () => {
      try {
        status.value = "Verifying credentials...";
        message.value = "Checking your admin access permissions.";

        // Wait for auth0 to finish processing
        if (isLoading.value) {
          await new Promise((resolve) => {
            const unwatch = isLoading.value.$watch("isLoading", (loading) => {
              if (!loading) {
                unwatch();
                resolve();
              }
            });
          });
        }

        if (auth0Error.value) {
          throw new Error(auth0Error.value.message || "Authentication failed");
        }

        if (!isAuthenticated.value) {
          throw new Error("Authentication was not completed successfully");
        }

        // Check if user has admin access
        if (!canAccessAdmin(user.value)) {
          throw new Error(
            "You do not have admin privileges. Contact your administrator."
          );
        }

        status.value = "Success!";
        message.value = "Redirecting to admin dashboard...";

        // Get the return URL from query params or localStorage
        const returnTo =
          route.query.returnTo ||
          route.query.state ||
          localStorage.getItem("admin_return_to") ||
          "/admin";

        // Clear any stored return URL
        localStorage.removeItem("admin_return_to");

        // Small delay to show success message
        setTimeout(() => {
          router.push(returnTo);
        }, 1000);
      } catch (err) {
        console.error("Auth callback error:", err);
        status.value = "Authentication Failed";
        message.value = "There was a problem with your login.";
        error.value =
          err.message || "An unexpected error occurred during authentication.";
      }
    };

    const goToLogin = () => {
      router.push("/admin/login");
    };

    onMounted(() => {
      // Handle URL errors first
      if (route.query.error) {
        status.value = "Authentication Failed";
        message.value = "Login was cancelled or failed.";
        error.value = route.query.error_description || route.query.error;
        return;
      }

      // Process successful auth
      handleAuthResult();
    });

    return {
      status,
      message,
      error,
      goToLogin,
    };
  },
};
</script>

<style scoped>
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
