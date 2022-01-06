package getviewer

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/rhuancaetano/enceladus/internal/config"
	"github.com/rhuancaetano/enceladus/internal/models/auth/dtos"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
)

type Repo interface {
	GetViewerByUserId(id string) (*dtos.ViewerDTO, error)
}

type GetViewerUseCase struct {
	repo Repo
}

func NewGetViewerUseCase(r Repo) *GetViewerUseCase {
	return &GetViewerUseCase{
		repo: r,
	}
}

func (uc *GetViewerUseCase) execute(accessToken string) (*dtos.ViewerDTO, *usecase.UseCaseError) {
	env := config.GetEnvironment()
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(env.JWT_SECRET), nil
	})
	if err != nil {
		return nil, usecase.BadRequestError("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, usecase.BadRequestError("invalid refresh token")
	}

	userId := claims["userId"].(string)

	viewer, err := uc.repo.GetViewerByUserId(userId)
	if err != nil {
		return nil, usecase.ServerError()
	}

	return viewer, nil
}
