package exercise

import (
	"fmt"
	"time"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
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
		return fmt.Errorf("failed to create exercise submission: %v", err)
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

		exerciseChoices, err := u.exerciseChoiceRepo.FindBy(map[string]interface{}{
			"question_id": exerciseQuestionIDs,
		}, 0, 0)
		if err != nil {
			return fmt.Errorf("failed to find exercise choices: %v", err)
		}

		mapChoices := make(map[int][]response.OptionItemSubmission)

		for _, exerciseChoice := range exerciseChoices {
			exerciseChoiceResponse := &response.OptionItemSubmission{
				ID:        exerciseChoice.ID,
				UUID:      exerciseChoice.UUID,
				IsCorrect: exerciseChoice.IsCorrect,
			}
			mapChoices[exerciseChoice.QuestionID] = append(mapChoices[exerciseChoice.QuestionID], *exerciseChoiceResponse)
		}

		totalRightAnswer := 0
		score := 0

		for _, question := range req.Questions {
			if question.SubModuleID == subModuleID {
				var rightAnswer string
				for _, option := range mapChoices[question.ID] {
					if option.IsCorrect {
						rightAnswer = option.UUID
					}
				}
				if question.UserAnswer == rightAnswer {
					totalRightAnswer++
					score = score + question.Score
				}
			}
		}

		newSubmissionSubModule := &model.ExerciseSubmissionsSubModule{
			UserID:              user.ID,
			ExerciseID:          exercise.ID,
			SubModuleID:         subModuleID,
			SubmissionsModuleID: submissionModule.ID,
			RightAnswer:         totalRightAnswer,
			Score:               score,
		}

		if err := tx.Create(newSubmissionSubModule).Error; err != nil {
			return fmt.Errorf("failed to create newSubmissionSubModule: %v", err)
		}

		// insert user exercise answer and start calculate the score
		for _, ques := range req.Questions {
			if ques.SubModuleID == subModuleID {
				if ques.UserAnswer != "" {
					exerciseChoice, err := u.exerciseChoiceRepo.FindOneBy(map[string]interface{}{
						"uuid": &ques.UserAnswer,
					})
					if err != nil {
						return fmt.Errorf("failed to find exercise answers: %v", err)
					}
					exerciseAnswer := &model.ExerciseAnswer{
						SubmissionSubModuleID: newSubmissionSubModule.ID,
						ChoiceID:              &exerciseChoice.ID,
					}
					_, err = u.exerciseAnswerRepo.Create(exerciseAnswer, tx)
					if err != nil {
						return fmt.Errorf("failed to create exercise answer: %v", err)
					}
				}
			}

		}

		// Mencetak newExerciseAnswers
		// for i, answer := range newExerciseAnswers {
		// 	fmt.Printf("Answer %d: %+v\n", i+1, *answer)
		// 	if answer.ChoiceID != nil {
		// 		fmt.Printf("ChoiceID: %d\n", *answer.ChoiceID)
		// 	} else {
		// 		fmt.Printf("ChoiceID: nil\n")
		// 	}
		// }

		// if err := tx.Create(newExerciseAnswers).Error; err != nil {
		// 	return fmt.Errorf("failed to create exercise answers: %v", err)
		// }

		totalScoreSubmisssionModule += score
		totalRightAnswerSubmisssionModule += totalRightAnswer

		// fmt.Printf("submoudule id %d\n", subModuleID)
		// for _, answer := range newExerciseAnswers {
		// 	fmt.Printf("ID: %d\n", answer.ID)
		// 	fmt.Printf("UUID: %s\n", answer.UUID)
		// 	fmt.Printf("SubmissionID: %d\n", answer.SubmissionSubModuleID)
		// 	if answer.ChoiceID != nil {
		// 		fmt.Printf("ChoiceID: %d\n", *answer.ChoiceID)
		// 	} else {
		// 		fmt.Printf("ChoiceID: nil\n")
		// 	}
		// 	fmt.Printf("CreatedAt: %s\n", answer.CreatedAt)
		// 	fmt.Printf("UpdatedAt: %s\n", answer.UpdatedAt)
		// 	fmt.Printf("DeletedAt: %v\n", answer.DeletedAt)
		// 	fmt.Println("-----")
		// }

		// fmt.Printf("((((((((((((((((((((((((((((((((((()))))))))))))))))))))))))))))))))))")

	}

	// fmt.Println("ini req.TimeRequired")
	// fmt.Println(req.TimeRequired)
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
		return fmt.Errorf("failed to create exercise submission: %v", err)
	}

	tx.Commit()

	return nil
}
