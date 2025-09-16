package constants

const (
	ErrMissingAuthHeader  = "Authorization token is required"
	ErrInvalidTokenFormat = "Invalid token format. Expected 'Bearer <token>'"
	ErrInvalidToken       = "Invalid or expired token"
)
