package helpers

import (
	"storage-service/internal/model"
	"time"
)

func CommentRequestToCommentDto(request model.CommentRequest) model.CommentDTO {
	return model.CommentDTO{
		Content:   request.Content,
		PostId:    request.PostId,
		Author:    request.Author,
		CreatedAt: time.Now(),
	}
}

func PostRequestToRequestDto(request model.PostRequest) model.PostDTO {
	return model.PostDTO{
		Title:     request.Title,
		Content:   request.Content,
		Author:    request.Author,
		CreatedAt: time.Now(),
	}
}
