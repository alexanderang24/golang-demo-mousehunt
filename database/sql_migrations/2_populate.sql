-- +migrate Up
-- +migrate StatementBegin

INSERT INTO trap(name, min_power, max_power, price, description, created_at, updated_at)
VALUES  ('Tacky Glue Trap', 100, 300, 0, 'If you believe the simplest answer is the best, then white glue and some cheese is the way to go!', now(), now()),
        ('High Tension Spring', 250, 500, 10000, 'Razor sharp barbs adorn a crushing bar held with over 100 kg of tension.', now(), now()),
        ('Spike Crusher', 550, 1000, 25000, 'A 500 pound weight with sharp spikes on the bottom precariously hangs by a thin rope, eagerly awaiting a mouse to wander below.', now(), now());

INSERT INTO location(name, description, travel_cost, created_at, updated_at)
VALUES  ('Backyard', 'Your backyard as a starting location.', 0, now(), now()),
        ('Urban Area', 'The neighborhood around your housing complex.', 5000, now(), now()),
        ('The Swamp', 'The terrain that surrounds the swamp is extremely unsteady, comprised of jagged rocks and slick moss.', 10000, now(), now());

INSERT INTO mouse(name, min_power, max_power, gold, location_id, description, created_at, updated_at)
VALUES  ('Grey Mouse', 50, 150, 500, 1, 'A very common mouse. They pose little challenge and reward.', now(), now()),
        ('Brown Mouse', 75, 200, 750, 1, 'Brown mice are the most challenging common mice.', now(), now()),
        ('Dwarf Mouse', 125, 275, 1000, 1, 'The dwarf mouse moves quickly, making it difficult to catch for new hunters.', now(), now()),
        ('Longtail Mouse', 150, 350, 2250, 2, 'Their tails grew to accommodate their lofty goals.', now(), now()),
        ('Flying Mouse', 275, 500, 2750, 2, ' Though the real birds have sought out more peaceful habitats, these winged pests don''t migrate often.', now(), now()),
        ('Gold Mouse', 325, 650, 5000, 2, 'This solid gold mouse fetches a handsome reward.', now(), now()),
        ('Speedy Mouse', 400, 750, 3000, 3, 'To the naked eye, the Speedy Mouse is less of a mouse, and more of a brown blurry comet with a pink tail.', now(), now()),
        ('Steel Mouse', 475, 850, 3500, 3, 'As the name suggests, the skin of this mouse is strong as steel.', now(), now()),
        ('Zombie Mouse', 700, 975, 250, 3, 'All mice have a weakness, but how do you kill something that''s already dead...?', now(), now());

INSERT INTO "user"(username, password, role, gold, location_id, trap_id, created_at, updated_at)
VALUES  ('admin', 'admin', 'admin', 0, 1, 1, now(), now()),
        ('john', 'admin', 'player', 500, 1, 1, now(), now());

-- +migrate StatementEnd
