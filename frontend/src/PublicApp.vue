<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <h1 class="text-xl font-bold text-gray-900">
              🏥 ABA & Medical Facilities Finder
            </h1>
          </div>
          <div class="flex items-center space-x-4">
            <div class="text-sm text-gray-500">
              Find ABA centers, resources, and support services near you
            </div>
            <!-- Admin Link -->
            <router-link
              to="/admin"
              class="text-sm text-blue-600 hover:text-blue-800 font-medium"
            >
              Admin
            </router-link>
          </div>
        </div>
      </div>
    </header>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Left Sidebar - Search and Filters -->
        <div class="lg:col-span-1 space-y-6">
          <!-- Location Search -->
          <LocationSearch
            @results-updated="onSearchResults"
            @location-selected="onLocationSelected"
          />

          <!-- Facility List -->
          <FacilityList
            :facilities="searchResults"
            :loading="loading"
            :user-location="userLocation"
            @facility-select="onFacilitySelect"
          />
        </div>

        <!-- Main Content Area -->
        <div class="lg:col-span-2">
          <!-- Welcome/Instructions -->
          <div
            v-if="!searchResults"
            class="bg-white rounded-lg shadow-md p-8 text-center"
          >
            <div class="text-6xl mb-4">🔍</div>
            <h2 class="text-2xl font-bold text-gray-900 mb-4">
              Find ABA Centers & Resources
            </h2>
            <p class="text-gray-600 mb-6 max-w-md mx-auto">
              Search for ABA therapy centers, resource centers, regional
              centers, and service providers in your area.
            </p>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 max-w-lg mx-auto">
              <div class="flex items-center p-4 bg-blue-50 rounded-lg">
                <span class="text-2xl mr-3">🏥</span>
                <div class="text-left">
                  <div class="font-medium text-blue-900">ABA Centers</div>
                  <div class="text-sm text-blue-700">Therapy services</div>
                </div>
              </div>
              <div class="flex items-center p-4 bg-green-50 rounded-lg">
                <span class="text-2xl mr-3">🏢</span>
                <div class="text-left">
                  <div class="font-medium text-green-900">Resource Centers</div>
                  <div class="text-sm text-green-700">Support services</div>
                </div>
              </div>
              <div class="flex items-center p-4 bg-purple-50 rounded-lg">
                <span class="text-2xl mr-3">🌍</span>
                <div class="text-left">
                  <div class="font-medium text-purple-900">
                    Regional Centers
                  </div>
                  <div class="text-sm text-purple-700">Area coordination</div>
                </div>
              </div>
              <div class="flex items-center p-4 bg-orange-50 rounded-lg">
                <span class="text-2xl mr-3">👥</span>
                <div class="text-left">
                  <div class="font-medium text-orange-900">Providers</div>
                  <div class="text-sm text-orange-700">Individual services</div>
                </div>
              </div>
            </div>

            <!-- Contribution Buttons -->
            <div class="mt-8 border-t border-gray-200 pt-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4 text-center">
                Help Us Improve Our Directory
              </h3>
              <div class="flex flex-col sm:flex-row gap-4 justify-center">
                <button
                  @click="openABASubmissionForm"
                  class="bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors duration-200 flex items-center justify-center space-x-2"
                >
                  <span>🏥</span>
                  <span>Submit New ABA Center</span>
                </button>
                <button
                  @click="openRegionalCenterUpdateForm"
                  class="bg-purple-600 text-white px-6 py-3 rounded-lg hover:bg-purple-700 transition-colors duration-200 flex items-center justify-center space-x-2"
                >
                  <span>📋</span>
                  <span>Update Regional Center Info</span>
                </button>
              </div>
              <p class="text-sm text-gray-500 text-center mt-3">
                Know of a center we're missing? Have updated information? Help
                us keep our directory accurate!
              </p>
            </div>
          </div>

          <!-- Interactive Map -->
          <FacilityMapbox
            v-if="searchResults"
            ref="mapboxComponent"
            :facilities="searchResults"
            :user-location="userLocation"
            :loading="loading"
            height="600px"
            @facility-select="onFacilitySelect"
          />

          <!-- Map Placeholder (when no search results) -->
          <div v-else class="bg-white rounded-lg shadow-md p-8 text-center">
            <div class="text-4xl mb-4">🗺️</div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">
              Interactive Map
            </h3>
            <p class="text-gray-600 mb-4">
              Search for facilities to view them on the map with color-coded
              markers.
            </p>
            <div class="text-sm text-gray-500">
              <strong>Interactive Mapbox map</strong> with color-coded markers,
              directions, and facility details.
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Facility Details Popup -->
    <FacilityPopup
      :facility="selectedFacility"
      :facility-type="selectedFacilityType"
      :is-visible="showFacilityPopup"
      :user-location="userLocation"
      @close="closeFacilityPopup"
      @contact-form="showContactForm"
      @get-directions="onGetDirections"
    />

    <!-- Contact Form -->
    <ContactForm
      :facility="contactFacility"
      :facility-type="contactFacilityType"
      :is-visible="showContactFormModal"
      @close="closeContactForm"
      @submitted="onContactFormSubmitted"
    />

    <!-- ABA Submission Form -->
    <ABASubmissionForm
      :is-visible="showABASubmissionForm"
      @close="closeABASubmissionForm"
      @submitted="onABASubmissionCompleted"
    />

    <!-- Regional Center Update Form -->
    <RegionalCenterUpdateForm
      :is-visible="showRegionalCenterUpdateForm"
      :pre-selected-center="selectedFacility"
      @close="closeRegionalCenterUpdateForm"
      @submitted="onRegionalCenterUpdateCompleted"
    />

    <!-- Toast Notifications -->
    <div v-if="showToast" class="fixed bottom-4 right-4 z-50">
      <div
        class="bg-green-500 text-white px-6 py-3 rounded-lg shadow-lg flex items-center space-x-2"
      >
        <span>✅</span>
        <span>{{ toastMessage }}</span>
        <button @click="showToast = false" class="ml-2 text-white">×</button>
      </div>
    </div>

    <!-- Error Messages -->
    <div
      v-if="error"
      class="fixed top-20 left-1/2 transform -translate-x-1/2 z-50"
    >
      <div
        class="bg-red-500 text-white px-6 py-3 rounded-lg shadow-lg flex items-center space-x-2"
      >
        <span>⚠️</span>
        <span>{{ error }}</span>
        <button @click="clearError" class="ml-2 text-white">×</button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import { useFacilities } from "@/composables/useFacilities";
import LocationSearch from "@/components/LocationSearch.vue";
import FacilityList from "@/components/FacilityList.vue";
import FacilityMapbox from "@/components/FacilityMapbox.vue";
import FacilityPopup from "@/components/FacilityPopup.vue";
import ContactForm from "@/components/ContactForm.vue";
import ABASubmissionForm from "@/components/ABASubmissionForm.vue";
import RegionalCenterUpdateForm from "@/components/RegionalCenterUpdateForm.vue";

export default {
  name: "PublicApp",
  components: {
    LocationSearch,
    FacilityList,
    FacilityMapbox,
    FacilityPopup,
    ContactForm,
    ABASubmissionForm,
    RegionalCenterUpdateForm,
  },
  setup() {
    const { loading, error, userLocation, searchResults, clearError } =
      useFacilities();

    // Local state
    const selectedFacility = ref(null);
    const selectedFacilityType = ref("");
    const showFacilityPopup = ref(false);
    const contactFacility = ref(null);
    const contactFacilityType = ref("");
    const showContactFormModal = ref(false);
    const showToast = ref(false);
    const toastMessage = ref("");
    const mapboxComponent = ref(null);

    // New form states
    const showABASubmissionForm = ref(false);
    const showRegionalCenterUpdateForm = ref(false);

    // Computed
    const totalResults = computed(() => {
      if (!searchResults.value) return 0;

      const counts = [
        searchResults.value.resource_centers?.length || 0,
        searchResults.value.regional_centers?.length || 0,
        searchResults.value.resources?.length || 0,
        searchResults.value.providers?.length || 0,
      ];

      return counts.reduce((total, count) => total + count, 0);
    });

    // Convert backend search results to component format
    const searchResultsFormatted = computed(() => {
      if (!searchResults.value) return null;

      return {
        abaCenters: searchResults.value.aba_centers || [],
        resourceCenters: searchResults.value.resource_centers || [],
        regionalCenters: searchResults.value.regional_centers || [],
        resources: searchResults.value.resources || [],
        providers: searchResults.value.providers || [],
      };
    });

    // Methods
    const onSearchResults = (results) => {
      console.log("Search results received:", results);
    };

    const onLocationSelected = (location) => {
      console.log("Location selected:", location);
    };

    const onFacilitySelect = (facility, facilityType) => {
      selectedFacility.value = facility;
      selectedFacilityType.value = facilityType;
      showFacilityPopup.value = true;
    };

    const closeFacilityPopup = () => {
      showFacilityPopup.value = false;
      selectedFacility.value = null;
      selectedFacilityType.value = "";
    };

    const showContactForm = (facility) => {
      contactFacility.value = facility;
      contactFacilityType.value = selectedFacilityType.value;
      showContactFormModal.value = true;
      closeFacilityPopup();
    };

    const closeContactForm = () => {
      showContactFormModal.value = false;
      contactFacility.value = null;
      contactFacilityType.value = "";
    };

    const onContactFormSubmitted = () => {
      toastMessage.value = "Your inquiry has been submitted successfully!";
      showToast.value = true;

      setTimeout(() => {
        showToast.value = false;
      }, 5000);

      closeContactForm();
    };

    const showToastMessage = (message) => {
      toastMessage.value = message;
      showToast.value = true;
      setTimeout(() => {
        showToast.value = false;
      }, 5000);
    };

    const onGetDirections = (facility) => {
      console.log("🗺️ App: Triggering in-app directions for:", facility);

      if (mapboxComponent.value) {
        mapboxComponent.value.getDirections(facility);
        showToastMessage(
          `🧭 Getting directions to ${
            facility.name || facility.regional_center || "facility"
          }`
        );
      } else {
        console.error("Mapbox component not available");
        showToastMessage("❌ Map not ready for directions");
      }
    };

    // Form submission handlers
    const openABASubmissionForm = () => {
      showABASubmissionForm.value = true;
    };

    const closeABASubmissionForm = () => {
      showABASubmissionForm.value = false;
    };

    const onABASubmissionCompleted = () => {
      showToastMessage(
        "🏥 ABA center submission received! Thank you for contributing."
      );
    };

    const openRegionalCenterUpdateForm = () => {
      showRegionalCenterUpdateForm.value = true;
    };

    const closeRegionalCenterUpdateForm = () => {
      showRegionalCenterUpdateForm.value = false;
    };

    const onRegionalCenterUpdateCompleted = () => {
      showToastMessage(
        "📋 Regional center update submitted! Thank you for helping us improve."
      );
    };

    onMounted(() => {
      console.log("Public App mounted");
    });

    return {
      loading,
      error,
      userLocation,
      searchResults: searchResultsFormatted,
      selectedFacility,
      selectedFacilityType,
      showFacilityPopup,
      contactFacility,
      contactFacilityType,
      showContactFormModal,
      showToast,
      toastMessage,
      mapboxComponent,
      showABASubmissionForm,
      showRegionalCenterUpdateForm,
      totalResults,
      onSearchResults,
      onLocationSelected,
      onFacilitySelect,
      closeFacilityPopup,
      showContactForm,
      closeContactForm,
      onContactFormSubmitted,
      showToastMessage,
      onGetDirections,
      clearError,
      openABASubmissionForm,
      closeABASubmissionForm,
      onABASubmissionCompleted,
      openRegionalCenterUpdateForm,
      closeRegionalCenterUpdateForm,
      onRegionalCenterUpdateCompleted,
    };
  },
};
</script>

<style scoped>
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
