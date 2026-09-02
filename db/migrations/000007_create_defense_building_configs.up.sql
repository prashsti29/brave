CREATE TABLE defense_building_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    damage_per_shot INT NOT NULL,
    attack_speed DECIMAL(5,2) NOT NULL,
    target_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_defense_building_configs_name ON defense_building_configs(name);
