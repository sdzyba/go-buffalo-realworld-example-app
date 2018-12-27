package repositories

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
)

var db *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	db, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}
