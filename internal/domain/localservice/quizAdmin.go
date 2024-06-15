package localservice

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"gorm.io/gorm"
)

type QuizAdminService struct {
	DB *gorm.DB
}

func NewQuizAdminService(
	DB *gorm.DB,
) localrepository.QuizAdmin {
	return &QuizAdminService{
		DB,
	}
}

func (svc *QuizAdminService) GetQuizAdminAll(searchKeywords string, testTypeID int, subModulID int, limit int, offset int) ([]*response.QuizAdminAllResponse, error) {
	quizAll := make([]*response.QuizAdminAllResponse, 0)
	query := `SELECT
		q.id, 
        q.uuid,
        sm.name AS sub_module_name,
        m.name AS module_name,
        s.name AS subject_name,
        tt.name AS test_type_name,
        q.title,
        q.description,
        q.time,
        q.open,
        q.close,
        q.max_score,
        q.attempt
        FROM 
            quizzes q
        JOIN 
            subjects s ON q.subject_id = s.id
        JOIN 
            sub_modules sm ON s.sub_module_id = sm.id
        JOIN 
            modules m ON sm.module_id = m.id
        JOIN 
            test_types tt ON q.test_type_id = tt.id
        WHERE 
            q.deleted_at IS NULL`

	params := []interface{}{}

	if testTypeID != 0 {
		query += ` AND q.test_type_id = ?`
		params = append(params, testTypeID)
	}

	if subModulID != 0 {
		query += ` AND sm.id = ?`
		params = append(params, subModulID)
	}

	if searchKeywords != "" {
		query += ` AND (q.title LIKE ? OR q.description LIKE ?)`
		searchPattern := "%" + searchKeywords + "%"
		params = append(params, searchPattern, searchPattern)
	}

	query += ` ORDER BY q.id DESC LIMIT ? OFFSET ?`
	params = append(params, limit, offset)

	res := svc.DB.Raw(query, params...).Scan(&quizAll)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find quizAll: %v", res.Error)
	}

	return quizAll, nil
}

func (svc *QuizAdminService) CountQuizAdminALL(searchKeywords string, testTypeID int, subModulID int) (int, error) {
	var totalCount int
	query := `SELECT COUNT(*) 
              FROM quizzes q
              JOIN subjects s ON q.subject_id = s.id
              JOIN sub_modules sm ON s.sub_module_id = sm.id
              WHERE q.deleted_at IS NULL`

	params := []interface{}{}

	if testTypeID != 0 {
		query += ` AND q.test_type_id = ?`
		params = append(params, testTypeID)
	}

	if subModulID != 0 {
		query += ` AND sm.id = ?`
		params = append(params, subModulID)
	}

	if searchKeywords != "" {
		query += ` AND (q.title LIKE ? OR q.description LIKE ?)`
		searchPattern := "%" + searchKeywords + "%"
		params = append(params, searchPattern, searchPattern)
	}

	res := svc.DB.Raw(query, params...).Scan(&totalCount)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to count quizAll: %v", res.Error)
	}

	return totalCount, nil
}
