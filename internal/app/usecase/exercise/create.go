package exercise

import (
	"fmt"
	"time"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Create(req request.ExerciseCreateRequest) error {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"uuid":    req.AuthenticatedUser,
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

	exerciseQuestions, err := u.exerciseQuestionRepo.FindBy(map[string]interface{}{
		"exercise_id": exercise.ID,
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
	userMarkedMap := make(map[int]bool)
	for i, choice := range req.Answers {
		userAnswerMap[i+1] = choice.ChoiceUUID
		userMarkedMap[i+1] = choice.IsMarked
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

	newSubmission := &model.ExerciseSubmission{
		UserID:       user.ID,
		ExerciseID:   exercise.ID,
		StartedAt:    time.Now().Add(-helper.ParseTimeString(req.TimeRequired)),
		FinishedAt:   time.Now(),
		TimeRequired: req.TimeRequired,
		RightAnswer:  rightAnswers,
		Score:        score,
	}

	tx := u.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(newSubmission).Error; err != nil {
		return fmt.Errorf("failed to create submission: %v", err)
	}

	var newExerciseAnswers []*model.ExerciseAnswer
	for i, answerUUID := range userAnswerMap {
		newExerciseAnswers = append(newExerciseAnswers, &model.ExerciseAnswer{
			SubmissionID: newSubmission.ID,
			ChoiceID:     exerciseAnswerMap[answerUUID],
			IsMarked:     userMarkedMap[i],
		})
	}

	if err := tx.CreateInBatches(newExerciseAnswers, 100).Error; err != nil {
		return fmt.Errorf("failed to create exercise answers: %v", err)
	}

	tx.Commit()
	return nil
}
