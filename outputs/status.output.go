package outputs

type StatusOutput struct {
	// Status code of the error.
	//
	// example: 200, 400, 401
	Status int `json:"status"`
}
