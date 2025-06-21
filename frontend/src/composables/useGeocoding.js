// composables/useGeocoding.js
import { ref } from 'vue';

export function useGeocoding(apiKey) {
  const isLoading = ref(false);
  const error = ref(null);

  const geocodeAddress = async (address) => {
    if (!address || !address.trim()) {
      throw new Error('Invalid address');
    }

    isLoading.value = true;
    error.value = null;

    try {
      const encodedAddress = encodeURIComponent(address.trim());
      const response = await fetch(
        `https://maps.googleapis.com/maps/api/geocode/json?address=${encodedAddress}&key=${apiKey}`
      );

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();

      if (data.status === "OK" && data.results?.[0]) {
        const result = data.results[0];
        return {
          lat: result.geometry.location.lat,
          lng: result.geometry.location.lng,
          display_name: result.formatted_address,
          place_id: result.place_id
        };
      }

      throw new Error(`Geocoding failed: ${data.status}`);
    } catch (err) {
      error.value = err;
      throw err;
    } finally {
      isLoading.value = false;
    }
  };

  const reverseGeocode = async (lat, lng) => {
    if (!lat || !lng) {
      throw new Error('Invalid coordinates');
    }

    isLoading.value = true;
    error.value = null;

    try {
      const response = await fetch(
        `https://maps.googleapis.com/maps/api/geocode/json?latlng=${lat},${lng}&key=${apiKey}`
      );

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();

      if (data.status === "OK" && data.results?.[0]) {
        return data.results[0].formatted_address;
      }

      throw new Error(`Geocoding failed: ${data.status}`);
    } catch (err) {
      error.value = err;
      throw err;
    } finally {
      isLoading.value = false;
    }
  };

  return {
    isLoading,
    error,
    geocodeAddress,
    reverseGeocode
  };
}