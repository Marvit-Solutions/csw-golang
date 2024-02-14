package pretest

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"

	"github.com/google/uuid"
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

	// belum ada handle jawaban tiap user
	err := pr.db.Preload("QuestionQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes.UserSubmittedAnswerQuizzes").Where("id = ? AND status = ?", pretestId, status).Find(&pretests).Error
	if err != nil {
		return err, pretests
	}

	return nil, pretests
}

func (pr pretestRepo) SubmitPretest(id string, req dto.PretestSubmitRequest) error {
	var testType datastruct.SubjectTestTypeQuizzes
	err := pr.db.Where("id = ?", id).First(&testType).Error
	if err != nil {
		return err
	}

	var submissionPretest = datastruct.UserTestSubmissionQuizzes{
		ID:             uuid.NewString(),
		UserID:         req.UserID,
		TestTypeQuizID: testType.ID,
	}

	tx := pr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, choice := range req.ChoiceQuizzes {
		submissionPretest.UserSubmittedAnswerQuizzes = append(submissionPretest.UserSubmittedAnswerQuizzes, datastruct.UserSubmittedAnswerQuizzes{
			ID:                       uuid.NewString(),
			ChoiceQuizID:             choice.ChoiceQuizID,
			QuestionQuizID:           choice.QuestionQuizID,
			UserTestSubmissionQuizID: submissionPretest.ID,
		})
	}

	if err := pr.db.Create(&submissionPretest).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
