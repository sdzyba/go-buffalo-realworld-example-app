package services

import (
	"gopkg.in/go-playground/validator.v9"

	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
	"github.com/sdzyba/go-buffalo-realworld-example-app/repositories"
)

type UserRepository interface {
	Create(*models.User) (*models.User, error)
}

var validate *validator.Validate
var userRepository UserRepository

func init() {
	validate = validator.New()
	userRepository = &repositories.User{}
}
