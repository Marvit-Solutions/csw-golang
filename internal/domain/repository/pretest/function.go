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
func (pr pretestRepo) GetPretestById(pretestId string) (error, datastruct.SubjectTestTypeQuizzes) {
	var pretests datastruct.SubjectTestTypeQuizzes

	err := pr.db.Preload("QuestionQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes").Where("id = ?", pretestId).Find(&pretests).Error
	if err != nil {
		return err, pretests
	}

	return nil, pretests
}
func (pr pretestRepo) GetPretestReview(pretestId, status string) (error, datastruct.SubjectTestTypeQuizzes) {
	var pretests datastruct.SubjectTestTypeQuizzes

	err := pr.db.Preload("QuestionQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes.UserSubmittedAnswerQuizzes").Where("id = ? AND status = ?", pretestId, status).Find(&pretests).Error
	if err != nil {	
		return err, pretests
	}

	return nil, pretests
}
