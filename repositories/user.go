package repositories

import (
	"github.com/sdzyba/go-buffalo-realworld-example-app/errutil"
	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
)

type User struct {
}

func (u *User) Create(user *models.User) (*models.User, error) {
	err := db.Create(user)
	if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_idx\"" {
		errResp := &errutil.ErrorResponse{}
		errResp.Errors = make(map[string]interface{})
		errResp.Errors["email"] = "already taken"

		return nil, errResp
	}
	if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_idx\"" {
		errResp := &errutil.ErrorResponse{}
		errResp.Errors = make(map[string]interface{})
		errResp.Errors["username"] = "already taken"

		return nil, errResp
	}

	return user, err
}
