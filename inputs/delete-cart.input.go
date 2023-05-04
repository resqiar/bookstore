package inputs

type DeleteCartInput struct {
	UserID int `json:"userID" validate:"required"`
	BookID int `json:"bookID" validate:"required"`
}
