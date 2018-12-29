package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/mw-tokenauth"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

// TODO: buffalo returns stacktrace on http error
//       check if the same in production environment

// TODO: output error with trace to stdout

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
		})

		uh := &Users{}

		app.Use(paramlogger.ParameterLogger)

		authMW := tokenauth.New(tokenauth.Options{
			SignMethod: jwt.SigningMethodHS256,
			GetKey: func(jwt.SigningMethod) (interface{}, error) {
				return []byte("!!SECRET!!"), nil
			},
			AuthScheme: "Token",
		})
		app.Use(authMW)
		app.Middleware.Skip(authMW, uh.Create, uh.Login)
		// app.Middleware.Skip(authMW, SignUp, Login)

		// app.GET("/", HomeHandler)
		api := app.Group("/api")
		guestUsers := api.Group("/users")
		guestUsers.POST("", uh.Create)
		guestUsers.POST("/login", uh.Login)

		// user := api.Group("/user")
		// user.GET("", h.CurrentUser)
		// user.PUT("", h.UpdateUser)
	}

	return app
}
