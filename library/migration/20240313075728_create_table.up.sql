CREATE TABLE IF NOT EXISTS public.addresses (
        id text COLLATE pg_catalog."default" NOT NULL,
        user_detail_id text COLLATE pg_catalog."default" NOT NULL,
        province text COLLATE pg_catalog."default" NOT NULL,
        regency_city text COLLATE pg_catalog."default" NOT NULL,
        sub_district text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT addresses_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.class_user (
        id text COLLATE pg_catalog."default" NOT NULL,
        name character varying(100) COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT class_user_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.class_user_plans (
        id text COLLATE pg_catalog."default" NOT NULL,
        user_id text COLLATE pg_catalog."default" NOT NULL,
        plan_id text COLLATE pg_catalog."default" NOT NULL,
        name character varying(50) COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT class_user_plan_pkey PRIMARY KEY (id),
        CONSTRAINT class_user_plan_uq UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.exercise_answers (
        id text COLLATE pg_catalog."default" NOT NULL,
        submission_id text COLLATE pg_catalog."default" NOT NULL,
        choice_id text COLLATE pg_catalog."default",
        is_marked boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT exercise_answer_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_choices (
        id text COLLATE pg_catalog."default" NOT NULL,
        question_id text COLLATE pg_catalog."default" NOT NULL,
        content text COLLATE pg_catalog."default" NOT NULL,
        point text COLLATE pg_catalog."default" NOT NULL,
        is_correct boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT exercise_choice_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_questions (
        id text COLLATE pg_catalog."default" NOT NULL,
        exercise_id text COLLATE pg_catalog."default" NOT NULL,
        content text COLLATE pg_catalog."default" NOT NULL,
        image text COLLATE pg_catalog."default",
        point integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT exercise_question_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_scores (
        id text COLLATE pg_catalog."default" NOT NULL,
        submission_id text COLLATE pg_catalog."default" NOT NULL,
        score integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT exercise_scores_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_submissions (
        id text COLLATE pg_catalog."default" NOT NULL,
        user_id text COLLATE pg_catalog."default" NOT NULL,
        exercise_id text COLLATE pg_catalog."default" NOT NULL,
        started_at timestamp with time zone NOT NULL,
        finished_at timestamp with time zone NOT NULL,
        time_required time without time zone NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT exercise_submissions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercises (
        id text COLLATE pg_catalog."default" NOT NULL,
        test_type_id text COLLATE pg_catalog."default" NOT NULL,
        title character varying(50) COLLATE pg_catalog."default" NOT NULL,
        description text COLLATE pg_catalog."default" NOT NULL,
        "time" integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT exercise_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.modules (
        id text COLLATE pg_catalog."default" NOT NULL,
        name text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT modules_pkey PRIMARY KEY (id),
        CONSTRAINT module_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.plans (
        id text COLLATE pg_catalog."default" NOT NULL,
        module_id text COLLATE pg_catalog."default" NOT NULL,
        name text COLLATE pg_catalog."default" NOT NULL,
        price numeric NOT NULL,
        grup_pejuang boolean NOT NULL,
        exercise bigint NOT NULL,
        access bigint NOT NULL,
        module boolean NOT NULL,
        try_out bigint NOT NULL,
        zoom boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT sub_plans_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_answers (
        id text COLLATE pg_catalog."default" NOT NULL,
        submission_id text COLLATE pg_catalog."default" NOT NULL,
        choice_id text COLLATE pg_catalog."default",
        is_marked boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT quiz_answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_choices (
        id text COLLATE pg_catalog."default" NOT NULL,
        question_id text COLLATE pg_catalog."default" NOT NULL,
        content text COLLATE pg_catalog."default" NOT NULL,
        point integer NOT NULL,
        is_correct boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT quiz_choices_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_questions (
        id text COLLATE pg_catalog."default" NOT NULL,
        quiz_id text COLLATE pg_catalog."default" NOT NULL,
        content text COLLATE pg_catalog."default" NOT NULL,
        image text COLLATE pg_catalog."default",
        point integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT quiz_questions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_scores (
        id text COLLATE pg_catalog."default" NOT NULL,
        submission_id text COLLATE pg_catalog."default" NOT NULL,
        score integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT quiz_scores_pket PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_submissions (
        id text COLLATE pg_catalog."default" NOT NULL,
        user_id text COLLATE pg_catalog."default" NOT NULL,
        quiz_id text COLLATE pg_catalog."default" NOT NULL,
        started_at timestamp with time zone NOT NULL,
        finished_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        time_required time without time zone NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT quiz_submissions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quizzes (
        id text COLLATE pg_catalog."default" NOT NULL,
        sub_subject_id text COLLATE pg_catalog."default" NOT NULL,
        test_type_id text COLLATE pg_catalog."default" NOT NULL,
        open timestamp with time zone NOT NULL,
        title character varying(50) COLLATE pg_catalog."default" NOT NULL,
        description text COLLATE pg_catalog."default" NOT NULL,
        "time" integer NOT NULL,
        point integer NOT NULL,
        attempt integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT quiz_pkey PRIMARY KEY (id),
        CONSTRAINT quiz_uq UNIQUE (title)
);

CREATE TABLE IF NOT EXISTS public.roles (
        id text COLLATE pg_catalog."default" NOT NULL,
        name character varying(20) COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT roles_pkey PRIMARY KEY (id),
        CONSTRAINT role_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.schedules (
        id text COLLATE pg_catalog."default" NOT NULL,
        sub_subject_id text COLLATE pg_catalog."default" NOT NULL,
        class_user_plan_id text COLLATE pg_catalog."default" NOT NULL,
        meeting_date timestamp with time zone NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT schedule_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.sub_modules (
        id text COLLATE pg_catalog."default" NOT NULL,
        module_id text COLLATE pg_catalog."default" NOT NULL,
        name text COLLATE pg_catalog."default" NOT NULL,
        description text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT sub_modules_pkey PRIMARY KEY (id),
        CONSTRAINT sub_module_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.sub_subjects (
        id text COLLATE pg_catalog."default" NOT NULL,
        subject_id text COLLATE pg_catalog."default" NOT NULL,
        name text COLLATE pg_catalog."default" NOT NULL,
        content text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT sub_subjects_pkey PRIMARY KEY (id),
        CONSTRAINT sub_subject_uq UNIQUE (name, content)
);

CREATE TABLE IF NOT EXISTS public.subjects (
        id text COLLATE pg_catalog."default" NOT NULL,
        sub_module_id text COLLATE pg_catalog."default" NOT NULL,
        name text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT subjects_pkey PRIMARY KEY (id),
        CONSTRAINT subject_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.test_types (
        id text COLLATE pg_catalog."default" NOT NULL,
        name text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT test_type_pkey PRIMARY KEY (id),
        CONSTRAINT test_type_uq UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.testimonials (
        id text COLLATE pg_catalog."default" NOT NULL,
        comment text COLLATE pg_catalog."default" NOT NULL,
        rating numeric NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT testimonials_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.user_details (
        id text COLLATE pg_catalog."default" NOT NULL,
        class_user_id text COLLATE pg_catalog."default" NOT NULL,
        user_id text COLLATE pg_catalog."default" NOT NULL,
        name character varying(100) COLLATE pg_catalog."default" NOT NULL,
        province character varying(100) NOT NULL,
        regency character varying(255) NOT NULL,
        district character varying(255) NOT NULL,
        phone_number character varying(50) COLLATE pg_catalog."default" NOT NULL,
        profile_picture text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT user_details_pkey PRIMARY KEY (id),
        CONSTRAINT user_details_uq UNIQUE (phone_number, profile_picture)
);

CREATE TABLE IF NOT EXISTS public.user_testimonials (
        id text COLLATE pg_catalog."default" NOT NULL,
        user_id text COLLATE pg_catalog."default" NOT NULL,
        testimonial_id text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT user_testimonials_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.users (
        id text COLLATE pg_catalog."default" NOT NULL,
        role_id text COLLATE pg_catalog."default" NOT NULL,
        google_id text COLLATE pg_catalog."default",
        facebook_id text COLLATE pg_catalog."default",
        email character varying(64) COLLATE pg_catalog."default" NOT NULL,
        password text COLLATE pg_catalog."default" NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT users_pkey PRIMARY KEY (id),
        CONSTRAINT users_email_key UNIQUE (email)
);