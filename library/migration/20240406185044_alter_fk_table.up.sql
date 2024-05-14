ALTER TABLE
        IF EXISTS public.class_user_plans
ADD
        CONSTRAINT plan_class_user_plan_fk FOREIGN KEY (plan_id) REFERENCES public.plans (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.class_user_plans
ADD
        CONSTRAINT user_class_user_plan_fk FOREIGN KEY (user_id) REFERENCES public.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercise_answers
ADD
        CONSTRAINT choice_exercise_answer_fk FOREIGN KEY (choice_id) REFERENCES public.exercise_choices (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercise_choices
ADD
        CONSTRAINT question_exercise_choice_fk FOREIGN KEY (question_id) REFERENCES public.exercise_questions (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercise_questions
ADD
        CONSTRAINT exercise_exercise_question_fk FOREIGN KEY (exercise_id) REFERENCES public.exercises (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercise_submissions
ADD
        CONSTRAINT user_exercise_submission FOREIGN KEY (user_id) REFERENCES public.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercise_submissions
ADD
        CONSTRAINT exercise_exercise_submissiion FOREIGN KEY (exercise_id) REFERENCES public.exercises (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercises
ADD
        CONSTRAINT test_type_exercise_fk FOREIGN KEY (test_type_id) REFERENCES public.test_types (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercises
ADD
        CONSTRAINT module_exercise_fk FOREIGN KEY (module_id) REFERENCES public.modules (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.mentors
ADD
        CONSTRAINT module_mentor_fk FOREIGN KEY (module_id) REFERENCES public.modules (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.mentors
ADD
        CONSTRAINT user_mentor_fk FOREIGN KEY (user_id) REFERENCES public.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.plans
ADD
        CONSTRAINT module_plan_fk FOREIGN KEY (module_id) REFERENCES public.modules (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.plans
ADD
        CONSTRAINT media_plan_fk FOREIGN KEY (media_id) REFERENCES public.medias (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_answers
ADD
        CONSTRAINT choice_quiz_answer_fk FOREIGN KEY (choice_id) REFERENCES public.quiz_choices (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_answers
ADD
        CONSTRAINT submission_quiz_answer_fk FOREIGN KEY (submission_id) REFERENCES public.quiz_submissions (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_choices
ADD
        CONSTRAINT question_quiz_choice_fk FOREIGN KEY (question_id) REFERENCES public.quiz_questions (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_questions
ADD
        CONSTRAINT quiz_quiz_question_fk FOREIGN KEY (quiz_id) REFERENCES public.quizzes (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_submissions
ADD
        CONSTRAINT quiz_quiz_submission_fk FOREIGN KEY (quiz_id) REFERENCES public.quizzes (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_submissions
ADD
        CONSTRAINT user_quiz_submission_fk FOREIGN KEY (user_id) REFERENCES public.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quizzes
ADD
        CONSTRAINT subject_quiz_fk FOREIGN KEY (subject_id) REFERENCES public.subjects (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quizzes
ADD
        CONSTRAINT test_type_quiz_fk FOREIGN KEY (test_type_id) REFERENCES public.test_types (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.schedules
ADD
        CONSTRAINT class_user_plan_schedule_fk FOREIGN KEY (class_user_plan_id) REFERENCES public.class_user_plans (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.schedules
ADD
        CONSTRAINT sub_subject_schedule_fk FOREIGN KEY (sub_subject_id) REFERENCES public.sub_subjects (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.sub_modules
ADD
        CONSTRAINT module_sub_module_fk FOREIGN KEY (module_id) REFERENCES public.modules (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.sub_subjects
ADD
        CONSTRAINT subject_sub_subject_fk FOREIGN KEY (subject_id) REFERENCES public.subjects (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.subjects
ADD
        CONSTRAINT sub_module_subject_fk FOREIGN KEY (sub_module_id) REFERENCES public.sub_modules (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.uniques
ADD
        CONSTRAINT mentor_unique_fk FOREIGN KEY (mentor_id) REFERENCES public.mentors (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.user_details
ADD
        CONSTRAINT class_user_user_detail_fk FOREIGN KEY (class_user_id) REFERENCES public.class_users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.user_details
ADD
        CONSTRAINT user_user_detail_fk FOREIGN KEY (user_id) REFERENCES public.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.user_details
ADD
        CONSTRAINT media_user_fk FOREIGN KEY (media_id) REFERENCES public.medias (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.user_mentor_testimonials
ADD
        CONSTRAINT mentor_user_mentor_testimonial_fk FOREIGN KEY (mentor_id) REFERENCES public.mentors (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.user_mentor_testimonials
ADD
        CONSTRAINT user_user_mentor_testimonial_fk FOREIGN KEY (user_id) REFERENCES public.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.user_testimonials
ADD
        CONSTRAINT user_user_testimonial_fk FOREIGN KEY (user_id) REFERENCES public.users (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.users
ADD
        CONSTRAINT role_user_fk FOREIGN KEY (role_id) REFERENCES public.roles (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_question_media
ADD
        CONSTRAINT quiz_question_quiz_question_media FOREIGN KEY (quiz_question_id) REFERENCES public.quiz_questions (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.quiz_question_media
ADD
        CONSTRAINT media_quiz_question_media FOREIGN KEY (media_id) REFERENCES public.medias (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercise_question_media
ADD
        CONSTRAINT exercise_question_exercise_question_media FOREIGN KEY (exercise_question_id) REFERENCES public.exercise_questions (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.exercise_question_media
ADD
        CONSTRAINT media_exercise_question_media FOREIGN KEY (media_id) REFERENCES public.medias (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.sub_subject_media
ADD
        CONSTRAINT sub_subject_sub_subject_media FOREIGN KEY (sub_subject_id) REFERENCES public.sub_subjects (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;

ALTER TABLE
        IF EXISTS public.sub_subject_media
ADD
        CONSTRAINT media_sub_subject_media FOREIGN KEY (media_id) REFERENCES public.medias (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;