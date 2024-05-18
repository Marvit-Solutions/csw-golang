package request

type ParamExercise struct {
	Module   string `form:"module"`
	TestType string `form:"test_type"`
}

type ExerciseDetailRequest struct {
	ExerciseUUID string `uri:"exercise_uuid"`
}
