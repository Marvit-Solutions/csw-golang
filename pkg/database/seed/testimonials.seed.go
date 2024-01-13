package seed

import ds "csw-golang/internal/domain/entity/datastruct"

func CreateTestimonials() []*ds.Testimonials {
	testimonials := []*ds.Testimonials{
		{
			Id:           "dc38d81f-3c96-46b1-b985-e6f91eb01cb7",
			Name:         "Annisa Pertiwi",
			Status:       "Mahasiswa STIS 2021",
			ProfilePhoto: "assets/img/users/profile/profile.png",
			Comment:      "Bimbel csw merupakan salah satu bimbel online kedinasan dengan sistem pembelajaran terbaik yang pernah saya temui. Setelah mengikuti bimbel ini alhamdulillah sekarang saya bisa diterima di lembaga yang saya inginkan.",
			Rating:       4.5,
		},
		{
			Id:           "some-other-id",
			Name:         "John Doe",
			Status:       "Working Professional",
			ProfilePhoto: "assets/img/users/profile/john.png",
			Comment:      "I found the online mentoring experience very enriching. It helped me gain insights into my field and improve my skills.",
			Rating:       4.2,
		},
		{
			Id:           "another-id",
			Name:         "Jane Smith",
			Status:       "College Student",
			ProfilePhoto: "assets/img/users/profile/jane.png",
			Comment:      "The mentorship program has been a game-changer for me. It provided valuable guidance for my academic and career pursuits.",
			Rating:       4.8,
		},
		{
			Id:           "yet-another-id",
			Name:         "Michael Johnson",
			Status:       "Software Engineer",
			ProfilePhoto: "assets/img/users/profile/michael.png",
			Comment:      "The mentoring sessions were instrumental in helping me navigate challenges in the tech industry. Highly recommended!",
			Rating:       4.7,
		},
		{
			Id:           "id-5",
			Name:         "Sara Williams",
			Status:       "Marketing Specialist",
			ProfilePhoto: "assets/img/users/profile/sara.png",
			Comment:      "As a marketing professional, the mentorship program provided valuable insights into the ever-evolving landscape of digital marketing.",
			Rating:       4.5,
		},
		{
			Id:           "id-6",
			Name:         "Alex Turner",
			Status:       "Graphic Designer",
			ProfilePhoto: "assets/img/users/profile/alex.png",
			Comment:      "Being a graphic designer, the mentorship sessions helped me refine my design skills and stay updated with industry trends.",
			Rating:       4.3,
		},
		{
			Id:           "id-7",
			Name:         "Emma Thompson",
			Status:       "High School Student",
			ProfilePhoto: "assets/img/users/profile/emma.png",
			Comment:      "Even as a high school student, the mentorship program opened my eyes to various career possibilities and helped me set goals for the future.",
			Rating:       4.6,
		},
		{
			Id:           "id-8",
			Name:         "David Miller",
			Status:       "Entrepreneur",
			ProfilePhoto: "assets/img/users/profile/david.png",
			Comment:      "As an entrepreneur, the mentorship program provided valuable insights into business strategies and effective leadership.",
			Rating:       4.9,
		},
		{
			Id:           "id-9",
			Name:         "Sophie Roberts",
			Status:       "PhD Candidate",
			ProfilePhoto: "assets/img/users/profile/sophie.png",
			Comment:      "The mentorship program played a crucial role in shaping my research direction and providing guidance throughout my PhD journey.",
			Rating:       4.7,
		},
		{
			Id:           "id-10",
			Name:         "Chris Harris",
			Status:       "IT Professional",
			ProfilePhoto: "assets/img/users/profile/chris.png",
			Comment:      "The mentorship sessions provided valuable insights into the IT industry, helping me advance my career and stay competitive.",
			Rating:       4.4,
		},
	}

	return testimonials
}
