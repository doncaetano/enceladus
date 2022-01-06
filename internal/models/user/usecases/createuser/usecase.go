package createuser

import (
	"log"
	"regexp"

	"github.com/rhuancaetano/enceladus/internal/models/user/dtos"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/internal/utils/encrypt"
)

type Repo interface {
	CreateUser(data *dtos.CreateUserDTO) (*dtos.UserDTO, error)
	FindByEmail(email string) (*dtos.UserDTO, error)
}

type CreateUserUseCase struct {
	repo Repo
}

func NewCreateUserUseCase(r Repo) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: r,
	}
}

func (uc *CreateUserUseCase) Execute(data *dtos.CreateUserDTO) (*dtos.UserDTO, *usecase.UseCaseError) {
	if data.FirstName == "" {
		return nil, usecase.BadRequestError("invalid user first name")
	}
	if data.LastName == "" {
		return nil, usecase.BadRequestError("invalid user last name")
	}
	reg := regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")
	if data.Email == "" || !reg.MatchString(data.Email) {
		return nil, usecase.BadRequestError("invalid user email")
	}
	if user, err := uc.repo.FindByEmail(data.Email); err != nil {
		return nil, usecase.ServerError()
	} else if user != nil {
		return nil, usecase.BadRequestError("the email is already taken")
	}
	if len(data.Password) < 8 || len(data.Password) > 16 {
		return nil, usecase.BadRequestError("invalid password")
	}

	if hashedPassword, err := encrypt.EncryptPassword(data.Password); err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	} else {
		data.Password = hashedPassword
	}

	user, err := uc.repo.CreateUser(data)
	if err != nil {
		return nil, usecase.ServerError()
	}

	return user, nil
}
