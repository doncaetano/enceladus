package getuser

import (
	"regexp"

	"github.com/rhuancaetano/enceladus/internal/models/user/dtos"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
)

type Repo interface {
	FindById(id string) (*dtos.UserDTO, error)
}

type GetUserUseCase struct {
	repo Repo
}

func NewGetUserUseCase(r Repo) *GetUserUseCase {
	return &GetUserUseCase{
		repo: r,
	}
}

func (uc *GetUserUseCase) execute(id string) (*dtos.UserDTO, *usecase.UseCaseError) {
	reg := regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
	if id == "" || !reg.MatchString(id) {
		return nil, usecase.BadRequestError("invalid user id")
	}

	user, err := uc.repo.FindById(id)
	if err != nil {
		return nil, usecase.ServerError()
	}

	return user, nil
}
