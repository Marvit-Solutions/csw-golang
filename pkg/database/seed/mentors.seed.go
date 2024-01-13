package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreateMentors() []*ds.Mentor {
	mentors := []*ds.Mentor{
		{
			Id:           "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Name:         "Ayu Lestari",
			Description:  "Matematika itu mudah, asik dan seru",
			ProfilePhoto: "assets/img/users/profile/profile.png",
			Rating:       4.5,
		},
		{
			Id:           "2daa6bd9-f39f-1ca8-9e8e-e692f687e123",
			Name:         "Bambang Sugiyono",
			Description:  "Bahasa Indonesia itu mudah, asik dan seru",
			ProfilePhoto: "assets/img/users/profile/profile.png",
			Rating:       4.2,
		},
		{
			Id:           "2daa6bd9-f19f-ca8-9e8e-e692f687e123",
			Name:         "Ayu Hafsah",
			Description:  "Sejarah itu mudah, asik dan seru",
			ProfilePhoto: "assets/img/users/profile/profile.png",
			Rating:       4.2,
		},
		{
			Id:           "1daa6bd9-f19f-ca8-9e8e-e692f687e123",
			Name:         "Megawati",
			Description:  "Geografis itu mudah, asik dan seru",
			ProfilePhoto: "assets/img/users/profile/profile.png",
			Rating:       4.7,
		},
		{
			Id:           "2daa6bd9-f19f-ca8-9e8e-e692f687e111",
			Name:         "Zikri Hanafi",
			Description:  "Bahasa Jepang itu mudah, asik dan seru",
			ProfilePhoto: "assets/img/users/profile/profile.png",
			Rating:       4.5,
		},
	}
	return mentors
}
