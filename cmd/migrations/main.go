package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/go-devs-ua/octagon/app/repository/pg"
	"github.com/go-devs-ua/octagon/lgr"
	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	logger := lgr.New()

	opt, err := pg.GetConfig()
	if err != nil {
		logger.Errorf("Failed to get config from .env: %+v\n", err)
		return
	}

	db, err := pg.ConnectDB(opt.DB)
	if err != nil {
		logger.Errorf("%+v\n", err)
		return
	}

	direction := flag.String("migrate", "", "applying migrations 'up/down'")
	flag.Parse()

	var dir migrate.MigrationDirection
	if *direction == "down" {
		dir = 1
	}

	if err = migrateDB(db, logger, dir); err != nil {
		logger.Errorf("Failed making migrations: %v\n", err)
	}
}

// MigrateDB executes migrations.
func migrateDB(db *sql.DB, logger *lgr.Logger, direction migrate.MigrationDirection) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	var dirLog string
	if direction == 0 {
		dirLog = "up"
	} else {
		dirLog = "down"
	}

	logger.Infof("Starting applying migrations '%s'...", dirLog)

	n, err := migrate.Exec(db, "postgres", migrations, direction)
	if err != nil {
		return fmt.Errorf("migration up failed: %w", err)
	}

	logger.Infof("The number of applied migration is: %d\n", n)

	return nil
}
