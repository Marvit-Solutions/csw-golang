INSERT INTO
        public.plans (module_id, name, slug, price, "group", exercise, access, module, try_out, zoom)
VALUES
        -- SKD
        (1, 'Paket Bronze SKD', 'skd-bronze', 50000, true, 660, 2, false, 0, false),
        (1, 'Paket Silver SKD', 'skd-silver', 150000, true, 660, 3, true, 2, false),
        (1, 'Paket Gold SKD', 'skd-gold', 250000, true, 660, 6, true, 4, true),
        (1, 'Paket Platinum SKD', 'skd-platinum', 300000, true, 1100, 6, true, 4, true),
        -- Matematika
        (2, 'Paket Bronze Matematika', 'matematika-bronze', 50000, true, 240, 2, false, 0, false),
        (2, 'Paket Silver Matematika', 'matematika-silver', 150000, true, 240, 3, true, 2, false),
        (2, 'Paket Gold Matematika', 'matematika-gold', 200000, true, 240, 6, true, 3, true),
        (2, 'Paket Platinum Matematika', 'matematika-platinum', 250000, true, 400, 6, true, 4, true);