package deleteuser

import (
	"regexp"

	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
)

type Repo interface {
	DeleteUser(id string) error
}

type DeleteUserUseCase struct {
	repo Repo
}

func NewDeleteUserUseCase(r Repo) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repo: r,
	}
}

func (uc *DeleteUserUseCase) Execute(id string) *usecase.UseCaseError {
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
