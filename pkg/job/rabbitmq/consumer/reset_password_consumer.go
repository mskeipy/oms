package consumer

import (
	"dropx/pkg/mail"
	"encoding/json"
	"html/template"
	"log"
)

type ResetPasswordPayload struct {
	Email     string       `json:"email"`
	Name      string       `json:"name"`
	ResetLink template.URL `json:"reset_link"`
}

func HandleResetPasswordMessage(msg []byte) {
	var payload ResetPasswordPayload
	if err := json.Unmarshal(msg, &payload); err != nil {
		log.Println("Failed to parse message:", err)
		return
	}

	body, err := mail.RenderTemplate("reset.html", payload)
	if err != nil {
		log.Println("Failed to render template:", err)
		return
	}

	mail.Send(payload.Email, "Reset your password", body)
}
