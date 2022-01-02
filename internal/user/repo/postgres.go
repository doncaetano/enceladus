package repo

import (
	"database/sql"
	"fmt"

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

func (ur *PostgresUserRepo) CreateUser() {
	fmt.Println("CREATE USER")
}
