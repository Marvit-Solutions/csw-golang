package modul

import (
	mc "csw-golang/internal/domain/usecase/modul"
)

type ModulHandler struct {
	modulUsecase mc.ModulUsecase
}

func New(modulUsecase mc.ModulUsecase) *ModulHandler {
	return &ModulHandler{
		modulUsecase,
	}
}
