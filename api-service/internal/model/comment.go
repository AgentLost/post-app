package model

import "time"

type Comment struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	PostId    int       `json:"postId"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
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
