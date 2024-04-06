INSERT INTO
        public.users (uuid, role_id, email, password)
VALUES
        -- admin
        (
                '668c1a70-d076-4fec-8852-948c40628933',
                1,
                'admin@gmail.com',
                '$2a$10$QlMk55EzxhCJK/1mm45fOePw8Zn35Mw/oafnrWHrZeg8n0V/v5wi.'
        ),
        -- mentor
        (
                '1824da37-f0c5-4a48-ac16-04c24ccca848',
                3,
                'mentor@gmail.com',
                '$2a$10$QlMk55EzxhCJK/1mm45fOePw8Zn35Mw/oafnrWHrZeg8n0V/v5wi.'
        ),
        -- user
        (
                'b8e7cb7f-3583-4c64-8e0a-3b30aa2f941b',
                2,
                'user@gmail.com',
                '$2a$10$QlMk55EzxhCJK/1mm45fOePw8Zn35Mw/oafnrWHrZeg8n0V/v5wi.'
        );