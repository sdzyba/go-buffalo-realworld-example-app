package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/sdzyba/go-buffalo-realworld-example-app/errutil"
	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
	"github.com/sdzyba/go-buffalo-realworld-example-app/services"
)

type UserService interface {
	Create(params *services.UserCreateParams) (*models.User, error)
}

var userService UserService

func init() {
	userService = &services.User{}
}

func handleError(err error, c buffalo.Context) error {
	switch e := errors.Cause(err).(type) {
	case *errutil.ErrorResponse:
		return c.Render(http.StatusUnprocessableEntity, r.JSON(e))
	case validator.ValidationErrors:
		errResp := &errutil.ErrorResponse{}
		errResp.Errors = make(map[string]interface{})
		for _, v := range e {
			errResp.Errors[v.Field()] = fmt.Sprintf("%v", v.Tag())
		}

		return c.Render(http.StatusUnprocessableEntity, r.JSON(errResp))
	default:
		return errors.WithStack(err)
	}
}
