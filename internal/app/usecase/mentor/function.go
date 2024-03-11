package mentor

import (
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (mr *mentorUsecase) GetListTopThreeMentors() (dto.ListMentor, error) {
	data, err := mr.mentorRepo.GetListTopThreeMentors()
	if err != nil {
		return dto.ListMentor{}, err
	}
	return data, nil
}

func (mr *mentorUsecase) GetAllMentors() (dto.ListMentor, error) {
	data, err := mr.mentorRepo.GetAllMentors()
	if err != nil {
		return dto.ListMentor{}, err
	}
	return data, nil
}
