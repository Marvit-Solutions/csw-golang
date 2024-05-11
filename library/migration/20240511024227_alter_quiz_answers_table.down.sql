ALTER TABLE
    public.quiz_answers DROP CONSTRAINT IF EXISTS question_quiz_answer_fk;

ALTER TABLE
    public.quiz_answers DROP COLUMN IF EXISTS question_id;