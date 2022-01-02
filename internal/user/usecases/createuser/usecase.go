package createuser

import "github.com/rhuancaetano/enceladus/internal/user/dtos"

type CreateUserUseCase struct {
	repo *dtos.Repo
}

func NewCreateUserUseCase(repo dtos.Repo) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: &repo,
	}
}

func (uc *CreateUserUseCase) execute(userData *dtos.CreateUserDTO) error {

	return nil
}
