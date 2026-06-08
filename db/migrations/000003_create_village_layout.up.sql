CREATE TABLE village_layout(
    player_id UUID,
    building_id UUID PRIMARY KEY,
    x INT NOT NULL,
    y INT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE,
    FOREIGN KEY (building_id) REFERENCES buildings(id) ON DELETE CASCADE
);
CREATE INDEX idx_village_layout_player_id ON village_layout(player_id);
CREATE INDEX idx_village_layout_building_id ON village_layout(building_id);