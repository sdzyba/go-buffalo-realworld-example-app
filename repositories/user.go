package repositories

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/sdzyba/go-buffalo-realworld-example-app/errmap"
	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
)

type User struct {
}

func (u *User) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := db.Where("email = ?", email).First(user)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (u *User) Create(user *models.User) (*models.User, error) {
	err := db.Create(user)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_idx\"" {
			errs := errmap.NewErrs()
			errs.Map["email"] = "already taken"

			return nil, errs
		}

		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_idx\"" {
			errs := errmap.NewErrs()
			errs.Map["username"] = "already taken"

			return nil, errs
		}

		return nil, err
	}

	return user, nil
}
