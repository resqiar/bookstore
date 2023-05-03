package inputs

type AddToCartInput struct {
	UserID   int `json:"userID" validate:"required"`
	BookID   int `json:"bookID" validate:"required"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}
