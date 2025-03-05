package bootstrap

import (
	"fmt"
	articlehandler "mycareerapp/internal/app/article/interface/rest"
	articlerepository "mycareerapp/internal/app/article/repository"
	articleusecase "mycareerapp/internal/app/article/usecase"
	userhandler "mycareerapp/internal/app/user/interface/rest"
	userrepository "mycareerapp/internal/app/user/repository"
	userusecase "mycareerapp/internal/app/user/usecase"
	"mycareerapp/internal/infra/env"
	"mycareerapp/internal/infra/fiber"
	"mycareerapp/internal/infra/jwt"
	"mycareerapp/internal/infra/mysql"
	"mycareerapp/internal/middleware"

	"github.com/go-playground/validator/v10"
)

func Start() error {

	config, err := env.New()
	if err != nil {
		panic(err)
	}

	database, err := mysql.New(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	))

	err = mysql.Migrate(database)
	if err != nil {
		panic(err)
	}

	val := validator.New()

	jwt := jwt.NewJwt(config)

	middleware := middleware.NewMiddleware(jwt)

	app := fiber.New()

	v1 := app.Group("/api/v1")

	articleRepository := articlerepository.NewArticleMySQL(database)
	articleUseCase := articleusecase.NewArticleUsecase(articleRepository)
	articlehandler.NewArticleHandler(v1, val, articleUseCase, middleware)

	userRepository := userrepository.NewUserMySQL(database)
	userUseCase := userusecase.NewUserUsecase(userRepository, jwt)
	userhandler.NewUserHandler(v1, val, userUseCase)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
