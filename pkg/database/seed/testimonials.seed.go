package seed

import ds "csw-golang/internal/domain/entity/datastruct"

func CreateTestimonials() []*ds.Testimonials {
	testimonials := []*ds.Testimonials{
		{
			ID:      "dc38d81f-3c96-46b1-b985-e6f91eb01cb7",
			Comment: "Bimbel csw merupakan salah satu bimbel online kedinasan dengan sistem pembelajaran terbaik yang pernah saya temui. Setelah mengikuti bimbel ini alhamdulillah sekarang saya bisa diterima di lembaga yang saya inginkan.",
			Rating:  4.5,
		},
		{
			ID:      "d7209984-1326-4872-ac47-09ea7c2dd1a8",
			Comment: "I found the online mentoring experience very enriching. It helped me gain insights into my field and improve my skills.",
			Rating:  4.2,
		},
		{
			ID:      "e8432fb3-e5b1-4568-8aa1-8cdbf185f948",
			Comment: "The mentorship program has been a game-changer for me. It provided valuable guidance for my academic and career pursuits.",
			Rating:  4.8,
		},
		{
			ID:      "3eccb5cd-a010-4fd2-8884-bdc1b0f6226e",
			Comment: "The mentoring sessions were instrumental in helping me navigate challenges in the tech industry. Highly recommended!",
			Rating:  4.7,
		},
		{
			ID:      "39e8d413-2e02-4ccd-b176-201ff9012c5c",
			Comment: "As a marketing professional, the mentorship program provided valuable insights into the ever-evolving landscape of digital marketing.",
			Rating:  4.5,
		},
		{
			ID:      "29d9990a-caa6-49ef-a261-fdca7a828b0e",
			Comment: "Being a graphic designer, the mentorship sessions helped me refine my design skills and stay updated with industry trends.",
			Rating:  4.3,
		},
	}

	return testimonials
}
