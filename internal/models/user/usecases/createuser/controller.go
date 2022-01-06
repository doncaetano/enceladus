package createuser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rhuancaetano/enceladus/internal/models/user/dtos"
	"github.com/rhuancaetano/enceladus/internal/models/user/repo"
	"github.com/rhuancaetano/enceladus/internal/request"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func CreateUserController(res http.ResponseWriter, req *http.Request, ctx *router.Context) {
	res.Header().Add("content-type", "application/json")

	if hasValidContentType := request.HasValidContentType(res, req); !hasValidContentType {
		return
	}

	var userData dtos.CreateUserDTO
	if success := request.GetStructFromJsonBody(res, req, &userData); !success {
		return
	}

	createUserUseCase := NewCreateUserUseCase(repo.NewPostgresUserRepo())
	user, err := createUserUseCase.Execute(&userData)
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
