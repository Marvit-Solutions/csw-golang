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
			SubPaketID:    "1",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "2",
			TransactionID: "2",
			SubPaketID:    "1",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "10",
			TransactionID: "10",
			SubPaketID:    "1",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "3",
			TransactionID: "3",
			SubPaketID:    "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "4",
			TransactionID: "4",
			SubPaketID:    "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "5",
			TransactionID: "5",
			SubPaketID:    "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "9",
			TransactionID: "9",
			SubPaketID:    "2",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "6",
			TransactionID: "6",
			SubPaketID:    "3",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "7",
			TransactionID: "8",
			SubPaketID:    "3",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
		{
			ID:            "8",
			TransactionID: "7",
			SubPaketID:    "4",
			UserID:        "6c8bcb83-a825-4df0-8c29-264402205b9b",
			Description:   "Paket ini berisi soal-soal SKD",
			DueDate:       time.Now(),
		},
	}

	return subscription
}
