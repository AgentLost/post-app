package service

import (
	"storage-service/internal/db"
	"storage-service/internal/model"
	"storage-service/internal/service/service_impl"
)

type Service struct {
	PostService
	CommentService
}

type PostService interface {
	Get(id int) (model.Post, error)
	GetAll() ([]model.Post, error)
	Save(post model.PostRequest) error
	Delete(id int) error
	Update(id int, post model.PostRequest) error
}

type CommentService interface {
	GetAll(postId int) ([]model.Comment, error)
	Get(id int) (model.Comment, error)
	Save(comment model.CommentRequest) error
	Delete(id int) error
	Update(id int, comment model.CommentRequest) error
}

func New(db *db.Postgres) *Service {
	return &Service{
		PostService:    service_impl.NewPostService(db),
		CommentService: service_impl.NewCommentService(db),
	}
}
