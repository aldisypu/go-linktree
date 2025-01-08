package repository

import (
	"go-linktree/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UrlRepository struct {
	Repository[entity.Url]
	Log *logrus.Logger
}

func NewUrlRepository(log *logrus.Logger) *UrlRepository {
	return &UrlRepository{
		Log: log,
	}
}

func (r *UrlRepository) FindByIdAndUsername(db *gorm.DB, url *entity.Url, id string, username string) error {
	return db.Where("id = ? AND username = ?", id, username).Take(url).Error
}
