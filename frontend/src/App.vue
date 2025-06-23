<template>
  <div id="app">
    <!-- Router view for all pages -->
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
/* Global styles */
#app {
  font-family: "Inter", -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
    "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

/* Ensure consistent box-sizing */
*,
*:before,
*:after {
  box-sizing: border-box;
}

/* Custom focus styles */
input:focus,
select:focus,
textarea:focus,
button:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

/* Smooth transitions */
.transition-all {
  transition: all 0.2s ease-in-out;
}

/* Toast animation */
.fixed.bottom-4.right-4 > div {
  animation: slideInRight 0.3s ease-out;
}

@keyframes slideInRight {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

/* Error animation */
.fixed.top-20 > div {
  animation: slideInDown 0.3s ease-out;
}

@keyframes slideInDown {
  from {
    transform: translateY(-100%);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
</style>
