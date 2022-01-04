package router

import (
	"log"

	"github.com/rhuancaetano/enceladus/internal/auth/usecases/createauthtoken"
	"github.com/rhuancaetano/enceladus/internal/auth/usecases/getviewer"
	"github.com/rhuancaetano/enceladus/internal/auth/usecases/updateauthtoken"
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

	if err = router.Post("/", createauthtoken.CreateAuthTokenController); err != nil {
		log.Fatal(err.Error())
	}
	if err = router.Put("/", updateauthtoken.UpdateAuthTokenController); err != nil {
		log.Fatal(err.Error())
	}
	if err = router.Get("/:accessToken", getviewer.GetViewerController); err != nil {
		log.Fatal(err.Error())
	}

	return router, nil
}
