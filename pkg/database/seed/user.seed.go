package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/helper/password"
	"time"
)

func CreateUsers() []*ds.User {
	hashPasswordUser1, _ := password.HashPassword("user1")
	hashPasswordUser2, _ := password.HashPassword("user2")
	hashPasswordAdmin, _ := password.HashPassword("admin123")
	users := []*ds.User{
		{
			ID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "admin@gmail.com",
			Username:  "admin",
			Password:  string(hashPasswordAdmin),
			RoleId:    "419a8a2d-0abe-4413-ac49-39d33cf9838d",
		},
		{
			ID:        "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user1@gmail.com",
			Username:  "user1",
			Password:  string(hashPasswordUser1),
			RoleId:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
		{
			ID:        "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user2@gmail.com",
			Username:  "user2",
			Password:  string(hashPasswordUser2),
			RoleId:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
	}
	return users
}
