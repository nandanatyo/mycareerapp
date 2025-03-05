package rest

import (
	"mycareerapp/internal/app/article/usecase"
	"mycareerapp/internal/domain/dto"
	"mycareerapp/internal/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ArticleHandler struct {
	Validator      *validator.Validate
	ArticleUseCase usecase.ArticleUsecaseItf
	Middleware     middleware.MiddleWareI
}

func NewArticleHandler(routerGroup fiber.Router, validator *validator.Validate, articleUseCase usecase.ArticleUsecaseItf, middleware middleware.MiddleWareI) {
	handler := ArticleHandler{
		Validator:      validator,
		ArticleUseCase: articleUseCase,
		Middleware: middleware,
	}

	routerGroup = routerGroup.Group("/articles")

	routerGroup.Get("/", middleware.Authentication, handler.GetAllArticles)
	routerGroup.Post("/", middleware.Authentication, middleware.Authorization, handler.CreateArticle)
	routerGroup.Get("/:id", middleware.Authentication, handler.GetSpecificArticle)
	routerGroup.Patch("/:id", middleware.Authentication, middleware.Authorization, handler.UpdateArticle)
	routerGroup.Delete("/:id", middleware.Authentication, middleware.Authorization, handler.DeleteArticle)
}

func (h ArticleHandler) GetSpecificArticle(ctx *fiber.Ctx) error {

	articleID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "articleID should an UUID")
	}

	res, err := h.ArticleUseCase.GetSpecificArticle(articleID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(res)
}

func (h ArticleHandler) GetAllArticles(ctx *fiber.Ctx) error {

	res, err := h.ArticleUseCase.GetAllArticles()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": res,
	})
}

func (h ArticleHandler) CreateArticle(ctx *fiber.Ctx) error {

	var request dto.RequestCreateArticle

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	//validasi
	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	res, err := h.ArticleUseCase.CreateArticle(request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "article berhasil dibuat",
		"payload": res,
	})
}

func (h ArticleHandler) UpdateArticle(ctx *fiber.Ctx) error {

	var request dto.RequestUpdateArticle

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	articleID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "articleID should an UUID")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	err = h.ArticleUseCase.UpdateArticle(articleID, request)
	if err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (h ArticleHandler) DeleteArticle(ctx *fiber.Ctx) error {

	articleID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "articleID should an UUID")
	}

	err = h.ArticleUseCase.DeleteArticle(articleID)
	if err != nil {
		return err
	}

	return nil
}
