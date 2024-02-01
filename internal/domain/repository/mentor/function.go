package mentor

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (m mentorRepo) GetListTopThreeMentors() (error, dto.ListMentor) {
	var topMentors dto.ListMentor
	var mentors []datastruct.Mentors
	_ = m.db.Order("rating desc, name").Limit(3).Find(&mentors)

	// Convert the database mentors to DTO mentors
	for _, mentor := range mentors {
		dtoMentor := struct {
			ID             string  `json:"ID"`
			Name           string  `json:"name"`
			Description    string  `json:"description"`
			ProfilePicture string  `json:"profile_photo"`
			Rating         float32 `json:"rating"`
		}{
			ID:             mentor.ID,
			Name:           mentor.Name,
			Description:    mentor.Description,
			ProfilePicture: mentor.ProfilePicture,
			Rating:         mentor.Rating,
			// Add other fields as needed
		}
		topMentors = append(topMentors, dtoMentor)
	}
	return nil, topMentors
}

func (m mentorRepo) GetAllMentors() (error, dto.ListMentor) {
	var allMentors dto.ListMentor
	var mentors []datastruct.Mentors

	// Fetch all mentors from the database
	tx := m.db.Find(&mentors)
	if tx.Error != nil {
		return tx.Error, allMentors
	}

	for _, mentor := range mentors {
		dtoMentor := struct {
			ID             string  `json:"ID"`
			Name           string  `json:"name"`
			Description    string  `json:"description"`
			ProfilePicture string  `json:"profile_photo"`
			Rating         float32 `json:"rating"`
		}{
			ID:             mentor.ID,
			Name:           mentor.Name,
			Description:    mentor.Description,
			ProfilePicture: mentor.ProfilePicture,
			Rating:         mentor.Rating,
			// Add other fields as needed
		}
		allMentors = append(allMentors, dtoMentor)
	}

	return nil, allMentors
}
