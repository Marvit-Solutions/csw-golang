package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateUserDetails() []*ds.UserDetails {
	testimonial1 := "dc38d81f-3c96-46b1-b985-e6f91eb01cb7"
	testimonial2 := "d7209984-1326-4872-ac47-09ea7c2dd1a8"
	testimonial3 := "e8432fb3-e5b1-4568-8aa1-8cdbf185f948"
	testimonial4 := "3eccb5cd-a010-4fd2-8884-bdc1b0f6226e"
	testimonial5 := "39e8d413-2e02-4ccd-b176-201ff9012c5c"
	testimonial6 := "29d9990a-caa6-49ef-a261-fdca7a828b0e"
	userDetails := []*ds.UserDetails{
		{
			ID:             "e36814fe-d479-4c5e-af49-4dede000f332",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "Administrator",
			Status:         "Admin",
			PhoneNumber:    "08917283129283",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "6c8bcb83-a825-4df0-8c29-264402205b9b",
		},
		{
			ID:             "ecf50842-f521-4291-a31f-9f10c48fd889",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "User 1",
			Status:         "Mahasiswa STIS 2020",
			PhoneNumber:    "08917283109283",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			TestimonialID:  &testimonial2,
		},
		{
			ID:             "a9131226-069a-43f9-b52d-d1c5bfc1b878",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "User 2",
			Status:         "Mahasiswa STIS 2022",
			PhoneNumber:    "0851728392716",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			TestimonialID:  &testimonial3,
		},
		{
			ID:             "931fce0e-83c9-422a-ad44-f4f11b8e51e9",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "User 3",
			Status:         "Alumni",
			PhoneNumber:    "0895475321896",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e",
			TestimonialID:  &testimonial4,
		},
		{
			ID:             "48e94d1e-217d-4f27-89d4-031bd44ee555",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "User 4",
			Status:         "Alumni",
			PhoneNumber:    "082546587612",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "8f23e5cd-0e96-4e1d-a72b-b8f85ca083ee",
			TestimonialID:  &testimonial5,
		},
		{
			ID:             "891c0769-e2df-4cbc-9bc9-2dc7344d0914",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "User 5",
			Status:         "Mahasiswa UGM 2020",
			PhoneNumber:    "0856487532158",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "fca8ef53-9f39-4cd6-940d-5df58110da59",
			TestimonialID:  &testimonial6,
		},
		{
			ID:             "236ef707-8fe1-4cc9-876c-2b9f6e49b5a8",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "User 6",
			Status:         "Kelas 11",
			PhoneNumber:    "08865486523215",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "66647904-fdee-4201-a4c7-fc83f5b136f7",
			TestimonialID:  &testimonial1,
		},
		{
			ID:             "612e746f-e0b9-4972-8f77-1ccda6306b6a",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           "User 7",
			Status:         "Kelas 12",
			PhoneNumber:    "082356987546",
			ProfilePicture: "assets/img/users/profile/profile.png",
			UserID:         "613db4f7-7fbd-4b82-97ed-1802368801b9",
		},
	}

	return userDetails
}
