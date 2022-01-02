package request

import (
	"net/http"

	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
)

func HandleErrorStatus(err usecase.UseCaseErrorType) int {
	switch err {
	case usecase.SERVER_ERROR:
		return http.StatusInternalServerError
	case usecase.BAD_REQUEST:
		return http.StatusBadRequest
	case usecase.UNAUTHORIZED:
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}
