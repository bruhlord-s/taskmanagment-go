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
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api/v1", h.userIdentity)
	{
		workspace := api.Group("/workspace")
		{
			workspace.GET("/:id", h.getWorkspaceById)
			workspace.GET("/", h.getAllWorkspaces)
			workspace.POST("/", h.createWorkspace)
			// workspace.PUT("/:id")
			// workspace.DELETE("/:id")
		}
	}

	return router
}
