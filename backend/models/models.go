package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Custom Point type for PostGIS geometry support (simplified)
type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Scan implements the sql.Scanner interface
func (p *Point) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	// For now, we'll let PostgreSQL handle the geometry/geography conversion
	// and just store lat/lng as separate fields in the database queries
	return nil
}

// Value implements the driver.Valuer interface
func (p Point) Value() (driver.Value, error) {
	if p.Lat == 0 && p.Lng == 0 {
		return nil, nil
	}
	// For now, return nil to avoid complex PostGIS integration
	// The database already has the geometry data
	return nil, nil
}

// ContactInfo represents JSON contact information
type ContactInfo map[string]interface{}

// Scan implements the sql.Scanner interface for ContactInfo
func (c *ContactInfo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, c)
	case string:
		return json.Unmarshal([]byte(v), c)
	}
	return nil
}

// Value implements the driver.Valuer interface for ContactInfo
func (c ContactInfo) Value() (driver.Value, error) {
	if len(c) == 0 {
		return nil, nil
	}
	return json.Marshal(c)
}

// ABACenter represents an ABA therapy center
type ABACenter struct {
	ID                   uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name                 string     `json:"name" gorm:"not null"`
	Street               string     `json:"street" gorm:"not null"`
	City                 string     `json:"city" gorm:"not null"`
	Zip                  string     `json:"zip" gorm:"not null"`
	Phone                string     `json:"phone" gorm:"not null"`
	ServiceType          string     `json:"service_type" gorm:"not null"`
	WaitlistAvailability *string    `json:"waitlist_availability"`
	WaitlistNotes        *string    `json:"waitlist_notes"`
	DxVerification       *string    `json:"dx_verification"`
	InsuranceAccepted    *string    `json:"insurance_accepted"`
	MediCalPlans         *string    `json:"medi_cal_plans"`
	Notes                *string    `json:"notes"`
	CreatedAt            *time.Time `json:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at"`
}

// Diagnosis represents a medical diagnosis
type Diagnosis struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string    `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
}

// FormSubmission represents form submissions from users
type FormSubmission struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// Permission represents system permissions
type Permission struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null;uniqueIndex"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`
}

// Provider represents service providers
type Provider struct {
	ID                  int      `json:"id" gorm:"primaryKey"`
	Name                string   `json:"name" gorm:"not null"`
	Phone               *string  `json:"phone"`
	CoverageAreas       *string  `json:"coverage_areas"`
	CenterBasedServices *string  `json:"center_based_services"`
	Areas               []string `json:"areas" gorm:"type:text[]"`
}

// RegionalCenter represents regional centers with geospatial data
type RegionalCenter struct {
	ID                       int      `json:"id" gorm:"primaryKey"`
	RegionalCenter           *string  `json:"regional_center"`
	OfficeType               *string  `json:"office_type"`
	Address                  *string  `json:"address"`
	Suite                    *string  `json:"suite"`
	City                     *string  `json:"city"`
	State                    *string  `json:"state"`
	ZipCode                  *string  `json:"zip_code"`
	Telephone                *string  `json:"telephone"`
	Website                  *string  `json:"website"`
	CountyServed             *string  `json:"county_served"`
	LosAngelesHealthDistrict *string  `json:"los_angeles_health_district"`
	LocationCoordinates      *string  `json:"location_coordinates"`
	Latitude                 *float64 `json:"latitude"`
	Longitude                *float64 `json:"longitude"`
	Location                 *Point   `json:"location" gorm:"type:geography(POINT,4326)"`
}

// ResourceCenter represents resource centers with geospatial data
type ResourceCenter struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string     `json:"name" gorm:"type:varchar(255);not null"`
	Description *string    `json:"description"`
	Address     *string    `json:"address"`
	Latitude    float64    `json:"latitude" gorm:"not null"`
	Longitude   float64    `json:"longitude" gorm:"not null"`
	Location    *Point     `json:"location" gorm:"type:geometry(POINT,4326)"`
	CreatedAt   *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	// Many-to-many relationship with diagnoses
	Diagnoses []Diagnosis `json:"diagnoses" gorm:"many2many:center_diagnoses;foreignKey:ID;joinForeignKey:center_id;References:ID;joinReferences:diagnosis_id"`
}

// Resource represents resources with geospatial data and diagnoses
type Resource struct {
	ID          uuid.UUID   `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string      `json:"name" gorm:"not null"`
	Description *string     `json:"description"`
	Latitude    float64     `json:"latitude" gorm:"type:numeric;not null"`
	Longitude   float64     `json:"longitude" gorm:"type:numeric;not null"`
	Diagnoses   []string    `json:"diagnoses" gorm:"type:text[]"`
	Address     *string     `json:"address"`
	ContactInfo ContactInfo `json:"contact_info" gorm:"type:jsonb"`
	CreatedAt   time.Time   `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Location    *Point      `json:"location" gorm:"type:geography(POINT,4326)"`
}

// Role represents user roles
type Role struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(50);not null;uniqueIndex"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:now()"`

	// Many-to-many relationship with permissions
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;foreignKey:ID;joinForeignKey:role_id;References:ID;joinReferences:permission_id"`
}

// User represents system users
type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"type:varchar(255);not null;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"type:varchar(255);not null"`
	FirstName    *string   `json:"first_name" gorm:"type:varchar(100)"`
	LastName     *string   `json:"last_name" gorm:"type:varchar(100)"`
	Name         *string   `json:"name" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:now()"`

	// Many-to-many relationship with roles
	Roles []Role `json:"roles" gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:user_id;References:ID;joinReferences:role_id"`
}

// Junction table models (GORM will create these automatically, but we can define them for explicit control)

// CenterDiagnosis represents the many-to-many relationship between resource centers and diagnoses
type CenterDiagnosis struct {
	CenterID    uuid.UUID `json:"center_id" gorm:"type:uuid;primaryKey"`
	DiagnosisID uuid.UUID `json:"diagnosis_id" gorm:"type:uuid;primaryKey"`
}

// RolePermission represents the many-to-many relationship between roles and permissions
type RolePermission struct {
	RoleID       int `json:"role_id" gorm:"primaryKey"`
	PermissionID int `json:"permission_id" gorm:"primaryKey"`
}

// UserRole represents the many-to-many relationship between users and roles
type UserRole struct {
	UserID int `json:"user_id" gorm:"primaryKey"`
	RoleID int `json:"role_id" gorm:"primaryKey"`
}

// Search and filter models for API queries
type SearchFilter struct {
	Diagnoses         []string `json:"diagnoses"`
	MaxDistance       float64  `json:"max_distance"`
	Latitude          float64  `json:"latitude"`
	Longitude         float64  `json:"longitude"`
	ServiceType       string   `json:"service_type"`
	InsuranceRequired bool     `json:"insurance_required"`
	WaitlistOnly      bool     `json:"waitlist_only"`
}

// UserPreferences for user settings (keeping this from original for compatibility)
type UserPreferences struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID string `json:"user_id" gorm:"unique;not null"`

	// Map preferences
	MapType        string `json:"map_type" gorm:"default:'roadmap'"`
	DefaultZoom    int    `json:"default_zoom" gorm:"default:10"`
	ShowFacilities bool   `json:"show_facilities" gorm:"default:true"`

	// Facility type filters
	ShowABACenters      bool `json:"show_aba_centers" gorm:"default:true"`
	ShowResourceCenters bool `json:"show_resource_centers" gorm:"default:true"`
	ShowRegionalCenters bool `json:"show_regional_centers" gorm:"default:true"`
	ShowProviders       bool `json:"show_providers" gorm:"default:true"`

	// Search preferences
	PreferredRadius    float64 `json:"preferred_radius" gorm:"default:25"` // miles
	RequireWaitlist    bool    `json:"require_waitlist" gorm:"default:false"`
	RequireInsurance   bool    `json:"require_insurance" gorm:"default:false"`
	PreferredDiagnoses string  `json:"preferred_diagnoses"` // JSON array of diagnosis IDs

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Table name overrides
func (ABACenter) TableName() string {
	return "aba_centers"
}

func (Diagnosis) TableName() string {
	return "diagnoses"
}

func (FormSubmission) TableName() string {
	return "form_submissions"
}

func (Permission) TableName() string {
	return "permissions"
}

func (Provider) TableName() string {
	return "providers"
}

func (RegionalCenter) TableName() string {
	return "regional_centers"
}

func (ResourceCenter) TableName() string {
	return "resource_centers"
}

func (Resource) TableName() string {
	return "resources"
}

func (Role) TableName() string {
	return "roles"
}

func (User) TableName() string {
	return "users"
}

func (CenterDiagnosis) TableName() string {
	return "center_diagnoses"
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

func (UserRole) TableName() string {
	return "user_roles"
}
