package usecase

import (
	"mycareerapp/internal/app/article/repository"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/domain/entity"

	"github.com/google/uuid"
)

type ArticleUsecaseItf interface {
	GetAllArticles() (*[]dto.ResponseGetArticle, error)
	CreateArticle(request dto.RequestCreateArticle) (dto.ResponseCreateArticle, error)
	GetSpecificArticle(articleID uuid.UUID) (dto.ResponseGetArticle, error)
	UpdateArticle(articleID uuid.UUID, request dto.RequestUpdateArticle) error
	DeleteArticle(articleID uuid.UUID) error
}

type ArticleUsecase struct {
	ArticleRepository repository.ArticleMySQLItf
}

func NewArticleUsecase(articleRepository repository.ArticleMySQLItf) ArticleUsecaseItf {
	return &ArticleUsecase{
		ArticleRepository: articleRepository,
	}
}

func (u ArticleUsecase) GetAllArticles() (*[]dto.ResponseGetArticle, error) {

	articles := new([]entity.Article)

	err := u.ArticleRepository.GetAllArticles(articles)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ResponseGetArticle, len(*articles))
	for i, article := range *articles {
		res[i] = article.ParseToDTOGet()
	}

	return &res, nil
}

func (u ArticleUsecase) CreateArticle(request dto.RequestCreateArticle) (dto.ResponseCreateArticle, error) {

	article := entity.Article{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		PhotoURL:    request.PhotoURL,
		ArticleURL:  request.ArticleURL,
	}

	err := u.ArticleRepository.Create(&article)
	if err != nil {
		return dto.ResponseCreateArticle{}, err
	}

	return article.ParseToDTO(), nil
}

func (u ArticleUsecase) GetSpecificArticle(articleID uuid.UUID) (dto.ResponseGetArticle, error) {

	article := &entity.Article{
		ID: articleID,
	}

	err := u.ArticleRepository.GetSpecificArticle(article)
	if err != nil {
		return dto.ResponseGetArticle{}, err
	}

	return article.ParseToDTOGet(), err
}

func (u ArticleUsecase) UpdateArticle(articleID uuid.UUID, request dto.RequestUpdateArticle) error {

	article := &entity.Article{
		ID:          articleID,
		Title:       request.Title,
		Description: request.Description,
		PhotoURL:    request.PhotoURL,
		ArticleURL:  request.ArticleURL,
	}

	err := u.ArticleRepository.Update(article)
	if err != nil {
		return err
	}

	return nil
}

func (u ArticleUsecase) DeleteArticle(articleID uuid.UUID) error {

	article := &entity.Article{
		ID: articleID,
	}

	return u.ArticleRepository.Delete(article)
}
