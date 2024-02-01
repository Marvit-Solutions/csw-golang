package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubscription() []*ds.Subscription {
	subscription := []*ds.Subscription{
		{
			ID:            "1",
			TransactionID: "1",
			SubPlanID:     "1",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "2",
			TransactionID: "2",
			SubPlanID:     "1",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "10",
			TransactionID: "10",
			SubPlanID:     "1",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "3",
			TransactionID: "3",
			SubPlanID:     "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "4",
			TransactionID: "4",
			SubPlanID:     "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "5",
			TransactionID: "5",
			SubPlanID:     "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "9",
			TransactionID: "9",
			SubPlanID:     "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "6",
			TransactionID: "6",
			SubPlanID:     "3",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "7",
			TransactionID: "8",
			SubPlanID:     "3",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "8",
			TransactionID: "7",
			SubPlanID:     "4",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
	}

	return subscription
}
