package repository

import (
	"go-linktree/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LinkRepository struct {
	Repository[entity.Link]
	Log *logrus.Logger
}

func NewLinkRepository(log *logrus.Logger) *LinkRepository {
	return &LinkRepository{
		Log: log,
	}
}

func (r *LinkRepository) FindByIdAndUsername(db *gorm.DB, link *entity.Link, id string, username string) error {
	return db.Where("id = ? AND username = ?", id, username).Take(link).Error
}
