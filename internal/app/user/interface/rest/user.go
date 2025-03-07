package rest

import (
	"mycareerapp/internal/app/user/usecase"
	"mycareerapp/internal/domain/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUseCase usecase.UserUsecaseItf
	Validator   *validator.Validate
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, userUseCase usecase.UserUsecaseItf) {
	UserHandler := UserHandler{
		Validator:   validator,
		UserUseCase: userUseCase,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", UserHandler.RegisterUser)
	routerGroup.Post("/login", UserHandler.LoginUser)
}

func (h *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	var register dto.Register

	if err := ctx.BodyParser(&register); err != nil {
		return err
	}

	if err := h.Validator.Struct(register); err != nil {
		return err
	}

	err := h.UserUseCase.Register(register)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated)
}

func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	var login dto.Login

	if err := ctx.BodyParser(&login); err != nil {
		return err
	}

	if err := h.Validator.Struct(login); err != nil {
		return err
	}

	token, err := h.UserUseCase.Login(login)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
