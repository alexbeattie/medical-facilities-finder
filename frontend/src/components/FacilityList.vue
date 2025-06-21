<template>
  <div class="facility-list bg-white rounded-lg shadow-md overflow-hidden">
    <!-- Header -->
    <div class="p-4 border-b border-gray-200 bg-gray-50">
      <h3 class="text-lg font-semibold text-gray-800">Search Results</h3>
      <p v-if="totalResults > 0" class="text-sm text-gray-600">
        Found {{ totalResults }} facilities
      </p>
      <p v-else class="text-sm text-gray-500">
        No facilities found. Try adjusting your search criteria.
      </p>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="p-8 text-center">
      <div
        class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"
      ></div>
      <p class="mt-2 text-gray-600">Searching for facilities...</p>
    </div>

    <!-- Results -->
    <div v-else-if="totalResults > 0" class="max-h-96 overflow-y-auto">
      <!-- ABA Centers -->
      <div
        v-if="facilities.abaCenters?.length"
        class="border-b border-gray-100 last:border-b-0"
      >
        <div class="p-3 bg-blue-50 border-b border-blue-100">
          <h4 class="text-sm font-semibold text-blue-900 flex items-center">
            <span class="mr-2">🏥</span>
            ABA Centers ({{ facilities.abaCenters.length }})
          </h4>
        </div>
        <div class="divide-y divide-gray-100">
          <FacilityListItem
            v-for="center in facilities.abaCenters"
            :key="`aba_${center.id}`"
            :facility="center"
            facility-type="aba_center"
            :user-location="userLocation"
            @select="onFacilitySelect"
          />
        </div>
      </div>

      <!-- Resource Centers -->
      <div
        v-if="facilities.resourceCenters?.length"
        class="border-b border-gray-100 last:border-b-0"
      >
        <div class="p-3 bg-green-50 border-b border-green-100">
          <h4 class="text-sm font-semibold text-green-900 flex items-center">
            <span class="mr-2">🏢</span>
            Resource Centers ({{ facilities.resourceCenters.length }})
          </h4>
        </div>
        <div class="divide-y divide-gray-100">
          <FacilityListItem
            v-for="center in facilities.resourceCenters"
            :key="`resource_${center.id}`"
            :facility="center"
            facility-type="resource_center"
            :user-location="userLocation"
            @select="onFacilitySelect"
          />
        </div>
      </div>

      <!-- Regional Centers -->
      <div
        v-if="facilities.regionalCenters?.length"
        class="border-b border-gray-100 last:border-b-0"
      >
        <div class="p-3 bg-purple-50 border-b border-purple-100">
          <h4 class="text-sm font-semibold text-purple-900 flex items-center">
            <span class="mr-2">🌍</span>
            Regional Centers ({{ facilities.regionalCenters.length }})
          </h4>
        </div>
        <div class="divide-y divide-gray-100">
          <FacilityListItem
            v-for="center in facilities.regionalCenters"
            :key="`regional_${center.id}`"
            :facility="center"
            facility-type="regional_center"
            :user-location="userLocation"
            @select="onFacilitySelect"
          />
        </div>
      </div>

      <!-- Resources -->
      <div
        v-if="facilities.resources?.length"
        class="border-b border-gray-100 last:border-b-0"
      >
        <div class="p-3 bg-yellow-50 border-b border-yellow-100">
          <h4 class="text-sm font-semibold text-yellow-900 flex items-center">
            <span class="mr-2">📚</span>
            Resources ({{ facilities.resources.length }})
          </h4>
        </div>
        <div class="divide-y divide-gray-100">
          <FacilityListItem
            v-for="resource in facilities.resources"
            :key="`resource_item_${resource.id}`"
            :facility="resource"
            facility-type="resource"
            :user-location="userLocation"
            @select="onFacilitySelect"
          />
        </div>
      </div>

      <!-- Providers -->
      <div
        v-if="facilities.providers?.length"
        class="border-b border-gray-100 last:border-b-0"
      >
        <div class="p-3 bg-orange-50 border-b border-orange-100">
          <h4 class="text-sm font-semibold text-orange-900 flex items-center">
            <span class="mr-2">👥</span>
            Service Providers ({{ facilities.providers.length }})
          </h4>
        </div>
        <div class="divide-y divide-gray-100">
          <FacilityListItem
            v-for="provider in facilities.providers"
            :key="`provider_${provider.id}`"
            :facility="provider"
            facility-type="provider"
            :user-location="userLocation"
            @select="onFacilitySelect"
          />
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="!loading" class="p-8 text-center text-gray-500">
      <div class="text-4xl mb-4">🔍</div>
      <h4 class="text-lg font-medium text-gray-600 mb-2">No Results Found</h4>
      <p class="text-sm">
        Try expanding your search radius or selecting different facility types.
      </p>
    </div>
  </div>
</template>

<script>
import { computed } from "vue";
import FacilityListItem from "./FacilityListItem.vue";

export default {
  name: "FacilityList",
  components: {
    FacilityListItem,
  },
  props: {
    facilities: {
      type: Object,
      default: () => ({}),
    },
    loading: {
      type: Boolean,
      default: false,
    },
    userLocation: {
      type: Object,
      default: null,
    },
  },
  emits: ["facility-select"],
  setup(props, { emit }) {
    // Computed
    const totalResults = computed(() => {
      if (!props.facilities) return 0;

      const counts = [
        props.facilities.abaCenters?.length || 0,
        props.facilities.resourceCenters?.length || 0,
        props.facilities.regionalCenters?.length || 0,
        props.facilities.resources?.length || 0,
        props.facilities.providers?.length || 0,
      ];

      return counts.reduce((total, count) => total + count, 0);
    });

    // Methods
    const onFacilitySelect = (facility, facilityType) => {
      emit("facility-select", facility, facilityType);
    };

    return {
      // Computed
      totalResults,

      // Methods
      onFacilitySelect,
    };
  },
};
</script>

<style scoped>
/* Custom scrollbar for facility list */
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
</style>
