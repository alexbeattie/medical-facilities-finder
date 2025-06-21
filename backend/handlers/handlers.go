package handlers

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/alexbeattie/medicalfacilities/models"
	"github.com/alexbeattie/medicalfacilities/services"
)

type Handler struct {
	service *services.Service
	db      *gorm.DB
}

func NewHandler(service *services.Service, db *gorm.DB) *Handler {
	return &Handler{
		service: service,
		db:      db,
	}
}

// ABA Centers Handlers

// GetABACenters retrieves ABA centers with optional filtering
func (h *Handler) GetABACenters(c *gin.Context) {
	log.Printf("[GET_ABA_CENTERS] Request received")

	// Parse query parameters
	city := c.Query("city")
	serviceType := c.Query("service_type")
	insurance := c.Query("insurance")
	waitlist := c.Query("waitlist") == "true"
	search := c.Query("search")

	var centers []models.ABACenter
	query := h.db.Model(&models.ABACenter{})

	// Apply filters
	if city != "" {
		query = query.Where("LOWER(city) LIKE ?", "%"+strings.ToLower(city)+"%")
	}

	if serviceType != "" {
		query = query.Where("service_type = ?", serviceType)
	}

	if insurance != "" {
		query = query.Where("LOWER(insurance_accepted) LIKE ?", "%"+strings.ToLower(insurance)+"%")
	}

	if waitlist {
		query = query.Where("waitlist_availability IS NOT NULL AND waitlist_availability != ''")
	}

	// Text search across name, street, and notes
	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(name) LIKE ? OR LOWER(street) LIKE ? OR LOWER(notes) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}

	// Execute query
	if err := query.Find(&centers).Error; err != nil {
		log.Printf("[GET_ABA_CENTERS] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ABA centers"})
		return
	}

	log.Printf("[GET_ABA_CENTERS] Returning %d centers", len(centers))
	c.JSON(http.StatusOK, centers)
}

// GetABACenter retrieves a single ABA center by ID
func (h *Handler) GetABACenter(c *gin.Context) {
	centerID := c.Param("id")
	log.Printf("[GET_ABA_CENTER] Request for center ID: %s", centerID)

	var center models.ABACenter
	if err := h.db.First(&center, "id = ?", centerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "ABA center not found"})
			return
		}
		log.Printf("[GET_ABA_CENTER] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ABA center"})
		return
	}

	c.JSON(http.StatusOK, center)
}

// CreateABACenter creates a new ABA center
func (h *Handler) CreateABACenter(c *gin.Context) {
	log.Printf("[CREATE_ABA_CENTER] Request received")

	var center models.ABACenter
	if err := c.ShouldBindJSON(&center); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if err := h.db.Create(&center).Error; err != nil {
		log.Printf("[CREATE_ABA_CENTER] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ABA center"})
		return
	}

	c.JSON(http.StatusCreated, center)
}

// Resource Centers Handlers

// GetResourceCenters retrieves resource centers with optional filtering
func (h *Handler) GetResourceCenters(c *gin.Context) {
	log.Printf("[GET_RESOURCE_CENTERS] Request received")

	search := c.Query("search")
	latStr := c.Query("lat")
	lngStr := c.Query("lng")
	radiusStr := c.Query("radius")

	var centers []models.ResourceCenter
	query := h.db.Model(&models.ResourceCenter{})

	// Text search
	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(name) LIKE ? OR LOWER(description) LIKE ? OR LOWER(address) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}

	// Execute query
	if err := query.Find(&centers).Error; err != nil {
		log.Printf("[GET_RESOURCE_CENTERS] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resource centers"})
		return
	}

	// Filter by distance if location is provided
	if latStr != "" && lngStr != "" && radiusStr != "" {
		lat, latErr := strconv.ParseFloat(latStr, 64)
		lng, lngErr := strconv.ParseFloat(lngStr, 64)
		radius, radiusErr := strconv.ParseFloat(radiusStr, 64)

		if latErr == nil && lngErr == nil && radiusErr == nil {
			filteredCenters := make([]models.ResourceCenter, 0)
			for _, center := range centers {
				distance := calculateDistance(lat, lng, center.Latitude, center.Longitude)
				if distance <= radius {
					filteredCenters = append(filteredCenters, center)
				}
			}
			centers = filteredCenters
		}
	}

	log.Printf("[GET_RESOURCE_CENTERS] Returning %d centers", len(centers))
	c.JSON(http.StatusOK, centers)
}

// GetResourceCenter retrieves a single resource center by ID
func (h *Handler) GetResourceCenter(c *gin.Context) {
	centerID := c.Param("id")
	log.Printf("[GET_RESOURCE_CENTER] Request for center ID: %s", centerID)

	var center models.ResourceCenter
	if err := h.db.First(&center, "id = ?", centerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource center not found"})
			return
		}
		log.Printf("[GET_RESOURCE_CENTER] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resource center"})
		return
	}

	c.JSON(http.StatusOK, center)
}

// Resources Handlers

// GetResources retrieves resources with optional filtering
func (h *Handler) GetResources(c *gin.Context) {
	log.Printf("[GET_RESOURCES] Request received")

	search := c.Query("search")
	diagnosis := c.Query("diagnosis")

	var resources []models.Resource
	query := h.db.Model(&models.Resource{})

	// Text search
	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(name) LIKE ? OR LOWER(description) LIKE ? OR LOWER(address) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}

	// Filter by diagnosis
	if diagnosis != "" {
		query = query.Where("? = ANY(diagnoses)", diagnosis)
	}

	// Execute query
	if err := query.Find(&resources).Error; err != nil {
		log.Printf("[GET_RESOURCES] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resources"})
		return
	}

	// Note: Resources don't have location data, so distance filtering is not available

	log.Printf("[GET_RESOURCES] Returning %d resources", len(resources))
	c.JSON(http.StatusOK, resources)
}

// GetResource retrieves a single resource by ID
func (h *Handler) GetResource(c *gin.Context) {
	resourceID := c.Param("id")
	log.Printf("[GET_RESOURCE] Request for resource ID: %s", resourceID)

	var resource models.Resource
	if err := h.db.First(&resource, "id = ?", resourceID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
			return
		}
		log.Printf("[GET_RESOURCE] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resource"})
		return
	}

	c.JSON(http.StatusOK, resource)
}

// Regional Centers Handlers

// GetRegionalCenters retrieves regional centers with optional filtering
func (h *Handler) GetRegionalCenters(c *gin.Context) {
	log.Printf("[GET_REGIONAL_CENTERS] Request received")

	county := c.Query("county")
	search := c.Query("search")
	latStr := c.Query("lat")
	lngStr := c.Query("lng")
	radiusStr := c.Query("radius")

	var centers []models.RegionalCenter
	query := h.db.Model(&models.RegionalCenter{})

	// Filter by county
	if county != "" {
		query = query.Where("LOWER(county_served) LIKE ?", "%"+strings.ToLower(county)+"%")
	}

	// Text search
	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(regional_center) LIKE ? OR LOWER(address) LIKE ? OR LOWER(city) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}

	// Execute query
	if err := query.Find(&centers).Error; err != nil {
		log.Printf("[GET_REGIONAL_CENTERS] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch regional centers"})
		return
	}

	// Filter by distance if location is provided
	if latStr != "" && lngStr != "" && radiusStr != "" {
		lat, latErr := strconv.ParseFloat(latStr, 64)
		lng, lngErr := strconv.ParseFloat(lngStr, 64)
		radius, radiusErr := strconv.ParseFloat(radiusStr, 64)

		if latErr == nil && lngErr == nil && radiusErr == nil {
			filteredCenters := make([]models.RegionalCenter, 0)
			for _, center := range centers {
				if center.Latitude != 0 && center.Longitude != 0 {
					distance := calculateDistance(lat, lng, center.Latitude, center.Longitude)
					if distance <= radius {
						filteredCenters = append(filteredCenters, center)
					}
				}
			}
			centers = filteredCenters
		}
	}

	log.Printf("[GET_REGIONAL_CENTERS] Returning %d centers", len(centers))
	c.JSON(http.StatusOK, centers)
}

// Providers Handlers

// GetProviders retrieves providers with optional filtering
func (h *Handler) GetProviders(c *gin.Context) {
	log.Printf("[GET_PROVIDERS] Request received")

	search := c.Query("search")
	area := c.Query("area")

	var providers []models.Provider
	query := h.db.Model(&models.Provider{})

	// Text search
	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(name) LIKE ? OR LOWER(coverage_areas) LIKE ? OR LOWER(center_based_services) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}

	// Filter by area
	if area != "" {
		query = query.Where("? = ANY(areas)", area)
	}

	// Execute query
	if err := query.Find(&providers).Error; err != nil {
		log.Printf("[GET_PROVIDERS] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch providers"})
		return
	}

	log.Printf("[GET_PROVIDERS] Returning %d providers", len(providers))
	c.JSON(http.StatusOK, providers)
}

// Diagnoses Handlers

// GetDiagnoses retrieves all diagnoses
func (h *Handler) GetDiagnoses(c *gin.Context) {
	log.Printf("[GET_DIAGNOSES] Request received")

	var diagnoses []models.Diagnosis
	if err := h.db.Find(&diagnoses).Error; err != nil {
		log.Printf("[GET_DIAGNOSES] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch diagnoses"})
		return
	}

	log.Printf("[GET_DIAGNOSES] Returning %d diagnoses", len(diagnoses))
	c.JSON(http.StatusOK, diagnoses)
}

// Form Submissions Handlers

// CreateFormSubmission creates a new form submission
func (h *Handler) CreateFormSubmission(c *gin.Context) {
	log.Printf("[CREATE_FORM_SUBMISSION] Request received")

	var submission models.FormSubmission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if err := h.db.Create(&submission).Error; err != nil {
		log.Printf("[CREATE_FORM_SUBMISSION] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form submission"})
		return
	}

	c.JSON(http.StatusCreated, submission)
}

// SearchNearby finds all types of facilities within a specified radius
func (h *Handler) SearchNearby(c *gin.Context) {
	latStr := c.Query("lat")
	lngStr := c.Query("lng")
	radiusStr := c.Query("radius")
	entityTypes := c.QueryArray("types") // e.g., ?types=aba_centers&types=resources

	if latStr == "" || lngStr == "" || radiusStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lat, lng, and radius parameters are required"})
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}

	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius"})
		return
	}

	log.Printf("[SEARCH_NEARBY] Searching near lat=%f, lng=%f, radius=%f", lat, lng, radius)

	result := make(map[string]interface{})

	// If no types specified, search all
	if len(entityTypes) == 0 {
		entityTypes = []string{"resource_centers", "regional_centers", "resources"}
	}

	for _, entityType := range entityTypes {
		switch entityType {
		case "resource_centers":
			var centers []models.ResourceCenter
			h.db.Model(&models.ResourceCenter{}).Find(&centers)
			nearby := make([]models.ResourceCenter, 0)
			for _, center := range centers {
				distance := calculateDistance(lat, lng, center.Latitude, center.Longitude)
				if distance <= radius {
					nearby = append(nearby, center)
				}
			}
			result["resource_centers"] = nearby

		case "regional_centers":
			var centers []models.RegionalCenter
			h.db.Find(&centers)
			nearby := make([]models.RegionalCenter, 0)
			for _, center := range centers {
				if center.Latitude != 0 && center.Longitude != 0 {
					distance := calculateDistance(lat, lng, center.Latitude, center.Longitude)
					if distance <= radius {
						nearby = append(nearby, center)
					}
				}
			}
			result["regional_centers"] = nearby

		case "resources":
			var resources []models.Resource
			h.db.Find(&resources)
			// Resources don't have location data, so return all
			result["resources"] = resources
		}
	}

	c.JSON(http.StatusOK, result)
}

// User Preferences (keeping from original)

// GetUserPreferences retrieves user preferences
func (h *Handler) GetUserPreferences(c *gin.Context) {
	userID := c.Param("userId")
	log.Printf("[GET_USER_PREFERENCES] Request for user: %s", userID)

	var preferences models.UserPreferences
	if err := h.db.Where("user_id = ?", userID).First(&preferences).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Return default preferences
			defaultPrefs := models.UserPreferences{
				UserID:      userID,
				Preferences: `{"mapType":"roadmap","defaultZoom":10,"showFacilities":true}`,
			}
			c.JSON(http.StatusOK, defaultPrefs)
			return
		}
		log.Printf("[GET_USER_PREFERENCES] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch preferences"})
		return
	}

	c.JSON(http.StatusOK, preferences)
}

// UpdateUserPreferences updates user preferences
func (h *Handler) UpdateUserPreferences(c *gin.Context) {
	userID := c.Param("userId")
	log.Printf("[UPDATE_USER_PREFERENCES] Request for user: %s", userID)

	var requestData models.UserPreferences
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Ensure the userID matches
	requestData.UserID = userID

	// Upsert the preferences
	if err := h.db.Where("user_id = ?", userID).Assign(requestData).FirstOrCreate(&requestData).Error; err != nil {
		log.Printf("[UPDATE_USER_PREFERENCES] Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update preferences"})
		return
	}

	c.JSON(http.StatusOK, requestData)
}

// Helper function to calculate distance between two coordinates using Haversine formula
func calculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 3959 // Earth's radius in miles

	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	deltaLatRad := (lat2 - lat1) * math.Pi / 180
	deltaLngRad := (lng2 - lng1) * math.Pi / 180

	a := math.Sin(deltaLatRad/2)*math.Sin(deltaLatRad/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLngRad/2)*math.Sin(deltaLngRad/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
