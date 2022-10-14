package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/lgr"
	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

const (
	up   = "up"
	down = "down"
)

func main() {
	logger := lgr.New()
	// defer logger.Flush()

	opt, err := cfg.GetConfig()
	if err != nil {
		logger.Errorf("Failed to get config from .env: %+v\n", err)
		return
	}

	db, err := connectDB(opt)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return
	}

	if err := migrateDB(db, logger); err != nil {
		logger.Errorf("Failed making migrations: %+v\n", err)
		return
	}

}

// MigrateDB executes migrations.
func migrateDB(db *sql.DB, logger *lgr.Logger) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	for _, arg := range os.Args {
		if arg == up {
			logger.Infof("Apllying migrations 'up'.")
			n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
			if err != nil {
				return fmt.Errorf("migration up failed: %w", err)
			}
			logger.Infof("The number of applied migration is: %d\n", n)
		}
		if arg == down {
			logger.Infof("Apllying migrations 'down'.")
			n, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
			if err != nil {
				return fmt.Errorf("migration up failed: %w", err)
			}
			logger.Infof("The number of applied migration is: %d\n", n)
		}
	}

	return nil
}

func connectDB(opt cfg.Options) (*sql.DB, error) {
	str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		opt.DB.Host, opt.DB.Port, opt.DB.Username, opt.DB.Password, opt.DB.DBName)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database failed: %w", err)
	}

	return db, nil
}
