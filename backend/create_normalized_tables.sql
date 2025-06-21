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

-- Legacy tables for backward compatibility (keep existing structure)
CREATE TABLE IF NOT EXISTS aba_centers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address TEXT,
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    phone VARCHAR(20),
    email VARCHAR(255),
    website VARCHAR(255),
    latitude DECIMAL(10,8),
    longitude DECIMAL(11,8),
    service_type VARCHAR(100),
    insurance_accepted TEXT,
    waitlist_availability VARCHAR(50),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS resource_centers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address TEXT,
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(10),
    phone VARCHAR(20),
    email VARCHAR(255),
    website VARCHAR(255),
    latitude DECIMAL(10,8),
    longitude DECIMAL(11,8),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS diagnoses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(20) UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS resource_center_diagnoses (
    resource_center_id UUID NOT NULL,
    diagnosis_id UUID NOT NULL,
    PRIMARY KEY (resource_center_id, diagnosis_id),
    FOREIGN KEY (resource_center_id) REFERENCES resource_centers(id) ON DELETE CASCADE,
    FOREIGN KEY (diagnosis_id) REFERENCES diagnoses(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS resources (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    url VARCHAR(500),
    resource_type VARCHAR(100),
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS form_submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    message TEXT NOT NULL,
    facility_id VARCHAR(255),
    facility_type VARCHAR(100),
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) NOT NULL UNIQUE,
    preferences JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes for legacy tables
CREATE INDEX IF NOT EXISTS idx_aba_centers_name ON aba_centers(name) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_aba_centers_location ON aba_centers(latitude, longitude) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_resource_centers_name ON resource_centers(name) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_resource_centers_location ON resource_centers(latitude, longitude) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_diagnoses_name ON diagnoses(name) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_diagnoses_code ON diagnoses(code) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_resources_title ON resources(title) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_resources_type ON resources(resource_type) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_form_submissions_email ON form_submissions(email) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_form_submissions_status ON form_submissions(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_form_submissions_created_at ON form_submissions(created_at) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_user_preferences_user_id ON user_preferences(user_id) WHERE deleted_at IS NULL;

-- Add trigger to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply trigger to all tables with updated_at column
CREATE TRIGGER update_providers_updated_at BEFORE UPDATE ON providers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_regional_centers_updated_at BEFORE UPDATE ON regional_centers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_coverage_areas_updated_at BEFORE UPDATE ON coverage_areas FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_insurance_plans_updated_at BEFORE UPDATE ON insurance_plans FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_services_updated_at BEFORE UPDATE ON services FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_specializations_updated_at BEFORE UPDATE ON specializations FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_aba_centers_updated_at BEFORE UPDATE ON aba_centers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_resource_centers_updated_at BEFORE UPDATE ON resource_centers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_diagnoses_updated_at BEFORE UPDATE ON diagnoses FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_resources_updated_at BEFORE UPDATE ON resources FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_form_submissions_updated_at BEFORE UPDATE ON form_submissions FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_user_preferences_updated_at BEFORE UPDATE ON user_preferences FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
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