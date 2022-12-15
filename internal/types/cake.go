package types

type CakeRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float32 `json:"rating" validate:"required,number"`
	Image       string  `json:"image" validate:"required,url"`
}

type CakeResponse struct {
	ID          int64   `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Rating      float32 `json:"rating,omitempty"`
	Image       string  `json:"image,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}

type CakesResponse struct {
	Cakes            []CakeResponse `json:"cakes"`
	TotalData        int64          `json:"total_data"`
	TotalDataPerPage int64          `json:"total_data_per_page"`
}

type Cake struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
