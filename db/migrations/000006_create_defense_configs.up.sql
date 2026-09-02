CREATE TABLE defense_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    build_time_seconds INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    health INT NOT NULL,
    damage INT NOT NULL,
    range INT NOT NULL,
    fire_rate DECIMAL(5,2) NOT NULL,
    dunbroch_level_required INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_defense_configs_name ON defense_configs(name);
