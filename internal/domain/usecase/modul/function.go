package module

import (
	"csw-golang/internal/domain/entity/dto"
	"fmt"
)

func (mod *moduleUsecase) GetListModules() (*[]dto.ModuleResponse, error) {
	fmt.Println("GetListModules usecase")
	modules, err := mod.moduleRepo.GetListModules()
	if err != nil {
		return nil, err
	}
	print("ListModule resultes")
	return modules, nil
}
