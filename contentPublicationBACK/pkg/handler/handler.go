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

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:4200/"},
	//	AllowMethods:     []string{"PUT", "PATCH", "GET"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:4200"}
	//router.Use(cors.New(config))

	mainRout := router.Group("/content")
	{
		auth := mainRout.Group("/auth")
		{
			auth.POST("/register", h.singUp)
			auth.GET("/login", h.singIn)
		}

		titles := mainRout.Group("/titles")
		//titles := mainRout.Group("/titles", h.userIdentity)
		{
			titles.GET("/", h.getTitles)
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
