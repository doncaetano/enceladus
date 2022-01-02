package router

import (
	"log"

	"github.com/rhuancaetano/enceladus/internal/user/usecases/createuser"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

type UserRouter struct {
	Router *router.Router
}

func GetRouter() (*router.Router, error) {
	router, err := router.New("/")
	if err != nil {
		return nil, err
	}

	if err = router.Post("/", createuser.CreateUserController); err != nil {
		log.Fatal(err.Error())
	}

	return router, nil
}
