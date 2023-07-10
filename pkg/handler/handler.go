package handler

import (
	"github.com/bruhlord-s/openboard-go/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/bruhlord-s/openboard-go/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			workspace.PUT("/:id", h.updateWorkspace)
			workspace.DELETE("/:id", h.deleteWorkspace)
		}
	}

	return router
}
