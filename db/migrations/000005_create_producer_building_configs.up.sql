CREATE TABLE producer_building_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    resource_type VARCHAR(50) NOT NULL,
    production_rate INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_producer_building_configs_type ON producer_building_configs(resource_type);
