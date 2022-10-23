package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

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
	config, err := cfg.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %+v", err)

		return
	}

	logger, err := lgr.New(config.LogLevel)
	if err != nil {
		log.Printf("failed to create logger: %v", err)

		return
	}

	db, err := pg.ConnectDB(config.DB)
	if err != nil {
		logger.Errorf("%+v", err)

		return
	}

	direction := flag.String("migrate", "", "applying migration direction")
	flag.Parse()

	if *direction != up && *direction != down {
		log.Printf("Wrong flag provided, choose '-migrate %s' or '-migrate %s'\n", up, down)

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
