package service_impl

import (
	"storage-service/helpers"
	"storage-service/internal/db"
	"storage-service/internal/model"
)

type Post struct {
	db *db.Postgres
}

func (p *Post) Get(id int) (model.Post, error) {
	var post model.Post
	err := p.db.DB.Get(&post, "select * from post where id = $1", id)

	return post, err
}

func (p *Post) GetAll() ([]model.Post, error) {
	var posts []model.Post
	err := p.db.DB.Select(&posts, "select * from post")

	return posts, err
}

func (p *Post) Save(post model.PostRequest) error {
	insert := helpers.PostRequestToRequestDto(post)

	_, err := p.db.DB.Exec("insert into post (title, content, author, created_at) values ($1, $2, $3, $4)",
		insert.Title, insert.Content, insert.Author, insert.CreatedAt)

	return err
}

func (p *Post) Delete(id int) error {
	_, err := p.db.DB.Exec("delete from post where id = $1", id)

	return err
}

func (p *Post) Update(id int, post model.PostRequest) error {
	_, err := p.db.DB.Exec("update post set title = $1, content = $2 where id = $3", post.Title, post.Content, id)

	return err
}

func NewPostService(db *db.Postgres) *Post {
	return &Post{db}
}
