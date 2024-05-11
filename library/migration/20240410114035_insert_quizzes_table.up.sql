INSERT INTO
        public.quizzes (
                subject_id,
                test_type_id,
                open,
                close,
                title,
                description,
                time,
                max_score,
                attempt
        )
VALUES
        -- Pancasila
        (
                1,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Pancasila',
                'Deskripsi latihan soal modul pancasila',
                10,
                0,
                0
        ),
        (
                1,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Pancasila',
                'Deskripsi pretest modul pancasila',
                10,
                0,
                1
        ),
        (
                1,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Pancasila',
                'Deskripsi posttest modul pancasila',
                10,
                0,
                1
        ),
        -- Sejarah Indonesia
        (
                2,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Sejarah Indonesia',
                'Deskripsi latihan soal modul Sejarah Indonesia',
                10,
                0,
                0
        ),
        (
                2,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Sejarah Indonesia',
                'Deskripsi pretest modul Sejarah Indonesia',
                10,
                0,
                1
        ),
        (
                2,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Sejarah Indonesia',
                'Deskripsi posttest modul Sejarah Indonesia',
                10,
                0,
                1
        ),
        -- Kebudayaan dan Keberagaman Indonesia
        (
                3,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Kebudayaan dan Keberagaman Indonesia',
                'Deskripsi latihan soal modul Kebudayaan dan Keberagaman Indonesia',
                10,
                0,
                0
        ),
        (
                3,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Kebudayaan dan Keberagaman Indonesia',
                'Deskripsi pretest modul Kebudayaan dan Keberagaman Indonesia',
                10,
                0,
                1
        ),
        (
                3,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Kebudayaan dan Keberagaman Indonesia',
                'Deskripsi posttest modul Kebudayaan dan Keberagaman Indonesia',
                10,
                0,
                1
        ),
        -- Kemampuan Verbal
        (
                4,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Kemampuan Verbal',
                'Deskripsi latihan soal modul Kemampuan Verbal',
                10,
                0,
                0
        ),
        (
                4,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Kemampuan Verbal',
                'Deskripsi pretest modul Kemampuan Verbal',
                10,
                0,
                1
        ),
        (
                4,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Kemampuan Verbal',
                'Deskripsi posttest modul Kemampuan Verbal',
                10,
                0,
                1
        ),
        -- Kemampuan Numerik
        (
                5,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Kemampuan Numerik',
                'Deskripsi latihan soal modul Kemampuan Numerik',
                10,
                0,
                0
        ),
        (
                5,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Kemampuan Numerik',
                'Deskripsi pretest modul Kemampuan Numerik',
                10,
                0,
                1
        ),
        (
                5,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Kemampuan Numerik',
                'Deskripsi posttest modul Kemampuan Numerik',
                10,
                0,
                1
        ),
        -- Kemampuan Logis
        (
                6,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Kemampuan Logis',
                'Deskripsi latihan soal modul Kemampuan Logis',
                10,
                0,
                0
        ),
        (
                6,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Kemampuan Logis',
                'Deskripsi pretest modul Kemampuan Logis',
                10,
                0,
                1
        ),
        (
                6,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Kemampuan Logis',
                'Deskripsi posttest modul Kemampuan Logis',
                10,
                0,
                1
        ),
        -- Etika dan Moralitas
        (
                7,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Etika dan Moralitas',
                'Deskripsi latihan soal modul Etika dan Moralitas',
                10,
                0,
                0
        ),
        (
                7,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Etika dan Moralitas',
                'Deskripsi pretest modul Etika dan Moralitas',
                10,
                0,
                1
        ),
        (
                7,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Etika dan Moralitas',
                'Deskripsi posttest modul Etika dan Moralitas',
                10,
                0,
                1
        ),
        -- Kerjasama Tim
        (
                8,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Kerjasama Tim',
                'Deskripsi latihan soal modul Kerjasama Tim',
                10,
                0,
                0
        ),
        (
                8,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Kerjasama Tim',
                'Deskripsi pretest modul Kerjasama Tim',
                10,
                0,
                1
        ),
        (
                8,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Kerjasama Tim',
                'Deskripsi posttest modul Kerjasama Tim',
                10,
                0,
                1
        ),
        -- Adaptabilitas
        (
                9,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Adaptabilitas',
                'Deskripsi latihan soal modul Adaptabilitas',
                10,
                0,
                0
        ),
        (
                9,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Adaptabilitas',
                'Deskripsi pretest modul Adaptabilitas',
                10,
                0,
                1
        ),
        (
                9,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Adaptabilitas',
                'Deskripsi posttest modul Adaptabilitas',
                10,
                0,
                1
        ),
        -- Aljabar
        (
                10,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Aljabar',
                'Deskripsi latihan soal modul Aljabar',
                10,
                0,
                0
        ),
        (
                10,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Aljabar',
                'Deskripsi pretest modul Aljabar',
                10,
                0,
                1
        ),
        (
                10,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Aljabar',
                'Deskripsi posttest modul Aljabar',
                10,
                0,
                1
        ),
        -- Geometri
        (
                11,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Geometri',
                'Deskripsi latihan soal modul Geometri',
                10,
                0,
                0
        ),
        (
                11,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Geometri',
                'Deskripsi pretest modul Geometri',
                10,
                0,
                1
        ),
        (
                11,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Geometri',
                'Deskripsi posttest modul Geometri',
                10,
                0,
                1
        ),
        -- Statistika
        (
                12,
                1,
                NULL,
                NULL,
                'Latihan Soal Modul Statistika',
                'Deskripsi latihan soal modul Statistika',
                10,
                0,
                0
        ),
        (
                12,
                2,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Pretest Modul Statistika',
                'Deskripsi pretest modul Statistika',
                10,
                0,
                1
        ),
        (
                12,
                3,
                CURRENT_TIMESTAMP,
                CURRENT_TIMESTAMP + INTERVAL '1 day',
                'Posttest Modul Statistika',
                'Deskripsi posttest modul Statistika',
                10,
                0,
                1
        );