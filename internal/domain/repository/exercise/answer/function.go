package answer

import "csw-golang/internal/domain/entity/datastruct"

func (e exercisesAnswerRepo) AddAnswer(answer []datastruct.UserSubmittedAnswerExercises) ([]datastruct.UserSubmittedAnswerExercises, error) {
	tx := e.db.CreateInBatches(&answer, 150)

	if tx.Error != nil {
		return []datastruct.UserSubmittedAnswerExercises{}, tx.Error
	}
	return answer, nil

}
