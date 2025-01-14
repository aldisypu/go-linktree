package converter

import (
	"go-linktree/internal/entity"
	"go-linktree/internal/model"
)

func LinkToResponse(link *entity.Link) *model.LinkResponse {
	return &model.LinkResponse{
		ID:        link.ID,
		Title:     link.Title,
		Url:       link.Url,
		CreatedAt: link.CreatedAt,
		UpdatedAt: link.UpdatedAt,
	}
}
