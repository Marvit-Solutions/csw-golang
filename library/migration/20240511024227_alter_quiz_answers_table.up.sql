ALTER TABLE
    public.quiz_answers
ADD
    COLUMN question_id integer NOT NULL
AFTER
    submission_id;

ALTER TABLE
    public.quiz_answers
ADD
    CONSTRAINT question_quiz_answer_fk FOREIGN KEY (question_id) REFERENCES public.quiz_questions (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID;