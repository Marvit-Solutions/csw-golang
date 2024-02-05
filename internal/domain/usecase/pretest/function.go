package pretest

import (
	"csw-golang/internal/domain/entity/dto"
)

func (pr *pretestUsecase) GetAllPretests() (error, []dto.GetAllPretestResponse) {
	err, pretest := pr.pretestRepo.GetAllPretests()
	if err != nil {
		return err, []dto.GetAllPretestResponse{}
	}
	return nil, pretest
}
