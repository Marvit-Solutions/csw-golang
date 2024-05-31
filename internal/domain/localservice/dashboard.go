package localservice

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"gorm.io/gorm"
)

type DashboardService struct {
	DB *gorm.DB
}

func NewDashboardService(
	DB *gorm.DB,
) localrepository.Dashboard {
	return &DashboardService{
		DB,
	}
}
func (svc *DashboardService) GetQuizAllDashboard(authenticatedUserID int, testTypeID int, modulID int) ([]*response.QuizItemAllDashboard, error) {
	quizAll := make([]*response.QuizItemAllDashboard, 0)
	// for future enhancement, use string concatenation
	query := ""

	query = `SELECT 
    q.*,
    sm.name AS sub_module_name,
    sm.uuid AS sub_module_uuid,
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
JOIN 
    modules m ON sm.module_id = m.id	
LEFT JOIN 
    quiz_submissions qs ON q.id = qs.quiz_id AND qs.user_id = ?
WHERE 
    q.test_type_id = ?
    AND m.id = ?
    AND q.deleted_at IS NULL
    AND qs.id IS NULL
ORDER BY 
    q.created_at DESC
LIMIT 5;
`

	res := svc.DB.Raw(query, authenticatedUserID, testTypeID, modulID).Scan(&quizAll)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find quizAll: %v", res.Error)
	}

	for _, quiz := range quizAll {
		fmt.Printf("quiz title: %s, quiz description: %s, quiz status pengerjaan: %s\n", quiz.Title, quiz.Description, quiz.StatusPengerjaan)
	}
	return quizAll, nil
}
