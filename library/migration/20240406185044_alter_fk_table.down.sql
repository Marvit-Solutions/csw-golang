ALTER TABLE
        IF EXISTS public.sub_subject_media DROP CONSTRAINT IF EXISTS media_sub_subject_media;

ALTER TABLE
        IF EXISTS public.sub_subject_media DROP CONSTRAINT IF EXISTS sub_subject_sub_subject_media;

ALTER TABLE
        IF EXISTS public.exercise_question_media DROP CONSTRAINT IF EXISTS media_exercise_question_media;

ALTER TABLE
        IF EXISTS public.exercise_question_media DROP CONSTRAINT IF EXISTS exercise_question_exercise_question_media;

ALTER TABLE
        IF EXISTS public.quiz_question_media DROP CONSTRAINT IF EXISTS media_quiz_question_media;

ALTER TABLE
        IF EXISTS public.quiz_question_media DROP CONSTRAINT IF EXISTS quiz_question_quiz_question_media;

ALTER TABLE
        IF EXISTS public.users DROP CONSTRAINT IF EXISTS role_user_fk;

ALTER TABLE
        IF EXISTS public.user_testimonials DROP CONSTRAINT IF EXISTS user_user_testimonial_fk;

ALTER TABLE
        IF EXISTS public.user_mentor_testimonials DROP CONSTRAINT IF EXISTS user_user_mentor_testimonial_fk;

ALTER TABLE
        IF EXISTS public.user_mentor_testimonials DROP CONSTRAINT IF EXISTS mentor_user_mentor_testimonial_fk;

ALTER TABLE
        IF EXISTS public.user_details DROP CONSTRAINT IF EXISTS media_user_fk;

ALTER TABLE
        IF EXISTS public.user_details DROP CONSTRAINT IF EXISTS user_user_detail_fk;

ALTER TABLE
        IF EXISTS public.user_details DROP CONSTRAINT IF EXISTS class_user_user_detail_fk;

ALTER TABLE
        IF EXISTS public.uniques DROP CONSTRAINT IF EXISTS mentor_unique_fk;

ALTER TABLE
        IF EXISTS public.subjects DROP CONSTRAINT IF EXISTS sub_module_subject_fk;

ALTER TABLE
        IF EXISTS public.sub_subjects DROP CONSTRAINT IF EXISTS subject_sub_subject_fk;

ALTER TABLE
        IF EXISTS public.sub_modules DROP CONSTRAINT IF EXISTS module_sub_module_fk;

ALTER TABLE
        IF EXISTS public.schedules DROP CONSTRAINT IF EXISTS sub_subject_schedule_fk;

ALTER TABLE
        IF EXISTS public.schedules DROP CONSTRAINT IF EXISTS class_user_plan_schedule_fk;

ALTER TABLE
        IF EXISTS public.quizzes DROP CONSTRAINT IF EXISTS test_type_quiz_fk;

ALTER TABLE
        IF EXISTS public.quizzes DROP CONSTRAINT IF EXISTS subject_quiz_fk;

ALTER TABLE
        IF EXISTS public.quiz_submissions DROP CONSTRAINT IF EXISTS user_quiz_submission_fk;

ALTER TABLE
        IF EXISTS public.quiz_submissions DROP CONSTRAINT IF EXISTS quiz_quiz_submission_fk;

ALTER TABLE
        IF EXISTS public.quiz_questions DROP CONSTRAINT IF EXISTS quiz_quiz_question_fk;

ALTER TABLE
        IF EXISTS public.quiz_choices DROP CONSTRAINT IF EXISTS question_quiz_choice_fk;

ALTER TABLE
        IF EXISTS public.quiz_answers DROP CONSTRAINT IF EXISTS submission_quiz_answer_fk;

ALTER TABLE
        IF EXISTS public.quiz_answers DROP CONSTRAINT IF EXISTS choice_quiz_answer_fk;

ALTER TABLE
        IF EXISTS public.plans DROP CONSTRAINT IF EXISTS module_sub_plan_fk;

ALTER TABLE
        IF EXISTS public.mentors DROP CONSTRAINT IF EXISTS user_mentor_fk;

ALTER TABLE
        IF EXISTS public.mentors DROP CONSTRAINT IF EXISTS module_mentor_fk;

ALTER TABLE
        IF EXISTS public.exercises DROP CONSTRAINT IF EXISTS test_type_exercise_fk;

ALTER TABLE
        IF EXISTS public.exercise_submissions DROP CONSTRAINT IF EXISTS exercise_exercise_submissiion;

ALTER TABLE
        IF EXISTS public.exercise_submissions DROP CONSTRAINT IF EXISTS user_exercise_submission;

ALTER TABLE
        IF EXISTS public.exercise_questions DROP CONSTRAINT IF EXISTS exercise_exercise_question_fk;

ALTER TABLE
        IF EXISTS public.exercise_choices DROP CONSTRAINT IF EXISTS question_exercise_choice_fk;

ALTER TABLE
        IF EXISTS public.exercise_answers DROP CONSTRAINT IF EXISTS choice_exercise_answer_fk;

ALTER TABLE
        IF EXISTS public.class_user_plans DROP CONSTRAINT IF EXISTS user_class_user_plan_fk;

ALTER TABLE
        IF EXISTS public.class_user_plans DROP CONSTRAINT IF EXISTS plan_class_user_plan_fk;

ALTER TABLE
        IF EXISTS public.quiz_answers DROP CONSTRAINT IF EXISTS choice_quiz_answer_fk;

ALTER TABLE
        IF EXISTS public.quiz_submissions DROP CONSTRAINT IF EXISTS submission_quiz_answer_fk;

ALTER TABLE
        IF EXISTS public.quiz_questions DROP CONSTRAINT IF EXISTS quiz_quiz_question_fk;

ALTER TABLE
        IF EXISTS public.quiz_choices DROP CONSTRAINT IF EXISTS question_quiz_choice_fk;

ALTER TABLE
        IF EXISTS public.quiz_submissions DROP CONSTRAINT IF EXISTS quiz_quiz_submission_fk;

ALTER TABLE
        IF EXISTS public.quiz_submissions DROP CONSTRAINT IF EXISTS user_quiz_submission_fk;

ALTER TABLE
        IF EXISTS public.quizzes DROP CONSTRAINT IF EXISTS subject_quiz_fk;

ALTER TABLE
        IF EXISTS public.quizzes DROP CONSTRAINT IF EXISTS test_type_quiz_fk;

ALTER TABLE
        IF EXISTS public.schedules DROP CONSTRAINT IF EXISTS class_user_plan_schedule_fk;

ALTER TABLE
        IF EXISTS public.schedules DROP CONSTRAINT IF EXISTS sub_subject_schedule_fk;

ALTER TABLE
        IF EXISTS public.sub_modules DROP CONSTRAINT IF EXISTS module_sub_module_fk;

ALTER TABLE
        IF EXISTS public.sub_subjects DROP CONSTRAINT IF EXISTS subject_sub_subject_fk;

ALTER TABLE
        IF EXISTS public.subjects DROP CONSTRAINT IF EXISTS sub_module_subject_fk;

ALTER TABLE
        IF EXISTS public.uniques DROP CONSTRAINT IF EXISTS mentor_unique_fk;

ALTER TABLE
        IF EXISTS public.user_details DROP CONSTRAINT IF EXISTS media_user_fk;

ALTER TABLE
        IF EXISTS public.user_details DROP CONSTRAINT IF EXISTS user_user_detail_fk;

ALTER TABLE
        IF EXISTS public.user_details DROP CONSTRAINT IF EXISTS class_user_user_detail_fk;

ALTER TABLE
        IF EXISTS public.user_mentor_testimonials DROP CONSTRAINT IF EXISTS mentor_user_mentor_testimonial_fk;

ALTER TABLE
        IF EXISTS public.user_mentor_testimonials DROP CONSTRAINT IF EXISTS user_user_mentor_testimonial_fk;

ALTER TABLE
        IF EXISTS public.user_testimonials DROP CONSTRAINT IF EXISTS user_user_testimonial_fk;

ALTER TABLE
        IF EXISTS public.users DROP CONSTRAINT IF EXISTS role_user_fk;

ALTER TABLE
        IF EXISTS public.quiz_question_media DROP CONSTRAINT IF EXISTS quiz_question_quiz_question_media;

ALTER TABLE
        IF EXISTS public.quiz_question_media DROP CONSTRAINT IF EXISTS media_quiz_question_media;

ALTER TABLE
        IF EXISTS public.exercise_question_media DROP CONSTRAINT IF EXISTS exercise_question_exercise_question_media;

ALTER TABLE
        IF EXISTS public.exercise_question_media DROP CONSTRAINT IF EXISTS media_exercise_question_media;

ALTER TABLE
        IF EXISTS public.sub_subject_media DROP CONSTRAINT IF EXISTS sub_subject_sub_subject_media;

ALTER TABLE
        IF EXISTS public.sub_subject_media DROP CONSTRAINT IF EXISTS media_sub_subject_media;

ALTER TABLE
        IF EXISTS public.medias DROP CONSTRAINT IF EXISTS media_sub_subject_media;

ALTER TABLE
        IF EXISTS public.medias DROP CONSTRAINT IF EXISTS media_exercise_question_media;

ALTER TABLE
        IF EXISTS public.medias DROP CONSTRAINT IF EXISTS media_quiz_question_media;

ALTER TABLE
        IF EXISTS public.medias DROP CONSTRAINT IF EXISTS media_quiz_question_media;

ALTER TABLE
        IF EXISTS public.medias DROP CONSTRAINT IF EXISTS media_sub_subject_media;

ALTER TABLE
        IF EXISTS public.medias DROP CONSTRAINT IF EXISTS media_quiz_question_media;

ALTER TABLE
        IF EXISTS public.class_user_plans DROP CONSTRAINT IF EXISTS class_user_plans_pkey;

ALTER TABLE
        IF EXISTS public.exercise_answers DROP CONSTRAINT IF EXISTS exercise_answers_pkey;

ALTER TABLE
        IF EXISTS public.exercise_choices DROP CONSTRAINT IF EXISTS exercise_choices_pkey;

ALTER TABLE
        IF EXISTS public.exercise_questions DROP CONSTRAINT IF EXISTS exercise_questions_pkey;

ALTER TABLE
        IF EXISTS public.exercise_submissions DROP CONSTRAINT IF EXISTS exercise_submissions_pkey;

ALTER TABLE
        IF EXISTS public.exercises DROP CONSTRAINT IF EXISTS exercises_pkey;

ALTER TABLE
        IF EXISTS public.mentors DROP CONSTRAINT IF EXISTS mentors_pkey;

ALTER TABLE
        IF EXISTS public.modules DROP CONSTRAINT IF EXISTS modules_pkey;

ALTER TABLE
        IF EXISTS public.plans DROP CONSTRAINT IF EXISTS plans_pkey;

ALTER TABLE
        IF EXISTS public.quiz_answers DROP CONSTRAINT IF EXISTS quiz_answers_pkey;

ALTER TABLE
        IF EXISTS public.quiz_choices DROP CONSTRAINT IF EXISTS quiz_choices_pkey;

ALTER TABLE
        IF EXISTS public.quiz_questions DROP CONSTRAINT IF EXISTS quiz_questions_pkey;

ALTER TABLE
        IF EXISTS public.quiz_submissions DROP CONSTRAINT IF EXISTS quiz_submissions_pkey;

ALTER TABLE
        IF EXISTS public.quizzes DROP CONSTRAINT IF EXISTS quizzes_pkey;

ALTER TABLE
        IF EXISTS public.roles DROP CONSTRAINT IF EXISTS roles_pkey;

ALTER TABLE
        IF EXISTS public.schedules DROP CONSTRAINT IF EXISTS schedules_pkey;

ALTER TABLE
        IF EXISTS public.sub_modules DROP CONSTRAINT IF EXISTS sub_modules_pkey;

ALTER TABLE
        IF EXISTS public.sub_subject_media DROP CONSTRAINT IF EXISTS sub_subject_media_pkey;

ALTER TABLE
        IF EXISTS public.sub_subjects DROP CONSTRAINT IF EXISTS sub_subjects_pkey;

ALTER TABLE
        IF EXISTS public.subjects DROP CONSTRAINT IF EXISTS subjects_pkey;

ALTER TABLE
        IF EXISTS public.test_types DROP CONSTRAINT IF EXISTS test_types_pkey;

ALTER TABLE
        IF EXISTS public.uniques DROP CONSTRAINT IF EXISTS uniques_pkey;

ALTER TABLE
        IF EXISTS public.user_details DROP CONSTRAINT IF EXISTS user_details_pkey;

ALTER TABLE
        IF EXISTS public.user_mentor_testimonials DROP CONSTRAINT IF EXISTS user_mentor_testimonials_pkey;

ALTER TABLE
        IF EXISTS public.user_testimonials DROP CONSTRAINT IF EXISTS user_testimonials_pkey;

ALTER TABLE
        IF EXISTS public.users DROP CONSTRAINT IF EXISTS users_pkey;

DROP TABLE IF EXISTS public.class_user_plans;

DROP TABLE IF EXISTS public.exercise_answers;

DROP TABLE IF EXISTS public.exercise_choices;

DROP TABLE IF EXISTS public.exercise_questions;

DROP TABLE IF EXISTS public.exercise_submissions;

DROP TABLE IF EXISTS public.exercises;

DROP TABLE IF EXISTS public.mentors;

DROP TABLE IF EXISTS public.modules;

DROP TABLE IF EXISTS public.plans;

DROP TABLE IF EXISTS public.quiz_answers;

DROP TABLE IF EXISTS public.quiz_choices;

DROP TABLE IF EXISTS public.quiz_questions;

DROP TABLE IF EXISTS public.quiz_submissions;

DROP TABLE IF EXISTS public.quizzes;

DROP TABLE IF EXISTS public.roles;

DROP TABLE IF EXISTS public.schedules;

DROP TABLE IF EXISTS public.sub_modules;

DROP TABLE IF EXISTS public.sub_subject_media;

DROP TABLE IF EXISTS public.sub_subjects;

DROP TABLE IF EXISTS public.subjects;

DROP TABLE IF EXISTS public.test_types;

DROP TABLE IF EXISTS public.uniques;

DROP TABLE IF EXISTS public.user_details;

DROP TABLE IF EXISTS public.user_mentor_testimonials;

DROP TABLE IF EXISTS public.user_testimonials;

DROP TABLE IF EXISTS public.users;

DROP TABLE IF EXISTS public.quiz_question_media;

DROP TABLE IF EXISTS public.medias;

DROP TABLE IF EXISTS public.quiz_question_media;

DROP TABLE IF EXISTS public.exercise_question_media;