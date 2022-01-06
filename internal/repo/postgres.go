package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/rhuancaetano/enceladus/internal/config"
)

var db *sql.DB

func InitDatabase() {
	fmt.Println("Connecting to Postgres...")
	db := GetPostgresDatabase()
	err := db.Ping()
	if err != nil {
		log.Println("Could not connect to Postgres")
		log.Fatal(err.Error())
	}
}

func GetPostgresDatabase() *sql.DB {
	var err error

	env := config.GetEnvironment()

	if db == nil {
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_HOST, env.POSTGRES_DB))
		if err != nil {
			log.Println("Could not connect to Postgres")
			log.Fatal(err.Error())
		}
	}

	return db
}
