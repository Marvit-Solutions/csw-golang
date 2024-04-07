ALTER TABLE
        IF EXISTS public.uniques 
        DROP CONSTRAINT IF EXISTS mentor_unique_fk;

ALTER TABLE
        IF EXISTS public.mentors 
        DROP CONSTRAINT IF EXISTS user_mentor_fk;

ALTER TABLE
        IF EXISTS public.users 
        DROP CONSTRAINT IF EXISTS role_user_fk;

ALTER TABLE
        IF EXISTS public.user_testimonials 
        DROP CONSTRAINT IF EXISTS user_user_testimonial_fk;

ALTER TABLE
        IF EXISTS public.user_details 
        DROP CONSTRAINT IF EXISTS user_user_detail_fk,
        DROP CONSTRAINT IF EXISTS class_user_user_detail_fk;

ALTER TABLE
        IF EXISTS public.subjects 
        DROP CONSTRAINT IF EXISTS sub_module_subject_fk;

ALTER TABLE
        IF EXISTS public.sub_modules 
        DROP CONSTRAINT IF EXISTS module_sub_module_fk;

ALTER TABLE
        IF EXISTS public.schedules 
        DROP CONSTRAINT IF EXISTS sub_subject_schedule_fk,
        DROP CONSTRAINT IF EXISTS class_user_plan_schedule_fk;

ALTER TABLE
        IF EXISTS public.quizzes 
        DROP CONSTRAINT IF EXISTS test_type_quiz_fk,
        DROP CONSTRAINT IF EXISTS sub_subject_quiz_fk;

ALTER TABLE
        IF EXISTS public.quiz_submissions 
        DROP CONSTRAINT IF EXISTS user_quiz_submission_fk,
        DROP CONSTRAINT IF EXISTS quiz_quiz_submission_fk;

ALTER TABLE
        IF EXISTS public.quiz_scores 
        DROP CONSTRAINT IF EXISTS submission_quiz_score_fk;

ALTER TABLE
        IF EXISTS public.quiz_questions 
        DROP CONSTRAINT IF EXISTS quiz_quiz_question_fk;

ALTER TABLE
        IF EXISTS public.quiz_choices 
        DROP CONSTRAINT IF EXISTS question_quiz_choice_fk;

ALTER TABLE
        IF EXISTS public.quiz_answers 
        DROP CONSTRAINT IF EXISTS submission_quiz_answer_fk,
        DROP CONSTRAINT IF EXISTS choice_quiz_answer_fk;

ALTER TABLE
        IF EXISTS public.plans 
        DROP CONSTRAINT IF EXISTS module_sub_plan_fk;

ALTER TABLE
        IF EXISTS public.exercises 
        DROP CONSTRAINT IF EXISTS test_type_exercise_fk;

ALTER TABLE
        IF EXISTS public.quiz_submissions 
        DROP CONSTRAINT IF EXISTS user_quiz_submission_fk,
        DROP CONSTRAINT IF EXISTS quiz_quiz_submission_fk;

ALTER TABLE
        IF EXISTS public.exercise_submissions 
        DROP CONSTRAINT IF EXISTS user_exercise_submission_fk,
        DROP CONSTRAINT IF EXISTS exercise_exercise_submission_fk;

ALTER TABLE
        IF EXISTS public.exercise_scores 
        DROP CONSTRAINT IF EXISTS submission_exercise_score_fk;

ALTER TABLE
        IF EXISTS public.exercise_questions 
        DROP CONSTRAINT IF EXISTS exercise_exercise_question_fk;

ALTER TABLE
        IF EXISTS public.exercise_choices 
        DROP CONSTRAINT IF EXISTS question_exercise_choice_fk;

ALTER TABLE
        IF EXISTS public.exercise_answers 
        DROP CONSTRAINT IF EXISTS submission_exercise_answer_fk,
        DROP CONSTRAINT IF EXISTS choice_exercise_answer_fk;

ALTER TABLE
        IF EXISTS public.class_user_plans 
        DROP CONSTRAINT IF EXISTS user_class_user_plan_fk,
        DROP CONSTRAINT IF EXISTS plan_class_user_plan_fk;