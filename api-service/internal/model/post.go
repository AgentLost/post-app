package model

import "time"

type Post struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

type PostDTO struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
