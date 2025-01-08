package converter

import (
	"go-linktree/internal/entity"
	"go-linktree/internal/model"
)

func UrlToResponse(url *entity.Url) *model.UrlResponse {
	return &model.UrlResponse{
		ID:        url.ID,
		Title:     url.Title,
		Url:       url.Url,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}
}
