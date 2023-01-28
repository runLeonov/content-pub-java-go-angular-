package handler

import (
	"contentPublicationBACK/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/rs/cors"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(CORSConfig()))

	mainRout := router.Group("/content")
	{
		auth := mainRout.Group("/auth")
		{
			auth.POST("/register", h.singUp)
			auth.GET("/login", h.singIn)
		}

		titles := mainRout.Group("/titles")
		{
			titles.GET("/", h.getTitles)
			titles.GET("/random", h.getRandom)
			titles.GET("/:id", h.getTitle)
			titles.POST("/", h.createTitle)
			titles.PUT("/:id", h.updateTitle)
			titles.DELETE("/:id", h.deleteTitle)
			titles.GET("/categories/", h.getTitlesByCategories)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return router
}

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}
