package module

import (
	"csw-golang/internal/domain/entity/datastruct"
)

func (mr *moduleRepo) GetListModules() ([]datastruct.SubModules, error) {

	var moduleList []datastruct.SubModules

	err := mr.db.Preload("Subjects").Find(&moduleList).Error
	if err != nil {
		return nil, err
	}

	return moduleList, nil
}

func (mr *moduleRepo) GetListSubjectQuiz() ([]datastruct.SubjectTestTypeQuizzes, error) {

	var quizzes []datastruct.SubjectTestTypeQuizzes

	err := mr.db.Where("test_type LIKE '%Module%'").Find(&quizzes).Error
	if err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (mr *moduleRepo) GetSubjectsBySubmoduleID(submoduleID string) ([]datastruct.Subjects, error) {

	var subjects []datastruct.Subjects

	err := mr.db.Preload("SubSubject").Where("sub_module_id = ?", submoduleID).Find(&subjects).Error

	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (mr *moduleRepo) GetTestByTestTypeID(testTypeID string) (datastruct.SubjectTestTypeQuizzes, error) {

	var test datastruct.SubjectTestTypeQuizzes

	err := mr.db.Preload("QuestionQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes").Where("id = ?", testTypeID).Find(&test).Error

	if err != nil {
		return test, err
	}

	return test, nil
}
