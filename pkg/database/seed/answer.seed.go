package seed

// import (
// 	ds "csw-golang/internal/domain/entity/datastruct"
// )

// func CreateJawaban() []*ds.Jawaban {
// 	jawaban := []*ds.Jawaban{
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e124",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
// 			Jenis:      "A",
// 			Konten:     "Mengurangi kegiatan-kegiatan sosial keagamaan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e125",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
// 			Jenis:      "B",
// 			Konten:     "Membatasi munculnya organisasi berlatar belakang kesukuan",
// 			Is_correct: true,
// 			Pembobotan: 5,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e126",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
// 			Jenis:      "C",
// 			Konten:     "Membatasi munculnya organisasi berlatar belakang keagamaan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e127",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
// 			Jenis:      "D",
// 			Konten:     "Menghindari pertemuan-pertemuan lintas suku/agama",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e128",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e123",
// 			Jenis:      "E",
// 			Konten:     "Menumbuhkan sikap inklusif dalam kehidupan masyarakat",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e129",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e124",
// 			Jenis:      "A",
// 			Konten:     "Mengurangi kegiatan-kegiatan sosial keagamaan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e130",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e124",
// 			Jenis:      "B",
// 			Konten:     "Membatasi munculnya organisasi berlatar belakang kesukuan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e131",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e124",
// 			Jenis:      "C",
// 			Konten:     "Membatasi munculnya organisasi berlatar belakang keagamaan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e132",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e124",
// 			Jenis:      "D",
// 			Konten:     "Menghindari pertemuan-pertemuan lintas suku/agama",
// 			Is_correct: true,
// 			Pembobotan: 5,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e133",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e124",
// 			Jenis:      "E",
// 			Konten:     "Menumbuhkan sikap inklusif dalam kehidupan masyarakat",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e134",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e125",
// 			Jenis:      "A",
// 			Konten:     "Mengurangi kegiatan-kegiatan sosial keagamaan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e135",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e125",
// 			Jenis:      "B",
// 			Konten:     "Membatasi munculnya organisasi berlatar belakang kesukuan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e136",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e125",
// 			Jenis:      "C",
// 			Konten:     "Membatasi munculnya organisasi berlatar belakang keagamaan",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e137",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e125",
// 			Jenis:      "D",
// 			Konten:     "Menghindari pertemuan-pertemuan lintas suku/agama",
// 			Is_correct: false,
// 			Pembobotan: 0,
// 		},
// 		{
// 			ID:         "2daa6bd9-f39f-4ca8-5e8e-e692f687e138",
// 			SoalID:     "2daa6bd9-f39f-4ca8-5e8e-e692f687e125",
// 			Jenis:      "E",
// 			Konten:     "Menumbuhkan sikap inklusif dalam kehidupan masyarakat",
// 			Is_correct: true,
// 			Pembobotan: 5,
// 		},
// 	}
// 	return jawaban
// }
