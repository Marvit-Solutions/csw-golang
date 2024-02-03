package module

import (
	"csw-golang/internal/domain/entity/dto"
	module "csw-golang/internal/domain/repository/modul"
)

type ModuleUsecase interface {
	GetListModules() (*[]dto.ModuleResponse, error)
}

type moduleUsecase struct {
	moduleRepo module.ModuleRepo
}

func New(moduleRepo module.ModuleRepo) ModuleUsecase {
	return &moduleUsecase{
		moduleRepo,
	}
}
