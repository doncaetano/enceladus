package createuser

import (
	"fmt"
	"net/http"

	"github.com/rhuancaetano/enceladus/internal/request"
	"github.com/rhuancaetano/enceladus/internal/user/dtos"
	"github.com/rhuancaetano/enceladus/internal/user/repo"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func CreateUserController(res http.ResponseWriter, req *http.Request, ctx *router.Context) {
	if hasValidContentType := request.HasValidContentType(res, req); hasValidContentType {
		return
	}

	var userData dtos.CreateUserDTO
	if success := request.GetStructFromJsonBody(res, req, &userData); !success {
		return
	}

	postgresUserRepo := repo.NewPostgresUserRepo()
	createUserUseCase := NewCreateUserUseCase(postgresUserRepo)
	err := createUserUseCase.execute(&userData)
	if err != nil {
		fmt.Println(err.Error())
	}

	res.Header().Add("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	// res.Write(jsonBytes)
}
