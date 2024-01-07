package handler

import (
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	paketHandler "csw-golang/internal/delivery/http/handler/paket"
)

type Handler struct {
	AuthHandler  *authHandler.AuthHandler
	PaketHandler *paketHandler.PaketHandler
}
