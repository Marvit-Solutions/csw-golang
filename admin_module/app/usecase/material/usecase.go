package materialAdmin

import (
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localrepository"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localservice"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	MaterialAdminAll(req request.ParamMaterialAdminAll) ([]*response.MaterialAdminAllResponse, int, int, error)
	Create(req request.MaterialAdminPayload) error
	Update(req request.MaterialAdminPayloadUpdate) error
	Delete(uuid string) error
	MaterialAdminUpdateDetail(req request.ParamMaterialAdminUpdateDetail) (*response.MaterialAdminUpdateDetailResponse, error)
}

type usecase struct {
	db                     *gorm.DB
	materialRepo           repository.QuizRepository
	subjectRepo            repository.SubjectRepository
	subSubjectRepo         repository.SubSubjectRepository
	testTypeRepo           repository.TestTypeRepository
	materialQuestionRepo   repository.QuizQuestionRepository
	materialChoiceRepo     repository.QuizChoiceRepository
	materialAnswerRepo     repository.QuizAnswerRepository
	materialSubmissionRepo repository.QuizSubmissionRepository
	materialAdminLocalRepo localrepository.MaterialAdmin
	subModulRepo           repository.SubModuleRepository
	modulRepo              repository.ModuleRepository
	userRepo               repository.UserRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                     db,
		materialRepo:           service.NewQuizService(db, nil),
		subjectRepo:            service.NewSubjectService(db, nil),
		subSubjectRepo:         service.NewSubSubjectService(db, nil),
		testTypeRepo:           service.NewTestTypeService(db, nil),
		materialQuestionRepo:   service.NewQuizQuestionService(db, nil),
		materialChoiceRepo:     service.NewQuizChoiceService(db, nil),
		materialAnswerRepo:     service.NewQuizAnswerService(db, nil),
		materialSubmissionRepo: service.NewQuizSubmissionService(db, nil),
		materialAdminLocalRepo: localservice.NewMaterialAdminService(db),
		subModulRepo:           service.NewSubModuleService(db, nil),
		modulRepo:              service.NewModuleService(db, nil),
		userRepo:               service.NewUserService(db, nil),
	}
}
