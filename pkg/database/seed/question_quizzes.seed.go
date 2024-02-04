package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateQuestionQuizzes() []*ds.QuestionQuizzes {
	questionQuizzes := []*ds.QuestionQuizzes{
		// Pretest Pancasila
		{
			ID:             "69551cd4-47f1-4fc1-b00f-a7df35a90077",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			Image:          "assets/img/questions/question.png",
			Content:        "Setiap warga masyarakat perlu mengupayakan persatuan dan kesatuan dalam bingkai kebhinekaan. Untuk itu yang perlu ditumbuhkan dalam diri setiap warga masyarakat adalah...",
			Weight:         5,
		},
		{
			ID:             "53ed22ce-0af5-453b-b87c-0bfb1565556a",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			Image:          "assets/img/questions/question.png",
			Content:        "Rumusan dan susunan Pancasila yang benar dan sah tercantum dalam ….",
			Weight:         5,
		},
		{
			ID:             "4f7c1f79-25fd-4e60-9219-073a414e9720",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			Image:          "assets/img/questions/question.png",
			Content:        " Landasan idiil negara kita adalah ….",
			Weight:         5,
		},
		{
			ID:             "0d1f4266-ad1e-41f0-b36a-9dc808e63c93",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			Image:          "assets/img/questions/question.png",
			Content:        "Pedoman Penghayatan dan Pengamalan Pancasila ditetapkan pada tanggal ….",
			Weight:         5,
		},
		{
			ID:             "0c301ace-1869-475c-99a4-7585a509e622",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8c88d1e5-eb9f-4d4d-a94e-94be94209cb0",
			Image:          "assets/img/questions/question.png",
			Content:        "Pancasila sebagai ideologi terbuka harus berorientasi pada masa depan dan mampu melihat semua kemungkinan yang dapat terjadi pada masa sekarang. Berikut ini yang merupakan salah satu unsur ideologi terbuka yang dimiliki oleh Pancasila adalah ….",
			Weight:         5,
		},
		// Posttest Pancasila
		{
			ID:             "2b45009a-e8cc-43f9-b957-c68fe137f9c2",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8049dc26-01c1-426a-803d-570b799d3810",
			Image:          "assets/img/questions/question.png",
			Content:        " Ciri khas paham integralistik Indonesia dapat dilihat dalam kehidupan …..",
			Weight:         5,
		},
		{
			ID:             "f4fad3b8-e6f0-48b3-8dd9-ba6982b90639",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8049dc26-01c1-426a-803d-570b799d3810",
			Image:          "assets/img/questions/question.png",
			Content:        " Asal mula tidak langsung Pancasila hakikatnya berasal dari bangsa Indonesia itu sendiri. Pernyataan tersebut merupakan uraian dari …",
			Weight:         5,
		},
		{
			ID:             "3fdcf4ac-e8ba-4bc8-bbdf-906fc756ebc0",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8049dc26-01c1-426a-803d-570b799d3810",
			Image:          "assets/img/questions/question.png",
			Content:        " Pancasila berasal dari kata panca yang berarti lima dan sila yang berarti sendi, asas, dasar atau peraturan tingkah laku yang penting dan baik. Dengan demikian, Pancasila merupakan lima dasar yang berisi pedoman atau aturan tentang tingkah laku yang penting dan baik, merupakan pengertian Pancasila yang diungkapkan oleh ….",
			Weight:         5,
		},
		{
			ID:             "b8456d7e-6940-485b-b32c-16207d8e1538",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8049dc26-01c1-426a-803d-570b799d3810",
			Image:          "assets/img/questions/question.png",
			Content:        " Pancasila digunakan sebagai dasar untuk mengatur penyelenggaraan ketatanegaraan negara, hal ini sesuai dengan kedudukan Pancasila sebagai ….",
			Weight:         5,
		},
		{
			ID:             "fe2d281e-fc4e-427e-b487-63d4065ef7a0",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			TestTypeQuizID: "8049dc26-01c1-426a-803d-570b799d3810",
			Image:          "assets/img/questions/question.png",
			Content:        "Mengembangkan sikap bahwa bangsa Indonesia merupakan bagian dari seluruh umat manusia merupakan perwujudan sila ke- …..",
			Weight:         5,
		},
	}

	return questionQuizzes
}
