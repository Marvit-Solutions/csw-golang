package seed

import "csw-golang/internal/domain/entity/datastruct"

func CreateUserTestimonials() []*datastruct.UserTestimonials {
	userTestimonials := []*datastruct.UserTestimonials{
		{
			ID:            "65d9b1c2-3d35-45b0-9ce7-c99666f1a98a",
			UserID:        "c7668a02-f1bd-4ebb-bf30-aaeebfedc16b",
			TestimonialID: "dc38d81f-3c96-46b1-b985-e6f91eb01cb7",
		},
		{
			ID:            "db8f3d41-1e01-4387-9a96-c0c8ed4c0936",
			UserID:        "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e",
			TestimonialID: "d7209984-1326-4872-ac47-09ea7c2dd1a8",
		},
		{
			ID:            "570e6b63-7c21-49f0-925e-c2cf6486ec01",
			UserID:        "41fb3d71-33bc-4a6e-9620-2d56f3090981",
			TestimonialID: "e8432fb3-e5b1-4568-8aa1-8cdbf185f948",
		},
		{
			ID:            "a89bfb42-6a7b-4014-b316-5323b1fb6ce4",
			UserID:        "8f23e5cd-0e96-4e1d-a72b-b8f85ca083ee",
			TestimonialID: "3eccb5cd-a010-4fd2-8884-bdc1b0f6226e",
		},
		{
			ID:            "29716b3b-40f9-4832-b0ee-61175d32f391",
			UserID:        "fca8ef53-9f39-4cd6-940d-5df58110da59",
			TestimonialID: "39e8d413-2e02-4ccd-b176-201ff9012c5c",
		},
		{
			ID:            "fec4655b-9f94-4e8e-a5ae-097eb56ecd55",
			UserID:        "66647904-fdee-4201-a4c7-fc83f5b136f7",
			TestimonialID: "29d9990a-caa6-49ef-a261-fdca7a828b0e",
		},
	}

	return userTestimonials
}
