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

type UrlUseCase struct {
	DB            *gorm.DB
	Log           *logrus.Logger
	Validate      *validator.Validate
	UrlRepository *repository.UrlRepository
}

func NewUrlUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	urlRepository *repository.UrlRepository) *UrlUseCase {
	return &UrlUseCase{
		DB:            db,
		Log:           logger,
		Validate:      validate,
		UrlRepository: urlRepository,
	}
}

func (c *UrlUseCase) Create(ctx context.Context, request *model.CreateUrlRequest) (*model.UrlResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	url := &entity.Url{
		ID:       uuid.New().String(),
		Title:    request.Title,
		Url:      request.Url,
		Username: request.Username,
	}

	if err := c.UrlRepository.Create(tx, url); err != nil {
		c.Log.WithError(err).Error("failed to creating url")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UrlToResponse(url), nil
}

func (c *UrlUseCase) Update(ctx context.Context, request *model.UpdateUrlRequest) (*model.UrlResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	url := new(entity.Url)
	if err := c.UrlRepository.FindByIdAndUsername(tx, url, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to getting url")
		return nil, fiber.ErrNotFound
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	url.Title = request.Title
	url.Url = request.Url

	if err := c.UrlRepository.Update(tx, url); err != nil {
		c.Log.WithError(err).Error("failed to updating url")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UrlToResponse(url), nil
}

func (c *UrlUseCase) Get(ctx context.Context, request *model.GetUrlRequest) (*model.UrlResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	url := new(entity.Url)
	if err := c.UrlRepository.FindByIdAndUsername(tx, url, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to getting url")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UrlToResponse(url), nil
}

func (c *UrlUseCase) Delete(ctx context.Context, request *model.DeleteUrlRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validating request body")
		return fiber.ErrBadRequest
	}

	url := new(entity.Url)
	if err := c.UrlRepository.FindByIdAndUsername(tx, url, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to getting url")
		return fiber.ErrNotFound
	}

	if err := c.UrlRepository.Delete(tx, url); err != nil {
		c.Log.WithError(err).Error("failed to deleting url")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}
