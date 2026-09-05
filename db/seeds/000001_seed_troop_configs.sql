-- Seeds for troop_configs
INSERT INTO troop_configs (name, training_time, damage_per_shot, health, unlocks_at_dunbroch_level, level, housing_space, cost_wisps, cost_embis) VALUES
('Archer', 30, 15, 50, 1, 1, 1, 10, 0),
('Archer Level 2', 35, 18, 55, 3, 2, 1, 12, 0),
('Goblin', 20, 10, 40, 1, 1, 1, 5, 0),
('Goblin Level 2', 25, 12, 45, 3, 2, 1, 7, 0),
('Golem', 60, 30, 200, 5, 1, 8, 50, 0),
('Golem Level 2', 70, 35, 250, 7, 2, 8, 60, 0),
('Dragon', 120, 60, 300, 7, 1, 10, 200, 100),
('Dragon Level 2', 130, 65, 320, 9, 2, 10, 220, 110);
