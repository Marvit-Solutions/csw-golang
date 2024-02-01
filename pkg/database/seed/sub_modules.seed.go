package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubModules() []*ds.SubModules {
	Submodules := []*ds.SubModules{
		{
			ID:          "65994d9d-9728-435c-bc6a-92ed2ec2b8c8",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ModuleID:    "97150a42-90bd-4b72-9b83-cbd7075f7662",
			Name:        "TWK",
			Description: "Berisi tentang materi - materi pembelajaran twk yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
		{
			ID:          "9159575a-440e-4e88-abfb-ad19b6715022",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ModuleID:    "97150a42-90bd-4b72-9b83-cbd7075f7662",
			Name:        "TIU",
			Description: "Berisi tentang materi - materi pembelajaran tiu yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
		{
			ID:          "2fadda64-888e-45cf-b16b-d7e1e2a0a1f3",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ModuleID:    "97150a42-90bd-4b72-9b83-cbd7075f7662",
			Name:        "TKP",
			Description: "Berisi tentang materi - materi pembelajaran tkp yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
		{
			ID:          "39dbb7c2-dbae-4877-8dca-709837f6ec46",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ModuleID:    "575282d4-9a34-47bf-8093-0c8c06520485",
			Name:        "Matematika",
			Description: "Berisi tentang materi - materi pembelajaran matematika yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
	}

	return Submodules
}
