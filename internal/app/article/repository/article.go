package repository

import (
	"mycareerapp/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ArticleMySQLItf interface {
	Create(article *entity.Article) error
	GetAllArticles(articles *[]entity.Article) error
	GetSpecificArticle(articles *entity.Article) error
	Update(article *entity.Article) error
	Delete(article *entity.Article) error
}

type ArticleMySQL struct {
	db *gorm.DB
}

func NewArticleMySQL(db *gorm.DB) ArticleMySQLItf {
	return &ArticleMySQL{db}
}

func (r ArticleMySQL) GetAllArticles(articles *[]entity.Article) error {

	return r.db.Debug().Find(articles).Error
}

func (r ArticleMySQL) GetSpecificArticle(article *entity.Article) error {

	return r.db.Debug().First(article).Error
}

func (r ArticleMySQL) Update(article *entity.Article) error {

	return r.db.Debug().Updates(article).Error
}

func (r ArticleMySQL) Create(article *entity.Article) error {
	return r.db.Debug().Create(article).Error
}

func (r ArticleMySQL) Delete(article *entity.Article) error {
	
	q := r.db.Debug().Delete(article).RowsAffected

	if q == 0 {
		return fiber.NewError(http.StatusNotFound, "article not found")
	}
	
	return nil
}