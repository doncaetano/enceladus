package deleteuser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rhuancaetano/enceladus/internal/request"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/internal/user/repo"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func DeleteUserController(res http.ResponseWriter, req *http.Request, ctx *router.Context) {
	res.Header().Add("content-type", "application/json")

	id, hasValue := ctx.Params["id"]
	if !hasValue {
		log.Println("could not get parameter 'id' on DeleteUserController")

		usecaseError := usecase.ServerError()
		res.WriteHeader(request.HandleErrorStatus(usecaseError.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, usecaseError.Message)))
		return
	}

	deleteUserUseCase := NewDeleteUserUseCase(repo.NewPostgresUserRepo())
	err := deleteUserUseCase.execute(id)
	if err != nil {
		res.WriteHeader(request.HandleErrorStatus(err.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, err.Message)))
		return
	}

	res.WriteHeader(http.StatusOK)
}
