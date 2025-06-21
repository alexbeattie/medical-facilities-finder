// src/services/api.js

class ApiService {
  constructor() {
    // Environment-based API configuration
    this.baseUrl = this.getApiBaseUrl();
    this.headers = {
      'Content-Type': 'application/json',
    };

    // Log the current environment for debugging
    console.log(`API Service initialized with baseUrl: ${this.baseUrl}`);
    console.log(`Environment mode: ${process.env.NODE_ENV}`);
  }

  getApiBaseUrl() {
    // Priority order for API base URL:
    // 1. Explicit environment variable
    // 2. Environment-specific defaults
    // 3. Production fallback

    if (process.env.VUE_APP_API_BASE_URL) {
      return process.env.VUE_APP_API_BASE_URL;
    }

    // Environment-specific defaults
    switch (process.env.NODE_ENV) {
      case 'development':
        return 'http://localhost:8080';
      case 'production':
        return 'https://medicalfacilities.com';
      case 'staging':
        return process.env.VUE_APP_STAGING_API_URL || 'https://staging.medicalfacilities.com';
      default:
        return 'http://localhost:8080';
    }
  }

  async handleResponse(response) {
    if (!response.ok) {
      const error = await response.json().catch(() => ({}));
      throw new Error(error.message || `HTTP error! status: ${response.status}`);
    }
    return response.json();
  }

  // ===== ABA CENTERS =====

  // Get all ABA centers with optional filtering
  async getABACenters(filters = {}) {
    try {
      const params = new URLSearchParams();

      if (filters.city) params.append('city', filters.city);
      if (filters.service_type) params.append('service_type', filters.service_type);
      if (filters.insurance) params.append('insurance', filters.insurance);
      if (filters.waitlist) params.append('waitlist', 'true');
      if (filters.search) params.append('search', filters.search);

      const queryString = params.toString();
      const url = `${this.baseUrl}/api/v1/aba-centers${queryString ? '?' + queryString : ''}`;

      console.log(`Fetching ABA centers from: ${url}`);
      const response = await fetch(url, {
        method: 'GET',
        headers: this.headers
      });

      return this.handleResponse(response);
    } catch (error) {
      console.error('Error fetching ABA centers:', error);
      throw error;
    }
  }

  // Get single ABA center by ID
  async getABACenter(centerId) {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/aba-centers/${centerId}`, {
        method: 'GET',
        headers: this.headers
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error(`Error fetching ABA center ${centerId}:`, error);
      throw error;
    }
  }

  // Create a new ABA center (admin function)
  async createABACenter(centerData) {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/aba-centers`, {
        method: 'POST',
        headers: this.headers,
        body: JSON.stringify(centerData)
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error('Error creating ABA center:', error);
      throw error;
    }
  }

  // ===== RESOURCE CENTERS =====

  // Get all resource centers with optional filtering
  async getResourceCenters(filters = {}) {
    try {
      const params = new URLSearchParams();

      if (filters.search) params.append('search', filters.search);
      if (filters.lat) params.append('lat', filters.lat);
      if (filters.lng) params.append('lng', filters.lng);
      if (filters.radius) params.append('radius', filters.radius);

      const queryString = params.toString();
      const url = `${this.baseUrl}/api/v1/resource-centers${queryString ? '?' + queryString : ''}`;

      console.log(`Fetching resource centers from: ${url}`);
      const response = await fetch(url, {
        method: 'GET',
        headers: this.headers
      });

      return this.handleResponse(response);
    } catch (error) {
      console.error('Error fetching resource centers:', error);
      throw error;
    }
  }

  // Get single resource center by ID
  async getResourceCenter(centerId) {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/resource-centers/${centerId}`, {
        method: 'GET',
        headers: this.headers
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error(`Error fetching resource center ${centerId}:`, error);
      throw error;
    }
  }

  // ===== RESOURCES =====

  // Get all resources with optional filtering
  async getResources(filters = {}) {
    try {
      const params = new URLSearchParams();

      if (filters.search) params.append('search', filters.search);
      if (filters.diagnosis) params.append('diagnosis', filters.diagnosis);
      if (filters.lat) params.append('lat', filters.lat);
      if (filters.lng) params.append('lng', filters.lng);
      if (filters.radius) params.append('radius', filters.radius);

      const queryString = params.toString();
      const url = `${this.baseUrl}/api/v1/resources${queryString ? '?' + queryString : ''}`;

      console.log(`Fetching resources from: ${url}`);
      const response = await fetch(url, {
        method: 'GET',
        headers: this.headers
      });

      return this.handleResponse(response);
    } catch (error) {
      console.error('Error fetching resources:', error);
      throw error;
    }
  }

  // Get single resource by ID
  async getResource(resourceId) {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/resources/${resourceId}`, {
        method: 'GET',
        headers: this.headers
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error(`Error fetching resource ${resourceId}:`, error);
      throw error;
    }
  }

  // ===== REGIONAL CENTERS =====

  // Get all regional centers with optional filtering
  async getRegionalCenters(filters = {}) {
    try {
      const params = new URLSearchParams();

      if (filters.county) params.append('county', filters.county);
      if (filters.search) params.append('search', filters.search);
      if (filters.lat) params.append('lat', filters.lat);
      if (filters.lng) params.append('lng', filters.lng);
      if (filters.radius) params.append('radius', filters.radius);

      const queryString = params.toString();
      const url = `${this.baseUrl}/api/v1/regional-centers${queryString ? '?' + queryString : ''}`;

      console.log(`Fetching regional centers from: ${url}`);
      const response = await fetch(url, {
        method: 'GET',
        headers: this.headers
      });

      return this.handleResponse(response);
    } catch (error) {
      console.error('Error fetching regional centers:', error);
      throw error;
    }
  }

  // ===== PROVIDERS =====

  // Get all providers with optional filtering
  async getProviders(filters = {}) {
    try {
      const params = new URLSearchParams();

      if (filters.search) params.append('search', filters.search);
      if (filters.area) params.append('area', filters.area);

      const queryString = params.toString();
      const url = `${this.baseUrl}/api/v1/providers${queryString ? '?' + queryString : ''}`;

      console.log(`Fetching providers from: ${url}`);
      const response = await fetch(url, {
        method: 'GET',
        headers: this.headers
      });

      return this.handleResponse(response);
    } catch (error) {
      console.error('Error fetching providers:', error);
      throw error;
    }
  }

  // ===== DIAGNOSES =====

  // Get all diagnoses
  async getDiagnoses() {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/diagnoses`, {
        method: 'GET',
        headers: this.headers
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error('Error fetching diagnoses:', error);
      throw error;
    }
  }

  // ===== SEARCH =====

  // Search for nearby facilities of all types
  async searchNearby(lat, lng, radius, types = []) {
    try {
      const params = new URLSearchParams({
        lat: lat.toString(),
        lng: lng.toString(),
        radius: radius.toString()
      });

      if (types.length > 0) {
        types.forEach(type => params.append('types', type));
      }

      const response = await fetch(
        `${this.baseUrl}/api/v1/search/nearby?${params.toString()}`,
        {
          method: 'GET',
          headers: this.headers
        }
      );

      return this.handleResponse(response);
    } catch (error) {
      console.error('Error searching nearby facilities:', error);
      throw error;
    }
  }

  // ===== FORM SUBMISSIONS =====

  // Submit a contact form
  async submitForm(formData) {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/form-submissions`, {
        method: 'POST',
        headers: this.headers,
        body: JSON.stringify(formData)
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error('Error submitting form:', error);
      throw error;
    }
  }

  // ===== USER PREFERENCES =====

  // Get user preferences
  async getUserPreferences(userId) {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/preferences/${userId}`, {
        method: 'GET',
        headers: this.headers
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error(`Error fetching preferences for user ${userId}:`, error);
      throw error;
    }
  }

  // Update user preferences
  async updateUserPreferences(userId, preferences) {
    try {
      const response = await fetch(`${this.baseUrl}/api/v1/preferences/${userId}`, {
        method: 'PUT',
        headers: this.headers,
        body: JSON.stringify(preferences)
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error(`Error updating preferences for user ${userId}:`, error);
      throw error;
    }
  }

  // ===== UTILITY =====

  // Health check
  async healthCheck() {
    try {
      const response = await fetch(`${this.baseUrl}/health`, {
        method: 'GET',
        headers: this.headers
      });
      return this.handleResponse(response);
    } catch (error) {
      console.error('Health check failed:', error);
      throw error;
    }
  }
}

export const apiService = new ApiService();
