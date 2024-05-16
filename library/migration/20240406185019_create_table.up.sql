CREATE TABLE IF NOT EXISTS public.class_user_plans (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        plan_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT class_user_plans_pkey PRIMARY KEY (id),
        CONSTRAINT class_user_plan_uq UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.class_users (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT class_users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_answers (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        submission_id integer NOT NULL,
        choice_id integer,
        is_marked boolean NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT exercise_answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_choices (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        question_id integer NOT NULL,
        content text NOT NULL,
        score integer NOT NULL,
        is_correct boolean NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT exercise_choices_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_questions (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        exercise_id integer NOT NULL,
        content text NOT NULL,
        score integer NOT NULL,
        explanation text NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT exercise_questions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_submissions (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        exercise_id integer NOT NULL,
        started_at timestamp with time zone NOT NULL,
        finished_at timestamp with time zone NOT NULL,
        time_required time without time zone NOT NULL,
        score integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT exercise_submissions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercises (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        test_type_id integer NOT NULL,
        module_id integer NOT NULL,
        title character varying(50) NOT NULL,
        total_question integer NOT NULL,
        description text NOT NULL,
        "time" integer NOT NULL,
        attempt integer NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT exercises_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.mentors (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        module_id integer NOT NULL,
        short_name character varying(100) NOT NULL,
        description text NOT NULL,
        motto character varying(255) NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT mentors_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.modules (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT modules_pkey PRIMARY KEY (id),
        CONSTRAINT module_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.plans (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        media_id integer NOT NULL,
        module_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        price numeric NOT NULL,
        "group" boolean NOT NULL,
        exercise bigint NOT NULL,
        access bigint NOT NULL,
        module boolean NOT NULL,
        try_out bigint NOT NULL,
        zoom boolean NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT plans_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_answers (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        submission_id integer NOT NULL,
        choice_id integer,
        is_marked boolean NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT quiz_answers_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_choices (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        question_id integer NOT NULL,
        content text NOT NULL,
        score integer NOT NULL,
        is_correct boolean NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT quiz_choices_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_questions (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        quiz_id integer NOT NULL,
        content text NOT NULL,
        score integer NOT NULL,
        explanation text NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT quiz_questions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_submissions (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        quiz_id integer NOT NULL,
        started_at timestamp with time zone NOT NULL,
        finished_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        time_required time without time zone NOT NULL,
        score integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT quiz_submissions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quizzes (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        subject_id integer NOT NULL,
        test_type_id integer NOT NULL,
        open timestamp with time zone NOT NULL,
        close timestamp with time zone NOT NULL,
        title character varying(100) NOT NULL,
        description text NOT NULL,
        "time" integer NOT NULL,
        max_score integer NOT NULL,
        attempt integer NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT quizzes_pkey PRIMARY KEY (id),
        CONSTRAINT quiz_uq UNIQUE (title)
);

CREATE TABLE IF NOT EXISTS public.roles (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT roles_pkey PRIMARY KEY (id),
        CONSTRAINT role_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.schedules (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        sub_subject_id integer NOT NULL,
        class_user_plan_id integer NOT NULL,
        meeting_date timestamp with time zone NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT schedules_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.sub_modules (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        module_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        description text NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT sub_modules_pkey PRIMARY KEY (id),
        CONSTRAINT sub_module_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.sub_subjects (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        subject_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        content text NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT sub_subjects_pkey PRIMARY KEY (id),
        CONSTRAINT sub_subject_uq UNIQUE (name, content)
);

CREATE TABLE IF NOT EXISTS public.subjects (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        sub_module_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT subjects_pkey PRIMARY KEY (id),
        CONSTRAINT subject_uq UNIQUE (name) INCLUDE(name)
);

CREATE TABLE IF NOT EXISTS public.test_types (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT test_types_pkey PRIMARY KEY (id),
        CONSTRAINT test_type_uq UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.uniques (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        mentor_id integer NOT NULL,
        uniqueness text NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT uniques_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.user_details (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        class_user_id integer NOT NULL,
        user_id integer NOT NULL,
        media_id integer,
        name character varying(1000) NOT NULL,
        province character varying(100) NOT NULL,
        regency character varying(255) NOT NULL,
        district character varying(255) NOT NULL,
        phone_number character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT user_details_pkey PRIMARY KEY (id),
        CONSTRAINT user_details_uq UNIQUE (phone_number)
);

CREATE TABLE IF NOT EXISTS public.user_mentor_testimonials (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        mentor_id integer NOT NULL,
        comment text NOT NULL,
        rating numeric NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT user_mentor_testimonials_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.user_testimonials (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        comment text NOT NULL,
        rating numeric NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT user_testimonials_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.users (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        role_id integer NOT NULL,
        google_id integer,
        facebook_id integer,
        email character varying(64) NOT NULL,
        password text NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT users_pkey PRIMARY KEY (id),
        CONSTRAINT users_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.medias (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        filename character varying(255) NOT NULL,
        mime character varying(50) NOT NULL,
        original_filename character varying(255) NOT NULL,
        description text NOT NULL,
        created_by integer NOT NULL,
        updated_by integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.quiz_question_media (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        quiz_question_id integer NOT NULL,
        media_id integer NOT NULL,
        index integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.exercise_question_media (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        exercise_question_id integer NOT NULL,
        media_id integer NOT NULL,
        index integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.sub_subject_media (
        id serial NOT NULL,
        uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
        sub_subject_id integer NOT NULL,
        media_id integer NOT NULL,
        index integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        PRIMARY KEY (id)
);