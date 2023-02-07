package handler

import (
	"contentPublicationBACK/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
			auth.POST("/login", h.singIn)
			auth.GET("/user", h.getUser)
			auth.GET("/logout", h.logOut)
			auth.POST("/check-exist", h.checkExist)
		}

		account := mainRout.Group("/account", h.userIdentity)
		{
			account.GET("/", h.getUserInfo)
			account.GET("/likes", h.getLikedTitles)
		}

		titles := mainRout.Group("/titles")
		{
			titles.GET("/", h.getTitles)
			titles.GET("/add-view/:id", h.updateViewsCount)
			titles.GET("/random", h.getRandom)
			titles.GET("/content-all", h.getAllPossibleContent)
			titles.GET("/:id", h.getTitle)
			titles.POST("/", h.createTitle)
			titles.POST("/like", h.likeTitle)
			titles.POST("/unlike", h.unlikeTitle)
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
