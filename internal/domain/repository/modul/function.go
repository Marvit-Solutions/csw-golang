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

func (mr *moduleRepo) GetQuestionsByTestTypeID(testTypeID string) ([]datastruct.QuestionQuizzes, error) {

	var questions []datastruct.QuestionQuizzes

	err := mr.db.Preload("ChoiceQuizzes").Where("test_type_id = ?", testTypeID).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	return questions, nil
}
