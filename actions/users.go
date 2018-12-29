package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/nulls"
	"github.com/pkg/errors"

	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
	"github.com/sdzyba/go-buffalo-realworld-example-app/services"
)

type Users struct {
}

func (u *Users) Create(c buffalo.Context) error {
	params := &services.UserCreateParams{}

	if err := c.Bind(params); err != nil {
		return errors.WithStack(err)
	}

	user, err := userService.Create(params)
	if err != nil {
		return handleError(err, c)
	}

	response, err := newUserResponse(user)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(http.StatusCreated, r.JSON(response))
}

func (u *Users) Login(c buffalo.Context) error {
	params := &services.UserLoginParams{}

	if err := c.Bind(params); err != nil {
		return errors.WithStack(err)
	}

	user, err := userService.Login(params)
	if err != nil {
		return handleError(err, c)
	}
	if user == nil {
		return c.Render(http.StatusUnauthorized, r.JSON("Unauthorized"))
	}

	response, err := newUserResponse(user)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

type userResponse struct {
	User struct {
		Username string       `json:"username"`
		Email    string       `json:"email"`
		Bio      nulls.String `json:"bio"`
		Image    nulls.String `json:"image"`
		Token    string       `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *models.User) (*userResponse, error) {
	ur := &userResponse{}
	ur.User.Username = u.Username
	ur.User.Email = u.Email
	ur.User.Bio = u.Bio
	ur.User.Image = u.Image

	token, err := u.GenerateToken()
	if err != nil {
		return nil, err
	}
	ur.User.Token = token

	return ur, nil
}
