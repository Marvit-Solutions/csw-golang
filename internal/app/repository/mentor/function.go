package mentor

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/datastruct"
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (m mentorRepository) GetListTopThreeMentors() (dto.ListMentor, error) {
	var data dto.ListMentor
	var mentors []datastruct.Mentors
	_ = m.db.Order("rating desc, name").Limit(3).Find(&mentors)

	// Convert the database mentors to DTO mentors
	for _, mentor := range mentors {
		dtoMentor := struct {
			ID             string  `json:"ID"`
			Name           string  `json:"Name"`
			Description    string  `json:"Description"`
			ProfilePicture string  `json:"ProfilePicture"`
			Rating         float32 `json:"Rating"`
		}{
			ID:             mentor.ID,
			Name:           mentor.Name,
			Description:    mentor.Description,
			ProfilePicture: mentor.ProfilePicture,
			Rating:         mentor.Rating,
			// Add other fields as needed
		}
		data = append(data, dtoMentor)
	}
	return data, nil
}

func (m mentorRepository) GetAllMentors() (dto.ListMentor, error) {
	var data dto.ListMentor
	var mentors []datastruct.Mentors

	// Fetch all mentors from the database
	tx := m.db.Find(&mentors)
	if tx.Error != nil {
		return data, fmt.Errorf("failed to fetch mentors: %v", tx.Error)
	}

	for _, mentor := range mentors {
		dtoMentor := struct {
			ID             string  `json:"ID"`
			Name           string  `json:"Name"`
			Description    string  `json:"Description"`
			ProfilePicture string  `json:"ProfilePicture"`
			Rating         float32 `json:"Rating"`
		}{
			ID:             mentor.ID,
			Name:           mentor.Name,
			Description:    mentor.Description,
			ProfilePicture: mentor.ProfilePicture,
			Rating:         mentor.Rating,
			// Add other fields as needed
		}
		data = append(data, dtoMentor)
	}

	return data, nil
}
