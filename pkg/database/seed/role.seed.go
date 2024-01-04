package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateRoles() []*ds.Role {
	roles := []*ds.Role{
		{
			ID:        "419a8a2d-0abe-4413-ac49-39d33cf9838d",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Role:      "Admin",
		},
		{
			ID:        "b29112fc-c7b5-4386-a31e-f2c040de7fcb",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Role:      "User",
		},
	}

	return roles
}
