ALTER TABLE
    public.quiz_questions RENAME point TO score;

ALTER TABLE
    public.quiz_questions
ADD
    COLUMN explanation text NOT NULL
AFTER
    score;