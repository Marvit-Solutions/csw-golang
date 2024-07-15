package localservice

import (
	"fmt"
	"strings"

	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localrepository"
	"gorm.io/gorm"
)

type ExerciseService struct {
	DB *gorm.DB
}

func NewExerciseService(
	DB *gorm.DB,
) localrepository.Exercise {
	return &ExerciseService{
		DB,
	}
}

func (svc *ExerciseService) FindSubModulesID() ([]int, error) {
	query := `SELECT ARRAY_AGG(DISTINCT sub_module_id ORDER BY sub_module_id) FROM exercise_questions`
	var subModuleIDsStr string

	res := svc.DB.Raw(query).Scan(&subModuleIDsStr)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find user answers: %v", res.Error)
	}

	subModuleIDsStr = strings.Trim(subModuleIDsStr, "{}")
	subModuleIDStrings := strings.Split(subModuleIDsStr, ",")
	subModuleIDs := make([]int, len(subModuleIDStrings))

	for i, idStr := range subModuleIDStrings {
		var id int
		_, err := fmt.Sscanf(idStr, "%d", &id)
		if err != nil {
			return nil, fmt.Errorf("failed to parse subModuleID %s: %v", idStr, err)
		}
		subModuleIDs[i] = id
	}

	for _, subModuleID := range subModuleIDs {
		fmt.Printf("subModuleID: %d\n", subModuleID)
	}
	return subModuleIDs, nil
}
