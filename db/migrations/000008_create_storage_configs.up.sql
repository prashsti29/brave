CREATE TABLE storage_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    resource_type VARCHAR(50) NOT NULL,
    max_capacity INT NOT NULL,
    build_time_seconds INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    health INT NOT NULL,
    dunbroch_level_required INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_storage_configs_resource_type ON storage_configs(resource_type);
