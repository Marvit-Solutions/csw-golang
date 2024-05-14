package response

type ExerciseResponse struct {
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	TestType      string `json:"test_type"`
	ModuleName    string `json:"module_name"`
	Title         string `json:"title"`
	TotalQuestion int    `json:"total_question"`
	Time          int    `json:"time"`
	Description   string `json:"description"`
}
