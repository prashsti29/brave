CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    dunbroch_level INT DEFAULT 1,
    gems BIGINT DEFAULT 10,
    wisps BIGINT DEFAULT 500,
    embis BIGINT DEFAULT 500,
    total_attacks INT DEFAULT 0,
    attacks_won INT DEFAULT 0,
    total_defenses INT DEFAULT 0,
    defenses_won INT DEFAULT 0
);
CREATE INDEX idx_players_email ON players(email);