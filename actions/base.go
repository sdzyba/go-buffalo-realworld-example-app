package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"

	"github.com/sdzyba/go-buffalo-realworld-example-app/errmap"
	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
	"github.com/sdzyba/go-buffalo-realworld-example-app/services"
)

type UserService interface {
	Create(params *services.UserCreateParams) (*models.User, error)
	Login(params *services.UserLoginParams) (*models.User, error)
}

var userService UserService

func init() {
	userService = &services.User{}
}

func handleError(err error, c buffalo.Context) error {
	switch e := errors.Cause(err).(type) {
	case *errmap.Errs:
		return c.Render(http.StatusUnprocessableEntity, r.JSON(e))
	default:
		return errors.WithStack(err)
	}
}
