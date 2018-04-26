package context

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenDB(config *Config) *sqlx.DB {
	db, err := sqlx.Connect(config.DB.Engine, fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.Name, config.DB.User, config.DB.Password))
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	return db
}
