package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alexbeattie/medicalfacilities/config"
	"github.com/alexbeattie/medicalfacilities/handlers"
	"github.com/alexbeattie/medicalfacilities/models"
	"github.com/alexbeattie/medicalfacilities/services"
)

func initLogger() (*os.File, error) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}
	logFile := filepath.Join("logs", fmt.Sprintf("app_%s.log", time.Now().Format("2006-01-02")))
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return f, nil
}

func loadEnvFile() {
	envFile := ".env.local" // Default to local
	if os.Getenv("APP_ENV") == "production" {
		envFile = ".env.production"
	}
	log.Printf("[LOAD_ENV] Time: %s | File: %s", time.Now().Format("2006-01-02 15:04:05"), envFile)
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("[LOAD_ENV] Warning: Could not load %s file: %v", envFile, err)
	} else {
		log.Printf("[LOAD_ENV] Successfully loaded environment variables from file: %s", envFile)
	}
}

func initDB(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("database connection string (DSN) is empty")
	}
	log.Printf("Attempting to connect to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Printf("Successfully connected to database")

	// Skip auto-migration since tables already exist
	// Only migrate UserPreference which is safe
	if err := db.AutoMigrate(&models.UserPreference{}); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}
	log.Printf("Database migrations completed successfully")
	return db, nil
}

func setupRouter(handler *handlers.Handler, service *services.Service, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173", // Vite default dev server
			"http://localhost:4173", // Vite preview
			"http://127.0.0.1:5173", // Alternative localhost
			"http://localhost:8080",
			"http://localhost:8081",
			"http://localhost:8082",
			"http://localhost:8083", // Vue dev server
			"http://localhost:8084", // Vue dev server (current)
			"https://medicalfacilities.com",
			"http://medicalfacilities.com",
			"https://www.medicalfacilities.com",
			"http://www.medicalfacilities.com",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve static files for the frontend
	r.Static("/assets", "./dist/assets")
	r.Static("/js", "./dist/js")
	r.Static("/css", "./dist/css")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")

	// API routes for ABA centers and resources
	api := r.Group("/api/v1")
	{
		// User preferences
		api.GET("/preferences/:userId", handler.GetUserPreferences)
		api.PUT("/preferences/:userId", handler.UpdateUserPreferences)

		// ABA Centers
		api.GET("/aba-centers", handler.GetABACenters)
		api.GET("/aba-centers/:id", handler.GetABACenter)
		api.POST("/aba-centers", handler.CreateABACenter)        // Admin function
		api.POST("/aba-centers/submit", handler.SubmitABACenter) // User submission

		// Resource Centers
		api.GET("/resource-centers", handler.GetResourceCenters)
		api.GET("/resource-centers/:id", handler.GetResourceCenter)

		// Resources
		api.GET("/resources", handler.GetResources)
		api.GET("/resources/:id", handler.GetResource)

		// Regional Centers
		api.GET("/regional-centers", handler.GetRegionalCenters)
		api.POST("/regional-centers/submit-update", handler.SubmitRegionalCenterUpdate) // User update submission

		// Providers
		api.GET("/providers", handler.GetProviders)

		// Diagnoses
		api.GET("/diagnoses", handler.GetDiagnoses)

		// Form submissions
		api.POST("/form-submissions", handler.CreateFormSubmission)

		// Search endpoints
		api.GET("/search/nearby", handler.SearchNearby)
	}

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
			"service":   "Medical Facilities API",
		})
	})

	// Seed data endpoint (for development)
	r.POST("/seed", func(c *gin.Context) {
		if err := service.SeedSampleData(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to seed data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Sample data seeded successfully"})
	})

	// Handle the root route and any unmatched routes with the Vue app
	r.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s | %d | %s | %s | %s %s | Host: %s | Origin: %s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Host,
			param.Request.Header.Get("Origin"),
		)
	}))

	return r
}

func main() {
	logFile, err := initLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logFile.Close()

	log.Println("Starting Medical Facilities API server...")

	loadEnvFile()

	// Check for required environment variables
	requiredEnvVars := []string{"DSN"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Required environment variable %s is not set", envVar)
		}
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN is not set")
	}

	cfg := &config.Config{
		GoogleMapsAPIKey: "", // Not used - app uses Mapbox
		DSN:              os.Getenv("DSN"),
	}

	db, err := initDB(cfg.DSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	service := services.NewService(db, cfg)
	handler := handlers.NewHandler(service, db)

	r := setupRouter(handler, service, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
		log.Printf("APP_PORT not set, defaulting to %s", port)
	}
	address := fmt.Sprintf("0.0.0.0:%s", port)

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	// Start HTTP server
	go func() {
		log.Printf("Starting Medical Facilities API server on %s", address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Seed sample data in development
	if os.Getenv("APP_ENV") != "production" {
		go func() {
			time.Sleep(2 * time.Second) // Wait for server to start
			if err := service.SeedSampleData(); err != nil {
				log.Printf("Failed to seed sample data: %v", err)
			}
		}()
	}

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	log.Printf("Received shutdown signal: %v", sig)
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	if sqlDB, err := db.DB(); err == nil {
		sqlDB.Close()
		log.Println("Database connection closed.")
	}

	log.Println("Medical Facilities API server exited properly")
}
