package localservice

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"gorm.io/gorm"
)

type QuizService struct {
	DB *gorm.DB
}

func NewQuizService(
	DB *gorm.DB,
) localrepository.Quiz {
	return &QuizService{
		DB,
	}
}

func (svc *QuizService) FindUserAnswerReview(quizSubmissionID int) ([]*response.UserAnswer, error) {
	userAnswers := make([]*response.UserAnswer, 0)
	query := `
		SELECT 
			qq.id AS question_id,
			qq.content AS question_content,
			qc.id AS choice_id,
			qc.content AS choices_content
		FROM 
			quiz_questions AS qq 
		INNER JOIN 
			quiz_choices AS qc 
		ON 
			qc.question_id = qq.id 
		WHERE 
			qc.id IN (
				SELECT DISTINCT(choice_id) 
				FROM quiz_answers 
				WHERE submission_id = ?
			)
	`

	res := svc.DB.Raw(query, quizSubmissionID).Scan(&userAnswers)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find user answers: %v", res.Error)
	}

	// for _, answer := range userAnswers {
	// 	fmt.Printf("Question ID: %d, Question Content: %s\n", answer.QuestionId, answer.QuestionContent)
	// 	fmt.Printf("Choice ID: %d, Choice Content: %s\n", answer.ChoiceId, answer.ChoiceContent)
	// }
	return userAnswers, nil
}

func (svc *QuizService) CountQuizzesGroupedBySubModule(moduleID int, testTypeID int) ([]*response.QuizzesGroupedBySubModule, error) {
	quizzesGroupedBySubModule := make([]*response.QuizzesGroupedBySubModule, 0)
	query := `
	SELECT 
    sm.id AS sub_module_id,
	sm.uuid AS sub_module_uuid,
    sm.name AS sub_module_name,
    COUNT(DISTINCT q.id) AS quiz_count,
    (
        SELECT COUNT(*)
        FROM (
            SELECT DISTINCT qs.user_id, qs.quiz_id
            FROM quiz_submissions qs
            JOIN quizzes q2 ON qs.quiz_id = q2.id
            JOIN subjects s2 ON q2.subject_id = s2.id
            WHERE s2.sub_module_id = sm.id AND q2.test_type_id = ?
        ) AS unique_submissions
    ) AS submission_count
	FROM 
		modules m
	JOIN 
		sub_modules sm ON sm.module_id = m.id
	JOIN 
		subjects s ON s.sub_module_id = sm.id
	JOIN 
		quizzes q ON q.subject_id = s.id
	WHERE 
		m.id = ? AND 
		q.test_type_id = ?
	GROUP BY 
		m.id, m.name, sm.id, sm.name
	ORDER BY 
		m.id, sm.id;
`

	res := svc.DB.Raw(query, testTypeID, moduleID, testTypeID).Scan(&quizzesGroupedBySubModule)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find quizzesGroupedBySubModule: %v", res.Error)
	}

	// for _, quizGroupedBySubModule := range quizzesGroupedBySubModule {
	// 	fmt.Printf("sub module id: %d, sub module name: %s, quiz count: %d, submission count : %d\n", quizGroupedBySubModule.SubModuleID, quizGroupedBySubModule.SubModuleName, quizGroupedBySubModule.QuizCount, quizGroupedBySubModule.SubmissionCount)
	// }
	return quizzesGroupedBySubModule, nil
}

func (svc *QuizService) GetQuizAll(testTypeID int, subModulID int, limit int, offset int) ([]*response.QuizItemAll, error) {
	quizAll := make([]*response.QuizItemAll, 0)
	// for future enhancement, use string concatenation
	query := ""

	query = `SELECT 
    q.*,
    qs.uuid AS quiz_submission_uuid,
    CASE 
        WHEN qs.id IS NOT NULL THEN 'sudah-dikerjakan'
        ELSE 'belum-dikerjakan'
    END AS status_pengerjaan
	FROM 
		quizzes q
	JOIN 
		subjects s ON q.subject_id = s.id
	JOIN 
		sub_modules sm ON s.sub_module_id = sm.id
	LEFT JOIN 
		quiz_submissions qs ON q.id = qs.quiz_id AND qs.user_id = 40
	WHERE 
		q.test_type_id = ?
		AND sm.id = ?
		AND q.deleted_at IS NULL
	ORDER BY 
		q.created_at DESC
	LIMIT ?
	OFFSET ?;`

	res := svc.DB.Raw(query, testTypeID, subModulID, limit, offset).Scan(&quizAll)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find quizAll: %v", res.Error)
	}

	// for _, quiz := range quizAll {
	// 	fmt.Printf("quiz title: %s, quiz description: %s, quiz status pengerjaan: %s\n", quiz.Title, quiz.Description, quiz.StatusPengerjaan)
	// }
	return quizAll, nil
}

func (svc *QuizService) CountQuizALL(testTypeID int, subModulID int) (int, error) {
	var count int
	query := `SELECT 
    COUNT(*)
	FROM 
		quizzes q
	JOIN 
		subjects s ON q.subject_id = s.id
	JOIN 
		sub_modules sm ON s.sub_module_id = sm.id
	LEFT JOIN 
		quiz_submissions qs ON q.id = qs.quiz_id AND qs.user_id = 40
	WHERE 
		q.test_type_id = ?
		AND sm.id = ?
		AND q.deleted_at IS NULL`

	res := svc.DB.Raw(query, testTypeID, subModulID).Scan(&count)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to count quizzes: %v", res.Error)
	}

	return count, nil
}
