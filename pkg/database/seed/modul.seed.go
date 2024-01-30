package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
)

func CreateModul() []*ds.Modul {
	moduls := []*ds.Modul{
		{
			ID:        "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
			Nama:      "TWK",
			Deskripsi: "Berisi tentang materi - materi pembelajaran TWK yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
		{
			ID:        "2daasdd9-f39f-4ca8-5e8e-e692f687e124",
			Nama:      "TIU",
			Deskripsi: "Berisi tentang materi - materi pembelajaran TIU yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
		{
			ID:        "2daa6bdq-f39f-4ca8-5e8e-e692f687e125",
			Nama:      "TKP",
			Deskripsi: "Berisi tentang materi - materi pembelajaran TKP yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
		{
			ID:        "212a6bd9-f39f-4ca8-5e8e-e692f68sddds",
			Nama:      "Matematika",
			Deskripsi: "Berisi tentang materi - materi pembelajaran Matematika yang dapat digunakan selama mengikuti pertemuan mulai dari 20 September - 30 Desember 2022",
		},
	}
	return moduls
}
