package response

import "time"

type RegisterResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
