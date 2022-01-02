package router

import (
	"log"

	userrouter "github.com/rhuancaetano/enceladus/internal/user/router"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func Router() *router.Router {
	rootRouter, err := router.New("/api/v1")
	if err != nil {
		log.Fatal(err.Error())
	}
	if r, e := userrouter.GetRouter(); e == nil {
		rootRouter.Use("/user", r)
	} else {
		log.Fatal(e.Error())
	}

	return rootRouter
}
