package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	"github.com/go-devs-ua/octagon/cfg"
	"github.com/go-devs-ua/octagon/lgr"
	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

const (
	up           = "up"
	down         = "down"
	migrationDir = "./migration"
)

func main() {
	logger := lgr.New()

	config, err := cfg.GetConfig()
	if err != nil {
		logger.Errorf("Failed to get config from .env: %+v", err)
		return
	}

	db, err := pg.ConnectDB(config.DB)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return
	}

	direction := flag.String("migrate", "", "applying migrations 'up/down'")
	flag.Parse()

	if *direction != up && *direction != down {
		fmt.Println("Wrong flag provided, choose '-migrate up' or '-migrate down'")
		return
	}

	if err := migrateDB(db, logger, *direction); err != nil {
		logger.Errorf("Failed making migrations: %v", err)
	}

}

// MigrateDB executes migrations.
func migrateDB(db *sql.DB, logger *lgr.Logger, direction string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationDir,
	}

	var dir migrate.MigrationDirection
	if direction == down {
		dir = 1
	}

	logger.Infof("Starting applying migrations '%s'...", direction)

	n, err := migrate.Exec(db, "postgres", migrations, dir)
	if err != nil {
		return fmt.Errorf("migration up failed: %w", err)
	}

	logger.Infof("The number of applied migration is: %d", n)

	return nil
}
