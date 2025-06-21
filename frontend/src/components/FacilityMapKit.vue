<template>
  <div class="facility-map bg-white rounded-lg shadow-md overflow-hidden">
    <!-- Map Header -->
    <div class="p-4 border-b border-gray-200 bg-gray-50">
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-semibold text-gray-800">Facility Map</h3>
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

    <!-- Directions Panel (when active) -->
    <div
      v-if="showingDirections && currentDirections"
      class="p-4 border-b border-gray-200 bg-blue-50 max-h-40 overflow-y-auto"
    >
      <div class="flex items-center justify-between mb-2">
        <h4 class="font-semibold text-blue-900">
          🧭 Directions to {{ currentDirections.destination_name }}
        </h4>
        <button
          @click="clearDirections"
          class="text-blue-600 hover:text-blue-800 p-1"
        >
          ✕
        </button>
      </div>
      <div class="text-sm text-blue-800 mb-2">
        📍 {{ currentDirections.distance }} • ⏱️
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
            class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mb-2"
          ></div>
          <p class="text-gray-600">Loading map...</p>
        </div>
      </div>

      <!-- MapKit Map -->
      <div
        ref="mapContainer"
        class="w-full h-full"
        :class="{ 'opacity-0': loading }"
      ></div>

      <!-- No MapKit Warning -->
      <div
        v-if="!mapKitAvailable"
        class="absolute inset-0 flex items-center justify-center bg-gray-100"
      >
        <div class="text-center p-6">
          <div class="text-4xl mb-4">🗺️</div>
          <h4 class="text-lg font-semibold text-gray-900 mb-2">
            MapKit Not Available
          </h4>
          <p class="text-sm text-gray-600">
            MapKit JS failed to load. Please check your internet connection.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, watch, onMounted, onUnmounted, nextTick } from "vue";
import { useFacilities } from "@/composables/useFacilities";

export default {
  name: "FacilityMapKit",
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

    // Map state
    const mapContainer = ref(null);
    const map = ref(null);
    const mapKitAvailable = ref(false);
    const selectedFacility = ref(null);
    const showingDirections = ref(false);
    const currentDirections = ref(null);
    const annotations = ref([]);
    const routeOverlay = ref(null);

    // Computed
    const mapHeight = computed(() => props.height);

    const allFacilities = computed(() => {
      const facilities = [];

      // Add ABA Centers
      if (props.facilities.abaCenters) {
        props.facilities.abaCenters.forEach((center) => {
          if (center.latitude && center.longitude) {
            facilities.push({ type: "aba_center", data: center });
          }
        });
      }

      // Add Resource Centers
      if (props.facilities.resourceCenters) {
        props.facilities.resourceCenters.forEach((center) => {
          if (center.latitude && center.longitude) {
            facilities.push({ type: "resource_center", data: center });
          }
        });
      }

      // Add Regional Centers
      if (props.facilities.regionalCenters) {
        props.facilities.regionalCenters.forEach((center) => {
          if (center.latitude && center.longitude) {
            facilities.push({ type: "regional_center", data: center });
          }
        });
      }

      // Add Resources
      if (props.facilities.resources) {
        props.facilities.resources.forEach((resource) => {
          if (resource.latitude && resource.longitude) {
            facilities.push({ type: "resource", data: resource });
          }
        });
      }

      // Add Providers
      if (props.facilities.providers) {
        props.facilities.providers.forEach((provider) => {
          if (provider.latitude && provider.longitude) {
            facilities.push({ type: "provider", data: provider });
          }
        });
      }

      return facilities;
    });

    // Colors for facility types
    const facilityColors = {
      aba_center: "#3B82F6", // Blue
      resource_center: "#10B981", // Green
      regional_center: "#8B5CF6", // Purple
      resource: "#F59E0B", // Yellow/Orange
      provider: "#F97316", // Orange
    };

    // Methods
    const initializeMapKit = async () => {
      try {
        if (!window.mapkit) {
          console.error("MapKit JS not loaded");
          return;
        }

        // Initialize MapKit (no auth required for basic usage)
        window.mapkit.init({
          authorizationCallback: function (done) {
            // For production, you'd get a proper JWT token from your server
            // For development, we can try without authentication
            done("");
          },
        });

        mapKitAvailable.value = true;
        await nextTick();
        createMap();
      } catch (error) {
        console.error("MapKit initialization failed:", error);
        mapKitAvailable.value = false;
      }
    };

    const createMap = () => {
      if (!mapContainer.value || !window.mapkit) return;

      // Create the map
      map.value = new window.mapkit.Map(mapContainer.value, {
        center: new window.mapkit.Coordinate(34.0522, -118.2437), // Default to LA
        span: new window.mapkit.CoordinateSpan(0.5, 0.5),
        mapType: window.mapkit.Map.MapTypes.Standard,
        showsMapTypeControl: true,
        showsZoomControl: true,
        showsUserLocationControl: true,
        isRotationEnabled: false,
      });

      console.log("✅ MapKit map created successfully");

      // Add facilities when map is ready
      addFacilityAnnotations();

      // Add user location if available
      if (props.userLocation) {
        addUserLocationAnnotation();
      }
    };

    const addFacilityAnnotations = () => {
      if (!map.value || allFacilities.value.length === 0) return;

      // Clear existing annotations
      clearAnnotations();

      allFacilities.value.forEach((facility) => {
        const annotation = new window.mapkit.MarkerAnnotation(
          new window.mapkit.Coordinate(
            facility.data.latitude,
            facility.data.longitude
          ),
          {
            color: facilityColors[facility.type] || "#6B7280",
            title:
              facility.data.name || facility.data.center_name || "Facility",
            subtitle: getFacilityTypeLabel(facility.type),
            data: { facility: facility },
          }
        );

        // Add click handler
        annotation.addEventListener("select", () => {
          selectFacility(facility);
        });

        annotations.value.push(annotation);
      });

      // Add all annotations to map
      map.value.addAnnotations(annotations.value);
    };

    const addUserLocationAnnotation = () => {
      if (!map.value || !props.userLocation) return;

      const userAnnotation = new window.mapkit.MarkerAnnotation(
        new window.mapkit.Coordinate(
          props.userLocation.lat,
          props.userLocation.lng
        ),
        {
          color: "#4285F4",
          title: "Your Location",
          subtitle: "Current position",
        }
      );

      map.value.addAnnotation(userAnnotation);
    };

    const clearAnnotations = () => {
      if (map.value && annotations.value.length > 0) {
        map.value.removeAnnotations(annotations.value);
        annotations.value = [];
      }
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

    const selectFacility = (facility) => {
      selectedFacility.value = facility;

      // Show facility details popup (you can customize this)
      const facilityName = facility.data.name || facility.data.center_name;
      const distance = props.userLocation
        ? calculateDistance(
            props.userLocation.lat,
            props.userLocation.lng,
            facility.data.latitude,
            facility.data.longitude
          ).toFixed(1)
        : null;

      // You could show a custom popup here or emit an event
      console.log(
        `Selected facility: ${facilityName}${
          distance ? ` (${distance} miles away)` : ""
        }`
      );

      // Emit facility selection
      emit("facility-select", facility.data, facility.type);
    };

    const getDirections = (facility) => {
      if (!props.userLocation || !facility.latitude || !facility.longitude) {
        alert(
          "Unable to get directions: User location or facility location not available"
        );
        return;
      }

      // For now, use a simple route (MapKit doesn't have built-in directions API)
      createSimpleRoute(facility);
    };

    const createSimpleRoute = (facility) => {
      if (!map.value) return;

      const start = new window.mapkit.Coordinate(
        props.userLocation.lat,
        props.userLocation.lng
      );
      const end = new window.mapkit.Coordinate(
        facility.latitude,
        facility.longitude
      );

      // Create a simple polyline route
      const coordinates = [start, end];

      if (routeOverlay.value) {
        map.value.removeOverlay(routeOverlay.value);
      }

      routeOverlay.value = new window.mapkit.PolylineOverlay(coordinates, {
        style: new window.mapkit.Style({
          strokeColor: "#4285F4",
          strokeWidth: 4,
          strokeOpacity: 0.8,
        }),
      });

      map.value.addOverlay(routeOverlay.value);

      // Create simple directions
      const distance = calculateDistance(
        props.userLocation.lat,
        props.userLocation.lng,
        facility.latitude,
        facility.longitude
      );

      const estimatedTime = Math.round((distance / 35) * 60); // 35 mph average

      currentDirections.value = {
        destination_name:
          facility.name || facility.center_name || "Selected Facility",
        distance: `${distance.toFixed(1)} miles`,
        duration:
          estimatedTime < 60
            ? `${estimatedTime} min`
            : `${Math.floor(estimatedTime / 60)}h ${estimatedTime % 60}m`,
        steps: [
          {
            instructions: `Head toward <strong>${
              facility.name || facility.center_name
            }</strong>`,
            distance: `${distance.toFixed(1)} mi`,
            duration:
              estimatedTime < 60
                ? `${estimatedTime} min`
                : `${Math.floor(estimatedTime / 60)}h ${estimatedTime % 60}m`,
          },
          {
            instructions: `Arrive at <strong>${
              facility.name || facility.center_name
            }</strong>`,
            distance: "0.0 mi",
            duration: "0 min",
          },
        ],
      };

      showingDirections.value = true;
    };

    const clearDirections = () => {
      if (map.value && routeOverlay.value) {
        map.value.removeOverlay(routeOverlay.value);
        routeOverlay.value = null;
      }
      showingDirections.value = false;
      currentDirections.value = null;
    };

    const fitMapToFacilities = () => {
      if (!map.value || allFacilities.value.length === 0) return;

      const coordinates = allFacilities.value.map(
        (facility) =>
          new window.mapkit.Coordinate(
            facility.data.latitude,
            facility.data.longitude
          )
      );

      // Add user location if available
      if (props.userLocation) {
        coordinates.push(
          new window.mapkit.Coordinate(
            props.userLocation.lat,
            props.userLocation.lng
          )
        );
      }

      if (coordinates.length > 0) {
        // Calculate bounding box manually
        let minLat = coordinates[0].latitude;
        let maxLat = coordinates[0].latitude;
        let minLng = coordinates[0].longitude;
        let maxLng = coordinates[0].longitude;

        coordinates.forEach((coord) => {
          minLat = Math.min(minLat, coord.latitude);
          maxLat = Math.max(maxLat, coord.latitude);
          minLng = Math.min(minLng, coord.longitude);
          maxLng = Math.max(maxLng, coord.longitude);
        });

        // Add padding
        const latPadding = (maxLat - minLat) * 0.1;
        const lngPadding = (maxLng - minLng) * 0.1;

        const center = new window.mapkit.Coordinate(
          (minLat + maxLat) / 2,
          (minLng + maxLng) / 2
        );

        const span = new window.mapkit.CoordinateSpan(
          Math.max(maxLat - minLat + latPadding, 0.01),
          Math.max(maxLng - minLng + lngPadding, 0.01)
        );

        const region = new window.mapkit.CoordinateRegion(center, span);
        map.value.setRegionAnimated(region, true);
      }
    };

    const centerOnUser = () => {
      if (!map.value || !props.userLocation) return;

      const userCoordinate = new window.mapkit.Coordinate(
        props.userLocation.lat,
        props.userLocation.lng
      );
      const region = new window.mapkit.CoordinateRegion(
        userCoordinate,
        new window.mapkit.CoordinateSpan(0.1, 0.1)
      );

      map.value.setRegionAnimated(region, true);
    };

    // Watch for changes
    watch(
      () => allFacilities.value,
      () => {
        if (map.value) {
          addFacilityAnnotations();
        }
      }
    );

    watch(
      () => props.userLocation,
      (newLocation) => {
        if (map.value && newLocation) {
          addUserLocationAnnotation();
        }
      }
    );

    // Lifecycle
    onMounted(() => {
      initializeMapKit();
    });

    onUnmounted(() => {
      if (map.value) {
        map.value.destroy();
      }
    });

    return {
      // Refs
      mapContainer,
      map,
      mapKitAvailable,
      selectedFacility,
      showingDirections,
      currentDirections,

      // Computed
      mapHeight,
      allFacilities,

      // Methods
      selectFacility,
      getDirections,
      clearDirections,
      fitMapToFacilities,
      centerOnUser,
    };
  },
};
</script>

<style scoped>
/* Custom styles for map container */
.facility-map {
  min-height: 400px;
}
</style>
