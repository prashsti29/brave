CREATE TABLE troop_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    unit_type VARCHAR(50) NOT NULL,
    training_time_seconds INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    health INT NOT NULL,
    damage INT NOT NULL,
    range INT NOT NULL,
    space_required INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_troop_configs_unit_type ON troop_configs(unit_type);
CREATE INDEX idx_troop_configs_name ON troop_configs(name);
