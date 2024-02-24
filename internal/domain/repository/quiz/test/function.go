package tests

import "csw-golang/internal/domain/entity/datastruct"

func (tr testRepo) GetAllTests() ([]datastruct.Modules, error) {
	var modules []datastruct.Modules

	err := tr.db.Preload("SubModules").Preload("SubModules.Subjects").Preload("SubModules.Subjects.SubjectTestTypeQuizzes").Find(&modules).Error
	if err != nil {
		return modules, err
	}

	return modules, nil
}
