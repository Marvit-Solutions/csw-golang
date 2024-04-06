INSERT INTO
        public.users (role_id, email, password)
VALUES
        -- admin
        (
                1,
                'admin@gmail.com',
                '$2a$10$QlMk55EzxhCJK/1mm45fOePw8Zn35Mw/oafnrWHrZeg8n0V/v5wi.'
        ),
        -- mentor
        (
                3,
                'mentor@gmail.com',
                '$2a$10$QlMk55EzxhCJK/1mm45fOePw8Zn35Mw/oafnrWHrZeg8n0V/v5wi.'
        ),
        -- user
        (
                2,
                'user@gmail.com',
                '$2a$10$QlMk55EzxhCJK/1mm45fOePw8Zn35Mw/oafnrWHrZeg8n0V/v5wi.'
        );