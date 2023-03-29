package outputs

// TokenOutput represents the output data containing a JWT token.
type TokenOutput struct {
	// Status code of the response.
	//
	// example: 200, 401
	Status int `json:"status"`

	// JWT token generated for the user.
	//
	// example: eyJhbGciOiJIUzI1NiIsInR5.XXX.XXX
	Token string `json:"token"`
}
