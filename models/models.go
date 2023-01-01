package models

import "time"

type Article struct {
	ID        int       `json:"article_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	UserName  string    `json:"user_name"`
	NiceNum   int       `json:"nice"`
	CreatedAt time.Time `json:"created_at"`
}
