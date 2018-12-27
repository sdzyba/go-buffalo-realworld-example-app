package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/nulls"
	"github.com/pkg/errors"
	"net/http"

	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
	"github.com/sdzyba/go-buffalo-realworld-example-app/services"
)

type Users struct {
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

func newUserResponse(u *models.User) (*userResponse, error) {
	r := &userResponse{}
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Bio = u.Bio
	r.User.Image = u.Image

	token, err := u.GenerateToken()
	if err != nil {
		return nil, err
	}
	r.User.Token = token

	return r, nil
}
