package services

import (
	english "github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/sdzyba/go-buffalo-realworld-example-app/errmap"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/en"
	"strings"

	"github.com/sdzyba/go-buffalo-realworld-example-app/models"
	"github.com/sdzyba/go-buffalo-realworld-example-app/repositories"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

var validate *validator.Validate
var trans ut.Translator
var userRepository UserRepository

func init() {
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	en.RegisterDefaultTranslations(validate, trans)

	userRepository = &repositories.User{}
}

func newValidationError(err error) *errmap.Errs {
	errs := errmap.NewErrs()
	for _, e := range err.(validator.ValidationErrors) {
		errs.Map[strings.ToLower(e.Field())] = strings.TrimPrefix(e.Translate(trans), e.Field()+" ")
	}

	return errs
}
