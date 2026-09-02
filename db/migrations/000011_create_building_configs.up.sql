CREATE TABLE building_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    level INT NOT NULL,
    upgrade_price INT NOT NULL,
    upgrade_time INT NOT NULL,
    currency VARCHAR(50) NOT NULL,
    is_upgrading BOOLEAN NOT NULL,
    upgrade_end_time TIMESTAMP NOT NULL,
    dunbroch_level INT NOT NULL,
    max_allowed INT NOT NULL,
    max_health INT NOT NULL,
    UNIQUE(name, level)
);
CREATE INDEX idx_building_configs_name ON building_configs(name);