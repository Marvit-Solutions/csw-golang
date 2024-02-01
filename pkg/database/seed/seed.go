package seed

import (
	"gorm.io/gorm"
)

type Seed struct {
	Seed interface{}
}

func RegisterSeed(db *gorm.DB) []Seed {
	return []Seed{
		{Seed: CreateRoles()},
		{Seed: CreateUsers()},
		{Seed: CreateUserDetails()},
		{Seed: CreateAddresses()},
		// {Seed: CreatePaket()},
		// {Seed: CreateSubPlan()},
		{Seed: CreateSubPlanDetail()},
		// {Seed: CreateModul()},
		// {Seed: CreateMateri()},
		// {Seed: CreateSubMateri()},
		// {Seed: CreateLatihanSoal()},
		// {Seed: CreateSoal()},
		// {Seed: CreateAnswer()},
		// {Seed: CreateStoredAnswer()},

		// {Seed: CreateSubscription()},
		// {Seed: CreateMentors()},
		// {Seed: CreateTestimonials()},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seed := range RegisterSeed(db) {
		var count int64
		if err := db.Model(seed.Seed).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			err := db.Debug().Create(seed.Seed).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
