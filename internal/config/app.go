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
	linkRepository := repository.NewLinkRepository(config.Log)

	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	linkUseCase := usecase.NewLinkUseCase(config.DB, config.Log, config.Validate, linkRepository)

	userController := http.NewUserController(userUseCase, config.Log)
	linkController := http.NewLinkController(linkUseCase, config.Log)

	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		LinkController: linkController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}
