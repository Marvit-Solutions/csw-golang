package mentor

import (
	"csw-golang/internal/domain/entity/dto"
)

func (mr *mentorUsecase) GetListTopThreeMentors() (error, dto.ListMentor) {
	err, mentor := mr.mentorRepo.GetListTopThreeMentors()
	if err != nil {
		return err, dto.ListMentor{}
	}
	return nil, mentor
}
