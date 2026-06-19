package service

import (
	"errors"
	"net/mail"

	"github.com/VelVit24/models"
)

func (s *Service) SendEmail(request *models.ContactMail) (int, error) {
	if request.Name == "" {
		return 400, errors.New("name required")
	}
	if request.Content == "" {
		return 400, errors.New("content required")
	}
	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return 400, errors.New("bad email format")
	}

	// отправка
	// err :=
	// if err != nil {
	// 	return 500, errors.New("failed to send mail")
	// }
}
