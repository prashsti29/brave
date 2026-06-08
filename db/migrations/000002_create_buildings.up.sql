
CREATE TABLE buildings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    level INT DEFAULT 1,
    upgrade_price INT DEFAULT 0,
    upgrade_time INT DEFAULT 0,
    currency VARCHAR(20),
    is_upgrading BOOLEAN DEFAULT FALSE,
    upgrade_end_time TIMESTAMP,
    dunbroch_level INT DEFAULT 1,
    max_allowed INT DEFAULT 1,
    max_health INT NOT NULL,
    current_health INT NOT NULL,
   created_at TIMESTAMP DEFAULT NOW()
);