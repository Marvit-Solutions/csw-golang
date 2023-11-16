package handler

import (
	authHandler "csw-golang/internal/delivery/http/handler/auth"
)

type Handler struct {
	AuthHandler *authHandler.AuthHandler
}
