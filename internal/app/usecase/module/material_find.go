package module

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) MaterialFind(req request.Material) (*response.MaterialFindResponse, error) {
	result := &response.MaterialFindResponse{}

	if req.Subject && req.SubSubject {
		subjects, err := u.subjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find subjects: %v", err)
		}
		subSubjects, err := u.subSubjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find sub subjects: %v", err)
		}

		result.Subjects = make([]*response.SubjectFindResponse, len(subjects))
		for i, subject := range subjects {
			result.Subjects[i] = &response.SubjectFindResponse{
				ID:          subject.ID,
				UUID:        subject.UUID,
				Name:        subject.Name,
				LastUpdated: helper.FormatTimeToStringPtr(subject.UpdatedAt),
			}
		}

		result.SubSubjects = make([]*response.SubSubjectFindResponse, len(subSubjects))
		for i, subSubject := range subSubjects {
			result.SubSubjects[i] = &response.SubSubjectFindResponse{
				ID:          subSubject.ID,
				UUID:        subSubject.UUID,
				Name:        subSubject.Name,
				LastUpdated: helper.FormatTimeToStringPtr(subSubject.UpdatedAt),
			}
		}

		sort.Slice(result.Subjects, func(i, j int) bool {
			return result.Subjects[i].ID < result.Subjects[j].ID
		})

		sort.Slice(result.SubSubjects, func(i, j int) bool {
			return result.SubSubjects[i].ID < result.SubSubjects[j].ID
		})

	} else if req.Subject {
		subjects, err := u.subjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find subjects: %v", err)
		}

		result.Subjects = make([]*response.SubjectFindResponse, len(subjects))
		for i, subject := range subjects {
			result.Subjects[i] = &response.SubjectFindResponse{
				ID:          subject.ID,
				UUID:        subject.UUID,
				Name:        subject.Name,
				LastUpdated: helper.FormatTimeToStringPtr(subject.UpdatedAt),
			}
		}

		sort.Slice(result.Subjects, func(i, j int) bool {
			return result.Subjects[i].ID < result.Subjects[j].ID
		})

	} else if req.SubSubject {
		subSubjects, err := u.subSubjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find sub subjects: %v", err)
		}

		result.SubSubjects = make([]*response.SubSubjectFindResponse, len(subSubjects))
		for i, subSubject := range subSubjects {
			result.SubSubjects[i] = &response.SubSubjectFindResponse{
				ID:          subSubject.ID,
				UUID:        subSubject.UUID,
				Name:        subSubject.Name,
				LastUpdated: helper.FormatTimeToStringPtr(subSubject.UpdatedAt),
			}
		}

		sort.Slice(result.SubSubjects, func(i, j int) bool {
			return result.SubSubjects[i].ID < result.SubSubjects[j].ID
		})
	}

	return result, nil
}
