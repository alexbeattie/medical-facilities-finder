<template>
  <div
    class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8"
  >
    <div class="max-w-md w-full space-y-8">
      <div>
        <div class="mx-auto h-12 w-12 text-red-500">
          <svg fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L5.636 5.636"
            />
          </svg>
        </div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Access Forbidden
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          You don't have permission to access this resource.
        </p>
      </div>
      <div class="text-center">
        <router-link
          to="/admin"
          class="font-medium text-blue-600 hover:text-blue-500"
        >
          Go to Admin Dashboard
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AdminForbidden",
};
</script>
