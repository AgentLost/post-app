package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage-service/internal/model"
	"strconv"
)

func (h *Handler) getAllPosts(context *gin.Context) {
	posts, err := h.services.PostService.GetAll()

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, posts)
}

func (h *Handler) getPost(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.services.PostService.Get(id)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, post)
}

func (h *Handler) deletePost(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.PostService.Delete(id)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "OK")
}

func (h *Handler) updatePost(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	var input model.PostRequest
	if err := context.BindJSON(input); err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.PostService.Update(id, input)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "OK")
}
