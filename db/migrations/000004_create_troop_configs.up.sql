CREATE TABLE troop_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    training_time INT NOT NULL,
    damage_per_shot INT NOT NULL,
    health INT NOT NULL,
    unlocks_at_dunbroch_level INT DEFAULT 1,
    level INT DEFAULT 1,
    housing_space INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_troop_configs_name ON troop_configs(name);
