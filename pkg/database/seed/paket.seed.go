package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreatePaket() []*ds.Paket {
	paket := []*ds.Paket{
		{
			ID:        "2710cf91-3113-4f4b-a499-cfb4625f3960",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "SKD",
		},
		{
			ID:        "4d08b495-67bc-459d-8dcd-5ad88b8c6e2b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "MATEMATIKA",
		},
	}

	return paket
}
