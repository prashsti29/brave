CREATE TABLE village_layout(
    player_id UUID,
    building_id UUID PRIMARY KEY,
    x INT NOT NULL,
    y INT NOT NULL,
    updated_at TIMESTAMPZONE DEFAULT NOW(),
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE,
    FOREIGN KEY (building_id) REFERENCES buildings(id) ON DELETE CASCADE
);