package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"gorm.io/gorm"
	"time"
)

func CreateQuestions() []*ds.QuestionExercises {
	QuestionsExercise := []*ds.QuestionExercises{
		{
			ID:                 "47669ec3-0483-4953-bd58-12b4166b9468",
			CreatedAt:          time.Time{},
			UpdatedAt:          time.Time{},
			TestTypeExerciseID: "9975116b-01a4-4978-90aa-a3a7525d504a",
			DeletedAt:          gorm.DeletedAt{},
			Content:            "Calculate the limit as x approaches 0 for (sin(x)/x) ?",
			Weight:             15,
			ChoiceExercises: []ds.ChoiceExercises{{
				ID:                           "90f91f8b-0a1d-4977-a7e8-09d27313cac3",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "0",
				IsCorrect:                    true,
				Weight:                       15,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "296f1912-c455-442c-9904-43917d258ebf",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "1",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "0f33c05a-f7c7-42bf-b9e5-a61bf6b8c3cd",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "Undefined",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "3c4cda6d-bb4f-4f97-963b-1e397a879200",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "Infinity",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "3c4cda6d-bb4f-4f97-963b-1e397a879200",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "phi",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}},
			UserSubmittedAnswerExercises: nil,
		},

		{
			ID:                 "e1d0f148-1863-4471-b8ab-90982fa70d97",
			CreatedAt:          time.Time{},
			UpdatedAt:          time.Time{},
			TestTypeExerciseID: "9975116b-01a4-4978-90aa-a3a7525d504a",
			DeletedAt:          gorm.DeletedAt{},
			Content:            "Ketika Dinda telah lulus SMA, ia punya pilihan untuk membuka usaha sendiri dengan penghasilan Rp 150.000,00/hari, bekerja di perusahaan orang tua dengan gaji per bulan Rp 5.500.000 atau bekerja di perusahaan teman dengan gaji Rp 6.000.000. Jika Dinda memilih membuka usaha sendiri, berapakah biaya peluang Dinda?",
			Weight:             15,
			ChoiceExercises: []ds.ChoiceExercises{{
				ID:                           "e8eb90c3-00fd-45e1-9f7e-bcd45f64d2f6",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "Rp 5.500.000",
				IsCorrect:                    true,
				Weight:                       15,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "1d2d5437-349e-489c-b0ce-6115a466393a",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "Rp 4.500.000",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "1f23c05a-f7c7-422f-b9e5-a61bf4b8c3c4",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "Rp 6.500.000",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "1c4cd26d-bb4f-4f91-963b-1e397a879201",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "Rp 150.000",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "334cda6d-1b4f-4f97-963b-1e397a81a300",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Content:                      "Rp 2.000.000",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}},
			UserSubmittedAnswerExercises: nil,
		},
	}

	return QuestionsExercise
}
