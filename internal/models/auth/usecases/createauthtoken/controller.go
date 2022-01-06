package createauthtoken

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rhuancaetano/enceladus/internal/models/auth/dtos"
	"github.com/rhuancaetano/enceladus/internal/models/auth/repo"
	"github.com/rhuancaetano/enceladus/internal/request"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func CreateAuthTokenController(res http.ResponseWriter, req *http.Request, ctx *router.Context) {
	res.Header().Add("content-type", "application/json")

	if hasValidContentType := request.HasValidContentType(res, req); !hasValidContentType {
		return
	}

	var authData dtos.CreateAuthTokenDTO
	if success := request.GetStructFromJsonBody(res, req, &authData); !success {
		return
	}

	createAuthTokenUseCase := NewCreateAuthTokenUseCase(repo.NewPostgresAuthRepo())
	authTokens, err := createAuthTokenUseCase.Execute(&authData)
	if err != nil {
		res.WriteHeader(request.HandleErrorStatus(err.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, err.Message)))
		return
	}

	jsonTokens, parserError := json.Marshal(authTokens)
	if parserError != nil {
		log.Println(parserError)

		usecaseError := usecase.ServerError()
		res.WriteHeader(request.HandleErrorStatus(usecaseError.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, usecaseError.Message)))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonTokens)
}
