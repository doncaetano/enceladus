package getuser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rhuancaetano/enceladus/internal/models/user/repo"
	"github.com/rhuancaetano/enceladus/internal/request"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func GetUserController(res http.ResponseWriter, req *http.Request, ctx *router.Context) {
	res.Header().Add("content-type", "application/json")

	id, hasValue := ctx.Params["id"]
	if !hasValue {
		log.Println("could not get parameter 'id' on GetUserController")

		usecaseError := usecase.ServerError()
		res.WriteHeader(request.HandleErrorStatus(usecaseError.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, usecaseError.Message)))
		return
	}

	getUserUseCase := NewGetUserUseCase(repo.NewPostgresUserRepo())
	user, err := getUserUseCase.Execute(id)
	if err != nil {
		res.WriteHeader(request.HandleErrorStatus(err.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, err.Message)))
		return
	}

	jsonUser, parserError := json.Marshal(user)
	if parserError != nil {
		log.Println(parserError)

		usecaseError := usecase.ServerError()
		res.WriteHeader(request.HandleErrorStatus(usecaseError.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, usecaseError.Message)))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonUser)

}
