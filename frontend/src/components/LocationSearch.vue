<template>
  <div class="location-search bg-white rounded-lg shadow-md p-4 mb-4">
    <h3 class="text-lg font-semibold mb-4 text-gray-800">
      Find Facilities Near You
    </h3>

    <!-- Search Type Toggle -->
    <div class="flex mb-4 bg-gray-100 rounded-lg p-1">
      <button
        @click="searchType = 'location'"
        :class="{
          'bg-white shadow-sm': searchType === 'location',
          'text-gray-600': searchType !== 'location',
        }"
        class="flex-1 py-2 px-4 rounded-md transition-all duration-200 text-sm font-medium"
      >
        📍 Current Location
      </button>
      <button
        @click="searchType = 'address'"
        :class="{
          'bg-white shadow-sm': searchType === 'address',
          'text-gray-600': searchType !== 'address',
        }"
        class="flex-1 py-2 px-4 rounded-md transition-all duration-200 text-sm font-medium"
      >
        🔍 Search Address
      </button>
    </div>

    <!-- Address Search -->
    <div v-if="searchType === 'address'" class="mb-4">
      <div class="relative">
        <input
          v-model="addressQuery"
          @input="onAddressInput"
          @keyup.enter="searchByAddress"
          type="text"
          placeholder="Enter address, city, or ZIP code..."
          class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
        />
        <button
          @click="searchByAddress"
          :disabled="!addressQuery.trim() || addressLoading"
          class="absolute right-2 top-1/2 transform -translate-y-1/2 bg-blue-600 text-white px-4 py-1.5 rounded-md hover:bg-blue-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-200"
        >
          <span v-if="addressLoading" class="inline-block animate-spin">⟳</span>
          <span v-else>Search</span>
        </button>
      </div>

      <!-- Address suggestions -->
      <div
        v-if="addressSuggestions.length > 0"
        class="absolute z-50 w-full bg-white border border-gray-200 rounded-lg shadow-lg mt-1 max-h-60 overflow-y-auto"
      >
        <button
          v-for="(suggestion, index) in addressSuggestions"
          :key="index"
          @click="selectAddressSuggestion(suggestion)"
          class="w-full text-left px-4 py-3 hover:bg-gray-50 border-b border-gray-100 last:border-b-0 transition-colors duration-150"
        >
          <div class="font-medium text-gray-900">
            {{ suggestion.display_name || suggestion.formatted_address }}
          </div>
          <div class="text-sm text-gray-500">
            {{ suggestion.type || "Address" }}
          </div>
        </button>
      </div>
    </div>

    <!-- Current Location -->
    <div v-if="searchType === 'location'" class="mb-4">
      <button
        @click="getCurrentLocation"
        :disabled="locationLoading"
        class="w-full bg-green-600 text-white py-3 px-4 rounded-lg hover:bg-green-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-200 flex items-center justify-center space-x-2"
      >
        <span v-if="locationLoading" class="inline-block animate-spin">⟳</span>
        <span v-else>📍</span>
        <span>{{
          locationLoading ? "Getting Location..." : "Use Current Location"
        }}</span>
      </button>

      <div v-if="currentLocation" class="mt-3 p-3 bg-green-50 rounded-lg">
        <div class="flex items-center text-green-800">
          <span class="text-lg mr-2">✅</span>
          <div>
            <div class="font-medium">Location found!</div>
            <div class="text-sm">
              {{
                currentLocationAddress ||
                `${currentLocation.lat.toFixed(
                  4
                )}, ${currentLocation.lng.toFixed(4)}`
              }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Search Radius -->
    <div class="mb-4">
      <label class="block text-sm font-medium text-gray-700 mb-2">
        Search Radius: {{ searchRadius }} miles
      </label>
      <input
        v-model.number="searchRadius"
        type="range"
        min="1"
        max="100"
        step="1"
        class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer slider"
      />
      <div class="flex justify-between text-xs text-gray-500 mt-1">
        <span>1 mile</span>
        <span>100 miles</span>
      </div>
    </div>

    <!-- Facility Types Filter -->
    <div class="mb-4">
      <label class="block text-sm font-medium text-gray-700 mb-2">
        Facility Types
      </label>
      <div class="grid grid-cols-2 gap-2">
        <label
          v-for="type in facilityTypes"
          :key="type.key"
          class="flex items-center space-x-2 p-2 border border-gray-200 rounded-lg hover:bg-gray-50 cursor-pointer"
        >
          <input
            v-model="selectedTypes"
            :value="type.key"
            type="checkbox"
            class="rounded text-blue-600 focus:ring-blue-500"
          />
          <span class="text-sm">{{ type.label }}</span>
        </label>
      </div>
    </div>

    <!-- Search Button -->
    <button
      @click="performSearch"
      :disabled="!canSearch || searchLoading"
      class="w-full bg-blue-600 text-white py-3 px-4 rounded-lg hover:bg-blue-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-200 flex items-center justify-center space-x-2"
    >
      <span v-if="searchLoading" class="inline-block animate-spin">⟳</span>
      <span v-else>🔍</span>
      <span>{{ searchLoading ? "Searching..." : "Find Facilities" }}</span>
    </button>

    <!-- Error Message -->
    <div
      v-if="error"
      class="mt-4 p-3 bg-red-50 border border-red-200 rounded-lg"
    >
      <div class="flex items-center text-red-800">
        <span class="text-lg mr-2">⚠️</span>
        <div>
          <div class="font-medium">Error</div>
          <div class="text-sm">{{ error }}</div>
        </div>
      </div>
    </div>

    <!-- Search Results Summary -->
    <div
      v-if="searchResults"
      class="mt-4 p-3 bg-blue-50 border border-blue-200 rounded-lg"
    >
      <div class="text-blue-800">
        <div class="font-medium">Search Results</div>
        <div class="text-sm mt-1">
          Found {{ totalResults }} facilities within {{ searchRadius }} miles
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, watch } from "vue";
import { useFacilities } from "@/composables/useFacilities";
import { useGeocoding } from "@/composables/useGeocoding";

export default {
  name: "LocationSearch",
  emits: ["results-updated", "location-selected"],
  setup(props, { emit }) {
    const {
      searchNearby,
      loading: facilitiesLoading,
      error: facilitiesError,
      searchResults,
      getUserLocation,
      setUserLocation,
      userLocation,
    } = useFacilities();

    // Get API key for geocoding
    const apiKey = process.env.VUE_APP_GOOGLE_MAPS_API_KEY || "";
    const { geocodeAddress, reverseGeocode } = useGeocoding(apiKey);

    // Search state
    const searchType = ref("location");
    const addressQuery = ref("");
    const addressSuggestions = ref([]);
    const addressLoading = ref(false);
    const locationLoading = ref(false);
    const searchLoading = ref(false);
    const searchRadius = ref(25);
    const selectedTypes = ref(["resource_centers", "regional_centers"]);
    const currentLocation = ref(null);
    const currentLocationAddress = ref("");
    const error = ref("");

    // Facility types
    const facilityTypes = [
      { key: "resource_centers", label: "Resource Centers" },
      { key: "regional_centers", label: "Regional Centers" },
      { key: "resources", label: "Resources" },
      { key: "providers", label: "Service Providers" },
    ];

    // Computed properties
    const canSearch = computed(() => {
      return currentLocation.value && selectedTypes.value.length > 0;
    });

    const totalResults = computed(() => {
      if (!searchResults.value) return 0;
      return Object.values(searchResults.value).reduce((total, results) => {
        return total + (Array.isArray(results) ? results.length : 0);
      }, 0);
    });

    // Watch for location changes
    watch(
      userLocation,
      (newLocation) => {
        if (newLocation) {
          currentLocation.value = newLocation;
          searchType.value = "location";
        }
      },
      { immediate: true }
    );

    // Methods
    const getCurrentLocation = async () => {
      try {
        locationLoading.value = true;
        error.value = "";

        const location = await getUserLocation();
        currentLocation.value = location;

        // Try to get readable address
        try {
          const address = await reverseGeocode(location.lat, location.lng);
          currentLocationAddress.value = address;
        } catch (err) {
          console.warn("Could not get address for current location:", err);
        }

        emit("location-selected", location);
      } catch (err) {
        error.value = `Could not get your location: ${err.message}`;
        console.error("Geolocation error:", err);
      } finally {
        locationLoading.value = false;
      }
    };

    const onAddressInput = async () => {
      if (addressQuery.value.length < 3) {
        addressSuggestions.value = [];
        return;
      }

      try {
        addressLoading.value = true;
        // This would typically call a geocoding service for suggestions
        // For now, we'll just search on enter
        addressSuggestions.value = [];
      } catch (err) {
        console.error("Error getting address suggestions:", err);
      } finally {
        addressLoading.value = false;
      }
    };

    const searchByAddress = async () => {
      if (!addressQuery.value.trim()) return;

      try {
        addressLoading.value = true;
        error.value = "";

        const result = await geocodeAddress(addressQuery.value);
        if (result && result.lat && result.lng) {
          currentLocation.value = {
            lat: result.lat,
            lng: result.lng,
          };
          currentLocationAddress.value =
            result.display_name || addressQuery.value;
          setUserLocation(currentLocation.value);
          emit("location-selected", currentLocation.value);
          addressSuggestions.value = [];
        } else {
          error.value =
            "Could not find that address. Please try a different search.";
        }
      } catch (err) {
        error.value = `Error searching for address: ${err.message}`;
        console.error("Geocoding error:", err);
      } finally {
        addressLoading.value = false;
      }
    };

    const selectAddressSuggestion = (suggestion) => {
      currentLocation.value = {
        lat: suggestion.lat,
        lng: suggestion.lng,
      };
      currentLocationAddress.value =
        suggestion.display_name || suggestion.formatted_address;
      addressQuery.value =
        suggestion.display_name || suggestion.formatted_address;
      addressSuggestions.value = [];
      setUserLocation(currentLocation.value);
      emit("location-selected", currentLocation.value);
    };

    const performSearch = async () => {
      if (!canSearch.value) return;

      try {
        searchLoading.value = true;
        error.value = "";

        const results = await searchNearby(
          currentLocation.value.lat,
          currentLocation.value.lng,
          searchRadius.value,
          selectedTypes.value
        );

        emit("results-updated", results);
      } catch (err) {
        error.value = `Search failed: ${err.message}`;
        console.error("Search error:", err);
      } finally {
        searchLoading.value = false;
      }
    };

    const clearError = () => {
      error.value = "";
    };

    return {
      // Data
      searchType,
      addressQuery,
      addressSuggestions,
      addressLoading,
      locationLoading,
      searchLoading: computed(
        () => searchLoading.value || facilitiesLoading.value
      ),
      searchRadius,
      selectedTypes,
      currentLocation,
      currentLocationAddress,
      error: computed(() => error.value || facilitiesError.value),
      facilityTypes,
      searchResults,

      // Computed
      canSearch,
      totalResults,

      // Methods
      getCurrentLocation,
      onAddressInput,
      searchByAddress,
      selectAddressSuggestion,
      performSearch,
      clearError,
    };
  },
};
</script>

<style scoped>
.slider::-webkit-slider-thumb {
  appearance: none;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.slider::-moz-range-thumb {
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
  border: none;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}
</style>
