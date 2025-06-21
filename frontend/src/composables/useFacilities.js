import { ref, reactive, computed } from 'vue'
import { apiService } from '@/services/api'

const state = reactive({
  abaCenters: [],
  resourceCenters: [],
  resources: [],
  regionalCenters: [],
  providers: [],
  diagnoses: [],
  loading: false,
  error: null,
  searchResults: null,
  userLocation: null,
  selectedFacility: null
})

export function useFacilities() {
  const loading = ref(false)
  const error = ref(null)

  // ===== DATA GETTERS =====

  const abaCenters = computed(() => state.abaCenters)
  const resourceCenters = computed(() => state.resourceCenters)
  const resources = computed(() => state.resources)
  const regionalCenters = computed(() => state.regionalCenters)
  const providers = computed(() => state.providers)
  const diagnoses = computed(() => state.diagnoses)
  const searchResults = computed(() => state.searchResults)
  const selectedFacility = computed(() => state.selectedFacility)
  const userLocation = computed(() => state.userLocation)

  // ===== ABA CENTERS =====

  const fetchABACenters = async (filters = {}) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getABACenters(filters)
      state.abaCenters = data || []
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching ABA centers:', err)
      return []
    } finally {
      loading.value = false
    }
  }

  const fetchABACenter = async (id) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getABACenter(id)
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching ABA center:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  // ===== RESOURCE CENTERS =====

  const fetchResourceCenters = async (filters = {}) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getResourceCenters(filters)
      state.resourceCenters = data || []
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching resource centers:', err)
      return []
    } finally {
      loading.value = false
    }
  }

  const fetchResourceCenter = async (id) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getResourceCenter(id)
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching resource center:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  // ===== RESOURCES =====

  const fetchResources = async (filters = {}) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getResources(filters)
      state.resources = data || []
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching resources:', err)
      return []
    } finally {
      loading.value = false
    }
  }

  const fetchResource = async (id) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getResource(id)
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching resource:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  // ===== REGIONAL CENTERS =====

  const fetchRegionalCenters = async (filters = {}) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getRegionalCenters(filters)
      state.regionalCenters = data || []
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching regional centers:', err)
      return []
    } finally {
      loading.value = false
    }
  }

  // ===== PROVIDERS =====

  const fetchProviders = async (filters = {}) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getProviders(filters)
      state.providers = data || []
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching providers:', err)
      return []
    } finally {
      loading.value = false
    }
  }

  // ===== DIAGNOSES =====

  const fetchDiagnoses = async () => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.getDiagnoses()
      state.diagnoses = data || []
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error fetching diagnoses:', err)
      return []
    } finally {
      loading.value = false
    }
  }

  // ===== SEARCH & LOCATION =====

  const searchNearby = async (lat, lng, radius = 25, types = []) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.searchNearby(lat, lng, radius, types)
      state.searchResults = data
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error searching nearby facilities:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  const setUserLocation = (location) => {
    state.userLocation = location
  }

  const getUserLocation = () => {
    return new Promise((resolve, reject) => {
      if (!navigator.geolocation) {
        reject(new Error('Geolocation is not supported by this browser.'))
        return
      }

      navigator.geolocation.getCurrentPosition(
        (position) => {
          const location = {
            lat: position.coords.latitude,
            lng: position.coords.longitude,
            accuracy: position.coords.accuracy
          }
          setUserLocation(location)
          resolve(location)
        },
        (error) => {
          reject(error)
        },
        {
          enableHighAccuracy: true,
          timeout: 10000,
          maximumAge: 300000 // 5 minutes
        }
      )
    })
  }

  // ===== FACILITY SELECTION =====

  const setSelectedFacility = (facility) => {
    state.selectedFacility = facility
  }

  const clearSelectedFacility = () => {
    state.selectedFacility = null
  }

  // ===== FORM SUBMISSION =====

  const submitContactForm = async (formData) => {
    try {
      loading.value = true
      error.value = null
      const data = await apiService.submitForm(formData)
      return data
    } catch (err) {
      error.value = err.message
      console.error('Error submitting form:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  // ===== UTILITY FUNCTIONS =====

  const calculateDistance = (lat1, lng1, lat2, lng2) => {
    const R = 3959 // Earth's radius in miles
    const dLat = (lat2 - lat1) * Math.PI / 180
    const dLng = (lng2 - lng1) * Math.PI / 180
    const a =
      Math.sin(dLat / 2) * Math.sin(dLat / 2) +
      Math.cos(lat1 * Math.PI / 180) * Math.cos(lat2 * Math.PI / 180) *
      Math.sin(dLng / 2) * Math.sin(dLng / 2)
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a))
    return R * c
  }

  const clearError = () => {
    error.value = null
  }

  const clearSearchResults = () => {
    state.searchResults = null
  }

  const clearAllData = () => {
    state.abaCenters = []
    state.resourceCenters = []
    state.resources = []
    state.regionalCenters = []
    state.providers = []
    state.searchResults = null
    state.selectedFacility = null
    error.value = null
  }

  return {
    // State
    loading: computed(() => loading.value),
    error: computed(() => error.value),
    abaCenters,
    resourceCenters,
    resources,
    regionalCenters,
    providers,
    diagnoses,
    searchResults,
    selectedFacility,
    userLocation,

    // ABA Centers
    fetchABACenters,
    fetchABACenter,

    // Resource Centers
    fetchResourceCenters,
    fetchResourceCenter,

    // Resources
    fetchResources,
    fetchResource,

    // Regional Centers
    fetchRegionalCenters,

    // Providers
    fetchProviders,

    // Diagnoses
    fetchDiagnoses,

    // Search & Location
    searchNearby,
    setUserLocation,
    getUserLocation,

    // Facility Selection
    setSelectedFacility,
    clearSelectedFacility,

    // Forms
    submitContactForm,

    // Utilities
    calculateDistance,
    clearError,
    clearSearchResults,
    clearAllData
  }
} 