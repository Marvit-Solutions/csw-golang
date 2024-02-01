package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"time"
)

func CreateSubPlanDetail() []*ds.SubPlanDetail {
	SubPlanDetail := []*ds.SubPlanDetail{
		// SKD Paket ABCD
		// Paket A
		{
			ID:          "73b3ea81-b9c7-4f83-bd3a-9a2af2b4c090",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "3bd1126e-9bc5-4b56-bb5e-4fcd098da64b",
			GrupPejuang: true,
			Exercise:    660,
			Access:      2,
			Module:      false,
			TryOut:      0,
			Zoom:        false,
		},
		// Paket B
		{
			ID:          "cd593e97-25d1-4e3f-9540-5fe2005e757b",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "fe40460b-b24c-466a-aa3b-cbc602bcba28",
			GrupPejuang: true,
			Exercise:    660,
			Access:      3,
			Module:      true,
			TryOut:      2,
			Zoom:        false,
		},
		// Paket C
		{
			ID:          "243b3db9-f3ca-4c7e-9338-c708f51977ca",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "1b209c62-c8ec-41e2-9af1-cab90c944695",
			GrupPejuang: true,
			Exercise:    660,
			Access:      6,
			Module:      true,
			TryOut:      3,
			Zoom:        true,
		},
		// Paket D
		{
			ID:          "d4505609-2f38-4cbf-bf0c-31e2abda928f",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "f127fe59-058e-44f5-a581-c3f903ffdb87",
			GrupPejuang: true,
			Exercise:    1100,
			Access:      6,
			Module:      true,
			TryOut:      4,
			Zoom:        true,
		},
		// MTK Paket ABCD
		// Paket A
		{
			ID:          "3d6053e5-31d3-4347-ae12-a3ddaa065872",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "12219bc1-813e-471c-b912-2dbdfd6b5356",
			GrupPejuang: true,
			Exercise:    660,
			Access:      2,
			Module:      false,
			TryOut:      0,
			Zoom:        false,
		},
		// Paket B
		{
			ID:          "4c696e20-0ec3-4e15-855a-952608cb9e85",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "cd45edc1-ab81-43ba-8645-e5b513dac832",
			GrupPejuang: true,
			Exercise:    240,
			Access:      6,
			Module:      true,
			TryOut:      2,
			Zoom:        false,
		},
		// Paket C
		{
			ID:          "604b6af7-5d03-47a8-a806-f9f50afa6bcb",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "122c9ba4-0474-453e-8b24-0fad03f5147c",
			GrupPejuang: true,
			Exercise:    240,
			Access:      6,
			Module:      true,
			TryOut:      3,
			Zoom:        true,
		},
		// Paket D
		{
			ID:          "e3db2abd-cfff-4030-a988-2075db8883c0",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			SubPlanID:   "1ec0e022-e834-45c9-807f-db2397cf4a8b",
			GrupPejuang: true,
			Exercise:    400,
			Access:      6,
			Module:      true,
			TryOut:      4,
			Zoom:        true,
		},
	}

	return SubPlanDetail
}
