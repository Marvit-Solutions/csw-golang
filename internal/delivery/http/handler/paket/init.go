package paket

import (
	pc "csw-golang/internal/domain/usecase/paket"
)

type PaketHandler struct {
	paketUsecase pc.PaketUsecase
}

func New(paketUsecase pc.PaketUsecase) *PaketHandler {
	return &PaketHandler{
		paketUsecase,
	}
}
