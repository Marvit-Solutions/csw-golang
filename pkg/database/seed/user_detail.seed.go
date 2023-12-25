package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreateUserDetails() []*ds.UserDetail {
	userDetail := []*ds.UserDetail{
		{
			Name:   "Administrator",
			Phone:  "08917283129283",
			UserId: "6c8bcb83-a825-4df0-8c29-264402205b9b",
		},
		{
			Name:   "User 1",
			Phone:  "08917283109283",
			UserId: "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
		},
		{
			Name:   "User 2",
			Phone:  "0851728392716",
			UserId: "41fb3d71-33bc-4a6e-9620-2d56f3090981",
		},
	}

	return userDetail
}
