package usecase

import (
	"context"
	"go-linktree/internal/entity"
	"go-linktree/internal/model"
	"go-linktree/internal/model/converter"
	"go-linktree/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LinkUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	LinkRepository *repository.LinkRepository
	UserRepository *repository.UserRepository
}

func NewLinkUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	linkRepository *repository.LinkRepository) *LinkUseCase {
	return &LinkUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		LinkRepository: linkRepository,
	}
}

func (c *LinkUseCase) Create(ctx context.Context, request *model.CreateLinkRequest) (*model.LinkResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	link := &entity.Link{
		ID:       uuid.New().String(),
		Title:    request.Title,
		Url:      request.Url,
		Username: request.Username,
	}

	if err := c.LinkRepository.Create(tx, link); err != nil {
		c.Log.WithError(err).Error("failed to creating link")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.LinkToResponse(link), nil
}

func (c *LinkUseCase) Update(ctx context.Context, request *model.UpdateLinkRequest) (*model.LinkResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	link := new(entity.Link)
	if err := c.LinkRepository.FindByIdAndUsername(tx, link, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to getting link")
		return nil, fiber.ErrNotFound
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	link.Title = request.Title
	link.Url = request.Url

	if err := c.LinkRepository.Update(tx, link); err != nil {
		c.Log.WithError(err).Error("failed to updating link")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.LinkToResponse(link), nil
}

func (c *LinkUseCase) Get(ctx context.Context, request *model.GetLinkRequest) (*model.LinkResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	link := new(entity.Link)
	if err := c.LinkRepository.FindByIdAndUsername(tx, link, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to getting link")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.LinkToResponse(link), nil
}

func (c *LinkUseCase) Delete(ctx context.Context, request *model.DeleteLinkRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return fiber.ErrBadRequest
	}

	link := new(entity.Link)
	if err := c.LinkRepository.FindByIdAndUsername(tx, link, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to getting link")
		return fiber.ErrNotFound
	}

	if err := c.LinkRepository.Delete(tx, link); err != nil {
		c.Log.WithError(err).Error("failed to deleting link")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *LinkUseCase) List(ctx context.Context, request *model.ListLinkRequest) ([]model.LinkResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := c.UserRepository.FindByUsername(tx, user, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to find user")
		return nil, fiber.ErrNotFound
	}

	links, err := c.LinkRepository.FindAllByUsername(tx, user.Username)
	if err != nil {
		c.Log.WithError(err).Error("failed to find links")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.LinkResponse, len(links))
	for i, link := range links {
		responses[i] = *converter.LinkToResponse(&link)
	}

	return responses, nil
}
