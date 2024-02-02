package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/helper/password"
	"time"
)

func CreateUsers() []*ds.Users {
	hashPasswordUser1, _ := password.HashPassword("user1")
	hashPasswordUser2, _ := password.HashPassword("user2")
	hashPasswordUser3, _ := password.HashPassword("user3")
	hashPasswordUser4, _ := password.HashPassword("user4")
	hashPasswordUser5, _ := password.HashPassword("user5")
	hashPasswordUser6, _ := password.HashPassword("user6")
	hashPasswordUser7, _ := password.HashPassword("user7")
	hashPasswordAdmin, _ := password.HashPassword("admin123")
	users := []*ds.Users{
		{
			ID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "admin@gmail.com",
			Password:  string(hashPasswordAdmin),
			RoleID:    "419a8a2d-0abe-4413-ac49-39d33cf9838d",
		},
		{
			ID:        "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user1@gmail.com",
			Password:  string(hashPasswordUser1),
			RoleID:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
		{
			ID:        "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user2@gmail.com",
			Password:  string(hashPasswordUser2),
			RoleID:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
		{
			ID:        "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user3@gmail.com",
			Password:  string(hashPasswordUser3),
			RoleID:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
		{
			ID:        "8f23e5cd-0e96-4e1d-a72b-b8f85ca083ee",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user4@gmail.com",
			Password:  string(hashPasswordUser4),
			RoleID:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
		{
			ID:        "fca8ef53-9f39-4cd6-940d-5df58110da59",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user5@gmail.com",
			Password:  string(hashPasswordUser5),
			RoleID:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
		{
			ID:        "66647904-fdee-4201-a4c7-fc83f5b136f7",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user6@gmail.com",
			Password:  string(hashPasswordUser6),
			RoleID:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
		{
			ID:        "613db4f7-7fbd-4b82-97ed-1802368801b9",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Email:     "user7@gmail.com",
			Password:  string(hashPasswordUser7),
			RoleID:    "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
		},
	}
	return users
}
