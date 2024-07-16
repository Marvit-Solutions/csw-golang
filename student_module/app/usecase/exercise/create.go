package exercise

import (
	"fmt"
	"time"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
)

func (u *usecase) Create(req request.ExerciseCreateRequest) error {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id":      req.AuthenticatedUser,
		"role_id": helper.PembeliPaketBimbel,
	})
	if user == nil {
		return helper.ErrAccessDenied
	}
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	exercise, err := u.exerciseRepo.FindOneBy(map[string]interface{}{
		"uuid": req.ExerciseUUID,
	})
	if err != nil {
		return fmt.Errorf("failed to find exercise: %v", err)
	}

	subModuleIDs, err := u.exerciseLocalRepo.FindSubModulesID()
	if err != nil {
		return fmt.Errorf("failed to find exercise: %v", err)
	}

	totalScoreSubmisssionModule := 0
	totalRightAnswerSubmisssionModule := 0

	// temporary data
	newSubmissionModuleBefore := &model.ExerciseSubmissionsModule{
		UserID:       user.ID,
		ExerciseID:   exercise.ID,
		RightAnswer:  totalRightAnswerSubmisssionModule, //tmp
		Score:        totalScoreSubmisssionModule,       //tmp
		StartedAt:    time.Now().Add(-helper.ParseTimeString(req.TimeRequired)),
		FinishedAt:   time.Now(),
		TimeRequired: req.TimeRequired,
	}

	tx := u.db.Begin()
	defer tx.Rollback()

	submissionModule, err := u.exerciseSubmissionsModuleRepo.Create(newSubmissionModuleBefore, tx)

	if err != nil {
		return fmt.Errorf("failed to create quiz submission: %v", err)
	}

	for _, subModuleID := range subModuleIDs {

		exerciseQuestions, err := u.exerciseQuestionRepo.FindBy(map[string]interface{}{
			"exercise_id":   exercise.ID,
			"sub_module_id": subModuleID,
		}, 0, 0)
		if err != nil {
			return fmt.Errorf("failed to find exercise questions: %v", err)
		}

		exerciseQuestionMap := make(map[int]int)
		exerciseQuestionIDs := make([]int, 0, len(exerciseQuestions))
		for _, question := range exerciseQuestions {
			exerciseQuestionMap[question.ID] = question.Score
			exerciseQuestionIDs = append(exerciseQuestionIDs, question.ID)
		}

		rightExerciseChoices, err := u.exerciseChoiceRepo.FindBy(map[string]interface{}{
			"question_id": exerciseQuestionIDs,
			"is_correct":  true,
		}, 0, 0)
		if err != nil {
			return fmt.Errorf("failed to find exercise choices: %v", err)
		}

		userAnswerMap := make(map[int]string)
		for i, choice := range req.Answers {
			if choice.SubModuleID == subModuleID {
				userAnswerMap[i+1] = choice.ChoiceUUID
			}
		}

		userAnswerUUIDs := make([]string, 0, len(userAnswerMap))
		for _, UUID := range userAnswerMap {
			if UUID != "" {
				userAnswerUUIDs = append(userAnswerUUIDs, UUID)
			}
		}

		exerciseAnswers, err := u.exerciseChoiceRepo.FindBy(map[string]interface{}{
			"uuid": userAnswerUUIDs,
		}, 0, 0)
		if err != nil {
			return fmt.Errorf("failed to find exercise answers: %v", err)
		}

		exerciseAnswerMap := make(map[string]*int)
		for _, answer := range exerciseAnswers {
			exerciseAnswerMap[answer.UUID] = &answer.ID
		}

		rightAnswerMap := make(map[int]string)
		for _, answer := range rightExerciseChoices {
			rightAnswerMap[answer.QuestionID] = answer.UUID
		}

		var score, rightAnswers int
		for questionID, userUUID := range userAnswerMap {
			if rightUUID, exists := rightAnswerMap[questionID]; exists && userUUID == rightUUID {
				score += exerciseQuestionMap[questionID]
				rightAnswers++
			}
		}

		newSubmissionSubModule := &model.ExerciseSubmissionsSubModule{
			UserID:              user.ID,
			ExerciseID:          exercise.ID,
			SubModuleID:         subModuleID,
			SubmissionsModuleID: submissionModule.ID,
			RightAnswer:         rightAnswers,
			Score:               score,
		}

		if err := tx.Create(newSubmissionSubModule).Error; err != nil {
			return fmt.Errorf("failed to create newSubmissionSubModule: %v", err)
		}

		var newExerciseAnswers []*model.ExerciseAnswer
		for _, answerUUID := range userAnswerMap {
			newExerciseAnswers = append(newExerciseAnswers, &model.ExerciseAnswer{
				SubmissionSubModuleID: newSubmissionSubModule.ID,
				ChoiceID:              exerciseAnswerMap[answerUUID],
			})
		}

		if err := tx.Create(newExerciseAnswers).Error; err != nil {
			return fmt.Errorf("failed to create exercise answers: %v", err)
		}

		totalScoreSubmisssionModule += score
		totalRightAnswerSubmisssionModule += rightAnswers

		fmt.Printf("submoudule id %d\n", subModuleID)
		for _, answer := range newExerciseAnswers {
			fmt.Printf("ID: %d\n", answer.ID)
			fmt.Printf("UUID: %s\n", answer.UUID)
			fmt.Printf("SubmissionID: %d\n", answer.SubmissionSubModuleID)
			if answer.ChoiceID != nil {
				fmt.Printf("ChoiceID: %d\n", *answer.ChoiceID)
			} else {
				fmt.Printf("ChoiceID: nil\n")
			}
			fmt.Printf("CreatedAt: %s\n", answer.CreatedAt)
			fmt.Printf("UpdatedAt: %s\n", answer.UpdatedAt)
			fmt.Printf("DeletedAt: %v\n", answer.DeletedAt)
			fmt.Println("-----")
		}

		fmt.Printf("((((((((((((((((((((((((((((((((((()))))))))))))))))))))))))))))))))))")

	}

	fmt.Println("ini req.TimeRequired")
	fmt.Println(req.TimeRequired)
	newSubmissionModuleAfter := &model.ExerciseSubmissionsModule{
		ID:           submissionModule.ID,
		UUID:         submissionModule.UUID,
		UserID:       user.ID,
		ExerciseID:   exercise.ID,
		RightAnswer:  totalRightAnswerSubmisssionModule, //tmp
		Score:        totalScoreSubmisssionModule,       //tmp
		StartedAt:    time.Now().Add(-helper.ParseTimeString(req.TimeRequired)),
		FinishedAt:   time.Now(),
		TimeRequired: req.TimeRequired,
	}

	err = u.exerciseSubmissionsModuleRepo.Update(newSubmissionModuleAfter, tx)

	if err != nil {
		return fmt.Errorf("failed to create quiz submission: %v", err)
	}

	tx.Commit()

	return nil
}
