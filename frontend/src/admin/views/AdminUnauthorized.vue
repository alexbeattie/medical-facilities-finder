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
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.732-.833-2.5 0L4.268 18.5c-.77.833.192 2.5 1.732 2.5z"
            />
          </svg>
        </div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Access Denied
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          You need to be logged in to access the admin panel.
        </p>
      </div>
      <div class="text-center">
        <router-link
          to="/admin/login"
          class="font-medium text-blue-600 hover:text-blue-500"
        >
          Go to Admin Login
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AdminUnauthorized",
};
</script>
