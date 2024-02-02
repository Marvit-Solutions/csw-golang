package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubjects() []*ds.Subjects {
	subjects := []*ds.Subjects{
		// TWK
		{
			ID:          "58ab4e13-e016-4c42-8123-74b220db465a",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "65994d9d-9728-435c-bc6a-92ed2ec2b8c8",
			Name:        "Pancasila",
		},
		{
			ID:          "b8aa7572-a9b6-40e4-9ef5-06cf13f79e49",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "65994d9d-9728-435c-bc6a-92ed2ec2b8c8",
			Name:        "Konstitusi yang berlaku di Indonesia",
		},
		{
			ID:          "94905d6b-fab1-43f5-ad1e-573b1feea26d",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "65994d9d-9728-435c-bc6a-92ed2ec2b8c8",
			Name:        "NKRI & Bhinneka Tunggal Ika",
		},
		{
			ID:          "58cdaedb-c3a0-47b6-96c3-bbb3bad537c7",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "65994d9d-9728-435c-bc6a-92ed2ec2b8c8",
			Name:        "Demokrasi Indonesia",
		},
		{
			ID:          "901b0827-8ae3-418f-9c40-e94be260ed9a",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "65994d9d-9728-435c-bc6a-92ed2ec2b8c8",
			Name:        "HAM",
		},
		// TIU
		{
			ID:          "3524550f-baf8-4583-aad5-e2ff24fb1926",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "9159575a-440e-4e88-abfb-ad19b6715022",
			Name:        "Kemampuan Verbal",
		},
		{
			ID:          "fc0ddb98-b124-4497-aecd-d13b47bc0682",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "9159575a-440e-4e88-abfb-ad19b6715022",
			Name:        "Kemampuan Figural",
		},
		{
			ID:          "07ac80c9-a1a7-4159-87fe-6fb030dcb063",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "9159575a-440e-4e88-abfb-ad19b6715022",
			Name:        "Kemampuan Numerik",
		},
		// TKP
		{
			ID:          "2317d952-3641-4210-bdda-5f2be1a47d0a",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "2fadda64-888e-45cf-b16b-d7e1e2a0a1f3",
			Name:        "Anti Radikalisme",
		},
		{
			ID:          "d7a8055e-7290-40ad-984b-7d04cf30fc2a",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "2fadda64-888e-45cf-b16b-d7e1e2a0a1f3",
			Name:        "Anti Radikalisme",
		},
		{
			ID:          "255a57c9-cf08-4d67-b52c-8889ce330624",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "2fadda64-888e-45cf-b16b-d7e1e2a0a1f3",
			Name:        "Pelayanan Publik",
		},
		{
			ID:          "e55f2051-2488-4a02-a3c7-d310ba3e03eb",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "2fadda64-888e-45cf-b16b-d7e1e2a0a1f3",
			Name:        "Sosial Budaya",
		},
		{
			ID:          "a73f70a1-ad87-4507-8970-df5626ec6c8c",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubModuleID: "2fadda64-888e-45cf-b16b-d7e1e2a0a1f3",
			Name:        "Jejaring Kerja",
		},
	}

	return subjects
}
