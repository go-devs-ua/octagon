package pg

import (
	"database/sql"
	"fmt"

	"github.com/go-devs-ua/octagon/cfg"
)

// ConnectDB is used to create connection to postgres DB.
func ConnectDB(cfg cfg.DB) (*sql.DB, error) {
	str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database failed: %w", err)
	}

	return db, nil
}
