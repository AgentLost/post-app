package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storage-service/internal/model"
	"strconv"
)

func (h *Handler) getComments(context *gin.Context) {
	param := context.Query("postId")
	postId, err := strconv.Atoi(param)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	comments, err := h.services.CommentService.GetAll(postId)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, comments)
}

func (h *Handler) deleteComment(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.CommentService.Delete(id)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}

func (h *Handler) updateComment(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	var request model.CommentRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CommentService.Update(id, request); err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "OK")
}

func (h *Handler) getComment(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	comments, err := h.services.CommentService.Get(id)

	if err != nil {
		NewErrorMessage(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, comments)
}
