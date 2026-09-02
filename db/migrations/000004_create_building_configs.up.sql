CREATE TABLE building_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    build_time_seconds INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    health_points INT NOT NULL,
    dunbroch_level_required INT DEFAULT 1,
    width INT NOT NULL,
    height INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_building_configs_name ON building_configs(name);
