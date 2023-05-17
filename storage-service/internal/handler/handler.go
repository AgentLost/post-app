package handler

import (
	"github.com/gin-gonic/gin"
	"storage-service/internal/service"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	api := router.Group("/api")

	api.GET("/comments", h.getComments)

	comments := api.Group("/comments")
	{
		comments.GET("/:id", h.getComment)
		comments.DELETE("/:id", h.deleteComment)
		comments.PUT("/:id", h.updateComment)
	}

	post := api.Group("/posts")

	{
		post.GET("/", h.getAllPosts)
		post.GET("/:id", h.getPost)
		post.DELETE("/:id", h.deletePost)
		post.PUT("/:id", h.updatePost)
	}

	return router
}

func New(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}
