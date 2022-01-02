package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rhuancaetano/enceladus/cmd/api/router"
	"github.com/rhuancaetano/enceladus/internal/config"
	"github.com/rhuancaetano/enceladus/internal/repo"
)

func main() {
	env := config.GetEnvironment()
	repo.InitDatabase()

	mux := http.NewServeMux()
	r := router.Router()
	r.AddHandlersIntoServeMux(mux)

	log.Println(fmt.Sprintf("Starting server on http://localhost:%s", env.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.PORT), mux))
}
