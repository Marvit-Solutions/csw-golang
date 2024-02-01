package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateAddresses() []*ds.Address {
	addresses := []*ds.Address{
		{
			ID:           "47750cc7-9868-46e5-8545-7ed87550920c",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Surabaya",
			Subdistrict:  "Gubeng",
			UserDetailID: "e36814fe-d479-4c5e-af49-4dede000f332",
		},
		{
			ID:           "414354cc-45a5-47f6-879e-bfb4476eb0d6",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Tengah",
			RegencyCity:  "Solo",
			Subdistrict:  "Ngagel",
			UserDetailID: "ecf50842-f521-4291-a31f-9f10c48fd889",
		},
		{
			ID:           "c8b974da-ccca-475f-8d2f-c8c505e152fa",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Province:     "Jawa Timur",
			RegencyCity:  "Mojokerto",
			Subdistrict:  "Tawang",
			UserDetailID: "a9131226-069a-43f9-b52d-d1c5bfc1b878",
		},
	}
	return addresses
}
