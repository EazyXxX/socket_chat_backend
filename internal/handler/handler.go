package handler

import (
	"os"
	"socket_chat_backend/internal/service"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONT_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	router.GET("/ws", h.handleWebSocket)

	// api := router.Group("/api", h.userIdentity)

	return router
}
