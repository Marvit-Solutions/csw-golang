package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateUserDetails() []*ds.UserDetail {
	userDetail := []*ds.UserDetail{
		{
			ID:        "e36814fe-d479-4c5e-af49-4dede000f332",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "Administrator",
			Phone:     "08917283129283",
			UserId:    "6c8bcb83-a825-4df0-8c29-264402205b9b",
		},
		{
			ID:        "ecf50842-f521-4291-a31f-9f10c48fd889",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "User 1",
			Phone:     "08917283109283",
			UserId:    "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		},
		{
			ID:        "a9131226-069a-43f9-b52d-d1c5bfc1b878",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "User 2",
			Phone:     "0851728392716",
			UserId:    "41fb3d71-33bc-4a6e-9620-2d56f3090981",
		},
	}

	return userDetail
}
