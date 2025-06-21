# Database Setup Guide

This guide will help you connect your Go backend to your PostgreSQL database with the ABA centers schema.

## Prerequisites

1. PostgreSQL database with PostGIS extension enabled
2. Your database should contain the tables defined in your provided schema
3. Go 1.22+ installed
4. Required Go dependencies (already included in go.mod)

## Environment Variables

Create a `.env.local` file in the backend directory with the following configuration:

```env
# Database Configuration
# Format: postgresql://username:password@hostname:port/database_name?sslmode=disable
DSN=postgresql://alexbeattie:your_password@localhost:5432/your_database_name?sslmode=disable

# Google Maps API Key (for geocoding and map services)
GOOGLE_MAPS_API_KEY=your_google_maps_api_key_here

# Server Configuration
APP_PORT=8080
APP_ENV=development

# Logging Configuration
LOG_LEVEL=info
```

## Database Connection String Examples

### Local PostgreSQL
```
DSN=postgresql://username:password@localhost:5432/database_name?sslmode=disable
```

### Remote PostgreSQL with SSL
```
DSN=postgresql://username:password@hostname:5432/database_name?sslmode=require
```

### PostgreSQL on Railway/Render/Heroku
```
DSN=postgresql://username:password@hostname:port/database_name?sslmode=require
```

## Testing the Connection

1. Make sure your PostgreSQL server is running
2. Verify your database contains the required tables:
   - aba_centers
   - diagnoses
   - providers
   - regional_centers
   - resource_centers
   - resources
   - users, roles, permissions (authentication tables)

3. Start the backend server:
```bash
cd backend
go run main.go
```

4. Test the API endpoints:
```bash
# Health check
curl http://localhost:8080/health

# Get ABA centers
curl http://localhost:8080/api/v1/aba-centers

# Get resource centers
curl http://localhost:8080/api/v1/resource-centers

# Get diagnoses
curl http://localhost:8080/api/v1/diagnoses
```

## Available API Endpoints

### ABA Centers
- `GET /api/v1/aba-centers` - List all ABA centers with optional filtering
- `GET /api/v1/aba-centers/:id` - Get specific ABA center
- `POST /api/v1/aba-centers` - Create new ABA center (admin)

### Resource Centers
- `GET /api/v1/resource-centers` - List resource centers
- `GET /api/v1/resource-centers/:id` - Get specific resource center

### Resources
- `GET /api/v1/resources` - List resources with diagnosis filtering
- `GET /api/v1/resources/:id` - Get specific resource

### Regional Centers
- `GET /api/v1/regional-centers` - List regional centers

### Providers
- `GET /api/v1/providers` - List providers

### Diagnoses
- `GET /api/v1/diagnoses` - List all diagnoses

### Search
- `GET /api/v1/search/nearby?lat=34.0522&lng=-118.2437&radius=25&types=resource_centers,resources` - Search nearby facilities

### User Preferences
- `GET /api/v1/preferences/:userId` - Get user preferences
- `PUT /api/v1/preferences/:userId` - Update user preferences

## Query Parameters

### ABA Centers
- `city` - Filter by city
- `service_type` - Filter by service type
- `insurance` - Filter by insurance accepted
- `waitlist=true` - Only show centers with waitlist availability
- `search` - Text search across name, street, notes

### Resource Centers & Resources
- `search` - Text search
- `lat`, `lng`, `radius` - Location-based filtering
- `diagnosis` - Filter resources by diagnosis (for resources endpoint)

### Regional Centers
- `county` - Filter by county served
- `search` - Text search
- `lat`, `lng`, `radius` - Location-based filtering

## Database Models

The backend now includes models for all your database tables:
- `ABACenter` - ABA therapy centers
- `ResourceCenter` - Resource centers with diagnosis relationships
- `Resource` - Individual resources with geospatial data
- `RegionalCenter` - Regional centers with geospatial data
- `Provider` - Service providers
- `Diagnosis` - Medical diagnoses
- `User`, `Role`, `Permission` - Authentication system
- `FormSubmission` - Contact form submissions

## PostGIS Support

The backend includes support for PostGIS geometry and geography types used in your database for geospatial queries and distance calculations.

## Troubleshooting

### Connection Issues
1. Verify PostgreSQL is running: `pg_ctl status`
2. Check if the database exists: `psql -l`
3. Verify credentials and permissions
4. Ensure PostGIS extension is installed: `SELECT PostGIS_version();`

### Missing Tables
If you get "table doesn't exist" errors, verify your database schema matches the provided SQL definitions.

### Permission Errors
Ensure your database user has SELECT/INSERT/UPDATE/DELETE permissions on all tables. 