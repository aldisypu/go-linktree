package model

type UrlResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateUrlRequest struct {
	Username string `json:"-" validate:"required"`
	Title    string `json:"title" validate:"required,max=100"`
	Url      string `json:"url" validate:"required,max=2048"`
}

type UpdateUrlRequest struct {
	Username string `json:"-" validate:"required"`
	ID       string `json:"-" validate:"required,max=100,uuid"`
	Title    string `json:"title" validate:"required,max=100"`
	Url      string `json:"url" validate:"required,max=2048"`
}

type GetUrlRequest struct {
	Username string `json:"-" validate:"required"`
	ID       string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteUrlRequest struct {
	Username string `json:"-" validate:"required"`
	ID       string `json:"-" validate:"required,max=100,uuid"`
}
