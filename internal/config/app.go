package config

import (
	"go-linktree/internal/delivery/http"
	"go-linktree/internal/delivery/http/middleware"
	"go-linktree/internal/delivery/http/route"
	"go-linktree/internal/repository"
	"go-linktree/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repository.NewUserRepository(config.Log)
	urlRepository := repository.NewUrlRepository(config.Log)

	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	urlUseCase := usecase.NewUrlUseCase(config.DB, config.Log, config.Validate, urlRepository)

	userController := http.NewUserController(userUseCase, config.Log)
	urlController := http.NewUrlController(urlUseCase, config.Log)

	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		UrlController:  urlController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}
