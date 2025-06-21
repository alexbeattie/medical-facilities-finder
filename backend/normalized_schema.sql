-- Create normalized database structure for ABA provider data
-- This matches the CSV structure with proper junction tables

-- Main tables
CREATE TABLE IF NOT EXISTS providers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(100),
    address TEXT,
    website_domain VARCHAR(255),
    latitude DECIMAL(10,8),
    longitude DECIMAL(11,8),
    google_address TEXT,
    google_website VARCHAR(255),
    google_rating DECIMAL(3,2),
    google_review_count DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS regional_centers (
    id SERIAL PRIMARY KEY,
    regional_center VARCHAR(255) NOT NULL,
    office_type VARCHAR(50),
    address VARCHAR(255),
    suite VARCHAR(50),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    telephone VARCHAR(20),
    website VARCHAR(255),
    county_served VARCHAR(100),
    los_angeles_health_district VARCHAR(100),
    location_coordinates TEXT,
    latitude DECIMAL(10,8),
    longitude DECIMAL(11,8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reference tables for many-to-many relationships
CREATE TABLE IF NOT EXISTS coverage_areas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS insurance_plans (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS specializations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Junction tables for many-to-many relationships
CREATE TABLE IF NOT EXISTS provider_coverage_areas (
    provider_id INTEGER NOT NULL,
    coverage_area_id INTEGER NOT NULL,
    PRIMARY KEY (provider_id, coverage_area_id),
    FOREIGN KEY (provider_id) REFERENCES providers(id) ON DELETE CASCADE,
    FOREIGN KEY (coverage_area_id) REFERENCES coverage_areas(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS provider_insurance_plans (
    provider_id INTEGER NOT NULL,
    insurance_plan_id INTEGER NOT NULL,
    PRIMARY KEY (provider_id, insurance_plan_id),
    FOREIGN KEY (provider_id) REFERENCES providers(id) ON DELETE CASCADE,
    FOREIGN KEY (insurance_plan_id) REFERENCES insurance_plans(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS provider_services (
    provider_id INTEGER NOT NULL,
    service_id INTEGER NOT NULL,
    PRIMARY KEY (provider_id, service_id),
    FOREIGN KEY (provider_id) REFERENCES providers(id) ON DELETE CASCADE,
    FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS provider_specializations (
    provider_id INTEGER NOT NULL,
    specialization_id INTEGER NOT NULL,
    PRIMARY KEY (provider_id, specialization_id),
    FOREIGN KEY (provider_id) REFERENCES providers(id) ON DELETE CASCADE,
    FOREIGN KEY (specialization_id) REFERENCES specializations(id) ON DELETE CASCADE
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_providers_name ON providers(name);
CREATE INDEX IF NOT EXISTS idx_providers_location ON providers(latitude, longitude);
CREATE INDEX IF NOT EXISTS idx_providers_created_at ON providers(created_at);

CREATE INDEX IF NOT EXISTS idx_regional_centers_name ON regional_centers(regional_center);
CREATE INDEX IF NOT EXISTS idx_regional_centers_location ON regional_centers(latitude, longitude);
CREATE INDEX IF NOT EXISTS idx_regional_centers_county ON regional_centers(county_served);

CREATE INDEX IF NOT EXISTS idx_coverage_areas_name ON coverage_areas(name);
CREATE INDEX IF NOT EXISTS idx_insurance_plans_name ON insurance_plans(name);
CREATE INDEX IF NOT EXISTS idx_services_name ON services(name);
CREATE INDEX IF NOT EXISTS idx_specializations_name ON specializations(name);

-- Junction table indexes
CREATE INDEX IF NOT EXISTS idx_provider_coverage_areas_provider ON provider_coverage_areas(provider_id);
CREATE INDEX IF NOT EXISTS idx_provider_coverage_areas_area ON provider_coverage_areas(coverage_area_id);

CREATE INDEX IF NOT EXISTS idx_provider_insurance_plans_provider ON provider_insurance_plans(provider_id);
CREATE INDEX IF NOT EXISTS idx_provider_insurance_plans_plan ON provider_insurance_plans(insurance_plan_id);

CREATE INDEX IF NOT EXISTS idx_provider_services_provider ON provider_services(provider_id);
CREATE INDEX IF NOT EXISTS idx_provider_services_service ON provider_services(service_id);

CREATE INDEX IF NOT EXISTS idx_provider_specializations_provider ON provider_specializations(provider_id);
CREATE INDEX IF NOT EXISTS idx_provider_specializations_spec ON provider_specializations(specialization_id);

-- Comments for documentation
COMMENT ON TABLE providers IS 'ABA therapy providers with normalized relationships';
COMMENT ON TABLE regional_centers IS 'California regional centers with location and contact information';
COMMENT ON TABLE coverage_areas IS 'Service coverage areas for providers';
COMMENT ON TABLE insurance_plans IS 'Insurance plans accepted by providers';
COMMENT ON TABLE services IS 'Services or service locations offered by providers';
COMMENT ON TABLE specializations IS 'Specializations or expertise areas of providers';
COMMENT ON TABLE provider_coverage_areas IS 'Junction table linking providers to their coverage areas';
COMMENT ON TABLE provider_insurance_plans IS 'Junction table linking providers to accepted insurance plans';
COMMENT ON TABLE provider_services IS 'Junction table linking providers to their services';
COMMENT ON TABLE provider_specializations IS 'Junction table linking providers to their specializations'; 