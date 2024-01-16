package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubPaket() []*ds.SubPaket {
	subPaket := []*ds.SubPaket{
		// SKD
		{
			ID:        "3bd1126e-9bc5-4b56-bb5e-4fcd098da64b",
			PaketID:   "2710cf91-3113-4f4b-a499-cfb4625f3960",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "SKD PAKET A",
			Harga:     50000,
		},
		{
			ID:        "fe40460b-b24c-466a-aa3b-cbc602bcba28",
			PaketID:   "2710cf91-3113-4f4b-a499-cfb4625f3960",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "SKD PAKET B",
			Harga:     150000,
		},
		{
			ID:        "1b209c62-c8ec-41e2-9af1-cab90c944695",
			PaketID:   "2710cf91-3113-4f4b-a499-cfb4625f3960",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "SKD PAKET C",
			Harga:     250000,
		},
		{
			ID:        "f127fe59-058e-44f5-a581-c3f903ffdb87",
			PaketID:   "2710cf91-3113-4f4b-a499-cfb4625f3960",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "SKD PAKET D",
			Harga:     300000,
		},
		// MTK
		{
			ID:        "12219bc1-813e-471c-b912-2dbdfd6b5356",
			PaketID:   "4d08b495-67bc-459d-8dcd-5ad88b8c6e2b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "MATEMATIKA PAKET A",
			Harga:     50000,
		},
		{
			ID:        "cd45edc1-ab81-43ba-8645-e5b513dac832",
			PaketID:   "4d08b495-67bc-459d-8dcd-5ad88b8c6e2b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "MATEMATIKA PAKET B",
			Harga:     150000,
		},
		{
			ID:        "122c9ba4-0474-453e-8b24-0fad03f5147c",
			PaketID:   "4d08b495-67bc-459d-8dcd-5ad88b8c6e2b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "MATEMATIKA PAKET C",
			Harga:     200000,
		},
		{
			ID:        "1ec0e022-e834-45c9-807f-db2397cf4a8b",
			PaketID:   "4d08b495-67bc-459d-8dcd-5ad88b8c6e2b",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Nama:      "MATEMATIKA PAKET D",
			Harga:     250000,
		},
	}

	return subPaket
}
