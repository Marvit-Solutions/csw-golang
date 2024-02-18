package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateModules() []*ds.Modules {
	modules := []*ds.Modules{
		{
			ID:        "97150a42-90bd-4b72-9b83-cbd7075f7662",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "SKD",
		},
		{
			ID:        "575282d4-9a34-47bf-8093-0c8c06520485",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "MATEMATIKA",
		},
	}

	return modules
}
