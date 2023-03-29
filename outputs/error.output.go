package outputs

// ErrorOutput represents the response data for an error message.
type ErrorOutput struct {
	// Status code of the error.
	//
	// example: 400, 500
	Status int `json:"status"`

	// Message explaining the error.
	//
	// example: Bad Request for ...
	Message string `json:"message"`
}
