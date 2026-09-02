CREATE TABLE producer_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    resource_type VARCHAR(50) NOT NULL,
    base_production_rate INT NOT NULL,
    level_multiplier DECIMAL(5,2) NOT NULL,
    max_storage INT NOT NULL,
    build_time_seconds INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_producer_configs_resource_type ON producer_configs(resource_type);
