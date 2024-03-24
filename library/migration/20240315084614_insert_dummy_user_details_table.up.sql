INSERT INTO
        public.user_details (
                id,
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
                '2461a9fd-b690-469f-b466-fa65dbab39ea',
                '668c1a70-d076-4fec-8852-948c40628933',
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
                '2461a9fd-b690-469f-b466-fa65dbab39ea',
                '1824da37-f0c5-4a48-ac16-04c24ccca848',
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
                '2b9be02c-1e0d-4430-913b-bbe0bf664c9a',
                'b8e7cb7f-3583-4c64-8e0a-3b30aa2f941b',
                'User',
                'Jawa Tengah',
                'Sukoharjo',
                'Bacem',
                '0853265478956',
                './assets/img/users/profiles/account.png'
        );