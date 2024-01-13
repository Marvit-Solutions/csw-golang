package handler

import (
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	paketHandler "csw-golang/internal/delivery/http/handler/paket"
)

type Handler struct {
	AuthHandler   *authHandler.AuthHandler
	PaketHandler  *paketHandler.PaketHandler
	MentorHandler *mentorHandler.MentorHandler
}
