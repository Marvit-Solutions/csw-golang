CREATE TABLE IF NOT EXISTS public.class_user_plans (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        plan_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT class_user_plan_uq UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.class_users (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.exercise_answers (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        submission_id integer NOT NULL,
        choice_id integer,
        is_marked boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.exercise_choices (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        question_id integer NOT NULL,
        content text NOT NULL,
        point text NOT NULL,
        is_correct boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.exercise_questions (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        exercise_id integer NOT NULL,
        content text NOT NULL,
        image text,
        point integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.exercise_scores (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        submission_id integer NOT NULL,
        score integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.exercise_submissions (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        exercise_id integer NOT NULL,
        started_at timestamp with time zone NOT NULL,
        finished_at timestamp with time zone NOT NULL,
        time_required time without time zone NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.exercises (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        test_type_id integer NOT NULL,
        title character varying(50) NOT NULL,
        description text NOT NULL,
        "time" integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.mentors (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        module_id integer NOT NULL,
        short_name character varying(100) NOT NULL,
        description text NOT NULL,
        motto character varying(255) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.modules (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT module_uq UNIQUE (name) INCLUDE (name)
);

CREATE TABLE IF NOT EXISTS public.plans (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
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
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.quiz_answers (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        submission_id integer NOT NULL,
        choice_id integer,
        is_marked boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.quiz_choices (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        question_id integer NOT NULL,
        content text NOT NULL,
        point integer NOT NULL,
        is_correct boolean NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.quiz_questions (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        quiz_id integer NOT NULL,
        content text NOT NULL,
        image text,
        point integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.quiz_scores (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        submission_id integer NOT NULL,
        score integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.quiz_submissions (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        quiz_id integer NOT NULL,
        started_at timestamp with time zone NOT NULL,
        finished_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        time_required time without time zone NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.quizzes (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        sub_subject_id integer NOT NULL,
        test_type_id integer NOT NULL,
        open timestamp with time zone NOT NULL,
        title character varying(50) NOT NULL,
        description text NOT NULL,
        "time" integer NOT NULL,
        point integer NOT NULL,
        attempt integer NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT quiz_uq UNIQUE (title)
);

CREATE TABLE IF NOT EXISTS public.roles (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT role_uq UNIQUE (name) INCLUDE (name)
);

CREATE TABLE IF NOT EXISTS public.schedules (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        sub_subject_id integer NOT NULL,
        class_user_plan_id integer NOT NULL,
        meeting_date timestamp with time zone NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.schema_migrations (
        version bigint NOT NULL,
        dirty boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS public.sub_modules (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        module_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        description text NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT sub_module_uq UNIQUE (name) INCLUDE (name)
);

CREATE TABLE IF NOT EXISTS public.sub_subjects (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        subject_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        content text NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT sub_subject_uq UNIQUE (name, content)
);

CREATE TABLE IF NOT EXISTS public.subjects (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        sub_module_id integer NOT NULL,
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT subject_uq UNIQUE (name) INCLUDE (name)
);

CREATE TABLE IF NOT EXISTS public.test_types (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        name character varying(100) NOT NULL,
        slug character varying(50) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT test_type_uq UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS public.uniques (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        mentor_id integer NOT NULL,
        name character varying(1000) NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.user_details (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        class_user_id integer NOT NULL,
        user_id integer NOT NULL,
        name character varying(1000) NOT NULL,
        province character varying(100) NOT NULL,
        regency character varying(255) NOT NULL,
        district character varying(255) NOT NULL,
        phone_number character varying(50) NOT NULL,
        profile_picture text NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT user_details_uq UNIQUE (phone_number, profile_picture)
);

CREATE TABLE IF NOT EXISTS public.user_mentor_testimonials (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        mentor_id integer NOT NULL,
        comment text NOT NULL,
        rating numeric NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.user_testimonials (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        user_id integer NOT NULL,
        comment text NOT NULL,
        rating numeric NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS public.users (
        id SERIAL PRIMARY KEY,
        uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
        role_id integer NOT NULL,
        google_id integer,
        facebook_id integer,
        email character varying(64) NOT NULL,
        password text NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at timestamp with time zone,
        CONSTRAINT users_email_key UNIQUE (email)
);