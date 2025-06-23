package models

import (
	"time"

	"github.com/google/uuid"
)

// Provider represents an ABA therapy provider
type Provider struct {
	ID                int       `json:"id" gorm:"primaryKey"`
	Name              string    `json:"name" gorm:"type:varchar(255);not null"`
	Phone             string    `json:"phone" gorm:"type:varchar(100)"`
	Address           string    `json:"address" gorm:"type:text"`
	WebsiteDomain     string    `json:"website_domain" gorm:"type:varchar(255)"`
	Latitude          float64   `json:"latitude" gorm:"type:decimal(10,8)"`
	Longitude         float64   `json:"longitude" gorm:"type:decimal(11,8)"`
	GoogleAddress     string    `json:"google_address" gorm:"type:text"`
	GoogleWebsite     string    `json:"google_website" gorm:"type:varchar(255)"`
	GoogleRating      float64   `json:"google_rating" gorm:"type:decimal(3,2)"`
	GoogleReviewCount float64   `json:"google_review_count" gorm:"type:decimal(10,2)"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	// Many-to-many relationships
	CoverageAreas   []CoverageArea   `json:"coverage_areas" gorm:"many2many:provider_coverage_areas;"`
	InsurancePlans  []InsurancePlan  `json:"insurance_plans" gorm:"many2many:provider_insurance_plans;"`
	Services        []Service        `json:"services" gorm:"many2many:provider_services;"`
	Specializations []Specialization `json:"specializations" gorm:"many2many:provider_specializations;"`
}

// RegionalCenter represents a regional center
type RegionalCenter struct {
	ID                       int       `json:"id" gorm:"primaryKey"`
	RegionalCenter           string    `json:"regional_center" gorm:"type:varchar(255);not null"`
	OfficeType               string    `json:"office_type" gorm:"type:varchar(50)"`
	Address                  string    `json:"address" gorm:"type:varchar(255)"`
	Suite                    string    `json:"suite" gorm:"type:varchar(50)"`
	City                     string    `json:"city" gorm:"type:varchar(100)"`
	State                    string    `json:"state" gorm:"type:varchar(2)"`
	ZipCode                  string    `json:"zip_code" gorm:"type:varchar(10)"`
	Telephone                string    `json:"telephone" gorm:"type:varchar(20)"`
	Website                  string    `json:"website" gorm:"type:varchar(255)"`
	CountyServed             string    `json:"county_served" gorm:"type:varchar(100)"`
	LosAngelesHealthDistrict string    `json:"los_angeles_health_district" gorm:"type:varchar(100)"`
	LocationCoordinates      string    `json:"location_coordinates" gorm:"type:text"`
	Latitude                 float64   `json:"latitude" gorm:"type:decimal(10,8)"`
	Longitude                float64   `json:"longitude" gorm:"type:decimal(11,8)"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

// CoverageArea represents a service coverage area
type CoverageArea struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Many-to-many relationships
	Providers []Provider `json:"providers" gorm:"many2many:provider_coverage_areas;"`
}

// InsurancePlan represents an insurance plan
type InsurancePlan struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Many-to-many relationships
	Providers []Provider `json:"providers" gorm:"many2many:provider_insurance_plans;"`
}

// Service represents a service location or service type
type Service struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Many-to-many relationships
	Providers []Provider `json:"providers" gorm:"many2many:provider_services;"`
}

// Specialization represents a specialization or service area
type Specialization struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Many-to-many relationships
	Providers []Provider `json:"providers" gorm:"many2many:provider_specializations;"`
}

// Junction table structs for explicit modeling if needed
type ProviderCoverageArea struct {
	ProviderID     int `json:"provider_id" gorm:"primaryKey"`
	CoverageAreaID int `json:"coverage_area_id" gorm:"primaryKey"`
}

type ProviderInsurancePlan struct {
	ProviderID      int `json:"provider_id" gorm:"primaryKey"`
	InsurancePlanID int `json:"insurance_plan_id" gorm:"primaryKey"`
}

type ProviderService struct {
	ProviderID int `json:"provider_id" gorm:"primaryKey"`
	ServiceID  int `json:"service_id" gorm:"primaryKey"`
}

type ProviderSpecialization struct {
	ProviderID       int `json:"provider_id" gorm:"primaryKey"`
	SpecializationID int `json:"specialization_id" gorm:"primaryKey"`
}

// Legacy models for backward compatibility (keeping existing structure)

// ABACenter represents an ABA therapy center
type ABACenter struct {
	ID                   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name                 string    `json:"name" gorm:"type:text;not null"`
	Street               string    `json:"street" gorm:"type:text;not null"`
	City                 string    `json:"city" gorm:"type:text;not null"`
	Zip                  string    `json:"zip" gorm:"type:text;not null"`
	Phone                string    `json:"phone" gorm:"type:text;not null"`
	ServiceType          string    `json:"service_type" gorm:"type:text;not null"`
	WaitlistAvailability string    `json:"waitlist_availability" gorm:"type:text"`
	WaitlistNotes        string    `json:"waitlist_notes" gorm:"type:text"`
	DxVerification       string    `json:"dx_verification" gorm:"type:text"`
	InsuranceAccepted    string    `json:"insurance_accepted" gorm:"type:text"`
	MediCalPlans         string    `json:"medi_cal_plans" gorm:"type:text"`
	Notes                string    `json:"notes" gorm:"type:text"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// ResourceCenter represents a resource center
type ResourceCenter struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Address     string    `json:"address" gorm:"type:text"`
	City        string    `json:"city" gorm:"type:varchar(100)"`
	State       string    `json:"state" gorm:"type:varchar(2)"`
	ZipCode     string    `json:"zip_code" gorm:"type:varchar(10)"`
	Phone       string    `json:"phone" gorm:"type:varchar(20)"`
	Email       string    `json:"email" gorm:"type:varchar(255)"`
	Website     string    `json:"website" gorm:"type:varchar(255)"`
	Latitude    float64   `json:"latitude" gorm:"type:decimal(10,8)"`
	Longitude   float64   `json:"longitude" gorm:"type:decimal(11,8)"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Diagnosis represents a diagnosis that resource centers support
type Diagnosis struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Code        string    `json:"code" gorm:"type:varchar(20)"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Resource represents a resource or document
type Resource struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"type:text;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Latitude    float64   `json:"latitude" gorm:"type:numeric;not null"`
	Longitude   float64   `json:"longitude" gorm:"type:numeric;not null"`
	Diagnoses   []string  `json:"diagnoses" gorm:"type:text[]"`
	Address     string    `json:"address" gorm:"type:text"`
	ContactInfo string    `json:"contact_info" gorm:"type:jsonb"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FormSubmission represents a contact form submission
type FormSubmission struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"type:varchar(255);not null"`
	Email        string    `json:"email" gorm:"type:varchar(255);not null"`
	Phone        string    `json:"phone" gorm:"type:varchar(20)"`
	Message      string    `json:"message" gorm:"type:text;not null"`
	FacilityID   string    `json:"facility_id" gorm:"type:varchar(255)"`
	FacilityType string    `json:"facility_type" gorm:"type:varchar(100)"`
	Status       string    `json:"status" gorm:"type:varchar(50);default:'pending'"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// UserPreference represents user preferences
type UserPreference struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID      string    `json:"user_id" gorm:"type:varchar(255);not null;uniqueIndex"`
	Preferences string    `json:"preferences" gorm:"type:jsonb"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserPreferences is an alias for backward compatibility
type UserPreferences = UserPreference

// SearchFilter represents search filtering options
type SearchFilter struct {
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Radius      float64  `json:"radius"`
	Types       []string `json:"types"`
	Keywords    string   `json:"keywords"`
	Insurance   []string `json:"insurance"`
	ServiceType string   `json:"service_type"`
}

// Table name overrides
func (Provider) TableName() string {
	return "providers"
}

func (RegionalCenter) TableName() string {
	return "regional_centers"
}

func (CoverageArea) TableName() string {
	return "coverage_areas"
}

func (InsurancePlan) TableName() string {
	return "insurance_plans"
}

func (Service) TableName() string {
	return "services"
}

func (Specialization) TableName() string {
	return "specializations"
}

func (ProviderCoverageArea) TableName() string {
	return "provider_coverage_areas"
}

func (ProviderInsurancePlan) TableName() string {
	return "provider_insurance_plans"
}

func (ProviderService) TableName() string {
	return "provider_services"
}

func (ProviderSpecialization) TableName() string {
	return "provider_specializations"
}

func (ABACenter) TableName() string {
	return "aba_centers"
}

func (ResourceCenter) TableName() string {
	return "resource_centers"
}

func (Diagnosis) TableName() string {
	return "diagnoses"
}

func (Resource) TableName() string {
	return "resources"
}

func (FormSubmission) TableName() string {
	return "form_submissions"
}

func (UserPreference) TableName() string {
	return "user_preferences"
}
