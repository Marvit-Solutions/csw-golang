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

	fmt.Println("SUBMISSION ID: ", grade.UserTestSubmissionQuizID)
	err := mr.db.Where("user_test_submission_quiz_id = ?", grade.UserTestSubmissionQuizID).Find(&submitted).Error
	if err != nil {
		fmt.Printf("ERROR: %v \n", err)
		return err
	}

	fmt.Printf("SUBMITTED: %v \n", submitted)

	var score int
	var mark int
	var quizzesArr []datastruct.ChoiceQuizzes

	for _, ss := range submitted {
		choiceId := ss.ChoiceQuizID
		var choice datastruct.ChoiceQuizzes
		if err := mr.db.Where("id = ?", choiceId).Find(&choice).Error; err != nil {
			fmt.Printf("ERROR retrieving choice: %v \n", err)
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

func (mr *moduleRepo) GetTop3EverySubject(userID string) ([]datastruct.SubModules, error) {
	var top3score []datastruct.SubModules

	whereClause := fmt.Sprintf("user_id = %v", userID)
	fmt.Println("WHERE CLAUSE: ", whereClause)

	err := mr.db.Preload("Subjects").Preload("Subjects.SubjectTestTypeQuizzes", "id = 'f8259b82-41ca-4d58-90e7-c2be3f4c235f' ").Preload("Subjects.SubjectTestTypeQuizzes.UserTestSubmissionQuizzes", "user_id = ?", userID).Preload("Subjects.SubjectTestTypeQuizzes.UserTestSubmissionQuizzes.GradeQuiz").Order("created_at DESC").Limit(3).Find(&top3score).Error
	if err != nil {
		return nil, err
	}
	return top3score, nil
}
