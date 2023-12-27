package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateUserDetails() []*ds.UserDetail {
	userDetail := []*ds.UserDetail{
		{
			ID:           "e36814fe-d479-4c5e-af49-4dede000f332",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Nama:         "Administrator",
			Telepon:      "08917283129283",
			TanggalLahir: time.Date(1998, time.December, 28, 0, 0, 0, 0, time.UTC),
			UserId:       "6c8bcb83-a825-4df0-8c29-264402205b9b",
		},
		{
			ID:           "ecf50842-f521-4291-a31f-9f10c48fd889",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Nama:         "User 1",
			Telepon:      "08917283109283",
			TanggalLahir: time.Date(2005, time.February, 20, 0, 0, 0, 0, time.UTC),
			UserId:       "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		},
		{
			ID:           "a9131226-069a-43f9-b52d-d1c5bfc1b878",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Nama:         "User 2",
			Telepon:      "0851728392716",
			TanggalLahir: time.Date(2001, time.August, 29, 0, 0, 0, 0, time.UTC),
			UserId:       "41fb3d71-33bc-4a6e-9620-2d56f3090981",
		},
	}

	return userDetail
}
