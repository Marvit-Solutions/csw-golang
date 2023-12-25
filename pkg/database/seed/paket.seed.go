package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreatePaket() []*ds.Paket {
	paket := []*ds.Paket{
		{
			NamaPaket:      "SKD",
			DeskripsiPaket: "Paket ini berisi soal-soal SKD",
		},
		{
			NamaPaket:      "MATEMATIKA",
			DeskripsiPaket: "Paket ini berisi soal-soal MATEMATIKA",
		},
	}

	return paket
}
