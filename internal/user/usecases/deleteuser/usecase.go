package deleteuser

import (
	"regexp"

	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/internal/user/dtos"
)

type DeleteUserUseCase struct {
	repo dtos.Repo
}

func NewDeleteUserUseCase(r dtos.Repo) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: r,
	}
}

func (uc *DeleteUserUseCase) execute(id string) *usecase.UseCaseError {
	reg := regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
	if id == "" || !reg.MatchString(id) {
		return usecase.BadRequestError("invalid user id")
	}

	err := uc.repo.DeleteUser(id)
	if err != nil {
		return usecase.ServerError()
	}

	return nil
}
