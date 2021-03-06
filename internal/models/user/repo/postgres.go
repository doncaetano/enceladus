package repo

import (
	"database/sql"
	"log"

	"github.com/rhuancaetano/enceladus/internal/models/user/dtos"
	"github.com/rhuancaetano/enceladus/internal/repo"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo() *PostgresUserRepo {
	return &PostgresUserRepo{
		db: repo.GetPostgresDatabase(),
	}
}

func (ur *PostgresUserRepo) FindByEmail(email string) (*dtos.UserDTO, error) {
	rows, err := ur.db.Query(`
  SELECT id, first_name, last_name, email, is_active, created_at, updated_at
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
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return &user, nil
	}

	return nil, nil
}

func (ur *PostgresUserRepo) FindById(id string) (*dtos.UserDTO, error) {
	rows, err := ur.db.Query(`
  SELECT id, first_name, last_name, email, is_active, created_at, updated_at
  FROM "user"
  WHERE id = $1;
`, id)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var user dtos.UserDTO
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return &user, nil
	}

	return nil, nil
}

func (ur *PostgresUserRepo) DeleteUser(id string) error {
	_, err := ur.db.Query(`DELETE FROM "user" WHERE id = $1;`, id)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (ur *PostgresUserRepo) CreateUser(data *dtos.CreateUserDTO) (*dtos.UserDTO, error) {
	rows, err := ur.db.Query(`
    INSERT INTO "user" (first_name, last_name, email, password)
    VALUES ($1, $2, $3, $4)
    RETURNING id, first_name, last_name, email, is_active, created_at, updated_at;
  `, data.FirstName, data.LastName, data.Email, data.Password)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var user dtos.UserDTO
	rows.Next()
	err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *PostgresUserRepo) UpdateUser(data *dtos.UpdateUserDTO) (*dtos.UserDTO, error) {
	rows, err := ur.db.Query(`
    UPDATE "user"
    SET first_name=$2,
        last_name=$3,
        email=$4,
        password=$5,
        is_active=$6
    WHERE id=$1
    RETURNING id, first_name, last_name, email, is_active, created_at, updated_at;
  `, data.Id, data.FirstName, data.LastName, data.Email, data.Password, data.IsActive)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var user dtos.UserDTO
	rows.Next()
	err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}
