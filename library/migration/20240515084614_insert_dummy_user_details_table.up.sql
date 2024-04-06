INSERT INTO
        public.user_details (
                class_user_id,
                user_id,
                name,
                province,
                regency,
                district,
                phone_number,
                profile_picture
        )
VALUES
        -- admin
        (
                4,
                1,
                'Administrator',
                'Jawa Tengah',
                'Wonogiri',
                'Selogiri',
                '081234567890',
                './assets/img/users/profiles/account.png'
        ),
        -- mentor
        (
                4,
                2,
                'Mentor',
                'Jawa Tengah',
                'Sukoharjo',
                'Carikan',
                '0853215465875',
                './assets/img/users/profiles/account.png'
        ),
        -- user
        (
                2,
                3,
                'User',
                'Jawa Tengah',
                'Sukoharjo',
                'Bacem',
                '0853265478956',
                './assets/img/users/profiles/account.png'
        );