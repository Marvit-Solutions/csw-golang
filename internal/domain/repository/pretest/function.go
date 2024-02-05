package pretest

import (
	"csw-golang/internal/domain/entity/datastruct"
)

func (pr pretestRepo) GetAllPretests() (error, []datastruct.Modules) {
	var modules []datastruct.Modules

	err := pr.db.Preload("SubModules").Preload("SubModules.Subjects").Preload("SubModules.Subjects.SubjectTestTypeQuizzes").Find(&modules).Error
	if err != nil {
		return err, modules
	}

	return nil, modules
}
