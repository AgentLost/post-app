package handler

import (
	"api-service/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UseCases interface {
	SendComment(request model.CommentRequest) error
	SendPost(request model.PostRequest) error
}

type Handler struct {
	use         UseCases
	storageAddr []string
}

func New(storageAddr string, use UseCases) *Handler {
	return &Handler{
		use:         use,
		storageAddr: []string{storageAddr},
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	api := router.Group("/api")

	api.GET("/comments", h.redirect)

	comments := api.Group("/comments")
	{
		comments.POST("/", h.saveComment)
		comments.GET("/:id", h.redirect)
		comments.DELETE("/:id", h.redirect)
		comments.PUT("/:id", h.redirect)
	}

	post := api.Group("/posts")

	{
		post.GET("/", h.redirect)
		post.POST("/", h.savePost)
		post.GET("/:id", h.redirect)
		post.DELETE("/:id", h.redirect)
		post.PUT("/:id", h.redirect)
	}

	return router
}

func (h *Handler) saveComment(c *gin.Context) {
	var input model.CommentRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println(input)
	err := h.use.SendComment(input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) savePost(c *gin.Context) {
	var input model.PostRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.use.SendPost(input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
