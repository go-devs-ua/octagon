package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Execute migrations.
// Use "migrate-up" argument after run of program. It will create new table
// Or use "migrate-down" - that will remove table
// Example: go run main.go migrate-up
func Migrate(db *sql.DB) error {
	for _, arg := range os.Args {
		if arg == "migrate-up" {
			log.Println("Migration starts up")

			if err := execMigrationQuery(db, "./migration/migrationsUp.sql"); err != nil {
				return fmt.Errorf("migration up failed: %w", err)
			}
		}

		if arg == "migrate-down" {
			log.Println("Migration starts down")

			if err := execMigrationQuery(db, "./migration/migrationsDown.sql"); err != nil {
				return fmt.Errorf("migration down failed: %w", err)
			}
		}
	}

	return nil
}

func execMigrationQuery(db *sql.DB, fileName string) error {
	byteQuery, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("read migration file: %w", err)
	}

	if _, err := db.Exec(string(byteQuery)); err != nil {
		return fmt.Errorf("cannot exec query: %w", err)
	}

	return nil
}
