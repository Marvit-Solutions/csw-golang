package submission

import "csw-golang/internal/domain/entity/datastruct"

func (e exercisesSubmissionRepo) AddSubmission(submission datastruct.UserTestSubmissionExercises, answer []datastruct.UserSubmittedAnswerExercises, grade datastruct.GradeExercises) (datastruct.UserTestSubmissionExercises, error) {

	// Start a transaction
	tx := e.db.Begin()

	// Defer a rollback in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tx = tx.Create(&submission)
	if tx.Error != nil {
		tx.Rollback()
		return datastruct.UserTestSubmissionExercises{}, tx.Error
	}

	tx = tx.CreateInBatches(&answer, 150)
	if tx.Error != nil {
		tx.Rollback()
		return datastruct.UserTestSubmissionExercises{}, tx.Error
	}
	//
	//var score int
	//var exercisesArr []datastruct.ChoiceExercises
	//for _, ss := range answer {
	//	choiceId := ss.ChoiceExerciseID
	//	var choice datastruct.ChoiceExercises
	//	_ = e.db.Where("id = ?", choiceId).Find(&choice)
	//	if choice.IsCorrect == true {
	//		score += choice.Weight
	//	}
	//	exercisesArr = append(exercisesArr, choice)
	//}
	//grade.Score = score
	//
	//tx = tx.Create(&grade)
	//if tx.Error != nil {
	//	tx.Rollback()
	//	return datastruct.UserTestSubmissionExercises{}, tx.Error
	//}

	//
	//// Calculate score and mark
	//var score int
	//var mark int
	//var quizzesArr []datastruct.ChoiceExercises
	//for _, ss := range answer {
	//	choiceID := ss.ChoiceExerciseID
	//	var choice datastruct.ChoiceExercises
	//	if err := tx.Where("id = ?", choiceID).Find(&choice).Error; err != nil {
	//		continue
	//	}
	//	if choice.IsCorrect {
	//		score += choice.Weight
	//		mark++
	//	}
	//	quizzesArr = append(quizzesArr, choice)
	//}
	//grade.Score = score

	//// Add grade
	//if err := tx.Create(&grade).Error; err != nil {
	//	tx.Rollback()
	//	return datastruct.UserTestSubmissionExercises{}, err
	//}

	// Commit transaction
	err := tx.Commit()
	if err.Error != nil {
		tx.Rollback()
		return datastruct.UserTestSubmissionExercises{}, err.Error
	}

	return submission, nil

}

func (e exercisesSubmissionRepo) GetSubmission(submissionId string) (datastruct.UserTestSubmissionExercises, error) {
	var exercises datastruct.UserTestSubmissionExercises

	tx := e.db.Where("id = ?", submissionId).First(&exercises)
	if tx.Error != nil {
		return datastruct.UserTestSubmissionExercises{}, tx.Error
	}
	return exercises, nil
}

func (e exercisesSubmissionRepo) GetAllSubmissions() ([]datastruct.UserTestSubmissionExercises, error) {
	var exercises []datastruct.UserTestSubmissionExercises

	tx := e.db.Find(&exercises)
	if tx.Error != nil {
		return []datastruct.UserTestSubmissionExercises{}, tx.Error
	}
	return exercises, nil
}
