package router

import (
	"log"

	authrouter "github.com/rhuancaetano/enceladus/internal/models/auth/router"
	userrouter "github.com/rhuancaetano/enceladus/internal/models/user/router"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func Router() *router.Router {
	rootRouter, err := router.New("/api/v1")
	if err != nil {
		log.Fatal(err.Error())
	}
	if r, e := userrouter.GetRouter(); e == nil {
		rootRouter.Use("/users", r)
	} else {
		log.Fatal(e.Error())
	}
	if r, e := authrouter.GetRouter(); e == nil {
		rootRouter.Use("/auth", r)
	} else {
		log.Fatal(e.Error())
	}

	return rootRouter
}
