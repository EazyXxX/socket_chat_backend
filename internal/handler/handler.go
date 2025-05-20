package handler

import (
	"socket_chat_backend/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	hub      *WebSocketHub
}

func NewHandler(services *service.Service) *Handler {
	hub := newHub()
	go hub.run()
	
	return &Handler{services: services, hub: hub}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	router.GET("/ws", h.handleWebSocket)

	// api := router.Group("/api", h.userIdentity)

	return router
}
