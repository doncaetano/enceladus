package repo

import (
	"database/sql"
	"log"

	"github.com/rhuancaetano/enceladus/internal/repo"
	"github.com/rhuancaetano/enceladus/internal/user/dtos"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo() *PostgresUserRepo {
	return &PostgresUserRepo{
		db: repo.GetPostgresDatabase(),
	}
}

func (ur *PostgresUserRepo) CreateUser(data *dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	rows, err := ur.db.Query(`
    INSERT INTO "user" (first_name, last_name, email)
    VALUES ($1, $2, $3)
    RETURNING id, first_name, last_name, email, is_active, created_at, updated_at;
  `, data.FirstName, data.LastName, data.Email)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var user dtos.UserDTO
	rows.Next()
	err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}
