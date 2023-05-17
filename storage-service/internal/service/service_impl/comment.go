package service_impl

import (
	"fmt"
	"storage-service/helpers"
	"storage-service/internal/db"
	"storage-service/internal/model"
)

type Comment struct {
	db *db.Postgres
}

func (c *Comment) GetAll(postId int) ([]model.Comment, error) {
	fmt.Println("Comment GetAll ", postId)
	var comments []model.Comment
	err := c.db.DB.Select(&comments, "select * from comment where post_id = $1", postId)

	return comments, err
}

func (c *Comment) Save(comment model.CommentRequest) error {
	insert := helpers.CommentRequestToCommentDto(comment)

	_, err := c.db.DB.Exec("insert into comment (post_id, content, author, created_at) values ($1, $2, $3, $4)",
		insert.PostId, insert.Content, insert.Author, insert.CreatedAt)

	return err
}

func (c *Comment) Delete(id int) error {
	_, err := c.db.DB.Exec("delete from comment where comment_id = $1", id)

	return err
}

func (c *Comment) Update(id int, comment model.CommentRequest) error {
	_, err := c.db.DB.Exec("update comment set content = $1 where comment_id = $2", comment.Content, id)

	return err
}

func (c *Comment) Get(id int) (model.Comment, error) {
	var comment model.Comment
	err := c.db.DB.Get(&comment, "select * from comment where comment_id = $1", id)

	return comment, err
}

func NewCommentService(db *db.Postgres) *Comment {
	return &Comment{db}
}
