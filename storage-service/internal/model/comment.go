package model

import "time"

type Comment struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	PostId    int       `json:"postId"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentDTO struct {
	Text      string    `json:"text"`
	PostId    int       `json:"postId"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentRequest struct {
	Text   string `json:"text"`
	PostId int    `json:"postId"`
	Author string `json:"author"`
}
