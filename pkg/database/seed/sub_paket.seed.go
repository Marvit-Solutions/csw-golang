package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreateSubPaket() []*ds.SubPaket {
	subPaket := []*ds.SubPaket{
		{
			ID:                "1",
			PaketID:           1,
			NamaSubPaket:      "PAKET A",
			DeskripsiSubPaket: "Paket ini berisi soal-soal SKD",
			Harga:             10000,
		},
		{
			ID:                "2",
			PaketID:           1,
			NamaSubPaket:      "PAKET B",
			DeskripsiSubPaket: "Paket ini berisi soal-soal SKD",
			Harga:             10000,
		},
		{
			ID:                "3",
			PaketID:           2,
			NamaSubPaket:      "PAKET C",
			DeskripsiSubPaket: "Paket ini berisi soal-soal MATEMATIKA",
			Harga:             90000,
		},
		{
			ID:                "4",
			PaketID:           2,
			NamaSubPaket:      "PAKET D",
			DeskripsiSubPaket: "Paket ini berisi soal-soal MATEMATIKA",
			Harga:             90000,
		},
	}

	return subPaket
}
