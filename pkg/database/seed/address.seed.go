package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateAddresses() []*ds.Addresses {
	addresses := []*ds.Addresses{
		{
			ID:           "47750cc7-9868-46e5-8545-7ed87550920c",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Surabaya",
			SubDistrict:  "Gubeng",
			UserDetailID: "e36814fe-d479-4c5e-af49-4dede000f332",
		},
		{
			ID:           "414354cc-45a5-47f6-879e-bfb4476eb0d6",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Tengah",
			RegencyCity:  "Solo",
			SubDistrict:  "Ngagel",
			UserDetailID: "ecf50842-f521-4291-a31f-9f10c48fd889",
		},
		{
			ID:           "c8b974da-ccca-475f-8d2f-c8c505e152fa",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Mojokerto",
			SubDistrict:  "Tawang",
			UserDetailID: "a9131226-069a-43f9-b52d-d1c5bfc1b878",
		},
		{
			ID:           "c7bb9f6c-4897-4af2-83f3-15f3ff157c8b",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Mojokerto",
			SubDistrict:  "Tawang",
			UserDetailID: "931fce0e-83c9-422a-ad44-f4f11b8e51e9",
		},
		{
			ID:           "ce82867a-952f-4792-80eb-a648c0361002",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Mojokerto",
			SubDistrict:  "Tawang",
			UserDetailID: "48e94d1e-217d-4f27-89d4-031bd44ee555",
		},
		{
			ID:           "58188414-1137-4871-950b-9e8395f7cea7",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Mojokerto",
			SubDistrict:  "Tawang",
			UserDetailID: "891c0769-e2df-4cbc-9bc9-2dc7344d0914",
		},
		{
			ID:           "1e4c1edb-fd5d-4f3d-8922-5736a287572c",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Mojokerto",
			SubDistrict:  "Tawang",
			UserDetailID: "236ef707-8fe1-4cc9-876c-2b9f6e49b5a8",
		},
		{
			ID:           "a378e996-b658-47bf-b1f6-6077ab3c3f31",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Mojokerto",
			SubDistrict:  "Tawang",
			UserDetailID: "612e746f-e0b9-4972-8f77-1ccda6306b6a",
		},
	}
	return addresses
}
