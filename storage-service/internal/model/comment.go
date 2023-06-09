package model

import "time"

type Comment struct {
	Id        int       `json:"id" db:"comment_id"`
	Content   string    `json:"content" db:"content"`
	PostId    int       `json:"postId" db:"post_id"`
	Author    string    `json:"author" db:"author"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type CommentDTO struct {
	Content   string    `json:"content"`
	PostId    int       `json:"postId"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentRequest struct {
	Content string `json:"content"`
	PostId  int    `json:"postId"`
	Author  string `json:"author"`
}
