package exercise

import (
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
	"gorm.io/gorm"
)

type Usecase interface {
	FindAll(req request.ParamExercise) ([]*response.Exercise, error)
	FindDetail(req request.ExerciseDetailRequest) (*response.ExerciseDetail, error)
	FindHistory(req request.ExerciseHistory) ([]*response.ExerciseHistory, error)
	Create(req request.ExerciseCreateRequest) error
	Review(req request.ExerciseReview) (*response.ExerciseReview, error)
}

type usecase struct {
	db                        *gorm.DB
	testTypeRepo              repository.TestTypeRepository
	moduleRepo                repository.ModuleRepository
	mediaRepo                 repository.MediaRepository
	exerciseRepo              repository.ExerciseRepository
	exerciseAnswerRepo        repository.ExerciseAnswerRepository
	exerciseSubmissionRepo    repository.ExerciseSubmissionRepository
	exerciseQuestionRepo      repository.ExerciseQuestionRepository
	exerciseChoiceRepo        repository.ExerciseChoiceRepository
	exerciseQuestionMediaRepo repository.ExerciseQuestionMediaRepository
	userRepo                  repository.UserRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                        db,
		testTypeRepo:              service.NewTestTypeService(db, nil),
		moduleRepo:                service.NewModuleService(db, nil),
		mediaRepo:                 service.NewMediaService(db, nil),
		exerciseRepo:              service.NewExerciseService(db, nil),
		exerciseAnswerRepo:        service.NewExerciseAnswerService(db, nil),
		exerciseSubmissionRepo:    service.NewExerciseSubmissionService(db, nil),
		exerciseQuestionRepo:      service.NewExerciseQuestionService(db, nil),
		exerciseChoiceRepo:        service.NewExerciseChoiceService(db, nil),
		exerciseQuestionMediaRepo: service.NewExerciseQuestionMediaService(db, nil),
		userRepo:                  service.NewUserService(db, nil),
	}
}
