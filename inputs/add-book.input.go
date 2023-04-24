package inputs

type AddBookInput struct {
	Title       string `validate:"required,max=100" json:"title"`
	Description string `validate:"required" json:"description"`
	ImageURL    string `validate:"omitempty,url" json:"imageURL"`
	ReleaseDate string `validate:"required,max=45" json:"releaseDate"`
	Author      string `validate:"required,max=45" json:"author"`
	Price       int    `validate:"required" json:"price"`
}
