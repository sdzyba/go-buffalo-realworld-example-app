package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/pop/nulls"
)

type User struct {
	ID                int64        `db:"id"`
	Email             string       `db:"email"`
	EncryptedPassword string       `db:"encrypted_password"`
	Username          string       `db:"username"`
	Image             nulls.String `db:"image"`
	Bio               nulls.String `db:"bio"`
	CreatedAt         time.Time    `db:"created_at"`
	UpdatedAt         time.Time    `db:"updated_at"`
}

func (u *User) GenerateToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 14).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("!!SECRET!!"))
}
