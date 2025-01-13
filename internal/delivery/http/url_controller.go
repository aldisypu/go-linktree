package http

import (
	"go-linktree/internal/delivery/http/middleware"
	"go-linktree/internal/model"
	"go-linktree/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UrlController struct {
	UseCase *usecase.UrlUseCase
	Log     *logrus.Logger
}

func NewUrlController(useCase *usecase.UrlUseCase, log *logrus.Logger) *UrlController {
	return &UrlController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *UrlController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateUrlRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}
	request.Username = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to creating url")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UrlResponse]{Data: response})
}

func (c *UrlController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUrlRequest{
		Username: auth.ID,
		ID:       ctx.Params("urlId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to getting url")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UrlResponse]{Data: response})
}

func (c *UrlController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateUrlRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	request.Username = auth.ID
	request.ID = ctx.Params("urlId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to updating url")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UrlResponse]{Data: response})
}

func (c *UrlController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	urlId := ctx.Params("urlId")

	request := &model.DeleteUrlRequest{
		Username: auth.ID,
		ID:       urlId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to deleting url")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
