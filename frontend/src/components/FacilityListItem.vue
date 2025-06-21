<template>
  <div
    @click="selectFacility"
    class="p-4 hover:bg-gray-50 cursor-pointer transition-colors duration-150 border-l-4"
    :class="getBorderColor(facilityType)"
  >
    <div class="flex items-start justify-between">
      <div class="flex-1 min-w-0">
        <!-- Name and Type -->
        <div class="flex items-center space-x-2 mb-2">
          <span class="text-lg" :class="getIconColor(facilityType)">
            {{ getFacilityIcon(facilityType) }}
          </span>
          <h4 class="text-base font-semibold text-gray-900 truncate">
            {{ getFacilityName(facility) }}
          </h4>
        </div>

        <!-- Address -->
        <div
          v-if="facility.address || facility.full_address"
          class="flex items-start space-x-2 mb-2"
        >
          <span class="text-xs text-gray-400 mt-0.5">📍</span>
          <p class="text-sm text-gray-600 line-clamp-2">
            {{ facility.address || facility.full_address }}
            <span v-if="facility.city || facility.state" class="block">
              {{ [facility.city, facility.state].filter(Boolean).join(", ") }}
            </span>
          </p>
        </div>

        <!-- Specific Info by Type -->
        <div class="space-y-1">
          <!-- ABA Center Info -->
          <div v-if="facilityType === 'aba_center'" class="text-xs space-y-1">
            <div
              v-if="facility.service_type"
              class="flex items-center space-x-1"
            >
              <span class="text-blue-600">Service:</span>
              <span class="text-gray-600">{{ facility.service_type }}</span>
            </div>
            <div
              v-if="facility.insurance_accepted"
              class="flex items-center space-x-1"
            >
              <span class="text-blue-600">Insurance:</span>
              <span class="text-gray-600">{{
                facility.insurance_accepted
              }}</span>
            </div>
          </div>

          <!-- Resource Center Info -->
          <div v-if="facilityType === 'resource_center'" class="text-xs">
            <div
              v-if="facility.diagnoses?.length"
              class="flex items-start space-x-1"
            >
              <span class="text-green-600 flex-shrink-0">Diagnoses:</span>
              <div class="flex flex-wrap gap-1">
                <span
                  v-for="diagnosis in facility.diagnoses.slice(0, 3)"
                  :key="diagnosis.id"
                  class="inline-block px-1.5 py-0.5 bg-green-100 text-green-700 rounded text-xs"
                >
                  {{ diagnosis.name }}
                </span>
                <span
                  v-if="facility.diagnoses.length > 3"
                  class="text-green-600 text-xs"
                >
                  +{{ facility.diagnoses.length - 3 }} more
                </span>
              </div>
            </div>
          </div>

          <!-- Regional Center Info -->
          <div
            v-if="facilityType === 'regional_center'"
            class="text-xs space-y-1"
          >
            <div v-if="facility.county" class="flex items-center space-x-1">
              <span class="text-purple-600">County:</span>
              <span class="text-gray-600">{{ facility.county }}</span>
            </div>
            <div
              v-if="facility.service_area"
              class="flex items-center space-x-1"
            >
              <span class="text-purple-600">Area:</span>
              <span class="text-gray-600">{{ facility.service_area }}</span>
            </div>
          </div>

          <!-- Resource Info -->
          <div v-if="facilityType === 'resource'" class="text-xs">
            <div v-if="facility.description" class="text-gray-600 line-clamp-2">
              {{ facility.description }}
            </div>
          </div>

          <!-- Provider Info -->
          <div v-if="facilityType === 'provider'" class="text-xs">
            <div v-if="facility.area" class="flex items-center space-x-1">
              <span class="text-orange-600">Service Area:</span>
              <span class="text-gray-600">{{ facility.area }}</span>
            </div>
          </div>
        </div>

        <!-- Contact Info -->
        <div class="flex items-center space-x-4 mt-2 text-xs text-gray-500">
          <div
            v-if="facility.phone || facility.contact_info?.phone"
            class="flex items-center space-x-1"
          >
            <span>📞</span>
            <span>{{ facility.phone || facility.contact_info?.phone }}</span>
          </div>
          <div v-if="distance" class="flex items-center space-x-1">
            <span>🗺️</span>
            <span>{{ distance }} mi</span>
          </div>
        </div>
      </div>

      <!-- Action Button -->
      <div class="flex-shrink-0 ml-4">
        <button
          @click.stop="selectFacility"
          class="px-3 py-1.5 text-xs font-medium rounded-lg transition-colors duration-150"
          :class="getButtonColor(facilityType)"
        >
          Details
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from "vue";
import { useFacilities } from "@/composables/useFacilities";

export default {
  name: "FacilityListItem",
  props: {
    facility: {
      type: Object,
      required: true,
    },
    facilityType: {
      type: String,
      required: true,
    },
    userLocation: {
      type: Object,
      default: null,
    },
  },
  emits: ["select"],
  setup(props, { emit }) {
    const { calculateDistance } = useFacilities();

    // Utility function to get facility name with fallbacks
    const getFacilityName = (facilityData) => {
      if (!facilityData) return "Unknown Facility";

      return (
        facilityData.name ||
        facilityData.regional_center || // For regional centers
        facilityData.center_name ||
        facilityData.facility_name ||
        facilityData.organization_name ||
        facilityData.provider_name ||
        facilityData.title ||
        facilityData.display_name ||
        "Unknown Facility"
      );
    };

    // Computed
    const distance = computed(() => {
      if (!props.userLocation || !props.facility) return null;

      const facilityLat = props.facility.latitude;
      const facilityLng = props.facility.longitude;

      if (!facilityLat || !facilityLng) return null;

      const dist = calculateDistance(
        props.userLocation.lat,
        props.userLocation.lng,
        facilityLat,
        facilityLng
      );

      return dist.toFixed(1);
    });

    // Methods
    const selectFacility = () => {
      emit("select", props.facility, props.facilityType);
    };

    const getFacilityIcon = (type) => {
      const icons = {
        aba_center: "🏥",
        resource_center: "🏢",
        regional_center: "🌍",
        resource: "📚",
        provider: "👥",
      };
      return icons[type] || "🏢";
    };

    const getBorderColor = (type) => {
      const colors = {
        aba_center: "border-blue-400",
        resource_center: "border-green-400",
        regional_center: "border-purple-400",
        resource: "border-yellow-400",
        provider: "border-orange-400",
      };
      return colors[type] || "border-gray-400";
    };

    const getIconColor = (type) => {
      const colors = {
        aba_center: "text-blue-500",
        resource_center: "text-green-500",
        regional_center: "text-purple-500",
        resource: "text-yellow-500",
        provider: "text-orange-500",
      };
      return colors[type] || "text-gray-500";
    };

    const getButtonColor = (type) => {
      const colors = {
        aba_center: "bg-blue-100 text-blue-700 hover:bg-blue-200",
        resource_center: "bg-green-100 text-green-700 hover:bg-green-200",
        regional_center: "bg-purple-100 text-purple-700 hover:bg-purple-200",
        resource: "bg-yellow-100 text-yellow-700 hover:bg-yellow-200",
        provider: "bg-orange-100 text-orange-700 hover:bg-orange-200",
      };
      return colors[type] || "bg-gray-100 text-gray-700 hover:bg-gray-200";
    };

    return {
      // Computed
      distance,

      // Methods
      selectFacility,
      getFacilityName,
      getFacilityIcon,
      getBorderColor,
      getIconColor,
      getButtonColor,
    };
  },
};
</script>

<style scoped>
/* Line clamp utility */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
