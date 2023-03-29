package config

const (
	// JWT maximum valid age
	TOKEN_EXPIRATION_TIME = 60 // value should be in minutes

	// This value indicates the barrier when the server
	// should renew the token to a new valid token before
	// the X expiration time.
	REFRESH_TOKEN_THRESHOLD = 10 // value should be in minutes
)
