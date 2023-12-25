package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreateSubPaket() []*ds.SubPaket {
	subPaket := []*ds.SubPaket{
		{
			IdPaket:           1,
			NamaSubPaket:      "PAKET A",
			DeskripsiSubPaket: "Paket ini berisi soal-soal SKD",
			Harga:             10000,
		},
		{
			IdPaket:           1,
			NamaSubPaket:      "PAKET B",
			DeskripsiSubPaket: "Paket ini berisi soal-soal SKD",
			Harga:             10000,
		},
		{
			IdPaket:           2,
			NamaSubPaket:      "PAKET C",
			DeskripsiSubPaket: "Paket ini berisi soal-soal MATEMATIKA",
			Harga:             90000,
		},
		{
			IdPaket:           2,
			NamaSubPaket:      "PAKET D",
			DeskripsiSubPaket: "Paket ini berisi soal-soal MATEMATIKA",
			Harga:             90000,
		},
	}

	return subPaket
}
