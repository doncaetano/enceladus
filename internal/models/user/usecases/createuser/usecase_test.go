package createuser_test

import (
	"errors"
	"testing"

	"github.com/rhuancaetano/enceladus/internal/models/user/dtos"
	. "github.com/rhuancaetano/enceladus/internal/models/user/usecases/createuser"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
)

type R struct {
	CreateUserReturn  *dtos.UserDTO
	CreateUserError   error
	FindByEmailReturn *dtos.UserDTO
	FindByEmailError  error
}

func (r *R) CreateUser(data *dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	return r.CreateUserReturn, r.CreateUserError
}

func (r *R) FindByEmail(email string) (*dtos.UserDTO, error) {
	return r.FindByEmailReturn, r.FindByEmailError
}

func TestCreateUserUseCaseEmptyFirstName(t *testing.T) {
	useCase := NewCreateUserUseCase(&R{})
	_, err := useCase.Execute(&dtos.CreateUserDTO{})

	if err == nil {
		t.Error("should return an error if the first name is empty")
		t.FailNow()
	}

	if err.Type != usecase.BAD_REQUEST {
		t.Error("should return a bad request error")
		t.FailNow()
	}

	if err.Message != "invalid user first name" {
		t.Error("message should be 'invalid user first name'")
	}
}

func TestCreateUserUseCaseEmptyLastName(t *testing.T) {
	useCase := NewCreateUserUseCase(&R{})
	_, err := useCase.Execute(&dtos.CreateUserDTO{
		FirstName: "first name",
	})

	if err == nil {
		t.Error("should return an error if the last name is empty")
		t.FailNow()
	}

	if err.Type != usecase.BAD_REQUEST {
		t.Error("should return a bad request error")
		t.FailNow()
	}

	if err.Message != "invalid user last name" {
		t.Error("message should be 'invalid user last name'")
	}
}

var invalidEmails = []string{
	"",
	"email",
	"email@",
	"email.com",
	"email@@email.com",
	"email@email@com",
}

func TestCreateUserUseCaseValidEmail(t *testing.T) {
	useCase := NewCreateUserUseCase(&R{})
	for _, email := range invalidEmails {
		_, err := useCase.Execute(&dtos.CreateUserDTO{
			FirstName: "first name",
			LastName:  "last name",
			Email:     email,
		})
		if err == nil {
			t.Error("should return an error if the email is invalid")
			t.FailNow()
		}

		if err.Type != usecase.BAD_REQUEST {
			t.Error("should return a bad request error")
			t.FailNow()
		}

		if err.Message != "invalid user email" {
			t.Error("message should be 'invalid user email'")
		}
	}

}

func TestCreateUserUseCaseFindByEmailError(t *testing.T) {
	useCase := NewCreateUserUseCase(&R{
		FindByEmailError: errors.New(""),
	})
	_, err := useCase.Execute(&dtos.CreateUserDTO{
		FirstName: "first name",
		LastName:  "last name",
		Email:     "email@email.com",
	})

	if err == nil {
		t.Error("should return an error")
		t.FailNow()
	}

	if err.Type != usecase.SERVER_ERROR {
		t.Error("should return a server error")
		t.FailNow()
	}

	if err.Message != usecase.ServerError().Message {
		t.Error("should return the default server error message")
	}
}

func TestCreateUserUseCaseFindByEmailUser(t *testing.T) {
	useCase := NewCreateUserUseCase(&R{
		FindByEmailReturn: &dtos.UserDTO{},
	})
	_, err := useCase.Execute(&dtos.CreateUserDTO{
		FirstName: "first name",
		LastName:  "last name",
		Email:     "email@email.com",
	})

	if err == nil {
		t.Error("should return an error")
		t.FailNow()
	}

	if err.Type != usecase.BAD_REQUEST {
		t.Error("should return a bad request error")
		t.FailNow()
	}

	if err.Message != "the email is already taken" {
		t.Error("message should be 'the email is already taken'")
	}
}

var invalidPasswords = []string{
	"",
	"1",
	"12",
	"123",
	"1234",
	"12345",
	"123456",
	"1234567",
	"0123456789ABCDEF0",
}

func TestCreateUserUseCasePasswordSize(t *testing.T) {
	useCase := NewCreateUserUseCase(&R{})
	for _, password := range invalidPasswords {
		_, err := useCase.Execute(&dtos.CreateUserDTO{
			FirstName: "first name",
			LastName:  "last name",
			Email:     "email@email.com",
			Password:  password,
		})

		if err == nil {
			t.Error("should return an error")
			t.FailNow()
		}

		if err.Type != usecase.BAD_REQUEST {
			t.Error("should return a bad request error")
			t.FailNow()
		}

		if err.Message != "invalid password" {
			t.Error("message should be 'invalid password'")
		}
	}
}

func TestCreateUserUseCaseCreateUserError(t *testing.T) {
	useCase := NewCreateUserUseCase(&R{
		CreateUserError: errors.New(""),
	})
	_, err := useCase.Execute(&dtos.CreateUserDTO{
		FirstName: "first name",
		LastName:  "last name",
		Email:     "email@email.com",
		Password:  "password123",
	})

	if err == nil {
		t.Error("should return an error")
		t.FailNow()
	}

	if err.Type != usecase.SERVER_ERROR {
		t.Error("should return a server error")
		t.FailNow()
	}

	if err.Message != usecase.ServerError().Message {
		t.Error("should return the default server error message")
	}
}

func TestCreateUserUseCaseSuccess(t *testing.T) {
	user := &dtos.UserDTO{
		Id:        "650b38a0-4453-40b9-872f-cf460dc4f005",
		FirstName: "first name",
		LastName:  "last name",
		Email:     "email@email",
		IsActive:  true,
		CreatedAt: "2022-01-06T04:32:15.030Z",
		UpdatedAt: "2022-01-06T04:32:15.030Z",
	}

	useCase := NewCreateUserUseCase(&R{
		CreateUserReturn: user,
	})
	createdUser, err := useCase.Execute(&dtos.CreateUserDTO{
		FirstName: "first name",
		LastName:  "last name",
		Email:     "email@email.com",
		Password:  "password123",
	})

	if err != nil {
		t.Error("should not return an error")
		t.FailNow()
	}

	if createdUser == nil {
		t.Error("should return the created user")
		t.FailNow()
	}

	if createdUser.Id != user.Id &&
		createdUser.FirstName != user.FirstName &&
		createdUser.LastName != user.LastName &&
		createdUser.Email != user.Email &&
		createdUser.IsActive != user.IsActive &&
		createdUser.CreatedAt != user.CreatedAt &&
		createdUser.UpdatedAt != user.UpdatedAt {
		t.Error("should return the created user data")
	}
}
