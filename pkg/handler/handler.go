package handler

import (
	"github.com/bruhlord-s/openboard-go/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("sign-in", h.signUp)
	}

	// TODO
	// api := router.Group("/api")
	// {

	// }

	return router
}
