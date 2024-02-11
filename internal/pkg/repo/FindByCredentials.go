package repo

import (
	"errors"

	"github.com/SharanyaSD/PayrollSystem.git/internal/pkg/models"
)

// db call
func FindByCredentials(email, password string) (*models.User, error) {
	//query database - given creds
	if email == "ssd174285@gmail.com" && password == "123456" {
		return &models.User{
			ID:       "501",
			Email:    "ssd174285@gmail.com",
			Password: "123456",
		}, nil
	}
	return nil, errors.New("User Not valid")
}
