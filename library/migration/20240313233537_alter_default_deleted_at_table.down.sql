ALTER TABLE
        public.addresses
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.class_users
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.class_user_plans
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.exercise_answers
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.exercise_choices
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.exercise_questions
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.exercise_scores
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.exercise_submissions
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.exercises
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.modules
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.plans
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.quiz_answers
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.quiz_choices
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.quiz_questions
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.quiz_scores
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.quiz_submissions
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.quizzes
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.roles
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.schedules
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.sub_modules
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.sub_subjects
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.subjects
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.test_types
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.testimonials
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.user_details
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.user_testimonials
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE
        public.users
ALTER COLUMN
        deleted_at
SET
        DEFAULT CURRENT_TIMESTAMP;