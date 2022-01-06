package createauthtoken

import (
	"log"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/rhuancaetano/enceladus/internal/config"
	"github.com/rhuancaetano/enceladus/internal/models/auth/dtos"
	"github.com/rhuancaetano/enceladus/internal/shared/usecase"
	"github.com/rhuancaetano/enceladus/internal/utils/encrypt"
)

type Repo interface {
	CreateRefreshToken(data *dtos.CreateRefreshTokenDTO) (*dtos.CreatedTokenDTO, error)
	CreateAccessToken(data *dtos.CreateAccessTokenDTO) (*dtos.CreatedTokenDTO, error)
	GetUserByEmail(email string) (*dtos.UserDTO, error)
	DeactivateUserTokensByUserId(userId string) error
}

type CreateAuthTokenUseCase struct {
	repo Repo
}

func NewCreateAuthTokenUseCase(r Repo) *CreateAuthTokenUseCase {
	return &CreateAuthTokenUseCase{
		repo: r,
	}
}

func (uc *CreateAuthTokenUseCase) execute(data *dtos.CreateAuthTokenDTO) (*dtos.AuthTokenDTO, *usecase.UseCaseError) {
	reg := regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")
	if data.Email == "" || !reg.MatchString(data.Email) {
		return nil, usecase.BadRequestError("invalid user email")
	}
	if len(data.Password) < 8 || len(data.Password) > 16 {
		return nil, usecase.BadRequestError("invalid password")
	}

	user, err := uc.repo.GetUserByEmail(data.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	if user == nil || !encrypt.IsCorrectPassword(data.Password, user.Password) {
		return nil, usecase.BadRequestError("email and password does not match")
	}

	if err := uc.repo.DeactivateUserTokensByUserId(user.Id); err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	refreshToken, err := uc.repo.CreateRefreshToken(&dtos.CreateRefreshTokenDTO{
		UserId: user.Id,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	expiresAt := time.Now().Add(time.Hour * 3)
	accessToken, err := uc.repo.CreateAccessToken(&dtos.CreateAccessTokenDTO{
		UserId:         user.Id,
		RefreshTokenId: refreshToken.Id,
		ExpiresAt:      expiresAt,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	env := config.GetEnvironment()

	jwtRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     refreshToken.Id,
		"userId": user.Id,
		"email":  user.Email,
	})
	refreshTokenString, err := jwtRefreshToken.SignedString([]byte(env.JWT_SECRET))
	if err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     accessToken.Id,
		"userId": user.Id,
		"email":  user.Email,
		"exp":    expiresAt.Unix(),
	})
	accessTokenString, err := jwtAccessToken.SignedString([]byte(env.JWT_SECRET))
	if err != nil {
		log.Println(err.Error())
		return nil, usecase.ServerError()
	}

	return &dtos.AuthTokenDTO{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
