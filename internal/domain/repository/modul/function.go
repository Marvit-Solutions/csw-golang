package module

import (
	"csw-golang/internal/domain/entity/datastruct"
	"fmt"
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

func (mr *moduleRepo) GetTestReview(moduleTestId string) (datastruct.SubjectTestTypeQuizzes, error) {
	var test datastruct.SubjectTestTypeQuizzes

	err := mr.db.Preload("QuestionQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes").Preload("QuestionQuizzes.ChoiceQuizzes.UserSubmittedAnswerQuizzes").Where("id = ?", moduleTestId).Find(&test).Error
	if err != nil {
		return test, err
	}

	return test, err
}

func (mr *moduleRepo) GetSubmittedReviewByID(moduleTestId string) (datastruct.UserTestSubmissionQuizzes, error) {
	var test datastruct.UserTestSubmissionQuizzes

	err := mr.db.Preload("UserSubmittedAnswerQuizzes").Preload("GradeQuiz").Where("id = ?", moduleTestId).Find(&test).Error
	if err != nil {
		return test, err
	}
	return test, err
}

func (mr *moduleRepo) AddQuizSubmission(submission datastruct.UserTestSubmissionQuizzes) (string, error) {
	err := mr.db.Create(&submission).Error
	if err != nil {
		return "", err
	}

	return submission.ID, nil

}

func (mr *moduleRepo) PostSubmittedQuizAnswer(testTypeID string, submittedAnswers []datastruct.UserSubmittedAnswerQuizzes) error {

	err := mr.db.Create(&submittedAnswers).Error
	if err != nil {
		return err
	}

	return nil
}

func (mr *moduleRepo) AddGrade(grade datastruct.GradeQuizzes) error {
	var submitted []datastruct.UserSubmittedAnswerQuizzes

	err := mr.db.Where("user_test_submission_quiz_id = ?", grade.UserTestSubmissionQuizID).Find(&submitted).Error
	if err != nil {
		fmt.Printf("ERROR: %v \n", err)
		return err
	}

	var score int
	var mark int
	var quizzesArr []datastruct.ChoiceQuizzes

	for _, ss := range submitted {
		choiceId := ss.ChoiceQuizID
		var choice datastruct.ChoiceQuizzes
		if err := mr.db.Where("id = ?", choiceId).Find(&choice).Error; err != nil {
			// Handle error or continue to next iteration
			continue
		}
		if choice.IsCorrect {
			score += choice.Weight
			mark += 1
		}
		quizzesArr = append(quizzesArr, choice)
	}

	grade.Score = score
	grade.Mark = mark

	if err := mr.db.Create(&grade).Error; err != nil {
		fmt.Printf("ERROR creating grade: %v \n", err)
		return err
	}
	return nil
}

func (mr *moduleRepo) GetTop3EverySubject(userID string, subjectTypeID string) ([]datastruct.SubModules, error) {
	var top3score []datastruct.SubModules

	// whereClause := fmt.Sprintf("subject_test_type_quizzes.id = %v AND grade_quizzes.user_id = %v", userID, subjectTypeID)

	err := mr.db.
		// Table("grade_quizzes").Table("user_test_submission_quizzes").Table("subject_test_type_quizzes").Table("subjects").Table("sub_modules").

		Preload("Subjects").
		Preload("Subjects.SubjectTestTypeQuizzes", "id = ?", subjectTypeID).
		Preload("Subjects.SubjectTestTypeQuizzes.UserTestSubmissionQuizzes", "user_id = ?", userID).
		Preload("Subjects.SubjectTestTypeQuizzes.UserTestSubmissionQuizzes.GradeQuiz").

		// Joins("JOIN subjects ON sub_modules.id = subjects.sub_module_id").
		// Joins("JOIN subject_test_type_quizzes ON subjects.id = subject_test_type_quizzes.subject_id").
		// Joins("JOIN user_test_submission_quizzes ON subject_test_type_quizzes.id = user_test_submission_quizzes.test_type_quiz_id").
		// Joins("JOIN grade_quizzes ON user_test_submission_quizzes.id = grade_quizzes.user_test_submission_quiz_id").
		// Where().

		Order("created_at DESC").
		Find(&top3score).Error

	if err != nil {
		return nil, err
	}
	return top3score, nil
}
