package localservice

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localrepository"
	"gorm.io/gorm"
)

type SubjectAdminService struct {
	DB *gorm.DB
}

func NewSubjectAdminService(
	DB *gorm.DB,
) localrepository.SubjectAdmin {
	return &SubjectAdminService{
		DB,
	}
}

func (svc *SubjectAdminService) GetSubjectAdminAll(searchKeywords string, subModulID int, limit int, offset int) ([]*response.SubjectAdminAllResponse, error) {
	subjectAll := make([]*response.SubjectAdminAllResponse, 0)
	query := `SELECT
		s.id, 
        s.uuid,
		s.name,
        sm.name AS sub_module_name,
        m.name AS module_name  
        FROM 
            subjects s
        JOIN 
            sub_modules sm ON s.sub_module_id = sm.id
        JOIN 
            modules m ON sm.module_id = m.id
        WHERE 
            s.deleted_at IS NULL`

	params := []interface{}{}

	if subModulID != 0 {
		query += ` AND sm.id = ?`
		params = append(params, subModulID)
	}

	if searchKeywords != "" {
		query += ` AND (LOWER(s.name) LIKE LOWER(?))`
		searchPattern := "%" + searchKeywords + "%"
		params = append(params, searchPattern)
	}

	query += ` ORDER BY s.id DESC LIMIT ? OFFSET ?`
	params = append(params, limit, offset)
	fmt.Println("ini query")

	res := svc.DB.Raw(query, params...).Scan(&subjectAll)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find subjectAll: %v", res.Error)
	}

	return subjectAll, nil
}

func (svc *SubjectAdminService) CountSubjectAdminALL(searchKeywords string, subModulID int) (int, error) {
	var totalCount int
	query := `SELECT COUNT(*) 
              FROM subjects s
              JOIN sub_modules sm ON s.sub_module_id = sm.id
              WHERE s.deleted_at IS NULL`

	params := []interface{}{}

	if subModulID != 0 {
		query += ` AND sm.id = ?`
		params = append(params, subModulID)
	}

	if searchKeywords != "" {
		query += ` AND LOWER(s.name) LIKE LOWER(?)`
		searchPattern := "%" + searchKeywords + "%"
		params = append(params, searchPattern)
	}

	res := svc.DB.Raw(query, params...).Scan(&totalCount)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to count subjectAll: %v", res.Error)
	}

	return totalCount, nil
}
