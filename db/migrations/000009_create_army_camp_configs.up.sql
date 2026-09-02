CREATE TABLE army_camp_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    troop_slot_capacity INT NOT NULL,
    build_time_seconds INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    health INT NOT NULL,
    dunbroch_level_required INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_army_camp_configs_name ON army_camp_configs(name);
