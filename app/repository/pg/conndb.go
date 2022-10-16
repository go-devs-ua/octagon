package pg

import (
	"database/sql"
	"fmt"

	"github.com/go-devs-ua/octagon/cfg"
)

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type DBOptions struct {
	DB DB
}

func ConnectDB(cfg cfg.Options) (*sql.DB, error) {
	str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database failed: %w", err)
	}

	return db, nil
}
