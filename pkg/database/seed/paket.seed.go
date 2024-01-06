package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreatePaket() []*ds.Paket {
	paket := []*ds.Paket{
		{
			ID:             "1",
			NamaPaket:      "SKD",
			DeskripsiPaket: "Paket ini berisi soal-soal SKD",
		},
		{
			ID:             "2",
			NamaPaket:      "MATEMATIKA",
			DeskripsiPaket: "Paket ini berisi soal-soal MATEMATIKA",
		},
	}

	return paket
}
