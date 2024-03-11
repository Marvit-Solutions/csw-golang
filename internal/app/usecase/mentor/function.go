package mentor

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (mr *mentorUsecase) GetListTopThreeMentors() (response.ListMentor, error) {
	data, err := mr.mentorRepo.GetListTopThreeMentors()
	if err != nil {
		return response.ListMentor{}, err
	}
	return data, nil
}

func (mr *mentorUsecase) GetAllMentors() (response.ListMentor, error) {
	data, err := mr.mentorRepo.GetAllMentors()
	if err != nil {
		return response.ListMentor{}, err
	}
	return data, nil
}
