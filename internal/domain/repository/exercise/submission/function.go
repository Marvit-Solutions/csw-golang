package submission

import "csw-golang/internal/domain/entity/datastruct"

func (e exercisesSubmissionRepo) AddSubmission(submission datastruct.UserTestSubmissionExercises) (datastruct.UserTestSubmissionExercises, error) {
	tx := e.db.Create(&submission)

	if tx.Error != nil {
		return datastruct.UserTestSubmissionExercises{}, tx.Error
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
