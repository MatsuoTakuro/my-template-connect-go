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

type Store struct {
	StoreCD   int     `json:"store_cd"   validate:"required"`
	CompanyCD int     `json:"company_cd"   validate:"required"`
	StoreName string  `json:"store_name" validate:"required"`
	Address   string  `json:"address"    validate:"required"`
	Latitude  float64 `json:"latitude"   validate:"required"`
	Longitude float64 `json:"longitude"  validate:"required"`
}
