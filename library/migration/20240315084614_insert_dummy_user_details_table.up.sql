INSERT INTO
        public.user_details (
                uuid,
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
                '8a7be0e9-905d-4639-9fb7-cc889930f778',
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
                'caee71d2-5e90-4dcc-a050-4e99fcd71995',
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
                '39c586d8-fc66-45c4-a744-1d4ace46611a',
                2,
                3,
                'User',
                'Jawa Tengah',
                'Sukoharjo',
                'Bacem',
                '0853265478956',
                './assets/img/users/profiles/account.png'
        );