# Medical Facilities Finder

A full-stack web application for finding and managing medical facilities like hospitals, clinics, urgent care centers, and pharmacies. Built with Vue.js frontend and Go backend, featuring interactive maps, filtering, and search capabilities.

## 🏥 Features

- **Interactive Map**: Google Maps integration with facility markers
- **Advanced Filtering**: Filter by facility type, rating, services, and location
- **Location-based Search**: Find facilities within a specified radius
- **Facility Information**: Detailed facility pages with services, reviews, and contact info
- **Responsive Design**: Mobile-friendly interface with sidebar navigation
- **Real-time Updates**: Dynamic facility information and availability

## 🏗️ Architecture

### Backend (Go)
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL with GORM ORM
- **Features**: RESTful API, CORS support, environment-based configuration

### Frontend (Vue.js)
- **Framework**: Vue.js 3 with Vuex for state management
- **Styling**: Tailwind CSS for responsive design
- **Maps**: Vue3-Google-Map for interactive mapping
- **Build Tool**: Vue CLI with environment-specific configurations

## 🚀 Quick Start

### Prerequisites

- Go 1.22+ 
- Node.js 16+
- PostgreSQL 12+
- Google Maps API Key

### Backend Setup

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Create environment file:**
   Create `.env.local` in the backend directory:
   ```bash
   GOOGLE_MAPS_API_KEY=your_google_maps_api_key_here
   DSN="host=localhost user=yourusername dbname=medical_facilities password=yourpassword port=5432 sslmode=disable TimeZone=America/Los_Angeles"
   APP_PORT=8080
   ```

4. **Set up PostgreSQL database:**
   ```sql
   CREATE DATABASE medical_facilities;
   ```

5. **Start the backend:**
   ```bash
   go run main.go
   ```

The backend will start on `http://localhost:8080` and automatically:
- Run database migrations
- Seed sample data in development mode
- Set up API endpoints

### Frontend Setup

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install dependencies:**
   ```bash
   yarn install --ignore-engines
   ```

3. **Create environment file:**
   Create `.env.local` in the frontend directory:
   ```bash
   VUE_APP_GOOGLE_MAPS_API_KEY=your_google_maps_api_key_here
   VUE_APP_API_BASE_URL=http://localhost:8080
   ```

4. **Start the development server:**
   ```bash
   yarn serve
   ```

The frontend will start on `http://localhost:8082` (or another available port).

## 📊 Database Schema

### Medical Facilities
- **id**: Primary key
- **name**: Facility name
- **facility_type**: hospital, clinic, urgent_care, pharmacy, specialist
- **address, city, state, zip_code**: Location information
- **latitude, longitude**: GPS coordinates
- **phone_number, website, email**: Contact information
- **is_operational**: Current operational status
- **emergency_services**: Provides emergency care
- **accepts_insurance**: Accepts insurance
- **rating, total_reviews**: User ratings
- **capacity, current_occupancy**: Capacity information
- **specialties, services**: JSON arrays of offered services

### User Preferences
- **user_id**: Unique user identifier
- **map_type**: Map display type (roadmap, satellite, etc.)
- **show_[facility_types]**: Display preferences for each facility type
- **preferred_radius**: Default search radius
- **minimum_rating**: Minimum rating filter

## 🔗 API Endpoints

### Facilities
- `GET /api/v1/facilities` - Get all facilities (with optional filters)
- `GET /api/v1/facilities/:id` - Get specific facility
- `POST /api/v1/facilities` - Create facility (admin)
- `PUT /api/v1/facilities/:id` - Update facility (admin)

### Search
- `GET /api/v1/search/nearby` - Find facilities within radius

### Facility Details
- `GET /api/v1/facilities/:id/services` - Get facility services
- `GET /api/v1/facilities/:id/facility_types` - Get facility facility types
<!-- - `GET /api/v1/facilities/:id/reviews` - Get facility reviews -->

### User Preferences
- `GET /api/v1/preferences/:userId` - Get user preferences
- `PUT /api/v1/preferences/:userId` - Update user preferences

### Utility
- `GET /health` - Health check
- `POST /seed` - Seed sample data (development)

## 🎛️ Environment Configuration

The application supports multiple environments:

### Development
```bash
yarn serve:dev    # Frontend development mode
NODE_ENV=development go run main.go
```

### Production
```bash
yarn build:prod   # Production build
APP_ENV=production go run main.go
```

### Environment Variables

**Backend (.env.local):**
```bash
GOOGLE_MAPS_API_KEY=your_api_key
DSN=postgresql_connection_string
APP_PORT=8080
APP_ENV=development
```

**Frontend (.env.local):**
```bash
VUE_APP_GOOGLE_MAPS_API_KEY=your_api_key
VUE_APP_API_BASE_URL=http://localhost:8080
```

## 🏥 Facility Types

- **hospital**: Full-service hospitals
- **clinic**: General medical clinics
- **urgent_care**: Urgent care centers
- **pharmacy**: Pharmacies and drug stores
- **specialist**: Specialist medical practices

## 🗺️ Features Overview

### Map Interface
- Interactive Google Maps with facility markers
- Clustered markers for better performance
- Custom info windows with facility details
- Map type switching (roadmap, satellite, hybrid, terrain)

### Filtering & Search
- Filter by facility type
- Search by name, address, or services
- Location-based radius search
- Rating and service filters
- Emergency services filter

### Facility Management
- Detailed facility profiles
- Service listings
- User reviews and ratings
- Contact information
- Operating hours and capacity

## 🔧 Development

### Adding New Facility Types
1. Update the `facility_type` enum in the database
2. Add the type to the frontend filters
3. Update the API filtering logic

### Customizing Map Appearance
Edit the Google Maps configuration in `frontend/src/App.vue`:
```javascript
:options="{
  mapTypeControl: false,
  streetViewControl: false,
  // Add your custom options
}"
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📝 License

MIT License - see LICENSE file for details.

