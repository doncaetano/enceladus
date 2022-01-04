package getviewer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rhuancaetano/enceladus/internal/models/auth/repo"
	"github.com/rhuancaetano/enceladus/internal/request"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/pkg/router"
)

func GetViewerController(res http.ResponseWriter, req *http.Request, ctx *router.Context) {
	res.Header().Add("content-type", "application/json")

	accessToken, hasValue := ctx.Params["accessToken"]
	if !hasValue {
		log.Println("could not get parameter 'accessToken' on GetViewerController")

		usecaseError := usecase.ServerError()
		res.WriteHeader(request.HandleErrorStatus(usecaseError.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, usecaseError.Message)))
		return
	}

	getViewerUseCase := NewGetViewerUseCase(repo.NewPostgresAuthRepo())
	viewer, err := getViewerUseCase.execute(accessToken)
	if err != nil {
		res.WriteHeader(request.HandleErrorStatus(err.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, err.Message)))
		return
	}

	jsonViewer, parserError := json.Marshal(viewer)
	if parserError != nil {
		log.Println(parserError)

		usecaseError := usecase.ServerError()
		res.WriteHeader(request.HandleErrorStatus(usecaseError.Type))
		res.Write([]byte(fmt.Sprintf(`{ "error": "%s" }`, usecaseError.Message)))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonViewer)
}
