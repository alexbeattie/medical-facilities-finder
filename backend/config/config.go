// config/config.go
package config

// Config holds the application configuration
type Config struct {
	// Database connection string
	DSN string

	// Google Maps API key for geocoding and map services
	GoogleMapsAPIKey string

	// Server configuration
	Port string
	Host string

	// Environment
	Environment string

	// Logging configuration
	LogLevel string
}
