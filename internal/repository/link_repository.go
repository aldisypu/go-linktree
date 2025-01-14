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

func (r *LinkRepository) FindAllByUsername(tx *gorm.DB, username string) ([]entity.Link, error) {
	var links []entity.Link
	if err := tx.Where("username = ?", username).Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}
