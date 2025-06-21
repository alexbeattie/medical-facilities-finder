import { createApp } from 'vue';
import App from './App.vue';
import store from './store';
import "core-js/stable";
import "regenerator-runtime/runtime";
import './assets/main.css';

console.log('🚀 Vue App Starting...');
console.log('Environment:', process.env.NODE_ENV);
console.log('API Base URL:', process.env.VUE_APP_API_BASE_URL);

const app = createApp(App);

// Initialize store settings with logging
console.log('📦 Initializing Vuex store...');
Promise.all([
  store.dispatch('mapSettings/loadSettings'),
  store.dispatch('devices/loadGeocodeCache')
]).then(() => {
  console.log('✅ Store initialization completed');
}).catch((error) => {
  console.error('❌ Store initialization failed:', error);
});

app.use(store);
console.log('🎯 Mounting Vue app to #app...');
app.mount('#app');
console.log('✅ Vue app mounted successfully!');

// Persist cache before unloading
window.addEventListener('beforeunload', () => {
  store.dispatch('devices/persistGeocodeCache');
});