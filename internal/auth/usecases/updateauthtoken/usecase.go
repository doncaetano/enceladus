package updateauthtoken

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/rhuancaetano/enceladus/internal/auth/dtos"
	"github.com/rhuancaetano/enceladus/internal/config"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
)

type UpdateAuthTokenUseCase struct {
	repo dtos.Repo
}

func NewUpdateAuthTokenUseCase(r dtos.Repo) *UpdateAuthTokenUseCase {
	return &UpdateAuthTokenUseCase{
		repo: r,
	}
}

func (uc *UpdateAuthTokenUseCase) execute(data *dtos.UpdateAuthDTO) (*dtos.AuthTokenDTO, *usecase.UseCaseError) {
	env := config.GetEnvironment()
	token, err := jwt.Parse(data.RefreshToken, func(token *jwt.Token) (interface{}, error) {
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

	refreshTokenId := claims["id"].(string)
	userId := claims["userId"].(string)
	email := claims["email"].(string)

	if ok, err := uc.repo.CheckIfActiveRefreshTokenExist(refreshTokenId); err != nil {
		return nil, usecase.ServerError()
	} else if !ok {
		return nil, usecase.BadRequestError("invalid refresh token")
	}

	if err := uc.repo.DeactivateAccessTokenByRefreshTokenId(refreshTokenId); err != nil {
		return nil, usecase.ServerError()
	}

	expiresAt := time.Now().Add(time.Hour * 3)
	accessToken, err := uc.repo.CreateAccessToken(&dtos.CreateAccessTokenDTO{
		UserId:         userId,
		RefreshTokenId: refreshTokenId,
		ExpiresAt:      expiresAt,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     accessToken.Id,
		"userId": userId,
		"email":  email,
		"exp":    expiresAt.Unix(),
	})
	accessTokenString, err := jwtAccessToken.SignedString([]byte(env.JWT_SECRET))
	if err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	return &dtos.AuthTokenDTO{
		AccessToken:  accessTokenString,
		RefreshToken: data.RefreshToken,
	}, nil
}
