package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateUserDetails() []*ds.UserDetails {
	userDetails := []*ds.UserDetails{
		{
			ID:          "e36814fe-d479-4c5e-af49-4dede000f332",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "Administrator",
			Class:       "Admin",
			PhoneNumber: "0895645875321",
			UserID:      "6c8bcb83-a825-4df0-8c29-264402205b9b",
		},
		{
			ID:          "ecf50842-f521-4291-a31f-9f10c48fd889",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "User 1",
			Class:       "Alumni",
			PhoneNumber: "082547865483",
			UserID:      "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		},
		{
			ID:          "a9131226-069a-43f9-b52d-d1c5bfc1b878",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "User 2",
			Class:       "Alumni",
			PhoneNumber: "0845654875321",
			UserID:      "41fb3d71-33bc-4a6e-9620-2d56f3090981",
		},
		{
			ID:          "931fce0e-83c9-422a-ad44-f4f11b8e51e9",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "User 3",
			Class:       "Alumni",
			PhoneNumber: "0885654878987",
			UserID:      "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e",
		},
		{
			ID:          "48e94d1e-217d-4f27-89d4-031bd44ee555",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "User 4",
			Class:       "Alumni",
			PhoneNumber: "085456875321",
			UserID:      "8f23e5cd-0e96-4e1d-a72b-b8f85ca083ee",
		},
		{
			ID:          "891c0769-e2df-4cbc-9bc9-2dc7344d0914",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "User 5",
			Class:       "Alumni",
			PhoneNumber: "0856452135487",
			UserID:      "fca8ef53-9f39-4cd6-940d-5df58110da59",
		},
		{
			ID:          "236ef707-8fe1-4cc9-876c-2b9f6e49b5a8",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "User 6",
			Class:       "Kelas 11",
			PhoneNumber: "083215478965",
			UserID:      "66647904-fdee-4201-a4c7-fc83f5b136f7",
		},
		{
			ID:          "612e746f-e0b9-4972-8f77-1ccda6306b6a",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        "User 7",
			Class:       "Kelas 12",
			PhoneNumber: "0895432102358",
			UserID:      "613db4f7-7fbd-4b82-97ed-1802368801b9",
		},
	}

	return userDetails
}
