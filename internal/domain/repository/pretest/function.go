package pretest

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"

	"github.com/google/uuid"
)

func (pr pretestRepo) GetAllPretests(userID, module string) (error, []datastruct.Modules) {
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
	err := pr.db.Preload("GradeQuizzes").Preload("QuestionQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes.UserSubmittedAnswerQuizzes").Where("id = ? AND status = ?", pretestId, status).Find(&pretests).Error
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

	for _, choice := range req.ChoiceQuizzes {
		submissionPretest.UserSubmittedAnswerQuizzes = append(submissionPretest.UserSubmittedAnswerQuizzes, datastruct.UserSubmittedAnswerQuizzes{
			ID:                       uuid.NewString(),
			ChoiceQuizID:             choice.ChoiceQuizID,
			QuestionQuizID:           choice.QuestionQuizID,
			UserTestSubmissionQuizID: submissionPretest.ID,
		})
	}

	tx := pr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := pr.db.Create(&submissionPretest).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
func (pr pretestRepo) GradingPretest(id string, req dto.PretestSubmitRequest) error {
	// Get Questions and Choices
	var questions []datastruct.QuestionQuizzes
	err := pr.db.Preload("ChoiceQuizzes").Where("test_type_quiz_id = ?", id).Find(&questions).Error
	if err != nil {
		return err
	}

	// Get User Answer
	var submissionPretest datastruct.UserTestSubmissionQuizzes
	err = pr.db.Preload("UserSubmittedAnswerQuizzes").Where("test_type_quiz_id = ? AND user_id = ?", id, req.UserID).Find(&submissionPretest).Error
	if err != nil {
		return err
	}

	score := 0
	for _, question := range questions {
		for _, choice := range question.ChoiceQuizzes {
			for _, userAnswer := range submissionPretest.UserSubmittedAnswerQuizzes {
				if userAnswer.ChoiceQuizID == choice.ID && choice.IsCorrect == true {
					score += choice.Weight
				}
			}
		}
	}

	// Grading
	grade := datastruct.GradeQuizzes{
		ID:                       uuid.NewString(),
		UserID:                   req.UserID,
		TestTypeQuizID:           id,
		Score:                    score,
		UserTestSubmissionQuizID: submissionPretest.ID,
	}

	tx := pr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := pr.db.Create(&grade).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
