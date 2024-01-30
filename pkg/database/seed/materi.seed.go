package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreateMateri() []*ds.Materi {
	materis := []*ds.Materi{
		{
			ID:      "2daa6bd9-f39g-4ca8-5e8e-e692f687e124",
			ModulID: "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Nama:    "Pancasila",
		},
		{
			ID:      "2daa6bd9-f39g-4ca8-5e8e-e692f687e125",
			ModulID: "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Nama:    "Konstitusi yang berlaku di Indonesia",
		},
		{
			ID:      "2daa6bd9-f39g-4ca8-5e8e-e692f687e126",
			ModulID: "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Nama:    "NKRI & Bhinneka Tunggal Ika",
		},
		{
			ID:      "2daa6bd9-f39g-4ca8-5e8e-e692f687e127",
			ModulID: "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Nama:    "Demokrasi Indonesia",
		},
		{
			ID:      "2daa6bd9-f39g-4ca8-5e8e-e692f687e128",
			ModulID: "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Nama:    "HAM",
		},
	}
	return materis
}
