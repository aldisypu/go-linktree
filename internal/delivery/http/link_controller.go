package http

import (
	"go-linktree/internal/delivery/http/middleware"
	"go-linktree/internal/model"
	"go-linktree/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LinkController struct {
	UseCase *usecase.LinkUseCase
	Log     *logrus.Logger
}

func NewLinkController(useCase *usecase.LinkUseCase, log *logrus.Logger) *LinkController {
	return &LinkController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *LinkController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateLinkRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}
	request.Username = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to creating link")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.LinkResponse]{Data: response})
}

func (c *LinkController) List(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	request := &model.ListLinkRequest{
		Username: username,
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list links")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.LinkResponse]{Data: responses})
}

func (c *LinkController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetLinkRequest{
		Username: auth.ID,
		ID:       ctx.Params("linkId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to getting link")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.LinkResponse]{Data: response})
}

func (c *LinkController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateLinkRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parsing request body")
		return fiber.ErrBadRequest
	}

	request.Username = auth.ID
	request.ID = ctx.Params("linkId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to updating link")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.LinkResponse]{Data: response})
}

func (c *LinkController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	linkId := ctx.Params("linkId")

	request := &model.DeleteLinkRequest{
		Username: auth.ID,
		ID:       linkId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to deleting link")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
