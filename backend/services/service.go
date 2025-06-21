// services/service.go
package services

import (
	"fmt"
	"log"
	"math"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/alexbeattie/medicalfacilities/config"
	"github.com/alexbeattie/medicalfacilities/models"
)

type Service struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewService(db *gorm.DB, cfg *config.Config) *Service {
	return &Service{
		db:  db,
		cfg: cfg,
	}
}

// ABA Centers Services

// GetABACenters retrieves ABA centers with filtering
func (s *Service) GetABACenters(filter *models.SearchFilter) ([]models.ABACenter, error) {
	var centers []models.ABACenter
	query := s.db.Model(&models.ABACenter{})

	if filter.ServiceType != "" {
		query = query.Where("service_type = ?", filter.ServiceType)
	}

	if filter.InsuranceRequired {
		query = query.Where("insurance_accepted IS NOT NULL AND insurance_accepted != ''")
	}

	if filter.WaitlistOnly {
		query = query.Where("waitlist_availability IS NOT NULL AND waitlist_availability != ''")
	}

	if err := query.Find(&centers).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch ABA centers: %w", err)
	}

	return centers, nil
}

// GetABACenterByID retrieves a single ABA center by ID
func (s *Service) GetABACenterByID(id uuid.UUID) (*models.ABACenter, error) {
	var center models.ABACenter
	if err := s.db.First(&center, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("ABA center not found")
		}
		return nil, fmt.Errorf("failed to fetch ABA center: %w", err)
	}
	return &center, nil
}

// CreateABACenter creates a new ABA center
func (s *Service) CreateABACenter(center *models.ABACenter) error {
	if err := s.db.Create(center).Error; err != nil {
		return fmt.Errorf("failed to create ABA center: %w", err)
	}
	return nil
}

// Resource Centers Services

// GetResourceCenters retrieves resource centers with filtering
func (s *Service) GetResourceCenters(filter *models.SearchFilter) ([]models.ResourceCenter, error) {
	var centers []models.ResourceCenter
	query := s.db.Model(&models.ResourceCenter{}).Preload("Diagnoses")

	if err := query.Find(&centers).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch resource centers: %w", err)
	}

	// Filter by distance if location is provided
	if filter.Latitude != 0 && filter.Longitude != 0 && filter.MaxDistance > 0 {
		filteredCenters := make([]models.ResourceCenter, 0)
		for _, center := range centers {
			distance := calculateDistance(filter.Latitude, filter.Longitude, center.Latitude, center.Longitude)
			if distance <= filter.MaxDistance {
				filteredCenters = append(filteredCenters, center)
			}
		}
		centers = filteredCenters
	}

	return centers, nil
}

// GetResourceCenterByID retrieves a single resource center by ID
func (s *Service) GetResourceCenterByID(id uuid.UUID) (*models.ResourceCenter, error) {
	var center models.ResourceCenter
	if err := s.db.Preload("Diagnoses").First(&center, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("resource center not found")
		}
		return nil, fmt.Errorf("failed to fetch resource center: %w", err)
	}
	return &center, nil
}

// Resources Services

// GetResources retrieves resources with filtering
func (s *Service) GetResources(filter *models.SearchFilter) ([]models.Resource, error) {
	var resources []models.Resource
	query := s.db.Model(&models.Resource{})

	// Filter by diagnosis
	if len(filter.Diagnoses) > 0 {
		for _, diagnosis := range filter.Diagnoses {
			query = query.Where("? = ANY(diagnoses)", diagnosis)
		}
	}

	if err := query.Find(&resources).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch resources: %w", err)
	}

	// Filter by distance if location is provided
	if filter.Latitude != 0 && filter.Longitude != 0 && filter.MaxDistance > 0 {
		filteredResources := make([]models.Resource, 0)
		for _, resource := range resources {
			distance := calculateDistance(filter.Latitude, filter.Longitude, resource.Latitude, resource.Longitude)
			if distance <= filter.MaxDistance {
				filteredResources = append(filteredResources, resource)
			}
		}
		resources = filteredResources
	}

	return resources, nil
}

// GetResourceByID retrieves a single resource by ID
func (s *Service) GetResourceByID(id uuid.UUID) (*models.Resource, error) {
	var resource models.Resource
	if err := s.db.First(&resource, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("resource not found")
		}
		return nil, fmt.Errorf("failed to fetch resource: %w", err)
	}
	return &resource, nil
}

// Regional Centers Services

// GetRegionalCenters retrieves regional centers with filtering
func (s *Service) GetRegionalCenters(filter *models.SearchFilter) ([]models.RegionalCenter, error) {
	var centers []models.RegionalCenter
	query := s.db.Model(&models.RegionalCenter{})

	if err := query.Find(&centers).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch regional centers: %w", err)
	}

	// Filter by distance if location is provided
	if filter.Latitude != 0 && filter.Longitude != 0 && filter.MaxDistance > 0 {
		filteredCenters := make([]models.RegionalCenter, 0)
		for _, center := range centers {
			if center.Latitude != nil && center.Longitude != nil {
				distance := calculateDistance(filter.Latitude, filter.Longitude, *center.Latitude, *center.Longitude)
				if distance <= filter.MaxDistance {
					filteredCenters = append(filteredCenters, center)
				}
			}
		}
		centers = filteredCenters
	}

	return centers, nil
}

// Providers Services

// GetProviders retrieves providers with filtering
func (s *Service) GetProviders() ([]models.Provider, error) {
	var providers []models.Provider
	if err := s.db.Find(&providers).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch providers: %w", err)
	}
	return providers, nil
}

// Diagnoses Services

// GetDiagnoses retrieves all diagnoses
func (s *Service) GetDiagnoses() ([]models.Diagnosis, error) {
	var diagnoses []models.Diagnosis
	if err := s.db.Find(&diagnoses).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch diagnoses: %w", err)
	}
	return diagnoses, nil
}

// Form Submissions Services

// CreateFormSubmission creates a new form submission
func (s *Service) CreateFormSubmission(submission *models.FormSubmission) error {
	if err := s.db.Create(submission).Error; err != nil {
		return fmt.Errorf("failed to create form submission: %w", err)
	}
	return nil
}

// Search Services

// SearchNearby finds various entities within a radius
func (s *Service) SearchNearby(lat, lng, radiusMiles float64, entityTypes []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// If no types specified, search all
	if len(entityTypes) == 0 {
		entityTypes = []string{"resource_centers", "regional_centers", "resources"}
	}

	for _, entityType := range entityTypes {
		switch entityType {
		case "resource_centers":
			var centers []models.ResourceCenter
			s.db.Model(&models.ResourceCenter{}).Preload("Diagnoses").Find(&centers)
			nearby := make([]models.ResourceCenter, 0)
			for _, center := range centers {
				distance := calculateDistance(lat, lng, center.Latitude, center.Longitude)
				if distance <= radiusMiles {
					nearby = append(nearby, center)
				}
			}
			result["resource_centers"] = nearby

		case "regional_centers":
			var centers []models.RegionalCenter
			s.db.Find(&centers)
			nearby := make([]models.RegionalCenter, 0)
			for _, center := range centers {
				if center.Latitude != nil && center.Longitude != nil {
					distance := calculateDistance(lat, lng, *center.Latitude, *center.Longitude)
					if distance <= radiusMiles {
						nearby = append(nearby, center)
					}
				}
			}
			result["regional_centers"] = nearby

		case "resources":
			var resources []models.Resource
			s.db.Find(&resources)
			nearby := make([]models.Resource, 0)
			for _, resource := range resources {
				distance := calculateDistance(lat, lng, resource.Latitude, resource.Longitude)
				if distance <= radiusMiles {
					nearby = append(nearby, resource)
				}
			}
			result["resources"] = nearby
		}
	}

	return result, nil
}

// User Preferences Services

// GetUserPreferences retrieves user preferences
func (s *Service) GetUserPreferences(userID string) (*models.UserPreferences, error) {
	var preferences models.UserPreferences
	if err := s.db.Where("user_id = ?", userID).First(&preferences).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Return default preferences
			return &models.UserPreferences{
				UserID:              userID,
				MapType:             "roadmap",
				DefaultZoom:         10,
				ShowFacilities:      true,
				ShowABACenters:      true,
				ShowResourceCenters: true,
				ShowRegionalCenters: true,
				ShowProviders:       true,
				PreferredRadius:     25,
				RequireWaitlist:     false,
				RequireInsurance:    false,
			}, nil
		}
		return nil, fmt.Errorf("failed to fetch preferences: %w", err)
	}
	return &preferences, nil
}

// UpdateUserPreferences updates user preferences
func (s *Service) UpdateUserPreferences(userID string, preferences *models.UserPreferences) error {
	preferences.UserID = userID
	if err := s.db.Where("user_id = ?", userID).Assign(preferences).FirstOrCreate(preferences).Error; err != nil {
		return fmt.Errorf("failed to update preferences: %w", err)
	}
	return nil
}

// SeedSampleData - simplified for ABA database (no seeding needed since database exists)
func (s *Service) SeedSampleData() error {
	log.Println("Using existing database - no sample data needed")
	return nil
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
