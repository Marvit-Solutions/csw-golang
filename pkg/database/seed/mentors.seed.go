package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreateMentors() []*ds.Mentors {
	mentors := []*ds.Mentors{
		{
			ID:             "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Name:           "Ayu Lestari",
			SubModuleID:    "65994d9d-9728-435c-bc6a-92ed2ec2b8c8",
			Description:    "TWK lebih banyak logika, jadi pahamilah setiap pertanyaan dan jawabannya",
			ProfilePicture: "assets/img/users/profile/profile.png",
			Rating:         4.5,
		},
		{
			ID:             "2daa6bd9-f39f-1ca8-9e8e-e692f687e123",
			Name:           "Bambang Sugiyono",
			SubModuleID:    "9159575a-440e-4e88-abfb-ad19b6715022",
			Description:    "Tes Intelegensi Umum merupakan salah satu test SKD yang gampang banget",
			ProfilePicture: "assets/img/users/profile/profile.png",
			Rating:         4.2,
		},
		{
			ID:             "2daa6bd9-f19f-ca8-9e8e-e692f687e123",
			Name:           "Ayu Hafsah",
			SubModuleID:    "2fadda64-888e-45cf-b16b-d7e1e2a0a1f3",
			Description:    "TKP itu mudah, asik dan seru",
			ProfilePicture: "assets/img/users/profile/profile.png",
			Rating:         4.2,
		},
		{
			ID:             "1daa6bd9-f19f-ca8-9e8e-e692f687e123",
			Name:           "Megawati",
			SubModuleID:    "39dbb7c2-dbae-4877-8dca-709837f6ec46",
			Description:    "Matematika itu mudah, asik dan seru",
			ProfilePicture: "assets/img/users/profile/profile.png",
			Rating:         4.7,
		},
		{
			ID:             "2daa6bd9-f19f-ca8-9e8e-e692f687e111",
			Name:           "Zikri Hanafi",
			SubModuleID:    "39dbb7c2-dbae-4877-8dca-709837f6ec46",
			Description:    "Matematika itu mudah, asik dan seru”",
			ProfilePicture: "assets/img/users/profile/profile.png",
			Rating:         4.5,
		},
	}
	return mentors
}
