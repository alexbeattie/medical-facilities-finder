<template>
  <!-- Overlay -->
  <div
    v-if="isVisible && facility"
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
    @click="closeModal"
  >
    <!-- Modal -->
    <div
      class="bg-white rounded-lg shadow-xl max-w-lg w-full max-h-[90vh] overflow-hidden"
      @click.stop
    >
      <!-- Header -->
      <div
        class="flex items-center justify-between p-6 border-b border-gray-200"
      >
        <div class="flex items-center space-x-3">
          <div
            class="w-10 h-10 rounded-full flex items-center justify-center text-white text-lg"
            :class="getFacilityTypeColor(facilityType)"
          >
            {{ getFacilityTypeIcon(facilityType) }}
          </div>
          <div>
            <h3 class="text-lg font-semibold text-gray-900">
              {{ facility.name || facility.center_name || "Facility" }}
            </h3>
            <p class="text-sm text-gray-500">
              {{ getFacilityTypeLabel(facilityType) }}
            </p>
          </div>
        </div>
        <button
          @click="closeModal"
          class="text-gray-400 hover:text-gray-600 transition-colors"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            ></path>
          </svg>
        </button>
      </div>

      <!-- Content -->
      <div class="overflow-y-auto max-h-[calc(90vh-140px)]">
        <!-- Basic Info -->
        <div class="p-6 space-y-4">
          <!-- Address -->
          <div
            v-if="facility.address || facility.full_address"
            class="flex items-start space-x-3"
          >
            <div class="w-5 h-5 text-gray-400 mt-0.5">📍</div>
            <div>
              <p class="text-sm font-medium text-gray-700">Address</p>
              <p class="text-sm text-gray-600">
                {{ facility.address || facility.full_address }}
              </p>
              <p
                v-if="facility.city || facility.state"
                class="text-sm text-gray-600"
              >
                {{ [facility.city, facility.state].filter(Boolean).join(", ") }}
              </p>
            </div>
          </div>

          <!-- Contact Info -->
          <div
            v-if="facility.phone || facility.contact_info?.phone"
            class="flex items-start space-x-3"
          >
            <div class="w-5 h-5 text-gray-400 mt-0.5">📞</div>
            <div>
              <p class="text-sm font-medium text-gray-700">Phone</p>
              <a
                :href="`tel:${facility.phone || facility.contact_info?.phone}`"
                class="text-sm text-blue-600 hover:text-blue-800"
              >
                {{ facility.phone || facility.contact_info?.phone }}
              </a>
            </div>
          </div>

          <!-- Email -->
          <div
            v-if="facility.email || facility.contact_info?.email"
            class="flex items-start space-x-3"
          >
            <div class="w-5 h-5 text-gray-400 mt-0.5">✉️</div>
            <div>
              <p class="text-sm font-medium text-gray-700">Email</p>
              <a
                :href="`mailto:${
                  facility.email || facility.contact_info?.email
                }`"
                class="text-sm text-blue-600 hover:text-blue-800"
              >
                {{ facility.email || facility.contact_info?.email }}
              </a>
            </div>
          </div>

          <!-- Website -->
          <div
            v-if="facility.website || facility.contact_info?.website"
            class="flex items-start space-x-3"
          >
            <div class="w-5 h-5 text-gray-400 mt-0.5">🌐</div>
            <div>
              <p class="text-sm font-medium text-gray-700">Website</p>
              <a
                :href="facility.website || facility.contact_info?.website"
                target="_blank"
                rel="noopener noreferrer"
                class="text-sm text-blue-600 hover:text-blue-800"
              >
                {{ facility.website || facility.contact_info?.website }}
              </a>
            </div>
          </div>

          <!-- Distance -->
          <div v-if="distance" class="flex items-start space-x-3">
            <div class="w-5 h-5 text-gray-400 mt-0.5">🗺️</div>
            <div>
              <p class="text-sm font-medium text-gray-700">Distance</p>
              <p class="text-sm text-gray-600">{{ distance }} miles away</p>
            </div>
          </div>
        </div>

        <!-- ABA Center Specific Info -->
        <div
          v-if="facilityType === 'aba_center'"
          class="px-6 py-4 bg-blue-50 border-y border-blue-100"
        >
          <h4 class="text-sm font-semibold text-blue-900 mb-3">
            ABA Center Details
          </h4>
          <div class="space-y-3">
            <!-- Service Type -->
            <div
              v-if="facility.service_type"
              class="flex items-center space-x-2"
            >
              <span class="text-sm font-medium text-blue-700"
                >Service Type:</span
              >
              <span class="text-sm text-blue-600">{{
                facility.service_type
              }}</span>
            </div>

            <!-- Insurance -->
            <div
              v-if="facility.insurance_accepted"
              class="flex items-center space-x-2"
            >
              <span class="text-sm font-medium text-blue-700">Insurance:</span>
              <span class="text-sm text-blue-600">{{
                facility.insurance_accepted
              }}</span>
            </div>

            <!-- Waitlist -->
            <div
              v-if="facility.waitlist_availability"
              class="flex items-center space-x-2"
            >
              <span class="text-sm font-medium text-blue-700">Waitlist:</span>
              <span class="text-sm text-blue-600">{{
                facility.waitlist_availability
              }}</span>
            </div>
          </div>
        </div>

        <!-- Resource Center Specific Info -->
        <div
          v-if="facilityType === 'resource_center'"
          class="px-6 py-4 bg-green-50 border-y border-green-100"
        >
          <h4 class="text-sm font-semibold text-green-900 mb-3">
            Resource Center Details
          </h4>
          <div class="space-y-3">
            <!-- Diagnoses -->
            <div
              v-if="facility.diagnoses && facility.diagnoses.length > 0"
              class="space-y-2"
            >
              <span class="text-sm font-medium text-green-700"
                >Supported Diagnoses:</span
              >
              <div class="flex flex-wrap gap-1">
                <span
                  v-for="diagnosis in facility.diagnoses"
                  :key="diagnosis.id"
                  class="inline-block px-2 py-1 bg-green-100 text-green-800 text-xs rounded-full"
                >
                  {{ diagnosis.name }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Regional Center Specific Info -->
        <div
          v-if="facilityType === 'regional_center'"
          class="px-6 py-4 bg-purple-50 border-y border-purple-100"
        >
          <h4 class="text-sm font-semibold text-purple-900 mb-3">
            Regional Center Details
          </h4>
          <div class="space-y-3">
            <!-- County -->
            <div v-if="facility.county" class="flex items-center space-x-2">
              <span class="text-sm font-medium text-purple-700">County:</span>
              <span class="text-sm text-purple-600">{{ facility.county }}</span>
            </div>

            <!-- Service Area -->
            <div
              v-if="facility.service_area"
              class="flex items-center space-x-2"
            >
              <span class="text-sm font-medium text-purple-700"
                >Service Area:</span
              >
              <span class="text-sm text-purple-600">{{
                facility.service_area
              }}</span>
            </div>
          </div>
        </div>

        <!-- Description -->
        <div v-if="facility.description" class="p-6">
          <h4 class="text-sm font-semibold text-gray-900 mb-2">About</h4>
          <p class="text-sm text-gray-600 leading-relaxed">
            {{ facility.description }}
          </p>
        </div>

        <!-- Hours -->
        <div v-if="facility.hours" class="px-6 pb-4">
          <h4 class="text-sm font-semibold text-gray-900 mb-2">Hours</h4>
          <p class="text-sm text-gray-600">{{ facility.hours }}</p>
        </div>
      </div>

      <!-- Actions -->
      <div class="p-6 border-t border-gray-200 bg-gray-50">
        <div class="flex space-x-3">
          <!-- Get Directions -->
          <button
            v-if="facility.latitude && facility.longitude"
            @click="getDirections"
            class="flex-1 bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 transition-colors duration-200 flex items-center justify-center space-x-2"
          >
            <span>🗺️</span>
            <span>Directions</span>
          </button>

          <!-- Call -->
          <button
            v-if="facility.phone || facility.contact_info?.phone"
            @click="callFacility"
            class="flex-1 bg-green-600 text-white py-2 px-4 rounded-lg hover:bg-green-700 transition-colors duration-200 flex items-center justify-center space-x-2"
          >
            <span>📞</span>
            <span>Call</span>
          </button>

          <!-- Contact Form -->
          <button
            @click="showContactForm"
            class="flex-1 bg-purple-600 text-white py-2 px-4 rounded-lg hover:bg-purple-700 transition-colors duration-200 flex items-center justify-center space-x-2"
          >
            <span>✉️</span>
            <span>Contact</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from "vue";
import { useFacilities } from "@/composables/useFacilities";

export default {
  name: "FacilityPopup",
  props: {
    facility: {
      type: Object,
      default: null,
    },
    facilityType: {
      type: String,
      required: true,
    },
    isVisible: {
      type: Boolean,
      default: false,
    },
    userLocation: {
      type: Object,
      default: null,
    },
  },
  emits: ["close", "contact-form"],
  setup(props, { emit }) {
    const { calculateDistance } = useFacilities();

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
    const closeModal = () => {
      emit("close");
    };

    const getFacilityTypeIcon = (type) => {
      const icons = {
        aba_center: "🏥",
        resource_center: "🏢",
        regional_center: "🌍",
        resource: "📚",
        provider: "👥",
      };
      return icons[type] || "🏢";
    };

    const getFacilityTypeLabel = (type) => {
      const labels = {
        aba_center: "ABA Center",
        resource_center: "Resource Center",
        regional_center: "Regional Center",
        resource: "Resource",
        provider: "Service Provider",
      };
      return labels[type] || "Facility";
    };

    const getFacilityTypeColor = (type) => {
      const colors = {
        aba_center: "bg-blue-500",
        resource_center: "bg-green-500",
        regional_center: "bg-purple-500",
        resource: "bg-yellow-500",
        provider: "bg-orange-500",
      };
      return colors[type] || "bg-gray-500";
    };

    const getDirections = () => {
      if (!props.facility.latitude || !props.facility.longitude) return;

      const url = `https://www.google.com/maps/dir/?api=1&destination=${props.facility.latitude},${props.facility.longitude}`;
      window.open(url, "_blank");
    };

    const callFacility = () => {
      const phone = props.facility.phone || props.facility.contact_info?.phone;
      if (phone) {
        window.location.href = `tel:${phone}`;
      }
    };

    const showContactForm = () => {
      emit("contact-form", props.facility);
    };

    return {
      // Computed
      distance,

      // Methods
      closeModal,
      getFacilityTypeIcon,
      getFacilityTypeLabel,
      getFacilityTypeColor,
      getDirections,
      callFacility,
      showContactForm,
    };
  },
};
</script>

<style scoped>
/* Custom scrollbar for modal content */
.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* Ensure modal appears above map */
.z-50 {
  z-index: 50;
}
</style>
