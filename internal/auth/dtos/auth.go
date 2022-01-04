package dtos

import "time"

type AuthTokenDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type CreateAuthTokenDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAuthDTO struct {
	RefreshToken string `json:"refreshToken"`
}

type UserDTO struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ViewerDTO struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreatedTokenDTO struct {
	Id string `json:"id"`
}

type CreateRefreshTokenDTO struct {
	UserId string `json:"userId"`
}

type CreateAccessTokenDTO struct {
	UserId         string    `json:"userId"`
	RefreshTokenId string    `json:"refreshTokenId"`
	ExpiresAt      time.Time `json:"expiresAt"`
}

type Repo interface {
	CreateRefreshToken(data *CreateRefreshTokenDTO) (*CreatedTokenDTO, error)
	CreateAccessToken(data *CreateAccessTokenDTO) (*CreatedTokenDTO, error)
	GetUserByEmail(email string) (*UserDTO, error)
	DeactivateUserTokensByUserId(userId string) error
	DeactivateAccessTokenByRefreshTokenId(refreshTokenId string) error
	CheckIfActiveRefreshTokenExist(refreshTokenId string) (bool, error)
	GetViewerByUserId(id string) (*ViewerDTO, error)
}
