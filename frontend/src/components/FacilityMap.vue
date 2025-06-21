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

      <!-- Google Map -->
      <GoogleMap
        v-if="!loading && apiKey"
        ref="map"
        :api-key="apiKey"
        :center="mapCenter"
        :zoom="mapZoom"
        :options="mapOptions"
        class="w-full h-full"
        @ready="onMapReady"
        @click="onMapClick"
      >
        <!-- User Location Marker -->
        <Marker
          v-if="userLocation && showUserLocation"
          :options="{
            position: userLocation,
            icon: userLocationIcon,
            title: 'Your Location',
            zIndex: 1000,
          }"
        />

        <!-- Facility Markers -->
        <template
          v-for="facility in allFacilities"
          :key="`${facility.type}_${facility.data.id}`"
        >
          <Marker
            :options="{
              position: {
                lat: facility.data.latitude,
                lng: facility.data.longitude,
              },
              icon: getFacilityIcon(facility.type),
              title: facility.data.name || facility.data.center_name,
              zIndex: 100,
            }"
            @click="selectFacility(facility)"
          />
        </template>

        <!-- Info Window for Selected Facility -->
        <InfoWindow
          v-if="selectedFacility"
          :options="{
            position: {
              lat: selectedFacility.data.latitude,
              lng: selectedFacility.data.longitude,
            },
          }"
          @closeclick="clearSelection"
        >
          <div class="p-3 max-w-xs">
            <div class="flex items-center space-x-2 mb-2">
              <span class="text-lg">{{
                getFacilityTypeIcon(selectedFacility.type)
              }}</span>
              <h4 class="font-semibold text-gray-900">
                {{
                  selectedFacility.data.name ||
                  selectedFacility.data.center_name
                }}
              </h4>
            </div>

            <div class="text-sm text-gray-600 space-y-1">
              <div
                v-if="
                  selectedFacility.data.address ||
                  selectedFacility.data.full_address
                "
              >
                📍
                {{
                  selectedFacility.data.address ||
                  selectedFacility.data.full_address
                }}
              </div>

              <div
                v-if="
                  selectedFacility.data.phone ||
                  selectedFacility.data.contact_info?.phone
                "
              >
                📞
                {{
                  selectedFacility.data.phone ||
                  selectedFacility.data.contact_info?.phone
                }}
              </div>

              <div
                v-if="
                  userLocation &&
                  selectedFacility.data.latitude &&
                  selectedFacility.data.longitude
                "
              >
                🗺️
                {{
                  calculateDistanceToFacility(selectedFacility.data).toFixed(1)
                }}
                miles away
              </div>
            </div>

            <div class="flex space-x-2 mt-3">
              <button
                @click="viewFacilityDetails(selectedFacility)"
                class="px-3 py-1 bg-blue-600 text-white text-xs rounded hover:bg-blue-700 transition-colors"
              >
                View Details
              </button>
              <button
                @click="getDirections(selectedFacility.data)"
                class="px-3 py-1 bg-green-600 text-white text-xs rounded hover:bg-green-700 transition-colors"
              >
                Directions
              </button>
            </div>
          </div>
        </InfoWindow>
      </GoogleMap>

      <!-- No API Key Warning -->
      <div
        v-else-if="!apiKey"
        class="absolute inset-0 flex items-center justify-center bg-gray-100"
      >
        <div class="text-center p-6">
          <div class="text-4xl mb-4">🗺️</div>
          <h4 class="text-lg font-semibold text-gray-900 mb-2">
            Map Not Available
          </h4>
          <p class="text-sm text-gray-600">
            Google Maps API key is required to display the map.
          </p>
        </div>
      </div>

      <!-- Directions API Notice (when needed) -->
      <div
        v-if="apiKey && googleMapInstance && !directionsService"
        class="absolute bottom-4 right-4 bg-blue-50 border border-blue-200 rounded-lg p-3 max-w-sm shadow-lg"
      >
        <div class="flex items-start space-x-2">
          <span class="text-blue-600 text-lg">ℹ️</span>
          <div>
            <h5 class="text-sm font-semibold text-blue-800">
              Enable Google Routes
            </h5>
            <p class="text-xs text-blue-700 mt-1">
              Enable "Directions API" in Google Cloud Console to get the real
              Google Maps blue route lines and turn-by-turn directions on this
              page.
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, watch, onMounted } from "vue";
import { GoogleMap, Marker, InfoWindow } from "vue3-google-map";
import { useFacilities } from "@/composables/useFacilities";

export default {
  name: "FacilityMap",
  components: {
    GoogleMap,
    Marker,
    InfoWindow,
  },
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
    const map = ref(null);
    const googleMapInstance = ref(null);
    const selectedFacility = ref(null);
    const showUserLocation = ref(true);
    const mapCenter = ref({ lat: 34.0522, lng: -118.2437 }); // Default to LA
    const mapZoom = ref(10);

    // Directions state
    const showingDirections = ref(false);
    const currentDirections = ref(null);
    const directionsService = ref(null);
    const customRouteLine = ref(null);

    // API Key - you'll need to set this in your environment
    const apiKey = process.env.VUE_APP_GOOGLE_MAPS_API_KEY || "";

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

    const mapOptions = computed(() => ({
      disableDefaultUI: false,
      zoomControl: true,
      mapTypeControl: true,
      scaleControl: true,
      streetViewControl: true,
      rotateControl: false,
      fullscreenControl: true,
      gestureHandling: "auto",
    }));

    const userLocationIcon = computed(() => ({
      path: window.google?.maps?.SymbolPath?.CIRCLE || 0,
      scale: 10,
      fillColor: "#4285F4",
      fillOpacity: 1,
      strokeColor: "#FFFFFF",
      strokeWeight: 3,
    }));

    // Watch for user location changes
    watch(
      () => props.userLocation,
      (newLocation) => {
        if (newLocation) {
          mapCenter.value = newLocation;
          mapZoom.value = 12;
        }
      },
      { immediate: true }
    );

    // Watch for facilities changes
    watch(
      () => allFacilities.value,
      (newFacilities) => {
        if (newFacilities.length > 0 && !props.userLocation) {
          // If no user location, center on facilities
          fitMapToFacilities();
        }
      }
    );

    // Methods
    const onMapReady = (mapInstance) => {
      map.value = mapInstance;
      googleMapInstance.value = mapInstance;

      // Initialize directions service
      if (window.google?.maps) {
        try {
          directionsService.value = new window.google.maps.DirectionsService();
          console.log("✅ Directions service initialized successfully");
        } catch (error) {
          console.warn(
            "⚠️ Directions API not available. Enable 'Directions API' in Google Cloud Console for real Google routes."
          );
          console.warn("Fallback: Local routing will be used.");
        }
      } else {
        console.warn("⚠️ Google Maps API not loaded");
      }

      console.log("Map ready:", mapInstance);
      console.log("Google Maps instance:", googleMapInstance.value);
    };

    const onMapClick = () => {
      clearSelection();
      // Optionally clear directions when clicking on map
      // clearDirections();
    };

    const selectFacility = (facility) => {
      selectedFacility.value = facility;
    };

    const clearSelection = () => {
      selectedFacility.value = null;
    };

    const viewFacilityDetails = (facility) => {
      emit("facility-select", facility.data, facility.type);
      clearSelection();
    };

    const getDirections = (facility) => {
      if (!props.userLocation || !facility.latitude || !facility.longitude) {
        alert(
          "Unable to get directions: User location or facility location not available"
        );
        return;
      }

      // Show loading state
      currentDirections.value = {
        destination_name:
          facility.name || facility.center_name || "Selected Facility",
        distance: "Calculating...",
        duration: "Calculating...",
        steps: [
          {
            instructions: "🔄 Getting directions...",
            distance: "",
            duration: "",
          },
        ],
      };
      showingDirections.value = true;
      clearSelection();

      // Always try Google Directions API first (if available)
      if (directionsService.value && window.google?.maps?.DirectionsService) {
        const request = {
          origin: props.userLocation,
          destination: { lat: facility.latitude, lng: facility.longitude },
          travelMode: window.google.maps.TravelMode.DRIVING,
        };

        directionsService.value.route(request, (result, status) => {
          if (status === "OK") {
            console.log("✅ Google Directions API success!");

            // Clear any existing custom routes
            clearCustomRoute();

            // DON'T use directionsRenderer - we'll draw our own polyline
            // Extract the route polyline from Google's response
            const route = result.routes[0];
            const leg = route.legs[0];

            // Get the encoded polyline and decode it
            const encodedPath = route.overview_polyline.points;

            let coordinates;
            if (window.google?.maps?.geometry?.encoding) {
              // Use Google's polyline decoder if available
              const decodedPath =
                window.google.maps.geometry.encoding.decodePath(encodedPath);
              coordinates = decodedPath.map((point) => [
                point.lng(),
                point.lat(),
              ]);
            } else {
              // Fallback: extract coordinates from legs path
              coordinates = [];
              leg.steps.forEach((step) => {
                step.path.forEach((point) => {
                  coordinates.push([point.lng(), point.lat()]);
                });
              });
            }

            drawRoutePolyline(coordinates);

            // Parse directions for the panel
            currentDirections.value = {
              destination_name:
                facility.name || facility.center_name || "Selected Facility",
              distance: leg.distance.text,
              duration: leg.duration.text,
              steps: leg.steps.map((step) => ({
                instructions: step.instructions,
                distance: step.distance.text,
                duration: step.duration.text,
              })),
            };

            return;
          } else {
            console.warn(
              `Google Directions failed: ${status}. Using local routing.`
            );
            createCustomDirections(facility);
          }
        });
      } else {
        console.log(
          "Google Directions API not available. Using local routing."
        );
        createCustomDirections(facility);
      }
    };

    const createCustomDirections = (facility) => {
      // Show loading state briefly
      currentDirections.value = {
        destination_name:
          facility.name || facility.center_name || "Selected Facility",
        distance: "Calculating...",
        duration: "Calculating...",
        steps: [
          {
            instructions: "🔄 Generating route...",
            distance: "",
            duration: "",
          },
        ],
      };
      showingDirections.value = true;
      clearSelection();

      // Generate local route with intelligent waypoints
      setTimeout(() => {
        const route = generateIntelligentRoute(props.userLocation, facility);
        drawRoutePolyline(route.coordinates);

        currentDirections.value = {
          destination_name:
            facility.name || facility.center_name || "Selected Facility",
          distance: route.totalDistance,
          duration: route.totalDuration,
          steps: route.directions,
        };
      }, 300); // Brief delay for better UX
    };

    const generateIntelligentRoute = (startLocation, facility) => {
      const start = { lat: startLocation.lat, lng: startLocation.lng };
      const end = { lat: facility.latitude, lng: facility.longitude };
      const facilityName =
        facility.name || facility.center_name || "Selected Facility";

      // Calculate straight-line distance
      const totalDistance = calculateDistance(
        start.lat,
        start.lng,
        end.lat,
        end.lng
      );
      const estimatedTimeMinutes = Math.round((totalDistance / 35) * 60); // 35 mph average

      // Generate intelligent waypoints
      const waypoints = generateWaypoints(start, end, totalDistance);

      // Create turn-by-turn directions
      const directions = generateDirections(
        waypoints,
        facilityName,
        totalDistance,
        estimatedTimeMinutes
      );

      // Convert waypoints to polyline coordinates
      const coordinates = waypoints.map((point) => [point.lng, point.lat]);

      return {
        coordinates: coordinates,
        totalDistance: `${totalDistance.toFixed(1)} miles`,
        totalDuration: formatDuration(estimatedTimeMinutes * 60),
        directions: directions,
      };
    };

    const generateWaypoints = (start, end, distance) => {
      const waypoints = [start];

      // For longer routes, add intermediate waypoints to simulate road following
      if (distance > 5) {
        const segments = Math.min(Math.floor(distance / 3), 4); // Max 4 segments

        for (let i = 1; i < segments; i++) {
          const ratio = i / segments;

          // Add some variation to simulate following roads
          const latOffset = (Math.random() - 0.5) * 0.01 * distance * 0.1;
          const lngOffset = (Math.random() - 0.5) * 0.01 * distance * 0.1;

          const waypoint = {
            lat: start.lat + (end.lat - start.lat) * ratio + latOffset,
            lng: start.lng + (end.lng - start.lng) * ratio + lngOffset,
          };

          waypoints.push(waypoint);
        }
      }

      waypoints.push(end);
      return waypoints;
    };

    const generateDirections = (
      waypoints,
      facilityName,
      totalDistance,
      totalTime
    ) => {
      const directions = [];

      if (waypoints.length < 2) return directions;

      // First instruction - starting direction
      const initialBearing = calculateBearing(
        waypoints[0].lat,
        waypoints[0].lng,
        waypoints[1].lat,
        waypoints[1].lng
      );

      const startDirection = getCompassDirection(initialBearing);

      // Add starting instruction
      directions.push({
        instructions: `Head <strong>${startDirection}</strong> toward your destination`,
        distance: `${(totalDistance * 0.3).toFixed(1)} mi`,
        duration: formatDuration(totalTime * 0.3 * 60),
      });

      // Add intermediate instructions for longer routes
      if (waypoints.length > 2) {
        for (let i = 1; i < waypoints.length - 1; i++) {
          const bearing = calculateBearing(
            waypoints[i].lat,
            waypoints[i].lng,
            waypoints[i + 1].lat,
            waypoints[i + 1].lng
          );
          const direction = getCompassDirection(bearing);
          const segmentRatio = 0.4 / (waypoints.length - 2);

          directions.push({
            instructions: `Continue <strong>${direction}</strong> on local roads`,
            distance: `${(totalDistance * segmentRatio).toFixed(1)} mi`,
            duration: formatDuration(totalTime * segmentRatio * 60),
          });
        }
      }

      // Add final instruction
      directions.push({
        instructions: `Arrive at <strong>${facilityName}</strong> on your right`,
        distance: `${(totalDistance * 0.1).toFixed(1)} mi`,
        duration: "1 min",
      });

      return directions;
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

    const drawRoutePolyline = (coordinates) => {
      if (!googleMapInstance.value || !window.google?.maps || !coordinates)
        return;

      clearCustomRoute(); // Clear any existing route

      // Convert coordinates to Google Maps format
      const path = coordinates.map((coord) => ({
        lat: coord[1], // latitude
        lng: coord[0], // longitude
      }));

      customRouteLine.value = new window.google.maps.Polyline({
        path: path,
        geodesic: true,
        strokeColor: "#4285F4",
        strokeOpacity: 0.9,
        strokeWeight: 6,
        icons: [
          {
            icon: {
              path: window.google.maps.SymbolPath.FORWARD_OPEN_ARROW,
              scale: 4,
              strokeColor: "#FFFFFF",
              fillColor: "#4285F4",
              fillOpacity: 1,
            },
            offset: "100%",
          },
        ],
      });

      customRouteLine.value.setMap(googleMapInstance.value);
    };

    const calculateBearing = (lat1, lng1, lat2, lng2) => {
      const dLng = ((lng2 - lng1) * Math.PI) / 180;
      const lat1Rad = (lat1 * Math.PI) / 180;
      const lat2Rad = (lat2 * Math.PI) / 180;

      const y = Math.sin(dLng) * Math.cos(lat2Rad);
      const x =
        Math.cos(lat1Rad) * Math.sin(lat2Rad) -
        Math.sin(lat1Rad) * Math.cos(lat2Rad) * Math.cos(dLng);

      const bearing = (Math.atan2(y, x) * 180) / Math.PI;
      return (bearing + 360) % 360;
    };

    const getCompassDirection = (bearing) => {
      const directions = [
        "North",
        "Northeast",
        "East",
        "Southeast",
        "South",
        "Southwest",
        "West",
        "Northwest",
      ];
      const index = Math.round(bearing / 45) % 8;
      return directions[index];
    };

    const clearCustomRoute = () => {
      if (customRouteLine.value) {
        customRouteLine.value.setMap(null);
        customRouteLine.value = null;
      }
    };

    const clearDirections = () => {
      // Clear any polyline route
      clearCustomRoute();

      showingDirections.value = false;
      currentDirections.value = null;
    };

    const getFacilityIcon = (type) => {
      const colors = {
        aba_center: "#3B82F6", // Blue
        resource_center: "#10B981", // Green
        regional_center: "#8B5CF6", // Purple
        resource: "#F59E0B", // Yellow/Orange
        provider: "#F97316", // Orange
      };

      // Create a reasonably sized, visible marker
      return {
        path: window.google?.maps?.SymbolPath?.CIRCLE || 0,
        scale: 8,
        fillColor: colors[type] || "#6B7280",
        fillOpacity: 0.9,
        strokeColor: "#FFFFFF",
        strokeWeight: 2,
      };
    };

    const calculateDistanceToFacility = (facility) => {
      if (!props.userLocation || !facility.latitude || !facility.longitude) {
        return 0;
      }

      return calculateDistance(
        props.userLocation.lat,
        props.userLocation.lng,
        facility.latitude,
        facility.longitude
      );
    };

    const getFacilityTypeIcon = (type) => {
      const icons = {
        aba_center: "🏥",
        resource_center: "🏢",
        regional_center: "🌍",
        resource: "📚",
        provider: "👥",
      };
      return icons[type] || "📍";
    };

    const fitMapToFacilities = () => {
      if (
        !googleMapInstance.value ||
        !window.google?.maps ||
        allFacilities.value.length === 0
      ) {
        console.warn(
          "Map or Google Maps API not ready, or no facilities to display"
        );
        return;
      }

      try {
        const bounds = new window.google.maps.LatLngBounds();

        // Add user location to bounds if available
        if (props.userLocation) {
          bounds.extend(props.userLocation);
        }

        // Add all facility locations to bounds
        allFacilities.value.forEach((facility) => {
          bounds.extend({
            lat: facility.data.latitude,
            lng: facility.data.longitude,
          });
        });

        // Check if map instance has fitBounds method
        if (typeof googleMapInstance.value.fitBounds === "function") {
          googleMapInstance.value.fitBounds(bounds);

          // Adjust zoom if too close
          window.google.maps.event.addListenerOnce(
            googleMapInstance.value,
            "bounds_changed",
            () => {
              if (
                googleMapInstance.value &&
                typeof googleMapInstance.value.getZoom === "function" &&
                googleMapInstance.value.getZoom() > 15
              ) {
                googleMapInstance.value.setZoom(15);
              }
            }
          );
        } else {
          console.error("Map instance does not have fitBounds method");
          console.log(
            "Available methods:",
            Object.getOwnPropertyNames(googleMapInstance.value)
          );
        }
      } catch (error) {
        console.error("Error fitting map to facilities:", error);
      }
    };

    const centerOnUser = () => {
      if (
        props.userLocation &&
        googleMapInstance.value &&
        typeof googleMapInstance.value.setCenter === "function"
      ) {
        try {
          googleMapInstance.value.setCenter(props.userLocation);
          googleMapInstance.value.setZoom(12);
        } catch (error) {
          console.error("Error centering on user location:", error);
        }
      }
    };

    // Lifecycle
    onMounted(() => {
      // Auto-fit to facilities when component mounts
      // Use a longer delay to ensure Google Maps API is fully loaded
      const checkAndFit = () => {
        if (
          window.google?.maps &&
          googleMapInstance.value &&
          allFacilities.value.length > 0
        ) {
          fitMapToFacilities();
        } else if (allFacilities.value.length > 0) {
          // If facilities exist but map isn't ready, try again
          setTimeout(checkAndFit, 500);
        }
      };

      setTimeout(checkAndFit, 1500);
    });

    return {
      // Refs
      map,
      googleMapInstance,
      selectedFacility,
      showUserLocation,
      mapCenter,
      mapZoom,
      apiKey,

      // Directions state
      showingDirections,
      currentDirections,
      directionsService,
      customRouteLine,

      // Computed
      mapHeight,
      allFacilities,
      mapOptions,
      userLocationIcon,

      // Methods
      onMapReady,
      onMapClick,
      selectFacility,
      clearSelection,
      viewFacilityDetails,
      getDirections,
      createCustomDirections,
      generateIntelligentRoute,
      generateWaypoints,
      generateDirections,
      formatDuration,
      drawRoutePolyline,
      calculateBearing,
      getCompassDirection,
      clearCustomRoute,
      clearDirections,
      getFacilityIcon,
      getFacilityTypeIcon,
      fitMapToFacilities,
      centerOnUser,
      calculateDistance,
      calculateDistanceToFacility,
    };
  },
};
</script>

<style scoped>
/* Custom styles for map container */
.facility-map {
  min-height: 400px;
}

/* Info window styling */
:deep(.gm-style-iw) {
  border-radius: 8px;
}

:deep(.gm-style-iw-d) {
  overflow: visible !important;
}
</style>
