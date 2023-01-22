package handler

import (
	"contentPublicationBACK/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.singUp)
		auth.GET("/login", h.singIn)
	}

	titles := router.Group("/titles", h.userIdentity)
	{
		titles.GET("/", h.getTitles)
		titles.GET("/:id", h.getTitle)
		titles.POST("/", h.createTitle)
		titles.PUT("/:id", h.updateTitle)
		titles.DELETE("/:id", h.deleteTitle)
		titles.GET("/categories/", h.getTitlesByCategories)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return router
}
