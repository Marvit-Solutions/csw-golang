package module

import (
	module "csw-golang/internal/domain/usecase/module"
)

type ModuleHandler struct {
	moduleUsecase module.ModuleUsecase
}

func New(moduleUsecase module.ModuleUsecase) *ModuleHandler {
	return &ModuleHandler{
		moduleUsecase,
	}
}
