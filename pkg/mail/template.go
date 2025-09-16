package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

func RenderTemplate(templateName string, data any) (string, error) {
	tmplPath := filepath.Join("pkg", "mail", "templates", templateName)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", templateName, err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	return body.String(), nil
}
