package services

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
)

type UserCreateParams struct {
	User *struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email"    validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user" validate:"required"`
}

type UserLoginParams struct {
	User *struct {
		Email    string `json:"email"    validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user" validate:"required"`
}

type User struct {
}

func (u *User) Create(params *UserCreateParams) (*models.User, error) {
	err := validate.Struct(params)
	if err != nil {
		return nil, newValidationError(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(params.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user := &models.User{
		Email:             params.User.Email,
		EncryptedPassword: string(password),
		Username:          params.User.Username,
	}

	return userRepository.Create(user)
}

func (u *User) Login(params *UserLoginParams) (*models.User, error) {
	err := validate.Struct(params)
	if err != nil {
		return nil, newValidationError(err)
	}

	user, err := userRepository.FindByEmail(params.User.Email)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if user == nil {
		return nil, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(params.User.Password))
	if err != nil {
		return nil, nil
	}

	return user, nil
}
