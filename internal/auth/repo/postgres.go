package repo

import (
	"database/sql"
	"log"

	"github.com/rhuancaetano/enceladus/internal/auth/dtos"
	"github.com/rhuancaetano/enceladus/internal/repo"
)

type PostgresAuthRepo struct {
	db *sql.DB
}

func NewPostgresAuthRepo() *PostgresAuthRepo {
	return &PostgresAuthRepo{
		db: repo.GetPostgresDatabase(),
	}
}

func (ur *PostgresAuthRepo) CreateRefreshToken(data *dtos.CreateRefreshTokenDTO) (*dtos.CreatedTokenDTO, error) {
	rows, err := ur.db.Query(`
    INSERT INTO refresh_token (user_id)
    VALUES ($1)
    RETURNING id;
  `, data.UserId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var token dtos.CreatedTokenDTO
	rows.Next()
	err = rows.Scan(&token.Id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &token, nil
}

func (ur *PostgresAuthRepo) CreateAccessToken(data *dtos.CreateAccessTokenDTO) (*dtos.CreatedTokenDTO, error) {
	rows, err := ur.db.Query(`
    INSERT INTO access_token (user_id, refresh_token_id, expires_at)
    VALUES ($1, $2, $3)
    RETURNING id;
  `, data.UserId, data.RefreshTokenId, data.ExpiresAt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	var token dtos.CreatedTokenDTO
	rows.Next()
	err = rows.Scan(&token.Id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &token, nil
}

func (ur *PostgresAuthRepo) CheckIfActiveRefreshTokenExist(refreshTokenId string) (bool, error) {
	rows, err := ur.db.Query(`
    SELECT COUNT(*)
    FROM refresh_token
    WHERE id=$1;
  `, refreshTokenId)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	defer rows.Close()

	var count int
	rows.Next()
	err = rows.Scan(&count)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return count > 0, nil
}

func (ur *PostgresAuthRepo) GetUserByEmail(email string) (*dtos.UserDTO, error) {
	rows, err := ur.db.Query(`
    SELECT id, email, password
    FROM "user"
    WHERE email = $1;
`, email)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var user dtos.UserDTO
		err = rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return &user, nil
	}

	return nil, nil
}

func (ur *PostgresAuthRepo) DeactivateUserTokensByUserId(userId string) error {
	_, err := ur.db.Exec(`
    UPDATE refresh_token
    SET is_active=false
    WHERE user_id=$1 AND is_active=true
`, userId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = ur.db.Exec(`
    UPDATE access_token
    SET is_active=false
    WHERE user_id=$1 AND is_active=true
`, userId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (ur *PostgresAuthRepo) DeactivateAccessTokenByRefreshTokenId(refreshTokenId string) error {
	_, err := ur.db.Exec(`
    UPDATE access_token
    SET is_active=false
    WHERE refresh_token_id=$1 AND is_active=true
`, refreshTokenId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
