<template>
  <div class="facility-map bg-white rounded-lg shadow-md overflow-hidden">
    <!-- Map Header -->
    <div class="p-4 border-b border-gray-200 bg-gray-50">
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-semibold text-gray-800">
          🗺️ Facility Map <span class="text-purple-600">(Mapbox)</span>
        </h3>
        <div class="flex items-center space-x-4">
          <!-- Legend -->
          <div class="flex items-center space-x-3 text-xs">
            <div class="flex items-center space-x-1">
              <div class="w-3 h-3 bg-blue-500 rounded-full"></div>
              <span>ABA Centers</span>
            </div>
            <div class="flex items-center space-x-1">
              <div class="w-3 h-3 bg-green-500 rounded-full"></div>
              <span>Resource Centers</span>
            </div>
            <div class="flex items-center space-x-1">
              <div class="w-3 h-3 bg-purple-500 rounded-full"></div>
              <span>Regional Centers</span>
            </div>
            <div class="flex items-center space-x-1">
              <div class="w-3 h-3 bg-orange-500 rounded-full"></div>
              <span>Providers</span>
            </div>
          </div>

          <!-- Map Controls -->
          <div class="flex items-center space-x-2">
            <button
              @click="fitMapToFacilities"
              class="px-3 py-1 bg-blue-100 text-blue-700 text-xs rounded hover:bg-blue-200 transition-colors"
            >
              📍 Fit All
            </button>
            <button
              @click="centerOnUser"
              v-if="userLocation"
              class="px-3 py-1 bg-green-100 text-green-700 text-xs rounded hover:bg-green-200 transition-colors"
            >
              🎯 My Location
            </button>
            <button
              @click="clearDirections"
              v-if="showingDirections"
              class="px-3 py-1 bg-red-100 text-red-700 text-xs rounded hover:bg-red-200 transition-colors"
            >
              🧭 Clear Route
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Directions Panel -->
    <div
      v-if="showingDirections && currentDirections"
      class="p-4 border-b border-gray-200 bg-blue-50 max-h-40 overflow-y-auto"
    >
      <div class="flex items-center justify-between mb-2">
        <h4 class="font-semibold text-blue-900">
          🧭 From Your Current Location →
          {{ currentDirections.destination_name }}
        </h4>
        <button
          @click="clearDirections"
          class="text-blue-600 hover:text-blue-800 p-1"
        >
          ✕
        </button>
      </div>
      <div class="text-sm text-blue-800 mb-2">
        📍 Distance: {{ currentDirections.distance }} • ⏱️ Time:
        {{ currentDirections.duration }}
      </div>
      <div class="space-y-1 text-xs text-blue-700">
        <div
          v-for="(step, index) in currentDirections.steps"
          :key="index"
          class="flex items-start space-x-2"
        >
          <span class="text-blue-500 mt-0.5">{{ index + 1 }}.</span>
          <span v-html="step.instructions"></span>
        </div>
      </div>
    </div>

    <!-- Map Container -->
    <div class="relative" :style="{ height: mapHeight }">
      <!-- Loading State -->
      <div
        v-if="loading"
        class="absolute inset-0 flex items-center justify-center bg-gray-100 z-10"
      >
        <div class="text-center">
          <div
            class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-purple-600 mb-2"
          ></div>
          <p class="text-gray-600">Loading Mapbox...</p>
        </div>
      </div>

      <!-- Mapbox Container -->
      <div
        ref="mapContainer"
        class="w-full h-full"
        :class="{ 'opacity-0': loading }"
      ></div>

      <!-- No Token Warning -->
      <div
        v-if="!mapboxToken"
        class="absolute inset-0 flex items-center justify-center bg-gray-100"
      >
        <div class="text-center p-6">
          <div class="text-4xl mb-4">🗝️</div>
          <h4 class="text-lg font-semibold text-gray-900 mb-2">
            Mapbox Token Required
          </h4>
          <p class="text-sm text-gray-600 mb-4">
            Add your Mapbox PUBLIC token to .env.local:
          </p>
          <div class="bg-gray-50 p-3 rounded text-sm font-mono mb-3">
            VUE_APP_MAPBOX_TOKEN=pk.your_token_here
          </div>
          <div class="text-xs text-orange-600 bg-orange-50 p-2 rounded">
            ⚠️ Use PUBLIC token (pk.*), NOT secret token (sk.*)
          </div>
        </div>
      </div>

      <!-- Error State -->
      <div
        v-if="mapError"
        class="absolute inset-0 flex items-center justify-center bg-red-50"
      >
        <div class="text-center p-6">
          <div class="text-4xl mb-4">⚠️</div>
          <h4 class="text-lg font-semibold text-red-900 mb-2">Map Error</h4>
          <p class="text-sm text-red-700 mb-4">{{ mapError }}</p>
          <button
            @click="initializeMapbox"
            class="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
          >
            Retry
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  ref,
  computed,
  watch,
  onMounted,
  onUnmounted,
  defineExpose,
} from "vue";
import { useFacilities } from "@/composables/useFacilities";
import mapboxgl from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";

export default {
  name: "FacilityMapbox",
  props: {
    facilities: {
      type: Object,
      default: () => ({}),
    },
    userLocation: {
      type: Object,
      default: null,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    height: {
      type: String,
      default: "500px",
    },
  },
  emits: ["facility-select"],
  setup(props, { emit }) {
    const { calculateDistance } = useFacilities();

    // State
    const mapContainer = ref(null);
    const map = ref(null);
    const mapError = ref("");
    const showingDirections = ref(false);
    const currentDirections = ref(null);
    const markers = ref([]);

    // Get Mapbox token
    const mapboxToken = process.env.VUE_APP_MAPBOX_TOKEN || "";

    // Computed
    const mapHeight = computed(() => props.height);

    const allFacilities = computed(() => {
      const facilities = [];

      if (props.facilities.abaCenters) {
        props.facilities.abaCenters.forEach((center) => {
          if (center.latitude && center.longitude) {
            facilities.push({ type: "aba_center", data: center });
          }
        });
      }

      if (props.facilities.resourceCenters) {
        props.facilities.resourceCenters.forEach((center) => {
          if (center.latitude && center.longitude) {
            facilities.push({ type: "resource_center", data: center });
          }
        });
      }

      if (props.facilities.regionalCenters) {
        props.facilities.regionalCenters.forEach((center) => {
          if (center.latitude && center.longitude) {
            facilities.push({ type: "regional_center", data: center });
          }
        });
      }

      if (props.facilities.resources) {
        props.facilities.resources.forEach((resource) => {
          if (resource.latitude && resource.longitude) {
            facilities.push({ type: "resource", data: resource });
          }
        });
      }

      if (props.facilities.providers) {
        props.facilities.providers.forEach((provider) => {
          if (provider.latitude && provider.longitude) {
            facilities.push({ type: "provider", data: provider });
          }
        });
      }

      return facilities;
    });

    const facilityColors = {
      aba_center: "#3B82F6",
      resource_center: "#10B981",
      regional_center: "#8B5CF6",
      resource: "#F59E0B",
      provider: "#F97316",
    };

    // Utility function to get facility name with fallbacks
    const getFacilityName = (facilityData, facilityType = "") => {
      return (
        facilityData.name ||
        facilityData.regional_center || // For regional centers
        facilityData.center_name ||
        facilityData.facility_name ||
        facilityData.organization_name ||
        facilityData.provider_name ||
        facilityData.title ||
        facilityData.display_name ||
        (facilityType
          ? `${getFacilityTypeLabel(facilityType)} - Location`
          : "Unknown Facility")
      );
    };

    // Methods
    const initializeMapbox = () => {
      if (!mapboxToken) {
        console.warn("No Mapbox token provided");
        return;
      }

      // Check token type
      if (mapboxToken.startsWith("sk.")) {
        mapError.value =
          "❌ You're using a SECRET token (sk.*). Please use a PUBLIC token (pk.*) for frontend applications.";
        return;
      }

      if (!mapboxToken.startsWith("pk.")) {
        mapError.value =
          "❌ Invalid Mapbox token format. Token should start with 'pk.' for public access.";
        return;
      }

      if (!mapContainer.value) return;

      try {
        mapboxgl.accessToken = mapboxToken;

        map.value = new mapboxgl.Map({
          container: mapContainer.value,
          style: "mapbox://styles/mapbox/streets-v12",
          center: [-118.2437, 34.0522],
          zoom: 10,
          // Disable external navigation behaviors
          hash: false,
          pitchWithRotate: false,
          attributionControl: false, // Remove Mapbox attribution links
        });

        map.value.addControl(new mapboxgl.NavigationControl());
        map.value.addControl(new mapboxgl.FullscreenControl());

        map.value.on("load", () => {
          console.log("✅ Mapbox loaded!");
          addFacilityMarkers();
          if (props.userLocation) {
            addUserLocationMarker();
          }
        });

        map.value.on("error", (e) => {
          mapError.value = `Mapbox error: ${
            e.error?.message || "Unknown error"
          }`;
        });
      } catch (error) {
        mapError.value = `Failed to initialize: ${error.message}`;
      }
    };

    const addFacilityMarkers = () => {
      if (!map.value) return;

      clearMarkers();

      // Clear old facility data
      if (window.mapboxFacilities) delete window.mapboxFacilities;
      if (window.mapboxFacilityData) delete window.mapboxFacilityData;

      allFacilities.value.forEach((facility, index) => {
        const color = facilityColors[facility.type] || "#6B7280";

        const marker = new mapboxgl.Marker({ color })
          .setLngLat([facility.data.longitude, facility.data.latitude])
          .addTo(map.value);

        const facilityName = getFacilityName(facility.data, facility.type);

        // Store facility data using index to avoid JSON parsing issues
        const facilityId = `facility_${index}_${Date.now()}`;
        window.mapboxFacilities = window.mapboxFacilities || {};
        window.mapboxFacilities[facilityId] = facility;
        window.mapboxFacilityData = window.mapboxFacilityData || {};
        window.mapboxFacilityData[facilityId] = facility.data;

        // Create safe, non-linking popup content
        const address =
          facility.data.address ||
          [facility.data.city, facility.data.state]
            .filter(Boolean)
            .join(", ") ||
          "Address not available";
        const phone = facility.data.phone || facility.data.telephone || "";

        const popup = new mapboxgl.Popup({
          offset: 25,
          closeButton: true,
          closeOnClick: false, // Prevent accidental closing
          maxWidth: "280px",
        }).setHTML(`
            <div class="p-3" style="font-family: inherit;">
              <h4 class="font-semibold text-gray-900 mb-1">${facilityName}</h4>
              <p class="text-xs text-gray-600 mb-2">${getFacilityTypeLabel(
                facility.type
              )}</p>
              
              ${
                address
                  ? `<div class="text-xs text-gray-700 mb-2">
                📍 <span style="user-select: none; pointer-events: none;">${address}</span>
              </div>`
                  : ""
              }
              
              ${
                phone
                  ? `<div class="text-xs text-gray-700 mb-2">
                📞 <span style="user-select: none; pointer-events: none;">${phone}</span>
              </div>`
                  : ""
              }
              
              <div class="flex space-x-2 mt-3">
                <button onclick="window.selectFacility('${facilityId}'); event.preventDefault(); event.stopPropagation(); return false;" 
                        class="px-3 py-1 bg-blue-500 text-white text-xs rounded hover:bg-blue-600 transition-colors"
                        style="text-decoration: none; color: white;">
                  📋 Details
                </button>
                <button onclick="window.getDirections('${facilityId}'); event.preventDefault(); event.stopPropagation(); return false;"
                        class="px-3 py-1 bg-green-500 text-white text-xs rounded hover:bg-green-600 transition-colors"
                        style="text-decoration: none; color: white;">
                  🧭 Directions
                </button>
              </div>
              
              <div class="text-xs text-gray-500 mt-2 text-center">
                Click buttons above - stay in app
              </div>
            </div>
          `);

        marker.setPopup(popup);
        markers.value.push(marker);
      });

      // Global functions for popup buttons - secured to prevent external navigation
      window.selectFacility = (facilityId) => {
        try {
          // Prevent any default browser behavior
          if (window.event) {
            window.event.preventDefault();
            window.event.stopPropagation();
          }

          const facility = window.mapboxFacilities?.[facilityId];
          if (facility) {
            console.log(
              "🏢 Opening facility details within app:",
              getFacilityName(facility.data)
            );
            emit("facility-select", facility.data, facility.type);
          }
          return false; // Prevent any navigation
        } catch (error) {
          console.error("Error opening facility details:", error);
          return false;
        }
      };

      window.getDirections = (facilityId) => {
        try {
          // Prevent any default browser behavior
          if (window.event) {
            window.event.preventDefault();
            window.event.stopPropagation();
          }

          const facilityData = window.mapboxFacilityData?.[facilityId];
          if (facilityData) {
            console.log(
              "🧭 Getting in-app directions to:",
              getFacilityName(facilityData)
            );
            getDirections(facilityData);
          }
          return false; // Prevent any navigation
        } catch (error) {
          console.error("Error getting directions:", error);
          return false;
        }
      };
    };

    const addUserLocationMarker = () => {
      if (!map.value || !props.userLocation) return;

      const marker = new mapboxgl.Marker({ color: "#4285F4" })
        .setLngLat([props.userLocation.lng, props.userLocation.lat])
        .addTo(map.value);

      const popup = new mapboxgl.Popup({
        offset: 25,
        closeButton: true,
        closeOnClick: false,
      }).setHTML(`
        <div class="p-3 text-center" style="font-family: inherit;">
          <div class="text-blue-600 font-semibold mb-1">📍 Your Current Location</div>
          <div class="text-xs text-gray-600">All directions start from here</div>
        </div>
      `);

      marker.setPopup(popup);
      markers.value.push(marker);
    };

    const clearMarkers = () => {
      markers.value.forEach((marker) => marker.remove());
      markers.value = [];
    };

    const getFacilityTypeLabel = (type) => {
      const labels = {
        aba_center: "ABA Center",
        resource_center: "Resource Center",
        regional_center: "Regional Center",
        resource: "Resource",
        provider: "Provider",
      };
      return labels[type] || "Facility";
    };

    const getDirections = async (facility) => {
      // Always ensure we have current user location
      if (!props.userLocation || !facility.latitude || !facility.longitude) {
        alert(
          "📍 Unable to get directions: Your current location is required. Please enable location services."
        );
        return;
      }

      // Clear any existing directions first
      clearDirections();

      console.log(
        "🧭 Getting directions from your current location to:",
        getFacilityName(facility)
      );

      try {
        // Always use current user location as starting point
        const startLng = props.userLocation.lng;
        const startLat = props.userLocation.lat;
        const endLng = facility.longitude;
        const endLat = facility.latitude;

        // Use Mapbox Directions API with explicit current location
        const response = await fetch(
          `https://api.mapbox.com/directions/v5/mapbox/driving/${startLng},${startLat};${endLng},${endLat}?steps=true&geometries=geojson&access_token=${mapboxToken}`
        );

        if (response.ok) {
          const data = await response.json();
          if (data.routes && data.routes.length > 0) {
            const route = data.routes[0];

            addRouteToMap(route.geometry);

            currentDirections.value = {
              destination_name: getFacilityName(facility),
              distance: `${(route.distance * 0.000621371).toFixed(1)} miles`,
              duration: formatDuration(route.duration),
              from_current_location: true,
              steps: route.legs[0].steps.map((step, index) => ({
                instructions: step.maneuver.instruction || `Step ${index + 1}`,
                distance: `${(step.distance * 0.000621371).toFixed(1)} mi`,
                duration: formatDuration(step.duration),
              })),
            };

            showingDirections.value = true;
            return;
          }
        }

        // Fallback to simple route
        createSimpleRoute(facility);
      } catch (error) {
        console.error("Directions failed:", error);
        createSimpleRoute(facility);
      }
    };

    const addRouteToMap = (geometry) => {
      if (!map.value) return;

      if (map.value.getLayer("route")) {
        map.value.removeLayer("route");
      }
      if (map.value.getSource("route")) {
        map.value.removeSource("route");
      }

      map.value.addSource("route", {
        type: "geojson",
        data: {
          type: "Feature",
          properties: {},
          geometry: geometry,
        },
      });

      map.value.addLayer({
        id: "route",
        type: "line",
        source: "route",
        layout: {
          "line-join": "round",
          "line-cap": "round",
        },
        paint: {
          "line-color": "#4285F4",
          "line-width": 5,
          "line-opacity": 0.8,
        },
      });
    };

    const createSimpleRoute = (facility) => {
      // Always use current user location as starting point
      const startLng = props.userLocation.lng;
      const startLat = props.userLocation.lat;
      const endLng = facility.longitude;
      const endLat = facility.latitude;

      const distance = calculateDistance(startLat, startLng, endLat, endLng);
      const estimatedTime = Math.round((distance / 35) * 60);

      currentDirections.value = {
        destination_name: getFacilityName(facility),
        distance: `${distance.toFixed(1)} miles`,
        duration: formatDuration(estimatedTime * 60),
        from_current_location: true,
        steps: [
          {
            instructions: `From your current location, head toward <strong>${getFacilityName(
              facility
            )}</strong>`,
            distance: `${distance.toFixed(1)} mi`,
            duration: formatDuration(estimatedTime * 60),
          },
        ],
      };

      // Create line from current location to destination
      const lineGeometry = {
        type: "LineString",
        coordinates: [
          [startLng, startLat], // Always start from current location
          [endLng, endLat], // End at facility
        ],
      };

      addRouteToMap(lineGeometry);
      showingDirections.value = true;
    };

    const formatDuration = (seconds) => {
      const minutes = Math.round(seconds / 60);
      if (minutes < 60) {
        return `${minutes} min`;
      }
      const hours = Math.floor(minutes / 60);
      const remainingMinutes = minutes % 60;
      return `${hours}h ${remainingMinutes}m`;
    };

    const clearDirections = () => {
      if (map.value) {
        if (map.value.getLayer("route")) {
          map.value.removeLayer("route");
        }
        if (map.value.getSource("route")) {
          map.value.removeSource("route");
        }
      }
      showingDirections.value = false;
      currentDirections.value = null;
    };

    const fitMapToFacilities = () => {
      if (!map.value || allFacilities.value.length === 0) return;

      const coordinates = allFacilities.value.map((facility) => [
        facility.data.longitude,
        facility.data.latitude,
      ]);

      if (props.userLocation) {
        coordinates.push([props.userLocation.lng, props.userLocation.lat]);
      }

      if (coordinates.length > 0) {
        const bounds = coordinates.reduce((bounds, coord) => {
          return bounds.extend(coord);
        }, new mapboxgl.LngLatBounds(coordinates[0], coordinates[0]));

        map.value.fitBounds(bounds, { padding: 50 });
      }
    };

    const centerOnUser = () => {
      if (!map.value || !props.userLocation) return;

      map.value.flyTo({
        center: [props.userLocation.lng, props.userLocation.lat],
        zoom: 14,
      });
    };

    // Watchers
    watch(
      () => allFacilities.value,
      () => {
        if (map.value && map.value.loaded()) {
          addFacilityMarkers();
        }
      }
    );

    watch(
      () => props.userLocation,
      (newLocation, oldLocation) => {
        if (map.value && map.value.loaded() && newLocation) {
          addUserLocationMarker();

          // If user location changed and we have active directions, clear them
          // so they can be recalculated from the new current location
          if (
            oldLocation &&
            showingDirections.value &&
            (newLocation.lat !== oldLocation.lat ||
              newLocation.lng !== oldLocation.lng)
          ) {
            console.log(
              "📍 User location changed - clearing directions for recalculation"
            );
            clearDirections();
          }
        }
      }
    );

    // Lifecycle
    onMounted(() => {
      // Disable auto-detection of phone numbers and addresses globally for this component
      const metaElements = document.querySelectorAll(
        'meta[name="format-detection"]'
      );
      metaElements.forEach((el) => el.remove());

      const metaTag = document.createElement("meta");
      metaTag.name = "format-detection";
      metaTag.content = "telephone=no, address=no";
      document.head.appendChild(metaTag);

      console.log(
        "🔒 Mapbox security: Disabled auto-linking for phone/address detection"
      );
      initializeMapbox();
    });

    onUnmounted(() => {
      if (map.value) {
        map.value.remove();
      }
      // Clean up global functions and data
      if (window.selectFacility) delete window.selectFacility;
      if (window.getDirections) delete window.getDirections;
      if (window.mapboxFacilities) delete window.mapboxFacilities;
      if (window.mapboxFacilityData) delete window.mapboxFacilityData;
    });

    // Expose methods for parent component access
    defineExpose({
      getDirections,
      clearDirections,
      fitMapToFacilities,
      centerOnUser,
    });

    return {
      mapContainer,
      map,
      mapError,
      showingDirections,
      currentDirections,
      mapboxToken,
      mapHeight,
      allFacilities,
      initializeMapbox,
      getDirections,
      clearDirections,
      fitMapToFacilities,
      centerOnUser,
    };
  },
};
</script>

<style scoped>
.facility-map {
  min-height: 400px;
}

:deep(.mapboxgl-popup-content) {
  border-radius: 8px;
  /* Prevent any auto-linking or external navigation */
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
}

/* Disable auto-detection of phone numbers and addresses in popups */
:deep(.mapboxgl-popup-content a) {
  pointer-events: none !important;
  text-decoration: none !important;
  color: inherit !important;
}

/* Ensure popup buttons don't navigate externally */
:deep(.mapboxgl-popup-content button) {
  cursor: pointer !important;
  text-decoration: none !important;
  border: none !important;
  background: transparent !important;
}

/* Prevent text selection in address/phone fields */
:deep(.mapboxgl-popup-content span[style*="pointer-events: none"]) {
  -webkit-touch-callout: none !important;
  -webkit-user-select: none !important;
  -khtml-user-select: none !important;
  -moz-user-select: none !important;
  -ms-user-select: none !important;
  user-select: none !important;
}
</style>
