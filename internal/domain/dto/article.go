package dto

import "github.com/google/uuid"

type RequestCreateArticle struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required"`
	PhotoURL    string `json:"photo_url"`
	ArticleURL  string `json:"article_url" validate:"required"`
}

type RequestUpdateArticle struct {
	Title       string `json:"title" validate:"omitempty,min=3"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url"`
	ArticleURL  string `json:"article_url"`
}

type ResponseCreateArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url"`
	ArticleURL  string `json:"article_url"`
}

type ResponseGetArticle struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PhotoURL    string    `json:"photo_url"`
	ArticleURL  string    `json:"article_url"`
}
