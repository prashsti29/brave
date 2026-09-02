CREATE TABLE laboratory_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    research_type VARCHAR(50) NOT NULL,
    research_time_seconds INT NOT NULL,
    cost_wisps INT NOT NULL,
    cost_embis INT NOT NULL,
    boost_percentage INT NOT NULL,
    dunbroch_level_required INT DEFAULT 1,
    max_level INT DEFAULT 30,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_laboratory_configs_research_type ON laboratory_configs(research_type);
