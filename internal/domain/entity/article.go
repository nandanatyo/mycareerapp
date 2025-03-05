package entity

import (
	"mycareerapp/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	PhotoURL    string    `gorm:"type:text"`
	ArticleURL  string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
}

func (p Article) ParseToDTO() dto.ResponseCreateArticle {
	return dto.ResponseCreateArticle{
		Title:       p.Title,
		Description: p.Description,
		PhotoURL:    p.PhotoURL,
		ArticleURL:  p.ArticleURL,
	}
}

func (p Article) ParseToDTOGet() dto.ResponseGetArticle {
	return dto.ResponseGetArticle{
		ID: p.ID,
		Title:       p.Title,
		Description: p.Description,
		PhotoURL:    p.PhotoURL,
		ArticleURL:  p.ArticleURL,
	}
}
