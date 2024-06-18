package localservice

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localrepository"
	"gorm.io/gorm"
)

type MaterialAdminService struct {
	DB *gorm.DB
}

func NewMaterialAdminService(
	DB *gorm.DB,
) localrepository.MaterialAdmin {
	return &MaterialAdminService{
		DB,
	}
}

func (svc *MaterialAdminService) GetMaterialAdminAll(searchKeywords string, subjectID int, subModulID int, limit int, offset int) ([]*response.MaterialAdminAllResponse, error) {
	materialAll := make([]*response.MaterialAdminAllResponse, 0)
	query := `SELECT
		ss.id, 
        ss.uuid,
        sm.name AS sub_module_name,
        m.name AS module_name,
        s.name AS subject_name,
        ss.name,
        ss.content
        FROM 
            sub_subjects ss
        JOIN 
            subjects s ON ss.subject_id = s.id
        JOIN 
            sub_modules sm ON s.sub_module_id = sm.id
        JOIN 
            modules m ON sm.module_id = m.id
        WHERE 
            ss.deleted_at IS NULL`

	params := []interface{}{}

	if subjectID != 0 {
		query += ` AND s.id = ?`
		params = append(params, subjectID)
	}

	if subModulID != 0 {
		query += ` AND sm.id = ?`
		params = append(params, subModulID)
	}

	if searchKeywords != "" {
		query += ` AND (LOWER(ss.name) LIKE LOWER(?) OR LOWER(ss.content) LIKE LOWER(?))`
		searchPattern := "%" + searchKeywords + "%"
		params = append(params, searchPattern, searchPattern)
	}

	query += ` ORDER BY ss.id DESC LIMIT ? OFFSET ?`
	params = append(params, limit, offset)
	fmt.Println("ini query")
	fmt.Println()

	res := svc.DB.Raw(query, params...).Scan(&materialAll)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find materialAll: %v", res.Error)
	}

	return materialAll, nil
}

func (svc *MaterialAdminService) CountMaterialAdminALL(searchKeywords string, subjectID int, subModulID int) (int, error) {
	var totalCount int
	query := `SELECT COUNT(*) 
              FROM sub_subjects ss
              JOIN subjects s ON ss.subject_id = s.id
              JOIN sub_modules sm ON s.sub_module_id = sm.id
              WHERE ss.deleted_at IS NULL`

	params := []interface{}{}

	if subjectID != 0 {
		query += ` AND s.id = ?`
		params = append(params, subjectID)
	}

	if subModulID != 0 {
		query += ` AND sm.id = ?`
		params = append(params, subModulID)
	}

	if searchKeywords != "" {
		query += ` AND (LOWER(ss.name) LIKE LOWER(?) OR LOWER(ss.content) LIKE LOWER(?))`
		searchPattern := "%" + searchKeywords + "%"
		params = append(params, searchPattern, searchPattern)
	}

	res := svc.DB.Raw(query, params...).Scan(&totalCount)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to count materialAll: %v", res.Error)
	}

	return totalCount, nil
}
